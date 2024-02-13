// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 网页类_test

import (
	"fmt"
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/guid"
)

func Test_Client_DoRequestObj(t *testing.T) {
	type UserCreateReq struct {
		g.Meta `path:"/user" method:"post"`
		Id     int
		Name   string
	}
	type UserCreateRes struct {
		Id int
	}
	type UserQueryReq struct {
		g.Meta `path:"/user/{id}" method:"get"`
		Id     int
	}
	type UserQueryRes struct {
		Id   int
		Name string
	}
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/user", func(group *http类.RouterGroup) {
		group.X绑定GET("/{id}", func(r *http类.Request) {
			r.Response.X写响应缓冲区JSON(g.Map{"id": r.Get别名("id").X取整数(), "name": "john"})
		})
		group.X绑定POST("/", func(r *http类.Request) {
			r.Response.X写响应缓冲区JSON(g.Map{"id": r.Get别名("Id")})
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		url := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类().X设置url前缀(url).X内容类型json()
		var (
			createRes *UserCreateRes
			createReq = UserCreateReq{
				Id:   1,
				Name: "john",
			}
		)
		err := client.DoRequestObj(ctx, createReq, &createRes)
		t.AssertNil(err)
		t.Assert(createRes.Id, 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		url := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类().X设置url前缀(url).X内容类型json()
		var (
			queryRes *UserQueryRes
			queryReq = UserQueryReq{
				Id: 1,
			}
		)
		err := client.DoRequestObj(ctx, queryReq, &queryRes)
		t.AssertNil(err)
		t.Assert(queryRes.Id, 1)
		t.Assert(queryRes.Name, "john")
	})
}
