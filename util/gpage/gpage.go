// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包gpage提供了针对网页的实用分页功能。
package 分页类

import (
	"fmt"
	"math"
	
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
)

// Page 是分页实现器。
// 所有的属性都是公开的，你可以在必要时更改它们。
type Page struct {
	X总数量      int    // Total size.
	X总页数      int    // 总页数，会自动计算得出。
	X当前页码    int    // 当前页码大于等于1。
	X自定义URL模板    string // 自定义URL模板用于生成页面URL。
	X链接标签css名称      string // CSS样式名称，用于HTML链接标签`a`。
	Span标签css名称      string // CSS样式名称，用于HTML span标签`span`，该标签用于首页、当前页和末页标签。
	X选择标签css名称    string // CSS样式名称，用于HTML选择标签`select`。
	X到下一页标签名称    string // 下一个p标签的名称
	X到前一页标签名称    string // 前一个p标签的名称
	X到第一页标签名称   string // 第一个p标签的名称
	X到最后一页标签名称    string // 上一个p标签的名称
	PrevBarTag     string // Tag字符串用于前一柱状图。
	NextBarTag     string // 下一个条形图的标签字符串。
	X分页栏显示页码     int    // 分页栏显示的页码
	Ajax函数名称 string // Ajax 函数名称。如果此属性不为空，则启用 Ajax。
}

const (
	X常量_默认页面名称        = "page"    // DefaultPageName 定义默认页面名称。
	X常量_默认模板占位符 = "{.page}" // DefaultPagePlaceHolder 定义了URL模板中的占位符。
)

// New 创建并返回一个分页管理器。
// 注意，参数`urlTemplate`指定了生成URL的模板，例如：
// /user/list/{.page}，/user/list/{.page}.html，/user/list?page={.page}&type=1 等等。
// `urlTemplate`中的内置变量"{.page}"表示页码，在生成时会被特定的页码替换。
func X创建(总数量, 分页大小, 当前页 int, url模板 string) *Page {
	p := &Page{
		X链接标签css名称:    "GPageLink",
		Span标签css名称:    "GPageSpan",
		X选择标签css名称:  "GPageSelect",
		X到前一页标签名称:  "<",
		X到下一页标签名称:  ">",
		X到第一页标签名称: "|<",
		X到最后一页标签名称:  ">|",
		PrevBarTag:   "<<",
		NextBarTag:   ">>",
		X总数量:    总数量,
		X总页数:    int(math.Ceil(float64(总数量) / float64(分页大小))),
		X当前页码:  当前页,
		X分页栏显示页码:   10,
		X自定义URL模板:  url模板,
	}
	if 当前页 == 0 {
		p.X当前页码 = 1
	}
	return p
}

// NextPage 返回下一页的 HTML 内容。
func (p *Page) X取下一页html() string {
	if p.X当前页码 < p.X总页数 {
		return p.X取链接html(p.X当前页码+1, p.X到下一页标签名称, "")
	}
	return fmt.Sprintf(`<span class="%s">%s</span>`, p.Span标签css名称, p.X到下一页标签名称)
}

// PrevPage 返回上一页的 HTML 内容。
func (p *Page) X取上一页html() string {
	if p.X当前页码 > 1 {
		return p.X取链接html(p.X当前页码-1, p.X到前一页标签名称, "")
	}
	return fmt.Sprintf(`<span class="%s">%s</span>`, p.Span标签css名称, p.X到前一页标签名称)
}

// FirstPage 返回首页面的 HTML 内容。
func (p *Page) X取首页html() string {
	if p.X当前页码 == 1 {
		return fmt.Sprintf(`<span class="%s">%s</span>`, p.Span标签css名称, p.X到第一页标签名称)
	}
	return p.X取链接html(1, p.X到第一页标签名称, "")
}

// LastPage 返回最后一页的 HTML 内容。
func (p *Page) X取最后一页html() string {
	if p.X当前页码 == p.X总页数 {
		return fmt.Sprintf(`<span class="%s">%s</span>`, p.Span标签css名称, p.X到最后一页标签名称)
	}
	return p.X取链接html(p.X总页数, p.X到最后一页标签名称, "")
}

