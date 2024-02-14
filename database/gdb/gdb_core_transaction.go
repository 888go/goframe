// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

import (
	"context"
	"database/sql"
	"reflect"
	
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/reflection"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/util/gconv"
)

// TXCore是用于事务管理的结构体。
type X基础事务 struct {
	db               DB              // db 是当前的 gdb 数据库管理器。
	tx               *sql.Tx         // tx 是原始且底层的事务管理器。
	ctx              context.Context // ctx 是本次事务的上下文。
	master           *sql.DB         // master 是原始的、底层的数据库管理器。
	transactionId    string          // transactionId 是由本对象为此次交易生成的唯一标识符。
	transactionCount int             // transactionCount 标记调用 Begins 的次数。
	isClosed         bool            // isClosed 标记此事务已经完成提交或回滚。
}

const (
	transactionPointerPrefix    = "transaction"
	contextTransactionKeyPrefix = "TransactionObjectForGroup_"
	transactionIdForLoggerCtx   = "TransactionId"
)

var transactionIdGenerator = 安全变量类.NewUint64()

// Begin 开始并返回事务对象。
// 如果不再使用该事务，你应该调用事务对象的 Commit 或 Rollback 函数。
// Commit 或 Rollback 函数也会自动关闭事务。
func (c *Core) X事务开启(上下文 context.Context) (事务对象 TX, 错误 error) {
	return c.doBeginCtx(上下文)
}

func (c *Core) doBeginCtx(ctx context.Context) (TX, error) {
	master, err := c.db.X取主节点对象()
	if err != nil {
		return nil, err
	}
	var out X输出
	out, err = c.db.X底层DoCommit(ctx, DoCommitInput{
		Db:            master,
		Sql:           "BEGIN",
		Type:          SqlTypeBegin,
		IsTransaction: true,
	})
	return out.Tx, err
}

// Transaction 通过函数 `f` 包装事务逻辑。如果函数 `f` 返回非空错误，它将回滚事务并返回该错误。若函数 `f` 返回空（nil）错误，它将提交事务并返回空。
// 注意：在函数 `f` 中不应手动调用 Commit 或 Rollback 方法处理事务，因为这些操作在此函数中已自动完成。
func (c *Core) X事务(上下文 context.Context, 回调函数 func(上下文 context.Context, 事务对象 TX) error) (错误 error) {
	if 上下文 == nil {
		上下文 = c.db.X取上下文对象()
	}
	上下文 = c.底层_InjectInternalCtxData(上下文)
	// 从上下文中检查交易对象。
	var tx TX
	tx = X事务从上下文取对象(上下文, c.db.X取配置组名称())
	if tx != nil {
		return tx.X事务(上下文, 回调函数)
	}
	tx, 错误 = c.doBeginCtx(上下文)
	if 错误 != nil {
		return 错误
	}
	// 将事务对象注入到上下文中。
	tx = tx.X设置上下文并取副本(底层WithTX(tx.X取上下文对象(), tx))
	defer func() {
		if 错误 == nil {
			if exception := recover(); exception != nil {
				if v, ok := exception.(error); ok && 错误类.X判断是否带堆栈(v) {
					错误 = v
				} else {
					错误 = 错误类.X创建错误码并格式化(错误码类.CodeInternalPanic, "%+v", exception)
				}
			}
		}
		if 错误 != nil {
			if e := tx.X事务回滚(); e != nil {
				错误 = e
			}
		} else {
			if e := tx.X事务提交(); e != nil {
				错误 = e
			}
		}
	}()
	错误 = 回调函数(tx.X取上下文对象(), tx)
	return
}

// WithTX 将给定的事务对象注入到上下文中并返回一个新的上下文。
func 底层WithTX(上下文 context.Context, 事务对象 TX) context.Context {
	if 事务对象 == nil {
		return 上下文
	}
	// 检查从给定参数中是否存在重复注入
	group := 事务对象.X取DB对象().X取配置组名称()
	if ctxTx := X事务从上下文取对象(上下文, group); ctxTx != nil && ctxTx.X取DB对象().X取配置组名称() == group {
		return 上下文
	}
	dbCtx := 事务对象.X取DB对象().X取上下文对象()
	if ctxTx := X事务从上下文取对象(dbCtx, group); ctxTx != nil && ctxTx.X取DB对象().X取配置组名称() == group {
		return dbCtx
	}
	// 将事务对象和ID注入到上下文中。
	上下文 = context.WithValue(上下文, transactionKeyForContext(group), 事务对象)
	return 上下文
}

