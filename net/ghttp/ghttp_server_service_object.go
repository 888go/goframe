// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	gfile "github.com/888go/goframe/os/gfile"
	gregex "github.com/888go/goframe/text/gregex"
	gstr "github.com/888go/goframe/text/gstr"
)

// BindObject 将对象绑定到具有给定模式的服务器路由。
//
// 可选参数 `method` 用于指定要注册的方法，支持多个方法名称；多个方法名称之间用字符 `,` 分隔，区分大小写。
// md5:224eaf0adfd81c84
func (s *Server) BindObject(pattern string, object interface{}, method ...string) {
	var bindMethod = ""
	if len(method) > 0 {
		bindMethod = method[0]
	}
	s.doBindObject(context.TODO(), doBindObjectInput{
		Prefix:     "",
		Pattern:    pattern,
		Object:     object,
		Method:     bindMethod,
		Middleware: nil,
		Source:     "",
	})
}

// BindObjectMethod 将指定对象的特定方法与给定模式的服务器路由绑定。
// 
// 可选参数 `method` 用于指定要注册的方法，它不支持多个方法名，仅支持一个，且区分大小写。
// md5:badb3f7323abfd11
func (s *Server) BindObjectMethod(pattern string, object interface{}, method string) {
	s.doBindObjectMethod(context.TODO(), doBindObjectMethodInput{
		Prefix:     "",
		Pattern:    pattern,
		Object:     object,
		Method:     method,
		Middleware: nil,
		Source:     "",
	})
}

// BindObjectRest 使用指定的模式将对象以REST API风格注册到服务器。 md5:e071850c88eb6751
func (s *Server) BindObjectRest(pattern string, object interface{}) {
	s.doBindObjectRest(context.TODO(), doBindObjectInput{
		Prefix:     "",
		Pattern:    pattern,
		Object:     object,
		Method:     "",
		Middleware: nil,
		Source:     "",
	})
}

type doBindObjectInput struct {
	Prefix     string
	Pattern    string
	Object     interface{}
	Method     string
	Middleware []HandlerFunc
	Source     string
}

