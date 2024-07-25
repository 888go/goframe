
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
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Transaction wraps the transaction logic using function `f`.
// It rollbacks the transaction and returns the error from function `f` if
// it returns non-nil error. It commits the transaction and returns nil if
// function `f` returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function `f`
// as it is automatically handled by this function.
<原文结束>

# <翻译开始>
// Transaction 包装了使用函数 `f` 执行的事务逻辑。
// 如果函数 `f` 返回非空错误，它将回滚事务并返回该错误。如果函数 `f` 返回 nil，它将提交事务并返回 nil。
//
// 注意，在函数 `f` 中不应手动提交或回滚事务，因为这些操作将由这个函数自动处理。 md5:8906440d4dbbef1f
# <翻译结束>

