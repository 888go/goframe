
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with gm file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者所有（https://goframe.org）。保留所有权利。
//
// 本源代码形式受MIT许可证条款的约束。如果gm文件中未附带MIT许可证的副本，
// 您可以从https://github.com/gogf/gf获取。
// md5:1d281c30cdc3423b
# <翻译结束>


<原文开始>
			// sample : close db conn
			// i.(DBConn).Conn.Close()
<原文结束>

# <翻译开始>
			// 示例：关闭数据库连接
			// 调用 i.(DBConn) 的 Conn 关闭方法
			// md5:1207f4943d8a98dc
# <翻译结束>


<原文开始>
// modify this conn's limit
<原文结束>

# <翻译开始>
// 修改这个连接的限制. md5:fbc2b791ac0ae7a0
# <翻译结束>


<原文开始>
	// example : do same db operation
	// conn.(*DBConn).Conn.QueryContext(context.Background(), "select * from user")
<原文结束>

# <翻译开始>
	// 示例：执行相同的数据库操作
	// 使用conn指向的*DBConn的Conn方法，以context.Background()为上下文，执行SQL查询"select * from user"
	// md5:92af4813b4267108
# <翻译结束>


<原文开始>
	// May Output:
	// Close The Pool
<原文结束>

# <翻译开始>
	// May Output:
	// Close The Pool
# <翻译结束>