// TXFromCtx 从上下文中检索并返回事务对象。
// 通常用于嵌套事务功能，如果没有先前设置，则返回nil。
func X事务从上下文取对象(上下文 context.Context, group string) TX {
	if 上下文 == nil {
		return nil
	}
	v := 上下文.Value(transactionKeyForContext(group))
	if v != nil {
		tx := v.(TX)
		if tx.X是否已关闭() {
			return nil
		}
		tx = tx.X设置上下文并取副本(上下文)
		return tx
	}
	return nil
}

// transactionKeyForContext 为特定数据库组的事务对象生成并返回一个字符串，以便将其存储到上下文中。
func transactionKeyForContext(group string) string {
	return contextTransactionKeyPrefix + group
}

// transactionKeyForNestedPoint 根据当前保存点构建并返回事务键。
func (tx *X基础事务) transactionKeyForNestedPoint() string {
	return tx.db.X取Core对象().X底层QuoteWord(transactionPointerPrefix + 转换类.String(tx.transactionCount))
}

// Ctx 设置当前事务的上下文。
func (tx *X基础事务) X设置上下文并取副本(上下文 context.Context) TX {
	tx.ctx = 上下文
	if tx.ctx != nil {
		tx.ctx = tx.db.X取Core对象().底层_InjectInternalCtxData(tx.ctx)
	}
	return tx
}

// GetCtx 返回当前事务的上下文。
func (tx *X基础事务) X取上下文对象() context.Context {
	return tx.ctx
}

// GetDB 返回当前事务的 DB。
func (tx *X基础事务) X取DB对象() DB {
	return tx.db
}

// GetSqlTX 返回当前事务的底层事务对象。
func (tx *X基础事务) X底层取事务对象() *sql.Tx {
	return tx.tx
}

// Commit 提交当前事务。
// 注意，如果它在一个嵌套事务过程中，会释放先前保存的事务点，
// 否则，它将提交整个事务。
func (tx *X基础事务) X事务提交() error {
	if tx.transactionCount > 0 {
		tx.transactionCount--
		_, err := tx.X原生SQL执行("RELEASE SAVEPOINT " + tx.transactionKeyForNestedPoint())
		return err
	}
	_, err := tx.db.X底层DoCommit(tx.ctx, DoCommitInput{
		Tx:            tx.tx,
		Sql:           "COMMIT",
		Type:          SqlTypeTXCommit,
		IsTransaction: true,
	})
	if err == nil {
		tx.isClosed = true
	}
	return err
}

// Rollback 回滚当前事务。
// 注意，如果当前处在嵌套事务过程中，它会回滚当前事务，
// 否则，它会回滚整个事务。
func (tx *X基础事务) X事务回滚() error {
	if tx.transactionCount > 0 {
		tx.transactionCount--
		_, err := tx.X原生SQL执行("ROLLBACK TO SAVEPOINT " + tx.transactionKeyForNestedPoint())
		return err
	}
	_, err := tx.db.X底层DoCommit(tx.ctx, DoCommitInput{
		Tx:            tx.tx,
		Sql:           "ROLLBACK",
		Type:          SqlTypeTXRollback,
		IsTransaction: true,
	})
	if err == nil {
		tx.isClosed = true
	}
	return err
}

// IsClosed 检查并返回此事务是否已提交或回滚。
func (tx *X基础事务) X是否已关闭() bool {
	return tx.isClosed
}

// Begin 开始一个嵌套事务过程。
func (tx *X基础事务) X事务开启() error {
	_, err := tx.X原生SQL执行("SAVEPOINT " + tx.transactionKeyForNestedPoint())
	if err != nil {
		return err
	}
	tx.transactionCount++
	return nil
}

// SavePoint 执行 `SAVEPOINT xxx` SQL 语句，用于保存当前事务点。
// 参数 `point` 指定了将被保存到服务器的事务点名称。
func (tx *X基础事务) X保存事务点(事务点名称 string) error {
	_, err := tx.X原生SQL执行("SAVEPOINT " + tx.db.X取Core对象().X底层QuoteWord(事务点名称))
	return err
}

