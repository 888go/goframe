// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 静态服务测试。

package http类_test

import (
	"fmt"
	"testing"
	"time"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/guid"
)

func Test_Error_Code(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := g.Http类(uid类.X生成())
		s.X创建分组路由("/", func(group *http类.RouterGroup) {
			group.X绑定中间件(func(r *http类.Request) {
				r.Middleware.Next()
				r.Response.X清空缓冲区()
				r.Response.X写响应缓冲区(错误类.X取错误码(r.X取错误信息()))
			})
			group.X绑定所有类型("/", func(r *http类.Request) {
				panic(错误类.X创建错误码(错误码类.New(10000, "", nil), "test error"))
			})
		})
		s.SetDumpRouterMap(false)
		s.X开始监听()
		defer s.X关闭当前服务()
		time.Sleep(100 * time.Millisecond)
		c := g.X网页类()
		c.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(c.Get文本(ctx, "/"), "10000")
	})
}
