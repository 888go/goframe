// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"
	"time"
	
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/guid"
)

func Test_Params_Basic(t *testing.T) {
	type User struct {
		Id    int
		Name  string
		Pass1 string `p:"password1"`
		Pass2 string `p:"password2"`
	}
	s := g.Http类(uid类.X生成())
	// GET
	s.X绑定("/get", func(r *http类.Request) {
		if r.X取查询参数到泛型类("array") != nil {
			r.Response.X写响应缓冲区(r.X取查询参数到泛型类("array"))
		}
		if r.X取查询参数到泛型类("slice") != nil {
			r.Response.X写响应缓冲区(r.X取查询参数到泛型类("slice"))
		}
		if r.X取查询参数到泛型类("bool") != nil {
			r.Response.X写响应缓冲区(r.X取查询参数到泛型类("bool").X取布尔())
		}
		if r.X取查询参数到泛型类("float32") != nil {
			r.Response.X写响应缓冲区(r.X取查询参数到泛型类("float32").X取小数32位())
		}
		if r.X取查询参数到泛型类("float64") != nil {
			r.Response.X写响应缓冲区(r.X取查询参数到泛型类("float64").X取小数64位())
		}
		if r.X取查询参数到泛型类("int") != nil {
			r.Response.X写响应缓冲区(r.X取查询参数到泛型类("int").X取整数())
		}
		if r.X取查询参数到泛型类("uint") != nil {
			r.Response.X写响应缓冲区(r.X取查询参数到泛型类("uint").X取正整数())
		}
		if r.X取查询参数到泛型类("string") != nil {
			r.Response.X写响应缓冲区(r.X取查询参数到泛型类("string").String())
		}
		if r.X取查询参数到泛型类("map") != nil {
			r.Response.X写响应缓冲区(r.X取查询参数到Map()["map"].(map[string]interface{})["b"])
		}
		if r.X取查询参数到泛型类("a") != nil {
			r.Response.X写响应缓冲区(r.X取查询参数到MapStrStr()["a"])
		}
	})
	// PUT
	s.X绑定("/put", func(r *http类.Request) {
		if r.Get别名("array") != nil {
			r.Response.X写响应缓冲区(r.Get别名("array"))
		}
		if r.Get别名("slice") != nil {
			r.Response.X写响应缓冲区(r.Get别名("slice"))
		}
		if r.Get别名("bool") != nil {
			r.Response.X写响应缓冲区(r.Get别名("bool").X取布尔())
		}
		if r.Get别名("float32") != nil {
			r.Response.X写响应缓冲区(r.Get别名("float32").X取小数32位())
		}
		if r.Get别名("float64") != nil {
			r.Response.X写响应缓冲区(r.Get别名("float64").X取小数64位())
		}
		if r.Get别名("int") != nil {
			r.Response.X写响应缓冲区(r.Get别名("int").X取整数())
		}
		if r.Get别名("uint") != nil {
			r.Response.X写响应缓冲区(r.Get别名("uint").X取正整数())
		}
		if r.Get别名("string") != nil {
			r.Response.X写响应缓冲区(r.Get别名("string").String())
		}
		if r.Get别名("map") != nil {
			r.Response.X写响应缓冲区(r.GetMap别名()["map"].(map[string]interface{})["b"])
		}
		if r.Get别名("a") != nil {
			r.Response.X写响应缓冲区(r.GetMapStrStr别名()["a"])
		}
	})

	// DELETE
	s.X绑定("/delete", func(r *http类.Request) {
		if r.Get别名("array") != nil {
			r.Response.X写响应缓冲区(r.Get别名("array"))
		}
		if r.Get别名("slice") != nil {
			r.Response.X写响应缓冲区(r.Get别名("slice"))
		}
		if r.Get别名("bool") != nil {
			r.Response.X写响应缓冲区(r.Get别名("bool").X取布尔())
		}
		if r.Get别名("float32") != nil {
			r.Response.X写响应缓冲区(r.Get别名("float32").X取小数32位())
		}
		if r.Get别名("float64") != nil {
			r.Response.X写响应缓冲区(r.Get别名("float64").X取小数64位())
		}
		if r.Get别名("int") != nil {
			r.Response.X写响应缓冲区(r.Get别名("int").X取整数())
		}
		if r.Get别名("uint") != nil {
			r.Response.X写响应缓冲区(r.Get别名("uint").X取正整数())
		}
		if r.Get别名("string") != nil {
			r.Response.X写响应缓冲区(r.Get别名("string").String())
		}
		if r.Get别名("map") != nil {
			r.Response.X写响应缓冲区(r.GetMap别名()["map"].(map[string]interface{})["b"])
		}
		if r.Get别名("a") != nil {
			r.Response.X写响应缓冲区(r.GetMapStrStr别名()["a"])
		}
	})
	// PATCH
	s.X绑定("/patch", func(r *http类.Request) {
		if r.Get别名("array") != nil {
			r.Response.X写响应缓冲区(r.Get别名("array"))
		}
		if r.Get别名("slice") != nil {
			r.Response.X写响应缓冲区(r.Get别名("slice"))
		}
		if r.Get别名("bool") != nil {
			r.Response.X写响应缓冲区(r.Get别名("bool").X取布尔())
		}
		if r.Get别名("float32") != nil {
			r.Response.X写响应缓冲区(r.Get别名("float32").X取小数32位())
		}
		if r.Get别名("float64") != nil {
			r.Response.X写响应缓冲区(r.Get别名("float64").X取小数64位())
		}
		if r.Get别名("int") != nil {
			r.Response.X写响应缓冲区(r.Get别名("int").X取整数())
		}
		if r.Get别名("uint") != nil {
			r.Response.X写响应缓冲区(r.Get别名("uint").X取正整数())
		}
		if r.Get别名("string") != nil {
			r.Response.X写响应缓冲区(r.Get别名("string").String())
		}
		if r.Get别名("map") != nil {
			r.Response.X写响应缓冲区(r.GetMap别名()["map"].(map[string]interface{})["b"])
		}
		if r.Get别名("a") != nil {
			r.Response.X写响应缓冲区(r.GetMapStrStr别名()["a"])
		}
	})
	// Form
	s.X绑定("/form", func(r *http类.Request) {
		if r.Get别名("array") != nil {
			r.Response.X写响应缓冲区(r.X取表单值到泛型类("array"))
		}
		if r.Get别名("slice") != nil {
			r.Response.X写响应缓冲区(r.X取表单值到泛型类("slice"))
		}
		if r.Get别名("bool") != nil {
			r.Response.X写响应缓冲区(r.X取表单值到泛型类("bool").X取布尔())
		}
		if r.Get别名("float32") != nil {
			r.Response.X写响应缓冲区(r.X取表单值到泛型类("float32").X取小数32位())
		}
		if r.Get别名("float64") != nil {
			r.Response.X写响应缓冲区(r.X取表单值到泛型类("float64").X取小数64位())
		}
		if r.Get别名("int") != nil {
			r.Response.X写响应缓冲区(r.X取表单值到泛型类("int").X取整数())
		}
		if r.Get别名("uint") != nil {
			r.Response.X写响应缓冲区(r.X取表单值到泛型类("uint").X取正整数())
		}
		if r.Get别名("string") != nil {
			r.Response.X写响应缓冲区(r.X取表单值到泛型类("string").String())
		}
		if r.Get别名("map") != nil {
			r.Response.X写响应缓冲区(r.X取表单值到Map()["map"].(map[string]interface{})["b"])
		}
		if r.Get别名("a") != nil {
			r.Response.X写响应缓冲区(r.X取表单值到MapStrStr()["a"])
		}
	})
	s.X绑定("/map", func(r *http类.Request) {
		if m := r.X取查询参数到Map(); len(m) > 0 {
			r.Response.X写响应缓冲区(m["name"])
			return
		}
		if m := r.GetMap别名(); len(m) > 0 {
			r.Response.X写响应缓冲区(m["name"])
			return
		}
	})
	s.X绑定("/raw", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.X取请求体字节集())
	})
	s.X绑定("/json", func(r *http类.Request) {
		j, err := r.X取请求体到json类()
		if err != nil {
			r.Response.X写响应缓冲区(err)
			return
		}
		r.Response.X写响应缓冲区(j.X取值("name"))
	})
	s.X绑定("/struct", func(r *http类.Request) {
		if m := r.X取查询参数到Map(); len(m) > 0 {
			user := new(User)
			r.X取查询参数到结构体(user)
			r.Response.X写响应缓冲区(user.Id, user.Name, user.Pass1, user.Pass2)
			return
		}
		if m := r.GetMap别名(); len(m) > 0 {
			user := new(User)
			r.GetStruct别名(user)
			r.Response.X写响应缓冲区(user.Id, user.Name, user.Pass1, user.Pass2)
			return
		}
	})
	s.X绑定("/struct-with-nil", func(r *http类.Request) {
		user := (*User)(nil)
		err := r.GetStruct别名(&user)
		r.Response.X写响应缓冲区(err)
	})
	s.X绑定("/struct-with-base", func(r *http类.Request) {
		type Base struct {
			Pass1 string `p:"password1"`
			Pass2 string `p:"password2"`
		}
		type UserWithBase1 struct {
			Id   int
			Name string
			Base
		}
		type UserWithBase2 struct {
			Id   int
			Name string
			Pass Base
		}
		if m := r.GetMap别名(); len(m) > 0 {
			user1 := new(UserWithBase1)
			user2 := new(UserWithBase2)
			r.GetStruct别名(user1)
			r.GetStruct别名(user2)
			r.Response.X写响应缓冲区(user1.Id, user1.Name, user1.Pass1, user1.Pass2)
			r.Response.X写响应缓冲区(user2.Id, user2.Name, user2.Pass.Pass1, user2.Pass.Pass2)
		}
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		// GET
		t.Assert(client.Get文本(ctx, "/get", "array[]=1&array[]=2"), `["1","2"]`)
		t.Assert(client.Get文本(ctx, "/get", "slice=1&slice=2"), `2`)
		t.Assert(client.Get文本(ctx, "/get", "bool=1"), `true`)
		t.Assert(client.Get文本(ctx, "/get", "bool=0"), `false`)
		t.Assert(client.Get文本(ctx, "/get", "float32=0.11"), `0.11`)
		t.Assert(client.Get文本(ctx, "/get", "float64=0.22"), `0.22`)
		t.Assert(client.Get文本(ctx, "/get", "int=-10000"), `-10000`)
		t.Assert(client.Get文本(ctx, "/get", "int=10000"), `10000`)
		t.Assert(client.Get文本(ctx, "/get", "uint=10000"), `10000`)
		t.Assert(client.Get文本(ctx, "/get", "uint=9"), `9`)
		t.Assert(client.Get文本(ctx, "/get", "string=key"), `key`)
		t.Assert(client.Get文本(ctx, "/get", "map[a]=1&map[b]=2"), `2`)
		t.Assert(client.Get文本(ctx, "/get", "a=1&b=2"), `1`)

		// PUT
		t.Assert(client.Put文本(ctx, "/put", "array[]=1&array[]=2"), `["1","2"]`)
		t.Assert(client.Put文本(ctx, "/put", "slice=1&slice=2"), `2`)
		t.Assert(client.Put文本(ctx, "/put", "bool=1"), `true`)
		t.Assert(client.Put文本(ctx, "/put", "bool=0"), `false`)
		t.Assert(client.Put文本(ctx, "/put", "float32=0.11"), `0.11`)
		t.Assert(client.Put文本(ctx, "/put", "float64=0.22"), `0.22`)
		t.Assert(client.Put文本(ctx, "/put", "int=-10000"), `-10000`)
		t.Assert(client.Put文本(ctx, "/put", "int=10000"), `10000`)
		t.Assert(client.Put文本(ctx, "/put", "uint=10000"), `10000`)
		t.Assert(client.Put文本(ctx, "/put", "uint=9"), `9`)
		t.Assert(client.Put文本(ctx, "/put", "string=key"), `key`)
		t.Assert(client.Put文本(ctx, "/put", "map[a]=1&map[b]=2"), `2`)
		t.Assert(client.Put文本(ctx, "/put", "a=1&b=2"), `1`)

		// DELETE
		t.Assert(client.Delete文本(ctx, "/delete", "array[]=1&array[]=2"), `["1","2"]`)
		t.Assert(client.Delete文本(ctx, "/delete", "slice=1&slice=2"), `2`)
		t.Assert(client.Delete文本(ctx, "/delete", "bool=1"), `true`)
		t.Assert(client.Delete文本(ctx, "/delete", "bool=0"), `false`)
		t.Assert(client.Delete文本(ctx, "/delete", "float32=0.11"), `0.11`)
		t.Assert(client.Delete文本(ctx, "/delete", "float64=0.22"), `0.22`)
		t.Assert(client.Delete文本(ctx, "/delete", "int=-10000"), `-10000`)
		t.Assert(client.Delete文本(ctx, "/delete", "int=10000"), `10000`)
		t.Assert(client.Delete文本(ctx, "/delete", "uint=10000"), `10000`)
		t.Assert(client.Delete文本(ctx, "/delete", "uint=9"), `9`)
		t.Assert(client.Delete文本(ctx, "/delete", "string=key"), `key`)
		t.Assert(client.Delete文本(ctx, "/delete", "map[a]=1&map[b]=2"), `2`)
		t.Assert(client.Delete文本(ctx, "/delete", "a=1&b=2"), `1`)

		// PATCH
		t.Assert(client.Patch文本(ctx, "/patch", "array[]=1&array[]=2"), `["1","2"]`)
		t.Assert(client.Patch文本(ctx, "/patch", "slice=1&slice=2"), `2`)
		t.Assert(client.Patch文本(ctx, "/patch", "bool=1"), `true`)
		t.Assert(client.Patch文本(ctx, "/patch", "bool=0"), `false`)
		t.Assert(client.Patch文本(ctx, "/patch", "float32=0.11"), `0.11`)
		t.Assert(client.Patch文本(ctx, "/patch", "float64=0.22"), `0.22`)
		t.Assert(client.Patch文本(ctx, "/patch", "int=-10000"), `-10000`)
		t.Assert(client.Patch文本(ctx, "/patch", "int=10000"), `10000`)
		t.Assert(client.Patch文本(ctx, "/patch", "uint=10000"), `10000`)
		t.Assert(client.Patch文本(ctx, "/patch", "uint=9"), `9`)
		t.Assert(client.Patch文本(ctx, "/patch", "string=key"), `key`)
		t.Assert(client.Patch文本(ctx, "/patch", "map[a]=1&map[b]=2"), `2`)
		t.Assert(client.Patch文本(ctx, "/patch", "a=1&b=2"), `1`)

		// Form
		t.Assert(client.Post文本(ctx, "/form", "array[]=1&array[]=2"), `["1","2"]`)
		t.Assert(client.Post文本(ctx, "/form", "slice=1&slice=2"), `2`)
		t.Assert(client.Post文本(ctx, "/form", "bool=1"), `true`)
		t.Assert(client.Post文本(ctx, "/form", "bool=0"), `false`)
		t.Assert(client.Post文本(ctx, "/form", "float32=0.11"), `0.11`)
		t.Assert(client.Post文本(ctx, "/form", "float64=0.22"), `0.22`)
		t.Assert(client.Post文本(ctx, "/form", "int=-10000"), `-10000`)
		t.Assert(client.Post文本(ctx, "/form", "int=10000"), `10000`)
		t.Assert(client.Post文本(ctx, "/form", "uint=10000"), `10000`)
		t.Assert(client.Post文本(ctx, "/form", "uint=9"), `9`)
		t.Assert(client.Post文本(ctx, "/form", "string=key"), `key`)
		t.Assert(client.Post文本(ctx, "/form", "map[a]=1&map[b]=2"), `2`)
		t.Assert(client.Post文本(ctx, "/form", "a=1&b=2"), `1`)

		// Map
		t.Assert(client.Get文本(ctx, "/map", "id=1&name=john"), `john`)
		t.Assert(client.Post文本(ctx, "/map", "id=1&name=john"), `john`)

		// Raw
		t.Assert(client.Put文本(ctx, "/raw", "id=1&name=john"), `id=1&name=john`)

		// Json
		t.Assert(client.Post文本(ctx, "/json", `{"id":1, "name":"john"}`), `john`)

		// Struct
		t.Assert(client.Get文本(ctx, "/struct", `id=1&name=john&password1=123&password2=456`), `1john123456`)
		t.Assert(client.Post文本(ctx, "/struct", `id=1&name=john&password1=123&password2=456`), `1john123456`)
		t.Assert(client.Post文本(ctx, "/struct-with-nil", ``), ``)
		t.Assert(client.Post文本(ctx, "/struct-with-base", `id=1&name=john&password1=123&password2=456`), "1john1234561john")
	})
}

