// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

// 包gjson提供了处理JSON/XML/INI/YAML/TOML数据的便捷API。 md5:ddbf6ad5d309a49c
package gjson

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/reflection"
	"github.com/gogf/gf/v2/internal/rwmutex"
	"github.com/gogf/gf/v2/internal/utils"
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
	defaultSplitChar = '.' // 用于层次数据访问的分隔符字符。 md5:3020966f087a4732
)

// Json 是自定义的JSON结构体。 md5:764b883e7cf79da7
type Json struct {
	mu rwmutex.RWMutex
	p  *interface{} // Pointer for hierarchical data access, it's the root of data in default.
	c  byte         // 字符分隔符（默认为'.'）。 md5:15307025b7ed9ae7
	vc bool         // 暴力检查（默认为false），用于在层次数据键包含分隔符字符时访问数据。 md5:465e099ccbdc4ca3
}

// 创建/加载Json对象的选项。 md5:d8614ea5dc358e89
type Options struct {
	Safe      bool        // 标记此对象适用于并发安全使用。这尤其适用于 Json 对象的创建。 md5:59d439559fecdc34
	Tags      string      // 自定义解码优先级标签，例如："json,yaml,MyTag"。这主要用于将结构体解析为Json对象。 md5:486cf257ddd06463
	Type      ContentType // 类型指定了数据内容类型，例如：json、xml、yaml、toml、ini等。 md5:afbae78560edde30
	StrNumber bool        // StrNumber 使得 Decoder 将数字解码为 interface{}` 作为字符串，而不是 float64。 md5:32e44e32c3cc37cc
}

// iInterfaces 用于类型断言接口，用于 Interfaces() 方法。 md5:711dc755f9cd4979
type iInterfaces interface {
	Interfaces() []interface{}
}

// iMapStrAny 是一个接口，支持将结构体参数转换为映射。 md5:cfd4642c77fca6ec
type iMapStrAny interface {
	MapStrAny() map[string]interface{}
}

// iVal是用于获取底层interface{}的接口。 md5:2915e3bd3d7e4f43
type iVal interface {
	Val() interface{}
}

// setValue 将`value`设置为`j`，按照`pattern`。
// 注意：
// 1. 如果`value`为nil且`removed`为true，表示删除这个值；
// 2. 在层次数据搜索、节点创建和数据赋值方面相当复杂。 md5:6aca091405b9da40
func (j *Json) setValue(pattern string, value interface{}, removed bool) error {
	var (
		err    error
		array  = strings.Split(pattern, string(j.c))
		length = len(array)
	)
	if value, err = j.convertValue(value); err != nil {
		return err
	}
	// Initialization checks.
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
					// Delete item from map.
					delete((*pointer).(map[string]interface{}), array[i])
				} else {
					if (*pointer).(map[string]interface{}) == nil {
						*pointer = map[string]interface{}{}
					}
					(*pointer).(map[string]interface{})[array[i]] = value
				}
			} else {
				// 如果键在映射中不存在。 md5:ba2af475e1347525
				if v, ok := (*pointer).(map[string]interface{})[array[i]]; !ok {
					if removed && value == nil {
						goto done
					}
					// Creating new node.
					if gstr.IsNumeric(array[i+1]) {
						// Creating array node.
						n, _ := strconv.Atoi(array[i+1])
						var v interface{} = make([]interface{}, n+1)
						pparent = j.setPointerWithValue(pointer, array[i], v)
						pointer = &v
					} else {
						// Creating map node.
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
						// It is the root node.
						j.setPointerWithValue(pointer, array[i], value)
					} else {
						// 它不是根节点。 md5:b90762478c5a92c6
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

		// 如果`pointer`指向的变量不是引用类型，那么它会通过其父对象（pparent）来修改该变量。 md5:aa59525c846686ce
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

// convertValue将"value"转换为map[string]interface{}或[]interface{}，这样可以支持层级数据访问。 md5:089e6e9291ed7aab
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

// setPointerWithValue 将 `key`:`value` 设置到 `pointer` 中，其中 `key` 可能是映射的键或切片的索引。
// 它返回新设置值的指针。 md5:2642aca0fd23f46c
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

// getPointerByPattern 根据指定的 `pattern` 返回值的指针。 md5:e5422879dc2c9285
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

// getPointerByPatternWithViolenceCheck 通过暴力检查返回指定`pattern`的值的指针。 md5:4ac204b4633753dc
func (j *Json) getPointerByPatternWithViolenceCheck(pattern string) *interface{} {
	if !j.vc {
		return j.getPointerByPatternWithoutViolenceCheck(pattern)
	}

	// 如果pattern为空，它将返回nil。 md5:8e2a6f56affd353a
	if pattern == "" {
		return nil
	}
	// 如果pattern是"."，则返回所有。 md5:1f0d65d517f332bd
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
			// 获取下一个分隔符字符的位置。 md5:7268bb1b6598460b
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

// getPointerByPatternWithoutViolenceCheck 返回指定`pattern`值的指针，不进行暴力检查。 md5:fd58f2cfd08f8751
func (j *Json) getPointerByPatternWithoutViolenceCheck(pattern string) *interface{} {
	if j.vc {
		return j.getPointerByPatternWithViolenceCheck(pattern)
	}

	// 如果pattern为空，它将返回nil。 md5:8e2a6f56affd353a
	if pattern == "" {
		return nil
	}
	// 如果pattern是"."，则返回所有。 md5:1f0d65d517f332bd
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

// checkPatternByPointer 检查指定`pointer`中是否存在键为`key`的值。它返回该值的指针。 md5:10f17307c0c6e052
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
