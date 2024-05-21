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

[func (r Result) IsEmpty() bool {]
ff=是否为空

[func (r Result) Len() int {]
ff=取数量

[func (r Result) Size() int {]
ff=Size别名

[func (r Result) Chunk(size int) #左中括号##右中括号#Result {]
ff=分割
size=数量

[func (r Result) Json() string {]
ff=取json

[func (r Result) Xml(rootTag ...string) string {]
ff=取xml
rootTag=根标记

[func (r Result) List() List {]
ff=取Map数组

[func (r Result) Array(field ...string) #左中括号##右中括号#Value {]
ff=取字段数组
field=字段名称

[func (r Result) MapKeyValue(key string) map#左中括号#string#右中括号#Value {]
ff=取字段Map泛型类
key=字段名称

[func (r Result) MapKeyStr(key string) map#左中括号#string#右中括号#Map {]
ff=取字段MapStr
key=字段名称

[func (r Result) MapKeyInt(key string) map#左中括号#int#右中括号#Map {]
ff=取字段MapInt
key=字段名称

[func (r Result) MapKeyUint(key string) map#左中括号#uint#右中括号#Map {]
ff=取字段MapUint
key=字段名称

[func (r Result) Structs(pointer interface{}) (err error) {]
ff=取数组结构体指针
err=错误
pointer=结构体指针
