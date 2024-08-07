// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/json"
	ghttp "github.com/888go/goframe/net/ghttp"
	gtest "github.com/888go/goframe/test/gtest"
	guid "github.com/888go/goframe/util/guid"
)

func Test_Params_Json_Request(t *testing.T) {
	type User struct {
		Id    int
		Name  string
		Time  *time.Time
		Pass1 string `p:"password1"`
		Pass2 string `p:"password2" v:"password2@required|length:2,20|password3|same:password1#||密码强度不足|两次密码不一致"`
	}
	s := g.Http类(guid.X生成())
	s.X绑定("/get", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区并退出(r.Get别名("id"), r.Get别名("name"))
	})
	s.X绑定("/map", func(r *ghttp.Request) {
		if m := r.GetMap别名(); len(m) > 0 {
			r.X响应.X写响应缓冲区并退出(m["id"], m["name"], m["password1"], m["password2"])
		}
	})
	s.X绑定("/parse", func(r *ghttp.Request) {
		if m := r.GetMap别名(); len(m) > 0 {
			var user *User
			if err := r.X解析参数到结构(&user); err != nil {
				r.X响应.X写响应缓冲区并退出(err)
			}
			r.X响应.X写响应缓冲区并退出(user.Id, user.Name, user.Pass1, user.Pass2)
		}
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/get", `{"id":1,"name":"john","password1":"123Abc!@#","password2":"123Abc!@#"}`), ``)
		t.Assert(client.Get文本(ctx, "/map", `{"id":1,"name":"john","password1":"123Abc!@#","password2":"123Abc!@#"}`), ``)
		t.Assert(client.Post文本(ctx, "/parse", `{"id":1,"name":"john","password1":"123Abc!@#","password2":"123Abc!@#"}`), `1john123Abc!@#123Abc!@#`)
		t.Assert(client.Post文本(ctx, "/parse", `{"id":1,"name":"john","password1":"123Abc!@#","password2":"123"}`), `密码强度不足`)
	})
}

func Test_Params_Json_Response(t *testing.T) {
	type User struct {
		Uid      int
		Name     string
		SiteUrl  string `json:"-"`
		NickName string `json:"nickname,omitempty"`
		Pass1    string `json:"password1"`
		Pass2    string `json:"password2"`
	}

	s := g.Http类(guid.X生成())
	s.X绑定("/json1", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区JSON(User{
			Uid:     100,
			Name:    "john",
			SiteUrl: "https://goframe.org",
			Pass1:   "123",
			Pass2:   "456",
		})
	})
	s.X绑定("/json2", func(r *ghttp.Request) {
		r.X响应.X写响应缓冲区JSON(&User{
			Uid:     100,
			Name:    "john",
			SiteUrl: "https://goframe.org",
			Pass1:   "123",
			Pass2:   "456",
		})
	})
	s.X绑定("/json3", func(r *ghttp.Request) {
		type Message struct {
			Code  int    `json:"code"`
			Body  string `json:"body,omitempty"`
			Error string `json:"error,omitempty"`
		}
		type ResponseJson struct {
			Success  bool        `json:"success"`
			Data     interface{} `json:"data,omitempty"`
			ExtData  interface{} `json:"ext_data,omitempty"`
			Paginate interface{} `json:"paginate,omitempty"`
			Message  Message     `json:"message,omitempty"`
		}
		responseJson := &ResponseJson{
			Success: true,
			Data:    nil,
			ExtData: nil,
			Message: Message{3, "测试", "error"},
		}
		r.X响应.X写响应缓冲区JSON(responseJson)
	})
	s.X绑定("/json4", func(r *ghttp.Request) {
		type Message struct {
			Code  int    `json:"code"`
			Body  string `json:"body,omitempty"`
			Error string `json:"error,omitempty"`
		}
		type ResponseJson struct {
			Success  bool        `json:"success"`
			Data     interface{} `json:"data,omitempty"`
			ExtData  interface{} `json:"ext_data,omitempty"`
			Paginate interface{} `json:"paginate,omitempty"`
			Message  *Message    `json:"message,omitempty"`
		}
		responseJson := ResponseJson{
			Success: true,
			Data:    nil,
			ExtData: nil,
			Message: &Message{3, "测试", "error"},
		}
		r.X响应.X写响应缓冲区JSON(responseJson)
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		map1 := make(map[string]interface{})
		err1 := json.UnmarshalUseNumber([]byte(client.Get文本(ctx, "/json1")), &map1)
		t.Assert(err1, nil)
		t.Assert(len(map1), 4)
		t.Assert(map1["Name"], "john")
		t.Assert(map1["Uid"], 100)
		t.Assert(map1["password1"], "123")
		t.Assert(map1["password2"], "456")

		map2 := make(map[string]interface{})
		err2 := json.UnmarshalUseNumber([]byte(client.Get文本(ctx, "/json2")), &map2)
		t.Assert(err2, nil)
		t.Assert(len(map2), 4)
		t.Assert(map2["Name"], "john")
		t.Assert(map2["Uid"], 100)
		t.Assert(map2["password1"], "123")
		t.Assert(map2["password2"], "456")

		map3 := make(map[string]interface{})
		err3 := json.UnmarshalUseNumber([]byte(client.Get文本(ctx, "/json3")), &map3)
		t.Assert(err3, nil)
		t.Assert(len(map3), 2)
		t.Assert(map3["success"], "true")
		t.Assert(map3["message"], g.Map{"body": "测试", "code": 3, "error": "error"})

		map4 := make(map[string]interface{})
		err4 := json.UnmarshalUseNumber([]byte(client.Get文本(ctx, "/json4")), &map4)
		t.Assert(err4, nil)
		t.Assert(len(map4), 2)
		t.Assert(map4["success"], "true")
		t.Assert(map4["message"], g.Map{"body": "测试", "code": 3, "error": "error"})
	})
}
