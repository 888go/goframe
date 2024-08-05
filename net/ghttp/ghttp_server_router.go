// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

import (
	"context"
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/gogf/gf/v2/container/glist"
	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/debug/gdebug"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/consts"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/gogf/gf/v2/util/gtag"
)

var (
		// handlerIdGenerator 是处理器项ID生成器。 md5:0a9e55b1609ba4df
	handlerIdGenerator = gtype.NewInt()
)

// routerMapKey 为给定参数创建并返回一个唯一的路由键。这个键用于 `Server.routerMap` 属性，主要用于检查重复的路由注册。
// md5:0a5f0d744a55d4ed
func (s *Server) routerMapKey(hook HookName, method, path, domain string) string {
	return string(hook) + "%" + s.serveHandlerKey(method, path, domain)
}

// parsePattern 将给定的模式解析为域名、方法和路径变量。 md5:9f5177d72b0e5cf6
func (s *Server) parsePattern(pattern string) (domain, method, path string, err error) {
	path = strings.TrimSpace(pattern)
	domain = DefaultDomainName
	method = defaultMethod
	if array, err := gregex.MatchString(`([a-zA-Z]+):(.+)`, pattern); len(array) > 1 && err == nil {
		path = strings.TrimSpace(array[2])
		if v := strings.TrimSpace(array[1]); v != "" {
			method = v
		}
	}
	if array, err := gregex.MatchString(`(.+)@([\w\.\-]+)`, path); len(array) > 1 && err == nil {
		path = strings.TrimSpace(array[1])
		if v := strings.TrimSpace(array[2]); v != "" {
			domain = v
		}
	}
	if path == "" {
		err = gerror.NewCode(gcode.CodeInvalidParameter, "invalid pattern: URI should not be empty")
	}
	if path != "/" {
		path = strings.TrimRight(path, "/")
	}
	return
}

type setHandlerInput struct {
	Prefix      string
	Pattern     string
	HandlerItem *HandlerItem
}

// setHandler 根据给定的处理器和模式创建路由项，并将处理器注册到路由树中。
// 路由树可以看作是一个多层哈希表，请参考下面代码中的注释。
// 此函数在服务器启动时被调用，此时对性能要求不高。真正重要的是设计良好的
// 路由存储结构，以便在处理请求时能够高效地进行路由查找。
// md5:325f9f3a1c077ca7
func (s *Server) setHandler(ctx context.Context, in setHandlerInput) {
	var (
		prefix  = in.Prefix
		pattern = in.Pattern
		handler = in.HandlerItem
	)
	if handler.Name == "" {
		handler.Name = runtime.FuncForPC(handler.Info.Value.Pointer()).Name()
	}
	if handler.Source == "" {
		_, file, line := gdebug.CallerWithFilter([]string{consts.StackFilterKeyForGoFrame})
		handler.Source = fmt.Sprintf(`%s:%d`, file, line)
	}
	domain, method, uri, err := s.parsePattern(pattern)
	if err != nil {
		s.Logger().Fatalf(ctx, `invalid pattern "%s", %+v`, pattern, err)
		return
	}
	// ====================================================================================
	// 根据请求结构体中的元信息更改注册的路由。
	// 它支持使用字符 ',' 连接的多个方法。
	// ====================================================================================
	// md5:1ec1db24aded2c53
	if handler.Info.Type != nil && handler.Info.Type.NumIn() == 2 {
		var objectReq = reflect.New(handler.Info.Type.In(1))
		if v := gmeta.Get(objectReq, gtag.Path); !v.IsEmpty() {
			uri = v.String()
		}
		if v := gmeta.Get(objectReq, gtag.Domain); !v.IsEmpty() {
			domain = v.String()
		}
		if v := gmeta.Get(objectReq, gtag.Method); !v.IsEmpty() {
			method = v.String()
		}
				// 使用字符`,`连接的多个方法注册。 md5:8edb2f5feed892c9
		if gstr.Contains(method, ",") {
			methods := gstr.SplitAndTrim(method, ",")
			for _, v := range methods {
								// 每个方法都有自己的处理器。 md5:006ab83dd8178a73
				clonedHandler := *handler
				s.doSetHandler(ctx, &clonedHandler, prefix, uri, pattern, v, domain)
			}
			return
		}
				// 将`all`转换为`ALL`。 md5:85c2f9ce5460fdd6
		if gstr.Equal(method, defaultMethod) {
			method = defaultMethod
		}
	}
	s.doSetHandler(ctx, handler, prefix, uri, pattern, method, domain)
}

