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
# re.F.String()
# }
# //zj:
# 备注结束

[func Chr(ascii int) string {]
ff=整数到ascii
ascii=整数

[func OctStr(str string) string {]
ff=八进制到文本
str=文本

[func Reverse(str string) string {]
ff=反转字符
str=文本

[func NumberFormat(number float64, decimals int, decPoint, thousandsSep string) string {]
ff=格式化数值
thousandsSep=千位分隔符
decPoint=小数点分隔符
decimals=小数点个数
number=数值

[func Shuffle(str string) string {]
ff=随机打散字符
str=文本

[func HideStr(str string, percent int, hide string) string {]
ff=替换中间字符
hide=替换符
percent=替换百分比
str=文本

[func Nl2Br(str string, isXhtml ...bool) string {]
ff=替换换行符
isXhtml=是否html
str=文本

[func WordWrap(str string, width int, br string) string {]
ff=按字符数量换行
br=换行符
width=字符数
str=文本
