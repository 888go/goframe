
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
// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as slice type.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter `where` is the same as the parameter of Model.Where function,
// see Model.Where.
<原文结束>

# <翻译开始>
// All 方法执行针对模型的 "SELECT FROM ..." 语句。
// 它从表中检索记录并以切片类型返回结果。
// 如果根据给定条件未能从表中检索到任何记录，则返回 nil。
//
// 可选参数 `where` 与 Model.Where 函数的参数相同，
// 请参阅 Model.Where。
# <翻译结束>


<原文开始>
// AllAndCount retrieves all records and the total count of records from the model.
// If useFieldForCount is true, it will use the fields specified in the model for counting;
// otherwise, it will use a constant value of 1 for counting.
// It returns the result as a slice of records, the total count of records, and an error if any.
// The where parameter is an optional list of conditions to use when retrieving records.
//
// Example:
//
//	var model Model
//	var result Result
//	var count int
//	where := []interface{}{"name = ?", "John"}
//	result, count, err := model.AllAndCount(true)
//	if err != nil {
//	    // Handle error.
//	}
//	fmt.Println(result, count)
<原文结束>

# <翻译开始>
// AllAndCount 从模型中检索所有记录以及记录的总数。
// 如果 useFieldForCount 设为 true，它将使用模型中指定的字段进行计数；
// 否则，它将以常数值 1 进行计数。
// 它以记录切片的形式返回结果、记录的总数量以及（如有）错误信息。
// where 参数是一个可选的条件列表，在检索记录时使用这些条件。
//
// 示例：
//
//	var model Model   // 假设 Model 是自定义的数据模型
//	var result Result // 假设 Result 是存储查询结果的数据结构
//	var count int
//	where := []interface{}{"name = ?", "John"} // 设置查询条件：name 等于 "John"
//	result, count, err := model.AllAndCount(true) // 调用 AllAndCount 方法并传入参数 useFieldForCount 为 true
//	if err != nil {
//// 处理错误
//	}
//	fmt.Println(result, count) // 输出查询结果和记录总数
# <翻译结束>


<原文开始>
// Clone the model for counting
<原文结束>

# <翻译开始>
// 对模型进行克隆以进行计数
# <翻译结束>


<原文开始>
// If useFieldForCount is false, set the fields to a constant value of 1 for counting
<原文结束>

# <翻译开始>
// 如果useFieldForCount为false，则将字段设置为常数值1用于计数
# <翻译结束>


<原文开始>
// Get the total count of records
<原文结束>

# <翻译开始>
// 获取记录的总数
# <翻译结束>


<原文开始>
// If the total count is 0, there are no records to retrieve, so return early
<原文结束>

# <翻译开始>
// 如果总记录数为0，则表示没有记录需要获取，所以提前返回
# <翻译结束>







<原文开始>
// Chunk iterates the query result with given `size` and `handler` function.
<原文结束>

# <翻译开始>
// Chunk 对查询结果进行迭代，指定 `size`（大小）和 `handler` 函数。
// 根据给定的 `size`，将查询结果分割成多个块，并对每个数据块应用 `handler` 函数进行处理。
# <翻译结束>


<原文开始>
// One retrieves one record from table and returns the result as map type.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter `where` is the same as the parameter of Model.Where function,
// see Model.Where.
<原文结束>

# <翻译开始>
// One 从表中检索一条记录并以map类型返回结果。
// 如果根据给定条件没有从表中检索到任何记录，则返回nil。
//
// 可选参数`where`与Model.Where函数的参数相同，
// 请参阅Model.Where。
# <翻译结束>


<原文开始>
// Array queries and returns data values as slice from database.
// Note that if there are multiple columns in the result, it returns just one column values randomly.
//
// If the optional parameter `fieldsAndWhere` is given, the fieldsAndWhere[0] is the selected fields
// and fieldsAndWhere[1:] is treated as where condition fields.
// Also see Model.Fields and Model.Where functions.
<原文结束>

# <翻译开始>
// 从数据库查询并以切片形式返回数据值。
// 注意，如果结果中有多个列，则随机返回其中一个列的值。
//
// 如果提供可选参数`fieldsAndWhere`，则fieldsAndWhere[0]表示选定的字段，
// 而fieldsAndWhere[1:]被视为where条件字段。
// 同时参阅Model.Fields和Model.Where函数。
# <翻译结束>


