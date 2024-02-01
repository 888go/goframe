
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
// iString is the type assert api for String.
<原文结束>

# <翻译开始>
// iString 是 String 类型断言的 API。
# <翻译结束>


<原文开始>
// iIterator is the type assert api for Iterator.
<原文结束>

# <翻译开始>
// iIterator 是 Iterator 的类型断言 API。
# <翻译结束>


<原文开始>
// iInterfaces is the type assert api for Interfaces.
<原文结束>

# <翻译开始>
// iInterfaces 是 Interfaces 的类型断言 API。
# <翻译结束>


<原文开始>
// iNil if the type assert api for IsNil.
<原文结束>

# <翻译开始>
// iNil 是用于类型断言的 IsNil 方法。
# <翻译结束>


<原文开始>
// iTableName is the interface for retrieving table name for struct.
<原文结束>

# <翻译开始>
// iTableName 是用于为结构体获取表名的接口。
# <翻译结束>


<原文开始>
// quoteWordReg is the regular expression object for a word check.
<原文结束>

# <翻译开始>
// quoteWordReg 是用于单词检查的正则表达式对象。
# <翻译结束>


<原文开始>
// structTagPriority tags for struct converting for orm field mapping.
<原文结束>

# <翻译开始>
// structTagPriority 结构体标签，用于在ORM字段映射时进行结构体转换的标记。
# <翻译结束>


<原文开始>
// WithDB injects given db object into context and returns a new context.
<原文结束>

# <翻译开始>
// WithDB 将给定的 db 对象注入到上下文中并返回一个新的上下文。
# <翻译结束>


<原文开始>
// DBFromCtx retrieves and returns DB object from context.
<原文结束>

# <翻译开始>
// DBFromCtx 从context中获取并返回DB对象。
# <翻译结束>


<原文开始>
// ToSQL formats and returns the last one of sql statements in given closure function
// WITHOUT TRULY EXECUTING IT.
// Be caution that, all the following sql statements should use the context object passing by function `f`.
<原文结束>

# <翻译开始>
// ToSQL 格式化并返回给定闭包函数中的最后一个 SQL 语句，
// 但**并不会真正执行它**。
// 注意，所有后续的 SQL 语句都应使用通过 `f` 函数传递的上下文对象。
# <翻译结束>


<原文开始>
// CatchSQL catches and returns all sql statements that are EXECUTED in given closure function.
// Be caution that, all the following sql statements should use the context object passing by function `f`.
<原文结束>

# <翻译开始>
// CatchSQL 在给定闭包函数中捕获并返回所有已执行的SQL语句。
// 注意，所有后续SQL语句都应使用通过`f`函数传递的上下文对象。
# <翻译结束>


<原文开始>
// isDoStruct checks and returns whether given type is a DO struct.
<原文结束>

# <翻译开始>
// isDoStruct 检查并返回给定类型是否为DO结构体。
# <翻译结束>


<原文开始>
	// It checks by struct name like "XxxForDao", to be compatible with old version.
	// TODO remove this compatible codes in future.
<原文结束>

# <翻译开始>
// 它通过检查结构体名称（如 "XxxForDao"），以兼容旧版本。
// TODO：未来删除这些兼容代码。
# <翻译结束>


<原文开始>
// It checks by struct meta for DO struct in version.
<原文结束>

# <翻译开始>
// 它通过结构体元数据在指定版本中检查DO（Data Object）结构体。
# <翻译结束>


<原文开始>
// getTableNameFromOrmTag retrieves and returns the table name from struct object.
<原文结束>

# <翻译开始>
// getTableNameFromOrmTag 从结构体对象中获取并返回表名。
# <翻译结束>












<原文开始>
// Use the struct name of snake case.
<原文结束>

# <翻译开始>
// 使用蛇形命名法为结构体命名。
# <翻译结束>


<原文开始>
// ListItemValues retrieves and returns the elements of all item struct/map with key `key`.
// Note that the parameter `list` should be type of slice which contains elements of map or struct,
// or else it returns an empty slice.
//
// The parameter `list` supports types like:
// []map[string]interface{}
// []map[string]sub-map
// []struct
// []struct:sub-struct
// Note that the sub-map/sub-struct makes sense only if the optional parameter `subKey` is given.
// See gutil.ListItemValues.
<原文结束>

