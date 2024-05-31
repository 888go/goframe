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

[func PtrAny(any interface{}) *interface{} {]
ff=取any指针
any=值

[func PtrString(any interface{}) *string {]
ff=取文本指针
any=值

[func PtrBool(any interface{}) *bool {]
ff=取布尔指针
any=值

[func PtrInt(any interface{}) *int {]
ff=取整数指针
any=值

[func PtrInt8(any interface{}) *int8 {]
ff=取整数8位指针
any=值

[func PtrInt16(any interface{}) *int16 {]
ff=取整数16位指针
any=值

[func PtrInt32(any interface{}) *int32 {]
ff=取整数32位指针
any=值

[func PtrInt64(any interface{}) *int64 {]
ff=取整数64位指针
any=值

[func PtrUint(any interface{}) *uint {]
ff=取正整数指针
any=值

[func PtrUint8(any interface{}) *uint8 {]
ff=取正整数8位指针
any=值

[func PtrUint16(any interface{}) *uint16 {]
ff=取正整数16位指针
any=值

[func PtrUint32(any interface{}) *uint32 {]
ff=取正整数32位指针
any=值

[func PtrUint64(any interface{}) *uint64 {]
ff=取正整数64位指针
any=值

[func PtrFloat32(any interface{}) *float32 {]
ff=取小数32位指针
any=值

[func PtrFloat64(any interface{}) *float64 {]
ff=取小数64位指针
any=值
