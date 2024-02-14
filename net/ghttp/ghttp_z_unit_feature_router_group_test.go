// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类_test

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/guid"
)

// 执行对象
type GroupObject struct{}

func (o *GroupObject) Init(r *http类.X请求) {
	r.X响应.X写响应缓冲区("1")
}

func (o *GroupObject) Shut(r *http类.X请求) {
	r.X响应.X写响应缓冲区("2")
}

func (o *GroupObject) Index(r *http类.X请求) {
	r.X响应.X写响应缓冲区("Object Index")
}

func (o *GroupObject) Show(r *http类.X请求) {
	r.X响应.X写响应缓冲区("Object Show")
}

func (o *GroupObject) Delete(r *http类.X请求) {
	r.X响应.X写响应缓冲区("Object Delete")
}

func Handler(r *http类.X请求) {
	r.X响应.X写响应缓冲区("Handler")
}

func Test_Router_GroupBasic1(t *testing.T) {
	s := g.Http类(uid类.X生成())
	obj := new(GroupObject)
	// 分组路由方法注册
	group := s.X创建分组路由("/api")
	group.X绑定所有类型("/handler", Handler)
	group.X绑定所有类型("/obj", obj)
	group.X绑定GET("/obj/my-show", obj, "Show")
	group.X绑定RESTfulAPI对象("/obj/rest", obj)
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/api/handler"), "Handler")

		t.Assert(client.Get文本(ctx, "/api/obj"), "1Object Index2")
		t.Assert(client.Get文本(ctx, "/api/obj/"), "1Object Index2")
		t.Assert(client.Get文本(ctx, "/api/obj/index"), "1Object Index2")
		t.Assert(client.Get文本(ctx, "/api/obj/delete"), "1Object Delete2")
		t.Assert(client.Get文本(ctx, "/api/obj/my-show"), "1Object Show2")
		t.Assert(client.Get文本(ctx, "/api/obj/show"), "1Object Show2")
		t.Assert(client.Delete文本(ctx, "/api/obj/rest"), "1Object Delete2")

		t.Assert(client.Delete文本(ctx, "/ThisDoesNotExist"), "Not Found")
		t.Assert(client.Delete文本(ctx, "/api/ThisDoesNotExist"), "Not Found")
	})
}

func Test_Router_GroupBuildInVar(t *testing.T) {
	s := g.Http类(uid类.X生成())
	obj := new(GroupObject)
	// 分组路由方法注册
	group := s.X创建分组路由("/api")
	group.X绑定所有类型("/{.struct}/{.method}", obj)
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/api/group-object/index"), "1Object Index2")
		t.Assert(client.Get文本(ctx, "/api/group-object/delete"), "1Object Delete2")
		t.Assert(client.Get文本(ctx, "/api/group-object/show"), "1Object Show2")

		t.Assert(client.Delete文本(ctx, "/ThisDoesNotExist"), "Not Found")
		t.Assert(client.Delete文本(ctx, "/api/ThisDoesNotExist"), "Not Found")
	})
}

func Test_Router_Group_Methods(t *testing.T) {
	s := g.Http类(uid类.X生成())
	obj := new(GroupObject)
	group := s.X创建分组路由("/")
	group.X绑定所有类型("/obj", obj, "Show, Delete")
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.Get文本(ctx, "/obj/show"), "1Object Show2")
		t.Assert(client.Get文本(ctx, "/obj/delete"), "1Object Delete2")
	})
}

func Test_Router_Group_MultiServer(t *testing.T) {
	s1 := g.Http类(uid类.X生成())
	s2 := g.Http类(uid类.X生成())
	s1.X创建分组路由("/", func(group *http类.X分组路由) {
		group.X绑定POST("/post", func(r *http类.X请求) {
			r.X响应.X写响应缓冲区("post1")
		})
	})
	s2.X创建分组路由("/", func(group *http类.X分组路由) {
		group.X绑定POST("/post", func(r *http类.X请求) {
			r.X响应.X写响应缓冲区("post2")
		})
	})
	s1.SetDumpRouterMap(false)
	s2.SetDumpRouterMap(false)
	单元测试类.Assert(s1.X开始监听(), nil)
	单元测试类.Assert(s2.X开始监听(), nil)
	defer s1.X关闭当前服务()
	defer s2.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c1 := g.X网页类()
		c1.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s1.X取已监听端口()))
		c2 := g.X网页类()
		c2.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s2.X取已监听端口()))
		t.Assert(c1.Post文本(ctx, "/post"), "post1")
		t.Assert(c2.Post文本(ctx, "/post"), "post2")
	})
}

func Test_Router_Group_Map(t *testing.T) {
	testFuncGet := func(r *http类.X请求) {
		r.X响应.X写响应缓冲区("get")
	}
	testFuncPost := func(r *http类.X请求) {
		r.X响应.X写响应缓冲区("post")
	}
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/", func(group *http类.X分组路由) {
		group.X绑定Map(map[string]interface{}{
			"Get: /test": testFuncGet,
			"Post:/test": testFuncPost,
		})
	})
	// s.SetDumpRouterMap(false) // 设置s（通常是一个HTTP服务器或路由器）不输出路由映射信息
	单元测试类.Assert(s.X开始监听(), nil)
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(c.Get文本(ctx, "/test"), "get")
		t.Assert(c.Post文本(ctx, "/test"), "post")
	})
}

type SafeBuffer struct {
	buffer *bytes.Buffer
	mu     sync.Mutex
}

func (b *SafeBuffer) Write(p []byte) (n int, err error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buffer.Write(p)
}

func (b *SafeBuffer) String() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buffer.String()
}

func Test_Router_OverWritten(t *testing.T) {
	var (
		s   = g.Http类(uid类.X生成())
		obj = new(GroupObject)
		buf = &SafeBuffer{
			buffer: bytes.NewBuffer(nil),
			mu:     sync.Mutex{},
		}
		logger = 日志类.X创建并按writer(buf)
	)
	logger.X设置是否同时输出到终端(false)
	s.X设置日志记录器(logger)
	s.X设置路由允许覆盖(true)
	s.X创建分组路由("/api", func(group *http类.X分组路由) {
		group.X绑定所有类型Map(g.Map{
			"/obj": obj,
		})
		group.X绑定所有类型Map(g.Map{
			"/obj": obj,
		})
	})
	s.X开始监听()
	defer s.X关闭当前服务()

	dumpContent := buf.String()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X统计次数(dumpContent, `/api/obj `), 1)
		t.Assert(文本类.X统计次数(dumpContent, `/api/obj/index`), 1)
		t.Assert(文本类.X统计次数(dumpContent, `/api/obj/show`), 1)
		t.Assert(文本类.X统计次数(dumpContent, `/api/obj/delete`), 1)
	})
}
