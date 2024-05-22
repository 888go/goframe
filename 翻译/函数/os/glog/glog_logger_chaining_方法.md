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

[func (l *Logger) To(writer io.Writer) *Logger {]
ff=重定向输出

[func (l *Logger) Path(path string) *Logger {]
ff=文件路径
path=文件路径

[func (l *Logger) Cat(category string) *Logger {]
ff=文件分类
category=类名称

[func (l *Logger) File(file string) *Logger {]
ff=文件名格式
file=文件名格式

[func (l *Logger) Level(level int) *Logger {]
ff=级别
level=级别

[func (l *Logger) LevelStr(levelStr string) *Logger {]
ff=文本级别
levelStr=文本级别

[func (l *Logger) Skip(skip int) *Logger {]
ff=堆栈偏移量
skip=偏移量

[func (l *Logger) Stack(enabled bool, skip ...int) *Logger {]
ff=堆栈选项
skip=偏移量
enabled=开启

[func (l *Logger) StackWithFilter(filter string) *Logger {]
ff=堆栈过滤
filter=过滤器

[func (l *Logger) Stdout(enabled ...bool) *Logger {]
ff=是否同时输出到终端
enabled=开启

[func (l *Logger) Header(enabled ...bool) *Logger {]
ff=是否输出头信息
enabled=开启

[func (l *Logger) Line(long ...bool) *Logger {]
ff=是否输出源文件路径与行号
long=开启

[func (l *Logger) Async(enabled ...bool) *Logger {]
ff=是否异步输出
enabled=开启
