// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package pgsql

import (
	"context"
	"database/sql"
	"strings"

	gdb "github.com/888go/goframe/database/gdb"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
)

// X底层原生SQL执行 通过给定的链接对象将 sql 字符串及其参数提交到底层驱动，并返回执行结果。
// md5:947bd2b83e751e10
func (d *Driver) X底层原生SQL执行(ctx context.Context, link gdb.Link, sql string, args ...interface{}) (result sql.Result, err error) {
	var (
		isUseCoreDoExec bool   = false // 检查是否需要使用默认方法. md5:5c2234665d43c8e2
		primaryKey      string = ""
		pkField         gdb.TableField
	)

	// Transaction checks.
	if link == nil {
		if tx := gdb.X事务从上下文取对象(ctx, d.X取配置组名称()); tx != nil {
						// 首先，从上下文中检查并获取交易链接。 md5:9ac4c60388fa960d
			link = tx
		} else if link, err = d.X底层MasterLink(); err != nil {
						// 否则，它将从主节点创建一个。 md5:4bd14606783b43fc
			return nil, err
		}
	} else if !link.IsTransaction() {
				// 如果当前链接不是事务链接，它会检查并从上下文中检索事务。 md5:e3c484ab061699a1
		if tx := gdb.X事务从上下文取对象(ctx, d.X取配置组名称()); tx != nil {
			link = tx
		}
	}

		// 检查是否为带有主键的插入操作。 md5:f7e52e42ae148c58
	if value := ctx.Value(internalPrimaryKeyInCtx); value != nil {
		var ok bool
		pkField, ok = value.(gdb.TableField)
		if !ok {
			isUseCoreDoExec = true
		}
	} else {
		isUseCoreDoExec = true
	}

		// 检查是否为插入操作。 md5:b3cb8d582bc267c9
	if !isUseCoreDoExec && pkField.X名称 != "" && strings.Contains(sql, "INSERT INTO") {
		primaryKey = pkField.X名称
		sql += " RETURNING " + primaryKey
	} else {
		// use default DoExec
		return d.Core.X底层原生SQL执行(ctx, link, sql, args...)
	}

			// 只有使用主键进行插入操作时，才能执行以下代码. md5:90f3f7f5e35bf09e

	if d.X取当前节点配置().ExecTimeout > 0 {
		var cancelFunc context.CancelFunc
		ctx, cancelFunc = context.WithTimeout(ctx, d.X取当前节点配置().ExecTimeout)
		defer cancelFunc()
	}

	// Sql filtering.
	sql, args = d.X格式化Sql(sql, args)
	sql, args, err = d.X底层DoFilter(ctx, link, sql, args)
	if err != nil {
		return nil, err
	}

	// Link execution.
	var out gdb.DoCommitOutput
	out, err = d.X底层DoCommit(ctx, gdb.DoCommitInput{
		Link:          link,
		Sql:           sql,
		Args:          args,
		Stmt:          nil,
		Type:          gdb.SqlTypeQueryContext,
		IsTransaction: link.IsTransaction(),
	})

	if err != nil {
		return nil, err
	}
	affected := len(out.X行记录切片)
	if affected > 0 {
		if !strings.Contains(pkField.X类型, "int") {
			return Result{
				affected:     int64(affected),
				lastInsertId: 0,
				lastInsertIdError: gerror.X创建错误码并格式化(
					gcode.CodeNotSupported,
					"LastInsertId is not supported by primary key type: %s", pkField.X类型),
			}, nil
		}

		if out.X行记录切片[affected-1][primaryKey] != nil {
			lastInsertId := out.X行记录切片[affected-1][primaryKey].X取整数64位()
			return Result{
				affected:     int64(affected),
				lastInsertId: lastInsertId,
			}, nil
		}
	}

	return Result{}, nil
}
