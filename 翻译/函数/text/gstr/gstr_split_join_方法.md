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

[func Split(str, delimiter string) #左中括号##右中括号#string {]
ff=分割
delimiter=用作分割的文本
str=文本

[func SplitAndTrim(str, delimiter string, characterMask ...string) #左中括号##右中括号#string {]
ff=分割并忽略空值
delimiter=用作分割的文本
str=文本

[func Join(array #左中括号##右中括号#string, sep string) string {]
ff=连接
sep=连接符
array=切片

[func JoinAny(array interface{}, sep string) string {]
ff=连接Any
sep=连接符
array=切片

[func Explode(delimiter, str string) #左中括号##右中括号#string {]
ff=Explode别名

[func Implode(glue string, pieces #左中括号##右中括号#string) string {]
ff=Implode别名

[func ChunkSplit(body string, chunkLen int, end string) string {]
ff=长度分割
end=分割符
chunkLen=分割长度
body=文本

[func Fields(str string) #左中括号##右中括号#string {]
ff=单词分割
str=文本
