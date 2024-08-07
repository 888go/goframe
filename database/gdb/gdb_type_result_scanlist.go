// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

import (
	"database/sql"
	"reflect"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/os/gstructs"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
	gutil "github.com/888go/goframe/util/gutil"
)

// X取指针列表 将 `r` 转换为包含其他复杂结构体属性的结构体切片。
// 注意，参数 `structSlicePointer` 应该是 *[]struct/*[]*struct 类型。
//
// 使用示例 1：普通属性结构体关系：
//
//	type EntityUser struct {
//		Uid  int
//		Name string
//	}
//
//	type EntityUserDetail struct {
//		Uid     int
//		Address string
//	}
//
//	type EntityUserScores struct {
//		Id     int
//		Uid    int
//		Score  int
//		Course string
//	}
//
//	type Entity struct {
//		User       *EntityUser
//		UserDetail *EntityUserDetail
//		UserScores []*EntityUserScores
//	}
//
//	var users []*Entity
//	X取指针列表(&users, "User")
//	X取指针列表(&users, "User", "uid")
//	X取指针列表(&users, "UserDetail", "User", "uid:Uid")
//	X取指针列表(&users, "UserScores", "User", "uid:Uid")
//	X取指针列表(&users, "UserScores", "User", "uid")
//
// 使用示例 2：嵌入属性结构体关系：
//
//	type EntityUser struct {
//		Uid  int
//		Name string
//	}
//
//	type EntityUserDetail struct {
//		Uid     int
//		Address string
//	}
//
//	type EntityUserScores struct {
//		Id    int
//		Uid   int
//		Score int
//	}
//
//	type Entity struct {
//		EntityUser
//		UserDetail EntityUserDetail
//		UserScores []EntityUserScores
//	}
//
//	var users []*Entity
//	X取指针列表(&users)
//	X取指针列表(&users, "UserDetail", "uid")
//	X取指针列表(&users, "UserScores", "uid")
//
// 示例代码中的 "User/UserDetail/UserScores" 参数指定了当前结果将绑定的目标属性结构体。
//
// 示例代码中的 "uid" 是结果表字段名，而 "Uid" 是相关结构体属性名，而不是绑定目标的属性名。
// 在示例代码中，它是 "Entity" 实体的 "User" 的属性名 "Uid"。它会根据给定的 `relation` 参数自动计算 HasOne/HasMany 关系。
//
// 可参考示例或单元测试用例以更清晰地理解此函数的工作方式。
// md5:d6997acc67d472c4
func (r Result) X取指针列表(结构体切片指针 interface{}, 绑定到结构体属性名称 string, 结构体属性关联 ...string) (错误 error) {
	out, 错误 := checkGetSliceElementInfoForScanList(结构体切片指针, 绑定到结构体属性名称)
	if 错误 != nil {
		return 错误
	}

	var (
		relationAttrName string
		relationFields   string
	)
	switch len(结构体属性关联) {
	case 2:
		relationAttrName = 结构体属性关联[0]
		relationFields = 结构体属性关联[1]
	case 1:
		relationFields = 结构体属性关联[0]
	}
	return doScanList(doScanListInput{
		Model:              nil,
		Result:             r,
		StructSlicePointer: 结构体切片指针,
		StructSliceValue:   out.SliceReflectValue,
		BindToAttrName:     绑定到结构体属性名称,
		RelationAttrName:   relationAttrName,
		RelationFields:     relationFields,
	})
}

type checkGetSliceElementInfoForScanListOutput struct {
	SliceReflectValue reflect.Value
	BindToAttrType    reflect.Type
}

