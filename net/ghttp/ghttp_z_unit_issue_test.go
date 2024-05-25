// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gtag"
	"github.com/gogf/gf/v2/util/guid"
)

// https://github.com/gogf/gf/issues/1609
// 
// 这段注释引用的是一个GitHub问题的链接，来自 "gf"（GoGF）项目，问题编号为1609。这可能是一个关于gf库的问题报告、讨论或者请求，具体内容需要查看链接才能得知。. md5:467f0f77d2ecbf4b
func Test_Issue1609(t *testing.T) {
	s := g.Server(guid.S())
	group := s.Group("/api/get")
	group.GET("/", func(r *ghttp.Request) {
		r.Response.Write("get")
	})
	s.SetDumpRouterMap(false)
	gtest.Assert(s.Start(), nil)
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(c.GetContent(ctx, "/api/get"), "get")
		t.Assert(c.PostContent(ctx, "/test"), "Not Found")
	})
}

func Test_Issue1611(t *testing.T) {
	s := g.Server(guid.S())
	v := g.View(guid.S())
	content := "This is header"
	gtest.AssertNil(v.SetPath(gtest.DataPath("issue1611")))
	s.SetView(v)
	s.BindHandler("/", func(r *ghttp.Request) {
		gtest.AssertNil(r.Response.WriteTpl("index/layout.html", g.Map{
			"header": content,
		}))
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(gstr.Contains(c.GetContent(ctx, "/"), content), true)
	})
}

