// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//

package db类

import (
	"context"
	"database/sql"
	"reflect"
	
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	
	"github.com/888go/goframe/util/gconv"
	
	"github.com/888go/goframe"
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/util/guid"
)

// Query 将一个查询SQL语句提交给底层驱动并返回执行结果。
// 这个方法最常用于数据查询。
func (c *Core) X原生SQL查询(上下文 context.Context, sql string, 参数 ...interface{}) (结果 Result, 错误 error) {
	return c.db.X底层原生SQL查询(上下文, nil, sql, 参数...)
}

// DoQuery 通过给定的link对象，将sql字符串及其参数提交到底层驱动，并返回执行结果。
func (c *Core) X底层原生SQL查询(上下文 context.Context, 链接 Link, sql string, 参数 ...interface{}) (结果 Result, 错误 error) {
	// 事务检查。
	if 链接 == nil {
		if tx := X事务从上下文取对象(上下文, c.db.X取配置组名称()); tx != nil {
			// 首先，从上下文检查并检索交易链接。
			链接 = &txLink{tx.X底层取事务对象()}
		} else if 链接, 错误 = c.X底层SlaveLink(); 错误 != nil {
			// 或者从主节点创建一个
			return nil, 错误
		}
	} else if !链接.IsTransaction() {
		// 如果当前链接不是事务链接，它会检查并从上下文中获取事务。
		if tx := X事务从上下文取对象(上下文, c.db.X取配置组名称()); tx != nil {
			链接 = &txLink{tx.X底层取事务对象()}
		}
	}

	if c.db.X取当前节点配置().X查询超时时长 > 0 {
		上下文, _ = context.WithTimeout(上下文, c.db.X取当前节点配置().X查询超时时长)
	}

	// Sql filtering.
	sql, 参数 = c.X格式化Sql(sql, 参数)
	sql, 参数, 错误 = c.db.X底层DoFilter(上下文, 链接, sql, 参数)
	if 错误 != nil {
		return nil, 错误
	}
	// SQL格式化并检索
	if v := 上下文.Value(ctxKeyCatchSQL); v != nil {
		var (
			manager      = v.(*CatchSQLManager)
			formattedSql = X格式化Sql(sql, 参数)
		)
		manager.SQLArray.Append别名(formattedSql)
		if !manager.DoCommit && 上下文.Value(ctxKeyInternalProducedSQL) == nil {
			return nil, nil
		}
	}
	// Link execution.
	var out X输出
	out, 错误 = c.db.X底层DoCommit(上下文, DoCommitInput{
		Link:          链接,
		Sql:           sql,
		Args:          参数,
		Stmt:          nil,
		Type:          SqlTypeQueryContext,
		IsTransaction: 链接.IsTransaction(),
	})
	return out.X行记录数组, 错误
}

// Exec方法将一个SQL查询语句提交给底层驱动执行并返回执行结果。
// 该方法主要用于数据的插入和更新操作。
func (c *Core) X原生SQL执行(上下文 context.Context, sql string, 参数 ...interface{}) (结果 sql.Result, 错误 error) {
	return c.db.X底层原生SQL执行(上下文, nil, sql, 参数...)
}

