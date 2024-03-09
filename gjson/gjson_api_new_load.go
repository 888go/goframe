// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package json类

import (
	"bytes"
	"reflect"
	
	"github.com/gogf/gf/v2/encoding/gini"
	"github.com/gogf/gf/v2/encoding/gproperties"
	"github.com/gogf/gf/v2/encoding/gtoml"
	"github.com/gogf/gf/v2/encoding/gxml"
	"github.com/gogf/gf/v2/encoding/gyaml"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/888go/goframe/gjson/internal/json"
	"github.com/888go/goframe/gjson/internal/reflection"
	"github.com/888go/goframe/gjson/internal/rwmutex"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// New 函数通过任意类型的 `data` 创建一个 Json 对象，但 `data` 应为 map 或 slice 类型以保证数据可访问性，否则创建此对象将无实际意义。
//
// 参数 `safe` 指定是否在并发安全的上下文中使用此 Json 对象，默认情况下 `safe` 为 false。
func X创建(值 interface{}, 并发安全 ...bool) *Json {
	return X创建并按类型标签(值, string(ContentTypeJson), 并发安全...)
}

// NewWithTag 创建一个Json对象，其数据类型可以是任意的 `data`，但为了方便数据访问，`data` 应该是一个 map 或 slice，否则将失去意义。
//
// 参数 `tags` 指定了在结构体转为 map 时使用的优先级标签，多个标签使用字符 ',' 连接。
//
// 参数 `safe` 指定了是否在并发安全的上下文中使用此 Json 对象，默认情况下为 false。
func X创建并按类型标签(值 interface{}, 类型标签 string, 并发安全 ...bool) *Json {
	option := Options{
		Tags: 类型标签,
	}
	if len(并发安全) > 0 && 并发安全[0] {
		option.Safe = true
	}
	return X创建并按选项(值, option)
}

// NewWithOptions 创建一个Json对象，其变量类型可以是 `data` 的任意类型，但为了能够访问数据，`data` 应该是一个 map 或 slice，否则将毫无意义。
func X创建并按选项(值 interface{}, 选项 Options) *Json {
	var j *Json
	switch 值.(type) {
	case string, []byte:
		if r, err := loadContentWithOptions(值, 选项); err == nil {
			j = r
		} else {
			j = &Json{
				p:  &值,
				c:  byte(defaultSplitChar),
				vc: false,
			}
		}
	default:
		var (
			pointedData interface{}
			reflectInfo = reflection.OriginValueAndKind(值)
		)
		switch reflectInfo.OriginKind {
		case reflect.Slice, reflect.Array:
			pointedData = gconv.Interfaces(值)

		case reflect.Map:
			pointedData = gconv.MapDeep(值, 选项.Tags)

		case reflect.Struct:
			if v, ok := 值.(iVal); ok {
				return X创建并按选项(v.Val(), 选项)
			}
			pointedData = gconv.MapDeep(值, 选项.Tags)

		default:
			pointedData = 值
		}
		j = &Json{
			p:  &pointedData,
			c:  byte(defaultSplitChar),
			vc: false,
		}
	}
	j.mu = rwmutex.Create(选项.Safe)
	return j
}

// Load 从指定的文件路径`path`加载内容，并根据其内容创建一个Json对象。
func X加载文件(路径 string, 并发安全 ...bool) (*Json, error) {
	if p, err := gfile.Search(路径); err != nil {
		return nil, err
	} else {
		路径 = p
	}
	options := Options{
		Type: ContentType(gfile.Ext(路径)),
	}
	if len(并发安全) > 0 && 并发安全[0] {
		options.Safe = true
	}
	return doLoadContentWithOptions(gfile.GetBytesWithCache(路径), options)
}

// LoadWithOptions 根据给定的JSON格式内容和选项，创建一个Json对象。
func X加载并按选项(值 interface{}, 选项 Options) (*Json, error) {
	return doLoadContentWithOptions(gconv.Bytes(值), 选项)
}

// LoadJson 从给定的 JSON 格式内容创建一个 Json 对象。
func X加载json(值 interface{}, 并发安全 ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeJson,
	}
	if len(并发安全) > 0 && 并发安全[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.Bytes(值), option)
}

