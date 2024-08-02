// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gdb为流行的关系型数据库提供ORM（对象关系映射）功能。
//
// 待办事项：将context.Context作为所有数据库操作的必需参数。
// md5:ed61b69bd00b7384
package db类

import (
	"context"
	"database/sql"
	"time"

	garray "github.com/888go/goframe/container/garray"
	gmap "github.com/888go/goframe/container/gmap"
	gtype "github.com/888go/goframe/container/gtype"
	gvar "github.com/888go/goframe/container/gvar"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	gcache "github.com/888go/goframe/os/gcache"
	gcmd "github.com/888go/goframe/os/gcmd"
	gctx "github.com/888go/goframe/os/gctx"
	glog "github.com/888go/goframe/os/glog"
	grand "github.com/888go/goframe/util/grand"
	gutil "github.com/888go/goframe/util/gutil"
)

// DB 定义 ORM 操作的接口。 md5:328f032182d38455
type DB interface {
	// =============================================================================
	// 模型创建。
	// =============================================================================
	// 这里是对一段Go代码中的注释进行翻译，"Model creation"指的是模型的创建过程。这部分代码可能是用于描述一个函数或部分代码的作用，即它负责构建或初始化某个模型。
	// md5:1c8c0b09089a9689

	// Model 根据给定的模式创建并返回一个新的 ORM 模型。
	// 参数 `table` 可以是多个表名，也可以包括别名，例如：
	// 1. 模型名称：
	//    Model("user")
	//    Model("user u") 	// u 作为 user 表的别名
	//    Model("user, user_detail") 	// 多个模型名称
	//    Model("user u, user_detail ud") 	// 多个模型名称及别名
	// 2. 带别名的模型名称：Model("user", "u") 	// 第二个参数指定别名
	// 参见 Core.Model 了解更多。
	// md5:61d3e6d835068122
	Model(tableNameOrStruct ...interface{}) *Model

		// Raw 根据原始SQL（而不是表格）创建并返回一个模型。 md5:96066a9d41296a2a
	Raw(rawSql string, args ...interface{}) *Model

	// Schema 创建并返回一个模式。
	// 参见 Core.Schema。
	// md5:0f4472ee79f06819
	Schema(schema string) *Schema

	// With根据给定对象的元数据创建并返回一个ORM模型。同时参见Core.With。
	// md5:78ab17ce6b00ce6e
	With(objects ...interface{}) *Model

	// Open 使用给定的节点配置为数据库创建一个原始连接对象。
	// 注意，不建议手动使用此函数。
	// 另请参阅 DriverMysql.Open。
	// md5:1021f26472df579e
	Open(config *ConfigNode) (*sql.DB, error)

	// Ctx 是一个链式函数，它创建并返回一个新的 DB，该 DB 是当前 DB 对象的浅拷贝，并在其中设置了给定的上下文。
	// 另请参阅 Core.Ctx。
	// md5:7eec5fab912764e7
	Ctx(ctx context.Context) DB

	// Close 关闭数据库并阻止新的查询开始。然后，Close 等待已经在服务器上开始处理的所有查询完成。
	// 
	// 通常不会关闭 DB，因为 DB句柄应该是长期存在的，并且在多个 goroutine 之间共享。
	// md5:0985fc8e558f83fc
	Close(ctx context.Context) error

	// =============================================================================
	// 查询接口。
	// =============================================================================
	// 这里是对查询相关的API进行的注释。
	// md5:06da8c4c9c8d957b

	Query(ctx context.Context, sql string, args ...interface{}) (Result, error)    // See Core.Query.
	Exec(ctx context.Context, sql string, args ...interface{}) (sql.Result, error) // See Core.Exec.
	Prepare(ctx context.Context, sql string, execOnMaster ...bool) (*Stmt, error)  // See Core.Prepare.

	// ===========================================================================
	// 用于CURD操作的通用API。
	// ===========================================================================
	// md5:781fc1b4ac386204

	Insert(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)                               // See Core.Insert.
	InsertIgnore(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)                         // See Core.InsertIgnore.
	InsertAndGetId(ctx context.Context, table string, data interface{}, batch ...int) (int64, error)                            // 参见Core.InsertAndGetId。 md5:b7dec69920da6e7a
	Replace(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)                              // See Core.Replace.
	Save(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)                                 // See Core.Save.
	Update(ctx context.Context, table string, data interface{}, condition interface{}, args ...interface{}) (sql.Result, error) // See Core.Update.
	Delete(ctx context.Context, table string, condition interface{}, args ...interface{}) (sql.Result, error)                   // See Core.Delete.

	// ===========================================================================
	// CURD的内部API，可以被自定义的CURD实现覆盖。
	// ===========================================================================
	// md5:02480feeb95bda1e

	DoSelect(ctx context.Context, link Link, sql string, args ...interface{}) (result Result, err error)                                           // See Core.DoSelect.
	DoInsert(ctx context.Context, link Link, table string, data List, option DoInsertOption) (result sql.Result, err error)                        // See Core.DoInsert.
	DoUpdate(ctx context.Context, link Link, table string, data interface{}, condition string, args ...interface{}) (result sql.Result, err error) // See Core.DoUpdate.
	DoDelete(ctx context.Context, link Link, table string, condition string, args ...interface{}) (result sql.Result, err error)                   // See Core.DoDelete.

	DoQuery(ctx context.Context, link Link, sql string, args ...interface{}) (result Result, err error)    // See Core.DoQuery.
	DoExec(ctx context.Context, link Link, sql string, args ...interface{}) (result sql.Result, err error) // See Core.DoExec.

	DoFilter(ctx context.Context, link Link, sql string, args []interface{}) (newSql string, newArgs []interface{}, err error) // See Core.DoFilter.
	DoCommit(ctx context.Context, in DoCommitInput) (out DoCommitOutput, err error)                                            // See Core.DoCommit.

	DoPrepare(ctx context.Context, link Link, sql string) (*Stmt, error) // See Core.DoPrepare.

	// ===========================================================================
	// 为了方便起见，提供查询API。
	// ===========================================================================
	// md5:be53a34b0863cf28

	GetAll(ctx context.Context, sql string, args ...interface{}) (Result, error)                // See Core.GetAll.
	GetOne(ctx context.Context, sql string, args ...interface{}) (Record, error)                // See Core.GetOne.
	GetValue(ctx context.Context, sql string, args ...interface{}) (Value, error)               // See Core.GetValue.
	GetArray(ctx context.Context, sql string, args ...interface{}) ([]Value, error)             // See Core.GetArray.
	GetCount(ctx context.Context, sql string, args ...interface{}) (int, error)                 // See Core.GetCount.
	GetScan(ctx context.Context, objPointer interface{}, sql string, args ...interface{}) error // See Core.GetScan.
	Union(unions ...*Model) *Model                                                              // See Core.Union.
	UnionAll(unions ...*Model) *Model                                                           // See Core.UnionAll.

	// ===========================================================================
	// 主从规范支持。
	// ===========================================================================
	// md5:f0ac82262204c704

	Master(schema ...string) (*sql.DB, error) // See Core.Master.
	Slave(schema ...string) (*sql.DB, error)  // See Core.Slave.

	// ===========================================================================
	// 乒乓游戏。
	// ===========================================================================
	// md5:548138891df7682f

	PingMaster() error // See Core.PingMaster.
	PingSlave() error  // See Core.PingSlave.

	// =============================================================================
	// 事务处理。
	// =============================================================================
	// 这里是对一个名为 "Transaction" 的部分或函数的注释，表示它与交易操作相关。
	// md5:98c80ce4a302c379

	Begin(ctx context.Context) (TX, error)                                           // See Core.Begin.
	Transaction(ctx context.Context, f func(ctx context.Context, tx TX) error) error // See Core.Transaction.

	// ===========================================================================
	// 配置方法。
	// ===========================================================================
	// md5:e4c7270c61398365

	GetCache() *gcache.Cache            // See Core.GetCache.
	SetDebug(debug bool)                // See Core.SetDebug.
	GetDebug() bool                     // See Core.GetDebug.
	GetSchema() string                  // See Core.GetSchema.
	GetPrefix() string                  // See Core.GetPrefix.
	GetGroup() string                   // See Core.GetGroup.
	SetDryRun(enabled bool)             // See Core.SetDryRun.
	GetDryRun() bool                    // See Core.GetDryRun.
	SetLogger(logger glog.ILogger)      // See Core.SetLogger.
	GetLogger() glog.ILogger            // See Core.GetLogger.
	GetConfig() *ConfigNode             // See Core.GetConfig.
	SetMaxIdleConnCount(n int)          // 参见 Core.SetMaxIdleConnCount。 md5:93a41dde27176210
	SetMaxOpenConnCount(n int)          // 参见Core.SetMaxOpenConnCount。 md5:781ba14245ef2d2f
	SetMaxConnLifeTime(d time.Duration) // 请参考Core.SetMaxConnLifeTime。 md5:9886c404ca6b5919

	// ===========================================================================
	// 辅助方法。
	// ===========================================================================
	// md5:0c5a132a773f89c0

	Stats(ctx context.Context) []StatsItem                                                                   // See Core.Stats.
	GetCtx() context.Context                                                                                 // See Core.GetCtx.
	GetCore() *Core                                                                                          // See Core.GetCore
	GetChars() (charLeft string, charRight string)                                                           // See Core.GetChars.
	Tables(ctx context.Context, schema ...string) (tables []string, err error)                               // 参见 Core.Tables。驱动程序必须实现此函数。 md5:d7f231f6b59af607
	TableFields(ctx context.Context, table string, schema ...string) (map[string]*TableField, error)         // 参见 Core.TableFields。驱动程序必须实现此函数。 md5:657c24bb39017da1
	ConvertValueForField(ctx context.Context, fieldType string, fieldValue interface{}) (interface{}, error) // 参见Core(ConvertValueForField). md5:cd3e4aabe989b5b0
	ConvertValueForLocal(ctx context.Context, fieldType string, fieldValue interface{}) (interface{}, error) // 参见Core.ConvertValueForLocal. md5:c5ed8f55d002cc9b
	CheckLocalTypeForField(ctx context.Context, fieldType string, fieldValue interface{}) (LocalType, error) // 参见 Core.CheckLocalTypeForField. md5:9dab404962da3137
	FormatUpsert(columns []string, list List, option DoInsertOption) (string, error)                         // 参见Core.DoFormatUpsert. md5:e28a610aead90684
}

