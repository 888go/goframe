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

[func (r *Request) WebSocket() (*WebSocket, error) {]
ff=升级为websocket请求

[func (r *Request) Exit() {]
ff=退出当前

[func (r *Request) ExitAll() {]
ff=退出全部

[func (r *Request) ExitHook() {]
ff=退出Hook

[func (r *Request) IsExited() bool {]
ff=是否已退出

[func (r *Request) GetHeader(key string) string {]
ff=取协议头值
key=名称

[func (r *Request) GetHost() string {]
ff=取主机名

[func (r *Request) IsFileRequest() bool {]
ff=是否为文件请求

[func (r *Request) IsAjaxRequest() bool {]
ff=是否为AJAX请求

[func (r *Request) GetClientIp() string {]
ff=取客户端IP地址

[func (r *Request) GetRemoteIp() string {]
ff=取远程IP地址

[func (r *Request) GetUrl() string {]
ff=取URL

[func (r *Request) GetSessionId() string {]
ff=取SessionId

[func (r *Request) GetReferer() string {]
ff=取引用来源URL

[func (r *Request) GetError() error {]
ff=取错误信息

[func (r *Request) SetError(err error) {]
ff=设置错误信息
err=错误

[func (r *Request) ReloadParam() {]
ff=重载请求参数

[func (r *Request) GetHandlerResponse() interface{} {]
ff=取响应对象及错误信息

[func (r *Request) GetServeHandler() *HandlerItemParsed {]
ff=取路由解析对象
