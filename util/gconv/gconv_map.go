// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 转换类

import (
	"reflect"
	"strings"
	
	"github.com/888go/goframe/internal/empty"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/utils"
)

type recursiveType string

const (
	recursiveTypeAuto recursiveType = "auto"
	recursiveTypeTrue recursiveType = "true"
)

// MapOption 定义了映射转换的选项。
type MapOption struct {
// Deep 标志表示递归地执行 Map 函数，这意味着如果给定转换值的属性也是一个结构体（struct/*struct），它会自动对该属性调用 Map 函数，将其转换为 map[string]interface{} 类型的变量。
	Deep bool

	// OmitEmpty 忽略具有 json `omitempty` 标签的属性。
	OmitEmpty bool

	// Tags 通过结构体标签名称指定转换后的映射键名称。
	Tags []string
}

// Map 将任意变量 `value` 转换为 map[string]interface{} 类型。如果参数 `value` 不是
// map、struct 或 *struct 类型，那么转换将会失败并返回 nil。
//
// 如果 `value` 是 struct 或 *struct 对象，第二个参数 `tags` 指定了优先检测的标签，
// 否则将按照 gconv、json、字段名的顺序进行检测。
func X取Map(值 interface{}, 选项 ...MapOption) map[string]interface{} {
	return doMapConvert(值, recursiveTypeAuto, false, 选项...)
}

// MapDeep 递归地执行 Map 函数，这意味着如果 `value` 的属性也是一个结构体（struct/*struct），则对该属性调用 Map 函数，并将其转换为 map[string]interface{} 类型的变量。
// 注意：已弃用，请改用 Map。
func X取Map_递归(值 interface{}, 值标签 ...string) map[string]interface{} {
	return doMapConvert(值, recursiveTypeTrue, false, MapOption{
		Tags: 值标签,
	})
}

