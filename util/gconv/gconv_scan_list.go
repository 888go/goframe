// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gconv

import (
	"database/sql"
	"reflect"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/utils"
	"github.com/gogf/gf/v2/os/gstructs"
)

// ScanList将`structSlice`转换为包含其他复杂结构体属性的结构体切片。
// 注意，参数`structSlicePointer`的类型应为*[]struct/*[]*struct。
//
// 使用示例1：普通属性结构关系：
//
//	type EntityUser struct {
//	    Uid  int
//	    Name string
//	}
//
//	type EntityUserDetail struct {
//	    Uid     int
//	    Address string
//	}
//
//	type EntityUserScores struct {
//	    Id     int
//	    Uid    int
//	    Score  int
//	    Course string
//	}
//
//	type Entity struct {
//	    User       *EntityUser
//	    UserDetail *EntityUserDetail
//	    UserScores []*EntityUserScores
//	}
//
//	var users []*Entity
//	var userRecords   = EntityUser{Uid: 1, Name:"john"}
//	var detailRecords = EntityUser{Uid: 1, Address: "chengdu"}
//	var scoresRecords = EntityUser{Id: 1, Uid: 1, Score: 100, Course: "math"}
//	ScanList(userRecords, &users, "User")
//	ScanList(userRecords, &users, "User", "uid")
//	ScanList(detailRecords, &users, "UserDetail", "User", "uid:Uid")
//	ScanList(scoresRecords, &users, "UserScores", "User", "uid:Uid")
//	ScanList(scoresRecords, &users, "UserScores", "User", "uid")
//
// 使用示例2：嵌入属性结构关系：
//
//	type EntityUser struct {
//		   Uid  int
//		   Name string
//	}
//
//	type EntityUserDetail struct {
//		   Uid     int
//		   Address string
//	}
//
//	type EntityUserScores struct {
//		   Id    int
//		   Uid   int
//		   Score int
//	}
//
//	type Entity struct {
//		   EntityUser
//		   UserDetail EntityUserDetail
//		   UserScores []EntityUserScores
//	}
//
//	var userRecords   = EntityUser{Uid: 1, Name:"john"}
//	var detailRecords = EntityUser{Uid: 1, Address: "chengdu"}
//	var scoresRecords = EntityUser{Id: 1, Uid: 1, Score: 100, Course: "math"}
//	ScanList(userRecords, &users)
//	ScanList(detailRecords, &users, "UserDetail", "uid")
//	ScanList(scoresRecords, &users, "UserScores", "uid")
//
// 示例代码中的"User/UserDetail/UserScores"指定当前结果将绑定的目标属性结构。
//
// 示例代码中的"uid"是结果的表字段名，而"Uid"是相关结构体属性名，而不是绑定到目标的属性名。在示例代码中，它是实体"Entity"的"User"属性的"Uid"属性名称。它会根据给定的`relation`参数自动计算HasOne/HasMany关系。
//
// 参考示例或单元测试用例以更清楚地理解此函数的工作原理。 md5:1e63a3d19a1b0060
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