// TX 定义 ORM 事务操作的接口。 md5:d71a7d0434928cac
type TX interface {
	Link

	Ctx(ctx context.Context) TX
	Raw(rawSql string, args ...interface{}) *Model
	Model(tableNameQueryOrStruct ...interface{}) *Model
	With(object interface{}) *Model

	// 如果需要，嵌套事务。
	// ===========================================================================
	// md5:96e249df6d75bc7f

	Begin() error
	Commit() error
	Rollback() error
	Transaction(ctx context.Context, f func(ctx context.Context, tx TX) error) (err error)

	// ===========================================================================
	// 核心方法。
	// ===========================================================================
	// md5:a10911bb5021107c

	Query(sql string, args ...interface{}) (result Result, err error)
	Exec(sql string, args ...interface{}) (sql.Result, error)
	Prepare(sql string) (*Stmt, error)

	// ===========================================================================
	// 查询
	// ===========================================================================
	// md5:4612a1ae72dd3cf5

	GetAll(sql string, args ...interface{}) (Result, error)
	GetOne(sql string, args ...interface{}) (Record, error)
	GetStruct(obj interface{}, sql string, args ...interface{}) error
	GetStructs(objPointerSlice interface{}, sql string, args ...interface{}) error
	GetScan(pointer interface{}, sql string, args ...interface{}) error
	GetValue(sql string, args ...interface{}) (Value, error)
	GetCount(sql string, args ...interface{}) (int64, error)

	// =============================================================================
	// CURD (Create, Read, Update, Delete) 操作。
	// =============================================================================
	// 这是Go语言中的注释，描述了一个与CRUD（创建、读取、更新、删除）操作相关的部分。在软件开发中，CURD通常用于数据库操作的基本操作。
	// md5:b9584d9a2373e908

	Insert(table string, data interface{}, batch ...int) (sql.Result, error)
	InsertIgnore(table string, data interface{}, batch ...int) (sql.Result, error)
	InsertAndGetId(table string, data interface{}, batch ...int) (int64, error)
	Replace(table string, data interface{}, batch ...int) (sql.Result, error)
	Save(table string, data interface{}, batch ...int) (sql.Result, error)
	Update(table string, data interface{}, condition interface{}, args ...interface{}) (sql.Result, error)
	Delete(table string, condition interface{}, args ...interface{}) (sql.Result, error)

	// ===========================================================================
	// 辅助方法。
	// ===========================================================================
	// md5:0c5a132a773f89c0

	GetCtx() context.Context
	GetDB() DB
	GetSqlTX() *sql.Tx
	IsClosed() bool

	// ===================================================================================
	// 保存点功能。
	// ===================================================================================
	// md5:54487b34337e4026

	SavePoint(point string) error
	RollbackTo(point string) error
}