// RollbackTo 执行 `ROLLBACK TO SAVEPOINT xxx` SQL 语句，用于回滚到指定的保存点。
// 参数 `point` 指定了先前保存的事务点名称。
func (tx *X基础事务) X回滚事务点(事务点名称 string) error {
	_, err := tx.X原生SQL执行("ROLLBACK TO SAVEPOINT " + tx.db.X取Core对象().X底层QuoteWord(事务点名称))
	return err
}

// Transaction 通过函数 `f` 包装事务逻辑。如果函数 `f` 返回非空错误，它将回滚事务并返回该错误。若函数 `f` 返回空（nil）错误，它将提交事务并返回空。
// 注意：在函数 `f` 中不应手动调用 Commit 或 Rollback 方法处理事务，因为这些操作在此函数中已自动完成。
func (tx *X基础事务) X事务(上下文 context.Context, 回调函数 func(上下文 context.Context, 事务对象 TX) error) (错误 error) {
	if 上下文 != nil {
		tx.ctx = 上下文
	}
	// 从上下文中检查交易对象。
	if X事务从上下文取对象(tx.ctx, tx.db.X取配置组名称()) == nil {
		// 将事务对象注入到上下文中。
		tx.ctx = 底层WithTX(tx.ctx, tx)
	}
	错误 = tx.X事务开启()
	if 错误 != nil {
		return 错误
	}
	defer func() {
		if 错误 == nil {
			if exception := recover(); exception != nil {
				if v, ok := exception.(error); ok && 错误类.X判断是否带堆栈(v) {
					错误 = v
				} else {
					错误 = 错误类.X创建错误码并格式化(错误码类.CodeInternalPanic, "%+v", exception)
				}
			}
		}
		if 错误 != nil {
			if e := tx.X事务回滚(); e != nil {
				错误 = e
			}
		} else {
			if e := tx.X事务提交(); e != nil {
				错误 = e
			}
		}
	}()
	错误 = 回调函数(tx.ctx, tx)
	return
}

// Query 在事务上执行查询操作。
// 参见 Core.Query.
func (tx *X基础事务) X原生SQL查询(sql string, 参数 ...interface{}) (结果 Result, 错误 error) {
	return tx.db.X底层原生SQL查询(tx.ctx, &txLink{tx.tx}, sql, 参数...)
}

// Exec 在事务上执行非查询操作。
// 请参阅 Core.Exec.
func (tx *X基础事务) X原生SQL执行(sql string, 参数 ...interface{}) (sql.Result, error) {
	return tx.db.X底层原生SQL执行(tx.ctx, &txLink{tx.tx}, sql, 参数...)
}

// Prepare 函数用于创建一个预编译的语句，以便后续进行查询或执行操作。
// 从返回的语句中可以并发地执行多个查询或执行操作。
// 当该语句不再被需要时，调用者必须调用该语句的 Close 方法。
func (tx *X基础事务) X原生sql取参数预处理对象(sql string) (*Stmt, error) {
	return tx.db.X底层原生sql参数预处理对象(tx.ctx, &txLink{tx.tx}, sql)
}

// GetAll 从数据库查询并返回数据记录。
// 2024-01-09 改成别名,功能与tx.Query()重复
func (tx *X基础事务) GetAll别名(sql string, 参数 ...interface{}) (Result, error) {
	return tx.X原生SQL查询(sql, 参数...)
}

// GetOne 从数据库查询并返回一条记录。
func (tx *X基础事务) X原生SQL查询单条记录(sql string, 参数 ...interface{}) (Record, error) {
	list, err := tx.GetAll别名(sql, 参数...)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	return nil, nil
}

// GetStruct 从数据库查询一条记录并将其转换为给定的结构体。
// 参数 `pointer` 应该是指向结构体的指针。
func (tx *X基础事务) X原生SQL查询单条到结构体指针(结构体指针 interface{}, sql string, 参数 ...interface{}) error {
	one, err := tx.X原生SQL查询单条记录(sql, 参数...)
	if err != nil {
		return err
	}
	return one.X取结构体指针(结构体指针)
}

// GetStructs 从数据库查询记录并将其转换为给定的结构体。
// 参数`pointer`应为结构体切片的类型：[]struct/[]*struct。
func (tx *X基础事务) X原生SQL查询到结构体数组指针(结构体指针 interface{}, sql string, 参数 ...interface{}) error {
	all, err := tx.GetAll别名(sql, 参数...)
	if err != nil {
		return err
	}
	return all.X取数组结构体指针(结构体指针)
}

