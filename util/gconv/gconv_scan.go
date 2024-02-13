// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 转换类

import (
	"database/sql"
	"reflect"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/utils"
	"github.com/888go/goframe/os/gstructs"
)

// Scan 自动检查 `pointer` 的类型并将 `params` 转换为 `pointer`。它支持将 `params` 转换为以下类型的 `pointer`：
// *map/*[]map/*[]*map/*struct/**struct/*[]struct/*[]*struct。
//
// 如果 `pointer` 类型为 *map，Scan 内部会调用函数 `doMapToMap` 进行转换。
// 如果 `pointer` 类型为 *[]map 或 *[]*map，Scan 内部会调用函数 `doMapToMaps` 进行转换。
// 如果 `pointer` 类型为 *struct 或 **struct，Scan 内部会调用函数 `doStruct` 进行转换。
// 如果 `pointer` 类型为 *[]struct 或 *[]*struct，Scan 内部会调用函数 `doStructs` 进行转换。
func Scan(params interface{}, pointer interface{}, paramKeyToAttrMap ...map[string]string) (err error) {
	var (
		pointerType  reflect.Type
		pointerKind  reflect.Kind
		pointerValue reflect.Value
	)
	if v, ok := pointer.(reflect.Value); ok {
		pointerValue = v
		pointerType = v.Type()
	} else {
		pointerValue = reflect.ValueOf(pointer)
		pointerType = reflect.TypeOf(pointer) // 不要使用pointerValue.Type()，因为pointerValue可能为零值。
	}

	if pointerType == nil {
		return 错误类.X创建错误码(错误码类.CodeInvalidParameter, "parameter pointer should not be nil")
	}
	pointerKind = pointerType.Kind()
	if pointerKind != reflect.Ptr {
		if pointerValue.CanAddr() {
			pointerValue = pointerValue.Addr()
			pointerType = pointerValue.Type()
			pointerKind = pointerType.Kind()
		} else {
			return 错误类.X创建错误码并格式化(
				错误码类.CodeInvalidParameter,
				"params should be type of pointer, but got type: %v",
				pointerType,
			)
		}
	}
	// 直接赋值检查！
	var (
		paramsType  reflect.Type
		paramsValue reflect.Value
	)
	if v, ok := params.(reflect.Value); ok {
		paramsValue = v
		paramsType = paramsValue.Type()
	} else {
		paramsValue = reflect.ValueOf(params)
		paramsType = reflect.TypeOf(params) // 不要使用 paramsValue.Type()，因为 paramsValue 可能为零值。
	}
// 如果`params`和`pointer`是相同类型，则直接进行赋值操作。
// 为了提升性能。
	var (
		pointerValueElem = pointerValue.Elem()
	)
	if pointerValueElem.CanSet() && paramsType == pointerValueElem.Type() {
		pointerValueElem.Set(paramsValue)
		return nil
	}

	// Converting.
	var (
		pointerElem               = pointerType.Elem()
		pointerElemKind           = pointerElem.Kind()
		keyToAttributeNameMapping map[string]string
	)
	if len(paramKeyToAttrMap) > 0 {
		keyToAttributeNameMapping = paramKeyToAttrMap[0]
	}
	switch pointerElemKind {
	case reflect.Map:
		return doMapToMap(params, pointer, paramKeyToAttrMap...)

	case reflect.Array, reflect.Slice:
		var (
			sliceElem     = pointerElem.Elem()
			sliceElemKind = sliceElem.Kind()
		)
		for sliceElemKind == reflect.Ptr {
			sliceElem = sliceElem.Elem()
			sliceElemKind = sliceElem.Kind()
		}
		if sliceElemKind == reflect.Map {
			return doMapToMaps(params, pointer, paramKeyToAttrMap...)
		}
		return doStructs(params, pointer, keyToAttributeNameMapping, "")

	default:
		return doStruct(params, pointer, keyToAttributeNameMapping, "")
	}
}

