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

[func All() #左中括号##右中括号#string {]
ff=取全部

[func Map() map#左中括号#string#右中括号#string {]
ff=取Map

[func Get(key string, def ...interface{}) *gvar.Var {]
ff=取值
def=默认值
key=名称

[func Set(key, value string) (err error) {]
ff=设置值
yx=true

[func SetMap(m map#左中括号#string#右中括号#string) (err error) {]
ff=设置Map值
err=错误

[func Contains(key string) bool {]
ff=是否存在
key=名称

[func Remove(key ...string) (err error) {]
ff=删除
err=错误
key=名称

[func GetWithCmd(key string, def ...interface{}) *gvar.Var {]
ff=取值或命令行
def=默认值
key=名称

[func Build(m map#左中括号#string#右中括号#string) #左中括号##右中括号#string {]
ff=Map到切片

[func MapFromEnv(envs #左中括号##右中括号#string) map#左中括号#string#右中括号#string {]
ff=切片到Map
envs=切片

[func MapToEnv(m map#左中括号#string#右中括号#string) #左中括号##右中括号#string {]
ff=MapToEnv别名

[func Filter(envs #左中括号##右中括号#string) #左中括号##右中括号#string {]
ff=切片去重
envs=切片
