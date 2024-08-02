// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

import (
	"time"

	gset "github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/internal/empty"
	gtime "github.com/888go/goframe/os/gtime"
	gregex "github.com/888go/goframe/text/gregex"
	gstr "github.com/888go/goframe/text/gstr"
	gutil "github.com/888go/goframe/util/gutil"
)

// QuoteWord 检查给定的字符串 `s` 是否为一个单词，
// 如果是，它将使用数据库的安全字符对 `s` 进行转义，并返回带引号的字符串；否则，返回原始字符串不做任何更改。
//
// 可以认为一个 `word` 表示列名。
// md5:71291615d7bcffe0
func (m *Model) QuoteWord(s string) string {
	return m.db.GetCore().QuoteWord(s)
}

// TableFields 获取并返回当前模式下指定表的字段信息。
//
// 参见 DriverMysql.TableFields。
// md5:61e256ba53f813cb
func (m *Model) TableFields(tableStr string, schema ...string) (fields map[string]*TableField, err error) {
	var (
		table      = m.db.GetCore().guessPrimaryTableName(tableStr)
		usedSchema = gutil.GetOrDefaultStr(m.schema, schema...)
	)
	return m.db.TableFields(m.GetCtx(), table, usedSchema)
}

// getModel 如果`safe`为真，则创建并返回当前模型的克隆，否则直接返回当前模型。
// md5:e4ae726aba6b01ab
func (m *Model) getModel() *Model {
	if !m.safe {
		return m
	} else {
		return m.Clone()
	}
}

// mappingAndFilterToTableFields：将给定的字段名映射并转换为实际的表格字段名。
// 例如：
// ID        -> id
// NICK_Name -> nickname.
// md5:35f1e9dc3d13c4f0
func (m *Model) mappingAndFilterToTableFields(table string, fields []string, filter bool) []string {
	var fieldsTable = table
	if fieldsTable != "" {
		hasTable, _ := m.db.GetCore().HasTable(fieldsTable)
		if !hasTable {
			fieldsTable = m.tablesInit
		}
	}
	if fieldsTable == "" {
		fieldsTable = m.tablesInit
	}

	fieldsMap, _ := m.TableFields(fieldsTable)
	if len(fieldsMap) == 0 {
		return fields
	}
	var outputFieldsArray = make([]string, 0)
	fieldsKeyMap := make(map[string]interface{}, len(fieldsMap))
	for k := range fieldsMap {
		fieldsKeyMap[k] = nil
	}
	for _, field := range fields {
		var inputFieldsArray []string
		if gregex.IsMatchString(regularFieldNameWithoutDotRegPattern, field) {
			inputFieldsArray = append(inputFieldsArray, field)
		} else if gregex.IsMatchString(regularFieldNameWithCommaRegPattern, field) {
			inputFieldsArray = gstr.SplitAndTrim(field, ",")
		} else {
			// 示例：
			// user.id, user.name
			// 将逗号分隔的字符串（格式：lpad(s.id, 6, '0')，s.name）替换为`code`
			// md5:5ee6374c41194bf3
			outputFieldsArray = append(outputFieldsArray, field)
			continue
		}
		for _, inputField := range inputFieldsArray {
			if !gregex.IsMatchString(regularFieldNameWithoutDotRegPattern, inputField) {
				outputFieldsArray = append(outputFieldsArray, inputField)
				continue
			}
			if _, ok := fieldsKeyMap[inputField]; !ok {
				// 示例：
				// id, 名称
				// md5:f16c15c62075a7aa
				if foundKey, _ := gutil.MapPossibleItemByKey(fieldsKeyMap, inputField); foundKey != "" {
					outputFieldsArray = append(outputFieldsArray, foundKey)
				} else if !filter {
					outputFieldsArray = append(outputFieldsArray, inputField)
				}
			} else {
				outputFieldsArray = append(outputFieldsArray, inputField)
			}
		}
	}
	return outputFieldsArray
}

// filterDataForInsertOrUpdate 对用于插入/更新操作的数据执行过滤功能。
// 请注意，它不会对列表项（也是一种映射类型）进行“忽略空值”处理。
// md5:ffc8a604eaec8a77
func (m *Model) filterDataForInsertOrUpdate(data interface{}) (interface{}, error) {
	var err error
	switch value := data.(type) {
	case List:
		var omitEmpty bool
		if m.option&optionOmitNilDataList > 0 {
			omitEmpty = true
		}
		for k, item := range value {
			value[k], err = m.doMappingAndFilterForInsertOrUpdateDataMap(item, omitEmpty)
			if err != nil {
				return nil, err
			}
		}
		return value, nil

	case Map:
		return m.doMappingAndFilterForInsertOrUpdateDataMap(value, true)

	default:
		return data, nil
	}
}

