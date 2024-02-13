// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

import (
	"database/sql"
	"fmt"
	"reflect"
	
	"github.com/888go/goframe/internal/intlog"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/reflection"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
)

// Update 执行针对模型的 "UPDATE ... " 语句。
//
// 如果提供可选参数 `dataAndWhere`，则 dataAndWhere[0] 被视为更新的数据字段，
// 而 dataAndWhere[1:] 被视为 WHERE 条件字段。
// 另请参阅 Model.Data 和 Model.Where 函数。
func (m *Model) X更新(数据或条件 ...interface{}) (结果 sql.Result, 错误 error) {
	var ctx = m.X取上下文对象()
	if len(数据或条件) > 0 {
		if len(数据或条件) > 2 {
			return m.X设置数据(数据或条件[0]).X条件(数据或条件[1], 数据或条件[2:]...).X更新()
		} else if len(数据或条件) == 2 {
			return m.X设置数据(数据或条件[0]).X条件(数据或条件[1]).X更新()
		} else {
			return m.X设置数据(数据或条件[0]).X更新()
		}
	}
	defer func() {
		if 错误 == nil {
			m.checkAndRemoveSelectCache(ctx)
		}
	}()
	if m.data == nil {
		return nil, 错误类.X创建错误码(错误码类.CodeMissingParameter, "updating table with empty data")
	}
	var (
		updateData                                    = m.data
		reflectInfo                                   = reflection.OriginTypeAndKind(updateData)
		fieldNameUpdate                               = m.getSoftFieldNameUpdated("", m.tablesInit)
		conditionWhere, conditionExtra, conditionArgs = m.formatCondition(ctx, false, false)
		conditionStr                                  = conditionWhere + conditionExtra
	)
	if m.unscoped {
		fieldNameUpdate = ""
	}

	switch reflectInfo.OriginKind {
	case reflect.Map, reflect.Struct:
		var dataMap = anyValueToMapBeforeToRecord(m.data)
		// 自动更新记录的更新时间。
		if fieldNameUpdate != "" {
			dataMap[fieldNameUpdate] = 时间类.X创建并按当前时间()
		}
		updateData = dataMap

	default:
		updates := 转换类.String(m.data)
		// 自动更新记录的更新时间。
		if fieldNameUpdate != "" {
			if fieldNameUpdate != "" && !文本类.X是否包含(updates, fieldNameUpdate) {
				updates += fmt.Sprintf(`,%s=?`, fieldNameUpdate)
				conditionArgs = append([]interface{}{时间类.X创建并按当前时间()}, conditionArgs...)
			}
		}
		updateData = updates
	}
	newData, 错误 := m.filterDataForInsertOrUpdate(updateData)
	if 错误 != nil {
		return nil, 错误
	}

	if !文本类.X是否包含并忽略大小写(conditionStr, " WHERE ") {
		intlog.Printf(
			ctx,
			`sql condition string "%s" has no WHERE for UPDATE operation, fieldNameUpdate: %s`,
			conditionStr, fieldNameUpdate,
		)
		return nil, 错误类.X创建错误码(
			错误码类.CodeMissingParameter,
			"there should be WHERE condition statement for UPDATE operation",
		)
	}

	in := &HookUpdateInput{
		internalParamHookUpdate: internalParamHookUpdate{
			internalParamHook: internalParamHook{
				link: m.getLink(true),
			},
			handler: m.hookHandler.Update,
		},
		Model:     m,
		Table:     m.tables,
		Data:      newData,
		Condition: conditionStr,
		Args:      m.mergeArguments(conditionArgs),
	}
	return in.Next(ctx)
}

// UpdateAndGetAffected 执行更新语句并返回受影响的行数。
func (m *Model) X更新并取影响行数(数据或条件 ...interface{}) (影响行数 int64, 错误 error) {
	result, 错误 := m.X更新(数据或条件...)
	if 错误 != nil {
		return 0, 错误
	}
	return result.RowsAffected()
}

// Increment 通过给定的量增加列的值。
// 参数 `amount` 可以是浮点数或整数类型。
func (m *Model) X更新增量(字段名称 string, 增量值 interface{}) (sql.Result, error) {
	return m.getModel().X设置数据(字段名称, &X增减{
		X字段名称: 字段名称,
		X增减值: 转换类.X取小数64位(增量值),
	}).X更新()
}

// Decrement 函数用于对某一列的值减去指定的数量。
// 参数 `amount` 可以是浮点数或整数类型。
func (m *Model) X更新减量(字段名称 string, 减量值 interface{}) (sql.Result, error) {
	return m.getModel().X设置数据(字段名称, &X增减{
		X字段名称: 字段名称,
		X增减值: -转换类.X取小数64位(减量值),
	}).X更新()
}