func Test_Params_Header(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/header", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.X取协议头值("test"))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		t.Assert(client.X协议头(g.MapStrStr{"test": "123456"}).Get文本(ctx, "/header"), "123456")
	})
}

func Test_Params_SupportChars(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/form-value", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.X取表单值到泛型类("test-value"))
	})
	s.X绑定("/form-array", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.X取表单值到泛型类("test-array"))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(c.Post文本(ctx, "/form-value", "test-value=100"), "100")
		t.Assert(c.Post文本(ctx, "/form-array", "test-array[]=1&test-array[]=2"), `["1","2"]`)
	})
}

func Test_Params_Priority(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/query", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.X取查询参数到泛型类("a"))
	})
	s.X绑定("/form", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.X取表单值到泛型类("a"))
	})
	s.X绑定("/request", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.Get别名("a"))
	})
	s.X绑定("/request-map", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.GetMap别名(g.Map{
			"a": 1,
			"b": 2,
		}))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		t.Assert(client.Get文本(ctx, "/query?a=1", "a=100"), "100")
		t.Assert(client.Post文本(ctx, "/form?a=1", "a=100"), "100")
		t.Assert(client.Put文本(ctx, "/form?a=1", "a=100"), "100")
		t.Assert(client.Get文本(ctx, "/request?a=1", "a=100"), "100")
		t.Assert(client.Get文本(ctx, "/request-map?a=1&b=2&c=3", "a=100&b=200&c=300"), `{"a":"100","b":"200"}`)
	})
}

