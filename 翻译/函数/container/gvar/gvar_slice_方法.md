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

[func (v *Var) Ints() #左中括号##右中括号#int {]
ff=取整数数组
yx=true

[func (v *Var) Int64s() #左中括号##右中括号#int64 {]
ff=取整数64位数组

[func (v *Var) Uints() #左中括号##右中括号#uint {]
ff=取正整数数组
yx=true

[func (v *Var) Uint64s() #左中括号##右中括号#uint64 {]
ff=取正整数64位数组

[func (v *Var) Floats() #左中括号##右中括号#float64 {]
ff=取小数数组
yx=true

[func (v *Var) Float32s() #左中括号##右中括号#float32 {]
ff=取小数32位数组

[func (v *Var) Float64s() #左中括号##右中括号#float64 {]
ff=取小数64位数组

[func (v *Var) Strings() #左中括号##右中括号#string {]
ff=取文本数组
yx=true

[func (v *Var) Interfaces() #左中括号##右中括号#interface{} {]
ff=取any数组
yx=true

[func (v *Var) Slice() #左中括号##右中括号#interface{} {]
ff=Slice别名

[func (v *Var) Array() #左中括号##右中括号#interface{} {]
ff=Array别名
[func (v *Var) Vars() #左中括号##右中括号#*Var {]
ff=取泛型类数组
