// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package goai

import (
	gstr "github.com/888go/goframe/text/gstr"
)

// XExtensions 存储了 `x-` 自定义扩展。 md5:e19838946aa45df7
type XExtensions map[string]string

func (oai *OpenApiV3) tagMapToXExtensions(tagMap map[string]string, extensions XExtensions) {
	for k, v := range tagMap {
		if gstr.X开头判断(k, "x-") || gstr.X开头判断(k, "X-") {
			extensions[k] = v
		}
	}
}