func Test_Params_GetRequestMap(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/map", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.X取参数到Map())
	})
	s.X绑定("/withKVMap", func(r *http类.Request) {
		m := r.X取参数到Map(map[string]interface{}{"id": 2})
		r.Response.X写响应缓冲区(m["id"])
	})
	s.X绑定("/paramsMapWithKVMap", func(r *http类.Request) {
		r.X设置自定义参数("name", "john")
		m := r.X取参数到Map(map[string]interface{}{"id": 2})
		r.Response.X写响应缓冲区(m["id"])
	})
	s.X绑定("/{name}.map", func(r *http类.Request) {
		m := r.X取参数到Map(map[string]interface{}{"id": 2})
		r.Response.X写响应缓冲区(m["id"])
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		t.Assert(
			client.Post文本(ctx,
				"/map",
				"time_end2020-04-18 16:11:58&returnmsg=Success&attach=",
			),
			`{"attach":"","returnmsg":"Success"}`,
		)

		t.Assert(client.Post文本(ctx, "/john.map", "name=john"), 2)

		t.Assert(client.Post文本(ctx, "/withKVMap", "name=john"), 2)

		t.Assert(client.Post文本(ctx, "/paramsMapWithKVMap"), 2)

		client.X设置内容类型("application/json")
		t.Assert(client.Get文本(ctx, "/withKVMap", "name=john"), 2)
	})
}

