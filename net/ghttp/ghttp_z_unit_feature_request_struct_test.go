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
	"github.com/888go/goframe/util/gvalid"
)

func Test_Params_Parse(t *testing.T) {
	type User struct {
		Id   int
		Name string
		Map  map[string]interface{}
	}
	s := g.Http类(uid类.X生成())
	s.X绑定("/parse", func(r *http类.X请求) {
		var user *User
		if err := r.X解析参数到结构(&user); err != nil {
			r.X响应.X写响应缓冲区并退出(err)
		}
		r.X响应.X写响应缓冲区并退出(user.Map["id"], user.Map["score"])
	})
	s.X绑定("/parseErr", func(r *http类.X请求) {
		var user User
		err := r.X解析参数到结构(user)
		r.X响应.X写响应缓冲区并退出(err != nil)
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.Post文本(ctx, "/parse", `{"id":1,"name":"john","map":{"id":1,"score":100}}`), `1100`)
		t.Assert(client.Post文本(ctx, "/parseErr", `{"id":1,"name":"john","map":{"id":1,"score":100}}`), true)
	})
}

func Test_Params_ParseQuery(t *testing.T) {
	type User struct {
		Id   int
		Name string
	}
	s := g.Http类(uid类.X生成())
	s.X绑定("/parse-query", func(r *http类.X请求) {
		var user *User
		if err := r.X解析URL到结构(&user); err != nil {
			r.X响应.X写响应缓冲区并退出(err)
		}
		r.X响应.X写响应缓冲区并退出(user.Id, user.Name)
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(c.Get文本(ctx, "/parse-query"), `0`)
		t.Assert(c.Get文本(ctx, "/parse-query?id=1&name=john"), `1john`)
		t.Assert(c.Post文本(ctx, "/parse-query"), `0`)
		t.Assert(c.Post文本(ctx, "/parse-query", g.Map{
			"id":   1,
			"name": "john",
		}), `0`)
	})
}

func Test_Params_ParseForm(t *testing.T) {
	type User struct {
		Id   int
		Name string
	}
	s := g.Http类(uid类.X生成())
	s.X绑定("/parse-form", func(r *http类.X请求) {
		var user *User
		if err := r.X解析表单到结构(&user); err != nil {
			r.X响应.X写响应缓冲区并退出(err)
		}
		r.X响应.X写响应缓冲区并退出(user.Id, user.Name)
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(c.Get文本(ctx, "/parse-form"), `0`)
		t.Assert(c.Get文本(ctx, "/parse-form", g.Map{
			"id":   1,
			"name": "john",
		}), 0)
		t.Assert(c.Post文本(ctx, "/parse-form"), `0`)
		t.Assert(c.Post文本(ctx, "/parse-form", g.Map{
			"id":   1,
			"name": "john",
		}), `1john`)
	})
}

