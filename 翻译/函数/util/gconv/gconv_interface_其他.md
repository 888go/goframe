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

[Val() interface{}]
qm=取值
cz=Val() interface{} yx:true

[Bool() bool]
qm=取布尔
cz=Bool() bool yx:true

[Int64() int64]
qm=取整数64位
cz=Int64() int64 yx:true

[Uint64() uint64]
qm=取正整数64位
cz=Uint64() uint64 yx:true

[Float32() float32]
qm=取小数32位
cz=Float32() float32 yx:true

[Float64() float64]
qm=取小数64位
cz=Float64() float64 yx:true

[Bytes() #左中括号##右中括号#byte]
qm=取字节集
cz=Bytes() []byte yx:true

[Interfaces() #左中括号##右中括号#interface{}]
qm=取any切片
cz=Interfaces() []interface{} yx:true

[Floats() #左中括号##右中括号#float64]
qm=取小数切片
cz=Floats() []float64 yx:true

[Ints() #左中括号##右中括号#int]
qm=取整数切片
cz=Ints() []int yx:true

[Strings() #左中括号##右中括号#string]
qm=取文本切片
cz=Strings() []string yx:true

[Uints() #左中括号##右中括号#uint]
qm=取正整数切片
cz=Uints() []uint yx:true

[MapStrAny() map#左中括号#string#右中括号#interface{}]
qm=取MapStrAny
cz=MapStrAny() map[string]interface{} yx:true

[Set(value interface{}) (old interface{})]
qm=设置值
cz=Set( yx:true

[GTime(format ...string) *gtime.Time]
qm=取gtime时间类
cz=GTime( yx:true
