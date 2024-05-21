# 备注开始
# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 如:
# //ff:取文本

# **_package.md 文件备注:
# bm= 包名,更换新的包名称 
# 如: 
# package gin //bm:gin类

# **_其他.md 文件备注:
# qm= 前面,跳转到前面进行重命名.文档内如果有多个相同的,会一起重命名.
# hm= 后面,跳转到后面进行重命名.文档内如果有多个相同的,会一起重命名.
# cz= 查找,配合前面/后面使用,
# 如:
# type Regexp struct {//qm:正则 cz:Regexp struct
#
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
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

[func (r *Request) SetForm(key string, value interface{}) {]
ff=设置表单值
value=值
key=名称

[func (r *Request) GetForm(key string, def ...interface{}) *gvar.Var {]
ff=取表单值到泛型类
def=默认值
key=名称

[func (r *Request) GetFormMap(kvMap ...map#左中括号#string#右中括号#interface{}) map#左中括号#string#右中括号#interface{} {]
ff=取表单值到Map

[func (r *Request) GetFormMapStrStr(kvMap ...map#左中括号#string#右中括号#interface{}) map#左中括号#string#右中括号#string {]
ff=取表单值到MapStrStr

[func (r *Request) GetFormMapStrVar(kvMap ...map#左中括号#string#右中括号#interface{}) map#左中括号#string#右中括号#*gvar.Var {]
ff=取表单值到Map泛型类

[func (r *Request) GetFormStruct(pointer interface{}, mapping ...map#左中括号#string#右中括号#string) error {]
ff=取表单值到结构
pointer=结构指针
