// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp

import (
	"bytes"
	"context"
	"reflect"
	"strings"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/os/gstructs"
	"github.com/888go/goframe/text/gstr"
)

// BindHandler 将一个处理函数注册到服务器，该函数与给定的模式关联。
//
// 注意参数 `handler` 可以是以下类型：
// 1. func(*ghttp.Request) // 类型为接收*ghttp.Request参数的函数
// 2. func(context.Context, BizRequest)(BizResponse, error) // 类型为接收context.Context和BizRequest参数，并返回BizResponse和error的函数
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

// doBindHandler 函数用于将指定模式的处理器函数注册到服务器。
//
// 参数 `pattern` 形如：
// /user/list, put:/user, delete:/user, post:/user@goframe.org
// 其中，这些模式用于定义HTTP请求的方法（如GET、PUT、DELETE等）以及对应的路由路径。
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

// bindHandlerByMap 使用map将处理器注册到服务器。
func (s *Server) bindHandlerByMap(ctx context.Context, prefix string, m map[string]*HandlerItem) {
	for pattern, handler := range m {
		s.setHandler(ctx, setHandlerInput{
			Prefix:      prefix,
			Pattern:     pattern,
			HandlerItem: handler,
		})
	}
}

// mergeBuildInNameToPattern 将内建名称按照以下规则合并到模式中，这些内建名称的命名格式为 "{.xxx}"。
// 规则1：若模式中的URI包含 {.struct} 关键字，则用结构体名称替换该关键字；
// 规则2：若模式中的URI包含 {.method} 关键字，则用方法名称替换该关键字；
// 规则3：如果未满足规则1，则将方法名称直接追加到模式中URI的末尾；
//
// 参数 `allowAppend` 指定是否允许将方法名称追加到模式末尾。
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
	// 检查域名参数。
	var (
		array = strings.Split(pattern, "@")
		uri   = strings.TrimRight(array[0], "/") + "/" + methodName
	)
	// 将domain参数追加到URI。
	if len(array) > 1 {
		return uri + "@" + array[1]
	}
	return uri
}

// nameToUri 将给定名称转换为URL格式，遵循以下规则：
// 规则0：将所有方法名转为小写，并在单词间添加字符'-'。
// 规则1：不转换方法名，使用原始方法名构建URI。
// 规则2：将所有方法名转为小写，单词间无连接符号。
// 规则3：使用驼峰命名法。
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

// 不要启用这段逻辑，因为许多用户已经使用非结构体指针类型作为第一个输出参数。
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

	// 请求结构体应命名为 `xxxReq`。
	reqStructName := trimGeneric(reflectType.In(1).String())
	if !gstr.HasSuffix(reqStructName, `Req`) {
		err = gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid struct naming for request: defined as "%s", but it should be named with "Req" suffix like "XxxReq"`,
			reqStructName,
		)
		return
	}

	// 响应结构体应当命名为 `xxxRes`。
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

	// 它检索并返回请求结构体的字段。
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
		// 使用动态创建的参数值调用处理器。
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

// trimGeneric 从响应类型名称中移除泛型的类型定义字符串（如果存在的话）
func trimGeneric(structName string) string {
	var (
		leftBraceIndex  = strings.LastIndex(structName, "[") // 对于泛型，从结尾开始遍历比从开头开始更快
		rightBraceIndex = strings.LastIndex(structName, "]")
	)
	if leftBraceIndex == -1 || rightBraceIndex == -1 {
		// 未找到 '[' 或 ']'
		return structName
	} else if leftBraceIndex+1 == rightBraceIndex {
// 可能是一个切片，因为泛型是 '[X]'，而不是 '[]'
// 为了兼容不规范的返回参数类型：[]XxxRes
		return structName
	}
	return structName[:leftBraceIndex]
}
