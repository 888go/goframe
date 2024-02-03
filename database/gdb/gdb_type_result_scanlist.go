// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb

import (
	"database/sql"
	"reflect"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/os/gstructs"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
)

// ScanList 将 `r` 转换为包含其他复杂结构体属性的结构体切片。
// 注意，参数 `structSlicePointer` 应该是 *[]struct 或 *[]*struct 类型。
// 参考关联模型: https://goframe.org/pages/viewpage.action?pageId=1114326
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
//	ScanList(&users, "User")
//	ScanList(&users, "User", "uid")
//	ScanList(&users, "UserDetail", "User", "uid:Uid")
//	ScanList(&users, "UserScores", "User", "uid:Uid")
//	ScanList(&users, "UserScores", "User", "uid")
//
// 使用示例 2：嵌入式属性结构体关系：
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
//	ScanList(&users)
//	ScanList(&users, "UserDetail", "uid")
//	ScanList(&users, "UserScores", "uid")
//
// 示例代码中的 "User/UserDetail/UserScores" 参数指定了当前结果将绑定的目标属性结构体。
//
// 示例代码中的 "uid" 是结果表字段名，而 "Uid" 是关联结构体属性名——而不是目标绑定的属性名。在示例中，它是实体 "Entity" 中 "User" 的属性名 "Uid"。它会根据给定的 `relation` 参数自动计算 HasOne/HasMany 关系。
//
// 为了清晰理解此函数，请参考示例或单元测试用例。
func (r Result) ScanList(structSlicePointer interface{}, bindToAttrName string, relationAttrNameAndFields ...string) (err error) {
	out, err := checkGetSliceElementInfoForScanList(structSlicePointer, bindToAttrName)
	if err != nil {
		return err
	}

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
	return doScanList(doScanListInput{
		Model:              nil,
		Result:             r,
		StructSlicePointer: structSlicePointer,
		StructSliceValue:   out.SliceReflectValue,
		BindToAttrName:     bindToAttrName,
		RelationAttrName:   relationAttrName,
		RelationFields:     relationFields,
	})
}

type checkGetSliceElementInfoForScanListOutput struct {
	SliceReflectValue reflect.Value
	BindToAttrType    reflect.Type
}

func checkGetSliceElementInfoForScanList(structSlicePointer interface{}, bindToAttrName string) (out *checkGetSliceElementInfoForScanListOutput, err error) {
	// 对参数进行必要的检查。
	if structSlicePointer == nil {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, `structSlicePointer cannot be nil`)
	}
	if bindToAttrName == "" {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, `bindToAttrName should not be empty`)
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
		return nil, gerror.NewCodef(
			gcode.CodeInvalidParameter,
			"structSlicePointer should be type of *[]struct/*[]*struct, but got: %s",
			reflect.TypeOf(structSlicePointer).String(),
		)
	}
	out = &checkGetSliceElementInfoForScanListOutput{
		SliceReflectValue: reflectValue.Elem(),
	}
	// 查找切片的元素结构体类型。
	reflectType = reflectValue.Type().Elem().Elem()
	reflectKind = reflectType.Kind()
	for reflectKind == reflect.Ptr {
		reflectType = reflectType.Elem()
		reflectKind = reflectType.Kind()
	}
	if reflectKind != reflect.Struct {
		err = gerror.NewCodef(
			gcode.CodeInvalidParameter,
			"structSlicePointer should be type of *[]struct/*[]*struct, but got: %s",
			reflect.TypeOf(structSlicePointer).String(),
		)
		return
	}
	// 根据给定名称查找目标字段。
	structField, ok := reflectType.FieldByName(bindToAttrName)
	if !ok {
		return nil, gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`field "%s" not found in element of "%s"`,
			bindToAttrName,
			reflect.TypeOf(structSlicePointer).String(),
		)
	}
	// 查找用于ORM字段过滤的属性结构体类型。
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

