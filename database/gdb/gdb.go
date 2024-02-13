// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gdb 提供针对主流关系型数据库的 ORM 功能。
package db类

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
	X创建Model对象(tableNameOrStruct ...interface{}) *Model

	// Raw 创建并返回一个基于原始SQL（非表）的模型。
// 通常用于嵌入原始sql语句,如:
// g.Model("user").WhereLT("created_at", gdb.Raw("now()")).All()  // SELECT * FROM `user` WHERE `created_at`<now()
// 参考文档:https://goframe.org/pages/viewpage.action?pageId=111911590&showComments=true
	X原生SQL(rawSql string, args ...interface{}) *Model

// Schema 创建并返回一个模式（Schema）。
// 另请参阅 Core.Schema。
	X切换数据库(schema string) *Schema

// With 根据给定对象的元数据创建并返回一个 ORM 模型。
// 也可以参考 Core.With。
	X关联对象(objects ...interface{}) *Model

// Open 通过给定的节点配置为数据库创建一个原始连接对象。
// 注意，不建议手动使用此函数。
// 另请参阅 DriverMysql.Open。
	X底层Open(config *X配置项) (*sql.DB, error)

// Ctx 是一个链式函数，它创建并返回一个新的 DB 对象，该对象是对当前 DB 对象的浅复制，并且其中包含给定的上下文。
// 也可参考 Core.Ctx。
	X设置上下文并取副本(ctx context.Context) DB

// Close 关闭数据库并阻止新的查询开始。
// Close 之后会等待所有已在服务器上开始处理的查询完成。
//
// 关闭 DB 是罕见的操作，因为 DB 连接句柄设计意图是长期存在且被多个 goroutine 共享。
	X关闭数据库(ctx context.Context) error

// ===========================================================================
// 查询API。
// ===========================================================================
// 这段注释是用英文书写的，翻译成中文后，其内容如下：
// ===========================================================================
// 查询相关的API接口。
// ===========================================================================
// 这里对代码段进行了概括性注释，表明该部分包含查询相关的API（应用程序接口）功能。

	X原生SQL查询(ctx context.Context, sql string, args ...interface{}) (X行记录数组, error)    // See Core.Query.
	X原生SQL执行(ctx context.Context, sql string, args ...interface{}) (sql.Result, error) // See Core.Exec.
	X原生sql取参数预处理对象(ctx context.Context, sql string, execOnMaster ...bool) (*Stmt, error)  // See Core.Prepare.

// ===========================================================================
// 常用的CURD API.
// ===========================================================================
// 这段注释是Go语言代码的一部分，用于描述该部分代码的功能。翻译成中文后，其含义如下：
// ===========================================================================
// 提供常用的创建（Create）、更新（Update）、读取（Read）和删除（Delete）操作的API。
// ===========================================================================

	X插入(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)                               // See Core.Insert.
	X插入并跳过已存在(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)                         // 参见 Core.InsertIgnore。
	X插入并取ID(ctx context.Context, table string, data interface{}, batch ...int) (int64, error)                            // 参见 Core.InsertAndGetId.
	X插入并替换已存在(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)                              // See Core.Replace.
	X插入并更新已存在(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)                                 // See Core.Save.
	X更新(ctx context.Context, table string, data interface{}, condition interface{}, args ...interface{}) (sql.Result, error) // See Core.Update.
	X删除(ctx context.Context, table string, condition interface{}, args ...interface{}) (sql.Result, error)                   // See Core.Delete.

// ===========================================================================
// 内部CURD API，可以被自定义的CURD实现覆盖。
// ===========================================================================

	X底层查询(ctx context.Context, link X底层链接, sql string, args ...interface{}) (result X行记录数组, err error)                                           // 参见 Core.DoSelect。
	X底层插入(ctx context.Context, link X底层链接, table string, data Map数组, option X底层输入) (result sql.Result, err error)                        // 参见 Core.DoInsert。
	X底层更新(ctx context.Context, link X底层链接, table string, data interface{}, condition string, args ...interface{}) (result sql.Result, err error) // 参见 Core.DoUpdate.
	X底层删除(ctx context.Context, link X底层链接, table string, condition string, args ...interface{}) (result sql.Result, err error)                   // 参见 Core.DoDelete。

	X底层原生SQL查询(ctx context.Context, link X底层链接, sql string, args ...interface{}) (result X行记录数组, err error)    // See Core.DoQuery.
	X底层原生SQL执行(ctx context.Context, link X底层链接, sql string, args ...interface{}) (result sql.Result, err error) // See Core.DoExec.

	X底层DoFilter(ctx context.Context, link X底层链接, sql string, args []interface{}) (newSql string, newArgs []interface{}, err error) // 参见 Core.DoFilter。
	X底层DoCommit(ctx context.Context, in X输入) (out X输出, err error)                                            // 参见 Core.DoCommit。

	X底层原生sql参数预处理对象(ctx context.Context, link X底层链接, sql string) (*Stmt, error) // 参见 Core.DoPrepare。