func Test_Params_Modify(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/param/modify", func(r *http类.Request) {
		param := r.GetMap别名()
		param["id"] = 2
		paramBytes, err := json类.X变量到json字节集(param)
		if err != nil {
			r.Response.X写响应缓冲区(err)
			return
		}
		r.Request.Body = io.NopCloser(bytes.NewReader(paramBytes))
		r.X重载请求参数()
		r.Response.X写响应缓冲区(r.GetMap别名())
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		t.Assert(
			client.Post文本(ctx,
				"/param/modify",
				`{"id":1}`,
			),
			`{"id":2}`,
		)
	})
}

func Test_Params_Parse_DefaultValueTag(t *testing.T) {
	type T struct {
		Name  string  `d:"john"`
		Score float32 `d:"60"`
	}
	s := g.Http类(uid类.X生成())
	s.X绑定("/parse", func(r *http类.Request) {
		var t *T
		if err := r.X解析参数到结构(&t); err != nil {
			r.Response.X写响应缓冲区并退出(err)
		}
		r.Response.X写响应缓冲区并退出(t)
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		t.Assert(client.Post文本(ctx, "/parse"), `{"Name":"john","Score":60}`)
		t.Assert(client.Post文本(ctx, "/parse", `{"name":"smith"}`), `{"Name":"smith","Score":60}`)
		t.Assert(client.Post文本(ctx, "/parse", `{"name":"smith", "score":100}`), `{"Name":"smith","Score":100}`)
	})
}

