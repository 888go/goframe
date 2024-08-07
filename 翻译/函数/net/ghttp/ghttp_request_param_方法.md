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

[func (r *Request) Parse(pointer interface{}) error {]
ff=解析参数到结构
pointer=结构指针

[func (r *Request) ParseQuery(pointer interface{}) error {]
ff=解析URL到结构
pointer=结构指针

[func (r *Request) ParseForm(pointer interface{}) error {]
ff=解析表单到结构
pointer=结构指针

[func (r *Request) Get(key string, def ...interface{}) *gvar.Var {]
ff=Get别名
def=默认值
key=名称

[func (r *Request) GetBody() #左中括号##右中括号#byte {]
ff=取请求体字节集

[func (r *Request) GetBodyString() string {]
ff=取请求体文本

[func (r *Request) GetJson() (*gjson.Json, error) {]
ff=取请求体到json类

[func (r *Request) GetMap(def ...map#左中括号#string#右中括号#interface{}) map#左中括号#string#右中括号#interface{} {]
ff=GetMap别名
def=默认值

[func (r *Request) GetMapStrStr(def ...map#左中括号#string#右中括号#interface{}) map#左中括号#string#右中括号#string {]
ff=GetMapStrStr别名
def=默认值

[func (r *Request) GetStruct(pointer interface{}, mapping ...map#左中括号#string#右中括号#string) error {]
ff=GetStruct别名
pointer=结构指针

[func (r *Request) GetMultipartForm() *multipart.Form {]
ff=取multipart表单对象

[func (r *Request) GetMultipartFiles(name string) #左中括号##右中括号#*multipart.FileHeader {]
ff=取multipart表单文件切片对象
name=名称
