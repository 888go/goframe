// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 网页类_test

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"
	
	"github.com/gorilla/websocket"
	
	"github.com/888go/goframe/debug/gdebug"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/gclient"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/guid"
)

var (
	ctx = context.TODO()
)

func Test_Client_Basic(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/hello", func(r *http类.Request) {
		r.Response.X写响应缓冲区("hello")
	})
	s.X绑定("/postForm", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.Get别名("key1"))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		url := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(url)

		t.Assert(g.X网页类().Get文本(ctx, ""), ``)
		t.Assert(client.Get文本(ctx, "/hello"), `hello`)

		_, err := g.X网页类().Post响应对象(ctx, "")
		t.AssertNE(err, nil)

		_, err = g.X网页类().Post表单响应对象(ctx, "/postForm", nil)
		t.AssertNE(err, nil)
		data, _ := g.X网页类().Post表单响应对象(ctx, url+"/postForm", map[string]string{
			"key1": "value1",
		})
		t.Assert(data.X取响应文本(), "value1")
	})
}

func Test_Client_BasicAuth(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/auth", func(r *http类.Request) {
		if r.X账号密码认证("admin", "admin") {
			r.Response.X写响应缓冲区("1")
		} else {
			r.Response.X写响应缓冲区("0")
		}
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(c.X账号密码("admin", "123").Get文本(ctx, "/auth"), "0")
		t.Assert(c.X账号密码("admin", "admin").Get文本(ctx, "/auth"), "1")
	})
}

func Test_Client_Cookie(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/cookie", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.Cookie.X取值("test"))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		c.X设置cookie("test", "0123456789")
		t.Assert(c.Post文本(ctx, "/cookie"), "0123456789")
	})
}

func Test_Client_MapParam(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/map", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.Get别名("test"))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(c.Get文本(ctx, "/map", g.Map{"test": "1234567890"}), "1234567890")
	})
}

func Test_Client_Cookies(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/cookie", func(r *http类.Request) {
		r.Cookie.X设置值("test1", "1")
		r.Cookie.X设置值("test2", "2")
		r.Response.X写响应缓冲区("ok")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		resp, err := c.Get响应对象(ctx, "/cookie")
		t.AssertNil(err)
		defer resp.X关闭()

		t.AssertNE(resp.Header.Get("Set-Cookie"), "")

		m := resp.X取CookieMap()
		t.Assert(len(m), 2)
		t.Assert(m["test1"], 1)
		t.Assert(m["test2"], 2)
		t.Assert(resp.X取Cookie("test1"), 1)
		t.Assert(resp.X取Cookie("test2"), 2)
	})
}

func Test_Client_Chain_Header(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/header1", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.Header.Get("test1"))
	})
	s.X绑定("/header2", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.Header.Get("test2"))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(c.X协议头(g.MapStrStr{"test1": "1234567890"}).Get文本(ctx, "/header1"), "1234567890")
		t.Assert(c.X原始协议头("test1: 1234567890\ntest2: abcdefg").Get文本(ctx, "/header1"), "1234567890")
		t.Assert(c.X原始协议头("test1: 1234567890\ntest2: abcdefg").Get文本(ctx, "/header2"), "abcdefg")
	})
}

func Test_Client_Chain_Context(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/context", func(r *http类.Request) {
		time.Sleep(1 * time.Second)
		r.Response.X写响应缓冲区("ok")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		ctx, _ := context.WithTimeout(context.Background(), 100*time.Millisecond)
		t.Assert(c.Get文本(ctx, "/context"), "")

		ctx, _ = context.WithTimeout(context.Background(), 2000*time.Millisecond)
		t.Assert(c.Get文本(ctx, "/context"), "ok")
	})
}

func Test_Client_Chain_Timeout(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/timeout", func(r *http类.Request) {
		time.Sleep(1 * time.Second)
		r.Response.X写响应缓冲区("ok")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(c.X超时(100*time.Millisecond).Get文本(ctx, "/timeout"), "")
		t.Assert(c.X超时(2000*time.Millisecond).Get文本(ctx, "/timeout"), "ok")
	})
}

