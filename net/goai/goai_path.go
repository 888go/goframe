// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package goai

import (
	"net/http"
	"reflect"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/os/gstructs"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gmeta"
	"github.com/888go/goframe/util/gtag"
)

// Path 是由 OpenAPI/Swagger 标准版本 3.0 规定的。
type Path struct {
	Ref         string      `json:"$ref,omitempty"`
	Summary     string      `json:"summary,omitempty"`
	Description string      `json:"description,omitempty"`
	Connect     *Operation  `json:"connect,omitempty"`
	Delete      *Operation  `json:"delete,omitempty"`
	Get         *Operation  `json:"get,omitempty"`
	Head        *Operation  `json:"head,omitempty"`
	Options     *Operation  `json:"options,omitempty"`
	Patch       *Operation  `json:"patch,omitempty"`
	Post        *Operation  `json:"post,omitempty"`
	Put         *Operation  `json:"put,omitempty"`
	Trace       *Operation  `json:"trace,omitempty"`
	Servers     Servers     `json:"servers,omitempty"`
	Parameters  Parameters  `json:"parameters,omitempty"`
	XExtensions XExtensions `json:"-"`
}

// 路径由OpenAPI/Swagger标准版本3.0指定。
type Paths map[string]Path

const (
	responseOkKey = `200`
)

type addPathInput struct {
	Path     string      // 精确路由路径。
	Prefix   string      // 路由路径前缀。
	Method   string      // Route method.
	Function interface{} // 统一化函数
}

func (oai *OpenApiV3) addPath(in addPathInput) error {
	if oai.Paths == nil {
		oai.Paths = map[string]Path{}
	}

	var reflectType = reflect.TypeOf(in.Function)
	if reflectType.NumIn() != 2 || reflectType.NumOut() != 2 {
		return gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`unsupported function "%s" for OpenAPI Path register, there should be input & output structures`,
			reflectType.String(),
		)
	}
	var (
		inputObject  reflect.Value
		outputObject reflect.Value
	)
	// 根据输入/输出类型创建实例。
	if reflectType.In(1).Kind() == reflect.Ptr {
		inputObject = reflect.New(reflectType.In(1).Elem()).Elem()
	} else {
		inputObject = reflect.New(reflectType.In(1)).Elem()
	}
	if reflectType.Out(0).Kind() == reflect.Ptr {
		outputObject = reflect.New(reflectType.Out(0).Elem()).Elem()
	} else {
		outputObject = reflect.New(reflectType.Out(0)).Elem()
	}

	var (
		mime                 string
		path                 = Path{XExtensions: make(XExtensions)}
		inputMetaMap         = gmeta.Data(inputObject.Interface())
		outputMetaMap        = gmeta.Data(outputObject.Interface())
		isInputStructEmpty   = oai.doesStructHasNoFields(inputObject.Interface())
		inputStructTypeName  = oai.golangTypeToSchemaName(inputObject.Type())
		outputStructTypeName = oai.golangTypeToSchemaName(outputObject.Type())
		operation            = Operation{
			Responses:   map[string]ResponseRef{},
			XExtensions: make(XExtensions),
		}
		seRequirement = SecurityRequirement{}
	)
	// Path check.
	if in.Path == "" {
		in.Path = gmeta.Get(inputObject.Interface(), gtag.Path).String()
		if in.Prefix != "" {
			in.Path = gstr.TrimRight(in.Prefix, "/") + "/" + gstr.TrimLeft(in.Path, "/")
		}
	}
	if in.Path == "" {
		return gerror.NewCodef(
			gcode.CodeMissingParameter,
			`missing necessary path parameter "%s" for input struct "%s", missing tag in attribute Meta?`,
			gtag.Path, inputStructTypeName,
		)
	}

	if v, ok := oai.Paths[in.Path]; ok {
		path = v
	}

	// Method check.
	if in.Method == "" {
		in.Method = gmeta.Get(inputObject.Interface(), gtag.Method).String()
	}
	if in.Method == "" {
		return gerror.NewCodef(
			gcode.CodeMissingParameter,
			`missing necessary method parameter "%s" for input struct "%s", missing tag in attribute Meta?`,
			gtag.Method, inputStructTypeName,
		)
	}

	if err := oai.addSchema(inputObject.Interface(), outputObject.Interface()); err != nil {
		return err
	}

	if len(inputMetaMap) > 0 {
// Path（路径）和 Operation（操作）不是同一概念，因此有必要从 Operation 中为 Path 复制一份 Meta 信息并进行编辑。
// 另外需要注意的是，我们在 Operation 上设置 Summary（摘要）和 Description（描述），而不是在 Path 上设置，所以我们需要将它们移除。
		inputMetaMapForPath := gmap.NewStrStrMapFrom(inputMetaMap).Clone()
		inputMetaMapForPath.Removes([]string{
			gtag.SummaryShort,
			gtag.SummaryShort2,
			gtag.Summary,
			gtag.DescriptionShort,
			gtag.DescriptionShort2,
			gtag.Description,
		})
		if err := oai.tagMapToPath(inputMetaMapForPath.Map(), &path); err != nil {
			return err
		}

		if err := oai.tagMapToOperation(inputMetaMap, &operation); err != nil {
			return err
		}
		// 允许的请求MIME类型。
		if mime = inputMetaMap[gtag.Mime]; mime == "" {
			mime = inputMetaMap[gtag.Consumes]
		}
	}

