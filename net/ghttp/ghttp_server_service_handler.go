// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package ghttp

import (
	"bytes"
	"context"
	"reflect"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gstructs"
	"github.com/gogf/gf/v2/text/gstr"
)

// BindHandler 将一个处理器函数注册到服务器，使用给定的模式。
//
// 注意，参数 `handler` 可以是以下两种类型之一：
// 1. func(*ghttp.Request)
// 2. func(context.Context, BizRequest) (BizResponse, error) md5:245b5139c4d933ad
func (s *Server) BindHandler(pattern string, handler interface{}) {
	var ctx = context.TODO()
	funcInfo, err := s.checkAndCreateFuncInfo(handler, "", "", "")
	if err != nil {
		s.Logger().Fatalf(ctx, `%+v`, err)
	}
	s.doBindHandler(ctx, doBindHandlerInput{
		Prefix:     "",
		Pattern:    pattern,
		FuncInfo:   funcInfo,
		Middleware: nil,
		Source:     "",
	})
}

type doBindHandlerInput struct {
	Prefix     string
	Pattern    string
	FuncInfo   handlerFuncInfo
	Middleware []HandlerFunc
	Source     string
}

// doBindHandler 使用给定的模式向服务器注册一个处理函数。
//
// 参数 `pattern` 的格式如下：
// /user/list, put:/user, delete:/user, post:/user@goframe.org md5:d71f121a1c2830d3
func (s *Server) doBindHandler(ctx context.Context, in doBindHandlerInput) {
	s.setHandler(ctx, setHandlerInput{
		Prefix:  in.Prefix,
		Pattern: in.Pattern,
		HandlerItem: &HandlerItem{
			Type:       HandlerTypeHandler,
			Info:       in.FuncInfo,
			Middleware: in.Middleware,
			Source:     in.Source,
		},
	})
}

// bindHandlerByMap 使用映射注册处理器到服务器。 md5:15729f837b1bc875
func (s *Server) bindHandlerByMap(ctx context.Context, prefix string, m map[string]*HandlerItem) {
	for pattern, handler := range m {
		s.setHandler(ctx, setHandlerInput{
			Prefix:      prefix,
			Pattern:     pattern,
			HandlerItem: handler,
		})
	}
}

// mergeBuildInNameToPattern 将内置名称合并到模式中，根据以下规则进行操作，内置名称的命名方式为"{.xxx}"。
// 规则 1：如果模式中的URI包含{.struct}关键字，它将替换该关键字为结构体名称；
// 规则 2：如果模式中的URI包含{.method}关键字，它将替换该关键字为方法名称；
// 规则 3：如果没有满足规则 1，那么直接在模式中的URI后添加方法名称。
//
// 参数 `allowAppend` 指定是否允许将方法名称追加到模式的末尾。 md5:1c79af7afc57b081
func (s *Server) mergeBuildInNameToPattern(pattern string, structName, methodName string, allowAppend bool) string {
	structName = s.nameToUri(structName)
	methodName = s.nameToUri(methodName)
	pattern = strings.ReplaceAll(pattern, "{.struct}", structName)
	if strings.Contains(pattern, "{.method}") {
		return strings.ReplaceAll(pattern, "{.method}", methodName)
	}
	if !allowAppend {
		return pattern
	}
	// 检查域名参数。 md5:1a963c36e4fee004
	var (
		array = strings.Split(pattern, "@")
		uri   = strings.TrimRight(array[0], "/") + "/" + methodName
	)
	// 将域名参数追加到URI。 md5:f94214453c1409c8
	if len(array) > 1 {
		return uri + "@" + array[1]
	}
	return uri
}

// nameToUri 使用以下规则将给定的名称转换为URL格式：
// 规则0：将所有方法名转换为小写，单词间添加字符'-'。
// 规则1：不转换方法名，使用原始方法名构建URI。
// 规则2：将所有方法名转换为小写，单词间不添加连接符号。
// 规则3：使用驼峰式命名。 md5:c9f350c3c6635668
func (s *Server) nameToUri(name string) string {
	switch s.config.NameToUriType {
	case UriTypeFullName:
		return name

	case UriTypeAllLower:
		return strings.ToLower(name)

	case UriTypeCamel:
		part := bytes.NewBuffer(nil)
		if gstr.IsLetterUpper(name[0]) {
			part.WriteByte(name[0] + 32)
		} else {
			part.WriteByte(name[0])
		}
		part.WriteString(name[1:])
		return part.String()

	case UriTypeDefault:
		fallthrough

	default:
		part := bytes.NewBuffer(nil)
		for i := 0; i < len(name); i++ {
			if i > 0 && gstr.IsLetterUpper(name[i]) {
				part.WriteByte('-')
			}
			if gstr.IsLetterUpper(name[i]) {
				part.WriteByte(name[i] + 32)
			} else {
				part.WriteByte(name[i])
			}
		}
		return part.String()
	}
}

