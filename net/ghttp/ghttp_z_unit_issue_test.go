// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类_test

import (
	"context"
	"fmt"
	"testing"
	"time"
	
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/encoding/gurl"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gtag"
	"github.com/888go/goframe/util/guid"
)

// 这是Go语言代码中的一行注释，其内容为一个URL链接地址，指向GitHub上gogf/gf项目的一个issue（问题）讨论页面，编号为1609。
// 中文翻译：
// 这是引用了GitHub上gogf/gf项目第1609号问题讨论页面的链接。
func Test_Issue1609(t *testing.T) {
	s := g.Http类(uid类.X生成())
	group := s.X创建分组路由("/api/get")
	group.X绑定GET("/", func(r *http类.X请求) {
		r.X响应.X写响应缓冲区("get")
	})
	s.SetDumpRouterMap(false)
	单元测试类.Assert(s.X开始监听(), nil)
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(c.Get文本(ctx, "/api/get"), "get")
		t.Assert(c.Post文本(ctx, "/test"), "Not Found")
	})
}

func Test_Issue1611(t *testing.T) {
	s := g.Http类(uid类.X生成())
	v := g.X模板类(uid类.X生成())
	content := "This is header"
	单元测试类.AssertNil(v.SetPath(单元测试类.DataPath("issue1611")))
	s.X设置默认模板对象(v)
	s.X绑定("/", func(r *http类.X请求) {
		单元测试类.AssertNil(r.X响应.X输出到模板文件("index/layout.html", g.Map{
			"header": content,
		}))
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(文本类.X是否包含(c.Get文本(ctx, "/"), content), true)
	})
}

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf项目的一个issue（问题）：#1626
// 中文翻译：
// 这是Go语言代码注释，它指向GitHub上gogf/gf项目中的一个问题编号1626。
func Test_Issue1626(t *testing.T) {
	type TestReq struct {
		Name string `v:"required"`
	}
	type TestRes struct {
		Name string
	}
	s := g.Http类(uid类.X生成())
	s.Use别名(
		http类.MiddlewareHandlerResponse,
		func(r *http类.X请求) {
			r.X中间件管理器.Next()
			if err := r.X取错误信息(); err != nil {
				r.X响应.X清空缓冲区()
				r.X响应.X写响应缓冲区(err.Error())
			}
		},
	)
	s.X绑定("/test", func(ctx context.Context, req *TestReq) (res *TestRes, err error) {
		return &TestRes{Name: req.Name}, nil
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(c.Get文本(ctx, "/test"), `The Name field is required`)
		t.Assert(
			文本类.X是否包含(c.Get文本(ctx, "/test?name=john"), `{"Name":"john"}`),
			true,
		)
	})
}

type Issue1653TestReq struct {
	g.Meta    `path:"/test" method:"post" summary:"执行报表查询" tags:""`
	UUID      string  `json:"uuid" v:"required#菜单唯一码不可为空" dc:""`
	Limit     int     `json:"limit"`
	Filter    []g.Map `json:"filter"`
	FilterMap g.Map   `json:"filter_map"`
}

type Issue1653TestRes struct {
	UUID     string      `json:"uuid"`
	FeedBack interface{} `json:"feed_back"`
}

type cIssue1653Foo struct{}

var Issue1653Foo = new(cIssue1653Foo)

func (r cIssue1653Foo) PostTest(ctx context.Context, req *Issue1653TestReq) (*Issue1653TestRes, error) {
	return &Issue1653TestRes{UUID: req.UUID, FeedBack: req.Filter[0]["code"]}, nil
}

