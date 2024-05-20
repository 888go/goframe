
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
// iString is the type assert api for String.
<原文结束>

# <翻译开始>
// iString 是用于字符串的类型断言接口。. md5:a51d442c33ac2cf4
# <翻译结束>


<原文开始>
// iIterator is the type assert api for Iterator.
<原文结束>

# <翻译开始>
// iIterator是Iterator的类型断言API。. md5:9e146fc0e640273e
# <翻译结束>


<原文开始>
// iInterfaces is the type assert api for Interfaces.
<原文结束>

# <翻译开始>
// iInterfaces是Interfaces类型的断言API。. md5:843230a1ccff49f5
# <翻译结束>


<原文开始>
// iNil if the type assert api for IsNil.
<原文结束>

# <翻译开始>
// iNil 用于类型断言接口以检查是否为Nil。. md5:49ddfde26b501402
# <翻译结束>


<原文开始>
// iTableName is the interface for retrieving table name for struct.
<原文结束>

# <翻译开始>
// iTableName 是用于获取结构体对应表名的接口。. md5:f3583f3a54701536
# <翻译结束>


<原文开始>
// quoteWordReg is the regular expression object for a word check.
<原文结束>

# <翻译开始>
// quoteWordReg是用于单词检查的正则表达式对象。. md5:99c41eabb9d23388
# <翻译结束>


<原文开始>
// structTagPriority tags for struct converting for orm field mapping.
<原文结束>

# <翻译开始>
// structTagPriority 是用于结构体转换为 ORM 字段映射的标签优先级。. md5:6e5a8632b6c8e48f
# <翻译结束>


<原文开始>
// WithDB injects given db object into context and returns a new context.
<原文结束>

# <翻译开始>
// WithDB 将给定的db对象注入到context中并返回一个新的context。. md5:e414408e96157a02
# <翻译结束>


<原文开始>
// DBFromCtx retrieves and returns DB object from context.
<原文结束>

# <翻译开始>
// DBFromCtx 从上下文中获取并返回DB对象。. md5:90c01e951db89218
# <翻译结束>


<原文开始>
// ToSQL formats and returns the last one of sql statements in given closure function
// WITHOUT TRULY EXECUTING IT.
// Be caution that, all the following sql statements should use the context object passing by function `f`.
<原文结束>

# <翻译开始>
// ToSQL 将给定闭包函数中的最后一个 SQL 语句格式化并返回，但并不会真正执行。
// 注意，所有后续的 SQL 语句都应该使用由函数 `f` 传递的上下文对象。
// md5:3fe82285d68728a0
# <翻译结束>


<原文开始>
// CatchSQL catches and returns all sql statements that are EXECUTED in given closure function.
// Be caution that, all the following sql statements should use the context object passing by function `f`.
<原文结束>

# <翻译开始>
// CatchSQL捕获并返回在给定闭包函数中执行的所有SQL语句。
// 注意，所有以下SQL语句都应使用由`f`函数传递的context对象。
// md5:1088111f1248173d
# <翻译结束>


<原文开始>
// isDoStruct checks and returns whether given type is a DO struct.
<原文结束>

# <翻译开始>
// isDoStruct 检查并返回给定类型是否为 DO 结构体。. md5:c235f077b3b3fcc5
# <翻译结束>


<原文开始>
	// It checks by struct name like "XxxForDao", to be compatible with old version.
	// TODO remove this compatible codes in future.
<原文结束>

# <翻译开始>
// 它通过结构体名称如 "XxxForDao" 进行检查，以兼容旧版本。
// TODO: 未来某个时候移除这个兼容性代码。
// md5:c8dc9518a3014aca
# <翻译结束>


<原文开始>
// It checks by struct meta for DO struct in version.
<原文结束>

# <翻译开始>
// 它通过结构体元数据检查version中的DO结构。. md5:1651dfa7d2770eb0
# <翻译结束>


<原文开始>
// getTableNameFromOrmTag retrieves and returns the table name from struct object.
<原文结束>

