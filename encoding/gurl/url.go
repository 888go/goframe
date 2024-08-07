// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gurl提供了处理URL的有用API。 md5:3954efb697af4a41
package url类

import (
	"net/url"
	"strings"

	gerror "github.com/888go/goframe/errors/gerror"
)

// X编码将字符串进行转义，以便安全地放置在URL查询中。
// md5:2e139b94de8d8e81
func X编码(文本 string) string {
	return url.QueryEscape(文本)
}

// X解码 执行与 Encode 相反的转换，
// 将形如 "%AB" 的每3字节编码子串转换为其十六进制解码字节 0xAB。
// 如果任何百分号（%）后面没有跟随两个十六进制数字，它将返回一个错误。
// md5:c8ff43c799b800c0
func X解码(文本 string) (string, error) {
	return url.QueryUnescape(文本)
}

// X编码RFC3986 按照 RFC 3986 标准对给定的字符串进行原始URL编码。
// 参考 http://php.net/manual/en/function.rawurlencode.php。
// md5:b116dd32b351afc8
func X编码RFC3986(文本 string) string {
	return strings.ReplaceAll(url.QueryEscape(文本), "+", "%20")
}

// X解码RFC3986 用于解码给定的字符串
// 解码 URL 编码的字符串。
// 参考：http://php.net/manual/zh/function.rawurldecode.php
// md5:ffbb20457d038fe3
func X解码RFC3986(文本 string) (string, error) {
	return url.QueryUnescape(strings.ReplaceAll(文本, "%20", "+"))
}

// X生成URL 生成 URL 编码的查询字符串。
// 参考：http://php.net/manual/zh/function.http-build-query.php。
// md5:f0e4222e29189a30
func X生成URL(查询参数 url.Values) string {
	return 查询参数.Encode()
}

// X解析 解析一个URL并返回其组成部分。
// 参数可选值：-1表示全部；1表示方案（scheme）；2表示主机（host）；4表示端口（port）；
// 8表示用户名（user）；16表示密码（pass）；32表示路径（path）；
// 64表示查询字符串（query）；128表示片段（fragment）。
// 参考 PHP 手册中的函数 parse-url：http://php.net/manual/en/function.parse-url.php。
// md5:ab33f23dd1fe61ca
func X解析(文本 string, 类型标签 int) (map[string]string, error) {
	u, err := url.Parse(文本)
	if err != nil {
		err = gerror.X多层错误并格式化(err, `url.Parse failed for URL "%s"`, 文本)
		return nil, err
	}
	if 类型标签 == -1 {
		类型标签 = 1 | 2 | 4 | 8 | 16 | 32 | 64 | 128
	}
	var components = make(map[string]string)
	if (类型标签 & 1) == 1 {
		components["scheme"] = u.Scheme
	}
	if (类型标签 & 2) == 2 {
		components["host"] = u.Hostname()
	}
	if (类型标签 & 4) == 4 {
		components["port"] = u.Port()
	}
	if (类型标签 & 8) == 8 {
		components["user"] = u.User.Username()
	}
	if (类型标签 & 16) == 16 {
		components["pass"], _ = u.User.Password()
	}
	if (类型标签 & 32) == 32 {
		components["path"] = u.Path
	}
	if (类型标签 & 64) == 64 {
		components["query"] = u.RawQuery
	}
	if (类型标签 & 128) == 128 {
		components["fragment"] = u.Fragment
	}
	return components, nil
}
