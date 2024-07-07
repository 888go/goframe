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

[func Str(haystack string, needle string) string {]
ff=取右边并含关键字
needle=欲寻找的文本
haystack=文本

[func StrEx(haystack string, needle string) string {]
ff=取右边
needle=欲寻找的文本
haystack=文本

[func StrTill(haystack string, needle string) string {]
ff=取左边并含关键字
needle=欲寻找的文本
haystack=文本

[func StrTillEx(haystack string, needle string) string {]
ff=取左边
needle=欲寻找的文本
haystack=文本

[func SubStr(str string, start int, length ...int) (substr string) {]
ff=按长度取文本
substr=返回
length=长度
start=起始位置
str=文本

[func SubStrRune(str string, start int, length ...int) (substr string) {]
ff=按长度取文本Unicode
substr=返回
length=长度
start=起始位置
str=文本

[func StrLimit(str string, length int, suffix ...string) string {]
ff=按长度取左边并带前缀
suffix=后缀
length=长度
str=文本

[func StrLimitRune(str string, length int, suffix ...string) string {]
ff=按长度取左边并带前缀Unicode
suffix=后缀
length=长度
str=文本

[func SubStrFrom(str string, need string) (substr string) {]
ff=SubStrFrom别名

[func SubStrFromEx(str string, need string) (substr string) {]
ff=SubStrFromEx别名

[func SubStrFromR(str string, need string) (substr string) {]
ff=取右边并倒找与含关键字
substr=文本结果
need=欲寻找的文本
str=文本

[func SubStrFromREx(str string, need string) (substr string) {]
ff=取右边并倒找
substr=文本结果
need=欲寻找的文本
str=文本
