// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gstructs提供了用于获取结构体信息的函数。 md5:ff6813ae5e3453ba
package gstructs

import (
	"reflect"

	"github.com/gogf/gf/v2/errors/gerror"
)

// Type 是 reflect.Type 的增强版本，提供了更多功能。 md5:8ebe2d126efacb49
type Type struct {
	reflect.Type
}

// Field包含一个结构字段的信息。 md5:937dc69e9da6620a
type Field struct {
	Value reflect.Value       // 字段的基础值。 md5:30c17ef0d945aeca
	Field reflect.StructField // 字段的底层字段。 md5:5d18dc4427e59bd7

	// 从TagValue中获取标签名。这取决于TagValue。 md5:2abab33cf7d9aa7a
	TagName string

// 获取标签值。
// 字段中可能有多个标签，但根据调用函数的规则，只能获取一个。
// md5:45a4365044272532
	TagValue string
}

// FieldsInput 是函数 Fields 的输入参数结构体类型。 md5:dea3d8d32792f883
type FieldsInput struct {
// Pointer 应该是 struct 类型的指针/*指向结构体的指针类型*/。
// TODO 这个属性名称不合适，可能会引起混淆。
// md5:7115141b84d46b4a
	Pointer interface{}

// RecursiveOption 定义了如果属性是一个嵌入的结构体，如何递归地检索字段。默认情况下为 RecursiveOptionNone。
// md5:ad0b9ef1d1f1f1e5
	RecursiveOption RecursiveOption
}

// FieldMapInput是FieldMap函数的输入参数结构体类型。 md5:6265e8efb4329ab9
type FieldMapInput struct {
// Pointer 应该是 struct 类型的指针/*指向结构体的指针类型*/。
// TODO 这个属性名称不合适，可能会引起混淆。
// md5:7115141b84d46b4a
	Pointer interface{}

// PriorityTagArray 用于指定优先级标签数组，按从高到低的顺序检索。
// 如果传入 `nil`，则返回 map[name]Field，其中 `name` 是属性名称。
// md5:454af14097a1e0a3
	PriorityTagArray []string

// RecursiveOption 定义了如果属性是一个嵌入的结构体，如何递归地检索字段。默认情况下为 RecursiveOptionNone。
// md5:ad0b9ef1d1f1f1e5
	RecursiveOption RecursiveOption
}

type RecursiveOption int

const (
	RecursiveOptionNone          RecursiveOption = iota // 如果字段是一个嵌入的结构体，不递归地将其字段作为映射获取。 md5:7e5b33b2b6192298
	RecursiveOptionEmbedded                             // 如果字段是一个嵌入的结构体，递归地获取其字段作为map。 md5:5c2576800c0efe83
	RecursiveOptionEmbeddedNoTag                        // 如果字段是嵌入的结构并且字段没有标签，递归地获取字段作为映射。 md5:24a441218c457b09
)

// Fields 获取并以切片形式返回 `pointer` 的字段。 md5:7856c4ee9e72f56f
// ff:
// in:
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
					// 当前级别字段可以覆盖具有相同名称的子结构体字段。 md5:e9bd19d3abe6f7e5
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

// FieldMap 从 `pointer` 获取并返回结构体字段，作为 map[name/tag]Field。
//
// 参数 `pointer` 应为 struct 或 *struct 类型。
//
// 参数 `priority` 指定了用于检索的优先级标签数组，从高到低。如果为空（`nil`），则返回 map[name]Field，其中 `name` 为属性名称。
//
// 参数 `recursive` 指定是否递归检索嵌入的结构体中的字段。
//
// 注意，它仅从结构体中检索首字母大写的导出属性。
// md5:deef4c5e31602259
// ff:
// in:
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
		// 只检索导出的属性。 md5:d8185f07060feffb
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

// StructType 获取并返回指定结构体的类型。
// 参数 `object` 应为结构体类型、指向结构体的指针类型、结构体切片类型或指向结构体的切片类型。
// md5:023b27218d435b61
// ff:
// object:
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
				// 如果指针是*struct类型且为nil，那么会自动创建一个临时的struct。 md5:23b5ebc131739e7d
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
