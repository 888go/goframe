// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package json类

import (
	"bytes"
	"reflect"

	gini "github.com/888go/goframe/encoding/gini"
	"github.com/888go/goframe/encoding/gproperties"
	gtoml "github.com/888go/goframe/encoding/gtoml"
	gxml "github.com/888go/goframe/encoding/gxml"
	gyaml "github.com/888go/goframe/encoding/gyaml"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/reflection"
	"github.com/888go/goframe/internal/rwmutex"
	gfile "github.com/888go/goframe/os/gfile"
	gregex "github.com/888go/goframe/text/gregex"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
)

// X创建 使用任何类型的`data`创建一个Json对象，但为了数据访问的原因，`data`应该是map或slice，
// 否则将失去意义。
//
// 参数`safe`指定是否在并发安全的上下文中使用此Json对象，默认值为false。
// md5:b84f401db24e69d8
func X创建(值 interface{}, 并发安全 ...bool) *Json {
	return X创建并按类型标签(值, string(ContentTypeJson), 并发安全...)
}

// X创建并按类型标签 创建一个Json对象，可以包含任何类型的`data`，但出于数据访问的原因，`data`应该是一个map或切片，否则将没有意义。
// 
// 参数`tags`用于指定结构体转换为map的优先标签，多个标签之间用逗号分隔。
// 
// 参数`safe`表示是否在并发安全上下文中使用这个Json对象，默认为false。
// md5:2558f08f4f082a16
func X创建并按类型标签(值 interface{}, 类型标签 string, 并发安全 ...bool) *Json {
	option := Options{
		Tags: 类型标签,
	}
	if len(并发安全) > 0 && 并发安全[0] {
		option.Safe = true
	}
	return X创建并按选项(值, option)
}

// X创建并按选项使用任何类型的'data'创建一个Json对象，但出于数据访问的原因，`data`应该是map或切片，否则将没有意义。
// md5:48be1828a6556518
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
			pointedData = gconv.X取any切片(值)

		case reflect.Map:
			pointedData = gconv.X取Map_递归(值, 选项.Tags)

		case reflect.Struct:
			if v, ok := 值.(iVal); ok {
				return X创建并按选项(v.X取值(), 选项)
			}
			pointedData = gconv.X取Map_递归(值, 选项.Tags)

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

// X加载文件 从指定的文件`path`加载内容，并根据其内容创建一个Json对象。 md5:fc26d8aa3d537173
func X加载文件(路径 string, 并发安全 ...bool) (*Json, error) {
	if p, err := gfile.X查找(路径); err != nil {
		return nil, err
	} else {
		路径 = p
	}
	options := Options{
		Type: ContentType(gfile.X路径取扩展名(路径)),
	}
	if len(并发安全) > 0 && 并发安全[0] {
		options.Safe = true
	}
	return doLoadContentWithOptions(gfile.X缓存读字节集(路径), options)
}

// X加载并按选项 根据给定的 JSON 格式内容和选项创建一个 Json 对象。 md5:77290b5f994f3ff1
func X加载并按选项(值 interface{}, 选项 Options) (*Json, error) {
	return doLoadContentWithOptions(gconv.X取字节集(值), 选项)
}

// X加载json 从给定的JSON格式内容创建一个Json对象。 md5:1f41cbc0a35bd390
func X加载json(值 interface{}, 并发安全 ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeJson,
	}
	if len(并发安全) > 0 && 并发安全[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.X取字节集(值), option)
}

// X加载xml 从给定的XML格式内容创建一个Json对象。 md5:a170d56aa371a2bb
func X加载xml(值 interface{}, 并发安全 ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeXml,
	}
	if len(并发安全) > 0 && 并发安全[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.X取字节集(值), option)
}

// X加载ini 从给定的INI格式内容创建一个Json对象。 md5:bf3225da0be4c26b
func X加载ini(值 interface{}, 并发安全 ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeIni,
	}
	if len(并发安全) > 0 && 并发安全[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.X取字节集(值), option)
}

// X加载Yaml 根据给定的 YAML 格式内容创建一个 Json 对象。 md5:d810aac213716b5a
func X加载Yaml(值 interface{}, 并发安全 ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeYaml,
	}
	if len(并发安全) > 0 && 并发安全[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.X取字节集(值), option)
}

// X加载Toml 从给定的TOML格式内容创建一个Json对象。 md5:a27ac84d2a7e5a70
func X加载Toml(值 interface{}, 并发安全 ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeToml,
	}
	if len(并发安全) > 0 && 并发安全[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.X取字节集(值), option)
}

// X加载Properties 从给定的TOML格式内容创建一个Json对象。 md5:aacff07e57605d82
func X加载Properties(值 interface{}, 并发安全 ...bool) (*Json, error) {
	option := Options{
		Type: ContentTypeProperties,
	}
	if len(并发安全) > 0 && 并发安全[0] {
		option.Safe = true
	}
	return doLoadContentWithOptions(gconv.X取字节集(值), option)
}

