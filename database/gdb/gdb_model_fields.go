// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

import (
	"fmt"
	
	"github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
)

// Fields 用于指定需要操作的表字段，包括查询字段、写入字段、更新字段等过滤；
// 参数 `fieldNamesOrMapStruct` 可以是 string/map/*map/struct/*struct 类型。
// 多个字段之间使用字符 ',' 连接。
//
// 查询过滤示例: 
// g.Model("user").Fields("uid, nickname").Order("uid asc").All()  //SELECT `uid`,`nickname` FROM `user` ORDER BY `uid` asc
//
// 写入过滤示例:
// m := g.Map{
// "uid"      : 10000,
// "nickname" : "John Guo",
// "passport" : "john",
// "password" : "123456",
// }
// g.Model(table).Fields("nickname,passport,password").Data(m).Insert()  //INSERT INTO `user`(`nickname`,`passport`,`password`) VALUES('John Guo','john','123456')
//
// 示例：
// Fields("id", "name", "age")    // 通过字符串直接指定字段名
// Fields([]string{"id", "name", "age"})   // 通过字符串切片指定字段名
// Fields(map[string]interface{}{"id":1, "name":"john", "age":18})  // 通过键值对映射指定字段和值
// Fields(User{ Id: 1, Name: "john", Age: 18})   // 通过结构体实例指定字段和值
func (m *Model) X字段保留过滤(字段名或Map结构体 ...interface{}) *Model {
	length := len(字段名或Map结构体)
	if length == 0 {
		return m
	}
	fields := m.getFieldsFrom(m.tablesInit, 字段名或Map结构体...)
	if len(fields) == 0 {
		return m
	}
	return m.appendFieldsByStr(文本类.X连接(fields, ","))
}

// FieldsPrefix 函数的功能与 Fields 相同，但会在每个字段前额外添加一个前缀。
func (m *Model) X字段保留过滤并带前缀(前缀或别名 string, 字段名或Map结构体 ...interface{}) *Model {
	fields := m.getFieldsFrom(m.getTableNameByPrefixOrAlias(前缀或别名), 字段名或Map结构体...)
	if len(fields) == 0 {
		return m
	}
	文本类.X数组加前缀(fields, 前缀或别名+".")
	return m.appendFieldsByStr(文本类.X连接(fields, ","))
}

// FieldsEx 将`fieldNamesOrMapStruct` 追加到模型的排除操作字段列表中，
// 多个字段之间使用逗号字符 ',' 连接。
// 注意，此函数仅支持单表操作。
// 参数 `fieldNamesOrMapStruct` 可以是 string、map、*map 或 struct、*struct 类型。
// 请同时参考 Fields 函数。
//
// 查询排除过滤例子
// g.Model("user").FieldsEx("passport, password").All()  //SELECT `uid`,`nickname` FROM `user`
//
// 写入排除过滤例子
// m := g.Map{
// "uid"      : 10000,
// "nickname" : "John Guo",
// "passport" : "john",
// "password" : "123456",
// }
// g.Model(table).FieldsEx("uid").Data(m).Insert()  // INSERT INTO `user`(`nickname`,`passport`,`password`) VALUES('John Guo','john','123456')
func (m *Model) X字段排除过滤(字段名或Map结构体 ...interface{}) *Model {
	return m.doFieldsEx(m.tablesInit, 字段名或Map结构体...)
}
func (m *Model) doFieldsEx(table string, fieldNamesOrMapStruct ...interface{}) *Model {
	length := len(fieldNamesOrMapStruct)
	if length == 0 {
		return m
	}
	fields := m.getFieldsFrom(table, fieldNamesOrMapStruct...)
	if len(fields) == 0 {
		return m
	}
	return m.appendFieldsExByStr(文本类.X连接(fields, ","))
}

// FieldsExPrefix 函数的功能与 FieldsEx 相同，但会在每个字段前额外添加一个前缀。
func (m *Model) X字段排除过滤并带前缀(前缀或别名 string, 字段名或Map结构体 ...interface{}) *Model {
	model := m.doFieldsEx(m.getTableNameByPrefixOrAlias(前缀或别名), 字段名或Map结构体...)
	array := 文本类.X分割并忽略空值(model.fieldsEx, ",")
	文本类.X数组加前缀(array, 前缀或别名+".")
	model.fieldsEx = 文本类.X连接(array, ",")
	return model
}