<原文开始>
// Struct retrieves one record from table and converts it into given struct.
// The parameter `pointer` should be type of *struct/**struct. If type **struct is given,
// it can create the struct internally during converting.
//
// The optional parameter `where` is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if the given parameter `pointer` pointed to a variable that has
// default value and there's no record retrieved with the given conditions from table.
//
// Example:
// user := new(User)
// err  := db.Model("user").Where("id", 1).Scan(user)
//
// user := (*User)(nil)
// err  := db.Model("user").Where("id", 1).Scan(&user).
<原文结束>

# <翻译开始>
// Struct 从表中获取一条记录并将其转换为给定的结构体。
// 参数 `pointer` 应为 *struct 或 **struct 类型。如果提供的是 **struct 类型，
// 它会在转换过程中自动创建该结构体。
//
// 可选参数 `where` 与 Model.Where 函数的参数相同，详情请参阅 Model.Where。
//
// 注意，如果给定的 `pointer` 指向一个具有默认值的变量，并且根据给定条件未能从表中检索到记录，
// 则返回 sql.ErrNoRows 错误。
//
// 示例：
// user := new(User)
// err  := db.Model("user").Where("id", 1).Scan(user)
//
// user := (*User)(nil)
// err  := db.Model("user").Where("id", 1).Scan(&user)
# <翻译结束>


<原文开始>
// Auto selecting fields by struct attributes.
<原文结束>

# <翻译开始>
// 根据结构体属性自动选择字段
# <翻译结束>


<原文开始>
// Structs retrieves records from table and converts them into given struct slice.
// The parameter `pointer` should be type of *[]struct/*[]*struct. It can create and fill the struct
// slice internally during converting.
//
// The optional parameter `where` is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if the given parameter `pointer` pointed to a variable that has
// default value and there's no record retrieved with the given conditions from table.
//
// Example:
// users := ([]User)(nil)
// err   := db.Model("user").Scan(&users)
//
// users := ([]*User)(nil)
// err   := db.Model("user").Scan(&users).
<原文结束>

# <翻译开始>
// Structs 从表中检索记录并将其转换为给定的结构体切片。
// 参数 `pointer` 应为 *[]struct 或 *[]*struct 类型，它在转换过程中可以内部创建并填充结构体切片。
//
// 可选参数 `where` 与 Model.Where 函数的参数相同，请参阅 Model.Where。
//
// 注意，如果给定的参数 `pointer` 指向一个具有默认值的变量，并且根据给定条件没有从表中检索到任何记录，则会返回 sql.ErrNoRows 错误。
//
// 示例：
// users := ([]User)(nil)
// err   := db.Model("user").Scan(&users)
//
// users := ([]*User)(nil)
// err   := db.Model("user").Scan(&users)
# <翻译结束>


<原文开始>
// Scan automatically calls Struct or Structs function according to the type of parameter `pointer`.
// It calls function doStruct if `pointer` is type of *struct/**struct.
// It calls function doStructs if `pointer` is type of *[]struct/*[]*struct.
//
// The optional parameter `where` is the same as the parameter of Model.Where function,  see Model.Where.
//
// Note that it returns sql.ErrNoRows if the given parameter `pointer` pointed to a variable that has
// default value and there's no record retrieved with the given conditions from table.
//
// Example:
// user := new(User)
// err  := db.Model("user").Where("id", 1).Scan(user)
//
// user := (*User)(nil)
// err  := db.Model("user").Where("id", 1).Scan(&user)
//
// users := ([]User)(nil)
// err   := db.Model("user").Scan(&users)
//
// users := ([]*User)(nil)
// err   := db.Model("user").Scan(&users).
<原文结束>

