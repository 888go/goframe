// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package goai_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	gjson "github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/net/goai"
	gtest "github.com/888go/goframe/test/gtest"
	gmeta "github.com/888go/goframe/util/gmeta"
	"github.com/888go/goframe/util/gtag"
)

func Test_Basic(t *testing.T) {
	type CommonReq struct {
		AppId      int64  `json:"appId" v:"required" in:"cookie" description:"应用Id"`
		ResourceId string `json:"resourceId" in:"query" description:"资源Id"`
	}
	type SetSpecInfo struct {
		StorageType string   `v:"required|in:CLOUD_PREMIUM,CLOUD_SSD,CLOUD_HSSD" description:"StorageType"`
		Shards      int32    `description:"shards 分片数"`
		Params      []string `description:"默认参数(json 串-ClickHouseParams)"`
	}
	type CreateResourceReq struct {
		CommonReq
		gmeta.Meta `path:"/CreateResourceReq" method:"POST" tags:"default"`
		Name       string                  `description:"实例名称"`
		Product    string                  `description:"业务类型"`
		Region     string                  `v:"required" description:"区域"`
		SetMap     map[string]*SetSpecInfo `v:"required" description:"配置Map"`
		SetSlice   []SetSpecInfo           `v:"required" description:"配置Slice"`
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
			req = new(CreateResourceReq)
		)
		err = oai.Add(goai.AddInput{
			Object: req,
		})
		t.AssertNil(err)
		// Schema asserts.
		t.Assert(len(oai.Components.Schemas.Map()), 2)
		t.Assert(oai.Components.Schemas.Get(`github.com.888go.goframe.net.goai_test.CreateResourceReq`).Value.Type, goai.TypeObject)
		t.Assert(len(oai.Components.Schemas.Get(`github.com.888go.goframe.net.goai_test.CreateResourceReq`).Value.Properties.Map()), 7)
		t.Assert(oai.Components.Schemas.Get(`github.com.888go.goframe.net.goai_test.CreateResourceReq`).Value.Properties.Get(`appId`).Value.Type, goai.TypeInteger)
		t.Assert(oai.Components.Schemas.Get(`github.com.888go.goframe.net.goai_test.CreateResourceReq`).Value.Properties.Get(`resourceId`).Value.Type, goai.TypeString)

		t.Assert(len(oai.Components.Schemas.Get(`github.com.888go.goframe.net.goai_test.SetSpecInfo`).Value.Properties.Map()), 3)
		t.Assert(oai.Components.Schemas.Get(`github.com.888go.goframe.net.goai_test.SetSpecInfo`).Value.Properties.Get(`Params`).Value.Type, goai.TypeArray)
	})
}

func TestOpenApiV3_Add(t *testing.T) {
	type CommonReq struct {
		AppId      int64  `json:"appId" v:"required" in:"path" description:"应用Id"`
		ResourceId string `json:"resourceId" in:"query" description:"资源Id"`
	}
	type SetSpecInfo struct {
		StorageType string   `v:"required|in:CLOUD_PREMIUM,CLOUD_SSD,CLOUD_HSSD" description:"StorageType"`
		Shards      int32    `description:"shards 分片数"`
		Params      []string `description:"默认参数(json 串-ClickHouseParams)"`
	}
	type CreateResourceReq struct {
		CommonReq
		gmeta.Meta `path:"/CreateResourceReq" method:"POST" tags:"default"`
		Name       string                  `description:"实例名称"`
		Product    string                  `description:"业务类型"`
		Region     string                  `v:"required" description:"区域"`
		SetMap     map[string]*SetSpecInfo `v:"required" description:"配置Map"`
		SetSlice   []SetSpecInfo           `v:"required" description:"配置Slice"`
	}

	type CreateResourceRes struct {
		gmeta.Meta `description:"Demo Response Struct"`
		FlowId     int64 `description:"创建实例流程id"`
	}

	f := func(ctx context.Context, req *CreateResourceReq) (res *CreateResourceRes, err error) {
		return
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
		)
		err = oai.Add(goai.AddInput{
			Path:   "/test1/{appId}",
			Method: http.MethodPut,
			Object: f,
		})
		t.AssertNil(err)

		err = oai.Add(goai.AddInput{
			Path:   "/test1/{appId}",
			Method: http.MethodPost,
			Object: f,
		})
		t.AssertNil(err)
		// fmt.Println(oai.String()) 打印oai的字符串表示形式
		// Schema断言
		// md5:ccde758f296b7f2e
		t.Assert(len(oai.Components.Schemas.Map()), 3)
		t.Assert(oai.Components.Schemas.Get(`github.com.888go.goframe.net.goai_test.CreateResourceReq`).Value.Type, goai.TypeObject)

		t.Assert(len(oai.Components.Schemas.Get(`github.com.888go.goframe.net.goai_test.CreateResourceReq`).Value.Properties.Map()), 7)
		t.Assert(len(oai.Paths["/test1/{appId}"].Put.RequestBody.Value.Content["application/json"].Schema.Value.Properties.Map()), 5)
		t.Assert(len(oai.Paths["/test1/{appId}"].Post.RequestBody.Value.Content["application/json"].Schema.Value.Properties.Map()), 5)

		t.Assert(oai.Paths["/test1/{appId}"].Post.Parameters[0].Value.Schema.Value.Type, goai.TypeInteger)
		t.Assert(oai.Paths["/test1/{appId}"].Post.Parameters[1].Value.Schema.Value.Type, goai.TypeString)

		t.Assert(len(oai.Components.Schemas.Get(`github.com.888go.goframe.net.goai_test.SetSpecInfo`).Value.Properties.Map()), 3)
		t.Assert(oai.Components.Schemas.Get(`github.com.888go.goframe.net.goai_test.SetSpecInfo`).Value.Properties.Get(`Params`).Value.Type, goai.TypeArray)

		// Paths.
		t.Assert(len(oai.Paths), 1)
		t.AssertNE(oai.Paths[`/test1/{appId}`].Put, nil)
		t.Assert(len(oai.Paths[`/test1/{appId}`].Put.Tags), 1)
		t.Assert(len(oai.Paths[`/test1/{appId}`].Put.Parameters), 2)
		t.AssertNE(oai.Paths[`/test1/{appId}`].Post, nil)
		t.Assert(len(oai.Paths[`/test1/{appId}`].Post.Tags), 1)
		t.Assert(len(oai.Paths[`/test1/{appId}`].Post.Parameters), 2)
	})
}

