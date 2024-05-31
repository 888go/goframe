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

[func ComparatorString(a, b interface{}) int {]
ff=比较文本

[func ComparatorInt(a, b interface{}) int {]
ff=比较整数

[func ComparatorInt8(a, b interface{}) int {]
ff=比较整数8位

[func ComparatorInt16(a, b interface{}) int {]
ff=比较整数16位

[func ComparatorInt32(a, b interface{}) int {]
ff=比较整数32位

[func ComparatorInt64(a, b interface{}) int {]
ff=比较整数64位

[func ComparatorUint(a, b interface{}) int {]
ff=比较正整数

[func ComparatorUint8(a, b interface{}) int {]
ff=比较正整数8位

[func ComparatorUint16(a, b interface{}) int {]
ff=比较正整数16位

[func ComparatorUint32(a, b interface{}) int {]
ff=比较正整数32位

[func ComparatorUint64(a, b interface{}) int {]
ff=比较正整数64位

[func ComparatorFloat32(a, b interface{}) int {]
ff=比较小数32位

[func ComparatorFloat64(a, b interface{}) int {]
ff=比较小数64位

[func ComparatorByte(a, b interface{}) int {]
ff=比较字节

[func ComparatorRune(a, b interface{}) int {]
ff=比较字符
[func ComparatorTime(a, b interface{}) int {]
ff=比较时间