func Test_Client_Chain_ContentJson(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/json", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.Get别名("name"), r.Get别名("score"))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(c.X内容类型json().Post文本(ctx, "/json", g.Map{
			"name":  "john",
			"score": 100,
		}), "john100")
		t.Assert(c.X内容类型json().Post文本(ctx, "/json", `{"name":"john", "score":100}`), "john100")

		type User struct {
			Name  string `json:"name"`
			Score int    `json:"score"`
		}
		t.Assert(c.X内容类型json().Post文本(ctx, "/json", User{"john", 100}), "john100")
	})
}

func Test_Client_Chain_ContentXml(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/xml", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.Get别名("name"), r.Get别名("score"))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(c.X内容类型xml().Post文本(ctx, "/xml", g.Map{
			"name":  "john",
			"score": 100,
		}), "john100")
		t.Assert(c.X内容类型xml().Post文本(ctx, "/xml", `{"name":"john", "score":100}`), "john100")

		type User struct {
			Name  string `json:"name"`
			Score int    `json:"score"`
		}
		t.Assert(c.X内容类型xml().Post文本(ctx, "/xml", User{"john", 100}), "john100")
	})
}

func Test_Client_Param_Containing_Special_Char(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/", func(r *http类.Request) {
		r.Response.X写响应缓冲区("k1=", r.Get别名("k1"), "&k2=", r.Get别名("k2"))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(c.Post文本(ctx, "/", "k1=MTIxMg==&k2=100"), "k1=MTIxMg==&k2=100")
		t.Assert(c.Post文本(ctx, "/", g.Map{
			"k1": "MTIxMg==",
			"k2": 100,
		}), "k1=MTIxMg==&k2=100")
	})
}

// 它在上传文件的同时发送数据。
// 它不会对参数进行URL编码。
func Test_Client_File_And_Param(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/", func(r *http类.Request) {
		tmpPath := 文件类.X取临时目录(uid类.X生成())
		err := 文件类.X创建目录(tmpPath)
		单元测试类.AssertNil(err)
		defer 文件类.X删除(tmpPath)

		file := r.X取上传文件对象("file")
		_, err = file.X保存(tmpPath)
		单元测试类.AssertNil(err)
		r.Response.X写响应缓冲区(
			r.Get别名("json"),
			文件类.X读文本(文件类.X路径生成(tmpPath, 文件类.X路径取文件名(file.Filename))),
		)
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	单元测试类.C(t, func(t *单元测试类.T) {
		path := 单元测试类.DataPath("upload", "file1.txt")
		data := g.Map{
			"file": "@file:" + path,
			"json": `{"uuid": "luijquiopm", "isRelative": false, "fileName": "test111.xls"}`,
		}
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(c.Post文本(ctx, "/", data), data["json"].(string)+文件类.X读文本(path))
	})
}

func Test_Client_Middleware(t *testing.T) {
	s := g.Http类(uid类.X生成())
	isServerHandler := false
	s.X绑定("/", func(r *http类.Request) {
		isServerHandler = true
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			str1 = ""
			str2 = "resp body"
		)
		c := g.X网页类().X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		c.X中间件(func(c *网页类.Client, r *http.Request) (resp *网页类.Response, err error) {
			str1 += "a"
			resp, err = c.Next(r)
			if err != nil {
				return nil, err
			}
			str1 += "b"
			return
		})
		c.X中间件(func(c *网页类.Client, r *http.Request) (resp *网页类.Response, err error) {
			str1 += "c"
			resp, err = c.Next(r)
			if err != nil {
				return nil, err
			}
			str1 += "d"
			return
		})
		c.X中间件(func(c *网页类.Client, r *http.Request) (resp *网页类.Response, err error) {
			str1 += "e"
			resp, err = c.Next(r)
			if err != nil {
				return nil, err
			}
			resp.Response.Body = io.NopCloser(bytes.NewBufferString(str2))
			str1 += "f"
			return
		})
		resp, err := c.Get响应对象(ctx, "/")
		t.Assert(str1, "acefdb")
		t.AssertNil(err)
		t.Assert(resp.X取响应文本(), str2)
		t.Assert(isServerHandler, true)

		// 测试中止，中止将不会发送
		var (
			str3     = ""
			abortStr = "abort request"
		)

		c = g.X网页类().X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		c.X中间件(func(c *网页类.Client, r *http.Request) (resp *网页类.Response, err error) {
			str3 += "a"
			resp, err = c.Next(r)
			str3 += "b"
			return
		})
		c.X中间件(func(c *网页类.Client, r *http.Request) (*网页类.Response, error) {
			str3 += "c"
			return nil, 错误类.X创建(abortStr)
		})
		c.X中间件(func(c *网页类.Client, r *http.Request) (resp *网页类.Response, err error) {
			str3 += "f"
			resp, err = c.Next(r)
			str3 += "g"
			return
		})
		resp, err = c.Get响应对象(ctx, "/")
		t.Assert(err, abortStr)
		t.Assert(str3, "acb")
		t.Assert(resp, nil)
	})
}

