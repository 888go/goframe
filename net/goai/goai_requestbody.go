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

// requestBody 是由 OpenAPI/Swagger 3.0 标准定义的。 md5:c7f34291e4ca763f
type RequestBody struct {
	Description string  `json:"description,omitempty"`
	Required    bool    `json:"required,omitempty"`
	Content     Content `json:"content,omitempty"`
}

type RequestBodyRef struct {
	Ref   string
	Value *RequestBody
}

// ff:
// r:
func (r RequestBodyRef) MarshalJSON() ([]byte, error) {
	if r.Ref != "" {
		return formatRefToBytes(r.Ref), nil
	}
	return json.Marshal(r.Value)
}

type getRequestSchemaRefInput struct {
	BusinessStructName string
	RequestObject      interface{}
	RequestDataField   string
}

func (oai *OpenApiV3) getRequestSchemaRef(in getRequestSchemaRefInput) (*SchemaRef, error) {
	if oai.Config.CommonRequest == nil {
		return &SchemaRef{
			Ref: in.BusinessStructName,
		}, nil
	}

	var (
		dataFieldsPartsArray      = gstr.Split(in.RequestDataField, ".")
		bizRequestStructSchemaRef = oai.Components.Schemas.Get(in.BusinessStructName)
		schema, err               = oai.structToSchema(in.RequestObject)
	)
	if err != nil {
		return nil, err
	}
	if in.RequestDataField == "" && bizRequestStructSchemaRef != nil {
		// Normal request.
		bizRequestStructSchemaRef.Value.Properties.Iterator(func(key string, ref SchemaRef) bool {
			schema.Properties.Set(key, ref)
			return true
		})
	} else {
		// Common request.
		structFields, _ := gstructs.Fields(gstructs.FieldsInput{
			Pointer:         in.RequestObject,
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
					if err = oai.tagMapToSchema(structField.TagMap(), bizRequestStructSchemaRef.Value); err != nil {
						return nil, err
					}
					schema.Properties.Set(fieldName, *bizRequestStructSchemaRef)
					break
				}
			default:
				if structField.Name() == dataFieldsPartsArray[0] {
					var structFieldInstance = reflect.New(structField.Type().Type).Elem()
					schemaRef, err := oai.getRequestSchemaRef(getRequestSchemaRefInput{
						BusinessStructName: in.BusinessStructName,
						RequestObject:      structFieldInstance,
						RequestDataField:   gstr.Join(dataFieldsPartsArray[1:], "."),
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