func Test_Params_ComplexJsonStruct(t *testing.T) {
	type ItemEnv struct {
		Type  string
		Key   string
		Value string
		Brief string
	}

	type ItemProbe struct {
		Type           string
		Port           int
		Path           string
		Brief          string
		Period         int
		InitialDelay   int
		TimeoutSeconds int
	}

	type ItemKV struct {
		Key   string
		Value string
	}

	type ItemPort struct {
		Port  int
		Type  string
		Alias string
		Brief string
	}

	type ItemMount struct {
		Type    string
		DstPath string
		Src     string
		SrcPath string
		Brief   string
	}

	type SaveRequest struct {
		AppId          uint
		Name           string
		Type           string
		Cluster        string
		Replicas       uint
		ContainerName  string
		ContainerImage string
		VersionTag     string
		Namespace      string
		Id             uint
		Status         uint
		Metrics        string
		InitImage      string
		CpuRequest     uint
		CpuLimit       uint
		MemRequest     uint
		MemLimit       uint
		MeshEnabled    uint
		ContainerPorts []ItemPort
		Labels         []ItemKV
		NodeSelector   []ItemKV
		EnvReserve     []ItemKV
		EnvGlobal      []ItemEnv
		EnvContainer   []ItemEnv
		Mounts         []ItemMount
		LivenessProbe  ItemProbe
		ReadinessProbe ItemProbe
	}

	s := g.Http类(uid类.X生成())
	s.X绑定("/parse", func(r *http类.X请求) {
		if m := r.GetMap别名(); len(m) > 0 {
			var data *SaveRequest
			if err := r.X解析参数到结构(&data); err != nil {
				r.X响应.X写响应缓冲区并退出(err)
			}
			r.X响应.X写响应缓冲区并退出(data)
		}
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		content := `
{
    "app_id": 5,
    "cluster": "test",
    "container_image": "nginx",
    "container_name": "test",
    "container_ports": [
        {
            "alias": "别名",
            "brief": "描述",
            "port": 80,
            "type": "tcp"
        }
    ],
    "cpu_limit": 100,
    "cpu_request": 10,
    "create_at": "2020-10-10 12:00:00",
    "creator": 1,
    "env_container": [
        {
            "brief": "用户环境变量",
            "key": "NAME",
            "type": "string",
            "value": "john"
        }
    ],
    "env_global": [
        {
            "brief": "数据数量",
            "key": "NUMBER",
            "type": "string",
            "value": "1"
        }
    ],
    "env_reserve": [
        {
            "key": "NODE_IP",
            "value": "status.hostIP"
        }
    ],
    "liveness_probe": {
        "brief": "存活探针",
        "initial_delay": 10,
        "path": "",
        "period": 5,
        "port": 80,
        "type": "tcpSocket"
    },
    "readiness_probe": {
        "brief": "就绪探针",
        "initial_delay": 10,
        "path": "",
        "period": 5,
        "port": 80,
        "type": "tcpSocket"
    },
    "id": 0,
    "init_image": "",
    "labels": [
        {
            "key": "app",
            "value": "test"
        }
    ],
    "mem_limit": 1000,
    "mem_request": 100,
    "mesh_enabled": 0,
    "metrics": "",
    "mounts": [],
    "name": "test",
    "namespace": "test",
    "node_selector": [
        {
            "key": "group",
            "value": "app"
        }
    ],
    "replicas": 1,
    "type": "test",
    "update_at": "2020-10-10 12:00:00",
    "version_tag": "test"
}
`
		t.Assert(client.Post文本(ctx, "/parse", content), `{"AppId":5,"Name":"test","Type":"test","Cluster":"test","Replicas":1,"ContainerName":"test","ContainerImage":"nginx","VersionTag":"test","Namespace":"test","Id":0,"Status":0,"Metrics":"","InitImage":"","CpuRequest":10,"CpuLimit":100,"MemRequest":100,"MemLimit":1000,"MeshEnabled":0,"ContainerPorts":[{"Port":80,"Type":"tcp","Alias":"别名","Brief":"描述"}],"Labels":[{"Key":"app","Value":"test"}],"NodeSelector":[{"Key":"group","Value":"app"}],"EnvReserve":[{"Key":"NODE_IP","Value":"status.hostIP"}],"EnvGlobal":[{"Type":"string","Key":"NUMBER","Value":"1","Brief":"数据数量"}],"EnvContainer":[{"Type":"string","Key":"NAME","Value":"john","Brief":"用户环境变量"}],"Mounts":[],"LivenessProbe":{"Type":"tcpSocket","Port":80,"Path":"","Brief":"存活探针","Period":5,"InitialDelay":10,"TimeoutSeconds":0},"ReadinessProbe":{"Type":"tcpSocket","Port":80,"Path":"","Brief":"就绪探针","Period":5,"InitialDelay":10,"TimeoutSeconds":0}}`)
	})
}