// doScanList 将 `structSlice` 转换为包含嵌套复杂结构体属性的 struct 切片。注意，参数 `structSlicePointer` 的类型应该是 `*[]struct` 或 `*[]*struct`。 md5:37e4682b243c9ef8
func doScanList(
	structSlice interface{}, structSlicePointer interface{}, bindToAttrName, relationAttrName, relationFields string,
) (err error) {
	var (
		maps = Maps(structSlice)
	)
	if len(maps) == 0 {
		return nil
	}
	// 对参数进行必要的检查。 md5:00bddba1a043bfdd
	if bindToAttrName == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, `bindToAttrName should not be empty`)
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
		return gerror.NewCodef(
			gcode.CodeInvalidParameter,
			"structSlicePointer should be type of *[]struct/*[]*struct, but got: %v",
			reflectKind,
		)
	}
	reflectValue = reflectValue.Elem()
	reflectKind = reflectValue.Kind()
	if reflectKind != reflect.Slice && reflectKind != reflect.Array {
		return gerror.NewCodef(
			gcode.CodeInvalidParameter,
			"structSlicePointer should be type of *[]struct/*[]*struct, but got: %v",
			reflectKind,
		)
	}
	length := len(maps)
	if length == 0 {
		// 指向的切片不为空。 md5:1348d4b6d686b8f3
		if reflectValue.Len() > 0 {
			// 这里检查是否具有已初始化的结构体项。
			// 然后返回错误以警告开发者其为空且无法进行转换。 md5:cd5f133a393c1157
			if v := reflectValue.Index(0); v.Kind() != reflect.Ptr {
				return sql.ErrNoRows
			}
		}
		// 对于空的结构体切片，什么也不做。 md5:f65a6d24cd42ca62
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

	// Slice element item.
	arrayItemType = arrayValue.Index(0).Type()

	// Relation variables.
	var (
		relationDataMap         map[string]interface{}
		relationFromFieldName   string // 例如：relationKV：id：uid -> id. md5:3732472417ccbf22
		relationBindToFieldName string // 例如：relationKV：id：uid -> uid. md5:dda263df86dc03a1
	)
	if len(relationFields) > 0 {
		// 表字段名与属性名之间的关联键字符串
		// 可以使用字符'='或':'进行连接。 md5:a3dd08343df8a7ac
		array := utils.SplitAndTrim(relationFields, "=")
		if len(array) == 1 {
			// 与旧的分隔字符':'兼容。 md5:21a764d3ea1e081b
			array = utils.SplitAndTrim(relationFields, ":")
		}
		if len(array) == 1 {
			// 关系名称是相同的。 md5:1075b6495b26357b
			array = []string{relationFields, relationFields}
		}
		if len(array) == 2 {
			// 定义表字段到关系属性名。
			// 例如：
			// uid:Uid
			// uid:UserId md5:029253159bee75d1
			relationFromFieldName = array[0]
			relationBindToFieldName = array[1]
			if key, _ := utils.MapPossibleItemByKey(maps[0], relationFromFieldName); key == "" {
				return gerror.NewCodef(
					gcode.CodeInvalidParameter,
					`cannot find possible related table field name "%s" from given relation fields "%s"`,
					relationFromFieldName,
					relationFields,
				)
			} else {
				relationFromFieldName = key
			}
		} else {
			return gerror.NewCode(
				gcode.CodeInvalidParameter,
				`parameter relationKV should be format of "ResultFieldName:BindToAttrName"`,
			)
		}
		if relationFromFieldName != "" {
			// 请注意，该值可能是切片类型。 md5:079de568e97881a6
			relationDataMap = utils.ListToMapByKey(maps, relationFromFieldName)
		}
		if len(relationDataMap) == 0 {
			return gerror.NewCodef(
				gcode.CodeInvalidParameter,
				`cannot find the relation data map, maybe invalid relation fields given "%v"`,
				relationFields,
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
		if bindToAttrField, ok = arrayItemType.Elem().FieldByName(bindToAttrName); !ok {
			return gerror.NewCodef(
				gcode.CodeInvalidParameter,
				`invalid parameter bindToAttrName: cannot find attribute with name "%s" from slice element`,
				bindToAttrName,
			)
		}
	} else {
		if bindToAttrField, ok = arrayItemType.FieldByName(bindToAttrName); !ok {
			return gerror.NewCodef(
				gcode.CodeInvalidParameter,
				`invalid parameter bindToAttrName: cannot find attribute with name "%s" from slice element`,
				bindToAttrName,
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
				// reflect.New(itemType.Elem()).Elem() => 实体 md5:0897d7c0e7467f9d
				arrayElemValue = reflect.New(arrayItemType.Elem()).Elem()
				arrayValue.Index(i).Set(arrayElemValue.Addr())
			}
		} else {
			// Like: []Entity
		}
		bindToAttrValue = arrayElemValue.FieldByName(bindToAttrName)
		if relationAttrName != "" {
			// 当前切片元素的属性值。 md5:b46440a93bb1ddaa
			relationFromAttrValue = arrayElemValue.FieldByName(relationAttrName)
			if relationFromAttrValue.Kind() == reflect.Ptr {
				relationFromAttrValue = relationFromAttrValue.Elem()
			}
		} else {
			// Current slice element.
			relationFromAttrValue = arrayElemValue
		}
		if len(relationDataMap) > 0 && !relationFromAttrValue.IsValid() {
			return gerror.NewCodef(gcode.CodeInvalidParameter, `invalid relation fields specified: "%v"`, relationFields)
		}
		// 检查并尝试找到可能与属性名绑定的位置。 md5:b1e1f2121b3b5f92
		if relationFields != "" && !relationBindToFieldNameChecked {
			relationFromAttrField = relationFromAttrValue.FieldByName(relationBindToFieldName)
			if !relationFromAttrField.IsValid() {
				var (
					fieldMap, _ = gstructs.FieldMap(gstructs.FieldMapInput{
						Pointer:         relationFromAttrValue,
						RecursiveOption: gstructs.RecursiveOptionEmbeddedNoTag,
					})
				)
				if key, _ := utils.MapPossibleItemByKey(Map(fieldMap), relationBindToFieldName); key == "" {
					return gerror.NewCodef(
						gcode.CodeInvalidParameter,
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
					// 将results初始化为一个长度为0的Result切片。 md5:676fe93ddada53e4
					results := make([]interface{}, 0)
					for _, v := range SliceAny(relationDataMap[String(relationFromAttrField.Interface())]) {
						item := v
						results = append(results, item)
					}
					if err = Structs(results, bindToAttrValue.Addr()); err != nil {
						return err
					}
				} else {
					// 可能属性还不存在。 md5:d7992076e8a1e5fe
					return gerror.NewCodef(gcode.CodeInvalidParameter, `invalid relation fields specified: "%v"`, relationFields)
				}
			} else {
				return gerror.NewCodef(
					gcode.CodeInvalidParameter,
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
						// 没有关联数据。 md5:4f76ca1525fb5005
						continue
					}
					if utils.IsSlice(v) {
						if err = Struct(SliceAny(v)[0], element); err != nil {
							return err
						}
					} else {
						if err = Struct(v, element); err != nil {
							return err
						}
					}
				} else {
					// 可能属性还不存在。 md5:d7992076e8a1e5fe
					return gerror.NewCodef(gcode.CodeInvalidParameter, `invalid relation fields specified: "%v"`, relationFields)
				}
			} else {
				if i >= len(maps) {
					// 没有关联数据。 md5:4f76ca1525fb5005
					continue
				}
				v := maps[i]
				if v == nil {
					// 没有关联数据。 md5:4f76ca1525fb5005
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
						// 没有关联数据。 md5:4f76ca1525fb5005
						continue
					}
					if utils.IsSlice(relationDataItem) {
						if err = Struct(SliceAny(relationDataItem)[0], bindToAttrValue); err != nil {
							return err
						}
					} else {
						if err = Struct(relationDataItem, bindToAttrValue); err != nil {
							return err
						}
					}
				} else {
					// 可能属性还不存在。 md5:d7992076e8a1e5fe
					return gerror.NewCodef(gcode.CodeInvalidParameter, `invalid relation fields specified: "%v"`, relationFields)
				}
			} else {
				if i >= len(maps) {
					// 没有关联数据。 md5:4f76ca1525fb5005
					continue
				}
				relationDataItem := maps[i]
				if relationDataItem == nil {
					// 没有关联数据。 md5:4f76ca1525fb5005
					continue
				}
				if err = Struct(relationDataItem, bindToAttrValue); err != nil {
					return err
				}
			}

		default:
			return gerror.NewCodef(gcode.CodeInvalidParameter, `unsupported attribute type: %s`, bindToAttrKind.String())
		}
	}
	reflect.ValueOf(structSlicePointer).Elem().Set(arrayValue)
	return nil
}