# <翻译开始>
// getTableNameFromOrmTag 从结构体对象中检索并返回表名。. md5:9da9a9980775dc38
# <翻译结束>


<原文开始>
// Use the interface value.
<原文结束>

# <翻译开始>
// 使用接口值。. md5:d769d38046ef266c
# <翻译结束>


<原文开始>
// User meta data tag "orm".
<原文结束>

# <翻译开始>
// User meta data 标签 "orm"。. md5:7269c54b8f9aa97f
# <翻译结束>


<原文开始>
// Use the struct name of snake case.
<原文结束>

# <翻译开始>
// 使用蛇形命名的结构体名称。. md5:02b76586ae24bd54
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
// ListItemValues 从所有具有键为 `key` 的映射或结构体元素中检索并返回。请注意，参数 `list` 应该是包含映射或结构体元素的切片，否则将返回一个空切片。
// 
// 参数 `list` 支持以下类型：
// []map[string]interface{}
// []map[string]子映射
// []struct
// []struct:子结构体
// 如果提供了可选参数 `subKey`，子映射/子结构体才有意义。请参阅 gutil.ListItemValues。
// md5:e67327bcbcd82096
# <翻译结束>


<原文开始>
// ListItemValuesUnique retrieves and returns the unique elements of all struct/map with key `key`.
// Note that the parameter `list` should be type of slice which contains elements of map or struct,
// or else it returns an empty slice.
// See gutil.ListItemValuesUnique.
<原文结束>

# <翻译开始>
// ListItemValuesUnique 获取并返回所有结构体/映射中键为`key`的唯一元素。
// 注意，参数`list`应为切片类型，且包含的元素为映射或结构体，
// 否则将返回一个空切片。
// 参见gutil.ListItemValuesUnique。
// md5:aa00cb15fafa41ba
# <翻译结束>


<原文开始>
// GetInsertOperationByOption returns proper insert option with given parameter `option`.
<原文结束>

# <翻译开始>
// GetInsertOperationByOption 根据给定的 `option` 参数返回合适的插入操作选项。. md5:19b87dd1244d55ec
# <翻译结束>


<原文开始>
// To be compatible with old version from v2.6.0.
<原文结束>

# <翻译开始>
// 为了与从v2.6.0版本兼容。. md5:c3311f248eacbf1e
# <翻译结束>


<原文开始>
// It here converts all struct/map slice attributes to json string.
<原文结束>

# <翻译开始>
// 这里将结构体/映射切片的所有属性转换为JSON字符串。. md5:0355d0233b27ae20
# <翻译结束>


<原文开始>
// Check map item slice item.
<原文结束>

# <翻译开始>
// 检查map中的项目切片项。. md5:e8ef9cf68f78aa92
# <翻译结束>


<原文开始>
// Check slice item type struct/map type.
<原文结束>

# <翻译开始>
// 检查切片元素类型为结构体/映射类型。. md5:8ae4ace59a5f1070
# <翻译结束>


<原文开始>
// MapOrStructToMapDeep converts `value` to map type recursively(if attribute struct is embedded).
// The parameter `value` should be type of *map/map/*struct/struct.
// It supports embedded struct definition for struct.
<原文结束>

# <翻译开始>
// MapOrStructToMapDeep 递归地将 `value` 转换为映射类型（如果属性结构体被嵌入）。
// 参数 `value` 应该是 *map、map、*struct 或 struct 类型。
// 它支持结构体中的嵌入式结构体定义。
// md5:3d9a4c7ad65d9fe1
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
// doQuoteTableName 在表名前添加前缀字符串和引号。它处理的表名类型包括：
// "user", "user u", "user,user_detail", "user u, user_detail ut", "user as u, user_detail as ut",
// "user.user u", "``user``.``user`` u"。
// 
// 注意，它会自动检查表名前缀是否已添加。如果已添加，则不进行任何操作；否则，会在表名前添加前缀，并返回带有前缀的新表名。
// md5:fc5ea60c27043ac8
# <翻译结束>


