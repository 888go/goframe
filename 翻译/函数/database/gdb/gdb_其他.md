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

[Model(tableNameOrStruct ...interface{}) *Model]
qm=创建Model对象
cz=Model(tableNameOrStruct ...interface{})

[Raw(rawSql string, args ...interface{}) *Model]
qm=原生SQL
cz=Raw(rawSql string, args ...interface{})
cf=2

[Schema(schema string) *Schema]
qm=切换数据库
cz=Schema(schema string)

[With(objects ...interface{}) *Model]
qm=关联对象
cz=With(objects ...interface{})

[Open(config *ConfigNode) (*sql.DB, error)]
qm=底层Open
cz=Open(config *ConfigNode)

[Ctx(ctx context.Context) DB]
qm=设置上下文并取副本
cz=Ctx(ctx context.Context)

[Close(ctx context.Context) error]
qm=关闭数据库
cz=Close(ctx context.Context)

[Query(ctx context.Context, sql string, args ...interface{}) (Result, error)]
qm=原生SQL查询
cz=Query(ctx context.Context, sql string, args ...interface{})

[Exec(ctx context.Context, sql string, args ...interface{}) (sql.Result, error)]
qm=原生SQL执行
cz=Exec(ctx context.Context, sql string, args ...interface{})

[Prepare(ctx context.Context, sql string, execOnMaster ...bool) (*Stmt, error)]
qm=原生sql取参数预处理对象
cz=Prepare(ctx context.Context, sql string, execOnMaster ...bool)

[Insert(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)]
qm=插入
cz=Insert(ctx context.Context, table string, data interface{}, batch ...int)

