// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包gstructs提供了用于获取结构体信息的函数。
// 这个不翻译了, 这是一个偏底层组件，一般业务上很少会用到，在框架、基础库、中间件编写中用到。
package gstructs
import (
	"reflect"
	
	"github.com/888go/goframe/errors/gerror"
	)
// Type 类型封装了 reflect.Type，提供了额外的功能。
type Type struct {
	reflect.Type
}

// Field 包含结构体字段的信息。
type Field struct {
	Value reflect.Value       // 字段的底层值
	Field reflect.StructField // 字段的底层字段。

	// 获取到的标签名称，其依赖于TagValue。
	TagName string

// 获取到的标签值。
// 在该字段中可能包含多个标签，
// 但根据调用函数规则，只能获取其中一个。
	TagValue string
}

// FieldsInput 是函数 Fields 的输入参数结构体类型。
type FieldsInput struct {
// Pointer 应为结构体类型。/*struct. */
// TODO 这个属性名称不合适，可能会引起混淆。
	Pointer interface{}

// RecursiveOption 指定当属性是一个嵌入式结构体时，以何种方式递归获取其字段。默认情况下，其值为 RecursiveOptionNone。
// ```go
// RecursiveOption 指定了在遇到嵌入式结构体属性时，如何进行递归地获取其字段的选项。默认设置为 RecursiveOptionNone。
	RecursiveOption RecursiveOption
}

// FieldMapInput 是函数 FieldMap 的输入参数结构体类型。
type FieldMapInput struct {
// Pointer 应为结构体类型。/*struct. */
// TODO 这个属性名称不合适，可能会引起混淆。
	Pointer interface{}

// PriorityTagArray 指定从高到低检索的优先级标签数组。
// 如果传入 `nil`，则返回 map[name]Field，其中 `name` 是属性名。
	PriorityTagArray []string

// RecursiveOption 指定当属性是一个嵌入式结构体时，以何种方式递归获取其字段。默认情况下，其值为 RecursiveOptionNone。
// ```go
// RecursiveOption 指定了在遇到嵌入式结构体属性时，如何进行递归地获取其字段的选项。默认设置为 RecursiveOptionNone。
	RecursiveOption RecursiveOption
}

type RecursiveOption int

const (
	RecursiveOptionNone          RecursiveOption = 0 // 如果字段是嵌入式结构体，则不递归地以 map 形式获取其字段。
	RecursiveOptionEmbedded      RecursiveOption = 1 // 如果字段是一个嵌入式结构体，则递归地将其字段作为映射获取。
	RecursiveOptionEmbeddedNoTag RecursiveOption = 2 // 如果字段是一个嵌入式结构体且该字段没有标签，则递归获取其字段并以映射形式返回。
)

// Fields 函数检索并返回 `pointer` 的字段作为一个切片。
func Fields(in FieldsInput) ([]Field, error) {
	var (
		ok                   bool
		fieldFilterMap       = make(map[string]struct{})
		retrievedFields      = make([]Field, 0)
		currentLevelFieldMap = make(map[string]Field)
		rangeFields, err     = getFieldValues(in.Pointer)
	)
	if err != nil {
		return nil, err
	}

	for index := 0; index < len(rangeFields); index++ {
		field := rangeFields[index]
		currentLevelFieldMap[field.Name()] = field
	}

	for index := 0; index < len(rangeFields); index++ {
		field := rangeFields[index]
		if _, ok = fieldFilterMap[field.Name()]; ok {
			continue
		}
		if field.IsEmbedded() {
			if in.RecursiveOption != RecursiveOptionNone {
				switch in.RecursiveOption {
				case RecursiveOptionEmbeddedNoTag:
					if field.TagStr() != "" {
						break
					}
					fallthrough

				case RecursiveOptionEmbedded:
					structFields, err := Fields(FieldsInput{
						Pointer:         field.Value,
						RecursiveOption: in.RecursiveOption,
					})
					if err != nil {
						return nil, err
					}
					// 当前层级字段可以覆盖具有相同名称的子结构体字段。
					for i := 0; i < len(structFields); i++ {
						var (
							structField = structFields[i]
							fieldName   = structField.Name()
						)
						if _, ok = fieldFilterMap[fieldName]; ok {
							continue
						}
						fieldFilterMap[fieldName] = struct{}{}
						if v, ok := currentLevelFieldMap[fieldName]; !ok {
							retrievedFields = append(retrievedFields, structField)
						} else {
							retrievedFields = append(retrievedFields, v)
						}
					}
					continue
				}
			}
			continue
		}
		fieldFilterMap[field.Name()] = struct{}{}
		retrievedFields = append(retrievedFields, field)
	}
	return retrievedFields, nil
}

