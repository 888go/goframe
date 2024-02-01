// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gdb 提供针对主流关系型数据库的 ORM 功能。
package gdb
import (
	"context"
	"database/sql"
	"fmt"
	"time"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/os/gcache"
	"github.com/888go/goframe/os/gcmd"
	"github.com/888go/goframe/os/gctx"
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/util/grand"
	"github.com/888go/goframe/util/gutil"
	)
// DB 定义了用于ORM操作的接口。
type DB interface {
// ===========================================================================
// 模型创建。
// ===========================================================================
// 这段注释是用Go语言编写的，其内容是对代码段的描述。翻译成中文后，其含义如下：
// ===========================================================================
// 模块创建部分。
// ===========================================================================
// 这里“Model creation”译为“模型创建”，表示接下来的代码将用于创建某种程序模型或数据模型。

// Model 根据给定的模式创建并返回一个新的 ORM 模型。
// 参数 `table` 可以是多个表名，也可以包含别名，例如：
// 1. 表名示例：
//    Model("user") // 用户表
//    Model("user u") // 用户表，并为表设置别名 "u"
//    Model("user, user_detail") // 用户表和用户详情表
//    Model("user u, user_detail ud") // 用户表（别名 u）和用户详情表（别名 ud）
// 2. 带有别名的表名：Model("user", "u")
// 有关更多信息，请参阅 Core.Model。
	Model(tableNameOrStruct ...interface{}) *Model

	// Raw 创建并返回一个基于原始SQL（非表）的模型。
// 通常用于嵌入原始sql语句,如:
// g.Model("user").WhereLT("created_at", gdb.Raw("now()")).All()  // SELECT * FROM `user` WHERE `created_at`<now()
// 参考文档:https://goframe.org/pages/viewpage.action?pageId=111911590&showComments=true
	Raw(rawSql string, args ...interface{}) *Model

// Schema 创建并返回一个模式（Schema）。
// 另请参阅 Core.Schema。
	Schema(schema string) *Schema

// With 根据给定对象的元数据创建并返回一个 ORM 模型。
// 也可以参考 Core.With。
	With(objects ...interface{}) *Model

// Open 通过给定的节点配置为数据库创建一个原始连接对象。
// 注意，不建议手动使用此函数。
// 另请参阅 DriverMysql.Open。
	Open(config *ConfigNode) (*sql.DB, error)

// Ctx 是一个链式函数，它创建并返回一个新的 DB 对象，该对象是对当前 DB 对象的浅复制，并且其中包含给定的上下文。
// 也可参考 Core.Ctx。
	Ctx(ctx context.Context) DB

// Close 关闭数据库并阻止新的查询开始。
// Close 之后会等待所有已在服务器上开始处理的查询完成。
//
// 关闭 DB 是罕见的操作，因为 DB 连接句柄设计意图是长期存在且被多个 goroutine 共享。
	Close(ctx context.Context) error

// ===========================================================================
// 查询API。
// ===========================================================================
// 这段注释是用英文书写的，翻译成中文后，其内容如下：
// ===========================================================================
// 查询相关的API接口。
// ===========================================================================
// 这里对代码段进行了概括性注释，表明该部分包含查询相关的API（应用程序接口）功能。

	Query(ctx context.Context, sql string, args ...interface{}) (Result, error)    // See Core.Query.
	Exec(ctx context.Context, sql string, args ...interface{}) (sql.Result, error) // See Core.Exec.
	Prepare(ctx context.Context, sql string, execOnMaster ...bool) (*Stmt, error)  // See Core.Prepare.

// ===========================================================================
// 常用的CURD API.
// ===========================================================================
// 这段注释是Go语言代码的一部分，用于描述该部分代码的功能。翻译成中文后，其含义如下：
// ===========================================================================
// 提供常用的创建（Create）、更新（Update）、读取（Read）和删除（Delete）操作的API。
// ===========================================================================

	Insert(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)                               // See Core.Insert.
	InsertIgnore(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)                         // 参见 Core.InsertIgnore。
	InsertAndGetId(ctx context.Context, table string, data interface{}, batch ...int) (int64, error)                            // 参见 Core.InsertAndGetId.
	Replace(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)                              // See Core.Replace.
	Save(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)                                 // See Core.Save.
	Update(ctx context.Context, table string, data interface{}, condition interface{}, args ...interface{}) (sql.Result, error) // See Core.Update.
	Delete(ctx context.Context, table string, condition interface{}, args ...interface{}) (sql.Result, error)                   // See Core.Delete.

// ===========================================================================
// 内部CURD API，可以被自定义的CURD实现覆盖。
// ===========================================================================

	DoSelect(ctx context.Context, link Link, sql string, args ...interface{}) (result Result, err error)                                           // 参见 Core.DoSelect。
	DoInsert(ctx context.Context, link Link, table string, data List, option DoInsertOption) (result sql.Result, err error)                        // 参见 Core.DoInsert。
	DoUpdate(ctx context.Context, link Link, table string, data interface{}, condition string, args ...interface{}) (result sql.Result, err error) // 参见 Core.DoUpdate.
	DoDelete(ctx context.Context, link Link, table string, condition string, args ...interface{}) (result sql.Result, err error)                   // 参见 Core.DoDelete。

	DoQuery(ctx context.Context, link Link, sql string, args ...interface{}) (result Result, err error)    // See Core.DoQuery.
	DoExec(ctx context.Context, link Link, sql string, args ...interface{}) (result sql.Result, err error) // See Core.DoExec.

	DoFilter(ctx context.Context, link Link, sql string, args []interface{}) (newSql string, newArgs []interface{}, err error) // 参见 Core.DoFilter。
	DoCommit(ctx context.Context, in DoCommitInput) (out DoCommitOutput, err error)                                            // 参见 Core.DoCommit。

	DoPrepare(ctx context.Context, link Link, sql string) (*Stmt, error) // 参见 Core.DoPrepare。

// ===========================================================================
// 为了方便起见，提供的查询APIs。
// ===========================================================================

	GetAll(ctx context.Context, sql string, args ...interface{}) (Result, error)                // See Core.GetAll.
	GetOne(ctx context.Context, sql string, args ...interface{}) (Record, error)                // See Core.GetOne.
	GetValue(ctx context.Context, sql string, args ...interface{}) (Value, error)               // 参见 Core.GetValue。
	GetArray(ctx context.Context, sql string, args ...interface{}) ([]Value, error)             // 参见 Core.GetArray.
	GetCount(ctx context.Context, sql string, args ...interface{}) (int, error)                 // 参见 Core.GetCount。
	GetScan(ctx context.Context, objPointer interface{}, sql string, args ...interface{}) error // See Core.GetScan.
	Union(unions ...*Model) *Model                                                              // See Core.Union.
	UnionAll(unions ...*Model) *Model                                                           // 参见 Core.UnionAll。

// ===========================================================================
// 主从模式支持。
// ===========================================================================
// 这段注释是用于描述Go语言代码中关于主从（Master/Slave）规范或模式的相关实现。主从模式通常是指在分布式系统中，存在一个主节点负责处理写入操作以及数据同步，而从节点则主要用于读取操作和备份数据的场景。

	Master(schema ...string) (*sql.DB, error) // See Core.Master.
	Slave(schema ...string) (*sql.DB, error)  // See Core.Slave.

// ===========================================================================
// 乒乓球.
// ===========================================================================
// 这段 Go 语言代码的注释表明这是一个关于“Ping-Pong”的模块或功能，但没有提供具体的代码实现细节。这里的注释翻译成中文后，其含义不变，仍然是对这一部分功能或模块的描述，表示与乒乓球游戏或者网络中的 Ping-Pong（心跳检测）机制相关的代码。

	PingMaster() error // 参见 Core.PingMaster.
	PingSlave() error  // 参见 Core.PingSlave.

// ===========================================================================
// 事务。
// ===========================================================================

	Begin(ctx context.Context) (TX, error)                                           // See Core.Begin.
	Transaction(ctx context.Context, f func(ctx context.Context, tx TX) error) error // 参见Core.Transaction.

// ===========================================================================
// 配置方法。
// ===========================================================================
// 这段 Go 语言代码注释表明接下来的代码是关于配置相关的方法，用于对程序或服务进行配置。

	GetCache() *gcache.Cache            // 参见 Core.GetCache。
	SetDebug(debug bool)                // 参见 Core.SetDebug.
	GetDebug() bool                     // 参见 Core.GetDebug。
	GetSchema() string                  // 参见 Core.GetSchema。
	GetPrefix() string                  // 参见 Core.GetPrefix。
	GetGroup() string                   // 参见 Core.GetGroup。
	SetDryRun(enabled bool)             // 参见 Core.SetDryRun。
	GetDryRun() bool                    // 参见 Core.GetDryRun。
	SetLogger(logger glog.ILogger)      // 参见 Core.SetLogger。
	GetLogger() glog.ILogger            // 参见 Core.GetLogger。
	GetConfig() *ConfigNode             // 参见 Core.GetConfig。
	SetMaxIdleConnCount(n int)          // 参见 Core.SetMaxIdleConnCount.
	SetMaxOpenConnCount(n int)          // 参见 Core.SetMaxOpenConnCount.
	SetMaxConnLifeTime(d time.Duration) // 参见 Core.SetMaxConnLifeTime.

// ===========================================================================
// 工具方法。
// ===========================================================================
// 这段注释是用英语书写的，以下是翻译成中文的版本：
// ===========================================================================
// 实用工具函数集合。
// ===========================================================================
// 这里是对Golang代码中一组工具方法的注释描述，表示这一部分包含一些通用、便捷的辅助函数。

	GetCtx() context.Context                                                                                 // See Core.GetCtx.
	GetCore() *Core                                                                                          // See Core.GetCore
	GetChars() (charLeft string, charRight string)                                                           // 参见 Core.GetChars。
	Tables(ctx context.Context, schema ...string) (tables []string, err error)                               // 查看Core.Tables。驱动程序必须实现这个函数。
	TableFields(ctx context.Context, table string, schema ...string) (map[string]*TableField, error)         // 查看 Core.TableFields。驱动程序必须实现此函数。
	ConvertValueForField(ctx context.Context, fieldType string, fieldValue interface{}) (interface{}, error) // 查看 Core.ConvertValueForField
	ConvertValueForLocal(ctx context.Context, fieldType string, fieldValue interface{}) (interface{}, error) // 参见 Core.ConvertValueForLocal
	CheckLocalTypeForField(ctx context.Context, fieldType string, fieldValue interface{}) (LocalType, error) // 查看 Core.CheckLocalTypeForField
}