func TestOpenApiV3_Add_Recursive(t *testing.T) {
	type CategoryTreeItem struct {
		Id       uint                `json:"id"`
		ParentId uint                `json:"parent_id"`
		Items    []*CategoryTreeItem `json:"items,omitempty"`
	}

	type CategoryGetTreeReq struct {
		gmeta.Meta  `path:"/category-get-tree" method:"GET" tags:"default"`
		ContentType string `in:"query"`
	}
	type CategoryGetTreeRes struct {
		List []*CategoryTreeItem
	}

	f := func(ctx context.Context, req *CategoryGetTreeReq) (res *CategoryGetTreeRes, err error) {
		return
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
		)
		err = oai.Add(goai.AddInput{
			Path:   "/tree",
			Object: f,
		})
		t.AssertNil(err)
		// Schema asserts.
		t.Assert(len(oai.Components.Schemas.Map()), 3)
		t.Assert(oai.Components.Schemas.Get(`github.com.888go.goframe.net.goai_test.CategoryTreeItem`).Value.Type, goai.TypeObject)
		t.Assert(len(oai.Components.Schemas.Get(`github.com.888go.goframe.net.goai_test.CategoryTreeItem`).Value.Properties.Map()), 3)
	})
}

func TestOpenApiV3_Add_EmptyReqAndRes(t *testing.T) {
	type CaptchaIndexReq struct {
		gmeta.Meta `method:"PUT" summary:"获取默认的验证码" description:"注意直接返回的是图片二进制内容" tags:"前台-验证码"`
	}
	type CaptchaIndexRes struct {
		gmeta.Meta `mime:"png" description:"验证码二进制内容" `
	}

	f := func(ctx context.Context, req *CaptchaIndexReq) (res *CaptchaIndexRes, err error) {
		return
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
		)
		err = oai.Add(goai.AddInput{
			Path:   "/tree",
			Object: f,
		})
		t.AssertNil(err)
		// Schema asserts.
		fmt.Println(oai.String())
	})
}

func TestOpenApiV3_Add_AutoDetectIn(t *testing.T) {
	type Req struct {
		gmeta.Meta `method:"get" tags:"default"`
		Name       string
		Product    string
		Region     string
	}

	type Res struct {
		gmeta.Meta `description:"Demo Response Struct"`
	}

	f := func(ctx context.Context, req *Req) (res *Res, err error) {
		return
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err  error
			oai  = goai.New()
			path = `/test/{product}/{name}`
		)
		err = oai.Add(goai.AddInput{
			Path:   path,
			Method: http.MethodGet,
			Object: f,
		})
		t.AssertNil(err)

		fmt.Println(oai.String())

		t.Assert(len(oai.Components.Schemas.Map()), 2)
		t.Assert(len(oai.Paths), 1)
		t.AssertNE(oai.Paths[path].Get, nil)
		t.Assert(len(oai.Paths[path].Get.Parameters), 3)
		t.Assert(oai.Paths[path].Get.Parameters[0].Value.Name, `Name`)
		t.Assert(oai.Paths[path].Get.Parameters[0].Value.In, goai.ParameterInPath)
		t.Assert(oai.Paths[path].Get.Parameters[1].Value.Name, `Product`)
		t.Assert(oai.Paths[path].Get.Parameters[1].Value.In, goai.ParameterInPath)
		t.Assert(oai.Paths[path].Get.Parameters[2].Value.Name, `Region`)
		t.Assert(oai.Paths[path].Get.Parameters[2].Value.In, goai.ParameterInQuery)
	})
}

func TestOpenApiV3_CommonRequest(t *testing.T) {
	type CommonRequest struct {
		Code    int         `json:"code"    description:"Error code"`
		Message string      `json:"message" description:"Error message"`
		Data    interface{} `json:"data"    description:"Result data for certain request according API definition"`
	}

	type Req struct {
		gmeta.Meta `method:"PUT"`
		Product    string `json:"product" v:"required" description:"Unique product key"`
		Name       string `json:"name"    v:"required" description:"Instance name"`
	}
	type Res struct {
		Product      string `json:"product"      v:"required" description:"Unique product key"`
		TemplateName string `json:"templateName" v:"required" description:"Workflow template name"`
		Version      string `json:"version"      v:"required" description:"Workflow template version"`
		TxID         string `json:"txID"         v:"required" description:"Transaction ID for creating instance"`
		Globals      string `json:"globals"                   description:"Globals"`
	}

	f := func(ctx context.Context, req *Req) (res *Res, err error) {
		return
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
		)

		oai.Config.CommonRequest = CommonRequest{}
		oai.Config.CommonRequestDataField = `Data`

		err = oai.Add(goai.AddInput{
			Path:   "/index",
			Object: f,
		})
		t.AssertNil(err)
		// Schema asserts.
		t.Assert(len(oai.Components.Schemas.Map()), 3)
		t.Assert(len(oai.Paths), 1)
		t.Assert(len(oai.Paths["/index"].Put.RequestBody.Value.Content["application/json"].Schema.Value.Properties.Map()), 3)
	})
}