func (s *Server) doBindObject(ctx context.Context, in doBindObjectInput) {
		// 将输入方法转换为映射，以便于进行高效便捷的搜索。 md5:116ad79ef3003f65
	var methodMap map[string]bool
	if len(in.Method) > 0 {
		methodMap = make(map[string]bool)
		for _, v := range strings.Split(in.Method, ",") {
			methodMap[strings.TrimSpace(v)] = true
		}
	}
	// 如果`pattern`中的`method`为`defaultMethod`，为了方便后续语句的控制，它会移除。
	// md5:08bf69a00eee9caa
	domain, method, path, err := s.parsePattern(in.Pattern)
	if err != nil {
		s.Logger().Fatalf(ctx, `%+v`, err)
		return
	}
	if gstr.Equal(method, defaultMethod) {
		in.Pattern = s.serveHandlerKey("", path, domain)
	}
	var (
		handlerMap   = make(map[string]*HandlerItem)
		reflectValue = reflect.ValueOf(in.Object)
		reflectType  = reflectValue.Type()
		initFunc     func(*Request)
		shutFunc     func(*Request)
	)
	// 如果给定的`object`不是指针，那么它会创建一个临时的，其值为`reflectValue`。
	// 然后它可以获取结构体/`*struct`的所有方法。
	// md5:1e216cd9c7839ef2
	if reflectValue.Kind() == reflect.Struct {
		newValue := reflect.New(reflectType)
		newValue.Elem().Set(reflectValue)
		reflectValue = newValue
		reflectType = reflectValue.Type()
	}
	structName := reflectType.Elem().Name()
	if reflectValue.MethodByName(specialMethodNameInit).IsValid() {
		initFunc = reflectValue.MethodByName(specialMethodNameInit).Interface().(func(*Request))
	}
	if reflectValue.MethodByName(specialMethodNameShut).IsValid() {
		shutFunc = reflectValue.MethodByName(specialMethodNameShut).Interface().(func(*Request))
	}
	pkgPath := reflectType.Elem().PkgPath()
	pkgName := gfile.Basename(pkgPath)
	for i := 0; i < reflectValue.NumMethod(); i++ {
		methodName := reflectType.Method(i).Name
		if methodMap != nil && !methodMap[methodName] {
			continue
		}
		if methodName == specialMethodNameInit || methodName == specialMethodNameShut {
			continue
		}
		objName := gstr.Replace(reflectType.String(), fmt.Sprintf(`%s.`, pkgName), "")
		if objName[0] == '*' {
			objName = fmt.Sprintf(`(%s)`, objName)
		}

		funcInfo, err := s.checkAndCreateFuncInfo(reflectValue.Method(i).Interface(), pkgPath, objName, methodName)
		if err != nil {
			s.Logger().Fatalf(ctx, `%+v`, err)
		}

		key := s.mergeBuildInNameToPattern(in.Pattern, structName, methodName, true)
		handlerMap[key] = &HandlerItem{
			Name:       fmt.Sprintf(`%s.%s.%s`, pkgPath, objName, methodName),
			Type:       HandlerTypeObject,
			Info:       funcInfo,
			InitFunc:   initFunc,
			ShutFunc:   shutFunc,
			Middleware: in.Middleware,
			Source:     in.Source,
		}
		// 如果存在"Index"方法，则会自动添加一个额外的路由来匹配主URI，例如：
		// 如果模式是"/user"，那么"/user"和"/user/index"都会被自动
		// 注册。
		//
		// 请注意，如果模式中包含内置变量，这条路由将不会被自动添加。
		// md5:96b4d9eca149582c
		var (
			isIndexMethod = strings.EqualFold(methodName, specialMethodNameIndex)
			hasBuildInVar = gregex.IsMatchString(`\{\.\w+\}`, in.Pattern)
			hashTwoParams = funcInfo.Type.NumIn() == 2
		)
		if isIndexMethod && !hasBuildInVar && !hashTwoParams {
			p := gstr.PosRI(key, "/index")
			k := key[0:p] + key[p+6:]
			if len(k) == 0 || k[0] == '@' {
				k = "/" + k
			}
			handlerMap[k] = &HandlerItem{
				Name:       fmt.Sprintf(`%s.%s.%s`, pkgPath, objName, methodName),
				Type:       HandlerTypeObject,
				Info:       funcInfo,
				InitFunc:   initFunc,
				ShutFunc:   shutFunc,
				Middleware: in.Middleware,
				Source:     in.Source,
			}
		}
	}
	s.bindHandlerByMap(ctx, in.Prefix, handlerMap)
}

type doBindObjectMethodInput struct {
	Prefix     string
	Pattern    string
	Object     interface{}
	Method     string
	Middleware []HandlerFunc
	Source     string
}

