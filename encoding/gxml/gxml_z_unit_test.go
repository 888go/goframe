// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package xml类_test

import (
	"bytes"
	"strings"
	"testing"

	gcharset "github.com/888go/goframe/encoding/gcharset"
	gjson "github.com/888go/goframe/encoding/gjson"
	gxml "github.com/888go/goframe/encoding/gxml"
	gtest "github.com/888go/goframe/test/gtest"
)

var testData = []struct {
	utf8, other, otherEncoding string
}{
	{"Hello 常用國字標準字體表", "Hello \xb1`\xa5\u03b0\xea\xa6r\xbc\u0437\u01e6r\xc5\xe9\xaa\xed", "big5"},
	{"Hello 常用國字標準字體表", "Hello \xb3\xa3\xd3\xc3\x87\xf8\xd7\xd6\x98\xcb\x9c\xca\xd7\xd6\xf3\x77\xb1\xed", "gbk"},
	{"Hello 常用國字標準字體表", "Hello \xb3\xa3\xd3\xc3\x87\xf8\xd7\xd6\x98\xcb\x9c\xca\xd7\xd6\xf3\x77\xb1\xed", "gb18030"},
}

var testErrData = []struct {
	utf8, other, otherEncoding string
}{
	{"Hello 常用國字標準字體表", "Hello \xb3\xa3\xd3\xc3\x87\xf8\xd7\xd6\x98\xcb\x9c\xca\xd7\xd6\xf3\x77\xb1\xed", "gbk"},
}

func buildXml(charset string, str string) (string, string) {
	head := `<?xml version="1.0" encoding="UTF-8"?>`
	srcXml := strings.Replace(head, "UTF-8", charset, -1)

	srcParser := gjson.X创建(nil)
	srcParser.X设置值("name", str)
	srcParser.X设置值("age", "12")

	s, err := srcParser.X取xml字节集()
	if err != nil {
		return "", ""
	}

	srcXml = srcXml + string(s)
	srcXml, err = gcharset.UTF8To(charset, srcXml)
	if err != nil {
		return "", ""
	}

	dstXml := head + string(s)

	return srcXml, dstXml
}

// 测试XML中字符集的转换
func Test_XmlToJson(t *testing.T) {
	for _, v := range testData {
		srcXml, dstXml := buildXml(v.otherEncoding, v.utf8)
		if len(srcXml) == 0 && len(dstXml) == 0 {
			t.Errorf("build xml string error. srcEncoding:%s, src:%s, utf8:%s", v.otherEncoding, v.other, v.utf8)
		}

		srcJson, err := gxml.ToJson([]byte(srcXml))
		if err != nil {
			t.Errorf("gxml.ToJson error. %s", srcXml)
		}

		dstJson, err := gxml.ToJson([]byte(dstXml))
		if err != nil {
			t.Errorf("dstXml to json error. %s", dstXml)
		}

		if !bytes.Equal(srcJson, dstJson) {
			t.Errorf("convert to json error. srcJson:%s, dstJson:%s", string(srcJson), string(dstJson))
		}

	}
}

func Test_Decode1(t *testing.T) {
	for _, v := range testData {
		srcXml, dstXml := buildXml(v.otherEncoding, v.utf8)
		if len(srcXml) == 0 && len(dstXml) == 0 {
			t.Errorf("build xml string error. srcEncoding:%s, src:%s, utf8:%s", v.otherEncoding, v.other, v.utf8)
		}

		srcMap, err := gxml.Decode([]byte(srcXml))
		if err != nil {
			t.Errorf("gxml.Decode error. %s", srcXml)
		}

		dstMap, err := gxml.Decode([]byte(dstXml))
		if err != nil {
			t.Errorf("gxml decode error. %s", dstXml)
		}
		s := srcMap["doc"].(map[string]interface{})
		d := dstMap["doc"].(map[string]interface{})
		for kk, vv := range s {
			if vv.(string) != d[kk].(string) {
				t.Errorf("convert to map error. src:%v, dst:%v", vv, d[kk])
			}
		}
	}
}

func Test_Decode2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		content := `
<?xml version="1.0" encoding="UTF-8"?><doc><username>johngcn</username><password1>123456</password1><password2>123456</password2></doc>
`
		m, err := gxml.Decode([]byte(content))
		t.AssertNil(err)
		t.Assert(m["doc"].(map[string]interface{})["username"], "johngcn")
		t.Assert(m["doc"].(map[string]interface{})["password1"], "123456")
		t.Assert(m["doc"].(map[string]interface{})["password2"], "123456")
	})
}

func Test_DecodeWitoutRoot(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		content := `
<?xml version="1.0" encoding="UTF-8"?><doc><username>johngcn</username><password1>123456</password1><password2>123456</password2></doc>
`
		m, err := gxml.DecodeWithoutRoot([]byte(content))
		t.AssertNil(err)
		t.Assert(m["username"], "johngcn")
		t.Assert(m["password1"], "123456")
		t.Assert(m["password2"], "123456")
	})
}

func Test_Encode(t *testing.T) {
	m := make(map[string]interface{})
	v := map[string]interface{}{
		"string": "hello world",
		"int":    123,
		"float":  100.92,
		"bool":   true,
	}
	m["root"] = interface{}(v)

	xmlStr, err := gxml.Encode(m)
	if err != nil {
		t.Errorf("encode error.")
	}
		// 使用t.Logf记录日志，格式化输出字符串xmlStr的内容，并在末尾添加换行符。 md5:06a73acff84b1307

	res := `<root><bool>true</bool><float>100.92</float><int>123</int><string>hello world</string></root>`
	if string(xmlStr) != res {
		t.Errorf("encode error. result: [%s], expect:[%s]", string(xmlStr), res)
	}
}

func Test_EncodeIndent(t *testing.T) {
	m := make(map[string]interface{})
	v := map[string]interface{}{
		"string": "hello world",
		"int":    123,
		"float":  100.92,
		"bool":   true,
	}
	m["root"] = interface{}(v)

	_, err := gxml.EncodeWithIndent(m, "xml")
	if err != nil {
		t.Errorf("encodeWithIndent error.")
	}

		// 使用t.Logf记录日志，格式化输出字符串xmlStr的内容，并在末尾添加换行符。 md5:06a73acff84b1307

}

func TestErrXml(t *testing.T) {
	for _, v := range testErrData {
		srcXml, dstXml := buildXml(v.otherEncoding, v.utf8)
		if len(srcXml) == 0 && len(dstXml) == 0 {
			t.Errorf("build xml string error. srcEncoding:%s, src:%s, utf8:%s", v.otherEncoding, v.other, v.utf8)
		}

		srcXml = strings.Replace(srcXml, "gbk", "XXX", -1)
		_, err := gxml.ToJson([]byte(srcXml))
		if err == nil {
			t.Errorf("srcXml to json should be failed. %s", srcXml)
		}

	}
}

func TestErrCase(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		errXml := `<root><bool>true</bool><float>100.92</float><int>123</int><string>hello world</string>`
		_, err := gxml.ToJson([]byte(errXml))
		if err == nil {
			t.Errorf("unexpected value: nil")
		}
	})

	gtest.C(t, func(t *gtest.T) {
		errXml := `<root><bool>true</bool><float>100.92</float><int>123</int><string>hello world</string>`
		_, err := gxml.Decode([]byte(errXml))
		if err == nil {
			t.Errorf("unexpected value: nil")
		}
	})
}