// 关于 issue #1626 的讨论，请访问：https://github.com/gogf/gf/issues/1626. md5:3ce099b2ed68c6a6
func Test_Issue1626(t *testing.T) {
	type TestReq struct {
		Name string `v:"required"`
	}
	type TestRes struct {
		Name string
	}
	s := g.Server(guid.S())
	s.Use(
		ghttp.MiddlewareHandlerResponse,
		func(r *ghttp.Request) {
			r.Middleware.Next()
			if err := r.GetError(); err != nil {
				r.Response.ClearBuffer()
				r.Response.Write(err.Error())
			}
		},
	)
	s.BindHandler("/test", func(ctx context.Context, req *TestReq) (res *TestRes, err error) {
		return &TestRes{Name: req.Name}, nil
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(c.GetContent(ctx, "/test"), `The Name field is required`)
		t.Assert(
			gstr.Contains(c.GetContent(ctx, "/test?name=john"), `{"Name":"john"}`),
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
	s := g.Server(guid.S())
	s.Use(ghttp.MiddlewareHandlerResponse)
	s.Group("/boot", func(grp *ghttp.RouterGroup) {
		grp.Bind(Issue1653Foo)
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()
	time.Sleep(1000 * time.Millisecond)
	// g.Client()测试：
	gtest.C(t, func(t *gtest.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
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
		resContent := c.PostContent(ctx, "/boot/test", dataReq)
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
	g.Dump(req)
	return
}

// https://github.com/gogf/gf/issues/1662
// 
// 这段注释是在引用一个GitHub上的问题链接，该问题是关于Go语言框架gf的。在翻译成中文后，它仍然是：
// 
// https://github.com/gogf/gf/issues/1662
// 
// 因为这是一个URL，直接保留原样即可，无需翻译。如果需要了解具体内容，可以访问该链接查看相关的讨论或问题描述。. md5:6f39757924ae8691
func Test_Issue662(t *testing.T) {
	s := g.Server(guid.S())
	s.Use(ghttp.MiddlewareHandlerResponse)
	s.Group("/boot", func(grp *ghttp.RouterGroup) {
		grp.Bind(Foo1)
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()
	time.Sleep(1000 * time.Millisecond)

	// g.Client()测试：
	// code字段传入空字符串时，校验没有提示
	gtest.C(t, func(t *gtest.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		dataReq := `
{"master":{"active":true,"code":"","created_at":"","created_by":"","created_by_text":"","datasource":"38b6f170-a584-43fc-8912-cc1e9bf1b1a9","description":"币种","folder":false,"path":"[\"XCUR\"]","preset":false,"sort":1000,"sql_text":"SELECT!!!!","superior":null,"updated_at":"","updated_by":"","updated_by_text":"","updated_tick":0,"uuid":""},"translation":[{"code":"zh_CN","text":"币种"},{"code":"en_US","text":"币种"}],"filters":null,"fields":[{"code":"F001","created_at":"2022-01-18 23:37:38","created_by":"3ed72aba-1622-4262-a61e-83581e020763","field":"value","hide":false,"min_width":120,"name":"value","parent":"296154bf-b718-4e8f-8b70-efb969b831ec","updated_at":"2022-01-18 23:37:38","updated_by":"3ed72aba-1622-4262-a61e-83581e020763","updated_tick":1,"uuid":"f2140b7a-044c-41c3-b70e-852e6160b21b"},{"code":"F002","created_at":"2022-01-18 23:37:38","created_by":"3ed72aba-1622-4262-a61e-83581e020763","field":"label","hide":false,"min_width":120,"name":"label","parent":"296154bf-b718-4e8f-8b70-efb969b831ec","updated_at":"2022-01-18 23:37:38","updated_by":"3ed72aba-1622-4262-a61e-83581e020763","updated_tick":1,"uuid":"2d3bba5d-308b-4dba-bcac-f093e6556eca"}],"limit":0}
`
		t.Assert(c.PostContent(ctx, "/boot/test", dataReq), `{"code":51,"message":"The code is required","data":null}`)
	})
}

type DemoReq struct {
	g.Meta `path:"/demo" method:"post"`
	Data   *gjson.Json
}

type DemoRes struct {
	Content string
}

type Api struct{}

func (a *Api) Demo(ctx context.Context, req *DemoReq) (res *DemoRes, err error) {
	return &DemoRes{
		Content: req.Data.MustToJsonString(),
	}, err
}

var api = Api{}

// https://github.com/gogf/gf/issues/2172
// 
// 这段注释引用的是一个GitHub问题（issue）的链接，来自 "gf"（Go Foundation）项目。gf是一个用Go语言编写的轻量级ORM（对象关系映射）库。这个注释可能是在讨论与issue号2172相关的某个问题或功能请求，具体内容需要查看GitHub页面以获取详细信息。. md5:65dacaec6e7eb465
func Test_Issue2172(t *testing.T) {
	s := g.Server(guid.S())
	s.Use(ghttp.MiddlewareHandlerResponse)
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Bind(api)
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()
	time.Sleep(1000 * time.Millisecond)

	gtest.C(t, func(t *gtest.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		dataReq := `{"data":{"asd":1}}`
		t.Assert(c.PostContent(ctx, "/demo", dataReq), `{"code":0,"message":"","data":{"Content":"{\"asd\":1}"}}`)
	})
}

// https://github.com/gogf/gf/issues/2334
// 
// 这段注释指的是在 GitHub 上的 "gogf/gf" 项目中，关于 issue（问题或建议）编号为 2334 的讨论或问题。"gf" 可能是指一个名为 "GF"（Go Framework，Go 框架）的开源项目。. md5:b25df7c1e97eea49
func Test_Issue2334(t *testing.T) {
	s := g.Server(guid.S())
	s.SetServerRoot(gtest.DataPath("static1"))
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()
	time.Sleep(1000 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		t.Assert(c.GetContent(ctx, "/index.html"), "index")

		c.SetHeader("If-Modified-Since", "Mon, 12 Dec 2040 05:53:35 GMT")
		res, _ := c.Get(ctx, "/index.html")
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

// 关于 issue #2482 的讨论，请访问：https://github.com/gogf/gf/issues/2482. md5:fcdf7bf6606e8f79
func Test_Issue2482(t *testing.T) {
	s := g.Server(guid.S())
	s.Group("/api/v2", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(OrderController{})
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()
	time.Sleep(1000 * time.Millisecond)

	c := g.Client()
	c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
	gtest.C(t, func(t *gtest.T) {
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
		t.Assert(c.PutContent(ctx, "/api/v2/order", content), `{"code":51,"message":"请输入物料名称","data":null}`)
	})
	gtest.C(t, func(t *gtest.T) {
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
		t.Assert(c.PutContent(ctx, "/api/v2/order", content), `{"code":51,"message":"请输入客户编号","data":null}`)
	})
	gtest.C(t, func(t *gtest.T) {
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
		t.Assert(c.PutContent(ctx, "/api/v2/order", content), `{"code":0,"message":"","data":null}`)
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
	g.RequestFromCtx(ctx).Response.Write(req.Enums)
	return
}

// https://github.com/gogf/gf/issues/2890
// 
// 这段注释是引用了GitHub上gf框架的一个问题链接，编号为2890。在Go语言中，这种注释用于提供外部资源的链接或者问题、功能的参考。翻译成中文后，它的含义不变，仍然是：“https://github.com/gogf/gf/issues/2890”。. md5:a335daa25707abc9
func Test_Issue2890(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		oldEnumsJson, err := gtag.GetGlobalEnums()
		t.AssertNil(err)
		defer t.AssertNil(gtag.SetGlobalEnums(oldEnumsJson))

		err = gtag.SetGlobalEnums(`{"github.com/gogf/gf/v2/net/ghttp_test.Issue2890Enum": ["a","b"]}`)
		t.AssertNil(err)

		s := g.Server(guid.S())
		s.Group("/api/v2", func(group *ghttp.RouterGroup) {
			group.Middleware(ghttp.MiddlewareHandlerResponse)
			group.Bind(Issue2890Controller{})
		})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()
		time.Sleep(1000 * time.Millisecond)

		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		t.Assert(
			c.PostContent(ctx, "/api/v2/issue2890", ``),
			`{"code":51,"message":"The Enums field is required","data":null}`,
		)
		t.Assert(
			c.PostContent(ctx, "/api/v2/issue2890", `{"Enums":"c"}`),
			"{\"code\":51,\"message\":\"The Enums value `c` should be in enums of: [\\\"a\\\",\\\"b\\\"]\",\"data\":null}",
		)
	})
}

// https://github.com/gogf/gf/issues/2963
// 
// 这段注释引用的是GitHub上的一个 issue，地址为：https://github.com/gogf/gf/issues/2963。gf可能是Go语言的一个库或者项目，而2963是一个issue的编号，可能是开发者社区中关于gf库的一个讨论或者bug报告/需求建议。. md5:526c7a4a48782ddd
func Test_Issue2963(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Server(guid.S())
		s.SetServerRoot(gtest.DataPath("issue2963"))
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)

		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		t.Assert(c.GetContent(ctx, "/1.txt"), `1`)
		t.Assert(c.GetContent(ctx, "/中文G146(1)-icon.txt"), `中文G146(1)-icon`)
		t.Assert(c.GetContent(ctx, "/"+gurl.Encode("中文G146(1)-icon.txt")), `中文G146(1)-icon`)
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
	g.RequestFromCtx(ctx).Response.Write(fmt.Sprintf("%v", req))
	return
}

// https://github.com/gogf/gf/issues/3077
// 
// 这段注释引用的是一个GitHub问题（issues）的链接，来自gogf（GoGF）项目。它表示这是对该项目Issue #3077的一个注释，可能是关于某个特定问题、错误报告、功能请求或者讨论的说明。在中文中，你可能会看到类似于“关于 Issue 3077 的讨论”或“参考 GitHub 问题 3077”的内容。. md5:5667623b3895dec8
func Test_Issue3077(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Server(guid.S())
		s.Group("/", func(group *ghttp.RouterGroup) {
			group.Bind(Issue3077V1{})
		})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)

		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		t.Assert(c.GetContent(ctx, "/echo?a=1&b=2"), `&{{} 1 2}`)
		t.Assert(c.GetContent(ctx, "/echo?"), `&{{} a }`)
	})
}

type ListMessageReq struct {
	g.Meta    `path:"/list" method:"get"`
	StartTime int64
	EndTime   int64
}
type ListMessageRes struct {
	g.Meta
	Title   string
	Content string
}
type BaseRes[T any] struct {
	g.Meta
	Code int
	Data T
	Msg  string
}
type cMessage struct{}

func (c *cMessage) List(ctx context.Context, req *ListMessageReq) (res *BaseRes[*ListMessageRes], err error) {
	res = &BaseRes[*ListMessageRes]{
		Code: 100,
		Data: &ListMessageRes{
			Title:   "title",
			Content: "hello",
		},
	}
	return res, err
}

// GitHub上的问题讨论：https://github.com/gogf/gf/issues/2457. md5:76c442b3ac4849d6
func Test_Issue2457(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Server(guid.S())
		s.Use(ghttp.MiddlewareHandlerResponse)
		s.Group("/", func(group *ghttp.RouterGroup) {
			group.Bind(
				new(cMessage),
			)
		})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)

		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		t.Assert(c.GetContent(ctx, "/list"), `{"code":0,"message":"","data":{"Code":100,"Data":{"Title":"title","Content":"hello"},"Msg":""}}`)
	})
}

// https://github.com/gogf/gf/issues/3245
// 
// 这段注释是引用了GitHub上gf框架的一个问题链接，编号为3245。在Go代码中，这种注释通常用于提供问题、讨论或相关资源的链接，以便其他开发者了解更多的背景信息或追踪代码变动的原因。. md5:d1176a3a58b932c6
type Issue3245Req struct {
	g.Meta      `path:"/hello" method:"get"`
	Name        string `p:"nickname" json:"name"`
	XHeaderName string `p:"Header-Name" in:"header" json:"X-Header-Name"`
	XHeaderAge  uint8  `p:"Header-Age" in:"cookie" json:"X-Header-Age"`
}
type Issue3245Res struct {
	Reply any
}

type Issue3245V1 struct{}

func (Issue3245V1) Hello(ctx context.Context, req *Issue3245Req) (res *Issue3245Res, err error) {
	res = &Issue3245Res{
		Reply: req,
	}
	return
}

func Test_Issue3245(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Server(guid.S())
		s.Use(ghttp.MiddlewareHandlerResponse)
		s.Group("/", func(group *ghttp.RouterGroup) {
			group.Bind(
				new(Issue3245V1),
			)
		})
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)

		c := g.Client()
		c.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
		c.SetHeader("Header-Name", "oldme")
		c.SetCookie("Header-Age", "25")

		expect := `{"code":0,"message":"","data":{"Reply":{"name":"oldme","X-Header-Name":"oldme","X-Header-Age":25}}}`
		t.Assert(c.GetContent(ctx, "/hello?nickname=oldme"), expect)
	})
}
