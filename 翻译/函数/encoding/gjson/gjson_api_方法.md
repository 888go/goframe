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

[func (j *Json) Var() *gvar.Var {]
ff=取泛型类

[func (j *Json) IsNil() bool {]
ff=是否为Nil

[func (j *Json) Get(pattern string, def ...interface{}) *gvar.Var {]
ff=取值
def=默认值
pattern=表达式

[func (j *Json) GetJson(pattern string, def ...interface{}) *Json {]
ff=取对象
def=默认值
pattern=表达式

[func (j *Json) GetJsons(pattern string, def ...interface{}) #左中括号##右中括号#*Json {]
ff=取对象数组
def=默认值
pattern=表达式

[func (j *Json) GetJsonMap(pattern string, def ...interface{}) map#左中括号#string#右中括号#*Json {]
ff=取对象Map
def=默认值
pattern=表达式

[func (j *Json) Set(pattern string, value interface{}) error {]
ff=设置值

[func (j *Json) MustSet(pattern string, value interface{}) {]
ff=设置值PANI
value=值
pattern=表达式

[func (j *Json) Remove(pattern string) error {]
ff=删除
pattern=表达式

[func (j *Json) MustRemove(pattern string) {]
ff=删除PANI
pattern=表达式

[func (j *Json) Contains(pattern string) bool {]
ff=是否存在
pattern=表达式

[func (j *Json) Len(pattern string) int {]
ff=取长度
pattern=表达式

[func (j *Json) Append(pattern string, value interface{}) error {]
ff=加入
value=值
pattern=表达式

[func (j *Json) MustAppend(pattern string, value interface{}) {]
ff=加入PANI
value=值
pattern=表达式

[func (j *Json) Map() map#左中括号#string#右中括号#interface{} {]
ff=取Map

[func (j *Json) Array() #左中括号##右中括号#interface{} {]
ff=取数组

[func (j *Json) Scan(pointer interface{}, mapping ...map#左中括号#string#右中括号#string) error {]
ff=取结构体指针
mapping=名称映射
pointer=结构体指针

[func (j *Json) Dump() {]
ff=调试输出