func Test_Params_Parse_Validation(t *testing.T) {
	type RegisterReq struct {
		Name  string `p:"username"  v:"required|length:6,30#请输入账号|账号长度为{min}到{max}位"`
		Pass  string `p:"password1" v:"required|length:6,30#请输入密码|密码长度不够"`
		Pass2 string `p:"password2" v:"required|length:6,30|same:password1#请确认密码|密码长度不够|两次密码不一致"`
	}

	s := g.Http类(uid类.X生成())
	s.X绑定("/parse", func(r *http类.Request) {
		var req *RegisterReq
		if err := r.X解析参数到结构(&req); err != nil {
			r.Response.X写响应缓冲区(err)
		} else {
			r.Response.X写响应缓冲区("ok")
		}
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		t.Assert(client.Get文本(ctx, "/parse"), `请输入账号`)
		t.Assert(client.Get文本(ctx, "/parse?name=john11&password1=123456&password2=123"), `密码长度不够`)
		t.Assert(client.Get文本(ctx, "/parse?name=john&password1=123456&password2=123456"), `账号长度为6到30位`)
		t.Assert(client.Get文本(ctx, "/parse?name=john11&password1=123456&password2=123456"), `ok`)
	})
}

func Test_Params_Parse_EmbeddedWithAliasName1(t *testing.T) {
	// 获取内容列表
	type ContentGetListInput struct {
		Type       string
		CategoryId uint
		Page       int
		Size       int
		Sort       int
		UserId     uint
	}
	// 获取内容列表
	type ContentGetListReq struct {
		ContentGetListInput
		CategoryId uint `p:"cate"`
		Page       int  `d:"1"  v:"min:0#分页号码错误"`
		Size       int  `d:"10" v:"max:50#分页数量最大50条"`
	}

	s := g.Http类(uid类.X生成())
	s.X绑定("/parse", func(r *http类.Request) {
		var req *ContentGetListReq
		if err := r.X解析参数到结构(&req); err != nil {
			r.Response.X写响应缓冲区(err)
		} else {
			r.Response.X写响应缓冲区(req.ContentGetListInput)
		}
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		t.Assert(client.Get文本(ctx, "/parse?cate=1&page=2&size=10"), `{"Type":"","CategoryId":0,"Page":2,"Size":10,"Sort":0,"UserId":0}`)
	})
}

