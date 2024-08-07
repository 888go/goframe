// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

package db类

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	gconv "github.com/888go/goframe/util/gconv"

	"github.com/888go/goframe"
	gvar "github.com/888go/goframe/container/gvar"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	gtime "github.com/888go/goframe/os/gtime"
	guid "github.com/888go/goframe/util/guid"
)

// X原生SQL查询 向底层驱动提交一个查询SQL并返回执行结果。
// 它最常用于数据查询。
// md5:06bbbfc29aa3894b
func (c *Core) X原生SQL查询(上下文 context.Context, sql string, 参数 ...interface{}) (结果 Result, 错误 error) {
	return c.db.X底层原生SQL查询(上下文, nil, sql, 参数...)
}

// X底层原生SQL查询 通过给定的链接对象将SQL字符串及其参数提交给底层驱动，并返回执行结果。
// md5:af7bdcd1a2074bc0
func (c *Core) X底层原生SQL查询(上下文 context.Context, 链接 Link, sql string, 参数 ...interface{}) (结果 Result, 错误 error) {
	// Transaction checks.
	if 链接 == nil {
		if tx := X事务从上下文取对象(上下文, c.db.X取配置组名称()); tx != nil {
						// 首先，从上下文中检查并获取交易链接。 md5:9ac4c60388fa960d
			链接 = &txLink{tx.X底层取事务对象()}
		} else if 链接, 错误 = c.X底层SlaveLink(); 错误 != nil {
						// 否则，它将从主节点创建一个。 md5:4bd14606783b43fc
			return nil, 错误
		}
	} else if !链接.IsTransaction() {
				// 如果当前链接不是事务链接，它会检查并从上下文中检索事务。 md5:e3c484ab061699a1
		if tx := X事务从上下文取对象(上下文, c.db.X取配置组名称()); tx != nil {
			链接 = &txLink{tx.X底层取事务对象()}
		}
	}

	if c.db.X取当前节点配置().QueryTimeout > 0 {
		上下文, _ = context.WithTimeout(上下文, c.db.X取当前节点配置().QueryTimeout)
	}

	// Sql filtering.
	sql, 参数 = c.X格式化Sql(sql, 参数)
	sql, 参数, 错误 = c.db.X底层DoFilter(上下文, 链接, sql, 参数)
	if 错误 != nil {
		return nil, 错误
	}
		// SQL格式化和获取。 md5:815f530302ec8a7e
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
	var out DoCommitOutput
	out, 错误 = c.db.X底层DoCommit(上下文, DoCommitInput{
		Link:          链接,
		Sql:           sql,
		Args:          参数,
		Stmt:          nil,
		Type:          SqlTypeQueryContext,
		IsTransaction: 链接.IsTransaction(),
	})
	return out.X行记录切片, 错误
}

// X原生SQL执行 将一个查询 SQL 执行到底层驱动并返回执行结果。它最常用于数据插入和更新。
// md5:6f9ddc85964b9797
func (c *Core) X原生SQL执行(上下文 context.Context, sql string, 参数 ...interface{}) (结果 sql.Result, 错误 error) {
	return c.db.X底层原生SQL执行(上下文, nil, sql, 参数...)
}

