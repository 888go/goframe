// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包gurl提供了用于URL处理的有用API。
package gurl

import (
	"net/url"
	"strings"
	
	"github.com/gogf/gf/v2/errors/gerror"
)

// Encode 对字符串进行转义，以便它可以安全地放置在
// URL 查询参数内。
func Encode(str string) string {
	return url.QueryEscape(str)
}

// Decode完成与Encode相反的转换操作，
// 将形式为"%AB"的每个3字节编码子串转化为十六进制解码后的字节0xAB。
// 如果存在百分号（%）后跟随的不是两个十六进制数字，则返回错误。
func Decode(str string) (string, error) {
	return url.QueryUnescape(str)
}

// RawEncode 对给定的字符串进行编码，
// 根据 RFC 3986 进行 URL 编码。
// 参见 http://php.net/manual/en/function.rawurlencode.php.
func RawEncode(str string) string {
	return strings.ReplaceAll(url.QueryEscape(str), "+", "%20")
}

// RawDecode用于解码给定的字符串
// 解码URL编码过的字符串。
// 参考：http://php.net/manual/en/function.rawurldecode.php.
func RawDecode(str string) (string, error) {
	return url.QueryUnescape(strings.ReplaceAll(str, "%20", "+"))
}

// BuildQuery 生成 URL 编码的查询字符串。
// 参考：http://php.net/manual/en/function.http-build-query.php.
func BuildQuery(queryData url.Values) string {
	return queryData.Encode()
}

// ParseURL 解析一个URL并返回其组成部分。
// 参数：
// -1：所有部分；1：方案（如http）；2：主机名；4：端口；8：用户名；16：密码；32：路径；64：查询字符串；128：片段（锚点）。
// 参考：http://php.net/manual/en/function.parse-url.php.
func ParseURL(str string, component int) (map[string]string, error) {
	u, err := url.Parse(str)
	if err != nil {
		err = gerror.Wrapf(err, `url.Parse failed for URL "%s"`, str)
		return nil, err
	}
	if component == -1 {
		component = 1 | 2 | 4 | 8 | 16 | 32 | 64 | 128
	}
	var components = make(map[string]string)
	if (component & 1) == 1 {
		components["scheme"] = u.Scheme
	}
	if (component & 2) == 2 {
		components["host"] = u.Hostname()
	}
	if (component & 4) == 4 {
		components["port"] = u.Port()
	}
	if (component & 8) == 8 {
		components["user"] = u.User.Username()
	}
	if (component & 16) == 16 {
		components["pass"], _ = u.User.Password()
	}
	if (component & 32) == 32 {
		components["path"] = u.Path
	}
	if (component & 64) == 64 {
		components["query"] = u.RawQuery
	}
	if (component & 128) == 128 {
		components["fragment"] = u.Fragment
	}
	return components, nil
}
