
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
// internalCtxData stores data in ctx for internal usage purpose.
<原文结束>

# <翻译开始>
// internalCtxData 为内部使用目的，在 ctx 中存储数据。 md5:95073898cc1f4772
# <翻译结束>


<原文开始>
// Used configuration node in current operation.
<原文结束>

# <翻译开始>
// 当前操作中使用的配置节点。 md5:85f106587581bb38
# <翻译结束>


<原文开始>
// column stores column data in ctx for internal usage purpose.
<原文结束>

# <翻译开始>
// column 用于内部目的，在ctx中存储列数据。 md5:12a8a80132bf8ae7
# <翻译结束>


<原文开始>
	// The first column in result response from database server.
	// This attribute is used for Value/Count selection statement purpose,
	// which is to avoid HOOK handler that might modify the result columns
	// that can confuse the Value/Count selection statement logic.
<原文结束>

# <翻译开始>
// 来自数据库服务器的响应结果中的第一列。
// 此属性用于值/计数选择语句的目的，以避免可能修改结果列的HOOK处理器，这可能会混淆值/计数选择语句的逻辑。
// md5:c678f20e25487136
# <翻译结束>


<原文开始>
	// `ignoreResultKeyInCtx` is a mark for some db drivers that do not support `RowsAffected` function,
	// for example: `clickhouse`. The `clickhouse` does not support fetching insert/update results,
	// but returns errors when execute `RowsAffected`. It here ignores the calling of `RowsAffected`
	// to avoid triggering errors, rather than ignoring errors after they are triggered.
<原文结束>

# <翻译开始>
// `ignoreResultKeyInCtx` 是为了一些不支持 `RowsAffected` 函数的数据库驱动（例如：`clickhouse`）设置的标记。`clickhouse` 不支持获取插入/更新的结果，但在执行 `RowsAffected` 时会返回错误。在这里，我们忽略对 `RowsAffected` 的调用，以避免触发错误，而不是在错误发生后忽略它们。
// md5:4a7864c37326a119
# <翻译结束>


<原文开始>
// If the internal data is already injected, it does nothing.
<原文结束>

# <翻译开始>
// 如果内部数据已经被注入，则不做任何操作。 md5:ae258e1c66cb106a
# <翻译结束>