func Test_Issue1653(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.Use别名(http类.MiddlewareHandlerResponse)
	s.X创建分组路由("/boot", func(grp *http类.X分组路由) {
		grp.X绑定(Issue1653Foo)
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(1000 * time.Millisecond)
	// g.Client()测试：
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		dataReq := `
{"uuid":"28ee701c-7daf-4cdc-9a62-6d6704e6112b","limit":0,"filter":
[
{
"code":"P00001","constraint":"",
"created_at":"2022-03-08 04:56:15","created_by":"3ed72aba-1622-4262-a61e-83581e020763","default_value":"MonthStart()",
"expression":"AND A.DLVDAT_0>='%v'","force":false,"frequent":true,"name":"发货日期起",
"parent":"13109602-0da3-49b9-827f-2f44183ab756","read_only":false,"reference":null,"type":"date",
"updated_at":"2022-03-08 04:56:15","updated_by":"3ed72aba-1622-4262-a61e-83581e020763","updated_tick":1,
"uuid":"e6cd3268-1d75-42e0-83f9-f1f7b29976e8"
},
{
"code":"P00002","constraint":"","created_at":"2022-03-08 04:56:15","created_by":
"3ed72aba-1622-4262-a61e-83581e020763","default_value":"MonthEnd()","expression":"AND A.DLVDAT_0<='%v'","force":false,"frequent":true,
"name":"发货日期止","parent":"13109602-0da3-49b9-827f-2f44183ab756","read_only":false,"reference":null,"type":"date","updated_at":
"2022-03-08 04:56:15","updated_by":"3ed72aba-1622-4262-a61e-83581e020763","updated_tick":1,"uuid":"dba005b5-655e-4ac4-8b22-898aa3ad2294"
}
],
"filter_map":{"P00001":1646064000000,"P00002":1648742399999},
"selector_template":""
}
`
		resContent := c.Post文本(ctx, "/boot/test", dataReq)
		t.Assert(resContent, `{"code":0,"message":"","data":{"uuid":"28ee701c-7daf-4cdc-9a62-6d6704e6112b","feed_back":"P00001"}}`)
	})
}

type LbseMasterHead struct {
	Code     string   `json:"code" v:"code@required|min-length:1#The code is required"`
	Active   bool     `json:"active"`
	Preset   bool     `json:"preset"`
	Superior string   `json:"superior"`
	Path     []string `json:"path"`
	Sort     int      `json:"sort"`
	Folder   bool     `json:"folder"`
	Test     string   `json:"test" v:"required"`
}

type Template struct {
	LbseMasterHead
	Datasource string `json:"datasource" v:"required|length:32,32#The datasource is required"`
	SQLText    string `json:"sql_text"`
}

type TemplateCreateReq struct {
	g.Meta `path:"/test" method:"post" summary:"Create template" tags:"Template"`
	Master Template `json:"master"`
}

type TemplateCreateRes struct{}

type cFoo1 struct{}

var Foo1 = new(cFoo1)

func (r cFoo1) PostTest1(ctx context.Context, req *TemplateCreateReq) (res *TemplateCreateRes, err error) {
	g.X调试输出(req)
	return
}

// 这是golang代码中的一行注释，其内容引用了GitHub上gogf/gf项目的一个问题：#1662
// 翻译为：
// 参考GitHub上gogf/gf项目的第1662号问题
func Test_Issue662(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.Use别名(http类.MiddlewareHandlerResponse)
	s.X创建分组路由("/boot", func(grp *http类.X分组路由) {
		grp.X绑定(Foo1)
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(1000 * time.Millisecond)

	// g.Client()测试：
	// code字段传入空字符串时，校验没有提示
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		dataReq := `
{"master":{"active":true,"code":"","created_at":"","created_by":"","created_by_text":"","datasource":"38b6f170-a584-43fc-8912-cc1e9bf1b1a9","description":"币种","folder":false,"path":"[\"XCUR\"]","preset":false,"sort":1000,"sql_text":"SELECT!!!!","superior":null,"updated_at":"","updated_by":"","updated_by_text":"","updated_tick":0,"uuid":""},"translation":[{"code":"zh_CN","text":"币种"},{"code":"en_US","text":"币种"}],"filters":null,"fields":[{"code":"F001","created_at":"2022-01-18 23:37:38","created_by":"3ed72aba-1622-4262-a61e-83581e020763","field":"value","hide":false,"min_width":120,"name":"value","parent":"296154bf-b718-4e8f-8b70-efb969b831ec","updated_at":"2022-01-18 23:37:38","updated_by":"3ed72aba-1622-4262-a61e-83581e020763","updated_tick":1,"uuid":"f2140b7a-044c-41c3-b70e-852e6160b21b"},{"code":"F002","created_at":"2022-01-18 23:37:38","created_by":"3ed72aba-1622-4262-a61e-83581e020763","field":"label","hide":false,"min_width":120,"name":"label","parent":"296154bf-b718-4e8f-8b70-efb969b831ec","updated_at":"2022-01-18 23:37:38","updated_by":"3ed72aba-1622-4262-a61e-83581e020763","updated_tick":1,"uuid":"2d3bba5d-308b-4dba-bcac-f093e6556eca"}],"limit":0}
`
		t.Assert(c.Post文本(ctx, "/boot/test", dataReq), `{"code":51,"message":"The code is required","data":null}`)
	})
}

