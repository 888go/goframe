// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package goai

import (
	"reflect"

	gmap "github.com/888go/goframe/container/gmap"
	gset "github.com/888go/goframe/container/gset"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/utils"
	"github.com/888go/goframe/os/gstructs"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
	gmeta "github.com/888go/goframe/util/gmeta"
	gvalid "github.com/888go/goframe/util/gvalid"
)

// Schema 按照 OpenAPI/Swagger 3.0 标准指定。 md5:6f1f02d1c3d44b09
type Schema struct {
	OneOf                SchemaRefs     `json:"oneOf,omitempty"`
	AnyOf                SchemaRefs     `json:"anyOf,omitempty"`
	AllOf                SchemaRefs     `json:"allOf,omitempty"`
	Not                  *SchemaRef     `json:"not,omitempty"`
	Type                 string         `json:"type,omitempty"`
	Title                string         `json:"title,omitempty"`
	Format               string         `json:"format,omitempty"`
	Description          string         `json:"description,omitempty"`
	Enum                 []interface{}  `json:"enum,omitempty"`
	Default              interface{}    `json:"default,omitempty"`
	Example              interface{}    `json:"example,omitempty"`
	ExternalDocs         *ExternalDocs  `json:"externalDocs,omitempty"`
	UniqueItems          bool           `json:"uniqueItems,omitempty"`
	ExclusiveMin         bool           `json:"exclusiveMinimum,omitempty"`
	ExclusiveMax         bool           `json:"exclusiveMaximum,omitempty"`
	Nullable             bool           `json:"nullable,omitempty"`
	ReadOnly             bool           `json:"readOnly,omitempty"`
	WriteOnly            bool           `json:"writeOnly,omitempty"`
	AllowEmptyValue      bool           `json:"allowEmptyValue,omitempty"`
	XML                  interface{}    `json:"xml,omitempty"`
	Deprecated           bool           `json:"deprecated,omitempty"`
	Min                  *float64       `json:"minimum,omitempty"`
	Max                  *float64       `json:"maximum,omitempty"`
	MultipleOf           *float64       `json:"multipleOf,omitempty"`
	MinLength            uint64         `json:"minLength,omitempty"`
	MaxLength            *uint64        `json:"maxLength,omitempty"`
	Pattern              string         `json:"pattern,omitempty"`
	MinItems             uint64         `json:"minItems,omitempty"`
	MaxItems             *uint64        `json:"maxItems,omitempty"`
	Items                *SchemaRef     `json:"items,omitempty"`
	Required             []string       `json:"required,omitempty"`
	Properties           Schemas        `json:"properties,omitempty"`
	MinProps             uint64         `json:"minProperties,omitempty"`
	MaxProps             *uint64        `json:"maxProperties,omitempty"`
	AdditionalProperties *SchemaRef     `json:"additionalProperties,omitempty"`
	Discriminator        *Discriminator `json:"discriminator,omitempty"`
	XExtensions          XExtensions    `json:"-"`
	ValidationRules      string         `json:"-"`
}

// 只有必要的属性会被克隆。
// TODO 克隆所有属性，或者改进包的深拷贝功能。
// md5:5648a4a6a90c8e18
func (s *Schema) Clone() *Schema {
	newSchema := *s
	newSchema.Required = make([]string, len(s.Required))
	copy(newSchema.Required, s.Required)
	newSchema.Properties = s.Properties.Clone()
	return &newSchema
}

func (s Schema) MarshalJSON() ([]byte, error) {
	var (
		b   []byte
		m   map[string]json.RawMessage
		err error
	)
	type tempSchema Schema // 为了防止JSON序列化时的递归错误。 md5:add9f5a47e638cc5
	if b, err = json.Marshal(tempSchema(s)); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	for k, v := range s.XExtensions {
		if b, err = json.Marshal(v); err != nil {
			return nil, err
		}
		m[k] = b
	}
	return json.Marshal(m)
}

// Discriminator是根据OpenAPI/Swagger标准版本3.0指定的。 md5:7587308535782993
type Discriminator struct {
	PropertyName string            `json:"propertyName"`
	Mapping      map[string]string `json:"mapping,omitempty"`
}

// addSchema 用于创建包含对象的模式。
// 注意，`object` 可以是数组别名，例如：`type Res []Item`。
// md5:c0f033836f564a8c
func (oai *OpenApiV3) addSchema(object ...interface{}) error {
	for _, v := range object {
		if err := oai.doAddSchemaSingle(v); err != nil {
			return err
		}
	}
	return nil
}

