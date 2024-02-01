// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp
import (
	"fmt"
	
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gpage"
	)
// GetPage 根据给定的`totalSize`（总数据量）和`pageSize`（每页大小）创建并返回分页对象。
// 注意：为了简化和方便，来自客户端的页参数名始终定义为 gpage.DefaultPageName。
func (r *Request) GetPage(totalSize, pageSize int) *gpage.Page {
	// 它必须具有Router对象属性。
	if r.Router == nil {
		panic("Router object not found")
	}
	var (
		url            = *r.URL
		urlTemplate    = url.Path
		uriHasPageName = false
	)
	// 检查URI中的page变量。
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
	// 检查查询字符串中的page变量。
	if !uriHasPageName {
		values := url.Query()
		values.Set(gpage.DefaultPageName, gpage.DefaultPagePlaceHolder)
		url.RawQuery = values.Encode()
		// 将编码后的“{.page}”替换为原始的“{.page}”。
		url.RawQuery = gstr.Replace(url.RawQuery, "%7B.page%7D", "{.page}")
	}
	if url.RawQuery != "" {
		urlTemplate += "?" + url.RawQuery
	}

	return gpage.New(totalSize, pageSize, r.Get(gpage.DefaultPageName).Int(), urlTemplate)
}