// X底层原生SQL执行 通过给定的链接对象将 sql 字符串及其参数提交到底层驱动，并返回执行结果。
// md5:947bd2b83e751e10
func (c *Core) X底层原生SQL执行(上下文 context.Context, 链接 Link, sql string, 参数 ...interface{}) (结果 sql.Result, 错误 error) {
	// Transaction checks.
	if 链接 == nil {
		if tx := X事务从上下文取对象(上下文, c.db.X取配置组名称()); tx != nil {
						// 首先，从上下文中检查并获取交易链接。 md5:9ac4c60388fa960d
			链接 = &txLink{tx.X底层取事务对象()}
		} else if 链接, 错误 = c.X底层MasterLink(); 错误 != nil {
						// 否则，它将从主节点创建一个。 md5:4bd14606783b43fc
			return nil, 错误
		}
	} else if !链接.IsTransaction() {
				// 如果当前链接不是事务链接，它会检查并从上下文中检索事务。 md5:e3c484ab061699a1
		if tx := X事务从上下文取对象(上下文, c.db.X取配置组名称()); tx != nil {
			链接 = &txLink{tx.X底层取事务对象()}
		}
	}

	if c.db.X取当前节点配置().ExecTimeout > 0 {
		var cancelFunc context.CancelFunc
		上下文, cancelFunc = context.WithTimeout(上下文, c.db.X取当前节点配置().ExecTimeout)
		defer cancelFunc()
	}

	// SQL filtering.
	sql, 参数 = c.X格式化Sql(sql, 参数)
	sql, 参数, 错误 = c.db.X底层DoFilter(上下文, 链接, sql, 参数)
	if 错误 != nil {
		return nil, 错误
	}
		// SQL格式化和获取。 md5:815f530302ec8a7e
	if v := 上下文.Value(ctxKeyCatchSQL); v != nil {
		var (
			manager      = v.(*CatchSQLManager)
			formattedSql = X格式化Sql(sql, 参数)
		)
		manager.SQLArray.Append别名(formattedSql)
		if !manager.DoCommit && 上下文.Value(ctxKeyInternalProducedSQL) == nil {
			return new(SqlResult), nil
		}
	}
	// Link execution.
	var out DoCommitOutput
	out, 错误 = c.db.X底层DoCommit(上下文, DoCommitInput{
		Link:          链接,
		Sql:           sql,
		Args:          参数,
		Stmt:          nil,
		Type:          SqlTypeExecContext,
		IsTransaction: 链接.IsTransaction(),
	})
	return out.Result, 错误
}

// X底层DoFilter 是一个钩子函数，它在 SQL 语句及其参数提交给底层驱动之前进行过滤。
// 参数 `link` 指定当前数据库连接的操作对象。在 SQL 语句 `sql` 及其参数 `args` 被提交给驱动之前，您可以根据需要随意修改它们。
// md5:41118fbc4e6c5562
func (c *Core) X底层DoFilter(ctx context.Context, link Link, sql string, args []interface{}) (newSql string, newArgs []interface{}, err error) {
	return sql, args, nil
}

// X底层DoCommit 将当前SQL和参数提交给底层SQL驱动程序。 md5:7cf9b1f6f4d9d2cb
func (c *Core) X底层DoCommit(ctx context.Context, in DoCommitInput) (out DoCommitOutput, err error) {
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
		timestampMilli1      = gtime.X取时间戳毫秒()
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
				transactionId: guid.X生成(),
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
			sqlResult = new(SqlResult)
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
			sqlResult = new(SqlResult)
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
		panic(gerror.X创建错误码并格式化(gcode.CodeInvalidParameter, `invalid SqlType "%s"`, in.Type))
	}
	// Result handling.
	switch {
	case sqlResult != nil && !c.X底层_GetIgnoreResultFromCtx(ctx):
		rowsAffected, err = sqlResult.RowsAffected()
		out.Result = sqlResult

	case sqlRows != nil:
		out.X行记录切片, err = c.X原生sql记录到行记录切片对象(ctx, sqlRows)
		rowsAffected = int64(len(out.X行记录切片))

	case sqlStmt != nil:
		out.X参数预处理 = &Stmt{
			Stmt: sqlStmt,
			core: c,
			link: in.Link,
			sql:  in.Sql,
		}
	}
	var (
		timestampMilli2 = gtime.X取时间戳毫秒()
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
		err = gerror.X多层错误码(
			gcode.CodeDbOperationError,
			err,
			X格式化Sql(in.Sql, in.Args),
		)
	}
	return out, err
}

// X原生sql取参数预处理对象 准备一个预编译语句，供后续查询或执行使用。
// 可以从返回的语句对象并发运行多个查询或执行。
// 当不再需要语句时，调用者必须调用 statement 的 Close 方法。
// 
// 参数 `execOnMaster` 指定是否在主节点上执行 SQL，如果配置了主从复制，则在从节点上执行。
// md5:639eebcae369b0a2
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

