// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//

package gdb

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
func (c *Core) Query(ctx context.Context, sql string, args ...interface{}) (result Result, err error) {
	return c.db.DoQuery(ctx, nil, sql, args...)
}

// DoQuery 通过给定的link对象，将sql字符串及其参数提交到底层驱动，并返回执行结果。
func (c *Core) DoQuery(ctx context.Context, link Link, sql string, args ...interface{}) (result Result, err error) {
	// 事务检查。
	if link == nil {
		if tx := TXFromCtx(ctx, c.db.GetGroup()); tx != nil {
			// 首先，从上下文检查并检索交易链接。
			link = &txLink{tx.GetSqlTX()}
		} else if link, err = c.SlaveLink(); err != nil {
			// 或者从主节点创建一个
			return nil, err
		}
	} else if !link.IsTransaction() {
		// 如果当前链接不是事务链接，它会检查并从上下文中获取事务。
		if tx := TXFromCtx(ctx, c.db.GetGroup()); tx != nil {
			link = &txLink{tx.GetSqlTX()}
		}
	}

	if c.db.GetConfig().QueryTimeout > 0 {
		ctx, _ = context.WithTimeout(ctx, c.db.GetConfig().QueryTimeout)
	}

	// Sql filtering.
	sql, args = c.FormatSqlBeforeExecuting(sql, args)
	sql, args, err = c.db.DoFilter(ctx, link, sql, args)
	if err != nil {
		return nil, err
	}
	// SQL格式化并检索
	if v := ctx.Value(ctxKeyCatchSQL); v != nil {
		var (
			manager      = v.(*CatchSQLManager)
			formattedSql = FormatSqlWithArgs(sql, args)
		)
		manager.SQLArray.Append(formattedSql)
		if !manager.DoCommit && ctx.Value(ctxKeyInternalProducedSQL) == nil {
			return nil, nil
		}
	}
	// Link execution.
	var out DoCommitOutput
	out, err = c.db.DoCommit(ctx, DoCommitInput{
		Link:          link,
		Sql:           sql,
		Args:          args,
		Stmt:          nil,
		Type:          SqlTypeQueryContext,
		IsTransaction: link.IsTransaction(),
	})
	return out.Records, err
}

// Exec方法将一个SQL查询语句提交给底层驱动执行并返回执行结果。
// 该方法主要用于数据的插入和更新操作。
func (c *Core) Exec(ctx context.Context, sql string, args ...interface{}) (result sql.Result, err error) {
	return c.db.DoExec(ctx, nil, sql, args...)
}

// DoExec通过给定的link对象，将SQL字符串及其参数提交给底层驱动，并返回执行结果。
func (c *Core) DoExec(ctx context.Context, link Link, sql string, args ...interface{}) (result sql.Result, err error) {
	// 事务检查。
	if link == nil {
		if tx := TXFromCtx(ctx, c.db.GetGroup()); tx != nil {
			// 首先，从上下文检查并检索交易链接。
			link = &txLink{tx.GetSqlTX()}
		} else if link, err = c.MasterLink(); err != nil {
			// 或者从主节点创建一个
			return nil, err
		}
	} else if !link.IsTransaction() {
		// 如果当前链接不是事务链接，它会检查并从上下文中获取事务。
		if tx := TXFromCtx(ctx, c.db.GetGroup()); tx != nil {
			link = &txLink{tx.GetSqlTX()}
		}
	}

	if c.db.GetConfig().ExecTimeout > 0 {
		var cancelFunc context.CancelFunc
		ctx, cancelFunc = context.WithTimeout(ctx, c.db.GetConfig().ExecTimeout)
		defer cancelFunc()
	}

	// SQL filtering.
	sql, args = c.FormatSqlBeforeExecuting(sql, args)
	sql, args, err = c.db.DoFilter(ctx, link, sql, args)
	if err != nil {
		return nil, err
	}
	// SQL格式化并检索
	if v := ctx.Value(ctxKeyCatchSQL); v != nil {
		var (
			manager      = v.(*CatchSQLManager)
			formattedSql = FormatSqlWithArgs(sql, args)
		)
		manager.SQLArray.Append(formattedSql)
		if !manager.DoCommit && ctx.Value(ctxKeyInternalProducedSQL) == nil {
			return new(SqlResult), nil
		}
	}
	// Link execution.
	var out DoCommitOutput
	out, err = c.db.DoCommit(ctx, DoCommitInput{
		Link:          link,
		Sql:           sql,
		Args:          args,
		Stmt:          nil,
		Type:          SqlTypeExecContext,
		IsTransaction: link.IsTransaction(),
	})
	return out.Result, err
}

