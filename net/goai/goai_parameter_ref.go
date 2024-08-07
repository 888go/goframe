// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package goai

import (
	"fmt"
	"net/http"

	gset "github.com/888go/goframe/container/gset"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/os/gstructs"
	gstr "github.com/888go/goframe/text/gstr"
)

// Parameters 是由 OpenAPI/Swagger 3.0 标准指定的。 md5:40f2fa7f283687bf
type Parameters []ParameterRef

type ParameterRef struct {
	Ref   string
	Value *Parameter
}

func (oai *OpenApiV3) newParameterRefWithStructMethod(field gstructs.Field, path, method string) (*ParameterRef, error) {
	var (
		tagMap    = field.TagMap()
		fieldName = field.TagPriorityName()
	)
	fieldName = gstr.X分割(gstr.X过滤首尾符并含空白(fieldName), ",")[0]
	if fieldName == "" {
		fieldName = field.Name()
	}
	var parameter = &Parameter{
		Name:        fieldName,
		XExtensions: make(XExtensions),
	}
	if len(tagMap) > 0 {
		if err := oai.tagMapToParameter(tagMap, parameter); err != nil {
			return nil, err
		}
	}
	if parameter.In == "" {
				// 自动检测其"in"属性。 md5:681a19858c274871
		if gstr.X是否包含并忽略大小写(path, fmt.Sprintf(`{%s}`, parameter.Name)) {
			parameter.In = ParameterInPath
		} else {
						// 如果方法为"GET/DELETE"，则将参数输入默认设置为"query"。 md5:e11ada17c61b4191
			switch gstr.X到大写(method) {
			case http.MethodGet, http.MethodDelete:
				parameter.In = ParameterInQuery

			default:
				return nil, nil
			}
		}
	}

	switch parameter.In {
	case ParameterInPath:
				// 路径参数所需的。 md5:55007a5c2ee8b9f3
		parameter.Required = true

	case ParameterInCookie, ParameterInHeader, ParameterInQuery:

	default:
		return nil, gerror.X创建错误码并格式化(gcode.CodeInvalidParameter, `invalid tag value "%s" for In`, parameter.In)
	}
		// 必要的模式或内容。 md5:35319ee51590f5ad
	schemaRef, err := oai.newSchemaRefWithGolangType(field.Type().Type, tagMap)
	if err != nil {
		return nil, err
	}
	parameter.Schema = schemaRef

	// Ignore parameter.
	if !isValidParameterName(parameter.Name) {
		return nil, nil
	}

	// Required check.
	if parameter.Schema.Value != nil && parameter.Schema.Value.ValidationRules != "" {
		validationRuleArray := gstr.X分割(parameter.Schema.Value.ValidationRules, "|")
		if gset.X创建文本并按值(validationRuleArray).X是否存在(validationRuleKeyForRequired) {
			parameter.Required = true
		}
	}

	return &ParameterRef{
		Ref:   "",
		Value: parameter,
	}, nil
}

func (r ParameterRef) MarshalJSON() ([]byte, error) {
	if r.Ref != "" {
		return formatRefToBytes(r.Ref), nil
	}
	return json.Marshal(r.Value)
}