func (s *Server) checkAndCreateFuncInfo(
	f interface{}, pkgPath, structName, methodName string,
) (funcInfo handlerFuncInfo, err error) {
	funcInfo = handlerFuncInfo{
		Type:  reflect.TypeOf(f),
		Value: reflect.ValueOf(f),
	}
	if handlerFunc, ok := f.(HandlerFunc); ok {
		funcInfo.Func = handlerFunc
		return
	}

	var (
		reflectType    = funcInfo.Type
		inputObject    reflect.Value
		inputObjectPtr interface{}
	)
	if reflectType.NumIn() != 2 || reflectType.NumOut() != 2 {
		if pkgPath != "" {
			err = gerror.NewCodef(
				gcode.CodeInvalidParameter,
				`invalid handler: %s.%s.%s defined as "%s", but "func(*ghttp.Request)" or "func(context.Context, *BizReq)(*BizRes, error)" is required`,
				pkgPath, structName, methodName, reflectType.String(),
			)
		} else {
			err = gerror.NewCodef(
				gcode.CodeInvalidParameter,
				`invalid handler: defined as "%s", but "func(*ghttp.Request)" or "func(context.Context, *BizReq)(*BizRes, error)" is required`,
				reflectType.String(),
			)
		}
		return
	}

	if !reflectType.In(0).Implements(reflect.TypeOf((*context.Context)(nil)).Elem()) {
		err = gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid handler: defined as "%s", but the first input parameter should be type of "context.Context"`,
			reflectType.String(),
		)
		return
	}

	if !reflectType.Out(1).Implements(reflect.TypeOf((*error)(nil)).Elem()) {
		err = gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid handler: defined as "%s", but the last output parameter should be type of "error"`,
			reflectType.String(),
		)
		return
	}

	if reflectType.In(1).Kind() != reflect.Ptr ||
		(reflectType.In(1).Kind() == reflect.Ptr && reflectType.In(1).Elem().Kind() != reflect.Struct) {
		err = gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid handler: defined as "%s", but the second input parameter should be type of pointer to struct like "*BizReq"`,
			reflectType.String(),
		)
		return
	}

	// 不要启用此逻辑，因为许多用户已经将非结构指针类型作为第一个输出参数使用。 md5:46785e26d27207d1
	/*
		if reflectType.Out(0).Kind() != reflect.Ptr ||
			(reflectType.Out(0).Kind() == reflect.Ptr && reflectType.Out(0).Elem().Kind() != reflect.Struct) {
			err = gerror.NewCodef(
				gcode.CodeInvalidParameter,
				`invalid handler: defined as "%s", but the first output parameter should be type of pointer to struct like "*BizRes"`,
				reflectType.String(),
			)
			return
		}
	*/

	// 请求结构体应该命名为 `xxxReq`。 md5:f366399bf3de35a1
	reqStructName := trimGeneric(reflectType.In(1).String())
	if !gstr.HasSuffix(reqStructName, `Req`) {
		err = gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid struct naming for request: defined as "%s", but it should be named with "Req" suffix like "XxxReq"`,
			reqStructName,
		)
		return
	}

	// 响应结构体应当命名为 `xxxRes`。 md5:0e837067ff972f27
	resStructName := trimGeneric(reflectType.Out(0).String())
	if !gstr.HasSuffix(resStructName, `Res`) {
		err = gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid struct naming for response: defined as "%s", but it should be named with "Res" suffix like "XxxRes"`,
			resStructName,
		)
		return
	}

	funcInfo.IsStrictRoute = true

	inputObject = reflect.New(funcInfo.Type.In(1).Elem())
	inputObjectPtr = inputObject.Interface()

	// 该函数获取并返回请求结构体的字段。 md5:25b3db67b1969d01
	fields, err := gstructs.Fields(gstructs.FieldsInput{
		Pointer:         inputObjectPtr,
		RecursiveOption: gstructs.RecursiveOptionEmbedded,
	})
	if err != nil {
		return funcInfo, err
	}
	funcInfo.ReqStructFields = fields
	funcInfo.Func = createRouterFunc(funcInfo)
	return
}

func createRouterFunc(funcInfo handlerFuncInfo) func(r *Request) {
	return func(r *Request) {
		var (
			ok          bool
			err         error
			inputValues = []reflect.Value{
				reflect.ValueOf(r.Context()),
			}
		)
		if funcInfo.Type.NumIn() == 2 {
			var inputObject reflect.Value
			if funcInfo.Type.In(1).Kind() == reflect.Ptr {
				inputObject = reflect.New(funcInfo.Type.In(1).Elem())
				r.error = r.Parse(inputObject.Interface())
			} else {
				inputObject = reflect.New(funcInfo.Type.In(1).Elem()).Elem()
				r.error = r.Parse(inputObject.Addr().Interface())
			}
			if r.error != nil {
				return
			}
			inputValues = append(inputValues, inputObject)
		}
		// 使用动态创建的参数值调用处理器。 md5:991efec71cdcc95a
		results := funcInfo.Value.Call(inputValues)
		switch len(results) {
		case 1:
			if !results[0].IsNil() {
				if err, ok = results[0].Interface().(error); ok {
					r.error = err
				}
			}

		case 2:
			r.handlerResponse = results[0].Interface()
			if !results[1].IsNil() {
				if err, ok = results[1].Interface().(error); ok {
					r.error = err
				}
			}
		}
	}
}

// trimGeneric 如果响应类型名称包含泛型定义，删除字符串类型的定义. md5:3c6ea03dfa650b71
func trimGeneric(structName string) string {
	var (
		leftBraceIndex  = strings.LastIndex(structName, "[") // 对于泛型来说，从末尾开始比从开头开始更快. md5:9e8730bfe1647d52
		rightBraceIndex = strings.LastIndex(structName, "]")
	)
	if leftBraceIndex == -1 || rightBraceIndex == -1 {
		// not found '[' or ']'
		return structName
	} else if leftBraceIndex+1 == rightBraceIndex {
		// 可能是一个切片，因为泛型是'[X]'而不是'[]'
		// 以兼容不良的返回参数类型：[]XxxRes md5:a521893d3e187a1a
		return structName
	}
	return structName[:leftBraceIndex]
}