// FieldCount 格式化并追加计数字段别名到模型的 select 字段中。
// 简单点说就是追加一个计数的别名字段
//
// 追加计数字段例子:
// db.Model(table).Fields("id").FieldCount("id", "total")  // COUNT(`id`) AS `total`
func (m *Model) X字段追加计数(需要计数的字段名称 string, 新字段别名 ...string) *Model {
	asStr := ""
	if len(新字段别名) > 0 && 新字段别名[0] != "" {
		asStr = fmt.Sprintf(` AS %s`, m.db.X取Core对象().X底层QuoteWord(新字段别名[0]))
	}
	return m.appendFieldsByStr(fmt.Sprintf(`COUNT(%s)%s`, m.底层QuoteWord(需要计数的字段名称), asStr))
}

// FieldSum 格式化并追加常用字段 `SUM(column)` 到模型的 select 字段中。
// 简单点说就是追加一个求和的别名字段
//
// 追加求和字段例子:
// db.Model(table).Fields("column").FieldSum("column", "total")  // SUM(`column`) AS `total`
func (m *Model) X字段追加求和(需要求和的字段名称 string, 新字段别名 ...string) *Model {
	asStr := ""
	if len(新字段别名) > 0 && 新字段别名[0] != "" {
		asStr = fmt.Sprintf(` AS %s`, m.db.X取Core对象().X底层QuoteWord(新字段别名[0]))
	}
	return m.appendFieldsByStr(fmt.Sprintf(`SUM(%s)%s`, m.底层QuoteWord(需要求和的字段名称), asStr))
}

// FieldMin 格式化并追加常用字段 `MIN(column)` 到模型的 select 字段中。
// 简单点说就是追加一个最小值的别名字段
//
// 追加最小值字段例子:
// db.Model(table).Fields("column").FieldMin("column", "total")  // MIN(`column`) AS `total`
func (m *Model) X字段追加最小值(最小值字段名称 string, 新字段别名 ...string) *Model {
	asStr := ""
	if len(新字段别名) > 0 && 新字段别名[0] != "" {
		asStr = fmt.Sprintf(` AS %s`, m.db.X取Core对象().X底层QuoteWord(新字段别名[0]))
	}
	return m.appendFieldsByStr(fmt.Sprintf(`MIN(%s)%s`, m.底层QuoteWord(最小值字段名称), asStr))
}

// FieldMax 格式化并追加常用字段 `MAX(column)` 到模型的 select 字段中。
// 简单点说就是追加一个最大值的别名字段
//
// 追加最大值字段例子:
// db.Model(table).Fields("column").FieldMax("column", "total")  // MAX(`column`) AS `total`
func (m *Model) X字段追加最大值(最大值字段名称 string, 新字段别名 ...string) *Model {
	asStr := ""
	if len(新字段别名) > 0 && 新字段别名[0] != "" {
		asStr = fmt.Sprintf(` AS %s`, m.db.X取Core对象().X底层QuoteWord(新字段别名[0]))
	}
	return m.appendFieldsByStr(fmt.Sprintf(`MAX(%s)%s`, m.底层QuoteWord(最大值字段名称), asStr))
}

// FieldAvg 格式化并追加常用字段 `AVG(column)` 到模型的 select 字段中。
// 简单点说就是追加一个平均值的别名字段
//
// 追加平均值字段例子:
// db.Model(table).Fields("column").FieldAvg("column", "total")  // AVG(`column`) AS `total`
func (m *Model) X字段追加平均值(求平均值字段名称 string, 新字段别名 ...string) *Model {
	asStr := ""
	if len(新字段别名) > 0 && 新字段别名[0] != "" {
		asStr = fmt.Sprintf(` AS %s`, m.db.X取Core对象().X底层QuoteWord(新字段别名[0]))
	}
	return m.appendFieldsByStr(fmt.Sprintf(`AVG(%s)%s`, m.底层QuoteWord(求平均值字段名称), asStr))
}