# <翻译开始>
// ListItemValues 函数用于获取并返回所有以 `key` 为键的项结构体或映射元素的值。
// 注意，参数 `list` 应为包含映射或结构体元素的切片类型，否则将返回一个空切片。
//
// 参数 `list` 支持以下类型：
// []map[string]interface{}
// []map[string]子映射
// []struct
// []struct:子结构体
// 注意，只有在提供可选参数 `subKey` 的情况下，子映射/子结构体才有意义。
// 参见 gutil.ListItemValues。
# <翻译结束>


<原文开始>
// ListItemValuesUnique retrieves and returns the unique elements of all struct/map with key `key`.
// Note that the parameter `list` should be type of slice which contains elements of map or struct,
// or else it returns an empty slice.
// See gutil.ListItemValuesUnique.
<原文结束>

# <翻译开始>
// ListItemValuesUnique 函数用于获取并返回所有具有键 `key` 的结构体或映射中的唯一元素。
// 注意，参数 `list` 应为包含映射或结构体元素的切片类型，否则将返回一个空切片。
// 参见 gutil.ListItemValuesUnique。
# <翻译结束>


<原文开始>
// GetInsertOperationByOption returns proper insert option with given parameter `option`.
<原文结束>

# <翻译开始>
// GetInsertOperationByOption 根据给定的参数 `option` 返回合适的插入选项。
# <翻译结束>


<原文开始>
// To be compatible with old version from v2.6.0.
<原文结束>

# <翻译开始>
// 为了与 v2.6.0 及其之前的旧版本兼容。
# <翻译结束>


<原文开始>
// DaToMapDeep is deprecated, use MapOrStructToMapDeep instead.
<原文结束>

# <翻译开始>
// DaToMapDeep 已废弃，请改用 MapOrStructToMapDeep。
# <翻译结束>


<原文开始>
// MapOrStructToMapDeep converts `value` to map type recursively(if attribute struct is embedded).
// The parameter `value` should be type of *map/map/*struct/struct.
// It supports embedded struct definition for struct.
<原文结束>

# <翻译开始>
// MapOrStructToMapDeep 递归地将`value`转换为map类型（如果属性结构体是嵌入的）。
// 参数`value`应为*map、map、*struct或struct类型。
// 它支持对结构体的嵌入式结构体定义。
# <翻译结束>


<原文开始>
// doQuoteTableName adds prefix string and quote chars for table name. It handles table string like:
// "user", "user u", "user,user_detail", "user u, user_detail ut", "user as u, user_detail as ut",
// "user.user u", "`user`.`user` u".
//
// Note that, this will automatically check the table prefix whether already added, if true it does
// nothing to the table name, or else adds the prefix to the table name and returns new table name with prefix.
<原文结束>

# <翻译开始>
// doQuoteTableName 为表名添加前缀字符串和引用字符。它处理诸如以下格式的表名：
// "user", "user u", "user,user_detail", "user u, user_detail ut", "user as u, user_detail as ut",
// "user.user u", "`user`.`user` u"。
// 注意，此函数会自动检查表名是否已包含前缀，如果已包含，则不对表名做任何操作；
// 否则，将前缀添加到表名中，并返回带有前缀的新表名。
# <翻译结束>







<原文开始>
// Check whether it has database name.
<原文结束>

# <翻译开始>
// 检查它是否包含数据库名。
# <翻译结束>


<原文开始>
// If the table name already has the prefix, skips the prefix adding.
<原文结束>

# <翻译开始>
// 如果表名已经包含前缀，则跳过添加前缀的操作。
# <翻译结束>







<原文开始>
// doQuoteWord checks given string `s` a word, if true quotes it with `charLeft` and `charRight`
// and returns the quoted string; or else returns `s` without any change.
<原文结束>

# <翻译开始>
// doQuoteWord 检查给定的字符串 `s` 是否为单词，如果是，则使用 `charLeft` 和 `charRight` 对其进行引号括起，
// 并返回引号包含的字符串；否则原样返回 `s` 不做任何改变。
# <翻译结束>


