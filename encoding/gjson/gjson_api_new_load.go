// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gjson
import (
	"bytes"
	"reflect"
	
	"github.com/888go/goframe/encoding/gini"
	"github.com/888go/goframe/encoding/gproperties"
	"github.com/888go/goframe/encoding/gtoml"
	"github.com/888go/goframe/encoding/gxml"
	"github.com/888go/goframe/encoding/gyaml"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/reflection"
	"github.com/888go/goframe/internal/rwmutex"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	)
// New 函数通过任意类型的 `data` 创建一个 Json 对象，但 `data` 应为 map 或 slice 类型以保证数据可访问性，否则创建此对象将无实际意义。
//
// 参数 `safe` 指定是否在并发安全的上下文中使用此 Json 对象，默认情况下 `safe` 为 false。
func New(data interface{}, safe ...bool) *Json {
	return NewWithTag(data, string(ContentTypeJson), safe...)
}

// NewWithTag 创建一个Json对象，其数据类型可以是任意的 `data`，但为了方便数据访问，`data` 应该是一个 map 或 slice，否则将失去意义。
//
// 参数 `tags` 指定了在结构体转为 map 时使用的优先级标签，多个标签使用字符 ',' 连接。
//
// 参数 `safe` 指定了是否在并发安全的上下文中使用此 Json 对象，默认情况下为 false。
func NewWithTag(data interface{}, tags string, safe ...bool) *Json {
	option := Options{
		Tags: tags,
	}
	if len(safe) > 0 && safe[0] {
		option.Safe = true
	}
	return NewWithOptions(data, option)
}

// NewWithOptions 创建一个Json对象，其变量类型可以是 `data` 的任意类型，但为了能够访问数据，`data` 应该是一个 map 或 slice，否则将毫无意义。
func NewWithOptions(data interface{}, options Options) *Json {
	var j *Json
	switch data.(type) {
	case string, []byte:
		if r, err := loadContentWithOptions(data, options); err == nil {
			j = r
		} else {
			j = &Json{
				p:  &data,
				c:  byte(defaultSplitChar),
				vc: false,
			}
		}
	default:
		var (
			pointedData interface{}
			reflectInfo = reflection.OriginValueAndKind(data)
		)
		switch reflectInfo.OriginKind {
		case reflect.Slice, reflect.Array:
			pointedData = gconv.Interfaces(data)

		case reflect.Map:
			pointedData = gconv.MapDeep(data, options.Tags)

		case reflect.Struct:
			if v, ok := data.(iVal); ok {
				return NewWithOptions(v.Val(), options)
			}
			pointedData = gconv.MapDeep(data, options.Tags)

		default:
			pointedData = data
		}
		j = &Json{
			p:  &pointedData,
			c:  byte(defaultSplitChar),
			vc: false,
		}
	}
	j.mu = rwmutex.Create(options.Safe)
	return j
}

// Load 从指定的文件路径`path`加载内容，并根据其内容创建一个Json对象。
func Load(path string, safe ...bool) (*Json, error) {
	if p, err := gfile.Search(path); err != nil {
		return nil, err
	} else {
		path = p
	}
	options := Options{
		Type: ContentType(gfile.Ext(path)),
	}
	if len(safe) > 0 && safe[0] {
		options.Safe = true
	}
	return doLoadContentWithOptions(gfile.GetBytesWithCache(path), options)
}

// LoadWithOptions 根据给定的JSON格式内容和选项，创建一个Json对象。
func LoadWithOptions(data interface{}, options Options) (*Json, error) {
	return doLoadContentWithOptions(gconv.Bytes(data), options)
}

// LoadJson 从给定的 JSON 格式内容创建一个 Json 对象。
func LoadJson(data interface{}, safe ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeJson,
	}
	if len(safe) > 0 && safe[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.Bytes(data), option)
}

// LoadXml 从给定的 XML 格式内容创建一个 Json 对象。
func LoadXml(data interface{}, safe ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeXml,
	}
	if len(safe) > 0 && safe[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.Bytes(data), option)
}

// LoadIni 从给定的 INI 格式内容创建一个 Json 对象。
func LoadIni(data interface{}, safe ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeIni,
	}
	if len(safe) > 0 && safe[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.Bytes(data), option)
}

// LoadYaml 从给定的YAML格式内容创建一个Json对象。
func LoadYaml(data interface{}, safe ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeYaml,
	}
	if len(safe) > 0 && safe[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.Bytes(data), option)
}

// LoadToml 从给定的TOML格式内容创建一个Json对象。
func LoadToml(data interface{}, safe ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeToml,
	}
	if len(safe) > 0 && safe[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.Bytes(data), option)
}

// LoadProperties 从给定的TOML格式内容创建一个Json对象。
func LoadProperties(data interface{}, safe ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeProperties,
	}
	if len(safe) > 0 && safe[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.Bytes(data), option)
}

// LoadContent 从给定的内容创建一个 Json 对象，它会自动检查 `content` 的数据类型，
// 支持以下数据内容类型：
// JSON、XML、INI、YAML 和 TOML。
func LoadContent(data interface{}, safe ...bool) (*Json, error) {
	content := gconv.Bytes(data)
	if len(content) == 0 {
		return New(nil, safe...), nil
	}
	return LoadContentType(checkDataType(content), content, safe...)
}

// LoadContentType 从给定的类型和内容创建一个 Json 对象，
// 支持以下数据内容类型：
// JSON、XML、INI、YAML 和 TOML。
func LoadContentType(dataType ContentType, data interface{}, safe ...bool) (*Json, error) {
	content := gconv.Bytes(data)
	if len(content) == 0 {
		return New(nil, safe...), nil
	}
	// ignore UTF8-BOM
	if content[0] == 0xEF && content[1] == 0xBB && content[2] == 0xBF {
		content = content[3:]
	}
	options := Options{
		Type:      dataType,
		StrNumber: true,
	}
	if len(safe) > 0 && safe[0] {
		options.Safe = true
	}
	return doLoadContentWithOptions(content, options)
}

// IsValidDataType 检查并返回给定的 `dataType` 是否为有效载入数据类型。
func IsValidDataType(dataType ContentType) bool {
	if dataType == "" {
		return false
	}
	if dataType[0] == '.' {
		dataType = dataType[1:]
	}
	switch dataType {
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
		return NewWithOptions(nil, options), nil
	}
	if options.Type == "" {
		options.Type = checkDataType(content)
	}
	return loadContentTypeWithOptions(content, options)
}

func loadContentTypeWithOptions(data interface{}, options Options) (*Json, error) {
	content := gconv.Bytes(data)
	if len(content) == 0 {
		return NewWithOptions(nil, options), nil
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
		return NewWithOptions(nil, options), nil
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
	return NewWithOptions(result, options), nil
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