// doMappingAndFilterForInsertOrUpdateDataMap 为映射类型的数据执行过滤功能。
// 注意，它不会对"忽略空"特性下的列表项（也是映射类型）进行过滤。
// md5:93fefbe3176f55de
func (m *Model) doMappingAndFilterForInsertOrUpdateDataMap(data Map, allowOmitEmpty bool) (Map, error) {
	var err error
	data, err = m.db.GetCore().mappingAndFilterData(
		m.GetCtx(), m.schema, m.tablesInit, data, m.filter,
	)
	if err != nil {
		return nil, err
	}
		// 删除值为nil的键值对。 md5:5219c3473c86d38c
	if allowOmitEmpty && m.option&optionOmitNilData > 0 {
		tempMap := make(Map, len(data))
		for k, v := range data {
			if empty.IsNil(v) {
				continue
			}
			tempMap[k] = v
		}
		data = tempMap
	}

		// 删除值为空的键值对。 md5:706fbf04684a1301
	if allowOmitEmpty && m.option&optionOmitEmptyData > 0 {
		tempMap := make(Map, len(data))
		for k, v := range data {
			if empty.IsEmpty(v) {
				continue
			}
						// 特殊类型的过滤。 md5:48598cd9e3395cfc
			switch r := v.(type) {
			case time.Time:
				if r.IsZero() {
					continue
				}
			case *time.Time:
				if r.IsZero() {
					continue
				}
			case gtime.Time:
				if r.IsZero() {
					continue
				}
			case *gtime.Time:
				if r.IsZero() {
					continue
				}
			}
			tempMap[k] = v
		}
		data = tempMap
	}

	if len(m.fields) > 0 && m.fields != "*" {
		// Keep specified fields.
		var (
			set          = gset.NewStrSetFrom(gstr.SplitAndTrim(m.fields, ","))
			charL, charR = m.db.GetChars()
			chars        = charL + charR
		)
		set.Walk(func(item string) string {
			return gstr.Trim(item, chars)
		})
		for k := range data {
			k = gstr.Trim(k, chars)
			if !set.Contains(k) {
				delete(data, k)
			}
		}
	} else if len(m.fieldsEx) > 0 {
				// 过滤指定字段。 md5:c1817e5f938542f0
		for _, v := range gstr.SplitAndTrim(m.fieldsEx, ",") {
			delete(data, v)
		}
	}
	return data, nil
}

// getLink 函数返回配置了 `linkType` 属性的底层数据库链接对象。
// 参数 `master` 指定是否在主从配置中使用主节点。
// md5:e8add2f9371393db
func (m *Model) getLink(master bool) Link {
	if m.tx != nil {
		return &txLink{m.tx.GetSqlTX()}
	}
	linkType := m.linkType
	if linkType == 0 {
		if master {
			linkType = linkTypeMaster
		} else {
			linkType = linkTypeSlave
		}
	}
	switch linkType {
	case linkTypeMaster:
		link, err := m.db.GetCore().MasterLink(m.schema)
		if err != nil {
			panic(err)
		}
		return link
	case linkTypeSlave:
		link, err := m.db.GetCore().SlaveLink(m.schema)
		if err != nil {
			panic(err)
		}
		return link
	}
	return nil
}

// getPrimaryKey 获取并返回模型表的主键名称。
// 它解析 m.tables 以获取主表名，支持如下的 m.tables 写法：
// "user", "user u", "user as u, user_detail as ud"。
// md5:07ea92a426e953d1
func (m *Model) getPrimaryKey() string {
	table := gstr.SplitAndTrim(m.tablesInit, " ")[0]
	tableFields, err := m.TableFields(table)
	if err != nil {
		return ""
	}
	for name, field := range tableFields {
		if gstr.ContainsI(field.Key, "pri") {
			return name
		}
	}
	return ""
}

// mergeArguments 将 `m.extraArgs` 和给定的 `args` 合并，创建并返回新的参数。 md5:80f949384113727a
func (m *Model) mergeArguments(args []interface{}) []interface{} {
	if len(m.extraArgs) > 0 {
		newArgs := make([]interface{}, len(m.extraArgs)+len(args))
		copy(newArgs, m.extraArgs)
		copy(newArgs[len(m.extraArgs):], args)
		return newArgs
	}
	return args
}
