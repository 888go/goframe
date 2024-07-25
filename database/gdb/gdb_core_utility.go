// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。 md5:a114f4bdd106ab31

package gdb

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

// GetDB 返回底层的DB。 md5:5ebeb6e695bd2a8a
func (c *Core) GetDB() DB {
	return c.db
}

// GetLink 创建并返回底层数据库链接对象，并进行事务检查。
// 参数 `master` 指定在配置了主从的情况下是否使用主节点。 md5:51315fe7b2e9a929
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

// MasterLink 类似于函数 Master，但增加了 `schema` 参数，用于指定连接的模式。
// 这个函数主要用于内部使用。同时也参考 Master 函数。 md5:ae74b996555aea95
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

// SlaveLink 行为类似于 Slave 函数，但增加了 `schema` 参数，用于指定连接的模式。它主要用于内部使用。
// 参阅 Slave。 md5:8a8929395882c04a
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

// QuoteWord 检查给定的字符串 `s` 是否为一个单词，
// 如果是，它将使用数据库的安全字符对 `s` 进行转义，并返回带引号的字符串；否则，返回原始字符串不做任何更改。
//
// 可以认为一个 `word` 表示列名。 md5:71291615d7bcffe0
func (c *Core) QuoteWord(s string) string {
	s = gstr.Trim(s)
	if s == "" {
		return s
	}
	charLeft, charRight := c.db.GetChars()
	return doQuoteWord(s, charLeft, charRight)
}

// QuoteString 使用引号字符对字符串进行引用。例如这样的字符串：
// "user", "user u", "user,user_detail", "user u, user_detail ut", "u.id asc".
//
// `string` 的含义可以理解为作为包含列名的语句字符串的一部分。 md5:09c5263950e9ed1a
func (c *Core) QuoteString(s string) string {
	if !gregex.IsMatchString(regularFieldNameWithCommaRegPattern, s) {
		return s
	}
	charLeft, charRight := c.db.GetChars()
	return doQuoteString(s, charLeft, charRight)
}

// QuotePrefixTableName 为表名添加前缀字符串并包围引号。
// 它可以处理如下形式的表字符串：
// "user", "user u",
// "user,user_detail",
// "user u, user_detail ut",
// "user as u, user_detail as ut".
//
// 请注意，此函数会自动检查表前缀是否已经添加，
// 如果是，则不对表名做任何处理，否则会在表名前添加前缀。 md5:46ab3c3833cc0124
func (c *Core) QuotePrefixTableName(table string) string {
	charLeft, charRight := c.db.GetChars()
	return doQuoteTableName(table, c.db.GetPrefix(), charLeft, charRight)
}

// GetChars 返回当前数据库的安全字符。在默认情况下，它不执行任何操作。 md5:681b4cc93b5adecd
func (c *Core) GetChars() (charLeft string, charRight string) {
	return "", ""
}

// Tables 获取并返回当前模式下的表格列表。
// 主要用于命令行工具链，用于自动生成模型。 md5:bce161ba95454bf5
func (c *Core) Tables(ctx context.Context, schema ...string) (tables []string, err error) {
	return
}

// TableFields 获取并返回当前模式指定表的字段信息。
//
// 参数 `link` 是可选的，如果为 nil，则自动获取一个原始 SQL 连接，用于执行必要的 SQL 查询。
//
// 它返回一个包含字段名及其对应字段的映射。由于映射是无序的，TableField 结构体有一个 "Index" 字段，标记其在字段中的顺序。
//
// 该方法使用缓存功能来提高性能，直到进程重启，缓存永不过期。 md5:c844572d5210b35e
func (c *Core) TableFields(ctx context.Context, table string, schema ...string) (fields map[string]*TableField, err error) {
	return
}

// ClearTableFields 清除当前配置组中特定的缓存表字段。 md5:061271b8a4f298a0
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

// ClearTableFieldsAll 删除当前配置组中所有缓存的表字段。 md5:2b2f2ebba86cfda6
func (c *Core) ClearTableFieldsAll(ctx context.Context) (err error) {
	var (
		keys        = tableFieldsMap.Keys()
		cachePrefix = fmt.Sprintf(`%s%s`, cachePrefixTableFields, c.db.GetGroup())
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

// ClearCache 删除特定表的缓存SQL结果。 md5:5849435c2385500b
func (c *Core) ClearCache(ctx context.Context, table string) (err error) {
	return c.db.GetCache().Clear(ctx)
}

// ClearCacheAll 从缓存中移除所有已缓存的SQL结果. md5:1cafe85ca7b9f62d
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

// HasField 用于判断该字段是否存在于表中。 md5:e26ad0ecb292096b
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

// guessPrimaryTableName 解析并返回主表名称。 md5:d6aaf3f09d0afaaa
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