// TX 定义了用于ORM事务操作的接口。
type TX interface {
	Link

	Ctx(ctx context.Context) TX
	Raw(rawSql string, args ...interface{}) *Model
	Model(tableNameQueryOrStruct ...interface{}) *Model
	With(object interface{}) *Model

// ===========================================================================
// 如果有必要，进行嵌套事务。
// ===========================================================================

	Begin() error
	Commit() error
	Rollback() error
	Transaction(ctx context.Context, f func(ctx context.Context, tx TX) error) (err error)

// ===========================================================================
// 核心方法
// ===========================================================================
// 这段代码中的注释是将英文注释翻译成中文，其含义如下：
// ```go
// ===========================================================================
// 核心方法
// ===========================================================================
// 这个注释是对接下来要定义或实现的 Go 语言代码段的描述，表示这部分代码是整个程序或模块的核心功能方法。

	Query(sql string, args ...interface{}) (result Result, err error)
	Exec(sql string, args ...interface{}) (sql.Result, error)
	Prepare(sql string) (*Stmt, error)

// ===========================================================================
// 查询。
// ===========================================================================

	GetAll(sql string, args ...interface{}) (Result, error)
	GetOne(sql string, args ...interface{}) (Record, error)
	GetStruct(obj interface{}, sql string, args ...interface{}) error
	GetStructs(objPointerSlice interface{}, sql string, args ...interface{}) error
	GetScan(pointer interface{}, sql string, args ...interface{}) error
	GetValue(sql string, args ...interface{}) (Value, error)
	GetCount(sql string, args ...interface{}) (int64, error)

// ===========================================================================
// CURD（增删改查操作）
// ===========================================================================

	Insert(table string, data interface{}, batch ...int) (sql.Result, error)
	InsertIgnore(table string, data interface{}, batch ...int) (sql.Result, error)
	InsertAndGetId(table string, data interface{}, batch ...int) (int64, error)
	Replace(table string, data interface{}, batch ...int) (sql.Result, error)
	Save(table string, data interface{}, batch ...int) (sql.Result, error)
	Update(table string, data interface{}, condition interface{}, args ...interface{}) (sql.Result, error)
	Delete(table string, condition interface{}, args ...interface{}) (sql.Result, error)

// ===========================================================================
// 工具方法。
// ===========================================================================
// 这段注释是用英语书写的，以下是翻译成中文的版本：
// ===========================================================================
// 实用工具函数集合。
// ===========================================================================
// 这里是对Golang代码中一组工具方法的注释描述，表示这一部分包含一些通用、便捷的辅助函数。

	GetCtx() context.Context
	GetDB() DB
	GetSqlTX() *sql.Tx
	IsClosed() bool

// ===========================================================================
// 保存点功能。
// ===========================================================================

	SavePoint(point string) error
	RollbackTo(point string) error
}

