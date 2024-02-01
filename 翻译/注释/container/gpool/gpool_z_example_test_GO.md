
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with gm file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随 gm 文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
# <翻译结束>


<原文开始>
			// sample : close db conn
			// i.(DBConn).Conn.Close()
<原文结束>

# <翻译开始>
// 示例：关闭数据库连接
// i.(DBConn).Conn.Close()
# <翻译结束>


<原文开始>
	// example : do same db operation
	// conn.(*DBConn).Conn.QueryContext(context.Background(), "select * from user")
<原文结束>

# <翻译开始>
// 示例：执行相同的数据库操作
// conn.(*DBConn).Conn.QueryContext(context.Background(), "select * from user")
// 翻译：
// 示例：进行同样的数据库操作
// conn.(*DBConn).Conn.QueryContext 使用 context.Background() 作为上下文，执行 SQL 查询语句："从 user 表中选择所有列"
# <翻译结束>


<原文开始>
// modify this conn's limit
<原文结束>

# <翻译开始>
// 修改此连接的限制
# <翻译结束>


<原文开始>
// wait for pool close
<原文结束>

# <翻译开始>
// 等待连接池关闭
# <翻译结束>

