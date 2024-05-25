// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457

// Package empty 提供检查空/空指针变量的函数。. md5:4be7b468c813f750
package empty

import (
	"reflect"
	"time"

	"github.com/gogf/gf/v2/internal/reflection"
)

// iString 用于类型断言API，用于String()。. md5:8ec0af717c4f530e
type iString interface {
	String() string
}

// iInterfaces 用于接口类型的断言API。. md5:9162512bdb64ee64
type iInterfaces interface {
	Interfaces() []interface{}
}

// iMapStrAny 是一个接口，支持将结构体参数转换为映射。. md5:cfd4642c77fca6ec
type iMapStrAny interface {
	MapStrAny() map[string]interface{}
}

type iTime interface {
	Date() (year int, month time.Month, day int)
	IsZero() bool
}

// IsEmpty 检查给定的 `value` 是否为空。
// 如果 `value` 为以下情况，函数返回 true：0, nil, false, "", slice/映射/通道的长度为0，
// 否则返回 false。
//
// 参数 `traceSource` 用于在 `value` 是指向指针的指针类型时追踪源变量。
// 当 `traceSource` 为 true 且源变量为空时，返回 true。
// 注意，这可能使用反射功能，可能会稍微影响性能。
// md5:343856f448e80aef
func IsEmpty(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}
// 它首先使用断言检查变量为常见类型，以提高性能，然后使用反射。
// md5:9722a28f813b5ddb
	switch result := value.(type) {
	case int:
		return result == 0
	case int8:
		return result == 0
	case int16:
		return result == 0
	case int32:
		return result == 0
	case int64:
		return result == 0
	case uint:
		return result == 0
	case uint8:
		return result == 0
	case uint16:
		return result == 0
	case uint32:
		return result == 0
	case uint64:
		return result == 0
	case float32:
		return result == 0
	case float64:
		return result == 0
	case bool:
		return !result
	case string:
		return result == ""
	case []byte:
		return len(result) == 0
	case []rune:
		return len(result) == 0
	case []int:
		return len(result) == 0
	case []string:
		return len(result) == 0
	case []float32:
		return len(result) == 0
	case []float64:
		return len(result) == 0
	case map[string]interface{}:
		return len(result) == 0

	default:
		// 最后，使用反射。. md5:e4ce8ad5b39b80cd
		var rv reflect.Value
		if v, ok := value.(reflect.Value); ok {
			rv = v
		} else {
			rv = reflect.ValueOf(value)
			if IsNil(rv) {
				return true
			}

// =========================
// 公共接口检查。
// =========================
// md5:e561bbb4afe04dee
			if f, ok := value.(iTime); ok {
				if f == (*time.Time)(nil) {
					return true
				}
				return f.IsZero()
			}
			if f, ok := value.(iString); ok {
				if f == nil {
					return true
				}
				return f.String() == ""
			}
			if f, ok := value.(iInterfaces); ok {
				if f == nil {
					return true
				}
				return len(f.Interfaces()) == 0
			}
			if f, ok := value.(iMapStrAny); ok {
				if f == nil {
					return true
				}
				return len(f.MapStrAny()) == 0
			}
		}

		switch rv.Kind() {
		case reflect.Bool:
			return !rv.Bool()

		case
			reflect.Int,
			reflect.Int8,
			reflect.Int16,
			reflect.Int32,
			reflect.Int64:
			return rv.Int() == 0

		case
			reflect.Uint,
			reflect.Uint8,
			reflect.Uint16,
			reflect.Uint32,
			reflect.Uint64,
			reflect.Uintptr:
			return rv.Uint() == 0

		case
			reflect.Float32,
			reflect.Float64:
			return rv.Float() == 0

		case reflect.String:
			return rv.Len() == 0

		case reflect.Struct:
			var fieldValueInterface interface{}
			for i := 0; i < rv.NumField(); i++ {
				fieldValueInterface, _ = reflection.ValueToInterface(rv.Field(i))
				if !IsEmpty(fieldValueInterface) {
					return false
				}
			}
			return true

		case
			reflect.Chan,
			reflect.Map,
			reflect.Slice,
			reflect.Array:
			return rv.Len() == 0

		case reflect.Ptr:
			if len(traceSource) > 0 && traceSource[0] {
				return IsEmpty(rv.Elem())
			}
			return rv.IsNil()

		case
			reflect.Func,
			reflect.Interface,
			reflect.UnsafePointer:
			return rv.IsNil()

		case reflect.Invalid:
			return true

		default:
			return false
		}
	}
}

// IsNil 检查给定的 `value` 是否为 nil，特别是对于 interface{} 类型的值。
// 参数 `traceSource` 用于在给定的 `value` 是指向指针的指针类型时，追踪到源变量。
// 当 `traceSource` 为真且源为 nil 时，它会返回 nil。
// 注意，该函数可能使用反射功能，这可能稍微影响性能。
// md5:c12efd8c176fc73a
func IsNil(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}
	var rv reflect.Value
	if v, ok := value.(reflect.Value); ok {
		rv = v
	} else {
		rv = reflect.ValueOf(value)
	}
	switch rv.Kind() {
	case reflect.Chan,
		reflect.Map,
		reflect.Slice,
		reflect.Func,
		reflect.Interface,
		reflect.UnsafePointer:
		return !rv.IsValid() || rv.IsNil()

	case reflect.Ptr:
		if len(traceSource) > 0 && traceSource[0] {
			for rv.Kind() == reflect.Ptr {
				rv = rv.Elem()
			}
			if !rv.IsValid() {
				return true
			}
			if rv.Kind() == reflect.Ptr {
				return rv.IsNil()
			}
		} else {
			return !rv.IsValid() || rv.IsNil()
		}

	default:
		return false
	}
	return false
}