// Core是数据库管理的基础结构体。
type Core struct {
	db            DB              // DB 接口对象。
	ctx           context.Context // 此上下文仅用于链式操作。请勿在 Core 初始化时设置默认值。
	group         string          // 配置组名称。
	schema        string          // 为此对象定制的自定义模式。
	debug         *gtype.Bool     // 启用数据库的调试模式，该模式可以在运行时进行更改。
	cache         *gcache.Cache   // 缓存管理器，仅用于SQL结果缓存。
	links         *gmap.StrAnyMap // links 缓存了所有已创建的按节点链接。
	logger        glog.ILogger    // Logger 用于提供日志记录功能。
	config        *ConfigNode     // 当前配置节点。
	dynamicConfig dynamicConfig   // 动态配置，可以在运行时进行更改。
}

type dynamicConfig struct {
	MaxIdleConnCount int
	MaxOpenConnCount int
	MaxConnLifeTime  time.Duration
}

// DoCommitInput 是函数 DoCommit 的输入参数。
type DoCommitInput struct {
	Db            *sql.DB
	Tx            *sql.Tx
	Stmt          *sql.Stmt
	Link          Link
	Sql           string
	Args          []interface{}
	Type          string
	IsTransaction bool
}

// DoCommitOutput 是函数 DoCommit 的输出参数。
type DoCommitOutput struct {
	Result    sql.Result  // Result 是执行语句的结果。
	Records   []Record    // Records 是查询语句的结果。
	Stmt      *Stmt       // Stmt是Prepare方法执行后返回的Statement对象结果。
	Tx        TX          // Tx是Begin方法返回的事务对象。
	RawResult interface{} // RawResult 是底层结果，它可能是 sql.Result、*sql.Rows 或 *sql.Row。
}

