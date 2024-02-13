// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"context"
	"fmt"
	"reflect"
	
	"github.com/888go/goframe/debug/gdebug"
	"github.com/888go/goframe/internal/consts"
	"github.com/888go/goframe/internal/reflection"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
)

type (
	// RouterGroup 是一个包裹了多个路由和中间件的分组。
	RouterGroup struct {
		parent     *RouterGroup  // Parent group.
		server     *Server       // Server.
		domain     *Domain       // Domain.
		prefix     string        // 子路由前缀。
		middleware []HandlerFunc // Middleware array.
	}

// preBindItem 是 router group 中用于懒加载注册功能的项目。当调用该组的路由函数时，preBindItem 并未真正注册到服务器上，
// 而是在服务器启动时才进行惰性注册。
	preBindItem struct {
		group    *RouterGroup
		bindType string
		pattern  string
		object   interface{}   // 可以是处理器、控制器或对象。
		params   []interface{} // 根据类型为路由注册提供额外的参数。
		source   string        // Handler 是在某个源文件路径：行处注册的处理器
		bound    bool          // 这个条目是否绑定到服务器？
	}
)

const (
	groupBindTypeHandler    = "HANDLER"
	groupBindTypeRest       = "REST"
	groupBindTypeHook       = "HOOK"
	groupBindTypeMiddleware = "MIDDLEWARE"
)

var (
	preBindItems = make([]*preBindItem, 0, 64)
)

// handlePreBindItems 在服务器启动时被调用，它负责向服务器进行实际的路由注册。
func (s *Server) handlePreBindItems(ctx context.Context) {
	if len(preBindItems) == 0 {
		return
	}
	for _, item := range preBindItems {
		if item.bound {
			continue
		}
		// 处理当前服务器的项目。
		if item.group.server != nil && item.group.server != s {
			continue
		}
		if item.group.domain != nil && item.group.domain.server != s {
			continue
		}
		item.group.doBindRoutersToServer(ctx, item)
		item.bound = true
	}
}

// Group 创建并返回一个 RouterGroup 对象。
func (s *Server) X创建分组路由(分组前缀 string, 分组函数 ...func(分组路由 *RouterGroup)) *RouterGroup {
	if len(分组前缀) > 0 && 分组前缀[0] != '/' {
		分组前缀 = "/" + 分组前缀
	}
	if 分组前缀 == "/" {
		分组前缀 = ""
	}
	group := &RouterGroup{
		server: s,
		prefix: 分组前缀,
	}
	if len(分组函数) > 0 {
		for _, v := range 分组函数 {
			v(group)
		}
	}
	return group
}

// Group 创建并返回一个 RouterGroup 对象，该对象与指定的域名绑定。
func (d *Domain) X创建分组路由(分组前缀 string, 分组函数 ...func(分组路由 *RouterGroup)) *RouterGroup {
	if len(分组前缀) > 0 && 分组前缀[0] != '/' {
		分组前缀 = "/" + 分组前缀
	}
	if 分组前缀 == "/" {
		分组前缀 = ""
	}
	routerGroup := &RouterGroup{
		domain: d,
		server: d.server,
		prefix: 分组前缀,
	}
	if len(分组函数) > 0 {
		for _, nestedGroup := range 分组函数 {
			nestedGroup(routerGroup)
		}
	}
	return routerGroup
}

// Group 创建并返回当前路由组的一个子组。
func (g *RouterGroup) X创建分组路由(分组前缀 string, 分组函数 ...func(分组路由 *RouterGroup)) *RouterGroup {
	if 分组前缀 == "/" {
		分组前缀 = ""
	}
	group := &RouterGroup{
		parent: g,
		server: g.server,
		domain: g.domain,
		prefix: 分组前缀,
	}
	if len(g.middleware) > 0 {
		group.middleware = make([]HandlerFunc, len(g.middleware))
		copy(group.middleware, g.middleware)
	}
	if len(分组函数) > 0 {
		for _, v := range 分组函数 {
			v(group)
		}
	}
	return group
}

