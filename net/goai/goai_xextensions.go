// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package goai

import (
	"github.com/888go/goframe/text/gstr"
)

// XExtensions 用于存储以 `x-` 开头的自定义扩展。
type XExtensions map[string]string

func (oai *OpenApiV3) tagMapToXExtensions(tagMap map[string]string, extensions XExtensions) {
	for k, v := range tagMap {
		if 文本类.X开头判断(k, "x-") || 文本类.X开头判断(k, "X-") {
			extensions[k] = v
		}
	}
}
