// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package ghttp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/guid"
)

func Test_Params_Page(t *testing.T) {
	s := g.Server(guid.S())
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/list", func(r *ghttp.Request) {
			page := r.GetPage(5, 2)
			r.Response.Write(page.GetContent(4))
		})
		group.GET("/list/{page}.html", func(r *ghttp.Request) {
			page := r.GetPage(5, 2)
			r.Response.Write(page.GetContent(4))
		})
	})
	s.SetDumpRouterMap(false)
	s.Start()
	defer s.Shutdown()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.Client()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))

		t.Assert(client.GetContent(ctx, "/list"), `<span class="GPageSpan">首页</span><span class="GPageSpan">上一页</span><span class="GPageSpan">1</span><a class="GPageLink" href="/list?page=2" title="2">2</a><a class="GPageLink" href="/list?page=3" title="3">3</a><a class="GPageLink" href="/list?page=2" title="">下一页</a><a class="GPageLink" href="/list?page=3" title="">尾页</a>`)
		t.Assert(client.GetContent(ctx, "/list?page=3"), `<a class="GPageLink" href="/list?page=1" title="">首页</a><a class="GPageLink" href="/list?page=2" title="">上一页</a><a class="GPageLink" href="/list?page=1" title="1">1</a><a class="GPageLink" href="/list?page=2" title="2">2</a><span class="GPageSpan">3</span><span class="GPageSpan">下一页</span><span class="GPageSpan">尾页</span>`)

		t.Assert(client.GetContent(ctx, "/list/1.html"), `<span class="GPageSpan">首页</span><span class="GPageSpan">上一页</span><span class="GPageSpan">1</span><a class="GPageLink" href="/list/2.html" title="2">2</a><a class="GPageLink" href="/list/3.html" title="3">3</a><a class="GPageLink" href="/list/2.html" title="">下一页</a><a class="GPageLink" href="/list/3.html" title="">尾页</a>`)
		t.Assert(client.GetContent(ctx, "/list/3.html"), `<a class="GPageLink" href="/list/1.html" title="">首页</a><a class="GPageLink" href="/list/2.html" title="">上一页</a><a class="GPageLink" href="/list/1.html" title="1">1</a><a class="GPageLink" href="/list/2.html" title="2">2</a><span class="GPageSpan">3</span><span class="GPageSpan">下一页</span><span class="GPageSpan">尾页</span>`)
	})
}