// Driver 是用于将 SQL 驱动程序集成到 gdb 包的接口。
type Driver interface {
	// New 创建并返回指定数据库服务器的数据库对象。
	New(core *Core, node *ConfigNode) (DB, error)
}

// Link 是一个通用的数据库函数包装器接口。
// 注意，使用 `Link` 进行的任何操作将不会有 SQL 日志记录。
type Link interface {
	QueryContext(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error)
	ExecContext(ctx context.Context, sql string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, sql string) (*sql.Stmt, error)
	IsOnMaster() bool
	IsTransaction() bool
}

// Sql 是用于记录 SQL 的结构体。
type Sql struct {
	Sql           string        // SQL 字符串（可能包含保留字符 '?'）。
	Type          string        // SQL操作类型。
	Args          []interface{} // 此SQL的参数。
	Format        string        // 格式化后的SQL，其中包含在SQL中的参数。
	Error         error         // Execution result.
	Start         int64         // 开始执行的时间戳（毫秒）。
	End           int64         // 结束执行的时间戳（毫秒）。
	Group         string        // Group 是执行 SQL 时所使用的配置组名称。
	Schema        string        // Schema 是执行 SQL 的配置的架构名称。
	IsTransaction bool          // IsTransaction 标记了这个SQL语句是否在事务中执行。
	RowsAffected  int64         // RowsAffected 标记了当前SQL语句执行后获取或影响的行数。
}

