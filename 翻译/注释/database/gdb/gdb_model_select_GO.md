
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
// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as slice type.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter `where` is the same as the parameter of Model.Where function,
// see Model.Where.
<原文结束>

# <翻译开始>
// 所有执行的是 "FROM ..." 语句针对该模型。
// 它从表中检索记录，并将结果作为切片类型返回。
// 如果根据给定条件从表中没有检索到任何记录，它将返回nil。
//
// 可选参数 `where` 和 Model.Where 函数的参数相同，
// 请参阅 Model.Where。 md5:fd88d2addfbe9655
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
// AllAndCount 从模型中检索所有记录以及记录的总数量。
// 如果 useFieldForCount 为真，它将使用模型中指定的字段进行计数；
// 否则，它将使用常数值1来进行计数。
// 此方法返回结果作为一个记录切片，记录的总数量，以及可能存在的错误。
// where 参数是一个可选的条件列表，用于在检索记录时应用。
//
// 示例：
//
//	var model Model
//	var result []Record
//	var count int
//	where := []interface{}{"name = ?", "John"}
//	result, count, err := model.AllAndCount(true, where...)
//	if err != nil {
//	    // 处理错误。
//	}
//	fmt.Println(result, count) md5:b631bbec9e186f68
# <翻译结束>


<原文开始>
// Clone the model for counting
<原文结束>

# <翻译开始>
	// 克隆模型用于计数. md5:662b7475962d2c44
# <翻译结束>


<原文开始>
// If useFieldForCount is false, set the fields to a constant value of 1 for counting
<原文结束>

# <翻译开始>
	// 如果useFieldForCount为false，将字段设置为计数的恒定值1. md5:2eea55571801d2ab
# <翻译结束>


<原文开始>
// Get the total count of records
<原文结束>

# <翻译开始>
	// 获取记录的总数. md5:d21517ef51fd67f3
# <翻译结束>


<原文开始>
// If the total count is 0, there are no records to retrieve, so return early
<原文结束>

# <翻译开始>
	// 如果总记录数为0，就没有需要检索的记录，因此提前返回. md5:ae90d44fd00f71aa
# <翻译结束>


<原文开始>
// Chunk iterates the query result with given `size` and `handler` function.
<原文结束>

# <翻译开始>
// Chunk 使用给定的 `size` 和 `handler` 函数来分块迭代查询结果。 md5:4c5d0d282b8e1fe4
# <翻译结束>


<原文开始>
// One retrieves one record from table and returns the result as map type.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter `where` is the same as the parameter of Model.Where function,
// see Model.Where.
<原文结束>

# <翻译开始>
// 从表中获取一条记录，并将结果作为map类型返回。如果使用给定条件从表中没有检索到记录，则返回nil。
//
// 可选参数`where`与Model.Where函数的参数相同，参见Model.Where。 md5:b48f8e0c5d07b484
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
// 从数据库查询并返回数据值作为切片。
// 注意，如果结果中有多个列，它会随机返回一列的值。
//
// 如果提供了可选参数 `fieldsAndWhere`，则 fieldsAndWhere[0] 是选择的字段，
// 而 fieldsAndWhere[1:] 则被视为 where 条件字段。
// 参见 Model.Fields 和 Model.Where 函数。 md5:1de6885dc1e83172
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
// Struct 从表中检索一条记录，并将其转换为给定的结构体。
// 参数 `pointer` 应为 *struct 或 **struct 类型。如果给出 **struct 类型，
// 则在转换过程中可以在内部创建该结构体。
//
// 可选参数 `where` 与 Model.Where 函数的参数相同，
// 详情请参阅 Model.Where。
//
// 注意，如果给定参数 `pointer` 指向一个具有默认值的变量，并且根据给定条件从表中没有检索到任何记录，
// 则它将返回 sql.ErrNoRows 错误。
//
// 示例：
// user := new(User)
// err  := db.Model("user").Where("id", 1).Scan(user)
//
// user := (*User)(nil)
// err  := db.Model("user").Where("id", 1).Scan(&user) md5:473a4005864a522f
# <翻译结束>


