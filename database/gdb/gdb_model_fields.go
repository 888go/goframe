// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

import (
	"fmt"

	gset "github.com/888go/goframe/container/gset"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
)

// `X字段保留过滤` 方法将 `fieldNamesOrMapStruct` 添加到模型的operation字段中，多个字段使用字符`,`连接。
// 参数 `fieldNamesOrMapStruct` 可以是字符串、映射（map）、*map、结构体或*结构体类型。
// 
// 示例：
// X字段保留过滤("id", "name", "age")
// X字段保留过滤([]string{"id", "name", "age"})
// X字段保留过滤(map[string]interface{}{"id":1, "name":"john", "age":18})
// X字段保留过滤(User{ Id: 1, Name: "john", Age: 18})
// md5:21db86fe96babad2
func (m *Model) X字段保留过滤(字段名或Map结构体 ...interface{}) *Model {
	length := len(字段名或Map结构体)
	if length == 0 {
		return m
	}
	fields := m.getFieldsFrom(m.tablesInit, 字段名或Map结构体...)
	if len(fields) == 0 {
		return m
	}
	return m.appendFieldsByStr(gstr.X连接(fields, ","))
}

// X字段保留过滤并带前缀 作为 Fields 函数，但为每个字段添加额外的前缀。 md5:8a672048e8753526
func (m *Model) X字段保留过滤并带前缀(前缀或别名 string, 字段名或Map结构体 ...interface{}) *Model {
	fields := m.getFieldsFrom(m.getTableNameByPrefixOrAlias(前缀或别名), 字段名或Map结构体...)
	if len(fields) == 0 {
		return m
	}
	gstr.X切片加前缀(fields, 前缀或别名+".")
	return m.appendFieldsByStr(gstr.X连接(fields, ","))
}

// X字段排除过滤 将 `fieldNamesOrMapStruct` 追加到模型的操作排除字段中，
// 多个字段使用逗号 ',' 连接。
// 注意，此函数仅支持单表操作。
// 参数 `fieldNamesOrMapStruct` 可以是字符串类型、映射类型（map）、映射指针类型（*map）、结构体类型或结构体指针类型（*struct）。
//
// 参见 Fields。
// md5:7b8ec243202549a8
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
	return m.appendFieldsExByStr(gstr.X连接(fields, ","))
}

// X字段排除过滤并带前缀 函数与 FieldsEx 功能相似，但在每个字段前添加额外的前缀。 md5:66ba7040b83e6e81
func (m *Model) X字段排除过滤并带前缀(前缀或别名 string, 字段名或Map结构体 ...interface{}) *Model {
	model := m.doFieldsEx(m.getTableNameByPrefixOrAlias(前缀或别名), 字段名或Map结构体...)
	array := gstr.X分割并忽略空值(model.fieldsEx, ",")
	gstr.X切片加前缀(array, 前缀或别名+".")
	model.fieldsEx = gstr.X连接(array, ",")
	return model
}

// X字段追加计数 将常用字段 `COUNT(column)` 格式化并添加到模型的 select 字段中。 md5:99439830c058a91f
func (m *Model) X字段追加计数(需要计数的字段名称 string, 新字段别名 ...string) *Model {
	asStr := ""
	if len(新字段别名) > 0 && 新字段别名[0] != "" {
		asStr = fmt.Sprintf(` AS %s`, m.db.X取Core对象().X底层QuoteWord(新字段别名[0]))
	}
	return m.appendFieldsByStr(fmt.Sprintf(`COUNT(%s)%s`, m.X底层QuoteWord(需要计数的字段名称), asStr))
}

// X字段追加求和 将常用字段 `SUM(column)` 格式化后添加到模型的 select 字段中。 md5:938249bb2923fa1f
func (m *Model) X字段追加求和(需要求和的字段名称 string, 新字段别名 ...string) *Model {
	asStr := ""
	if len(新字段别名) > 0 && 新字段别名[0] != "" {
		asStr = fmt.Sprintf(` AS %s`, m.db.X取Core对象().X底层QuoteWord(新字段别名[0]))
	}
	return m.appendFieldsByStr(fmt.Sprintf(`SUM(%s)%s`, m.X底层QuoteWord(需要求和的字段名称), asStr))
}

