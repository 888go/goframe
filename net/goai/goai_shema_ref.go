// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package goai

import (
	"fmt"
	"reflect"

	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gtag"
)

type SchemaRefs []SchemaRef

type SchemaRef struct {
	Ref   string
	Value *Schema
}

// isEmbeddedStructDefinition 检查并返回给定的 Go 语言类型是否是嵌入的结构体定义，例如：
// 
//	```go
//	type A struct {
//	    B struct {
//	        // ...
//	    }
//	}
//	```
// 
// 在 `A` 中的 `B` 被称为 `嵌入的结构体定义`。
// md5:45d275bc85e98290
func (oai *OpenApiV3) isEmbeddedStructDefinition(golangType reflect.Type) bool {
	s := golangType.String()
	return gstr.Contains(s, `struct {`)
}

// newSchemaRefWithGolangType 创建一个新的架构（Schema）并返回其架构引用（SchemaRef）。 md5:024c7a7371946acb
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
			schemaRef.Value.Default = gconv.Int64(schemaRef.Value.Default)
		}
		// 将默认值保留为nil。 md5:a85d623b66e78405

		// 示例值需要像默认值一样进行转换. md5:6442bc8222d0ea95
		if schemaRef.Value.Example != nil {
			schemaRef.Value.Example = gconv.Int64(schemaRef.Value.Example)
		}
		// 保持示例值为 nil。 md5:236a31f4aed61b8c
	case TypeNumber:
		if schemaRef.Value.Default != nil {
			schemaRef.Value.Default = gconv.Float64(schemaRef.Value.Default)
		}
		// 将默认值保留为nil。 md5:a85d623b66e78405

		// 示例值需要像默认值一样进行转换. md5:6442bc8222d0ea95
		if schemaRef.Value.Example != nil {
			schemaRef.Value.Example = gconv.Float64(schemaRef.Value.Example)
		}
		// 保持示例值为 nil。 md5:236a31f4aed61b8c
	case TypeBoolean:
		if schemaRef.Value.Default != nil {
			schemaRef.Value.Default = gconv.Bool(schemaRef.Value.Default)
		}
		// 将默认值保留为nil。 md5:a85d623b66e78405

		// 示例值需要像默认值一样进行转换. md5:6442bc8222d0ea95
		if schemaRef.Value.Example != nil {
			schemaRef.Value.Example = gconv.Bool(schemaRef.Value.Example)
		}
		// 保持示例值为 nil。 md5:236a31f4aed61b8c
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
			// 特别针对映射类型。 md5:e7fa8512d15545fb
			subSchemaRef, err := oai.newSchemaRefWithGolangType(golangType.Elem(), nil)
			if err != nil {
				return nil, err
			}
			schema.AdditionalProperties = subSchemaRef
			return schemaRef, nil

		case reflect.Interface:
			// 专门用于接口类型。 md5:dbaf9c5bd34e0ea8
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

// ff:
// r:
func (r SchemaRef) MarshalJSON() ([]byte, error) {
	if r.Ref != "" {
		return formatRefToBytes(r.Ref), nil
	}
	return json.Marshal(r.Value)
}
