// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp_test
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
	s := g.Server(guid.S())
	s.BindHandler("/get", func(r *ghttp.Request) {
		r.Response.WriteExit(r.Get("id"), r.Get("name"))
	})
	s.BindHandler("/map", func(r *ghttp.Request) {
		if m := r.GetMap(); len(m) > 0 {
			r.Response.WriteExit(m["id"], m["name"], m["password1"], m["password2"])
		}
	})
	s.BindHandler("/parse", func(r *ghttp.Request) {
		if m := r.GetMap(); len(m) > 0 {
			var user *User
			if err := r.Parse(&user); err != nil {
				r.Response.WriteExit(err)
			}
			r.Response.WriteExit(user.Id, user.Name, user.Pass1, user.Pass2)
		}
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		content1 := `<doc><id>1</id><name>john</name><password1>123Abc!@#</password1><password2>123Abc!@#</password2></doc>`
		content2 := `<doc><id>1</id><name>john</name><password1>123Abc!@#</password1><password2>123</password2></doc>`
		t.Assert(client.GetContent(ctx, "/get", content1), ``)
		t.Assert(client.PostContent(ctx, "/get", content1), `1john`)
		t.Assert(client.GetContent(ctx, "/map", content1), ``)
		t.Assert(client.PostContent(ctx, "/map", content1), `1john123Abc!@#123Abc!@#`)
		t.Assert(client.PostContent(ctx, "/parse", content1), `1john123Abc!@#123Abc!@#`)
		t.Assert(client.PostContent(ctx, "/parse", content2), `密码强度不足`)
	})
}