func checkGetSliceElementInfoForScanList(structSlicePointer interface{}, bindToAttrName string) (out *checkGetSliceElementInfoForScanListOutput, err error) {
		// 对参数进行必要的检查。 md5:00bddba1a043bfdd
	if structSlicePointer == nil {
		return nil, gerror.X创建错误码(gcode.CodeInvalidParameter, `structSlicePointer cannot be nil`)
	}
	if bindToAttrName == "" {
		return nil, gerror.X创建错误码(gcode.CodeInvalidParameter, `bindToAttrName should not be empty`)
	}
	var (
		reflectType  reflect.Type
		reflectValue = reflect.ValueOf(structSlicePointer)
		reflectKind  = reflectValue.Kind()
	)
	if reflectKind == reflect.Interface {
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Kind()
	}
	if reflectKind != reflect.Ptr {
		return nil, gerror.X创建错误码并格式化(
			gcode.CodeInvalidParameter,
			"structSlicePointer should be type of *[]struct/*[]*struct, but got: %s",
			reflect.TypeOf(structSlicePointer).String(),
		)
	}
	out = &checkGetSliceElementInfoForScanListOutput{
		SliceReflectValue: reflectValue.Elem(),
	}
		// 找到切片的元素结构类型。 md5:a55c378f6fa3b326
	reflectType = reflectValue.Type().Elem().Elem()
	reflectKind = reflectType.Kind()
	for reflectKind == reflect.Ptr {
		reflectType = reflectType.Elem()
		reflectKind = reflectType.Kind()
	}
	if reflectKind != reflect.Struct {
		err = gerror.X创建错误码并格式化(
			gcode.CodeInvalidParameter,
			"structSlicePointer should be type of *[]struct/*[]*struct, but got: %s",
			reflect.TypeOf(structSlicePointer).String(),
		)
		return
	}
		// 通过给定的名称查找目标字段。 md5:8fe292d32e17dba0
	structField, ok := reflectType.FieldByName(bindToAttrName)
	if !ok {
		return nil, gerror.X创建错误码并格式化(
			gcode.CodeInvalidParameter,
			`field "%s" not found in element of "%s"`,
			bindToAttrName,
			reflect.TypeOf(structSlicePointer).String(),
		)
	}
		// 查找用于ORM字段过滤的属性结构体类型。 md5:1b98a4f65808a146
	reflectType = structField.Type
	reflectKind = reflectType.Kind()
	for reflectKind == reflect.Ptr {
		reflectType = reflectType.Elem()
		reflectKind = reflectType.Kind()
	}
	if reflectKind == reflect.Slice || reflectKind == reflect.Array {
		reflectType = reflectType.Elem()
		reflectKind = reflectType.Kind()
	}
	out.BindToAttrType = reflectType
	return
}

type doScanListInput struct {
	Model              *Model
	Result             Result
	StructSlicePointer interface{}
	StructSliceValue   reflect.Value
	BindToAttrName     string
	RelationAttrName   string
	RelationFields     string
}