<原文开始>
// doQuoteString quotes string with quote chars.
// For example, if quote char is '`':
// "null"                             => "NULL"
// "user"                             => "`user`"
// "user u"                           => "`user` u"
// "user,user_detail"                 => "`user`,`user_detail`"
// "user u, user_detail ut"           => "`user` u,`user_detail` ut"
// "user.user u, user.user_detail ut" => "`user`.`user` u,`user`.`user_detail` ut"
// "u.id, u.name, u.age"              => "`u`.`id`,`u`.`name`,`u`.`age`"
// "u.id asc"                         => "`u`.`id` asc".
<原文结束>

# <翻译开始>
// doQuoteString 用引号字符对字符串进行引用。
// 例如，如果引号字符是 '`'：
// "null"                             => "NULL"
// "user"                             => "`user`"
// "user u"                           => "`user` u"
// "user,user_detail"                 => "`user`,`user_detail`"
// "user u, user_detail ut"           => "`user` u,`user_detail` ut"
// "user.user u, user.user_detail ut" => "`user`.`user` u,`user`.`user_detail` ut"
// "u.id, u.name, u.age"              => "`u`.`id`,`u`.`name`,`u`.`age`"
// "u.id asc"                         => "`u`.`id` asc"
// 此函数用于将输入的字符串中的字段名用指定的引号（本例中为反引号 `）包裹起来，以符合SQL语句中字段名引用的规范。对于包含点号（.）的情况，会分别对表名和字段名进行引号包裹，同时保留原有的空格和逗号等分隔符及排序关键字（如“asc”）。
# <翻译结束>


<原文开始>
			// Note:
			// mysql: u.uid
			// mssql double dots: Database..Table
<原文结束>

# <翻译开始>
// 注意：
// mysql: u.uid
// mssql 双点表示法：Database..Table
// （注释翻译如下）
// 注释：
// 在 MySQL 中，使用 `u.uid` 表示用户ID
// 在 MSSQL 中，双点（..）用于表示特定的表引用方式，即“Database..Table”，表示从指定数据库引用某个表
# <翻译结束>


<原文开始>
// GetPrimaryKeyCondition returns a new where condition by primary field name.
// The optional parameter `where` is like follows:
// 123                             => primary=123
// []int{1, 2, 3}                  => primary IN(1,2,3)
// "john"                          => primary='john'
// []string{"john", "smith"}       => primary IN('john','smith')
// g.Map{"id": g.Slice{1,2,3}}     => id IN(1,2,3)
// g.Map{"id": 1, "name": "john"}  => id=1 AND name='john'
// etc.
//
// Note that it returns the given `where` parameter directly if the `primary` is empty
// or length of `where` > 1.
<原文结束>

# <翻译开始>
// GetPrimaryKeyCondition 通过主键字段名称返回一个新的条件表达式。
// 可选参数 `where` 形如：
// 123                             => 主键=123
// []int{1, 2, 3}                  => 主键 IN(1,2,3)
// "john"                          => 主键='john'
// []string{"john", "smith"}       => 主键 IN('john','smith')
// g.Map{"id": g.Slice{1,2,3}}     => id IN(1,2,3)
// g.Map{"id": 1, "name": "john"}  => id=1 AND name='john'
// 等等
//
// 注意，如果 `primary` 为空或者 `where` 参数长度大于1，则直接返回给定的 `where` 参数。
# <翻译结束>


<原文开始>
// Ignore the parameter `primary`.
<原文结束>

# <翻译开始>
// 忽略参数`primary`。
# <翻译结束>


<原文开始>
// Table is used for fields mapping and filtering internally.
<原文结束>

# <翻译开始>
// Table 用于内部字段的映射和筛选。
# <翻译结束>


<原文开始>
	// Eg:
	// Where("id", []int{}).All()             -> SELECT xxx FROM xxx WHERE 0=1
	// Where("name", "").All()                -> SELECT xxx FROM xxx WHERE `name`=''
	// OmitEmpty().Where("id", []int{}).All() -> SELECT xxx FROM xxx
	// OmitEmpty().Where("name", "").All()    -> SELECT xxx FROM xxx
	// OmitEmpty().Where("1").All()           -> SELECT xxx FROM xxx WHERE 1