// StatsItem 定义了配置节点的统计信息。 md5:95acda1876ad44fa
type StatsItem interface {
		// Node 返回配置节点信息。 md5:868005c0df3fa483
	Node() ConfigNode

		// Stats 返回当前节点的连接状态统计信息。 md5:b497e68c5fce778b
	Stats() sql.DBStats
}

// Core 是数据库管理的基础结构。 md5:9f4e8f0e026a368e
type Core struct {
	db            DB              // DB interface object.
	ctx           context.Context // 只用于链式操作的上下文。不要在核心初始化时设置默认值。 md5:4d22c5c34f3c1128
	group         string          // 配置组名称。 md5:50dbd64908990986
	schema        string          // 该对象的自定义架构。 md5:dd54ddfd9c22e232
	debug         *gtype.Bool     // 启用数据库的调试模式，该模式可以在运行时更改。 md5:31b723dcdbabbcb3
	cache         *gcache.Cache   // 缓存管理器，仅用于SQL结果缓存。 md5:480937caf34cae3b
	links         *gmap.Map       // links 缓存由节点创建的所有链接。 md5:a89aca4e23df0139
	logger        glog.ILogger    // 用于记录功能的日志记录器。 md5:ff375387d5036677
	config        *ConfigNode     // Current config node.
	dynamicConfig dynamicConfig   // 动态配置，这些配置可以在运行时进行更改。 md5:11c382c381ba12fc
}

