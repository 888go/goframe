// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
//

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

// GetDB returns the underlying DB.
// ff:取DB对象
// c:
func (c *Core) GetDB() DB {
	return c.db
}

// GetLink creates and returns the underlying database link object with transaction checks.
// The parameter `master` specifies whether using the master node if master-slave configured.
// ff:取数据库链接对象
// c:
// ctx:上下文
// master:主节点
// schema:
// Link:
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

// MasterLink acts like function Master but with additional `schema` parameter specifying
// the schema for the connection. It is defined for internal usage.
// Also see Master.
// ff:底层MasterLink
// c:
// schema:
// Link:
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

// SlaveLink acts like function Slave but with additional `schema` parameter specifying
// the schema for the connection. It is defined for internal usage.
// Also see Slave.
// ff:底层SlaveLink
// c:
// schema:
// Link:
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

// QuoteWord checks given string `s` a word,
// if true it quotes `s` with security chars of the database
// and returns the quoted string; or else it returns `s` without any change.
//
// The meaning of a `word` can be considered as a column name.
// ff:底层QuoteWord
// c:
// s:
func (c *Core) QuoteWord(s string) string {
	s = gstr.Trim(s)
	if s == "" {
		return s
	}
	charLeft, charRight := c.db.GetChars()
	return doQuoteWord(s, charLeft, charRight)
}

// QuoteString quotes string with quote chars. Strings like:
// "user", "user u", "user,user_detail", "user u, user_detail ut", "u.id asc".
//
// The meaning of a `string` can be considered as part of a statement string including columns.
// ff:底层QuoteString
// c:
// s:
func (c *Core) QuoteString(s string) string {
	if !gregex.IsMatchString(regularFieldNameWithCommaRegPattern, s) {
		return s
	}
	charLeft, charRight := c.db.GetChars()
	return doQuoteString(s, charLeft, charRight)
}

// QuotePrefixTableName adds prefix string and quotes chars for the table.
// It handles table string like:
// "user", "user u",
// "user,user_detail",
// "user u, user_detail ut",
// "user as u, user_detail as ut".
//
// Note that, this will automatically checks the table prefix whether already added,
// if true it does nothing to the table name, or else adds the prefix to the table name.
// ff:底层添加前缀字符和引用字符
// c:
// table:表名称
func (c *Core) QuotePrefixTableName(table string) string {
	charLeft, charRight := c.db.GetChars()
	return doQuoteTableName(table, c.db.GetPrefix(), charLeft, charRight)
}

// GetChars returns the security char for current database.
// It does nothing in default.
// ff:底层取数据库安全字符
// c:
// charLeft:左字符
// charRight:右字符
func (c *Core) GetChars() (charLeft string, charRight string) {
	return "", ""
}

// Tables retrieves and returns the tables of current schema.
// It's mainly used in cli tool chain for automatically generating the models.
// ff:取表名称切片
// c:
// ctx:上下文
// schema:
// tables:表名称切片
// err:错误
func (c *Core) Tables(ctx context.Context, schema ...string) (tables []string, err error) {
	return
}

// TableFields retrieves and returns the fields' information of specified table of current
// schema.
//
// The parameter `link` is optional, if given nil it automatically retrieves a raw sql connection
// as its link to proceed necessary sql query.
//
// Note that it returns a map containing the field name and its corresponding fields.
// As a map is unsorted, the TableField struct has an "Index" field marks its sequence in
// the fields.
//
// It's using cache feature to enhance the performance, which is never expired util the
// process restarts.
// ff:取表字段信息Map
// c:
// ctx:上下文
// table:表名称
// schema:
// fields:字段信息Map
// err:错误
func (c *Core) TableFields(ctx context.Context, table string, schema ...string) (fields map[string]*TableField, err error) {
	return
}

// ClearTableFields removes certain cached table fields of current configuration group.
// ff:删除表字段缓存
// c:
// ctx:上下文
// table:表名称
// schema:
// err:错误
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

// ClearTableFieldsAll removes all cached table fields of current configuration group.
// ff:删除表字段所有缓存
// c:
// ctx:上下文
// err:错误
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

// ClearCache removes cached sql result of certain table.
// ff:删除表查询缓存
// c:
// ctx:上下文
// table:表名称
// err:错误
func (c *Core) ClearCache(ctx context.Context, table string) (err error) {
	return c.db.GetCache().Clear(ctx)
}

// ClearCacheAll removes all cached sql result from cache
// ff:删除所有表查询缓存
// c:
// ctx:上下文
// err:错误
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

// HasField determine whether the field exists in the table.
// ff:是否存在字段
// c:
// ctx:上下文
// table:表名称
// field:字段名称
// schema:
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

// guessPrimaryTableName parses and returns the primary table name.
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
