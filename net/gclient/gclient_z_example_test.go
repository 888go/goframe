// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 网页类_test

import (
	"context"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"
	
	"github.com/888go/goframe/debug/gdebug"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/gclient"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/os/gctx"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/util/guid"
)

var (
	crtFile = 文件类.X路径取父目录(gdebug.CallerFilePath()) + 文件类.Separator + "testdata/server.crt"
	keyFile = 文件类.X路径取父目录(gdebug.CallerFilePath()) + 文件类.Separator + "testdata/server.key"
)

func init() {
	// 默认的客户端服务器。
	p := 8999
	s := g.Http类(p)
	// HTTP方法处理器。
	s.X创建分组路由("/", func(group *http类.X分组路由) {
		group.X绑定GET("/", func(r *http类.X请求) {
			r.X响应.X写响应缓冲区并格式化(
				"GET: query: %d, %s",
				r.X取查询参数到泛型类("id").X取整数(),
				r.X取查询参数到泛型类("name").String(),
			)
		})
		group.X绑定PUT("/", func(r *http类.X请求) {
			r.X响应.X写响应缓冲区并格式化(
				"PUT: form: %d, %s",
				r.X取表单值到泛型类("id").X取整数(),
				r.X取表单值到泛型类("name").String(),
			)
		})
		group.X绑定POST("/", func(r *http类.X请求) {
			r.X响应.X写响应缓冲区并格式化(
				"POST: form: %d, %s",
				r.X取表单值到泛型类("id").X取整数(),
				r.X取表单值到泛型类("name").String(),
			)
		})
		group.X绑定DELETE("/", func(r *http类.X请求) {
			r.X响应.X写响应缓冲区并格式化(
				"DELETE: form: %d, %s",
				r.X取表单值到泛型类("id").X取整数(),
				r.X取表单值到泛型类("name").String(),
			)
		})
		group.X绑定HEAD("/", func(r *http类.X请求) {
			r.X响应.X写响应缓冲区并格式化(
				"HEAD: form: %d, %s",
				r.X取表单值到泛型类("id").X取整数(),
				r.X取表单值到泛型类("name").String(),
			)
		})
		group.X绑定PATCH("/", func(r *http类.X请求) {
			r.X响应.X写响应缓冲区并格式化(
				"PATCH: form: %d, %s",
				r.X取表单值到泛型类("id").X取整数(),
				r.X取表单值到泛型类("name").String(),
			)
		})
		group.X绑定CONNECT("/", func(r *http类.X请求) {
			r.X响应.X写响应缓冲区并格式化(
				"CONNECT: form: %d, %s",
				r.X取表单值到泛型类("id").X取整数(),
				r.X取表单值到泛型类("name").String(),
			)
		})
		group.X绑定OPTIONS("/", func(r *http类.X请求) {
			r.X响应.X写响应缓冲区并格式化(
				"OPTIONS: form: %d, %s",
				r.X取表单值到泛型类("id").X取整数(),
				r.X取表单值到泛型类("name").String(),
			)
		})
		group.X绑定TRACE("/", func(r *http类.X请求) {
			r.X响应.X写响应缓冲区并格式化(
				"TRACE: form: %d, %s",
				r.X取表单值到泛型类("id").X取整数(),
				r.X取表单值到泛型类("name").String(),
			)
		})
	})
	// 客户端链式操作处理程序。
	s.X创建分组路由("/", func(group *http类.X分组路由) {
		group.X绑定所有类型("/header", func(r *http类.X请求) {
			r.X响应.X写响应缓冲区并格式化(
				"Span-Id: %s, Trace-Id: %s",
				r.Header.Get("Span-Id"),
				r.Header.Get("Trace-Id"),
			)
		})
		group.X绑定所有类型("/cookie", func(r *http类.X请求) {
			r.X响应.X写响应缓冲区并格式化(
				"SessionId: %s",
				r.Cookie.X取值("SessionId"),
			)
		})
		group.X绑定所有类型("/json", func(r *http类.X请求) {
			r.X响应.X写响应缓冲区并格式化(
				"Content-Type: %s, id: %d",
				r.Header.Get("Content-Type"),
				r.Get别名("id").X取整数(),
			)
		})
	})
	// 其他测试处理程序。
	s.X创建分组路由("/var", func(group *http类.X分组路由) {
		group.X绑定所有类型("/json", func(r *http类.X请求) {
			r.X响应.X写响应缓冲区(`{"id":1,"name":"john"}`)
		})
		group.X绑定所有类型("/jsons", func(r *http类.X请求) {
			r.X响应.X写响应缓冲区(`[{"id":1,"name":"john"}, {"id":2,"name":"smith"}]`)
		})
	})
	s.X设置日志开启访客记录(false)
	s.SetDumpRouterMap(false)
	s.X设置监听端口(p)
	err := s.X开始监听()
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Millisecond * 500)
}