// ===========================================================================
// 为了方便起见，提供的查询APIs。
// ===========================================================================

	GetAll别名(ctx context.Context, sql string, args ...interface{}) (X行记录数组, error)                // See Core.GetAll.
	X原生SQL查询单条记录(ctx context.Context, sql string, args ...interface{}) (X行记录, error)                // See Core.GetOne.
	X原生SQL查询字段值(ctx context.Context, sql string, args ...interface{}) (X字段值, error)               // 参见 Core.GetValue。
	X原生SQL查询数组(ctx context.Context, sql string, args ...interface{}) ([]X字段值, error)             // 参见 Core.GetArray.
	X原生SQL查询字段计数(ctx context.Context, sql string, args ...interface{}) (int, error)                 // 参见 Core.GetCount。
	X原生SQL查询到结构体指针(ctx context.Context, objPointer interface{}, sql string, args ...interface{}) error // See Core.GetScan.
	X多表去重查询(unions ...*Model) *Model                                                              // See Core.Union.
	X多表查询(unions ...*Model) *Model                                                           // 参见 Core.UnionAll。

// ===========================================================================
// 主从模式支持。
// ===========================================================================
// 这段注释是用于描述Go语言代码中关于主从（Master/Slave）规范或模式的相关实现。主从模式通常是指在分布式系统中，存在一个主节点负责处理写入操作以及数据同步，而从节点则主要用于读取操作和备份数据的场景。

	X取主节点对象(schema ...string) (*sql.DB, error) // See Core.Master.
	X取从节点对象(schema ...string) (*sql.DB, error)  // See Core.Slave.

// ===========================================================================
// 乒乓球.
// ===========================================================================
// 这段 Go 语言代码的注释表明这是一个关于“Ping-Pong”的模块或功能，但没有提供具体的代码实现细节。这里的注释翻译成中文后，其含义不变，仍然是对这一部分功能或模块的描述，表示与乒乓球游戏或者网络中的 Ping-Pong（心跳检测）机制相关的代码。

	X向主节点发送心跳() error // 参见 Core.PingMaster.
	X向从节点发送心跳() error  // 参见 Core.PingSlave.

// ===========================================================================
// 事务。
// ===========================================================================

	X事务开启(ctx context.Context) (X事务, error)                                           // See Core.Begin.
	X事务(ctx context.Context, f func(ctx context.Context, tx X事务) error) error // 参见Core.Transaction.

// ===========================================================================
// 配置方法。
// ===========================================================================
// 这段 Go 语言代码注释表明接下来的代码是关于配置相关的方法，用于对程序或服务进行配置。

	X取缓存对象() *缓存类.Cache            // 参见 Core.GetCache。
	X设置调试模式(debug bool)                // 参见 Core.SetDebug.
	X取调试模式() bool                     // 参见 Core.GetDebug。
	X取默认数据库名称() string                  // 参见 Core.GetSchema。
	X取表前缀() string                  // 参见 Core.GetPrefix。
	X取配置组名称() string                   // 参见 Core.GetGroup。
	X设置空跑特性(enabled bool)             // 参见 Core.SetDryRun。
	X取空跑特性() bool                    // 参见 Core.GetDryRun。
	X设置日志记录器(logger 日志类.ILogger)      // 参见 Core.SetLogger。
	X取日志记录器() 日志类.ILogger            // 参见 Core.GetLogger。
	X取当前节点配置() *X配置项             // 参见 Core.GetConfig。
	X设置最大闲置连接数(n int)          // 参见 Core.SetMaxIdleConnCount.
	X设置最大打开连接数(n int)          // 参见 Core.SetMaxOpenConnCount.
	X设置最大空闲时长(d time.Duration) // 参见 Core.SetMaxConnLifeTime.

