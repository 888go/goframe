// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/888go/goframe/frame/g"
	ghttp "github.com/888go/goframe/net/ghttp"
	gtest "github.com/888go/goframe/test/gtest"
	guid "github.com/888go/goframe/util/guid"
)

func Test_Params_Page(t *testing.T) {
	s := g.Http类(guid.X生成())
	s.X创建分组路由("/", func(group *ghttp.X分组路由) {
		group.X绑定GET("/list", func(r *ghttp.Request) {
			page := r.X取分页类(5, 2)
			r.X响应.X写响应缓冲区(page.X取预定义模式html(4))
		})
		group.X绑定GET("/list/{page}.html", func(r *ghttp.Request) {
			page := r.X取分页类(5, 2)
			r.X响应.X写响应缓冲区(page.X取预定义模式html(4))
		})
	})
	s.SetDumpRouterMap(false)
	s.X开始监听()
	defer s.X关闭当前服务()

	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		client := g.X网页类()
		client.X设置url前缀(fmt.Sprintf("http://127.0.0.1:%d", s.X取已监听端口()))

		t.Assert(client.Get文本(ctx, "/list"), `<span class="GPageSpan">首页</span><span class="GPageSpan">上一页</span><span class="GPageSpan">1</span><a class="GPageLink" href="/list?page=2" title="2">2</a><a class="GPageLink" href="/list?page=3" title="3">3</a><a class="GPageLink" href="/list?page=2" title="">下一页</a><a class="GPageLink" href="/list?page=3" title="">尾页</a>`)
		t.Assert(client.Get文本(ctx, "/list?page=3"), `<a class="GPageLink" href="/list?page=1" title="">首页</a><a class="GPageLink" href="/list?page=2" title="">上一页</a><a class="GPageLink" href="/list?page=1" title="1">1</a><a class="GPageLink" href="/list?page=2" title="2">2</a><span class="GPageSpan">3</span><span class="GPageSpan">下一页</span><span class="GPageSpan">尾页</span>`)

		t.Assert(client.Get文本(ctx, "/list/1.html"), `<span class="GPageSpan">首页</span><span class="GPageSpan">上一页</span><span class="GPageSpan">1</span><a class="GPageLink" href="/list/2.html" title="2">2</a><a class="GPageLink" href="/list/3.html" title="3">3</a><a class="GPageLink" href="/list/2.html" title="">下一页</a><a class="GPageLink" href="/list/3.html" title="">尾页</a>`)
		t.Assert(client.Get文本(ctx, "/list/3.html"), `<a class="GPageLink" href="/list/1.html" title="">首页</a><a class="GPageLink" href="/list/2.html" title="">上一页</a><a class="GPageLink" href="/list/1.html" title="1">1</a><a class="GPageLink" href="/list/2.html" title="2">2</a><span class="GPageSpan">3</span><span class="GPageSpan">下一页</span><span class="GPageSpan">尾页</span>`)
	})
}
