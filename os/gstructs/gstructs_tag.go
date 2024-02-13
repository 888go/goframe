// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gstructs

import (
	"reflect"
	"strconv"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/util/gtag"
)

// ParseTag函数用于解析标签字符串并转换为映射（map）。
// 例如：
// ParseTag(`v:"required" p:"id" d:"1"`) => map[v:required p:id d:1].
func ParseTag(tag string) map[string]string {
	var (
		key  string
		data = make(map[string]string)
	)
	for tag != "" {
		// 跳过前面的空格。
		i := 0
		for i < len(tag) && tag[i] == ' ' {
			i++
		}
		tag = tag[i:]
		if tag == "" {
			break
		}
// 扫描到冒号。空格、引号或控制字符都是语法错误。
// 严格来讲，控制字符包括范围 [0x7f, 0x9f]，而不只是 [0x00, 0x1f]，
// 但在实际操作中，我们忽略了多字节的控制字符，
// 因为检查标签字节比检查标签符文更为简单。
		i = 0
		for i < len(tag) && tag[i] > ' ' && tag[i] != ':' && tag[i] != '"' && tag[i] != 0x7f {
			i++
		}
		if i == 0 || i+1 >= len(tag) || tag[i] != ':' || tag[i+1] != '"' {
			break
		}
		key = tag[:i]
		tag = tag[i+1:]

		// 扫描带引号的字符串以查找值。
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
			panic(错误类.X多层错误码并格式化(错误码类.CodeInvalidParameter, err, `error parsing tag "%s"`, tag))
		}
		data[key] = gtag.Parse(value)
	}
	return data
}

// TagFields 从`pointer`中获取并返回以[]Field形式的结构体标签。
//
// 参数`pointer`应为struct或*struct类型。
//
// 注意：
// 1. 它仅从结构体中获取首字母大写的导出属性（即公开字段）。
// 2. 参数`priority`应提供，它只检索具有给定标签的字段。
func TagFields(pointer interface{}, priority []string) ([]Field, error) {
	return getFieldValuesByTagPriority(pointer, priority, map[string]struct{}{})
}

// TagMapName 从`pointer`中获取并返回以map[tag]attribute形式的结构体标签。
//
// 参数`pointer`应为struct或*struct类型。
//
// 注意：
// 1. 它只从结构体中获取首字母大写的导出属性。
// 2. 应提供参数`priority`，它只检索具有给定标签的字段。
// 3. 如果某个字段没有指定标签，则使用其字段名作为结果映射键。
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

// TagMapField 从`pointer`中获取结构体标签并以map[tag]Field的形式返回。参数`object`应为struct/*struct/[]struct/[]*struct类型。
//
// 注意：
// 1. 它只检索结构体中首字母大写的导出属性（即公开字段）。
// 2. 参数`priority`必须给出，它只检索具有该给定标签的字段。
// 3. 如果某个字段没有指定标签，则使用其字段名称作为结果映射键。
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
				// 如果指针是结构体类型且为nil，则自动创建一个临时结构体。
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
		return nil, 错误类.X创建错误码(
			错误码类.CodeInvalidParameter,
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
		// 仅获取导出的属性。
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
			// 过滤重复的标签。
			if _, ok := repeatedTagFilteringMap[tagValue]; ok {
				continue
			}
			tagField := field
			tagField.TagName = tagName
			tagField.TagValue = tagValue
			tagFields = append(tagFields, tagField)
		}
		// 如果这是一个嵌入式属性，它会递归地获取标签。
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
