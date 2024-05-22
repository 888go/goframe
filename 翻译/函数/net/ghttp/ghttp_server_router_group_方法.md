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

[func (s *Server) Group(prefix string, groups ...func(group *RouterGroup)) *RouterGroup {]
ff=创建分组路由
groups=分组函数
group=分组路由
prefix=分组前缀

[func (d *Domain) Group(prefix string, groups ...func(group *RouterGroup)) *RouterGroup {]
ff=创建分组路由
groups=分组函数
group=分组路由
prefix=分组前缀

[func (g *RouterGroup) Group(prefix string, groups ...func(group *RouterGroup)) *RouterGroup {]
ff=创建分组路由
groups=分组函数
group=分组路由
prefix=分组前缀

[func (g *RouterGroup) Clone() *RouterGroup {]
ff=取副本

[func (g *RouterGroup) Bind(handlerOrObject ...interface{}) *RouterGroup {]
ff=X绑定
handlerOrObject=处理对象

[func (g *RouterGroup) ALL(pattern string, object interface{}, params ...interface{}) *RouterGroup {]
ff=绑定所有类型
params=额外参数
object=处理函数
pattern=路由规则

[func (g *RouterGroup) ALLMap(m map#左中括号#string#右中括号#interface{}) {]
ff=绑定所有类型Map

[func (g *RouterGroup) Map(m map#左中括号#string#右中括号#interface{}) {]
ff=绑定Map

[func (g *RouterGroup) GET(pattern string, object interface{}, params ...interface{}) *RouterGroup {]
ff=绑定GET
params=额外参数
object=处理函数
pattern=路由规则

[func (g *RouterGroup) PUT(pattern string, object interface{}, params ...interface{}) *RouterGroup {]
ff=绑定PUT
params=额外参数
object=处理函数
pattern=路由规则

[func (g *RouterGroup) POST(pattern string, object interface{}, params ...interface{}) *RouterGroup {]
ff=绑定POST
params=额外参数
object=处理函数
pattern=路由规则

[func (g *RouterGroup) DELETE(pattern string, object interface{}, params ...interface{}) *RouterGroup {]
ff=绑定DELETE
params=额外参数
object=处理函数
pattern=路由规则

[func (g *RouterGroup) PATCH(pattern string, object interface{}, params ...interface{}) *RouterGroup {]
ff=绑定PATCH
params=额外参数
object=处理函数
pattern=路由规则

[func (g *RouterGroup) HEAD(pattern string, object interface{}, params ...interface{}) *RouterGroup {]
ff=绑定HEAD
params=额外参数
object=处理函数
pattern=路由规则

[func (g *RouterGroup) CONNECT(pattern string, object interface{}, params ...interface{}) *RouterGroup {]
ff=绑定CONNECT
params=额外参数
object=处理函数
pattern=路由规则

[func (g *RouterGroup) OPTIONS(pattern string, object interface{}, params ...interface{}) *RouterGroup {]
ff=绑定OPTIONS
params=额外参数
object=处理函数
pattern=路由规则

[func (g *RouterGroup) TRACE(pattern string, object interface{}, params ...interface{}) *RouterGroup {]
ff=绑定TRACE
params=额外参数
object=处理函数
pattern=路由规则

[func (g *RouterGroup) REST(pattern string, object interface{}) *RouterGroup {]
ff=绑定RESTfulAPI对象
object=处理对象
pattern=路由规则

[func (g *RouterGroup) Hook(pattern string, hook HookName, handler HandlerFunc) *RouterGroup {]
ff=绑定Hook
handler=处理函数
hook=触发时机
pattern=路由规则

[func (g *RouterGroup) Middleware(handlers ...HandlerFunc) *RouterGroup {]
ff=绑定中间件
handlers=处理函数
