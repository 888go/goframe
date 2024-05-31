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

[func (vs Vars) Strings() (s #左中括号##右中括号#string) {]
ff=取文本数组
yx=true

[func (vs Vars) Interfaces() (s #左中括号##右中括号#interface{}) {]
ff=取any数组
yx=true

[func (vs Vars) Float32s() (s #左中括号##右中括号#float32) {]
ff=取小数32位数组
s=数组

[func (vs Vars) Float64s() (s #左中括号##右中括号#float64) {]
ff=取小数64位数组
s=数组

[func (vs Vars) Ints() (s #左中括号##右中括号#int) {]
ff=取整数数组
yx=true

[func (vs Vars) Int8s() (s #左中括号##右中括号#int8) {]
ff=取整数8位数组
s=数组

[func (vs Vars) Int16s() (s #左中括号##右中括号#int16) {]
ff=取整数16位数组
s=数组

[func (vs Vars) Int32s() (s #左中括号##右中括号#int32) {]
ff=取整数32位数组
s=数组

[func (vs Vars) Int64s() (s #左中括号##右中括号#int64) {]
ff=取整数64位数组
s=数组

[func (vs Vars) Uints() (s #左中括号##右中括号#uint) {]
ff=取正整数数组
yx=true

[func (vs Vars) Uint8s() (s #左中括号##右中括号#uint8) {]
ff=取正整数8位数组
s=数组

[func (vs Vars) Uint16s() (s #左中括号##右中括号#uint16) {]
ff=取正整数16位数组
s=数组

[func (vs Vars) Uint32s() (s #左中括号##右中括号#uint32) {]
ff=取正整数32位数组
s=数组

[func (vs Vars) Uint64s() (s #左中括号##右中括号#uint64) {]
ff=取正整数64位数组
s=数组

[func (vs Vars) Scan(pointer interface{}, mapping ...map#左中括号#string#右中括号#string) error {]
ff=取结构体指针
mapping=名称映射
pointer=结构体指针