type DemoReq struct {
	g.Meta `path:"/demo" method:"post"`
	Data   *json类.Json
}

type DemoRes struct {
	Content string
}

type Api struct{}

func (a *Api) Demo(ctx context.Context, req *DemoReq) (res *DemoRes, err error) {
	return &DemoRes{
		Content: req.Data.X取json文本PANI(),
	}, err
}

var api = Api{}

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf项目的一个问题链接：https://github.com/gogf/gf/issues/2172
// 翻译为：
// 该注释指向GitHub上gogf/gf项目中的第2172号议题。
func Test_Issue2172(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.Use别名(http类.MiddlewareHandlerResponse)
	s.X创建分组路由("/", func(group *http类.X分组路由) {
		group.X绑定(api)
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(1000 * time.Millisecond)

	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		dataReq := `{"data":{"asd":1}}`
		t.Assert(c.Post文本(ctx, "/demo", dataReq), `{"code":0,"message":"","data":{"Content":"{\"asd\":1}"}}`)
	})
}

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf项目的一个问题链接：https://github.com/gogf/gf/issues/2334
// 翻译成中文：
// 这里引用了GitHub上gogf/gf项目的一个问题（编号为2334）的链接：https://github.com/gogf/gf/issues/2334
func Test_Issue2334(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X设置静态文件根目录(单元测试类.DataPath("static1"))
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(1000 * time.Millisecond)
	单元测试类.C(t, func(t *单元测试类.T) {
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(c.Get文本(ctx, "/index.html"), "index")

		c.X设置协议头("If-Modified-Since", "Mon, 12 Dec 2040 05:53:35 GMT")
		res, _ := c.Get响应对象(ctx, "/index.html")
		t.Assert(res.StatusCode, 304)
	})
}

type CreateOrderReq struct {
	g.Meta  `path:"/order" tags:"订单" method:"put" summary:"创建订单"`
	Details []*OrderDetail `p:"detail" v:"required#请输入订单详情" dc:"订单详情"`
}

type OrderDetail struct {
	Name   string  `p:"name" v:"required#请输入物料名称" dc:"物料名称"`
	Sn     string  `p:"sn" v:"required#请输入客户编号" dc:"客户编号"`
	Images string  `p:"images" dc:"图片"`
	Desc   string  `p:"desc" dc:"备注"`
	Number int     `p:"number" v:"required#请输入数量" dc:"数量"`
	Price  float64 `p:"price" v:"required" dc:"单价"`
}

type CreateOrderRes struct{}
type OrderController struct{}

