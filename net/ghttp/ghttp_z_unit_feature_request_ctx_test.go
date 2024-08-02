// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	garray "github.com/888go/goframe/container/garray"
	gbase64 "github.com/888go/goframe/encoding/gbase64"
	"github.com/888go/goframe/frame/g"
	ghttp "github.com/888go/goframe/net/ghttp"
	gtest "github.com/888go/goframe/test/gtest"
	guid "github.com/888go/goframe/util/guid"
)

func Test_Request_IsFileRequest(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Server(guid.S())
		s.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/", func(r *ghttp.Request) {
				r.Response.Write(r.IsFileRequest())
			})
		})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		time.Sleep(100 * time.Millisecond)

		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(c.GetContent(ctx, "/"), false)
	})
}

func Test_Request_IsAjaxRequest(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Server(guid.S())
		s.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/", func(r *ghttp.Request) {
				r.Response.Write(r.IsAjaxRequest())
			})
		})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		time.Sleep(100 * time.Millisecond)

		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(c.GetContent(ctx, "/"), false)
	})
}

func Test_Request_GetClientIp(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Server(guid.S())
		s.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/", func(r *ghttp.Request) {
				r.Response.Write(r.GetClientIp())
			})
		})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		time.Sleep(100 * time.Millisecond)

		c := g.Client()
		c.SetHeader("X-Forwarded-For", "192.168.0.1")
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(c.GetContent(ctx, "/"), "192.168.0.1")
	})
}

func Test_Request_GetUrl(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Server(guid.S())
		s.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/", func(r *ghttp.Request) {
				r.Response.Write(r.GetUrl())
			})
		})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		time.Sleep(100 * time.Millisecond)

		c := g.Client()
		c.SetHeader("X-Forwarded-Proto", "https")
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(c.GetContent(ctx, "/"), fmt.Sprintf("https://127.0.0.1:%d/", s.GetListenedPort()))
	})
}

func Test_Request_GetReferer(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Server(guid.S())
		s.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/", func(r *ghttp.Request) {
				r.Response.Write(r.GetReferer())
			})
		})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		time.Sleep(100 * time.Millisecond)

		c := g.Client()
		c.SetHeader("Referer", "Referer")
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(c.GetContent(ctx, "/"), "Referer")
	})
}

func Test_Request_GetServeHandler(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Server(guid.S())
		s.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/", func(r *ghttp.Request) {
				r.Response.Write(r.GetServeHandler() != nil)
			})
		})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()

		time.Sleep(100 * time.Millisecond)

		c := g.Client()
		c.SetHeader("Referer", "Referer")
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(c.GetContent(ctx, "/"), true)
	})
}

func Test_Request_BasicAuth(t *testing.T) {
	const (
		user      = "root"
		pass      = "123456"
		wrongPass = "12345"
	)

	s := g.Server(guid.S())
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/auth1", func(r *ghttp.Request) {
			r.BasicAuth(user, pass, "tips")
		})
		group.ALL("/auth2", func(r *ghttp.Request) {
			r.BasicAuth(user, pass)
		})
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)

	gtest.C(t, func(t *gtest.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		rsp, err := c.Get(ctx, "/auth1")
		t.AssertNil(err)
		t.Assert(rsp.Header.Get("WWW-Authenticate"), "Basic realm=\"tips\"")
		t.Assert(rsp.StatusCode, http.StatusUnauthorized)

		rsp, err = c.SetHeader("Authorization", user+pass).Get(ctx, "/auth1")
		t.AssertNil(err)
		t.Assert(rsp.StatusCode, http.StatusForbidden)

		rsp, err = c.SetHeader("Authorization", "Test "+user+pass).Get(ctx, "/auth1")
		t.AssertNil(err)
		t.Assert(rsp.StatusCode, http.StatusForbidden)

		rsp, err = c.SetHeader("Authorization", "Basic "+user+pass).Get(ctx, "/auth1")
		t.AssertNil(err)
		t.Assert(rsp.StatusCode, http.StatusForbidden)

		rsp, err = c.SetHeader("Authorization", "Basic "+gbase64.EncodeString(user+pass)).Get(ctx, "/auth1")
		t.AssertNil(err)
		t.Assert(rsp.StatusCode, http.StatusForbidden)

		rsp, err = c.SetHeader("Authorization", "Basic "+gbase64.EncodeString(user+":"+wrongPass)).Get(ctx, "/auth1")
		t.AssertNil(err)
		t.Assert(rsp.StatusCode, http.StatusUnauthorized)

		rsp, err = c.BasicAuth(user, pass).Get(ctx, "/auth1")
		t.AssertNil(err)
		t.Assert(rsp.StatusCode, http.StatusOK)

		rsp, err = c.Get(ctx, "/auth2")
		t.AssertNil(err)
		t.Assert(rsp.Header.Get("WWW-Authenticate"), "Basic realm=\"Need Login\"")
		t.Assert(rsp.StatusCode, http.StatusUnauthorized)
	})
}