// DoInsertOption 是函数 DoInsert 的输入结构体。
type DoInsertOption struct {
	OnDuplicateStr string                 // 自定义用于`on duplicated`语句的字符串。
	OnDuplicateMap map[string]interface{} // `OnDuplicateEx`函数为`on duplicated`语句提供的自定义键值对映射
	InsertOption   InsertOption           // 在常数值中执行插入操作。
	BatchCount     int                    // 批量插入的批次数量
}

// TableField 是用于表示表格字段的结构体。
type TableField struct {
	Index   int         // 用于排序目的，因为映射（map）是无序的。
	Name    string      // Field name.
	Type    string      // 字段类型。例如：'int(10) unsigned', 'varchar(64)'。
// 这段注释是对Go语言代码中某个表示字段类型的变量或常量的解释，该字段在数据库表结构设计中使用，比如MySQL等关系型数据库。'int(10) unsigned' 表示一个无符号整数类型，长度为10位；'varchar(64)' 则表示变长字符串类型，最大长度为64个字符。
	Null    bool        // 字段可以为空或非空
	Key     string      // The index information(empty if it's not an index). Eg: PRI, MUL.
	Default interface{} // 字段的默认值。
	Extra   string      // 额外信息。例如：自动增长。
	Comment string      // Field comment.
}

// Counter 是用于更新计数的类型。
type Counter struct {
	Field string
	Value float64
}

type (
	Raw    string                   // Raw 是一个原始SQL语句，它不会被视为参数处理，而是直接作为SQL部分。
// 通常用于嵌入原始sql语句,如:
// g.Model("user").WhereLT("created_at", gdb.Raw("now()")).All()  // SELECT * FROM `user` WHERE `created_at`<now()
// 参考文档:https://goframe.org/pages/viewpage.action?pageId=111911590&showComments=true
	Value  = *gvar.Var              // Value 是字段值类型。
	Record map[string]Value         // Record 是表格的行记录。
	Result []Record                 // Result 是行记录数组。
	Map    = map[string]interface{} // Map 是 map[string]interface{} 的别名，这是最常用的映射类型。
	List   = []Map                  // List 是映射数组的类型。
)

type CatchSQLManager struct {
	SQLArray *garray.StrArray
	DoCommit bool // DoCommit 标记是否将提交到底层驱动。
}