// X加载并自动识别格式 根据给定的内容创建一个Json对象，它会自动检查`content`的数据类型，
// 支持如下数据内容类型：
// JSON、XML、INI、YAML和TOML。
// md5:e930374f4ac3b32e
func X加载并自动识别格式(值 interface{}, 并发安全 ...bool) (*Json, error) {
	content := gconv.X取字节集(值)
	if len(content) == 0 {
		return X创建(nil, 并发安全...), nil
	}
	return X加载并按格式(checkDataType(content), content, 并发安全...)
}

// X加载并按格式 根据给定的类型和内容创建一个 JSON 对象，支持以下数据内容类型：
// JSON, XML, INI, YAML 和 TOML.
// md5:7db5bd0b429fea01
func X加载并按格式(类型标签 ContentType, 值 interface{}, 并发安全 ...bool) (*Json, error) {
	content := gconv.X取字节集(值)
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

// X检查类型 检查并返回给定的 `dataType` 是否是用于加载的有效数据类型。 md5:3cc6cab5a2631a3e
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
	content := gconv.X取字节集(data)
	if len(content) == 0 {
		return X创建并按选项(nil, options), nil
	}
	if options.Type == "" {
		options.Type = checkDataType(content)
	}
	return loadContentTypeWithOptions(content, options)
}

func loadContentTypeWithOptions(data interface{}, options Options) (*Json, error) {
	content := gconv.X取字节集(data)
	if len(content) == 0 {
		return X创建并按选项(nil, options), nil
	}
	// ignore UTF8-BOM
	if content[0] == 0xEF && content[1] == 0xBB && content[2] == 0xBF {
		content = content[3:]
	}
	return doLoadContentWithOptions(content, options)
}

// doLoadContent 从给定内容创建一个Json对象。
// 它支持以下数据内容类型：
// JSON、XML、INI、YAML和TOML。
// md5:a1daf6666c64b0bc
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
	options.Type = ContentType(gstr.X过滤首字符并含空白(
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
		if data, err = gini.X取json(data); err != nil {
			return nil, err
		}
	case ContentTypeProperties:
		if data, err = gproperties.ToJson(data); err != nil {
			return nil, err
		}

	default:
		err = gerror.X创建错误码并格式化(
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
		return nil, gerror.X创建并格式化(`json decoding failed for content: %s`, data)
	}
	return X创建并按选项(result, options), nil
}

// checkDataType 会自动检查并返回`content`的数据类型。
// 注意，它使用正则表达式进行宽松的检查，你可以根据需要使用LoadXXX/LoadContentType
// 函数来为特定内容类型加载内容。
// md5:faa69696c8f02af2
func checkDataType(content []byte) ContentType {
	if json.Valid(content) {
		return ContentTypeJson
	} else if gregex.X是否匹配字节集(`^<.+>[\S\s]+<.+>\s*$`, content) {
		return ContentTypeXml
	} else if !gregex.X是否匹配字节集(`[\n\r]*[\s\t\w\-\."]+\s*=\s*"""[\s\S]+"""`, content) &&
		!gregex.X是否匹配字节集(`[\n\r]*[\s\t\w\-\."]+\s*=\s*'''[\s\S]+'''`, content) &&
		((gregex.X是否匹配字节集(`^[\n\r]*[\w\-\s\t]+\s*:\s*".+"`, content) || gregex.X是否匹配字节集(`^[\n\r]*[\w\-\s\t]+\s*:\s*\w+`, content)) ||
			(gregex.X是否匹配字节集(`[\n\r]+[\w\-\s\t]+\s*:\s*".+"`, content) || gregex.X是否匹配字节集(`[\n\r]+[\w\-\s\t]+\s*:\s*\w+`, content))) {
		return ContentTypeYaml
	} else if !gregex.X是否匹配字节集(`^[\s\t\n\r]*;.+`, content) &&
		!gregex.X是否匹配字节集(`[\s\t\n\r]+;.+`, content) &&
		!gregex.X是否匹配字节集(`[\n\r]+[\s\t\w\-]+\.[\s\t\w\-]+\s*=\s*.+`, content) &&
		(gregex.X是否匹配字节集(`[\n\r]*[\s\t\w\-\."]+\s*=\s*".+"`, content) || gregex.X是否匹配字节集(`[\n\r]*[\s\t\w\-\."]+\s*=\s*\w+`, content)) {
		return ContentTypeToml
	} else if gregex.X是否匹配字节集(`\[[\w\.]+\]`, content) &&
		(gregex.X是否匹配字节集(`[\n\r]*[\s\t\w\-\."]+\s*=\s*".+"`, content) || gregex.X是否匹配字节集(`[\n\r]*[\s\t\w\-\."]+\s*=\s*\w+`, content)) {
				// 必须包含"xxx"部分。 md5:6dc6d0a6d417b6a6
		return ContentTypeIni
	} else if gregex.X是否匹配字节集(`[\n\r]*[\s\t\w\-\."]+\s*=\s*\w+`, content) {
		return ContentTypeProperties
	} else {
		return ""
	}
}
