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

[func DefaultOptions() TimerOptions {]
ff=取单例对象

[func SetTimeout(ctx context.Context, delay time.Duration, job JobFunc) {]
ff=SetTimeout别名

[func SetInterval(ctx context.Context, interval time.Duration, job JobFunc) {]
ff=SetInterval别名

[func Add(ctx context.Context, interval time.Duration, job JobFunc) *Entry {]
ff=加入循环任务
job=任务函数
interval=间隔时长
ctx=上下文

[func AddEntry(ctx context.Context, interval time.Duration, job JobFunc, isSingleton bool, times int, status int) *Entry {]
ff=加入详细循环任务
status=任务状态
times=次数
isSingleton=是否单例模式
job=任务函数
interval=间隔时长
ctx=上下文

[func AddSingleton(ctx context.Context, interval time.Duration, job JobFunc) *Entry {]
ff=加入单例循环任务
job=任务函数
interval=间隔时长
ctx=上下文

[func AddOnce(ctx context.Context, interval time.Duration, job JobFunc) *Entry {]
ff=加入单次任务
job=任务函数
interval=间隔时长
ctx=上下文

[func AddTimes(ctx context.Context, interval time.Duration, times int, job JobFunc) *Entry {]
ff=加入指定次数任务
job=任务函数
times=次数
interval=间隔时长
ctx=上下文

[func DelayAdd(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc) {]
ff=延时加入循环任务
job=任务函数
interval=间隔时长
delay=延时加入
ctx=上下文

[func DelayAddEntry(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc, isSingleton bool, times int, status int) {]
ff=延时加入详细循环任务
status=任务状态
times=次数
isSingleton=是否单例模式
job=任务函数
interval=间隔时长
delay=延时加入
ctx=上下文

[func DelayAddSingleton(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc) {]
ff=延时加入单例循环任务
job=任务函数
interval=间隔时长
delay=延时加入
ctx=上下文

[func DelayAddOnce(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc) {]
ff=延时加入单次任务
job=任务函数
interval=间隔时长
delay=延时加入
ctx=上下文

[func DelayAddTimes(ctx context.Context, delay time.Duration, interval time.Duration, times int, job JobFunc) {]
ff=延时加入指定次数任务
job=任务函数
times=次数
interval=间隔时长
delay=延时加入
ctx=上下文