func TestOpenApiV3_CommonRequest_WithoutDataField_Setting(t *testing.T) {
	type CommonRequest struct {
		Code    int         `json:"code"    description:"Error code"`
		Message string      `json:"message" description:"Error message"`
		Data    interface{} `json:"data"    description:"Result data for certain request according API definition"`
	}

	type PutReq struct {
		gmeta.Meta `method:"PUT"`
		Product    string `json:"product" in:"query" v:"required" description:"Unique product key"`
		Name       string `json:"name"    in:"query"  v:"required" description:"Instance name"`
	}

	type PostReq struct {
		gmeta.Meta `method:"POST"`
		Product    string `json:"product" v:"required" description:"Unique product key"`
		Name       string `json:"name"    v:"required" description:"Instance name"`
	}

	type Res struct {
		Product      string `json:"product"      v:"required" description:"Unique product key"`
		TemplateName string `json:"templateName" v:"required" description:"Workflow template name"`
		Version      string `json:"version"      v:"required" description:"Workflow template version"`
		TxID         string `json:"txID"         v:"required" description:"Transaction ID for creating instance"`
		Globals      string `json:"globals"                   description:"Globals"`
	}

	f := func(ctx context.Context, req *PutReq) (res *Res, err error) {
		return
	}
	f2 := func(ctx context.Context, req *PostReq) (res *Res, err error) {
		return
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
		)

		oai.Config.CommonRequest = CommonRequest{}

		err = oai.Add(goai.AddInput{
			Path:   "/index",
			Object: f,
		})
		t.AssertNil(err)

		err = oai.Add(goai.AddInput{
			Path:   "/index",
			Object: f2,
		})
		t.AssertNil(err)
		// Schema断言。
		// 打印OAI的字符串表示。
		// md5:91f27bfa251a3680
		t.Assert(len(oai.Components.Schemas.Map()), 4)
		t.Assert(len(oai.Paths), 1)
		t.Assert(len(oai.Paths["/index"].Put.Parameters), 2)
		t.Assert(len(oai.Paths["/index"].Put.RequestBody.Value.Content["application/json"].Schema.Value.Properties.Map()), 3)
		t.Assert(len(oai.Paths["/index"].Post.Parameters), 0)
		t.Assert(len(oai.Paths["/index"].Post.RequestBody.Value.Content["application/json"].Schema.Value.Properties.Map()), 5)
	})
}

func TestOpenApiV3_CommonRequest_EmptyRequest(t *testing.T) {
	type CommonRequest struct {
		Code    int         `json:"code"    description:"Error code"`
		Message string      `json:"message" description:"Error message"`
		Data    interface{} `json:"data"    description:"Result data for certain request according API definition"`
	}

	type Req struct {
		gmeta.Meta `method:"Put"`
		Product    string `json:"product" in:"query" v:"required" description:"Unique product key"`
		Name       string `json:"name"    in:"query"  v:"required" description:"Instance name"`
	}
	type Res struct{}

	f := func(ctx context.Context, req *Req) (res *Res, err error) {
		return
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
		)

		oai.Config.CommonRequest = CommonRequest{}
		oai.Config.CommonRequestDataField = `Data`

		err = oai.Add(goai.AddInput{
			Path:   "/index",
			Object: f,
		})
		t.AssertNil(err)
		// Schema断言。
		// 打印OAI的字符串表示。
		// md5:91f27bfa251a3680
		t.Assert(len(oai.Components.Schemas.Map()), 3)
		t.Assert(len(oai.Paths), 1)
		t.Assert(len(oai.Paths["/index"].Put.RequestBody.Value.Content["application/json"].Schema.Value.Properties.Map()), 3)
	})
}

func TestOpenApiV3_CommonRequest_SubDataField(t *testing.T) {
	type CommonReqError struct {
		Code    string `description:"错误码"`
		Message string `description:"错误描述"`
	}

	type CommonReqRequest struct {
		RequestId string          `description:"RequestId"`
		Error     *CommonReqError `json:",omitempty" description:"执行错误信息"`
	}

	type CommonReq struct {
		Request CommonReqRequest
	}

	type PutReq struct {
		gmeta.Meta `method:"Put"`
		Product    string `json:"product" in:"query" v:"required" description:"Unique product key"`
		Name       string `json:"name"    in:"query"  v:"required" description:"Instance name"`
	}

	type PostReq struct {
		gmeta.Meta `method:"Post"`
		Product    string `json:"product" v:"required" description:"Unique product key"`
		Name       string `json:"name"    v:"required" description:"Instance name"`
	}

	type Res struct {
		Product      string `json:"product"      v:"required" description:"Unique product key"`
		TemplateName string `json:"templateName" v:"required" description:"Workflow template name"`
		Version      string `json:"version"      v:"required" description:"Workflow template version"`
		TxID         string `json:"txID"         v:"required" description:"Transaction ID for creating instance"`
		Globals      string `json:"globals"                   description:"Globals"`
	}

	f := func(ctx context.Context, req *PutReq) (res *Res, err error) {
		return
	}
	f2 := func(ctx context.Context, req *PostReq) (res *Res, err error) {
		return
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
		)

		oai.Config.CommonRequest = CommonReq{}
		oai.Config.CommonRequestDataField = `Request.`

		err = oai.Add(goai.AddInput{
			Path:   "/index",
			Object: f,
		})
		t.AssertNil(err)

		err = oai.Add(goai.AddInput{
			Path:   "/index",
			Object: f2,
		})
		t.AssertNil(err)

		// Schema断言。
		// 打印OAI的字符串表示。
		// md5:91f27bfa251a3680
		t.Assert(len(oai.Components.Schemas.Map()), 5)
		t.Assert(len(oai.Paths), 1)
		t.Assert(len(oai.Paths["/index"].Put.RequestBody.Value.Content["application/json"].Schema.Value.Properties.Map()), 1)
		t.Assert(len(oai.Paths["/index"].Put.RequestBody.Value.Content["application/json"].Schema.Value.Properties.Get(`Request`).Value.Properties.Map()), 2)
		t.Assert(len(oai.Paths["/index"].Post.RequestBody.Value.Content["application/json"].Schema.Value.Properties.Map()), 1)
		t.Assert(len(oai.Paths["/index"].Post.RequestBody.Value.Content["application/json"].Schema.Value.Properties.Get(`Request`).Value.Properties.Map()), 4)
	})
}