// doScanList 将 `result` 转换为包含其他复杂结构体属性的结构体切片，并且这个转换过程是递归的。
// 参数 `model` 用于递归扫描目的，这意味着它可以递归地扫描结构体/结构体数组，
// 但需要 Model 以便进行数据库访问。
// 注意参数 `structSlicePointer` 的类型应为 *[]struct 或 *[]*struct。
func doScanList(in doScanListInput) (err error) {
	if in.Result.IsEmpty() {
		return nil
	}
	if in.BindToAttrName == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, `bindToAttrName should not be empty`)
	}

	length := len(in.Result)
	if length == 0 {
		// 指向的切片不为空。
		if in.StructSliceValue.Len() > 0 {
// 在这里检查它是否包含已初始化的结构体项。
// 如果为空且无法进行转换，则返回错误以警告开发者。
			if v := in.StructSliceValue.Index(0); v.Kind() != reflect.Ptr {
				return sql.ErrNoRows
			}
		}
		// 对于空的结构体切片，不执行任何操作。
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

	// 切片元素项。
	arrayItemType = arrayValue.Index(0).Type()

	// 关系变量。
	var (
		relationDataMap         map[string]Value
		relationFromFieldName   string // 示例：relationKV: id:uid -> id
// 这个注释表明了一个键值对的示例，其中关系（relationKV）的键是"id:uid"，对应的值为"id"。在实际应用中，这可能表示一个映射关系，通过用户ID(uid)可以找到对应的ID(id)。
		relationBindToFieldName string // 示例：relationKV: id:uid  -> uid
// （注释翻译：这个字段或者变量表示一种键值对关系，其中键是"id:uid"，值是"uid"）
	)
	if len(in.RelationFields) > 0 {
// 表字段名与属性名之间的关联键字符串，可以使用字符'='或':'连接。
		array := gstr.SplitAndTrim(in.RelationFields, "=")
		if len(array) == 1 {
			// 与旧的分隔符 ':' 兼容。
			array = gstr.SplitAndTrim(in.RelationFields, ":")
		}
		if len(array) == 1 {
			// 关系名称是相同的。
			array = []string{in.RelationFields, in.RelationFields}
		}
		if len(array) == 2 {
// 定义表格字段到关联属性名称的映射。
// 例如：
// uid:Uid
// uid:UserId
// 这段代码的作用是将数据库表中的字段名（如uid）映射到程序中使用的关联属性名称，以实现字段名称在代码逻辑中的语义化表达。
			relationFromFieldName = array[0]
			relationBindToFieldName = array[1]
			if key, _ := gutil.MapPossibleItemByKey(in.Result[0].Map(), relationFromFieldName); key == "" {
				return gerror.NewCodef(
					gcode.CodeInvalidParameter,
					`cannot find possible related table field name "%s" from given relation fields "%s"`,
					relationFromFieldName,
					in.RelationFields,
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
			// 注意，该值可能是切片类型。
			relationDataMap = in.Result.MapKeyValue(relationFromFieldName)
		}
		if len(relationDataMap) == 0 {
			return gerror.NewCodef(
				gcode.CodeInvalidParameter,
				`cannot find the relation data map, maybe invalid relation fields given "%v"`,
				in.RelationFields,
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
		if bindToAttrField, ok = arrayItemType.Elem().FieldByName(in.BindToAttrName); !ok {
			return gerror.NewCodef(
				gcode.CodeInvalidParameter,
				`invalid parameter bindToAttrName: cannot find attribute with name "%s" from slice element`,
				in.BindToAttrName,
			)
		}
	} else {
		if bindToAttrField, ok = arrayItemType.FieldByName(in.BindToAttrName); !ok {
			return gerror.NewCodef(
				gcode.CodeInvalidParameter,
				`invalid parameter bindToAttrName: cannot find attribute with name "%s" from slice element`,
				in.BindToAttrName,
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
		bindToAttrValue = arrayElemValue.FieldByName(in.BindToAttrName)
		if in.RelationAttrName != "" {
			// 当前切片元素的属性值。
			relationFromAttrValue = arrayElemValue.FieldByName(in.RelationAttrName)
			if relationFromAttrValue.Kind() == reflect.Ptr {
				relationFromAttrValue = relationFromAttrValue.Elem()
			}
		} else {
			// 当前切片元素。
			relationFromAttrValue = arrayElemValue
		}
		if len(relationDataMap) > 0 && !relationFromAttrValue.IsValid() {
			return gerror.NewCodef(gcode.CodeInvalidParameter, `invalid relation fields specified: "%v"`, in.RelationFields)
		}
		// 检查并查找可能绑定到属性名称的地方。
		if in.RelationFields != "" && !relationBindToFieldNameChecked {
			relationFromAttrField = relationFromAttrValue.FieldByName(relationBindToFieldName)
			if !relationFromAttrField.IsValid() {
				fieldMap, _ := gstructs.FieldMap(gstructs.FieldMapInput{
					Pointer:         relationFromAttrValue,
					RecursiveOption: gstructs.RecursiveOptionEmbeddedNoTag,
				})
				if key, _ := gutil.MapPossibleItemByKey(gconv.Map(fieldMap), relationBindToFieldName); key == "" {
					return gerror.NewCodef(
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
					for _, v := range relationDataMap[gconv.String(relationFromAttrField.Interface())].Slice() {
						results = append(results, v.(Record))
					}
					if err = results.Structs(bindToAttrValue.Addr()); err != nil {
						return err
					}
					// Recursively Scan.
					if in.Model != nil {
						if err = in.Model.doWithScanStructs(bindToAttrValue.Addr()); err != nil {
							return nil
						}
					}
				} else {
					// 可能该属性尚不存在。
					return gerror.NewCodef(gcode.CodeInvalidParameter, `invalid relation fields specified: "%v"`, in.RelationFields)
				}
			} else {
				return gerror.NewCodef(
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
						// 没有关联数据。
						continue
					}
					if v.IsSlice() {
						if err = v.Slice()[0].(Record).Struct(element); err != nil {
							return err
						}
					} else {
						if err = v.Val().(Record).Struct(element); err != nil {
							return err
						}
					}
				} else {
					// 可能该属性尚不存在。
					return gerror.NewCodef(gcode.CodeInvalidParameter, `invalid relation fields specified: "%v"`, in.RelationFields)
				}
			} else {
				if i >= len(in.Result) {
					// 没有关联数据。
					continue
				}
				v := in.Result[i]
				if v == nil {
					// 没有关联数据。
					continue
				}
				if err = v.Struct(element); err != nil {
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
						// 没有关联数据。
						continue
					}
					if relationDataItem.IsSlice() {
						if err = relationDataItem.Slice()[0].(Record).Struct(bindToAttrValue); err != nil {
							return err
						}
					} else {
						if err = relationDataItem.Val().(Record).Struct(bindToAttrValue); err != nil {
							return err
						}
					}
				} else {
					// 可能该属性尚不存在。
					return gerror.NewCodef(gcode.CodeInvalidParameter, `invalid relation fields specified: "%v"`, in.RelationFields)
				}
			} else {
				if i >= len(in.Result) {
					// 没有关联数据。
					continue
				}
				relationDataItem := in.Result[i]
				if relationDataItem == nil {
					// 没有关联数据。
					continue
				}
				if err = relationDataItem.Struct(bindToAttrValue); err != nil {
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
			return gerror.NewCodef(gcode.CodeInvalidParameter, `unsupported attribute type: %s`, bindToAttrKind.String())
		}
	}
	reflect.ValueOf(in.StructSlicePointer).Elem().Set(arrayValue)
	return nil
}
