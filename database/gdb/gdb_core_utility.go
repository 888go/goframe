// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

package db类

import (
	"context"
	"fmt"

	gmd5 "github.com/888go/goframe/crypto/gmd5"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	gregex "github.com/888go/goframe/text/gregex"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
	gutil "github.com/888go/goframe/util/gutil"
)

// X取DB对象 返回底层的DB。 md5:5ebeb6e695bd2a8a
func (c *Core) X取DB对象() DB {
	return c.db
}

// X取数据库链接对象 创建并返回底层数据库链接对象，并进行事务检查。
// 参数 `master` 指定在配置了主从的情况下是否使用主节点。
// md5:51315fe7b2e9a929
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

// X底层MasterLink 类似于函数 Master，但增加了 `schema` 参数，用于指定连接的模式。
// 这个函数主要用于内部使用。同时也参考 Master 函数。
// md5:ae74b996555aea95
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

// X底层SlaveLink 行为类似于 Slave 函数，但增加了 `schema` 参数，用于指定连接的模式。它主要用于内部使用。
// 参阅 Slave。
// md5:8a8929395882c04a
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

// X底层QuoteWord 检查给定的字符串 `s` 是否为一个单词，
// 如果是，它将使用数据库的安全字符对 `s` 进行转义，并返回带引号的字符串；否则，返回原始字符串不做任何更改。
//
// 可以认为一个 `word` 表示列名。
// md5:71291615d7bcffe0
func (c *Core) X底层QuoteWord(s string) string {
	s = gstr.X过滤首尾符并含空白(s)
	if s == "" {
		return s
	}
	charLeft, charRight := c.db.X底层取数据库安全字符()
	return doQuoteWord(s, charLeft, charRight)
}

// X底层QuoteString 使用引号字符对字符串进行引用。例如这样的字符串：
// "user", "user u", "user,user_detail", "user u, user_detail ut", "u.id asc".
//
// `string` 的含义可以理解为作为包含列名的语句字符串的一部分。
// md5:09c5263950e9ed1a
func (c *Core) X底层QuoteString(s string) string {
	if !gregex.X是否匹配文本(regularFieldNameWithCommaRegPattern, s) {
		return s
	}
	charLeft, charRight := c.db.X底层取数据库安全字符()
	return doQuoteString(s, charLeft, charRight)
}

// X底层添加前缀字符和引用字符 为表名添加前缀字符串并包围引号。
// 它可以处理如下形式的表字符串：
// "user", "user u",
// "user,user_detail",
// "user u, user_detail ut",
// "user as u, user_detail as ut".
//
// 请注意，此函数会自动检查表前缀是否已经添加，
// 如果是，则不对表名做任何处理，否则会在表名前添加前缀。
// md5:46ab3c3833cc0124
func (c *Core) X底层添加前缀字符和引用字符(表名称 string) string {
	charLeft, charRight := c.db.X底层取数据库安全字符()
	return doQuoteTableName(表名称, c.db.X取表前缀(), charLeft, charRight)
}

// X底层取数据库安全字符 返回当前数据库的安全字符。在默认情况下，它不执行任何操作。
// md5:681b4cc93b5adecd
func (c *Core) X底层取数据库安全字符() (左字符 string, 右字符 string) {
	return "", ""
}

// X取表名称切片 获取并返回当前模式下的表格列表。
//主要用于命令行工具链，用于自动生成模型。
// md5:bce161ba95454bf5
func (c *Core) X取表名称切片(上下文 context.Context, schema ...string) (表名称切片 []string, 错误 error) {
	return
}

