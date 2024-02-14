// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

import (
	"time"
	
	"github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/internal/empty"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gutil"
)

// 2024-01-09 改成内部方法,此方法属于底层, 几乎用不到.
// QuoteWord 检查给定字符串 `s` 是否为一个单词，
// 如果是，它会使用数据库的安全字符对 `s` 进行引用，并返回引述后的字符串；
// 否则，它将直接返回未经修改的 `s`。
//
// 这里的“单词”可以理解为列名。
func (m *Model) 底层QuoteWord(s string) string {
	return m.db.X取Core对象().X底层QuoteWord(s)
}

// TableFields 获取并返回当前模式下指定表的字段信息。
//
// 另请参阅 DriverMysql.TableFields。
func (m *Model) X取表字段信息Map(表名称 string, schema ...string) (字段信息Map map[string]*TableField, 错误 error) {
	var (
		table      = m.db.X取Core对象().guessPrimaryTableName(表名称)
		usedSchema = 工具类.X取文本值或取默认值(m.schema, schema...)
	)
	return m.db.X取表字段信息Map(m.X取上下文对象(), table, usedSchema)
}

// getModel函数如果`safe`为真，则创建并返回当前模型的一个克隆副本，否则直接返回当前模型。
func (m *Model) getModel() *Model {
	if !m.safe {
		return m
	} else {
		return m.X取副本()
	}
}

// mappingAndFilterToTableFields 将给定的字段名映射并转换为实际的数据库表字段名。
// 例如：
// ID        -> id
// NICK_Name -> nickname.
func (m *Model) mappingAndFilterToTableFields(table string, fields []string, filter bool) []string {
	var fieldsTable = table
	if fieldsTable != "" {
		hasTable, _ := m.db.X取Core对象().X是否存在表名(fieldsTable)
		if !hasTable {
			fieldsTable = m.tablesInit
		}
	}
	if fieldsTable == "" {
		fieldsTable = m.tablesInit
	}

	fieldsMap, _ := m.X取表字段信息Map(fieldsTable)
	if len(fieldsMap) == 0 {
		return fields
	}
	var (
		inputFieldsArray  = 文本类.X分割并忽略空值(文本类.X连接(fields, ","), ",")
		outputFieldsArray = make([]string, 0, len(inputFieldsArray))
	)
	fieldsKeyMap := make(map[string]interface{}, len(fieldsMap))
	for k := range fieldsMap {
		fieldsKeyMap[k] = nil
	}
	for _, field := range inputFieldsArray {
		if _, ok := fieldsKeyMap[field]; !ok {
			if !正则类.X是否匹配文本(regularFieldNameWithoutDotRegPattern, field) {
				// 示例：user.id, user.name
				outputFieldsArray = append(outputFieldsArray, field)
				continue
			} else {
				// Eg: id, name
				if foundKey, _ := 工具类.MapPossibleItemByKey(fieldsKeyMap, field); foundKey != "" {
					outputFieldsArray = append(outputFieldsArray, foundKey)
				} else if !filter {
					outputFieldsArray = append(outputFieldsArray, field)
				}
			}
		} else {
			outputFieldsArray = append(outputFieldsArray, field)
		}
	}
	return outputFieldsArray
}

// filterDataForInsertOrUpdate 对用于插入/更新操作的数据执行过滤功能。
// 注意，它不针对“忽略空值”特性过滤列表项（该列表项也是映射类型）。
func (m *Model) filterDataForInsertOrUpdate(data interface{}) (interface{}, error) {
	var err error
	switch value := data.(type) {
	case Map数组:
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

// doMappingAndFilterForInsertOrUpdateDataMap 对 map 执行映射和过滤功能。
// 注意，对于“忽略空值”特性，它不会过滤列表项（其类型也为 map）。
func (m *Model) doMappingAndFilterForInsertOrUpdateDataMap(data Map, allowOmitEmpty bool) (Map, error) {
	var err error
	data, err = m.db.X取Core对象().mappingAndFilterData(
		m.X取上下文对象(), m.schema, m.tablesInit, data, m.filter,
	)
	if err != nil {
		return nil, err
	}
	// 删除值为nil的键值对。
	if allowOmitEmpty && m.option&optionOmitNilData > 0 {
		tempMap := make(Map, len(data))
		for k, v := range data {
			if empty.X是否为Nil(v) {
				continue
			}
			tempMap[k] = v
		}
		data = tempMap
	}

	// 删除值为空的键值对。
	if allowOmitEmpty && m.option&optionOmitEmptyData > 0 {
		tempMap := make(Map, len(data))
		for k, v := range data {
			if empty.IsEmpty(v) {
				continue
			}
			// 特殊类型过滤
			switch r := v.(type) {
			case time.Time:
				if r.IsZero() {
					continue
				}
			case *time.Time:
				if r.IsZero() {
					continue
				}
			case 时间类.Time:
				if r.IsZero() {
					continue
				}
			case *时间类.Time:
				if r.IsZero() {
					continue
				}
			}
			tempMap[k] = v
		}
		data = tempMap
	}

	if len(m.fields) > 0 && m.fields != "*" {
		// 保留指定字段。
		var (
			set          = 集合类.X创建文本并按值(文本类.X分割并忽略空值(m.fields, ","))
			charL, charR = m.db.X底层取数据库安全字符()
			chars        = charL + charR
		)
		set.X遍历修改(func(item string) string {
			return 文本类.X过滤首尾符并含空白(item, chars)
		})
		for k := range data {
			k = 文本类.X过滤首尾符并含空白(k, chars)
			if !set.X是否存在(k) {
				delete(data, k)
			}
		}
	} else if len(m.fieldsEx) > 0 {
		// 过滤指定字段。
		for _, v := range 文本类.X分割并忽略空值(m.fieldsEx, ",") {
			delete(data, v)
		}
	}
	return data, nil
}

// getLink 返回配置了 `linkType` 属性的基础数据库连接对象。
// 参数 `master` 指定在主从配置时是否使用主节点。
func (m *Model) getLink(master bool) Link {
	if m.tx != nil {
		return &txLink{m.tx.X底层取事务对象()}
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
		link, err := m.db.X取Core对象().X底层MasterLink(m.schema)
		if err != nil {
			panic(err)
		}
		return link
	case linkTypeSlave:
		link, err := m.db.X取Core对象().X底层SlaveLink(m.schema)
		if err != nil {
			panic(err)
		}
		return link
	}
	return nil
}

// getPrimaryKey 获取并返回模型表的主键名称。
// 它通过解析 m.tables 来检索主表名，支持如下形式的 m.tables：
// "user", "user u", "user as u, user_detail as ud"。
func (m *Model) getPrimaryKey() string {
	table := 文本类.X分割并忽略空值(m.tablesInit, " ")[0]
	tableFields, err := m.X取表字段信息Map(table)
	if err != nil {
		return ""
	}
	for name, field := range tableFields {
		if 文本类.X是否包含并忽略大小写(field.Key, "pri") {
			return name
		}
	}
	return ""
}

// mergeArguments通过合并`m.extraArgs`和给定的`args`创建并返回新的参数。
func (m *Model) mergeArguments(args []interface{}) []interface{} {
	if len(m.extraArgs) > 0 {
		newArgs := make([]interface{}, len(m.extraArgs)+len(args))
		copy(newArgs, m.extraArgs)
		copy(newArgs[len(m.extraArgs):], args)
		return newArgs
	}
	return args
}