// 路径安全
// 注：安全模式类型仅支持http和apiKey，不支持oauth2和openIdConnect。
// 多种模式之间用逗号分隔，例如 `security: apiKey1,apiKey2`
	TagNameSecurity := gmeta.Get(inputObject.Interface(), gtag.Security).String()
	securities := gstr.SplitAndTrim(TagNameSecurity, ",")
	for _, sec := range securities {
		seRequirement[sec] = []string{}
	}
	if len(securities) > 0 {
		operation.Security = &SecurityRequirements{seRequirement}
	}

// =================================================================================================================
// 请求参数.
// =================================================================================================================
	structFields, _ := gstructs.Fields(gstructs.FieldsInput{
		Pointer:         inputObject.Interface(),
		RecursiveOption: gstructs.RecursiveOptionEmbeddedNoTag,
	})
	for _, structField := range structFields {
		if operation.Parameters == nil {
			operation.Parameters = []ParameterRef{}
		}
		parameterRef, err := oai.newParameterRefWithStructMethod(structField, in.Path, in.Method)
		if err != nil {
			return err
		}
		if parameterRef != nil {
			operation.Parameters = append(operation.Parameters, *parameterRef)
		}
	}

// =================================================================================================================
// 请求体.
// =================================================================================================================
	if operation.RequestBody == nil {
		operation.RequestBody = &RequestBodyRef{}
	}
	if operation.RequestBody.Value == nil {
		var (
			requestBody = RequestBody{
				Required: true,
				Content:  map[string]MediaType{},
			}
		)
		// 请求支持的MIME类型。
		var (
			contentTypes = oai.Config.ReadContentTypes
			tagMimeValue = gmeta.Get(inputObject.Interface(), gtag.Mime).String()
		)
		if tagMimeValue != "" {
			contentTypes = gstr.SplitAndTrim(tagMimeValue, ",")
		}
		for _, v := range contentTypes {
			if isInputStructEmpty {
				requestBody.Content[v] = MediaType{}
			} else {
				schemaRef, err := oai.getRequestSchemaRef(getRequestSchemaRefInput{
					BusinessStructName: inputStructTypeName,
					RequestObject:      oai.Config.CommonRequest,
					RequestDataField:   oai.Config.CommonRequestDataField,
				})
				if err != nil {
					return err
				}
				requestBody.Content[v] = MediaType{
					Schema: schemaRef,
				}
			}
		}
		operation.RequestBody = &RequestBodyRef{
			Value: &requestBody,
		}
	}

// =================================================================================================================
// 响应.
// =================================================================================================================
	if _, ok := operation.Responses[responseOkKey]; !ok {
		var (
			response = Response{
				Content:     map[string]MediaType{},
				XExtensions: make(XExtensions),
			}
		)
		if len(outputMetaMap) > 0 {
			if err := oai.tagMapToResponse(outputMetaMap, &response); err != nil {
				return err
			}
		}
		// 响应支持的MIME类型。
		var (
			contentTypes = oai.Config.ReadContentTypes
			tagMimeValue = gmeta.Get(outputObject.Interface(), gtag.Mime).String()
			refInput     = getResponseSchemaRefInput{
				BusinessStructName:      outputStructTypeName,
				CommonResponseObject:    oai.Config.CommonResponse,
				CommonResponseDataField: oai.Config.CommonResponseDataField,
			}
		)
		if tagMimeValue != "" {
			contentTypes = gstr.SplitAndTrim(tagMimeValue, ",")
		}
		for _, v := range contentTypes {
			// 如果指定了自定义的响应MIME类型，则会忽略通用的响应特性。
			if tagMimeValue != "" {
				refInput.CommonResponseObject = nil
				refInput.CommonResponseDataField = ""
			}
			schemaRef, err := oai.getResponseSchemaRef(refInput)
			if err != nil {
				return err
			}
			response.Content[v] = MediaType{
				Schema: schemaRef,
			}
		}
		operation.Responses[responseOkKey] = ResponseRef{Value: &response}
	}

	// 移除操作体中重复的属性。
	oai.removeOperationDuplicatedProperties(operation)

	// 给特定操作属性赋值。
	switch gstr.ToUpper(in.Method) {
	case http.MethodGet:
		// GET方法不能包含请求体。
		operation.RequestBody = nil
		path.Get = &operation

	case http.MethodPut:
		path.Put = &operation

	case http.MethodPost:
		path.Post = &operation

	case http.MethodDelete:
		// DELETE操作不能包含请求体。
		operation.RequestBody = nil
		path.Delete = &operation

	case http.MethodConnect:
		// 对于Connect无需执行任何操作。

	case http.MethodHead:
		path.Head = &operation

	case http.MethodOptions:
		path.Options = &operation

	case http.MethodPatch:
		path.Patch = &operation

	case http.MethodTrace:
		path.Trace = &operation

	default:
		return gerror.NewCodef(gcode.CodeInvalidParameter, `invalid method "%s"`, in.Method)
	}
	oai.Paths[in.Path] = path
	return nil
}

