// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

import (
	"fmt"

	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gpage"
)

// GetPage 根据给定的`totalSize`和`pageSize`创建并返回分页对象。
// 请注意，来自客户端的分页参数名称常量定义为gpage.DefaultPageName，以简化和方便使用。
// md5:4d3bd97d937b25b8
func (r *Request) GetPage(totalSize, pageSize int) *gpage.Page {
	// 它必须有 Router 对象属性。 md5:8cc7be190bf78663
	if r.Router == nil {
		panic("Router object not found")
	}
	var (
		url            = *r.URL
		urlTemplate    = url.Path
		uriHasPageName = false
	)
	// 检查URI中的page变量。 md5:7e5a6c48958e8612
	if len(r.Router.RegNames) > 0 {
		for _, name := range r.Router.RegNames {
			if name == gpage.DefaultPageName {
				uriHasPageName = true
				break
			}
		}
		if uriHasPageName {
			if match, err := gregex.MatchString(r.Router.RegRule, url.Path); err == nil && len(match) > 0 {
				if len(match) > len(r.Router.RegNames) {
					urlTemplate = r.Router.Uri
					for i, name := range r.Router.RegNames {
						rule := fmt.Sprintf(`[:\*]%s|\{%s\}`, name, name)
						if name == gpage.DefaultPageName {
							urlTemplate, err = gregex.ReplaceString(rule, gpage.DefaultPagePlaceHolder, urlTemplate)
						} else {
							urlTemplate, err = gregex.ReplaceString(rule, match[i+1], urlTemplate)
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
		values.Set(gpage.DefaultPageName, gpage.DefaultPagePlaceHolder)
		url.RawQuery = values.Encode()
		// 将编码的 "{.page}" 替换为原始的 "{.page}"。 md5:755f7a81273710ea
		url.RawQuery = gstr.Replace(url.RawQuery, "%7B.page%7D", "{.page}")
	}
	if url.RawQuery != "" {
		urlTemplate += "?" + url.RawQuery
	}

	return gpage.New(totalSize, pageSize, r.Get(gpage.DefaultPageName).Int(), urlTemplate)
}
