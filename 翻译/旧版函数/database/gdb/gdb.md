
# <翻译开始>
Model(tableNameOrStruct ...interface{})
X创建Model对象
<跳到行首>
# <翻译结束>

# <翻译开始>
Raw(rawSql string, args ...interface{})
X原生SQL
<跳到行首>
<忽略重复>
# <翻译结束>

# <翻译开始>
Schema(schema string)
X切换数据库
<跳到行首>
# <翻译结束>

# <翻译开始>
With(objects ...interface{})
X关联对象
<跳到行首>
# <翻译结束>

# <翻译开始>
Open(config *ConfigNode)
X底层Open
<跳到行首>
# <翻译结束>

# <翻译开始>
Ctx(ctx context.Context)
X设置上下文并取副本
<跳到行首>
<忽略重复>
# <翻译结束>

# <翻译开始>
Close(ctx context.Context)
X关闭数据库
<跳到行首>
# <翻译结束>

# <翻译开始>
Query(ctx context.Context, sql string, args ...interface{})
X原生SQL查询
<跳到行首>
# <翻译结束>

# <翻译开始>
Exec(ctx context.Context, sql string, args ...interface{})
X原生SQL执行
<跳到行首>
# <翻译结束>

# <翻译开始>
Prepare(ctx context.Context, sql string, execOnMaster ...bool)
X原生sql取参数预处理对象
<跳到行首>
# <翻译结束>

# <翻译开始>
Insert(ctx context.Context, table string, data interface{}, batch ...int)
X插入
<跳到行首>
# <翻译结束>

