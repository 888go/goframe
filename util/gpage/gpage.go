// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gpage为网页提供有用的分页功能。 md5:21def24b73b57b89
package 分页类

import (
	"fmt"
	"math"

	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
)

// Page 是分页实现者。
// 所有属性都是公开的，您可以在需要时更改它们。
// md5:78469553f947138e
type Page struct {
	X总数量      int    // Total size.
	X总页数      int    // 总页数，由系统自动计算。 md5:f193103cf068ac01
	X当前页码    int    // 当前页面编号大于或等于 1。 md5:9b8199029dd6b8ba
	X自定义URL模板    string // 用于自定义页面URL生成的模板。 md5:0b8b624ee864eb73
	X链接标签css名称      string // HTML链接标签`a`的CSS样式名称。 md5:36ef7ca8f8320560
	Span标签css名称      string // CSS样式名称，用于HTML的`span`标签，该标签用于首页、当前页和最后一页的标记。 md5:4986703e502e7951
	X选择标签css名称    string // 用于HTML选择标签`select`的CSS样式名称。 md5:f8bef1c5304f46e0
	X到下一页标签名称    string // Tag name for next p.
	X到前一页标签名称    string // Tag name for prev p.
	X到第一页标签名称   string // Tag name for first p.
	X到最后一页标签名称    string // Tag name for last p.
	PrevBarTag     string // 用于前一个柱状图的标签字符串。 md5:aa5bb3a271974cb9
	NextBarTag     string // 下一个条形的标签字符串。 md5:bc5dad3acc8e9dc8
	X分页栏显示页码     int    // 用于显示的分页条编号。 md5:663d7e28f2e8da2f
	Ajax函数名称 string // Ajax 函数名。如果此属性不为空，则启用了Ajax。 md5:b0e509303a20d45a
}

const (
	X常量_默认页面名称        = "page"    // DefaultPageName 定义了默认的页面名称。 md5:4cdd682d15037e0c
	X常量_默认模板占位符 = "{.page}" // DefaultPagePlaceHolder 定义了URL模板中占位符的默认值。 md5:ee96d0a865392462
)

// X创建 创建并返回一个分页管理器。
// 请注意，参数 `urlTemplate` 指定用于生成URL的模板，例如：
// /user/list/{.page}, /user/list/{.page}.html, /user/list?page={.page}&type=1 等。
// 在 `urlTemplate` 中内置的变量 "{.page}" 表示页码，在生成时将被替换为特定的页码。
// md5:019378bcadf783f6
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

// X取下一页html 返回下一页的HTML内容。 md5:5e3d6534a771a5bc
func (p *Page) X取下一页html() string {
	if p.X当前页码 < p.X总页数 {
		return p.X取链接html(p.X当前页码+1, p.X到下一页标签名称, "")
	}
	return fmt.Sprintf(`<span class="%s">%s</span>`, p.Span标签css名称, p.X到下一页标签名称)
}

// X取上一页html 返回前一页的HTML内容。 md5:f1834cdd39f5958f
func (p *Page) X取上一页html() string {
	if p.X当前页码 > 1 {
		return p.X取链接html(p.X当前页码-1, p.X到前一页标签名称, "")
	}
	return fmt.Sprintf(`<span class="%s">%s</span>`, p.Span标签css名称, p.X到前一页标签名称)
}

// X取首页html 返回首页的HTML内容。 md5:3da5d9517addeef9
func (p *Page) X取首页html() string {
	if p.X当前页码 == 1 {
		return fmt.Sprintf(`<span class="%s">%s</span>`, p.Span标签css名称, p.X到第一页标签名称)
	}
	return p.X取链接html(1, p.X到第一页标签名称, "")
}

// X取最后一页html 返回最后一页的HTML内容。 md5:7b9da4335fd7cabf
func (p *Page) X取最后一页html() string {
	if p.X当前页码 == p.X总页数 {
		return fmt.Sprintf(`<span class="%s">%s</span>`, p.Span标签css名称, p.X到最后一页标签名称)
	}
	return p.X取链接html(p.X总页数, p.X到最后一页标签名称, "")
}

// X取分页栏html 返回带有链接和span标签的HTML页面栏内容。 md5:fdbe83a2ac56364b
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
				barText := gconv.String(i)
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

// X取下拉框html 用于返回分页的 select HTML 内容。 md5:675aaa94bd2abde3
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

// X取预定义模式html 返回预定义模式下的页面内容。
// 这些预定义内容主要为了中文本地化的目的。您可以根据此函数的实现定义自己的
// 页面函数来获取页面内容。
// md5:36d242b683a4fb96
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

// X取链接 使用给定的页面号解析UrlTemplate，并返回URL字符串。
// 注意，UrlTemplate属性可以是一个URL，也可以是一个包含"{.page}"占位符的URI字符串，
// 该占位符指定了页面号的位置。
// md5:f7db6853b1f3a681
func (p *Page) X取链接(页码编号 int) string {
	return gstr.X替换(p.X自定义URL模板, X常量_默认模板占位符, gconv.String(页码编号))
}

// X取链接html 函数根据给定的页码，返回 HTML 链接标签（`a`）的内容。 md5:fe3c9d8f012c0f0c
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
