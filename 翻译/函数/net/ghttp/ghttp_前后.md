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

[RegRule string]
qm=正则路由规则
cz=RegRule string

[RegNames #左中括号##右中括号#string]
qm=路由参数名称
cz=RegNames []string

[Server string]
qm=服务器名称
cz=Server string

[Address string]
qm=监听地址
cz=Address string

[Middleware string]
qm=中间件名称
cz=Middleware string

[Route string]
qm=路由URI
cz=Route string

[IsServiceHandler bool]
qm=是否为服务处理器
cz=IsServiceHandler bool

[Name string]
qm=处理器名称
cz=Name string

[Info handlerFuncInfo]
qm=处理器函数信息
cz=Info handlerFuncInfo

[InitFunc HandlerFunc]
qm=初始化回调函数
cz=InitFunc HandlerFunc

[ShutFunc HandlerFunc]
qm=关闭回调函数
cz=ShutFunc HandlerFunc

[Middleware #左中括号##右中括号#HandlerFunc]
qm=中间件切片
cz=Middleware []HandlerFunc

[HookName HookName]
qm=Hook名称
cz=HookName HookName

[Router *Router]
qm=路由
cz=Router *

[Source string]
qm=注册来源
cz=Source string

[Values map#左中括号#string#右中括号#string]
qm=路由值
cz=Values map[string]string

[ServerStatus = int]
qm=服务状态
cz=ServerStatus #等号# int

[HookName string]
qm=Hook名称
cz=HookName string

[HandlerType string]
qm=路由处理器类型
cz=HandlerType string

[FreePortAddress = ":0"]
qm=空闲端口地址
cz=FreePortAddress #等号# ":0"

[ErrNeedJsonBody = gerror.NewWithOption(gerror.Option{]
qm=ERR请求体必须json格式
cz=ErrNeedJsonBody #等号#
