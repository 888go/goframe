// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

import (
	"context"
	"database/sql"
	"fmt"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
)

type (
	HookFuncSelect func(ctx context.Context, in *HookSelectInput) (result Result, err error)
	HookFuncInsert func(ctx context.Context, in *HookInsertInput) (result sql.Result, err error)
	HookFuncUpdate func(ctx context.Context, in *HookUpdateInput) (result sql.Result, err error)
	HookFuncDelete func(ctx context.Context, in *HookDeleteInput) (result sql.Result, err error)
)

// HookHandler 管理 Model 支持的所有钩子函数。
type HookHandler struct {
	Select HookFuncSelect
	Insert HookFuncInsert
	Update HookFuncUpdate
	Delete HookFuncDelete
}

// internalParamHook 管理 hook 操作的所有内部参数。
// `internal` 显然意味着你无法在本包外部访问这些参数。
type internalParamHook struct {
	link               Link      // Connection 对象来自第三方 SQL 驱动程序。
	handlerCalled      bool      // 简单标记用于自定义处理程序被调用的情况，以防递归调用。
	removedWhere       bool      // 移除了带有`WHERE`前缀的条件字符串的标记
	originalTableName  *泛型类.Var // 原始表名。
	originalSchemaName *泛型类.Var // 原始模式名称。
}

type internalParamHookSelect struct {
	internalParamHook
	handler HookFuncSelect
}

type internalParamHookInsert struct {
	internalParamHook
	handler HookFuncInsert
}

type internalParamHookUpdate struct {
	internalParamHook
	handler HookFuncUpdate
}

type internalParamHookDelete struct {
	internalParamHook
	handler HookFuncDelete
}

// HookSelectInput 保存了 select 钩子操作的参数。
// 注意，此特性也会对 COUNT 语句进行钩子处理，
// 而这通常对于上层业务钩子处理器来说并不有趣（即可能不需要处理）。
type HookSelectInput struct {
	internalParamHookSelect
	Model  *Model        // 当前操作模型
	Table  string        // 将要使用的表名。更新此属性以更改目标表名。
	Schema string        // 要使用的架构名称。更新此属性以更改目标架构名称。
	Sql    string        // 需要执行提交的SQL字符串。
	Args   []interface{} // sql的参数
}

// HookInsertInput 用于存储插入钩子操作的参数。
type HookInsertInput struct {
	internalParamHookInsert
	Model  *Model         // 当前操作模型
	Table  string         // 将要使用的表名。更新此属性以更改目标表名。
	Schema string         // 要使用的架构名称。更新此属性以更改目标架构名称。
	Data   Map数组           // 待插入/保存到表中的数据记录列表
	Option DoInsertOption // 数据插入时的额外选项。
}

// HookUpdateInput 用于保存更新钩子操作的参数。
type HookUpdateInput struct {
	internalParamHookUpdate
	Model     *Model        // 当前操作模型
	Table     string        // 将要使用的表名。更新此属性以更改目标表名。
	Schema    string        // 要使用的架构名称。更新此属性以更改目标架构名称。
	Data      interface{}   // Data 的类型可以是：map[string]interface{}/string。你可以对 `Data` 使用类型断言。
	Condition string        // 更新时的条件字符串。
	Args      []interface{} // 用于SQL占位符的参数。
}

// HookDeleteInput 用于持有删除钩子操作的参数。
type HookDeleteInput struct {
	internalParamHookDelete
	Model     *Model        // 当前操作模型
	Table     string        // 将要使用的表名。更新此属性以更改目标表名。
	Schema    string        // 要使用的架构名称。更新此属性以更改目标架构名称。
	Condition string        // 删除操作的条件字符串
	Args      []interface{} // 用于SQL占位符的参数。
}

const (
	whereKeyInCondition = " WHERE "
)

// IsTransaction 检查并返回当前操作是否在事务中进行。
func (h *internalParamHook) X是否为事务() bool {
	return h.link.IsTransaction()
}