// GetScan 从数据库查询一个或多个记录，并将它们转换为给定的结构体或结构体数组。
//
// 如果参数`pointer`是结构体指针类型，它会内部调用 GetStruct 进行转换。如果参数 `pointer` 是切片类型，则会内部调用 GetStructs 进行转换。
func (tx *X基础事务) X原生SQL查询到结构体指针(结构体指针 interface{}, sql string, 参数 ...interface{}) error {
	reflectInfo := reflection.OriginTypeAndKind(结构体指针)
	if reflectInfo.InputKind != reflect.Ptr {
		return 错误类.X创建错误码并格式化(
			错误码类.CodeInvalidParameter,
			"params should be type of pointer, but got: %v",
			reflectInfo.InputKind,
		)
	}
	switch reflectInfo.OriginKind {
	case reflect.Array, reflect.Slice:
		return tx.X原生SQL查询到结构体数组指针(结构体指针, sql, 参数...)

	case reflect.Struct:
		return tx.X原生SQL查询单条到结构体指针(结构体指针, sql, 参数...)
	}
	return 错误类.X创建错误码并格式化(
		错误码类.CodeInvalidParameter,
		`in valid parameter type "%v", of which element type should be type of struct/slice`,
		reflectInfo.InputType,
	)
}

// GetValue 从数据库查询并返回字段值。
// SQL语句应当只查询数据库中的一个字段，否则它将仅返回结果中的一个字段。
func (tx *X基础事务) X原生SQL查询字段值(sql string, 参数 ...interface{}) (Value, error) {
	one, err := tx.X原生SQL查询单条记录(sql, 参数...)
	if err != nil {
		return nil, err
	}
	for _, v := range one {
		return v, nil
	}
	return nil, nil
}

// GetCount 从数据库查询并返回计数。
func (tx *X基础事务) X原生SQL查询字段计数(sql string, 参数 ...interface{}) (int64, error) {
	if !正则类.X是否匹配文本(`(?i)SELECT\s+COUNT\(.+\)\s+FROM`, sql) {
		sql, _ = 正则类.X替换文本(`(?i)(SELECT)\s+(.+)\s+(FROM)`, `$1 COUNT($2) $3`, sql)
	}
	value, err := tx.X原生SQL查询字段值(sql, 参数...)
	if err != nil {
		return 0, err
	}
	return value.X取整数64位(), nil
}

// Insert 执行针对该表的 "INSERT INTO ..." 语句。
// 如果表中已经存在一条相同数据的唯一记录，则返回错误。
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 示例：
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// 当给定的数据为切片时，参数 `batch` 指定了批量操作的数量。
func (tx *X基础事务) X插入(表名称 string, 值 interface{}, 批量操作行数 ...int) (sql.Result, error) {
	if len(批量操作行数) > 0 {
		return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X设置数据(值).X设置批量操作行数(批量操作行数[0]).X插入()
	}
	return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X设置数据(值).X插入()
}

// InsertIgnore 执行针对表的 "INSERT IGNORE INTO ..." 语句。
// 如果表中已存在一条相同的数据记录，它将忽略插入操作。
//
// 参数 `data` 可以为 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 例如：
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// 当给定数据为切片时，参数 `batch` 指定批量操作的数量。
func (tx *X基础事务) X插入并跳过已存在(表名称 string, 值 interface{}, 批量操作行数 ...int) (sql.Result, error) {
	if len(批量操作行数) > 0 {
		return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X设置数据(值).X设置批量操作行数(批量操作行数[0]).X插入并跳过已存在()
	}
	return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X设置数据(值).X插入并跳过已存在()
}

// InsertAndGetId 执行插入操作，并返回自动生成的最后一个插入ID。
func (tx *X基础事务) X插入并取ID(表名称 string, 值 interface{}, 批量操作行数 ...int) (int64, error) {
	if len(批量操作行数) > 0 {
		return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X设置数据(值).X设置批量操作行数(批量操作行数[0]).X插入并取ID()
	}
	return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X设置数据(值).X插入并取ID()
}

// Replace 执行针对该表的 "REPLACE INTO ..." 语句。
// 如果表中已存在一条唯一数据记录，它会先删除这条记录，然后插入一条新的记录。
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 示例：
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 若给定的数据是切片类型，它将执行批量替换操作，可选参数
// `batch` 指定了批量操作的数量。
func (tx *X基础事务) X插入并替换已存在(表名称 string, 值 interface{}, 批量操作行数 ...int) (sql.Result, error) {
	if len(批量操作行数) > 0 {
		return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X设置数据(值).X设置批量操作行数(批量操作行数[0]).X插入并替换已存在()
	}
	return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X设置数据(值).X插入并替换已存在()
}

