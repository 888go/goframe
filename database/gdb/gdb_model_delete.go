// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gdb

import (
	"database/sql"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/text/gstr"
)

// Delete 执行针对模型的 "DELETE FROM ... " 语句。
// 可选参数 `where` 与 Model.Where 函数的参数相同，
// 请参考 Model.Where 查看详细用法。
// md5:efc496574e0829d8
func (m *Model) Delete(where ...interface{}) (result sql.Result, err error) {
	var ctx = m.GetCtx()
	if len(where) > 0 {
		return m.Where(where[0], where[1:]...).Delete()
	}
	defer func() {
		if err == nil {
			m.checkAndRemoveSelectCache(ctx)
		}
	}()
	var (
		conditionWhere, conditionExtra, conditionArgs = m.formatCondition(ctx, false, false)
		conditionStr                                  = conditionWhere + conditionExtra
		fieldNameDelete, fieldTypeDelete              = m.softTimeMaintainer().GetFieldNameAndTypeForDelete(
			ctx, "", m.tablesInit,
		)
	)
	if m.unscoped {
		fieldNameDelete = ""
	}
	if !gstr.ContainsI(conditionStr, " WHERE ") || (fieldNameDelete != "" && !gstr.ContainsI(conditionStr, " AND ")) {
		intlog.Printf(
			ctx,
			`sql condition string "%s" has no WHERE for DELETE operation, fieldNameDelete: %s`,
			conditionStr, fieldNameDelete,
		)
		return nil, gerror.NewCode(
			gcode.CodeMissingParameter,
			"there should be WHERE condition statement for DELETE operation",
		)
	}

	// Soft deleting.
	if fieldNameDelete != "" {
		dataHolder, dataValue := m.softTimeMaintainer().GetDataByFieldNameAndTypeForDelete(
			ctx, "", fieldNameDelete, fieldTypeDelete,
		)
		in := &HookUpdateInput{
			internalParamHookUpdate: internalParamHookUpdate{
				internalParamHook: internalParamHook{
					link: m.getLink(true),
				},
				handler: m.hookHandler.Update,
			},
			Model:     m,
			Table:     m.tables,
			Data:      dataHolder,
			Condition: conditionStr,
			Args:      append([]interface{}{dataValue}, conditionArgs...),
		}
		return in.Next(ctx)
	}

	in := &HookDeleteInput{
		internalParamHookDelete: internalParamHookDelete{
			internalParamHook: internalParamHook{
				link: m.getLink(true),
			},
			handler: m.hookHandler.Delete,
		},
		Model:     m,
		Table:     m.tables,
		Condition: conditionStr,
		Args:      conditionArgs,
	}
	return in.Next(ctx)
}
