// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gdb

import (
	"context"
	"database/sql"
	"reflect"

	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/reflection"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/util/gconv"
)

// TXCore是事务管理的结构体。 md5:c2173551528f4399
type TXCore struct {
	db               DB              // db是当前的gdb数据库管理器。 md5:cf7449b67dd32334
	tx               *sql.Tx         // tx 是原始且底层的交易管理器。 md5:ca359da6d7cbfd5b
	ctx              context.Context // ctx 是此次交易专用的上下文。 md5:029c6f2bb9191f37
	master           *sql.DB         // master是原始的和底层的数据库管理器。 md5:cf639ffb6a4872a2
	transactionId    string          // transactionId是此对象为此次交易自动生成的唯一标识符。 md5:1837a379fa0972f8
	transactionCount int             // transactionCount 标记了Begin操作执行的次数。 md5:b733593df5711420
	isClosed         bool            // isClosed 标记该事务已经提交或回滚。 md5:4a5014ffe4a470ba
}

const (
	transactionPointerPrefix    = "transaction"
	contextTransactionKeyPrefix = "TransactionObjectForGroup_"
	transactionIdForLoggerCtx   = "TransactionId"
)

var transactionIdGenerator = gtype.NewUint64()

// Begin 启动并返回事务对象。
// 如果不再使用事务，你应该调用事务对象的Commit或Rollback方法。
// Commit或Rollback方法也会自动关闭事务。
// md5:cca0e58680665343
func (c *Core) Begin(ctx context.Context) (tx TX, err error) {
	return c.doBeginCtx(ctx)
}

func (c *Core) doBeginCtx(ctx context.Context) (TX, error) {
	master, err := c.db.Master()
	if err != nil {
		return nil, err
	}
	var out DoCommitOutput
	out, err = c.db.DoCommit(ctx, DoCommitInput{
		Db:            master,
		Sql:           "BEGIN",
		Type:          SqlTypeBegin,
		IsTransaction: true,
	})
	return out.Tx, err
}

// Transaction 包装了使用函数 `f` 执行的事务逻辑。
// 如果函数 `f` 返回非空错误，它将回滚事务并返回该错误。如果函数 `f` 返回 nil，它将提交事务并返回 nil。
// 
// 注意，在函数 `f` 中不应手动提交或回滚事务，因为这些操作将由这个函数自动处理。
// md5:8906440d4dbbef1f
func (c *Core) Transaction(ctx context.Context, f func(ctx context.Context, tx TX) error) (err error) {
	if ctx == nil {
		ctx = c.db.GetCtx()
	}
	ctx = c.injectInternalCtxData(ctx)
	// 从上下文中检查交易对象。 md5:98b621386407ef30
	var tx TX
	tx = TXFromCtx(ctx, c.db.GetGroup())
	if tx != nil {
		return tx.Transaction(ctx, f)
	}
	tx, err = c.doBeginCtx(ctx)
	if err != nil {
		return err
	}
	// 将事务对象注入上下文。 md5:f5ae21debffd107d
	tx = tx.Ctx(WithTX(tx.GetCtx(), tx))
	defer func() {
		if err == nil {
			if exception := recover(); exception != nil {
				if v, ok := exception.(error); ok && gerror.HasStack(v) {
					err = v
				} else {
					err = gerror.NewCodef(gcode.CodeInternalPanic, "%+v", exception)
				}
			}
		}
		if err != nil {
			if e := tx.Rollback(); e != nil {
				err = e
			}
		} else {
			if e := tx.Commit(); e != nil {
				err = e
			}
		}
	}()
	err = f(tx.GetCtx(), tx)
	return
}

// WithTX 将给定的事务对象注入到上下文中，并返回一个新的上下文。 md5:b4c3c1077b95f681
func WithTX(ctx context.Context, tx TX) context.Context {
	if tx == nil {
		return ctx
	}
	// 检查从给定的输入中是否存在重复注入。 md5:46e37fcbcbc508b5
	group := tx.GetDB().GetGroup()
	if ctxTx := TXFromCtx(ctx, group); ctxTx != nil && ctxTx.GetDB().GetGroup() == group {
		return ctx
	}
	dbCtx := tx.GetDB().GetCtx()
	if ctxTx := TXFromCtx(dbCtx, group); ctxTx != nil && ctxTx.GetDB().GetGroup() == group {
		return dbCtx
	}
	// 向上下文中注入交易对象和ID。 md5:b9cf191f9e07b60d
	ctx = context.WithValue(ctx, transactionKeyForContext(group), tx)
	return ctx
}

