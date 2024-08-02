// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	glist "github.com/888go/goframe/container/glist"
	gurl "github.com/888go/goframe/encoding/gurl"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/json"
	gregex "github.com/888go/goframe/text/gregex"
)

// handlerCacheItem 是仅用于内部路由器搜索缓存的项。 md5:bff6700a37e67c6b
type handlerCacheItem struct {
	parsedItems []*HandlerItemParsed
	serveItem   *HandlerItemParsed
	hasHook     bool
	hasServe    bool
}

// serveHandlerKey 为路由器创建并返回一个处理器键。 md5:a4cf69fa7df9d5ac
func (s *Server) serveHandlerKey(method, path, domain string) string {
	if len(domain) > 0 {
		domain = "@" + domain
	}
	if method == "" {
		return path + strings.ToLower(domain)
	}
	return strings.ToUpper(method) + ":" + path + strings.ToLower(domain)
}

// getHandlersWithCache 为给定的请求搜索具有缓存功能的路由项。 md5:00b96b129dd9a5f8
func (s *Server) getHandlersWithCache(r *Request) (parsedItems []*HandlerItemParsed, serveItem *HandlerItemParsed, hasHook, hasServe bool) {
	var (
		ctx    = r.Context()
		method = r.Method
		path   = r.URL.Path
		host   = r.GetHost()
	)
	// 在以下情况中，例如：
	// 情况1：
	// 		GET /net/http
	// 		r.URL.Path    : /net/http
	// 		r.URL.RawPath : （空字符串）
	// 情况2：
	// 		GET /net%2Fhttp
	// 		r.URL.Path    : /net/http
	// 		r.URL.RawPath : /net%2Fhttp
	// md5:97750eaa6ac9d07d
	if r.URL.RawPath != "" {
		path = r.URL.RawPath
	}
	// 专门处理 HTTP 方法 OPTIONS。
	// 它会使用请求方法搜索处理器，而不是 OPTIONS 方法。
	// md5:2704b13524189224
	if method == http.MethodOptions {
		if v := r.Request.Header.Get("Access-Control-Request-Method"); v != "" {
			method = v
		}
	}
		// 搜索并缓存路由器处理器。 md5:7263da0c149e9280
	if xUrlPath := r.Header.Get(HeaderXUrlPath); xUrlPath != "" {
		path = xUrlPath
	}
	var handlerCacheKey = s.serveHandlerKey(method, path, host)
	value, err := s.serveCache.GetOrSetFunc(ctx, handlerCacheKey, func(ctx context.Context) (interface{}, error) {
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
		item := value.Val().(*handlerCacheItem)
		return item.parsedItems, item.serveItem, item.hasHook, item.hasServe
	}
	return
}

// searchHandlers 根据给定的参数检索并返回路由器。
// 注意，返回的路由器包含了服务处理程序、中间件处理程序和钩子处理程序。
// md5:c8f076ede0fbe806
func (s *Server) searchHandlers(method, path, domain string) (parsedItems []*HandlerItemParsed, serveItem *HandlerItemParsed, hasHook, hasServe bool) {
	if len(path) == 0 {
		return nil, nil, false, false
	}
	// 对于包含连续'/'的URI，例如：
	// /user	//index, 	//user/index, 	//user	//index	//
	// md5:fb272e4928c6b465
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
		// 将URL.path分割成多个部分。 md5:421153d9f0413872
	var array []string
	if strings.EqualFold("/", path) {
		array = []string{"/"}
	} else {
		array = strings.Split(path[1:], "/")
	}
	var (
		lastMiddlewareElem    *glist.Element
		parsedItemList        = glist.New()
		repeatHandlerCheckMap = make(map[int]struct{}, 16)
	)

	// 当迭代时，默认域具有最高优先级。如果您想了解serveTree的结构，请参阅doSetHandler。
	// md5:8bc20bbd07335cfd
	for _, domainItem := range []string{DefaultDomainName, domain} {
		p, ok := s.serveTree[domainItem]
		if !ok {
			continue
		}
				// 创建一个容量为16的列表数组。 md5:9ce7c6d246550dea
		lists := make([]*glist.List, 0, 16)
		for i, part := range array {
						// 将每个节点的所有列表添加到列表数组中。 md5:82a101859541f52e
			if v, ok := p.(map[string]interface{})["*list"]; ok {
				lists = append(lists, v.(*glist.List))
			}
			if v, ok := p.(map[string]interface{})[part]; ok {
								// 通过指定的关键字名称，循环到下一个节点。 md5:e9bddd4258b62329
				p = v
				if i == len(array)-1 {
					if v, ok := p.(map[string]interface{})["*list"]; ok {
						lists = append(lists, v.(*glist.List))
						break
					}
				}
			} else if v, ok := p.(map[string]interface{})["*fuzz"]; ok {
								// 通过模糊节点项遍历到下一个节点。 md5:fe1b87e1d17d2d0f
				p = v
			}
			if i == len(array)-1 {
				// 这里同时检查模糊项，
				// 适用于诸如规则情况："/user/*action" 匹配到 "/user"。
				// md5:89d31460cfdbd8e6
				if v, ok := p.(map[string]interface{})["*fuzz"]; ok {
					p = v
				}
								// 叶子节点必须有一个列表项。它将这个列表添加到列表数组中。 md5:dbc074656c73362f
				if v, ok := p.(map[string]interface{})["*list"]; ok {
					lists = append(lists, v.(*glist.List))
				}
			}
		}

		// 好的，让我们遍历结果列表数组，将处理项添加到结果处理器结果数组中。由于列表数组的尾部优先级最高，所以我们从数组尾部开始向前遍历。
		// md5:1f7f116128551404
		for i := len(lists) - 1; i >= 0; i-- {
			for e := lists[i].Front(); e != nil; e = e.Next() {
				item := e.Value.(*HandlerItem)
				// 过滤重复的处理器项，特别是中间件和钩子处理器。
				// 这是必要的，除非你非常清楚为什么需要移除这个检查逻辑，否则请不要删除。
				//
				// `repeatHandlerCheckMap` 用于在搜索处理器时进行重复处理器过滤。由于存在模糊节点，这些模糊节点既有子节点也有子列表节点，因此可能会在子节点和子列表节点中出现重复的处理器项。
				//
				// 同一个处理器项是指使用 `doSetHandler` 函数在同一函数中注册的处理器。需要注意的是，一个处理器函数（中间件或钩子函数）可能通过 `doSetHandler` 函数以不同的处理器项多次注册，并且它们有不同的处理器项 ID。
				//
				// 另外需要注意，同一种处理器函数可能由于不同的处理目的而被多次注册为不同的处理器项。
				// md5:6e4536c4e013b86a
				if _, isRepeatedHandler := repeatHandlerCheckMap[item.Id]; isRepeatedHandler {
					continue
				} else {
					repeatHandlerCheckMap[item.Id] = struct{}{}
				}
				// 服务处理程序只能添加到处理器数组中一次。
				// 列表中的第一个路由项比其余项具有更高的优先级。
				// 此忽略功能可以实现路由覆盖功能。
				// md5:6e93290e1cdad8d9
				if hasServe {
					switch item.Type {
					case HandlerTypeHandler, HandlerTypeObject:
						continue
					}
				}
				if item.Router.Method == defaultMethod || item.Router.Method == method {
										// 注意没有模糊规则的规则：match 的长度等于 1. md5:c26d1818ce3f384e
					if match, err := gregex.MatchString(item.Router.RegRule, path); err == nil && len(match) > 0 {
						parsedItem := &HandlerItemParsed{item, nil}
						// 如果规则包含模糊名称（fuzzy names），
						// 需要对URL进行切分以获取名称的值。
						// md5:022aca8d52d2dc1f
						if len(item.Router.RegNames) > 0 {
							if len(match) > len(item.Router.RegNames) {
								parsedItem.Values = make(map[string]string)
																//如果有重复的名称，它就会覆盖相同的名称。 md5:afb894e9dbad1062
								for i, name := range item.Router.RegNames {
									parsedItem.Values[name], _ = gurl.Decode(match[i+1])
								}
							}
						}
						switch item.Type {
												// 服务处理程序只能添加一次。 md5:fef3170c186d44cb
						case HandlerTypeHandler, HandlerTypeObject:
							hasServe = true
							serveItem = parsedItem
							parsedItemList.PushBack(parsedItem)

						// 中间件在服务处理器之前插入。
						// 如果有多个中间件，它们会按照注册的顺序插入到结果列表中。
						// 中间件也会按照注册的顺序执行。
						// md5:3ae9ef9a044965f3
						case HandlerTypeMiddleware:
							if lastMiddlewareElem == nil {
								lastMiddlewareElem = parsedItemList.PushFront(parsedItem)
							} else {
								lastMiddlewareElem = parsedItemList.InsertAfter(lastMiddlewareElem, parsedItem)
							}

												// HOOK处理器，只是将其推回列表。 md5:5c3afbdb8ce6826c
						case HandlerTypeHook:
							hasHook = true
							parsedItemList.PushBack(parsedItem)

						default:
							panic(gerror.Newf(`invalid handler type %s`, item.Type))
						}
					}
				}
			}
		}
	}
	if parsedItemList.Len() > 0 {
		var index = 0
		parsedItems = make([]*HandlerItemParsed, parsedItemList.Len())
		for e := parsedItemList.Front(); e != nil; e = e.Next() {
			parsedItems[index] = e.Value.(*HandlerItemParsed)
			index++
		}
	}
	return
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (item HandlerItem) MarshalJSON() ([]byte, error) {
	switch item.Type {
	case HandlerTypeHook:
		return json.Marshal(
			fmt.Sprintf(
				`%s %s:%s (%s)`,
				item.Router.Uri,
				item.Router.Domain,
				item.Router.Method,
				item.HookName,
			),
		)
	case HandlerTypeMiddleware:
		return json.Marshal(
			fmt.Sprintf(
				`%s %s:%s (MIDDLEWARE)`,
				item.Router.Uri,
				item.Router.Domain,
				item.Router.Method,
			),
		)
	default:
		return json.Marshal(
			fmt.Sprintf(
				`%s %s:%s`,
				item.Router.Uri,
				item.Router.Domain,
				item.Router.Method,
			),
		)
	}
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (item HandlerItemParsed) MarshalJSON() ([]byte, error) {
	return json.Marshal(item.Handler)
}
