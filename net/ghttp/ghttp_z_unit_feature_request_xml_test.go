// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类_test

import (
	"fmt"
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/guid"
)

func Test_Params_Xml_Request(t *testing.T) {
	type User struct {
		Id    int
		Name  string
		Time  *time.Time
		Pass1 string `p:"password1"`
		Pass2 string `p:"password2" v:"password2@required|length:2,20|password3|same:password1#||密码强度不足|两次密码不一致"`
	}
	s := g.Http类(uid类.X生成())
	s.X绑定("/get", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区并退出(r.Get别名("id"), r.Get别名("name"))
	})
	s.X绑定("/map", func(r *http类.X请求) {
		if m := r.GetMap别名(); len(m) > 0 {
			r.X响应.X写响应缓冲区并退出(m["id"], m["name"], m["password1"], m["password2"])
		}
	})
	s.X绑定("/parse", func(r *http类.X请求) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		content1 := `<doc><id>1</id><name>john</name><password1>123Abc!@#</password1><password2>123Abc!@#</password2></doc>`
		content2 := `<doc><id>1</id><name>john</name><password1>123Abc!@#</password1><password2>123</password2></doc>`
		t.Assert(client.Get文本(ctx, "/get", content1), ``)
		t.Assert(client.Post文本(ctx, "/get", content1), `1john`)
		t.Assert(client.Get文本(ctx, "/map", content1), ``)
		t.Assert(client.Post文本(ctx, "/map", content1), `1john123Abc!@#123Abc!@#`)
		t.Assert(client.Post文本(ctx, "/parse", content1), `1john123Abc!@#123Abc!@#`)
		t.Assert(client.Post文本(ctx, "/parse", content2), `密码强度不足`)
	})
}
