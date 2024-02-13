// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package goai

import (
	"fmt"
	"net/http"
	
	"github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/os/gstructs"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
)

// 参数遵循OpenAPI/Swagger 3.0标准进行指定。
type Parameters []ParameterRef

type ParameterRef struct {
	Ref   string
	Value *Parameter
}

func (oai *OpenApiV3) newParameterRefWithStructMethod(field gstructs.Field, path, method string) (*ParameterRef, error) {
	var (
		tagMap    = field.TagMap()
		fieldName = field.Name()
	)
	for _, tagName := range 转换类.StructTagPriority {
		if tagValue := field.Tag(tagName); tagValue != "" {
			fieldName = tagValue
			break
		}
	}
	fieldName = 文本类.X分割(文本类.X过滤首尾符并含空白(fieldName), ",")[0]
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
		// 自动检测其 "in" 属性。
		if 文本类.X是否包含并忽略大小写(path, fmt.Sprintf(`{%s}`, parameter.Name)) {
			parameter.In = ParameterInPath
		} else {
			// 如果请求方法为 "GET" 或 "DELETE"，则将参数输入默认设置为 "query"。
			switch 文本类.X到大写(method) {
			case http.MethodGet, http.MethodDelete:
				parameter.In = ParameterInQuery

			default:
				return nil, nil
			}
		}
	}

	switch parameter.In {
	case ParameterInPath:
		// 用于路径参数，这是必需的。
		parameter.Required = true

	case ParameterInCookie, ParameterInHeader, ParameterInQuery:

	default:
		return nil, 错误类.X创建错误码并格式化(错误码类.CodeInvalidParameter, `invalid tag value "%s" for In`, parameter.In)
	}
	// 必要的架构或内容。
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
		validationRuleArray := 文本类.X分割(parameter.Schema.Value.ValidationRules, "|")
		if 集合类.X创建文本并按值(validationRuleArray).X是否存在(validationRuleKeyForRequired) {
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