// PageBar 函数返回带有链接（link标签）和段落（span标签）的HTML分页栏内容。
func (p *Page) X取分页栏html() string {
	plus := int(math.Ceil(float64(p.X分页栏显示页码 / 2)))
	if p.X分页栏显示页码-plus+p.X当前页码 > p.X总页数 {
		plus = p.X分页栏显示页码 - p.X总页数 + p.X当前页码
	}
	begin := p.X当前页码 - plus + 1
	if begin < 1 {
		begin = 1
	}
	barContent := ""
	for i := begin; i < begin+p.X分页栏显示页码; i++ {
		if i <= p.X总页数 {
			if i != p.X当前页码 {
				barText := 转换类.String(i)
				barContent += p.X取链接html(i, barText, barText)
			} else {
				barContent += fmt.Sprintf(`<span class="%s">%d</span>`, p.Span标签css名称, i)
			}
		} else {
			break
		}
	}
	return barContent
}

// SelectBar 返回用于分页的 select HTML 内容。
func (p *Page) X取下拉框html() string {
	barContent := fmt.Sprintf(`<select name="%s" onchange="window.location.href=this.value">`, p.X选择标签css名称)
	for i := 1; i <= p.X总页数; i++ {
		if i == p.X当前页码 {
			barContent += fmt.Sprintf(`<option value="%s" selected>%d</option>`, p.X取链接(i), i)
		} else {
			barContent += fmt.Sprintf(`<option value="%s">%d</option>`, p.X取链接(i), i)
		}
	}
	barContent += "</select>"
	return barContent
}

// GetContent 返回预定义模式的页面内容。
// 这些预定义的内容主要用于中文本地化目的。您可以根据此函数的实现来自定义
// 页面函数以获取页面内容。
func (p *Page) X取预定义模式html(预定义编号 int) string {
	switch 预定义编号 {
	case 1:
		p.X到下一页标签名称 = "下一页"
		p.X到前一页标签名称 = "上一页"
		return fmt.Sprintf(
			`%s <span class="current">%d</span> %s`,
			p.X取上一页html(),
			p.X当前页码,
			p.X取下一页html(),
		)

	case 2:
		p.X到下一页标签名称 = "下一页>>"
		p.X到前一页标签名称 = "<<上一页"
		p.X到第一页标签名称 = "首页"
		p.X到最后一页标签名称 = "尾页"
		return fmt.Sprintf(
			`%s%s<span class="current">[第%d页]</span>%s%s第%s页`,
			p.X取首页html(),
			p.X取上一页html(),
			p.X当前页码,
			p.X取下一页html(),
			p.X取最后一页html(),
			p.X取下拉框html(),
		)

	case 3:
		p.X到下一页标签名称 = "下一页"
		p.X到前一页标签名称 = "上一页"
		p.X到第一页标签名称 = "首页"
		p.X到最后一页标签名称 = "尾页"
		pageStr := p.X取首页html()
		pageStr += p.X取上一页html()
		pageStr += p.X取分页栏html()
		pageStr += p.X取下一页html()
		pageStr += p.X取最后一页html()
		pageStr += fmt.Sprintf(
			`<span>当前页%d/%d</span> <span>共%d条</span>`,
			p.X当前页码,
			p.X总页数,
			p.X总数量,
		)
		return pageStr

	case 4:
		p.X到下一页标签名称 = "下一页"
		p.X到前一页标签名称 = "上一页"
		p.X到第一页标签名称 = "首页"
		p.X到最后一页标签名称 = "尾页"
		pageStr := p.X取首页html()
		pageStr += p.X取上一页html()
		pageStr += p.X取分页栏html()
		pageStr += p.X取下一页html()
		pageStr += p.X取最后一页html()
		return pageStr
	}
	return ""
}

// GetUrl 根据给定的页码解析 UrlTemplate，并返回 URL 字符串。
// 注意，UrlTemplate 属性可以是 URL 或包含 "{.page}" 占位符的 URI 字符串，该占位符用于指定页码的位置。
func (p *Page) X取链接(页码编号 int) string {
	return 文本类.X替换(p.X自定义URL模板, X常量_默认模板占位符, 转换类.String(页码编号))
}

// GetLink 返回给定页码的 HTML 链接标签 `a` 的内容。
func (p *Page) X取链接html(页码编号 int, 内容, 标题 string) string {
	if len(p.Ajax函数名称) > 0 {
		return fmt.Sprintf(
			`<a class="%s" href="javascript:%s('%s')" title="%s">%s</a>`,
			p.X链接标签css名称, p.Ajax函数名称, p.X取链接(页码编号), 标题, 内容,
		)
	} else {
		return fmt.Sprintf(
			`<a class="%s" href="%s" title="%s">%s</a>`,
			p.X链接标签css名称, p.X取链接(页码编号), 标题, 内容,
		)
	}
}
