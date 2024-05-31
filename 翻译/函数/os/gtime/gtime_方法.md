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

[func Timestamp() int64 {]
ff=取时间戳秒

[func TimestampMilli() int64 {]
ff=取时间戳毫秒

[func TimestampMicro() int64 {]
ff=取时间戳微秒

[func TimestampNano() int64 {]
ff=取时间戳纳秒

[func TimestampStr() string {]
ff=取文本时间戳秒

[func TimestampMilliStr() string {]
ff=取文本时间戳毫秒

[func TimestampMicroStr() string {]
ff=取文本时间戳微秒

[func TimestampNanoStr() string {]
ff=取文本时间戳纳秒

[func Datetime() string {]
ff=取当前日期时间

[func ISO8601() string {]
ff=取当前日期时间ISO8601

[func RFC822() string {]
ff=取当前日期时间RFC822

[func StrToTime(str string, format ...string) (*Time, error) {]
ff=转换文本
format=格式
str=文本时间

[func ConvertZone(strTime string, toZone string, fromZone ...string) (*Time, error) {]
ff=转换时区
fromZone=旧时区
toZone=新时区
strTime=文本时间

[func StrToTimeFormat(str string, format string) (*Time, error) {]
ff=StrToTimeFormat别名

[func StrToTimeLayout(str string, layout string) (*Time, error) {]
ff=转换文本Layout
layout=格式
str=文本时间

[func ParseTimeFromContent(content string, format ...string) *Time {]
ff=解析文本
format=格式
content=文本

[func ParseDuration(s string) (duration time.Duration, err error) {]
ff=文本取时长
err=错误
duration=纳秒
s=文本

[func FuncCost(f func()) time.Duration {]
ff=取函数执行时长
f=执行函数
