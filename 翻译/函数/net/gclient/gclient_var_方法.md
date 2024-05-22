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

[func (c *Client) GetVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {]
ff=Get泛型类
data=参数
ctx=上下文

[func (c *Client) PutVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {]
ff=Put泛型类
data=参数
ctx=上下文

[func (c *Client) PostVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {]
ff=Post泛型类
data=参数
ctx=上下文

[func (c *Client) DeleteVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {]
ff=Delete泛型类
data=参数
ctx=上下文

[func (c *Client) HeadVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {]
ff=Head泛型类
data=参数
ctx=上下文

[func (c *Client) PatchVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {]
ff=Patch泛型类
data=参数
ctx=上下文

[func (c *Client) ConnectVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {]
ff=Connect泛型类
data=参数
ctx=上下文

[func (c *Client) OptionsVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {]
ff=Options泛型类
data=参数
ctx=上下文

[func (c *Client) TraceVar(ctx context.Context, url string, data ...interface{}) *gvar.Var {]
ff=Trace泛型类
data=参数
ctx=上下文

[func (c *Client) RequestVar(ctx context.Context, method string, url string, data ...interface{}) *gvar.Var {]
ff=请求泛型类
data=参数
method=方法
ctx=上下文
