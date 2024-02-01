// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gconv
import (
	"reflect"
	"time"
	
	"github.com/888go/goframe/os/gtime"
	)
// Convert 将变量 `fromValue` 转换为类型 `toTypeName`，其中 `toTypeName` 由字符串指定。
//
// 可选参数 `extraParams` 用于提供此次转换所需的额外必要参数。
// 它支持基于类型名称字符串的基本常见类型的转换。
func Convert(fromValue interface{}, toTypeName string, extraParams ...interface{}) interface{} {
	return doConvert(doConvertInput{
		FromValue:  fromValue,
		ToTypeName: toTypeName,
		ReferValue: nil,
		Extra:      extraParams,
	})
}

// ConvertWithRefer 将变量 `fromValue` 转换为由值 `referValue` 所引用的类型。
//
// 可选参数 `extraParams` 用于提供本次转换所需的额外必要参数。
// 它支持基于类型名称字符串的基本常见类型的转换。
func ConvertWithRefer(fromValue interface{}, referValue interface{}, extraParams ...interface{}) interface{} {
	var referValueRf reflect.Value
	if v, ok := referValue.(reflect.Value); ok {
		referValueRf = v
	} else {
		referValueRf = reflect.ValueOf(referValue)
	}
	return doConvert(doConvertInput{
		FromValue:  fromValue,
		ToTypeName: referValueRf.Type().String(),
		ReferValue: referValue,
		Extra:      extraParams,
	})
}

type doConvertInput struct {
	FromValue  interface{}   // 需要转换的原始值。
	ToTypeName string        // 字符串形式的目标值类型名称
	ReferValue interface{}   // 指针引用的值，类型为 `ToTypeName` 的值。注意，其实际类型可能为 reflect.Value。
	Extra      []interface{} // 用于实现转换功能的额外值。
// 标记该值已转换并设置为`ReferValue`。调用者可以忽略返回的结果。
// 这是一个用于内部使用的属性。
	alreadySetToReferValue bool
}