// DoExec通过给定的link对象，将SQL字符串及其参数提交给底层驱动，并返回执行结果。
func (c *Core) X底层原生SQL执行(上下文 context.Context, 链接 Link, sql string, 参数 ...interface{}) (结果 sql.Result, 错误 error) {
	// 事务检查。
	if 链接 == nil {
		if tx := X事务从上下文取对象(上下文, c.db.X取配置组名称()); tx != nil {
			// 首先，从上下文检查并检索交易链接。
			链接 = &txLink{tx.X底层取事务对象()}
		} else if 链接, 错误 = c.X底层MasterLink(); 错误 != nil {
			// 或者从主节点创建一个
			return nil, 错误
		}
	} else if !链接.IsTransaction() {
		// 如果当前链接不是事务链接，它会检查并从上下文中获取事务。
		if tx := X事务从上下文取对象(上下文, c.db.X取配置组名称()); tx != nil {
			链接 = &txLink{tx.X底层取事务对象()}
		}
	}

	if c.db.X取当前节点配置().X执行超时时长 > 0 {
		var cancelFunc context.CancelFunc
		上下文, cancelFunc = context.WithTimeout(上下文, c.db.X取当前节点配置().X执行超时时长)
		defer cancelFunc()
	}

	// SQL filtering.
	sql, 参数 = c.X格式化Sql(sql, 参数)
	sql, 参数, 错误 = c.db.X底层DoFilter(上下文, 链接, sql, 参数)
	if 错误 != nil {
		return nil, 错误
	}
	// SQL格式化并检索
	if v := 上下文.Value(ctxKeyCatchSQL); v != nil {
		var (
			manager      = v.(*CatchSQLManager)
			formattedSql = X格式化Sql(sql, 参数)
		)
		manager.SQLArray.Append别名(formattedSql)
		if !manager.DoCommit && 上下文.Value(ctxKeyInternalProducedSQL) == nil {
			return new(Sql执行结果), nil
		}
	}
	// Link execution.
	var out X输出
	out, 错误 = c.db.X底层DoCommit(上下文, DoCommitInput{
		Link:          链接,
		Sql:           sql,
		Args:          参数,
		Stmt:          nil,
		Type:          SqlTypeExecContext,
		IsTransaction: 链接.IsTransaction(),
	})
	return out.X原生sql行记录, 错误
}

// DoFilter 是一个钩子函数，在 SQL 语句及其参数提交给底层驱动程序之前对其进行过滤。
// 参数 `link` 指定了当前数据库连接操作对象。您可以在 SQL 字符串 `sql` 和其参数 `args` 提交给驱动程序之前，根据需要自由修改它们。
func (c *Core) X底层DoFilter(ctx context.Context, link Link, sql string, args []interface{}) (newSql string, newArgs []interface{}, err error) {
	return sql, args, nil
}

