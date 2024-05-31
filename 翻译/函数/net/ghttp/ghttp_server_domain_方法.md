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

[func (s *Server) Domain(domains string) *Domain {]
ff=创建域名路由
domains=域名

[func (d *Domain) BindHandler(pattern string, handler interface{}) {]
ff=X绑定
handler=处理函数
pattern=路由规则

[func (d *Domain) BindObject(pattern string, obj interface{}, methods ...string) {]
ff=绑定对象
methods=方法名
obj=处理对象
pattern=路由规则

[func (d *Domain) BindObjectMethod(pattern string, obj interface{}, method string) {]
ff=绑定对象方法
method=方法
obj=处理对象
pattern=路由规则

[func (d *Domain) BindObjectRest(pattern string, obj interface{}) {]
ff=绑定RESTfulAPI对象
obj=处理对象
pattern=路由规则

[func (d *Domain) BindHookHandler(pattern string, hook HookName, handler HandlerFunc) {]
ff=绑定Hook
handler=处理函数
hook=触发时机
pattern=路由规则

[func (d *Domain) BindHookHandlerByMap(pattern string, hookMap map#左中括号#HookName#右中括号#HandlerFunc) {]
ff=绑定HookMap
hookMap=HookMap
pattern=路由规则

[func (d *Domain) BindStatusHandler(status int, handler HandlerFunc) {]
ff=绑定状态码中间件
handler=处理函数
status=状态码

[func (d *Domain) BindStatusHandlerByMap(handlerMap map#左中括号#int#右中括号#HandlerFunc) {]
ff=绑定状态码中间件Map
handlerMap=中间件Map

[func (d *Domain) BindMiddleware(pattern string, handlers ...HandlerFunc) {]
ff=绑定中间件
handlers=处理函数
pattern=路由规则

[func (d *Domain) BindMiddlewareDefault(handlers ...HandlerFunc) {]
ff=绑定默认中间件
handlers=处理函数

[func (d *Domain) Use(handlers ...HandlerFunc) {]
ff=Use别名
handlers=处理函数
