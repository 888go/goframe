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

[func New(value interface{}, safe ...bool) *Var {]
ff=创建
safe=并发安全
value=值

[func (v *Var) Copy() *Var {]
ff=深拷贝

[func (v *Var) Clone() *Var {]
ff=浅拷贝

[func (v *Var) Set(value interface{}) (old interface{}) {]
ff=设置值
yx=true

[func (v *Var) Val() interface{} {]
ff=取值
yx=true

[func (v *Var) Bytes() #左中括号##右中括号#byte {]
ff=取字节集
yx=true

[func (v *Var) Bool() bool {]
ff=取布尔
yx=true

[func (v *Var) Int() int {]
ff=取整数

[func (v *Var) Int8() int8 {]
ff=取整数8位

[func (v *Var) Int16() int16 {]
ff=取整数16位

[func (v *Var) Int32() int32 {]
ff=取整数32位

[func (v *Var) Int64() int64 {]
ff=取整数64位
yx=true

[func (v *Var) Uint() uint {]
ff=取正整数

[func (v *Var) Uint8() uint8 {]
ff=取正整数8位

[func (v *Var) Uint16() uint16 {]
ff=取正整数16位

[func (v *Var) Uint32() uint32 {]
ff=取正整数32位

[func (v *Var) Uint64() uint64 {]
ff=取正整数64位
yx=true

[func (v *Var) Float32() float32 {]
ff=取小数32位
yx=true

[func (v *Var) Float64() float64 {]
ff=取小数64位
yx=true

[func (v *Var) Time(format ...string) time.Time {]
ff=取时间类
format=格式

[func (v *Var) Duration() time.Duration {]
ff=取时长

[func (v *Var) GTime(format ...string) *gtime.Time {]
ff=取gtime时间类
yx=true