<原文开始>
// Trim the security chars.
<原文结束>

# <翻译开始>
// 去除安全字符。. md5:13d69b898c5635c6
# <翻译结束>


<原文开始>
// Check whether it has database name.
<原文结束>

# <翻译开始>
// 检查是否具有数据库名称。. md5:8d7ee8b347dfbbac
# <翻译结束>


<原文开始>
// If the table name already has the prefix, skips the prefix adding.
<原文结束>

# <翻译开始>
// 如果表名已经包含前缀，则跳过添加前缀。. md5:047ee36460a1c519
# <翻译结束>


<原文开始>
// Add the security chars.
<原文结束>

# <翻译开始>
// 添加安全字符。. md5:b299af34cac147b9
# <翻译结束>


<原文开始>
// doQuoteWord checks given string `s` a word, if true quotes it with `charLeft` and `charRight`
// and returns the quoted string; or else returns `s` without any change.
<原文结束>

# <翻译开始>
// doQuoteWord 检查给定的字符串 `s` 是否为一个单词，如果是，则使用 `charLeft` 和 `charRight` 对其进行引用并返回被引用的字符串；否则原样返回 `s`，不做任何改变。
// md5:ac0c8a621b951784
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
// doQuoteString 使用引号字符对字符串进行引用。
// 例如，如果引用字符是 '`'：
// "null"                             => "NULL"
// "user"                             => "`user`"
// "user u"                           => "`user` u"
// "user,user_detail"                 => "`user`,`user_detail`"
// "user u, user_detail ut"           => "`user` u,`user_detail` ut"
// "user.user u, user.user_detail ut" => "`user`.`user` u,`user`.`user_detail` ut"
// "u.id, u.name, u.age"              => "`u`.`id`,`u`.`name`,`u`.`age`"
// "u.id asc"                         => "`u`.`id` asc"
// md5:556d2b4db186afe8
# <翻译结束>


<原文开始>
			// Note:
			// mysql: u.uid
			// mssql double dots: Database..Table
<原文结束>

# <翻译开始>
// 注释：
// mysql：u.uid
// mssql：使用双点表示法 Database..Table
// md5:66df82a8563f168b
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
// GetPrimaryKeyCondition 通过主键字段名返回一个新的WHERE条件。可选参数`where`的格式如下：
// 123                            => primary=123
// []int{1, 2, 3}                  => primary IN(1,2,3)
// "john"                          => primary='john'
// []string{"john", "smith"}       => primary IN('john','smith')
// g.Map{"id": g.Slice{1,2,3}}     => id IN(1,2,3)
// g.Map{"id": 1, "name": "john"}  => id=1 AND name='john'
// 等等。
//
// 注意，如果主键为空或者`where`参数长度大于1，它会直接返回给定的`where`参数。
// md5:748dfa9c0f0d93b5
# <翻译结束>


<原文开始>
// Ignore the parameter `primary`.
<原文结束>

# <翻译开始>
// 忽略参数 `primary`。. md5:be747ae45f6887e1
# <翻译结束>


<原文开始>
// Table is used for fields mapping and filtering internally.
<原文结束>

# <翻译开始>
// Table 用于内部字段映射和过滤。. md5:6f0adcc9f806782a
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
// 例如：
// Where("id", []int{}) .All()             -> 选择xxx FROM xxx WHERE 0=1
// Where("name", "") .All()                -> 选择xxx FROM xxx WHERE `name`= ''
// OmitEmpty() .Where("id", []int{}) .All() -> 选择xxx FROM xxx
// OmitEmpty() .Where("name", "") .All()    -> 选择xxx FROM xxx
// OmitEmpty() .Where("1") .All()           -> 选择xxx FROM xxx WHERE 1
// md5:13a1baa59d83a9fe
# <翻译结束>


