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
	"reflect"

	"github.com/888go/goframe/debug/gdebug"
	"github.com/888go/goframe/internal/consts"
	"github.com/888go/goframe/internal/reflection"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
)

type (
		// X分组路由 是一个包裹多个路由和中间件的分组。 md5:609e7eb75d8a51f0
	X分组路由 struct {
		parent     *X分组路由  // Parent group.
		server     *X服务       // Server.
		domain     *Domain       // Domain.
		prefix     string        // Prefix for sub-route.
		middleware []HandlerFunc // Middleware array.
	}

	// preBindItem 是用于路由器组延迟注册功能的项目。当路由组的路由函数被调用时，preBindItem 并没有真正注册到服务器，
	// 而是在服务器启动时进行懒惰注册。
	// md5:4255b2f4d61ba05c
	preBindItem struct {
		group    *X分组路由
		bindType string
		pattern  string
		object   interface{}   // 可以是处理器、控制器或对象。 md5:0c53d2880dc0aafc
		params   []interface{} // 根据类型，为路由注册额外的参数。 md5:9b2faaa25e40fcdb
		source   string        // Handler 是在特定源文件路径和行号处的注册处理器。 md5:0cf92074c14f8d58
		bound    bool          // 是否将此项目绑定到服务器？. md5:7da1001818a8ed44
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

// handlePreBindItems 在服务器启动时被调用，它确实将路由注册到服务器上。 md5:4f0e019b6d905274
func (s *X服务) handlePreBindItems(ctx context.Context) {
	if len(preBindItems) == 0 {
		return
	}
	for _, item := range preBindItems {
		if item.bound {
			continue
		}
				// 处理当前服务器的项目。 md5:5caf897b6c7b073a
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

// X创建分组路由 创建并返回一个 RouterGroup 对象。 md5:ab811975f9ba0334
func (s *X服务) X创建分组路由(分组前缀 string, 分组函数 ...func(分组路由 *X分组路由)) *X分组路由 {
	if len(分组前缀) > 0 && 分组前缀[0] != '/' {
		分组前缀 = "/" + 分组前缀
	}
	if 分组前缀 == "/" {
		分组前缀 = ""
	}
	group := &X分组路由{
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

// X创建分组路由 创建并返回一个 RouterGroup 对象，该对象绑定到指定的域名。 md5:bd60cfbd62234fcd
func (d *Domain) X创建分组路由(分组前缀 string, 分组函数 ...func(分组路由 *X分组路由)) *X分组路由 {
	if len(分组前缀) > 0 && 分组前缀[0] != '/' {
		分组前缀 = "/" + 分组前缀
	}
	if 分组前缀 == "/" {
		分组前缀 = ""
	}
	routerGroup := &X分组路由{
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

// X创建分组路由 创建并返回当前路由器组的一个子组。 md5:9706484677759d8f
func (g *X分组路由) X创建分组路由(分组前缀 string, 分组函数 ...func(分组路由 *X分组路由)) *X分组路由 {
	if 分组前缀 == "/" {
		分组前缀 = ""
	}
	group := &X分组路由{
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

// X取副本 返回一个新的路由组，它是当前组的克隆。 md5:a3328662d1da7f5f
func (g *X分组路由) X取副本() *X分组路由 {
	newGroup := &X分组路由{
		parent:     g.parent,
		server:     g.server,
		domain:     g.domain,
		prefix:     g.prefix,
		middleware: make([]HandlerFunc, len(g.middleware)),
	}
	copy(newGroup.middleware, g.middleware)
	return newGroup
}

// X绑定 为路由器组提供了批量路由注册的功能。 md5:16fbec330e17cafe
func (g *X分组路由) X绑定(处理对象 ...interface{}) *X分组路由 {
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

// X绑定所有类型 注册一个http处理器，用于处理给定路由模式的所有HTTP方法。 md5:06f3f9b3c30b17f0
func (g *X分组路由) X绑定所有类型(路由规则 string, 处理函数 interface{}, 额外参数 ...interface{}) *X分组路由 {
	return g.X取副本().preBindToLocalArray(
		groupBindTypeHandler,
		defaultMethod+":"+路由规则,
		处理函数,
		额外参数...,
	)
}

// X绑定所有类型Map 使用映射注册HTTP处理程序，针对HTTP方法。 md5:4baef0383348c469
func (g *X分组路由) X绑定所有类型Map(m map[string]interface{}) {
	for pattern, object := range m {
		g.X绑定所有类型(pattern, object)
	}
}

// X绑定Map使用映射注册HTTP方法的处理器。 md5:234d05d7bb247514
func (g *X分组路由) X绑定Map(m map[string]interface{}) {
	for pattern, object := range m {
		g.preBindToLocalArray(groupBindTypeHandler, pattern, object)
	}
}

// X绑定GET 函数用于注册一个HTTP处理程序，该程序根据给定的路由模式和HTTP方法（X绑定GET）进行处理。 md5:28790c458e1b962d
func (g *X分组路由) X绑定GET(路由规则 string, 处理函数 interface{}, 额外参数 ...interface{}) *X分组路由 {
	return g.X取副本().preBindToLocalArray(groupBindTypeHandler, "GET:"+路由规则, 处理函数, 额外参数...)
}

// X绑定PUT 注册一个 HTTP 处理器，用于处理给定的路由模式和 HTTP 方法：X绑定PUT。 md5:28ecbdff64685060
func (g *X分组路由) X绑定PUT(路由规则 string, 处理函数 interface{}, 额外参数 ...interface{}) *X分组路由 {
	return g.X取副本().preBindToLocalArray(groupBindTypeHandler, "PUT:"+路由规则, 处理函数, 额外参数...)
}

// X绑定POST 注册一个http处理器，用于给路由模式和HTTP方法：X绑定POST。 md5:a251027c1c7a1d8c
func (g *X分组路由) X绑定POST(路由规则 string, 处理函数 interface{}, 额外参数 ...interface{}) *X分组路由 {
	return g.X取副本().preBindToLocalArray(groupBindTypeHandler, "POST:"+路由规则, 处理函数, 额外参数...)
}

// X绑定DELETE 注册一个 http 处理器，用于给路由模式（pattern）和 http 方法：X绑定DELETE。 md5:b493fe2a753e0422
func (g *X分组路由) X绑定DELETE(路由规则 string, 处理函数 interface{}, 额外参数 ...interface{}) *X分组路由 {
	return g.X取副本().preBindToLocalArray(groupBindTypeHandler, "DELETE:"+路由规则, 处理函数, 额外参数...)
}

// X绑定PATCH 注册一个HTTP处理器，给定路由模式和HTTP方法：X绑定PATCH。 md5:6662f45a2e57a836
func (g *X分组路由) X绑定PATCH(路由规则 string, 处理函数 interface{}, 额外参数 ...interface{}) *X分组路由 {
	return g.X取副本().preBindToLocalArray(groupBindTypeHandler, "PATCH:"+路由规则, 处理函数, 额外参数...)
}

// X绑定HEAD 注册一个http处理器，用于指定路由模式和HTTP方法：X绑定HEAD。 md5:c1e170eaa1fe60b7
func (g *X分组路由) X绑定HEAD(路由规则 string, 处理函数 interface{}, 额外参数 ...interface{}) *X分组路由 {
	return g.X取副本().preBindToLocalArray(groupBindTypeHandler, "HEAD:"+路由规则, 处理函数, 额外参数...)
}

// X绑定CONNECT 注册一个 http 处理器，用于指定路由模式和方法：X绑定CONNECT。 md5:01352b24b5b15d84
func (g *X分组路由) X绑定CONNECT(路由规则 string, 处理函数 interface{}, 额外参数 ...interface{}) *X分组路由 {
	return g.X取副本().preBindToLocalArray(groupBindTypeHandler, "CONNECT:"+路由规则, 处理函数, 额外参数...)
}

// X绑定OPTIONS 注册一个 http 处理器，用于指定路由模式和方法：X绑定OPTIONS。 md5:7c22cd8904d32b99
func (g *X分组路由) X绑定OPTIONS(路由规则 string, 处理函数 interface{}, 额外参数 ...interface{}) *X分组路由 {
	return g.X取副本().preBindToLocalArray(groupBindTypeHandler, "OPTIONS:"+路由规则, 处理函数, 额外参数...)
}

// X绑定TRACE 注册一个HTTP处理程序，用于提供路由模式和HTTP方法：X绑定TRACE。 md5:530929842b31c7fa
func (g *X分组路由) X绑定TRACE(路由规则 string, 处理函数 interface{}, 额外参数 ...interface{}) *X分组路由 {
	return g.X取副本().preBindToLocalArray(groupBindTypeHandler, "TRACE:"+路由规则, 处理函数, 额外参数...)
}

// X绑定RESTfulAPI对象 根据 X绑定RESTfulAPI对象 规则注册一个 HTTP 处理器，以提供路由模式。 md5:b89313386e2f52de
func (g *X分组路由) X绑定RESTfulAPI对象(路由规则 string, 处理对象 interface{}) *X分组路由 {
	return g.X取副本().preBindToLocalArray(groupBindTypeRest, 路由规则, 处理对象)
}

// X绑定Hook 注册一个钩子到给定的路由模式。 md5:1b98e351ffc870a2
func (g *X分组路由) X绑定Hook(路由规则 string, 触发时机 Hook名称, 处理函数 HandlerFunc) *X分组路由 {
	return g.X取副本().preBindToLocalArray(groupBindTypeHandler, 路由规则, 处理函数, 触发时机)
}

// X绑定中间件 将一个或多个中间件绑定到路由器组。 md5:ba25a44638f73d20
func (g *X分组路由) X绑定中间件(处理函数 ...HandlerFunc) *X分组路由 {
	g.middleware = append(g.middleware, 处理函数...)
	return g
}

// preBindToLocalArray 将路由注册参数添加到内部变量数组中，以便于惰性注册特性。 md5:0b2a8a31bb20bca1
func (g *X分组路由) preBindToLocalArray(bindType string, pattern string, object interface{}, params ...interface{}) *X分组路由 {
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

// getPrefix 返回该组的路由前缀，该方法会递归地获取其父组的前缀。 md5:0d086cd9e63f6697
func (g *X分组路由) getPrefix() string {
	prefix := g.prefix
	parent := g.parent
	for parent != nil {
		prefix = parent.prefix + prefix
		parent = parent.parent
	}
	return prefix
}

// doBindRoutersToServer 确实为该组注册。 md5:436447cc3534e54c
func (g *X分组路由) doBindRoutersToServer(ctx context.Context, item *preBindItem) *X分组路由 {
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
				// 如果已经有域，那么在模式中清除域字段。 md5:e02751d36da77b97
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
		// 过滤重复的字符 '/'。 md5:9b9a7539f6ae7305
	pattern = gstr.X替换(pattern, "//", "/")

		// 将参数转换为字符串数组。 md5:8388b98c9b261cad
	extras := gconv.X取文本切片(params)

		// 检查它是否是钩子处理器。 md5:f6b816a5e567ae34
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
				if gstr.X是否包含(extras[0], ",") {
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
								// 最后，它将`object`视为注册类型的对象。 md5:1175240ff3996b3d
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
				HookName: Hook名称(extras[0]),
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