// X底层原生sql参数预处理对象 会调用给定链接对象上的prepare函数，并返回语句对象。 md5:bae03ede256987bd
func (c *Core) X底层原生sql参数预处理对象(上下文 context.Context, 链接 Link, sql string) (参数预处理 *Stmt, 错误 error) {
	// Transaction checks.
	if 链接 == nil {
		if tx := X事务从上下文取对象(上下文, c.db.X取配置组名称()); tx != nil {
						// 首先，从上下文中检查并获取交易链接。 md5:9ac4c60388fa960d
			链接 = &txLink{tx.X底层取事务对象()}
		} else {
						// 否则，它将从主节点创建一个。 md5:4bd14606783b43fc
			var err error
			if 链接, err = c.X底层MasterLink(); err != nil {
				return nil, err
			}
		}
	} else if !链接.IsTransaction() {
				// 如果当前链接不是事务链接，它会检查并从上下文中检索事务。 md5:e3c484ab061699a1
		if tx := X事务从上下文取对象(上下文, c.db.X取配置组名称()); tx != nil {
			链接 = &txLink{tx.X底层取事务对象()}
		}
	}

	if c.db.X取当前节点配置().PrepareTimeout > 0 {
				// 不要在预处理语句中使用取消函数。 md5:5e529fe5094c7942
		上下文, _ = context.WithTimeout(上下文, c.db.X取当前节点配置().PrepareTimeout)
	}

	// Link execution.
	var out DoCommitOutput
	out, 错误 = c.db.X底层DoCommit(上下文, DoCommitInput{
		Link:          链接,
		Sql:           sql,
		Type:          SqlTypePrepareContext,
		IsTransaction: 链接.IsTransaction(),
	})
	return out.X参数预处理, 错误
}

// FormatUpsert 格式化并返回用于 UPSERT 语句的 SQL 子句部分。
// 在默认实现中，此函数执行类似 MySQL 的 UPSERT 语句：
// `INSERT INTO ... ON DUPLICATE KEY UPDATE x=VALUES(z),m=VALUES(y)...`
// md5:c1c6d7b14661682b
func (c *Core) FormatUpsert(columns []string, list Map切片, option DoInsertOption) (string, error) {
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
					c.X底层QuoteWord(k),
					v,
				)
			default:
				onDuplicateStr += fmt.Sprintf(
					"%s=VALUES(%s)",
					c.X底层QuoteWord(k),
					c.X底层QuoteWord(gconv.String(v)),
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
				c.X底层QuoteWord(column),
				c.X底层QuoteWord(column),
			)
		}
	}

	return InsertOnDuplicateKeyUpdate + " " + onDuplicateStr, nil
}

// X原生sql记录到行记录切片对象 将底层的 sql.Rows 数据记录类型转换为 Result 类型。 md5:ae9065176ef07b2e
func (c *Core) X原生sql记录到行记录切片对象(上下文 context.Context, 底层数据记录 *sql.Rows) (Result, error) {
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
		// 列名和类型。 md5:51cafb00c4482aba
	columnTypes, err := 底层数据记录.ColumnTypes()
	if err != nil {
		return nil, err
	}

	if len(columnTypes) > 0 {
		if internalData := c.getInternalColumnFromCtx(上下文); internalData != nil {
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
				// 不要在这里使用 `gvar.New(nil)`，因为它会创建一个已初始化的对象，
				// 这将导致结构体转换问题。
				// md5:4df778f025fefd53
				record[columnTypes[i].Name()] = nil
			} else {
				var convertedValue interface{}
				if convertedValue, err = c.columnValueToLocalValue(上下文, value, columnTypes[i]); err != nil {
					return nil, err
				}
				record[columnTypes[i].Name()] = gvar.X创建(convertedValue)
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
				// 常见的基本内置类型。 md5:4c57bcc430188806
		switch scanType.Kind() {
		case
			reflect.Bool,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			return gconv.X按名称转换(
				gconv.String(value),
				columnType.ScanType().String(),
			), nil
		}
	}
		// 其他复杂类型，特别是自定义类型。 md5:5d9bae215068a0c1
	return c.db.X底层ConvertValueForLocal(ctx, columnType.DatabaseTypeName(), value)
}
