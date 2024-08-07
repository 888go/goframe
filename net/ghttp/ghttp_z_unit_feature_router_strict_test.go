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
	"testing"
	"time"

	gjson "github.com/888go/goframe/encoding/gjson"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/json"
	ghttp "github.com/888go/goframe/net/ghttp"
	gtest "github.com/888go/goframe/test/gtest"
	guid "github.com/888go/goframe/util/guid"
)

func Test_Router_Handler_Strict_WithObject(t *testing.T) {
	type TestReq struct {
		Age  int
		Name string
	}
	type TestRes struct {
		Id   int
		Age  int
		Name string
	}
	s := g.Http类(guid.X生成())
	s.Use别名(ghttp.MiddlewareHandlerResponse)
	s.X绑定("/test", func(ctx context.Context, req *TestReq) (res *TestRes, err error) {
		return &TestRes{
			Id:   1,
			Age:  req.Age,
			Name: req.Name,
		}, nil
	})
	s.X绑定("/test/error", func(ctx context.Context, req *TestReq) (res *TestRes, err error) {
		return &TestRes{
			Id:   1,
			Age:  req.Age,
			Name: req.Name,
		}, gerror.X创建("error")
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/test?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Age":18,"Name":"john"}}`)
		t.Assert(client.Get文本(ctx, "/test/error"), `{"code":50,"message":"error","data":{"Id":1,"Age":0,"Name":""}}`)
	})
}

type TestForHandlerWithObjectAndMeta1Req struct {
	g.Meta `path:"/custom-test1" method:"get"`
	Age    int
	Name   string
}

type TestForHandlerWithObjectAndMeta1Res struct {
	Id  int
	Age int
}

type TestForHandlerWithObjectAndMeta2Req struct {
	g.Meta `path:"/custom-test2" method:"get"`
	Age    int
	Name   string
}

type TestForHandlerWithObjectAndMeta2Res struct {
	Id   int
	Name string
}

type ControllerForHandlerWithObjectAndMeta1 struct{}

func (ControllerForHandlerWithObjectAndMeta1) Index(ctx context.Context, req *TestForHandlerWithObjectAndMeta1Req) (res *TestForHandlerWithObjectAndMeta1Res, err error) {
	return &TestForHandlerWithObjectAndMeta1Res{
		Id:  1,
		Age: req.Age,
	}, nil
}

func (ControllerForHandlerWithObjectAndMeta1) Test2(ctx context.Context, req *TestForHandlerWithObjectAndMeta2Req) (res *TestForHandlerWithObjectAndMeta2Res, err error) {
	return &TestForHandlerWithObjectAndMeta2Res{
		Id:   1,
		Name: req.Name,
	}, nil
}

type TestForHandlerWithObjectAndMeta3Req struct {
	g.Meta `path:"/custom-test3" method:"get"`
	Age    int
	Name   string
}

type TestForHandlerWithObjectAndMeta3Res struct {
	Id  int
	Age int
}

type TestForHandlerWithObjectAndMeta4Req struct {
	g.Meta `path:"/custom-test4" method:"get"`
	Age    int
	Name   string
}

type TestForHandlerWithObjectAndMeta4Res struct {
	Id   int
	Name string
}

type ControllerForHandlerWithObjectAndMeta2 struct{}

func (ControllerForHandlerWithObjectAndMeta2) Test3(ctx context.Context, req *TestForHandlerWithObjectAndMeta3Req) (res *TestForHandlerWithObjectAndMeta3Res, err error) {
	return &TestForHandlerWithObjectAndMeta3Res{
		Id:  1,
		Age: req.Age,
	}, nil
}

func (ControllerForHandlerWithObjectAndMeta2) Test4(ctx context.Context, req *TestForHandlerWithObjectAndMeta4Req) (res *TestForHandlerWithObjectAndMeta4Res, err error) {
	return &TestForHandlerWithObjectAndMeta4Res{
		Id:   1,
		Name: req.Name,
	}, nil
}

