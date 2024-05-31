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

[func (s *Server) SetSessionMaxAge(ttl time.Duration) {]
ff=设置Session最大存活时长
ttl=时长

[func (s *Server) SetSessionIdName(name string) {]
ff=设置SessionID名称
name=名称

[func (s *Server) SetSessionStorage(storage gsession.Storage) {]
ff=设置Session存储对象
storage=Session存储对象

[func (s *Server) SetSessionCookieOutput(enabled bool) {]
ff=设置SessionID输出到Cookie
enabled=开启

[func (s *Server) SetSessionCookieMaxAge(maxAge time.Duration) {]
ff=设置SessionCookie存活时长
maxAge=最大时长

[func (s *Server) GetSessionMaxAge() time.Duration {]
ff=取Session最大存活时长

[func (s *Server) GetSessionIdName() string {]
ff=取SessionID名称
[func (s *Server) GetSessionCookieMaxAge() time.Duration {]
ff=取SessionCookie存活时长
