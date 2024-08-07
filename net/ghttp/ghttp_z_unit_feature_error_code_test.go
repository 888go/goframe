// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 静态服务测试。 md5:2105c089651008de

package http类_test

import (
	"fmt"
	"testing"
	"time"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/frame/g"
	ghttp "github.com/888go/goframe/net/ghttp"
	gtest "github.com/888go/goframe/test/gtest"
	guid "github.com/888go/goframe/util/guid"
)

func Test_Error_Code(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := g.Http类(guid.X生成())
		s.X创建分组路由("/", func(group *ghttp.X分组路由) {
			group.X绑定中间件(func(r *ghttp.Request) {
				r.X中间件管理器.Next()
				r.X响应.ClearBuffer()
				r.X响应.X写响应缓冲区(gerror.X取错误码(r.X取错误信息()))
			})
			group.X绑定所有类型("/", func(r *ghttp.Request) {
				panic(gerror.X创建错误码(gcode.New(10000, "", nil), "test error"))
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