// ===========================================================================
// 工具方法。
// ===========================================================================
// 这段注释是用英语书写的，以下是翻译成中文的版本：
// ===========================================================================
// 实用工具函数集合。
// ===========================================================================
// 这里是对Golang代码中一组工具方法的注释描述，表示这一部分包含一些通用、便捷的辅助函数。

	X取上下文对象() context.Context                                                                                 // See Core.GetCtx.
	X取Core对象() *Core                                                                                          // See Core.GetCore
	X底层取数据库安全字符() (charLeft string, charRight string)                                                           // 参见 Core.GetChars。
	X取表名称数组(ctx context.Context, schema ...string) (tables []string, err error)                               // 查看Core.Tables。驱动程序必须实现这个函数。
	X取表字段信息Map(ctx context.Context, table string, schema ...string) (map[string]*X字段信息, error)         // 查看 Core.TableFields。驱动程序必须实现此函数。
	X底层ConvertValueForField(ctx context.Context, fieldType string, fieldValue interface{}) (interface{}, error) // 查看 Core.ConvertValueForField
	X底层ConvertValueForLocal(ctx context.Context, fieldType string, fieldValue interface{}) (interface{}, error) // 参见 Core.ConvertValueForLocal
	X底层CheckLocalTypeForField(ctx context.Context, fieldType string, fieldValue interface{}) (LocalType, error) // 查看 Core.CheckLocalTypeForField
}

// TX 定义了用于ORM事务操作的接口。
type X事务 interface {
	X底层链接

	X设置上下文并取副本(ctx context.Context) X事务
	X原生SQL(rawSql string, args ...interface{}) *Model
	X创建Model对象(tableNameQueryOrStruct ...interface{}) *Model
	X关联对象(object interface{}) *Model

// ===========================================================================
// 如果有必要，进行嵌套事务。
// ===========================================================================

	X事务开启() error
	X事务提交() error
	X事务回滚() error
	X事务(ctx context.Context, f func(ctx context.Context, tx X事务) error) (err error)

// ===========================================================================
// 核心方法
// ===========================================================================
// 这段代码中的注释是将英文注释翻译成中文，其含义如下：
// ```go
// ===========================================================================
// 核心方法
// ===========================================================================
// 这个注释是对接下来要定义或实现的 Go 语言代码段的描述，表示这部分代码是整个程序或模块的核心功能方法。

	X原生SQL查询(sql string, args ...interface{}) (result X行记录数组, err error)
	X原生SQL执行(sql string, args ...interface{}) (sql.Result, error)
	X原生sql取参数预处理对象(sql string) (*Stmt, error)

// ===========================================================================
// 查询。
// ===========================================================================

	GetAll别名(sql string, args ...interface{}) (X行记录数组, error)
	X原生SQL查询单条记录(sql string, args ...interface{}) (X行记录, error)
	X原生SQL查询单条到结构体指针(obj interface{}, sql string, args ...interface{}) error
	X原生SQL查询到结构体数组指针(objPointerSlice interface{}, sql string, args ...interface{}) error
	X原生SQL查询到结构体指针(pointer interface{}, sql string, args ...interface{}) error
	X原生SQL查询字段值(sql string, args ...interface{}) (X字段值, error)
	X原生SQL查询字段计数(sql string, args ...interface{}) (int64, error)

// ===========================================================================
// CURD（增删改查操作）
// ===========================================================================

	X插入(table string, data interface{}, batch ...int) (sql.Result, error)
	X插入并跳过已存在(table string, data interface{}, batch ...int) (sql.Result, error)
	X插入并取ID(table string, data interface{}, batch ...int) (int64, error)
	X插入并替换已存在(table string, data interface{}, batch ...int) (sql.Result, error)
	X插入并更新已存在(table string, data interface{}, batch ...int) (sql.Result, error)
	X更新(table string, data interface{}, condition interface{}, args ...interface{}) (sql.Result, error)
	X删除(table string, condition interface{}, args ...interface{}) (sql.Result, error)

// ===========================================================================
// 工具方法。
// ===========================================================================
// 这段注释是用英语书写的，以下是翻译成中文的版本：
// ===========================================================================
// 实用工具函数集合。
// ===========================================================================
// 这里是对Golang代码中一组工具方法的注释描述，表示这一部分包含一些通用、便捷的辅助函数。

	X取上下文对象() context.Context
	X取DB对象() DB
	X底层取事务对象() *sql.Tx
	X是否已关闭() bool

// ===========================================================================
// 保存点功能。
// ===========================================================================

	X保存事务点(point string) error
	X回滚事务点(point string) error
}

