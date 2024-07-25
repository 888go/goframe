// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gxml提供了访问和转换XML内容的功能。 md5:7f72b127efb49044
package gxml

import (
	"strings"

	"github.com/clbanning/mxj/v2"

	"github.com/gogf/gf/v2/encoding/gcharset"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gregex"
)

// Decode将`content`解析为映射并返回。 md5:09afa737da32e1d3
func Decode(content []byte) (map[string]interface{}, error) {
	res, err := convert(content)
	if err != nil {
		return nil, err
	}
	m, err := mxj.NewMapXml(res)
	if err != nil {
		err = gerror.Wrapf(err, `mxj.NewMapXml failed`)
	}
	return m, err
}

// DecodeWithoutRoot 将 `content` 解析为一个映射，然后返回不包含根级别的映射。 md5:3210d3b75da05efb
func DecodeWithoutRoot(content []byte) (map[string]interface{}, error) {
	res, err := convert(content)
	if err != nil {
		return nil, err
	}
	m, err := mxj.NewMapXml(res)
	if err != nil {
		err = gerror.Wrapf(err, `mxj.NewMapXml failed`)
		return nil, err
	}
	for _, v := range m {
		if r, ok := v.(map[string]interface{}); ok {
			return r, nil
		}
	}
	return m, nil
}

// Encode 将字典 `m` 编码为 XML 格式的字节内容。
// 可选参数 `rootTag` 用于指定 XML 的根标签。
// md5:b83a924118f435fb
func Encode(m map[string]interface{}, rootTag ...string) ([]byte, error) {
	b, err := mxj.Map(m).Xml(rootTag...)
	if err != nil {
		err = gerror.Wrapf(err, `mxj.Map.Xml failed`)
	}
	return b, err
}

// EncodeWithIndent 将映射 `m` 编码为带缩进的 XML 格式的字节内容。
// 可选参数 `rootTag` 用于指定 XML 根标签。
// md5:4ce035684ef6a0cc
func EncodeWithIndent(m map[string]interface{}, rootTag ...string) ([]byte, error) {
	b, err := mxj.Map(m).XmlIndent("", "\t", rootTag...)
	if err != nil {
		err = gerror.Wrapf(err, `mxj.Map.XmlIndent failed`)
	}
	return b, err
}

// ToJson 将 `content` 作为 XML 格式转换为 JSON 格式的字节。 md5:ba7a69746da22ea8
func ToJson(content []byte) ([]byte, error) {
	res, err := convert(content)
	if err != nil {
		return nil, err
	}
	mv, err := mxj.NewMapXml(res)
	if err == nil {
		return mv.Json()
	}
	err = gerror.Wrap(err, `mxj.NewMapXml failed`)
	return nil, err
}

// convert 将给定XML内容的编码从XML根标签转换为UTF-8编码内容。 md5:c37aa75d79ed6c1b
func convert(xml []byte) (res []byte, err error) {
	var (
		patten      = `<\?xml.*encoding\s*=\s*['|"](.*?)['|"].*\?>`
		matchStr, _ = gregex.MatchString(patten, string(xml))
		xmlEncode   = "UTF-8"
	)
	if len(matchStr) == 2 {
		xmlEncode = matchStr[1]
	}
	xmlEncode = strings.ToUpper(xmlEncode)
	res, err = gregex.Replace(patten, []byte(""), xml)
	if err != nil {
		return nil, err
	}
	if xmlEncode != "UTF-8" && xmlEncode != "UTF8" {
		dst, err := gcharset.Convert("UTF-8", xmlEncode, string(res))
		if err != nil {
			return nil, err
		}
		res = []byte(dst)
	}
	return res, nil
}