type dynamicConfig struct {
	MaxIdleConnCount int
	MaxOpenConnCount int
	MaxConnLifeTime  time.Duration
}

// DoCommitInput是DoCommit函数的输入参数。 md5:151d182ffc05e6f3
type DoCommitInput struct {
	Db            *sql.DB
	Tx            *sql.Tx
	Stmt          *sql.Stmt
	Link          Link
	Sql           string
	Args          []interface{}
	Type          SqlType
	IsTransaction bool
}

// DoCommitOutput是DoCommit函数的输出参数。 md5:bb154a9d2f960894
type DoCommitOutput struct {
	Result    sql.Result  // Result 是执行语句的结果。 md5:92181818237c3bdd
	Records   []Record    // Records 是查询语句的结果。 md5:3ab79979d5bb7a15
	Stmt      *Stmt       // Stmt是Prepare的结果，是一个Statement对象。 md5:f7d8689435820710
	Tx        TX          // Tx是Begin操作的结果交易对象。 md5:388468f78948bf40
	RawResult interface{} // RawResult 是底层结果，可能是 sql.Result/*sql.Rows/*sql.Row。 md5:8f6721571bd4ebc3
}

// Driver 是将 sql 驱动程序集成到 gdb 包的接口。 md5:739e8c3911355df2
type Driver interface {
		// New 为指定的数据库服务器创建并返回一个数据库对象。 md5:27bb5dc9ab2ddbdf
	New(core *Core, node *ConfigNode) (DB, error)
}

// Link 是一个常见的数据库函数包装接口。
// 注意，使用 `Link` 进行的任何操作都不会有 SQL 日志记录。
// md5:d84360a9ae77a1de
type Link interface {
	QueryContext(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error)
	ExecContext(ctx context.Context, sql string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, sql string) (*sql.Stmt, error)
	IsOnMaster() bool
	IsTransaction() bool
}

// Sql 是用于记录 SQL 的结构体。 md5:00be6a573786ce1b
type Sql struct {
	Sql           string        // SQL字符串（可能包含保留字符'?'）。 md5:8c5af3ca9410a41d
	Type          SqlType       // SQL operation type.
	Args          []interface{} // 这个SQL的参数。 md5:b4c3922763e4978a
	Format        string        // 包含在SQL中的参数格式化的SQL。 md5:555ebca4fe98838b
	Error         error         // Execution result.
	Start         int64         // 开始执行的时间戳（以毫秒为单位）。 md5:90bcb553c20da17f
	End           int64         // 结束执行的时间戳，以毫秒为单位。 md5:a69dc8d121bd65c6
	Group         string        // Group是执行SQL的配置组的名称。 md5:ae6f0deafeed211b
	Schema        string        // Schema是执行SQL时的配置模式名称。 md5:ac146287e6f60f27
	IsTransaction bool          // IsTransaction 标记此 SQL 语句是否在事务中执行。 md5:df029c1ffc72fbf7
	RowsAffected  int64         // RowsAffected 标记了当前 SQL 语句所影响或检索到的数量。 md5:24de0b6ae028d942
}

// DoInsertOption是用于DoInsert函数的输入结构体。 md5:de18cdcf6449aa9a
type DoInsertOption struct {
	OnDuplicateStr string                 // 自定义的`on duplicated`语句字符串。 md5:1076eb09195a063c
	OnDuplicateMap map[string]interface{} // 从`OnDuplicateEx`函数中自定义的键值映射，用于`on duplicated`语句。 md5:11d4cf3d337093bb
	OnConflict     []string               // 自定义的更新或插入语句中的冲突键，如果数据库需要的话。 md5:9f554ee78fb66a28
	InsertOption   InsertOption           // 在常量值中插入操作。 md5:e5ca1b47e1d66f7a
	BatchCount     int                    // 批量插入的批次计数。 md5:015bd9ee24bd1f5c
}

