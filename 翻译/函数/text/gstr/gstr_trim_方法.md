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

[func Trim(str string, characterMask ...string) string {]
ff=过滤首尾符并含空白
characterMask=可选过滤符号
str=文本

[func TrimStr(str string, cut string, count ...int) string {]
ff=过滤首尾
count=可选过滤次数
cut=过滤符号
str=文本

[func TrimLeft(str string, characterMask ...string) string {]
ff=过滤首字符并含空白
characterMask=可选过滤符号
str=文本

[func TrimLeftStr(str string, cut string, count ...int) string {]
ff=过滤首字符
count=可选过滤次数
cut=过滤符号
str=文本

[func TrimRight(str string, characterMask ...string) string {]
ff=过滤尾字符并含空白
characterMask=可选过滤符号
str=文本

[func TrimRightStr(str string, cut string, count ...int) string {]
ff=过滤尾字符
count=可选过滤次数
cut=过滤符号
str=文本

[func TrimAll(str string, characterMask ...string) string {]
ff=过滤所有字符并含空白
characterMask=可选过滤符号
str=文本

[func HasPrefix(s, prefix string) bool {]
ff=开头判断
prefix=开头文本
s=文本

[func HasSuffix(s, suffix string) bool {]
ff=末尾判断
suffix=末尾文本
s=文本