func Test_Params_Parse_EmbeddedWithAliasName2(t *testing.T) {
	// 获取内容列表
	type ContentGetListInput struct {
		Type       string
		CategoryId uint `p:"cate"`
		Page       int
		Size       int
		Sort       int
		UserId     uint
	}
	// 获取内容列表
	type ContentGetListReq struct {
		ContentGetListInput
		CategoryId uint `p:"cate"`
		Page       int  `d:"1"  v:"min:0#分页号码错误"`
		Size       int  `d:"10" v:"max:50#分页数量最大50条"`
	}

	s := g.Http类(uid类.X生成())
	s.X绑定("/parse", func(r *http类.Request) {
		var req *ContentGetListReq
		if err := r.X解析参数到结构(&req); err != nil {
			r.Response.X写响应缓冲区(err)
		} else {
			r.Response.X写响应缓冲区(req.ContentGetListInput)
		}
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		t.Assert(client.Get文本(ctx, "/parse?cate=1&page=2&size=10"), `{"Type":"","CategoryId":1,"Page":2,"Size":10,"Sort":0,"UserId":0}`)
	})
}

func Test_Params_GetParam(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.X取自定义参数到泛型类("key", "val"))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		t.Assert(client.Post文本(ctx, "/"), "val")
	})
}

func Test_Params_SetQuery(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/SetQuery", func(r *http类.Request) {
		r.X设置查询参数("a", 100)
		r.Response.X写响应缓冲区(r.X取查询参数到泛型类("a"))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		t.Assert(client.Get文本(ctx, "/SetQuery"), "100")
		t.Assert(client.Get文本(ctx, "/SetQuery?a=1"), "100")
	})
}