func TestOpenApiV3_CommonResponse(t *testing.T) {
	type CommonResponse struct {
		Code    int         `json:"code"    description:"Error code"`
		Message string      `json:"message" description:"Error message"`
		Data    interface{} `json:"data"    description:"Result data for certain request according API definition"`
	}

	type Req struct {
		gmeta.Meta `method:"GET"`
		Product    string `json:"product" in:"query" v:"required" description:"Unique product key"`
		Name       string `json:"name"    in:"query"  v:"required" description:"Instance name"`
	}
	type Res struct {
		Product      string `json:"product"      v:"required" description:"Unique product key"`
		TemplateName string `json:"templateName" v:"required" description:"Workflow template name"`
		Version      string `json:"version"      v:"required" description:"Workflow template version"`
		TxID         string `json:"txID"         v:"required" description:"Transaction ID for creating instance"`
		Globals      string `json:"globals"                   description:"Globals"`
	}

	f := func(ctx context.Context, req *Req) (res *Res, err error) {
		return
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
		)

		oai.Config.CommonResponse = CommonResponse{}
		oai.Config.CommonResponseDataField = `Data`

		err = oai.Add(goai.AddInput{
			Path:   "/index",
			Object: f,
		})
		t.AssertNil(err)

		// 使用g.Dump函数，获取oai.Paths中"/index"的GET请求的"200"响应值中的Content部分，针对"application/json"媒体类型，然后进一步获取其Schema属性的值，最后处理得到的Properties映射（Map()）。
		// 进行Schema断言。
		// md5:e78f685e8fcc2fc0
		t.Assert(len(oai.Components.Schemas.Map()), 3)
		t.Assert(len(oai.Paths), 1)
		t.Assert(len(oai.Paths["/index"].Get.Responses["200"].Value.Content["application/json"].Schema.Value.Properties.Map()), 3)
		t.Assert(
			oai.Paths["/index"].Get.Responses["200"].Value.Content["application/json"].Schema.Value.Properties.Get("data").Value.Description,
			`Result data for certain request according API definition`,
		)
	})
}

func TestOpenApiV3_CommonResponse_WithoutDataField_Setting(t *testing.T) {
	type CommonResponse struct {
		Code    int         `json:"code"    description:"Error code"`
		Message string      `json:"message" description:"Error message"`
		Data    interface{} `json:"data"    description:"Result data for certain request according API definition"`
	}

	type Req struct {
		gmeta.Meta `method:"GET"`
		Product    string `json:"product" in:"query" v:"required" description:"Unique product key"`
		Name       string `json:"name"    in:"query"  v:"required" description:"Instance name"`
	}
	type Res struct {
		Product      string `json:"product"      v:"required" description:"Unique product key"`
		TemplateName string `json:"templateName" v:"required" description:"Workflow template name"`
		Version      string `json:"version"      v:"required" description:"Workflow template version"`
		TxID         string `json:"txID"         v:"required" description:"Transaction ID for creating instance"`
		Globals      string `json:"globals"                   description:"Globals"`
	}

	f := func(ctx context.Context, req *Req) (res *Res, err error) {
		return
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
		)

		oai.Config.CommonResponse = CommonResponse{}

		err = oai.Add(goai.AddInput{
			Path:   "/index",
			Object: f,
		})
		t.AssertNil(err)
		// Schema asserts.
		fmt.Println(oai.String())
		t.Assert(len(oai.Components.Schemas.Map()), 3)
		t.Assert(len(oai.Paths), 1)
		t.Assert(len(oai.Paths["/index"].Get.Responses["200"].Value.Content["application/json"].Schema.Value.Properties.Map()), 8)
	})
}

func TestOpenApiV3_CommonResponse_EmptyResponse(t *testing.T) {
	type CommonResponse struct {
		Code    int         `json:"code"    description:"Error code"`
		Message string      `json:"message" description:"Error message"`
		Data    interface{} `json:"data"    description:"Result data for certain request according API definition"`
	}

	type Req struct {
		gmeta.Meta `method:"PUT"`
		Product    string `json:"product" v:"required" description:"Unique product key"`
		Name       string `json:"name"    v:"required" description:"Instance name"`
	}
	type Res struct{}

	f := func(ctx context.Context, req *Req) (res *Res, err error) {
		return
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
		)

		oai.Config.CommonResponse = CommonResponse{}
		oai.Config.CommonResponseDataField = `Data`

		err = oai.Add(goai.AddInput{
			Path:   "/index",
			Object: f,
		})
		t.AssertNil(err)
		// Schema断言。
		// 打印OAI的字符串表示。
		// md5:91f27bfa251a3680
		t.Assert(len(oai.Components.Schemas.Map()), 3)
		t.Assert(len(oai.Paths), 1)
		t.Assert(oai.Paths["/index"].Put.RequestBody.Value.Content["application/json"].Schema.Ref, `github.com.888go.goframe.net.goai_test.Req`)
		t.Assert(len(oai.Paths["/index"].Put.Responses["200"].Value.Content["application/json"].Schema.Value.Properties.Map()), 3)
	})
}