[InsertIgnore(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)]
qm=插入并跳过已存在
cz=InsertIgnore(ctx context.Context, table string, data interface{

[InsertAndGetId(ctx context.Context, table string, data interface{}, batch ...int) (int64, error)]
qm=插入并取ID
cz=InsertAndGetId(ctx context.Context, table string, data interface{}

[Replace(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)]
qm=插入并替换已存在
cz=Replace(ctx context.Context, table string, data interface{}, batch ...int)

[Save(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)]
qm=插入并更新已存在
cz=Save(ctx context.Context, table string, data interface{}, batch ...int)

[Update(ctx context.Context, table string, data interface{}, condition interface{}, args ...interface{}) (sql.Result, error)]
qm=更新
cz=Update(ctx context.Context, table string, data interface{}, condition interface{}, args ...interface{})

[Delete(ctx context.Context, table string, condition interface{}, args ...interface{}) (sql.Result, error)]
qm=删除
cz=Delete(ctx context.Context, table string, condition interface{}, args ...interface{})

[DoSelect(ctx context.Context, link Link, sql string, args ...interface{}) (result Result, err error)]
qm=底层查询
cz=DoSelect(ctx context.Context, link Link, sql string, args ...interface{})

[DoInsert(ctx context.Context, link Link, table string, data List, option DoInsertOption) (result sql.Result, err error)]
qm=原生sql行记录
cz=DoInsert(ctx context.Context, link Link, table string, data List, option DoInsertOption)

[DoUpdate(ctx context.Context, link Link, table string, data interface{}, condition string, args ...interface{}) (result sql.Result, err error)]
qm=底层更新
cz=DoUpdate(ctx context.Context, link Link, table string, data interface{}, condition string, args ...interface{})

[DoDelete(ctx context.Context, link Link, table string, condition string, args ...interface{}) (result sql.Result, err error)]
qm=底层删除
cz=DoDelete(ctx context.Context, link Link, table string, condition string, args ...interface{})

[DoQuery(ctx context.Context, link Link, sql string, args ...interface{}) (result Result, err error)]
qm=底层原生SQL查询
cz=DoQuery(ctx context.Context, link Link, sql string, args ...interface{})

[DoExec(ctx context.Context, link Link, sql string, args ...interface{}) (result sql.Result, err error)]
qm=底层原生SQL执行
cz=DoExec(ctx context.Context, link Link, sql string, args ...interface{})

[DoFilter(ctx context.Context, link Link, sql string, args #左中括号##右中括号#interface{}) (newSql string, newArgs #左中括号##右中括号#interface{}, err error)]
qm=底层DoFilter
cz=DoFilter(ctx context.Context, link Link, sql string, args []interface{})

[DoCommit(ctx context.Context, in DoCommitInput) (out DoCommitOutput, err error)]
qm=底层DoCommit
cz=DoCommit(ctx context.Context, in DoCommitInput)

[DoPrepare(ctx context.Context, link Link, sql string) (*Stmt, error)]
qm=底层原生sql参数预处理对象
cz=DoPrepare(ctx context.Context, link Link, sql string)

[GetAll(ctx context.Context, sql string, args ...interface{}) (Result, error)]
qm=GetAll别名
cz=GetAll(ctx context.Context, sql string, args ...interface{})

[GetOne(ctx context.Context, sql string, args ...interface{}) (Record, error)]
qm=原生SQL查询单条记录
cz=GetOne(ctx context.Context, sql string, args ...interface{})

[GetValue(ctx context.Context, sql string, args ...interface{}) (Value, error)]
qm=原生SQL查询字段值
cz=GetValue(ctx context.Context, sql string, args ...interface{})

[GetArray(ctx context.Context, sql string, args ...interface{}) (#左中括号##右中括号#Value, error)]
qm=原生SQL查询切片
cz=GetArray(ctx context.Context, sql string, args ...interface{})

[GetCount(ctx context.Context, sql string, args ...interface{}) (int, error)]
qm=原生SQL查询字段计数
cz=GetCount(ctx context.Context, sql string, args ...interface{})

[GetScan(ctx context.Context, objPointer interface{}, sql string, args ...interface{}) error]
qm=原生SQL查询到结构体指针
cz=GetScan(ctx context.Context, objPointer interface{}, sql string, args ...interface{})

[Union(unions ...*Model) *Model]
qm=多表去重查询
cz=Union(unions ...*Model)

[UnionAll(unions ...*Model) *Model]
qm=多表查询
cz=UnionAll(unions ...*Model)

[Master(schema ...string) (*sql.DB, error)]
qm=取主节点对象
cz=Master(schema ...string)

[Slave(schema ...string) (*sql.DB, error)]
qm=取从节点对象
cz=Slave(schema ...string)

[PingMaster() error]
qm=向主节点发送心跳
cz=PingMaster()

[PingSlave() error]
qm=向从节点发送心跳
cz=PingSlave()

[Begin(ctx context.Context) (TX, error)]
qm=事务开启
cz=Begin(ctx context.Context)

[Transaction(ctx context.Context, f func(ctx context.Context, tx TX) error) error]
qm=事务
cz=Transaction(ctx context.Context, f func(ctx context.Context, tx TX) error)

[GetCache() *gcache.Cache]
qm=取缓存对象
cz=GetCache()

[SetDebug(debug bool)]
qm=设置调试模式
cz=SetDebug(debug bool)

[GetDebug() bool]
qm=取调试模式
cz=GetDebug() bool

[GetSchema() string]
qm=取默认数据库名称
cz=GetSchema() string

[GetPrefix() string]
qm=取表前缀
cz=GetPrefix() string

[GetGroup() string]
qm=取配置组名称
cz=GetGroup() string

[SetDryRun(enabled bool)]
qm=设置空跑特性
cz=SetDryRun(enabled bool)

[GetDryRun() bool]
qm=取空跑特性
cz=GetDryRun()

[SetLogger(logger glog.ILogger)]
qm=设置日志记录器
cz=SetLogger(

[GetLogger() glog.ILogger]
qm=取日志记录器
cz=GetLogger()

[GetConfig() *ConfigNode]
qm=取当前节点配置
cz=GetConfig() *ConfigNode

[SetMaxIdleConnCount(n int)]
qm=设置最大闲置连接数
cz=SetMaxIdleConnCount(n int)

[SetMaxOpenConnCount(n int)]
qm=设置最大打开连接数
cz=SetMaxOpenConnCount(n int)

[SetMaxConnLifeTime(d time.Duration)]
qm=设置最大空闲时长
cz=SetMaxConnLifeTime(d time.Duration)

[GetCtx() context.Context]
qm=取上下文对象
cz=GetCtx() context.Context

[GetCore() *Core]
qm=取Core对象
cz=GetCore() *Core

[GetChars() (charLeft string, charRight string)]
qm=底层取数据库安全字符
cz=GetChars() (charLeft string, charRight string)

[Tables(ctx context.Context, schema ...string) (tables #左中括号##右中括号#string, err error)]
qm=取表名称切片
cz=Tables(ctx context.Context, schema ...string)

[TableFields(ctx context.Context, table string, schema ...string) (map#左中括号#string#右中括号#*TableField, error)]
qm=取表字段信息Map
cz=TableFields(ctx context.Context, table string, schema ...string)

[ConvertValueForField(ctx context.Context, fieldType string, fieldValue interface{}) (interface{}, error)]
qm=底层ConvertValueForField
cz=ConvertValueForField(ctx context.Context, fieldType string, fieldValue interface{})

[ConvertValueForLocal(ctx context.Context, fieldType string, fieldValue interface{}) (interface{}, error)]
qm=底层ConvertValueForLocal
cz=ConvertValueForLocal(ctx context.Context, fieldType string, fieldValue interface{})

[CheckLocalTypeForField(ctx context.Context, fieldType string, fieldValue interface{}) (LocalType, error)]
qm=底层CheckLocalTypeForField
cz=CheckLocalTypeForField(ctx context.Context, fieldType string, fieldValue interface{})

[Model(tableNameQueryOrStruct ...interface{}) *Model]
qm=创建Model对象
cz=Model(tableNameQueryOrStruct ...interface{})

[With(object interface{}) *Model]
qm=关联对象
cz=With(object interface{})

[Begin() error]
qm=事务开启
cz=Begin() error

[Commit() error]
qm=事务提交
cz=Commit() error

[Rollback() error]
qm=事务回滚
cz=Rollback() error

[Query(sql string, args ...interface{}) (result Result, err error)]
qm=原生SQL查询
cz=Query(sql string, args ...interface{})

[Exec(sql string, args ...interface{}) (sql.Result, error)]
qm=原生SQL执行
cz=Exec(sql string, args ...interface{})

[Prepare(sql string) (*Stmt, error)]
qm=原生sql取参数预处理对象
cz=Prepare(sql string)

[GetAll(sql string, args ...interface{}) (Result, error)]
qm=GetAll别名
cz=GetAll(sql string, args ...interface{})

[GetOne(sql string, args ...interface{}) (Record, error)]
qm=原生SQL查询单条记录
cz=GetOne(sql string, args ...interface{})

[GetStruct(obj interface{}, sql string, args ...interface{}) error]
qm=原生SQL查询单条到结构体指针
cz=GetStruct(obj interface{}, sql string, args ...interface{})

[GetStructs(objPointerSlice interface{}, sql string, args ...interface{}) error]
qm=原生SQL查询到结构体切片指针
cz=GetStructs(objPointerSlice interface{}, sql string, args ...interface{})

[GetScan(pointer interface{}, sql string, args ...interface{}) error]
qm=原生SQL查询到结构体指针
cz=GetScan(pointer interface{}, sql string, args ...interface{})

[GetValue(sql string, args ...interface{}) (Value, error)]
qm=原生SQL查询字段值
cz=GetValue(sql string, args ...interface{})

[GetCount(sql string, args ...interface{}) (int64, error)]
qm=原生SQL查询字段计数
cz=GetCount(sql string, args ...interface{})

[Insert(table string, data interface{}, batch ...int) (sql.Result, error)]
qm=插入
cz=Insert(table string, data interface{}, batch ...int)

[InsertIgnore(table string, data interface{}, batch ...int) (sql.Result, error)]
qm=插入并跳过已存在
cz=InsertIgnore(table string, data interface{}, batch ...int)

[InsertAndGetId(table string, data interface{}, batch ...int) (int64, error)]
qm=插入并取ID
cz=InsertAndGetId(table string, data interface{}, batch ...int)

[Replace(table string, data interface{}, batch ...int) (sql.Result, error)]
qm=插入并替换已存在
cz=Replace(table string, data interface{}, batch ...int)

[Save(table string, data interface{}, batch ...int) (sql.Result, error)]
qm=插入并更新已存在
cz=Save(table string, data interface{}, batch ...int)

[Update(table string, data interface{}, condition interface{}, args ...interface{}) (sql.Result, error)]
qm=更新
cz=Update(table string, data interface{}, condition interface{}, args ...interface{})

[Delete(table string, condition interface{}, args ...interface{}) (sql.Result, error)]
qm=删除
cz=Delete(table string, condition interface{}, args ...interface{})

[GetDB() DB]
qm=取DB对象
cz=GetDB() DB

[GetSqlTX() *sql.Tx]
qm=底层取事务对象
cz=GetSqlTX() *sql.Tx

[IsClosed() bool]
qm=是否已关闭
cz=IsClosed() bool

[SavePoint(point string) error]
qm=保存事务点
cz=SavePoint(point string) error

[RollbackTo(point string) error]
qm=回滚事务点
cz=RollbackTo(point string) error

[Records   #左中括号##右中括号#Record]
qm=行记录切片
cz=Records   []

[Stmt      *Stmt]
qm=参数预处理
cz=Stmt      *Stmt

[RawResult interface{}]
qm=底层结果
cz=RawResult interface{}

[Field string]
qm=字段名称
cz=Field string

[Value float64]
qm=增减值
cz=Value float64

[List   = #左中括号##右中括号#Map]
qm=Map切片
cz=List #等号# []Map
