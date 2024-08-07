// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gconv实现了对任何类型变量的强大而便捷的转换功能。
//
// 此包应尽量减少与其他包的依赖关系。
// md5:b18f07aca2be5125
package 转换类

import (
	"context"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"

	gbinary "github.com/888go/goframe/encoding/gbinary"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/reflection"
	gtime "github.com/888go/goframe/os/gtime"
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

	// StructTagPriority 定义了Map*/Struct*函数的默认优先级标签。
	// 注意，`gconv/param` 标签由旧版本的包使用。强烈建议未来改用简短的标签 `c/p`。
	// md5:c4b7d2fe8905ed52
	StructTagPriority = gtag.StructTagPriority
)

// X取字节将`any`转换为byte。 md5:aeef919e3fba4f95
func X取字节(值 interface{}) byte {
	if v, ok := 值.(byte); ok {
		return v
	}
	return X取正整数8位(值)
}

// X取字节集 将 `any` 转换为 []byte。 md5:06125d6ba5f449a5
func X取字节集(any interface{}) []byte {
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
			return f.X取字节集()
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
				int32Value := X取整数32位(originValueAndKind.OriginValue.Index(i).Interface())
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

// X取字符 将 `any` 转换为 rune。 md5:3459f7528861cc23
func X取字符(值 interface{}) rune {
	if v, ok := 值.(rune); ok {
		return v
	}
	return X取整数32位(值)
}

// X取字符切片将`any`转换为[]rune。 md5:25552cd961d1d6bb
func X取字符切片(值 interface{}) []rune {
	if v, ok := 值.([]rune); ok {
		return v
	}
	return []rune(String(值))
}

// String 将 `any` 转换为字符串。它是最常用的转换函数。
// md5:722d0704c061781b
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
			// 则使用该接口来进行转换
			// md5:08e76021f60d81ed
			return f.String()
		}
		if f, ok := value.(iError); ok {
			// /* 如果该变量实现了Error()接口，
			//    则使用该接口进行转换 */
			// md5:7c7c512864a0b034
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
				// 最后，我们使用json.Marshal进行转换。 md5:57829b67798bbc93
		if jsonContent, err := json.Marshal(value); err != nil {
			return fmt.Sprint(value)
		} else {
			return string(jsonContent)
		}
	}
}

// X取布尔 将 `any` 转换为布尔值。
// 如果 `any` 是：false，""，0，"false"，"off"，"no"，空切片/映射，则返回 false。
// md5:b9d150a8798a274a
func X取布尔(any interface{}) bool {
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
			return f.X取布尔()
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

// checkJsonAndUnmarshalUseNumber 检查给定的 `any` 是否为 JSON 格式的字符串值，并使用 `json.UnmarshalUseNumber` 进行转换。 md5:ce3edf33e8eea76c
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
