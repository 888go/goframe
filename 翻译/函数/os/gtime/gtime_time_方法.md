# 备注开始
# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 如:
# //ff:取文本

# **_package.md 文件备注:
# bm= 包名,更换新的包名称 
# 如: 
# package gin //bm:gin类

# **_其他.md 文件备注:
# qm= 前面,跳转到前面进行重命名.文档内如果有多个相同的,会一起重命名.
# hm= 后面,跳转到后面进行重命名.文档内如果有多个相同的,会一起重命名.
# cz= 查找,配合前面/后面使用,
# 如:
# type Regexp struct {//qm:正则 cz:Regexp struct
#
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
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

[func New(param ...interface{}) *Time {]
ff=创建
param=参数

[func Now() *Time {]
ff=创建并按当前时间

[func NewFromTime(t time.Time) *Time {]
ff=创建并按Time

[func NewFromStr(str string) *Time {]
ff=创建并从文本
str=文本时间

[func NewFromStrFormat(str string, format string) *Time {]
ff=创建并按给定格式文本
format=格式
str=文本时间

[func NewFromStrLayout(str string, layout string) *Time {]
ff=创建并按Layout格式文本
layout=格式
str=文本时间

[func NewFromTimeStamp(timestamp int64) *Time {]
ff=创建并从时间戳
timestamp=时间戳

[func (t *Time) Timestamp() int64 {]
ff=取时间戳秒

[func (t *Time) TimestampMilli() int64 {]
ff=取时间戳毫秒

[func (t *Time) TimestampMicro() int64 {]
ff=取时间戳微秒

[func (t *Time) TimestampNano() int64 {]
ff=取时间戳纳秒

[func (t *Time) TimestampStr() string {]
ff=取文本时间戳秒

[func (t *Time) TimestampMilliStr() string {]
ff=取文本时间戳毫秒

[func (t *Time) TimestampMicroStr() string {]
ff=取文本时间戳微秒

[func (t *Time) TimestampNanoStr() string {]
ff=取文本时间戳纳秒

[func (t *Time) Month() int {]
ff=取月份

[func (t *Time) Second() int {]
ff=取秒

[func (t *Time) Millisecond() int {]
ff=取毫秒

[func (t *Time) Microsecond() int {]
ff=取微秒

[func (t *Time) Nanosecond() int {]
ff=取纳秒

[func (t *Time) Clone() *Time {]
ff=取副本

[func (t *Time) Add(d time.Duration) *Time {]
ff=增加时长
d=时长

[func (t *Time) AddStr(duration string) (*Time, error) {]
ff=增加文本时长
duration=时长

[func (t *Time) UTC() *Time {]
ff=取UTC时区

[func (t *Time) ISO8601() string {]
ff=取文本时间ISO8601

[func (t *Time) RFC822() string {]
ff=取文本时间RFC822

[func (t *Time) AddDate(years int, months int, days int) *Time {]
ff=增加时间
days=日
months=月
years=年

[func (t *Time) Round(d time.Duration) *Time {]
ff=向上舍入
d=时长

[func (t *Time) Truncate(d time.Duration) *Time {]
ff=向下舍入
d=时长

[func (t *Time) Equal(u *Time) bool {]
ff=是否相等

[func (t *Time) Before(u *Time) bool {]
ff=是否之前

[func (t *Time) After(u *Time) bool {]
ff=是否之后

[func (t *Time) Sub(u *Time) time.Duration {]
ff=取纳秒时长

[func (t *Time) StartOfMinute() *Time {]
ff=取副本忽略秒

[func (t *Time) StartOfHour() *Time {]
ff=取副本忽略分钟秒

[func (t *Time) StartOfDay() *Time {]
ff=取副本忽略小时分钟秒

[func (t *Time) StartOfWeek() *Time {]
ff=取副本周第一天

[func (t *Time) StartOfMonth() *Time {]
ff=取副本月第一天

[func (t *Time) StartOfQuarter() *Time {]
ff=取副本季度第一天

[func (t *Time) StartOfHalf() *Time {]
ff=取副本半年第一天

[func (t *Time) StartOfYear() *Time {]
ff=取副本年第一天

[func (t *Time) EndOfMinute(withNanoPrecision ...bool) *Time {]
ff=取副本59秒
withNanoPrecision=纳秒精度

[func (t *Time) EndOfHour(withNanoPrecision ...bool) *Time {]
ff=取副本59分59秒
withNanoPrecision=纳秒精度

[func (t *Time) EndOfDay(withNanoPrecision ...bool) *Time {]
ff=取副本23点59分59秒
withNanoPrecision=纳秒精度

[func (t *Time) EndOfWeek(withNanoPrecision ...bool) *Time {]
ff=取副本周末23点59分59秒
withNanoPrecision=纳秒精度

[func (t *Time) EndOfMonth(withNanoPrecision ...bool) *Time {]
ff=取副本月末23点59分59秒
withNanoPrecision=纳秒精度

[func (t *Time) EndOfQuarter(withNanoPrecision ...bool) *Time {]
ff=取副本季末23点59分59秒
withNanoPrecision=纳秒精度

[func (t *Time) EndOfHalf(withNanoPrecision ...bool) *Time {]
ff=取副本半年末23点59分59秒
withNanoPrecision=纳秒精度

[func (t *Time) EndOfYear(withNanoPrecision ...bool) *Time {]
ff=取副本年末23点59分59秒
withNanoPrecision=纳秒精度