func ExampleNew() {
	var (
		ctx    = 上下文类.X创建()
		client = 网页类.X创建()
	)

	if r, err := client.Get响应对象(ctx, "http://127.0.0.1:8999/var/json"); err != nil {
		panic(err)
	} else {
		defer r.X关闭()
		fmt.Println(r.X取响应文本())
	}

	// Output:
	// {"id":1,"name":"john"}
}

func ExampleClient_Clone() {
	var (
		ctx    = 上下文类.X创建()
		client = 网页类.X创建()
	)

	client.X设置cookie("key", "value")
	cloneClient := client.X取副本()

	if r, err := cloneClient.Get响应对象(ctx, "http://127.0.0.1:8999/var/json"); err != nil {
		panic(err)
	} else {
		defer r.X关闭()
		fmt.Println(r.X取响应文本())
	}

	// Output:
	// {"id":1,"name":"john"}
}

func fromHex(s string) []byte {
	b, _ := hex.DecodeString(s)
	return b
}

func ExampleNew_MultiConn_Recommend() {
	var (
		ctx    = 上下文类.X创建()
		client = g.X网页类()
	)

	// 控制每个主机保持的最大空闲(keep-alive)连接数
	client.Transport.(*http.Transport).MaxIdleConnsPerHost = 5

	for i := 0; i < 5; i++ {
		if r, err := client.Get响应对象(ctx, "http://127.0.0.1:8999/var/json"); err != nil {
			panic(err)
		} else {
			fmt.Println(r.X取响应文本())
			r.X关闭()
		}
	}

	// Output:
	//{"id":1,"name":"john"}
	//{"id":1,"name":"john"}
	//{"id":1,"name":"john"}
	//{"id":1,"name":"john"}
	//{"id":1,"name":"john"}
}

func ExampleClient_Header() {
	var (
		url    = "http://127.0.0.1:8999/header"
		header = g.MapStrStr{
			"Span-Id":  "0.1",
			"Trace-Id": "123456789",
		}
	)
	content := g.X网页类().X协议头(header).Post文本(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})
	fmt.Println(content)

	// Output:
	// Span-Id: 0.1, Trace-Id: 123456789
}

func ExampleClient_HeaderRaw() {
	var (
		url       = "http://127.0.0.1:8999/header"
		headerRaw = `
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3950.0 Safari/537.36
Span-Id: 0.1
Trace-Id: 123456789
`
	)
	content := g.X网页类().X原始协议头(headerRaw).Post文本(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})
	fmt.Println(content)

	// Output:
	// Span-Id: 0.1, Trace-Id: 123456789
}

func ExampleClient_Cookie() {
	var (
		url    = "http://127.0.0.1:8999/cookie"
		cookie = g.MapStrStr{
			"SessionId": "123",
		}
	)
	content := g.X网页类().Cookie(cookie).Post文本(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})
	fmt.Println(content)

	// Output:
	// SessionId: 123
}

func ExampleClient_ContentJson() {
	var (
		url     = "http://127.0.0.1:8999/json"
		jsonStr = `{"id":10000,"name":"john"}`
		jsonMap = g.Map{
			"id":   10000,
			"name": "john",
		}
	)
	// 使用JSON字符串进行POST请求。
	fmt.Println(g.X网页类().X内容类型json().Post文本(ctx, url, jsonStr))
	// 使用JSON映射进行POST请求。
	fmt.Println(g.X网页类().X内容类型json().Post文本(ctx, url, jsonMap))

	// Output:
	// Content-Type: application/json, id: 10000
	// Content-Type: application/json, id: 10000
}

