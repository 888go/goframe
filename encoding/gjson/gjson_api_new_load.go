// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gjson

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
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/internal/reflection"
	"github.com/gogf/gf/v2/internal/rwmutex"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// New 使用任何类型的`data`创建一个Json对象，但为了数据访问的原因，`data`应该是map或slice，
// 否则将失去意义。
//
// 参数`safe`指定是否在并发安全的上下文中使用此Json对象，默认值为false。 md5:b84f401db24e69d8
func New(data interface{}, safe ...bool) *Json {
	return NewWithTag(data, string(ContentTypeJson), safe...)
}

// NewWithTag 创建一个Json对象，可以包含任何类型的`data`，但出于数据访问的原因，`data`应该是一个map或切片，否则将没有意义。
//
// 参数`tags`用于指定结构体转换为map的优先标签，多个标签之间用逗号分隔。
//
// 参数`safe`表示是否在并发安全上下文中使用这个Json对象，默认为false。 md5:2558f08f4f082a16
func NewWithTag(data interface{}, tags string, safe ...bool) *Json {
	option := Options{
		Tags: tags,
	}
	if len(safe) > 0 && safe[0] {
		option.Safe = true
	}
	return NewWithOptions(data, option)
}

// NewWithOptions使用任何类型的'data'创建一个Json对象，但出于数据访问的原因，`data`应该是map或切片，否则将没有意义。 md5:48be1828a6556518
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

// Load 从指定的文件`path`加载内容，并根据其内容创建一个Json对象。 md5:fc26d8aa3d537173
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

// LoadWithOptions 根据给定的 JSON 格式内容和选项创建一个 Json 对象。 md5:77290b5f994f3ff1
func LoadWithOptions(data interface{}, options Options) (*Json, error) {
	return doLoadContentWithOptions(gconv.Bytes(data), options)
}

// LoadJson 从给定的JSON格式内容创建一个Json对象。 md5:1f41cbc0a35bd390
func LoadJson(data interface{}, safe ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeJson,
	}
	if len(safe) > 0 && safe[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.Bytes(data), option)
}

// LoadXml 从给定的XML格式内容创建一个Json对象。 md5:a170d56aa371a2bb
func LoadXml(data interface{}, safe ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeXml,
	}
	if len(safe) > 0 && safe[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.Bytes(data), option)
}

// LoadIni 从给定的INI格式内容创建一个Json对象。 md5:bf3225da0be4c26b
func LoadIni(data interface{}, safe ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeIni,
	}
	if len(safe) > 0 && safe[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.Bytes(data), option)
}

// LoadYaml 根据给定的 YAML 格式内容创建一个 Json 对象。 md5:d810aac213716b5a
func LoadYaml(data interface{}, safe ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeYaml,
	}
	if len(safe) > 0 && safe[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.Bytes(data), option)
}

// LoadToml 从给定的TOML格式内容创建一个Json对象。 md5:a27ac84d2a7e5a70
func LoadToml(data interface{}, safe ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeToml,
	}
	if len(safe) > 0 && safe[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.Bytes(data), option)
}

// LoadProperties 从给定的TOML格式内容创建一个Json对象。 md5:aacff07e57605d82
func LoadProperties(data interface{}, safe ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeProperties,
	}
	if len(safe) > 0 && safe[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.Bytes(data), option)
}

// LoadContent 根据给定的内容创建一个Json对象，它会自动检查`content`的数据类型，
// 支持如下数据内容类型：
// JSON、XML、INI、YAML和TOML。 md5:e930374f4ac3b32e
func LoadContent(data interface{}, safe ...bool) (*Json, error) {
	content := gconv.Bytes(data)
	if len(content) == 0 {
		return New(nil, safe...), nil
	}
	return LoadContentType(checkDataType(content), content, safe...)
}

// LoadContentType 根据给定的类型和内容创建一个 JSON 对象，支持以下数据内容类型：
// JSON, XML, INI, YAML 和 TOML. md5:7db5bd0b429fea01
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

// IsValidDataType 检查并返回给定的 `dataType` 是否是用于加载的有效数据类型。 md5:3cc6cab5a2631a3e
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

// doLoadContent 从给定内容创建一个Json对象。
// 它支持以下数据内容类型：
// JSON、XML、INI、YAML和TOML。 md5:a1daf6666c64b0bc
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

// checkDataType 会自动检查并返回`content`的数据类型。
// 注意，它使用正则表达式进行宽松的检查，你可以根据需要使用LoadXXX/LoadContentType
// 函数来为特定内容类型加载内容。 md5:faa69696c8f02af2
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
		// 必须包含"xxx"部分。 md5:6dc6d0a6d417b6a6
		return ContentTypeIni
	} else if gregex.IsMatch(`[\n\r]*[\s\t\w\-\."]+\s*=\s*\w+`, content) {
		return ContentTypeProperties
	} else {
		return ""
	}
}