// DoCommit 将当前SQL语句及其参数提交给底层SQL驱动执行。
func (c *Core) X底层DoCommit(ctx context.Context, in DoCommitInput) (out X输出, err error) {
	// 将内部数据注入到ctx中，特别是用于创建事务。
	ctx = c.底层_InjectInternalCtxData(ctx)

	var (
		sqlTx                *sql.Tx
		sqlStmt              *sql.Stmt
		sqlRows              *sql.Rows
		sqlResult            sql.Result
		stmtSqlRows          *sql.Rows
		stmtSqlRow           *sql.Row
		rowsAffected         int64
		cancelFuncForTimeout context.CancelFunc
		formattedSql         = X格式化Sql(in.Sql, in.Args)
		timestampMilli1      = 时间类.X取时间戳毫秒()
	)

	// Trace span start.
	tr := otel.GetTracerProvider().Tracer(traceInstrumentName, trace.WithInstrumentationVersion(gf.VERSION))
	ctx, span := tr.Start(ctx, in.Type, trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// 根据类型执行。
	switch in.Type {
	case SqlTypeBegin:
		if sqlTx, err = in.Db.Begin(); err == nil {
			out.Tx = &X基础事务{
				db:            c.db,
				tx:            sqlTx,
				ctx:           context.WithValue(ctx, transactionIdForLoggerCtx, transactionIdGenerator.Add(1)),
				master:        in.Db,
				transactionId: uid类.X生成(),
			}
			ctx = out.Tx.X取上下文对象()
		}
		out.X底层结果 = sqlTx

	case SqlTypeTXCommit:
		err = in.Tx.Commit()

	case SqlTypeTXRollback:
		err = in.Tx.Rollback()

	case SqlTypeExecContext:
		if c.db.X取空跑特性() {
			sqlResult = new(Sql执行结果)
		} else {
			sqlResult, err = in.Link.ExecContext(ctx, in.Sql, in.Args...)
		}
		out.X底层结果 = sqlResult

	case SqlTypeQueryContext:
		sqlRows, err = in.Link.QueryContext(ctx, in.Sql, in.Args...)
		out.X底层结果 = sqlRows

	case SqlTypePrepareContext:
		sqlStmt, err = in.Link.PrepareContext(ctx, in.Sql)
		out.X底层结果 = sqlStmt

	case SqlTypeStmtExecContext:
		ctx, cancelFuncForTimeout = c.X取超时上下文对象(ctx, ctxTimeoutTypeExec)
		defer cancelFuncForTimeout()
		if c.db.X取空跑特性() {
			sqlResult = new(Sql执行结果)
		} else {
			sqlResult, err = in.Stmt.ExecContext(ctx, in.Args...)
		}
		out.X底层结果 = sqlResult

	case SqlTypeStmtQueryContext:
		ctx, cancelFuncForTimeout = c.X取超时上下文对象(ctx, ctxTimeoutTypeQuery)
		defer cancelFuncForTimeout()
		stmtSqlRows, err = in.Stmt.QueryContext(ctx, in.Args...)
		out.X底层结果 = stmtSqlRows

	case SqlTypeStmtQueryRowContext:
		ctx, cancelFuncForTimeout = c.X取超时上下文对象(ctx, ctxTimeoutTypeQuery)
		defer cancelFuncForTimeout()
		stmtSqlRow = in.Stmt.QueryRowContext(ctx, in.Args...)
		out.X底层结果 = stmtSqlRow

	default:
		panic(错误类.X创建错误码并格式化(错误码类.CodeInvalidParameter, `invalid SqlType "%s"`, in.Type))
	}
	// Result handling.
	switch {
	case sqlResult != nil && !c.底层_GetIgnoreResultFromCtx(ctx):
		rowsAffected, err = sqlResult.RowsAffected()
		out.X原生sql行记录 = sqlResult

	case sqlRows != nil:
		out.X行记录数组, err = c.X原生sql记录到行记录数组对象(ctx, sqlRows)
		rowsAffected = int64(len(out.X行记录数组))

	case sqlStmt != nil:
		out.X参数预处理 = &Stmt{
			Stmt: sqlStmt,
			core: c,
			link: in.Link,
			sql:  in.Sql,
		}
	}
	var (
		timestampMilli2 = 时间类.X取时间戳毫秒()
		sqlObj          = &Sql{
			Sql:           in.Sql,
			Type:          in.Type,
			Args:          in.Args,
			Format:        formattedSql,
			Error:         err,
			Start:         timestampMilli1,
			End:           timestampMilli2,
			Group:         c.db.X取配置组名称(),
			Schema:        c.db.X取默认数据库名称(),
			RowsAffected:  rowsAffected,
			IsTransaction: in.IsTransaction,
		}
	)

	// Tracing.
	c.traceSpanEnd(ctx, span, sqlObj)

	// Logging.
	if c.db.X取调试模式() {
		c.writeSqlToLogger(ctx, sqlObj)
	}
	if err != nil && err != sql.ErrNoRows {
		err = 错误类.X多层错误码(
			错误码类.CodeDbOperationError,
			err,
			X格式化Sql(in.Sql, in.Args),
		)
	}
	return out, err
}

// Prepare 函数用于为后续查询或执行创建预编译语句。
// 从返回的语句可以并发地运行多个查询或执行操作。
// 当该语句不再需要时，调用者必须调用该语句的 Close 方法。
//
// 参数 `execOnMaster` 指定是否在主节点上执行 SQL，如果配置了主从模式，
// 则此参数为 false 时将在从节点上执行 SQL。
func (c *Core) X原生sql取参数预处理对象(上下文 context.Context, sql string, 是否主节点执行 ...bool) (*Stmt, error) {
	var (
		err  error
		link Link
	)
	if len(是否主节点执行) > 0 && 是否主节点执行[0] {
		if link, err = c.X底层MasterLink(); err != nil {
			return nil, err
		}
	} else {
		if link, err = c.X底层SlaveLink(); err != nil {
			return nil, err
		}
	}
	return c.db.X底层原生sql参数预处理对象(上下文, link, sql)
}

// DoPrepare在给定的link对象上调用prepare函数，并返回statement对象。
func (c *Core) X底层原生sql参数预处理对象(上下文 context.Context, 链接 Link, sql string) (参数预处理 *Stmt, 错误 error) {
	// 事务检查。
	if 链接 == nil {
		if tx := X事务从上下文取对象(上下文, c.db.X取配置组名称()); tx != nil {
			// 首先，从上下文检查并检索交易链接。
			链接 = &txLink{tx.X底层取事务对象()}
		} else {
			// 或者从主节点创建一个
			var err error
			if 链接, err = c.X底层MasterLink(); err != nil {
				return nil, err
			}
		}
	} else if !链接.IsTransaction() {
		// 如果当前链接不是事务链接，它会检查并从上下文中获取事务。
		if tx := X事务从上下文取对象(上下文, c.db.X取配置组名称()); tx != nil {
			链接 = &txLink{tx.X底层取事务对象()}
		}
	}

	if c.db.X取当前节点配置().X预准备SQL超时时长 > 0 {
		// **请勿在预处理语句中使用取消函数。**
		上下文, _ = context.WithTimeout(上下文, c.db.X取当前节点配置().X预准备SQL超时时长)
	}

	// Link execution.
	var out X输出
	out, 错误 = c.db.X底层DoCommit(上下文, DoCommitInput{
		Link:          链接,
		Sql:           sql,
		Type:          SqlTypePrepareContext,
		IsTransaction: 链接.IsTransaction(),
	})
	return out.X参数预处理, 错误
}

// RowsToResult 将底层数据记录类型 sql.Rows 转换为 Result 类型。
func (c *Core) X原生sql记录到行记录数组对象(上下文 context.Context, 底层数据记录 *sql.Rows) (Result, error) {
	if 底层数据记录 == nil {
		return nil, nil
	}
	defer func() {
		if err := 底层数据记录.Close(); err != nil {
			intlog.Errorf(上下文, `%+v`, err)
		}
	}()
	if !底层数据记录.Next() {
		return nil, nil
	}
	// 列名和类型。
	columnTypes, err := 底层数据记录.ColumnTypes()
	if err != nil {
		return nil, err
	}

	if len(columnTypes) > 0 {
		if internalData := c.底层_GetInternalCtxDataFromCtx(上下文); internalData != nil {
			internalData.FirstResultColumn = columnTypes[0].Name()
		}
	}
	var (
		values   = make([]interface{}, len(columnTypes))
		result   = make(Result, 0)
		scanArgs = make([]interface{}, len(values))
	)
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for {
		if err = 底层数据记录.Scan(scanArgs...); err != nil {
			return result, err
		}
		record := Record{}
		for i, value := range values {
			if value == nil {
// **注意**：在此处不要使用 `gvar.New(nil)`，因为它会创建一个已初始化的对象，
// 这将会导致结构体转换问题。
				record[columnTypes[i].Name()] = nil
			} else {
				var convertedValue interface{}
				if convertedValue, err = c.columnValueToLocalValue(上下文, value, columnTypes[i]); err != nil {
					return nil, err
				}
				record[columnTypes[i].Name()] = 泛型类.X创建(convertedValue)
			}
		}
		result = append(result, record)
		if !底层数据记录.Next() {
			break
		}
	}
	return result, nil
}

func (c *Core) columnValueToLocalValue(ctx context.Context, value interface{}, columnType *sql.ColumnType) (interface{}, error) {
	var scanType = columnType.ScanType()
	if scanType != nil {
		// 常见的基本内置类型。
		switch scanType.Kind() {
		case
			reflect.Bool,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			return 转换类.X按名称转换(
				转换类.String(value),
				columnType.ScanType().String(),
			), nil
		}
	}
	// 其他复杂类型，特别是自定义类型。
	return c.db.X底层ConvertValueForLocal(ctx, columnType.DatabaseTypeName(), value)
}