<原文结束>

# <翻译开始>
// 示例：
// Where("id", []int{}).All()             -> SELECT xxx FROM xxx WHERE 0=1
// 当条件字段为"id"且传入一个空的整数切片时，生成SQL语句为：从xxx表中选择所有列，但WHERE子句恒为假（0=1）
// Where("name", "").All()                -> SELECT xxx FROM xxx WHERE `name`=''
// 当条件字段为"name"且传入空字符串时，生成SQL语句为：从xxx表中选择所有列，其中"name"字段为空
// OmitEmpty().Where("id", []int{}).All() -> SELECT xxx FROM xxx
// 使用OmitEmpty()方法后，在条件字段为"id"且传入一个空的整数切片时，忽略无效的条件，生成SQL语句为：从xxx表中选择所有列，不带任何WHERE子句
// OmitEmpty().Where("name", "").All()    -> SELECT xxx FROM xxx
// 使用OmitEmpty()方法后，在条件字段为"name"且传入空字符串时，忽略无效的条件，生成SQL语句为：从xxx表中选择所有列，不带任何WHERE子句
// OmitEmpty().Where("1").All()           -> SELECT xxx FROM xxx WHERE 1
// 使用OmitEmpty()方法后，当条件为"1"时，生成SQL语句为：从xxx表中选择所有列，并在WHERE子句中使用恒为真的条件（1）
// 以上代码中的注释是描述了一个用于生成SQL查询语句的方法链式调用。根据不同的输入参数，会生成不同WHERE条件的SQL查询语句，而OmitEmpty()方法的作用是忽略无效或空的条件表达式。
# <翻译结束>


<原文开始>
// formatWhereHolder formats where statement and its arguments for `Where` and `Having` statements.
<原文结束>

# <翻译开始>
// formatWhereHolder 格式化 where 语句及其参数，用于 `Where` 和 `Having` 语句。
# <翻译结束>


<原文开始>
		// If the `where` parameter is `DO` struct, it then adds `OmitNil` option for this condition,
		// which will filter all nil parameters in `where`.
<原文结束>

# <翻译开始>
// 如果`where`参数是`DO`结构体，那么它会为这个条件添加`OmitNil`选项，
// 这将会过滤`where`中所有为nil的参数。
# <翻译结束>


<原文开始>
		// If `where` struct implements `iIterator` interface,
		// it then uses its Iterate function to iterate its key-value pairs.
		// For example, ListMap and TreeMap are ordered map,
		// which implement `iIterator` interface and are index-friendly for where conditions.
<原文结束>

# <翻译开始>
// 如果`where`结构体实现了`iIterator`接口，
// 那么它会使用其Iterate函数遍历自身的键值对。
// 例如，ListMap和TreeMap是有序映射，
// 它们实现了`iIterator`接口，对于where条件查询时更加友好。
# <翻译结束>


<原文开始>
// Automatically mapping and filtering the struct attribute.
<原文结束>

# <翻译开始>
// 自动映射并过滤结构体属性。
# <翻译结束>


<原文开始>
// If `Prefix` is given, it checks and retrieves the table name.
<原文结束>

# <翻译开始>
// 如果`Prefix`已给出，它会检查并检索表名。
# <翻译结束>


<原文开始>
// Mapping and filtering fields if `Table` is given.
<原文结束>

# <翻译开始>
// 如果提供了`Table`，则对字段进行映射和过滤。
# <翻译结束>


<原文开始>
// Put the struct attributes in sequence in Where statement.
<原文结束>

# <翻译开始>
// 在Where语句中按顺序放置结构体属性。
# <翻译结束>


<原文开始>
// Use tag value from `orm` as field name if specified.
<原文结束>

# <翻译开始>
// 如果已指定，则使用来自`orm`标签的值作为字段名。
# <翻译结束>







<原文开始>
		// Is `whereStr` a field name which composed as a key-value condition?
		// Eg:
		// Where("id", 1)
		// Where("id", g.Slice{1,2,3})
