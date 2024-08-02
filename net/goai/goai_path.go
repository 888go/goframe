// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package goai

import (
	"net/http"
	"reflect"

	garray "github.com/888go/goframe/container/garray"
	gmap "github.com/888go/goframe/container/gmap"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/os/gstructs"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
	gmeta "github.com/888go/goframe/util/gmeta"
	"github.com/888go/goframe/util/gtag"
)

// Path遵循OpenAPI/Swagger标准版本3.0。 md5:26b252ebd7fb17bd
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

// 路径按照OpenAPI/Swagger标准版本3.0进行指定。 md5:77c53887ba9bfc0f
type Paths map[string]Path

const (
	responseOkKey = `200`
)

type addPathInput struct {
	Path     string      // Precise route path.
	Prefix   string      // Route path prefix.
	Method   string      // Route method.
	Function interface{} // Uniformed function.
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
		// 根据输入/输出类型创建实例。 md5:f07c2f3124391e08
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
		// 路径（Path）和操作（Operation）不是同一概念，因此需要从操作中复制一个元信息（Meta）到路径，并进行编辑。
		// 你知道的，我们是在操作上设置Summary和Description，而不是在路径上，所以我们需要将它们移除。
		// md5:82d486896b1d65b3
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
		// Allowed request mime.
		if mime = inputMetaMap[gtag.Mime]; mime == "" {
			mime = inputMetaMap[gtag.Consumes]
		}
	}

	// 路径安全
	// 注意：安全模式类型仅支持http和apiKey；不支持oauth2和openIdConnect。
	// 多个模式使用逗号分隔，例如：`security: apiKey1,apiKey2`
	// md5:b64ffa4261f0711d
	TagNameSecurity := gmeta.Get(inputObject.Interface(), gtag.Security).String()
	securities := gstr.SplitAndTrim(TagNameSecurity, ",")
	for _, sec := range securities {
		seRequirement[sec] = []string{}
	}
	if len(securities) > 0 {
		operation.Security = &SecurityRequirements{seRequirement}
	}

	// =================================================================================================================
	// 请求参数。
	// =================================================================================================================
	// md5:c70d5376eecf5c01
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
	// 请求体
	// =================================================================================================================
	// md5:c70baaeba9963b54
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
				// 支持的请求MIME类型。 md5:fd32e8079c221b58
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

	// =======================================================
	// 响应。
	// =======================================================
	// md5:ceb9c442cfbdefa1
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
				// 支持的响应MIME类型。 md5:aefcf019c3abea83
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
						// 如果指定了自定义的响应MIME类型，则会忽略通用响应特性。 md5:c0c25e2bd38f6d7b
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

		// 移除操作体中重复的属性。 md5:976053e0b8002715
	oai.removeOperationDuplicatedProperties(operation)

		// 为特定操作分配属性。 md5:2e40ddbde8a1317e
	switch gstr.ToUpper(in.Method) {
	case http.MethodGet:
				// GET 操作不能有请求体。 md5:efd94c634a1773f9
		operation.RequestBody = nil
		path.Get = &operation

	case http.MethodPut:
		path.Put = &operation

	case http.MethodPost:
		path.Post = &operation

	case http.MethodDelete:
				// DELETE操作不能有requestBody。 md5:29660405e268d3ca
		operation.RequestBody = nil
		path.Delete = &operation

	case http.MethodConnect:
				// 对于Connect，无需执行任何操作。 md5:200e0639d4f11b33

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

		// 检查操作请求体是否包含通用请求数据字段。 md5:3e4ccc578046cc45
	dataFields := gstr.Split(oai.Config.CommonRequestDataField, ".")
	if len(dataFields) > 0 && dataFields[0] != "" {
		dataField = dataFields[0]
	}

	for _, requestBodyContent := range operation.RequestBody.Value.Content {
				// 检查请求体架构. md5:dab7ff5a79f31000
		if requestBodyContent.Schema == nil {
			continue
		}

				// 检查请求体架构. md5:dab7ff5a79f31000 ref.
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

				// 检查请求体中的 Value 公共字段。 md5:dd0253ff15259e4b
		if commonRequest := requestBodyContent.Schema.Value.Properties.Get(dataField); commonRequest != nil {
			commonRequest.Value.Required = oai.removeItemsFromArray(commonRequest.Value.Required, duplicatedParameterNames)
			commonRequest.Value.Properties.Removes(duplicatedParameterNames)
			continue
		}

						// 检查请求体架构. md5:dab7ff5a79f31000
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

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (p Path) MarshalJSON() ([]byte, error) {
	var (
		b   []byte
		m   map[string]json.RawMessage
		err error
	)
	type tempPath Path // 为了防止JSON序列化时的递归错误。 md5:add9f5a47e638cc5
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
