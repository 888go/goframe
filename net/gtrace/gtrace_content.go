// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtrace

import (
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/encoding/gcompress"

	"github.com/gogf/gf/v2/text/gstr"
)

// SafeContentForHttp 通过 `MaxContentLogSize` 剪切并返回给定内容。如果内容大小超过 `MaxContentLogSize`，它会在结果尾部添加字符串 `...`。
// md5:ba3a657223c70042
// ff:
// data:
// header:
func SafeContentForHttp(data []byte, header http.Header) (string, error) {
	var err error
	if gzipAccepted(header) {
		if data, err = gcompress.UnGzip(data); err != nil {
			return string(data), err
		}
	}

	return SafeContent(data), nil
}

// SafeContent 通过 `MaxContentLogSize` 剪切并返回给定的内容。如果内容大小超过 `MaxContentLogSize`，则在结果的尾部添加字符串 `...`。
// md5:7ea5429876707ef6
// ff:
// data:
func SafeContent(data []byte) string {
	content := string(data)
	if gstr.LenRune(content) > MaxContentLogSize() {
		content = gstr.StrLimitRune(content, MaxContentLogSize(), "...")
	}

	return content
}

// gzipAccepted 返回客户端是否接受gzip编码的内容。 md5:19ef5390dbb76f53
func gzipAccepted(header http.Header) bool {
	a := header.Get("Content-Encoding")
	parts := strings.Split(a, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "gzip" || strings.HasPrefix(part, "gzip;") {
			return true
		}
	}

	return false
}