<原文结束>

# <翻译开始>
// `whereStr` 是否是由键值对构成的字段名？
// 例如：
// Where("id", 1) // 指定id为1的条件
// Where("id", g.Slice{1,2,3}) // 指定id在[1, 2, 3]范围内的条件
# <翻译结束>


<原文开始>
// If the first part is column name, it automatically adds prefix to the column.
<原文结束>

# <翻译开始>
// 如果第一部分是列名，它会自动为该列添加前缀。
# <翻译结束>


<原文开始>
		// Regular string and parameter place holder handling.
		// Eg:
		// Where("id in(?) and name=?", g.Slice{1,2,3}, "john")
<原文结束>

# <翻译开始>
// 正常字符串及参数占位符处理
// 示例：
// Where("id in(?) and name=?", g.Slice{1,2,3}, "john")
// 表示：当id在(1,2,3)中且name为"john"时的条件语句
# <翻译结束>


<原文开始>
			// ===============================================================
			// Sub query, which is always used along with a string condition.
			// ===============================================================
<原文结束>

# <翻译开始>
// ===============================================================
// 子查询，始终与字符串条件一起使用。
// ===============================================================
# <翻译结束>


<原文开始>
// Automatically adding the brackets.
<原文结束>

# <翻译开始>
// 自动添加括号
# <翻译结束>


<原文开始>
// Eg: Where/And/Or("uid>=", 1)
<原文结束>

# <翻译开始>
// 示例：Where/And/Or("uid>=", 1)
// （译注：在Go语言中，这段代码可能是用于构建SQL查询条件的方法调用，表示查询条件为“uid大于等于1”）
// Where: 设置或添加查询条件，如“uid>=”
// And: 在已有的查询条件下追加一个与（AND）关系的条件，此处表示“并且uid大于等于1”
// Or: 在已有的查询条件下追加一个或（OR）关系的条件，但根据示例实际未使用到OR操作
// 整体来看，这段代码片段是展示如何通过链式调用构建复杂查询条件的一种方式。
# <翻译结束>


<原文开始>
						// Eg:
						// Where("id", []int{1,2,3})
						// Where("user.id", []int{1,2,3})
<原文结束>

# <翻译开始>
// 示例：
// Where("id", []int{1,2,3}) // 根据id为1、2、3进行查询
// Where("user.id", []int{1,2,3}) // 根据user表中的id为1、2、3进行查询
// 以上Go语言代码的注释翻译成中文如下：
// ```go
// 例如：
// Where("id", []int{1,2,3}) // 用于指定id字段分别在1、2、3时的条件查询
// Where("user.id", []int{1,2,3}) // 用于指定user表中id字段分别在1、2、3时的条件查询
# <翻译结束>


<原文开始>
						// Eg:
						// Where("id", nil)
						// Where("user.id", nil)
<原文结束>

# <翻译开始>
// 示例：
// Where("id", nil) // 根据id查询
// Where("user.id", nil) // 根据user表中的id字段查询
# <翻译结束>


<原文开始>
						// Eg:
						// Where/And/Or("uid", 1)
						// Where/And/Or("user.uid", 1)
<原文结束>

# <翻译开始>
// 示例：
// Where/And/Or("uid", 1) // 指定条件：uid 等于 1
// Where/And/Or("user.uid", 1) // 指定条件：user.uid 等于 1
// 这段 Go 代码的注释翻译成中文为：
// ```go
// 示例：
// Where/And/Or 函数用于设置查询条件，例如：
// Where/And/Or("uid", 1) // 设置条件为 uid 等于 1
// Where/And/Or("user.uid", 1) // 设置条件为 user 表中的 uid 字段等于 1
// 这里的 `Where`、`And` 和 `Or` 都是可能的方法名，表示SQL语句中的“WHERE”子句以及逻辑连接符“AND”、“OR”，用于构建查询条件。在实际的数据库操作中，这些方法通常用于链式调用，形成更复杂的查询表达式。
# <翻译结束>


<原文开始>
// formatWhereInterfaces formats `where` as []interface{}.
<原文结束>