const (
	defaultModelSafe                      = false
	defaultCharset                        = `utf8`
	defaultProtocol                       = `tcp`
	unionTypeNormal                       = 0
	unionTypeAll                          = 1
	defaultMaxIdleConnCount               = 10               // 在连接池中的最大空闲连接数。
	defaultMaxOpenConnCount               = 0                // 在连接池中的最大打开连接数。默认是没有限制。
	defaultMaxConnLifeTime                = 30 * time.Second // 连接池中每个连接的最大生命周期，单位为秒。
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

	// 定义数据库连接格式：
// type: [用户名[:密码]@][协议[(地址)]]/数据库名[?参数1=值1&...&参数N=值N]
// 其中各部分的含义：
// - type：表示数据库类型，如mysql、postgres等
// - username: 可选，用于登录数据库的用户名
// - password: 可选，对应用户名的密码，通常会进行编码处理
// - protocol: 数据库访问协议，如tcp、unix等
// - address: 协议对应的服务器地址或socket路径
// - dbname: 需要连接的数据库名称
// - param1=value1,...,paramN=valueN: 可选，一系列键值对形式的连接参数，例如charset=utf8、sslmode=disable等
	linkPattern = `(\w+):([\w\-\$]*):(.*?)@(\w+?)\((.+?)\)/{0,1}([^\?]*)\?{0,1}(.*)`
)

type queryType int

const (
	queryTypeNormal queryType = 0
	queryTypeCount  queryType = 1
	queryTypeValue  queryType = 2
)

type joinOperator string

const (
	joinOperatorLeft  joinOperator = "LEFT"
	joinOperatorRight joinOperator = "RIGHT"
	joinOperatorInner joinOperator = "INNER"
)

type InsertOption int

const (
	InsertOptionDefault        InsertOption = 0
	InsertOptionReplace        InsertOption = 1
	InsertOptionSave           InsertOption = 2
	InsertOptionIgnore         InsertOption = 3
	InsertOperationInsert                   = "INSERT"
	InsertOperationReplace                  = "REPLACE"
	InsertOperationIgnore                   = "INSERT IGNORE"
	InsertOnDuplicateKeyUpdate              = "ON DUPLICATE KEY UPDATE"
)

const (
	SqlTypeBegin               = "DB.Begin"
	SqlTypeTXCommit            = "TX.Commit"
	SqlTypeTXRollback          = "TX.Rollback"
	SqlTypeExecContext         = "DB.ExecContext"
	SqlTypeQueryContext        = "DB.QueryContext"
	SqlTypePrepareContext      = "DB.PrepareContext"
	SqlTypeStmtExecContext     = "DB.Statement.ExecContext"
	SqlTypeStmtQueryContext    = "DB.Statement.QueryContext"
	SqlTypeStmtQueryRowContext = "DB.Statement.QueryRowContext"
)

type LocalType string

const (
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
	// instances 是用于实例管理的映射（map）。
	instances = gmap.NewStrAnyMap(true)

	// driverMap 管理所有已注册的自定义驱动。
	driverMap = map[string]Driver{}

// lastOperatorRegPattern 是正则表达式模式，用于表示字符串尾部包含操作符的字符串。
	lastOperatorRegPattern = `[<>=]+\s*$`

// regularFieldNameRegPattern 是用于表示表中常规字段名的字符串的正则表达式模式。
	regularFieldNameRegPattern = `^[\w\.\-]+$`

// regularFieldNameWithoutDotRegPattern 与 regularFieldNameRegPattern 类似，但不允许包含“.”字符。
// 注意，尽管某些数据库允许字段名中包含字符'.'，但在这里不允许可字段名中出现'.'，因为在某些情况下它会与 "db.table.field" 的模式产生冲突。
	regularFieldNameWithoutDotRegPattern = `^[\w\-]+$`

// allDryRun 为所有数据库连接设置模拟执行（dry-run）功能。
// 通常为了方便，它被用于命令行选项中。
	allDryRun = false

	// tableFieldsMap 缓存从数据库获取的表信息。
	tableFieldsMap = gmap.NewStrAnyMap(true)
)

func init() {
	// allDryRun 从环境变量或命令选项中初始化。
	allDryRun = gcmd.GetOptWithEnv(commandEnvKeyForDryRun, false).Bool()
}

// Register 注册自定义数据库驱动到 gdb。
func Register(name string, driver Driver) error {
	driverMap[name] = newDriverWrapper(driver)
	return nil
}