func ExampleClient_Post() {
	url := "http://127.0.0.1:8999"
	// 使用字符串参数作为请求体发送。
	r1, err := g.X网页类().Post响应对象(ctx, url, "id=10000&name=john")
	if err != nil {
		panic(err)
	}
	defer r1.X关闭()
	fmt.Println(r1.X取响应文本())

	// 使用map参数发送。
	r2, err := g.X网页类().Post响应对象(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})
	if err != nil {
		panic(err)
	}
	defer r2.X关闭()
	fmt.Println(r2.X取响应文本())

	// Output:
	// POST: form: 10000, john
	// POST: form: 10000, john
}

func ExampleClient_PostBytes() {
	url := "http://127.0.0.1:8999"
	fmt.Println(string(g.X网页类().Post字节集(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})))

	// Output:
	// POST: form: 10000, john
}

func ExampleClient_DeleteBytes() {
	url := "http://127.0.0.1:8999"
	fmt.Println(string(g.X网页类().Delete字节集(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})))

	// Output:
	// DELETE: form: 10000, john
}

func ExampleClient_HeadBytes() {
	url := "http://127.0.0.1:8999"
	fmt.Println(string(g.X网页类().Head字节集(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})))

	// Output:
}

func ExampleClient_PatchBytes() {
	url := "http://127.0.0.1:8999"
	fmt.Println(string(g.X网页类().Patch字节集(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})))

	// Output:
	// PATCH: form: 10000, john
}

func ExampleClient_ConnectBytes() {
	url := "http://127.0.0.1:8999"
	fmt.Println(string(g.X网页类().Connect字节集(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})))

	// Output:
	// CONNECT: form: 10000, john
}

func ExampleClient_OptionsBytes() {
	url := "http://127.0.0.1:8999"
	fmt.Println(string(g.X网页类().Options字节集(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})))

	// Output:
	// OPTIONS: form: 10000, john
}

func ExampleClient_TraceBytes() {
	url := "http://127.0.0.1:8999"
	fmt.Println(string(g.X网页类().Trace字节集(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})))

	// Output:
	// TRACE: form: 10000, john
}

func ExampleClient_PostContent() {
	url := "http://127.0.0.1:8999"
	fmt.Println(g.X网页类().Post文本(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	}))

	// Output:
	// POST: form: 10000, john
}

func ExampleClient_PostVar() {
	type User struct {
		Id   int
		Name string
	}
	var (
		users []User
		url   = "http://127.0.0.1:8999/var/jsons"
	)
	err := g.X网页类().Post泛型类(ctx, url).X取结构体指针(&users)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)

	// Output:
	// [{1 john} {2 smith}]
}

func ExampleClient_Get() {
	var (
		ctx = context.Background()
		url = "http://127.0.0.1:8999"
	)

	// 使用字符串参数并通过URL发送。
	r1, err := g.X网页类().Get响应对象(ctx, url+"?id=10000&name=john")
	if err != nil {
		panic(err)
	}
	defer r1.X关闭()
	fmt.Println(r1.X取响应文本())

	// 使用字符串参数作为请求体发送。
	r2, err := g.X网页类().Get响应对象(ctx, url, "id=10000&name=john")
	if err != nil {
		panic(err)
	}
	defer r2.X关闭()
	fmt.Println(r2.X取响应文本())

	// 使用map参数发送。
	r3, err := g.X网页类().Get响应对象(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})
	if err != nil {
		panic(err)
	}
	defer r3.X关闭()
	fmt.Println(r3.X取响应文本())

	// Output:
	// GET: query: 10000, john
	// GET: query: 10000, john
	// GET: query: 10000, john
}

func ExampleClient_Put() {
	url := "http://127.0.0.1:8999"
	r, _ := g.X网页类().Put响应对象(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})
	defer r.X关闭()
	fmt.Println(r.X取响应文本())

	// Output:
	// PUT: form: 10000, john
}

func ExampleClient_GetBytes() {
	var (
		ctx = context.Background()
		url = "http://127.0.0.1:8999"
	)
	fmt.Println(string(g.X网页类().Get字节集(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})))

	// Output:
	// GET: query: 10000, john
}

func ExampleClient_PutBytes() {
	var (
		ctx = context.Background()
		url = "http://127.0.0.1:8999"
	)
	fmt.Println(string(g.X网页类().Put字节集(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})))

	// Output:
	// PUT: form: 10000, john
}