func Test_Router_Handler_Strict_WithObjectAndMeta(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.Use别名(ghttp.MiddlewareHandlerResponse)
	s.X创建分组路由("/", func(group *ghttp.X分组路由) {
		group.X绑定所有类型("/", new(ControllerForHandlerWithObjectAndMeta1))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), `{"code":65,"message":"Not Found","data":null}`)
		t.Assert(client.Get文本(ctx, "/custom-test1?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Age":18}}`)
		t.Assert(client.Get文本(ctx, "/custom-test2?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Name":"john"}}`)
		t.Assert(client.Post文本(ctx, "/custom-test2?age=18&name=john"), `{"code":65,"message":"Not Found","data":null}`)
	})
}

func Test_Router_Handler_Strict_Group_Bind(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.Use别名(ghttp.MiddlewareHandlerResponse)
	s.X创建分组路由("/api/v1", func(group *ghttp.X分组路由) {
		group.X绑定(
			new(ControllerForHandlerWithObjectAndMeta1),
			new(ControllerForHandlerWithObjectAndMeta2),
		)
	})
	s.X创建分组路由("/api/v2", func(group *ghttp.X分组路由) {
		group.X绑定(
			new(ControllerForHandlerWithObjectAndMeta1),
			new(ControllerForHandlerWithObjectAndMeta2),
		)
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/"), `{"code":65,"message":"Not Found","data":null}`)
		t.Assert(client.Get文本(ctx, "/api/v1/custom-test1?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Age":18}}`)
		t.Assert(client.Get文本(ctx, "/api/v1/custom-test2?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Name":"john"}}`)
		t.Assert(client.Post文本(ctx, "/api/v1/custom-test2?age=18&name=john"), `{"code":65,"message":"Not Found","data":null}`)

		t.Assert(client.Get文本(ctx, "/api/v1/custom-test3?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Age":18}}`)
		t.Assert(client.Get文本(ctx, "/api/v1/custom-test4?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Name":"john"}}`)

		t.Assert(client.Get文本(ctx, "/api/v2/custom-test1?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Age":18}}`)
		t.Assert(client.Get文本(ctx, "/api/v2/custom-test2?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Name":"john"}}`)
		t.Assert(client.Get文本(ctx, "/api/v2/custom-test3?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Age":18}}`)
		t.Assert(client.Get文本(ctx, "/api/v2/custom-test4?age=18&name=john"), `{"code":0,"message":"","data":{"Id":1,"Name":"john"}}`)
	})
}

func Test_Issue1708(t *testing.T) {
	type Test struct {
		Name string `json:"name"`
	}
	type Req struct {
		Page       int      `json:"page"       dc:"分页码"`
		Size       int      `json:"size"       dc:"分页数量"`
		TargetType string   `json:"targetType" v:"required#评论内容类型错误" dc:"评论类型: topic/ask/article/reply"`
		TargetId   uint     `json:"targetId"   v:"required#评论目标ID错误" dc:"对应内容ID"`
		Test       [][]Test `json:"test"`
	}
	type Res struct {
		Page       int      `json:"page"       dc:"分页码"`
		Size       int      `json:"size"       dc:"分页数量"`
		TargetType string   `json:"targetType" v:"required#评论内容类型错误" dc:"评论类型: topic/ask/article/reply"`
		TargetId   uint     `json:"targetId"   v:"required#评论目标ID错误" dc:"对应内容ID"`
		Test       [][]Test `json:"test"`
	}

	s := g.Http类(guid.X生成())
	s.Use别名(ghttp.MiddlewareHandlerResponse)
	s.X绑定("/test", func(ctx context.Context, req *Req) (res *Res, err error) {
		return &Res{
			Page:       req.Page,
			Size:       req.Size,
			TargetType: req.TargetType,
			TargetId:   req.TargetId,
			Test:       req.Test,
		}, nil
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		content := `
{
    "targetType":"topic",
    "targetId":10785,
    "test":[
        [
            {
                "name":"123"
            }
        ]
    ]
}
`
		t.Assert(
			client.Post文本(ctx, "/test", content),
			`{"code":0,"message":"","data":{"page":0,"size":0,"targetType":"topic","targetId":10785,"test":[[{"name":"123"}]]}}`,
		)
	})
}

func Test_Custom_Slice_Type_Attribute(t *testing.T) {
	type (
		WhiteListKey    string
		WhiteListValues []string
		WhiteList       map[WhiteListKey]WhiteListValues
	)
	type Req struct {
		Id   int
		List WhiteList
	}
	type Res struct {
		Content string
	}

	s := g.Http类(guid.X生成())
	s.Use别名(ghttp.MiddlewareHandlerResponse)
	s.X绑定("/test", func(ctx context.Context, req *Req) (res *Res, err error) {
		return &Res{
			Content: gjson.X变量到json文本PANI(req),
		}, nil
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		content := `
{
    "id":1,
	"list":{
		"key": ["a", "b", "c"]
	}
}
`
		t.Assert(
			client.Post文本(ctx, "/test", content),
			`{"code":0,"message":"","data":{"Content":"{\"Id\":1,\"List\":{\"key\":[\"a\",\"b\",\"c\"]}}"}}`,
		)
	})
}

func Test_Router_Handler_Strict_WithGeneric(t *testing.T) {
	type TestReq struct {
		Age int
	}
	type TestGeneric[T any] struct {
		Test T
	}
	type Test1Res struct {
		Age TestGeneric[int]
	}
	type Test2Res TestGeneric[int]
	type TestGenericRes[T any] struct {
		Test T
	}

	s := g.Http类(guid.X生成())
	s.Use别名(ghttp.MiddlewareHandlerResponse)
	s.X绑定("/test1", func(ctx context.Context, req *TestReq) (res *Test1Res, err error) {
		return &Test1Res{
			Age: TestGeneric[int]{
				Test: req.Age,
			},
		}, nil
	})
	s.X绑定("/test1_slice", func(ctx context.Context, req *TestReq) (res []Test1Res, err error) {
		return []Test1Res{
			Test1Res{
				Age: TestGeneric[int]{
					Test: req.Age,
				},
			},
		}, nil
	})
	s.X绑定("/test2", func(ctx context.Context, req *TestReq) (res *Test2Res, err error) {
		return &Test2Res{
			Test: req.Age,
		}, nil
	})

	s.X绑定("/test2_slice", func(ctx context.Context, req *TestReq) (res []Test2Res, err error) {
		return []Test2Res{
			Test2Res{
				Test: req.Age,
			},
		}, nil
	})

	s.X绑定("/test3", func(ctx context.Context, req *TestReq) (res *TestGenericRes[int], err error) {
		return &TestGenericRes[int]{
			Test: req.Age,
		}, nil
	})

	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	s.X绑定("/test3_slice", func(ctx context.Context, req *TestReq) (res []TestGenericRes[int], err error) {
		return []TestGenericRes[int]{
			TestGenericRes[int]{
				Test: req.Age,
			},
		}, nil
	})

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/test1?age=1"), `{"code":0,"message":"","data":{"Age":{"Test":1}}}`)
		t.Assert(client.Get文本(ctx, "/test1_slice?age=1"), `{"code":0,"message":"","data":[{"Age":{"Test":1}}]}`)
		t.Assert(client.Get文本(ctx, "/test2?age=2"), `{"code":0,"message":"","data":{"Test":2}}`)
		t.Assert(client.Get文本(ctx, "/test2_slice?age=2"), `{"code":0,"message":"","data":[{"Test":2}]}`)
		t.Assert(client.Get文本(ctx, "/test3?age=3"), `{"code":0,"message":"","data":{"Test":3}}`)
		t.Assert(client.Get文本(ctx, "/test3_slice?age=3"), `{"code":0,"message":"","data":[{"Test":3}]}`)
	})
}

type ParameterCaseSensitiveController struct{}

type ParameterCaseSensitiveControllerPathReq struct {
	g.Meta `path:"/api/*path" method:"post"`
	Path   string
}

type ParameterCaseSensitiveControllerPathRes struct {
	Path string
}

func (c *ParameterCaseSensitiveController) Path(
	ctx context.Context,
	req *ParameterCaseSensitiveControllerPathReq,
) (res *ParameterCaseSensitiveControllerPathRes, err error) {
	return &ParameterCaseSensitiveControllerPathRes{Path: req.Path}, nil
}

func Test_Router_Handler_Strict_ParameterCaseSensitive(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.Use别名(ghttp.MiddlewareHandlerResponse)
	s.X创建分组路由("/", func(group *ghttp.X分组路由) {
		group.X绑定(&ParameterCaseSensitiveController{})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		for i := 0; i < 1000; i++ {
			t.Assert(
				client.Post文本(ctx, "/api/111", `{"Path":"222"}`),
				`{"code":0,"message":"","data":{"Path":"222"}}`,
			)
		}
	})
}

type testJsonRawMessageIssue3449Req struct {
	g.Meta `path:"/test" method:"POST" sm:"hello" tags:"示例"`

	Name    string          `json:"name" v:"required" dc:"名称"`
	JSONRaw json.RawMessage `json:"jsonRaw" dc:"原始JSON"`
}
type testJsonRawMessageIssue3449Res struct {
	Name    string          `json:"name" v:"required" dc:"名称"`
	JSONRaw json.RawMessage `json:"jsonRaw" dc:"原始JSON"`
}

type testJsonRawMessageIssue3449 struct {
}

func (t *testJsonRawMessageIssue3449) Test(ctx context.Context, req *testJsonRawMessageIssue3449Req) (res *testJsonRawMessageIssue3449Res, err error) {
	return &testJsonRawMessageIssue3449Res{
		Name:    req.Name,
		JSONRaw: req.JSONRaw,
	}, nil
}

// 这段代码是链接到一个GitHub问题的注释，问题编号为3449。在Go语言中，这种注释用于提供有关代码的额外信息，比如引用外部资源、说明问题或者提供待解决的任务。由于内容是URL，它可能指向一个关于代码库gf的讨论、错误报告或改进请求。要查看具体的内容，需要访问该链接。 md5:641e41ce7485cc00
func Test_JsonRawMessage_Issue3449(t *testing.T) {

	s := g.Http类(guid.X生成())
	s.Use别名(ghttp.MiddlewareHandlerResponse)
	s.X创建分组路由("/", func(group *ghttp.X分组路由) {
		group.X绑定(new(testJsonRawMessageIssue3449))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		v1 := map[string]any{
			"jkey1": "11",
			"jkey2": "12",
		}

		v2 := map[string]any{
			"jkey1": "21",
			"jkey2": "22",
		}
		data := map[string]any{
			"Name": "test",
			"jsonRaw": []any{
				v1, v2,
			},
		}

		expect1 := `{"code":0,"message":"","data":{"name":"test","jsonRaw":[{"jkey1":"11","jkey2":"12"},{"jkey1":"21","jkey2":"22"}]}}`
		t.Assert(client.Post文本(ctx, "/test", data), expect1)

		expect2 := `{"code":0,"message":"","data":{"name":"test","jsonRaw":{"jkey1":"11","jkey2":"12"}}}`
		t.Assert(client.Post文本(ctx, "/test", map[string]any{
			"Name":    "test",
			"jsonRaw": v1,
		}), expect2)

	})
}

type testNullStringIssue3465Req struct {
	g.Meta `path:"/test" method:"get" sm:"hello" tags:"示例"`
	Name   []string `json:"name" v:"required"`
}
type testNullStringIssue3465Res struct {
	Name []string `json:"name" v:"required" `
}

type testNullStringIssue3465 struct {
}

func (t *testNullStringIssue3465) Test(ctx context.Context, req *testNullStringIssue3465Req) (res *testNullStringIssue3465Res, err error) {
	return &testNullStringIssue3465Res{
		Name: req.Name,
	}, nil
}

// 这段注释引用的是GitHub上的一个 issue，gf（Go Foundation）是一个用Go语言编写的开源框架。3465号issue可能是指该框架中的某个问题或讨论的编号。具体的内容需要查看相关链接才能了解详情。 md5:53810ebfb659d15e
func Test_NullString_Issue3465(t *testing.T) {

	s := g.Http类(guid.X生成())
	s.Use别名(ghttp.MiddlewareHandlerResponse)
	s.X创建分组路由("/", func(group *ghttp.X分组路由) {
		group.X绑定(new(testNullStringIssue3465))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		data1 := map[string]any{
			"name": "null",
		}

		expect1 := `{"code":0,"message":"","data":{"name":["null"]}}`
		t.Assert(client.Get文本(ctx, "/test", data1), expect1)

		data2 := map[string]any{
			"name": []string{"null", "null"},
		}
		expect2 := `{"code":0,"message":"","data":{"name":["null","null"]}}`
		t.Assert(client.Get文本(ctx, "/test", data2), expect2)

	})
}