func TestOpenApiV3_CommonResponse_SubDataField(t *testing.T) {
	type CommonResError struct {
		Code    string `description:"错误码"`
		Message string `description:"错误描述"`
	}

	type CommonResResponse struct {
		RequestId string          `description:"RequestId"`
		Error     *CommonResError `json:",omitempty" description:"执行错误信息"`
	}

	type CommonRes struct {
		Response CommonResResponse
	}

	type Req struct {
		gmeta.Meta `method:"GET"`
		Product    string `json:"product" in:"query" v:"required" description:"Unique product key"`
		Name       string `json:"name"    in:"query"  v:"required" description:"Instance name"`
	}

	type Res struct {
		Product      string `json:"product"      v:"required" description:"Unique product key"`
		TemplateName string `json:"templateName" v:"required" description:"Workflow template name"`
		Version      string `json:"version"      v:"required" description:"Workflow template version"`
		TxID         string `json:"txID"         v:"required" description:"Transaction ID for creating instance"`
		Globals      string `json:"globals"                   description:"Globals"`
	}

	f := func(ctx context.Context, req *Req) (res *Res, err error) {
		return
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
		)

		oai.Config.CommonResponse = CommonRes{}
		oai.Config.CommonResponseDataField = `Response.`

		err = oai.Add(goai.AddInput{
			Path:   "/index",
			Object: f,
		})
		t.AssertNil(err)
		// Schema断言。
		// 打印OAI的字符串表示。
		// md5:91f27bfa251a3680
		t.Assert(len(oai.Components.Schemas.Map()), 4)
		t.Assert(len(oai.Paths), 1)
		t.Assert(len(oai.Paths["/index"].Get.Responses["200"].Value.Content["application/json"].Schema.Value.Properties.Map()), 1)
		t.Assert(len(oai.Paths["/index"].Get.Responses["200"].Value.Content["application/json"].Schema.Value.Properties.Get(`Response`).Value.Properties.Map()), 7)
	})
}

func TestOpenApiV3_ShortTags(t *testing.T) {
	type CommonReq struct {
		AppId      int64  `json:"appId" v:"required" in:"path" dc:"应用Id" sm:"应用Id Summary"`
		ResourceId string `json:"resourceId" in:"query" dc:"资源Id" sm:"资源Id Summary"`
	}
	type SetSpecInfo struct {
		StorageType string   `v:"required|in:CLOUD_PREMIUM,CLOUD_SSD,CLOUD_HSSD" dc:"StorageType"`
		Shards      int32    `dc:"shards 分片数" sm:"Shards Summary"`
		Params      []string `dc:"默认参数(json 串-ClickHouseParams)" sm:"Params Summary"`
	}
	type CreateResourceReq struct {
		CommonReq
		gmeta.Meta `path:"/CreateResourceReq" method:"POST" tags:"default" sm:"CreateResourceReq sum" dc:"CreateResourceReq des"`
		Name       string                  `dc:"实例名称"`
		Product    string                  `dc:"业务类型"`
		Region     string                  `v:"required" dc:"区域"`
		SetMap     map[string]*SetSpecInfo `v:"required" dc:"配置Map"`
		SetSlice   []SetSpecInfo           `v:"required" dc:"配置Slice"`
	}

	type CreateResourceRes struct {
		gmeta.Meta `dc:"Demo Response Struct"`
		FlowId     int64 `dc:"创建实例流程id"`
	}

	f := func(ctx context.Context, req *CreateResourceReq) (res *CreateResourceRes, err error) {
		return
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
		)
		err = oai.Add(goai.AddInput{
			Path:   "/test1/{appId}",
			Method: http.MethodPut,
			Object: f,
		})
		t.AssertNil(err)

		err = oai.Add(goai.AddInput{
			Path:   "/test1/{appId}",
			Method: http.MethodPost,
			Object: f,
		})
		t.AssertNil(err)
		// fmt.Println(oai.String()) 打印oai的字符串表示形式
		// Schema断言
		// md5:ccde758f296b7f2e
		t.Assert(len(oai.Components.Schemas.Map()), 3)
		t.Assert(oai.Paths[`/test1/{appId}`].Summary, ``)
		t.Assert(oai.Paths[`/test1/{appId}`].Description, ``)
		t.Assert(oai.Paths[`/test1/{appId}`].Put.Summary, `CreateResourceReq sum`)
		t.Assert(oai.Paths[`/test1/{appId}`].Put.Description, `CreateResourceReq des`)
		t.Assert(oai.Paths[`/test1/{appId}`].Put.Parameters[1].Value.Schema.Value.Description, `资源Id`)
		t.Assert(oai.Components.Schemas.Get(`github.com.888go.goframe.net.goai_test.CreateResourceReq`).Value.Properties.Get(`Name`).Value.Description, `实例名称`)
	})
}

func TestOpenApiV3_HtmlResponse(t *testing.T) {
	type Req struct {
		g.Meta `path:"/test" method:"get" summary:"展示内容详情页面" tags:"内容"`
		Id     uint `json:"id" v:"min:1#请选择查看的内容" dc:"内容id"`
	}
	type Res struct {
		g.Meta `mime:"text/html" type:"string" example:"<html/>"`
	}

	f := func(ctx context.Context, req *Req) (res *Res, err error) {
		return
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
		)
		err = oai.Add(goai.AddInput{
			Path:   "/test",
			Method: http.MethodGet,
			Object: f,
		})
		t.AssertNil(err)

		// fmt.Println(oai.String())
		t.Assert(oai.Components.Schemas.Get(`github.com.888go.goframe.net.goai_test.Res`).Value.Type, goai.TypeString)
	})
}

