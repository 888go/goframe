
# <翻译开始>
TotalSize int
X总数量
<跳到行首>
# <翻译结束>

# <翻译开始>
TotalPage int
X总页数
<跳到行首>
# <翻译结束>

# <翻译开始>
CurrentPage int
X当前页码
<跳到行首>
# <翻译结束>

# <翻译开始>
UrlTemplate string
X自定义URL模板
<跳到行首>
# <翻译结束>

# <翻译开始>
LinkStyle string
X链接标签css名称
<跳到行首>
# <翻译结束>

# <翻译开始>
SpanStyle string
Span标签css名称
<跳到行首>
# <翻译结束>

# <翻译开始>
SelectStyle string
X选择标签css名称
<跳到行首>
# <翻译结束>

# <翻译开始>
NextPageTag string
X到下一页标签名称
<跳到行首>
# <翻译结束>

# <翻译开始>
PrevPageTag string
X到前一页标签名称
<跳到行首>
# <翻译结束>

# <翻译开始>
FirstPageTag string
X到第一页标签名称
<跳到行首>
# <翻译结束>

# <翻译开始>
LastPageTag string
X到最后一页标签名称
<跳到行首>
# <翻译结束>

# <翻译开始>
PageBarNum int
X分页栏显示页码
<跳到行首>
# <翻译结束>

# <翻译开始>
AjaxActionName string
Ajax函数名称
<跳到行首>
# <翻译结束>

# <翻译开始>
DefaultPageName = "page"
X常量_默认页面名称
<跳到行首>
# <翻译结束>

# <翻译开始>
DefaultPagePlaceHolder = "{.page}"
X常量_默认模板占位符
<跳到行首>
# <翻译结束>

# <翻译开始>
func New(totalSize, pageSize, currentPage int, urlTemplate
url模板
# <翻译结束>

# <翻译开始>
func New(totalSize, pageSize, currentPage
当前页
# <翻译结束>

# <翻译开始>
func New(totalSize, pageSize
分页大小
# <翻译结束>

# <翻译开始>
func New(totalSize
总数量
# <翻译结束>

# <翻译开始>
func New
X创建
# <翻译结束>

# <翻译开始>
) NextPage
X取下一页html
# <翻译结束>

# <翻译开始>
) PrevPage
X取上一页html
# <翻译结束>

# <翻译开始>
) FirstPage
X取首页html
# <翻译结束>

# <翻译开始>
) LastPage
X取最后一页html
# <翻译结束>

# <翻译开始>
) PageBar
X取分页栏html
# <翻译结束>

# <翻译开始>
) SelectBar
X取下拉框html
# <翻译结束>

# <翻译开始>
) GetContent(mode
预定义编号
# <翻译结束>

# <翻译开始>
) GetContent
X取预定义模式html
# <翻译结束>

# <翻译开始>
) GetUrl(page
页码编号
# <翻译结束>

# <翻译开始>
) GetUrl
X取链接
# <翻译结束>

# <翻译开始>
) GetLink(page int, text, title
标题
# <翻译结束>

# <翻译开始>
) GetLink(page int, text
内容
# <翻译结束>

# <翻译开始>
) GetLink(page
页码编号
# <翻译结束>

# <翻译开始>
) GetLink
X取链接html
# <翻译结束>