<原文开始>
// formatWhereHolder formats where statement and its arguments for `Where` and `Having` statements.
<原文结束>

# <翻译开始>
// formatWhereHolder 格式化 WHERE 和 HAVING 语句及其参数。. md5:bd64f5b4ad435946
# <翻译结束>


<原文开始>
		// If the `where` parameter is `DO` struct, it then adds `OmitNil` option for this condition,
		// which will filter all nil parameters in `where`.
<原文结束>

# <翻译开始>
// 如果`where`参数是`DO`结构体，那么它会为这个条件添加`OmitNil`选项，
// 这将会过滤掉`where`中的所有nil参数。
// md5:dc90f650b5a33b25
# <翻译结束>


<原文开始>
		// If `where` struct implements `iIterator` interface,
		// it then uses its Iterate function to iterate its key-value pairs.
		// For example, ListMap and TreeMap are ordered map,
		// which implement `iIterator` interface and are index-friendly for where conditions.
<原文结束>

# <翻译开始>
// 如果`where`结构体实现了`iIterator`接口，
// 则使用其Iterate函数来遍历键值对。
// 例如，ListMap和TreeMap是有序映射，
// 它们实现了`iIterator`接口，并且对where条件的索引友好。
// md5:d2bd42ea2a41d114
# <翻译结束>


<原文开始>
// Automatically mapping and filtering the struct attribute.
<原文结束>

# <翻译开始>
// 自动映射和过滤结构体属性。. md5:8dc7e982c45f4e9d
# <翻译结束>


<原文开始>
// If `Prefix` is given, it checks and retrieves the table name.
<原文结束>

# <翻译开始>
// 如果提供了`Prefix`，则检查并获取表名。. md5:68af0bee501f583b
# <翻译结束>


<原文开始>
// Mapping and filtering fields if `Table` is given.
<原文结束>

# <翻译开始>
// 如果提供了`Table`，则对字段进行映射和过滤。. md5:369e7f3da1245a29
# <翻译结束>


<原文开始>
// Put the struct attributes in sequence in Where statement.
<原文结束>

# <翻译开始>
// 将结构体属性按顺序放入Where语句中。. md5:e18e7534d834dd8a
# <翻译结束>


<原文开始>
// Use tag value from `orm` as field name if specified.
<原文结束>

# <翻译开始>
// 如果指定了，使用`orm`标签的值作为字段名。. md5:0e761199d5be562b
# <翻译结束>


<原文开始>
		// Is `whereStr` a field name which composed as a key-value condition?
		// Eg:
		// Where("id", 1)
		// Where("id", g.Slice{1,2,3})
<原文结束>

# <翻译开始>
// `whereStr` 是用作键值条件的字段名称吗？
// 例如：
// Where("id", 1)
// Where("id", g.Slice{1,2,3}) 
// 
// 这段Go代码中的注释是在询问`whereStr`是否是一个用作键值对条件的字段名。它举例说明了如何使用`where`函数，其中第一个参数是字段名（如"id"），第二个参数可以是单个值（如1）或一个包含多个值的切片（如g.Slice{1,2,3}）。
// md5:3e3e293b8d2b6e27
# <翻译结束>


<原文开始>
// If the first part is column name, it automatically adds prefix to the column.
<原文结束>

# <翻译开始>
// 如果第一部分是列名，它会自动为列添加前缀。. md5:8174a130580bbf74
# <翻译结束>


<原文开始>
		// Regular string and parameter place holder handling.
		// Eg:
		// Where("id in(?) and name=?", g.Slice{1,2,3}, "john")
<原文结束>

# <翻译开始>
// 普通字符串和参数占位符处理。
// 例如：
// Where("id in(?) and name=?", g.Slice{1,2,3}, "john")
// md5:8a2be53569a9ada1
# <翻译结束>


<原文开始>
			// ===============================================================
			// Sub query, which is always used along with a string condition.
			// ===============================================================
<原文结束>