func TestOpenApiV3_HtmlResponseWithCommonResponse(t *testing.T) {
	type CommonResError struct {
		Code    string `description:"错误码"`
		Message string `description:"错误描述"`
	}

	type CommonResResponse struct {
		RequestId string          `description:"RequestId"`
		Error     *CommonResError `json:",omitempty" description:"执行错误信息"`
	}

	type CommonRes struct {
		Response CommonResResponse
	}

	type Req struct {
		g.Meta `path:"/test" method:"get" summary:"展示内容详情页面" tags:"内容"`
		Id     uint `json:"id" v:"min:1#请选择查看的内容" dc:"内容id"`
	}
	type Res struct {
		g.Meta `mime:"text/html" type:"string" example:"<html/>"`
	}

	f := func(ctx context.Context, req *Req) (res *Res, err error) {
		return
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
		)
		oai.Config.CommonResponse = CommonRes{}
		oai.Config.CommonResponseDataField = `Response.`

		err = oai.Add(goai.AddInput{
			Path:   "/test",
			Method: http.MethodGet,
			Object: f,
		})
		t.AssertNil(err)

		// fmt.Println(oai.String())
		t.Assert(oai.Components.Schemas.Get(`github.com.888go.goframe.net.goai_test.Res`).Value.Type, goai.TypeString)
	})
}

func Test_Required_In_Schema(t *testing.T) {
	type CommonReq struct {
		AppId      int64  `json:"appId" v:"required" in:"cookie" description:"应用Id"`
		ResourceId string `json:"resourceId" in:"query" description:"资源Id"`
	}
	type SetSpecInfo struct {
		StorageType string   `v:"required|in:CLOUD_PREMIUM,CLOUD_SSD,CLOUD_HSSD" description:"StorageType"`
		Shards      int32    `description:"shards 分片数"`
		Params      []string `description:"默认参数(json 串-ClickHouseParams)"`
	}
	type CreateResourceReq struct {
		CommonReq
		gmeta.Meta `path:"/CreateResourceReq" method:"POST" tags:"default"`
		Name       string                  `description:"实例名称"`
		Product    string                  `description:"业务类型"`
		Region     string                  `v:"required|min:1" description:"区域"`
		SetMap     map[string]*SetSpecInfo `v:"required|min:1" description:"配置Map"`
		SetSlice   []SetSpecInfo           `v:"required|min:1" description:"配置Slice"`
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
			req = new(CreateResourceReq)
		)
		err = oai.Add(goai.AddInput{
			Object: req,
		})
		t.AssertNil(err)
		var (
			schemaKey1 = `github.com.888go.goframe.net.goai_test.CreateResourceReq`
			schemaKey2 = `github.com.888go.goframe.net.goai_test.SetSpecInfo`
		)
		t.Assert(oai.Components.Schemas.Map()[schemaKey1].Value.Required, g.Slice别名{
			"appId",
			"Region",
			"SetMap",
			"SetSlice",
		})
		t.Assert(oai.Components.Schemas.Map()[schemaKey2].Value.Required, g.Slice别名{
			"StorageType",
		})
		t.Assert(oai.Components.Schemas.Map()[schemaKey2].Value.Properties.Map()["StorageType"].Value.Enum, g.Slice别名{
			"CLOUD_PREMIUM",
			"CLOUD_SSD",
			"CLOUD_HSSD",
		})
	})
}

//func Test_Properties_In_Sequence(t *testing.T) {
//	type ResourceCreateReq struct {
//		g.Meta           `path:"/resource" tags:"OSS Resource" method:"put" x-sort:"1" summary:"创建实例(发货)"`
//		AppId            uint64 `v:"required" dc:"应用Id"`
//		Uin              string `v:"required" dc:"主用户账号，该资源隶属于的账号"`
//		CreateUin        string `v:"required" dc:"创建实例的用户账号"`
//		Product          string `v:"required" dc:"业务类型" eg:"tdach"`
//		Region           string `v:"required" dc:"地域" eg:"ap-guangzhou"`
//		Zone             string `v:"required" dc:"区域" eg:"ap-guangzhou-1"`
//		Tenant           string `v:"required" dc:"业务自定义数据，透传到底层"`
//		VpcId            string `dc:"业务Vpc Id, TCS场景下非必须"`
//		SubnetId         string `dc:"业务Vpc子网Id"`
//		Name             string `dc:"自定义实例名称，默认和ResourceId一致"`
//		ClusterPreset    string `dc:"业务自定义Cluster定义，透传到底层"`
//		Engine           string `dc:"引擎名称，例如：TxLightning"`
//		Version          string `dc:"引擎版本，例如：10.3.213 (兼容ClickHouse 21.3.12)"`
//		SkipUpdateStatus bool   `dc:"是否跳过状态更新，继续保持creating" ed:"http://goframe.org"`
//	}
//	gtest.C(t, func(t *gtest.T) {
//		var (
//			err error
//			oai = goai.New()
//			req = new(ResourceCreateReq)
//		)
//		err = oai.Add(goai.AddInput{
//			Object: req,
//		})
//		t.AssertNil(err)
//		fmt.Println(oai)
//	})
//}