// doScanList 将 `result` 转换为包含嵌套复杂结构体属性的切片。参数 `model` 用于递归扫描，即它可以递归地扫描结构体/结构体的属性，但需要数据库访问模型。
// 注意参数 `structSlicePointer` 应该是 *[]struct 或 *[]*struct 类型。
// md5:b32c3ddd7d2b8656
func doScanList(in doScanListInput) (err error) {
	if in.Result.X是否为空() {
		return nil
	}
	if in.BindToAttrName == "" {
		return gerror.X创建错误码(gcode.CodeInvalidParameter, `bindToAttrName should not be empty`)
	}

	length := len(in.Result)
	if length == 0 {
				// 指向的切片不为空。 md5:1348d4b6d686b8f3
		if in.StructSliceValue.Len() > 0 {
			// 这里检查是否具有已初始化的结构体项。
			// 然后返回错误以警告开发者其为空且无法进行转换。
			// md5:cd5f133a393c1157
			if v := in.StructSliceValue.Index(0); v.Kind() != reflect.Ptr {
				return sql.ErrNoRows
			}
		}
				// 对于空的结构体切片，什么也不做。 md5:f65a6d24cd42ca62
		return nil
	}
	var (
		arrayValue    reflect.Value // Like: []*Entity
		arrayItemType reflect.Type  // Like: *Entity
		reflectType   = reflect.TypeOf(in.StructSlicePointer)
	)
	if in.StructSliceValue.Len() > 0 {
		arrayValue = in.StructSliceValue
	} else {
		arrayValue = reflect.MakeSlice(reflectType.Elem(), length, length)
	}

	// Slice element item.
	arrayItemType = arrayValue.Index(0).Type()

	// Relation variables.
	var (
		relationDataMap         map[string]Value
		relationFromFieldName   string // 例如：relationKV：id：uid -> id. md5:3732472417ccbf22
		relationBindToFieldName string // 例如：relationKV：id：uid -> uid. md5:dda263df86dc03a1
	)
	if len(in.RelationFields) > 0 {
		// 表字段名与属性名之间的关联键字符串
		// 可以使用字符'='或':'进行连接。
		// md5:a3dd08343df8a7ac
		array := gstr.X分割并忽略空值(in.RelationFields, "=")
		if len(array) == 1 {
						// 与旧的分隔字符':'兼容。 md5:21a764d3ea1e081b
			array = gstr.X分割并忽略空值(in.RelationFields, ":")
		}
		if len(array) == 1 {
						// 关系名称是相同的。 md5:1075b6495b26357b
			array = []string{in.RelationFields, in.RelationFields}
		}
		if len(array) == 2 {
			// 定义表字段到关系属性名。
			// 例如：
			// uid:Uid
			// uid:UserId
			// md5:029253159bee75d1
			relationFromFieldName = array[0]
			relationBindToFieldName = array[1]
			if key, _ := gutil.MapPossibleItemByKey(in.Result[0].X取Map(), relationFromFieldName); key == "" {
				return gerror.X创建错误码并格式化(
					gcode.CodeInvalidParameter,
					`cannot find possible related table field name "%s" from given relation fields "%s"`,
					relationFromFieldName,
					in.RelationFields,
				)
			} else {
				relationFromFieldName = key
			}
		} else {
			return gerror.X创建错误码(
				gcode.CodeInvalidParameter,
				`parameter relationKV should be format of "ResultFieldName:BindToAttrName"`,
			)
		}
		if relationFromFieldName != "" {
						// 请注意，该值可能是切片类型。 md5:079de568e97881a6
			relationDataMap = in.Result.X取字段Map泛型类(relationFromFieldName)
		}
		if len(relationDataMap) == 0 {
			return gerror.X创建错误码并格式化(
				gcode.CodeInvalidParameter,
				`cannot find the relation data map, maybe invalid relation fields given "%v"`,
				in.RelationFields,
			)
		}
	}
		// 将其绑定到目标属性。 md5:6248a034de9b08e4
	var (
		ok              bool
		bindToAttrValue reflect.Value
		bindToAttrKind  reflect.Kind
		bindToAttrType  reflect.Type
		bindToAttrField reflect.StructField
	)
	if arrayItemType.Kind() == reflect.Ptr {
		if bindToAttrField, ok = arrayItemType.Elem().FieldByName(in.BindToAttrName); !ok {
			return gerror.X创建错误码并格式化(
				gcode.CodeInvalidParameter,
				`invalid parameter bindToAttrName: cannot find attribute with name "%s" from slice element`,
				in.BindToAttrName,
			)
		}
	} else {
		if bindToAttrField, ok = arrayItemType.FieldByName(in.BindToAttrName); !ok {
			return gerror.X创建错误码并格式化(
				gcode.CodeInvalidParameter,
				`invalid parameter bindToAttrName: cannot find attribute with name "%s" from slice element`,
				in.BindToAttrName,
			)
		}
	}
	bindToAttrType = bindToAttrField.Type
	bindToAttrKind = bindToAttrType.Kind()

		// 绑定关系条件。 md5:1d13e1ebe0b47bd2
	var (
		relationFromAttrValue          reflect.Value
		relationFromAttrField          reflect.Value
		relationBindToFieldNameChecked bool
	)
	for i := 0; i < arrayValue.Len(); i++ {
		arrayElemValue := arrayValue.Index(i)
				// 应该在非指针的reflect.Value上调用FieldByName。 md5:1343ff0ec0419e1f
		if arrayElemValue.Kind() == reflect.Ptr {
			// Like: []*Entity
			arrayElemValue = arrayElemValue.Elem()
			if !arrayElemValue.IsValid() {
				// 如果元素为nil，则创建一个并将其设置到切片中。
				// "reflect.New(itemType.Elem())" 用于创建一个新的元素，并返回该元素的地址。
				// 例如：
				// reflect.New(itemType.Elem())        => *实体
				// reflect.New(itemType.Elem()).Elem() => 实体
				// md5:0897d7c0e7467f9d
				arrayElemValue = reflect.New(arrayItemType.Elem()).Elem()
				arrayValue.Index(i).Set(arrayElemValue.Addr())
			}
		} else {
			// Like: []Entity
		}
		bindToAttrValue = arrayElemValue.FieldByName(in.BindToAttrName)
		if in.RelationAttrName != "" {
						// 当前切片元素的属性值。 md5:b46440a93bb1ddaa
			relationFromAttrValue = arrayElemValue.FieldByName(in.RelationAttrName)
			if relationFromAttrValue.Kind() == reflect.Ptr {
				relationFromAttrValue = relationFromAttrValue.Elem()
			}
		} else {
			// Current slice element.
			relationFromAttrValue = arrayElemValue
		}
		if len(relationDataMap) > 0 && !relationFromAttrValue.IsValid() {
			return gerror.X创建错误码并格式化(gcode.CodeInvalidParameter, `invalid relation fields specified: "%v"`, in.RelationFields)
		}
				// 检查并尝试找到可能与属性名绑定的位置。 md5:b1e1f2121b3b5f92
		if in.RelationFields != "" && !relationBindToFieldNameChecked {
			relationFromAttrField = relationFromAttrValue.FieldByName(relationBindToFieldName)
			if !relationFromAttrField.IsValid() {
				fieldMap, _ := gstructs.FieldMap(gstructs.FieldMapInput{
					Pointer:         relationFromAttrValue,
					RecursiveOption: gstructs.RecursiveOptionEmbeddedNoTag,
				})
				if key, _ := gutil.MapPossibleItemByKey(gconv.X取Map(fieldMap), relationBindToFieldName); key == "" {
					return gerror.X创建错误码并格式化(
						gcode.CodeInvalidParameter,
						`cannot find possible related attribute name "%s" from given relation fields "%s"`,
						relationBindToFieldName,
						in.RelationFields,
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
					results := make(Result, 0)
					for _, v := range relationDataMap[gconv.String(relationFromAttrField.Interface())].Slice别名() {
						results = append(results, v.(Record))
					}
					if err = results.X取切片结构体指针(bindToAttrValue.Addr()); err != nil {
						return err
					}
					// Recursively Scan.
					if in.Model != nil {
						if err = in.Model.doWithScanStructs(bindToAttrValue.Addr()); err != nil {
							return nil
						}
					}
				} else {
										// 可能属性还不存在。 md5:d7992076e8a1e5fe
					return gerror.X创建错误码并格式化(gcode.CodeInvalidParameter, `invalid relation fields specified: "%v"`, in.RelationFields)
				}
			} else {
				return gerror.X创建错误码并格式化(
					gcode.CodeInvalidParameter,
					`relationKey should not be empty as field "%s" is slice`,
					in.BindToAttrName,
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
					v := relationDataMap[gconv.String(relationFromAttrField.Interface())]
					if v == nil {
											// 没有关联数据。 md5:4f76ca1525fb5005
						continue
					}
					if v.X是否为切片() {
						if err = v.Slice别名()[0].(Record).X取结构体指针(element); err != nil {
							return err
						}
					} else {
						if err = v.X取值().(Record).X取结构体指针(element); err != nil {
							return err
						}
					}
				} else {
										// 可能属性还不存在。 md5:d7992076e8a1e5fe
					return gerror.X创建错误码并格式化(gcode.CodeInvalidParameter, `invalid relation fields specified: "%v"`, in.RelationFields)
				}
			} else {
				if i >= len(in.Result) {
										// 没有关联数据。 md5:4f76ca1525fb5005
					continue
				}
				v := in.Result[i]
				if v == nil {
										// 没有关联数据。 md5:4f76ca1525fb5005
					continue
				}
				if err = v.X取结构体指针(element); err != nil {
					return err
				}
			}
			// Recursively Scan.
			if in.Model != nil {
				if err = in.Model.doWithScanStruct(element); err != nil {
					return err
				}
			}
			bindToAttrValue.Set(element.Addr())

		case reflect.Struct:
			if len(relationDataMap) > 0 {
				relationFromAttrField = relationFromAttrValue.FieldByName(relationBindToFieldName)
				if relationFromAttrField.IsValid() {
					relationDataItem := relationDataMap[gconv.String(relationFromAttrField.Interface())]
					if relationDataItem == nil {
											// 没有关联数据。 md5:4f76ca1525fb5005
						continue
					}
					if relationDataItem.X是否为切片() {
						if err = relationDataItem.Slice别名()[0].(Record).X取结构体指针(bindToAttrValue); err != nil {
							return err
						}
					} else {
						if err = relationDataItem.X取值().(Record).X取结构体指针(bindToAttrValue); err != nil {
							return err
						}
					}
				} else {
										// 可能属性还不存在。 md5:d7992076e8a1e5fe
					return gerror.X创建错误码并格式化(gcode.CodeInvalidParameter, `invalid relation fields specified: "%v"`, in.RelationFields)
				}
			} else {
				if i >= len(in.Result) {
										// 没有关联数据。 md5:4f76ca1525fb5005
					continue
				}
				relationDataItem := in.Result[i]
				if relationDataItem == nil {
										// 没有关联数据。 md5:4f76ca1525fb5005
					continue
				}
				if err = relationDataItem.X取结构体指针(bindToAttrValue); err != nil {
					return err
				}
			}
			// Recursively Scan.
			if in.Model != nil {
				if err = in.Model.doWithScanStruct(bindToAttrValue); err != nil {
					return err
				}
			}

		default:
			return gerror.X创建错误码并格式化(gcode.CodeInvalidParameter, `unsupported attribute type: %s`, bindToAttrKind.String())
		}
	}
	reflect.ValueOf(in.StructSlicePointer).Elem().Set(arrayValue)
	return nil
}