func (c *OrderController) CreateOrder(ctx context.Context, req *CreateOrderReq) (res *CreateOrderRes, err error) {
	return
}

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf仓库下的第2482号问题。
func Test_Issue2482(t *testing.T) {
	s := g.Http类(uid类.X生成())
	s.X创建分组路由("/api/v2", func(group *http类.X分组路由) {
		group.X绑定中间件(http类.MiddlewareHandlerResponse)
		group.X绑定(OrderController{})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()
	time.Sleep(1000 * time.Millisecond)

	c := g.X网页类()
	c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
	单元测试类.C(t, func(t *单元测试类.T) {
		content := `
{
    "detail": [
      {
        "images": "string",
        "desc": "string",
        "number": 0,
        "price": 0
      }
    ]
  }
`
		t.Assert(c.Put文本(ctx, "/api/v2/order", content), `{"code":51,"message":"请输入物料名称","data":null}`)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		content := `
{
    "detail": [
      {
        "images": "string",
        "desc": "string",
        "number": 0,
		"name": "string",
        "price": 0
      }
    ]
  }
`
		t.Assert(c.Put文本(ctx, "/api/v2/order", content), `{"code":51,"message":"请输入客户编号","data":null}`)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		content := `
{
    "detail": [
      {
        "images": "string",
        "desc": "string",
        "number": 0,
		"name": "string",
		"sn": "string",
        "price": 0
      }
    ]
  }
`
		t.Assert(c.Put文本(ctx, "/api/v2/order", content), `{"code":0,"message":"","data":null}`)
	})
}

type Issue2890Enum string

const (
	Issue2890EnumA Issue2890Enum = "a"
	Issue2890EnumB Issue2890Enum = "b"
)

type Issue2890Req struct {
	g.Meta `path:"/issue2890" method:"post"`
	Id     int
	Enums  Issue2890Enum `v:"required|enums"`
}

type Issue2890Res struct{}
type Issue2890Controller struct{}

func (c *Issue2890Controller) Post(ctx context.Context, req *Issue2890Req) (res *Issue2890Res, err error) {
	g.Http类上下文取请求对象(ctx).X响应.X写响应缓冲区(req.Enums)
	return
}

// 这是GitHub上gogf/gf仓库的第2890个issue链接
func Test_Issue2890(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		oldEnumsJson, err := gtag.GetGlobalEnums()
		t.AssertNil(err)
		defer t.AssertNil(gtag.SetGlobalEnums(oldEnumsJson))

		err = gtag.SetGlobalEnums(`{"github.com/888go/goframe/net/ghttp_test.Issue2890Enum": ["a","b"]}`)
		t.AssertNil(err)

		s := g.Http类(uid类.X生成())
		s.X创建分组路由("/api/v2", func(group *http类.X分组路由) {
			group.X绑定中间件(http类.MiddlewareHandlerResponse)
			group.X绑定(Issue2890Controller{})
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(1000 * time.Millisecond)

		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(
			c.Post文本(ctx, "/api/v2/issue2890", ``),
			`{"code":51,"message":"The Enums field is required","data":null}`,
		)
		t.Assert(
			c.Post文本(ctx, "/api/v2/issue2890", `{"Enums":"c"}`),
			"{\"code\":51,\"message\":\"The Enums value `c` should be in enums of: [\\\"a\\\",\\\"b\\\"]\",\"data\":null}",
		)
	})
}

// 这是Go语言代码中的一行注释，其内容为一个GitHub仓库的issues链接地址。
// 翻译为：
// 参考GitHub上gogf/gf项目的问题2963
func Test_Issue2963(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		s.X设置静态文件根目录(单元测试类.DataPath("issue2963"))
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)

		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(c.Get文本(ctx, "/1.txt"), `1`)
		t.Assert(c.Get文本(ctx, "/中文G146(1)-icon.txt"), `中文G146(1)-icon`)
		t.Assert(c.Get文本(ctx, "/"+url类.X编码("中文G146(1)-icon.txt")), `中文G146(1)-icon`)
	})
}

type Issue3077Req struct {
	g.Meta `path:"/echo" method:"get"`
	A      string `default:"a"`
	B      string `default:""`
}
type Issue3077Res struct {
	g.Meta `mime:"text/html"`
}

type Issue3077V1 struct{}

func (c *Issue3077V1) Hello(ctx context.Context, req *Issue3077Req) (res *Issue3077Res, err error) {
	g.Http类上下文取请求对象(ctx).X响应.X写响应缓冲区(fmt.Sprintf("%v", req))
	return
}

// 这是Go语言代码中的一行注释，其内容引用了GitHub上gogf/gf项目的一个问题编号3077。 
// 翻译成中文：
// 引用了GitHub上gogf/gf项目的问题 #3077
func Test_Issue3077(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		s.X创建分组路由("/", func(group *http类.X分组路由) {
			group.X绑定(Issue3077V1{})
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)

		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))
		t.Assert(c.Get文本(ctx, "/echo?a=1&b=2"), `&{{} 1 2}`)
		t.Assert(c.Get文本(ctx, "/echo?"), `&{{} a }`)
	})
}