# <翻译开始>
// ===============================================================
// 子查询，总是与字符串条件一起使用。
// ===============================================================
// md5:3cd7047ec77cba30
# <翻译结束>


<原文开始>
// Automatically adding the brackets.
<原文结束>

# <翻译开始>
// 自动添加括号。. md5:4b202bb8e8a55e8b
# <翻译结束>


<原文开始>
// Eg: Where/And/Or("uid>=", 1)
<原文结束>

# <翻译开始>
// 示例：Where/And/Or("uid>=", 1) 
// 
// 翻译为中文：
// 
// 例如：Where/And/Or("uid>=", 1). md5:97d0042da730e39a
# <翻译结束>


<原文开始>
						// Eg:
						// Where("id", []int{1,2,3})
						// Where("user.id", []int{1,2,3})
<原文结束>

# <翻译开始>
// 例如：
// Where("id", []int{1,2,3})
// Where("user.id", []int{1,2,3})
// 
// 这段代码的注释表示示例用法，其中`Where`是一个函数，它接受两个参数：一个是要查询的字段（如"id"或"user.id"），另一个是一组值（如包含1, 2, 3的整数切片）。这通常用于在数据库查询中设置条件，比如筛选id为1, 2或3的记录，或者在"user"表中的"id"字段匹配这些值。
// md5:5688161e5a37e690
# <翻译结束>


<原文开始>
						// Eg:
						// Where("id", nil)
						// Where("user.id", nil)
<原文结束>

# <翻译开始>
// 例如：
// Where("id", nil) // 指定 "id" 字段的查询条件为 nil（空）
// Where("user.id", nil) // 指定 "user.id" 字段的查询条件为 nil（空）
// md5:7b874349f1af2dd8
# <翻译结束>


<原文开始>
						// Eg:
						// Where/And/Or("uid", 1)
						// Where/And/Or("user.uid", 1)
<原文结束>

# <翻译开始>
// 例如：
// Where/And/Or("uid", 1) // 指定 "uid" 字段等于 1 的条件
// Where/And/Or("user.uid", 1) // 指定 "user" 对象下的 "uid" 字段等于 1 的条件
// md5:0809c46d1c195714
# <翻译结束>


<原文开始>
// formatWhereInterfaces formats `where` as []interface{}.
<原文结束>

# <翻译开始>
// formatWhereInterfaces 将 `where` 格式化为 []interface{}。. md5:6fb34f9561771cdc
# <翻译结束>


<原文开始>
// Db is the underlying DB object for current operation.
<原文结束>

# <翻译开始>
// Db 是当前操作的底层数据库对象。. md5:f4b4c46633cb4235
# <翻译结束>


<原文开始>
// Buffer is the sql statement string without Args for current operation.
<原文结束>

# <翻译开始>
// Buffer是当前操作的SQL语句字符串，不包含Args。. md5:fc54d627bfb62054
# <翻译结束>


<原文开始>
// Args is the full arguments of current operation.
<原文结束>

# <翻译开始>
// Args是当前操作的全部参数。. md5:e962690161726419
# <翻译结束>


<原文开始>
// The field name, eg: "id", "name", etc.
<原文结束>

# <翻译开始>
// 字段名称，例如："id"，"name"等。. md5:26a2c4cbd9f18aa7
# <翻译结束>


<原文开始>
// The field value, can be any types.
<原文结束>

# <翻译开始>
// 字段值，可以是任何类型。. md5:1edf6819770a85b8
# <翻译结束>


<原文开始>
// The value in Where type.
<原文结束>

# <翻译开始>
// Where类型中的值。. md5:8becc6ca3981308b
# <翻译结束>


<原文开始>
// Ignores current condition key if `value` is empty.
<原文结束>

# <翻译开始>
// 如果`value`为空，忽略当前条件键。. md5:a7b83f4b09b6f499
# <翻译结束>


<原文开始>
// Field prefix, eg: "user", "order", etc.
<原文结束>

