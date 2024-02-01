
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
// DriverWrapperDB is a DB wrapper for extending features with embedded DB.
<原文结束>

# <翻译开始>
// DriverWrapperDB 是一个数据库（DB）包装器，用于通过嵌入式数据库扩展功能。
# <翻译结束>


<原文开始>
// Open creates and returns an underlying sql.DB object for pgsql.
// https://pkg.go.dev/github.com/lib/pq
<原文结束>

# <翻译开始>
// Open创建并返回一个用于pgsql的底层sql.DB对象。
// 参考文档：https://pkg.go.dev/github.com/lib/pq
# <翻译结束>


<原文开始>
// Tables retrieves and returns the tables of current schema.
// It's mainly used in cli tool chain for automatically generating the models.
<原文结束>

# <翻译开始>
// Tables 获取并返回当前模式的表。
// 它主要用于cli工具链中，用于自动生成模型。
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
// TableFields 获取并返回当前模式下指定表的字段信息。
//
// 参数 `link` 是可选的，如果给出 nil，则会自动获取一个原始的 SQL 连接作为其链接以执行必要的 SQL 查询。
//
// 注意，它返回一个包含字段名及其对应字段信息的映射。由于映射是无序的，TableField 结构体有一个 "Index" 字段用于标记其在所有字段中的顺序。
//
// 为了提高性能，该函数使用了缓存特性，缓存有效期直到进程重启才会过期。
# <翻译结束>


<原文开始>
// DoInsert inserts or updates data forF given table.
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
// DoInsert 对给定表插入或更新数据。
// 该函数通常用于自定义接口定义，您无需手动调用它。
// 参数`data`的类型可以是 map/gmap/struct/*struct/[]map/[]struct 等。
// 示例：
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
// 参数 `option` 的取值如下：
// InsertOptionDefault：仅插入，如果数据中存在唯一/主键，则返回错误；
// InsertOptionReplace：如果数据中存在唯一/主键，先从表中删除并插入新的记录；
// InsertOptionSave：如果数据中存在唯一/主键，则更新记录，否则插入新记录；
// InsertOptionIgnore：如果数据中存在唯一/主键，则忽略插入操作。
# <翻译结束>


<原文开始>
// Convert data type before commit it to underlying db driver.
<原文结束>

# <翻译开始>
// 在提交给底层数据库驱动之前，转换数据类型。
# <翻译结束>