// X字段追加最小值 格式化并追加常用的字段 `MIN(column)` 到模型的选择字段中。 md5:fd1204ad66608451
func (m *Model) X字段追加最小值(最小值字段名称 string, 新字段别名 ...string) *Model {
	asStr := ""
	if len(新字段别名) > 0 && 新字段别名[0] != "" {
		asStr = fmt.Sprintf(` AS %s`, m.db.X取Core对象().X底层QuoteWord(新字段别名[0]))
	}
	return m.appendFieldsByStr(fmt.Sprintf(`MIN(%s)%s`, m.X底层QuoteWord(最小值字段名称), asStr))
}

// X字段追加最大值 格式化并追加常用的字段 `MAX(column)` 到模型的选择字段中。 md5:77150e433b0d44c4
func (m *Model) X字段追加最大值(最大值字段名称 string, 新字段别名 ...string) *Model {
	asStr := ""
	if len(新字段别名) > 0 && 新字段别名[0] != "" {
		asStr = fmt.Sprintf(` AS %s`, m.db.X取Core对象().X底层QuoteWord(新字段别名[0]))
	}
	return m.appendFieldsByStr(fmt.Sprintf(`MAX(%s)%s`, m.X底层QuoteWord(最大值字段名称), asStr))
}

// X字段追加平均值 将常用字段 `AVG(column)` 格式化并添加到模型的 select 字段中。 md5:0b09ffae1b0cbabe
func (m *Model) X字段追加平均值(求平均值字段名称 string, 新字段别名 ...string) *Model {
	asStr := ""
	if len(新字段别名) > 0 && 新字段别名[0] != "" {
		asStr = fmt.Sprintf(` AS %s`, m.db.X取Core对象().X底层QuoteWord(新字段别名[0]))
	}
	return m.appendFieldsByStr(fmt.Sprintf(`AVG(%s)%s`, m.X底层QuoteWord(求平均值字段名称), asStr))
}

// X取所有字段名称 从表中检索并返回所有字段，以逗号分隔。
// 可选参数 `prefix` 指定每个字段的前缀，例如：X取所有字段名称("u.")。
// md5:c76f2f45c8680a27
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
		fieldsArray[v.Index] = k
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

// X取所有字段名称并排除 从表中获取并返回那些不在参数`fields`中的字段，这些字段通过逗号','连接。
// 参数`fields`指定了需要排除的字段。
// 可选参数`prefix`为每个字段指定前缀，例如：FieldsExStr("id", "u.")。
// md5:57698a0c43f54ec9
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
	fieldsExSet := gset.X创建文本并按值(gstr.X分割并忽略空值(需要排除字段, ","))
	fieldsArray := make([]string, len(tableFields))
	for k, v := range tableFields {
		fieldsArray[v.Index] = k
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

// X是否存在字段 用于判断该字段是否存在于表中。 md5:e26ad0ecb292096b
func (m *Model) X是否存在字段(字段名称 string) (bool, error) {
	return m.db.X取Core对象().X是否存在字段(m.X取上下文对象(), m.tablesInit, 字段名称)
}

// getFieldsFrom 从表格`table`中获取、过滤并返回字段名。 md5:9a2c6dffbdfe3d24
func (m *Model) getFieldsFrom(table string, fieldNamesOrMapStruct ...interface{}) []string {
	length := len(fieldNamesOrMapStruct)
	if length == 0 {
		return nil
	}
	switch {
	// String slice.
	case length >= 2:
		return m.mappingAndFilterToTableFields(
			table, gconv.X取文本切片(fieldNamesOrMapStruct), true,
		)

		// 需要类型断言。 md5:ec336d143828f70d
	case length == 1:
		structOrMap := fieldNamesOrMapStruct[0]
		switch r := structOrMap.(type) {
		case string:
			return m.mappingAndFilterToTableFields(table, []string{r}, false)

		case []string:
			return m.mappingAndFilterToTableFields(table, r, true)

		case Raw, *Raw:
			return []string{gconv.String(structOrMap)}

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
