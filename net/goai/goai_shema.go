// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package goai

import (
	"reflect"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/utils"
	"github.com/888go/goframe/os/gstructs"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gmeta"
	"github.com/888go/goframe/util/gvalid"
)

// Schema遵循OpenAPI/Swagger 3.0标准进行定义。
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

// Clone仅克隆必要的属性。
// TODO：克隆所有属性，或改进包deepcopy。
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
	type tempSchema Schema // 为防止JSON序列化时出现递归错误
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

// Discriminator 是由 OpenAPI/Swagger 3.0 标准指定的。
type Discriminator struct {
	PropertyName string            `json:"propertyName"`
	Mapping      map[string]string `json:"mapping,omitempty"`
}

// addSchema 创建包含对象的模式。
// 注意，`object` 可能是数组别名，例如：`type Res []Item`。
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
		oai.Components.Schemas.refs = map类.X创建链表mp()
	}

	var (
		reflectType    = reflect.TypeOf(object)
		structTypeName = oai.golangTypeToSchemaName(reflectType)
	)

	// Already added.
	if oai.Components.Schemas.Get(structTypeName) != nil {
		return nil
	}
	// 首先获取持有者。
	oai.Components.Schemas.X设置值(structTypeName, SchemaRef{})

	schema, err := oai.structToSchema(object)
	if err != nil {
		return err
	}

	oai.Components.Schemas.X设置值(structTypeName, SchemaRef{
		Ref:   "",
		Value: schema,
	})
	return nil
}

// structToSchema将给定的结构体对象转换并返回为Schema。
func (oai *OpenApiV3) structToSchema(object interface{}) (*Schema, error) {
	var (
		tagMap = 元数据类.Data(object)
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
		if !文本类.X是否大写字符(structField.Name()[0]) {
			continue
		}
		var fieldName = structField.Name()
		for _, tagName := range 转换类.StructTagPriority {
			if tagValue := structField.Tag(tagName); tagValue != "" {
				fieldName = tagValue
				break
			}
		}
		fieldName = 文本类.X分割(文本类.X过滤首尾符并含空白(fieldName), ",")[0]
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
		schema.Properties.X设置值(fieldName, *schemaRef)
	}

	schema.Properties.X遍历(func(key string, ref SchemaRef) bool {
		if ref.Value != nil && ref.Value.ValidationRules != "" {
			validationRuleSet := 集合类.X创建文本并按值(文本类.X分割(ref.Value.ValidationRules, "|"))
			if validationRuleSet.X是否存在(validationRuleKeyForRequired) {
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
	if err := 转换类.Struct(mergedTagMap, schema); err != nil {
		return 错误类.X多层错误(err, `mapping struct tags to Schema failed`)
	}
	oai.tagMapToXExtensions(mergedTagMap, schema.XExtensions)
	// 将验证信息转换为OpenAPI模式规范
	for _, tag := range 效验类.GetTags() {
		if validationTagValue, ok := tagMap[tag]; ok {
			_, validationRules, _ := 效验类.ParseTagValue(validationTagValue)
			schema.ValidationRules = validationRules
			// Enum checks.
			if len(schema.Enum) == 0 {
				for _, rule := range 文本类.X分割并忽略空值(validationRules, "|") {
					if 文本类.X开头判断(rule, validationRuleKeyForIn) {
						var (
							isAllEnumNumber = true
							enumArray       = 文本类.X分割并忽略空值(rule[len(validationRuleKeyForIn):], ",")
						)
						for _, enum := range enumArray {
							if !文本类.X是否为数字(enum) {
								isAllEnumNumber = false
								break
							}
						}
						if isAllEnumNumber {
							schema.Enum = 转换类.X取any数组(转换类.X取整数64位数组(enumArray))
						} else {
							schema.Enum = 转换类.X取any数组(enumArray)
						}
					}
				}
			}
			break
		}
	}
	return nil
}
