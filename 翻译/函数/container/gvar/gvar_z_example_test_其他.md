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

[// *gvar.Var(11) "fisrt hello"]
th=*泛型类.Var(11) "fisrt hello"
cz=*gvar.Var(11) "fisrt hello"

[// *gtime.Time(19) "2021-11-11 00:00:00"]
th=*时间类.Time(19) "2021-11-11 00:00:00"
cz=*gtime.Time(19) "2021-11-11 00:00:00"

[// #左中括号##右中括号#gvar_test.tartget(2) #左中括号#]
th=[]泛型类_test.tartget(2) [
cz=[]gvar_test.tartget(2) [

[// gvar_test.tartget(2) {]
th=泛型类_test.tartget(2) {
cz=gvar_test.tartget(2) {

[// gvar_test.Student(3) {]
th=泛型类_test.Student(3)
cz=gvar_test.Student(3)

[// Id: *gvar.Var(1) "1",]
th=*泛型类.Var(1) "1",
cz=*gvar.Var(1) "1",

[// Name: *gvar.Var(4) "john",]
th=*泛型类.Var(4) "john",
cz=*gvar.Var(4) "john",

[// Scores: *gvar.Var(11) "#左中括号#100,99,98#右中括号#",]
th=*泛型类.Var(11) "[100,99,98]",
cz=*gvar.Var(11) "[100,99,98]",