// Save 执行针对表的 "INSERT INTO ... ON DUPLICATE KEY UPDATE..." 语句。
// 如果保存数据中存在主键或唯一索引，它将更新记录，否则将在表中插入新的记录。
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 例如：
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// 如果给定的数据是切片类型，那么它将执行批量保存操作，可选参数
// `batch` 指定了批量操作的数量。
func (tx *X基础事务) X插入并更新已存在(表名称 string, 值 interface{}, 批量操作行数 ...int) (sql.Result, error) {
	if len(批量操作行数) > 0 {
		return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X设置数据(值).X设置批量操作行数(批量操作行数[0]).X插入并更新已存在()
	}
	return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X设置数据(值).X插入并更新已存在()
}

// Update 执行针对该表的 "UPDATE ... " 语句。
//
// 参数 `data` 可以为 string/map/gmap/struct/*struct 等类型。
// 例如："uid=10000"、"uid", 10000、g.Map{"uid": 10000, "name":"john"}
//
// 参数 `condition` 可以为 string/map/gmap/slice/struct/*struct 等类型，通常与参数 `args` 一起使用。
// 例如：
// "uid=10000"
// "uid", 10000
// "money>? AND name like ?", 99999, "vip_%"
// "status IN (?)", g.Slice{1,2,3}
// "age IN(?,?)", 18, 50
// User{ Id : 1, UserName : "john"}
// 注：这里的 `g.Map` 和 `g.Slice` 是一种特定的 Go 语言数据结构（可能是自定义类型），分别代表映射和切片。
func (tx *X基础事务) X更新(表名称 string, 值 interface{}, 条件 interface{}, 参数 ...interface{}) (sql.Result, error) {
	return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X设置数据(值).X条件(条件, 参数...).X更新()
}

// Delete 执行针对该表的 "DELETE FROM ... " 语句。
//
// 参数 `condition` 可以是字符串、映射（map/gmap）、切片、结构体或指针类型等。
// 常常与参数 `args` 一起使用。例如：
// "uid=10000"，
// "uid", 10000
// "money>? AND name like ?", 99999, "vip_%"
// "status IN (?)", g.Slice{1,2,3}
// "age IN(?,?)", 18, 50
// User{ Id : 1, UserName : "john"}
// 中文注释：
// Delete 方法用于执行对该数据表执行 "DELETE FROM ... " SQL语句。
//
// 参数 `condition` 支持多种类型，如字符串、字典(map/gmap)、数组、结构体、结构体指针等。
// 通常会结合参数 `args` 使用，例如以下示例：
// "uid=10000"，（条件为 uid 等于 10000）
// "uid", 10000
// "money>? AND name like ?", 99999, "vip_%" （条件为 money 大于 99999 并且 name 字段匹配 "vip_%"）
// "status IN (?)", g.Slice{1,2,3} （条件为 status 字段在数组 [1,2,3] 中）
// "age IN(?,?)", 18, 50 （条件为 age 字段在范围 18 到 50 内）
// User{ Id : 1, UserName : "john"} （根据结构体定义的字段作为条件）
func (tx *X基础事务) X删除(表名称 string, 条件 interface{}, 参数 ...interface{}) (sql.Result, error) {
	return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X条件(条件, 参数...).X删除()
}

// QueryContext 实现了接口函数 Link.QueryContext。
func (tx *X基础事务) QueryContext(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error) {
	return tx.tx.QueryContext(ctx, sql, args...)
}

// ExecContext 实现了接口函数 Link.ExecContext。
func (tx *X基础事务) ExecContext(ctx context.Context, sql string, args ...interface{}) (sql.Result, error) {
	return tx.tx.ExecContext(ctx, sql, args...)
}

// PrepareContext 实现了接口函数 Link.PrepareContext。
func (tx *X基础事务) PrepareContext(ctx context.Context, sql string) (*sql.Stmt, error) {
	return tx.tx.PrepareContext(ctx, sql)
}

// IsOnMaster 实现了接口函数 Link.IsOnMaster。
func (tx *X基础事务) IsOnMaster() bool {
	return true
}

// IsTransaction 实现了接口函数 Link.IsTransaction。
func (tx *X基础事务) IsTransaction() bool {
	return true
}
