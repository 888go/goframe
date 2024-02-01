
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
// Transaction wraps the transaction logic using function `f`.
// It rollbacks the transaction and returns the error from function `f` if
// it returns non-nil error. It commits the transaction and returns nil if
// function `f` returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function `f`
// as it is automatically handled by this function.
<原文结束>

# <翻译开始>
// Transaction 通过函数 `f` 包装事务逻辑。如果函数 `f` 返回非空错误，它将回滚事务并返回该错误。若函数 `f` 返回空（nil）错误，它将提交事务并返回空。
// 注意：在函数 `f` 中不应手动调用 Commit 或 Rollback 方法处理事务，因为这些操作在此函数中已自动完成。
# <翻译结束>

