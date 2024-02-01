// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gconv 提供了强大且方便的任意类型变量转换功能。
//
// 本包应尽量减少对其他包的依赖。
package gconv
import (
	"context"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
	
	"github.com/888go/goframe/encoding/gbinary"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/reflection"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/util/gtag"
	)
var (
	// Empty strings.
	emptyStringMap = map[string]struct{}{
		"":      {},
		"0":     {},
		"no":    {},
		"off":   {},
		"false": {},
	}

// StructTagPriority 定义了 Map*/Struct* 函数的默认优先级标签。
// 注意，`gconv/param` 标签在旧版本包中使用。
// 强烈建议在未来使用简写标签 `c/p` 代替。
	StructTagPriority = []string{
		gtag.GConv, gtag.Param, gtag.GConvShort, gtag.ParamShort, gtag.Json,
	}
)

// Byte将`any`转换为字节。
func Byte(any interface{}) byte {
	if v, ok := any.(byte); ok {
		return v
	}
	return Uint8(any)
}

// Bytes 将 `any` 类型转换为 []byte 类型。
func Bytes(any interface{}) []byte {
	if any == nil {
		return nil
	}
	switch value := any.(type) {
	case string:
		return []byte(value)

	case []byte:
		return value

	default:
		if f, ok := value.(iBytes); ok {
			return f.Bytes()
		}
		originValueAndKind := reflection.OriginValueAndKind(any)
		switch originValueAndKind.OriginKind {
		case reflect.Map:
			bytes, err := json.Marshal(any)
			if err != nil {
				intlog.Errorf(context.TODO(), `%+v`, err)
			}
			return bytes

		case reflect.Array, reflect.Slice:
			var (
				ok    = true
				bytes = make([]byte, originValueAndKind.OriginValue.Len())
			)
			for i := range bytes {
				int32Value := Int32(originValueAndKind.OriginValue.Index(i).Interface())
				if int32Value < 0 || int32Value > math.MaxUint8 {
					ok = false
					break
				}
				bytes[i] = byte(int32Value)
			}
			if ok {
				return bytes
			}
		}
		return gbinary.Encode(any)
	}
}

// Rune将`any`转换为rune类型。
func Rune(any interface{}) rune {
	if v, ok := any.(rune); ok {
		return v
	}
	return Int32(any)
}

// Runes 将 `any` 转换为 []rune 类型。
func Runes(any interface{}) []rune {
	if v, ok := any.([]rune); ok {
		return v
	}
	return []rune(String(any))
}

// String 将`any`转换为字符串。
// 这是最常用的转换函数。
func String(any interface{}) string {
	if any == nil {
		return ""
	}
	switch value := any.(type) {
	case int:
		return strconv.Itoa(value)
	case int8:
		return strconv.Itoa(int(value))
	case int16:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.FormatInt(value, 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	case string:
		return value
	case []byte:
		return string(value)
	case time.Time:
		if value.IsZero() {
			return ""
		}
		return value.String()
	case *time.Time:
		if value == nil {
			return ""
		}
		return value.String()
	case gtime.Time:
		if value.IsZero() {
			return ""
		}
		return value.String()
	case *gtime.Time:
		if value == nil {
			return ""
		}
		return value.String()
	default:
		// Empty checks.
		if value == nil {
			return ""
		}
		if f, ok := value.(iString); ok {
// 如果变量实现了String()接口，
// 那么就使用该接口来执行转换
			return f.String()
		}
		if f, ok := value.(iError); ok {
// 如果变量实现了Error()接口，
// 那么使用该接口进行转换
			return f.Error()
		}
		// Reflect checks.
		var (
			rv   = reflect.ValueOf(value)
			kind = rv.Kind()
		)
		switch kind {
		case reflect.Chan,
			reflect.Map,
			reflect.Slice,
			reflect.Func,
			reflect.Ptr,
			reflect.Interface,
			reflect.UnsafePointer:
			if rv.IsNil() {
				return ""
			}
		case reflect.String:
			return rv.String()
		}
		if kind == reflect.Ptr {
			return String(rv.Elem().Interface())
		}
		// 最后，我们使用json.Marshal将数据转换。
		if jsonContent, err := json.Marshal(value); err != nil {
			return fmt.Sprint(value)
		} else {
			return string(jsonContent)
		}
	}
}

// Bool将`any`转换为布尔值。
// 当`any`为：false、空字符串、0、"false"、"off"、"no"或空切片/映射时，返回false。
func Bool(any interface{}) bool {
	if any == nil {
		return false
	}
	switch value := any.(type) {
	case bool:
		return value
	case []byte:
		if _, ok := emptyStringMap[strings.ToLower(string(value))]; ok {
			return false
		}
		return true
	case string:
		if _, ok := emptyStringMap[strings.ToLower(value)]; ok {
			return false
		}
		return true
	default:
		if f, ok := value.(iBool); ok {
			return f.Bool()
		}
		rv := reflect.ValueOf(any)
		switch rv.Kind() {
		case reflect.Ptr:
			return !rv.IsNil()
		case reflect.Map:
			fallthrough
		case reflect.Array:
			fallthrough
		case reflect.Slice:
			return rv.Len() != 0
		case reflect.Struct:
			return true
		default:
			s := strings.ToLower(String(any))
			if _, ok := emptyStringMap[s]; ok {
				return false
			}
			return true
		}
	}
}

// checkJsonAndUnmarshalUseNumber 检查给定的 `any` 是否为格式化的 JSON 字符串值，并使用 `json.UnmarshalUseNumber` 进行转换。
func checkJsonAndUnmarshalUseNumber(any interface{}, target interface{}) bool {
	switch r := any.(type) {
	case []byte:
		if json.Valid(r) {
			if err := json.UnmarshalUseNumber(r, &target); err != nil {
				return false
			}
			return true
		}

	case string:
		anyAsBytes := []byte(r)
		if json.Valid(anyAsBytes) {
			if err := json.UnmarshalUseNumber(anyAsBytes, &target); err != nil {
				return false
			}
			return true
		}
	}
	return false
}