func Test_Request_SetCtx(t *testing.T) {
	type ctxKey string
	const testkey ctxKey = "test"
	s := g.Server(guid.S())
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(func(r *ghttp.Request) {
			ctx := context.WithValue(r.Context(), testkey, 1)
			r.SetCtx(ctx)
			r.Middleware.Next()
		})
		group.ALL("/", func(r *ghttp.Request) {
			r.Response.Write(r.Context().Value(testkey))
		})
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(c.GetContent(ctx, "/"), "1")
	})
}

func Test_Request_GetCtx(t *testing.T) {
	type ctxKey string
	const testkey ctxKey = "test"
	s := g.Server(guid.S())
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(func(r *ghttp.Request) {
			ctx := context.WithValue(r.GetCtx(), testkey, 1)
			r.SetCtx(ctx)
			r.Middleware.Next()
		})
		group.ALL("/", func(r *ghttp.Request) {
			r.Response.Write(r.Context().Value(testkey))
		})
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(c.GetContent(ctx, "/"), "1")
	})
}

func Test_Request_GetCtxVar(t *testing.T) {
	s := g.Server(guid.S())
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(func(r *ghttp.Request) {
			r.Middleware.Next()
		})
		group.GET("/", func(r *ghttp.Request) {
			r.Response.Write(r.GetCtxVar("key", "val"))
		})
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/"), "val")
	})
}

func Test_Request_Form(t *testing.T) {
	type User struct {
		Id   int
		Name string
	}
	s := g.Server(guid.S())
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", func(r *ghttp.Request) {
			r.SetForm("key", "val")
			r.Response.Write(r.GetForm("key"))
		})
		group.ALL("/useDef", func(r *ghttp.Request) {
			r.Response.Write(r.GetForm("key", "defVal"))
		})
		group.ALL("/GetFormMap", func(r *ghttp.Request) {
			r.Response.Write(r.GetFormMap(map[string]interface{}{"key": "val"}))
		})
		group.ALL("/GetFormMap1", func(r *ghttp.Request) {
			r.Response.Write(r.GetFormMap(map[string]interface{}{"array": "val"}))
		})
		group.ALL("/GetFormMapStrVar", func(r *ghttp.Request) {
			if r.Get("a") != nil {
				r.Response.Write(r.GetFormMapStrVar()["a"])
			}
		})
		group.ALL("/GetFormStruct", func(r *ghttp.Request) {
			var user User
			if err := r.GetFormStruct(&user); err != nil {
				r.Response.Write(err.Error())
			} else {
				r.Response.Write(user.Name)
			}
		})
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/"), "val")
		t.Assert(client.GetContent(ctx, "/useDef"), "defVal")
		t.Assert(client.PostContent(ctx, "/GetFormMap"), "{\"key\":\"val\"}")
		t.Assert(client.PostContent(ctx, "/GetFormMap", "array[]=1&array[]=2"), "{\"key\":\"val\"}")
		t.Assert(client.PostContent(ctx, "/GetFormMap1", "array[]=1&array[]=2"), "{\"array\":[\"1\",\"2\"]}")
		t.Assert(client.GetContent(ctx, "/GetFormMapStrVar", "a=1&b=2"), nil)
		t.Assert(client.PostContent(ctx, "/GetFormMapStrVar", "a=1&b=2"), `1`)
		t.Assert(client.PostContent(ctx, "/GetFormStruct", g.Map{
			"id":   1,
			"name": "john",
		}), "john")
	})
}

func Test_Request_NeverDoneCtx_Done(t *testing.T) {
	var array = garray.New(true)
	s := g.Server(guid.S())
	s.BindHandler("/done", func(r *ghttp.Request) {
		var (
			ctx    = r.Context()
			ticker = time.NewTimer(time.Millisecond * 1500)
		)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				array.Append(1)
				return
			case <-ticker.C:
				array.Append(1)
				return
			}
		}
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)

	c := g.Client()
	c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
	gtest.C(t, func(t *gtest.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		go func() {
			result := c.GetContent(ctx, "/done")
			fmt.Println(result)
		}()
		time.Sleep(time.Millisecond * 100)

		t.Assert(array.Len(), 0)
		cancel()

		time.Sleep(time.Millisecond * 500)
		t.Assert(array.Len(), 1)
	})
}

func Test_Request_NeverDoneCtx_NeverDone(t *testing.T) {
	var array = garray.New(true)
	s := g.Server(guid.S())
	s.Use(ghttp.MiddlewareNeverDoneCtx)
	s.BindHandler("/never-done", func(r *ghttp.Request) {
		var (
			ctx    = r.Context()
			ticker = time.NewTimer(time.Millisecond * 1500)
		)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				array.Append(1)
				return
			case <-ticker.C:
				array.Append(1)
				return
			}
		}
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)

	c := g.Client()
	c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
	gtest.C(t, func(t *gtest.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		go func() {
			result := c.GetContent(ctx, "/never-done")
			fmt.Println(result)
		}()
		time.Sleep(time.Millisecond * 100)

		t.Assert(array.Len(), 0)
		cancel()

		time.Sleep(time.Millisecond * 1500)
		t.Assert(array.Len(), 1)
	})
}