// LoadXml 从给定的 XML 格式内容创建一个 Json 对象。
func X加载xml(值 interface{}, 并发安全 ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeXml,
	}
	if len(并发安全) > 0 && 并发安全[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.Bytes(值), option)
}

// LoadIni 从给定的 INI 格式内容创建一个 Json 对象。
func X加载ini(值 interface{}, 并发安全 ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeIni,
	}
	if len(并发安全) > 0 && 并发安全[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.Bytes(值), option)
}

// LoadYaml 从给定的YAML格式内容创建一个Json对象。
func X加载Yaml(值 interface{}, 并发安全 ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeYaml,
	}
	if len(并发安全) > 0 && 并发安全[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.Bytes(值), option)
}

// LoadToml 从给定的TOML格式内容创建一个Json对象。
func X加载Toml(值 interface{}, 并发安全 ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeToml,
	}
	if len(并发安全) > 0 && 并发安全[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.Bytes(值), option)
}

// LoadProperties 从给定的TOML格式内容创建一个Json对象。
func X加载Properties(值 interface{}, 并发安全 ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeProperties,
	}
	if len(并发安全) > 0 && 并发安全[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.Bytes(值), option)
}

// LoadContent 从给定的内容创建一个 Json 对象，它会自动检查 `content` 的数据类型，
// 支持以下数据内容类型：
// JSON、XML、INI、YAML 和 TOML。
func X加载并自动识别格式(值 interface{}, 并发安全 ...bool) (*Json, error) {
	content := gconv.Bytes(值)
	if len(content) == 0 {
		return X创建(nil, 并发安全...), nil
	}
	return X加载并按格式(checkDataType(content), content, 并发安全...)
}

// LoadContentType 从给定的类型和内容创建一个 Json 对象，
// 支持以下数据内容类型：
// JSON、XML、INI、YAML 和 TOML。
func X加载并按格式(类型标签 ContentType, 值 interface{}, 并发安全 ...bool) (*Json, error) {
	content := gconv.Bytes(值)
	if len(content) == 0 {
		return X创建(nil, 并发安全...), nil
	}
	// ignore UTF8-BOM
	if content[0] == 0xEF && content[1] == 0xBB && content[2] == 0xBF {
		content = content[3:]
	}
	options := Options{
		Type:      类型标签,
		StrNumber: true,
	}
	if len(并发安全) > 0 && 并发安全[0] {
		options.Safe = true
	}
	return doLoadContentWithOptions(content, options)
}

// IsValidDataType 检查并返回给定的 `dataType` 是否为有效载入数据类型。
func X检查类型(待判断值 ContentType) bool {
	if 待判断值 == "" {
		return false
	}
	if 待判断值[0] == '.' {
		待判断值 = 待判断值[1:]
	}
	switch 待判断值 {
	case
		ContentTypeJson,
		ContentTypeJs,
		ContentTypeXml,
		ContentTypeYaml,
		ContentTypeYml,
		ContentTypeToml,
		ContentTypeIni,
		ContentTypeProperties:
		return true
	}
	return false
}

func loadContentWithOptions(data interface{}, options Options) (*Json, error) {
	content := gconv.Bytes(data)
	if len(content) == 0 {
		return X创建并按选项(nil, options), nil
	}
	if options.Type == "" {
		options.Type = checkDataType(content)
	}
	return loadContentTypeWithOptions(content, options)
}

func loadContentTypeWithOptions(data interface{}, options Options) (*Json, error) {
	content := gconv.Bytes(data)
	if len(content) == 0 {
		return X创建并按选项(nil, options), nil
	}
	// ignore UTF8-BOM
	if content[0] == 0xEF && content[1] == 0xBB && content[2] == 0xBF {
		content = content[3:]
	}
	return doLoadContentWithOptions(content, options)
}