<原文开始>
// Auto selecting fields by struct attributes.
<原文结束>

# <翻译开始>
	// 自动通过结构体属性选择字段。 md5:25f031330d67c88b
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
// 参数 `pointer` 应该是类型为 *[]struct 或 *[]*struct。在转换过程中，它可以内部创建和填充结构体切片。
//
// 可选参数 `where` 和 Model.Where 函数的参数相同，参见 Model.Where。
//
// 注意，如果给定的 `pointer` 指向一个具有默认值的变量，并且根据给定条件从表中没有检索到任何记录，则返回 sql.ErrNoRows。
//
// 示例：
// users := ([]User)(nil)
// err   := db.Model("user").Scan(&users)
//
// users := ([]*User)(nil)
// err   := db.Model("user").Scan(&users) md5:bd3102709ae8c192
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
// Scan会根据参数`pointer`的类型自动调用Struct或Structs函数。
// 如果`pointer`是类型`*struct`或`**struct`，它会调用doStruct函数。
// 如果`pointer`是类型`*[]struct`或`*[]*struct`，它会调用doStructs函数。
//
// 可选参数`where`与Model.Where函数的参数相同，参见Model.Where。
//
// 注意，如果给定的`pointer`指向一个具有默认值的变量，并且在表中没有满足条件的记录，它将返回sql.ErrNoRows错误。
//
// 示例：
// user := new(User)
// err := db.Model("user").Where("id", 1).Scan(user)
//
// user := (*User)(nil)
// err := db.Model("user").Where("id", 1).Scan(&user)
//
// users := ([]User)(nil)
// err := db.Model("user").Scan(&users)
//
// users := ([]*User)(nil)
// err := db.Model("user").Scan(&users) md5:a6df07ddafe5975a
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
// ScanAndCount 扫描与给定条件匹配的单条记录或记录数组，并计算符合这些条件的总记录数。
// 如果 useFieldForCount 为 true，它将使用模型中指定的字段进行计数；
// pointer 参数是一个指向结构体的指针，用于存储扫描到的数据。
// pointerCount 参数是一个指向整数的指针，将被设置为符合给定条件的总记录数。
// where 参数是可选的条件列表，用于在检索记录时使用。
//
// 示例：
//
//	var count int
//	user := new(User)
//	err  := db.Model("user").Where("id", 1).ScanAndCount(user, &count, true)
//	fmt.Println(user, count)
//
// 示例（联接）：
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
//		ScanAndCount(&users, &count, false) md5:984fa8f0e50708f4
# <翻译结束>


<原文开始>
// support Fields with *, example: .Fields("a.*, b.name"). Count sql is select count(1) from xxx
<原文结束>

# <翻译开始>
	// 支持使用 * 的字段，例如：.Fields("a.*, b.name")。计数SQL为：select count(1) from xxx. md5:a3fc56bcc1dcba76
# <翻译结束>


<原文开始>
// ScanList converts `r` to struct slice which contains other complex struct attributes.
// Note that the parameter `listPointer` should be type of *[]struct/*[]*struct.
//
// See Result.ScanList.
<原文结束>

# <翻译开始>
// ScanList 将 `r` 转换为包含其他复杂结构体属性的切片。请注意，参数 `listPointer` 的类型应该是 `*[]struct` 或 `*[]*struct`。
//
// 参见 Result.ScanList。 md5:4116492a123661b5
# <翻译结束>


<原文开始>
// There are custom fields.
<原文结束>

# <翻译开始>
		// 有自定义字段。 md5:57eb1cc07145128c
# <翻译结束>


<原文开始>
// Filter fields using temporary created struct using reflect.New.
<原文结束>