# <翻译开始>
// 字段前缀，例如："用户"，"订单"等。. md5:53032e7ee552fd8f
# <翻译结束>


<原文开始>
// formatWhereKeyValue handles each key-value pair of the parameter map.
<原文结束>

# <翻译开始>
// formatWhereKeyValue 处理参数映射中的每一组键值对。. md5:5d6f9d3dee346d1a
# <翻译结束>


<原文开始>
	// If the value is type of slice, and there's only one '?' holder in
	// the key string, it automatically adds '?' holder chars according to its arguments count
	// and converts it to "IN" statement.
<原文结束>

# <翻译开始>
// 如果值是切片类型，并且键字符串中只有一个'?'占位符，它会根据参数数量自动添加占位符字符，并将其转换为"In"语句。
// md5:10f5c168c92db7c7
# <翻译结束>


<原文开始>
// The key is a single field name.
<原文结束>

# <翻译开始>
// 键是一个单个字段名。. md5:637e2145bd0e0f57
# <翻译结束>


<原文开始>
// The key may have operation chars.
<原文结束>

# <翻译开始>
// 键可能包含操作字符。. md5:1e21edc23189eeb0
# <翻译结束>


<原文开始>
// It also supports "LIKE" statement, which we consider it an operator.
<原文结束>

# <翻译开始>
// 它还支持"LIKE"语句，我们将其视为一种运算符。. md5:71ae1896f0d3b4fd
# <翻译结束>


<原文开始>
// Eg: Where(g.Map{"name like": "john%"})
<原文结束>

# <翻译开始>
// 例如：Where(g.Map{"name like": "john%"}) 
// 
// 这段Go语言的注释表示一个示例用法，其中`Where`是一个函数，它接受一个映射（Map）作为参数，这个映射的键值对是`"name like"` 和 `"john%"`。这意味着在查询时，将对"name"字段进行模糊匹配，查找以"john"开头的记录。. md5:a6037088e14ea97a
# <翻译结束>


<原文开始>
// Eg: Where(g.Map{"age > ": 16})
<原文结束>

# <翻译开始>
// 例如：Where(g.Map{"age > ": 16}) 
// 
// 这段Go语言代码的注释表示这是一个示例（Eg），它使用了一个谓词（Where）和一个映射（Map），这个映射中键值对为 "age > " : 16，意思是筛选出年龄大于16的项。. md5:2b3b5668547eafe7
# <翻译结束>


<原文开始>
// The key is a regular field name.
<原文结束>

# <翻译开始>
// 键是一个常规的字段名。. md5:6088ec7a69f84698
# <翻译结束>


<原文开始>
					// The key is not a regular field name.
					// Eg: Where(g.Map{"age > 16": nil})
					// Issue: https://github.com/gogf/gf/issues/765
<原文结束>

# <翻译开始>
// 键不是一个常规的字段名。
// 例如：Where(g.Map{"age > 16": nil})
// 问题链接：https://github.com/gogf/gf/issues/765
// md5:79107f0b28e8b612
# <翻译结束>


<原文开始>
// handleSliceAndStructArgsForSql is an important function, which handles the sql and all its arguments
// before committing them to underlying driver.
<原文结束>

# <翻译开始>
// handleSliceAndStructArgsForSql 是一个重要的函数，它在将 sql 和其所有参数提交给底层驱动程序之前处理它们。
// md5:a6e05a5f78b51a2b
# <翻译结束>


<原文开始>
// insertHolderCount is used to calculate the inserting position for the '?' holder.
<原文结束>

# <翻译开始>
// insertHolderCount 用于计算 "?" 持有者的插入位置。. md5:878313cc9b1fa5c3
# <翻译结束>


<原文开始>
// Handles the slice and struct type argument item.
<原文结束>

# <翻译开始>
// 处理切片和结构体类型的参数项。. md5:ce55342863e73d8f
# <翻译结束>


<原文开始>
			// It does not split the type of []byte.
			// Eg: table.Where("name = ?", []byte("john"))