// Core是数据库管理的基础结构体。
type Core struct {
	db            DB              // DB 接口对象。
	ctx           context.Context // 此上下文仅用于链式操作。请勿在 Core 初始化时设置默认值。
	group         string          // 配置组名称。
	schema        string          // 为此对象定制的自定义模式。
	debug         *安全变量类.Bool     // 启用数据库的调试模式，该模式可以在运行时进行更改。
	cache         *缓存类.Cache   // 缓存管理器，仅用于SQL结果缓存。
	links         *map类.StrAnyMap // links 缓存了所有已创建的按节点链接。
	logger        日志类.ILogger    // Logger 用于提供日志记录功能。
	config        *X配置项     // 当前配置节点。
	dynamicConfig dynamicConfig   // 动态配置，可以在运行时进行更改。
}

type dynamicConfig struct {
	MaxIdleConnCount int
	MaxOpenConnCount int
	MaxConnLifeTime  time.Duration
}

// DoCommitInput 是函数 DoCommit 的输入参数。
type X输入 struct {
	Db            *sql.DB
	Tx            *sql.Tx
	Stmt          *sql.Stmt
	Link          X底层链接
	Sql           string
	Args          []interface{}
	X类型          string
	IsTransaction bool
}

// DoCommitOutput 是函数 DoCommit 的输出参数。
type X输出 struct {
	X原生sql行记录    sql.Result  // Result 是执行语句的结果。
	X行记录数组   []X行记录    // Records 是查询语句的结果。
	X参数预处理      *Stmt       // Stmt是Prepare方法执行后返回的Statement对象结果。
	Tx        X事务          // Tx是Begin方法返回的事务对象。
	X底层结果 interface{} // RawResult 是底层结果，它可能是 sql.Result、*sql.Rows 或 *sql.Row。
}

// Driver 是用于将 SQL 驱动程序集成到 gdb 包的接口。
type X驱动 interface {
	// New 创建并返回指定数据库服务器的数据库对象。
	New(core *Core, node *X配置项) (DB, error)
}

// Link 是一个通用的数据库函数包装器接口。
// 注意，使用 `Link` 进行的任何操作将不会有 SQL 日志记录。
type X底层链接 interface {
	QueryContext(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error)
	ExecContext(ctx context.Context, sql string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, sql string) (*sql.Stmt, error)
	IsOnMaster() bool
	IsTransaction() bool
}

// Sql 是用于记录 SQL 的结构体。
type Sql struct {
	SQL字符串           string        // SQL 字符串（可能包含保留字符 '?'）。
	X类型          string        // SQL操作类型。
	SQL参数          []interface{} // 此SQL的参数。
	SQL格式化后        string        // 格式化后的SQL，其中包含在SQL中的参数。
	X执行错误         error         // Execution result.
	X开始时间戳         int64         // 开始执行的时间戳（毫秒）。
	X结束时间戳           int64         // 结束执行的时间戳（毫秒）。
	X配置组名称         string        // Group 是执行 SQL 时所使用的配置组名称。
	X架构名称        string        // Schema 是执行 SQL 的配置的架构名称。
	X是否为事务 bool          // IsTransaction 标记了这个SQL语句是否在事务中执行。
	X影响行数  int64         // RowsAffected 标记了当前SQL语句执行后获取或影响的行数。
}

// DoInsertOption 是函数 DoInsert 的输入结构体。
type X底层输入 struct {
	OnDuplicateStr string                 // 自定义用于`on duplicated`语句的字符串。
	OnDuplicateMap map[string]interface{} // `OnDuplicateEx`函数为`on duplicated`语句提供的自定义键值对映射
	InsertOption   X插入选项           // 在常数值中执行插入操作。
	BatchCount     int                    // 批量插入的批次数量
}

// TableField 是用于表示表格字段的结构体。
type X字段信息 struct {
	X排序   int         // 用于排序目的，因为映射（map）是无序的。
	X名称    string      // Field name.
	X类型    string      // 字段类型。例如：'int(10) unsigned', 'varchar(64)'。
// 这段注释是对Go语言代码中某个表示字段类型的变量或常量的解释，该字段在数据库表结构设计中使用，比如MySQL等关系型数据库。'int(10) unsigned' 表示一个无符号整数类型，长度为10位；'varchar(64)' 则表示变长字符串类型，最大长度为64个字符。
	X是否为null    bool        // 字段可以为空或非空
	X索引信息     string      // The index information(empty if it's not an index). Eg: PRI, MUL.
	X字段默认值 interface{} // 字段的默认值。
	X额外   string      // 额外信息。例如：自动增长。
	X字段注释 string      // Field comment.
}

