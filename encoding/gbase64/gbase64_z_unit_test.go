// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 编码base64类_test

import (
	"testing"

	gbase64 "github.com/888go/goframe/encoding/gbase64"
	gtest "github.com/888go/goframe/test/gtest"
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

	// Wikipedia examples
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
			t.Assert(gbase64.X字节集编码([]byte(pairs[k].decoded)), []byte(pairs[k].encoded))
			t.Assert(gbase64.X字节集编码到文本([]byte(pairs[k].decoded)), pairs[k].encoded)
			t.Assert(gbase64.X文本编码(pairs[k].decoded), pairs[k].encoded)

			// Decode
			r1, _ := gbase64.X字节集解码([]byte(pairs[k].encoded))
			t.Assert(r1, []byte(pairs[k].decoded))

			r2, _ := gbase64.X文本解码到字节集(pairs[k].encoded)
			t.Assert(r2, []byte(pairs[k].decoded))

			r3, _ := gbase64.X文本解码(pairs[k].encoded)
			t.Assert(r3, pairs[k].decoded)
		}
	})
}

func Test_File(t *testing.T) {
	path := gtest.DataPath("test")
	expect := "dGVzdA=="
	gtest.C(t, func(t *gtest.T) {
		b, err := gbase64.X文件编码到字节集(path)
		t.AssertNil(err)
		t.Assert(string(b), expect)
	})
	gtest.C(t, func(t *gtest.T) {
		s, err := gbase64.X文件编码到文本(path)
		t.AssertNil(err)
		t.Assert(s, expect)
	})
}

func Test_File_Error(t *testing.T) {
	path := "none-exist-file"
	expect := ""
	gtest.C(t, func(t *gtest.T) {
		b, err := gbase64.X文件编码到字节集(path)
		t.AssertNE(err, nil)
		t.Assert(string(b), expect)
	})
	gtest.C(t, func(t *gtest.T) {
		s, err := gbase64.X文件编码到文本(path)
		t.AssertNE(err, nil)
		t.Assert(s, expect)
	})
}
