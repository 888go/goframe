// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gpage provides useful paging functionality for web pages.
package gpage//bm:分页类

import (
	"fmt"
	"math"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// Page is the pagination implementer.
// All the attributes are public, you can change them when necessary.
type Page struct {
	TotalSize      int//qm:总数量  cz:TotalSize int      // Total size.
	TotalPage      int//qm:总页数  cz:TotalPage int      // Total page, which is automatically calculated.
	CurrentPage    int//qm:当前页码  cz:CurrentPage int      // Current page number >= 1.
	UrlTemplate    string//qm:自定义URL模板  cz:UrlTemplate string   // Custom url template for page url producing.
	LinkStyle      string//qm:链接标签css名称  cz:LinkStyle string   // CSS style name for HTML link tag `a`.
	SpanStyle      string//qm:Span标签css名称  cz:SpanStyle string   // CSS style name for HTML span tag `span`, which is used for first, current and last page tag.
	SelectStyle    string//qm:选择标签css名称  cz:SelectStyle string   // CSS style name for HTML select tag `select`.
	NextPageTag    string//qm:到下一页标签名称  cz:NextPageTag string   // Tag name for next p.
	PrevPageTag    string//qm:到前一页标签名称  cz:PrevPageTag string   // Tag name for prev p.
	FirstPageTag   string//qm:到第一页标签名称  cz:FirstPageTag string   // Tag name for first p.
	LastPageTag    string//qm:到最后一页标签名称  cz:LastPageTag string   // Tag name for last p.
	PrevBarTag     string // Tag string for prev bar.
	NextBarTag     string // Tag string for next bar.
	PageBarNum     int//qm:分页栏显示页码  cz:PageBarNum int      // Page bar number for displaying.
	AjaxActionName string//qm:Ajax函数名称  cz:AjaxActionName string   // Ajax function name. Ajax is enabled if this attribute is not empty.
}

const (
	DefaultPageName        = "page"//qm:常量_默认页面名称  cz:DefaultPageName = "page"      // DefaultPageName defines the default page name.
	DefaultPagePlaceHolder = "{.page}"//qm:常量_默认模板占位符  cz:DefaultPagePlaceHolder = "{.page}"   // DefaultPagePlaceHolder defines the place holder for the url template.
)

// New creates and returns a pagination manager.
// Note that the parameter `urlTemplate` specifies the URL producing template, like:
// /user/list/{.page}, /user/list/{.page}.html, /user/list?page={.page}&type=1, etc.
// The build-in variable in `urlTemplate` "{.page}" specifies the page number, which will be replaced by certain
// page number when producing.
// ff:创建
// totalSize:总数量
// pageSize:分页大小
// currentPage:当前页
// urlTemplate:url模板
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

// NextPage returns the HTML content for the next page.
// ff:取下一页html
// p:
func (p *Page) NextPage() string {
	if p.CurrentPage < p.TotalPage {
		return p.GetLink(p.CurrentPage+1, p.NextPageTag, "")
	}
	return fmt.Sprintf(`<span class="%s">%s</span>`, p.SpanStyle, p.NextPageTag)
}

// PrevPage returns the HTML content for the previous page.
// ff:取上一页html
// p:
func (p *Page) PrevPage() string {
	if p.CurrentPage > 1 {
		return p.GetLink(p.CurrentPage-1, p.PrevPageTag, "")
	}
	return fmt.Sprintf(`<span class="%s">%s</span>`, p.SpanStyle, p.PrevPageTag)
}

// FirstPage returns the HTML content for the first page.
// ff:取首页html
// p:
func (p *Page) FirstPage() string {
	if p.CurrentPage == 1 {
		return fmt.Sprintf(`<span class="%s">%s</span>`, p.SpanStyle, p.FirstPageTag)
	}
	return p.GetLink(1, p.FirstPageTag, "")
}

// LastPage returns the HTML content for the last page.
// ff:取最后一页html
// p:
func (p *Page) LastPage() string {
	if p.CurrentPage == p.TotalPage {
		return fmt.Sprintf(`<span class="%s">%s</span>`, p.SpanStyle, p.LastPageTag)
	}
	return p.GetLink(p.TotalPage, p.LastPageTag, "")
}

// PageBar returns the HTML page bar content with link and span tags.
// ff:取分页栏html
// p:
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

// SelectBar returns the select HTML content for pagination.
// ff:取下拉框html
// p:
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

// GetContent returns the page content for predefined mode.
// These predefined contents are mainly for chinese localization purpose. You can defines your own
// page function retrieving the page content according to the implementation of this function.
// ff:取预定义模式html
// p:
// mode:预定义编号
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

// GetUrl parses the UrlTemplate with given page number and returns the URL string.
// Note that the UrlTemplate attribute can be either an URL or an URI string with "{.page}"
// place holder specifying the page number position.
// ff:取链接
// p:
// page:页码编号
func (p *Page) GetUrl(page int) string {
	return gstr.Replace(p.UrlTemplate, DefaultPagePlaceHolder, gconv.String(page))
}

// GetLink returns the HTML link tag `a` content for given page number.
// ff:取链接html
// p:
// page:页码编号
// text:内容
// title:标题
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