# <翻译开始>
InsertIgnore(ctx context.Context, table string, data interface{
X插入并跳过已存在
<跳到行首>
# <翻译结束>

# <翻译开始>
InsertAndGetId(ctx context.Context, table string, data interface{}
X插入并取ID
<跳到行首>
# <翻译结束>

# <翻译开始>
Replace(ctx context.Context, table string, data interface{}, batch ...int)
X插入并替换已存在
<跳到行首>
# <翻译结束>

# <翻译开始>
Save(ctx context.Context, table string, data interface{}, batch ...int)
X插入并更新已存在
<跳到行首>
# <翻译结束>

# <翻译开始>
Update(ctx context.Context, table string, data interface{}, condition interface{}, args ...interface{})
X更新
<跳到行首>
# <翻译结束>

# <翻译开始>
Delete(ctx context.Context, table string, condition interface{}, args ...interface{})
X删除
<跳到行首>
# <翻译结束>

# <翻译开始>
DoSelect(ctx context.Context, link Link, sql string, args ...interface{})
X底层查询
<跳到行首>
# <翻译结束>

# <翻译开始>
DoInsert(ctx context.Context, link Link, table string, data List, option DoInsertOption)
X底层插入
<跳到行首>
# <翻译结束>

# <翻译开始>
DoUpdate(ctx context.Context, link Link, table string, data interface{}, condition string, args ...interface{})
X底层更新
<跳到行首>
# <翻译结束>

# <翻译开始>
DoDelete(ctx context.Context, link Link, table string, condition string, args ...interface{})
X底层删除
<跳到行首>
# <翻译结束>

# <翻译开始>
DoQuery(ctx context.Context, link Link, sql string, args ...interface{})
X底层原生SQL查询
<跳到行首>
# <翻译结束>

# <翻译开始>
DoExec(ctx context.Context, link Link, sql string, args ...interface{})
X底层原生SQL执行
<跳到行首>
# <翻译结束>

# <翻译开始>
DoFilter(ctx context.Context, link Link, sql string, args []interface{})
X底层DoFilter
<跳到行首>
# <翻译结束>

# <翻译开始>
DoCommit(ctx context.Context, in DoCommitInput)
X底层DoCommit
<跳到行首>
# <翻译结束>

# <翻译开始>
DoPrepare(ctx context.Context, link Link, sql string)
X底层原生sql参数预处理对象
<跳到行首>
# <翻译结束>

# <翻译开始>
GetAll(ctx context.Context, sql string, args ...interface{})
GetAll别名
<跳到行首>
# <翻译结束>

# <翻译开始>
GetOne(ctx context.Context, sql string, args ...interface{})
X原生SQL查询单条记录
<跳到行首>
# <翻译结束>

# <翻译开始>
GetValue(ctx context.Context, sql string, args ...interface{})
X原生SQL查询字段值
<跳到行首>
# <翻译结束>

# <翻译开始>
GetArray(ctx context.Context, sql string, args ...interface{})
X原生SQL查询切片
<跳到行首>
# <翻译结束>

# <翻译开始>
GetCount(ctx context.Context, sql string, args ...interface{})
X原生SQL查询字段计数
<跳到行首>
# <翻译结束>

# <翻译开始>
GetScan(ctx context.Context, objPointer interface{}, sql string, args ...interface{})
X原生SQL查询到结构体指针
<跳到行首>
# <翻译结束>

# <翻译开始>
Union(unions ...*Model)
X多表去重查询
<跳到行首>
# <翻译结束>

# <翻译开始>
UnionAll(unions ...*Model)
X多表查询
<跳到行首>
# <翻译结束>

# <翻译开始>
Master(schema ...string)
X取主节点对象
<跳到行首>
# <翻译结束>

# <翻译开始>
Slave(schema ...string)
X取从节点对象
<跳到行首>
# <翻译结束>

# <翻译开始>
PingMaster()
X向主节点发送心跳
<跳到行首>
# <翻译结束>

# <翻译开始>
PingSlave()
X向从节点发送心跳
<跳到行首>
# <翻译结束>

# <翻译开始>
Begin(ctx context.Context)
X事务开启
<跳到行首>
# <翻译结束>

# <翻译开始>
Transaction(ctx context.Context, f func(ctx context.Context, tx TX) error)
X事务
<跳到行首>
<忽略重复>
# <翻译结束>

# <翻译开始>
GetCache()
X取缓存对象
<跳到行首>
# <翻译结束>

# <翻译开始>
SetDebug(debug bool)
X设置调试模式
<跳到行首>
# <翻译结束>

# <翻译开始>
GetDebug() bool
X取调试模式
<跳到行首>
# <翻译结束>

# <翻译开始>
GetSchema() string
X取默认数据库名称
<跳到行首>
# <翻译结束>

# <翻译开始>
GetPrefix() string
X取表前缀
<跳到行首>
# <翻译结束>

# <翻译开始>
GetGroup() string
X取配置组名称
<跳到行首>
# <翻译结束>

# <翻译开始>
SetDryRun(enabled bool)
X设置空跑特性
<跳到行首>
# <翻译结束>

# <翻译开始>
GetDryRun()
X取空跑特性
<跳到行首>
# <翻译结束>

# <翻译开始>
SetLogger(
X设置日志记录器
<跳到行首>
# <翻译结束>

# <翻译开始>
GetLogger()
X取日志记录器
<跳到行首>
# <翻译结束>

# <翻译开始>
GetConfig() *ConfigNode
X取当前节点配置
<跳到行首>
# <翻译结束>

# <翻译开始>
SetMaxIdleConnCount(n int)
X设置最大闲置连接数
<跳到行首>
# <翻译结束>

# <翻译开始>
SetMaxOpenConnCount(n int)
X设置最大打开连接数
<跳到行首>
# <翻译结束>

# <翻译开始>
SetMaxConnLifeTime(d time.Duration)
X设置最大空闲时长
<跳到行首>
# <翻译结束>

# <翻译开始>
GetCtx() context.Context
X取上下文对象
<跳到行首>
<忽略重复>
# <翻译结束>

# <翻译开始>
GetCore() *Core
X取Core对象
<跳到行首>
# <翻译结束>

# <翻译开始>
GetChars() (charLeft string, charRight string)
X底层取数据库安全字符
<跳到行首>
# <翻译结束>

# <翻译开始>
Tables(ctx context.Context, schema ...string)
X取表名称切片
<跳到行首>
# <翻译结束>

# <翻译开始>
TableFields(ctx context.Context, table string, schema ...string)
X取表字段信息Map
<跳到行首>
# <翻译结束>

# <翻译开始>
ConvertValueForField(ctx context.Context, fieldType string, fieldValue interface{})
X底层ConvertValueForField
<跳到行首>
# <翻译结束>

# <翻译开始>
ConvertValueForLocal(ctx context.Context, fieldType string, fieldValue interface{})
X底层ConvertValueForLocal
<跳到行首>
# <翻译结束>

# <翻译开始>
CheckLocalTypeForField(ctx context.Context, fieldType string, fieldValue interface{})
X底层CheckLocalTypeForField
<跳到行首>
# <翻译结束>

# <翻译开始>
Ctx(ctx context.Context)
X设置上下文并取副本
<跳到行首>
<忽略重复>
# <翻译结束>

# <翻译开始>
Raw(rawSql string, args ...interface{})
X原生SQL
<跳到行首>
<忽略重复>
# <翻译结束>

# <翻译开始>
Model(tableNameQueryOrStruct ...interface{})
X创建Model对象
<跳到行首>
# <翻译结束>

# <翻译开始>
With(object interface{})
X关联对象
<跳到行首>
# <翻译结束>

# <翻译开始>
Begin() error
X事务开启
<跳到行首>
# <翻译结束>

# <翻译开始>
Commit() error
X事务提交
<跳到行首>
# <翻译结束>

# <翻译开始>
Rollback() error
X事务回滚
<跳到行首>
# <翻译结束>

# <翻译开始>
Transaction(ctx context.Context, f func(ctx context.Context, tx TX)
X事务
<跳到行首>
<忽略重复>
# <翻译结束>

# <翻译开始>
Query(sql string, args ...interface{})
X原生SQL查询
<跳到行首>
# <翻译结束>

# <翻译开始>
Exec(sql string, args ...interface{})
X原生SQL执行
<跳到行首>
# <翻译结束>

# <翻译开始>
Prepare(sql string)
X原生sql取参数预处理对象
<跳到行首>
# <翻译结束>

# <翻译开始>
GetAll(sql string, args ...interface{})
GetAll别名
<跳到行首>
# <翻译结束>

# <翻译开始>
GetOne(sql string, args ...interface{})
X原生SQL查询单条记录
<跳到行首>
# <翻译结束>

# <翻译开始>
GetStruct(obj interface{}, sql string, args ...interface{})
X原生SQL查询单条到结构体指针
<跳到行首>
# <翻译结束>

# <翻译开始>
GetStructs(objPointerSlice interface{}, sql string, args ...interface{})
X原生SQL查询到结构体切片指针
<跳到行首>
# <翻译结束>

# <翻译开始>
GetScan(pointer interface{}, sql string, args ...interface{})
X原生SQL查询到结构体指针
<跳到行首>
# <翻译结束>

# <翻译开始>
GetValue(sql string, args ...interface{})
X原生SQL查询字段值
<跳到行首>
# <翻译结束>

# <翻译开始>
GetCount(sql string, args ...interface{})
X原生SQL查询字段计数
<跳到行首>
# <翻译结束>

# <翻译开始>
Insert(table string, data interface{}, batch ...int)
X插入
<跳到行首>
# <翻译结束>

# <翻译开始>
InsertIgnore(table string, data interface{}, batch ...int)
X插入并跳过已存在
<跳到行首>
# <翻译结束>

# <翻译开始>
InsertAndGetId(table string, data interface{}, batch ...int)
X插入并取ID
<跳到行首>
# <翻译结束>

# <翻译开始>
Replace(table string, data interface{}, batch ...int)
X插入并替换已存在
<跳到行首>
# <翻译结束>

# <翻译开始>
Save(table string, data interface{}, batch ...int)
X插入并更新已存在
<跳到行首>
# <翻译结束>

# <翻译开始>
Update(table string, data interface{}, condition interface{}, args ...interface{})
X更新
<跳到行首>
# <翻译结束>

# <翻译开始>
Delete(table string, condition interface{}, args ...interface{})
X删除
<跳到行首>
# <翻译结束>

# <翻译开始>
GetCtx() context.Context
X取上下文对象
<跳到行首>
<忽略重复>
# <翻译结束>

# <翻译开始>
GetDB() DB
X取DB对象
<跳到行首>
# <翻译结束>

# <翻译开始>
GetSqlTX() *sql.Tx
X底层取事务对象
<跳到行首>
# <翻译结束>

# <翻译开始>
IsClosed() bool
X是否已关闭
<跳到行首>
# <翻译结束>

# <翻译开始>
SavePoint(point string) error
X保存事务点
<跳到行首>
# <翻译结束>

# <翻译开始>
RollbackTo(point string) error
X回滚事务点
<跳到行首>
# <翻译结束>

# <翻译开始>
func Register(name string, driver
驱动
# <翻译结束>

# <翻译开始>
func Register(name
名称
# <翻译结束>

# <翻译开始>
func Register
X注册驱动
# <翻译结束>

# <翻译开始>
func New(node ConfigNode) (db DB, err
错误
# <翻译结束>

# <翻译开始>
func New(node ConfigNode) (db
DB对象
# <翻译结束>

# <翻译开始>
func New(node
配置项
# <翻译结束>

# <翻译开始>
func New
X创建DB对象
# <翻译结束>

# <翻译开始>
func NewByGroup(group ...string) (db DB, err
错误
# <翻译结束>

# <翻译开始>
func NewByGroup(group ...string) (db
DB对象
# <翻译结束>

# <翻译开始>
func NewByGroup(group
配置组名称
# <翻译结束>

# <翻译开始>
func NewByGroup
X创建DB对象并按配置组
# <翻译结束>

# <翻译开始>
func Instance(name ...string) (db DB, err
错误
# <翻译结束>

# <翻译开始>
func Instance(name ...string) (db
DB对象
# <翻译结束>

# <翻译开始>
func Instance(name
配置组名称
# <翻译结束>

# <翻译开始>
func Instance
X取单例对象
# <翻译结束>

# <翻译开始>
List = []Map
Map切片
<跳到行首>
# <翻译结束>
