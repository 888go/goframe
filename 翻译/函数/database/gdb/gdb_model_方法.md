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

[func (c *Core) Model(tableNameQueryOrStruct ...interface{}) *Model {]
ff=创建Model对象
tableNameQueryOrStruct=表名或结构体

[func (c *Core) Raw(rawSql string, args ...interface{}) *Model {]
ff=原生SQL
args=参数
rawSql=原生Sql

[func (m *Model) Raw(rawSql string, args ...interface{}) *Model {]
ff=原生SQL
args=参数
rawSql=原生Sql

[func (tx *TXCore) Raw(rawSql string, args ...interface{}) *Model {]
ff=原生SQL
args=参数
rawSql=原生Sql

[func (c *Core) With(objects ...interface{}) *Model {]
ff=关联对象
objects=关联结构体

[func (m *Model) Partition(partitions ...string) *Model {]
ff=设置分区名称
partitions=分区名称

[func (tx *TXCore) Model(tableNameQueryOrStruct ...interface{}) *Model {]
ff=创建Model对象
tableNameQueryOrStruct=表名或结构体

[func (tx *TXCore) With(object interface{}) *Model {]
ff=关联对象
object=关联结构体

[func (m *Model) Ctx(ctx context.Context) *Model {]
ff=设置上下文并取副本
ctx=上下文

[func (m *Model) GetCtx() context.Context {]
ff=取上下文对象

[func (m *Model) As(as string) *Model {]
ff=设置表别名
as=别名

[func (m *Model) DB(db DB) *Model {]
ff=设置DB对象
db=DB对象

[func (m *Model) TX(tx TX) *Model {]
ff=设置事务对象
tx=事务对象

[func (m *Model) Schema(schema string) *Model {]
ff=切换数据库
schema=数据库名

[func (m *Model) Clone() *Model {]
ff=取副本

[func (m *Model) Master() *Model {]
ff=取主节点对象

[func (m *Model) Slave() *Model {]
ff=取从节点对象

[func (m *Model) Safe(safe ...bool) *Model {]
ff=链式安全
safe=开启

[func (m *Model) Args(args ...interface{}) *Model {]
ff=底层Args
args=参数

[func (m *Model) Handler(handlers ...ModelHandler) *Model {]
ff=处理函数
handlers=处理函数