func ExampleClient_GetContent() {
	url := "http://127.0.0.1:8999"
	fmt.Println(g.X网页类().Get文本(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	}))

	// Output:
	// GET: query: 10000, john
}

func ExampleClient_GetVar() {
	type User struct {
		Id   int
		Name string
	}
	var (
		user *User
		ctx  = context.Background()
		url  = "http://127.0.0.1:8999/var/json"
	)
	err := g.X网页类().Get泛型类(ctx, url).X取结构体指针(&user)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

	// Output:
	// &{1 john}
}

// ExampleClient_SetProxy 是 `gclient.Client.SetProxy` 方法的示例。
// 请在运行此示例前准备好两个代理服务器。
// HTTP 代理服务器监听 `127.0.0.1:1081`
// SOCKS5 代理服务器监听 `127.0.0.1:1080`
func ExampleClient_SetProxy() {
	// 连接到一个HTTP代理服务器
	client := g.X网页类()
	client.X设置代理("http://127.0.0.1:1081")
	client.X设置超时(5 * time.Second) // it's suggested to set http client timeout
	resp, err := client.Get响应对象(ctx, "http://127.0.0.1:8999")
	if err != nil {
// 当你的代理服务器无法访问时，err将不为nil。
// 例如：尝试获取"http://127.0.0.1:8999"时，出现如下错误：
//       通过代理连接tcp：拨号tcp到127.0.0.1:1087时出错：连接被拒绝
// （注：原文中的 "proxyconnect tcp: dial tcp" 是golang中通过代理连接远程地址时可能出现的错误信息，意指在建立TCP连接至代理服务器的过程中出现了问题，具体表现为连接被拒绝。）
	}
	fmt.Println(err != nil)
	resp.X关闭()

	// 连接到一个HTTP代理服务器 which needs auth
	client.X设置代理("http://user:password:127.0.0.1:1081")
	client.X设置超时(5 * time.Second) // it's suggested to set http client timeout
	resp, err = client.Get响应对象(ctx, "http://127.0.0.1:8999")
	if err != nil {
// 当你的代理服务器无法访问时，err将不为nil。
// 例如：尝试获取"http://127.0.0.1:8999"时，出现如下错误：
//       通过代理连接tcp：拨号tcp到127.0.0.1:1087时出错：连接被拒绝
// （注：原文中的 "proxyconnect tcp: dial tcp" 是golang中通过代理连接远程地址时可能出现的错误信息，意指在建立TCP连接至代理服务器的过程中出现了问题，具体表现为连接被拒绝。）
	}
	fmt.Println(err != nil)
	resp.X关闭()

	// 连接到一个SOCKS5代理服务器
	client.X设置代理("socks5://127.0.0.1:1080")
	client.X设置超时(5 * time.Second) // it's suggested to set http client timeout
	resp, err = client.Get响应对象(ctx, "http://127.0.0.1:8999")
	if err != nil {
// 当您的代理服务器不可用时，err将不为nil。
// 例如：获取"http://127.0.0.1:8999"时出错，错误信息为：
// "socks连接tcp 127.0.0.1:1087->api.ip.sb:443时发生错误：拨号tcp 127.0.0.1:1087失败，原因：连接被拒绝"
// 这段Go语言代码的注释翻译成中文是：
// ```go
// 如果您的代理服务器宕机，err将不会是nil（即会返回一个非空错误）。
// 比如：尝试获取"http://127.0.0.1:8999"时，
// 出现如下错误："通过SOCKS协议连接到tcp 127.0.0.1:1087并试图转发到api.ip.sb:443的过程中发生错误，
// 具体原因为：尝试连接到tcp 127.0.0.1:1087时被拒绝，连接无法建立。"
	}
	fmt.Println(err != nil)
	resp.X关闭()

	// 连接到一个SOCKS5代理服务器 which needs auth
	client.X设置代理("socks5://user:password@127.0.0.1:1080")
	client.X设置超时(5 * time.Second) // it's suggested to set http client timeout
	resp, err = client.Get响应对象(ctx, "http://127.0.0.1:8999")
	if err != nil {
// 当您的代理服务器不可用时，err将不为nil。
// 例如：获取"http://127.0.0.1:8999"时出错，错误信息为：
// "socks连接tcp 127.0.0.1:1087->api.ip.sb:443时发生错误：拨号tcp 127.0.0.1:1087失败，原因：连接被拒绝"
// 这段Go语言代码的注释翻译成中文是：
// ```go
// 如果您的代理服务器宕机，err将不会是nil（即会返回一个非空错误）。
// 比如：尝试获取"http://127.0.0.1:8999"时，
// 出现如下错误："通过SOCKS协议连接到tcp 127.0.0.1:1087并试图转发到api.ip.sb:443的过程中发生错误，
// 具体原因为：尝试连接到tcp 127.0.0.1:1087时被拒绝，连接无法建立。"
	}
	fmt.Println(err != nil)
	resp.X关闭()

	// Output:
	// true
	// true
	// true
	// true
}

