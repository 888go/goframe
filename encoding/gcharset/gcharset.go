// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// Package gcharset implements character-set conversion functionality.
//
// Supported Character Set:
//
// Chinese : GBK/GB18030/GB2312/Big5
//
//
// Korean  : EUCKR
//
// Unicode : UTF-8/UTF-16/UTF-16BE/UTF-16LE
//
// Other   : macintosh/IBM*/Windows*/ISO-*
package gcharset//bm:编码字符集类

import (
	"bytes"
	"context"
	"io"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/intlog"
)

var (
	// Alias for charsets.
	charsetAlias = map[string]string{
		"HZGB2312": "HZ-GB-2312",
		"hzgb2312": "HZ-GB-2312",
		"GB2312":   "HZ-GB-2312",
		"gb2312":   "HZ-GB-2312",
	}
)

// Supported 返回字符集 `charset` 是否受支持。 md5:ecb209536b99e114
// ff:
// charset:
func Supported(charset string) bool {
	return getEncoding(charset) != nil
}

// Convert 将 `src` 字符串的编码从 `srcCharset` 转换为 `dstCharset`，并返回转换后的字符串。
// 如果转换失败，则返回原 `src` 作为 `dst`。
// md5:d579c6167a34081f
// ff:
// dstCharset:
// srcCharset:
// src:
// dst:
// err:
func Convert(dstCharset string, srcCharset string, src string) (dst string, err error) {
	if dstCharset == srcCharset {
		return src, nil
	}
	dst = src
	// 将 `src` 转换为 UTF-8 编码。 md5:345cd013199770a3
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
	// 将UTF-8转换为`dstCharset`。 md5:4caf3880c33fb49d
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

// ToUTF8 将 `src` 字符串的字符集编码从 `srcCharset` 转换为 UTF-8，
// 并返回转换后的字符串。
// md5:ed113e096f11dcee
// ff:
// srcCharset:
// src:
// dst:
// err:
func ToUTF8(srcCharset string, src string) (dst string, err error) {
	return Convert("UTF-8", srcCharset, src)
}

// UTF8To 将 `src` 字符集编码从 UTF-8 转换为 `dstCharset`，
// 并返回转换后的字符串。
// md5:6d376918eb2969a6
// ff:
// dstCharset:
// src:
// dst:
// err:
func UTF8To(dstCharset string, src string) (dst string, err error) {
	return Convert(dstCharset, "UTF-8", src)
}

	// getEncoding 返回与 `charset` 对应的 encoding.Encoding 接口对象。如果 `charset` 不被支持，它将返回 nil。
	// md5:8770abf28a404b1b
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
