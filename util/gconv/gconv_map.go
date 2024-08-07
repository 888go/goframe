// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 转换类

import (
	"reflect"
	"strings"

	"github.com/888go/goframe/internal/empty"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/utils"
	"github.com/888go/goframe/util/gtag"
)

type recursiveType string

const (
	recursiveTypeAuto recursiveType = "auto"
	recursiveTypeTrue recursiveType = "true"
)

// MapOption 定义了映射转换的选项。 md5:8dc53d6fdc486bf8
type MapOption struct {
	// Deep 标记表示递归地执行 Map 函数，这意味着如果给定转换值的属性也是一个结构体（struct），它会自动对这个属性调用 Map 函数，将其转换为 map[string]interface{} 类型变量。
	// md5:3653359965fb222d
	Deep bool

		// OmitEmpty 忽略具有 json `omitempty` 标签的属性。 md5:ce80b66cfe17a0ba
	OmitEmpty bool

		// Tags 指定了通过结构体标签名转换后的映射键名称。 md5:b08e40ad043d7120
	Tags []string
}

// X取Map 将任何变量 `value` 转换为 map[string]interface{}。如果参数 `value` 不是
// 类型为 map/struct/*struct，转换将会失败并返回 nil。
//
// 如果 `value` 是一个 struct/*struct 对象，第二个参数 `tags` 指定了具有最高优先级的
// 将被检测的标签，否则它会按照以下顺序检测标签：
// gconv, json, 字段名称。
// md5:34498665a6393f82
func X取Map(值 interface{}, 选项 ...MapOption) map[string]interface{} {
	return doMapConvert(值, recursiveTypeAuto, false, 选项...)
}

// X取Map_递归递归地执行Map函数，这意味着如果`value`的属性也是一个`struct/*struct`，则会在这个属性上调用Map函数，并将其转换为map[string]interface{}类型的变量。
// 警告：建议使用Map替代。
// md5:dc0620a4d15b4389
func X取Map_递归(值 interface{}, 值标签 ...string) map[string]interface{} {
	return doMapConvert(值, recursiveTypeTrue, false, MapOption{
		Deep: true,
		Tags: 值标签,
	})
}

