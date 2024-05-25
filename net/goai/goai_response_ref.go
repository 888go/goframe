// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package goai

import (
	"reflect"

	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/os/gstructs"
	"github.com/gogf/gf/v2/text/gstr"
)

type ResponseRef struct {
	Ref   string
	Value *Response
}

// Responses 是由 OpenAPI/Swagger 3.0 标准指定的。. md5:6a58747083cc2ced
type Responses map[string]ResponseRef

func (r ResponseRef) MarshalJSON() ([]byte, error) {
	if r.Ref != "" {
		return formatRefToBytes(r.Ref), nil
	}
	return json.Marshal(r.Value)
}

type getResponseSchemaRefInput struct {
	BusinessStructName      string      // 业务结构体名称。. md5:189f56b364de535f
	CommonResponseObject    interface{} // 共享响应对象。. md5:ba66a231cf8c9a50
	CommonResponseDataField string      // 公共响应数据字段。. md5:55294824f52f2730
}

func (oai *OpenApiV3) getResponseSchemaRef(in getResponseSchemaRefInput) (*SchemaRef, error) {
	if in.CommonResponseObject == nil {
		return &SchemaRef{
			Ref: in.BusinessStructName,
		}, nil
	}

	var (
		dataFieldsPartsArray       = gstr.Split(in.CommonResponseDataField, ".")
		bizResponseStructSchemaRef = oai.Components.Schemas.Get(in.BusinessStructName)
		schema, err                = oai.structToSchema(in.CommonResponseObject)
	)
	if err != nil {
		return nil, err
	}
	if in.CommonResponseDataField == "" && bizResponseStructSchemaRef != nil {
		// Normal response.
		bizResponseStructSchemaRef.Value.Properties.Iterator(func(key string, ref SchemaRef) bool {
			schema.Properties.Set(key, ref)
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
					schema.Properties.Set(fieldName, *bizResponseStructSchemaRef)
					break
				}
			default:
				// 递归创建通用响应对象 schema。. md5:5e2478af86757d58
				if structField.Name() == dataFieldsPartsArray[0] {
					var structFieldInstance = reflect.New(structField.Type().Type).Elem()
					schemaRef, err := oai.getResponseSchemaRef(getResponseSchemaRefInput{
						BusinessStructName:      in.BusinessStructName,
						CommonResponseObject:    structFieldInstance,
						CommonResponseDataField: gstr.Join(dataFieldsPartsArray[1:], "."),
					})
					if err != nil {
						return nil, err
					}
					schema.Properties.Set(fieldName, *schemaRef)
					break
				}
			}
		}
	}

	return &SchemaRef{
		Value: schema,
	}, nil
}
