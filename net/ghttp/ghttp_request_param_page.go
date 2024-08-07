// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"fmt"

	gregex "github.com/888go/goframe/text/gregex"
	gstr "github.com/888go/goframe/text/gstr"
	gpage "github.com/888go/goframe/util/gpage"
)

// X取分页类 根据给定的`totalSize`和`pageSize`创建并返回分页对象。
// 请注意，来自客户端的分页参数名称常量定义为gpage.DefaultPageName，以简化和方便使用。
// md5:4d3bd97d937b25b8
func (r *Request) X取分页类(总数据量, 每页大小 int) *gpage.Page {
		// 它必须有 Router 对象属性。 md5:8cc7be190bf78663
	if r.X路由 == nil {
		panic("Router object not found")
	}
	var (
		url            = *r.URL
		urlTemplate    = url.Path
		uriHasPageName = false
	)
		// 检查URI中的page变量。 md5:7e5a6c48958e8612
	if len(r.X路由.X路由参数名称) > 0 {
		for _, name := range r.X路由.X路由参数名称 {
			if name == gpage.X常量_默认页面名称 {
				uriHasPageName = true
				break
			}
		}
		if uriHasPageName {
			if match, err := gregex.X匹配文本(r.X路由.X正则路由规则, url.Path); err == nil && len(match) > 0 {
				if len(match) > len(r.X路由.X路由参数名称) {
					urlTemplate = r.X路由.Uri
					for i, name := range r.X路由.X路由参数名称 {
						rule := fmt.Sprintf(`[:\*]%s|\{%s\}`, name, name)
						if name == gpage.X常量_默认页面名称 {
							urlTemplate, err = gregex.X替换文本(rule, gpage.X常量_默认模板占位符, urlTemplate)
						} else {
							urlTemplate, err = gregex.X替换文本(rule, match[i+1], urlTemplate)
						}
						if err != nil {
							panic(err)
						}
					}
				}
			} else {
				panic(err)
			}
		}
	}
		// 检查查询字符串中的page变量。 md5:4c851fb33e38ee9b
	if !uriHasPageName {
		values := url.Query()
		values.Set(gpage.X常量_默认页面名称, gpage.X常量_默认模板占位符)
		url.RawQuery = values.Encode()
				// 将编码的 "{.page}" 替换为原始的 "{.page}"。 md5:755f7a81273710ea
		url.RawQuery = gstr.X替换(url.RawQuery, "%7B.page%7D", "{.page}")
	}
	if url.RawQuery != "" {
		urlTemplate += "?" + url.RawQuery
	}

	return gpage.X创建(总数据量, 每页大小, r.Get别名(gpage.X常量_默认页面名称).X取整数(), urlTemplate)
}
