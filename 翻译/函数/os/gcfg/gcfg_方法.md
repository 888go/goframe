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

[func NewWithAdapter(adapter Adapter) *Config {]
ff=创建并按适配器
adapter=适配器

[func Instance(name ...string) *Config {]
ff=取单例对象
name=名称

[func (c *Config) SetAdapter(adapter Adapter) {]
ff=设置适配器
adapter=适配器

[func (c *Config) GetAdapter() Adapter {]
ff=取适配器

[func (c *Config) Available(ctx context.Context, resource ...string) (ok bool) {]
ff=是否可用
ok=可用
ctx=上下文

[func (c *Config) Get(ctx context.Context, pattern string, def ...interface{}) (*gvar.Var, error) {]
ff=取值
def=默认值
pattern=表达式
ctx=上下文

[func (c *Config) GetWithEnv(ctx context.Context, pattern string, def ...interface{}) (*gvar.Var, error) {]
ff=取值并从环境变量
def=默认值
pattern=表达式
ctx=上下文

[func (c *Config) GetWithCmd(ctx context.Context, pattern string, def ...interface{}) (*gvar.Var, error) {]
ff=取值并从启动命令
def=默认值
pattern=表达式
ctx=上下文

[func (c *Config) Data(ctx context.Context) (data map#左中括号#string#右中括号#interface{}, err error) {]
ff=取Map
err=错误
data=值
ctx=上下文

[func (c *Config) MustGet(ctx context.Context, pattern string, def ...interface{}) *gvar.Var {]
ff=取值PANI
def=默认值
pattern=表达式
ctx=上下文

[func (c *Config) MustGetWithEnv(ctx context.Context, pattern string, def ...interface{}) *gvar.Var {]
ff=取值并从环境变量PANI
def=默认值
pattern=表达式
ctx=上下文

[func (c *Config) MustGetWithCmd(ctx context.Context, pattern string, def ...interface{}) *gvar.Var {]
ff=取值并从启动命令PANI_有bug
def=默认值
pattern=表达式
ctx=上下文

[func (c *Config) MustData(ctx context.Context) map#左中括号#string#右中括号#interface{} {]
ff=取MapPANI
ctx=上下文