# <翻译开始>
// Scan 根据参数 `pointer` 的类型自动调用 Struct 或 Structs 函数。
// 若 `pointer` 类型为 *struct/**struct，则调用 doStruct 函数。
// 若 `pointer` 类型为 *[]struct/*[]*struct，则调用 doStructs 函数。
// 可选参数 `where` 与 Model.Where 函数的参数相同，具体可参考 Model.Where。
// 注意：如果给定的参数 `pointer` 指向一个具有默认值的变量，并且根据给定条件未能从表中检索到任何记录时，
// 此函数将返回 sql.ErrNoRows 错误。
// 示例：
// 创建一个新的 User 实例
// user := new(User)
// 使用查询条件从 "user" 表中扫描数据并赋值给 user
// err  := db.Model("user").Where("id", 1).Scan(user)
// 创建一个空指针类型的 User 实例
// user := (*User)(nil)
// 使用查询条件从 "user" 表中扫描数据并赋值给 user
// err  := db.Model("user").Where("id", 1).Scan(&user)
// 创建一个空切片类型的 User 实例
// users := ([]User)(nil)
// 从 "user" 表中扫描数据并将其填充到 users 切片
// err   := db.Model("user").Scan(&users)
// 创建一个空指针切片类型的 User 实例
// users := ([]*User)(nil)
// 从 "user" 表中扫描数据并将其填充到 users 指针切片
// err   := db.Model("user").Scan(&users)
# <翻译结束>


<原文开始>
// ScanAndCount scans a single record or record array that matches the given conditions and counts the total number of records that match those conditions.
// If useFieldForCount is true, it will use the fields specified in the model for counting;
// The pointer parameter is a pointer to a struct that the scanned data will be stored in.
// The pointerCount parameter is a pointer to an integer that will be set to the total number of records that match the given conditions.
// The where parameter is an optional list of conditions to use when retrieving records.
//
// Example:
//
//	var count int
//	user := new(User)
//	err  := db.Model("user").Where("id", 1).ScanAndCount(user,&count,true)
//	fmt.Println(user, count)
//
// Example Join:
//
//	type User struct {
//		Id       int
//		Passport string
//		Name     string
//		Age      int
//	}
//	var users []User
//	var count int
//	db.Model(table).As("u1").
//		LeftJoin(tableName2, "u2", "u2.id=u1.id").
//		Fields("u1.passport,u1.id,u2.name,u2.age").
//		Where("u1.id<2").
//		ScanAndCount(&users, &count, false)
<原文结束>

# <翻译开始>
// ScanAndCount 扫描满足给定条件的单个记录或记录数组，并计算匹配这些条件的记录总数。
// 如果 useFieldForCount 为 true，则会使用模型中指定的字段进行计数；
// pointer 参数是一个指向结构体的指针，扫描的数据将存储在这个结构体中。
// pointerCount 参数是一个指向整数的指针，该整数将被设置为匹配给定条件的总记录数。
// where 参数是可选的检索记录时使用的条件列表。
// 示例：
//
//	var count int
//	user := new(User)
//	err  := db.Model("user").Where("id", 1).ScanAndCount(user, &count, true)
//	fmt.Println(user, count)
// 示例（联接查询）：
//
//	type User struct {
//		Id       int
//		Passport string
//		Name     string
//		Age      int
//	}
//	var users []User
//	var count int
//	db.Model(table).As("u1").
//		LeftJoin(tableName2, "u2", "u2.id=u1.id").
//		Fields("u1.passport,u1.id,u2.name,u2.age").
//		Where("u1.id<2").
//		ScanAndCount(&users, &count, false) // 对于计数不考虑特定字段，因此设为 false
// 此函数用于根据提供的条件执行数据库查询，同时获取满足条件的记录数量。当 `useFieldForCount` 设为 `true` 时，计数会基于模型中定义的某些字段；否则，计数的是所有满足条件的记录条数。查询结果可以填充到传入的结构体实例（单例或切片）中，同时返回满足条件的记录总数。
# <翻译结束>


<原文开始>
// support Fields with *, example: .Fields("a.*, b.name"). Count sql is select count(1) from xxx
<原文结束>

# <翻译开始>
// 支持使用 * 通配符的 Fields，例如：.Fields("a.*, b.name")。Count SQL 为：从 xxx 中选择 count(1)
# <翻译结束>


<原文开始>
// ScanList converts `r` to struct slice which contains other complex struct attributes.
// Note that the parameter `listPointer` should be type of *[]struct/*[]*struct.
//
// See Result.ScanList.
<原文结束>