func (s *Server) doSetHandler(
	ctx context.Context, handler *HandlerItem,
	prefix, uri, pattern, method, domain string,
) {
	if !s.isValidMethod(method) {
		s.Logger().Fatalf(
			ctx,
			`invalid method value "%s", should be in "%s" or "%s"`,
			method, supportedHttpMethods, defaultMethod,
		)
	}
		// URI功能的前缀。 md5:7ec2c5614dbd89a6
	if prefix != "" {
		uri = prefix + "/" + strings.TrimLeft(uri, "/")
	}
	uri = strings.TrimRight(uri, "/")
	if uri == "" {
		uri = "/"
	}

	if len(uri) == 0 || uri[0] != '/' {
		s.Logger().Fatalf(ctx, `invalid pattern "%s", URI should lead with '/'`, pattern)
	}

		// 重复的路由检查，这个功能可以通过服务器配置来禁用。 md5:16d9ca5ef5f6ce27
	var routerKey = s.routerMapKey(handler.HookName, method, uri, domain)
	if !s.config.RouteOverWrite {
		switch handler.Type {
		case HandlerTypeHandler, HandlerTypeObject:
			if items, ok := s.routesMap[routerKey]; ok {
				var duplicatedHandler *HandlerItem
				for i, item := range items {
					switch item.Type {
					case HandlerTypeHandler, HandlerTypeObject:
						duplicatedHandler = items[i]
					}
					if duplicatedHandler != nil {
						break
					}
				}
				if duplicatedHandler != nil {
					s.Logger().Fatalf(
						ctx,
						`duplicated route registry "%s" at %s , already registered at %s`,
						pattern, handler.Source, duplicatedHandler.Source,
					)
				}
			}
		}
	}
		// 每个处理器的唯一标识符。 md5:d5cdd6ccf90c625e
	handler.Id = handlerIdGenerator.Add(1)
		// 根据给定的参数创建一个新的路由。 md5:a6e213f025b1718b
	handler.Router = &Router{
		Uri:      uri,
		Domain:   domain,
		Method:   strings.ToUpper(method),
		Priority: strings.Count(uri[1:], "/"),
	}
	handler.Router.RegRule, handler.Router.RegNames = s.patternToRegular(uri)

	if _, ok := s.serveTree[domain]; !ok {
		s.serveTree[domain] = make(map[string]interface{})
	}
	// List数组，对路由器注册非常重要。
	// 在从根到叶的搜索过程中，可能会有多个列表添加到这个数组中。
	// md5:7ddaff62bcec3109
	var (
		array []string
		lists = make([]*glist.List, 0)
	)
	if strings.EqualFold("/", uri) {
		array = []string{"/"}
	} else {
		array = strings.Split(uri[1:], "/")
	}
	// 多层哈希表：
	// 1. 表中的每个节点由以字符 '/' 分割的 URI 路径标识。
	// 2. 键 "*fuzz" 指示这是一个模糊节点，它没有确定的名字。
	// 3. 键 "*list" 是节点的列表项，大多数节点都有这个项，特别是模糊节点。注意：模糊节点必须有 "*list" 项，叶子节点也有 "*list" 项。如果节点既不是模糊节点也不是叶子节点，则不包含 "*list" 项。
	// 4. "*list" 项是一个按优先级从高到低排序的已注册路由项的列表。如果是模糊节点，该模糊节点的所有子路由项也会添加到其 "*list" 项中。
	// 5. 路由列表中可能存在重复的路由项。从根到叶的列表优先级是从低到高。
	// md5:3b9d86c224bf6153
	var p = s.serveTree[domain]
	for i, part := range array {
						//index. md5:44ed3114aa11886a
		if part == "" {
			continue
		}
				// 检查是否为模糊节点。 md5:ea4491ebe7a6c626
		if gregex.IsMatchString(`^[:\*]|\{[\w\.\-]+\}|\*`, part) {
			part = "*fuzz"
			// 如果它是一个模糊节点，它会在哈希映射中创建一个"*list"项，这实际上是一个列表。
			// 该模糊节点下的所有子路由器项也将被添加到它的"*list"项中。
			// md5:31e4feee2e295113
			if v, ok := p.(map[string]interface{})["*list"]; !ok {
				newListForFuzzy := glist.New()
				p.(map[string]interface{})["*list"] = newListForFuzzy
				lists = append(lists, newListForFuzzy)
			} else {
				lists = append(lists, v.(*glist.List))
			}
		}
				// 为当前节点创建一个新的桶。 md5:a99064964800f461
		if _, ok := p.(map[string]interface{})[part]; !ok {
			p.(map[string]interface{})[part] = make(map[string]interface{})
		}
		// Loop to next bucket.
		p = p.(map[string]interface{})[part]
		// 叶节点是一个哈希映射，必须包含一个名为"*list"的项，其中包含路由项。
		// 叶节点可以通过在其映射中添加更多的键值对来进一步扩展。
		// 请注意，需要进行 `v != "*fuzz"` 的比较，因为列表可能在先前的模糊检查中被添加。
		// md5:0a1026e07b9b2544
		if i == len(array)-1 && part != "*fuzz" {
			if v, ok := p.(map[string]interface{})["*list"]; !ok {
				leafList := glist.New()
				p.(map[string]interface{})["*list"] = leafList
				lists = append(lists, leafList)
			} else {
				lists = append(lists, v.(*glist.List))
			}
		}
	}
	// 它遍历`lists`的列表数组，比较优先级，并将新的路由项插入到每个列表的适当位置。
	// 列表的优先级从高到低排序。
	// md5:f7e3738ec2e01b79
	var item *HandlerItem
	for _, l := range lists {
		pushed := false
		for e := l.Front(); e != nil; e = e.Next() {
			item = e.Value.(*HandlerItem)
			// 检查是否应在当前项之前插入路由项，即它具有更高的优先级。
			// md5:0e6fc2994f00bc96
			if s.compareRouterPriority(handler, item) {
				l.InsertBefore(e, handler)
				pushed = true
				goto end
			}
		}
	end:
				// 就默认再推回去。 md5:eba32b1a9fbdfd1f
		if !pushed {
			l.PushBack(handler)
		}
	}
		// 初始化路由项。 md5:6ae2dff0c163c17e
	if _, ok := s.routesMap[routerKey]; !ok {
		s.routesMap[routerKey] = make([]*HandlerItem, 0)
	}

	// Append the route.
	s.routesMap[routerKey] = append(s.routesMap[routerKey], handler)
}

