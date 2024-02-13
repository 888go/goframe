// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package 分页类_test

import (
	"testing"
	
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gpage"
)

func Test_New(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		page := 分页类.X创建(9, 2, 1, `/user/list?page={.page}`)
		t.Assert(page.X总数量, 9)
		t.Assert(page.X总页数, 5)
		t.Assert(page.X当前页码, 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		page := 分页类.X创建(9, 2, 0, `/user/list?page={.page}`)
		t.Assert(page.X总数量, 9)
		t.Assert(page.X总页数, 5)
		t.Assert(page.X当前页码, 1)
	})
}

func Test_Basic(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		page := 分页类.X创建(9, 2, 1, `/user/list?page={.page}`)
		t.Assert(page.X取下一页html(), `<a class="GPageLink" href="/user/list?page=2" title="">></a>`)
		t.Assert(page.X取上一页html(), `<span class="GPageSpan"><</span>`)
		t.Assert(page.X取首页html(), `<span class="GPageSpan">|<</span>`)
		t.Assert(page.X取最后一页html(), `<a class="GPageLink" href="/user/list?page=5" title="">>|</a>`)
		t.Assert(page.X取分页栏html(), `<span class="GPageSpan">1</span><a class="GPageLink" href="/user/list?page=2" title="2">2</a><a class="GPageLink" href="/user/list?page=3" title="3">3</a><a class="GPageLink" href="/user/list?page=4" title="4">4</a><a class="GPageLink" href="/user/list?page=5" title="5">5</a>`)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		page := 分页类.X创建(9, 2, 3, `/user/list?page={.page}`)
		t.Assert(page.X取下一页html(), `<a class="GPageLink" href="/user/list?page=4" title="">></a>`)
		t.Assert(page.X取上一页html(), `<a class="GPageLink" href="/user/list?page=2" title=""><</a>`)
		t.Assert(page.X取首页html(), `<a class="GPageLink" href="/user/list?page=1" title="">|<</a>`)
		t.Assert(page.X取最后一页html(), `<a class="GPageLink" href="/user/list?page=5" title="">>|</a>`)
		t.Assert(page.X取分页栏html(), `<a class="GPageLink" href="/user/list?page=1" title="1">1</a><a class="GPageLink" href="/user/list?page=2" title="2">2</a><span class="GPageSpan">3</span><a class="GPageLink" href="/user/list?page=4" title="4">4</a><a class="GPageLink" href="/user/list?page=5" title="5">5</a>`)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		page := 分页类.X创建(9, 2, 5, `/user/list?page={.page}`)
		t.Assert(page.X取下一页html(), `<span class="GPageSpan">></span>`)
		t.Assert(page.X取上一页html(), `<a class="GPageLink" href="/user/list?page=4" title=""><</a>`)
		t.Assert(page.X取首页html(), `<a class="GPageLink" href="/user/list?page=1" title="">|<</a>`)
		t.Assert(page.X取最后一页html(), `<span class="GPageSpan">>|</span>`)
		t.Assert(page.X取分页栏html(), `<a class="GPageLink" href="/user/list?page=1" title="1">1</a><a class="GPageLink" href="/user/list?page=2" title="2">2</a><a class="GPageLink" href="/user/list?page=3" title="3">3</a><a class="GPageLink" href="/user/list?page=4" title="4">4</a><span class="GPageSpan">5</span>`)
	})
}

func Test_CustomTag(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		page := 分页类.X创建(5, 1, 2, `/user/list/{.page}`)
		page.X到前一页标签名称 = "《"
		page.X到下一页标签名称 = "》"
		page.X到第一页标签名称 = "|《"
		page.X到最后一页标签名称 = "》|"
		page.PrevBarTag = "《《"
		page.NextBarTag = "》》"
		t.Assert(page.X取下一页html(), `<a class="GPageLink" href="/user/list/3" title="">》</a>`)
		t.Assert(page.X取上一页html(), `<a class="GPageLink" href="/user/list/1" title="">《</a>`)
		t.Assert(page.X取首页html(), `<a class="GPageLink" href="/user/list/1" title="">|《</a>`)
		t.Assert(page.X取最后一页html(), `<a class="GPageLink" href="/user/list/5" title="">》|</a>`)
		t.Assert(page.X取分页栏html(), `<a class="GPageLink" href="/user/list/1" title="1">1</a><span class="GPageSpan">2</span><a class="GPageLink" href="/user/list/3" title="3">3</a><a class="GPageLink" href="/user/list/4" title="4">4</a><a class="GPageLink" href="/user/list/5" title="5">5</a>`)
	})
}