# <翻译开始>
// ScanList 将 `r` 转换为包含其他复杂结构体属性的结构体切片。
// 注意，参数 `listPointer` 应该是指向 []struct 或 []*struct 类型的指针。
// 参考关联模型: https://goframe.org/pages/viewpage.action?pageId=1114326
//
// 参见 Result.ScanList。
# <翻译结束>







<原文开始>
// Filter fields using temporary created struct using reflect.New.
<原文结束>

# <翻译开始>
// 使用reflect.New创建的临时结构体来过滤字段。
# <翻译结束>


<原文开始>
// Value retrieves a specified record value from table and returns the result as interface type.
// It returns nil if there's no record found with the given conditions from table.
//
// If the optional parameter `fieldsAndWhere` is given, the fieldsAndWhere[0] is the selected fields
// and fieldsAndWhere[1:] is treated as where condition fields.
// Also see Model.Fields and Model.Where functions.
<原文结束>

# <翻译开始>
// Value 从表中检索指定记录的值并以 interface 类型返回结果。
// 如果根据给定条件在表中未找到记录，则返回 nil。
//
// 如果提供了可选参数 `fieldsAndWhere`，则 fieldsAndWhere[0] 表示选择的字段，
// 而 fieldsAndWhere[1:] 将被视为 where 条件字段。
// 请参阅 Model.Fields 和 Model.Where 函数。
# <翻译结束>


<原文开始>
// Count does "SELECT COUNT(x) FROM ..." statement for the model.
// The optional parameter `where` is the same as the parameter of Model.Where function,
// see Model.Where.
<原文结束>

# <翻译开始>
// Count 对模型执行 "SELECT COUNT(x) FROM ..." 语句。
// 可选参数 `where` 与 Model.Where 函数的参数相同，
// 请参阅 Model.Where。
# <翻译结束>


<原文开始>
// CountColumn does "SELECT COUNT(x) FROM ..." statement for the model.
<原文结束>

# <翻译开始>
// CountColumn 对模型执行 "SELECT COUNT(x) FROM ..." 语句。
# <翻译结束>


<原文开始>
// Min does "SELECT MIN(x) FROM ..." statement for the model.
<原文结束>

# <翻译开始>
// Min 为该模型执行“SELECT MIN(x) FROM ...”语句。
# <翻译结束>


<原文开始>
// Max does "SELECT MAX(x) FROM ..." statement for the model.
<原文结束>

# <翻译开始>
// Max 为给定的模型执行“SELECT MAX(x) FROM ...”语句。
# <翻译结束>


<原文开始>
// Avg does "SELECT AVG(x) FROM ..." statement for the model.
<原文结束>

# <翻译开始>
// Avg 对模型执行 "SELECT AVG(x) FROM ..." 语句，计算平均值。
# <翻译结束>


<原文开始>
// Sum does "SELECT SUM(x) FROM ..." statement for the model.
<原文结束>

# <翻译开始>
// Sum 对模型执行 "SELECT SUM(x) FROM ..." 语句，计算求和。
# <翻译结束>


<原文开始>
// Union does "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." statement for the model.
<原文结束>

# <翻译开始>
// Union 为给定的模型执行 "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." 类似的SQL语句查询。
# <翻译结束>


<原文开始>
// UnionAll does "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ..." statement for the model.
<原文结束>

# <翻译开始>
// UnionAll 对模型执行“(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ...”语句。
# <翻译结束>


<原文开始>
// Limit sets the "LIMIT" statement for the model.
// The parameter `limit` can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
<原文结束>

# <翻译开始>
// Limit 为模型设置 "LIMIT" 语句。
// 参数 `limit` 可以是一个或两个数字，如果传入两个数字，
// 则为模型设置 "LIMIT limit[0], limit[1]" 语句，否则设置 "LIMIT limit[0]" 语句。
# <翻译结束>


<原文开始>
// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
<原文结束>

# <翻译开始>
// Offset 设置模型的 "OFFSET" 语句。
// 它只对某些数据库有意义，如 SQLServer、PostgreSQL 等。
# <翻译结束>


<原文开始>
// Distinct forces the query to only return distinct results.
<原文结束>

# <翻译开始>
// Distinct 用于强制查询只返回不重复的结果。
# <翻译结束>


