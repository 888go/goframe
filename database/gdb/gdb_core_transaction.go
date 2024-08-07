// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

import (
	"context"
	"database/sql"
	"reflect"

	gtype "github.com/888go/goframe/container/gtype"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/reflection"
	gregex "github.com/888go/goframe/text/gregex"
	gconv "github.com/888go/goframe/util/gconv"
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

// X事务开启 启动并返回事务对象。
// 如果不再使用事务，你应该调用事务对象的Commit或Rollback方法。
// Commit或Rollback方法也会自动关闭事务。
// md5:cca0e58680665343
func (c *Core) X事务开启(上下文 context.Context) (事务对象 TX, 错误 error) {
	return c.doBeginCtx(上下文)
}

func (c *Core) doBeginCtx(ctx context.Context) (TX, error) {
	master, err := c.db.X取主节点对象()
	if err != nil {
		return nil, err
	}
	var out DoCommitOutput
	out, err = c.db.X底层DoCommit(ctx, DoCommitInput{
		Db:            master,
		Sql:           "BEGIN",
		Type:          SqlTypeBegin,
		IsTransaction: true,
	})
	return out.Tx, err
}

// X事务 包装了使用函数 `f` 执行的事务逻辑。
// 如果函数 `f` 返回非空错误，它将回滚事务并返回该错误。如果函数 `f` 返回 nil，它将提交事务并返回 nil。
// 
// 注意，在函数 `f` 中不应手动提交或回滚事务，因为这些操作将由这个函数自动处理。
// md5:8906440d4dbbef1f
func (c *Core) X事务(上下文 context.Context, 回调函数 func(上下文 context.Context, tx TX) error) (错误 error) {
	if 上下文 == nil {
		上下文 = c.db.X取上下文对象()
	}
	上下文 = c.injectInternalCtxData(上下文)
		// 从上下文中检查交易对象。 md5:98b621386407ef30
	var tx TX
	tx = X事务从上下文取对象(上下文, c.db.X取配置组名称())
	if tx != nil {
		return tx.X事务(上下文, 回调函数)
	}
	tx, 错误 = c.doBeginCtx(上下文)
	if 错误 != nil {
		return 错误
	}
			// 将事务对象注入上下文。 md5:f5ae21debffd107d
	tx = tx.X设置上下文并取副本(X底层WithTX(tx.X取上下文对象(), tx))
	defer func() {
		if 错误 == nil {
			if exception := recover(); exception != nil {
				if v, ok := exception.(error); ok && gerror.X判断是否带堆栈(v) {
					错误 = v
				} else {
					错误 = gerror.X创建错误码并格式化(gcode.CodeInternalPanic, "%+v", exception)
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

// X底层WithTX 将给定的事务对象注入到上下文中，并返回一个新的上下文。 md5:b4c3c1077b95f681
func X底层WithTX(上下文 context.Context, 事务对象 TX) context.Context {
	if 事务对象 == nil {
		return 上下文
	}
		// 检查从给定的输入中是否存在重复注入。 md5:46e37fcbcbc508b5
	group := 事务对象.X取DB对象().X取配置组名称()
	if ctxTx := X事务从上下文取对象(上下文, group); ctxTx != nil && ctxTx.X取DB对象().X取配置组名称() == group {
		return 上下文
	}
	dbCtx := 事务对象.X取DB对象().X取上下文对象()
	if ctxTx := X事务从上下文取对象(dbCtx, group); ctxTx != nil && ctxTx.X取DB对象().X取配置组名称() == group {
		return dbCtx
	}
		// 向上下文中注入交易对象和ID。 md5:b9cf191f9e07b60d
	上下文 = context.WithValue(上下文, transactionKeyForContext(group), 事务对象)
	return 上下文
}

// X事务从上下文取对象 从上下文中获取并返回事务对象。
// 它通常用于嵌套事务功能，如果之前未设置，则返回nil。
// md5:21e22b68139fc8b6
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

// transactionKeyForContext 为特定数据库组生成并返回一个字符串，用于将事务对象存储在上下文中。 md5:1dc9cbe3d8e29f02
func transactionKeyForContext(group string) string {
	return contextTransactionKeyPrefix + group
}

// transactionKeyForNestedPoint 在当前保存点处生成并返回事务键。 md5:ca81c7094d96c9fc
func (tx *TXCore) transactionKeyForNestedPoint() string {
	return tx.db.X取Core对象().X底层QuoteWord(transactionPointerPrefix + gconv.String(tx.transactionCount))
}

// X设置上下文并取副本为当前事务设置上下文。 md5:da0e9ba442dc74f9
func (tx *TXCore) X设置上下文并取副本(上下文 context.Context) TX {
	tx.ctx = 上下文
	if tx.ctx != nil {
		tx.ctx = tx.db.X取Core对象().injectInternalCtxData(tx.ctx)
	}
	return tx
}

// X取上下文对象 返回当前事务的上下文。 md5:e3cb35516cebab84
func (tx *TXCore) X取上下文对象() context.Context {
	return tx.ctx
}

// X取DB对象 返回当前事务的DB。 md5:26a64f5fed9954b6
func (tx *TXCore) X取DB对象() DB {
	return tx.db
}

// X底层取事务对象 返回当前事务的底层事务对象。 md5:31b14245dcb30833
func (tx *TXCore) X底层取事务对象() *sql.Tx {
	return tx.tx
}

// X事务提交 提交当前事务。
// 注意，如果处于嵌套事务过程中，它会释放之前的保存事务点，
// 否则，它将提交整个事务。
// md5:9ca50fd58870ed9e
func (tx *TXCore) X事务提交() error {
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

// X事务回滚 会回滚当前事务。
// 注意，在嵌套事务过程中，它会回滚当前的事务；否则，它将回滚整个事务。
// md5:0c483721f8447f53
func (tx *TXCore) X事务回滚() error {
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

// X是否已关闭检查并返回此事务是否已经提交或回滚。 md5:cecc2f01ef3e3556
func (tx *TXCore) X是否已关闭() bool {
	return tx.isClosed
}

// X事务开启 启动一个嵌套事务过程。 md5:1b04e48800ebefdd
func (tx *TXCore) X事务开启() error {
	_, err := tx.X原生SQL执行("SAVEPOINT " + tx.transactionKeyForNestedPoint())
	if err != nil {
		return err
	}
	tx.transactionCount++
	return nil
}

// X保存事务点 执行 `SAVEPOINT xxx` SQL 语句，该语句在当前点保存事务。
// 参数 `point` 指定将被保存到服务器的保存点名称。
// md5:f4061450298afabd
func (tx *TXCore) X保存事务点(事务点名称 string) error {
	_, err := tx.X原生SQL执行("SAVEPOINT " + tx.db.X取Core对象().X底层QuoteWord(事务点名称))
	return err
}

// X回滚事务点 执行 `ROLLBACK TO SAVEPOINT xxx` SQL语句，回滚到指定的保存点事务。
// 参数 `point` 指定了之前保存的保存点名称。
// md5:e347c163ad8fefa7
func (tx *TXCore) X回滚事务点(事务点名称 string) error {
	_, err := tx.X原生SQL执行("ROLLBACK TO SAVEPOINT " + tx.db.X取Core对象().X底层QuoteWord(事务点名称))
	return err
}

// X事务 包装了使用函数 `f` 执行的事务逻辑。
// 如果函数 `f` 返回非空错误，它将回滚事务并返回该错误。如果函数 `f` 返回 nil，它将提交事务并返回 nil。
// 
// 注意，在函数 `f` 中不应手动提交或回滚事务，因为这些操作将由这个函数自动处理。
// md5:8906440d4dbbef1f
func (tx *TXCore) X事务(上下文 context.Context, 回调函数 func(上下文 context.Context, tx TX) error) (错误 error) {
	if 上下文 != nil {
		tx.ctx = 上下文
	}
		// 从上下文中检查交易对象。 md5:98b621386407ef30
	if X事务从上下文取对象(tx.ctx, tx.db.X取配置组名称()) == nil {
				// 将事务对象注入上下文。 md5:f5ae21debffd107d
		tx.ctx = X底层WithTX(tx.ctx, tx)
	}
	错误 = tx.X事务开启()
	if 错误 != nil {
		return 错误
	}
	defer func() {
		if 错误 == nil {
			if exception := recover(); exception != nil {
				if v, ok := exception.(error); ok && gerror.X判断是否带堆栈(v) {
					错误 = v
				} else {
					错误 = gerror.X创建错误码并格式化(gcode.CodeInternalPanic, "%+v", exception)
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

// X原生SQL查询 在事务上执行查询操作。
// 请参阅Core.X原生SQL查询。
// md5:0d7503cceb0dc1d6
func (tx *TXCore) X原生SQL查询(sql string, 参数 ...interface{}) (结果 Result, 错误 error) {
	return tx.db.X底层原生SQL查询(tx.ctx, &txLink{tx.tx}, sql, 参数...)
}

// X原生SQL执行 在事务上执行非查询操作。
// 参见Core.X原生SQL执行。
// md5:043edf012223f310
func (tx *TXCore) X原生SQL执行(sql string, 参数 ...interface{}) (sql.Result, error) {
	return tx.db.X底层原生SQL执行(tx.ctx, &txLink{tx.tx}, sql, 参数...)
}

// X原生sql取参数预处理对象 创建一个预处理语句，以便后续的查询或执行。
// 可以从返回的语句中并发地运行多个查询或执行。
// 调用者必须在不再需要该语句时调用语句的 Close 方法。
// md5:16334dc7db1c37a9
func (tx *TXCore) X原生sql取参数预处理对象(sql string) (*Stmt, error) {
	return tx.db.X底层原生sql参数预处理对象(tx.ctx, &txLink{tx.tx}, sql)
}

// GetAll别名 从数据库中查询并返回数据记录。 md5:dfdcfddaa70ab1d4
func (tx *TXCore) GetAll别名(sql string, 参数 ...interface{}) (Result, error) {
	return tx.X原生SQL查询(sql, 参数...)
}

// X原生SQL查询单条记录 从数据库中查询并返回一条记录。 md5:9552f7e095f58141
func (tx *TXCore) X原生SQL查询单条记录(sql string, 参数 ...interface{}) (Record, error) {
	list, err := tx.GetAll别名(sql, 参数...)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	return nil, nil
}

// X原生SQL查询单条到结构体指针 从数据库中查询一条记录，并将其转换为给定的结构体。
// 参数 `pointer` 应该是指向结构体的指针。
// md5:7ddc0d419d6cd2aa
func (tx *TXCore) X原生SQL查询单条到结构体指针(结构体指针 interface{}, sql string, 参数 ...interface{}) error {
	one, err := tx.X原生SQL查询单条记录(sql, 参数...)
	if err != nil {
		return err
	}
	return one.X取结构体指针(结构体指针)
}

// X原生SQL查询到结构体切片指针 从数据库查询记录，并将它们转换为给定的结构体。参数 `pointer` 应该是结构体切片的类型：[]struct 或 []*struct。
// md5:af7dfbf46c6660c6
func (tx *TXCore) X原生SQL查询到结构体切片指针(结构体指针 interface{}, sql string, 参数 ...interface{}) error {
	all, err := tx.GetAll别名(sql, 参数...)
	if err != nil {
		return err
	}
	return all.X取切片结构体指针(结构体指针)
}

// X原生SQL查询到结构体指针 从数据库查询一个或多个记录，并将它们转换为给定的结构体或结构体数组。
//
// 如果参数 `pointer` 是结构体指针类型，它内部会调用 GetStruct 进行转换。如果参数 `pointer` 是切片类型，它内部会调用 GetStructs 进行转换。
// md5:c1dbdab7a7c29c51
func (tx *TXCore) X原生SQL查询到结构体指针(结构体指针 interface{}, sql string, 参数 ...interface{}) error {
	reflectInfo := reflection.OriginTypeAndKind(结构体指针)
	if reflectInfo.InputKind != reflect.Ptr {
		return gerror.X创建错误码并格式化(
			gcode.CodeInvalidParameter,
			"params should be type of pointer, but got: %v",
			reflectInfo.InputKind,
		)
	}
	switch reflectInfo.OriginKind {
	case reflect.Array, reflect.Slice:
		return tx.X原生SQL查询到结构体切片指针(结构体指针, sql, 参数...)

	case reflect.Struct:
		return tx.X原生SQL查询单条到结构体指针(结构体指针, sql, 参数...)
	}
	return gerror.X创建错误码并格式化(
		gcode.CodeInvalidParameter,
		`in valid parameter type "%v", of which element type should be type of struct/slice`,
		reflectInfo.InputType,
	)
}

// X原生SQL查询字段值 从数据库查询并返回字段值。
// SQL 应该只查询数据库中的一个字段，否则它将只返回结果中的一个字段。
// md5:96794360fadbc288
func (tx *TXCore) X原生SQL查询字段值(sql string, 参数 ...interface{}) (Value, error) {
	one, err := tx.X原生SQL查询单条记录(sql, 参数...)
	if err != nil {
		return nil, err
	}
	for _, v := range one {
		return v, nil
	}
	return nil, nil
}

// X原生SQL查询字段计数 从数据库中查询并返回计数。 md5:a8368d39f4a58979
func (tx *TXCore) X原生SQL查询字段计数(sql string, 参数 ...interface{}) (int64, error) {
	if !gregex.X是否匹配文本(`(?i)SELECT\s+COUNT\(.+\)\s+FROM`, sql) {
		sql, _ = gregex.X替换文本(`(?i)(SELECT)\s+(.+)\s+(FROM)`, `$1 COUNT($2) $3`, sql)
	}
	value, err := tx.X原生SQL查询字段值(sql, 参数...)
	if err != nil {
		return 0, err
	}
	return value.X取整数64位(), nil
}

// X插入 执行 "INSERT INTO..." 语句来操作表。
// 如果表中已经存在数据的唯一记录，它会返回错误。
//
// 参数 `data` 可以是 map、gmap、struct、*struct、[]map 或 []struct 等类型。
// 例如：
// Data(g.Map{"uid": 10000, "name": "john"})
// Data(g.Slice{g.Map{"uid": 10000, "name": "john"}, g.Map{"uid": 20000, "name": "smith"}})
//
// 参数 `batch` 在给定数据为切片时，指定批量操作的次数。
// md5:fd75d343f485b8dc
func (tx *TXCore) X插入(表名称 string, 值 interface{}, 批量操作行数 ...int) (sql.Result, error) {
	if len(批量操作行数) > 0 {
		return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X设置数据(值).X设置批量操作行数(批量操作行数[0]).X插入()
	}
	return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X设置数据(值).X插入()
}

// X插入并跳过已存在 执行 "INSERT IGNORE INTO ..." 语句针对该表。
// 如果表中已存在该数据的唯一记录，则忽略插入操作。
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 例如：
// Data(g.Map{"uid": 10000, "name": "john"})
// Data(g.Slice{g.Map{"uid": 10000, "name": "john"}, g.Map{"uid": 20000, "name": "smith"})
//
// 当给定的数据为切片时，参数 `batch` 指定批处理操作的计数。
// md5:49f76901041c9819
func (tx *TXCore) X插入并跳过已存在(表名称 string, 值 interface{}, 批量操作行数 ...int) (sql.Result, error) {
	if len(批量操作行数) > 0 {
		return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X设置数据(值).X设置批量操作行数(批量操作行数[0]).X插入并跳过已存在()
	}
	return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X设置数据(值).X插入并跳过已存在()
}

// X插入并取ID 执行插入操作，并返回自动生成的最后一个插入id。 md5:8d00b40a35fa48a5
func (tx *TXCore) X插入并取ID(表名称 string, 值 interface{}, 批量操作行数 ...int) (int64, error) {
	if len(批量操作行数) > 0 {
		return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X设置数据(值).X设置批量操作行数(批量操作行数[0]).X插入并取ID()
	}
	return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X设置数据(值).X插入并取ID()
}

// X插入并替换已存在 用于执行针对该表的 "REPLACE INTO..." 语句。如果表中已经存在数据的唯一记录，它会删除该记录并插入新的。
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。例如：
// Data(g.Map{"uid": 10000, "name": "john"})
// Data(g.Slice{g.Map{"uid": 10000, "name": "john"}, g.Map{"uid": 20000, "name": "smith"}})
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。如果给定的数据是切片类型，它将进行批量替换，并可选地通过参数 `batch` 指定批量操作次数。
// md5:69ecd0994eab5bbb
func (tx *TXCore) X插入并替换已存在(表名称 string, 值 interface{}, 批量操作行数 ...int) (sql.Result, error) {
	if len(批量操作行数) > 0 {
		return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X设置数据(值).X设置批量操作行数(批量操作行数[0]).X插入并替换已存在()
	}
	return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X设置数据(值).X插入并替换已存在()
}

// X插入并更新已存在 执行 "INSERT INTO ... ON DUPLICATE KEY UPDATE..." 语句来操作表。
// 如果保存的数据中存在主键或唯一索引，它将更新记录；否则，将在表中插入新记录。
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 例如：
// Data(g.Map{"uid": 10000, "name": "john"})
// Data(g.Slice{g.Map{"uid": 10000, "name": "john"}, g.Map{"uid": 20000, "name": "smith"}})
//
// 如果给定的数据是切片类型，它将进行批量保存。可选参数 `batch` 指定了批量操作的次数。
// md5:c76721f5e0b01424
func (tx *TXCore) X插入并更新已存在(表名称 string, 值 interface{}, 批量操作行数 ...int) (sql.Result, error) {
	if len(批量操作行数) > 0 {
		return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X设置数据(值).X设置批量操作行数(批量操作行数[0]).X插入并更新已存在()
	}
	return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X设置数据(值).X插入并更新已存在()
}

// X更新 执行表的 "UPDATE ... " 语句。
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
func (tx *TXCore) X更新(表名称 string, 值 interface{}, 条件 interface{}, 参数 ...interface{}) (sql.Result, error) {
	return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X设置数据(值).X条件(条件, 参数...).X更新()
}

// X删除 执行 "DELETE FROM ... " 语句针对该表。
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
func (tx *TXCore) X删除(表名称 string, 条件 interface{}, 参数 ...interface{}) (sql.Result, error) {
	return tx.X创建Model对象(表名称).X设置上下文并取副本(tx.ctx).X条件(条件, 参数...).X删除()
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
