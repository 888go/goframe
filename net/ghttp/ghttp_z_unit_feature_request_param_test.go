package http类_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/guid"
)

type UserReq struct {
	g.Meta `path:"/user" tags:"User" method:"post" summary:"user api" title:"api title"`
	Id     int    `v:"required" d:"1"`
	Name   string `v:"required" in:"cookie"`
	Age    string `v:"required" in:"header"`
	// header: 头部信息
// query: 查询参数
// cookie: Cookie信息
// form: 表单数据
}

type UserRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

var (
	User = cUser{}
)

type cUser struct{}

func (c *cUser) User(ctx context.Context, req *UserReq) (res *UserRes, err error) {
	g.Http类上下文取请求对象(ctx).Response.X写响应缓冲区JSON(req)
	return
}

func Test_Params_Tag(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/", func(group *http类.RouterGroup) {
		group.X绑定中间件(http类.MiddlewareHandlerResponse)
		group.X绑定(User)
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)
		client.X设置cookie("name", "john")
		client.X设置协议头("age", "18")

		t.Assert(client.Post文本(ctx, "/user"), `{"Id":1,"Name":"john","Age":"18"}`)
		t.Assert(client.Post文本(ctx, "/user", "name=&age=&id="), `{"Id":1,"Name":"john","Age":"18"}`)
	})
}

func Benchmark_ParamTag(b *testing.B) {
	b.StopTimer()

	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/", func(group *http类.RouterGroup) {
		group.X绑定中间件(http类.MiddlewareHandlerResponse)
		group.X绑定(User)
	})
	s.SetDumpRouterMap(false)
	s.X设置日志开启访客记录(false)
	s.X设置日志开启错误记录(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
	client := g.X网页类()
	client.X设置url前缀(prefix)
	client.X设置cookie("name", "john")
	client.X设置协议头("age", "18")

	b.StartTimer()
	for i := 1; i < b.N; i++ {
		client.Post文本(ctx, "/user", "id="+strconv.Itoa(i))
	}
}