// doMapConvert 实现了映射转换。
// 如果 `value` 是字符串或[]byte，它会自动检查并将其转换为map。
// 
// TODO 完全实现所有类型的递归转换，特别是map。
// md5:f55eadf34b47fad4
func doMapConvert(value interface{}, recursive recursiveType, mustMapReturn bool, option ...MapOption) map[string]interface{} {
	if value == nil {
		return nil
	}
		// 如果它已经实现了iVal接口，那么它会重定向到其底层值。 md5:fb13fb87762a52a2
	if v, ok := value.(iVal); ok {
		value = v.X取值()
	}

	var (
		usedOption = getUsedMapOption(option...)
		newTags    = gtag.StructTagPriority
	)
	if usedOption.Deep {
		recursive = recursiveTypeTrue
	}
	switch len(usedOption.Tags) {
	case 0:
		// No need handling.
	case 1:
		newTags = append(strings.Split(usedOption.Tags[0], ","), gtag.StructTagPriority...)
	default:
		newTags = append(usedOption.Tags, gtag.StructTagPriority...)
	}
		// 断言常见的类型组合，并最终使用反射。 md5:28d02793b273a6c1
	dataMap := make(map[string]interface{})
	switch r := value.(type) {
	case string:
						// 如果它是一个JSON字符串，自动反序列化它！. md5:2da2afc6ee11f379
		if len(r) > 0 && r[0] == '{' && r[len(r)-1] == '}' {
			if err := json.UnmarshalUseNumber([]byte(r), &dataMap); err != nil {
				return nil
			}
		} else {
			return nil
		}
	case []byte:
						// 如果它是一个JSON字符串，自动反序列化它！. md5:2da2afc6ee11f379
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
			// A copy of current map.
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
						// 它直接返回映射，不做任何更改。 md5:fa0b6b4232562112
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
				// 并非常见类型，因此它使用反射来进行转换。 md5:a4126e9dfe7a56bd
		var reflectValue reflect.Value
		if v, ok := value.(reflect.Value); ok {
			reflectValue = v
		} else {
			reflectValue = reflect.ValueOf(value)
		}
		reflectKind := reflectValue.Kind()
			// 如果它是一个指针，我们应该找到其实际的数据类型。 md5:db4733e40015c40e
		for reflectKind == reflect.Ptr {
			reflectValue = reflectValue.Elem()
			reflectKind = reflectValue.Kind()
		}
		switch reflectKind {
		// 如果`value`是数组类型，它将偶数索引的值作为键，奇数索引的值作为对应的值。例如：
		// `[]string{"k1","v1","k2","v2"}` => `map[string]interface{}{"k1":"v1", "k2":"v2"}`
		// `[]string{"k1","v1","k2"}`       => `map[string]interface{}{"k1":"v1", "k2":nil}`
		// md5:5e90ff5bc08f2638
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
	IsRoot          bool          // 如果它不是根并且没有递归转换，它会直接返回。 md5:70679f33f48f5a89
	Value           interface{}   // 当前操作值。 md5:2dcb5cbb4a76dbe7
	RecursiveType   recursiveType // 从顶级函数入口的类型。 md5:6fd96f3dbc57d815
	RecursiveOption bool          // 是否为当前操作`current`进行递归转换。 md5:d915897a37c59c4a
	Option          MapOption     // Map converting option.
	MustMapReturn   bool          // 当空时，必须返回map而不是Value。 md5:e49001a917ef93fb
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
		// 如果它是一个指针，我们应该找到其实际的数据类型。 md5:db4733e40015c40e
	for reflectKind == reflect.Ptr {
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Kind()
	}
	switch reflectKind {
	case reflect.Map:
		var (
			mapIter = reflectValue.MapRange()
			dataMap = make(map[string]interface{})
		)
		for mapIter.Next() {
			var (
				mapKeyValue = mapIter.Value()
				mapValue    interface{}
			)
			switch {
			case mapKeyValue.IsZero():
				if utils.CanCallIsNil(mapKeyValue) && mapKeyValue.IsNil() {
										// 快速检查值是否为nil。 md5:93138802a95bcbf7
					mapValue = nil
				} else {
					// 在出现以下情况时：
					// 异常恢复：reflect: 对零值调用reflect.Value.Interface
					// md5:e32f0249911d4dde
					mapValue = reflect.New(mapKeyValue.Type()).Elem().Interface()
				}
			default:
				mapValue = mapKeyValue.Interface()
			}
			dataMap[String(mapIter.Key().Interface())] = doMapConvertForMapOrStructValue(
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
				// 转换接口检查的映射。 md5:e4adcda9bbeec1fc
		if v, ok := in.Value.(iMapStrAny); ok {
						// 为了并发安全，进行值复制。 md5:57f6f9976b1be5ca
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
				// 使用反射进行转换。 md5:ffcf9b71ef0563af
		var (
			rtField     reflect.StructField
			rvField     reflect.Value
			reflectType = reflectValue.Type() // attribute value type.
			mapKey      = ""                  // mapKey 可能是标签名或结构体属性名。 md5:7798f495f1f4211d
		)
		for i := 0; i < reflectValue.NumField(); i++ {
			rtField = reflectType.Field(i)
			rvField = reflectValue.Field(i)
						// 只转换公共属性。 md5:090d3eafbff3ac6e
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
												// 支持json标签特性：-，omitempty. md5:89511416feac7bb4
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
								// 递归地进行映射转换。 md5:1676b5bed955fd64
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
					// 内嵌结构体且没有字段，将忽略它。
					// 例如：gmeta.Meta
					// md5:8505cb87a6269724
					if rvAttrField.Type().NumField() == 0 {
						continue
					}
					var (
						hasNoTag = mapKey == fieldName
						// 不要在这里使用rvAttrField.Interface()，因为它可能会从指针转换为结构体。
						// md5:5cd6517f328dfd1c
						rvInterface = rvField.Interface()
					)
					switch {
					case hasNoTag && rtField.Anonymous:
						// 这意味着这个属性字段没有标签。
						// 使用子结构体的属性字段覆盖该属性。
						// md5:525f64e84a599d2d
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

										// 这意味着该属性字段具有期望的标签。 md5:e6252ec8be3f90cb
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

								// 该结构体属性是切片类型。 md5:e1a646d8191abc2f
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
						mapIter   = rvAttrField.MapRange()
						nestedMap = make(map[string]interface{})
					)
					for mapIter.Next() {
						nestedMap[String(mapIter.Key().Interface())] = doMapConvertForMapOrStructValue(
							doMapConvertForMapOrStructValueInput{
								IsRoot:          false,
								Value:           mapIter.Value().Interface(),
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
								// 不进行递归地将映射值转换. md5:fd213a1b3835dd97
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

		// 给定的值是切片类型。 md5:fb5a502257cf9a01
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

// X取文本Map 将 `value` 转换为 map[string]string 类型。
// 注意，对于这种映射类型转换，可能会有数据复制的情况发生。
// md5:a1ec9ce0d856cd1e
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

// X取文本Map_递归 递归地将`value`转换为map[string]string。
// 请注意，这种映射类型的转换可能会涉及数据复制。
// 已弃用：请使用MapStrStr代替。
// md5:79528a85e8ff4c82
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
