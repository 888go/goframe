// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gdb

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/gogf/gf/v2/internal/intlog"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/reflection"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// Update 为模型执行"UPDATE ... "语句。
//
// 如果提供了可选参数 `dataAndWhere`，则 dataAndWhere[0] 是更新的数据字段，dataAndWhere[1:] 被视为 WHERE 条件字段。同时参考 Model.Data 和 Model.Where 函数。 md5:06a16ce16f9da0c0
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
		stm                                           = m.softTimeMaintainer()
		updateData                                    = m.data
		reflectInfo                                   = reflection.OriginTypeAndKind(updateData)
		conditionWhere, conditionExtra, conditionArgs = m.formatCondition(ctx, false, false)
		conditionStr                                  = conditionWhere + conditionExtra
		fieldNameUpdate, fieldTypeUpdate              = stm.GetFieldNameAndTypeForUpdate(
			ctx, "", m.tablesInit,
		)
	)
	if m.unscoped {
		fieldNameUpdate = ""
	}

	switch reflectInfo.OriginKind {
	case reflect.Map, reflect.Struct:
		var dataMap = anyValueToMapBeforeToRecord(m.data)
		// 自动更新记录的更新时间。 md5:cf60b195a97c5bfa
		if fieldNameUpdate != "" {
			dataValue := stm.GetValueByFieldTypeForCreateOrUpdate(ctx, fieldTypeUpdate, false)
			dataMap[fieldNameUpdate] = dataValue
		}
		updateData = dataMap

	default:
		updates := gconv.String(m.data)
		// 自动更新记录的更新时间。 md5:cf60b195a97c5bfa
		if fieldNameUpdate != "" {
			dataValue := stm.GetValueByFieldTypeForCreateOrUpdate(ctx, fieldTypeUpdate, false)
			if fieldNameUpdate != "" && !gstr.Contains(updates, fieldNameUpdate) {
				updates += fmt.Sprintf(`,%s=?`, fieldNameUpdate)
				conditionArgs = append([]interface{}{dataValue}, conditionArgs...)
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

// UpdateAndGetAffected 执行更新语句并返回受影响的行数。 md5:2f9b42bc238e70a4
func (m *Model) UpdateAndGetAffected(dataAndWhere ...interface{}) (affected int64, err error) {
	result, err := m.Update(dataAndWhere...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// Increment 函数通过给定的数量增加某列的值。
// 参数 `amount` 可以是浮点数或整数类型。 md5:31e7e26d28456940
func (m *Model) Increment(column string, amount interface{}) (sql.Result, error) {
	return m.getModel().Data(column, &Counter{
		Field: column,
		Value: gconv.Float64(amount),
	}).Update()
}

// Decrement 函数通过给定的数量减小某一列的值。
// 参数 `amount` 可以是浮点数或整数类型。 md5:e9b9ca17fcd1d042
func (m *Model) Decrement(column string, amount interface{}) (sql.Result, error) {
	return m.getModel().Data(column, &Counter{
		Field: column,
		Value: -gconv.Float64(amount),
	}).Update()
}