// TXFromCtx 从上下文中获取并返回事务对象。
// 它通常用于嵌套事务功能，如果之前未设置，则返回nil。
// md5:21e22b68139fc8b6
func TXFromCtx(ctx context.Context, group string) TX {
	if ctx == nil {
		return nil
	}
	v := ctx.Value(transactionKeyForContext(group))
	if v != nil {
		tx := v.(TX)
		if tx.IsClosed() {
			return nil
		}
		tx = tx.Ctx(ctx)
		return tx
	}
	return nil
}

// transactionKeyForContext 为特定数据库组生成并返回一个字符串，用于将事务对象存储在上下文中。 md5:1dc9cbe3d8e29f02
func transactionKeyForContext(group string) string {
	return contextTransactionKeyPrefix + group
}

// transactionKeyForNestedPoint 在当前保存点处生成并返回事务键。 md5:ca81c7094d96c9fc
func (tx *TXCore) transactionKeyForNestedPoint() string {
	return tx.db.GetCore().QuoteWord(transactionPointerPrefix + gconv.String(tx.transactionCount))
}

// Ctx为当前事务设置上下文。 md5:da0e9ba442dc74f9
func (tx *TXCore) Ctx(ctx context.Context) TX {
	tx.ctx = ctx
	if tx.ctx != nil {
		tx.ctx = tx.db.GetCore().injectInternalCtxData(tx.ctx)
	}
	return tx
}

// GetCtx 返回当前事务的上下文。 md5:e3cb35516cebab84
func (tx *TXCore) GetCtx() context.Context {
	return tx.ctx
}

// GetDB 返回当前事务的DB。 md5:26a64f5fed9954b6
func (tx *TXCore) GetDB() DB {
	return tx.db
}

// GetSqlTX 返回当前事务的底层事务对象。 md5:31b14245dcb30833
func (tx *TXCore) GetSqlTX() *sql.Tx {
	return tx.tx
}