func Test_Client_Agent(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.UserAgent())
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类().X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		c.X设置UA("test")
		t.Assert(c.Get文本(ctx, "/"), "test")
	})
}

func Test_Client_Request_13_Dump(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/hello", func(r *http类.Request) {
		r.Response.WriteHeader(200)
		r.Response.X写响应缓冲区JSON(g.Map{"field": "test_for_response_body"})
	})
	s.X绑定("/hello2", func(r *http类.Request) {
		r.Response.WriteHeader(200)
		r.Response.X写响应缓冲区并换行(g.Map{"field": "test_for_response_body"})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		url := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类().X设置url前缀(url).X内容类型json()
		r, err := client.Post响应对象(ctx, "/hello", g.Map{"field": "test_for_request_body"})
		t.AssertNil(err)
		dumpedText := r.X取请求原始文本()
		t.Assert(文本类.X是否包含(dumpedText, "test_for_request_body"), true)
		dumpedText2 := r.X取响应原始文本()
		fmt.Println(dumpedText2)
		t.Assert(文本类.X是否包含(dumpedText2, "test_for_response_body"), true)

		client2 := g.X网页类().X设置url前缀(url).X内容类型("text/html")
		r2, err := client2.Post响应对象(ctx, "/hello2", g.Map{"field": "test_for_request_body"})
		t.AssertNil(err)
		dumpedText3 := r2.X取请求原始文本()
		t.Assert(文本类.X是否包含(dumpedText3, "test_for_request_body"), true)
		dumpedText4 := r2.X取响应原始文本()
		t.Assert(文本类.X是否包含(dumpedText4, "test_for_request_body"), false)
		r2 = nil
		t.Assert(r2.X取请求原始文本(), "")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		url := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		response, _ := g.X网页类().Get响应对象(ctx, url, g.Map{
			"id":   10000,
			"name": "john",
		})
		response = nil
		t.Assert(response.X取请求原始文本(), "")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		url := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		response, _ := g.X网页类().Get响应对象(ctx, url, g.Map{
			"id":   10000,
			"name": "john",
		})
		response.X请求和响应输出终端()
		t.AssertGT(len(response.X取请求和响应原始文本()), 0)
	})
}

func Test_WebSocketClient(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/ws", func(r *http类.Request) {
		ws, err := r.X升级为websocket请求()
		if err != nil {
			r.X退出当前()
		}
		for {
			msgType, msg, err := ws.ReadMessage()
			if err != nil {
				return
			}
			if err = ws.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
// 由于WebSocket保持连接的特性，可能存在数据竞争情况，因此不在此处关闭。// 延迟调用 s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := 网页类.X创建WebSocket()
		client.Proxy = http.ProxyFromEnvironment
		client.HandshakeTimeout = time.Minute

		conn, _, err := client.Dial(fmt.Sprintf("ws://127.0.0.1:%d/ws", s.X取已监听端口()), nil)
		t.AssertNil(err)
		defer conn.Close()

		msg := []byte("hello")
		err = conn.WriteMessage(websocket.TextMessage, msg)
		t.AssertNil(err)

		mt, data, err := conn.ReadMessage()
		t.AssertNil(err)
		t.Assert(mt, websocket.TextMessage)
		t.Assert(data, msg)
	})
}