func TestOpenApiV3_Ignore_Parameter(t *testing.T) {
	type CommonResponse struct {
		Code    int         `json:"code"    description:"Error code"`
		Message string      `json:"message" description:"Error message"`
		Data    interface{} `json:"data"    description:"Result data for certain request according API definition"`
	}
	type ProductSearchReq struct {
		gmeta.Meta `path:"/test" method:"get"`
		Product    string `json:"-" in:"query" v:"required|max:5" description:"Unique product key"`
		Name       string `json:"name"    in:"query"  v:"required" description:"Instance name"`
	}
	type ProductSearchRes struct {
		ID           int64  `json:"-"                         description:"Product ID"`
		Product      string `json:"product"      v:"required" description:"Unique product key"`
		TemplateName string `json:"templateName" v:"required" description:"Workflow template name"`
		Version      string `json:"version"      v:"required" description:"Workflow template version"`
		TxID         string `json:"txID"         v:"required" description:"Transaction ID for creating instance"`
		Globals      string `json:"globals"                   description:"Globals"`
	}

	f := func(ctx context.Context, req *ProductSearchReq) (res *ProductSearchRes, err error) {
		return
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
		)

		oai.Config.CommonResponse = CommonResponse{}

		err = oai.Add(goai.AddInput{
			Object: f,
		})
		t.AssertNil(err)
		// Schema断言。
		// 打印OAI的字符串表示。
		// md5:91f27bfa251a3680
		t.Assert(len(oai.Components.Schemas.Map()), 3)
		t.Assert(len(oai.Paths), 1)
		t.Assert(len(oai.Paths["/test"].Get.Responses["200"].Value.Content["application/json"].Schema.Value.Properties.Map()), 8)
	})
}

func Test_EnumOfSchemaItems(t *testing.T) {
	type CreateResourceReq struct {
		gmeta.Meta `path:"/CreateResourceReq" method:"POST"`
		Members    []string `v:"required|in:a,b,c"`
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
			req = new(CreateResourceReq)
		)
		err = oai.Add(goai.AddInput{
			Object: req,
		})
		t.AssertNil(err)

		t.Assert(
			oai.Components.Schemas.Get(`github.com.888go.goframe.net.goai_test.CreateResourceReq`).Value.
				Properties.Get(`Members`).Value.
				Items.Value.Enum,
			g.Slice别名{"a", "b", "c"},
		)
	})
}

func Test_AliasNameOfAttribute(t *testing.T) {
	type CreateResourceReq struct {
		gmeta.Meta `path:"/CreateResourceReq" method:"POST"`
		Name       string `p:"n"`
		Age        string `json:"a"`
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
			req = new(CreateResourceReq)
		)
		err = oai.Add(goai.AddInput{
			Object: req,
		})
		t.AssertNil(err)

		t.Assert(
			oai.Components.Schemas.Get(`github.com.888go.goframe.net.goai_test.CreateResourceReq`).Value.
				Properties.Get(`Name`), nil,
		)
		t.Assert(
			oai.Components.Schemas.Get(`github.com.888go.goframe.net.goai_test.CreateResourceReq`).Value.
				Properties.Get(`Age`), nil,
		)
		t.AssertNE(
			oai.Components.Schemas.Get(`github.com.888go.goframe.net.goai_test.CreateResourceReq`).Value.
				Properties.Get(`n`), nil,
		)
		t.AssertNE(
			oai.Components.Schemas.Get(`github.com.888go.goframe.net.goai_test.CreateResourceReq`).Value.
				Properties.Get(`a`), nil,
		)
	})
}

func Test_EmbeddedStructAttribute(t *testing.T) {
	type CreateResourceReq struct {
		gmeta.Meta `path:"/CreateResourceReq" method:"POST"`
		Name       string `dc:"This is name."`
		Embedded   struct {
			Age uint `dc:"This is embedded age."`
		}
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
			req = new(CreateResourceReq)
		)
		err = oai.Add(goai.AddInput{
			Object: req,
		})
		t.AssertNil(err)

		b, err := json.Marshal(oai)
		t.AssertNil(err)
		t.Assert(b, `{"openapi":"3.0.0","components":{"schemas":{"github.com.888go.goframe.net.goai_test.CreateResourceReq":{"properties":{"Name":{"description":"This is name.","format":"string","properties":{},"type":"string"},"Embedded":{"properties":{"Age":{"description":"This is embedded age.","format":"uint","properties":{},"type":"integer"}},"type":"object"}},"type":"object"}}},"info":{"title":"","version":""},"paths":null}`)
	})
}

func Test_NameFromJsonTag(t *testing.T) {
	// POST
	gtest.C(t, func(t *gtest.T) {
		type CreateReq struct {
			gmeta.Meta `path:"/CreateReq" method:"POST"`
			Name       string `json:"nick_name, omitempty"`
		}

		var (
			err error
			oai = goai.New()
			req = new(CreateReq)
		)
		err = oai.Add(goai.AddInput{
			Object: req,
		})
		t.AssertNil(err)

		b, err := json.Marshal(oai)
		t.AssertNil(err)
		t.Assert(b, `{"openapi":"3.0.0","components":{"schemas":{"github.com.888go.goframe.net.goai_test.CreateReq":{"properties":{"nick_name":{"format":"string","properties":{},"type":"string"}},"type":"object"}}},"info":{"title":"","version":""},"paths":null}`)
	})
	// GET
	gtest.C(t, func(t *gtest.T) {
		type CreateReq struct {
			gmeta.Meta `path:"/CreateReq" method:"GET"`
			Name       string `json:"nick_name, omitempty" in:"header"`
		}
		var (
			err error
			oai = goai.New()
			req = new(CreateReq)
		)
		err = oai.Add(goai.AddInput{
			Object: req,
		})
		t.AssertNil(err)

		b, err := json.Marshal(oai)
		t.AssertNil(err)
		fmt.Println(string(b))
		t.Assert(b, `{"openapi":"3.0.0","components":{"schemas":{"github.com.888go.goframe.net.goai_test.CreateReq":{"properties":{"nick_name":{"format":"string","properties":{},"type":"string"}},"type":"object"}}},"info":{"title":"","version":""},"paths":null}`)
	})
}

