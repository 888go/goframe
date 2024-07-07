# 备注开始
# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 如://ff:取文本
#
# yx=true,此方法优先翻译
# 如: //yx=true

# **_package.md 文件备注:
# bm= 包名,更换新的包名称 
# 如: package gin //bm:gin类

# **_其他.md 文件备注:
# qm= 前面,跳转到前面进行重命名.文档内如果有多个相同的,会一起重命名.
# hm= 后面,跳转到后面进行重命名.文档内如果有多个相同的,会一起重命名.
# cz= 查找,配合前面/后面使用,
# zz= 正则查找,配合前面/后面使用, 有设置正则查找,就不用设置上面的查找
# 如: type Regexp struct {//qm:正则 cz:Regexp struct
#
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# 如:
# type Regexp struct {//th:type Regexp222 struct
#
# cf= 重复,用于重命名多次,
# 如: 
# 一个文档内有2个"One(result interface{}) error"需要重命名.
# 但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

# **_追加.md 文件备注:
# 在代码内追加代码,如:
# //zj:前面一行的代码,如果为空,追加到末尾行
# func (re *Regexp) X取文本() string { 
# re.F.String()
# }
# //zj:
# 备注结束

[TotalSize int]
qm=总数量
cz=TotalSize int

[TotalPage int]
qm=总页数
cz=TotalPage int

[CurrentPage int]
qm=当前页码
cz=CurrentPage int

[UrlTemplate string]
qm=自定义URL模板
cz=UrlTemplate string

[LinkStyle string]
qm=链接标签css名称
cz=LinkStyle string

[SpanStyle string]
qm=Span标签css名称
cz=SpanStyle string

[SelectStyle string]
qm=选择标签css名称
cz=SelectStyle string

[NextPageTag string]
qm=到下一页标签名称
cz=NextPageTag string

[PrevPageTag string]
qm=到前一页标签名称
cz=PrevPageTag string

[FirstPageTag string]
qm=到第一页标签名称
cz=FirstPageTag string

[LastPageTag string]
qm=到最后一页标签名称
cz=LastPageTag string

[PageBarNum int]
qm=分页栏显示页码
cz=PageBarNum int

[AjaxActionName string]
qm=Ajax函数名称
cz=AjaxActionName string

[DefaultPageName = "page"]
qm=常量_默认页面名称
cz=DefaultPageName #等号# "page"

[DefaultPagePlaceHolder = "{.page}"]
qm=常量_默认模板占位符
cz=DefaultPagePlaceHolder #等号# "{.page}"