<原文结束>

# <翻译开始>
// 它不会分割 []byte 类型。
// 例如：table.Where("name = ?", []byte("john"))
// md5:05dcc823e289de42
# <翻译结束>


<原文开始>
				// Empty slice argument, it converts the sql to a false sql.
				// Example:
				// Query("select * from xxx where id in(?)", g.Slice{}) -> select * from xxx where 0=1
				// Where("id in(?)", g.Slice{}) -> WHERE 0=1
<原文结束>

# <翻译开始>
// 空切片参数，它将SQL转换为一个假的SQL。
// 示例：
// Query("select * from xxx where id in(?)", g.Slice{}) -> select * from xxx where 0=1
// Where("id in(?)", g.Slice{}) -> WHERE 0=1
// 
// 这里的注释说明了当使用空切片（`g.Slice{}`）作为参数时，Go的某些函数（如`Query`和`Where`）会将SQL中的条件改变为等价于`false`的形式，例如将`in`条件替换为`0=1`，从而使得查询结果为空。
// md5:020597c0f38437e4
# <翻译结束>


<原文开始>
				// Example:
				// Query("SELECT ?+?", g.Slice{1,2})
				// WHERE("id=?", g.Slice{1,2})
<原文结束>

# <翻译开始>
// 示例：
// Query("SELECT ?+?", g.Slice{1,2}) // 查询语句，参数为1和2
// WHERE("id=?", g.Slice{1,2}) // WHERE子句，参数为1和2
// md5:4f9ae718d40ffb8b
# <翻译结束>


<原文开始>
			// If the '?' holder count equals the length of the slice,
			// it does not implement the arguments splitting logic.
			// Eg: db.Query("SELECT ?+?", g.Slice{1, 2})
<原文结束>

# <翻译开始>
// 如果 "?" 占位符的数量等于切片的长度，
// 则不执行参数分割的逻辑。
// 例如：db.Query("SELECT ?+?", g.Slice{1, 2})
// md5:aac31c8c27bdcf7d
# <翻译结束>


<原文开始>
// counter is used to finding the inserting position for the '?' holder.
<原文结束>

# <翻译开始>
// counter 用于查找 '?' 占位符的插入位置。. md5:22bff4ac2bdd0f47
# <翻译结束>


<原文开始>
// Special struct handling.
<原文结束>

# <翻译开始>
// 特殊的结构体处理。. md5:8911bc2424fd10eb
# <翻译结束>


<原文开始>
// The underlying driver supports time.Time/*time.Time types.
<原文结束>

# <翻译开始>
// 基础驱动程序支持time.Time类型。. md5:9143055892307413
# <翻译结束>


<原文开始>
				// It converts the struct to string in default
				// if it has implemented the String interface.
<原文结束>

# <翻译开始>
// 如果结构体实现了String接口，它将默认将结构体转换为字符串。
// md5:59ba6afad009bc6a
# <翻译结束>


<原文开始>
// FormatSqlWithArgs binds the arguments to the sql string and returns a complete
// sql string, just for debugging.
<原文结束>

# <翻译开始>
// FormatSqlWithArgs 将参数绑定到SQL字符串并返回一个完整的SQL字符串，仅用于调试。
// md5:1453466956e418ba
# <翻译结束>


<原文开始>
// Parameters of type Raw do not require special treatment
<原文结束>

# <翻译开始>
// 类型为Raw的参数不需要特殊处理. md5:aa477b3ebc58b939
# <翻译结束>


<原文开始>
// FormatMultiLineSqlToSingle formats sql template string into one line.
<原文结束>

# <翻译开始>
// FormatMultiLineSqlToSingle 将多行SQL模板字符串格式化为单行。. md5:cb1180487fd5c495
# <翻译结束>


<原文开始>
// format sql template string.
<原文结束>

# <翻译开始>
// 格式化SQL模板字符串。. md5:77bcc0fc1c095ebd
# <翻译结束>