// Clone 返回一个新的路由组，它是当前组的克隆副本。
func (g *RouterGroup) X取副本() *RouterGroup {
	newGroup := &RouterGroup{
		parent:     g.parent,
		server:     g.server,
		domain:     g.domain,
		prefix:     g.prefix,
		middleware: make([]HandlerFunc, len(g.middleware)),
	}
	copy(newGroup.middleware, g.middleware)
	return newGroup
}

// Bind 为一个路由组提供批量注册路由的功能。
func (g *RouterGroup) X绑定(处理对象 ...interface{}) *RouterGroup {
	var (
		ctx   = context.TODO()
		group = g.X取副本()
	)
	for _, v := range 处理对象 {
		var (
			item               = v
			originValueAndKind = reflection.OriginValueAndKind(item)
		)

		switch originValueAndKind.OriginKind {
		case reflect.Func, reflect.Struct:
			group = group.preBindToLocalArray(
				groupBindTypeHandler,
				"/",
				item,
			)

		default:
			g.server.Logger别名().X输出并格式化FATA(
				ctx, "invalid bind parameter type: %v, should be route function or struct object",
				originValueAndKind.InputValue.Type(),
			)
		}
	}
	return group
}

// ALL 注册一个HTTP处理器，通过给定路由模式和所有HTTP方法。
func (g *RouterGroup) X绑定所有类型(路由规则 string, 处理函数 interface{}, 额外参数 ...interface{}) *RouterGroup {
	return g.X取副本().preBindToLocalArray(
		groupBindTypeHandler,
		defaultMethod+":"+路由规则,
		处理函数,
		额外参数...,
	)
}

// ALLMap 使用map为HTTP方法注册HTTP处理程序。
func (g *RouterGroup) X绑定所有类型Map(m map[string]interface{}) {
	for pattern, object := range m {
		g.X绑定所有类型(pattern, object)
	}
}

// Map 通过使用映射表注册HTTP方法对应的HTTP处理器。
func (g *RouterGroup) X绑定Map(m map[string]interface{}) {
	for pattern, object := range m {
		g.preBindToLocalArray(groupBindTypeHandler, pattern, object)
	}
}

// GET 注册一个 HTTP 处理器，用于给定的路由模式和 HTTP 方法：GET。
func (g *RouterGroup) X绑定GET(路由规则 string, 处理函数 interface{}, 额外参数 ...interface{}) *RouterGroup {
	return g.X取副本().preBindToLocalArray(groupBindTypeHandler, "GET:"+路由规则, 处理函数, 额外参数...)
}

// PUT 注册一个 HTTP 处理器，用于给定路由模式和 HTTP 方法：PUT。
func (g *RouterGroup) X绑定PUT(路由规则 string, 处理函数 interface{}, 额外参数 ...interface{}) *RouterGroup {
	return g.X取副本().preBindToLocalArray(groupBindTypeHandler, "PUT:"+路由规则, 处理函数, 额外参数...)
}

// POST 注册一个 HTTP 处理器，用于给定路由模式和 HTTP 方法：POST。
func (g *RouterGroup) X绑定POST(路由规则 string, 处理函数 interface{}, 额外参数 ...interface{}) *RouterGroup {
	return g.X取副本().preBindToLocalArray(groupBindTypeHandler, "POST:"+路由规则, 处理函数, 额外参数...)
}

// DELETE 注册一个 HTTP 处理器，根据给定的路由模式和 HTTP 方法（DELETE）进行处理。
func (g *RouterGroup) X绑定DELETE(路由规则 string, 处理函数 interface{}, 额外参数 ...interface{}) *RouterGroup {
	return g.X取副本().preBindToLocalArray(groupBindTypeHandler, "DELETE:"+路由规则, 处理函数, 额外参数...)
}