func Test_CustomStyle(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		page := 分页类.X创建(5, 1, 2, `/user/list/{.page}`)
		page.X链接标签css名称 = "MyPageLink"
		page.Span标签css名称 = "MyPageSpan"
		page.X选择标签css名称 = "MyPageSelect"
		t.Assert(page.X取下一页html(), `<a class="MyPageLink" href="/user/list/3" title="">></a>`)
		t.Assert(page.X取上一页html(), `<a class="MyPageLink" href="/user/list/1" title=""><</a>`)
		t.Assert(page.X取首页html(), `<a class="MyPageLink" href="/user/list/1" title="">|<</a>`)
		t.Assert(page.X取最后一页html(), `<a class="MyPageLink" href="/user/list/5" title="">>|</a>`)
		t.Assert(page.X取分页栏html(), `<a class="MyPageLink" href="/user/list/1" title="1">1</a><span class="MyPageSpan">2</span><a class="MyPageLink" href="/user/list/3" title="3">3</a><a class="MyPageLink" href="/user/list/4" title="4">4</a><a class="MyPageLink" href="/user/list/5" title="5">5</a>`)
		t.Assert(page.X取下拉框html(), `<select name="MyPageSelect" onchange="window.location.href=this.value"><option value="/user/list/1">1</option><option value="/user/list/2" selected>2</option><option value="/user/list/3">3</option><option value="/user/list/4">4</option><option value="/user/list/5">5</option></select>`)
	})
}

func Test_Ajax(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		page := 分页类.X创建(5, 1, 2, `/user/list/{.page}`)
		page.Ajax函数名称 = "LoadPage"
		t.Assert(page.X取下一页html(), `<a class="GPageLink" href="javascript:LoadPage('/user/list/3')" title="">></a>`)
		t.Assert(page.X取上一页html(), `<a class="GPageLink" href="javascript:LoadPage('/user/list/1')" title=""><</a>`)
		t.Assert(page.X取首页html(), `<a class="GPageLink" href="javascript:LoadPage('/user/list/1')" title="">|<</a>`)
		t.Assert(page.X取最后一页html(), `<a class="GPageLink" href="javascript:LoadPage('/user/list/5')" title="">>|</a>`)
		t.Assert(page.X取分页栏html(), `<a class="GPageLink" href="javascript:LoadPage('/user/list/1')" title="1">1</a><span class="GPageSpan">2</span><a class="GPageLink" href="javascript:LoadPage('/user/list/3')" title="3">3</a><a class="GPageLink" href="javascript:LoadPage('/user/list/4')" title="4">4</a><a class="GPageLink" href="javascript:LoadPage('/user/list/5')" title="5">5</a>`)
	})
}

func Test_PredefinedContent(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		page := 分页类.X创建(5, 1, 2, `/user/list/{.page}`)
		page.Ajax函数名称 = "LoadPage"
		t.Assert(page.X取预定义模式html(1), `<a class="GPageLink" href="javascript:LoadPage('/user/list/1')" title="">上一页</a> <span class="current">2</span> <a class="GPageLink" href="javascript:LoadPage('/user/list/3')" title="">下一页</a>`)
		t.Assert(page.X取预定义模式html(2), `<a class="GPageLink" href="javascript:LoadPage('/user/list/1')" title="">首页</a><a class="GPageLink" href="javascript:LoadPage('/user/list/1')" title=""><<上一页</a><span class="current">[第2页]</span><a class="GPageLink" href="javascript:LoadPage('/user/list/3')" title="">下一页>></a><a class="GPageLink" href="javascript:LoadPage('/user/list/5')" title="">尾页</a>第<select name="GPageSelect" onchange="window.location.href=this.value"><option value="/user/list/1">1</option><option value="/user/list/2" selected>2</option><option value="/user/list/3">3</option><option value="/user/list/4">4</option><option value="/user/list/5">5</option></select>页`)
		t.Assert(page.X取预定义模式html(3), `<a class="GPageLink" href="javascript:LoadPage('/user/list/1')" title="">首页</a><a class="GPageLink" href="javascript:LoadPage('/user/list/1')" title="">上一页</a><a class="GPageLink" href="javascript:LoadPage('/user/list/1')" title="1">1</a><span class="GPageSpan">2</span><a class="GPageLink" href="javascript:LoadPage('/user/list/3')" title="3">3</a><a class="GPageLink" href="javascript:LoadPage('/user/list/4')" title="4">4</a><a class="GPageLink" href="javascript:LoadPage('/user/list/5')" title="5">5</a><a class="GPageLink" href="javascript:LoadPage('/user/list/3')" title="">下一页</a><a class="GPageLink" href="javascript:LoadPage('/user/list/5')" title="">尾页</a><span>当前页2/5</span> <span>共5条</span>`)
		t.Assert(page.X取预定义模式html(4), `<a class="GPageLink" href="javascript:LoadPage('/user/list/1')" title="">首页</a><a class="GPageLink" href="javascript:LoadPage('/user/list/1')" title="">上一页</a><a class="GPageLink" href="javascript:LoadPage('/user/list/1')" title="1">1</a><span class="GPageSpan">2</span><a class="GPageLink" href="javascript:LoadPage('/user/list/3')" title="3">3</a><a class="GPageLink" href="javascript:LoadPage('/user/list/4')" title="4">4</a><a class="GPageLink" href="javascript:LoadPage('/user/list/5')" title="5">5</a><a class="GPageLink" href="javascript:LoadPage('/user/list/3')" title="">下一页</a><a class="GPageLink" href="javascript:LoadPage('/user/list/5')" title="">尾页</a>`)
		t.Assert(page.X取预定义模式html(5), ``)
	})
}
