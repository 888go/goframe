
# <翻译开始>
Server struct {
X服务
<跳到行首>
# <翻译结束>

# <翻译开始>
Router struct {
X路由
<跳到行首>
# <翻译结束>

# <翻译开始>
RegRule string
X正则路由规则
<跳到行首>
# <翻译结束>

# <翻译开始>
RegNames []string
X路由参数名称
<跳到行首>
# <翻译结束>

# <翻译开始>
Server string
X服务器名称
<跳到行首>
# <翻译结束>

# <翻译开始>
Address string
X监听地址
<跳到行首>
# <翻译结束>

# <翻译开始>
Middleware string
X中间件名称
<跳到行首>
# <翻译结束>

# <翻译开始>
Route string
X路由URI
<跳到行首>
# <翻译结束>

# <翻译开始>
IsServiceHandler bool
X是否为服务处理器
<跳到行首>
# <翻译结束>


# <翻译开始>
HandlerItem struct {
X路由处理函数
<跳到行首>
# <翻译结束>

# <翻译开始>
Name string
X处理器名称
<跳到行首>
# <翻译结束>

# <翻译开始>
Info handlerFuncInfo
X处理器函数信息
<跳到行首>
# <翻译结束>

# <翻译开始>
InitFunc HandlerFunc
X初始化回调函数
<跳到行首>
# <翻译结束>

# <翻译开始>
ShutFunc HandlerFunc
X关闭回调函数
<跳到行首>
# <翻译结束>

# <翻译开始>
Middleware []HandlerFunc
X中间件切片
<跳到行首>
# <翻译结束>

# <翻译开始>
HookName HookName
Hook名称
<跳到行首>
# <翻译结束>

# <翻译开始>
Router *
X路由
<跳到行首>
# <翻译结束>

# <翻译开始>
Source string
X注册来源
<跳到行首>
# <翻译结束>

# <翻译开始>
HandlerItemParsed struct {
X路由解析
<跳到行首>
# <翻译结束>

# <翻译开始>
Values map[string]string
X路由值
<跳到行首>
# <翻译结束>

# <翻译开始>
ServerStatus = int
X服务状态
<跳到行首>
# <翻译结束>

# <翻译开始>
HookName string
Hook名称
<跳到行首>
# <翻译结束>

# <翻译开始>
HandlerType string
X路由处理器类型
<跳到行首>
# <翻译结束>

# <翻译开始>
FreePortAddress = ":0"
X空闲端口地址
<跳到行首>
# <翻译结束>

# <翻译开始>
ErrNeedJsonBody =
ERR请求体必须json格式
<跳到行首>
# <翻译结束>