// PATCH 注册一个 HTTP 处理器，用于给定的路由模式和 HTTP 方法：PATCH。
func (g *RouterGroup) X绑定PATCH(路由规则 string, 处理函数 interface{}, 额外参数 ...interface{}) *RouterGroup {
	return g.X取副本().preBindToLocalArray(groupBindTypeHandler, "PATCH:"+路由规则, 处理函数, 额外参数...)
}

// HEAD 注册一个 HTTP 处理器，用于指定路由模式和 HTTP 方法：HEAD。
func (g *RouterGroup) X绑定HEAD(路由规则 string, 处理函数 interface{}, 额外参数 ...interface{}) *RouterGroup {
	return g.X取副本().preBindToLocalArray(groupBindTypeHandler, "HEAD:"+路由规则, 处理函数, 额外参数...)
}

// CONNECT 注册一个 HTTP 处理器，用于给定路由模式和 HTTP 方法：CONNECT。
func (g *RouterGroup) X绑定CONNECT(路由规则 string, 处理函数 interface{}, 额外参数 ...interface{}) *RouterGroup {
	return g.X取副本().preBindToLocalArray(groupBindTypeHandler, "CONNECT:"+路由规则, 处理函数, 额外参数...)
}

// OPTIONS 注册一个HTTP处理器，用于指定路由模式和HTTP方法：OPTIONS。
func (g *RouterGroup) X绑定OPTIONS(路由规则 string, 处理函数 interface{}, 额外参数 ...interface{}) *RouterGroup {
	return g.X取副本().preBindToLocalArray(groupBindTypeHandler, "OPTIONS:"+路由规则, 处理函数, 额外参数...)
}

// TRACE 注册一个 HTTP 处理器，用于提供路由模式和 HTTP 方法：TRACE。
func (g *RouterGroup) X绑定TRACE(路由规则 string, 处理函数 interface{}, 额外参数 ...interface{}) *RouterGroup {
	return g.X取副本().preBindToLocalArray(groupBindTypeHandler, "TRACE:"+路由规则, 处理函数, 额外参数...)
}

// REST 根据REST规则注册一个HTTP处理器，以便给定路由模式。
func (g *RouterGroup) X绑定RESTfulAPI对象(路由规则 string, 处理对象 interface{}) *RouterGroup {
	return g.X取副本().preBindToLocalArray(groupBindTypeRest, 路由规则, 处理对象)
}

// Hook 将钩子注册到给定的路由模式。
func (g *RouterGroup) X绑定Hook(路由规则 string, 触发时机 HookName, 处理函数 HandlerFunc) *RouterGroup {
	return g.X取副本().preBindToLocalArray(groupBindTypeHandler, 路由规则, 处理函数, 触发时机)
}

// Middleware 将一个或多个中间件绑定到路由组。
func (g *RouterGroup) X绑定中间件(处理函数 ...HandlerFunc) *RouterGroup {
	g.middleware = append(g.middleware, 处理函数...)
	return g
}

// preBindToLocalArray 将路由注册参数预先添加到内部变量数组中，以便于进行惰性注册功能。
func (g *RouterGroup) preBindToLocalArray(bindType string, pattern string, object interface{}, params ...interface{}) *RouterGroup {
	_, file, line := gdebug.CallerWithFilter([]string{consts.StackFilterKeyForGoFrame})
	preBindItems = append(preBindItems, &preBindItem{
		group:    g,
		bindType: bindType,
		pattern:  pattern,
		object:   object,
		params:   params,
		source:   fmt.Sprintf(`%s:%d`, file, line),
	})
	return g
}

// getPrefix 返回当前组的路由前缀，该方法会递归获取其父级的前缀。
func (g *RouterGroup) getPrefix() string {
	prefix := g.prefix
	parent := g.parent
	for parent != nil {
		prefix = parent.prefix + prefix
		parent = parent.parent
	}
	return prefix
}