# <翻译开始>
// formatWhereInterfaces 将 `where` 格式化为 []interface{} 类型的切片。
# <翻译结束>


<原文开始>
// Db is the underlying DB object for current operation.
<原文结束>

# <翻译开始>
// Db 是当前操作所基于的底层 DB 对象。
# <翻译结束>


<原文开始>
// Buffer is the sql statement string without Args for current operation.
<原文结束>

# <翻译开始>
// Buffer 是当前操作中不包含 Args 的 SQL 语句字符串。
# <翻译结束>


<原文开始>
// Args is the full arguments of current operation.
<原文结束>

# <翻译开始>
// Args 是当前操作的完整参数。
# <翻译结束>


<原文开始>
// The field value, can be any types.
<原文结束>

# <翻译开始>
// 字段值，可以是任意类型。
# <翻译结束>







<原文开始>
// formatWhereKeyValue handles each key-value pair of the parameter map.
<原文结束>

# <翻译开始>
// formatWhereKeyValue 处理参数映射中的每一组键值对。
# <翻译结束>


<原文开始>
	// If the value is type of slice, and there's only one '?' holder in
	// the key string, it automatically adds '?' holder chars according to its arguments count
	// and converts it to "IN" statement.
<原文结束>

# <翻译开始>
// 如果值的类型是切片，并且在键字符串中只有一个 '?' 占位符，
// 那么它会根据其参数个数自动添加 '?' 占位符，并将其转换为 "IN" 语句。
# <翻译结束>


<原文开始>
// The key is a single field name.
<原文结束>

# <翻译开始>
// 键是一个单独的字段名称。
# <翻译结束>


<原文开始>
// The key may have operation chars.
<原文结束>

# <翻译开始>
// 密钥可能包含操作字符。
# <翻译结束>


<原文开始>
// It also supports "LIKE" statement, which we consider it an operator.
<原文结束>

# <翻译开始>
// 它还支持 "LIKE" 语句，我们认为它是一个运算符。
# <翻译结束>


<原文开始>
// Eg: Where(g.Map{"name like": "john%"})
<原文结束>

# <翻译开始>
// 示例：Where(g.Map{"name like": "john%"})
// （注：此代码片段使用了golang编写的数据库操作语句，其中"Where"表示SQL中的WHERE子句，用于设置查询条件。这里传入了一个g.Map类型的参数，它是一个键值对映射，其中"name like"是SQL的模糊查询条件，"john%"代表查询名字以"john"开头的所有记录。）
# <翻译结束>


<原文开始>
// Eg: Where(g.Map{"age > ": 16})
<原文结束>

# <翻译开始>
// 示例：Where(g.Map{"age > ": 16})
// （译注：此处代码为Go语言中使用g.Map进行条件筛选的示例，其中"g.Map"是一个自定义的映射类型，"age > "表示年龄大于，整体即表示筛选出年龄大于16的记录。）
# <翻译结束>


<原文开始>
// The key is a regular field name.
<原文结束>

# <翻译开始>
// key 是一个普通的字段名称。
# <翻译结束>


<原文开始>
					// The key is not a regular field name.
					// Eg: Where(g.Map{"age > 16": nil})
					// Issue: https://github.com/gogf/gf/issues/765
<原文结束>

# <翻译开始>
// 这个键不是常规的字段名称。
// 例如：Where(g.Map{"age > 16": nil})
// 相关问题：https://github.com/gogf/gf/issues/765
# <翻译结束>


<原文开始>
// handleArguments is an important function, which handles the sql and all its arguments
// before committing them to underlying driver.
<原文结束>

# <翻译开始>
// handleArguments 是一个重要的函数，它在将 SQL 及其所有参数提交给底层驱动之前，负责处理这些 SQL 和参数。
# <翻译结束>


<原文开始>
// insertHolderCount is used to calculate the inserting position for the '?' holder.
<原文结束>

# <翻译开始>
// insertHolderCount 用于计算 '?' 占位符的插入位置。
# <翻译结束>


<原文开始>
// Handles the slice arguments.
<原文结束>

# <翻译开始>
// 处理切片参数。
# <翻译结束>


