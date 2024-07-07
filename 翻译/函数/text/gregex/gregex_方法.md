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

[func Quote(s string) string {]
ff=转义特殊符号
s=文本

[func Validate(pattern string) error {]
ff=表达式验证
pattern=表达式

[func IsMatch(pattern string, src #左中括号##右中括号#byte) bool {]
ff=是否匹配字节集
src=字节集
pattern=表达式

[func IsMatchString(pattern string, src string) bool {]
ff=是否匹配文本
src=字节集
pattern=表达式

[func Match(pattern string, src #左中括号##右中括号#byte) (#左中括号##右中括号##左中括号##右中括号#byte, error) {]
ff=匹配字节集
src=字节集
pattern=表达式

[func MatchString(pattern string, src string) (#左中括号##右中括号#string, error) {]
ff=匹配文本
src=文本
pattern=表达式

[func MatchAll(pattern string, src #左中括号##右中括号#byte) (#左中括号##右中括号##左中括号##右中括号##左中括号##右中括号#byte, error) {]
ff=匹配全部字节集
src=字节集
pattern=表达式

[func MatchAllString(pattern string, src string) (#左中括号##右中括号##左中括号##右中括号#string, error) {]
ff=匹配全部文本
src=文本
pattern=表达式

[func Replace(pattern string, replace, src #左中括号##右中括号#byte) (#左中括号##右中括号#byte, error) {]
ff=替换字节集
src=字节集
replace=替换字节集
pattern=表达式

[func ReplaceString(pattern, replace, src string) (string, error) {]
ff=替换文本
src=文本
replace=替换文本
pattern=表达式

[func ReplaceFunc(pattern string, src #左中括号##右中括号#byte, replaceFunc func(b #左中括号##右中括号#byte) #左中括号##右中括号#byte) (#左中括号##右中括号#byte, error) {]
ff=替换字节集_函数
replaceFunc=回调函数
src=字节集
pattern=表达式

[func ReplaceStringFunc(pattern string, src string, replaceFunc func(s string) string) (string, error) {]
ff=替换文本_函数
replaceFunc=回调函数
src=文本
pattern=表达式

[func Split(pattern string, src string) #左中括号##右中括号#string {]
ff=分割
src=文本
pattern=表达式