// ExampleClientChain_Proxy 是 `gclient.Client.Proxy` 方法的链式版本示例。
// 在运行此示例前，请确保已准备好两个代理服务器。
// HTTP 代理服务器监听地址为 `127.0.0.1:1081`
// SOCKS5 代理服务器监听地址为 `127.0.0.1:1080`
// 更多详情，请参考 ExampleClient_SetProxy 示例。
func ExampleClient_Proxy() {
	var (
		ctx = context.Background()
	)
	client := g.X网页类()
	_, err := client.X代理("http://127.0.0.1:1081").Get响应对象(ctx, "http://127.0.0.1:8999")
	fmt.Println(err != nil)

	client2 := g.X网页类()
	_, err = client2.X代理("socks5://127.0.0.1:1080").Get响应对象(ctx, "http://127.0.0.1:8999")
	fmt.Println(err != nil)

	client3 := g.X网页类()
	_, err = client3.X代理("").Get响应对象(ctx, "http://127.0.0.1:8999")
	fmt.Println(err != nil)

	client4 := g.X网页类()
	url := "http://127.0.0.1:1081" + string([]byte{0x7f})
	_, err = client4.X代理(url).Get响应对象(ctx, "http://127.0.0.1:8999")
	fmt.Println(err != nil)

	// Output:
	// true
	// true
	// false
	// false
}

func ExampleClient_Prefix() {
	var (
		ctx = 上下文类.X创建()
	)

	s := g.Http类(uid类.X生成())
	// HTTP方法处理器。
	s.X创建分组路由("/api", func(group *http类.X分组路由) {
		group.X绑定GET("/v1/prefix", func(r *http类.X请求) {
			r.X响应.X写响应缓冲区("this is v1 prefix")
		})
		group.X绑定GET("/v1/hello", func(r *http类.X请求) {
			r.X响应.X写响应缓冲区("this is v1 hello")
		})
	})
	s.X设置日志开启访客记录(false)
	s.SetDumpRouterMap(false)
	s.X开始监听()
	time.Sleep(time.Millisecond * 100)

	// 添加客户端URI前缀
	client := g.X网页类().Url前缀(fmt.Sprintf(
		"http://127.0.0.1:%d/api/v1/", s.X取已监听端口(),
	))

	fmt.Println(string(client.Get字节集(ctx, "prefix")))
	fmt.Println(string(client.Get字节集(ctx, "hello")))

	// Output:
	// this is v1 prefix
	// this is v1 hello
}

func ExampleClient_Retry() {
	var (
		ctx = 上下文类.X创建()
		url = "http://127.0.0.1:8999"
	)
	client := g.X网页类().X重试与间隔(2, time.Second)

	fmt.Println(string(client.Get字节集(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})))

	// Output:
	// GET: query: 10000, john
}

func ExampleClient_RedirectLimit() {
	var (
		ctx = 上下文类.X创建()
		url = "http://127.0.0.1:8999"
	)
	client := g.X网页类().X重定向次数限制(1)

	fmt.Println(string(client.Get字节集(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})))

	// Output:
	// GET: query: 10000, john
}

func ExampleClient_SetBrowserMode() {
	var (
		ctx = 上下文类.X创建()
		url = "http://127.0.0.1:8999"
	)
	client := g.X网页类().X启用浏览器模式(true)

	fmt.Println(string(client.Get字节集(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})))

	// Output:
	// GET: query: 10000, john
}

func ExampleClient_SetHeader() {
	var (
		ctx = 上下文类.X创建()
		url = "http://127.0.0.1:8999"
	)
	client := g.X网页类()
	client.X设置协议头("Server", "GoFrameServer")
	client.X设置协议头("Client", "g.Client()")

	fmt.Println(string(client.Get字节集(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})))

	// Output:
	// GET: query: 10000, john
}

