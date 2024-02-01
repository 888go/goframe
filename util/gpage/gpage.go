// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包gpage提供了针对网页的实用分页功能。
package gpage
import (
	"fmt"
	"math"
	
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	)
// Page 是分页实现器。
// 所有的属性都是公开的，你可以在必要时更改它们。
type Page struct {
	TotalSize      int    // Total size.
	TotalPage      int    // 总页数，会自动计算得出。
	CurrentPage    int    // 当前页码大于等于1。
	UrlTemplate    string // 自定义URL模板用于生成页面URL。
	LinkStyle      string // CSS样式名称，用于HTML链接标签`a`。
	SpanStyle      string // CSS样式名称，用于HTML span标签`span`，该标签用于首页、当前页和末页标签。
	SelectStyle    string // CSS样式名称，用于HTML选择标签`select`。
	NextPageTag    string // 下一个p标签的名称
	PrevPageTag    string // 前一个p标签的名称
	FirstPageTag   string // 第一个p标签的名称
	LastPageTag    string // 上一个p标签的名称
	PrevBarTag     string // Tag字符串用于前一柱状图。
	NextBarTag     string // 下一个条形图的标签字符串。
	PageBarNum     int    // 分页栏显示的页码
	AjaxActionName string // Ajax 函数名称。如果此属性不为空，则启用 Ajax。
}

const (
	DefaultPageName        = "page"    // DefaultPageName 定义默认页面名称。
	DefaultPagePlaceHolder = "{.page}" // DefaultPagePlaceHolder 定义了URL模板中的占位符。
)

// New 创建并返回一个分页管理器。
// 注意，参数`urlTemplate`指定了生成URL的模板，例如：
// /user/list/{.page}，/user/list/{.page}.html，/user/list?page={.page}&type=1 等等。
// `urlTemplate`中的内置变量"{.page}"表示页码，在生成时会被特定的页码替换。
func New(totalSize, pageSize, currentPage int, urlTemplate string) *Page {
	p := &Page{
		LinkStyle:    "GPageLink",
		SpanStyle:    "GPageSpan",
		SelectStyle:  "GPageSelect",
		PrevPageTag:  "<",
		NextPageTag:  ">",
		FirstPageTag: "|<",
		LastPageTag:  ">|",
		PrevBarTag:   "<<",
		NextBarTag:   ">>",
		TotalSize:    totalSize,
		TotalPage:    int(math.Ceil(float64(totalSize) / float64(pageSize))),
		CurrentPage:  currentPage,
		PageBarNum:   10,
		UrlTemplate:  urlTemplate,
	}
	if currentPage == 0 {
		p.CurrentPage = 1
	}
	return p
}

// NextPage 返回下一页的 HTML 内容。
func (p *Page) NextPage() string {
	if p.CurrentPage < p.TotalPage {
		return p.GetLink(p.CurrentPage+1, p.NextPageTag, "")
	}
	return fmt.Sprintf(`<span class="%s">%s</span>`, p.SpanStyle, p.NextPageTag)
}

// PrevPage 返回上一页的 HTML 内容。
func (p *Page) PrevPage() string {
	if p.CurrentPage > 1 {
		return p.GetLink(p.CurrentPage-1, p.PrevPageTag, "")
	}
	return fmt.Sprintf(`<span class="%s">%s</span>`, p.SpanStyle, p.PrevPageTag)
}

// FirstPage 返回首页面的 HTML 内容。
func (p *Page) FirstPage() string {
	if p.CurrentPage == 1 {
		return fmt.Sprintf(`<span class="%s">%s</span>`, p.SpanStyle, p.FirstPageTag)
	}
	return p.GetLink(1, p.FirstPageTag, "")
}

// LastPage 返回最后一页的 HTML 内容。
func (p *Page) LastPage() string {
	if p.CurrentPage == p.TotalPage {
		return fmt.Sprintf(`<span class="%s">%s</span>`, p.SpanStyle, p.LastPageTag)
	}
	return p.GetLink(p.TotalPage, p.LastPageTag, "")
}

