// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/888go/goframe/internal/intlog"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/reflection"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
)

// X更新 为模型执行"UPDATE ... "语句。
// 
// 如果提供了可选参数 `dataAndWhere`，则 dataAndWhere[0] 是更新的数据字段，dataAndWhere[1:] 被视为 WHERE 条件字段。同时参考 Model.Data 和 Model.Where 函数。
// md5:06a16ce16f9da0c0
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
		return nil, gerror.X创建错误码(gcode.CodeMissingParameter, "updating table with empty data")
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
			if fieldNameUpdate != "" && !gstr.X是否包含(updates, fieldNameUpdate) {
				updates += fmt.Sprintf(`,%s=?`, fieldNameUpdate)
				conditionArgs = append([]interface{}{dataValue}, conditionArgs...)
			}
		}
		updateData = updates
	}
	newData, 错误 := m.filterDataForInsertOrUpdate(updateData)
	if 错误 != nil {
		return nil, 错误
	}

	if !gstr.X是否包含并忽略大小写(conditionStr, " WHERE ") {
		intlog.Printf(
			ctx,
			`sql condition string "%s" has no WHERE for UPDATE operation, fieldNameUpdate: %s`,
			conditionStr, fieldNameUpdate,
		)
		return nil, gerror.X创建错误码(
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

// X更新并取影响行数 执行更新语句并返回受影响的行数。 md5:2f9b42bc238e70a4
func (m *Model) X更新并取影响行数(数据或条件 ...interface{}) (影响行数 int64, 错误 error) {
	result, 错误 := m.X更新(数据或条件...)
	if 错误 != nil {
		return 0, 错误
	}
	return result.RowsAffected()
}

// X更新增量 函数通过给定的数量增加某列的值。
// 参数 `amount` 可以是浮点数或整数类型。
// md5:31e7e26d28456940
func (m *Model) X更新增量(字段名称 string, 增量值 interface{}) (sql.Result, error) {
	return m.getModel().X设置数据(字段名称, &Counter{
		X字段名称: 字段名称,
		X增减值: gconv.X取小数64位(增量值),
	}).X更新()
}

// X更新减量 函数通过给定的数量减小某一列的值。
// 参数 `amount` 可以是浮点数或整数类型。
// md5:e9b9ca17fcd1d042
func (m *Model) X更新减量(字段名称 string, 减量值 interface{}) (sql.Result, error) {
	return m.getModel().X设置数据(字段名称, &Counter{
		X字段名称: 字段名称,
		X增减值: -gconv.X取小数64位(减量值),
	}).X更新()
}