// X取表字段信息Map 获取并返回当前模式指定表的字段信息。
// 
// 参数 `link` 是可选的，如果为 nil，则自动获取一个原始 SQL 连接，用于执行必要的 SQL 查询。
// 
// 它返回一个包含字段名及其对应字段的映射。由于映射是无序的，TableField 结构体有一个 "Index" 字段，标记其在字段中的顺序。
// 
// 该方法使用缓存功能来提高性能，直到进程重启，缓存永不过期。
// md5:c844572d5210b35e
func (c *Core) X取表字段信息Map(上下文 context.Context, 表名称 string, schema ...string) (字段信息Map map[string]*TableField, 错误 error) {
	return
}

// X删除表字段缓存 清除当前配置组中特定的缓存表字段。 md5:061271b8a4f298a0
func (c *Core) X删除表字段缓存(上下文 context.Context, 表名称 string, schema ...string) (错误 error) {
	tableFieldsMap.X删除(fmt.Sprintf(
		`%s%s@%s#%s`,
		cachePrefixTableFields,
		c.db.X取配置组名称(),
		gutil.X取文本值或取默认值(c.db.X取默认数据库名称(), schema...),
		表名称,
	))
	return
}

// X删除表字段所有缓存 删除当前配置组中所有缓存的表字段。 md5:2b2f2ebba86cfda6
func (c *Core) X删除表字段所有缓存(上下文 context.Context) (错误 error) {
	var (
		keys        = tableFieldsMap.X取所有名称()
		cachePrefix = fmt.Sprintf(`%s%s`, cachePrefixTableFields, c.db.X取配置组名称())
		removedKeys = make([]string, 0)
	)
	for _, key := range keys {
		if gstr.X开头判断(key, cachePrefix) {
			removedKeys = append(removedKeys, key)
		}
	}
	if len(removedKeys) > 0 {
		tableFieldsMap.X删除多个值(removedKeys)
	}
	return
}

// X删除表查询缓存 删除特定表的缓存SQL结果。 md5:5849435c2385500b
func (c *Core) X删除表查询缓存(上下文 context.Context, 表名称 string) (错误 error) {
	return c.db.X取缓存对象().X清空(上下文)
}

// X删除所有表查询缓存 从缓存中移除所有已缓存的SQL结果. md5:1cafe85ca7b9f62d
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
			gmd5.X加密文本PANI(sql+", @PARAMS:"+gconv.String(args)),
		)
	}
	return fmt.Sprintf(`%s%s`, cachePrefixSelectCache, name)
}

// X是否存在字段 用于判断该字段是否存在于表中。 md5:e26ad0ecb292096b
func (c *Core) X是否存在字段(上下文 context.Context, 表名称, 字段名称 string, schema ...string) (bool, error) {
	表名称 = c.guessPrimaryTableName(表名称)
	tableFields, err := c.db.X取表字段信息Map(上下文, 表名称, schema...)
	if err != nil {
		return false, err
	}
	if len(tableFields) == 0 {
		return false, gerror.X创建错误码并格式化(
			gcode.CodeNotFound,
			`empty table fields for table "%s"`, 表名称,
		)
	}
	fieldsArray := make([]string, len(tableFields))
	for k, v := range tableFields {
		fieldsArray[v.Index] = k
	}
	charLeft, charRight := c.db.X底层取数据库安全字符()
	字段名称 = gstr.X过滤首尾符并含空白(字段名称, charLeft+charRight)
	for _, f := range fieldsArray {
		if f == 字段名称 {
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
		array1           = gstr.X分割并忽略空值(tableStr, ",")
		array2           = gstr.X分割并忽略空值(array1[0], " ")
		array3           = gstr.X分割并忽略空值(array2[0], ".")
	)
	if len(array3) >= 2 {
		guessedTableName = array3[1]
	} else {
		guessedTableName = array3[0]
	}
	charL, charR := c.db.X底层取数据库安全字符()
	if charL != "" || charR != "" {
		guessedTableName = gstr.X过滤首尾符并含空白(guessedTableName, charL+charR)
	}
	if !gregex.X是否匹配文本(regularFieldNameRegPattern, guessedTableName) {
		return ""
	}
	return guessedTableName
}
