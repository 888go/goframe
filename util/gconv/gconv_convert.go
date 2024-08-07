// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 转换类

import (
	"context"
	"reflect"
	"time"

	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/json"
	gtime "github.com/888go/goframe/os/gtime"
)

// X按名称转换 将变量 `fromValue` 转换为类型 `toTypeName`，其中 `toTypeName` 由字符串指定。
//
// 可选参数 `extraParams` 用于提供此转换所需的额外参数。
// 它支持基于类型名称字符串的常见基本类型转换。
// md5:e081c8fc6552be4c
func X按名称转换(值 interface{}, 类型名称 string, extraParams ...interface{}) interface{} {
	return doConvert(doConvertInput{
		FromValue:  值,
		ToTypeName: 类型名称,
		ReferValue: nil,
		Extra:      extraParams,
	})
}

// X按参考值类型转换 将变量 `fromValue` 转换为由 `referValue` 指定的类型。
//
// 可选参数 `extraParams` 用于此转换所需的额外参数。它支持基于类型名称字符串的常见基本类型转换。
// md5:0badd37157c72db1
func X按参考值类型转换(值 interface{}, 参考值 interface{}, extraParams ...interface{}) interface{} {
	var referValueRf reflect.Value
	if v, ok := 参考值.(reflect.Value); ok {
		referValueRf = v
	} else {
		referValueRf = reflect.ValueOf(参考值)
	}
	return doConvert(doConvertInput{
		FromValue:  值,
		ToTypeName: referValueRf.Type().String(),
		ReferValue: 参考值,
		Extra:      extraParams,
	})
}

type doConvertInput struct {
	FromValue  interface{}   // 要转换的值。 md5:b9384f7def81e56d
	ToTypeName string        // 目标值类型名称（字符串形式）。 md5:56863f5417d5b24f
	ReferValue interface{}   // 引用的值，类型为`ToTypeName`。请注意，它的类型可能是`reflect.Value`。 md5:7e9c4375ec4d26f3
	Extra      []interface{} // 用于实现转换的额外值。 md5:c5e0f680118ba627
	// 标记该值已经转换并设置为`ReferValue`。调用者可以忽略返回的结果。
	// 这是一个用于内部使用的属性。
	// md5:91187d21c0d0ac16
	alreadySetToReferValue bool
}