// TableField 是用于表示表字段的结构体。 md5:dad00a23ddbc4525
type TableField struct {
	Index   int         // 用于排序，因为map是无序的。 md5:2c2b51c0f42d0aa5
	Name    string      // Field name.
	Type    string      // 字段类型。例如：'int(10) unsigned'，'varchar(64)'。 md5:c5cb4af28fd84cc4
	Null    bool        // 字段可以为null，也可以不为null。 md5:eecc03ab53cc06c9
	Key     string      // The index information(empty if it's not an index). Eg: PRI, MUL.
	Default interface{} // 字段的默认值。 md5:e9e6c4fce349ba5e
	Extra   string      //额外的信息。例如：auto_increment。 md5:706ba1f1653042e0
	Comment string      // Field comment.
}

// Counter 是更新计数的类型。 md5:d05ba3f2911b8013
type Counter struct {
	Field string
	Value float64
}

type (
	Raw    string                   // Raw是一个原始SQL，它不会被当作参数对待，而是直接作为SQL的一部分。 md5:cd876e2c863a974f
	Value  = *gvar.Var              // Value 是字段值的类型。 md5:f2d488e083da124a
	Record map[string]Value         // Record 是表格的行记录。 md5:f13ea17db2caaf51
	Result []Record                 // Result 是行记录数组。 md5:cfc50f7c3d051e4e
	Map    = map[string]interface{} // Map是map[string]interface{}的别名，这是最常用的映射类型。 md5:d30ae3cb84b9285e
	List   = []Map                  // List 是映射数组的类型。 md5:a6dda10906f1d599
)

type CatchSQLManager struct {
	SQLArray *garray.StrArray
	DoCommit bool // DoCommit 标记是否将提交给底层驱动器。 md5:f117644c40f63234
}

const (
	defaultModelSafe                      = false
	defaultCharset                        = `utf8`
	defaultProtocol                       = `tcp`
	unionTypeNormal                       = 0
	unionTypeAll                          = 1
	defaultMaxIdleConnCount               = 10               // 连接池中的最大空闲连接数。 md5:2d7d30a1c51e849a
	defaultMaxOpenConnCount               = 0                // 连接池中的最大打开连接数。默认无限制。 md5:aef415d0d3363e03
	defaultMaxConnLifeTime                = 30 * time.Second // 连接池中每个连接的最大生存时间，单位为秒。 md5:5b2ee66fdff9b7f6
	ctxTimeoutTypeExec                    = 0
	ctxTimeoutTypeQuery                   = 1
	ctxTimeoutTypePrepare                 = 2
	cachePrefixTableFields                = `TableFields:`
	cachePrefixSelectCache                = `SelectCache:`
	commandEnvKeyForDryRun                = "gf.gdb.dryrun"
	modelForDaoSuffix                     = `ForDao`
	dbRoleSlave                           = `slave`
	ctxKeyForDB               gctx.StrKey = `CtxKeyForDB`
	ctxKeyCatchSQL            gctx.StrKey = `CtxKeyCatchSQL`
	ctxKeyInternalProducedSQL gctx.StrKey = `CtxKeyInternalProducedSQL`

		// 类型: [用户名[:密码]@][协议[(地址)]]/数据库名[?参数1=值1&...&参数N=值N]. md5:ccc4969581f025ac
	linkPattern = `(\w+):([\w\-\$]*):(.*?)@(\w+?)\((.+?)\)/{0,1}([^\?]*)\?{0,1}(.*)`
)

type queryType int

const (
	queryTypeNormal queryType = iota
	queryTypeCount
	queryTypeValue
)

type joinOperator string

const (
	joinOperatorLeft  joinOperator = "LEFT"
	joinOperatorRight joinOperator = "RIGHT"
	joinOperatorInner joinOperator = "INNER"
)

type InsertOption int

const (
	InsertOptionDefault InsertOption = iota
	InsertOptionReplace
	InsertOptionSave
	InsertOptionIgnore
)

const (
	InsertOperationInsert      = "INSERT"
	InsertOperationReplace     = "REPLACE"
	InsertOperationIgnore      = "INSERT IGNORE"
	InsertOnDuplicateKeyUpdate = "ON DUPLICATE KEY UPDATE"
)

type SqlType string

const (
	SqlTypeBegin               SqlType = "DB.Begin"
	SqlTypeTXCommit            SqlType = "TX.Commit"
	SqlTypeTXRollback          SqlType = "TX.Rollback"
	SqlTypeExecContext         SqlType = "DB.ExecContext"
	SqlTypeQueryContext        SqlType = "DB.QueryContext"
	SqlTypePrepareContext      SqlType = "DB.PrepareContext"
	SqlTypeStmtExecContext     SqlType = "DB.Statement.ExecContext"
	SqlTypeStmtQueryContext    SqlType = "DB.Statement.QueryContext"
	SqlTypeStmtQueryRowContext SqlType = "DB.Statement.QueryRowContext"
)

