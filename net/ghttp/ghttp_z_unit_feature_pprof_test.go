// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

// 静态服务测试。 md5:2105c089651008de

package ghttp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	. "github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/guid"
)

func TestServer_EnablePProf(t *testing.T) {
	C(t, func(t *T) {
		s := g.Server(guid.S())
		s.EnablePProf("/pprof")
		s.SetDumpRouterMap(false)
		s.Start()
		defer s.Shutdown()
		time.Sleep(100 * time.Millisecond)
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		r, err := client.Get(ctx, "/pprof/index")
		Assert(err, nil)
		Assert(r.StatusCode, 200)
		r.Close()

		r, err = client.Get(ctx, "/pprof/cmdline")
		Assert(err, nil)
		Assert(r.StatusCode, 200)
		r.Close()

		//r, err = client.Get(ctx, "/pprof/profile") 		// 将客户端的GET请求翻译为中文：r, 错误 = 客户端在上下文中获取"/pprof/profile"
		//Assert(err, nil)                           		// 断言错误应为nil：断言错误，应为空
		//Assert(r.StatusCode, 200)                  		// 断言响应状态码应为200：断言响应的状态码，应为200
		//r.Close()                                   		// 关闭响应：关闭r md5:629678dd0441cb92

		r, err = client.Get(ctx, "/pprof/symbol")
		Assert(err, nil)
		Assert(r.StatusCode, 200)
		r.Close()

		r, err = client.Get(ctx, "/pprof/trace")
		Assert(err, nil)
		Assert(r.StatusCode, 200)
		r.Close()
	})

}