func (s *Server) isValidMethod(method string) bool {
	if gstr.Equal(method, defaultMethod) {
		return true
	}
	_, ok := methodsMap[strings.ToUpper(method)]
	return ok
}

// compareRouterPriority 比较 `newItem` 和 `oldItem` 之间的优先级。如果 `newItem` 的优先级高于 `oldItem`，则返回 true；否则返回 false。优先级较高的项将被插入到路由列表的前面。
//
// 比较规则：
// 1. 中间件具有最高的优先级。
// 2. URI：深度越深，优先级越高（简单地检查 URI 中字符 '/' 的数量）。
// 3. 路由类型：{xxx} > :xxx > *xxx。
// md5:d3f2e1aac7e71a05
func (s *Server) compareRouterPriority(newItem *HandlerItem, oldItem *HandlerItem) bool {
		// 如果它们都是中间件类型，则按照注册的顺序决定优先级。 md5:3a53e273b3f3566f
	if newItem.Type == HandlerTypeMiddleware && oldItem.Type == HandlerTypeMiddleware {
		return false
	}
		// 中间件具有最高的优先级。 md5:0ca14429f8f96e06
	if newItem.Type == HandlerTypeMiddleware && oldItem.Type != HandlerTypeMiddleware {
		return true
	}
		// URI：深度越深（只需检查URI中'/'字符的数量）。 md5:cf10a2d60b6808df
	if newItem.Router.Priority > oldItem.Router.Priority {
		return true
	}
	if newItem.Router.Priority < oldItem.Router.Priority {
		return false
	}

	// 比较它们的URI长度，
	// 但URI中的模糊部分和命名部分不计入结果中。
	// md5:55bd5729f8c0352a

	// 示例：
	// /admin-goods-{分页}  > /admin-{分页}
	// /{哈希}.{类型}      > /{哈希}
	// md5:482c38c410b3c591
	var uriNew, uriOld string
	uriNew, _ = gregex.ReplaceString(`\{[^/]+?\}`, "", newItem.Router.Uri)
	uriOld, _ = gregex.ReplaceString(`\{[^/]+?\}`, "", oldItem.Router.Uri)
	uriNew, _ = gregex.ReplaceString(`:[^/]+?`, "", uriNew)
	uriOld, _ = gregex.ReplaceString(`:[^/]+?`, "", uriOld)
	uriNew, _ = gregex.ReplaceString(`\*[^/]*`, "", uriNew) // 替换 "/*" 和 "任何字符串"。 md5:4bbaf5031e185545
	uriOld, _ = gregex.ReplaceString(`\*[^/]*`, "", uriOld) // 替换 "/*" 和 "任何字符串"。 md5:4bbaf5031e185545
	if len(uriNew) > len(uriOld) {
		return true
	}
	if len(uriNew) < len(uriOld) {
		return false
	}

	// 路由类型检查：{xxx} > :xxx > *xxx。
	// 例子：
	// /name/act > /{name}/:act
	// 
	// 这段注释的意思是，它在描述Go语言中的路由类型检查规则。`{xxx}`、`:xxx`和`*xxx`是路由匹配模式：
	// 
	// - `{xxx}` 表示路径中可以包含任意字符的占位符，但需要与实际请求中的某个参数匹配。
	// - `:xxx` 表示路径中的命名参数，这些参数将在路由处理函数中作为变量传递。
	// - `*xxx` 表示零个或多个重复的前面的模式，通常用于处理路径中的可选组件。
	// 
	// 举例来说，路由`/name/act` 使用了`{name}`和`:act`，表示请求的URL可以形式为`/具体名称/操作名`，`{name}`会被替换为实际请求中的名称，`:act`则是一个动态的操作标识。
	// md5:5fc64b4a4a78b2aa
	var (
		fuzzyCountFieldNew int
		fuzzyCountFieldOld int
		fuzzyCountNameNew  int
		fuzzyCountNameOld  int
		fuzzyCountAnyNew   int
		fuzzyCountAnyOld   int
		fuzzyCountTotalNew int
		fuzzyCountTotalOld int
	)
	for _, v := range newItem.Router.Uri {
		switch v {
		case '{':
			fuzzyCountFieldNew++
		case ':':
			fuzzyCountNameNew++
		case '*':
			fuzzyCountAnyNew++
		}
	}
	for _, v := range oldItem.Router.Uri {
		switch v {
		case '{':
			fuzzyCountFieldOld++
		case ':':
			fuzzyCountNameOld++
		case '*':
			fuzzyCountAnyOld++
		}
	}
	fuzzyCountTotalNew = fuzzyCountFieldNew + fuzzyCountNameNew + fuzzyCountAnyNew
	fuzzyCountTotalOld = fuzzyCountFieldOld + fuzzyCountNameOld + fuzzyCountAnyOld
	if fuzzyCountTotalNew < fuzzyCountTotalOld {
		return true
	}
	if fuzzyCountTotalNew > fuzzyCountTotalOld {
		return false
	}

		// 如果它们的模糊规则数量相等。 md5:0a1cd4da270f5da3

		// 例如：/name/{act} > /name/:act. md5:14051818a0cea80c
	if fuzzyCountFieldNew > fuzzyCountFieldOld {
		return true
	}
	if fuzzyCountFieldNew < fuzzyCountFieldOld {
		return false
	}
		// 这段注释的意思是，当URL路径匹配模式"/name/:act"时，它会被转换或重写为"/name/*act"。这里的":act"是一个占位符，表示可以包含任意字符的动态部分，"*act"则表示任何字符序列（包括零个字符）。这是一种路由或路径匹配的规则，在Go或其他支持类似语法的语言中常见于处理URL路由。 md5:6f9e027f06c4b833
	if fuzzyCountNameNew > fuzzyCountNameOld {
		return true
	}
	if fuzzyCountNameNew < fuzzyCountNameOld {
		return false
	}

	// 然后，它会比较它们的HTTP方法的准确性，越准确优先级越高。
	// md5:19e263d51107b5cb
	if newItem.Router.Method != defaultMethod {
		return true
	}
	if oldItem.Router.Method != defaultMethod {
		return true
	}

	// 如果它们具有不同的路由类型，
	// 那么新的路由项比其他项具有更高的优先级。
	// md5:63dfba3b91db8cc4
	if newItem.Type == HandlerTypeHandler || newItem.Type == HandlerTypeObject {
		return true
	}

	// 其他情况，如HOOK项目，
	// 旧的路由项具有更高的优先级。
	// md5:53b3ce09282d12db
	return false
}

