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

[func (l *Logger) Print(ctx context.Context, v ...interface{}) {]
ff=输出
v=值
ctx=上下文

[func (l *Logger) Printf(ctx context.Context, format string, v ...interface{}) {]
ff=输出并格式化
v=值
format=格式
ctx=上下文

[func (l *Logger) Fatal(ctx context.Context, v ...interface{}) {]
ff=输出FATA
v=值
ctx=上下文

[func (l *Logger) Fatalf(ctx context.Context, format string, v ...interface{}) {]
ff=输出并格式化FATA
v=值
format=格式
ctx=上下文

[func (l *Logger) Panic(ctx context.Context, v ...interface{}) {]
ff=输出PANI
v=值
ctx=上下文

[func (l *Logger) Panicf(ctx context.Context, format string, v ...interface{}) {]
ff=输出并格式化PANI
v=值
format=格式
ctx=上下文

[func (l *Logger) Info(ctx context.Context, v ...interface{}) {]
ff=输出INFO
v=值
ctx=上下文

[func (l *Logger) Infof(ctx context.Context, format string, v ...interface{}) {]
ff=输出并格式化INFO
v=值
format=格式
ctx=上下文

[func (l *Logger) Debug(ctx context.Context, v ...interface{}) {]
ff=输出DEBU
v=值
ctx=上下文

[func (l *Logger) Debugf(ctx context.Context, format string, v ...interface{}) {]
ff=输出并格式化DEBU
v=值
format=格式
ctx=上下文

[func (l *Logger) Notice(ctx context.Context, v ...interface{}) {]
ff=输出NOTI
v=值
ctx=上下文

[func (l *Logger) Noticef(ctx context.Context, format string, v ...interface{}) {]
ff=输出并格式化NOTI
v=值
format=格式
ctx=上下文

[func (l *Logger) Warning(ctx context.Context, v ...interface{}) {]
ff=输出WARN
v=值
ctx=上下文

[func (l *Logger) Warningf(ctx context.Context, format string, v ...interface{}) {]
ff=输出并格式化WARN
v=值
format=格式
ctx=上下文

[func (l *Logger) Error(ctx context.Context, v ...interface{}) {]
v=值
ctx=上下文

[func (l *Logger) Errorf(ctx context.Context, format string, v ...interface{}) {]
ff=输出并格式化ERR
v=值
format=格式
ctx=上下文

[func (l *Logger) Critical(ctx context.Context, v ...interface{}) {]
ff=输出CRIT
v=值
ctx=上下文

[func (l *Logger) Criticalf(ctx context.Context, format string, v ...interface{}) {]
ff=输出并格式化CRIT
v=值
format=格式
ctx=上下文
