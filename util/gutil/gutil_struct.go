// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 工具类

import (
	"reflect"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/os/gstructs"
	gconv "github.com/888go/goframe/util/gconv"
)

// X结构体到切片 将结构体转换为一个键值对作为元素的切片。
// 例如：{"K1": "v1", "K2": "v2"} => ["K1", "v1", "K2", "v2"]
// md5:ca8c34ec711fb0de
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
		// 如果在结构体属性中找到了gconv标签，它将使用gconv标签名而不是属性名。
		// md5:697077ff458895f0
		for k, v := range gconv.X取Map(reflectValue) {
			array = append(array, k)
			array = append(array, v)
		}
		return array
	}
	return nil
}

// FillStructWithDefault 使用`default/d`标签的值填充指向的结构体的属性。参数`structPtr`应该是`*struct`或`[]*struct`类型的一种。
// md5:5777fe6fdb6efa8a
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
			return gerror.X创建错误码并格式化(
				gcode.CodeInvalidParameter,
				`invalid parameter "%s", the element of slice should be type of pointer of struct, but given "%s"`,
				reflectValue.Type().String(), reflectValue.Elem().Type().String(),
			)
		}
	default:
		return gerror.X创建错误码并格式化(
			gcode.CodeInvalidParameter,
			`invalid parameter "%s", should be type of pointer of struct`,
			reflectValue.Type().String(),
		)
	}
	if reflectValue.IsNil() {
		return gerror.X创建错误码(
			gcode.CodeInvalidParameter,
			`the pointed struct object should not be nil`,
		)
	}
	if !reflectValue.Elem().IsValid() {
		return gerror.X创建错误码(
			gcode.CodeInvalidParameter,
			`the pointed struct object should be valid`,
		)
	}
	fields, err := gstructs.Fields(gstructs.FieldsInput{
		Pointer:         reflectValue,
		RecursiveOption: gstructs.RecursiveOptionEmbedded,
	})
	if err != nil {
		return err
	}
	for _, field := range fields {
		if field.OriginalKind() == reflect.Struct {
			err := FillStructWithDefault(field.OriginalValue().Addr())
			if err != nil {
				return err
			}
			continue
		}

		if defaultValue := field.TagDefault(); defaultValue != "" {
			if field.IsEmpty() {
				field.Value.Set(reflect.ValueOf(
					gconv.X按参考值类型转换(defaultValue, field.Value),
				))
			}
		}
	}

	return nil
}
