# 备注开始
# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 如:
# //ff:取文本

# **_package.md 文件备注:
# bm= 包名,更换新的包名称 
# 如: 
# package gin //bm:gin类

# **_其他.md 文件备注:
# qm= 前面,跳转到前面进行重命名.文档内如果有多个相同的,会一起重命名.
# hm= 后面,跳转到后面进行重命名.文档内如果有多个相同的,会一起重命名.
# cz= 查找,配合前面/后面使用,
# 如:
# type Regexp struct {//qm:正则 cz:Regexp struct
#
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
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

[func (c *Client) SetBrowserMode(enabled bool) *Client {]
ff=启用浏览器模式
enabled=启用

[func (c *Client) SetHeader(key, value string) *Client {]
ff=设置协议头
value=值
key=名称

[func (c *Client) SetHeaderMap(m map#左中括号#string#右中括号#string) *Client {]
ff=设置Map协议头
m=map协议头

[func (c *Client) SetAgent(agent string) *Client {]
ff=设置UA
agent=UA值

[func (c *Client) SetContentType(contentType string) *Client {]
ff=设置内容类型
contentType=内容类型

[func (c *Client) SetHeaderRaw(headers string) *Client {]
ff=设置原始协议头
headers=原始协议头

[func (c *Client) SetCookie(key, value string) *Client {]
ff=设置cookie
value=值
key=名称

[func (c *Client) SetCookieMap(m map#左中括号#string#右中括号#string) *Client {]
ff=设置CookieMap
m=MapCookie

[func (c *Client) SetPrefix(prefix string) *Client {]
ff=设置url前缀
prefix=前缀

[func (c *Client) SetTimeout(t time.Duration) *Client {]
ff=设置超时
t=时长

[func (c *Client) SetBasicAuth(user, pass string) *Client {]
ff=设置账号密码
pass=密码
user=账号

[func (c *Client) SetRetry(retryCount int, retryInterval time.Duration) *Client {]
ff=设置重试与间隔
retryInterval=重试间隔时长
retryCount=重试计数

[func (c *Client) SetRedirectLimit(redirectLimit int) *Client {]
ff=设置重定向次数限制
redirectLimit=次数

[func (c *Client) SetNoUrlEncode(noUrlEncode bool) *Client {]
ff=设置请求参数禁止URL编码
noUrlEncode=禁止编码

[func (c *Client) SetProxy(proxyURL string) {]
ff=设置代理
proxyURL=代理地址

[func (c *Client) SetTLSKeyCrt(crtFile, keyFile string) error {]
ff=设置证书
keyFile=key路径
crtFile=crt路径

[func (c *Client) SetTLSConfig(tlsConfig *tls.Config) error {]
ff=设置TLS配置
tlsConfig=TLS配置
