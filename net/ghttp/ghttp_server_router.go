// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"context"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	
	"github.com/888go/goframe/container/glist"
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/debug/gdebug"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/consts"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gmeta"
	"github.com/888go/goframe/util/gtag"
)

var (
	// handlerIdGenerator 是处理器项 ID 生成器。
	handlerIdGenerator = 安全变量类.NewInt()
)

// routerMapKey 根据给定的参数创建并返回一个唯一的路由键。
// 此键用于 Server.routerMap 属性，主要用于检查重复的路由注册情况。
func (s *Server) routerMapKey(hook HookName, method, path, domain string) string {
	return string(hook) + "%" + s.serveHandlerKey(method, path, domain)
}

// parsePattern 将给定的模式解析为域名、方法和路径变量。
func (s *Server) parsePattern(pattern string) (domain, method, path string, err error) {
	path = strings.TrimSpace(pattern)
	domain = DefaultDomainName
	method = defaultMethod
	if array, err := 正则类.X匹配文本(`([a-zA-Z]+):(.+)`, pattern); len(array) > 1 && err == nil {
		path = strings.TrimSpace(array[2])
		if v := strings.TrimSpace(array[1]); v != "" {
			method = v
		}
	}
	if array, err := 正则类.X匹配文本(`(.+)@([\w\.\-]+)`, path); len(array) > 1 && err == nil {
		path = strings.TrimSpace(array[1])
		if v := strings.TrimSpace(array[2]); v != "" {
			domain = v
		}
	}
	if path == "" {
		err = 错误类.X创建错误码(错误码类.CodeInvalidParameter, "invalid pattern: URI should not be empty")
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
// 路由树可以被视为一个多层哈希表，请参考下文中的注释。
// 此函数在服务器启动时调用，对性能要求不高。真正重要的是
// 当请求处于服务状态时，用于路由搜索的良好设计的路由存储结构。
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
		s.Logger别名().X输出并格式化FATA(ctx, `invalid pattern "%s", %+v`, pattern, err)
		return
	}
// ====================================================================================
// 根据请求结构中的元信息更改已注册的路由。
// 它支持使用逗号 `,` 连接的多种方法。
// ====================================================================================
	if handler.Info.Type != nil && handler.Info.Type.NumIn() == 2 {
		var objectReq = reflect.New(handler.Info.Type.In(1))
		if v := 元数据类.Get(objectReq, gtag.Path); !v.X是否为空() {
			uri = v.String()
		}
		if v := 元数据类.Get(objectReq, gtag.Domain); !v.X是否为空() {
			domain = v.String()
		}
		if v := 元数据类.Get(objectReq, gtag.Method); !v.X是否为空() {
			method = v.String()
		}
		// 多个方法注册，使用字符 `,` 连接。
		if 文本类.X是否包含(method, ",") {
			methods := 文本类.X分割并忽略空值(method, ",")
			for _, v := range methods {
				// 每个方法都有自己的处理程序。
				clonedHandler := *handler
				s.doSetHandler(ctx, &clonedHandler, prefix, uri, pattern, v, domain)
			}
			return
		}
		// 将`all`转换为`ALL`。
		if 文本类.X相等比较并忽略大小写(method, defaultMethod) {
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
		s.Logger别名().X输出并格式化FATA(
			ctx,
			`invalid method value "%s", should be in "%s" or "%s"`,
			method, supportedHttpMethods, defaultMethod,
		)
	}
	// URI特征的前缀。
	if prefix != "" {
		uri = prefix + "/" + strings.TrimLeft(uri, "/")
	}
	uri = strings.TrimRight(uri, "/")
	if uri == "" {
		uri = "/"
	}

	if len(uri) == 0 || uri[0] != '/' {
		s.Logger别名().X输出并格式化FATA(ctx, `invalid pattern "%s", URI should lead with '/'`, pattern)
	}

	// 重复路由检查，可以通过服务器配置禁用此功能。
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
					s.Logger别名().X输出并格式化FATA(
						ctx,
						`duplicated route registry "%s" at %s , already registered at %s`,
						pattern, handler.Source, duplicatedHandler.Source,
					)
				}
			}
		}
	}
	// 每个处理器的唯一标识 ID
	handler.Id = handlerIdGenerator.Add(1)
	// 根据给定参数创建一个新的路由器。
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
// List 数组，对于路由器注册非常重要。
// 从根节点到叶子节点搜索过程中，可能会有多个列表添加到此数组中。
	var (
		array []string
		lists = make([]*链表类.List, 0)
	)
	if strings.EqualFold("/", uri) {
		array = []string{"/"}
	} else {
		array = strings.Split(uri[1:], "/")
	}
