// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

package gdb

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
)

// Query 向底层驱动提交一个查询SQL并返回执行结果。
// 它最常用于数据查询。
// md5:06bbbfc29aa3894b
func (c *Core) Query(ctx context.Context, sql string, args ...interface{}) (result Result, err error) {
	return c.db.DoQuery(ctx, nil, sql, args...)
}

// DoQuery 通过给定的链接对象将SQL字符串及其参数提交给底层驱动，并返回执行结果。
// md5:af7bdcd1a2074bc0
func (c *Core) DoQuery(ctx context.Context, link Link, sql string, args ...interface{}) (result Result, err error) {
	// Transaction checks.
	if link == nil {
		if tx := TXFromCtx(ctx, c.db.GetGroup()); tx != nil {
						// 首先，从上下文中检查并获取交易链接。 md5:9ac4c60388fa960d
			link = &txLink{tx.GetSqlTX()}
		} else if link, err = c.SlaveLink(); err != nil {
						// 否则，它将从主节点创建一个。 md5:4bd14606783b43fc
			return nil, err
		}
	} else if !link.IsTransaction() {
				// 如果当前链接不是事务链接，它会检查并从上下文中检索事务。 md5:e3c484ab061699a1
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
		// SQL格式化和获取。 md5:815f530302ec8a7e
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

// Exec 将一个查询 SQL 执行到底层驱动并返回执行结果。它最常用于数据插入和更新。
// md5:6f9ddc85964b9797
func (c *Core) Exec(ctx context.Context, sql string, args ...interface{}) (result sql.Result, err error) {
	return c.db.DoExec(ctx, nil, sql, args...)
}

// DoExec 通过给定的链接对象将 sql 字符串及其参数提交到底层驱动，并返回执行结果。
// md5:947bd2b83e751e10
func (c *Core) DoExec(ctx context.Context, link Link, sql string, args ...interface{}) (result sql.Result, err error) {
	// Transaction checks.
	if link == nil {
		if tx := TXFromCtx(ctx, c.db.GetGroup()); tx != nil {
						// 首先，从上下文中检查并获取交易链接。 md5:9ac4c60388fa960d
			link = &txLink{tx.GetSqlTX()}
		} else if link, err = c.MasterLink(); err != nil {
						// 否则，它将从主节点创建一个。 md5:4bd14606783b43fc
			return nil, err
		}
	} else if !link.IsTransaction() {
				// 如果当前链接不是事务链接，它会检查并从上下文中检索事务。 md5:e3c484ab061699a1
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
		// SQL格式化和获取。 md5:815f530302ec8a7e
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

// DoFilter 是一个钩子函数，它在 SQL 语句及其参数提交给底层驱动之前进行过滤。
// 参数 `link` 指定当前数据库连接的操作对象。在 SQL 语句 `sql` 及其参数 `args` 被提交给驱动之前，您可以根据需要随意修改它们。
// md5:41118fbc4e6c5562
func (c *Core) DoFilter(ctx context.Context, link Link, sql string, args []interface{}) (newSql string, newArgs []interface{}, err error) {
	return sql, args, nil
}

// DoCommit 将当前SQL和参数提交给底层SQL驱动程序。 md5:7cf9b1f6f4d9d2cb
func (c *Core) DoCommit(ctx context.Context, in DoCommitInput) (out DoCommitOutput, err error) {
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
	ctx, span := tr.Start(ctx, string(in.Type), trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

		// 根据类型执行。 md5:4f8bf756ef2da0c5
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

// Prepare 准备一个预编译语句，供后续查询或执行使用。
// 可以从返回的语句对象并发运行多个查询或执行。
// 当不再需要语句时，调用者必须调用 statement 的 Close 方法。
// 
// 参数 `execOnMaster` 指定是否在主节点上执行 SQL，如果配置了主从复制，则在从节点上执行。
// md5:639eebcae369b0a2
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

// DoPrepare 会调用给定链接对象上的prepare函数，并返回语句对象。 md5:bae03ede256987bd
func (c *Core) DoPrepare(ctx context.Context, link Link, sql string) (stmt *Stmt, err error) {
	// Transaction checks.
	if link == nil {
		if tx := TXFromCtx(ctx, c.db.GetGroup()); tx != nil {
						// 首先，从上下文中检查并获取交易链接。 md5:9ac4c60388fa960d
			link = &txLink{tx.GetSqlTX()}
		} else {
						// 否则，它将从主节点创建一个。 md5:4bd14606783b43fc
			var err error
			if link, err = c.MasterLink(); err != nil {
				return nil, err
			}
		}
	} else if !link.IsTransaction() {
				// 如果当前链接不是事务链接，它会检查并从上下文中检索事务。 md5:e3c484ab061699a1
		if tx := TXFromCtx(ctx, c.db.GetGroup()); tx != nil {
			link = &txLink{tx.GetSqlTX()}
		}
	}

	if c.db.GetConfig().PrepareTimeout > 0 {
				// 不要在预处理语句中使用取消函数。 md5:5e529fe5094c7942
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

// FormatUpsert 格式化并返回用于 UPSERT 语句的 SQL 子句部分。
// 在默认实现中，此函数执行类似 MySQL 的 UPSERT 语句：
// `INSERT INTO ... ON DUPLICATE KEY UPDATE x=VALUES(z),m=VALUES(y)...`
// md5:c1c6d7b14661682b
func (c *Core) FormatUpsert(columns []string, list List, option DoInsertOption) (string, error) {
	var onDuplicateStr string
	if option.OnDuplicateStr != "" {
		onDuplicateStr = option.OnDuplicateStr
	} else if len(option.OnDuplicateMap) > 0 {
		for k, v := range option.OnDuplicateMap {
			if len(onDuplicateStr) > 0 {
				onDuplicateStr += ","
			}
			switch v.(type) {
			case Raw, *Raw:
				onDuplicateStr += fmt.Sprintf(
					"%s=%s",
					c.QuoteWord(k),
					v,
				)
			default:
				onDuplicateStr += fmt.Sprintf(
					"%s=VALUES(%s)",
					c.QuoteWord(k),
					c.QuoteWord(gconv.String(v)),
				)
			}
		}
	} else {
		for _, column := range columns {
						// 如果是SAVE操作，不要自动更新创建时间。 md5:409c9c162d30afae
			if c.IsSoftCreatedFieldName(column) {
				continue
			}
			if len(onDuplicateStr) > 0 {
				onDuplicateStr += ","
			}
			onDuplicateStr += fmt.Sprintf(
				"%s=VALUES(%s)",
				c.QuoteWord(column),
				c.QuoteWord(column),
			)
		}
	}

	return InsertOnDuplicateKeyUpdate + " " + onDuplicateStr, nil
}

// RowsToResult 将底层的 sql.Rows 数据记录类型转换为 Result 类型。 md5:ae9065176ef07b2e
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
		// 列名和类型。 md5:51cafb00c4482aba
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	if len(columnTypes) > 0 {
		if internalData := c.getInternalColumnFromCtx(ctx); internalData != nil {
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
				// 不要在这里使用 `gvar.New(nil)`，因为它会创建一个已初始化的对象，
				// 这将导致结构体转换问题。
				// md5:4df778f025fefd53
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
				// 常见的基本内置类型。 md5:4c57bcc430188806
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
		// 其他复杂类型，特别是自定义类型。 md5:5d9bae215068a0c1
	return c.db.ConvertValueForLocal(ctx, columnType.DatabaseTypeName(), value)
}