// ScanList 将 `structSlice` 转换为包含其他复杂结构体属性的结构体切片。
// 注意，参数 `structSlicePointer` 应该是 *[]struct 或 *[]*struct 类型。
//
// 使用示例 1：普通属性结构体关联：
//
//	定义 EntityUser 结构体，包含 Uid 和 Name 属性
//	定义 EntityUserDetail 结构体，包含 Uid 和 Address 属性
//	定义 EntityUserScores 结构体，包含 Id、Uid、Score 和 Course 属性
//	定义 Entity 结构体，包含 User（指向 EntityUser 的指针）、UserDetail（指向 EntityUserDetail 的指针）和 UserScores（EntityUserScores 的指针切片）
//
//	var users []*Entity
//	var userRecords = EntityUser{Uid: 1, Name:"john"}
//	var detailRecords = EntityUserDetail{Uid: 1, Address: "chengdu"}
//	var scoresRecords = EntityUserScores{Id: 1, Uid: 1, Score: 100, Course: "math"}
//	ScanList(userRecords, &users, "User")
//	ScanList(detailRecords, &users, "User", "uid")
//	ScanList(scoresRecords, &users, "UserScores", "User", "uid:Uid")
//	ScanList(scoresRecords, &users, "UserScores", "User", "uid")
//
// 使用示例 2：嵌入式属性结构体关联：
//
//	重新定义 EntityUser、EntityUserDetail 和 EntityUserScores 结构体
//	定义 Entity 结构体，其中包含嵌入的 EntityUser、UserDetail（EntityUserDetail 类型）和 UserScores（EntityUserScores 切片类型）
//
//	var userRecords = EntityUser{Uid: 1, Name:"john"}
//	var detailRecords = EntityUserDetail{Uid: 1, Address: "chengdu"}
//	var scoresRecords = EntityUserScores{Id: 1, Uid: 1, Score: 100}
//	ScanList(userRecords, &users)
//	ScanList(detailRecords, &users, "UserDetail", "uid")
//	ScanList(scoresRecords, &users, "UserScores", "uid")
//
// 示例代码中的 "User/UserDetail/UserScores" 参数用于指定当前结果将绑定的目标属性结构体。
//
// 示例代码中的 "uid" 是结果中的表字段名，而 "Uid" 是相关结构体属性名——不是目标绑定的属性名。在示例中，它是实体 "Entity" 中 "User" 的属性名 "Uid"。它会根据给定的 `relation` 参数自动计算 HasOne/HasMany 关系。
//
// 为了清晰理解此函数，请参阅示例或单元测试用例。
func ScanList(structSlice interface{}, structSlicePointer interface{}, bindToAttrName string, relationAttrNameAndFields ...string) (err error) {
	var (
		relationAttrName string
		relationFields   string
	)
	switch len(relationAttrNameAndFields) {
	case 2:
		relationAttrName = relationAttrNameAndFields[0]
		relationFields = relationAttrNameAndFields[1]
	case 1:
		relationFields = relationAttrNameAndFields[0]
	}
	return doScanList(structSlice, structSlicePointer, bindToAttrName, relationAttrName, relationFields)
}

