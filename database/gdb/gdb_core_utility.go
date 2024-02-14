// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//

package db类

import (
	"context"
	"fmt"
	
	"github.com/888go/goframe/crypto/gmd5"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
)

// GetDB 返回底层的 DB（数据库）
func (c *Core) X取DB对象() DB {
	return c.db
}

// GetLink 创建并返回底层的数据库连接对象，同时进行事务检查。
// 参数 `master` 指定在主从配置的情况下是否使用主节点。
func (c *Core) X取数据库链接对象(上下文 context.Context, 主节点 bool, schema string) (Link, error) {
	tx := X事务从上下文取对象(上下文, c.db.X取配置组名称())
	if tx != nil {
		return &txLink{tx.X底层取事务对象()}, nil
	}
	if 主节点 {
		link, err := c.db.X取Core对象().X底层MasterLink(schema)
		if err != nil {
			return nil, err
		}
		return link, nil
	}
	link, err := c.db.X取Core对象().X底层SlaveLink(schema)
	if err != nil {
		return nil, err
	}
	return link, nil
}

// MasterLink 表现得像函数 Master，但额外添加了一个 `schema` 参数用于指定连接的模式。它被定义为内部使用。
// 有关更多信息，请参阅 Master。
func (c *Core) X底层MasterLink(schema ...string) (Link, error) {
	db, err := c.db.X取主节点对象(schema...)
	if err != nil {
		return nil, err
	}
	return &dbLink{
		DB:         db,
		isOnMaster: true,
	}, nil
}

// SlaveLink 表现得像函数 Slave，但额外添加了一个 `schema` 参数用于指定连接的模式。它被定义为内部使用。
// 有关更多信息，请参阅 Slave。
func (c *Core) X底层SlaveLink(schema ...string) (Link, error) {
	db, err := c.db.X取从节点对象(schema...)
	if err != nil {
		return nil, err
	}
	return &dbLink{
		DB:         db,
		isOnMaster: false,
	}, nil
}

// 2024-01-09 改成内部方法,此方法属于底层, 几乎用不到.
// QuoteWord 检查给定字符串 `s` 是否为一个单词，
// 如果是，它会使用数据库的安全字符对 `s` 进行引用，并返回引述后的字符串；
// 否则，它将直接返回未经修改的 `s`。
//
// 这里的“单词”可以理解为列名。
func (c *Core) X底层QuoteWord(s string) string {
	s = 文本类.X过滤首尾符并含空白(s)
	if s == "" {
		return s
	}
	charLeft, charRight := c.db.X底层取数据库安全字符()
	return doQuoteWord(s, charLeft, charRight)
}

// QuoteString 用引号字符对字符串进行引用。例如以下字符串：
// "user", "user u", "user,user_detail", "user u, user_detail ut", "u.id asc".
//
// 可以将 `string` 的含义视为包含列部分的语句字符串中的一部分。
func (c *Core) X底层QuoteString(s string) string {
	charLeft, charRight := c.db.X底层取数据库安全字符()
	return doQuoteString(s, charLeft, charRight)
}

// QuotePrefixTableName 为表名添加前缀字符串和引用字符。
// 它处理诸如以下格式的表名：
// "user", "user u",
// "user,user_detail",
// "user u, user_detail ut",
// "user as u, user_detail as ut"。
//
// 注意，此函数会自动检查表名是否已添加了前缀，
// 如果已经添加，则不对表名做任何处理；否则，将前缀添加到表名中。
func (c *Core) X底层添加前缀字符和引用字符(表名称 string) string {
	charLeft, charRight := c.db.X底层取数据库安全字符()
	return doQuoteTableName(表名称, c.db.X取表前缀(), charLeft, charRight)
}

// GetChars 返回当前数据库的安全字符。
// 默认情况下，此方法不做任何操作。
func (c *Core) X底层取数据库安全字符() (左字符 string, 右字符 string) {
	return "", ""
}

// Tables 获取并返回当前模式的表。
// 它主要用于cli工具链中，用于自动生成模型。
func (c *Core) X取表名称数组(上下文 context.Context, schema ...string) (表名称数组 []string, 错误 error) {
	return
}

