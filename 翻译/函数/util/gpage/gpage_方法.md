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
# //zj:
# func (re *Regexp) X取文本() string { 
#    re.F.String()
# }
# //zj:
# 备注结束

[func New(totalSize, pageSize, currentPage int, urlTemplate string) *Page {]
ff=创建
urlTemplate=url模板
currentPage=当前页
pageSize=分页大小
totalSize=总数量

[func (p *Page) NextPage() string {]
ff=取下一页html

[func (p *Page) PrevPage() string {]
ff=取上一页html

[func (p *Page) FirstPage() string {]
ff=取首页html

[func (p *Page) LastPage() string {]
ff=取最后一页html

[func (p *Page) PageBar() string {]
ff=取分页栏html

[func (p *Page) SelectBar() string {]
ff=取下拉框html

[func (p *Page) GetContent(mode int) string {]
ff=取预定义模式html
mode=预定义编号

[func (p *Page) GetUrl(page int) string {]
ff=取链接
page=页码编号

[func (p *Page) GetLink(page int, text, title string) string {]
ff=取链接html
title=标题
text=内容
page=页码编号
