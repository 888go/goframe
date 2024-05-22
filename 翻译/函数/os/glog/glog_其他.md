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

[Print(ctx context.Context, v ...interface{})]
qm=输出
cz=Print(

[Printf(ctx context.Context, format string, v ...interface{})]
qm=输出并格式化
cz=Printf(

[Debug(ctx context.Context, v ...interface{})]
qm=输出DEBU
cz=Debug(

[Debugf(ctx context.Context, format string, v ...interface{})]
qm=输出并格式化DEBU
cz=Debugf(

[Info(ctx context.Context, v ...interface{})]
qm=输出INFO
cz=Info(

[Infof(ctx context.Context, format string, v ...interface{})]
qm=输出并格式化INFO
cz=Infof(

[Notice(ctx context.Context, v ...interface{})]
qm=输出NOTI
cz=Notice(

[Noticef(ctx context.Context, format string, v ...interface{})]
qm=输出并格式化NOTI
cz=Noticef(

[Warning(ctx context.Context, v ...interface{})]
qm=输出WARN
cz=Warning(

[Warningf(ctx context.Context, format string, v ...interface{})]
qm=输出并格式化WARN
cz=Warningf(

[Errorf(ctx context.Context, format string, v ...interface{})]
qm=输出并格式化ERR
cz=Errorf(

[Critical(ctx context.Context, v ...interface{})]
qm=输出CRIT
cz=Critical(

[Criticalf(ctx context.Context, format string, v ...interface{})]
qm=输出并格式化CRIT
cz=Criticalf(

[Panic(ctx context.Context, v ...interface{})]
qm=输出PANI
cz=Panic(

[Panicf(ctx context.Context, format string, v ...interface{})]
qm=输出并格式化PANI
cz=Panicf(

[Fatal(ctx context.Context, v ...interface{})]
qm=输出FATA
cz=Fatal(

[Fatalf(ctx context.Context, format string, v ...interface{})]
qm=输出并格式化FATA
cz=Fatalf(