// doConvert 执行常用类型转换。 md5:a4f52e85ed63dbe3
func doConvert(in doConvertInput) (convertedValue interface{}) {
	switch in.ToTypeName {
	case "int":
		return X取整数(in.FromValue)
	case "*int":
		if _, ok := in.FromValue.(*int); ok {
			return in.FromValue
		}
		v := X取整数(in.FromValue)
		return &v

	case "int8":
		return X取整数8位(in.FromValue)
	case "*int8":
		if _, ok := in.FromValue.(*int8); ok {
			return in.FromValue
		}
		v := X取整数8位(in.FromValue)
		return &v

	case "int16":
		return X取整数16位(in.FromValue)
	case "*int16":
		if _, ok := in.FromValue.(*int16); ok {
			return in.FromValue
		}
		v := X取整数16位(in.FromValue)
		return &v

	case "int32":
		return X取整数32位(in.FromValue)
	case "*int32":
		if _, ok := in.FromValue.(*int32); ok {
			return in.FromValue
		}
		v := X取整数32位(in.FromValue)
		return &v

	case "int64":
		return X取整数64位(in.FromValue)
	case "*int64":
		if _, ok := in.FromValue.(*int64); ok {
			return in.FromValue
		}
		v := X取整数64位(in.FromValue)
		return &v

	case "uint":
		return X取正整数(in.FromValue)
	case "*uint":
		if _, ok := in.FromValue.(*uint); ok {
			return in.FromValue
		}
		v := X取正整数(in.FromValue)
		return &v

	case "uint8":
		return X取正整数8位(in.FromValue)
	case "*uint8":
		if _, ok := in.FromValue.(*uint8); ok {
			return in.FromValue
		}
		v := X取正整数8位(in.FromValue)
		return &v

	case "uint16":
		return X取正整数16位(in.FromValue)
	case "*uint16":
		if _, ok := in.FromValue.(*uint16); ok {
			return in.FromValue
		}
		v := X取正整数16位(in.FromValue)
		return &v

	case "uint32":
		return X取正整数32位(in.FromValue)
	case "*uint32":
		if _, ok := in.FromValue.(*uint32); ok {
			return in.FromValue
		}
		v := X取正整数32位(in.FromValue)
		return &v

	case "uint64":
		return X取正整数64位(in.FromValue)
	case "*uint64":
		if _, ok := in.FromValue.(*uint64); ok {
			return in.FromValue
		}
		v := X取正整数64位(in.FromValue)
		return &v

	case "float32":
		return X取小数32位(in.FromValue)
	case "*float32":
		if _, ok := in.FromValue.(*float32); ok {
			return in.FromValue
		}
		v := X取小数32位(in.FromValue)
		return &v

	case "float64":
		return X取小数64位(in.FromValue)
	case "*float64":
		if _, ok := in.FromValue.(*float64); ok {
			return in.FromValue
		}
		v := X取小数64位(in.FromValue)
		return &v

	case "bool":
		return X取布尔(in.FromValue)
	case "*bool":
		if _, ok := in.FromValue.(*bool); ok {
			return in.FromValue
		}
		v := X取布尔(in.FromValue)
		return &v

	case "string":
		return String(in.FromValue)
	case "*string":
		if _, ok := in.FromValue.(*string); ok {
			return in.FromValue
		}
		v := String(in.FromValue)
		return &v

	case "[]byte":
		return X取字节集(in.FromValue)
	case "[]int":
		return X取整数切片(in.FromValue)
	case "[]int32":
		return X取整数32位切片(in.FromValue)
	case "[]int64":
		return X取整数64位切片(in.FromValue)
	case "[]uint":
		return X取正整数切片(in.FromValue)
	case "[]uint8":
		return X取字节集(in.FromValue)
	case "[]uint32":
		return X取正整数32位切片(in.FromValue)
	case "[]uint64":
		return X取正整数64位切片(in.FromValue)
	case "[]float32":
		return X取小数32位切片(in.FromValue)
	case "[]float64":
		return X取小数64位切片(in.FromValue)
	case "[]string":
		return X取文本切片(in.FromValue)

	case "Time", "time.Time":
		if len(in.Extra) > 0 {
			return X取时间(in.FromValue, String(in.Extra[0]))
		}
		return X取时间(in.FromValue)
	case "*time.Time":
		var v time.Time
		if len(in.Extra) > 0 {
			v = X取时间(in.FromValue, String(in.Extra[0]))
		} else {
			if _, ok := in.FromValue.(*time.Time); ok {
				return in.FromValue
			}
			v = X取时间(in.FromValue)
		}
		return &v

	case "GTime", "gtime.Time", "时间类", "时间类.Time":
		if len(in.Extra) > 0 {
			if v := X取gtime时间类(in.FromValue, String(in.Extra[0])); v != nil {
				return *v
			} else {
				return *gtime.X创建()
			}
		}
		if v := X取gtime时间类(in.FromValue); v != nil {
			return *v
		} else {
			return *gtime.X创建()
		}
	case "*gtime.Time", "*时间类.Time":
		if len(in.Extra) > 0 {
			if v := X取gtime时间类(in.FromValue, String(in.Extra[0])); v != nil {
				return v
			} else {
				return gtime.X创建()
			}
		}
		if v := X取gtime时间类(in.FromValue); v != nil {
			return v
		} else {
			return gtime.X创建()
		}

	case "Duration", "time.Duration":
		return X取时长(in.FromValue)
	case "*time.Duration":
		if _, ok := in.FromValue.(*time.Duration); ok {
			return in.FromValue
		}
		v := X取时长(in.FromValue)
		return &v

	case "map[string]string":
		return X取文本Map(in.FromValue)

	case "map[string]interface{}":
		return X取Map(in.FromValue)

	case "[]map[string]interface{}":
		return X取Map切片(in.FromValue)

	case "RawMessage", "json.RawMessage":
		// issue 3449
		bytes, err := json.Marshal(in.FromValue)
		if err != nil {
			intlog.Errorf(context.TODO(), `%+v`, err)
		}
		return bytes

	default:
		if in.ReferValue != nil {
			var referReflectValue reflect.Value
			if v, ok := in.ReferValue.(reflect.Value); ok {
				referReflectValue = v
			} else {
				referReflectValue = reflect.ValueOf(in.ReferValue)
			}
			var fromReflectValue reflect.Value
			if v, ok := in.FromValue.(reflect.Value); ok {
				fromReflectValue = v
			} else {
				fromReflectValue = reflect.ValueOf(in.FromValue)
			}

			// custom converter.
			if dstReflectValue, ok, _ := callCustomConverterWithRefer(fromReflectValue, referReflectValue); ok {
				return dstReflectValue.Interface()
			}

			defer func() {
				if recover() != nil {
					in.alreadySetToReferValue = false
					if err := bindVarToReflectValue(referReflectValue, in.FromValue, nil); err == nil {
						in.alreadySetToReferValue = true
						convertedValue = referReflectValue.Interface()
					}
				}
			}()
			switch referReflectValue.Kind() {
			case reflect.Ptr:
				// 自定义类型指针的类型转换。
				// 例如：
				// type PayMode int
				// type Req struct{
				//     Mode *PayMode
				// }
				// 
				// Struct(`{"Mode": 1000}`, &req)
				// md5:d218e7f3f409c5f7
				originType := referReflectValue.Type().Elem()
				switch originType.Kind() {
				case reflect.Struct:
										// 不支持某些类型。 md5:74a7c80d66154fb9
				default:
					in.ToTypeName = originType.Kind().String()
					in.ReferValue = nil
					refElementValue := reflect.ValueOf(doConvert(in))
					originTypeValue := reflect.New(refElementValue.Type()).Elem()
					originTypeValue.Set(refElementValue)
					in.alreadySetToReferValue = true
					return originTypeValue.Addr().Convert(referReflectValue.Type()).Interface()
				}

			case reflect.Map:
				var targetValue = reflect.New(referReflectValue.Type()).Elem()
				if err := doMapToMap(in.FromValue, targetValue); err == nil {
					in.alreadySetToReferValue = true
				}
				return targetValue.Interface()
			}
			in.ToTypeName = referReflectValue.Kind().String()
			in.ReferValue = nil
			in.alreadySetToReferValue = true
			convertedValue = reflect.ValueOf(doConvert(in)).Convert(referReflectValue.Type()).Interface()
			return convertedValue
		}
		return in.FromValue
	}
}

func doConvertWithReflectValueSet(reflectValue reflect.Value, in doConvertInput) {
	convertedValue := doConvert(in)
	if !in.alreadySetToReferValue {
		reflectValue.Set(reflect.ValueOf(convertedValue))
	}
}