func (s *Server) doBindObjectMethod(ctx context.Context, in doBindObjectMethodInput) {
	var (
		handlerMap   = make(map[string]*HandlerItem)
		reflectValue = reflect.ValueOf(in.Object)
		reflectType  = reflectValue.Type()
		initFunc     func(*Request)
		shutFunc     func(*Request)
	)
	// 如果给定的`object`不是指针，那么它会创建一个临时的指针，
	// 其值为`v`。
	// md5:ea1cbad8bfbac476
	if reflectValue.Kind() == reflect.Struct {
		newValue := reflect.New(reflectType)
		newValue.Elem().Set(reflectValue)
		reflectValue = newValue
		reflectType = reflectValue.Type()
	}
	var (
		structName  = reflectType.Elem().Name()
		methodName  = strings.TrimSpace(in.Method)
		methodValue = reflectValue.MethodByName(methodName)
	)
	if !methodValue.IsValid() {
		s.Logger().Fatalf(ctx, "invalid method name: %s", methodName)
		return
	}
	if reflectValue.MethodByName(specialMethodNameInit).IsValid() {
		initFunc = reflectValue.MethodByName(specialMethodNameInit).Interface().(func(*Request))
	}
	if reflectValue.MethodByName(specialMethodNameShut).IsValid() {
		shutFunc = reflectValue.MethodByName(specialMethodNameShut).Interface().(func(*Request))
	}
	var (
		pkgPath = reflectType.Elem().PkgPath()
		pkgName = gfile.Basename(pkgPath)
		objName = gstr.Replace(reflectType.String(), fmt.Sprintf(`%s.`, pkgName), "")
	)
	if objName[0] == '*' {
		objName = fmt.Sprintf(`(%s)`, objName)
	}

	funcInfo, err := s.checkAndCreateFuncInfo(methodValue.Interface(), pkgPath, objName, methodName)
	if err != nil {
		s.Logger().Fatalf(ctx, `%+v`, err)
	}

	key := s.mergeBuildInNameToPattern(in.Pattern, structName, methodName, false)
	handlerMap[key] = &HandlerItem{
		Name:       fmt.Sprintf(`%s.%s.%s`, pkgPath, objName, methodName),
		Type:       HandlerTypeObject,
		Info:       funcInfo,
		InitFunc:   initFunc,
		ShutFunc:   shutFunc,
		Middleware: in.Middleware,
		Source:     in.Source,
	}

	s.bindHandlerByMap(ctx, in.Prefix, handlerMap)
}

func (s *Server) doBindObjectRest(ctx context.Context, in doBindObjectInput) {
	var (
		handlerMap   = make(map[string]*HandlerItem)
		reflectValue = reflect.ValueOf(in.Object)
		reflectType  = reflectValue.Type()
		initFunc     func(*Request)
		shutFunc     func(*Request)
	)
	// 如果给定的`object`不是指针，那么它会创建一个临时的指针，
	// 其值为`v`。
	// md5:ea1cbad8bfbac476
	if reflectValue.Kind() == reflect.Struct {
		newValue := reflect.New(reflectType)
		newValue.Elem().Set(reflectValue)
		reflectValue = newValue
		reflectType = reflectValue.Type()
	}
	structName := reflectType.Elem().Name()
	if reflectValue.MethodByName(specialMethodNameInit).IsValid() {
		initFunc = reflectValue.MethodByName(specialMethodNameInit).Interface().(func(*Request))
	}
	if reflectValue.MethodByName(specialMethodNameShut).IsValid() {
		shutFunc = reflectValue.MethodByName(specialMethodNameShut).Interface().(func(*Request))
	}
	pkgPath := reflectType.Elem().PkgPath()
	for i := 0; i < reflectValue.NumMethod(); i++ {
		methodName := reflectType.Method(i).Name
		if _, ok := methodsMap[strings.ToUpper(methodName)]; !ok {
			continue
		}
		pkgName := gfile.Basename(pkgPath)
		objName := gstr.Replace(reflectType.String(), fmt.Sprintf(`%s.`, pkgName), "")
		if objName[0] == '*' {
			objName = fmt.Sprintf(`(%s)`, objName)
		}

		funcInfo, err := s.checkAndCreateFuncInfo(
			reflectValue.Method(i).Interface(),
			pkgPath,
			objName,
			methodName,
		)
		if err != nil {
			s.Logger().Fatalf(ctx, `%+v`, err)
		}

		key := s.mergeBuildInNameToPattern(methodName+":"+in.Pattern, structName, methodName, false)
		handlerMap[key] = &HandlerItem{
			Name:       fmt.Sprintf(`%s.%s.%s`, pkgPath, objName, methodName),
			Type:       HandlerTypeObject,
			Info:       funcInfo,
			InitFunc:   initFunc,
			ShutFunc:   shutFunc,
			Middleware: in.Middleware,
			Source:     in.Source,
		}
	}
	s.bindHandlerByMap(ctx, in.Prefix, handlerMap)
}