<原文开始>
				// It does not split the type of []byte.
				// Eg: table.Where("name = ?", []byte("john"))
<原文结束>

# <翻译开始>
// 它不会分割[]byte类型的数据。
// 例如：table.Where("name = ?", []byte("john"))
// 翻译为：
// 此处的处理不会对[]byte类型的值进行分割操作。
// 举例说明：在调用table.Where方法时，可以传入一个如"name = ?"的条件字符串以及一个[]byte类型的值，如：[]byte("john")。
# <翻译结束>


<原文开始>
					// Empty slice argument, it converts the sql to a false sql.
					// Eg:
					// Query("select * from xxx where id in(?)", g.Slice{}) -> select * from xxx where 0=1
					// Where("id in(?)", g.Slice{}) -> WHERE 0=1
<原文结束>

# <翻译开始>
// 当传入空切片作为参数时，它会将SQL转换为一个永假的SQL语句。
// 例如：
// Query("select * from xxx where id in(?)", g.Slice{}) 将转换为 -> select * from xxx where 0=1
// Where("id in(?)", g.Slice{}) 将转换为 -> WHERE 0=1
# <翻译结束>


<原文开始>
				// If the '?' holder count equals the length of the slice,
				// it does not implement the arguments splitting logic.
				// Eg: db.Query("SELECT ?+?", g.Slice{1, 2})
<原文结束>

# <翻译开始>
// 如果'?'占位符的数量等于切片的长度，
// 则它不会执行参数分割逻辑。
// 例如：db.Query("SELECT ?+?", g.Slice{1, 2})
# <翻译结束>


<原文开始>
// counter is used to finding the inserting position for the '?' holder.
<原文结束>

# <翻译开始>
// counter 用于计算 '?' 占位符的插入位置。
# <翻译结束>







<原文开始>
// The underlying driver supports time.Time/*time.Time types.
<原文结束>

# <翻译开始>
// 底层驱动程序支持 time.Time 类型（注：time.Time 是 Go 语言中的时间类型）。
# <翻译结束>


<原文开始>
					// It converts the struct to string in default
					// if it has implemented the String interface.
<原文结束>

# <翻译开始>
// 如果结构体实现了 String 接口，它会默认将该结构体转换为字符串。
# <翻译结束>


<原文开始>
// FormatSqlWithArgs binds the arguments to the sql string and returns a complete
// sql string, just for debugging.
<原文结束>

# <翻译开始>
// FormatSqlWithArgs 将参数绑定到sql字符串，并返回一个完整的sql字符串，仅用于调试。
# <翻译结束>


<原文开始>
// Parameters of type Raw do not require special treatment
<原文结束>

# <翻译开始>
// 类型为Raw的参数不需要进行特殊处理
# <翻译结束>


<原文开始>
// The field name, eg: "id", "name", etc.
<原文结束>

# <翻译开始>
// 字段名称，例如："id"、"name"等。
# <翻译结束>


<原文开始>
// Ignores current condition key if `value` is empty.
<原文结束>

# <翻译开始>
// 如果`value`为空，则忽略当前条件键。
# <翻译结束>


<原文开始>
// Field prefix, eg: "user", "order", etc.
<原文结束>

# <翻译开始>
// 字段前缀，例如："user"、"order"等。
# <翻译结束>







<原文开始>
// Use the interface value.
<原文结束>

# <翻译开始>
// 使用接口值。
# <翻译结束>


<原文开始>
// User meta data tag "orm".
<原文结束>

# <翻译开始>
// 用户元数据标签 "orm"。
# <翻译结束>


<原文开始>
// Trim the security chars.
<原文结束>

# <翻译开始>
// 剔除安全字符。
# <翻译结束>


<原文开始>
// Add the security chars.
<原文结束>

# <翻译开始>
// 添加安全字符。
# <翻译结束>


<原文开始>
// The value in Where type.
<原文结束>

# <翻译开始>
// Where 类型中的值。
# <翻译结束>


<原文开始>
// Special struct handling.
<原文结束>

# <翻译开始>
// 特殊结构体处理。
# <翻译结束>