// doLoadContent 从给定的内容创建一个 Json 对象。
// 它支持以下数据内容类型：
// JSON、XML、INI、YAML 和 TOML。
func doLoadContentWithOptions(data []byte, options Options) (*Json, error) {
	var (
		err    error
		result interface{}
	)
	if len(data) == 0 {
		return X创建并按选项(nil, options), nil
	}
	if options.Type == "" {
		options.Type = checkDataType(data)
	}
	options.Type = ContentType(gstr.TrimLeft(
		string(options.Type), "."),
	)
	switch options.Type {
	case ContentTypeJson, ContentTypeJs:

	case ContentTypeXml:
		if data, err = gxml.ToJson(data); err != nil {
			return nil, err
		}

	case ContentTypeYaml, ContentTypeYml:
		if data, err = gyaml.ToJson(data); err != nil {
			return nil, err
		}

	case ContentTypeToml:
		if data, err = gtoml.ToJson(data); err != nil {
			return nil, err
		}

	case ContentTypeIni:
		if data, err = gini.ToJson(data); err != nil {
			return nil, err
		}
	case ContentTypeProperties:
		if data, err = gproperties.ToJson(data); err != nil {
			return nil, err
		}

	default:
		err = gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`unsupported type "%s" for loading`,
			options.Type,
		)
	}
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	if options.StrNumber {
		decoder.UseNumber()
	}
	if err = decoder.Decode(&result); err != nil {
		return nil, err
	}
	switch result.(type) {
	case string, []byte:
		return nil, gerror.Newf(`json decoding failed for content: %s`, data)
	}
	return X创建并按选项(result, options), nil
}

// checkDataType 自动检查并返回 `content` 的数据类型。
// 注意，它使用正则表达式进行宽松的检查，你可以使用 LoadXXX/LoadContentType
// 函数来按特定内容类型加载内容。
func checkDataType(content []byte) ContentType {
	if json.Valid(content) {
		return ContentTypeJson
	} else if gregex.IsMatch(`^<.+>[\S\s]+<.+>\s*$`, content) {
		return ContentTypeXml
	} else if !gregex.IsMatch(`[\n\r]*[\s\t\w\-\."]+\s*=\s*"""[\s\S]+"""`, content) &&
		!gregex.IsMatch(`[\n\r]*[\s\t\w\-\."]+\s*=\s*'''[\s\S]+'''`, content) &&
		((gregex.IsMatch(`^[\n\r]*[\w\-\s\t]+\s*:\s*".+"`, content) || gregex.IsMatch(`^[\n\r]*[\w\-\s\t]+\s*:\s*\w+`, content)) ||
			(gregex.IsMatch(`[\n\r]+[\w\-\s\t]+\s*:\s*".+"`, content) || gregex.IsMatch(`[\n\r]+[\w\-\s\t]+\s*:\s*\w+`, content))) {
		return ContentTypeYaml
	} else if !gregex.IsMatch(`^[\s\t\n\r]*;.+`, content) &&
		!gregex.IsMatch(`[\s\t\n\r]+;.+`, content) &&
		!gregex.IsMatch(`[\n\r]+[\s\t\w\-]+\.[\s\t\w\-]+\s*=\s*.+`, content) &&
		(gregex.IsMatch(`[\n\r]*[\s\t\w\-\."]+\s*=\s*".+"`, content) || gregex.IsMatch(`[\n\r]*[\s\t\w\-\."]+\s*=\s*\w+`, content)) {
		return ContentTypeToml
	} else if gregex.IsMatch(`\[[\w\.]+\]`, content) &&
		(gregex.IsMatch(`[\n\r]*[\s\t\w\-\."]+\s*=\s*".+"`, content) || gregex.IsMatch(`[\n\r]*[\s\t\w\-\."]+\s*=\s*\w+`, content)) {
		// 必须包含 "[xxx]" 部分。
		return ContentTypeIni
	} else if gregex.IsMatch(`[\n\r]*[\s\t\w\-\."]+\s*=\s*\w+`, content) {
		return ContentTypeProperties
	} else {
		return ""
	}
}
