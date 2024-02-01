
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// HookHandler manages all supported hook functions for Model.
<原文结束>

# <翻译开始>
// HookHandler 管理 Model 支持的所有钩子函数。
# <翻译结束>


<原文开始>
// internalParamHook manages all internal parameters for hook operations.
// The `internal` obviously means you cannot access these parameters outside this package.
<原文结束>

# <翻译开始>
// internalParamHook 管理 hook 操作的所有内部参数。
// `internal` 显然意味着你无法在本包外部访问这些参数。
# <翻译结束>


<原文开始>
// Connection object from third party sql driver.
<原文结束>

# <翻译开始>
// Connection 对象来自第三方 SQL 驱动程序。
# <翻译结束>


<原文开始>
// Simple mark for custom handler called, in case of recursive calling.
<原文结束>

# <翻译开始>
// 简单标记用于自定义处理程序被调用的情况，以防递归调用。
# <翻译结束>












<原文开始>
// HookSelectInput holds the parameters for select hook operation.
// Note that, COUNT statement will also be hooked by this feature,
// which is usually not be interesting for upper business hook handler.
<原文结束>

# <翻译开始>
// HookSelectInput 保存了 select 钩子操作的参数。
// 注意，此特性也会对 COUNT 语句进行钩子处理，
// 而这通常对于上层业务钩子处理器来说并不有趣（即可能不需要处理）。
# <翻译结束>







<原文开始>
// The table name that to be used. Update this attribute to change target table name.
<原文结束>

# <翻译开始>
// 将要使用的表名。更新此属性以更改目标表名。
# <翻译结束>


<原文开始>
// The schema name that to be used. Update this attribute to change target schema name.
<原文结束>

# <翻译开始>
// 要使用的架构名称。更新此属性以更改目标架构名称。
# <翻译结束>


<原文开始>
// The sql string that to be committed.
<原文结束>

# <翻译开始>
// 需要执行提交的SQL字符串。
# <翻译结束>







<原文开始>
// HookInsertInput holds the parameters for insert hook operation.
<原文结束>

# <翻译开始>
// HookInsertInput 用于存储插入钩子操作的参数。
# <翻译结束>


<原文开始>
// The data records list to be inserted/saved into table.
<原文结束>

# <翻译开始>
// 待插入/保存到表中的数据记录列表
# <翻译结束>


<原文开始>
// The extra option for data inserting.
<原文结束>

# <翻译开始>
// 数据插入时的额外选项。
# <翻译结束>


<原文开始>
// HookUpdateInput holds the parameters for update hook operation.
<原文结束>

# <翻译开始>
// HookUpdateInput 用于保存更新钩子操作的参数。
# <翻译结束>


<原文开始>
// The where condition string for updating.
<原文结束>

# <翻译开始>
// 更新时的条件字符串。
# <翻译结束>


<原文开始>
// The arguments for sql place-holders.
<原文结束>

# <翻译开始>
// 用于SQL占位符的参数。
# <翻译结束>


<原文开始>
// HookDeleteInput holds the parameters for delete hook operation.
<原文结束>

# <翻译开始>
// HookDeleteInput 用于持有删除钩子操作的参数。
# <翻译结束>


<原文开始>
// The where condition string for deleting.
<原文结束>

# <翻译开始>
// 删除操作的条件字符串
# <翻译结束>


<原文开始>
// IsTransaction checks and returns whether current operation is during transaction.
<原文结束>

# <翻译开始>
// IsTransaction 检查并返回当前操作是否在事务中进行。
# <翻译结束>


<原文开始>
// Next calls the next hook handler.
<原文结束>

# <翻译开始>
// Next调用下一个钩子处理器。
# <翻译结束>







<原文开始>
// Hook sets the hook functions for current model.
<原文结束>

# <翻译开始>
// Hook 设置当前模型的钩子函数。
# <翻译结束>


<原文开始>
// Removed mark for condition string that was removed `WHERE` prefix.
<原文结束>

# <翻译开始>
// 移除了带有`WHERE`前缀的条件字符串的标记
# <翻译结束>


<原文开始>
// Data can be type of: map[string]interface{}/string. You can use type assertion on `Data`.
<原文结束>

# <翻译开始>
// Data 的类型可以是：map[string]interface{}/string。你可以对 `Data` 使用类型断言。
# <翻译结束>


<原文开始>
// The original table name.
<原文结束>

# <翻译开始>
// 原始表名。
# <翻译结束>


<原文开始>
// The original schema name.
<原文结束>

# <翻译开始>
// 原始模式名称。
# <翻译结束>


<原文开始>
// Current operation Model.
<原文结束>

# <翻译开始>
// 当前操作模型
# <翻译结束>


<原文开始>
// The arguments of sql.
<原文结束>

# <翻译开始>
// sql的参数
# <翻译结束>


<原文开始>
// Custom hook handler call.
<原文结束>

# <翻译开始>
// 自定义钩子处理器调用。
# <翻译结束>