// doBindRoutersToServer 是真正执行将路由器绑定到服务器的注册操作。
func (g *RouterGroup) doBindRoutersToServer(ctx context.Context, item *preBindItem) *RouterGroup {
	var (
		bindType = item.bindType
		pattern  = item.pattern
		object   = item.object
		params   = item.params
		source   = item.source
	)
	prefix := g.getPrefix()
	// Route check.
	if len(prefix) > 0 {
		domain, method, path, err := g.server.parsePattern(pattern)
		if err != nil {
			g.server.Logger别名().X输出并格式化FATA(ctx, "invalid route pattern: %s", pattern)
		}
		// 如果已经有域名，那么在模式中取消设置域名字段。
		if g.domain != nil {
			domain = ""
		}
		if bindType == groupBindTypeRest {
			pattern = path
		} else {
			pattern = g.server.serveHandlerKey(
				method, path, domain,
			)
		}
	}
	// 过滤重复的字符 '/'。
	pattern = 文本类.X替换(pattern, "//", "/")

	// 将参数转换为字符串数组。
	extras := 转换类.X取文本数组(params)

	// 检查是否为钩子处理器。
	if _, ok := object.(HandlerFunc); ok && len(extras) > 0 {
		bindType = groupBindTypeHook
	}
	switch bindType {
	case groupBindTypeHandler:
		if reflect.ValueOf(object).Kind() == reflect.Func {
			funcInfo, err := g.server.checkAndCreateFuncInfo(object, "", "", "")
			if err != nil {
				g.server.Logger别名().X输出FATA(ctx, err.Error())
				return g
			}
			in := doBindHandlerInput{
				Prefix:     prefix,
				Pattern:    pattern,
				FuncInfo:   funcInfo,
				Middleware: g.middleware,
				Source:     source,
			}
			if g.domain != nil {
				g.domain.doBindHandler(ctx, in)
			} else {
				g.server.doBindHandler(ctx, in)
			}
		} else {
			if len(extras) > 0 {
				if 文本类.X是否包含(extras[0], ",") {
					in := doBindObjectInput{
						Prefix:     prefix,
						Pattern:    pattern,
						Object:     object,
						Method:     extras[0],
						Middleware: g.middleware,
						Source:     source,
					}
					if g.domain != nil {
						g.domain.doBindObject(ctx, in)
					} else {
						g.server.doBindObject(ctx, in)
					}
				} else {
					in := doBindObjectMethodInput{
						Prefix:     prefix,
						Pattern:    pattern,
						Object:     object,
						Method:     extras[0],
						Middleware: g.middleware,
						Source:     source,
					}
					if g.domain != nil {
						g.domain.doBindObjectMethod(ctx, in)
					} else {
						g.server.doBindObjectMethod(ctx, in)
					}
				}
			} else {
				in := doBindObjectInput{
					Prefix:     prefix,
					Pattern:    pattern,
					Object:     object,
					Method:     "",
					Middleware: g.middleware,
					Source:     source,
				}
				// 最后，它将`object`视为正在注册的类型对象。
				if g.domain != nil {
					g.domain.doBindObject(ctx, in)
				} else {
					g.server.doBindObject(ctx, in)
				}
			}
		}

	case groupBindTypeRest:
		in := doBindObjectInput{
			Prefix:     prefix,
			Pattern:    pattern,
			Object:     object,
			Method:     "",
			Middleware: g.middleware,
			Source:     source,
		}
		if g.domain != nil {
			g.domain.doBindObjectRest(ctx, in)
		} else {
			g.server.doBindObjectRest(ctx, in)
		}

	case groupBindTypeHook:
		if handler, ok := object.(HandlerFunc); ok {
			in := doBindHookHandlerInput{
				Prefix:   prefix,
				Pattern:  pattern,
				HookName: HookName(extras[0]),
				Handler:  handler,
				Source:   source,
			}
			if g.domain != nil {
				g.domain.doBindHookHandler(ctx, in)
			} else {
				g.server.doBindHookHandler(ctx, in)
			}
		} else {
			g.server.Logger别名().X输出并格式化FATA(ctx, "invalid hook handler for pattern: %s", pattern)
		}
	}
	return g
}
