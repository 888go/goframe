
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// Package gpage provides useful paging functionality for web pages.
<原文结束>

# <翻译开始>
// 包gpage提供了针对网页的实用分页功能。
# <翻译结束>


<原文开始>
// Page is the pagination implementer.
// All the attributes are public, you can change them when necessary.
<原文结束>

# <翻译开始>
// Page 是分页实现器。
// 所有的属性都是公开的，你可以在必要时更改它们。
# <翻译结束>


<原文开始>
// Total page, which is automatically calculated.
<原文结束>

# <翻译开始>
// 总页数，会自动计算得出。
# <翻译结束>







<原文开始>
// Custom url template for page url producing.
<原文结束>

# <翻译开始>
// 自定义URL模板用于生成页面URL。
# <翻译结束>
































<原文开始>
// Page bar number for displaying.
<原文结束>

# <翻译开始>
// 分页栏显示的页码
# <翻译结束>


<原文开始>
// Ajax function name. Ajax is enabled if this attribute is not empty.
<原文结束>

# <翻译开始>
// Ajax 函数名称。如果此属性不为空，则启用 Ajax。
# <翻译结束>


<原文开始>
// DefaultPageName defines the default page name.
<原文结束>

# <翻译开始>
// DefaultPageName 定义默认页面名称。
# <翻译结束>


<原文开始>
// DefaultPagePlaceHolder defines the place holder for the url template.
<原文结束>

# <翻译开始>
// DefaultPagePlaceHolder 定义了URL模板中的占位符。
# <翻译结束>


<原文开始>
// New creates and returns a pagination manager.
// Note that the parameter `urlTemplate` specifies the URL producing template, like:
// /user/list/{.page}, /user/list/{.page}.html, /user/list?page={.page}&type=1, etc.
// The build-in variable in `urlTemplate` "{.page}" specifies the page number, which will be replaced by certain
// page number when producing.
<原文结束>

# <翻译开始>
// New 创建并返回一个分页管理器。
// 注意，参数`urlTemplate`指定了生成URL的模板，例如：
// /user/list/{.page}，/user/list/{.page}.html，/user/list?page={.page}&type=1 等等。
// `urlTemplate`中的内置变量"{.page}"表示页码，在生成时会被特定的页码替换。
# <翻译结束>


<原文开始>
// NextPage returns the HTML content for the next page.
<原文结束>

# <翻译开始>
// NextPage 返回下一页的 HTML 内容。
# <翻译结束>


<原文开始>
// PrevPage returns the HTML content for the previous page.
<原文结束>

# <翻译开始>
// PrevPage 返回上一页的 HTML 内容。
# <翻译结束>


<原文开始>
// FirstPage returns the HTML content for the first page.
<原文结束>

# <翻译开始>
// FirstPage 返回首页面的 HTML 内容。
# <翻译结束>


<原文开始>
// LastPage returns the HTML content for the last page.
<原文结束>

# <翻译开始>
// LastPage 返回最后一页的 HTML 内容。
# <翻译结束>


<原文开始>
// PageBar returns the HTML page bar content with link and span tags.
<原文结束>

# <翻译开始>
// PageBar 函数返回带有链接（link标签）和段落（span标签）的HTML分页栏内容。
# <翻译结束>


<原文开始>
// SelectBar returns the select HTML content for pagination.
<原文结束>

# <翻译开始>
// SelectBar 返回用于分页的 select HTML 内容。
# <翻译结束>


<原文开始>
// GetContent returns the page content for predefined mode.
// These predefined contents are mainly for chinese localization purpose. You can defines your own
// page function retrieving the page content according to the implementation of this function.
<原文结束>

# <翻译开始>
// GetContent 返回预定义模式的页面内容。
// 这些预定义的内容主要用于中文本地化目的。您可以根据此函数的实现来自定义
// 页面函数以获取页面内容。
# <翻译结束>


<原文开始>
// GetUrl parses the UrlTemplate with given page number and returns the URL string.
// Note that the UrlTemplate attribute can be either an URL or an URI string with "{.page}"
// place holder specifying the page number position.
<原文结束>

# <翻译开始>
// GetUrl 根据给定的页码解析 UrlTemplate，并返回 URL 字符串。
// 注意，UrlTemplate 属性可以是 URL 或包含 "{.page}" 占位符的 URI 字符串，该占位符用于指定页码的位置。
# <翻译结束>


<原文开始>
// GetLink returns the HTML link tag `a` content for given page number.
<原文结束>

# <翻译开始>
// GetLink 返回给定页码的 HTML 链接标签 `a` 的内容。
# <翻译结束>


<原文开始>
// CSS style name for HTML link tag `a`.
<原文结束>

# <翻译开始>
// CSS样式名称，用于HTML链接标签`a`。
# <翻译结束>


<原文开始>
// CSS style name for HTML span tag `span`, which is used for first, current and last page tag.
<原文结束>

# <翻译开始>
// CSS样式名称，用于HTML span标签`span`，该标签用于首页、当前页和末页标签。
# <翻译结束>


<原文开始>
// CSS style name for HTML select tag `select`.
<原文结束>

# <翻译开始>
// CSS样式名称，用于HTML选择标签`select`。
# <翻译结束>


<原文开始>
// Current page number >= 1.
<原文结束>

# <翻译开始>
// 当前页码大于等于1。
# <翻译结束>


<原文开始>
// Tag name for next p.
<原文结束>

# <翻译开始>
// 下一个p标签的名称
# <翻译结束>


<原文开始>
// Tag name for prev p.
<原文结束>

# <翻译开始>
// 前一个p标签的名称
# <翻译结束>


<原文开始>
// Tag name for first p.
<原文结束>

# <翻译开始>
// 第一个p标签的名称
# <翻译结束>


<原文开始>
// Tag name for last p.
<原文结束>

# <翻译开始>
// 上一个p标签的名称
# <翻译结束>


<原文开始>
// Tag string for prev bar.
<原文结束>

# <翻译开始>
// Tag字符串用于前一柱状图。
# <翻译结束>


<原文开始>
// Tag string for next bar.
<原文结束>

# <翻译开始>
// 下一个条形图的标签字符串。
# <翻译结束>

