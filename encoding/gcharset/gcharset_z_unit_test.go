// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 编码字符集类_test

import (
	"testing"

	gcharset "github.com/888go/goframe/encoding/gcharset"
	gtest "github.com/888go/goframe/test/gtest"
)

var testData = []struct {
	utf8, other, otherEncoding string
}{
	{"Résumé", "Résumé", "utf-8"},
	//{"Résumé", "R\xe9sum\xe9", "latin-1"},
	{"これは漢字です。", "S0\x8c0o0\"oW[g0Y0\x020", "UTF-16LE"},
	{"これは漢字です。", "0S0\x8c0oo\"[W0g0Y0\x02", "UTF-16BE"},
	{"これは漢字です。", "\xfe\xff0S0\x8c0oo\"[W0g0Y0\x02", "UTF-16"},
	{"𝄢𝄞𝄪𝄫", "\xfe\xff\xd8\x34\xdd\x22\xd8\x34\xdd\x1e\xd8\x34\xdd\x2a\xd8\x34\xdd\x2b", "UTF-16"},
		// 这段Go语言注释翻译成中文是：一个包含三个元素的字符串切片，分别是"Hello, world"、"Hello, world"和"ASCII"。 md5:06bc99dbdccc93f3
	{"Gdańsk", "Gda\xf1sk", "ISO-8859-2"},
	{"Ââ Čč Đđ Ŋŋ Õõ Šš Žž Åå Ää", "\xc2\xe2 \xc8\xe8 \xa9\xb9 \xaf\xbf \xd5\xf5 \xaa\xba \xac\xbc \xc5\xe5 \xc4\xe4", "ISO-8859-10"},
	//{"สำหรับ", "\xca\xd3\xcb\xc3\u047a", "ISO-8859-11"},
	{"latviešu", "latvie\xf0u", "ISO-8859-13"},
	{"Seònaid", "Se\xf2naid", "ISO-8859-14"},
	{"€1 is cheap", "\xa41 is cheap", "ISO-8859-15"},
	{"românește", "rom\xe2ne\xbate", "ISO-8859-16"},
	{"nutraĵo", "nutra\xbco", "ISO-8859-3"},
	{"Kalâdlit", "Kal\xe2dlit", "ISO-8859-4"},
	{"русский", "\xe0\xe3\xe1\xe1\xda\xd8\xd9", "ISO-8859-5"},
	{"ελληνικά", "\xe5\xeb\xeb\xe7\xed\xe9\xea\xdc", "ISO-8859-7"},
	{"Kağan", "Ka\xf0an", "ISO-8859-9"},
	{"Résumé", "R\x8esum\x8e", "macintosh"},
	{"Gdańsk", "Gda\xf1sk", "windows-1250"},
	{"русский", "\xf0\xf3\xf1\xf1\xea\xe8\xe9", "windows-1251"},
	{"Résumé", "R\xe9sum\xe9", "windows-1252"},
	{"ελληνικά", "\xe5\xeb\xeb\xe7\xed\xe9\xea\xdc", "windows-1253"},
	{"Kağan", "Ka\xf0an", "windows-1254"},
	{"עִבְרִית", "\xf2\xc4\xe1\xc0\xf8\xc4\xe9\xfa", "windows-1255"},
	{"العربية", "\xc7\xe1\xda\xd1\xc8\xed\xc9", "windows-1256"},
	{"latviešu", "latvie\xf0u", "windows-1257"},
	{"Việt", "Vi\xea\xf2t", "windows-1258"},
	{"สำหรับ", "\xca\xd3\xcb\xc3\u047a", "windows-874"},
	{"русский", "\xd2\xd5\xd3\xd3\xcb\xc9\xca", "KOI8-R"},
	{"українська", "\xd5\xcb\xd2\xc1\xa7\xce\xd3\xd8\xcb\xc1", "KOI8-U"},
	{"Hello 常用國字標準字體表", "Hello \xb1`\xa5\u03b0\xea\xa6r\xbc\u0437\u01e6r\xc5\xe9\xaa\xed", "big5"},
	{"Hello 常用國字標準字體表", "Hello \xb3\xa3\xd3\xc3\x87\xf8\xd7\xd6\x98\xcb\x9c\xca\xd7\xd6\xf3\x77\xb1\xed", "gbk"},
	{"Hello 常用國字標準字體表", "Hello \xb3\xa3\xd3\xc3\x87\xf8\xd7\xd6\x98\xcb\x9c\xca\xd7\xd6\xf3\x77\xb1\xed", "gb18030"},
	{"花间一壶酒，独酌无相亲。", "~{;(<dR;:x>F#,6@WCN^O`GW!#", "GB2312"},
	{"花间一壶酒，独酌无相亲。", "~{;(<dR;:x>F#,6@WCN^O`GW!#", "HZGB2312"},
	{"עִבְרִית", "\x81\x30\xfb\x30\x81\x30\xf6\x34\x81\x30\xf9\x33\x81\x30\xf6\x30\x81\x30\xfb\x36\x81\x30\xf6\x34\x81\x30\xfa\x31\x81\x30\xfb\x38", "gb18030"},
	{"㧯", "\x82\x31\x89\x38", "gb18030"},
	{"㧯", "㧯", "UTF-8"},
	//{"これは漢字です。", "\x82\xb1\x82\xea\x82\xcd\x8a\xbf\x8e\x9a\x82\xc5\x82\xb7\x81B", "SJIS"},
	{"これは漢字です。", "\xa4\xb3\xa4\xec\xa4\u03f4\xc1\xbb\xfa\xa4\u01e4\xb9\xa1\xa3", "EUC-JP"},
}

