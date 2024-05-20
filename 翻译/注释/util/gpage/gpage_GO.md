
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Package gpage provides useful paging functionality for web pages.
<原文结束>

# <翻译开始>
// 包gpage为网页提供有用的分页功能。. md5:21def24b73b57b89
# <翻译结束>


<原文开始>
// Page is the pagination implementer.
// All the attributes are public, you can change them when necessary.
<原文结束>

# <翻译开始>
// Page 是分页实现者。
// 所有属性都是公开的，您可以在需要时更改它们。
// md5:78469553f947138e
# <翻译结束>


<原文开始>
// Total page, which is automatically calculated.
<原文结束>

# <翻译开始>
// 总页数，由系统自动计算。. md5:f193103cf068ac01
# <翻译结束>


<原文开始>
// Current page number >= 1.
<原文结束>

# <翻译开始>
// 当前页面编号大于或等于 1。. md5:9b8199029dd6b8ba
# <翻译结束>


<原文开始>
// Custom url template for page url producing.
<原文结束>

# <翻译开始>
// 用于自定义页面URL生成的模板。. md5:0b8b624ee864eb73
# <翻译结束>


<原文开始>
// CSS style name for HTML link tag `a`.
<原文结束>

# <翻译开始>
// HTML链接标签`a`的CSS样式名称。. md5:36ef7ca8f8320560
# <翻译结束>


<原文开始>
// CSS style name for HTML span tag `span`, which is used for first, current and last page tag.
<原文结束>

# <翻译开始>
// CSS样式名称，用于HTML的`span`标签，该标签用于首页、当前页和最后一页的标记。. md5:4986703e502e7951
# <翻译结束>


<原文开始>
// CSS style name for HTML select tag `select`.
<原文结束>

# <翻译开始>
// 用于HTML选择标签`select`的CSS样式名称。. md5:f8bef1c5304f46e0
# <翻译结束>


<原文开始>
// Tag string for prev bar.
<原文结束>

# <翻译开始>
// 用于前一个柱状图的标签字符串。. md5:aa5bb3a271974cb9
# <翻译结束>


<原文开始>
// Tag string for next bar.
<原文结束>

# <翻译开始>
// 下一个条形的标签字符串。. md5:bc5dad3acc8e9dc8
# <翻译结束>


<原文开始>
// Page bar number for displaying.
<原文结束>

# <翻译开始>
// 用于显示的分页条编号。. md5:663d7e28f2e8da2f
# <翻译结束>


<原文开始>
// Ajax function name. Ajax is enabled if this attribute is not empty.
<原文结束>

# <翻译开始>
// Ajax 函数名。如果此属性不为空，则启用了Ajax。. md5:b0e509303a20d45a
# <翻译结束>


<原文开始>
// DefaultPageName defines the default page name.
<原文结束>

# <翻译开始>
// DefaultPageName 定义了默认的页面名称。. md5:4cdd682d15037e0c
# <翻译结束>


<原文开始>
// DefaultPagePlaceHolder defines the place holder for the url template.
<原文结束>

# <翻译开始>
// DefaultPagePlaceHolder 定义了URL模板中占位符的默认值。. md5:ee96d0a865392462
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
// 请注意，参数 `urlTemplate` 指定用于生成URL的模板，例如：
// /user/list/{.page}, /user/list/{.page}.html, /user/list?page={.page}&type=1 等。
// 在 `urlTemplate` 中内置的变量 "{.page}" 表示页码，在生成时将被替换为特定的页码。
// md5:019378bcadf783f6
# <翻译结束>


<原文开始>
// NextPage returns the HTML content for the next page.
<原文结束>

# <翻译开始>
// NextPage 返回下一页的HTML内容。. md5:5e3d6534a771a5bc
# <翻译结束>


<原文开始>
// PrevPage returns the HTML content for the previous page.
<原文结束>

# <翻译开始>
// PrevPage 返回前一页的HTML内容。. md5:f1834cdd39f5958f
# <翻译结束>


<原文开始>
// FirstPage returns the HTML content for the first page.
<原文结束>

# <翻译开始>
// FirstPage 返回首页的HTML内容。. md5:3da5d9517addeef9
# <翻译结束>


<原文开始>
// LastPage returns the HTML content for the last page.
<原文结束>

# <翻译开始>
// LastPage 返回最后一页的HTML内容。. md5:7b9da4335fd7cabf
# <翻译结束>


<原文开始>
// PageBar returns the HTML page bar content with link and span tags.
<原文结束>

# <翻译开始>
// PageBar 返回带有链接和span标签的HTML页面栏内容。. md5:fdbe83a2ac56364b
# <翻译结束>


<原文开始>
// SelectBar returns the select HTML content for pagination.
<原文结束>

# <翻译开始>
// SelectBar 用于返回分页的 select HTML 内容。. md5:675aaa94bd2abde3
# <翻译结束>


<原文开始>
// GetContent returns the page content for predefined mode.
// These predefined contents are mainly for chinese localization purpose. You can defines your own
// page function retrieving the page content according to the implementation of this function.
<原文结束>

# <翻译开始>
// GetContent 返回预定义模式下的页面内容。
// 这些预定义内容主要为了中文本地化的目的。您可以根据此函数的实现定义自己的
// 页面函数来获取页面内容。
// md5:36d242b683a4fb96
# <翻译结束>


<原文开始>
// GetUrl parses the UrlTemplate with given page number and returns the URL string.
// Note that the UrlTemplate attribute can be either an URL or an URI string with "{.page}"
// place holder specifying the page number position.
<原文结束>

# <翻译开始>
// GetUrl 使用给定的页面号解析UrlTemplate，并返回URL字符串。
// 注意，UrlTemplate属性可以是一个URL，也可以是一个包含"{.page}"占位符的URI字符串，
// 该占位符指定了页面号的位置。
// md5:f7db6853b1f3a681
# <翻译结束>


<原文开始>
// GetLink returns the HTML link tag `a` content for given page number.
<原文结束>

# <翻译开始>
// GetLink 函数根据给定的页码，返回 HTML 链接标签（`a`）的内容。. md5:fe3c9d8f012c0f0c
# <翻译结束>