func Test_Params_Parse_Attr_Pointer1(t *testing.T) {
	type User struct {
		Id   *int
		Name *string
	}
	s := g.Http类(uid类.X生成())
	s.X绑定("/parse1", func(r *http类.X请求) {
		if m := r.GetMap别名(); len(m) > 0 {
			var user *User
			if err := r.X解析参数到结构(&user); err != nil {
				r.X响应.X写响应缓冲区并退出(err)
			}
			r.X响应.X写响应缓冲区并退出(user.Id, user.Name)
		}
	})
	s.X绑定("/parse2", func(r *http类.X请求) {
		if m := r.GetMap别名(); len(m) > 0 {
			var user = new(User)
			if err := r.X解析参数到结构(user); err != nil {
				r.X响应.X写响应缓冲区并退出(err)
			}
			r.X响应.X写响应缓冲区并退出(user.Id, user.Name)
		}
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.Post文本(ctx, "/parse1", `{"id":1,"name":"john"}`), `1john`)
		t.Assert(client.Post文本(ctx, "/parse2", `{"id":1,"name":"john"}`), `1john`)
		t.Assert(client.Post文本(ctx, "/parse2?id=1&name=john"), `1john`)
		t.Assert(client.Post文本(ctx, "/parse2", `id=1&name=john`), `1john`)
	})
}

func Test_Params_Parse_Attr_Pointer2(t *testing.T) {
	type User struct {
		Id *int `v:"required"`
	}
	s := g.Http类(uid类.X生成())
	s.X绑定("/parse", func(r *http类.X请求) {
		var user *User
		if err := r.X解析参数到结构(&user); err != nil {
			r.X响应.X写响应缓冲区并退出(err.Error())
		}
		r.X响应.X写响应缓冲区并退出(user.Id)
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.Post文本(ctx, "/parse"), `The Id field is required`)
		t.Assert(client.Post文本(ctx, "/parse?id=1"), `1`)
	})
}

// 目前还不支持这种类型的转换。
// func Test_Params_Parse_Attr_SliceSlice(t *testing.T) {
// 定义一个User结构体，包含Id（整型）、Name（字符串）和Scores（二维整数数组）属性
// }
//// 创建gServer实例（注释已被省略）
// s := g.Server(guid.S())
// 为"/parse"路径绑定处理器函数
// s.BindHandler("/parse", func(r *ghttp.Request) {
// 获取请求参数映射
// 如果参数映射不为空
// 定义一个指向User类型的指针变量user
// 尝试解析请求参数到user变量中，如果解析出错
// 则向响应写入错误信息并结束处理
// 解析成功则向响应写入user的Scores属性值并结束处理
// })
//// 设置不输出路由映射信息（注释已被省略）
// s.SetDumpRouterMap(false)
// 启动服务器
// 在程序结束时关闭服务器
// s.Start()
// defer s.Shutdown()
// 暂停100毫秒
// time.Sleep(100 * time.Millisecond)
// 使用gtest进行测试
// gtest.C(t, func(t *gtest.T) {
// 创建gClient实例
// 设置客户端请求前缀为服务器地址与端口
// 发送POST请求到"/parse"，内容为指定的JSON格式数据
// 断言返回的内容应为"1100"
// })
// }