// FieldMap 从`pointer`获取并返回结构体字段为map[name/tag]Field。
//
// 参数`pointer`应为struct/*struct类型。
//
// 参数`priority`指定了按从高到低优先级获取的标签数组。如果给出`nil`，则返回map[name]Field，其中`name`是属性名称。
//
// 参数`recursive`指定了当属性是一个嵌入式结构体时，是否递归地获取其字段。
//
// 注意，它仅从结构体中获取首字母大写的导出属性（即公开属性）。
func FieldMap(in FieldMapInput) (map[string]Field, error) {
	fields, err := getFieldValues(in.Pointer)
	if err != nil {
		return nil, err
	}
	var (
		tagValue string
		mapField = make(map[string]Field)
	)
	for _, field := range fields {
		// 仅获取导出的属性。
		if !field.IsExported() {
			continue
		}
		tagValue = ""
		for _, p := range in.PriorityTagArray {
			tagValue = field.Tag(p)
			if tagValue != "" && tagValue != "-" {
				break
			}
		}
		tempField := field
		tempField.TagValue = tagValue
		if tagValue != "" {
			mapField[tagValue] = tempField
		} else {
			if in.RecursiveOption != RecursiveOptionNone && field.IsEmbedded() {
				switch in.RecursiveOption {
				case RecursiveOptionEmbeddedNoTag:
					if field.TagStr() != "" {
						mapField[field.Name()] = tempField
						break
					}
					fallthrough

				case RecursiveOptionEmbedded:
					m, err := FieldMap(FieldMapInput{
						Pointer:          field.Value,
						PriorityTagArray: in.PriorityTagArray,
						RecursiveOption:  in.RecursiveOption,
					})
					if err != nil {
						return nil, err
					}
					for k, v := range m {
						if _, ok := mapField[k]; !ok {
							tempV := v
							mapField[k] = tempV
						}
					}
				}
			} else {
				mapField[field.Name()] = tempField
			}
		}
	}
	return mapField, nil
}

// StructType 函数检索并返回指定结构体的结构体类型。
// 参数 `object` 应为 struct 类型、指针到 struct 类型、struct 切片或指针到 struct 切片类型。
func StructType(object interface{}) (*Type, error) {
	var (
		reflectValue reflect.Value
		reflectKind  reflect.Kind
		reflectType  reflect.Type
	)
	if rv, ok := object.(reflect.Value); ok {
		reflectValue = rv
	} else {
		reflectValue = reflect.ValueOf(object)
	}
	reflectKind = reflectValue.Kind()
	for {
		switch reflectKind {
		case reflect.Ptr:
			if !reflectValue.IsValid() || reflectValue.IsNil() {
				// 如果指针是结构体类型且为nil，则自动创建一个临时结构体。
				reflectValue = reflect.New(reflectValue.Type().Elem()).Elem()
				reflectKind = reflectValue.Kind()
			} else {
				reflectValue = reflectValue.Elem()
				reflectKind = reflectValue.Kind()
			}

		case reflect.Array, reflect.Slice:
			reflectValue = reflect.New(reflectValue.Type().Elem()).Elem()
			reflectKind = reflectValue.Kind()

		default:
			goto exitLoop
		}
	}

exitLoop:
	if reflectKind != reflect.Struct {
		return nil, gerror.Newf(
			`invalid object kind "%s", kind of "struct" is required`,
			reflectKind,
		)
	}
	reflectType = reflectValue.Type()
	return &Type{
		Type: reflectType,
	}, nil
}