type LocalType string

const (
	LocalTypeUndefined   LocalType = ""
	LocalTypeString      LocalType = "string"
	LocalTypeDate        LocalType = "date"
	LocalTypeDatetime    LocalType = "datetime"
	LocalTypeInt         LocalType = "int"
	LocalTypeUint        LocalType = "uint"
	LocalTypeInt64       LocalType = "int64"
	LocalTypeUint64      LocalType = "uint64"
	LocalTypeIntSlice    LocalType = "[]int"
	LocalTypeInt64Slice  LocalType = "[]int64"
	LocalTypeUint64Slice LocalType = "[]uint64"
	LocalTypeInt64Bytes  LocalType = "int64-bytes"
	LocalTypeUint64Bytes LocalType = "uint64-bytes"
	LocalTypeFloat32     LocalType = "float32"
	LocalTypeFloat64     LocalType = "float64"
	LocalTypeBytes       LocalType = "[]byte"
	LocalTypeBool        LocalType = "bool"
	LocalTypeJson        LocalType = "json"
	LocalTypeJsonb       LocalType = "jsonb"
)

const (
	fieldTypeBinary     = "binary"
	fieldTypeVarbinary  = "varbinary"
	fieldTypeBlob       = "blob"
	fieldTypeTinyblob   = "tinyblob"
	fieldTypeMediumblob = "mediumblob"
	fieldTypeLongblob   = "longblob"
	fieldTypeInt        = "int"
	fieldTypeTinyint    = "tinyint"
	fieldTypeSmallInt   = "small_int"
	fieldTypeSmallint   = "smallint"
	fieldTypeMediumInt  = "medium_int"
	fieldTypeMediumint  = "mediumint"
	fieldTypeSerial     = "serial"
	fieldTypeBigInt     = "big_int"
	fieldTypeBigint     = "bigint"
	fieldTypeBigserial  = "bigserial"
	fieldTypeReal       = "real"
	fieldTypeFloat      = "float"
	fieldTypeDouble     = "double"
	fieldTypeDecimal    = "decimal"
	fieldTypeMoney      = "money"
	fieldTypeNumeric    = "numeric"
	fieldTypeSmallmoney = "smallmoney"
	fieldTypeBool       = "bool"
	fieldTypeBit        = "bit"
	fieldTypeDate       = "date"
	fieldTypeDatetime   = "datetime"
	fieldTypeTimestamp  = "timestamp"
	fieldTypeTimestampz = "timestamptz"
	fieldTypeJson       = "json"
	fieldTypeJsonb      = "jsonb"
)

var (
		// instances 是实例的管理映射。 md5:4600091cea2428de
	instances = gmap.NewStrAnyMap(true)

		// driverMap管理所有自定义注册的驱动程序。 md5:625ff37f5c3fb23d
	driverMap = map[string]Driver{}

	// lastOperatorRegPattern 是一个正则表达式模式，用于匹配字符串末尾带有操作符的字符串。
	// md5:6a05c1a2b57c687b
	lastOperatorRegPattern = `[<>=]+\s*$`

	// regularFieldNameRegPattern 是一个正则表达式模式，用于匹配表格的普通字段名称的字符串。
	// md5:d18bc9e2bf6112ed
	regularFieldNameRegPattern = `^[\w\.\-]+$`

	// regularFieldNameWithCommaRegPattern 是用于匹配一个或多个表的常规字段名的正则表达式模式，这些字段名由字符','连接。
	// md5:90a0d75039f03540
	regularFieldNameWithCommaRegPattern = `^[\w\.\-,\s]+$`

	// regularFieldNameWithoutDotRegPattern 与 regularFieldNameRegPattern 类似，但不允许使用点（.）。
	// 注意，虽然一些数据库允许字段名中包含字符 '.', 但在某些情况下，这里不允许在字段名中使用 '.'，因为它与 "db.table.field" 的模式冲突。
	// md5:4a7a4427aab61aa8
	regularFieldNameWithoutDotRegPattern = `^[\w\-]+$`

	// allDryRun 为所有数据库连接设置了 dry-run 特性。
	// 它通常用于命令选项，以便于使用时带来方便。
	// md5:038bcc87fc3093b6
	allDryRun = false

		// tableFieldsMap 缓存从数据库获取的表信息。 md5:5ae26e45c71e9a09
	tableFieldsMap = gmap.NewStrAnyMap(true)
)

