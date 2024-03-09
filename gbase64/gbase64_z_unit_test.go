// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 编码base64类_test

import (
	"testing"
	
	"github.com/888go/goframe/gbase64"
	"github.com/gogf/gf/v2/test/gtest"
)

type testPair struct {
	decoded, encoded string
}

var pairs = []testPair{
	// RFC 3548 examples
	{"\x14\xfb\x9c\x03\xd9\x7e", "FPucA9l+"},
	{"\x14\xfb\x9c\x03\xd9", "FPucA9k="},
	{"\x14\xfb\x9c\x03", "FPucAw=="},

	// RFC 4648 examples
	{"", ""},
	{"f", "Zg=="},
	{"fo", "Zm8="},
	{"foo", "Zm9v"},
	{"foob", "Zm9vYg=="},
	{"fooba", "Zm9vYmE="},
	{"foobar", "Zm9vYmFy"},

	// 维基百科示例
	{"sure.", "c3VyZS4="},
	{"sure", "c3VyZQ=="},
	{"sur", "c3Vy"},
	{"su", "c3U="},
	{"leasure.", "bGVhc3VyZS4="},
	{"easure.", "ZWFzdXJlLg=="},
	{"asure.", "YXN1cmUu"},
	{"sure.", "c3VyZS4="},
}

func Test_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for k := range pairs {
			// Encode
			t.Assert(编码base64类.X字节集编码([]byte(pairs[k].decoded)), []byte(pairs[k].encoded))
			t.Assert(编码base64类.X字节集编码到文本([]byte(pairs[k].decoded)), pairs[k].encoded)
			t.Assert(编码base64类.X文本编码(pairs[k].decoded), pairs[k].encoded)

			// Decode
			r1, _ := 编码base64类.X字节集解码([]byte(pairs[k].encoded))
			t.Assert(r1, []byte(pairs[k].decoded))

			r2, _ := 编码base64类.X文本解码到字节集(pairs[k].encoded)
			t.Assert(r2, []byte(pairs[k].decoded))

			r3, _ := 编码base64类.X文本解码(pairs[k].encoded)
			t.Assert(r3, pairs[k].decoded)
		}
	})
}

func Test_File(t *testing.T) {
	path := gtest.DataPath("test")
	expect := "dGVzdA=="
	gtest.C(t, func(t *gtest.T) {
		b, err := 编码base64类.X文件编码到字节集(path)
		t.AssertNil(err)
		t.Assert(string(b), expect)
	})
	gtest.C(t, func(t *gtest.T) {
		s, err := 编码base64类.X文件编码到文本(path)
		t.AssertNil(err)
		t.Assert(s, expect)
	})
}

func Test_File_Error(t *testing.T) {
	path := "none-exist-file"
	expect := ""
	gtest.C(t, func(t *gtest.T) {
		b, err := 编码base64类.X文件编码到字节集(path)
		t.AssertNE(err, nil)
		t.Assert(string(b), expect)
	})
	gtest.C(t, func(t *gtest.T) {
		s, err := 编码base64类.X文件编码到文本(path)
		t.AssertNE(err, nil)
		t.Assert(s, expect)
	})
}
