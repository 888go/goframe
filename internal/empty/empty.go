// 版权所有 GoFrame gf 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一个。

// 包 empty 提供了检查空/nil 变量的函数。
package empty
import (
	"reflect"
	"time"
	
	"github.com/888go/goframe/internal/reflection"
	)
// iString 用于在进行类型断言时，配合 String() 方法使用。
type iString interface {
	String() string
}

// iInterfaces 用于对 Interfaces 进行类型断言的 API。
type iInterfaces interface {
	Interfaces() []interface{}
}

// iMapStrAny 是支持将结构体参数转换为映射的接口。
type iMapStrAny interface {
	MapStrAny() map[string]interface{}
}

type iTime interface {
	Date() (year int, month time.Month, day int)
	IsZero() bool
}

// IsEmpty 检查给定的 `value` 是否为空。
// 当 `value` 为以下情形之一时，返回 true：0, nil, false, "", 切片/映射/通道长度为0，
// 否则返回 false。
//
// 参数 `traceSource` 用于在 `value` 类型为指针且指向另一个指针时追踪到源变量。
// 如果 `traceSource` 为 true 并且源变量为空，则它会返回 true。
// 注意，这可能会使用 reflect 特性，对性能稍有影响。
func IsEmpty(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}
// 它首先通过断言检查变量作为常见类型以提升性能，
// 然后再使用反射进行处理。
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
		// 最后，使用 reflect 包
		var rv reflect.Value
		if v, ok := value.(reflect.Value); ok {
			rv = v
		} else {
// ========================================
// 常用接口检查。
// ========================================
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

			rv = reflect.ValueOf(value)
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
		}
	}
	return false
}

// IsNil 检查给定的 `value` 是否为 nil，特别是对 interface{} 类型的值。
// 参数 `traceSource` 用于在 `value` 是指针类型且指向另一个指针时，追踪到源变量。如果源变量为 nil 并且 `traceSource` 为真，则返回 nil。
// 注意，该函数可能会使用 reflect 特性，这会对性能造成一定的影响。
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
	}
	return false
}