func init() {
		// allDryRun 从环境或命令选项中初始化。 md5:1dffa2ad4982da25
	allDryRun = gcmd.GetOptWithEnv(commandEnvKeyForDryRun, false).Bool()
}

// Register 注册自定义数据库驱动到gdb。 md5:d889e7374da12918
func Register(name string, driver Driver) error {
	driverMap[name] = newDriverWrapper(driver)
	return nil
}

// New 根据给定的配置节点创建并返回一个ORM对象。 md5:c6039a0817062f9e
func New(node ConfigNode) (db DB, err error) {
	return newDBByConfigNode(&node, "")
}

// NewByGroup 根据指定的配置组名称创建并返回一个ORM对象，带有全局配置。
// 参数`name`指定了配置组的名称，默认为DefaultGroupName。
// md5:a15dd30e999d29e5
func NewByGroup(group ...string) (db DB, err error) {
	groupName := configs.group
	if len(group) > 0 && group[0] != "" {
		groupName = group[0]
	}
	configs.RLock()
	defer configs.RUnlock()

	if len(configs.config) < 1 {
		return nil, gerror.NewCode(
			gcode.CodeInvalidConfiguration,
			"database configuration is empty, please set the database configuration before using",
		)
	}
	if _, ok := configs.config[groupName]; ok {
		var node *ConfigNode
		if node, err = getConfigNodeByGroup(groupName, true); err == nil {
			return newDBByConfigNode(node, groupName)
		}
		return nil, err
	}
	return nil, gerror.NewCodef(
		gcode.CodeInvalidConfiguration,
		`database configuration node "%s" is not found, did you misspell group name "%s" or miss the database configuration?`,
		groupName, groupName,
	)
}

// newDBByConfigNode 使用给定的配置节点和组名创建并返回一个ORM对象。
// 
// 非常注意：
// 参数`node`用于数据库的创建，而不是底层连接的创建。因此，同一组中的所有数据库类型配置应该相同。
// md5:b916b78d0af6a875
func newDBByConfigNode(node *ConfigNode, group string) (db DB, err error) {
	if node.Link != "" {
		node = parseConfigNodeLink(node)
	}
	c := &Core{
		group:  group,
		debug:  gtype.NewBool(),
		cache:  gcache.New(),
		links:  gmap.New(true),
		logger: glog.New(),
		config: node,
		dynamicConfig: dynamicConfig{
			MaxIdleConnCount: node.MaxIdleConnCount,
			MaxOpenConnCount: node.MaxOpenConnCount,
			MaxConnLifeTime:  node.MaxConnLifeTime,
		},
	}
	if v, ok := driverMap[node.Type]; ok {
		if c.db, err = v.New(c, node); err != nil {
			return nil, err
		}
		return c.db, nil
	}
	errorMsg := `cannot find database driver for specified database type "%s"`
	errorMsg += `, did you misspell type name "%s" or forget importing the database driver? `
	errorMsg += `possible reference: https://github.com/gogf/gf/tree/master/contrib/drivers`
	return nil, gerror.NewCodef(gcode.CodeInvalidConfiguration, errorMsg, node.Type, node.Type)
}

// Instance 返回用于数据库操作的实例。
// 参数 `name` 指定配置组名称，默认为 DefaultGroupName。
// md5:06c22232a9c53a60
func Instance(name ...string) (db DB, err error) {
	group := configs.group
	if len(name) > 0 && name[0] != "" {
		group = name[0]
	}
	v := instances.GetOrSetFuncLock(group, func() interface{} {
		db, err = NewByGroup(group)
		return db
	})
	if v != nil {
		return v.(DB), nil
	}
	return
}

// getConfigNodeByGroup 计算并返回给定组的配置节点。它使用权重算法内部计算值，以实现负载均衡。
//
// 参数 `master` 指定是否获取主节点，如果配置了主从结构，则在非主节点情况下获取从节点。
// md5:0e8709cfd78ceae4
func getConfigNodeByGroup(group string, master bool) (*ConfigNode, error) {
	if list, ok := configs.config[group]; ok {
				// 分离主节点和从节点配置数组。 md5:0aea1639f2f64823
		var (
			masterList = make(ConfigGroup, 0)
			slaveList  = make(ConfigGroup, 0)
		)
		for i := 0; i < len(list); i++ {
			if list[i].Role == dbRoleSlave {
				slaveList = append(slaveList, list[i])
			} else {
				masterList = append(masterList, list[i])
			}
		}
		if len(masterList) < 1 {
			return nil, gerror.NewCode(
				gcode.CodeInvalidConfiguration,
				"at least one master node configuration's need to make sense",
			)
		}
		if len(slaveList) < 1 {
			slaveList = masterList
		}
		if master {
			return getConfigNodeByWeight(masterList), nil
		} else {
			return getConfigNodeByWeight(slaveList), nil
		}
	}
	return nil, gerror.NewCodef(
		gcode.CodeInvalidConfiguration,
		"empty database configuration for item name '%s'",
		group,
	)
}