// GetFieldsStr 函数从表中检索并返回所有字段，各字段之间用字符 ',' 连接。
// 可选参数 `prefix` 用于指定每个字段的前缀，例如：GetFieldsStr("u.").
func (m *Model) X取所有字段名称(字段前缀 ...string) string {
	prefixStr := ""
	if len(字段前缀) > 0 {
		prefixStr = 字段前缀[0]
	}
	tableFields, err := m.X取表字段信息Map(m.tablesInit)
	if err != nil {
		panic(err)
	}
	if len(tableFields) == 0 {
		panic(fmt.Sprintf(`empty table fields for table "%s"`, m.tables))
	}
	fieldsArray := make([]string, len(tableFields))
	for k, v := range tableFields {
		fieldsArray[v.X排序] = k
	}
	newFields := ""
	for _, k := range fieldsArray {
		if len(newFields) > 0 {
			newFields += ","
		}
		newFields += prefixStr + k
	}
	newFields = m.db.X取Core对象().X底层QuoteString(newFields)
	return newFields
}

// GetFieldsExStr 从表中检索并返回不在参数 `fields` 中的字段，并使用字符 ',' 连接这些字段。
// 参数 `fields` 指定需要排除的字段。
// 可选参数 `prefix` 用于指定每个字段的前缀，例如：FieldsExStr("id", "u.")。
func (m *Model) X取所有字段名称并排除(需要排除字段 string, 字段前缀 ...string) string {
	prefixStr := ""
	if len(字段前缀) > 0 {
		prefixStr = 字段前缀[0]
	}
	tableFields, err := m.X取表字段信息Map(m.tablesInit)
	if err != nil {
		panic(err)
	}
	if len(tableFields) == 0 {
		panic(fmt.Sprintf(`empty table fields for table "%s"`, m.tables))
	}
	fieldsExSet := 集合类.X创建文本并按值(文本类.X分割并忽略空值(需要排除字段, ","))
	fieldsArray := make([]string, len(tableFields))
	for k, v := range tableFields {
		fieldsArray[v.X排序] = k
	}
	newFields := ""
	for _, k := range fieldsArray {
		if fieldsExSet.X是否存在(k) {
			continue
		}
		if len(newFields) > 0 {
			newFields += ","
		}
		newFields += prefixStr + k
	}
	newFields = m.db.X取Core对象().X底层QuoteString(newFields)
	return newFields
}

// HasField 判断字段是否在表中存在。
func (m *Model) X是否存在字段(字段名称 string) (bool, error) {
	return m.db.X取Core对象().X是否存在字段(m.X取上下文对象(), m.tablesInit, 字段名称)
}

// getFieldsFrom 从表 `table` 中检索、过滤并返回字段名称。
func (m *Model) getFieldsFrom(table string, fieldNamesOrMapStruct ...interface{}) []string {
	length := len(fieldNamesOrMapStruct)
	if length == 0 {
		return nil
	}
	switch {
	// String slice.
	case length >= 2:
		return m.mappingAndFilterToTableFields(
			table, 转换类.X取文本数组(fieldNamesOrMapStruct), true,
		)

	// 需要进行类型断言。
	case length == 1:
		structOrMap := fieldNamesOrMapStruct[0]
		switch r := structOrMap.(type) {
		case string:
			return m.mappingAndFilterToTableFields(table, []string{r}, false)

		case []string:
			return m.mappingAndFilterToTableFields(table, r, true)

		case X原生sql, *X原生sql:
			return []string{转换类.String(structOrMap)}

		default:
			return m.mappingAndFilterToTableFields(table, getFieldsFromStructOrMap(structOrMap), true)
		}

	default:
		return nil
	}
}

func (m *Model) appendFieldsByStr(fields string) *Model {
	if fields != "" {
		model := m.getModel()
		if model.fields == defaultFields {
			model.fields = ""
		}
		if model.fields != "" {
			model.fields += ","
		}
		model.fields += fields
		return model
	}
	return m
}

func (m *Model) appendFieldsExByStr(fieldsEx string) *Model {
	if fieldsEx != "" {
		model := m.getModel()
		if model.fieldsEx != "" {
			model.fieldsEx += ","
		}
		model.fieldsEx += fieldsEx
		return model
	}
	return m
}
