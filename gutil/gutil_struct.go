// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 工具类

import (
	"reflect"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gstructs"
	"github.com/gogf/gf/v2/util/gconv"
)

// StructToSlice 将结构体转换为键值对构成的切片。
// 例如：{"K1": "v1", "K2": "v2"} => ["K1", "v1", "K2", "v2"]
func X结构体到切片(结构体指针 interface{}) []interface{} {
	var (
		reflectValue = reflect.ValueOf(结构体指针)
		reflectKind  = reflectValue.Kind()
	)
	for reflectKind == reflect.Ptr {
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Kind()
	}
	switch reflectKind {
	case reflect.Struct:
		array := make([]interface{}, 0)
		// 注意，如果在结构体属性中找到gconv标签，则它使用gconv标签名称而非属性名称。
		for k, v := range gconv.Map(reflectValue) {
			array = append(array, k)
			array = append(array, v)
		}
		return array
	}
	return nil
}

// FillStructWithDefault 用 `default/d` 标签的值填充指针指向的结构体属性。
// 参数 `structPtr` 应为 *struct 或 []*struct 类型。
func FillStructWithDefault(structPtr interface{}) error {
	var (
		reflectValue reflect.Value
	)
	if rv, ok := structPtr.(reflect.Value); ok {
		reflectValue = rv
	} else {
		reflectValue = reflect.ValueOf(structPtr)
	}
	switch reflectValue.Kind() {
	case reflect.Ptr:
		// Nothing to do.
	case reflect.Array, reflect.Slice:
		if reflectValue.Elem().Kind() != reflect.Ptr {
			return gerror.NewCodef(
				gcode.CodeInvalidParameter,
				`invalid parameter "%s", the element of slice should be type of pointer of struct, but given "%s"`,
				reflectValue.Type().String(), reflectValue.Elem().Type().String(),
			)
		}
	default:
		return gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid parameter "%s", should be type of pointer of struct`,
			reflectValue.Type().String(),
		)
	}
	if reflectValue.IsNil() {
		return gerror.NewCode(
			gcode.CodeInvalidParameter,
			`the pointed struct object should not be nil`,
		)
	}
	if !reflectValue.Elem().IsValid() {
		return gerror.NewCode(
			gcode.CodeInvalidParameter,
			`the pointed struct object should be valid`,
		)
	}
	fields, err := gstructs.Fields(gstructs.FieldsInput{
		Pointer:         reflectValue,
		RecursiveOption: gstructs.RecursiveOptionNone,
	})
	if err != nil {
		return err
	}
	for _, field := range fields {
		if defaultValue := field.TagDefault(); defaultValue != "" {
			if field.IsEmpty() {
				field.Value.Set(reflect.ValueOf(
					gconv.ConvertWithRefer(defaultValue, field.Value),
				))
			}
		}
	}
	return nil
}