// DoFilter 是一个钩子函数，在 SQL 语句及其参数提交给底层驱动程序之前对其进行过滤。
// 参数 `link` 指定了当前数据库连接操作对象。您可以在 SQL 字符串 `sql` 和其参数 `args` 提交给驱动程序之前，根据需要自由修改它们。
func (c *Core) DoFilter(ctx context.Context, link Link, sql string, args []interface{}) (newSql string, newArgs []interface{}, err error) {
	return sql, args, nil
}

// DoCommit 将当前SQL语句及其参数提交给底层SQL驱动执行。
func (c *Core) DoCommit(ctx context.Context, in DoCommitInput) (out DoCommitOutput, err error) {
	// 将内部数据注入到ctx中，特别是用于创建事务。
	ctx = c.InjectInternalCtxData(ctx)

	var (
		sqlTx                *sql.Tx
		sqlStmt              *sql.Stmt
		sqlRows              *sql.Rows
		sqlResult            sql.Result
		stmtSqlRows          *sql.Rows
		stmtSqlRow           *sql.Row
		rowsAffected         int64
		cancelFuncForTimeout context.CancelFunc
		formattedSql         = FormatSqlWithArgs(in.Sql, in.Args)
		timestampMilli1      = gtime.TimestampMilli()
	)

	// Trace span start.
	tr := otel.GetTracerProvider().Tracer(traceInstrumentName, trace.WithInstrumentationVersion(gf.VERSION))
	ctx, span := tr.Start(ctx, in.Type, trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// 根据类型执行。
	switch in.Type {
	case SqlTypeBegin:
		if sqlTx, err = in.Db.Begin(); err == nil {
			out.Tx = &TXCore{
				db:            c.db,
				tx:            sqlTx,
				ctx:           context.WithValue(ctx, transactionIdForLoggerCtx, transactionIdGenerator.Add(1)),
				master:        in.Db,
				transactionId: guid.S(),
			}
			ctx = out.Tx.GetCtx()
		}
		out.RawResult = sqlTx

	case SqlTypeTXCommit:
		err = in.Tx.Commit()

	case SqlTypeTXRollback:
		err = in.Tx.Rollback()

	case SqlTypeExecContext:
		if c.db.GetDryRun() {
			sqlResult = new(SqlResult)
		} else {
			sqlResult, err = in.Link.ExecContext(ctx, in.Sql, in.Args...)
		}
		out.RawResult = sqlResult

	case SqlTypeQueryContext:
		sqlRows, err = in.Link.QueryContext(ctx, in.Sql, in.Args...)
		out.RawResult = sqlRows

	case SqlTypePrepareContext:
		sqlStmt, err = in.Link.PrepareContext(ctx, in.Sql)
		out.RawResult = sqlStmt

	case SqlTypeStmtExecContext:
		ctx, cancelFuncForTimeout = c.GetCtxTimeout(ctx, ctxTimeoutTypeExec)
		defer cancelFuncForTimeout()
		if c.db.GetDryRun() {
			sqlResult = new(SqlResult)
		} else {
			sqlResult, err = in.Stmt.ExecContext(ctx, in.Args...)
		}
		out.RawResult = sqlResult

	case SqlTypeStmtQueryContext:
		ctx, cancelFuncForTimeout = c.GetCtxTimeout(ctx, ctxTimeoutTypeQuery)
		defer cancelFuncForTimeout()
		stmtSqlRows, err = in.Stmt.QueryContext(ctx, in.Args...)
		out.RawResult = stmtSqlRows

	case SqlTypeStmtQueryRowContext:
		ctx, cancelFuncForTimeout = c.GetCtxTimeout(ctx, ctxTimeoutTypeQuery)
		defer cancelFuncForTimeout()
		stmtSqlRow = in.Stmt.QueryRowContext(ctx, in.Args...)
		out.RawResult = stmtSqlRow

	default:
		panic(gerror.NewCodef(gcode.CodeInvalidParameter, `invalid SqlType "%s"`, in.Type))
	}
	// Result handling.
	switch {
	case sqlResult != nil && !c.GetIgnoreResultFromCtx(ctx):
		rowsAffected, err = sqlResult.RowsAffected()
		out.Result = sqlResult

	case sqlRows != nil:
		out.Records, err = c.RowsToResult(ctx, sqlRows)
		rowsAffected = int64(len(out.Records))

	case sqlStmt != nil:
		out.Stmt = &Stmt{
			Stmt: sqlStmt,
			core: c,
			link: in.Link,
			sql:  in.Sql,
		}
	}
	var (
		timestampMilli2 = gtime.TimestampMilli()
		sqlObj          = &Sql{
			Sql:           in.Sql,
			Type:          in.Type,
			Args:          in.Args,
			Format:        formattedSql,
			Error:         err,
			Start:         timestampMilli1,
			End:           timestampMilli2,
			Group:         c.db.GetGroup(),
			Schema:        c.db.GetSchema(),
			RowsAffected:  rowsAffected,
			IsTransaction: in.IsTransaction,
		}
	)

	// Tracing.
	c.traceSpanEnd(ctx, span, sqlObj)

	// Logging.
	if c.db.GetDebug() {
		c.writeSqlToLogger(ctx, sqlObj)
	}
	if err != nil && err != sql.ErrNoRows {
		err = gerror.WrapCode(
			gcode.CodeDbOperationError,
			err,
			FormatSqlWithArgs(in.Sql, in.Args),
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
func (c *Core) Prepare(ctx context.Context, sql string, execOnMaster ...bool) (*Stmt, error) {
	var (
		err  error
		link Link
	)
	if len(execOnMaster) > 0 && execOnMaster[0] {
		if link, err = c.MasterLink(); err != nil {
			return nil, err
		}
	} else {
		if link, err = c.SlaveLink(); err != nil {
			return nil, err
		}
	}
	return c.db.DoPrepare(ctx, link, sql)
}

// DoPrepare在给定的link对象上调用prepare函数，并返回statement对象。
func (c *Core) DoPrepare(ctx context.Context, link Link, sql string) (stmt *Stmt, err error) {
	// 事务检查。
	if link == nil {
		if tx := TXFromCtx(ctx, c.db.GetGroup()); tx != nil {
			// 首先，从上下文检查并检索交易链接。
			link = &txLink{tx.GetSqlTX()}
		} else {
			// 或者从主节点创建一个
			var err error
			if link, err = c.MasterLink(); err != nil {
				return nil, err
			}
		}
	} else if !link.IsTransaction() {
		// 如果当前链接不是事务链接，它会检查并从上下文中获取事务。
		if tx := TXFromCtx(ctx, c.db.GetGroup()); tx != nil {
			link = &txLink{tx.GetSqlTX()}
		}
	}

	if c.db.GetConfig().PrepareTimeout > 0 {
		// **请勿在预处理语句中使用取消函数。**
		ctx, _ = context.WithTimeout(ctx, c.db.GetConfig().PrepareTimeout)
	}

	// Link execution.
	var out DoCommitOutput
	out, err = c.db.DoCommit(ctx, DoCommitInput{
		Link:          link,
		Sql:           sql,
		Type:          SqlTypePrepareContext,
		IsTransaction: link.IsTransaction(),
	})
	return out.Stmt, err
}

// RowsToResult 将底层数据记录类型 sql.Rows 转换为 Result 类型。
func (c *Core) RowsToResult(ctx context.Context, rows *sql.Rows) (Result, error) {
	if rows == nil {
		return nil, nil
	}
	defer func() {
		if err := rows.Close(); err != nil {
			intlog.Errorf(ctx, `%+v`, err)
		}
	}()
	if !rows.Next() {
		return nil, nil
	}
	// 列名和类型。
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	if len(columnTypes) > 0 {
		if internalData := c.GetInternalCtxDataFromCtx(ctx); internalData != nil {
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
		if err = rows.Scan(scanArgs...); err != nil {
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
				if convertedValue, err = c.columnValueToLocalValue(ctx, value, columnTypes[i]); err != nil {
					return nil, err
				}
				record[columnTypes[i].Name()] = gvar.New(convertedValue)
			}
		}
		result = append(result, record)
		if !rows.Next() {
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
			return gconv.Convert(
				gconv.String(value),
				columnType.ScanType().String(),
			), nil
		}
	}
	// 其他复杂类型，特别是自定义类型。
	return c.db.ConvertValueForLocal(ctx, columnType.DatabaseTypeName(), value)
}
