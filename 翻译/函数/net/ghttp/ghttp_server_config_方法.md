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

[func NewConfig() ServerConfig {]
ff=创建默认配置项

[func ConfigFromMap(m map#左中括号#string#右中括号#interface{}) (ServerConfig, error) {]
ff=创建配置对象Map
m=配置

[func (s *Server) SetConfigWithMap(m map#左中括号#string#右中括号#interface{}) error {]
ff=设置配置项Map
m=配置

[func (s *Server) SetConfig(c ServerConfig) error {]
ff=设置配置项

[func (s *Server) SetAddr(address string) {]
ff=设置监听地址
address=地址

[func (s *Server) SetPort(port ...int) {]
ff=设置监听端口
port=端口

[func (s *Server) SetHTTPSAddr(address string) {]
ff=设置HTTPS监听地址
address=地址

[func (s *Server) SetHTTPSPort(port ...int) {]
ff=设置HTTPS监听端口
port=端口

[func (s *Server) SetListener(listeners ...net.Listener) error {]
ff=设置自定义监听器
listeners=监听器

[func (s *Server) EnableHTTPS(certFile, keyFile string, tlsConfig ...*tls.Config) {]
ff=启用HTTPS
tlsConfig=tls配置
keyFile=密钥路径
certFile=证书路径

[func (s *Server) SetTLSConfig(tlsConfig *tls.Config) {]
ff=设置TLS配置
tlsConfig=tls配置

[func (s *Server) SetReadTimeout(t time.Duration) {]
ff=设置读取超时
t=时长

[func (s *Server) SetWriteTimeout(t time.Duration) {]
ff=设置写入超时
t=时长

[func (s *Server) SetIdleTimeout(t time.Duration) {]
ff=设置长连接超时
t=时长

[func (s *Server) SetMaxHeaderBytes(b int) {]
ff=设置协议头最大长度
b=最大长度

[func (s *Server) SetServerAgent(agent string) {]
ff=设置服务器代理标识
agent=代理标识

[func (s *Server) SetKeepAlive(enabled bool) {]
ff=设置开启长连接
enabled=开启

[func (s *Server) SetView(view *gview.View) {]
ff=设置默认模板对象
view=模板对象

[func (s *Server) GetName() string {]
ff=取服务名称

[func (s *Server) SetName(name string) {]
ff=设置服务名称
name=名称

[func (s *Server) SetHandler(h func(w http.ResponseWriter, r *http.Request)) {]
ff=设置请求处理器

[func (s *Server) GetHandler() func(w http.ResponseWriter, r *http.Request) {]
ff=取请求处理器

[func (s *Server) SetRegistrar(registrar gsvc.Registrar) {]
ff=设置注册发现对象
registrar=注册发现对象

[func (s *Server) GetRegistrar() gsvc.Registrar {]
ff=取注册发现对象