func ExampleClient_SetRedirectLimit() {
	go func() {
		s := g.Http类()
		s.X绑定("/hello", func(r *http类.X请求) {
			r.X响应.X写响应缓冲区并换行("hello world")
		})
		s.X绑定("/back", func(r *http类.X请求) {
			r.X响应.X重定向到来源页面()
		})
		s.SetDumpRouterMap(false)
		s.X设置监听端口(8199)
		s.X启动服务()
	}()

	time.Sleep(time.Second)

	var (
		ctx      = 上下文类.X创建()
		urlHello = "http://127.0.0.1:8199/hello"
		urlBack  = "http://127.0.0.1:8199/back"
	)
	client := g.X网页类().X设置重定向次数限制(1)
	client.X设置协议头("Referer", urlHello)

	resp, err := client.X请求响应对象(ctx, http.MethodGet, urlBack, g.Map{
		"id":   10000,
		"name": "john",
	})
	if err == nil {
		fmt.Println(resp.X取响应文本())
		resp.X关闭()
	}

	client.X设置重定向次数限制(2)
	resp, err = client.X请求响应对象(ctx, http.MethodGet, urlBack, g.Map{
		"id":   10000,
		"name": "john",
	})
	if err == nil {
		fmt.Println(resp.X取响应文本())
		resp.X关闭()
	}

	// Output:
	// Found
	// hello world
}

func ExampleClient_SetTLSKeyCrt() {
	var (
		ctx         = 上下文类.X创建()
		url         = "http://127.0.0.1:8999"
		testCrtFile = 文件类.X路径取父目录(gdebug.CallerFilePath()) + 文件类.Separator + "testdata/upload/file1.txt"
		testKeyFile = 文件类.X路径取父目录(gdebug.CallerFilePath()) + 文件类.Separator + "testdata/upload/file2.txt"
	)
	client := g.X网页类()
	client.X设置证书(testCrtFile, testKeyFile)
	client.X设置证书(crtFile, keyFile)
	fmt.Println(string(client.Get字节集(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})))

	// Output:
	// GET: query: 10000, john
}

func ExampleClient_SetTLSConfig() {
	var (
		ctx       = 上下文类.X创建()
		url       = "http://127.0.0.1:8999"
		tlsConfig = &tls.Config{}
	)
	client := g.X网页类()
	client.X设置TLS配置(tlsConfig)
	fmt.Println(string(client.Get字节集(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})))

	// Output:
	// GET: query: 10000, john
}

func ExampleClient_PutContent() {
	url := "http://127.0.0.1:8999"
	fmt.Println(g.X网页类().Put文本(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	}))

	// Output:
	// PUT: form: 10000, john
}

func ExampleClient_DeleteContent() {
	url := "http://127.0.0.1:8999"
	fmt.Println(g.X网页类().Delete文本(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	}))

	// Output:
	// DELETE: form: 10000, john
}

func ExampleClient_HeadContent() {
	url := "http://127.0.0.1:8999"
	fmt.Println(g.X网页类().Head文本(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	}))

	// Output:
}

func ExampleClient_PatchContent() {
	url := "http://127.0.0.1:8999"
	fmt.Println(g.X网页类().Patch文本(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	}))

	// Output:
	// PATCH: form: 10000, john
}

func ExampleClient_ConnectContent() {
	url := "http://127.0.0.1:8999"
	fmt.Println(g.X网页类().Connect文本(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	}))

	// Output:
	// CONNECT: form: 10000, john
}

func ExampleClient_OptionsContent() {
	url := "http://127.0.0.1:8999"
	fmt.Println(g.X网页类().Options文本(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	}))

	// Output:
	// OPTIONS: form: 10000, john
}

func ExampleClient_TraceContent() {
	url := "http://127.0.0.1:8999"
	fmt.Println(g.X网页类().Trace文本(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	}))

	// Output:
	// TRACE: form: 10000, john
}

func ExampleClient_RequestContent() {
	url := "http://127.0.0.1:8999"
	fmt.Println(g.X网页类().X请求文本(ctx, http.MethodGet, url, g.Map{
		"id":   10000,
		"name": "john",
	}))

	// Output:
	// GET: query: 10000, john
}