// Commit 提交当前事务。
// 注意，如果处于嵌套事务过程中，它会释放之前的保存事务点，
// 否则，它将提交整个事务。
// md5:9ca50fd58870ed9e
func (tx *TXCore) Commit() error {
	if tx.transactionCount > 0 {
		tx.transactionCount--
		_, err := tx.Exec("RELEASE SAVEPOINT " + tx.transactionKeyForNestedPoint())
		return err
	}
	_, err := tx.db.DoCommit(tx.ctx, DoCommitInput{
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

// Rollback 会回滚当前事务。
// 注意，在嵌套事务过程中，它会回滚当前的事务；否则，它将回滚整个事务。
// md5:0c483721f8447f53
func (tx *TXCore) Rollback() error {
	if tx.transactionCount > 0 {
		tx.transactionCount--
		_, err := tx.Exec("ROLLBACK TO SAVEPOINT " + tx.transactionKeyForNestedPoint())
		return err
	}
	_, err := tx.db.DoCommit(tx.ctx, DoCommitInput{
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

// IsClosed检查并返回此事务是否已经提交或回滚。 md5:cecc2f01ef3e3556
func (tx *TXCore) IsClosed() bool {
	return tx.isClosed
}

// Begin 启动一个嵌套事务过程。 md5:1b04e48800ebefdd
func (tx *TXCore) Begin() error {
	_, err := tx.Exec("SAVEPOINT " + tx.transactionKeyForNestedPoint())
	if err != nil {
		return err
	}
	tx.transactionCount++
	return nil
}

// SavePoint 执行 `SAVEPOINT xxx` SQL 语句，该语句在当前点保存事务。
// 参数 `point` 指定将被保存到服务器的保存点名称。
// md5:f4061450298afabd
func (tx *TXCore) SavePoint(point string) error {
	_, err := tx.Exec("SAVEPOINT " + tx.db.GetCore().QuoteWord(point))
	return err
}

// RollbackTo 执行 `ROLLBACK TO SAVEPOINT xxx` SQL语句，回滚到指定的保存点事务。
// 参数 `point` 指定了之前保存的保存点名称。
// md5:e347c163ad8fefa7
func (tx *TXCore) RollbackTo(point string) error {
	_, err := tx.Exec("ROLLBACK TO SAVEPOINT " + tx.db.GetCore().QuoteWord(point))
	return err
}

// Transaction 包装了使用函数 `f` 执行的事务逻辑。
// 如果函数 `f` 返回非空错误，它将回滚事务并返回该错误。如果函数 `f` 返回 nil，它将提交事务并返回 nil。
// 
// 注意，在函数 `f` 中不应手动提交或回滚事务，因为这些操作将由这个函数自动处理。
// md5:8906440d4dbbef1f
func (tx *TXCore) Transaction(ctx context.Context, f func(ctx context.Context, tx TX) error) (err error) {
	if ctx != nil {
		tx.ctx = ctx
	}
	// 从上下文中检查交易对象。 md5:98b621386407ef30
	if TXFromCtx(tx.ctx, tx.db.GetGroup()) == nil {
		// 将事务对象注入上下文。 md5:f5ae21debffd107d
		tx.ctx = WithTX(tx.ctx, tx)
	}
	err = tx.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err == nil {
			if exception := recover(); exception != nil {
				if v, ok := exception.(error); ok && gerror.HasStack(v) {
					err = v
				} else {
					err = gerror.NewCodef(gcode.CodeInternalPanic, "%+v", exception)
				}
			}
		}
		if err != nil {
			if e := tx.Rollback(); e != nil {
				err = e
			}
		} else {
			if e := tx.Commit(); e != nil {
				err = e
			}
		}
	}()
	err = f(tx.ctx, tx)
	return
}

// Query 在事务上执行查询操作。
// 请参阅Core.Query。
// md5:0d7503cceb0dc1d6
func (tx *TXCore) Query(sql string, args ...interface{}) (result Result, err error) {
	return tx.db.DoQuery(tx.ctx, &txLink{tx.tx}, sql, args...)
}

// Exec 在事务上执行非查询操作。
// 参见Core.Exec。
// md5:043edf012223f310
func (tx *TXCore) Exec(sql string, args ...interface{}) (sql.Result, error) {
	return tx.db.DoExec(tx.ctx, &txLink{tx.tx}, sql, args...)
}

// Prepare 创建一个预处理语句，以便后续的查询或执行。
// 可以从返回的语句中并发地运行多个查询或执行。
// 调用者必须在不再需要该语句时调用语句的 Close 方法。
// md5:16334dc7db1c37a9
func (tx *TXCore) Prepare(sql string) (*Stmt, error) {
	return tx.db.DoPrepare(tx.ctx, &txLink{tx.tx}, sql)
}

// GetAll 从数据库中查询并返回数据记录。 md5:dfdcfddaa70ab1d4
func (tx *TXCore) GetAll(sql string, args ...interface{}) (Result, error) {
	return tx.Query(sql, args...)
}

// GetOne 从数据库中查询并返回一条记录。 md5:9552f7e095f58141
func (tx *TXCore) GetOne(sql string, args ...interface{}) (Record, error) {
	list, err := tx.GetAll(sql, args...)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	return nil, nil
}

// GetStruct 从数据库中查询一条记录，并将其转换为给定的结构体。
// 参数 `pointer` 应该是指向结构体的指针。
// md5:7ddc0d419d6cd2aa
func (tx *TXCore) GetStruct(obj interface{}, sql string, args ...interface{}) error {
	one, err := tx.GetOne(sql, args...)
	if err != nil {
		return err
	}
	return one.Struct(obj)
}

// GetStructs 从数据库查询记录，并将它们转换为给定的结构体。参数 `pointer` 应该是结构体切片的类型：[]struct 或 []*struct。
// md5:af7dfbf46c6660c6
func (tx *TXCore) GetStructs(objPointerSlice interface{}, sql string, args ...interface{}) error {
	all, err := tx.GetAll(sql, args...)
	if err != nil {
		return err
	}
	return all.Structs(objPointerSlice)
}

// GetScan 从数据库查询一个或多个记录，并将它们转换为给定的结构体或结构体数组。
//
// 如果参数 `pointer` 是结构体指针类型，它内部会调用 GetStruct 进行转换。如果参数 `pointer` 是切片类型，它内部会调用 GetStructs 进行转换。
// md5:c1dbdab7a7c29c51
func (tx *TXCore) GetScan(pointer interface{}, sql string, args ...interface{}) error {
	reflectInfo := reflection.OriginTypeAndKind(pointer)
	if reflectInfo.InputKind != reflect.Ptr {
		return gerror.NewCodef(
			gcode.CodeInvalidParameter,
			"params should be type of pointer, but got: %v",
			reflectInfo.InputKind,
		)
	}
	switch reflectInfo.OriginKind {
	case reflect.Array, reflect.Slice:
		return tx.GetStructs(pointer, sql, args...)

	case reflect.Struct:
		return tx.GetStruct(pointer, sql, args...)
	}
	return gerror.NewCodef(
		gcode.CodeInvalidParameter,
		`in valid parameter type "%v", of which element type should be type of struct/slice`,
		reflectInfo.InputType,
	)
}

// GetValue 从数据库查询并返回字段值。
// SQL 应该只查询数据库中的一个字段，否则它将只返回结果中的一个字段。
// md5:96794360fadbc288
func (tx *TXCore) GetValue(sql string, args ...interface{}) (Value, error) {
	one, err := tx.GetOne(sql, args...)
	if err != nil {
		return nil, err
	}
	for _, v := range one {
		return v, nil
	}
	return nil, nil
}

// GetCount 从数据库中查询并返回计数。 md5:a8368d39f4a58979
func (tx *TXCore) GetCount(sql string, args ...interface{}) (int64, error) {
	if !gregex.IsMatchString(`(?i)SELECT\s+COUNT\(.+\)\s+FROM`, sql) {
		sql, _ = gregex.ReplaceString(`(?i)(SELECT)\s+(.+)\s+(FROM)`, `$1 COUNT($2) $3`, sql)
	}
	value, err := tx.GetValue(sql, args...)
	if err != nil {
		return 0, err
	}
	return value.Int64(), nil
}

// Insert 执行 "INSERT INTO..." 语句来操作表。
// 如果表中已经存在数据的唯一记录，它会返回错误。
//
// 参数 `data` 可以是 map、gmap、struct、*struct、[]map 或 []struct 等类型。
// 例如：
// Data(g.Map{"uid": 10000, "name": "john"})
// Data(g.Slice{g.Map{"uid": 10000, "name": "john"}, g.Map{"uid": 20000, "name": "smith"}})
//
// 参数 `batch` 在给定数据为切片时，指定批量操作的次数。
// md5:fd75d343f485b8dc
func (tx *TXCore) Insert(table string, data interface{}, batch ...int) (sql.Result, error) {
	if len(batch) > 0 {
		return tx.Model(table).Ctx(tx.ctx).Data(data).Batch(batch[0]).Insert()
	}
	return tx.Model(table).Ctx(tx.ctx).Data(data).Insert()
}

// InsertIgnore 执行 "INSERT IGNORE INTO ..." 语句针对该表。
// 如果表中已存在该数据的唯一记录，则忽略插入操作。
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 例如：
// Data(g.Map{"uid": 10000, "name": "john"})
// Data(g.Slice{g.Map{"uid": 10000, "name": "john"}, g.Map{"uid": 20000, "name": "smith"})
//
// 当给定的数据为切片时，参数 `batch` 指定批处理操作的计数。
// md5:49f76901041c9819
func (tx *TXCore) InsertIgnore(table string, data interface{}, batch ...int) (sql.Result, error) {
	if len(batch) > 0 {
		return tx.Model(table).Ctx(tx.ctx).Data(data).Batch(batch[0]).InsertIgnore()
	}
	return tx.Model(table).Ctx(tx.ctx).Data(data).InsertIgnore()
}

// InsertAndGetId 执行插入操作，并返回自动生成的最后一个插入id。 md5:8d00b40a35fa48a5
func (tx *TXCore) InsertAndGetId(table string, data interface{}, batch ...int) (int64, error) {
	if len(batch) > 0 {
		return tx.Model(table).Ctx(tx.ctx).Data(data).Batch(batch[0]).InsertAndGetId()
	}
	return tx.Model(table).Ctx(tx.ctx).Data(data).InsertAndGetId()
}

// Replace 用于执行针对该表的 "REPLACE INTO..." 语句。如果表中已经存在数据的唯一记录，它会删除该记录并插入新的。
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。例如：
// Data(g.Map{"uid": 10000, "name": "john"})
// Data(g.Slice{g.Map{"uid": 10000, "name": "john"}, g.Map{"uid": 20000, "name": "smith"}})
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。如果给定的数据是切片类型，它将进行批量替换，并可选地通过参数 `batch` 指定批量操作次数。
// md5:69ecd0994eab5bbb
func (tx *TXCore) Replace(table string, data interface{}, batch ...int) (sql.Result, error) {
	if len(batch) > 0 {
		return tx.Model(table).Ctx(tx.ctx).Data(data).Batch(batch[0]).Replace()
	}
	return tx.Model(table).Ctx(tx.ctx).Data(data).Replace()
}

// Save 执行 "INSERT INTO ... ON DUPLICATE KEY UPDATE..." 语句来操作表。
// 如果保存的数据中存在主键或唯一索引，它将更新记录；否则，将在表中插入新记录。
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 例如：
// Data(g.Map{"uid": 10000, "name": "john"})
// Data(g.Slice{g.Map{"uid": 10000, "name": "john"}, g.Map{"uid": 20000, "name": "smith"}})
//
// 如果给定的数据是切片类型，它将进行批量保存。可选参数 `batch` 指定了批量操作的次数。
// md5:c76721f5e0b01424
func (tx *TXCore) Save(table string, data interface{}, batch ...int) (sql.Result, error) {
	if len(batch) > 0 {
		return tx.Model(table).Ctx(tx.ctx).Data(data).Batch(batch[0]).Save()
	}
	return tx.Model(table).Ctx(tx.ctx).Data(data).Save()
}

// Update 执行表的 "UPDATE ... " 语句。
//
// 参数 `data` 可以是字符串、映射、gmap、结构体或指向结构体的指针等类型。
// 例如："uid=10000", "uid", 10000, g.Map{"uid": 10000, "name":"john"}
//
// 参数 `condition` 也可以是字符串、映射、gmap、切片或结构体及指向结构体的指针等类型。
// 常与参数 `args` 配合使用。
// 例如：
// "uid=10000",
// "uid", 10000
// "money>? AND name like ?", 99999, "vip_%"
// "status IN (?)", g.Slice{1,2,3}
// "age IN(?,?)", 18, 50
// User{ Id : 1, UserName : "john"}.
// md5:8651eb1bd7e10da0
func (tx *TXCore) Update(table string, data interface{}, condition interface{}, args ...interface{}) (sql.Result, error) {
	return tx.Model(table).Ctx(tx.ctx).Data(data).Where(condition, args...).Update()
}

// Delete 执行 "DELETE FROM ... " 语句针对该表。
//
// `condition` 参数可以是字符串、映射、gmap、切片、结构体或指向结构体的指针等多种类型。
// 它常与参数 `args` 一起使用。
// 例如：
// "uid=10000",
// "uid", 10000
// "money>? AND name like ?", 99999, "vip_%"
// "status IN (?)", g.Slice{1,2,3}
// "age IN(?,?)", 18, 50
// User{ Id : 1, UserName : "john"}.
// md5:c6c87830434eba7d
func (tx *TXCore) Delete(table string, condition interface{}, args ...interface{}) (sql.Result, error) {
	return tx.Model(table).Ctx(tx.ctx).Where(condition, args...).Delete()
}

// QueryContext实现了Link.QueryContext接口函数。 md5:f42e7710688a27fc
func (tx *TXCore) QueryContext(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error) {
	return tx.tx.QueryContext(ctx, sql, args...)
}

// ExecContext 实现了 Link.ExecContext 接口函数。 md5:9bd9a386cc5fc878
func (tx *TXCore) ExecContext(ctx context.Context, sql string, args ...interface{}) (sql.Result, error) {
	return tx.tx.ExecContext(ctx, sql, args...)
}

// PrepareContext 实现了接口 Link 的 PrepareContext 函数。 md5:b08e1c50bfb8f8e8
func (tx *TXCore) PrepareContext(ctx context.Context, sql string) (*sql.Stmt, error) {
	return tx.tx.PrepareContext(ctx, sql)
}

// IsOnMaster 实现接口函数 Link.IsOnMaster。 md5:4fddd5d2ad612d30
func (tx *TXCore) IsOnMaster() bool {
	return true
}

// IsTransaction 实现了 Link 接口中的函数 IsTransaction。 md5:692b7be612be52bd
func (tx *TXCore) IsTransaction() bool {
	return true
}
