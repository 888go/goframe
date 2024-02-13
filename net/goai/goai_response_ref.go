// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package goai

import (
	"reflect"
	
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/os/gstructs"
	"github.com/888go/goframe/text/gstr"
)

type ResponseRef struct {
	Ref   string
	Value *Response
}

// Responses 按照 OpenAPI/Swagger 3.0 标准进行指定。
type Responses map[string]ResponseRef

func (r ResponseRef) MarshalJSON() ([]byte, error) {
	if r.Ref != "" {
		return formatRefToBytes(r.Ref), nil
	}
	return json.Marshal(r.Value)
}

type getResponseSchemaRefInput struct {
	BusinessStructName      string      // 业务结构体名称。
	CommonResponseObject    interface{} // 常用响应对象。
	CommonResponseDataField string      // 公共响应数据字段。
}

func (oai *OpenApiV3) getResponseSchemaRef(in getResponseSchemaRefInput) (*SchemaRef, error) {
	if in.CommonResponseObject == nil {
		return &SchemaRef{
			Ref: in.BusinessStructName,
		}, nil
	}

	var (
		dataFieldsPartsArray       = 文本类.X分割(in.CommonResponseDataField, ".")
		bizResponseStructSchemaRef = oai.Components.Schemas.Get(in.BusinessStructName)
		schema, err                = oai.structToSchema(in.CommonResponseObject)
	)
	if err != nil {
		return nil, err
	}
	if in.CommonResponseDataField == "" && bizResponseStructSchemaRef != nil {
		// Normal response.
		bizResponseStructSchemaRef.Value.Properties.X遍历(func(key string, ref SchemaRef) bool {
			schema.Properties.X设置值(key, ref)
			return true
		})
	} else {
		// Common response.
		structFields, _ := gstructs.Fields(gstructs.FieldsInput{
			Pointer:         in.CommonResponseObject,
			RecursiveOption: gstructs.RecursiveOptionEmbeddedNoTag,
		})
		for _, structField := range structFields {
			var fieldName = structField.Name()
			if jsonName := structField.TagJsonName(); jsonName != "" {
				fieldName = jsonName
			}
			switch len(dataFieldsPartsArray) {
			case 1:
				if structField.Name() == dataFieldsPartsArray[0] {
					if err = oai.tagMapToSchema(structField.TagMap(), bizResponseStructSchemaRef.Value); err != nil {
						return nil, err
					}
					schema.Properties.X设置值(fieldName, *bizResponseStructSchemaRef)
					break
				}
			default:
				// 递归创建通用响应对象模式
				if structField.Name() == dataFieldsPartsArray[0] {
					var structFieldInstance = reflect.New(structField.Type().Type).Elem()
					schemaRef, err := oai.getResponseSchemaRef(getResponseSchemaRefInput{
						BusinessStructName:      in.BusinessStructName,
						CommonResponseObject:    structFieldInstance,
						CommonResponseDataField: 文本类.X连接(dataFieldsPartsArray[1:], "."),
					})
					if err != nil {
						return nil, err
					}
					schema.Properties.X设置值(fieldName, *schemaRef)
					break
				}
			}
		}
	}

	return &SchemaRef{
		Value: schema,
	}, nil
}
