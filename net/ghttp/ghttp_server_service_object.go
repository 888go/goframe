// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
)

// BindObject 将对象注册到服务器路由上，给定特定的模式。
//
// 可选参数 `method` 用于指定要注册的方法，该方法支持多个方法名；
// 多个方法之间用字符 ',' 分隔，大小写敏感。
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

// BindObjectMethod 将指定对象的方法注册到服务器路由中，使用给定的模式。
//
// 可选参数 `method` 用于指定要注册的方法，该参数不支持多个方法名，仅支持单个、大小写敏感的方法名。
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

// BindObjectRest 以指定模式将符合REST API风格的对象注册到服务器。
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
	// 将输入方法转换为映射以便于实现高效检索
	var methodMap map[string]bool
	if len(in.Method) > 0 {
		methodMap = make(map[string]bool)
		for _, v := range strings.Split(in.Method, ",") {
			methodMap[strings.TrimSpace(v)] = true
		}
	}
// 如果`pattern`中的`method`是`defaultMethod`，
// 为了方便后续语句的控制，将其移除。
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
// 如果给定的`object`不是指针，它会创建一个临时指针，
// 其指向值为`reflectValue`。
// 然后可以获取结构体（包括结构体指针）的所有方法。
// 这段代码注释的翻译如下：
// ```go
// 如果传入的`object`不是一个指针类型，
// 则会创建一个临时指针变量，该指针指向`reflectValue`。
// 这样就可以获取到结构体及其指针类型的全部方法。
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
// 如果存在"Index"方法，则会自动添加一个附加路由以匹配主URI，例如：
// 如果模式是"/user"，那么"/user"和"/user/index"都会被自动注册。
//
// 注意，如果模式中存在内置变量，则此路由不会被自动添加。
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
// 如果给定的`object`不是指针，则创建一个临时指针，
// 其值为`v`。
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
// 如果给定的`object`不是指针，则创建一个临时指针，
// 其值为`v`。
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
