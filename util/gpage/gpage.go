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
	TotalSize      int    // Total size.
	TotalPage      int    // 总页数，由系统自动计算。 md5:f193103cf068ac01
	CurrentPage    int    // 当前页面编号大于或等于 1。 md5:9b8199029dd6b8ba
	UrlTemplate    string // 用于自定义页面URL生成的模板。 md5:0b8b624ee864eb73
	LinkStyle      string // HTML链接标签`a`的CSS样式名称。 md5:36ef7ca8f8320560
	SpanStyle      string // CSS样式名称，用于HTML的`span`标签，该标签用于首页、当前页和最后一页的标记。 md5:4986703e502e7951
	SelectStyle    string // 用于HTML选择标签`select`的CSS样式名称。 md5:f8bef1c5304f46e0
	NextPageTag    string // Tag name for next p.
	PrevPageTag    string // Tag name for prev p.
	FirstPageTag   string // Tag name for first p.
	LastPageTag    string // Tag name for last p.
	PrevBarTag     string // 用于前一个柱状图的标签字符串。 md5:aa5bb3a271974cb9
	NextBarTag     string // 下一个条形的标签字符串。 md5:bc5dad3acc8e9dc8
	PageBarNum     int    // 用于显示的分页条编号。 md5:663d7e28f2e8da2f
	AjaxActionName string // Ajax 函数名。如果此属性不为空，则启用了Ajax。 md5:b0e509303a20d45a
}

const (
	DefaultPageName        = "page"    // DefaultPageName 定义了默认的页面名称。 md5:4cdd682d15037e0c
	DefaultPagePlaceHolder = "{.page}" // DefaultPagePlaceHolder 定义了URL模板中占位符的默认值。 md5:ee96d0a865392462
)

// New 创建并返回一个分页管理器。
// 请注意，参数 `urlTemplate` 指定用于生成URL的模板，例如：
// /user/list/{.page}, /user/list/{.page}.html, /user/list?page={.page}&type=1 等。
// 在 `urlTemplate` 中内置的变量 "{.page}" 表示页码，在生成时将被替换为特定的页码。
// md5:019378bcadf783f6
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

// NextPage 返回下一页的HTML内容。 md5:5e3d6534a771a5bc
func (p *Page) NextPage() string {
	if p.CurrentPage < p.TotalPage {
		return p.GetLink(p.CurrentPage+1, p.NextPageTag, "")
	}
	return fmt.Sprintf(`<span class="%s">%s</span>`, p.SpanStyle, p.NextPageTag)
}

// PrevPage 返回前一页的HTML内容。 md5:f1834cdd39f5958f
func (p *Page) PrevPage() string {
	if p.CurrentPage > 1 {
		return p.GetLink(p.CurrentPage-1, p.PrevPageTag, "")
	}
	return fmt.Sprintf(`<span class="%s">%s</span>`, p.SpanStyle, p.PrevPageTag)
}

// FirstPage 返回首页的HTML内容。 md5:3da5d9517addeef9
func (p *Page) FirstPage() string {
	if p.CurrentPage == 1 {
		return fmt.Sprintf(`<span class="%s">%s</span>`, p.SpanStyle, p.FirstPageTag)
	}
	return p.GetLink(1, p.FirstPageTag, "")
}

// LastPage 返回最后一页的HTML内容。 md5:7b9da4335fd7cabf
func (p *Page) LastPage() string {
	if p.CurrentPage == p.TotalPage {
		return fmt.Sprintf(`<span class="%s">%s</span>`, p.SpanStyle, p.LastPageTag)
	}
	return p.GetLink(p.TotalPage, p.LastPageTag, "")
}

// PageBar 返回带有链接和span标签的HTML页面栏内容。 md5:fdbe83a2ac56364b
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

// SelectBar 用于返回分页的 select HTML 内容。 md5:675aaa94bd2abde3
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

// GetContent 返回预定义模式下的页面内容。
// 这些预定义内容主要为了中文本地化的目的。您可以根据此函数的实现定义自己的
// 页面函数来获取页面内容。
// md5:36d242b683a4fb96
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

// GetUrl 使用给定的页面号解析UrlTemplate，并返回URL字符串。
// 注意，UrlTemplate属性可以是一个URL，也可以是一个包含"{.page}"占位符的URI字符串，
// 该占位符指定了页面号的位置。
// md5:f7db6853b1f3a681
func (p *Page) GetUrl(page int) string {
	return gstr.Replace(p.UrlTemplate, DefaultPagePlaceHolder, gconv.String(page))
}

// GetLink 函数根据给定的页码，返回 HTML 链接标签（`a`）的内容。 md5:fe3c9d8f012c0f0c
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