func (oai *OpenApiV3) removeOperationDuplicatedProperties(operation Operation) {
	if len(operation.Parameters) == 0 {
		// Nothing to do.
		return
	}

	var (
		duplicatedParameterNames []interface{}
		dataField                string
	)

	for _, parameter := range operation.Parameters {
		duplicatedParameterNames = append(duplicatedParameterNames, parameter.Value.Name)
	}

	// 检查操作请求体中是否包含通用请求数据字段。
	dataFields := gstr.Split(oai.Config.CommonRequestDataField, ".")
	if len(dataFields) > 0 && dataFields[0] != "" {
		dataField = dataFields[0]
	}

	for _, requestBodyContent := range operation.RequestBody.Value.Content {
		// Check request body schema
		if requestBodyContent.Schema == nil {
			continue
		}

		// 检查请求体schema引用。
		if requestBodyContent.Schema.Ref != "" {
			if schema := oai.Components.Schemas.Get(requestBodyContent.Schema.Ref); schema != nil {
				newSchema := schema.Value.Clone()
				requestBodyContent.Schema.Ref = ""
				requestBodyContent.Schema.Value = newSchema
				newSchema.Required = oai.removeItemsFromArray(newSchema.Required, duplicatedParameterNames)
				newSchema.Properties.Removes(duplicatedParameterNames)
				continue
			}
		}

		// 检查请求体中Value公共字段的值。
		if commonRequest := requestBodyContent.Schema.Value.Properties.Get(dataField); commonRequest != nil {
			commonRequest.Value.Required = oai.removeItemsFromArray(commonRequest.Value.Required, duplicatedParameterNames)
			commonRequest.Value.Properties.Removes(duplicatedParameterNames)
			continue
		}

		// 检查请求体中 schema 值。
		if requestBodyContent.Schema.Value != nil {
			requestBodyContent.Schema.Value.Required = oai.removeItemsFromArray(requestBodyContent.Schema.Value.Required, duplicatedParameterNames)
			requestBodyContent.Schema.Value.Properties.Removes(duplicatedParameterNames)
			continue
		}
	}
}

func (oai *OpenApiV3) removeItemsFromArray(array []string, items []interface{}) []string {
	arr := garray.NewStrArrayFrom(array)
	for _, item := range items {
		if value, ok := item.(string); ok {
			arr.RemoveValue(value)
		}
	}
	return arr.Slice()
}

func (oai *OpenApiV3) doesStructHasNoFields(s interface{}) bool {
	return reflect.TypeOf(s).NumField() == 0
}

func (oai *OpenApiV3) tagMapToPath(tagMap map[string]string, path *Path) error {
	var mergedTagMap = oai.fillMapWithShortTags(tagMap)
	if err := gconv.Struct(mergedTagMap, path); err != nil {
		return gerror.Wrap(err, `mapping struct tags to Path failed`)
	}
	oai.tagMapToXExtensions(mergedTagMap, path.XExtensions)
	return nil
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (p Path) MarshalJSON() ([]byte, error) {
	var (
		b   []byte
		m   map[string]json.RawMessage
		err error
	)
	type tempPath Path // 为防止JSON序列化时出现递归错误
	if b, err = json.Marshal(tempPath(p)); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	for k, v := range p.XExtensions {
		if b, err = json.Marshal(v); err != nil {
			return nil, err
		}
		m[k] = b
	}
	return json.Marshal(m)
}