func (oai *OpenApiV3) doAddSchemaSingle(object interface{}) error {
	if oai.Components.Schemas.refs == nil {
		oai.Components.Schemas.refs = gmap.NewListMap()
	}

	var (
		reflectType    = reflect.TypeOf(object)
		structTypeName = oai.golangTypeToSchemaName(reflectType)
	)

	// Already added.
	if oai.Components.Schemas.Get(structTypeName) != nil {
		return nil
	}
	// Take the holder first.
	oai.Components.Schemas.Set(structTypeName, SchemaRef{})

	schema, err := oai.structToSchema(object)
	if err != nil {
		return err
	}

	oai.Components.Schemas.Set(structTypeName, SchemaRef{
		Ref:   "",
		Value: schema,
	})
	return nil
}

// structToSchema 将给定的结构体对象转换并返回为Schema。 md5:c3023fab7f0fbf3b
func (oai *OpenApiV3) structToSchema(object interface{}) (*Schema, error) {
	var (
		tagMap = gmeta.Data(object)
		schema = &Schema{
			Properties:  createSchemas(),
			XExtensions: make(XExtensions),
		}
		ignoreProperties []interface{}
	)
	if len(tagMap) > 0 {
		if err := oai.tagMapToSchema(tagMap, schema); err != nil {
			return nil, err
		}
	}
	if schema.Type != "" && schema.Type != TypeObject {
		return schema, nil
	}
	// []struct.
	if utils.IsArray(object) {
		schema.Type = TypeArray
		subSchemaRef, err := oai.newSchemaRefWithGolangType(reflect.TypeOf(object).Elem(), nil)
		if err != nil {
			return nil, err
		}
		schema.Items = subSchemaRef
		if len(schema.Enum) > 0 {
			schema.Items.Value.Enum = schema.Enum
			schema.Enum = nil
		}
		return schema, nil
	}
	// struct.
	structFields, _ := gstructs.Fields(gstructs.FieldsInput{
		Pointer:         object,
		RecursiveOption: gstructs.RecursiveOptionEmbeddedNoTag,
	})
	schema.Type = TypeObject
	for _, structField := range structFields {
		if !gstr.IsLetterUpper(structField.Name()[0]) {
			continue
		}
		var fieldName = structField.TagPriorityName()
		fieldName = gstr.Split(gstr.Trim(fieldName), ",")[0]
		if fieldName == "" {
			fieldName = structField.Name()
		}
		schemaRef, err := oai.newSchemaRefWithGolangType(
			structField.Type().Type,
			structField.TagMap(),
		)
		if err != nil {
			return nil, err
		}
		schema.Properties.Set(fieldName, *schemaRef)
	}

	schema.Properties.Iterator(func(key string, ref SchemaRef) bool {
		if ref.Value != nil && ref.Value.ValidationRules != "" {
			validationRuleSet := gset.NewStrSetFrom(gstr.Split(ref.Value.ValidationRules, "|"))
			if validationRuleSet.Contains(validationRuleKeyForRequired) {
				schema.Required = append(schema.Required, key)
			}
		}
		if !isValidParameterName(key) {
			ignoreProperties = append(ignoreProperties, key)
		}
		return true
	})

	if len(ignoreProperties) > 0 {
		schema.Properties.Removes(ignoreProperties)
	}

	return schema, nil
}

func (oai *OpenApiV3) tagMapToSchema(tagMap map[string]string, schema *Schema) error {
	var mergedTagMap = oai.fillMapWithShortTags(tagMap)
	if err := gconv.Struct(mergedTagMap, schema); err != nil {
		return gerror.Wrap(err, `mapping struct tags to Schema failed`)
	}
	oai.tagMapToXExtensions(mergedTagMap, schema.XExtensions)
		// 验证信息到OpenAPI规范模式。 md5:8caca50de8e752c8
	for _, tag := range gvalid.GetTags() {
		if validationTagValue, ok := tagMap[tag]; ok {
			_, validationRules, _ := gvalid.ParseTagValue(validationTagValue)
			schema.ValidationRules = validationRules
			// Enum checks.
			if len(schema.Enum) == 0 {
				for _, rule := range gstr.SplitAndTrim(validationRules, "|") {
					if gstr.HasPrefix(rule, validationRuleKeyForIn) {
						var (
							isAllEnumNumber = true
							enumArray       = gstr.SplitAndTrim(rule[len(validationRuleKeyForIn):], ",")
						)
						for _, enum := range enumArray {
							if !gstr.IsNumeric(enum) {
								isAllEnumNumber = false
								break
							}
						}
						if isAllEnumNumber {
							schema.Enum = gconv.Interfaces(gconv.Int64s(enumArray))
						} else {
							schema.Enum = gconv.Interfaces(enumArray)
						}
					}
				}
			}
			break
		}
	}
	return nil
}