// 多层哈希表:
// 1. 表中的每个节点通过以字符 '/' 分割的URI路径进行区分。
// 2. 键 "*fuzz" 指定该节点是一个模糊节点，没有特定名称。
// 3. 键 "*list" 是节点的列表项，**大多数节点都有此项**，特别是模糊节点。注意，模糊节点必须包含"*list"项，并且叶节点也具有"*list"项。如果节点既不是模糊节点也不是叶节点，则它不包含"*list"项。
// 4. "*list" 项是一个按其优先级从高到低排序的已注册路由项列表。如果是模糊节点，从此模糊节点开始的所有子路由项也将被添加到其"*list"项中。
// 5. 路由列表中可能存在重复的路由项。从根节点到叶节点的列表优先级是从低到高。
	var p = s.serveTree[domain]
	for i, part := range array {
		// 忽略空的URI部分，例如：/user//index
		if part == "" {
			continue
		}
		// 检查是否为模糊节点。
		if 正则类.X是否匹配文本(`^[:\*]|\{[\w\.\-]+\}|\*`, part) {
			part = "*fuzz"
// 如果这是一个模糊节点，它会在哈希映射中创建一个“*list”项——这是一个列表。
// 从此模糊节点派生的所有子路由项也将被添加到其“*list”项中。
			if v, ok := p.(map[string]interface{})["*list"]; !ok {
				newListForFuzzy := 链表类.New()
				p.(map[string]interface{})["*list"] = newListForFuzzy
				lists = append(lists, newListForFuzzy)
			} else {
				lists = append(lists, v.(*链表类.List))
			}
		}
		// 为当前节点新建一个桶。
		if _, ok := p.(map[string]interface{})[part]; !ok {
			p.(map[string]interface{})[part] = make(map[string]interface{})
		}
		// 循环到下一个桶。
		p = p.(map[string]interface{})[part]
// 叶子节点是一个哈希表，且必须包含一个名为"*list"的项目，其中存储着路由项。
// 通过向其映射中添加更多键值对，叶子节点可以进一步扩展。
// 注意，由于在之前的模糊检查中可能已添加了列表，所以需要进行 `v != "*fuzz"` 的比较。
		if i == len(array)-1 && part != "*fuzz" {
			if v, ok := p.(map[string]interface{})["*list"]; !ok {
				leafList := 链表类.New()
				p.(map[string]interface{})["*list"] = leafList
				lists = append(lists, leafList)
			} else {
				lists = append(lists, v.(*链表类.List))
			}
		}
	}
// 它遍历`lists`列表数组，比较优先级并将新的路由项插入到每个列表中的适当位置。列表的优先级从高到低排序。
	var item *HandlerItem
	for _, l := range lists {
		pushed := false
		for e := l.Front(); e != nil; e = e.Next() {
			item = e.Value.(*HandlerItem)
// 检查优先级，是否在当前项之前插入路由项，
// 这意味着它具有更高的优先级。
			if s.compareRouterPriority(handler, item) {
				l.InsertBefore(e, handler)
				pushed = true
				goto end
			}
		}
	end:
		// 默认情况下，仅向后推入
		if !pushed {
			l.PushBack(handler)
		}
	}
	// 初始化路由映射项。
	if _, ok := s.routesMap[routerKey]; !ok {
		s.routesMap[routerKey] = make([]*HandlerItem, 0)
	}

	// Append the route.
	s.routesMap[routerKey] = append(s.routesMap[routerKey], handler)
}

func (s *Server) isValidMethod(method string) bool {
	if 文本类.X相等比较并忽略大小写(method, defaultMethod) {
		return true
	}
	_, ok := methodsMap[strings.ToUpper(method)]
	return ok
}