func Test_Params_GetQuery(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/GetQuery", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.X取查询参数到泛型类("a", 200))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		t.Assert(client.Get文本(ctx, "/GetQuery"), 200)
		t.Assert(client.X设置内容类型("application/json").Get文本(ctx, "/GetQuery", "a=100"), 100)
	})
}

func Test_Params_GetQueryMap(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/GetQueryMap", func(r *http类.Request) {
		if m := r.X取查询参数到Map(); len(m) > 0 {
			r.Response.X写响应缓冲区(m["name"])
		}
	})
	s.X绑定("/GetQueryMapWithKVMap", func(r *http类.Request) {
		if m := r.X取查询参数到Map(map[string]interface{}{"id": 1}); len(m) > 0 {
			r.Response.X写响应缓冲区(m["id"])
		}
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)

	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)
		client.X设置内容类型("application/json")
		t.Assert(client.Get文本(ctx, "/GetQueryMap", "id=1&name=john"), `john`)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)
		t.Assert(client.Get文本(ctx, "/GetQueryMapWithKVMap"), 1)
		t.Assert(client.Get文本(ctx, "/GetQueryMapWithKVMap", "name=john"), 1)
		t.Assert(client.Get文本(ctx, "/GetQueryMapWithKVMap", "id=2&name=john"), 2)
		client.X设置内容类型("application/json")
		t.Assert(client.Get文本(ctx, "/GetQueryMapWithKVMap", "name=john"), 1)
		t.Assert(client.Get文本(ctx, "/GetQueryMapWithKVMap", "id=2&name=john"), 2)
	})
}

func Test_Params_GetQueryMapStrStr(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/GetQueryMapStrStr", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.X取查询参数到MapStrStr())
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		t.Assert(client.Get文本(ctx, "/GetQueryMapStrStr"), "")
	})
}

func Test_Params_GetQueryMapStrVar(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/GetQueryMapStrVar", func(r *http类.Request) {
		m := r.X取查询参数到Map泛型类数组()
		r.Response.X写响应缓冲区(m["id"])
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		t.Assert(client.Get文本(ctx, "/GetQueryMapStrVar"), "")
		t.Assert(client.Get文本(ctx, "/GetQueryMapStrVar", "id=1"), 1)
	})
}

func Test_Params_GetRequest(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/GetRequest", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.X取参数("id"))
	})
	s.X绑定("/GetRequestWithDef", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.X取参数("id", 2))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		t.Assert(client.Get文本(ctx, "/GetRequestWithDef"), 2)

		client.X设置内容类型("application/json")
		t.Assert(client.Get文本(ctx, "/GetRequest", "id=1"), 1)
	})
}

func Test_Params_GetRequestMapStrStr(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/GetRequestMapStrStr", func(r *http类.Request) {
		r.Response.X写响应缓冲区(r.X取参数到MapStrStr())
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		t.Assert(client.Get文本(ctx, "/GetRequestMapStrStr"), "")
	})
}

func Test_Params_GetRequestMapStrVar(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X绑定("/GetRequestMapStrVar", func(r *http类.Request) {
		m := r.X取参数到Map泛型类()
		r.Response.X写响应缓冲区(m["id"])
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口())
		client := g.X网页类()
		client.X设置url前缀(prefix)

		t.Assert(client.Get文本(ctx, "/GetRequestMapStrVar"), "")
		t.Assert(client.Get文本(ctx, "/GetRequestMapStrVar", "id=1"), 1)
	})
}