// doMapConvert 实现了映射转换功能。
// 它会自动检查并转换，如果 `value` 为字符串或 []byte 类型，则将其转换为 map。
//
// TODO 完全实现对所有类型的递归转换，特别是对 map 的转换。
func doMapConvert(value interface{}, recursive recursiveType, mustMapReturn bool, option ...MapOption) map[string]interface{} {
	if value == nil {
		return nil
	}
	// 如果实现了接口 iVal，则重定向到其底层值。
	if v, ok := value.(iVal); ok {
		value = v.X取值()
	}

	var (
		usedOption = getUsedMapOption(option...)
		newTags    = StructTagPriority
	)
	switch len(usedOption.Tags) {
	case 0:
		// No need handling.
	case 1:
		newTags = append(strings.Split(usedOption.Tags[0], ","), StructTagPriority...)
	default:
		newTags = append(usedOption.Tags, StructTagPriority...)
	}
	// 断言常见类型的组合，并最终使用反射。
	dataMap := make(map[string]interface{})
	switch r := value.(type) {
	case string:
		// 如果这是一个JSON字符串，自动进行反序列化操作！
		if len(r) > 0 && r[0] == '{' && r[len(r)-1] == '}' {
			if err := json.UnmarshalUseNumber([]byte(r), &dataMap); err != nil {
				return nil
			}
		} else {
			return nil
		}
	case []byte:
		// 如果这是一个JSON字符串，自动进行反序列化操作！
		if len(r) > 0 && r[0] == '{' && r[len(r)-1] == '}' {
			if err := json.UnmarshalUseNumber(r, &dataMap); err != nil {
				return nil
			}
		} else {
			return nil
		}
	case map[interface{}]interface{}:
		recursiveOption := usedOption
		recursiveOption.Tags = newTags
		for k, v := range r {
			dataMap[String(k)] = doMapConvertForMapOrStructValue(
				doMapConvertForMapOrStructValueInput{
					IsRoot:          false,
					Value:           v,
					RecursiveType:   recursive,
					RecursiveOption: recursive == recursiveTypeTrue,
					Option:          recursiveOption,
				},
			)
		}
	case map[interface{}]string:
		for k, v := range r {
			dataMap[String(k)] = v
		}
	case map[interface{}]int:
		for k, v := range r {
			dataMap[String(k)] = v
		}
	case map[interface{}]uint:
		for k, v := range r {
			dataMap[String(k)] = v
		}
	case map[interface{}]float32:
		for k, v := range r {
			dataMap[String(k)] = v
		}
	case map[interface{}]float64:
		for k, v := range r {
			dataMap[String(k)] = v
		}
	case map[string]bool:
		for k, v := range r {
			dataMap[k] = v
		}
	case map[string]int:
		for k, v := range r {
			dataMap[k] = v
		}
	case map[string]uint:
		for k, v := range r {
			dataMap[k] = v
		}
	case map[string]float32:
		for k, v := range r {
			dataMap[k] = v
		}
	case map[string]float64:
		for k, v := range r {
			dataMap[k] = v
		}
	case map[string]string:
		for k, v := range r {
			dataMap[k] = v
		}
	case map[string]interface{}:
		if recursive == recursiveTypeTrue {
			recursiveOption := usedOption
			recursiveOption.Tags = newTags
			// 当前映射的一份副本。
			for k, v := range r {
				dataMap[k] = doMapConvertForMapOrStructValue(
					doMapConvertForMapOrStructValueInput{
						IsRoot:          false,
						Value:           v,
						RecursiveType:   recursive,
						RecursiveOption: recursive == recursiveTypeTrue,
						Option:          recursiveOption,
					},
				)
			}
		} else {
			// 它直接返回该映射，不作任何更改。
			return r
		}
	case map[int]interface{}:
		recursiveOption := usedOption
		recursiveOption.Tags = newTags
		for k, v := range r {
			dataMap[String(k)] = doMapConvertForMapOrStructValue(
				doMapConvertForMapOrStructValueInput{
					IsRoot:          false,
					Value:           v,
					RecursiveType:   recursive,
					RecursiveOption: recursive == recursiveTypeTrue,
					Option:          recursiveOption,
				},
			)
		}
	case map[int]string:
		for k, v := range r {
			dataMap[String(k)] = v
		}
	case map[uint]string:
		for k, v := range r {
			dataMap[String(k)] = v
		}

	default:
		// 如果不是常见类型，它将使用反射进行转换。
		var reflectValue reflect.Value
		if v, ok := value.(reflect.Value); ok {
			reflectValue = v
		} else {
			reflectValue = reflect.ValueOf(value)
		}
		reflectKind := reflectValue.Kind()
		// 如果是指针，我们应该找到它指向的实际数据类型。
		for reflectKind == reflect.Ptr {
			reflectValue = reflectValue.Elem()
			reflectKind = reflectValue.Kind()
		}
		switch reflectKind {
// 如果`value`的类型是数组，它会将偶数索引位置的值作为键，
// 奇数索引位置的值作为对应的值。例如：
// []string{"k1","v1","k2","v2"} 转换为 map[string]interface{}{"k1":"v1", "k2":"v2"}
// []string{"k1","v1","k2"}      转换为 map[string]interface{}{"k1":"v1", "k2":nil}
// 以下是中文注释：
// ```go
// 若变量`value`的类型为数组，该函数将数组中偶数下标的元素作为键，
// 奇数下标的元素作为对应的值进行转换，例如：
// []string{"k1","v1","k2","v2"} 将被转化为 map[string]interface{}{"k1":"v1", "k2":"v2"}
// []string{"k1","v1","k2"}      将被转化为 map[string]interface{}{"k1":"v1", "k2":nil}
		case reflect.Slice, reflect.Array:
			length := reflectValue.Len()
			for i := 0; i < length; i += 2 {
				if i+1 < length {
					dataMap[String(reflectValue.Index(i).Interface())] = reflectValue.Index(i + 1).Interface()
				} else {
					dataMap[String(reflectValue.Index(i).Interface())] = nil
				}
			}
		case reflect.Map, reflect.Struct, reflect.Interface:
			recursiveOption := usedOption
			recursiveOption.Tags = newTags
			convertedValue := doMapConvertForMapOrStructValue(
				doMapConvertForMapOrStructValueInput{
					IsRoot:          true,
					Value:           value,
					RecursiveType:   recursive,
					RecursiveOption: recursive == recursiveTypeTrue,
					Option:          recursiveOption,
					MustMapReturn:   mustMapReturn,
				},
			)
			if m, ok := convertedValue.(map[string]interface{}); ok {
				return m
			}
			return nil
		default:
			return nil
		}
	}
	return dataMap
}