// doScanList 将 `structSlice` 递归地转换为包含其他复杂结构体属性的结构体切片。
// 注意，参数 `structSlicePointer` 的类型应为 *[]struct 或 *[]*struct。
// 这段代码注释的中文翻译如下：
// ```go
// doScanList 函数将 `structSlice` 转换为一个结构体切片，该切片会递归地包含其他的复杂结构体属性。
// 需要注意的是，传入参数 `structSlicePointer` 的类型应当是指向结构体切片的指针，即 *[]struct 或者 *[]*struct 类型。
func doScanList(
	structSlice interface{}, structSlicePointer interface{}, bindToAttrName, relationAttrName, relationFields string,
) (err error) {
	var (
		maps = X取Map数组(structSlice)
	)
	if len(maps) == 0 {
		return nil
	}
	// 对参数进行必要的检查。
	if bindToAttrName == "" {
		return 错误类.X创建错误码(错误码类.CodeInvalidParameter, `bindToAttrName should not be empty`)
	}

	if relationAttrName == "." {
		relationAttrName = ""
	}

	var (
		reflectValue = reflect.ValueOf(structSlicePointer)
		reflectKind  = reflectValue.Kind()
	)
	if reflectKind == reflect.Interface {
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Kind()
	}
	if reflectKind != reflect.Ptr {
		return 错误类.X创建错误码并格式化(
			错误码类.CodeInvalidParameter,
			"structSlicePointer should be type of *[]struct/*[]*struct, but got: %v",
			reflectKind,
		)
	}
	reflectValue = reflectValue.Elem()
	reflectKind = reflectValue.Kind()
	if reflectKind != reflect.Slice && reflectKind != reflect.Array {
		return 错误类.X创建错误码并格式化(
			错误码类.CodeInvalidParameter,
			"structSlicePointer should be type of *[]struct/*[]*struct, but got: %v",
			reflectKind,
		)
	}
	length := len(maps)
	if length == 0 {
		// 指向的切片不为空。
		if reflectValue.Len() > 0 {
// 在这里检查它是否包含已初始化的结构体项。
// 如果为空且无法进行转换，则返回错误以警告开发者。
			if v := reflectValue.Index(0); v.Kind() != reflect.Ptr {
				return sql.ErrNoRows
			}
		}
		// 对于空的结构体切片，不执行任何操作。
		return nil
	}
	var (
		arrayValue    reflect.Value // Like: []*Entity
		arrayItemType reflect.Type  // Like: *Entity
		reflectType   = reflect.TypeOf(structSlicePointer)
	)
	if reflectValue.Len() > 0 {
		arrayValue = reflectValue
	} else {
		arrayValue = reflect.MakeSlice(reflectType.Elem(), length, length)
	}

	// 切片元素项。
	arrayItemType = arrayValue.Index(0).Type()

	// 关系变量。
	var (
		relationDataMap         map[string]interface{}
		relationFromFieldName   string // 示例：relationKV: id:uid -> id
// 这个注释表明了一个键值对的示例，其中关系（relationKV）的键是"id:uid"，对应的值为"id"。在实际应用中，这可能表示一个映射关系，通过用户ID(uid)可以找到对应的ID(id)。
		relationBindToFieldName string // 示例：relationKV: id:uid  -> uid
// （注释翻译：这个字段或者变量表示一种键值对关系，其中键是"id:uid"，值是"uid"）
	)
	if len(relationFields) > 0 {
// 表字段名与属性名之间的关联键字符串，可以使用字符'='或':'连接。
		array := utils.SplitAndTrim(relationFields, "=")
		if len(array) == 1 {
			// 与旧的分隔符 ':' 兼容。
			array = utils.SplitAndTrim(relationFields, ":")
		}
		if len(array) == 1 {
			// 关系名称是相同的。
			array = []string{relationFields, relationFields}
		}
		if len(array) == 2 {
// 定义表格字段到关联属性名称的映射。
// 例如：
// uid:Uid
// uid:UserId
// 这段代码的作用是将数据库表中的字段名（如uid）映射到程序中使用的关联属性名称，以实现字段名称在代码逻辑中的语义化表达。
			relationFromFieldName = array[0]
			relationBindToFieldName = array[1]
			if key, _ := utils.MapPossibleItemByKey(maps[0], relationFromFieldName); key == "" {
				return 错误类.X创建错误码并格式化(
					错误码类.CodeInvalidParameter,
					`cannot find possible related table field name "%s" from given relation fields "%s"`,
					relationFromFieldName,
					relationFields,
				)
			} else {
				relationFromFieldName = key
			}
		} else {
			return 错误类.X创建错误码(
				错误码类.CodeInvalidParameter,
				`parameter relationKV should be format of "ResultFieldName:BindToAttrName"`,
			)
		}
		if relationFromFieldName != "" {
			// 注意，该值可能是切片类型。
			relationDataMap = utils.ListToMapByKey(maps, relationFromFieldName)
		}
		if len(relationDataMap) == 0 {
			return 错误类.X创建错误码并格式化(
				错误码类.CodeInvalidParameter,
				`cannot find the relation data map, maybe invalid relation fields given "%v"`,
				relationFields,
			)
		}
	}
	// 绑定到目标属性。
	var (
		ok              bool
		bindToAttrValue reflect.Value
		bindToAttrKind  reflect.Kind
		bindToAttrType  reflect.Type
		bindToAttrField reflect.StructField
	)
	if arrayItemType.Kind() == reflect.Ptr {
		if bindToAttrField, ok = arrayItemType.Elem().FieldByName(bindToAttrName); !ok {
			return 错误类.X创建错误码并格式化(
				错误码类.CodeInvalidParameter,
				`invalid parameter bindToAttrName: cannot find attribute with name "%s" from slice element`,
				bindToAttrName,
			)
		}
	} else {
		if bindToAttrField, ok = arrayItemType.FieldByName(bindToAttrName); !ok {
			return 错误类.X创建错误码并格式化(
				错误码类.CodeInvalidParameter,
				`invalid parameter bindToAttrName: cannot find attribute with name "%s" from slice element`,
				bindToAttrName,
			)
		}
	}
	bindToAttrType = bindToAttrField.Type
	bindToAttrKind = bindToAttrType.Kind()

	// 绑定到关联条件。
	var (
		relationFromAttrValue          reflect.Value
		relationFromAttrField          reflect.Value
		relationBindToFieldNameChecked bool
	)
	for i := 0; i < arrayValue.Len(); i++ {
		arrayElemValue := arrayValue.Index(i)
		// 应在非指针 reflect.Value 上调用 FieldByName。
		if arrayElemValue.Kind() == reflect.Ptr {
			// Like: []*Entity
			arrayElemValue = arrayElemValue.Elem()
			if !arrayElemValue.IsValid() {
// 如果元素为nil，则创建一个新元素并将其设置到切片中。
// "reflect.New(itemType.Elem())" 用于创建一个新的元素，并返回该元素的地址。
// 例如：
// reflect.New(itemType.Elem())        // => *Entity （返回指向新创建实体类型的指针）
// reflect.New(itemType.Elem()).Elem() // => Entity （获取新创建实体类型的值）
				arrayElemValue = reflect.New(arrayItemType.Elem()).Elem()
				arrayValue.Index(i).Set(arrayElemValue.Addr())
			}
		} else {
			// Like: []Entity
		}
		bindToAttrValue = arrayElemValue.FieldByName(bindToAttrName)
		if relationAttrName != "" {
			// 当前切片元素的属性值。
			relationFromAttrValue = arrayElemValue.FieldByName(relationAttrName)
			if relationFromAttrValue.Kind() == reflect.Ptr {
				relationFromAttrValue = relationFromAttrValue.Elem()
			}
		} else {
			// 当前切片元素。
			relationFromAttrValue = arrayElemValue
		}
		if len(relationDataMap) > 0 && !relationFromAttrValue.IsValid() {
			return 错误类.X创建错误码并格式化(错误码类.CodeInvalidParameter, `invalid relation fields specified: "%v"`, relationFields)
		}
		// 检查并查找可能绑定到属性名称的地方。
		if relationFields != "" && !relationBindToFieldNameChecked {
			relationFromAttrField = relationFromAttrValue.FieldByName(relationBindToFieldName)
			if !relationFromAttrField.IsValid() {
				var (
					fieldMap, _ = gstructs.FieldMap(gstructs.FieldMapInput{
						Pointer:         relationFromAttrValue,
						RecursiveOption: gstructs.RecursiveOptionEmbeddedNoTag,
					})
				)
				if key, _ := utils.MapPossibleItemByKey(X取Map(fieldMap), relationBindToFieldName); key == "" {
					return 错误类.X创建错误码并格式化(
						错误码类.CodeInvalidParameter,
						`cannot find possible related attribute name "%s" from given relation fields "%s"`,
						relationBindToFieldName,
						relationFields,
					)
				} else {
					relationBindToFieldName = key
				}
			}
			relationBindToFieldNameChecked = true
		}
		switch bindToAttrKind {
		case reflect.Array, reflect.Slice:
			if len(relationDataMap) > 0 {
				relationFromAttrField = relationFromAttrValue.FieldByName(relationBindToFieldName)
				if relationFromAttrField.IsValid() {
					// 创建一个Result类型的切片，初始长度为0
// results := make(Result, 0)
					results := make([]interface{}, 0)
					for _, v := range SliceAny别名(relationDataMap[String(relationFromAttrField.Interface())]) {
						item := v
						results = append(results, item)
					}
					if err = Structs(results, bindToAttrValue.Addr()); err != nil {
						return err
					}
				} else {
					// 可能该属性尚不存在。
					return 错误类.X创建错误码并格式化(错误码类.CodeInvalidParameter, `invalid relation fields specified: "%v"`, relationFields)
				}
			} else {
				return 错误类.X创建错误码并格式化(
					错误码类.CodeInvalidParameter,
					`relationKey should not be empty as field "%s" is slice`,
					bindToAttrName,
				)
			}

		case reflect.Ptr:
			var element reflect.Value
			if bindToAttrValue.IsNil() {
				element = reflect.New(bindToAttrType.Elem()).Elem()
			} else {
				element = bindToAttrValue.Elem()
			}
			if len(relationDataMap) > 0 {
				relationFromAttrField = relationFromAttrValue.FieldByName(relationBindToFieldName)
				if relationFromAttrField.IsValid() {
					v := relationDataMap[String(relationFromAttrField.Interface())]
					if v == nil {
						// 没有关联数据。
						continue
					}
					if utils.IsSlice(v) {
						if err = Struct(SliceAny别名(v)[0], element); err != nil {
							return err
						}
					} else {
						if err = Struct(v, element); err != nil {
							return err
						}
					}
				} else {
					// 可能该属性尚不存在。
					return 错误类.X创建错误码并格式化(错误码类.CodeInvalidParameter, `invalid relation fields specified: "%v"`, relationFields)
				}
			} else {
				if i >= len(maps) {
					// 没有关联数据。
					continue
				}
				v := maps[i]
				if v == nil {
					// 没有关联数据。
					continue
				}
				if err = Struct(v, element); err != nil {
					return err
				}
			}
			bindToAttrValue.Set(element.Addr())

		case reflect.Struct:
			if len(relationDataMap) > 0 {
				relationFromAttrField = relationFromAttrValue.FieldByName(relationBindToFieldName)
				if relationFromAttrField.IsValid() {
					relationDataItem := relationDataMap[String(relationFromAttrField.Interface())]
					if relationDataItem == nil {
						// 没有关联数据。
						continue
					}
					if utils.IsSlice(relationDataItem) {
						if err = Struct(SliceAny别名(relationDataItem)[0], bindToAttrValue); err != nil {
							return err
						}
					} else {
						if err = Struct(relationDataItem, bindToAttrValue); err != nil {
							return err
						}
					}
				} else {
					// 可能该属性尚不存在。
					return 错误类.X创建错误码并格式化(错误码类.CodeInvalidParameter, `invalid relation fields specified: "%v"`, relationFields)
				}
			} else {
				if i >= len(maps) {
					// 没有关联数据。
					continue
				}
				relationDataItem := maps[i]
				if relationDataItem == nil {
					// 没有关联数据。
					continue
				}
				if err = Struct(relationDataItem, bindToAttrValue); err != nil {
					return err
				}
			}

		default:
			return 错误类.X创建错误码并格式化(错误码类.CodeInvalidParameter, `unsupported attribute type: %s`, bindToAttrKind.String())
		}
	}
	reflect.ValueOf(structSlicePointer).Elem().Set(arrayValue)
	return nil
}