// getConfigNodeByWeight 计算配置权重并随机返回一个节点。
//
// 算法简述：
// 1. 如果我们有两个节点，它们的权重都是 1，那么权重范围是 [0, 199]；
// 2. 节点1的权重范围是 [0, 99]，节点2的权重范围是 [100, 199]，比例为 1:1；
// 3. 如果随机数是 99，那么它会选择并返回节点1。
// md5:dc1548f9e38ff89b
func getConfigNodeByWeight(cg ConfigGroup) *ConfigNode {
	if len(cg) < 2 {
		return &cg[0]
	}
	var total int
	for i := 0; i < len(cg); i++ {
		total += cg[i].Weight * 100
	}
	// 如果total为0，表示所有节点都没有配置权重属性。在这种情况下，将为每个节点的权重属性默认设置为1。
	// md5:a8625af7b996c9a2
	if total == 0 {
		for i := 0; i < len(cg); i++ {
			cg[i].Weight = 1
			total += cg[i].Weight * 100
		}
	}
		// 不包括右侧边界值。 md5:660dcac461d09c8d
	var (
		min    = 0
		max    = 0
		random = grand.N(0, total-1)
	)
	for i := 0; i < len(cg); i++ {
		max = min + cg[i].Weight*100
		if random >= min && random < max {
			// ====================================================
			// 返回ConfigNode的一个副本。
			// ====================================================
			// md5:c9cfb887df88f931
			node := ConfigNode{}
			node = cg[i]
			return &node
		}
		min = max
	}
	return nil
}

// getSqlDb 获取并返回底层的数据库连接对象。
// 参数 `master` 指定是否获取主节点连接，如果配置了主从节点。
// md5:fb885ef2d5264cdc
func (c *Core) getSqlDb(master bool, schema ...string) (sqlDb *sql.DB, err error) {
	var (
		node *ConfigNode
		ctx  = c.db.GetCtx()
	)
	if c.group != "" {
		// Load balance.
		configs.RLock()
		defer configs.RUnlock()
		// Value COPY for node.
		node, err = getConfigNodeByGroup(c.group, master)
		if err != nil {
			return nil, err
		}
	} else {
		// Value COPY for node.
		n := *c.db.GetConfig()
		node = &n
	}
	if node.Charset == "" {
		node.Charset = defaultCharset
	}
	// Changes the schema.
	nodeSchema := gutil.GetOrDefaultStr(c.schema, schema...)
	if nodeSchema != "" {
		node.Name = nodeSchema
	}
		// 更新内部数据中的配置对象。 md5:9cbb8bdfb84aa63f
	if err = c.setConfigNodeToCtx(ctx, node); err != nil {
		return
	}

		// 按节点缓存底层连接池对象。 md5:5ba47140febabd5e
	var (
		instanceCacheFunc = func() interface{} {
			if sqlDb, err = c.db.Open(node); err != nil {
				return nil
			}
			if sqlDb == nil {
				return nil
			}
			if c.dynamicConfig.MaxIdleConnCount > 0 {
				sqlDb.SetMaxIdleConns(c.dynamicConfig.MaxIdleConnCount)
			} else {
				sqlDb.SetMaxIdleConns(defaultMaxIdleConnCount)
			}
			if c.dynamicConfig.MaxOpenConnCount > 0 {
				sqlDb.SetMaxOpenConns(c.dynamicConfig.MaxOpenConnCount)
			} else {
				sqlDb.SetMaxOpenConns(defaultMaxOpenConnCount)
			}
			if c.dynamicConfig.MaxConnLifeTime > 0 {
				sqlDb.SetConnMaxLifetime(c.dynamicConfig.MaxConnLifeTime)
			} else {
				sqlDb.SetConnMaxLifetime(defaultMaxConnLifeTime)
			}
			return sqlDb
		}
				// 这里使用节点值而不是指针作为缓存键，以防出现Oracle ORA-12516错误。 md5:404d8e507e0c4548
		instanceValue = c.links.GetOrSetFuncLock(*node, instanceCacheFunc)
	)
	if instanceValue != nil && sqlDb == nil {
				// 从实例映射中读取。 md5:9cd258c405d8d50f
		sqlDb = instanceValue.(*sql.DB)
	}
	if node.Debug {
		c.db.SetDebug(node.Debug)
	}
	if node.DryRun {
		c.db.SetDryRun(node.DryRun)
	}
	return
}
