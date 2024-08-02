// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gstructs

import (
	"reflect"
	"strconv"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/util/gtag"
)

// ParseTag 将标签字符串解析为映射。
// 例如：
// ParseTag(`v:"required" p:"id" d:"1"`)) => map[v:required p:id d:1]。
// md5:967d381052c3a2d8
func ParseTag(tag string) map[string]string {
	var (
		key  string
		data = make(map[string]string)
	)
	for tag != "" {
		// Skip leading space.
		i := 0
		for i < len(tag) && tag[i] == ' ' {
			i++
		}
		tag = tag[i:]
		if tag == "" {
			break
		}
		// 扫描到冒号。空格、引号或控制字符都是语法错误。
		// 严格来说，控制字符包括范围 [0x7f, 0x9f]，而不仅仅是 [0x00, 0x1f]。但在实践中，我们忽略多字节控制字符，因为检查标签的字节比检查标签的 rune 更简单。
		// md5:2b37f6b6cf4e8415
		i = 0
		for i < len(tag) && tag[i] > ' ' && tag[i] != ':' && tag[i] != '"' && tag[i] != 0x7f {
			i++
		}
		if i == 0 || i+1 >= len(tag) || tag[i] != ':' || tag[i+1] != '"' {
			break
		}
		key = tag[:i]
		tag = tag[i+1:]

				// 扫描带引号的字符串以找到值。 md5:022e03f120cb2054
		i = 1
		for i < len(tag) && tag[i] != '"' {
			if tag[i] == '\\' {
				i++
			}
			i++
		}
		if i >= len(tag) {
			break
		}
		quotedValue := tag[:i+1]
		tag = tag[i+1:]
		value, err := strconv.Unquote(quotedValue)
		if err != nil {
			panic(gerror.WrapCodef(gcode.CodeInvalidParameter, err, `error parsing tag "%s"`, tag))
		}
		data[key] = gtag.Parse(value)
	}
	return data
}

// TagFields 从`pointer`获取并返回结构体标签作为[]Field。
//
// 参数`pointer`应为struct/*struct类型。
//
// 请注意：
// 1. 它只从结构体中检索首字母大写的导出属性。
// 2. 应提供参数`priority`，它只检索具有给定标签的字段。
// md5:55390bfc1f5537f2
func TagFields(pointer interface{}, priority []string) ([]Field, error) {
	return getFieldValuesByTagPriority(pointer, priority, map[string]struct{}{})
}

// TagMapName从`pointer`获取并返回结构体标签作为map[tag]attribute。
// 
// 参数`pointer`应为结构体或*struct类型。
// 
// 注意：
// 1. 它仅从结构体中检索首字母大写的导出属性。
// 2. 需要提供参数`priority`，它只检索具有给定标签的字段。
// 3. 如果一个字段没有指定标签，它将使用其字段名称作为结果映射的键。
// md5:0eb7c62c8a6f7e09
func TagMapName(pointer interface{}, priority []string) (map[string]string, error) {
	fields, err := TagFields(pointer, priority)
	if err != nil {
		return nil, err
	}
	tagMap := make(map[string]string, len(fields))
	for _, field := range fields {
		tagMap[field.TagValue] = field.Name()
	}
	return tagMap, nil
}

// TagMapField 从 `pointer` 中获取结构体标签作为 map[tag]Field，然后返回它。
// 参数 `object` 应该是 struct 类型、*struct 类型、struct 切片或 []*struct 类型之一。
// 
// 注意：
// 1. 它只会从结构体中检索首字母大写的导出属性。
// 2. 需要提供参数 `priority`，只检索具有给定标签的字段。
// 3. 如果一个字段没有指定标签，它将使用其字段名称作为结果映射的键。
// md5:ba865b4214b27332
func TagMapField(object interface{}, priority []string) (map[string]Field, error) {
	fields, err := TagFields(object, priority)
	if err != nil {
		return nil, err
	}
	tagMap := make(map[string]Field, len(fields))
	for _, field := range fields {
		tagField := field
		tagMap[field.TagValue] = tagField
	}
	return tagMap, nil
}

func getFieldValues(structObject interface{}) ([]Field, error) {
	var (
		reflectValue reflect.Value
		reflectKind  reflect.Kind
	)
	if v, ok := structObject.(reflect.Value); ok {
		reflectValue = v
		reflectKind = reflectValue.Kind()
	} else {
		reflectValue = reflect.ValueOf(structObject)
		reflectKind = reflectValue.Kind()
	}
	for {
		switch reflectKind {
		case reflect.Ptr:
			if !reflectValue.IsValid() || reflectValue.IsNil() {
								// 如果指针是*struct类型且为nil，那么会自动创建一个临时的struct。 md5:23b5ebc131739e7d
				reflectValue = reflect.New(reflectValue.Type().Elem()).Elem()
				reflectKind = reflectValue.Kind()
			} else {
				reflectValue = reflectValue.Elem()
				reflectKind = reflectValue.Kind()
			}
		case reflect.Array, reflect.Slice:
			reflectValue = reflect.New(reflectValue.Type().Elem()).Elem()
			reflectKind = reflectValue.Kind()
		default:
			goto exitLoop
		}
	}

exitLoop:
	for reflectKind == reflect.Ptr {
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Kind()
	}
	if reflectKind != reflect.Struct {
		return nil, gerror.NewCode(
			gcode.CodeInvalidParameter,
			"given value should be either type of struct/*struct/[]struct/[]*struct",
		)
	}
	var (
		structType = reflectValue.Type()
		length     = reflectValue.NumField()
		fields     = make([]Field, length)
	)
	for i := 0; i < length; i++ {
		fields[i] = Field{
			Value: reflectValue.Field(i),
			Field: structType.Field(i),
		}
	}
	return fields, nil
}

func getFieldValuesByTagPriority(
	pointer interface{}, priority []string, repeatedTagFilteringMap map[string]struct{},
) ([]Field, error) {
	fields, err := getFieldValues(pointer)
	if err != nil {
		return nil, err
	}
	var (
		tagName   string
		tagValue  string
		tagFields = make([]Field, 0)
	)
	for _, field := range fields {
				// 只检索导出的属性。 md5:d8185f07060feffb
		if !field.IsExported() {
			continue
		}
		tagValue = ""
		for _, p := range priority {
			tagName = p
			tagValue = field.Tag(p)
			if tagValue != "" && tagValue != "-" {
				break
			}
		}
		if tagValue != "" {
			// Filter repeated tag.
			if _, ok := repeatedTagFilteringMap[tagValue]; ok {
				continue
			}
			tagField := field
			tagField.TagName = tagName
			tagField.TagValue = tagValue
			tagFields = append(tagFields, tagField)
		}
				// 如果这是一个嵌入属性，它将递归地获取标签。 md5:ed1233074f938682
		if field.IsEmbedded() && field.OriginalKind() == reflect.Struct {
			subTagFields, err := getFieldValuesByTagPriority(field.Value, priority, repeatedTagFilteringMap)
			if err != nil {
				return nil, err
			} else {
				tagFields = append(tagFields, subTagFields...)
			}
		}
	}
	return tagFields, nil
}
