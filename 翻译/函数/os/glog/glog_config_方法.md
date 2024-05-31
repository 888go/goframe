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

[func SetConfig(config Config) error {]
ff=设置配置项
config=配置项

[func SetConfigWithMap(m map#左中括号#string#右中括号#interface{}) error {]
ff=设置配置Map

[func SetPath(path string) error {]
ff=设置文件路径
path=文件路径

[func GetPath() string {]
ff=取文件路径

[func SetFile(pattern string) {]
ff=设置文件名格式
pattern=文件名格式

[func SetLevel(level int) {]
ff=设置级别
level=级别

[func GetLevel() int {]
ff=取级别

[func SetWriter(writer io.Writer) {]
ff=设置Writer

[func GetWriter() io.Writer {]
ff=取Writer

[func SetDebug(debug bool) {]
ff=设置debug
debug=开启

[func SetAsync(enabled bool) {]
ff=设置异步输出
enabled=开启

[func SetStdoutPrint(enabled bool) {]
ff=设置是否同时输出到终端
enabled=开启

[func SetHeaderPrint(enabled bool) {]
ff=设置是否输出头信息
enabled=开启

[func SetPrefix(prefix string) {]
ff=设置前缀
prefix=前缀

[func SetFlags(flags int) {]
ff=设置额外标识
flags=标识

[func GetFlags() int {]
ff=取标识

[func SetCtxKeys(keys ...interface{}) {]
ff=设置上下文名称
keys=名称

[func GetCtxKeys() #左中括号##右中括号#interface{} {]
ff=取上下文名称

[func PrintStack(ctx context.Context, skip ...int) {]
ff=输出堆栈信息
skip=偏移量
ctx=上下文

[func GetStack(skip ...int) string {]
ff=取堆栈信息
skip=偏移量

[func SetStack(enabled bool) {]
ff=设置堆栈跟踪
enabled=开启

[func SetLevelStr(levelStr string) error {]
ff=设置文本级别
levelStr=级别

[func SetLevelPrefix(level int, prefix string) {]
ff=设置级别前缀
prefix=前缀
level=级别

[func SetLevelPrefixes(prefixes map#左中括号#int#右中括号#string) {]
ff=设置级别前缀Map
prefixes=前缀Map

[func GetLevelPrefix(level int) string {]
ff=取级别前缀
level=级别

[func SetHandlers(handlers ...Handler) {]
ff=设置中间件
handlers=处理函数

[func SetWriterColorEnable(enabled bool) {]
ff=设置文件是否输出颜色
enabled=开启