// New 根据给定的配置节点创建并返回一个ORM对象。
func New(node ConfigNode) (db DB, err error) {
	return newDBByConfigNode(&node, "")
}

// NewByGroup 根据全局配置创建并返回一个 ORM 对象。
// 参数 `name` 指定配置组名称，默认为 DefaultGroupName。
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

// newDBByConfigNode 根据给定的配置节点和组名创建并返回一个ORM对象。
//
// **非常重要**：
// 参数`node`用于数据库创建，而非底层连接创建，
// 因此在同一组内的所有数据库类型配置应当保持一致。
func newDBByConfigNode(node *ConfigNode, group string) (db DB, err error) {
	if node.Link != "" {
		node = parseConfigNodeLink(node)
	}
	c := &Core{
		group:  group,
		debug:  gtype.NewBool(),
		cache:  gcache.New(),
		links:  gmap.NewStrAnyMap(true),
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

// Instance 返回一个用于数据库操作的实例。
// 参数 `name` 指定了配置组名称，默认为 DefaultGroupName。
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

// 根据给定组获取配置节点，通过内部计算并返回。它使用权重算法进行负载均衡计算。
//
// 参数`master`指定是否获取主节点，如果不是，则在主从配置情况下获取从节点。
func getConfigNodeByGroup(group string, master bool) (*ConfigNode, error) {
	if list, ok := configs.config[group]; ok {
		// 将主节点和从节点配置数组分离。
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

// getConfigNodeByWeight 根据权重计算配置，并随机返回一个节点。
//
// 计算算法简述：
// 1. 若我们有2个节点，它们的权重都为1，则权重范围是 [0, 199]；
// 2. 节点1的权重范围是 [0, 99]，节点2的权重范围是 [100, 199]，比例为 1:1；
// 3. 如果随机数是99，则选择并返回节点1。
func getConfigNodeByWeight(cg ConfigGroup) *ConfigNode {
	if len(cg) < 2 {
		return &cg[0]
	}
	var total int
	for i := 0; i < len(cg); i++ {
		total += cg[i].Weight * 100
	}
// 如果total为0，表示所有节点都没有配置权重属性。
// 此时，默认将每个节点的权重属性设为1。
	if total == 0 {
		for i := 0; i < len(cg); i++ {
			cg[i].Weight = 1
			total += cg[i].Weight * 100
		}
	}
	// 排除右侧边界值。
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
			node := ConfigNode{}
			node = cg[i]
			return &node
		}
		min = max
	}
	return nil
}

// getSqlDb 获取并返回底层数据库连接对象。
// 参数 `master` 指定在配置了主从节点的情况下，是否获取主节点的连接。
func (c *Core) getSqlDb(master bool, schema ...string) (sqlDb *sql.DB, err error) {
	var (
		node *ConfigNode
		ctx  = c.db.GetCtx()
	)
	if c.group != "" {
		// Load balance.
		configs.RLock()
		defer configs.RUnlock()
		// Value COPY 表示节点的复制值。
		node, err = getConfigNodeByGroup(c.group, master)
		if err != nil {
			return nil, err
		}
	} else {
		// Value COPY 表示节点的复制值。
		n := *c.db.GetConfig()
		node = &n
	}
	if node.Charset == "" {
		node.Charset = defaultCharset
	}
	// 修改模式
	nodeSchema := gutil.GetOrDefaultStr(c.schema, schema...)
	if nodeSchema != "" {
		node.Name = nodeSchema
	}
	// 更新内部数据中的配置对象。
	internalData := c.GetInternalCtxDataFromCtx(ctx)
	if internalData != nil {
		internalData.ConfigNode = node
	}
	// 通过节点缓存底层连接池对象。
	instanceNameByNode := fmt.Sprintf(`%+v`, node)
	instanceValue := c.links.GetOrSetFuncLock(instanceNameByNode, func() interface{} {
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
	})
	if instanceValue != nil && sqlDb == nil {
		// 它从实例映射中读取。
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
