
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// HookHandler manages all supported hook functions for Model.
<原文结束>

# <翻译开始>
// HookHandler 管理Model支持的所有钩子函数。 md5:bc5db27f3bf00d12
# <翻译结束>


<原文开始>
// internalParamHook manages all internal parameters for hook operations.
// The `internal` obviously means you cannot access these parameters outside this package.
<原文结束>

# <翻译开始>
// internalParamHook 管理所有用于钩子操作的内部参数。
// `internal` 显然意味着您无法在此包之外访问这些参数。
// md5:25a7b0a478a19a4b
# <翻译结束>


<原文开始>
// Connection object from third party sql driver.
<原文结束>

# <翻译开始>
// 来自第三方sql驱动的连接对象。 md5:8c0e18a3b7135850
# <翻译结束>


<原文开始>
// Simple mark for custom handler called, in case of recursive calling.
<原文结束>

# <翻译开始>
// 用于自定义处理器调用的简单标记，如果存在递归调用。 md5:8a70de5e368bfa75
# <翻译结束>


<原文开始>
// Removed mark for condition string that was removed `WHERE` prefix.
<原文结束>

# <翻译开始>
// 删除了已移除`WHERE`前缀的条件字符串标记。 md5:65b20530f0b91cf9
# <翻译结束>


<原文开始>
// The original table name.
<原文结束>

# <翻译开始>
// 原始表名。 md5:4a73dda3a3e91183
# <翻译结束>


<原文开始>
// The original schema name.
<原文结束>

# <翻译开始>
// 原始的模式名称。 md5:bea72de299f2aa4d
# <翻译结束>


<原文开始>
// HookSelectInput holds the parameters for select hook operation.
// Note that, COUNT statement will also be hooked by this feature,
// which is usually not be interesting for upper business hook handler.
<原文结束>

# <翻译开始>
// HookSelectInput 存储选择操作的参数。
// 注意，COUNT 语句也会被此功能捕获，这通常对上层业务钩子处理程序不感兴趣。
// md5:c5f22bccaae80481
# <翻译结束>


<原文开始>
// Current operation Model.
<原文结束>

# <翻译开始>
// 当前操作模型。 md5:d9c5abcf43d4a0c5
# <翻译结束>


<原文开始>
// The table name that to be used. Update this attribute to change target table name.
<原文结束>

# <翻译开始>
// 将要使用的表名。更新此属性以更改目标表名。 md5:b5d4582f7fa65327
# <翻译结束>


<原文开始>
// The schema name that to be used. Update this attribute to change target schema name.
<原文结束>

# <翻译开始>
// 将要使用的模式名称。更新此属性以更改目标模式名称。 md5:40385c83e27c8a07
# <翻译结束>


<原文开始>
// The sql string that to be committed.
<原文结束>

# <翻译开始>
// 将要提交的SQL字符串。 md5:7c6c74bdd4ed9bb2
# <翻译结束>


<原文开始>
// HookInsertInput holds the parameters for insert hook operation.
<原文结束>

# <翻译开始>
// HookInsertInput 插入钩子操作的参数。 md5:76f9069cc685c571
# <翻译结束>


<原文开始>
// The data records list to be inserted/saved into table.
<原文结束>

# <翻译开始>
// 要插入/保存到表中的数据记录列表。 md5:af6867e8ee9b8dd5
# <翻译结束>


<原文开始>
// The extra option for data inserting.
<原文结束>

# <翻译开始>
// 用于数据插入的额外选项。 md5:ffac0ff130d3b693
# <翻译结束>


<原文开始>
// HookUpdateInput holds the parameters for update hook operation.
<原文结束>

# <翻译开始>
// HookUpdateInput 表示更新钩子操作的参数。 md5:a9d35fc8f42cd434
# <翻译结束>


<原文开始>
// Data can be type of: map[string]interface{}/string. You can use type assertion on `Data`.
<原文结束>

# <翻译开始>
// `Data` 可以是类型：map[string]interface{} 或 string。你可以对 `Data` 进行类型断言。 md5:f92fddf82f17883a
# <翻译结束>


<原文开始>
// The where condition string for updating.
<原文结束>

# <翻译开始>
// 用于更新的条件字符串。 md5:4bcf07b70ed87d5a
# <翻译结束>


<原文开始>
// The arguments for sql place-holders.
<原文结束>

# <翻译开始>
// sql占位符的参数。 md5:aed81f2b97f42d86
# <翻译结束>


<原文开始>
// HookDeleteInput holds the parameters for delete hook operation.
<原文结束>

# <翻译开始>
// HookDeleteInput包含删除钩子操作的参数。 md5:f7d586e1f75c0a3e
# <翻译结束>


<原文开始>
// The where condition string for deleting.
<原文结束>

# <翻译开始>
// 删除操作的WHERE条件字符串。 md5:63d65a2af6b3c2b9
# <翻译结束>


<原文开始>
// IsTransaction checks and returns whether current operation is during transaction.
<原文结束>

# <翻译开始>
// IsTransaction 检查并返回当前操作是否处于事务中。 md5:689b943de611f296
# <翻译结束>


<原文开始>
// Next calls the next hook handler.
<原文结束>

# <翻译开始>
// Next 调用下一个钩子处理器。 md5:7348deede95e47b0
# <翻译结束>


<原文开始>
// Custom hook handler call.
<原文结束>

# <翻译开始>
	// 自定义钩子处理器调用。 md5:edb1c6e5a718f78e
# <翻译结束>


<原文开始>
// Hook sets the hook functions for current model.
<原文结束>

# <翻译开始>
// Hook 设置当前模型的钩子函数。 md5:a324f56d597fd873
# <翻译结束>

