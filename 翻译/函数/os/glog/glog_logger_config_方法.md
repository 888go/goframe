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

[func DefaultConfig() Config {]
ff=生成默认配置

[func (l *Logger) GetConfig() Config {]
ff=取配置项

[func (l *Logger) SetConfig(config Config) error {]
ff=设置配置项
config=配置项

[func (l *Logger) SetConfigWithMap(m map#左中括号#string#右中括号#interface{}) error {]
ff=设置配置Map

[func (l *Logger) SetDebug(debug bool) {]
ff=设置debug
debug=开启

[func (l *Logger) SetAsync(enabled bool) {]
ff=设置异步输出
enabled=开启

[func (l *Logger) SetFlags(flags int) {]
ff=设置额外标识
flags=标识

[func (l *Logger) GetFlags() int {]
ff=取标识

[func (l *Logger) SetStack(enabled bool) {]
ff=设置堆栈跟踪
enabled=开启

[func (l *Logger) SetStackSkip(skip int) {]
ff=设置堆栈偏移量
skip=偏移量

[func (l *Logger) SetStackFilter(filter string) {]
ff=设置堆栈过滤
filter=过滤器

[func (l *Logger) SetCtxKeys(keys ...interface{}) {]
ff=设置上下文名称
keys=名称

[func (l *Logger) GetCtxKeys() #左中括号##右中括号#interface{} {]
ff=取上下文名称

[func (l *Logger) SetWriter(writer io.Writer) {]
ff=设置Writer

[func (l *Logger) GetWriter() io.Writer {]
ff=取Writer

[func (l *Logger) SetPath(path string) error {]
ff=设置文件路径
path=文件路径

[func (l *Logger) GetPath() string {]
ff=取文件路径

[func (l *Logger) SetFile(pattern string) {]
ff=设置文件名格式
pattern=文件名格式

[func (l *Logger) SetTimeFormat(timeFormat string) {]
ff=设置时间格式
timeFormat=时间格式

[func (l *Logger) SetStdoutPrint(enabled bool) {]
ff=设置是否同时输出到终端
enabled=开启

[func (l *Logger) SetHeaderPrint(enabled bool) {]
ff=设置是否输出头信息
enabled=开启

[func (l *Logger) SetLevelPrint(enabled bool) {]
ff=设置是否输出级别
enabled=开启

[func (l *Logger) SetPrefix(prefix string) {]
ff=设置前缀
prefix=前缀

[func (l *Logger) SetHandlers(handlers ...Handler) {]
ff=设置中间件
handlers=处理函数

[func (l *Logger) SetWriterColorEnable(enabled bool) {]
ff=设置文件是否输出颜色
enabled=开启

[func (l *Logger) SetStdoutColorDisabled(disabled bool) {]
ff=设置关闭终端颜色输出
disabled=关闭
