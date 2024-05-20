
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
// DoFilter handles the sql before posts it to database.
<原文结束>

# <翻译开始>
// DoFilter 在将SQL提交到数据库之前进行处理。. md5:e56455a7432db765
# <翻译结束>


<原文开始>
// Convert placeholder char '?' to string "$x".
<原文结束>

# <翻译开始>
// 将占位符字符'?'转换为字符串 "$x"。. md5:a1e39f745b49128a
# <翻译结束>


<原文开始>
// Only SQL generated through the framework is processed.
<原文结束>

# <翻译开始>
// 只有通过框架生成的SQL才会被处理。. md5:fe793f1543ad8481
# <翻译结束>


<原文开始>
// replace STD SQL to Clickhouse SQL grammar
<原文结束>

# <翻译开始>
// 将标准SQL替换为ClickHouse SQL语法. md5:eb710cc2dce6880f
# <翻译结束>


<原文开始>
// Only delete/ UPDATE statements require filter
<原文结束>

# <翻译开始>
// 只有删除/更新语句需要过滤条件. md5:a43c6e48c79fc525
# <翻译结束>


<原文开始>
		// MySQL eg: UPDATE table_name SET field1=new-value1, field2=new-value2 [WHERE Clause]
		// Clickhouse eg: ALTER TABLE [db.]table UPDATE column1 = expr1 [, ...] WHERE filter_expr
<原文结束>

# <翻译开始>
// MySQL 示例：UPDATE table_name SET field1=new-value1, field2=new-value2 [WHERE 条件]
// Clickhouse 示例：ALTER TABLE [db.]table UPDATE column1 = expr1[, ...] WHERE filter_expr
// 
// 这段代码是针对两种数据库系统的更新语句的注释。在MySQL中，`UPDATE` 用于更新表中的数据，设置指定字段的新值，并可选地使用 `WHERE` 子句来限制更新的行。在Clickhouse中，`ALTER TABLE` 用于更新表中的列，将列的值设置为表达式（expr1），并且需要一个过滤表达式（filter_expr）来确定哪些行会被更新。其中 `[db.]table` 表示可以包含数据库名的表名。
// md5:d201a8d0c4df9319
# <翻译结束>


<原文开始>
		// MySQL eg: DELETE FROM table_name [WHERE Clause]
		// Clickhouse eg: ALTER TABLE [db.]table [ON CLUSTER cluster] DELETE WHERE filter_expr
<原文结束>

# <翻译开始>
// MySQL 示例：DELETE FROM 表名 [WHERE 子句]
// Clickhouse 示例：ALTER TABLE [db.]表名 [ON CLUSTER 集群名称] DELETE WHERE 过滤表达式
// md5:b83ff5334ed70fb6
# <翻译结束>