# <翻译开始>
		// 使用反射创建的临时结构体过滤字段。 md5:6873597e9de7f128
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
// Value 从表中获取指定记录的值，并将结果作为接口类型返回。
// 如果根据给定条件在表中找不到记录，它将返回nil。
//
// 如果提供了可选参数 `fieldsAndWhere`，其中 fieldsAndWhere[0] 是选择的字段，
// 而 fieldsAndWhere[1:] 用作 WHERE 条件字段。
// 另请参阅 Model.Fields 和 Model.Where 函数。 md5:e6b48ca188d3d208
# <翻译结束>


<原文开始>
// Count does "SELECT COUNT(x) FROM ..." statement for the model.
// The optional parameter `where` is the same as the parameter of Model.Where function,
// see Model.Where.
<原文结束>

# <翻译开始>
// Count 对于该模型执行 "SELECT COUNT(x) FROM ..." 语句。
// 可选参数 `where` 和 Model.Where 函数的参数相同，参见 Model.Where。 md5:52b3d2e0e43bb2af
# <翻译结束>


<原文开始>
// CountColumn does "SELECT COUNT(x) FROM ..." statement for the model.
<原文结束>

# <翻译开始>
// CountColumn 执行对模型的 "SELECT COUNT(x) FROM ..." 语句。 md5:150abf4737c4588c
# <翻译结束>


<原文开始>
// Min does "SELECT MIN(x) FROM ..." statement for the model.
<原文结束>

# <翻译开始>
// Min 为模型执行 "SELECT MIN(x) FROM ..." 语句。 md5:e2fc098c542503d1
# <翻译结束>


<原文开始>
// Max does "SELECT MAX(x) FROM ..." statement for the model.
<原文结束>

# <翻译开始>
// Max 对模型执行 "SELECT MAX(x) FROM ..." 语句。 md5:bb6b4d0dc51fbfaf
# <翻译结束>


<原文开始>
// Avg does "SELECT AVG(x) FROM ..." statement for the model.
<原文结束>

# <翻译开始>
// Avg 对于该模型执行"SELECT AVG(x) FROM ..." 语句。 md5:9b360a11d26d6fca
# <翻译结束>


<原文开始>
// Sum does "SELECT SUM(x) FROM ..." statement for the model.
<原文结束>

# <翻译开始>
// Sum 对于该模型执行 "SELECT SUM(x) FROM ..." 语句。 md5:bcbe9e29cd168603
# <翻译结束>


<原文开始>
// Union does "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." statement for the model.
<原文结束>

# <翻译开始>
// Union 为模型执行 "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." 语句。 md5:97431dccd533414e
# <翻译结束>


<原文开始>
// UnionAll does "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ..." statement for the model.
<原文结束>

# <翻译开始>
// UnionAll 对模型执行 "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ..." 语句。 md5:d112aec0d1929661
# <翻译结束>


<原文开始>
// Limit sets the "LIMIT" statement for the model.
// The parameter `limit` can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
<原文结束>

# <翻译开始>
// Limit 设置模型的 "LIMIT" 语句。
// 参数 `limit` 可以是一个或两个数字。如果传递两个数字，它将为模型设置 "LIMIT limit[0],limit[1]" 语句；否则，它将设置 "LIMIT limit[0]" 语句。 md5:fd06ed75a128d403
# <翻译结束>


<原文开始>
// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
<原文结束>

# <翻译开始>
// Offset 设置模型的“OFFSET”语句。它只对某些数据库（如 SQLServer、PostgreSQL 等）有意义。 md5:5a99cab6ce558c69
# <翻译结束>


<原文开始>
// Distinct forces the query to only return distinct results.
<原文结束>

# <翻译开始>
// Distinct 强制查询仅返回不重复的结果。 md5:ead62c0e4b4795ab
# <翻译结束>


<原文开始>
// Page sets the paging number for the model.
// The parameter `page` is started from 1 for paging.
// Note that, it differs that the Limit function starts from 0 for "LIMIT" statement.
<原文结束>

