// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb
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
func (m *Model) Update(dataAndWhere ...interface{}) (result sql.Result, err error) {
	var ctx = m.GetCtx()
	if len(dataAndWhere) > 0 {
		if len(dataAndWhere) > 2 {
			return m.Data(dataAndWhere[0]).Where(dataAndWhere[1], dataAndWhere[2:]...).Update()
		} else if len(dataAndWhere) == 2 {
			return m.Data(dataAndWhere[0]).Where(dataAndWhere[1]).Update()
		} else {
			return m.Data(dataAndWhere[0]).Update()
		}
	}
	defer func() {
		if err == nil {
			m.checkAndRemoveSelectCache(ctx)
		}
	}()
	if m.data == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "updating table with empty data")
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
			dataMap[fieldNameUpdate] = gtime.Now()
		}
		updateData = dataMap

	default:
		updates := gconv.String(m.data)
		// 自动更新记录的更新时间。
		if fieldNameUpdate != "" {
			if fieldNameUpdate != "" && !gstr.Contains(updates, fieldNameUpdate) {
				updates += fmt.Sprintf(`,%s=?`, fieldNameUpdate)
				conditionArgs = append([]interface{}{gtime.Now()}, conditionArgs...)
			}
		}
		updateData = updates
	}
	newData, err := m.filterDataForInsertOrUpdate(updateData)
	if err != nil {
		return nil, err
	}

	if !gstr.ContainsI(conditionStr, " WHERE ") {
		intlog.Printf(
			ctx,
			`sql condition string "%s" has no WHERE for UPDATE operation, fieldNameUpdate: %s`,
			conditionStr, fieldNameUpdate,
		)
		return nil, gerror.NewCode(
			gcode.CodeMissingParameter,
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
func (m *Model) UpdateAndGetAffected(dataAndWhere ...interface{}) (affected int64, err error) {
	result, err := m.Update(dataAndWhere...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// Increment 通过给定的量增加列的值。
// 参数 `amount` 可以是浮点数或整数类型。
func (m *Model) Increment(column string, amount interface{}) (sql.Result, error) {
	return m.getModel().Data(column, &Counter{
		Field: column,
		Value: gconv.Float64(amount),
	}).Update()
}

// Decrement 函数用于对某一列的值减去指定的数量。
// 参数 `amount` 可以是浮点数或整数类型。
func (m *Model) Decrement(column string, amount interface{}) (sql.Result, error) {
	return m.getModel().Data(column, &Counter{
		Field: column,
		Value: -gconv.Float64(amount),
	}).Update()
}