// doConvert 执行常用类型的转换。
func doConvert(in doConvertInput) (convertedValue interface{}) {
	switch in.ToTypeName {
	case "int":
		return Int(in.FromValue)
	case "*int":
		if _, ok := in.FromValue.(*int); ok {
			return in.FromValue
		}
		v := Int(in.FromValue)
		return &v

	case "int8":
		return Int8(in.FromValue)
	case "*int8":
		if _, ok := in.FromValue.(*int8); ok {
			return in.FromValue
		}
		v := Int8(in.FromValue)
		return &v

	case "int16":
		return Int16(in.FromValue)
	case "*int16":
		if _, ok := in.FromValue.(*int16); ok {
			return in.FromValue
		}
		v := Int16(in.FromValue)
		return &v

	case "int32":
		return Int32(in.FromValue)
	case "*int32":
		if _, ok := in.FromValue.(*int32); ok {
			return in.FromValue
		}
		v := Int32(in.FromValue)
		return &v

	case "int64":
		return Int64(in.FromValue)
	case "*int64":
		if _, ok := in.FromValue.(*int64); ok {
			return in.FromValue
		}
		v := Int64(in.FromValue)
		return &v

	case "uint":
		return Uint(in.FromValue)
	case "*uint":
		if _, ok := in.FromValue.(*uint); ok {
			return in.FromValue
		}
		v := Uint(in.FromValue)
		return &v

	case "uint8":
		return Uint8(in.FromValue)
	case "*uint8":
		if _, ok := in.FromValue.(*uint8); ok {
			return in.FromValue
		}
		v := Uint8(in.FromValue)
		return &v

	case "uint16":
		return Uint16(in.FromValue)
	case "*uint16":
		if _, ok := in.FromValue.(*uint16); ok {
			return in.FromValue
		}
		v := Uint16(in.FromValue)
		return &v

	case "uint32":
		return Uint32(in.FromValue)
	case "*uint32":
		if _, ok := in.FromValue.(*uint32); ok {
			return in.FromValue
		}
		v := Uint32(in.FromValue)
		return &v

	case "uint64":
		return Uint64(in.FromValue)
	case "*uint64":
		if _, ok := in.FromValue.(*uint64); ok {
			return in.FromValue
		}
		v := Uint64(in.FromValue)
		return &v

	case "float32":
		return Float32(in.FromValue)
	case "*float32":
		if _, ok := in.FromValue.(*float32); ok {
			return in.FromValue
		}
		v := Float32(in.FromValue)
		return &v

	case "float64":
		return Float64(in.FromValue)
	case "*float64":
		if _, ok := in.FromValue.(*float64); ok {
			return in.FromValue
		}
		v := Float64(in.FromValue)
		return &v

	case "bool":
		return Bool(in.FromValue)
	case "*bool":
		if _, ok := in.FromValue.(*bool); ok {
			return in.FromValue
		}
		v := Bool(in.FromValue)
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
		return Bytes(in.FromValue)
	case "[]int":
		return Ints(in.FromValue)
	case "[]int32":
		return Int32s(in.FromValue)
	case "[]int64":
		return Int64s(in.FromValue)
	case "[]uint":
		return Uints(in.FromValue)
	case "[]uint8":
		return Bytes(in.FromValue)
	case "[]uint32":
		return Uint32s(in.FromValue)
	case "[]uint64":
		return Uint64s(in.FromValue)
	case "[]float32":
		return Float32s(in.FromValue)
	case "[]float64":
		return Float64s(in.FromValue)
	case "[]string":
		return Strings(in.FromValue)

	case "Time", "time.Time":
		if len(in.Extra) > 0 {
			return Time(in.FromValue, String(in.Extra[0]))
		}
		return Time(in.FromValue)
	case "*time.Time":
		var v time.Time
		if len(in.Extra) > 0 {
			v = Time(in.FromValue, String(in.Extra[0]))
		} else {
			if _, ok := in.FromValue.(*time.Time); ok {
				return in.FromValue
			}
			v = Time(in.FromValue)
		}
		return &v

	case "GTime", "gtime.Time":
		if len(in.Extra) > 0 {
			if v := GTime(in.FromValue, String(in.Extra[0])); v != nil {
				return *v
			} else {
				return *gtime.New()
			}
		}
		if v := GTime(in.FromValue); v != nil {
			return *v
		} else {
			return *gtime.New()
		}
	case "*gtime.Time":
		if len(in.Extra) > 0 {
			if v := GTime(in.FromValue, String(in.Extra[0])); v != nil {
				return v
			} else {
				return gtime.New()
			}
		}
		if v := GTime(in.FromValue); v != nil {
			return v
		} else {
			return gtime.New()
		}

	case "Duration", "time.Duration":
		return Duration(in.FromValue)
	case "*time.Duration":
		if _, ok := in.FromValue.(*time.Duration); ok {
			return in.FromValue
		}
		v := Duration(in.FromValue)
		return &v

	case "map[string]string":
		return MapStrStr(in.FromValue)

	case "map[string]interface{}":
		return Map(in.FromValue)

	case "[]map[string]interface{}":
		return Maps(in.FromValue)

	case "RawMessage", "json.RawMessage":
		return Bytes(in.FromValue)

	default:
		if in.ReferValue != nil {
			var referReflectValue reflect.Value
			if v, ok := in.ReferValue.(reflect.Value); ok {
				referReflectValue = v
			} else {
				referReflectValue = reflect.ValueOf(in.ReferValue)
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
// 自定义类型指针的类型转换
// 示例：
// 定义自定义类型 PayMode，其为 int 类型的别名
// 定义结构体 Req，其中包含一个指向 PayMode 类型的指针 Mode
//
// type PayMode int
// type Req struct{
//     Mode *PayMode
// }
//
// 通过 Struct 函数将 `{"Mode": 1000}` 转换并解析到 req 指针所指向的结构体中
// Struct(`{"Mode": 1000}`, &req)
				originType := referReflectValue.Type().Elem()
				switch originType.Kind() {
				case reflect.Struct:
					// 不支持某些类型。
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
