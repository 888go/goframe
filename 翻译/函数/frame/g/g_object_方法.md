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

[func Client() *gclient.Client {]
ff=网页类

[func Server(name ...interface{}) *ghttp.Server {]
ff=Http类
name=名称

[func TCPServer(name ...interface{}) *gtcp.Server {]
ff=TCP类
name=名称

[func UDPServer(name ...interface{}) *gudp.Server {]
ff=UDP类
name=名称

[func View(name ...string) *gview.View {]
ff=模板类
name=名称

[func Config(name ...string) *gcfg.Config {]
ff=配置类
name=名称

[func Cfg(name ...string) *gcfg.Config {]
ff=Cfg别名
name=名称

[func Resource(name ...string) *gres.Resource {]
ff=资源类
name=名称

[func I18n(name ...string) *gi18n.Manager {]
ff=多语言类
name=名称

[func Res(name ...string) *gres.Resource {]
ff=Res别名
name=名称

[func Log(name ...string) *glog.Logger {]
ff=日志类
name=名称

[func DB(name ...string) gdb.DB {]
ff=DB类
name=名称

[func Model(tableNameOrStruct ...interface{}) *gdb.Model {]
ff=DB类Model
tableNameOrStruct=表名或结构体

[func ModelRaw(rawSql string, args ...interface{}) *gdb.Model {]
ff=DB类原生SQL
args=参数
rawSql=原生Sql

[func Redis(name ...string) *gredis.Redis {]
ff=Redis类
name=名称
[func Validator() *gvalid.Validator {]
ff=效验类