func Test_Params_Struct(t *testing.T) {
	type User struct {
		Id    int
		Name  string
		Time  *time.Time
		Pass1 string `p:"password1"`
		Pass2 string `p:"password2" v:"password2 @required|length:2,20|password3#||密码强度不足"`
	}
	s := g.Http类(uid类.X生成())
	s.X绑定("/struct1", func(r *http类.X请求) {
		if m := r.GetMap别名(); len(m) > 0 {
			user := new(User)
			if err := r.GetStruct别名(user); err != nil {
				r.X响应.X写响应缓冲区并退出(err)
			}
			r.X响应.X写响应缓冲区并退出(user.Id, user.Name, user.Pass1, user.Pass2)
		}
	})
	s.X绑定("/struct2", func(r *http类.X请求) {
		if m := r.GetMap别名(); len(m) > 0 {
			user := (*User)(nil)
			if err := r.GetStruct别名(&user); err != nil {
				r.X响应.X写响应缓冲区并退出(err)
			}
			if user != nil {
				r.X响应.X写响应缓冲区并退出(user.Id, user.Name, user.Pass1, user.Pass2)
			}
		}
	})
	s.X绑定("/struct-valid", func(r *http类.X请求) {
		if m := r.GetMap别名(); len(m) > 0 {
			user := new(User)
			if err := r.GetStruct别名(user); err != nil {
				r.X响应.X写响应缓冲区并退出(err)
			}
			if err := 效验类.New().Data(user).Run(r.Context别名()); err != nil {
				r.X响应.X写响应缓冲区并退出(err)
			}
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
		t.Assert(client.Get文本(ctx, "/struct1", `id=1&name=john&password1=123&password2=456`), `1john123456`)
		t.Assert(client.Post文本(ctx, "/struct1", `id=1&name=john&password1=123&password2=456`), `1john123456`)
		t.Assert(client.Post文本(ctx, "/struct2", `id=1&name=john&password1=123&password2=456`), `1john123456`)
		t.Assert(client.Post文本(ctx, "/struct2", ``), ``)
		t.Assert(client.Post文本(ctx, "/struct-valid", `id=1&name=john&password1=123&password2=0`), "The password2 value `0` length must be between 2 and 20; 密码强度不足")
		t.Assert(client.Post文本(ctx, "/parse", `id=1&name=john&password1=123&password2=0`), "The password2 value `0` length must be between 2 and 20")
		t.Assert(client.Post文本(ctx, "/parse", `{"id":1,"name":"john","password1":"123Abc!@#","password2":"123Abc!@#"}`), `1john123Abc!@#123Abc!@#`)
	})
}

func Test_Params_Structs(t *testing.T) {
	type User struct {
		Id    int
		Name  string
		Time  *time.Time
		Pass1 string `p:"password1"`
		Pass2 string `p:"password2" v:"password2 @required|length:2,20|password3#||密码强度不足"`
	}
	s := g.Http类(uid类.X生成())
	s.X绑定("/parse1", func(r *http类.X请求) {
		var users []*User
		if err := r.X解析参数到结构(&users); err != nil {
			r.X响应.X写响应缓冲区并退出(err)
		}
		r.X响应.X写响应缓冲区并退出(users[0].Id, users[1].Id)
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(client.Post文本(ctx,
			"/parse1",
			`[{"id":1,"name":"john","password1":"123Abc!@#","password2":"123Abc!@#"}, {"id":2,"name":"john","password1":"123Abc!@#","password2":"123Abc!@#"}]`),
			`12`,
		)
	})
}

func Test_Params_Struct_Validation(t *testing.T) {
	type User struct {
		Id   int    `v:"required"`
		Name string `v:"name@required-with:id"`
	}
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/", func(group *http类.X分组路由) {
		group.X绑定所有类型("/", func(r *http类.X请求) {
			var (
				err  error
				user *User
			)
			err = r.X解析参数到结构(&user)
			if err != nil {
				r.X响应.X写响应缓冲区并退出(err)
			}
			r.X响应.X写响应缓冲区并退出(user.Id, user.Name)
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(c.Get文本(ctx, "/", ``), `The Id field is required`)
		t.Assert(c.Get文本(ctx, "/", `id=1&name=john`), `1john`)
		t.Assert(c.Post文本(ctx, "/", `id=1&name=john&password1=123&password2=456`), `1john`)
		t.Assert(c.Post文本(ctx, "/", `id=1`), `The name field is required`)
	})
}

// 这是GitHub上gogf/gf仓库中第1488号问题的链接
func Test_Params_Parse_Issue1488(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/", func(group *http类.X分组路由) {
		group.X绑定所有类型("/", func(r *http类.X请求) {
			type Request struct {
				Type         []int  `p:"type"`
				Keyword      string `p:"keyword"`
				Limit        int    `p:"per_page" d:"10"`
				Page         int    `p:"page" d:"1"`
				Order        string
				CreatedAtLte string
				CreatedAtGte string
				CreatorID    []int
			}
			for i := 0; i < 10; i++ {
				r.X设置自定义参数Map(g.Map{
					"type[]":           0,
					"keyword":          "",
					"t_start":          "",
					"t_end":            "",
					"reserve_at_start": "",
					"reserve_at_end":   "",
					"user_name":        "",
					"flag":             "",
					"per_page":         6,
				})
				var parsed Request
				_ = r.X解析参数到结构(&parsed)
				r.X响应.X写响应缓冲区(parsed.Page, parsed.Limit)
			}
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(c.Get文本(ctx, "/", ``), `16161616161616161616`)
	})
}
