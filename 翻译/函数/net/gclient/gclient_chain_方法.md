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

[func (c *Client) Prefix(prefix string) *Client {]
ff=Url前缀
prefix=前缀

[func (c *Client) Header(m map#左中括号#string#右中括号#string) *Client {]
ff=协议头
m=map协议头

[func (c *Client) HeaderRaw(headers string) *Client {]
ff=原始协议头
headers=原始协议头

[func (c *Client) Cookie(m map#左中括号#string#右中括号#string) *Client {]
m=MapCookie

[func (c *Client) ContentType(contentType string) *Client {]
ff=内容类型

[func (c *Client) ContentJson() *Client {]
ff=内容类型json

[func (c *Client) ContentXml() *Client {]
ff=内容类型xml

[func (c *Client) Timeout(t time.Duration) *Client {]
ff=超时
t=时长

[func (c *Client) BasicAuth(user, pass string) *Client {]
ff=账号密码
pass=密码
user=账号

[func (c *Client) Retry(retryCount int, retryInterval time.Duration) *Client {]
ff=重试与间隔
retryInterval=重试间隔时长
retryCount=重试次数

[func (c *Client) Proxy(proxyURL string) *Client {]
ff=代理
proxyURL=代理地址

[func (c *Client) RedirectLimit(redirectLimit int) *Client {]
ff=重定向次数限制
redirectLimit=次数

[func (c *Client) NoUrlEncode() *Client {]
ff=请求参数禁止URL编码