func getUsedMapOption(option ...MapOption) MapOption {
	var usedOption MapOption
	if len(option) > 0 {
		usedOption = option[0]
	}
	return usedOption
}

type doMapConvertForMapOrStructValueInput struct {
	IsRoot          bool          // 如果当前不是根目录且不需要递归转换，则直接返回。
	Value           interface{}   // 当前操作值。
	RecursiveType   recursiveType // 从顶级函数入口处的类型。
	RecursiveOption bool          // 是否对`current`操作进行递归转换
	Option          MapOption     // 映射转换选项。
	MustMapReturn   bool          // 当为空时，必须返回 map 而不是 Value。
}

func doMapConvertForMapOrStructValue(in doMapConvertForMapOrStructValueInput) interface{} {
	if !in.IsRoot && !in.RecursiveOption {
		return in.Value
	}

	var reflectValue reflect.Value
	if v, ok := in.Value.(reflect.Value); ok {
		reflectValue = v
		in.Value = v.Interface()
	} else {
		reflectValue = reflect.ValueOf(in.Value)
	}
	reflectKind := reflectValue.Kind()
	// 如果是指针，我们应该找到它指向的实际数据类型。
	for reflectKind == reflect.Ptr {
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Kind()
	}
	switch reflectKind {
	case reflect.Map:
		var (
			mapKeys = reflectValue.MapKeys()
			dataMap = make(map[string]interface{})
		)
		for _, k := range mapKeys {
			var (
				mapKeyValue = reflectValue.MapIndex(k)
				mapValue    interface{}
			)
			switch {
			case mapKeyValue.IsZero():
				if mapKeyValue.IsNil() {
					// 快速检查空值。
					mapValue = nil
				} else {
// 在出现以下情况时：
// 恢复的异常：reflect: 对零值调用了reflect.Value.Interface
					mapValue = reflect.New(mapKeyValue.Type()).Elem().Interface()
				}
			default:
				mapValue = mapKeyValue.Interface()
			}
			dataMap[String(k.Interface())] = doMapConvertForMapOrStructValue(
				doMapConvertForMapOrStructValueInput{
					IsRoot:          false,
					Value:           mapValue,
					RecursiveType:   in.RecursiveType,
					RecursiveOption: in.RecursiveType == recursiveTypeTrue,
					Option:          in.Option,
				},
			)
		}
		return dataMap

	case reflect.Struct:
		var dataMap = make(map[string]interface{})
		// 接口检查映射转换
		if v, ok := in.Value.(iMapStrAny); ok {
			// 值复制，以保证并发安全。
			for mapK, mapV := range v.X取MapStrAny() {
				if in.RecursiveOption {
					dataMap[mapK] = doMapConvertForMapOrStructValue(
						doMapConvertForMapOrStructValueInput{
							IsRoot:          false,
							Value:           mapV,
							RecursiveType:   in.RecursiveType,
							RecursiveOption: in.RecursiveType == recursiveTypeTrue,
							Option:          in.Option,
						},
					)
				} else {
					dataMap[mapK] = mapV
				}
			}
			if len(dataMap) > 0 {
				return dataMap
			}
		}
		// 使用reflect进行转换
		var (
			rtField     reflect.StructField
			rvField     reflect.Value
			reflectType = reflectValue.Type() // 属性值类型
			mapKey      = ""                  // mapKey 可能是标签名称或结构体属性名称。
		)
		for i := 0; i < reflectValue.NumField(); i++ {
			rtField = reflectType.Field(i)
			rvField = reflectValue.Field(i)
			// 只转换公共属性。
			fieldName := rtField.Name
			if !utils.IsLetterUpper(fieldName[0]) {
				continue
			}
			mapKey = ""
			fieldTag := rtField.Tag
			for _, tag := range in.Option.Tags {
				if mapKey = fieldTag.Get(tag); mapKey != "" {
					break
				}
			}
			if mapKey == "" {
				mapKey = fieldName
			} else {
				// 支持json标签特性：-, omitempty
// 这行注释说明了Go语言中对于结构体字段的JSON标签功能支持两种特殊标记：
// `-`：表示忽略该字段，即在进行JSON编码（Marshal）时，不会将该字段包含到生成的JSON数据中。
// `omitempty`：表示当该字段值为空（如零值、空字符串、长度为0的数组/切片/映射等）时，在进行JSON编码时不包含此字段。
				mapKey = strings.TrimSpace(mapKey)
				if mapKey == "-" {
					continue
				}
				array := strings.Split(mapKey, ",")
				if len(array) > 1 {
					switch strings.TrimSpace(array[1]) {
					case "omitempty":
						if in.Option.OmitEmpty && empty.IsEmpty(rvField.Interface()) {
							continue
						} else {
							mapKey = strings.TrimSpace(array[0])
						}
					default:
						mapKey = strings.TrimSpace(array[0])
					}
				}
				if mapKey == "" {
					mapKey = fieldName
				}
			}
			if in.RecursiveOption || rtField.Anonymous {
				// 递归地执行映射转换。
				var (
					rvAttrField = rvField
					rvAttrKind  = rvField.Kind()
				)
				if rvAttrKind == reflect.Ptr {
					rvAttrField = rvField.Elem()
					rvAttrKind = rvAttrField.Kind()
				}
				switch rvAttrKind {
				case reflect.Struct:
// 嵌入式结构体，无字段，仅忽略它。
// 例如：gmeta.Meta
					if rvAttrField.Type().NumField() == 0 {
						continue
					}
					var (
						hasNoTag = mapKey == fieldName
// **注意**：在此处不要使用rvAttrField.Interface()，
// 因为它可能从指针变为结构体。
						rvInterface = rvField.Interface()
					)
					switch {
					case hasNoTag && rtField.Anonymous:
// 这意味着该属性字段没有标签。
// 用子结构体的属性字段覆盖该属性。
// 这段Go语言代码注释翻译成中文为：
// ```go
// 这表示该属性字段未设置标签。
// 使用子结构体的属性字段来重写该属性。
						anonymousValue := doMapConvertForMapOrStructValue(doMapConvertForMapOrStructValueInput{
							IsRoot:          false,
							Value:           rvInterface,
							RecursiveType:   in.RecursiveType,
							RecursiveOption: true,
							Option:          in.Option,
						})
						if m, ok := anonymousValue.(map[string]interface{}); ok {
							for k, v := range m {
								dataMap[k] = v
							}
						} else {
							dataMap[mapKey] = rvInterface
						}

					// 这意味着该属性字段具有期望的标签。
					case !hasNoTag && rtField.Anonymous:
						dataMap[mapKey] = doMapConvertForMapOrStructValue(doMapConvertForMapOrStructValueInput{
							IsRoot:          false,
							Value:           rvInterface,
							RecursiveType:   in.RecursiveType,
							RecursiveOption: true,
							Option:          in.Option,
						})

					default:
						dataMap[mapKey] = doMapConvertForMapOrStructValue(doMapConvertForMapOrStructValueInput{
							IsRoot:          false,
							Value:           rvInterface,
							RecursiveType:   in.RecursiveType,
							RecursiveOption: in.RecursiveType == recursiveTypeTrue,
							Option:          in.Option,
						})
					}

				// 结构体属性的类型为切片。
				case reflect.Array, reflect.Slice:
					length := rvAttrField.Len()
					if length == 0 {
						dataMap[mapKey] = rvAttrField.Interface()
						break
					}
					array := make([]interface{}, length)
					for arrayIndex := 0; arrayIndex < length; arrayIndex++ {
						array[arrayIndex] = doMapConvertForMapOrStructValue(
							doMapConvertForMapOrStructValueInput{
								IsRoot:          false,
								Value:           rvAttrField.Index(arrayIndex).Interface(),
								RecursiveType:   in.RecursiveType,
								RecursiveOption: in.RecursiveType == recursiveTypeTrue,
								Option:          in.Option,
							},
						)
					}
					dataMap[mapKey] = array
				case reflect.Map:
					var (
						mapKeys   = rvAttrField.MapKeys()
						nestedMap = make(map[string]interface{})
					)
					for _, k := range mapKeys {
						nestedMap[String(k.Interface())] = doMapConvertForMapOrStructValue(
							doMapConvertForMapOrStructValueInput{
								IsRoot:          false,
								Value:           rvAttrField.MapIndex(k).Interface(),
								RecursiveType:   in.RecursiveType,
								RecursiveOption: in.RecursiveType == recursiveTypeTrue,
								Option:          in.Option,
							},
						)
					}
					dataMap[mapKey] = nestedMap
				default:
					if rvField.IsValid() {
						dataMap[mapKey] = reflectValue.Field(i).Interface()
					} else {
						dataMap[mapKey] = nil
					}
				}
			} else {
				// 不进行递归的映射值转换
				if rvField.IsValid() {
					dataMap[mapKey] = reflectValue.Field(i).Interface()
				} else {
					dataMap[mapKey] = nil
				}
			}
		}
		if !in.MustMapReturn && len(dataMap) == 0 {
			return in.Value
		}
		return dataMap

	// 给定的值是切片类型。
	case reflect.Array, reflect.Slice:
		length := reflectValue.Len()
		if length == 0 {
			break
		}
		array := make([]interface{}, reflectValue.Len())
		for i := 0; i < length; i++ {
			array[i] = doMapConvertForMapOrStructValue(doMapConvertForMapOrStructValueInput{
				IsRoot:          false,
				Value:           reflectValue.Index(i).Interface(),
				RecursiveType:   in.RecursiveType,
				RecursiveOption: in.RecursiveType == recursiveTypeTrue,
				Option:          in.Option,
			})
		}
		return array
	}
	return in.Value
}

// MapStrStr 将`value`转换为map[string]string类型。
// 注意，这种映射类型的转换可能会导致数据复制。
func X取文本Map(值 interface{}, 选项 ...MapOption) map[string]string {
	if r, ok := 值.(map[string]string); ok {
		return r
	}
	m := X取Map(值, 选项...)
	if len(m) > 0 {
		vMap := make(map[string]string, len(m))
		for k, v := range m {
			vMap[k] = String(v)
		}
		return vMap
	}
	return nil
}

// MapStrStrDeep递归地将`value`转换为map[string]string类型。
// 注意，这种映射类型转换可能会涉及数据复制。
// 废弃: 请改用MapStrStr。
func X取文本Map_递归(值 interface{}, 值标签 ...string) map[string]string {
	if r, ok := 值.(map[string]string); ok {
		return r
	}
	m := X取Map_递归(值, 值标签...)
	if len(m) > 0 {
		vMap := make(map[string]string, len(m))
		for k, v := range m {
			vMap[k] = String(v)
		}
		return vMap
	}
	return nil
}