func TestDecode(t *testing.T) {
	for _, data := range testData {
		str := ""
		str, err := gcharset.Convert("UTF-8", data.otherEncoding, data.other)
		if err != nil {
			t.Errorf("Could not create decoder for %v", err)
			continue
		}

		if str != data.utf8 {
			t.Errorf("Unexpected value: %#v (expected %#v) %v", str, data.utf8, data.otherEncoding)
		}
	}
}

func TestUTF8To(t *testing.T) {
	for _, data := range testData {
		str := ""
		str, err := gcharset.UTF8To(data.otherEncoding, data.utf8)
		if err != nil {
			t.Errorf("Could not create decoder for %v", err)
			continue
		}

		if str != data.other {
			t.Errorf("Unexpected value: %#v (expected %#v) %v", str, data.other, data.otherEncoding)
		}
	}
}

func TestToUTF8(t *testing.T) {
	for _, data := range testData {
		str := ""
		str, err := gcharset.ToUTF8(data.otherEncoding, data.other)
		if err != nil {
			t.Errorf("Could not create decoder for %v", err)
			continue
		}

		if str != data.utf8 {
			t.Errorf("Unexpected value: %#v (expected %#v)", str, data.utf8)
		}
	}
}

func TestEncode(t *testing.T) {
	for _, data := range testData {
		str := ""
		str, err := gcharset.Convert(data.otherEncoding, "UTF-8", data.utf8)
		if err != nil {
			t.Errorf("Could not create decoder for %v", err)
			continue
		}

		if str != data.other {
			t.Errorf("Unexpected value: %#v (expected %#v)", str, data.other)
		}
	}
}

func TestConvert(t *testing.T) {
	srcCharset := "big5"
	src := "Hello \xb1`\xa5\u03b0\xea\xa6r\xbc\u0437\u01e6r\xc5\xe9\xaa\xed"
	dstCharset := "gbk"
	dst := "Hello \xb3\xa3\xd3\xc3\x87\xf8\xd7\xd6\x98\xcb\x9c\xca\xd7\xd6\xf3\x77\xb1\xed"

	str, err := gcharset.Convert(dstCharset, srcCharset, src)
	if err != nil {
		t.Errorf("convert error. %v", err)
		return
	}

	if str != dst {
		t.Errorf("unexpected value:%#v (expected %#v)", str, dst)
	}
}

func TestConvertErr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		srcCharset := "big5"
		dstCharset := "gbk"
		src := "Hello \xb1`\xa5\u03b0\xea\xa6r\xbc\u0437\u01e6r\xc5\xe9\xaa\xed"

		s1, e1 := gcharset.Convert(srcCharset, srcCharset, src)
		t.Assert(e1, nil)
		t.Assert(s1, src)

		s2, e2 := gcharset.Convert(dstCharset, "no this charset", src)
		t.AssertNE(e2, nil)
		t.Assert(s2, src)

		s3, e3 := gcharset.Convert("no this charset", srcCharset, src)
		t.AssertNE(e3, nil)
		t.Assert(s3, src)
	})
}

func TestSupported(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gcharset.Supported("UTF-8"), true)
		t.Assert(gcharset.Supported("UTF-80"), false)
	})
}