func TestLoadKeyCrt(t *testing.T) {
	var (
		testCrtFile = 文件类.X路径取父目录(gdebug.CallerFilePath()) + 文件类.Separator + "testdata/upload/file1.txt"
		testKeyFile = 文件类.X路径取父目录(gdebug.CallerFilePath()) + 文件类.Separator + "testdata/upload/file2.txt"
		crtFile     = 文件类.X路径取父目录(gdebug.CallerFilePath()) + 文件类.Separator + "testdata/server.crt"
		keyFile     = 文件类.X路径取父目录(gdebug.CallerFilePath()) + 文件类.Separator + "testdata/server.key"
		tlsConfig   = &tls.Config{}
	)

	单元测试类.C(t, func(t *单元测试类.T) {
		tlsConfig, _ = 网页类.X创建TLS配置("crtFile", "keyFile")
		t.AssertNil(tlsConfig)

		tlsConfig, _ = 网页类.X创建TLS配置(crtFile, "keyFile")
		t.AssertNil(tlsConfig)

		tlsConfig, _ = 网页类.X创建TLS配置(testCrtFile, testKeyFile)
		t.AssertNil(tlsConfig)

		tlsConfig, _ = 网页类.X创建TLS配置(crtFile, keyFile)
		t.AssertNE(tlsConfig, nil)
	})
}

func TestClient_DoRequest(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/hello", func(r *http类.Request) {
		r.Response.WriteHeader(200)
		r.Response.X写响应缓冲区JSON(g.Map{"field": "test_for_response_body"})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		url := fmt.Sprintf("127.0.0.1:%d/hello", s.X取已监听端口())
		resp, err := c.X请求响应对象(ctx, http.MethodGet, url)
		t.AssertNil(err)
		t.AssertNE(resp, nil)
		t.Assert(resp.X取响应文本(), "{\"field\":\"test_for_response_body\"}")

		resp.Response = nil
		bytes := resp.X取响应字节集()
		t.Assert(bytes, []byte{})
		resp.X关闭()
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		url := "127.0.0.1:99999/hello"
		resp, err := c.X请求响应对象(ctx, http.MethodGet, url)
		t.AssertNil(resp.Response)
		t.AssertNE(err, nil)
	})
}

func TestClient_RequestVar(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			url = "http://127.0.0.1:99999/var/jsons"
		)
		varValue := g.X网页类().X请求泛型类(ctx, http.MethodGet, url)
		t.AssertNil(varValue)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Id   int
			Name string
		}
		var (
			users []User
			url   = "http://127.0.0.1:8999/var/jsons"
		)
		err := g.X网页类().X请求泛型类(ctx, http.MethodGet, url).X取结构体指针(&users)
		t.AssertNil(err)
		t.AssertNE(users, nil)
	})
}

func TestClient_SetBodyContent(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/", func(r *http类.Request) {
		r.Response.X写响应缓冲区("hello")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		res, err := c.Get响应对象(ctx, "/")
		t.AssertNil(err)
		defer res.X关闭()
		t.Assert(res.X取响应文本(), "hello")
		res.X覆盖响应内容([]byte("world"))
		t.Assert(res.X取响应文本(), "world")
	})
}

func TestClient_SetNoUrlEncode(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.URL.RawQuery)
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		var params = g.Map{
			"path": "/data/binlog",
		}
		t.Assert(c.Get文本(ctx, `/`, params), `path=%2Fdata%2Fbinlog`)

		c.X设置请求参数禁止URL编码(true)
		t.Assert(c.Get文本(ctx, `/`, params), `path=/data/binlog`)

		c.X设置请求参数禁止URL编码(false)
		t.Assert(c.Get文本(ctx, `/`, params), `path=%2Fdata%2Fbinlog`)
	})
}

func TestClient_NoUrlEncode(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.URL.RawQuery)
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		var params = g.Map{
			"path": "/data/binlog",
		}
		t.Assert(c.Get文本(ctx, `/`, params), `path=%2Fdata%2Fbinlog`)

		t.Assert(c.X请求参数禁止URL编码().Get文本(ctx, `/`, params), `path=/data/binlog`)
	})
}
