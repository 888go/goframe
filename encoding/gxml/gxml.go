// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gxml 提供了对 XML 内容的访问和转换功能。
package xml类

import (
	"strings"
	
	"github.com/clbanning/mxj/v2"
	
	"github.com/888go/goframe/encoding/gcharset"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/text/gregex"
)

// Decode 解析 `content` 并将其转换为 map 后返回。
func Decode(content []byte) (map[string]interface{}, error) {
	res, err := convert(content)
	if err != nil {
		return nil, err
	}
	m, err := mxj.NewMapXml(res)
	if err != nil {
		err = 错误类.X多层错误并格式化(err, `mxj.NewMapXml failed`)
	}
	return m, err
}

// DecodeWithoutRoot 将 `content` 解析为一个映射（map），并返回没有根级别的映射。
func DecodeWithoutRoot(content []byte) (map[string]interface{}, error) {
	res, err := convert(content)
	if err != nil {
		return nil, err
	}
	m, err := mxj.NewMapXml(res)
	if err != nil {
		err = 错误类.X多层错误并格式化(err, `mxj.NewMapXml failed`)
		return nil, err
	}
	for _, v := range m {
		if r, ok := v.(map[string]interface{}); ok {
			return r, nil
		}
	}
	return m, nil
}

// Encode 将字典 `m` 编码为 XML 格式的内容并以字节形式输出。
// 可选参数 `rootTag` 用于指定 XML 根标签。
func Encode(m map[string]interface{}, rootTag ...string) ([]byte, error) {
	b, err := mxj.Map(m).Xml(rootTag...)
	if err != nil {
		err = 错误类.X多层错误并格式化(err, `mxj.Map.Xml failed`)
	}
	return b, err
}

// EncodeWithIndent 将 map `m` 编码为带有缩进的 XML 格式字节内容。
// 可选参数 `rootTag` 用于指定 XML 根标签。
func EncodeWithIndent(m map[string]interface{}, rootTag ...string) ([]byte, error) {
	b, err := mxj.Map(m).XmlIndent("", "\t", rootTag...)
	if err != nil {
		err = 错误类.X多层错误并格式化(err, `mxj.Map.XmlIndent failed`)
	}
	return b, err
}

// ToJson 将`content`以XML格式转换为JSON格式的字节流。
func ToJson(content []byte) ([]byte, error) {
	res, err := convert(content)
	if err != nil {
		return nil, err
	}
	mv, err := mxj.NewMapXml(res)
	if err == nil {
		return mv.Json()
	}
	err = 错误类.X多层错误(err, `mxj.NewMapXml failed`)
	return nil, err
}

// convert 将给定XML内容从XML根标签转换为UTF-8编码的内容。
func convert(xml []byte) (res []byte, err error) {
	var (
		patten      = `<\?xml.*encoding\s*=\s*['|"](.*?)['|"].*\?>`
		matchStr, _ = 正则类.X匹配文本(patten, string(xml))
		xmlEncode   = "UTF-8"
	)
	if len(matchStr) == 2 {
		xmlEncode = matchStr[1]
	}
	xmlEncode = strings.ToUpper(xmlEncode)
	res, err = 正则类.X替换字节集(patten, []byte(""), xml)
	if err != nil {
		return nil, err
	}
	if xmlEncode != "UTF-8" && xmlEncode != "UTF8" {
		dst, err := 编码字符集类.Convert("UTF-8", xmlEncode, string(res))
		if err != nil {
			return nil, err
		}
		res = []byte(dst)
	}
	return res, nil
}