<原文开始>
// Page sets the paging number for the model.
// The parameter `page` is started from 1 for paging.
// Note that, it differs that the Limit function starts from 0 for "LIMIT" statement.
<原文结束>

# <翻译开始>
// Page 为模型设置分页号。
// 参数 `page` 的分页从1开始计数。
// 注意，它与 Limit 函数从0开始为 "LIMIT" 语句设置偏移量有所不同。
# <翻译结束>


<原文开始>
// Having sets the having statement for the model.
// The parameters of this function usage are as the same as function Where.
// See Where.
<原文结束>

# <翻译开始>
// Having 设置模型的 having 子句。
// 该函数的使用参数与 Where 函数相同。
// 参见 Where。
# <翻译结束>


<原文开始>
// doGetAll does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as slice type.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The parameter `limit1` specifies whether limits querying only one record if m.limit is not set.
// The optional parameter `where` is the same as the parameter of Model.Where function,
// see Model.Where.
<原文结束>

# <翻译开始>
// doGetAll 执行针对模型的 "SELECT FROM ..." 语句。
// 它从表中检索记录并以切片类型返回结果。
// 如果根据给定条件没有从表中检索到任何记录，则返回 nil。
//
// 参数 `limit1` 指定在 m.limit 未设置时是否限制仅查询一条记录。
// 可选参数 `where` 与 Model.Where 函数的参数相同，
// 请参阅 Model.Where。
# <翻译结束>


<原文开始>
// doGetAllBySql does the select statement on the database.
<原文结束>

# <翻译开始>
// doGetAllBySql 对数据库执行 select 语句。
# <翻译结束>


<原文开始>
			// DO NOT quote the m.fields here, in case of fields like:
			// DISTINCT t.user_id uid
<原文结束>

# <翻译开始>
// **注意**：在此处不要引用m.fields，以防字段类似以下情况：
// DISTINCT t.user_id uid
# <翻译结束>


<原文开始>
// Raw SQL Model, especially for UNION/UNION ALL featured SQL.
<原文结束>

# <翻译开始>
// 原生SQL模型，特别适用于包含UNION/UNION ALL特性的SQL语句。
# <翻译结束>


<原文开始>
		// DO NOT quote the m.fields where, in case of fields like:
		// DISTINCT t.user_id uid
<原文结束>

# <翻译开始>
// **不要**对m.fields进行引用，特别是在处理类似以下字段时：
// DISTINCT t.user_id AS uid
# <翻译结束>


<原文开始>
// getFieldsFiltered checks the fields and fieldsEx attributes, filters and returns the fields that will
// really be committed to underlying database driver.
<原文结束>

# <翻译开始>
// getFieldsFiltered 检查 fields 和 fieldsEx 属性，进行过滤并返回真正将提交到底层数据库驱动的字段。
# <翻译结束>


<原文开始>
// No filtering, containing special chars.
<原文结束>

# <翻译开始>
// 不进行过滤，包含特殊字符。
# <翻译结束>


<原文开始>
// Filter custom fields with fieldEx.
<原文结束>

# <翻译开始>
// 使用fieldEx过滤自定义字段。
# <翻译结束>


<原文开始>
// Filter table fields with fieldEx.
<原文结束>

# <翻译开始>
// 使用fieldEx过滤表字段。
# <翻译结束>


<原文开始>
// formatCondition formats where arguments of the model and returns a new condition sql and its arguments.
// Note that this function does not change any attribute value of the `m`.
//
// The parameter `limit1` specifies whether limits querying only one record if m.limit is not set.
<原文结束>

# <翻译开始>
// formatCondition 格式化模型的 where 条件参数，并返回一个新的条件 SQL 语句及其参数。
// 注意，此函数不会改变 `m` 的任何属性值。
//
// 参数 `limit1` 指定在 `m.limit` 未设置时是否限制查询仅一条记录。
# <翻译结束>


<原文开始>
// The count statement of sqlserver cannot contain the order by statement
<原文结束>

# <翻译开始>
// SQL Server 中的 count 语句不能包含 order by 子句
# <翻译结束>


<原文开始>
// Retrieve all records
<原文结束>

# <翻译开始>
// 获取所有记录
# <翻译结束>


<原文开始>
// There are custom fields.
<原文结束>

# <翻译开始>
// 存在自定义字段。
# <翻译结束>

