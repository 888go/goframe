// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

import (
	"context"
	"database/sql"
	"fmt"

	gvar "github.com/888go/goframe/container/gvar"
	gregex "github.com/888go/goframe/text/gregex"
	gstr "github.com/888go/goframe/text/gstr"
)

type (
	HookFuncSelect func(ctx context.Context, in *HookSelectInput) (result Result, err error)
	HookFuncInsert func(ctx context.Context, in *HookInsertInput) (result sql.Result, err error)
	HookFuncUpdate func(ctx context.Context, in *HookUpdateInput) (result sql.Result, err error)
	HookFuncDelete func(ctx context.Context, in *HookDeleteInput) (result sql.Result, err error)
)

// HookHandler 管理Model支持的所有钩子函数。 md5:bc5db27f3bf00d12
type HookHandler struct {
	Select HookFuncSelect
	Insert HookFuncInsert
	Update HookFuncUpdate
	Delete HookFuncDelete
}

// internalParamHook 管理所有用于钩子操作的内部参数。
// `internal` 显然意味着您无法在此包之外访问这些参数。
// md5:25a7b0a478a19a4b
type internalParamHook struct {
	link               Link      // 来自第三方sql驱动的连接对象。 md5:8c0e18a3b7135850
	handlerCalled      bool      // 用于自定义处理器调用的简单标记，如果存在递归调用。 md5:8a70de5e368bfa75
	removedWhere       bool      // 删除了已移除`WHERE`前缀的条件字符串标记。 md5:65b20530f0b91cf9
	originalTableName  *gvar.Var // 原始表名。 md5:4a73dda3a3e91183
	originalSchemaName *gvar.Var // 原始的模式名称。 md5:bea72de299f2aa4d
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

// HookSelectInput 存储选择操作的参数。
// 注意，COUNT 语句也会被此功能捕获，这通常对上层业务钩子处理程序不感兴趣。
// md5:c5f22bccaae80481
type HookSelectInput struct {
	internalParamHookSelect
	Model  *Model        // 当前操作模型。 md5:d9c5abcf43d4a0c5
	Table  string        // 将要使用的表名。更新此属性以更改目标表名。 md5:b5d4582f7fa65327
	Schema string        // 将要使用的模式名称。更新此属性以更改目标模式名称。 md5:40385c83e27c8a07
	Sql    string        // 将要提交的SQL字符串。 md5:7c6c74bdd4ed9bb2
	Args   []interface{} // The arguments of sql.
}

// HookInsertInput 插入钩子操作的参数。 md5:76f9069cc685c571
type HookInsertInput struct {
	internalParamHookInsert
	Model  *Model         // 当前操作模型。 md5:d9c5abcf43d4a0c5
	Table  string         // 将要使用的表名。更新此属性以更改目标表名。 md5:b5d4582f7fa65327
	Schema string         // 将要使用的模式名称。更新此属性以更改目标模式名称。 md5:40385c83e27c8a07
	Data   Map切片           // 要插入/保存到表中的数据记录列表。 md5:af6867e8ee9b8dd5
	Option DoInsertOption // 用于数据插入的额外选项。 md5:ffac0ff130d3b693
}

// HookUpdateInput 表示更新钩子操作的参数。 md5:a9d35fc8f42cd434
type HookUpdateInput struct {
	internalParamHookUpdate
	Model     *Model        // 当前操作模型。 md5:d9c5abcf43d4a0c5
	Table     string        // 将要使用的表名。更新此属性以更改目标表名。 md5:b5d4582f7fa65327
	Schema    string        // 将要使用的模式名称。更新此属性以更改目标模式名称。 md5:40385c83e27c8a07
	Data      interface{}   // `Data` 可以是类型：map[string]interface{} 或 string。你可以对 `Data` 进行类型断言。 md5:f92fddf82f17883a
	Condition string        // 用于更新的条件字符串。 md5:4bcf07b70ed87d5a
	Args      []interface{} // sql占位符的参数。 md5:aed81f2b97f42d86
}

// HookDeleteInput包含删除钩子操作的参数。 md5:f7d586e1f75c0a3e
type HookDeleteInput struct {
	internalParamHookDelete
	Model     *Model        // 当前操作模型。 md5:d9c5abcf43d4a0c5
	Table     string        // 将要使用的表名。更新此属性以更改目标表名。 md5:b5d4582f7fa65327
	Schema    string        // 将要使用的模式名称。更新此属性以更改目标模式名称。 md5:40385c83e27c8a07
	Condition string        // 删除操作的WHERE条件字符串。 md5:63d65a2af6b3c2b9
	Args      []interface{} // sql占位符的参数。 md5:aed81f2b97f42d86
}

const (
	whereKeyInCondition = " WHERE "
)

// X是否为事务 检查并返回当前操作是否处于事务中。 md5:689b943de611f296
func (h *internalParamHook) X是否为事务() bool {
	return h.link.IsTransaction()
}

// Next 调用下一个钩子处理器。 md5:7348deede95e47b0
func (h *HookSelectInput) Next(上下文 context.Context) (行记录切片 Result, 错误 error) {
	if h.originalTableName.X是否为Nil() {
		h.originalTableName = gvar.X创建(h.Table)
	}
	if h.originalSchemaName.X是否为Nil() {
		h.originalSchemaName = gvar.X创建(h.Schema)
	}
		// 自定义钩子处理器调用。 md5:edb1c6e5a718f78e
	if h.handler != nil && !h.handlerCalled {
		h.handlerCalled = true
		return h.handler(上下文, h)
	}
	var toBeCommittedSql = h.Sql
	// Table change.
	if h.Table != h.originalTableName.String() {
		toBeCommittedSql, 错误 = gregex.ReplaceStringFuncMatch(
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

// Next 调用下一个钩子处理器。 md5:7348deede95e47b0
func (h *HookInsertInput) Next(上下文 context.Context) (行记录切片 sql.Result, 错误 error) {
	if h.originalTableName.X是否为Nil() {
		h.originalTableName = gvar.X创建(h.Table)
	}
	if h.originalSchemaName.X是否为Nil() {
		h.originalSchemaName = gvar.X创建(h.Schema)
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

// Next 调用下一个钩子处理器。 md5:7348deede95e47b0
func (h *HookUpdateInput) Next(上下文 context.Context) (行记录切片 sql.Result, 错误 error) {
	if h.originalTableName.X是否为Nil() {
		h.originalTableName = gvar.X创建(h.Table)
	}
	if h.originalSchemaName.X是否为Nil() {
		h.originalSchemaName = gvar.X创建(h.Schema)
	}

	if h.handler != nil && !h.handlerCalled {
		h.handlerCalled = true
		if gstr.X开头判断(h.Condition, whereKeyInCondition) {
			h.removedWhere = true
			h.Condition = gstr.X过滤首字符(h.Condition, whereKeyInCondition)
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

// Next 调用下一个钩子处理器。 md5:7348deede95e47b0
func (h *HookDeleteInput) Next(ctx context.Context) (result sql.Result, err error) {
	if h.originalTableName.X是否为Nil() {
		h.originalTableName = gvar.X创建(h.Table)
	}
	if h.originalSchemaName.X是否为Nil() {
		h.originalSchemaName = gvar.X创建(h.Schema)
	}

	if h.handler != nil && !h.handlerCalled {
		h.handlerCalled = true
		if gstr.X开头判断(h.Condition, whereKeyInCondition) {
			h.removedWhere = true
			h.Condition = gstr.X过滤首字符(h.Condition, whereKeyInCondition)
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

// Hook 设置当前模型的钩子函数。 md5:a324f56d597fd873
func (m *Model) Hook(hook HookHandler) *Model {
	model := m.getModel()
	model.hookHandler = hook
	return model
}