// Counter 是用于更新计数的类型。
type X增减 struct {
	X字段名称 string
	X增减值 float64
}

type (
	X原生sql    string                   // Raw 是一个原始SQL语句，它不会被视为参数处理，而是直接作为SQL部分。
// 通常用于嵌入原始sql语句,如:
// g.Model("user").WhereLT("created_at", gdb.Raw("now()")).All()  // SELECT * FROM `user` WHERE `created_at`<now()
// 参考文档:https://goframe.org/pages/viewpage.action?pageId=111911590&showComments=true
	X字段值  = *泛型类.Var              // Value 是字段值类型。
	X行记录 map[string]X字段值         // Record 是表格的行记录。
	X行记录数组 []X行记录                 // Result 是行记录数组。
	Map    = map[string]interface{} // Map 是 map[string]interface{} 的别名，这是最常用的映射类型。
	Map数组   = []Map                  // List 是映射数组的类型。
)

type CatchSQLManager struct {
	SQLArray *数组类.StrArray
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
	ctxKeyForDB               上下文类.StrKey = `CtxKeyForDB`
	ctxKeyCatchSQL            上下文类.StrKey = `CtxKeyCatchSQL`
	ctxKeyInternalProducedSQL 上下文类.StrKey = `CtxKeyInternalProducedSQL`

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

type X插入选项 int

const (
	InsertOptionDefault        X插入选项 = 0
	InsertOptionReplace        X插入选项 = 1
	InsertOptionSave           X插入选项 = 2
	InsertOptionIgnore         X插入选项 = 3
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
	instances = map类.X创建StrAny(true)

	// driverMap 管理所有已注册的自定义驱动。
	driverMap = map[string]X驱动{}

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
	tableFieldsMap = map类.X创建StrAny(true)
)

func init() {
	// allDryRun 从环境变量或命令选项中初始化。
	allDryRun = cmd类.GetOptWithEnv(commandEnvKeyForDryRun, false).X取布尔()
}

// Register 注册自定义数据库驱动到 gdb。
func X注册驱动(名称 string, 驱动 X驱动) error {
	driverMap[名称] = newDriverWrapper(驱动)
	return nil
}

// New 根据给定的配置节点创建并返回一个ORM对象。
func X创建DB对象(配置项 X配置项) (DB对象 DB, 错误 error) {
	return newDBByConfigNode(&配置项, "")
}

// NewByGroup 根据全局配置创建并返回一个 ORM 对象。
// 参数 `name` 指定配置组名称，默认为 DefaultGroupName。
func X创建DB对象并按配置组(配置组名称 ...string) (DB对象 DB, 错误 error) {
	groupName := configs.group
	if len(配置组名称) > 0 && 配置组名称[0] != "" {
		groupName = 配置组名称[0]
	}
	configs.RLock()
	defer configs.RUnlock()

	if len(configs.config) < 1 {
		return nil, 错误类.X创建错误码(
			错误码类.CodeInvalidConfiguration,
			"database configuration is empty, please set the database configuration before using",
		)
	}
	if _, ok := configs.config[groupName]; ok {
		var node *X配置项
		if node, 错误 = getConfigNodeByGroup(groupName, true); 错误 == nil {
			return newDBByConfigNode(node, groupName)
		}
		return nil, 错误
	}
	return nil, 错误类.X创建错误码并格式化(
		错误码类.CodeInvalidConfiguration,
		`database configuration node "%s" is not found, did you misspell group name "%s" or miss the database configuration?`,
		groupName, groupName,
	)
}

// newDBByConfigNode 根据给定的配置节点和组名创建并返回一个ORM对象。
//
// **非常重要**：
// 参数`node`用于数据库创建，而非底层连接创建，
// 因此在同一组内的所有数据库类型配置应当保持一致。
func newDBByConfigNode(node *X配置项, group string) (db DB, err error) {
	if node.X自定义链接信息 != "" {
		node = parseConfigNodeLink(node)
	}
	c := &Core{
		group:  group,
		debug:  安全变量类.NewBool(),
		cache:  缓存类.X创建(),
		links:  map类.X创建StrAny(true),
		logger: 日志类.X创建(),
		config: node,
		dynamicConfig: dynamicConfig{
			MaxIdleConnCount: node.X最大闲置连接数,
			MaxOpenConnCount: node.X最大打开连接数,
			MaxConnLifeTime:  node.X最大空闲时长,
		},
	}
	if v, ok := driverMap[node.X类型]; ok {
		if c.db, err = v.New(c, node); err != nil {
			return nil, err
		}
		return c.db, nil
	}
	errorMsg := `cannot find database driver for specified database type "%s"`
	errorMsg += `, did you misspell type name "%s" or forget importing the database driver? `
	errorMsg += `possible reference: https://github.com/gogf/gf/tree/master/contrib/drivers`
	return nil, 错误类.X创建错误码并格式化(错误码类.CodeInvalidConfiguration, errorMsg, node.X类型, node.X类型)
}

// Instance 返回一个用于数据库操作的实例。
// 参数 `name` 指定了配置组名称，默认为 DefaultGroupName。
func X取单例对象(配置组名称 ...string) (DB对象 DB, 错误 error) {
	group := configs.group
	if len(配置组名称) > 0 && 配置组名称[0] != "" {
		group = 配置组名称[0]
	}
	v := instances.X取值或设置值_函数带锁(group, func() interface{} {
		DB对象, 错误 = X创建DB对象并按配置组(group)
		return DB对象
	})
	if v != nil {
		return v.(DB), nil
	}
	return
}

// 根据给定组获取配置节点，通过内部计算并返回。它使用权重算法进行负载均衡计算。
//
// 参数`master`指定是否获取主节点，如果不是，则在主从配置情况下获取从节点。
func getConfigNodeByGroup(group string, master bool) (*X配置项, error) {
	if list, ok := configs.config[group]; ok {
		// 将主节点和从节点配置数组分离。
		var (
			masterList = make(X配置组, 0)
			slaveList  = make(X配置组, 0)
		)
		for i := 0; i < len(list); i++ {
			if list[i].X节点角色 == dbRoleSlave {
				slaveList = append(slaveList, list[i])
			} else {
				masterList = append(masterList, list[i])
			}
		}
		if len(masterList) < 1 {
			return nil, 错误类.X创建错误码(
				错误码类.CodeInvalidConfiguration,
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
	return nil, 错误类.X创建错误码并格式化(
		错误码类.CodeInvalidConfiguration,
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
func getConfigNodeByWeight(cg X配置组) *X配置项 {
	if len(cg) < 2 {
		return &cg[0]
	}
	var total int
	for i := 0; i < len(cg); i++ {
		total += cg[i].X负载均衡权重 * 100
	}
// 如果total为0，表示所有节点都没有配置权重属性。
// 此时，默认将每个节点的权重属性设为1。
	if total == 0 {
		for i := 0; i < len(cg); i++ {
			cg[i].X负载均衡权重 = 1
			total += cg[i].X负载均衡权重 * 100
		}
	}
	// 排除右侧边界值。
	var (
		min    = 0
		max    = 0
		random = 随机类.X区间整数(0, total-1)
	)
	for i := 0; i < len(cg); i++ {
		max = min + cg[i].X负载均衡权重*100
		if random >= min && random < max {
// ====================================================
// 返回ConfigNode的一个副本。
// ====================================================
			node := X配置项{}
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
		node *X配置项
		ctx  = c.db.X取上下文对象()
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
		n := *c.db.X取当前节点配置()
		node = &n
	}
	if node.X字符集 == "" {
		node.X字符集 = defaultCharset
	}
	// 修改模式
	nodeSchema := 工具类.X取文本值或取默认值(c.schema, schema...)
	if nodeSchema != "" {
		node.X名称 = nodeSchema
	}
	// 更新内部数据中的配置对象。
	internalData := c.底层_GetInternalCtxDataFromCtx(ctx)
	if internalData != nil {
		internalData.ConfigNode = node
	}
	// 通过节点缓存底层连接池对象。
	instanceNameByNode := fmt.Sprintf(`%+v`, node)
	instanceValue := c.links.X取值或设置值_函数带锁(instanceNameByNode, func() interface{} {
		if sqlDb, err = c.db.X底层Open(node); err != nil {
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
	if node.X调试模式 {
		c.db.X设置调试模式(node.X调试模式)
	}
	if node.X空跑特性 {
		c.db.X设置空跑特性(node.X空跑特性)
	}
	return
}