// PageBar 函数返回带有链接（link标签）和段落（span标签）的HTML分页栏内容。
func (p *Page) PageBar() string {
	plus := int(math.Ceil(float64(p.PageBarNum / 2)))
	if p.PageBarNum-plus+p.CurrentPage > p.TotalPage {
		plus = p.PageBarNum - p.TotalPage + p.CurrentPage
	}
	begin := p.CurrentPage - plus + 1
	if begin < 1 {
		begin = 1
	}
	barContent := ""
	for i := begin; i < begin+p.PageBarNum; i++ {
		if i <= p.TotalPage {
			if i != p.CurrentPage {
				barText := gconv.String(i)
				barContent += p.GetLink(i, barText, barText)
			} else {
				barContent += fmt.Sprintf(`<span class="%s">%d</span>`, p.SpanStyle, i)
			}
		} else {
			break
		}
	}
	return barContent
}

// SelectBar 返回用于分页的 select HTML 内容。
func (p *Page) SelectBar() string {
	barContent := fmt.Sprintf(`<select name="%s" onchange="window.location.href=this.value">`, p.SelectStyle)
	for i := 1; i <= p.TotalPage; i++ {
		if i == p.CurrentPage {
			barContent += fmt.Sprintf(`<option value="%s" selected>%d</option>`, p.GetUrl(i), i)
		} else {
			barContent += fmt.Sprintf(`<option value="%s">%d</option>`, p.GetUrl(i), i)
		}
	}
	barContent += "</select>"
	return barContent
}

// GetContent 返回预定义模式的页面内容。
// 这些预定义的内容主要用于中文本地化目的。您可以根据此函数的实现来自定义
// 页面函数以获取页面内容。
func (p *Page) GetContent(mode int) string {
	switch mode {
	case 1:
		p.NextPageTag = "下一页"
		p.PrevPageTag = "上一页"
		return fmt.Sprintf(
			`%s <span class="current">%d</span> %s`,
			p.PrevPage(),
			p.CurrentPage,
			p.NextPage(),
		)

	case 2:
		p.NextPageTag = "下一页>>"
		p.PrevPageTag = "<<上一页"
		p.FirstPageTag = "首页"
		p.LastPageTag = "尾页"
		return fmt.Sprintf(
			`%s%s<span class="current">[第%d页]</span>%s%s第%s页`,
			p.FirstPage(),
			p.PrevPage(),
			p.CurrentPage,
			p.NextPage(),
			p.LastPage(),
			p.SelectBar(),
		)

	case 3:
		p.NextPageTag = "下一页"
		p.PrevPageTag = "上一页"
		p.FirstPageTag = "首页"
		p.LastPageTag = "尾页"
		pageStr := p.FirstPage()
		pageStr += p.PrevPage()
		pageStr += p.PageBar()
		pageStr += p.NextPage()
		pageStr += p.LastPage()
		pageStr += fmt.Sprintf(
			`<span>当前页%d/%d</span> <span>共%d条</span>`,
			p.CurrentPage,
			p.TotalPage,
			p.TotalSize,
		)
		return pageStr

	case 4:
		p.NextPageTag = "下一页"
		p.PrevPageTag = "上一页"
		p.FirstPageTag = "首页"
		p.LastPageTag = "尾页"
		pageStr := p.FirstPage()
		pageStr += p.PrevPage()
		pageStr += p.PageBar()
		pageStr += p.NextPage()
		pageStr += p.LastPage()
		return pageStr
	}
	return ""
}

// GetUrl 根据给定的页码解析 UrlTemplate，并返回 URL 字符串。
// 注意，UrlTemplate 属性可以是 URL 或包含 "{.page}" 占位符的 URI 字符串，该占位符用于指定页码的位置。
func (p *Page) GetUrl(page int) string {
	return gstr.Replace(p.UrlTemplate, DefaultPagePlaceHolder, gconv.String(page))
}

// GetLink 返回给定页码的 HTML 链接标签 `a` 的内容。
func (p *Page) GetLink(page int, text, title string) string {
	if len(p.AjaxActionName) > 0 {
		return fmt.Sprintf(
			`<a class="%s" href="javascript:%s('%s')" title="%s">%s</a>`,
			p.LinkStyle, p.AjaxActionName, p.GetUrl(page), title, text,
		)
	} else {
		return fmt.Sprintf(
			`<a class="%s" href="%s" title="%s">%s</a>`,
			p.LinkStyle, p.GetUrl(page), title, text,
		)
	}
}