func ExampleClient_RawRequest() {
	url := "http://127.0.0.1:8999"
	response, _ := g.X网页类().Get响应对象(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})
	fmt.Println(len(response.X取响应原始文本()) > 100)

	// Output:
	// true
}

func ExampleClient_Delete() {
	url := "http://127.0.0.1:8999"
	r, _ := g.X网页类().Delete响应对象(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})
	defer r.X关闭()
	fmt.Println(r.X取响应文本())

	// Output:
	// DELETE: form: 10000, john
}

func ExampleClient_Head() {
	url := "http://127.0.0.1:8999"
	r, _ := g.X网页类().Head响应对象(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})
	defer r.X关闭()
	fmt.Println(r.X取响应文本())

	// Output:
	//
}

func ExampleClient_Patch() {
	url := "http://127.0.0.1:8999"
	r, _ := g.X网页类().Patch响应对象(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})
	defer r.X关闭()
	fmt.Println(r.X取响应文本())

	// Output:
	// PATCH: form: 10000, john
}

func ExampleClient_Connect() {
	url := "http://127.0.0.1:8999"
	r, _ := g.X网页类().Connect响应对象(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})
	defer r.X关闭()
	fmt.Println(r.X取响应文本())

	// Output:
	// CONNECT: form: 10000, john
}

func ExampleClient_Options() {
	url := "http://127.0.0.1:8999"
	r, _ := g.X网页类().Options响应对象(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})
	defer r.X关闭()
	fmt.Println(r.X取响应文本())

	// Output:
	// OPTIONS: form: 10000, john
}

func ExampleClient_Trace() {
	url := "http://127.0.0.1:8999"
	r, _ := g.X网页类().Trace响应对象(ctx, url, g.Map{
		"id":   10000,
		"name": "john",
	})
	defer r.X关闭()
	fmt.Println(r.X取响应文本())

	// Output:
	// TRACE: form: 10000, john
}

func ExampleClient_PutVar() {
	type User struct {
		Id   int
		Name string
	}
	var (
		user *User
		ctx  = context.Background()
		url  = "http://127.0.0.1:8999/var/json"
	)
	err := g.X网页类().Put泛型类(ctx, url).X取结构体指针(&user)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

	// Output:
	// &{1 john}
}

func ExampleClient_DeleteVar() {
	type User struct {
		Id   int
		Name string
	}
	var (
		users []User
		url   = "http://127.0.0.1:8999/var/jsons"
	)
	err := g.X网页类().Delete泛型类(ctx, url).X取结构体指针(&users)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)

	// Output:
	// [{1 john} {2 smith}]
}

func ExampleClient_HeadVar() {
	type User struct {
		Id   int
		Name string
	}
	var (
		users []User
		url   = "http://127.0.0.1:8999/var/jsons"
	)
	err := g.X网页类().Head泛型类(ctx, url).X取结构体指针(&users)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)

	// Output:
	// []
}

func ExampleClient_PatchVar() {
	type User struct {
		Id   int
		Name string
	}
	var (
		users []User
		url   = "http://127.0.0.1:8999/var/jsons"
	)
	err := g.X网页类().Patch泛型类(ctx, url).X取结构体指针(&users)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)

	// Output:
	// [{1 john} {2 smith}]
}

func ExampleClient_ConnectVar() {
	type User struct {
		Id   int
		Name string
	}
	var (
		users []User
		url   = "http://127.0.0.1:8999/var/jsons"
	)
	err := g.X网页类().Connect泛型类(ctx, url).X取结构体指针(&users)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)

	// Output:
	// [{1 john} {2 smith}]
}

func ExampleClient_OptionsVar() {
	type User struct {
		Id   int
		Name string
	}
	var (
		users []User
		url   = "http://127.0.0.1:8999/var/jsons"
	)
	err := g.X网页类().Options泛型类(ctx, url).X取结构体指针(&users)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)

	// Output:
	// [{1 john} {2 smith}]
}

func ExampleClient_TraceVar() {
	type User struct {
		Id   int
		Name string
	}
	var (
		users []User
		url   = "http://127.0.0.1:8999/var/jsons"
	)
	err := g.X网页类().Trace泛型类(ctx, url).X取结构体指针(&users)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)

	// Output:
	// [{1 john} {2 smith}]
}
