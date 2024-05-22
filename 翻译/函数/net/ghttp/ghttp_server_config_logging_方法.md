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

[func (s *Server) SetLogPath(path string) error {]
ff=设置日志存储目录
path=目录

[func (s *Server) SetLogger(logger *glog.Logger) {]
ff=设置日志记录器
logger=日志记录器

[func (s *Server) Logger() *glog.Logger {]
ff=Logger别名

[func (s *Server) SetLogLevel(level string) {]
ff=设置日志开启记录等级
level=等级

[func (s *Server) SetLogStdout(enabled bool) {]
ff=设置日志开启输出到CMD
enabled=开启

[func (s *Server) SetAccessLogEnabled(enabled bool) {]
ff=设置日志开启访客记录
enabled=开启

[func (s *Server) SetErrorLogEnabled(enabled bool) {]
ff=设置日志开启错误记录
enabled=开启

[func (s *Server) SetErrorStack(enabled bool) {]
ff=设置日志开启错误堆栈记录
enabled=开启

[func (s *Server) GetLogPath() string {]
ff=取日志存储目录

[func (s *Server) IsAccessLogEnabled() bool {]
ff=日志访客记录是否已开启

[func (s *Server) IsErrorLogEnabled() bool {]
ff=日志错误记录是否已开启
