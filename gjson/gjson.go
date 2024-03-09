// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gjson 提供了方便的 API 用于处理 JSON/XML/INI/YAML/TOML 数据。
package json类

import (
	"reflect"
	"strconv"
	"strings"
	
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/888go/goframe/gjson/internal/reflection"
	"github.com/888go/goframe/gjson/internal/rwmutex"
	"github.com/888go/goframe/gjson/internal/utils"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type ContentType string

const (
	ContentTypeJson       ContentType = `json`
	ContentTypeJs         ContentType = `js`
	ContentTypeXml        ContentType = `xml`
	ContentTypeIni        ContentType = `ini`
	ContentTypeYaml       ContentType = `yaml`
	ContentTypeYml        ContentType = `yml`
	ContentTypeToml       ContentType = `toml`
	ContentTypeProperties ContentType = `properties`
)

const (
	defaultSplitChar = '.' // 分隔符字符，用于分层数据访问。
)

// Json 是自定义的 JSON 结构体。
type Json struct {
	mu rwmutex.RWMutex
	p  *interface{} // Pointer for hierarchical data access, it's the root of data in default.
	c  byte         // 字符分隔符（默认为'.'）
	vc bool         // 暴力检查（默认为false），用于在分层数据键包含分隔符字符时访问数据。
}

// Json对象创建/加载的选项。
type Options struct {
	Safe      bool        // 标记此对象适用于并发安全的使用场景。这尤其针对Json对象的创建。
	Tags      string      // 自定义解析优先级标签，例如："json,yaml,MyTag"。这主要用于将结构体解析为Json对象时。
	Type      ContentType // Type 指定数据内容类型，例如：json、xml、yaml、toml、ini。
	StrNumber bool        // StrNumber 使 Decoder 在将数字反序列化到 interface{} 时，将其作为字符串处理而非 float64。
}

// iInterfaces 用于对 Interfaces() 方法进行类型断言。
type iInterfaces interface {
	Interfaces() []interface{}
}

// iMapStrAny 是支持将结构体参数转换为映射的接口。
type iMapStrAny interface {
	MapStrAny() map[string]interface{}
}

// iVal 是用于获取底层 interface{} 的接口。
type iVal interface {
	Val() interface{}
}

// setValue 通过 `pattern` 将 `value` 设置为 `j`。
// 注意：
// 1. 如果 value 为 nil 且 removed 为 true，表示删除这个值；
// 2. 在层次数据搜索、节点创建和数据赋值过程中较为复杂。
func (j *Json) setValue(pattern string, value interface{}, removed bool) error {
	var (
		err    error
		array  = strings.Split(pattern, string(j.c))
		length = len(array)
	)
	if value, err = j.convertValue(value); err != nil {
		return err
	}
	// 初始化检查。
	if *j.p == nil {
		if gstr.IsNumeric(array[0]) {
			*j.p = make([]interface{}, 0)
		} else {
			*j.p = make(map[string]interface{})
		}
	}
	var (
		pparent *interface{} = nil // Parent pointer.
		pointer *interface{} = j.p // Current pointer.
	)
	j.mu.Lock()
	defer j.mu.Unlock()
	for i := 0; i < length; i++ {
		switch (*pointer).(type) {
		case map[string]interface{}:
			if i == length-1 {
				if removed && value == nil {
					// 从map中删除项目。
					delete((*pointer).(map[string]interface{}), array[i])
				} else {
					if (*pointer).(map[string]interface{}) == nil {
						*pointer = map[string]interface{}{}
					}
					(*pointer).(map[string]interface{})[array[i]] = value
				}
			} else {
				// 如果键在映射中不存在。
				if v, ok := (*pointer).(map[string]interface{})[array[i]]; !ok {
					if removed && value == nil {
						goto done
					}
					// 创建新的节点。
					if gstr.IsNumeric(array[i+1]) {
						// 创建数组节点
						n, _ := strconv.Atoi(array[i+1])
						var v interface{} = make([]interface{}, n+1)
						pparent = j.setPointerWithValue(pointer, array[i], v)
						pointer = &v
					} else {
						// 创建映射节点。
						var v interface{} = make(map[string]interface{})
						pparent = j.setPointerWithValue(pointer, array[i], v)
						pointer = &v
					}
				} else {
					pparent = pointer
					pointer = &v
				}
			}

		case []interface{}:
			// A string key.
			if !gstr.IsNumeric(array[i]) {
				if i == length-1 {
					*pointer = map[string]interface{}{array[i]: value}
				} else {
					var v interface{} = make(map[string]interface{})
					*pointer = v
					pparent = pointer
					pointer = &v
				}
				continue
			}
			// Numeric index.
			valueNum, err := strconv.Atoi(array[i])
			if err != nil {
				err = gerror.WrapCodef(gcode.CodeInvalidParameter, err, `strconv.Atoi failed for string "%s"`, array[i])
				return err
			}

			if i == length-1 {
				// Leaf node.
				if len((*pointer).([]interface{})) > valueNum {
					if removed && value == nil {
						// Deleting element.
						if pparent == nil {
							*pointer = append((*pointer).([]interface{})[:valueNum], (*pointer).([]interface{})[valueNum+1:]...)
						} else {
							j.setPointerWithValue(pparent, array[i-1], append((*pointer).([]interface{})[:valueNum], (*pointer).([]interface{})[valueNum+1:]...))
						}
					} else {
						(*pointer).([]interface{})[valueNum] = value
					}
				} else {
					if removed && value == nil {
						goto done
					}
					if pparent == nil {
						// 这是根节点。
						j.setPointerWithValue(pointer, array[i], value)
					} else {
						// 它不是根节点。
						s := make([]interface{}, valueNum+1)
						copy(s, (*pointer).([]interface{}))
						s[valueNum] = value
						j.setPointerWithValue(pparent, array[i-1], s)
					}
				}
			} else {
				// Branch node.
				if gstr.IsNumeric(array[i+1]) {
					n, _ := strconv.Atoi(array[i+1])
					pSlice := (*pointer).([]interface{})
					if len(pSlice) > valueNum {
						item := pSlice[valueNum]
						if s, ok := item.([]interface{}); ok {
							for i := 0; i < n-len(s); i++ {
								s = append(s, nil)
							}
							pparent = pointer
							pointer = &pSlice[valueNum]
						} else {
							if removed && value == nil {
								goto done
							}
							var v interface{} = make([]interface{}, n+1)
							pparent = j.setPointerWithValue(pointer, array[i], v)
							pointer = &v
						}
					} else {
						if removed && value == nil {
							goto done
						}
						var v interface{} = make([]interface{}, n+1)
						pparent = j.setPointerWithValue(pointer, array[i], v)
						pointer = &v
					}
				} else {
					pSlice := (*pointer).([]interface{})
					if len(pSlice) > valueNum {
						pparent = pointer
						pointer = &(*pointer).([]interface{})[valueNum]
					} else {
						s := make([]interface{}, valueNum+1)
						copy(s, pSlice)
						s[valueNum] = make(map[string]interface{})
						if pparent != nil {
							// i > 0
							j.setPointerWithValue(pparent, array[i-1], s)
							pparent = pointer
							pointer = &s[valueNum]
						} else {
							// i = 0
							var v interface{} = s
							*pointer = v
							pparent = pointer
							pointer = &s[valueNum]
						}
					}
				}
			}

// 如果通过`pointer`指向的变量不是引用类型，
// 那么它将通过其父级（即：pparent）来修改该变量。
// 在Go语言中，这段注释描述了如果给定的指针`pointer`不指向一个引用类型，那么对变量的修改会通过其上级父级指针`pparent`间接完成。
		default:
			if removed && value == nil {
				goto done
			}
			if gstr.IsNumeric(array[i]) {
				n, _ := strconv.Atoi(array[i])
				s := make([]interface{}, n+1)
				if i == length-1 {
					s[n] = value
				}
				if pparent != nil {
					pparent = j.setPointerWithValue(pparent, array[i-1], s)
				} else {
					*pointer = s
					pparent = pointer
				}
			} else {
				var v1, v2 interface{}
				if i == length-1 {
					v1 = map[string]interface{}{
						array[i]: value,
					}
				} else {
					v1 = map[string]interface{}{
						array[i]: nil,
					}
				}
				if pparent != nil {
					pparent = j.setPointerWithValue(pparent, array[i-1], v1)
				} else {
					*pointer = v1
					pparent = pointer
				}
				v2 = v1.(map[string]interface{})[array[i]]
				pointer = &v2
			}
		}
	}
done:
	return nil
}

// convertValue 将 `value` 转换为 map[string]interface{} 或 []interface{}，
// 这样就可以支持对层级数据的访问。
func (j *Json) convertValue(value interface{}) (convertedValue interface{}, err error) {
	if value == nil {
		return
	}

	switch value.(type) {
	case map[string]interface{}:
		return value, nil

	case []interface{}:
		return value, nil

	default:
		var (
			reflectInfo = reflection.OriginValueAndKind(value)
		)
		switch reflectInfo.OriginKind {
		case reflect.Array:
			return gconv.Interfaces(value), nil

		case reflect.Slice:
			return gconv.Interfaces(value), nil

		case reflect.Map:
			return gconv.Map(value), nil

		case reflect.Struct:
			if v, ok := value.(iMapStrAny); ok {
				convertedValue = v.MapStrAny()
			}
			if utils.IsNil(convertedValue) {
				if v, ok := value.(iInterfaces); ok {
					convertedValue = v.Interfaces()
				}
			}
			if utils.IsNil(convertedValue) {
				convertedValue = gconv.Map(value)
			}
			if utils.IsNil(convertedValue) {
				err = gerror.NewCodef(gcode.CodeInvalidParameter, `unsupported value type "%s"`, reflect.TypeOf(value))
			}
			return

		default:
			return value, nil
		}
	}
}

// setPointerWithValue 将 `key`:`value` 设置到 `pointer` 中，其中 `key` 可能是 map 的键或 slice 的索引。
// 它返回指向新设置值的指针。
func (j *Json) setPointerWithValue(pointer *interface{}, key string, value interface{}) *interface{} {
	switch (*pointer).(type) {
	case map[string]interface{}:
		(*pointer).(map[string]interface{})[key] = value
		return &value
	case []interface{}:
		n, _ := strconv.Atoi(key)
		if len((*pointer).([]interface{})) > n {
			(*pointer).([]interface{})[n] = value
			return &(*pointer).([]interface{})[n]
		} else {
			s := make([]interface{}, n+1)
			copy(s, (*pointer).([]interface{}))
			s[n] = value
			*pointer = s
			return &s[n]
		}
	default:
		*pointer = value
	}
	return pointer
}

// getPointerByPattern 根据指定的`pattern`返回一个指向该值的指针。
func (j *Json) getPointerByPattern(pattern string) *interface{} {
	if j.p == nil {
		return nil
	}
	if j.vc {
		return j.getPointerByPatternWithViolenceCheck(pattern)
	} else {
		return j.getPointerByPatternWithoutViolenceCheck(pattern)
	}
}

// getPointerByPatternWithViolenceCheck 函数通过暴力检查的方式，返回指定 `pattern` 的值的指针。
func (j *Json) getPointerByPatternWithViolenceCheck(pattern string) *interface{} {
	if !j.vc {
		return j.getPointerByPatternWithoutViolenceCheck(pattern)
	}

	// 如果模式为空，则返回nil。
	if pattern == "" {
		return nil
	}
	// 如果模式为"."，则返回所有内容。
	if pattern == "." {
		return j.p
	}

	var (
		index   = len(pattern)
		start   = 0
		length  = 0
		pointer = j.p
	)
	if index == 0 {
		return pointer
	}
	for {
		if r := j.checkPatternByPointer(pattern[start:index], pointer); r != nil {
			if length += index - start; start > 0 {
				length += 1
			}
			start = index + 1
			index = len(pattern)
			if length == len(pattern) {
				return r
			} else {
				pointer = r
			}
		} else {
			// 获取下一个分隔符字符的位置。
			index = strings.LastIndexByte(pattern[start:index], j.c)
			if index != -1 && length > 0 {
				index += length + 1
			}
		}
		if start >= index {
			break
		}
	}
	return nil
}

// getPointerByPatternWithoutViolenceCheck 根据指定的`pattern`返回一个指向其值的指针，且不进行暴力检查。
func (j *Json) getPointerByPatternWithoutViolenceCheck(pattern string) *interface{} {
	if j.vc {
		return j.getPointerByPatternWithViolenceCheck(pattern)
	}

	// 如果模式为空，则返回nil。
	if pattern == "" {
		return nil
	}
	// 如果模式为"."，则返回所有内容。
	if pattern == "." {
		return j.p
	}

	pointer := j.p
	if len(pattern) == 0 {
		return pointer
	}
	array := strings.Split(pattern, string(j.c))
	for k, v := range array {
		if r := j.checkPatternByPointer(v, pointer); r != nil {
			if k == len(array)-1 {
				return r
			} else {
				pointer = r
			}
		} else {
			break
		}
	}
	return nil
}

// checkPatternByPointer 检查在指定的 `pointer` 中是否存在通过 `key` 访问的值。
// 它会返回该值的指针。
func (j *Json) checkPatternByPointer(key string, pointer *interface{}) *interface{} {
	switch (*pointer).(type) {
	case map[string]interface{}:
		if v, ok := (*pointer).(map[string]interface{})[key]; ok {
			return &v
		}
	case []interface{}:
		if gstr.IsNumeric(key) {
			n, err := strconv.Atoi(key)
			if err == nil && len((*pointer).([]interface{})) > n {
				return &(*pointer).([]interface{})[n]
			}
		}
	}
	return nil
}
