// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	
	"github.com/888go/goframe/container/glist"
	"github.com/888go/goframe/encoding/gurl"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/text/gregex"
)

// handlerCacheItem 是一个仅用于内部路由器搜索缓存的项目。
type handlerCacheItem struct {
	parsedItems []*X路由解析
	serveItem   *X路由解析
	hasHook     bool
	hasServe    bool
}

// serveHandlerKey 为路由器创建并返回一个处理程序键。
func (s *X服务) serveHandlerKey(method, path, domain string) string {
	if len(domain) > 0 {
		domain = "@" + domain
	}
	if method == "" {
		return path + strings.ToLower(domain)
	}
	return strings.ToUpper(method) + ":" + path + strings.ToLower(domain)
}

// getHandlersWithCache 根据给定的请求，搜索具有缓存功能的路由器项。
func (s *X服务) getHandlersWithCache(r *X请求) (parsedItems []*X路由解析, serveItem *X路由解析, hasHook, hasServe bool) {
	var (
		ctx    = r.Context别名()
		method = r.Method
		path   = r.URL.Path
		host   = r.X取主机名()
	)
// 举例说明：
// 情况1：
//   GET /net/http
//   r.URL.Path    : /net/http
//   r.URL.RawPath : （空字符串）
// 情况2：
//   GET /net%2Fhttp
//   r.URL.Path    : /net/http
//   r.URL.RawPath : /net%2Fhttp
// 在上述情况下，r 是一个 *http.Request 类型的对象，
// r.URL.Path 表示已解码的请求路径（在情况1和情况2中都显示为"/net/http"）。
// 而 r.URL.RawPath 则表示原始编码格式的请求路径，
// 在情况1中由于未进行编码，所以是空字符串；
// 在情况2中，原始路径经过URL编码，"%2F"代表"/"，因此RawPath为"/net%2Fhttp"。
	if r.URL.RawPath != "" {
		path = r.URL.RawPath
	}
// 特殊处理HTTP方法OPTIONS。
// 它会按照请求的方法而非OPTIONS方法去搜索处理器。
	if method == http.MethodOptions {
		if v := r.Request.Header.Get("Access-Control-Request-Method"); v != "" {
			method = v
		}
	}
	// 搜索并缓存路由处理器。
	if xUrlPath := r.Header.Get(HeaderXUrlPath); xUrlPath != "" {
		path = xUrlPath
	}
	var handlerCacheKey = s.serveHandlerKey(method, path, host)
	value, err := s.serveCache.X取值或设置值_函数(ctx, handlerCacheKey, func(ctx context.Context) (interface{}, error) {
		parsedItems, serveItem, hasHook, hasServe = s.searchHandlers(method, path, host)
		if parsedItems != nil {
			return &handlerCacheItem{parsedItems, serveItem, hasHook, hasServe}, nil
		}
		return nil, nil
	}, routeCacheDuration)
	if err != nil {
		intlog.Errorf(ctx, `%+v`, err)
	}
	if value != nil {
		item := value.X取值().(*handlerCacheItem)
		return item.parsedItems, item.serveItem, item.hasHook, item.hasServe
	}
	return
}

