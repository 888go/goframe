// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gclient_test

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
	s := g.Server(guid.S())
	s.Group("/user", func(group *ghttp.RouterGroup) {
		group.GET("/{id}", func(r *ghttp.Request) {
			r.Response.WriteJson(g.Map{"id": r.Get("id").Int(), "name": "john"})
		})
		group.POST("/", func(r *ghttp.Request) {
			r.Response.WriteJson(g.Map{"id": r.Get("Id")})
		})
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		url := fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort())
		client := g.Client().SetPrefix(url).ContentJson()
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
	gtest.C(t, func(t *gtest.T) {
		url := fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort())
		client := g.Client().SetPrefix(url).ContentJson()
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