// compareRouterPriority 比较 `newItem` 和 `oldItem` 之间的优先级。如果 `newItem` 的优先级高于 `oldItem`，则返回 true，否则返回 false。优先级较高的项将被插入到路由器列表的前面。
//
// 比较规则：
// 1. 中间件具有最高的优先级。
// 2. URI：路径越深，优先级越高（简单地检查 URI 中字符 '/' 的数量）。
// 3. 路由类型：{xxx} > :xxx > *xxx。
func (s *Server) compareRouterPriority(newItem *HandlerItem, oldItem *HandlerItem) bool {
	// 如果它们都是中间件类型，则优先级根据其注册顺序决定。
	if newItem.Type == HandlerTypeMiddleware && oldItem.Type == HandlerTypeMiddleware {
		return false
	}
	// 该中间件具有最高优先级。
	if newItem.Type == HandlerTypeMiddleware && oldItem.Type != HandlerTypeMiddleware {
		return true
	}
	// URI：URI中'/'字符出现次数越多，级别越高。
	if newItem.Router.Priority > oldItem.Router.Priority {
		return true
	}
	if newItem.Router.Priority < oldItem.Router.Priority {
		return false
	}

// 比较它们URI的长度，
// 但URI中的模糊部分和命名部分不计算到结果中。

// 示例：
// /admin-goods-{page} > /admin-{page}
// /{hash}.{type}      > /{hash}
// 上面的注释是用于示例URL路径重写规则：
// 第一条规则表示将 "/admin-goods-任意页码" 重写为 "/admin-任意页码"，其中 {page} 是一个占位符，代表任何数字页码。
// 第二条规则表示将 "/任意哈希值.任意类型" 重写为 "/任意哈希值"，其中 {hash} 和 {type} 分别是占位符，代表任何哈希值和文件类型。
	var uriNew, uriOld string
	uriNew, _ = 正则类.X替换文本(`\{[^/]+?\}`, "", newItem.Router.Uri)
	uriOld, _ = 正则类.X替换文本(`\{[^/]+?\}`, "", oldItem.Router.Uri)
	uriNew, _ = 正则类.X替换文本(`:[^/]+?`, "", uriNew)
	uriOld, _ = 正则类.X替换文本(`:[^/]+?`, "", uriOld)
	uriNew, _ = 正则类.X替换文本(`\*[^/]*`, "", uriNew) // 将 "/*" 和 "/*any" 进行替换。
	uriOld, _ = 正则类.X替换文本(`\*[^/]*`, "", uriOld) // 将 "/*" 和 "/*any" 进行替换。
	if len(uriNew) > len(uriOld) {
		return true
	}
	if len(uriNew) < len(uriOld) {
		return false
	}

// 路由类型检查规则：{xxx} > :xxx > *xxx。
// 示例：
// /name/act 对应于 /{name}/:act
// （注释翻译：这段Go语言代码的注释描述了路由路径匹配的优先级规则，其中花括号 `{}`、冒号 `:` 以及星号 `*` 分别用于表示路径参数的不同格式。按照优先级从高到低排列为：`{xxx}`（命名路径参数）、`:xxx`（动态路径参数）和 `*xxx`（任意长度路径参数）。示例说明了这种对应关系，在路由 `/name/act` 中，`name` 可以映射为 `{name}` 形式的命名路径参数，而 `act` 可以映射为 `:act` 形式的动态路径参数。）
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

	// 如果它们的模糊规则计数相等。

	// 示例：/name/{act} 转换为 /name/:act
	if fuzzyCountFieldNew > fuzzyCountFieldOld {
		return true
	}
	if fuzzyCountFieldNew < fuzzyCountFieldOld {
		return false
	}
	// 示例：/name/:act > /name/*act
// 注释翻译：该注释用于表示一个路由映射规则的示例，其中"/name/:act"是一个动态路由模板，":act"是一个参数占位符，可以匹配任何非空字符串。在实际应用中，它将被映射到类似"/name/任意动作名"的实际路由路径。例如，如果":act"为"edit"，则此规则可匹配路径"/name/edit"。而"*act"则通常用于匹配任意后缀，包括"/"字符及其后面的所有内容。但在给定的代码片段中，并没有明确指出"*act"的行为，可能需要更多上下文信息来准确解释。
	if fuzzyCountNameNew > fuzzyCountNameOld {
		return true
	}
	if fuzzyCountNameNew < fuzzyCountNameOld {
		return false
	}

// 然后比较它们的HTTP方法的准确性，
// 越准确则优先级越高。
	if newItem.Router.Method != defaultMethod {
		return true
	}
	if oldItem.Router.Method != defaultMethod {
		return true
	}

// 如果它们具有不同的路由类型，
// 那么新的路由项比另一个具有更高的优先级。
	if newItem.Type == HandlerTypeHandler || newItem.Type == HandlerTypeObject {
		return true
	}

// 其他情况，如 HOOK 项，
// 则旧的路由项比其他项具有更高的优先级。
	return false
}

// patternToRegular 将路由规则转换为相应的正则表达式。
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
			// 特殊字符替换。
			v = 文本类.Map替换(v, map[string]string{
				`.`: `\.`,
				`+`: `\+`,
				`*`: `.*`,
			})
			s, _ := 正则类.X替换文本_函数(`\{[\w\.\-]+\}`, v, func(s string) string {
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
