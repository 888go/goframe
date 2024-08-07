// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

import (
	"database/sql"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	gstr "github.com/888go/goframe/text/gstr"
)

// X删除 执行针对模型的 "DELETE FROM ... " 语句。
// 可选参数 `where` 与 Model.Where 函数的参数相同，
// 请参考 Model.Where 查看详细用法。
// md5:efc496574e0829d8
func (m *Model) X删除(条件 ...interface{}) (结果 sql.Result, 错误 error) {
	var ctx = m.X取上下文对象()
	if len(条件) > 0 {
		return m.X条件(条件[0], 条件[1:]...).X删除()
	}
	defer func() {
		if 错误 == nil {
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
	if !gstr.X是否包含并忽略大小写(conditionStr, " WHERE ") || (fieldNameDelete != "" && !gstr.X是否包含并忽略大小写(conditionStr, " AND ")) {
		intlog.Printf(
			ctx,
			`sql condition string "%s" has no WHERE for DELETE operation, fieldNameDelete: %s`,
			conditionStr, fieldNameDelete,
		)
		return nil, gerror.X创建错误码(
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