// Next调用下一个钩子处理器。
func (h *HookSelectInput) Next(上下文 context.Context) (行记录数组 Result, 错误 error) {
	if h.originalTableName.X是否为Nil() {
		h.originalTableName = 泛型类.X创建(h.Table)
	}
	if h.originalSchemaName.X是否为Nil() {
		h.originalSchemaName = 泛型类.X创建(h.Schema)
	}
	// 自定义钩子处理器调用。
	if h.handler != nil && !h.handlerCalled {
		h.handlerCalled = true
		return h.handler(上下文, h)
	}
	var toBeCommittedSql = h.Sql
	// Table change.
	if h.Table != h.originalTableName.String() {
		toBeCommittedSql, 错误 = 正则类.ReplaceStringFuncMatch(
			`(?i) FROM ([\S]+)`,
			toBeCommittedSql,
			func(match []string) string {
				charL, charR := h.Model.db.X底层取数据库安全字符()
				return fmt.Sprintf(` FROM %s%s%s`, charL, h.Table, charR)
			},
		)
	}
	// Schema change.
	if h.Schema != "" && h.Schema != h.originalSchemaName.String() {
		h.link, 错误 = h.Model.db.X取Core对象().X底层SlaveLink(h.Schema)
		if 错误 != nil {
			return
		}
	}
	return h.Model.db.X底层查询(上下文, h.link, toBeCommittedSql, h.Args...)
}

// Next调用下一个钩子处理器。
func (h *HookInsertInput) Next(上下文 context.Context) (行记录数组 sql.Result, 错误 error) {
	if h.originalTableName.X是否为Nil() {
		h.originalTableName = 泛型类.X创建(h.Table)
	}
	if h.originalSchemaName.X是否为Nil() {
		h.originalSchemaName = 泛型类.X创建(h.Schema)
	}

	if h.handler != nil && !h.handlerCalled {
		h.handlerCalled = true
		return h.handler(上下文, h)
	}

	// Schema change.
	if h.Schema != "" && h.Schema != h.originalSchemaName.String() {
		h.link, 错误 = h.Model.db.X取Core对象().X底层MasterLink(h.Schema)
		if 错误 != nil {
			return
		}
	}
	return h.Model.db.X底层插入(上下文, h.link, h.Table, h.Data, h.Option)
}

// Next调用下一个钩子处理器。
func (h *HookUpdateInput) Next(上下文 context.Context) (行记录数组 sql.Result, 错误 error) {
	if h.originalTableName.X是否为Nil() {
		h.originalTableName = 泛型类.X创建(h.Table)
	}
	if h.originalSchemaName.X是否为Nil() {
		h.originalSchemaName = 泛型类.X创建(h.Schema)
	}

	if h.handler != nil && !h.handlerCalled {
		h.handlerCalled = true
		if 文本类.X开头判断(h.Condition, whereKeyInCondition) {
			h.removedWhere = true
			h.Condition = 文本类.X过滤首字符(h.Condition, whereKeyInCondition)
		}
		return h.handler(上下文, h)
	}
	if h.removedWhere {
		h.Condition = whereKeyInCondition + h.Condition
	}
	// Schema change.
	if h.Schema != "" && h.Schema != h.originalSchemaName.String() {
		h.link, 错误 = h.Model.db.X取Core对象().X底层MasterLink(h.Schema)
		if 错误 != nil {
			return
		}
	}
	return h.Model.db.X底层更新(上下文, h.link, h.Table, h.Data, h.Condition, h.Args...)
}

// Next调用下一个钩子处理器。
func (h *HookDeleteInput) Next(ctx context.Context) (result sql.Result, err error) {
	if h.originalTableName.X是否为Nil() {
		h.originalTableName = 泛型类.X创建(h.Table)
	}
	if h.originalSchemaName.X是否为Nil() {
		h.originalSchemaName = 泛型类.X创建(h.Schema)
	}

	if h.handler != nil && !h.handlerCalled {
		h.handlerCalled = true
		if 文本类.X开头判断(h.Condition, whereKeyInCondition) {
			h.removedWhere = true
			h.Condition = 文本类.X过滤首字符(h.Condition, whereKeyInCondition)
		}
		return h.handler(ctx, h)
	}
	if h.removedWhere {
		h.Condition = whereKeyInCondition + h.Condition
	}
	// Schema change.
	if h.Schema != "" && h.Schema != h.originalSchemaName.String() {
		h.link, err = h.Model.db.X取Core对象().X底层MasterLink(h.Schema)
		if err != nil {
			return
		}
	}
	return h.Model.db.X底层删除(ctx, h.link, h.Table, h.Condition, h.Args...)
}

// Hook 设置当前模型的钩子函数。
func (m *Model) Hook(hook HookHandler) *Model {
	model := m.getModel()
	model.hookHandler = hook
	return model
}
