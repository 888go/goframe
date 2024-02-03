// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//

package gdb

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
func (c *Core) GetDB() DB {
	return c.db
}

// GetLink 创建并返回底层的数据库连接对象，同时进行事务检查。
// 参数 `master` 指定在主从配置的情况下是否使用主节点。
func (c *Core) GetLink(ctx context.Context, master bool, schema string) (Link, error) {
	tx := TXFromCtx(ctx, c.db.GetGroup())
	if tx != nil {
		return &txLink{tx.GetSqlTX()}, nil
	}
	if master {
		link, err := c.db.GetCore().MasterLink(schema)
		if err != nil {
			return nil, err
		}
		return link, nil
	}
	link, err := c.db.GetCore().SlaveLink(schema)
	if err != nil {
		return nil, err
	}
	return link, nil
}

// MasterLink 表现得像函数 Master，但额外添加了一个 `schema` 参数用于指定连接的模式。它被定义为内部使用。
// 有关更多信息，请参阅 Master。
func (c *Core) MasterLink(schema ...string) (Link, error) {
	db, err := c.db.Master(schema...)
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
func (c *Core) SlaveLink(schema ...string) (Link, error) {
	db, err := c.db.Slave(schema...)
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
func (c *Core) QuoteWord(s string) string {
	s = gstr.Trim(s)
	if s == "" {
		return s
	}
	charLeft, charRight := c.db.GetChars()
	return doQuoteWord(s, charLeft, charRight)
}

// QuoteString 用引号字符对字符串进行引用。例如以下字符串：
// "user", "user u", "user,user_detail", "user u, user_detail ut", "u.id asc".
//
// 可以将 `string` 的含义视为包含列部分的语句字符串中的一部分。
func (c *Core) QuoteString(s string) string {
	charLeft, charRight := c.db.GetChars()
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
func (c *Core) QuotePrefixTableName(table string) string {
	charLeft, charRight := c.db.GetChars()
	return doQuoteTableName(table, c.db.GetPrefix(), charLeft, charRight)
}

// GetChars 返回当前数据库的安全字符。
// 默认情况下，此方法不做任何操作。
func (c *Core) GetChars() (charLeft string, charRight string) {
	return "", ""
}

// Tables 获取并返回当前模式的表。
// 它主要用于cli工具链中，用于自动生成模型。
func (c *Core) Tables(ctx context.Context, schema ...string) (tables []string, err error) {
	return
}

// TableFields 获取并返回当前模式下指定表的字段信息。
//
// 参数 `link` 是可选的，如果给出 nil，则会自动获取一个原始的 SQL 连接作为其链接以执行必要的 SQL 查询。
//
// 注意，它返回一个包含字段名及其对应字段信息的映射。由于映射是无序的，TableField 结构体有一个 "Index" 字段用于标记其在所有字段中的顺序。
//
// 为了提高性能，该函数使用了缓存特性，缓存有效期直到进程重启才会过期。
func (c *Core) TableFields(ctx context.Context, table string, schema ...string) (fields map[string]*TableField, err error) {
	return
}

// ClearTableFields 删除当前配置组中特定的缓存表字段。
func (c *Core) ClearTableFields(ctx context.Context, table string, schema ...string) (err error) {
	tableFieldsMap.Remove(fmt.Sprintf(
		`%s%s@%s#%s`,
		cachePrefixTableFields,
		c.db.GetGroup(),
		gutil.GetOrDefaultStr(c.db.GetSchema(), schema...),
		table,
	))
	return
}

// ClearTableFieldsAll 清除当前配置组中所有已缓存的表字段。
func (c *Core) ClearTableFieldsAll(ctx context.Context) (err error) {
	var (
		keys        = tableFieldsMap.Keys()
		cachePrefix = fmt.Sprintf(`%s@%s`, cachePrefixTableFields, c.db.GetGroup())
		removedKeys = make([]string, 0)
	)
	for _, key := range keys {
		if gstr.HasPrefix(key, cachePrefix) {
			removedKeys = append(removedKeys, key)
		}
	}
	if len(removedKeys) > 0 {
		tableFieldsMap.Removes(removedKeys)
	}
	return
}

// ClearCache 清除特定表的缓存SQL结果。
func (c *Core) ClearCache(ctx context.Context, table string) (err error) {
	return c.db.GetCache().Clear(ctx)
}

// ClearCacheAll 从缓存中移除所有已缓存的SQL查询结果
func (c *Core) ClearCacheAll(ctx context.Context) (err error) {
	return c.db.GetCache().Clear(ctx)
}

func (c *Core) makeSelectCacheKey(name, schema, table, sql string, args ...interface{}) string {
	if name == "" {
		name = fmt.Sprintf(
			`%s@%s#%s:%s`,
			c.db.GetGroup(),
			schema,
			table,
			gmd5.MustEncryptString(sql+", @PARAMS:"+gconv.String(args)),
		)
	}
	return fmt.Sprintf(`%s%s`, cachePrefixSelectCache, name)
}

// HasField 判断字段是否在表中存在。
func (c *Core) HasField(ctx context.Context, table, field string, schema ...string) (bool, error) {
	table = c.guessPrimaryTableName(table)
	tableFields, err := c.db.TableFields(ctx, table, schema...)
	if err != nil {
		return false, err
	}
	if len(tableFields) == 0 {
		return false, gerror.NewCodef(
			gcode.CodeNotFound,
			`empty table fields for table "%s"`, table,
		)
	}
	fieldsArray := make([]string, len(tableFields))
	for k, v := range tableFields {
		fieldsArray[v.Index] = k
	}
	charLeft, charRight := c.db.GetChars()
	field = gstr.Trim(field, charLeft+charRight)
	for _, f := range fieldsArray {
		if f == field {
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
		array1           = gstr.SplitAndTrim(tableStr, ",")
		array2           = gstr.SplitAndTrim(array1[0], " ")
		array3           = gstr.SplitAndTrim(array2[0], ".")
	)
	if len(array3) >= 2 {
		guessedTableName = array3[1]
	} else {
		guessedTableName = array3[0]
	}
	charL, charR := c.db.GetChars()
	if charL != "" || charR != "" {
		guessedTableName = gstr.Trim(guessedTableName, charL+charR)
	}
	if !gregex.IsMatchString(regularFieldNameRegPattern, guessedTableName) {
		return ""
	}
	return guessedTableName
}
