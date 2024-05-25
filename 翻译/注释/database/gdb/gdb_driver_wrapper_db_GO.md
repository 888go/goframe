
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
// DriverWrapperDB is a DB wrapper for extending features with embedded DB.
<原文结束>

# <翻译开始>
// DriverWrapperDB是一个DB包装器，用于通过嵌入式DB扩展功能。 md5:a926644143c69c76
# <翻译结束>


<原文开始>
// Open creates and returns an underlying sql.DB object for pgsql.
// https://pkg.go.dev/github.com/lib/pq
<原文结束>

# <翻译开始>
// Open 创建并返回一个用于pgsql的底层sql.DB对象。
// 参考链接：https://pkg.go.dev/github.com/lib/pq
// md5:9889bcb899248a2b
# <翻译结束>


<原文开始>
// Tables retrieves and returns the tables of current schema.
// It's mainly used in cli tool chain for automatically generating the models.
<原文结束>

# <翻译开始>
// Tables 获取并返回当前模式下的表格列表。
//主要用于命令行工具链，用于自动生成模型。
// md5:bce161ba95454bf5
# <翻译结束>


<原文开始>
// TableFields retrieves and returns the fields' information of specified table of current
// schema.
//
// The parameter `link` is optional, if given nil it automatically retrieves a raw sql connection
// as its link to proceed necessary sql query.
//
// Note that it returns a map containing the field name and its corresponding fields.
// As a map is unsorted, the TableField struct has an "Index" field marks its sequence in
// the fields.
//
// It's using cache feature to enhance the performance, which is never expired util the
// process restarts.
<原文结束>

# <翻译开始>
// TableFields 获取并返回当前模式指定表的字段信息。
// 
// 参数 `link` 是可选的，如果为 nil，则自动获取一个原始 SQL 连接，用于执行必要的 SQL 查询。
// 
// 它返回一个包含字段名及其对应字段的映射。由于映射是无序的，TableField 结构体有一个 "Index" 字段，标记其在字段中的顺序。
// 
// 该方法使用缓存功能来提高性能，直到进程重启，缓存永不过期。
// md5:c844572d5210b35e
# <翻译结束>


<原文开始>
// prefix:group@schema#table
<原文结束>

# <翻译开始>
// 前缀:组@模式#表. md5:b22e67d9da02a91b
# <翻译结束>


<原文开始>
// DoInsert inserts or updates data for given table.
// This function is usually used for custom interface definition, you do not need call it manually.
// The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc.
// Eg:
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// The parameter `option` values are as follows:
// InsertOptionDefault:  just insert, if there's unique/primary key in the data, it returns error;
// InsertOptionReplace: if there's unique/primary key in the data, it deletes it from table and inserts a new one;
// InsertOptionSave:    if there's unique/primary key in the data, it updates it or else inserts a new one;
// InsertOptionIgnore:  if there's unique/primary key in the data, it ignores the inserting;
<原文结束>

# <翻译开始>
// DoInsert 用于插入或更新给定表的数据。
// 此函数通常用于自定义接口定义，您无需手动调用。
// 参数 `data` 可以为 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 例如：
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// 参数 `option` 的值如下：
// InsertOptionDefault：仅插入，如果数据中包含唯一键或主键，则返回错误；
// InsertOptionReplace：如果数据中包含唯一键或主键，先从表中删除原有记录，再插入新记录；
// InsertOptionSave：如果数据中包含唯一键或主键，进行更新，否则插入新记录；
// InsertOptionIgnore：如果数据中包含唯一键或主键，忽略插入操作。
// md5:9fab32fdc41df179
# <翻译结束>


<原文开始>
// Convert data type before commit it to underlying db driver.
<原文结束>

# <翻译开始>
// 在将数据类型提交给底层数据库驱动程序之前进行转换。 md5:58b56ae1ed22196f
# <翻译结束>