// patternToRegular 将路由规则转换为相应的正则表达式。 md5:c212402d9fd8cb59
func (s *Server) patternToRegular(rule string) (regular string, names []string) {
	if len(rule) < 2 {
		return rule, nil
	}
	regular = "^"
	var array = strings.Split(rule[1:], "/")
	for _, v := range array {
		if len(v) == 0 {
			continue
		}
		switch v[0] {
		case ':':
			if len(v) > 1 {
				regular += `/([^/]+)`
				names = append(names, v[1:])
			} else {
				regular += `/[^/]+`
			}
		case '*':
			if len(v) > 1 {
				regular += `/{0,1}(.*)`
				names = append(names, v[1:])
			} else {
				regular += `/{0,1}.*`
			}
		default:
						// 特殊字符替换。 md5:fe1b718da00180dd
			v = gstr.ReplaceByMap(v, map[string]string{
				`.`: `\.`,
				`+`: `\+`,
				`*`: `.*`,
			})
			s, _ := gregex.ReplaceStringFunc(`\{[\w\.\-]+\}`, v, func(s string) string {
				names = append(names, s[1:len(s)-1])
				return `([^/]+)`
			})
			if strings.EqualFold(s, v) {
				regular += "/" + v
			} else {
				regular += "/" + s
			}
		}
	}
	regular += `$`
	return
}
