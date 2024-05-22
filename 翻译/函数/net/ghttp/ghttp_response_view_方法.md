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

[func (r *Response) WriteTpl(tpl string, params ...gview.Params) error {]
ff=输出到模板文件
params=模板变量
tpl=模板文件路径

[func (r *Response) WriteTplDefault(params ...gview.Params) error {]
ff=输出到默认模板文件
params=模板变量

[func (r *Response) WriteTplContent(content string, params ...gview.Params) error {]
ff=输出文本模板
params=模板变量
content=文本模板

[func (r *Response) ParseTpl(tpl string, params ...gview.Params) (string, error) {]
ff=解析模板文件
params=模板变量
tpl=模板文件路径

[func (r *Response) ParseTplDefault(params ...gview.Params) (string, error) {]
ff=解析默认模板文件
params=模板变量

[func (r *Response) ParseTplContent(content string, params ...gview.Params) (string, error) {]
ff=解析文本模板
params=模板变量
content=文本模板
