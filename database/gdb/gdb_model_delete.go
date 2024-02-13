// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

import (
	"database/sql"
	"fmt"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/text/gstr"
	
	"github.com/888go/goframe/os/gtime"
)

// Delete 执行 "DELETE FROM ..." 语句用于该模型。
// 可选参数 `where` 与 Model.Where 函数的参数相同，
// 请参阅 Model.Where。
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
		fieldNameDelete                               = m.getSoftFieldNameDeleted("", m.tablesInit)
		conditionWhere, conditionExtra, conditionArgs = m.formatCondition(ctx, false, false)
		conditionStr                                  = conditionWhere + conditionExtra
	)
	if m.unscoped {
		fieldNameDelete = ""
	}
	if !文本类.X是否包含并忽略大小写(conditionStr, " WHERE ") || (fieldNameDelete != "" && !文本类.X是否包含并忽略大小写(conditionStr, " AND ")) {
		intlog.Printf(
			ctx,
			`sql condition string "%s" has no WHERE for DELETE operation, fieldNameDelete: %s`,
			conditionStr, fieldNameDelete,
		)
		return nil, 错误类.X创建错误码(
			错误码类.CodeMissingParameter,
			"there should be WHERE condition statement for DELETE operation",
		)
	}

	// Soft deleting.
	if fieldNameDelete != "" {
		in := &HookUpdateInput{
			internalParamHookUpdate: internalParamHookUpdate{
				internalParamHook: internalParamHook{
					link: m.getLink(true),
				},
				handler: m.hookHandler.Update,
			},
			Model:     m,
			Table:     m.tables,
			Data:      fmt.Sprintf(`%s=?`, m.db.X取Core对象().X底层QuoteString(fieldNameDelete)),
			Condition: conditionStr,
			Args:      append([]interface{}{时间类.X创建并按当前时间()}, conditionArgs...),
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