// searchHandlers 函数根据给定的参数获取并返回路由器。
// 注意，返回的路由器中包含了服务处理程序、中间件处理程序以及钩子处理程序。
func (s *X服务) searchHandlers(method, path, domain string) (parsedItems []*X路由解析, serveItem *X路由解析, hasHook, hasServe bool) {
	if len(path) == 0 {
		return nil, nil, false, false
	}
// 在出现双斜线（/）的URI情况下，例如：
// /user//index,/user/index,/user//index//
// （注：这些示例中的双斜线表示连续的路径分隔符，Go语言中会进行规范化处理，将多个连续的斜线视为一个斜线。）
// ```go
// 对于如下的含有双斜线（/）的URI情况：
// 比如 "/user//index", "/user/index", "/user//index//" 等
	var previousIsSep = false
	for i := 0; i < len(path); {
		if path[i] == '/' {
			if previousIsSep {
				path = path[:i] + path[i+1:]
				continue
			} else {
				previousIsSep = true
			}
		} else {
			previousIsSep = false
		}
		i++
	}
	// 将URL.path拆分为多个部分。
	var array []string
	if strings.EqualFold("/", path) {
		array = []string{"/"}
	} else {
		array = strings.Split(path[1:], "/")
	}
	var (
		lastMiddlewareElem    *链表类.Element
		parsedItemList        = 链表类.New()
		repeatHandlerCheckMap = make(map[int]struct{}, 16)
	)

// 当进行迭代时，默认域具有最高优先级。
// 如果您想了解serveTree的结构，请参阅doSetHandler。
	for _, domainItem := range []string{DefaultDomainName, domain} {
		p, ok := s.serveTree[domainItem]
		if !ok {
			continue
		}
		// 创建一个容量为16的list数组
		lists := make([]*链表类.List, 0, 16)
		for i, part := range array {
			// 将每个节点的所有列表添加到list数组中。
			if v, ok := p.(map[string]interface{})["*list"]; ok {
				lists = append(lists, v.(*链表类.List))
			}
			if v, ok := p.(map[string]interface{})[part]; ok {
				// 通过指定的键名循环到下一个节点。
				p = v
				if i == len(array)-1 {
					if v, ok := p.(map[string]interface{})["*list"]; ok {
						lists = append(lists, v.(*链表类.List))
						break
					}
				}
			} else if v, ok := p.(map[string]interface{})["*fuzz"]; ok {
				// 通过模糊节点项循环到下一个节点。
				p = v
			}
			if i == len(array)-1 {
// 此处同时检查模糊匹配项，
// 例如：规则 "/user/*action" 匹配到 "/user" 的情况。
				if v, ok := p.(map[string]interface{})["*fuzz"]; ok {
					p = v
				}
				// 叶子节点必须包含一个列表项。它将该列表添加到列表数组中。
				if v, ok := p.(map[string]interface{})["*list"]; ok {
					lists = append(lists, v.(*链表类.List))
				}
			}
		}

// 好的，让我们循环遍历结果列表数组，
// 将处理器项添加到结果处理器结果数组中。
// 由于列表数组的尾部具有最高的优先级，
// 因此从尾部到头部遍历列表数组。
		for i := len(lists) - 1; i >= 0; i-- {
			for e := lists[i].Front(); e != nil; e = e.Next() {
				item := e.Value.(*X路由处理函数)
// 过滤重复的处理器项，特别是中间件和钩子处理器。
// 这是必要的，请不要移除这个检查逻辑，除非你确实了解其必要性。
//
// `repeatHandlerCheckMap` 用于在搜索处理器时进行重复处理器项过滤。
// 由于存在模糊节点，且模糊节点同时包含子节点和子列表节点，在子节点和子列表节点中可能存在重复的处理器项。这里使用处理器项ID来识别注册时相同的处理器项。
//
// 同样的处理器项是指在同一个 doSetHandler 函数中注册的项。
// 注意，一个处理器函数（中间件或钩子函数）可能通过 function doSetHandler 以不同的处理器项形式被多次注册，它们具有不同的处理器项ID。
//
// 特别注意，处理器函数可能作为不同的处理器项被多次注册。
				if _, isRepeatedHandler := repeatHandlerCheckMap[item.Id]; isRepeatedHandler {
					continue
				} else {
					repeatHandlerCheckMap[item.Id] = struct{}{}
				}
// 服务处理程序只能一次性添加到处理器数组中。
// 列表中的第一条路由项比其余项具有更高的优先级。
// 这种忽略方式可以实现路由覆盖功能。
				if hasServe {
					switch item.Type {
					case HandlerTypeHandler, HandlerTypeObject:
						continue
					}
				}
				if item.X路由.Method == defaultMethod || item.X路由.Method == method {
					// 注意这个规则：没有模糊规则时，匹配项的长度为1
					if match, err := 正则类.X匹配文本(item.X路由.X正则路由规则, path); err == nil && len(match) > 0 {
						parsedItem := &X路由解析{item, nil}
// 如果规则中包含模糊名称，
// 则需要解析URL以获取这些名称的值。
						if len(item.X路由.X路由参数名称) > 0 {
							if len(match) > len(item.X路由.X路由参数名称) {
								parsedItem.X路由值 = make(map[string]string)
								// 如果存在重复的名字，它只会覆盖相同的那个。
								for i, name := range item.X路由.X路由参数名称 {
									parsedItem.X路由值[name], _ = url类.X解码(match[i+1])
								}
							}
						}
						switch item.Type {
						// 服务处理程序只能添加一次。
						case HandlerTypeHandler, HandlerTypeObject:
							hasServe = true
							serveItem = parsedItem
							parsedItemList.PushBack(parsedItem)

// 中文注释：
// 中间件在服务处理程序之前插入。
// 如果存在多个中间件，它们会按照注册顺序插入到结果列表中。
// 中间件也会按照注册时的顺序执行。
// 这段Go语言代码注释翻译成中文后为：
// ```markdown
// 中间件在实际服务处理程序之前被插入。
// 若存在多个中间件，它们将根据注册顺序依次插入到结果列表中。
// 同样地，这些中间件也是按照其注册时的顺序来执行的。
						case HandlerTypeMiddleware:
							if lastMiddlewareElem == nil {
								lastMiddlewareElem = parsedItemList.PushFront(parsedItem)
							} else {
								lastMiddlewareElem = parsedItemList.InsertAfter(lastMiddlewareElem, parsedItem)
							}

						// HOOK 处理函数，只需将其推回列表中。
						case HandlerTypeHook:
							hasHook = true
							parsedItemList.PushBack(parsedItem)

						default:
							panic(错误类.X创建并格式化(`invalid handler type %s`, item.Type))
						}
					}
				}
			}
		}
	}
	if parsedItemList.Len() > 0 {
		var index = 0
		parsedItems = make([]*X路由解析, parsedItemList.Len())
		for e := parsedItemList.Front(); e != nil; e = e.Next() {
			parsedItems[index] = e.Value.(*X路由解析)
			index++
		}
	}
	return
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (item X路由处理函数) MarshalJSON() ([]byte, error) {
	switch item.Type {
	case HandlerTypeHook:
		return json.Marshal(
			fmt.Sprintf(
				`%s %s:%s (%s)`,
				item.X路由.Uri,
				item.X路由.Domain,
				item.X路由.Method,
				item.Hook名称,
			),
		)
	case HandlerTypeMiddleware:
		return json.Marshal(
			fmt.Sprintf(
				`%s %s:%s (MIDDLEWARE)`,
				item.X路由.Uri,
				item.X路由.Domain,
				item.X路由.Method,
			),
		)
	default:
		return json.Marshal(
			fmt.Sprintf(
				`%s %s:%s`,
				item.X路由.Uri,
				item.X路由.Domain,
				item.X路由.Method,
			),
		)
	}
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (item X路由解析) MarshalJSON() ([]byte, error) {
	return json.Marshal(item.Handler)
}