# <翻译开始>
// Page 设置模型的分页号。
// 参数 `page` 的起始值为1，用于分页。
// 注意，这与Limit函数在"LIMIT"语句中从0开始不同。 md5:02b920e99951ce53
# <翻译结束>


<原文开始>
// Having sets the having statement for the model.
// The parameters of this function usage are as the same as function Where.
// See Where.
<原文结束>

# <翻译开始>
// Having 设置模型的having语句。
// 该函数的使用参数与Where函数相同。
// 参见Where。 md5:b4e737511765f79f
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
// doGetAll 对应于 "SELECT FROM ..." 语句，用于模型。
// 它从表中检索记录，并以切片类型返回结果。如果根据给定条件从表中没有检索到记录，则返回 nil。
//
// 参数 `limit1` 指定当模型的 `limit` 未设置时，是否只查询一条记录。
// 可选参数 `where` 的用法与 Model.Where 函数的参数相同，参见 Model.Where。 md5:d4f7ecca6c5aaa48
# <翻译结束>


<原文开始>
// doGetAllBySql does the select statement on the database.
<原文结束>

# <翻译开始>
// doGetAllBySql 在数据库上执行选择语句。 md5:b9498c08926ceb6a
# <翻译结束>


<原文开始>
			// DO NOT quote the m.fields here, in case of fields like:
			// DISTINCT t.user_id uid
<原文结束>

# <翻译开始>
			// 不要在这里引用m.fields，以防出现如下的字段情况：
			// DISTINCT t.user_id uid md5:97ff3b5639a12242
# <翻译结束>


<原文开始>
// Raw SQL Model, especially for UNION/UNION ALL featured SQL.
<原文结束>

# <翻译开始>
		// 原生SQL模型，特别适用于包含UNION/UNION ALL特性的SQL。 md5:03942fe59d08c0b4
# <翻译结束>


<原文开始>
		// DO NOT quote the m.fields where, in case of fields like:
		// DISTINCT t.user_id uid
<原文结束>

# <翻译开始>
		// 请不要在 m.fields 中引用，例如：
		// 如果字段为：
		// DISTINCT t.user_id uid md5:e3b773558c54f2eb
# <翻译结束>


<原文开始>
// getFieldsFiltered checks the fields and fieldsEx attributes, filters and returns the fields that will
// really be committed to underlying database driver.
<原文结束>

# <翻译开始>
// getFieldsFiltered 检查字段和字段排除属性，过滤并返回那些将真正被提交到底层数据库驱动的字段。 md5:e8c5bf23790637e0
# <翻译结束>


<原文开始>
// No filtering, containing special chars.
<原文结束>

# <翻译开始>
		// 没有过滤，包含特殊字符。 md5:f2ccc24dfd015b85
# <翻译结束>


<原文开始>
// Filter custom fields with fieldEx.
<原文结束>

# <翻译开始>
		// 使用fieldEx过滤自定义字段。 md5:edee7113e1c2daf9
# <翻译结束>


<原文开始>
// Filter table fields with fieldEx.
<原文结束>

# <翻译开始>
		// 使用fieldEx过滤表格字段。 md5:e15e7d68ef0a3c54
# <翻译结束>


<原文开始>
// formatCondition formats where arguments of the model and returns a new condition sql and its arguments.
// Note that this function does not change any attribute value of the `m`.
//
// The parameter `limit1` specifies whether limits querying only one record if m.limit is not set.
<原文结束>

# <翻译开始>
// formatCondition 格式化模型的where参数，并返回一个新的条件SQL及其参数。
// 注意，此函数不会更改`m`的任何属性值。
//
// 参数 `limit1` 指定如果m.limit未设置，是否限制只查询一条记录。 md5:d251ca8a182de4ff
# <翻译结束>


<原文开始>
// The count statement of sqlserver cannot contain the order by statement
<原文结束>

# <翻译开始>
// SQLServer 的 count 语句中不能包含 order by 子句. md5:a176c1f7165860e0
# <翻译结束>

