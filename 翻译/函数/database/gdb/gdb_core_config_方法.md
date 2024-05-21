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

[func SetConfig(config Config) {]
ff=设置全局配置
config=配置

[func SetConfigGroup(group string, nodes ConfigGroup) {]
ff=设置组配置
nodes=配置
group=配置组名称

[func AddConfigNode(group string, node ConfigNode) {]
ff=添加配置组节点
node=配置
group=配置组名称

[func AddDefaultConfigNode(node ConfigNode) {]
ff=添加默认配置组节点
node=配置

[func AddDefaultConfigGroup(nodes ConfigGroup) {]
ff=添加默认配置组
nodes=配置组

[func GetConfig(group string) ConfigGroup {]
ff=取配置组配置
group=配置组名称

[func SetDefaultGroup(name string) {]
ff=设置默认组名称
name=配置组名称

[func GetDefaultGroup() string {]
ff=获取默认组名称

[func IsConfigured() bool {]
ff=是否已配置数据库

[func (c *Core) SetLogger(logger glog.ILogger) {]
ff=设置日志记录器
logger=日志记录器

[func (c *Core) GetLogger() glog.ILogger {]
ff=取日志记录器

[func (c *Core) SetMaxIdleConnCount(n int) {]
ff=设置最大闲置连接数
n=连接数

[func (c *Core) SetMaxOpenConnCount(n int) {]
ff=设置最大打开连接数
n=连接数

[func (c *Core) SetMaxConnLifeTime(d time.Duration) {]
ff=设置最大空闲时长
d=时长

[func (c *Core) GetConfig() *ConfigNode {]
ff=取当前节点配置

[func (c *Core) SetDebug(debug bool) {]
ff=设置调试模式
debug=开启

[func (c *Core) GetDebug() bool {]
ff=取调试模式

[func (c *Core) GetCache() *gcache.Cache {]
ff=取缓存对象

[func (c *Core) GetGroup() string {]
ff=取配置组名称

[func (c *Core) SetDryRun(enabled bool) {]
ff=设置空跑特性
enabled=开启

[func (c *Core) GetDryRun() bool {]
ff=取空跑特性

[func (c *Core) GetPrefix() string {]
ff=取表前缀

[func (c *Core) GetSchema() string {]
ff=取默认数据库名称
