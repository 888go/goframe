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

[func NewVar(i interface{}, safe ...bool) *Var {]
ff=X泛型类
safe=并发安全
i=值

[func Wait() {]
ff=Http类等待所有服务完成

[func Dump(values ...interface{}) {]
ff=调试输出
values=值s

[func DumpTo(writer io.Writer, value interface{}, option gutil.DumpOption) {]
ff=调试输出到Writer
option=选项
value=值

[func DumpWithType(values ...interface{}) {]
ff=调试输出并带类型
values=值s

[func DumpWithOption(value interface{}, option gutil.DumpOption) {]
ff=调试输出并带选项
option=选项
value=值s

[func DumpJson(value any) {]
ff=调试输出json

[func Throw(exception interface{}) {]
ff=异常输出
exception=消息

[func Try(ctx context.Context, try func(ctx context.Context)) (err error) {]
ff=异常捕捉
err=错误
try=处理函数
ctx=上下文

[func TryCatch(ctx context.Context, try func(ctx context.Context), catch func(ctx context.Context, exception error)) {]
ff=异常捕捉并带异常处理
catch=异常处理函数
exception=错误
ctx=上下文
try=处理函数

[func IsNil(value interface{}, traceSource ...bool) bool {]
ff=是否为Nil

[func IsEmpty(value interface{}, traceSource ...bool) bool {]
ff=是否为空
traceSource=追踪到源变量
value=值

[func RequestFromCtx(ctx context.Context) *ghttp.Request {]
ff=Http类上下文取请求对象
ctx=上下文
