// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package goai

import (
	"fmt"
	"reflect"
	
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gtag"
)

type SchemaRefs []SchemaRef

type SchemaRef struct {
	Ref   string
	Value *Schema
}

// isEmbeddedStructDefinition 检查并返回给定的 Go 语言类型是否为嵌入式结构体定义，例如：
//
//	struct A struct{
//	    B struct{
//	   // ...
//	    }
//	}
//
// 在 `A` 中的 `B` 被称为 `嵌入式结构体定义`。
func (oai *OpenApiV3) isEmbeddedStructDefinition(golangType reflect.Type) bool {
	s := golangType.String()
	return 文本类.X是否包含(s, `struct {`)
}

// newSchemaRefWithGolangType 根据 Go 语言类型创建一个新的 Schema，并返回其对应的 SchemaRef。
func (oai *OpenApiV3) newSchemaRefWithGolangType(golangType reflect.Type, tagMap map[string]string) (*SchemaRef, error) {
	var (
		err       error
		oaiType   = oai.golangTypeToOAIType(golangType)
		oaiFormat = oai.golangTypeToOAIFormat(golangType)
		typeName  = golangType.Name()
		pkgPath   = golangType.PkgPath()
		schemaRef = &SchemaRef{}
		schema    = &Schema{
			Type:        oaiType,
			Format:      oaiFormat,
			XExtensions: make(XExtensions),
		}
	)
	if pkgPath == "" {
		switch golangType.Kind() {
		case reflect.Ptr, reflect.Array, reflect.Slice:
			pkgPath = golangType.Elem().PkgPath()
			typeName = golangType.Elem().Name()
		}
	}

	// Type enums.
	var typeId = fmt.Sprintf(`%s.%s`, pkgPath, typeName)
	if enums := gtag.GetEnumsByType(typeId); enums != "" {
		schema.Enum = make([]interface{}, 0)
		if err = json.Unmarshal([]byte(enums), &schema.Enum); err != nil {
			return nil, err
		}
	}

	if len(tagMap) > 0 {
		if err := oai.tagMapToSchema(tagMap, schema); err != nil {
			return nil, err
		}
	}
	schemaRef.Value = schema
	switch oaiType {
	case TypeString, TypeFile:
	// Nothing to do.
	case TypeInteger:
		if schemaRef.Value.Default != nil {
			schemaRef.Value.Default = 转换类.X取整数64位(schemaRef.Value.Default)
		}
		// 保留默认值为 nil。

		// 示例值需要像默认值一样进行转换
		if schemaRef.Value.Example != nil {
			schemaRef.Value.Example = 转换类.X取整数64位(schemaRef.Value.Example)
		}
		// 保留示例值为nil。
	case TypeNumber:
		if schemaRef.Value.Default != nil {
			schemaRef.Value.Default = 转换类.X取小数64位(schemaRef.Value.Default)
		}
		// 保留默认值为 nil。

		// 示例值需要像默认值一样进行转换
		if schemaRef.Value.Example != nil {
			schemaRef.Value.Example = 转换类.X取小数64位(schemaRef.Value.Example)
		}
		// 保留示例值为nil。
	case TypeBoolean:
		if schemaRef.Value.Default != nil {
			schemaRef.Value.Default = 转换类.X取布尔(schemaRef.Value.Default)
		}
		// 保留默认值为 nil。

		// 示例值需要像默认值一样进行转换
		if schemaRef.Value.Example != nil {
			schemaRef.Value.Example = 转换类.X取布尔(schemaRef.Value.Example)
		}
		// 保留示例值为nil。
	case
		TypeArray:
		subSchemaRef, err := oai.newSchemaRefWithGolangType(golangType.Elem(), nil)
		if err != nil {
			return nil, err
		}
		schema.Items = subSchemaRef
		if len(schema.Enum) > 0 {
			schema.Items.Value.Enum = schema.Enum
			schema.Enum = nil
		}

	case
		TypeObject:
		for golangType.Kind() == reflect.Ptr {
			golangType = golangType.Elem()
		}
		switch golangType.Kind() {
		case reflect.Map:
			// 特别为 map 类型设计
			subSchemaRef, err := oai.newSchemaRefWithGolangType(golangType.Elem(), nil)
			if err != nil {
				return nil, err
			}
			schema.AdditionalProperties = subSchemaRef
			return schemaRef, nil

		case reflect.Interface:
			// 特别用于接口类型。
			var (
				structTypeName = oai.golangTypeToSchemaName(golangType)
			)
			if oai.Components.Schemas.Get(structTypeName) == nil {
				if err := oai.addSchema(reflect.New(golangType).Interface()); err != nil {
					return nil, err
				}
			}
			schemaRef.Ref = structTypeName
			schemaRef.Value = nil

		default:
			golangTypeInstance := reflect.New(golangType).Elem().Interface()
			if oai.isEmbeddedStructDefinition(golangType) {
				schema, err = oai.structToSchema(golangTypeInstance)
				if err != nil {
					return nil, err
				}
				schemaRef.Ref = ""
				schemaRef.Value = schema
			} else {
				var structTypeName = oai.golangTypeToSchemaName(golangType)
				if oai.Components.Schemas.Get(structTypeName) == nil {
					if err := oai.addSchema(golangTypeInstance); err != nil {
						return nil, err
					}
				}
				schemaRef.Ref = structTypeName
				schemaRef.Value = nil
			}
		}
	}
	return schemaRef, nil
}

func (r SchemaRef) MarshalJSON() ([]byte, error) {
	if r.Ref != "" {
		return formatRefToBytes(r.Ref), nil
	}
	return json.Marshal(r.Value)
}