// TableFields 获取并返回当前模式下指定表的字段信息。
//
// 参数 `link` 是可选的，如果给出 nil，则会自动获取一个原始的 SQL 连接作为其链接以执行必要的 SQL 查询。
//
// 注意，它返回一个包含字段名及其对应字段信息的映射。由于映射是无序的，TableField 结构体有一个 "Index" 字段用于标记其在所有字段中的顺序。
//
// 为了提高性能，该函数使用了缓存特性，缓存有效期直到进程重启才会过期。
func (c *Core) X取表字段信息Map(上下文 context.Context, 表名称 string, schema ...string) (字段信息Map map[string]*TableField, 错误 error) {
	return
}

// ClearTableFields 删除当前配置组中特定的缓存表字段。
func (c *Core) X删除表字段缓存(上下文 context.Context, 表名称 string, schema ...string) (错误 error) {
	tableFieldsMap.X删除(fmt.Sprintf(
		`%s%s@%s#%s`,
		cachePrefixTableFields,
		c.db.X取配置组名称(),
		工具类.X取文本值或取默认值(c.db.X取默认数据库名称(), schema...),
		表名称,
	))
	return
}

// ClearTableFieldsAll 清除当前配置组中所有已缓存的表字段。
func (c *Core) X删除表字段所有缓存(上下文 context.Context) (错误 error) {
	var (
		keys        = tableFieldsMap.X取所有名称()
		cachePrefix = fmt.Sprintf(`%s@%s`, cachePrefixTableFields, c.db.X取配置组名称())
		removedKeys = make([]string, 0)
	)
	for _, key := range keys {
		if 文本类.X开头判断(key, cachePrefix) {
			removedKeys = append(removedKeys, key)
		}
	}
	if len(removedKeys) > 0 {
		tableFieldsMap.X删除多个值(removedKeys)
	}
	return
}

// ClearCache 清除特定表的缓存SQL结果。
func (c *Core) X删除表查询缓存(上下文 context.Context, 表名称 string) (错误 error) {
	return c.db.X取缓存对象().X清空(上下文)
}

// ClearCacheAll 从缓存中移除所有已缓存的SQL查询结果
func (c *Core) X删除所有表查询缓存(上下文 context.Context) (错误 error) {
	return c.db.X取缓存对象().X清空(上下文)
}

func (c *Core) makeSelectCacheKey(name, schema, table, sql string, args ...interface{}) string {
	if name == "" {
		name = fmt.Sprintf(
			`%s@%s#%s:%s`,
			c.db.X取配置组名称(),
			schema,
			table,
			加密md5类.X加密文本PANI(sql+", @PARAMS:"+转换类.String(args)),
		)
	}
	return fmt.Sprintf(`%s%s`, cachePrefixSelectCache, name)
}

// HasField 判断字段是否在表中存在。
func (c *Core) X是否存在字段(上下文 context.Context, 表名称, 字段名称 string, schema ...string) (bool, error) {
	表名称 = c.guessPrimaryTableName(表名称)
	tableFields, err := c.db.X取表字段信息Map(上下文, 表名称, schema...)
	if err != nil {
		return false, err
	}
	if len(tableFields) == 0 {
		return false, 错误类.X创建错误码并格式化(
			错误码类.CodeNotFound,
			`empty table fields for table "%s"`, 表名称,
		)
	}
	fieldsArray := make([]string, len(tableFields))
	for k, v := range tableFields {
		fieldsArray[v.Index] = k
	}
	charLeft, charRight := c.db.X底层取数据库安全字符()
	字段名称 = 文本类.X过滤首尾符并含空白(字段名称, charLeft+charRight)
	for _, f := range fieldsArray {
		if f == 字段名称 {
			return true, nil
		}
	}
	return false, nil
}

// guessPrimaryTableName 解析并返回主表名称。
func (c *Core) guessPrimaryTableName(tableStr string) string {
	if tableStr == "" {
		return ""
	}
	var (
		guessedTableName string
		array1           = 文本类.X分割并忽略空值(tableStr, ",")
		array2           = 文本类.X分割并忽略空值(array1[0], " ")
		array3           = 文本类.X分割并忽略空值(array2[0], ".")
	)
	if len(array3) >= 2 {
		guessedTableName = array3[1]
	} else {
		guessedTableName = array3[0]
	}
	charL, charR := c.db.X底层取数据库安全字符()
	if charL != "" || charR != "" {
		guessedTableName = 文本类.X过滤首尾符并含空白(guessedTableName, charL+charR)
	}
	if !正则类.X是否匹配文本(regularFieldNameRegPattern, guessedTableName) {
		return ""
	}
	return guessedTableName
}
