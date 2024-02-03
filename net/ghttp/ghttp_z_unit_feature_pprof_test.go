// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 静态服务测试。

package ghttp_test

import (
	"fmt"
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/g"
	. "github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/guid"
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

// r, err = client.Get(ctx, "/pprof/profile")
// 翻译：使用client发起一个GET请求，获取"/pprof/profile"资源，将响应赋值给r，错误信息赋值给err
// Assert(err, nil)
// 翻译：断言err为nil，即判断此次请求是否发生错误，如果没有错误则继续执行
// Assert(r.StatusCode, 200)
// 翻译：断言HTTP响应状态码r.StatusCode为200，即判断请求是否成功
// r.Close()
// 翻译：关闭HTTP响应体r，释放相关资源
// 在实际的Go语言项目中，通常不会有一个名为Assert的函数，这可能是一个自定义的错误检查函数。如果是在测试代码中，这可能是模拟了类似testing包中的assert.NoError或require.HTTPStatusEqual等断言行为。

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
