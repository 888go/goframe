// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gcharset 提供字符集转换功能的实现。
//
// 支持的字符集：
//
// 中文 : GBK/GB18030/GB2312/Big5
//
// 日文: EUCJP/ISO2022JP/ShiftJIS
//
// 韩文  : EUCKR
//
// Unicode : UTF-8/UTF-16/UTF-16BE/UTF-16LE
//
// 其他   : macintosh/IBM*/Windows*/ISO-*
// 这段注释描述了一个名为 `gcharset` 的 Go 语言包，它实现了字符集转换功能，并列出了该包支持的多种字符集，包括中文、日文、韩文对应的常见编码以及 Unicode 和其他一些通用或特定平台的字符集。
package gcharset

import (
	"bytes"
	"context"
	"io"
	
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
)

var (
	// 对字符集的别名。
	charsetAlias = map[string]string{
		"HZGB2312": "HZ-GB-2312",
		"hzgb2312": "HZ-GB-2312",
		"GB2312":   "HZ-GB-2312",
		"gb2312":   "HZ-GB-2312",
	}
)

// Supported 返回 `charset` 字符集是否被支持。
func Supported(charset string) bool {
	return getEncoding(charset) != nil
}

// Convert 将 `src` 字符串从 `srcCharset` 编码转换为 `dstCharset` 编码，
// 并返回转换后的字符串。
// 如果转换失败，则直接返回原字符串 `src` 作为结果。
func Convert(dstCharset string, srcCharset string, src string) (dst string, err error) {
	if dstCharset == srcCharset {
		return src, nil
	}
	dst = src
	// 将 `src` 转换为 UTF-8。
	if srcCharset != "UTF-8" {
		if e := getEncoding(srcCharset); e != nil {
			tmp, err := io.ReadAll(
				transform.NewReader(bytes.NewReader([]byte(src)), e.NewDecoder()),
			)
			if err != nil {
				return "", gerror.Wrapf(err, `convert string "%s" to utf8 failed`, srcCharset)
			}
			src = string(tmp)
		} else {
			return dst, gerror.NewCodef(gcode.CodeInvalidParameter, `unsupported srcCharset "%s"`, srcCharset)
		}
	}
	// 将UTF-8转换为`dstCharset`。
	if dstCharset != "UTF-8" {
		if e := getEncoding(dstCharset); e != nil {
			tmp, err := io.ReadAll(
				transform.NewReader(bytes.NewReader([]byte(src)), e.NewEncoder()),
			)
			if err != nil {
				return "", gerror.Wrapf(err, `convert string from utf8 to "%s" failed`, dstCharset)
			}
			dst = string(tmp)
		} else {
			return dst, gerror.NewCodef(gcode.CodeInvalidParameter, `unsupported dstCharset "%s"`, dstCharset)
		}
	} else {
		dst = src
	}
	return dst, nil
}

// ToUTF8 将 `src` 字符集从 `srcCharset` 转换为 UTF-8，
// 并返回转换后的字符串。
func ToUTF8(srcCharset string, src string) (dst string, err error) {
	return Convert("UTF-8", srcCharset, src)
}

// UTF8To将`src`的字符集编码从UTF-8转换为`dstCharset`，
// 并返回转换后的字符串。
func UTF8To(dstCharset string, src string) (dst string, err error) {
	return Convert(dstCharset, "UTF-8", src)
}

// getEncoding 函数根据 `charset` 参数返回对应的 encoding.Encoding 接口对象。
// 如果 `charset` 不被支持，则返回 nil。
func getEncoding(charset string) encoding.Encoding {
	if c, ok := charsetAlias[charset]; ok {
		charset = c
	}
	enc, err := ianaindex.MIB.Encoding(charset)
	if err != nil {
		intlog.Errorf(context.TODO(), `%+v`, err)
	}
	return enc
}
