// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"fmt"
	
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gpage"
)

// GetPage 根据给定的`totalSize`（总数据量）和`pageSize`（每页大小）创建并返回分页对象。
// 注意：为了简化和方便，来自客户端的页参数名始终定义为 gpage.DefaultPageName。
func (r *X请求) X取分页类(总数据量, 每页大小 int) *分页类.Page {
	// 它必须具有Router对象属性。
	if r.X路由 == nil {
		panic("Router object not found")
	}
	var (
		url            = *r.URL
		urlTemplate    = url.Path
		uriHasPageName = false
	)
	// 检查URI中的page变量。
	if len(r.X路由.X路由参数名称) > 0 {
		for _, name := range r.X路由.X路由参数名称 {
			if name == 分页类.X常量_默认页面名称 {
				uriHasPageName = true
				break
			}
		}
		if uriHasPageName {
			if match, err := 正则类.X匹配文本(r.X路由.X正则路由规则, url.Path); err == nil && len(match) > 0 {
				if len(match) > len(r.X路由.X路由参数名称) {
					urlTemplate = r.X路由.Uri
					for i, name := range r.X路由.X路由参数名称 {
						rule := fmt.Sprintf(`[:\*]%s|\{%s\}`, name, name)
						if name == 分页类.X常量_默认页面名称 {
							urlTemplate, err = 正则类.X替换文本(rule, 分页类.X常量_默认模板占位符, urlTemplate)
						} else {
							urlTemplate, err = 正则类.X替换文本(rule, match[i+1], urlTemplate)
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
	// 检查查询字符串中的page变量。
	if !uriHasPageName {
		values := url.Query()
		values.Set(分页类.X常量_默认页面名称, 分页类.X常量_默认模板占位符)
		url.RawQuery = values.Encode()
		// 将编码后的“{.page}”替换为原始的“{.page}”。
		url.RawQuery = 文本类.X替换(url.RawQuery, "%7B.page%7D", "{.page}")
	}
	if url.RawQuery != "" {
		urlTemplate += "?" + url.RawQuery
	}

	return 分页类.X创建(总数据量, 每页大小, r.Get别名(分页类.X常量_默认页面名称).X取整数(), urlTemplate)
}