func TestOpenApiV3_PathSecurity(t *testing.T) {
	type CommonResponse struct {
		Code    int         `json:"code"    description:"Error code"`
		Message string      `json:"message" description:"Error message"`
		Data    interface{} `json:"data"    description:"Result data for certain request according API definition"`
	}

	type Req struct {
		gmeta.Meta `method:"PUT" security:"apiKey"` // 这里的apiKey要和openApi定义的key一致
		Product    string                           `json:"product" v:"required" description:"Unique product key"`
		Name       string                           `json:"name"    v:"required" description:"Instance name"`
	}
	type Res struct{}

	f := func(ctx context.Context, req *Req) (res *Res, err error) {
		return
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
		)

		oai.Config.CommonResponse = CommonResponse{}
		oai.Components = goai.Components{
			SecuritySchemes: goai.SecuritySchemes{
				"apiKey": goai.SecuritySchemeRef{
					Ref: "",
					Value: &goai.SecurityScheme{
						// 此处type是openApi的规定，详见 https://swagger.io/docs/specification/authentication/api-keys/
						Type: "apiKey",
						In:   "header",
						Name: "X-API-KEY",
					},
				},
			},
		}
		err = oai.Add(goai.AddInput{
			Path:   "/index",
			Object: f,
		})
		t.AssertNil(err)
		// Schema asserts.
		fmt.Println(oai.String())
		t.Assert(len(oai.Components.Schemas.Map()), 3)
		t.Assert(len(oai.Components.SecuritySchemes), 1)
		t.Assert(oai.Components.SecuritySchemes["apiKey"].Value.Type, "apiKey")
		t.Assert(len(oai.Paths), 1)
		t.Assert(len(oai.Paths["/index"].Put.Responses["200"].Value.Content["application/json"].Schema.Value.Properties.Map()), 3)
	})
}

func Test_EmptyJsonNameWithOmitEmpty(t *testing.T) {
	type CreateResourceReq struct {
		gmeta.Meta `path:"/CreateResourceReq" method:"POST" tags:"default"`
		Name       string `description:"实例名称" json:",omitempty"`
		Product    string `description:"业务类型" json:",omitempty"`
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
			req = new(CreateResourceReq)
		)
		err = oai.Add(goai.AddInput{
			Object: req,
		})
		t.AssertNil(err)
		var reqKey = "github.com.888go.goframe.net.goai_test.CreateResourceReq"
		t.AssertNE(oai.Components.Schemas.Get(reqKey).Value.Properties.Get("Name"), nil)
		t.AssertNE(oai.Components.Schemas.Get(reqKey).Value.Properties.Get("Product"), nil)
		t.Assert(oai.Components.Schemas.Get(reqKey).Value.Properties.Get("None"), nil)
	})
}

func Test_Enums(t *testing.T) {
	type Status string
	const (
		StatusA Status = "a"
		StatusB Status = "b"
	)
	type Req struct {
		gmeta.Meta `path:"/CreateResourceReq" method:"POST" tags:"default"`
		Name       string    `dc:"实例名称" json:",omitempty"`
		Status1    Status    `dc:"状态1" json:",omitempty"`
		Status2    *Status   `dc:"状态2" json:",omitempty"`
		Status3    []Status  `dc:"状态2" json:",omitempty"`
		Status4    []*Status `dc:"状态2" json:",omitempty"`
	}

	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
			req = new(Req)
		)
		err = gtag.SetGlobalEnums(gjson.X变量到json文本PANI(g.Map{
			"github.com/888go/goframe/net/goai_test.Status": []interface{}{StatusA, StatusB},
		}))
		t.AssertNil(err)

		err = oai.Add(goai.AddInput{
			Object: req,
		})
		t.AssertNil(err)
		var reqKey = "github.com.888go.goframe.net.goai_test.Req"
		t.AssertNE(oai.Components.Schemas.Get(reqKey).Value.Properties.Get("Name"), nil)
		t.Assert(oai.Components.Schemas.Get(reqKey).Value.Properties.Get("Status1").Value.Enum, g.Slice别名{"a", "b"})
		t.Assert(oai.Components.Schemas.Get(reqKey).Value.Properties.Get("Status2").Value.Enum, g.Slice别名{"a", "b"})
		t.Assert(oai.Components.Schemas.Get(reqKey).Value.Properties.Get("Status3").Value.Items.Value.Enum, g.Slice别名{"a", "b"})
		t.Assert(oai.Components.Schemas.Get(reqKey).Value.Properties.Get("Status4").Value.Items.Value.Enum, g.Slice别名{"a", "b"})
	})
}

func Test_XExtension(t *testing.T) {
	type GetListReq struct {
		g.Meta `path:"/user" tags:"User" method:"get" x-group:"User/Info" summary:"Get user list with basic info."`
		Page   int `dc:"Page number" d:"1" x-sort:"1"`
		Size   int `dc:"Size for per page." d:"10" x-sort:"2"`
	}
	gtest.C(t, func(t *gtest.T) {
		var (
			err error
			oai = goai.New()
			req = new(GetListReq)
		)
		err = oai.Add(goai.AddInput{
			Object: req,
		})
		t.AssertNil(err)
		t.Assert(oai.String(), `{"openapi":"3.0.0","components":{"schemas":{"github.com.888go.goframe.net.goai_test.GetListReq":{"properties":{"Page":{"default":1,"description":"Page number","format":"int","properties":{},"type":"integer","x-sort":"1"},"Size":{"default":10,"description":"Size for per page.","format":"int","properties":{},"type":"integer","x-sort":"2"}},"type":"object","x-group":"User/Info"}}},"info":{"title":"","version":""},"paths":null}`)
	})
}
