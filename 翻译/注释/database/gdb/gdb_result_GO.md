
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
// SqlResult is execution result for sql operations.
// It also supports batch operation result for rowsAffected.
<原文结束>

# <翻译开始>
// SqlResult是SQL操作的执行结果。它还支持行影响的批量操作结果。
// md5:c89c5ab24627c936
# <翻译结束>


<原文开始>
// MustGetAffected returns the affected rows count, if any error occurs, it panics.
<原文结束>

# <翻译开始>
// MustGetAffected 返回受影响的行数，如果发生任何错误，则会引发恐慌。 md5:be151685a0da2b44
# <翻译结束>


<原文开始>
// MustGetInsertId returns the last insert id, if any error occurs, it panics.
<原文结束>

# <翻译开始>
// MustGetInsertId 返回最后的插入ID，如果发生任何错误，它将引发恐慌。 md5:bd23d169a4cb6738
# <翻译结束>


<原文开始>
// RowsAffected returns the number of rows affected by an
// update, insert, or delete. Not every database or database
// driver may support this.
// Also, See sql.Result.
<原文结束>

# <翻译开始>
// RowsAffected 返回更新、插入或删除操作影响的行数。并非所有数据库或数据库驱动程序都支持此功能。
// 参见 sql.Result。
// md5:f41c8ccbf7344301
# <翻译结束>


<原文开始>
// LastInsertId returns the integer generated by the database
// in response to a command. Typically, this will be from an
// "auto increment" column when inserting a new row. Not all
// databases support this feature, and the syntax of such
// statements varies.
// Also, See sql.Result.
<原文结束>

# <翻译开始>
// LastInsertId返回数据库对命令的响应生成的整数。通常，这将是在插入新行时来自“自动递增”列的。并非所有数据库都支持此功能，且此类语句的语法各不相同。
// 参见sql.Result。
// md5:7236c1ac3f4fc094
# <翻译结束>

