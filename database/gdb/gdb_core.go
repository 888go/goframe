// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

package db类

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	gmap "github.com/888go/goframe/container/gmap"
	gset "github.com/888go/goframe/container/gset"
	gvar "github.com/888go/goframe/container/gvar"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/reflection"
	"github.com/888go/goframe/internal/utils"
	gregex "github.com/888go/goframe/text/gregex"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
	gutil "github.com/888go/goframe/util/gutil"
)

// X取Core对象 返回底层的 *Core 对象。 md5:b7d2ff344b9a6a33
func (c *Core) X取Core对象() *Core {
	return c
}

// X设置上下文并取副本是一个链式函数，它创建并返回一个新的DB对象，它是当前DB对象的浅拷贝，并且包含了给定的上下文。
// 注意，这个返回的DB对象只能使用一次，所以不要将其分配给全局变量或长期使用的包变量。
// md5:9dfddec16d5df9f5
func (c *Core) X设置上下文并取副本(上下文 context.Context) DB {
	if 上下文 == nil {
		return c.db
	}
		// 它会浅复制当前的数据库，并为下一个链式操作更改其上下文。 md5:bf3a0a0f1f30a496
	var (
		err        error
		newCore    = &Core{}
		configNode = c.db.X取当前节点配置()
	)
	*newCore = *c
		// 它创建了一个新的DB对象（不是新连接），通常是对`Core`对象的封装。 md5:6cd230087401c98f
	newCore.db, err = driverMap[configNode.Type].New(newCore, configNode)
	if err != nil {
		// 这里确实是一个严重的错误。
		// 不要让它继续下去。
		// md5:790820a929dc0bfd
		panic(err)
	}
	newCore.ctx = X底层WithDB(上下文, newCore.db)
	newCore.ctx = c.injectInternalCtxData(newCore.ctx)
	return newCore.db
}

// X取上下文对象 返回当前数据库的上下文。
// 如果之前没有设置上下文，则返回 `context.Background()`。
// md5:9b56f79a5eaa891e
func (c *Core) X取上下文对象() context.Context {
	ctx := c.ctx
	if ctx == nil {
		ctx = context.TODO()
	}
	return c.injectInternalCtxData(ctx)
}

// X取超时上下文对象 返回指定超时类型的上下文和取消函数。 md5:5d0be7078de61c6d
func (c *Core) X取超时上下文对象(上下文 context.Context, 超时类型 int) (context.Context, context.CancelFunc) {
	if 上下文 == nil {
		上下文 = c.db.X取上下文对象()
	} else {
		上下文 = context.WithValue(上下文, "WrappedByGetCtxTimeout", nil)
	}
	switch 超时类型 {
	case ctxTimeoutTypeExec:
		if c.db.X取当前节点配置().ExecTimeout > 0 {
			return context.WithTimeout(上下文, c.db.X取当前节点配置().ExecTimeout)
		}
	case ctxTimeoutTypeQuery:
		if c.db.X取当前节点配置().QueryTimeout > 0 {
			return context.WithTimeout(上下文, c.db.X取当前节点配置().QueryTimeout)
		}
	case ctxTimeoutTypePrepare:
		if c.db.X取当前节点配置().PrepareTimeout > 0 {
			return context.WithTimeout(上下文, c.db.X取当前节点配置().PrepareTimeout)
		}
	default:
		panic(gerror.X创建错误码并格式化(gcode.CodeInvalidParameter, "invalid context timeout type: %d", 超时类型))
	}
	return 上下文, func() {}
}

// X关闭数据库 方法关闭数据库并阻止新的查询开始。
// 随后，X关闭数据库 会等待所有已在服务器上开始处理的查询完成。
//
// 关闭 DB 实例是很少见的操作，因为 DB 处理句柄设计为长生命周期的，
// 并且旨在多个 goroutine 间共享。
// md5:39e5c90e1da0ee5e
func (c *Core) X关闭数据库(上下文 context.Context) (错误 error) {
	if 错误 = c.cache.X关闭(上下文); 错误 != nil {
		return 错误
	}
	c.links.X遍历写锁定(func(m map[any]any) {
		for k, v := range m {
			if db, ok := v.(*sql.DB); ok {
				错误 = db.Close()
				if 错误 != nil {
					错误 = gerror.X多层错误码(gcode.CodeDbOperationError, 错误, `db.Close failed`)
				}
				intlog.Printf(上下文, `close link: %s, err: %v`, k, 错误)
				if 错误 != nil {
					return
				}
				delete(m, k)
			}
		}
	})
	return
}

// X取主节点对象 如果配置了主从节点，则创建并返回一个与主节点的连接。
// 如果未配置主从节点，则返回默认连接。
// md5:0bd77b596cdae9a3
func (c *Core) X取主节点对象(数据库名称 ...string) (*sql.DB, error) {
	var (
		usedSchema   = gutil.X取文本值或取默认值(c.schema, 数据库名称...)
		charL, charR = c.db.X底层取数据库安全字符()
	)
	return c.getSqlDb(true, gstr.X过滤首尾符并含空白(usedSchema, charL+charR))
}

// 如果配置了主从模式，X取从节点对象 会创建并返回一个从节点连接。如果没有配置主从模式，则返回默认连接。
// md5:d92640050cf063d3
func (c *Core) X取从节点对象(数据库名称 ...string) (*sql.DB, error) {
	var (
		usedSchema   = gutil.X取文本值或取默认值(c.schema, 数据库名称...)
		charL, charR = c.db.X底层取数据库安全字符()
	)
	return c.getSqlDb(false, gstr.X过滤首尾符并含空白(usedSchema, charL+charR))
}

// GetAll别名 从数据库中查询并返回数据记录。 md5:dfdcfddaa70ab1d4
func (c *Core) GetAll别名(上下文 context.Context, sql string, 参数 ...interface{}) (Result, error) {
	return c.db.X底层查询(上下文, nil, sql, 参数...)
}

// X底层查询 从数据库查询并返回数据记录。 md5:82b06146b8d539d1
func (c *Core) X底层查询(上下文 context.Context, 链接 Link, sql string, 参数 ...interface{}) (结果 Result, 错误 error) {
	return c.db.X底层原生SQL查询(上下文, 链接, sql, 参数...)
}

// X原生SQL查询单条记录 从数据库中查询并返回一条记录。 md5:9552f7e095f58141
func (c *Core) X原生SQL查询单条记录(上下文 context.Context, sql string, 参数 ...interface{}) (Record, error) {
	list, err := c.db.GetAll别名(上下文, sql, 参数...)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	return nil, nil
}

// X原生SQL查询切片 从数据库查询并返回数据值作为切片。
// 注意，如果结果中有多个列，它会随机返回一列的值。
// md5:b81cd4c5e063a6f2
func (c *Core) X原生SQL查询切片(上下文 context.Context, sql string, 参数 ...interface{}) ([]Value, error) {
	all, err := c.db.X底层查询(上下文, nil, sql, 参数...)
	if err != nil {
		return nil, err
	}
	return all.X取字段切片(), nil
}

// doGetStruct 从数据库查询一条记录，并将其转换为给定的结构体。
// 参数 `pointer` 应该是指向结构体的指针。
// md5:9260d4f62deef626
func (c *Core) doGetStruct(ctx context.Context, pointer interface{}, sql string, args ...interface{}) error {
	one, err := c.db.X原生SQL查询单条记录(ctx, sql, args...)
	if err != nil {
		return err
	}
	return one.X取结构体指针(pointer)
}

// doGetStructs 从数据库查询记录并将其转换为给定的结构体。
// 参数 `pointer` 应为结构体切片类型：[]struct/[]*struct。
// md5:4ce864edda9b9231
func (c *Core) doGetStructs(ctx context.Context, pointer interface{}, sql string, args ...interface{}) error {
	all, err := c.db.GetAll别名(ctx, sql, args...)
	if err != nil {
		return err
	}
	return all.X取切片结构体指针(pointer)
}

// X原生SQL查询到结构体指针 从数据库查询一个或多个记录，并将它们转换为给定的结构体或结构体数组。
//
// 如果参数 `pointer` 是结构体指针类型，它内部会调用 GetStruct 进行转换。如果参数 `pointer` 是切片类型，它内部会调用 GetStructs 进行转换。
// md5:c1dbdab7a7c29c51
func (c *Core) X原生SQL查询到结构体指针(上下文 context.Context, 结构体指针 interface{}, sql string, 参数 ...interface{}) error {
	reflectInfo := reflection.OriginTypeAndKind(结构体指针)
	if reflectInfo.InputKind != reflect.Ptr {
		return gerror.X创建错误码并格式化(
			gcode.CodeInvalidParameter,
			"params should be type of pointer, but got: %v",
			reflectInfo.InputKind,
		)
	}
	switch reflectInfo.OriginKind {
	case reflect.Array, reflect.Slice:
		return c.db.X取Core对象().doGetStructs(上下文, 结构体指针, sql, 参数...)

	case reflect.Struct:
		return c.db.X取Core对象().doGetStruct(上下文, 结构体指针, sql, 参数...)
	}
	return gerror.X创建错误码并格式化(
		gcode.CodeInvalidParameter,
		`in valid parameter type "%v", of which element type should be type of struct/slice`,
		reflectInfo.InputType,
	)
}

// X原生SQL查询字段值 从数据库查询并返回字段值。
// SQL 应该只查询数据库中的一个字段，否则它将只返回结果中的一个字段。
// md5:96794360fadbc288
func (c *Core) X原生SQL查询字段值(上下文 context.Context, sql string, 参数 ...interface{}) (Value, error) {
	one, err := c.db.X原生SQL查询单条记录(上下文, sql, 参数...)
	if err != nil {
		return gvar.X创建(nil), err
	}
	for _, v := range one {
		return v, nil
	}
	return gvar.X创建(nil), nil
}

// X原生SQL查询字段计数 从数据库中查询并返回计数。 md5:a8368d39f4a58979
func (c *Core) X原生SQL查询字段计数(上下文 context.Context, sql string, 参数 ...interface{}) (int, error) {
	// 如果查询字段中不包含"COUNT"函数，
	// 则替换SQL字符串，并在字段中添加"COUNT"函数。
	// md5:624b6da82fb9facd
	if !gregex.X是否匹配文本(`(?i)SELECT\s+COUNT\(.+\)\s+FROM`, sql) {
		sql, _ = gregex.X替换文本(`(?i)(SELECT)\s+(.+)\s+(FROM)`, `$1 COUNT($2) $3`, sql)
	}
	value, err := c.db.X原生SQL查询字段值(上下文, sql, 参数...)
	if err != nil {
		return 0, err
	}
	return value.X取整数(), nil
}

// X多表去重查询 执行 "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." 语句。 md5:6a2f9809c172cb31
func (c *Core) X多表去重查询(Model对象 ...*Model) *Model {
	var ctx = c.db.X取上下文对象()
	return c.doUnion(ctx, unionTypeNormal, Model对象...)
}

// X多表查询 生成 "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ... " 语句。 md5:5a15c8720fcb152f
func (c *Core) X多表查询(Model对象 ...*Model) *Model {
	var ctx = c.db.X取上下文对象()
	return c.doUnion(ctx, unionTypeAll, Model对象...)
}

func (c *Core) doUnion(ctx context.Context, unionType int, unions ...*Model) *Model {
	var (
		unionTypeStr   string
		composedSqlStr string
		composedArgs   = make([]interface{}, 0)
	)
	if unionType == unionTypeAll {
		unionTypeStr = "UNION ALL"
	} else {
		unionTypeStr = "UNION"
	}
	for _, v := range unions {
		sqlWithHolder, holderArgs := v.getFormattedSqlAndArgs(ctx, queryTypeNormal, false)
		if composedSqlStr == "" {
			composedSqlStr += fmt.Sprintf(`(%s)`, sqlWithHolder)
		} else {
			composedSqlStr += fmt.Sprintf(` %s (%s)`, unionTypeStr, sqlWithHolder)
		}
		composedArgs = append(composedArgs, holderArgs...)
	}
	return c.db.X原生SQL(composedSqlStr, composedArgs...)
}

// X向主节点发送心跳 向主节点发送请求以检查身份验证或保持连接活动。 md5:47a7df55cbee8583
func (c *Core) X向主节点发送心跳() error {
	var ctx = c.db.X取上下文对象()
	if master, err := c.db.X取主节点对象(); err != nil {
		return err
	} else {
		if err = master.PingContext(ctx); err != nil {
			err = gerror.X多层错误码(gcode.CodeDbOperationError, err, `master.Ping failed`)
		}
		return err
	}
}

// X向从节点发送心跳 调用ping命令检查从节点的认证或者维持连接。 md5:62272b38d874eda6
func (c *Core) X向从节点发送心跳() error {
	var ctx = c.db.X取上下文对象()
	if slave, err := c.db.X取从节点对象(); err != nil {
		return err
	} else {
		if err = slave.PingContext(ctx); err != nil {
			err = gerror.X多层错误码(gcode.CodeDbOperationError, err, `slave.Ping failed`)
		}
		return err
	}
}

// X插入 执行 "INSERT INTO..." 语句来操作表。
// 如果表中已经存在数据的唯一记录，它会返回错误。
//
// 参数 `data` 可以是 map、gmap、struct、*struct、[]map 或 []struct 等类型。
// 例如：
// Data(g.Map{"uid": 10000, "name": "john"})
// Data(g.Slice{g.Map{"uid": 10000, "name": "john"}, g.Map{"uid": 20000, "name": "smith"}})
//
// 参数 `batch` 在给定数据为切片时，指定批量操作的次数。
// md5:fd75d343f485b8dc
func (c *Core) X插入(上下文 context.Context, 表名称 string, 值 interface{}, 批量操作行数 ...int) (sql.Result, error) {
	if len(批量操作行数) > 0 {
		return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X设置数据(值).X设置批量操作行数(批量操作行数[0]).X插入()
	}
	return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X设置数据(值).X插入()
}

// X插入并跳过已存在 执行 "INSERT IGNORE INTO ..." 语句针对该表。
// 如果表中已存在该数据的唯一记录，则忽略插入操作。
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 例如：
// Data(g.Map{"uid": 10000, "name": "john"})
// Data(g.Slice{g.Map{"uid": 10000, "name": "john"}, g.Map{"uid": 20000, "name": "smith"})
//
// 当给定的数据为切片时，参数 `batch` 指定批处理操作的计数。
// md5:49f76901041c9819
func (c *Core) X插入并跳过已存在(上下文 context.Context, 表名称 string, 值 interface{}, 批量操作行数 ...int) (sql.Result, error) {
	if len(批量操作行数) > 0 {
		return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X设置数据(值).X设置批量操作行数(批量操作行数[0]).X插入并跳过已存在()
	}
	return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X设置数据(值).X插入并跳过已存在()
}

// X插入并取ID 执行插入操作，并返回自动生成的最后一个插入id。 md5:8d00b40a35fa48a5
func (c *Core) X插入并取ID(上下文 context.Context, 表名称 string, 值 interface{}, 批量操作行数 ...int) (int64, error) {
	if len(批量操作行数) > 0 {
		return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X设置数据(值).X设置批量操作行数(批量操作行数[0]).X插入并取ID()
	}
	return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X设置数据(值).X插入并取ID()
}

// X插入并替换已存在 用于执行针对该表的 "REPLACE INTO..." 语句。如果表中已经存在数据的唯一记录，它会删除该记录并插入新的。
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。例如：
// Data(g.Map{"uid": 10000, "name": "john"})
// Data(g.Slice{g.Map{"uid": 10000, "name": "john"}, g.Map{"uid": 20000, "name": "smith"}})
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。如果给定的数据是切片类型，它将进行批量替换，并可选地通过参数 `batch` 指定批量操作次数。
// md5:69ecd0994eab5bbb
func (c *Core) X插入并替换已存在(上下文 context.Context, 表名称 string, 值 interface{}, 批量操作行数 ...int) (sql.Result, error) {
	if len(批量操作行数) > 0 {
		return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X设置数据(值).X设置批量操作行数(批量操作行数[0]).X插入并替换已存在()
	}
	return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X设置数据(值).X插入并替换已存在()
}

// X插入并更新已存在 执行 "INSERT INTO ... ON DUPLICATE KEY UPDATE..." 语句来操作表。
// 如果保存的数据中存在主键或唯一索引，它将更新记录；否则，将在表中插入新记录。
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 例如：
// Data(g.Map{"uid": 10000, "name": "john"})
// Data(g.Slice{g.Map{"uid": 10000, "name": "john"}, g.Map{"uid": 20000, "name": "smith"}})
//
// 如果给定的数据是切片类型，它将进行批量保存。可选参数 `batch` 指定了批量操作的次数。
// md5:c76721f5e0b01424
func (c *Core) X插入并更新已存在(上下文 context.Context, 表名称 string, 值 interface{}, 批量操作行数 ...int) (sql.Result, error) {
	if len(批量操作行数) > 0 {
		return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X设置数据(值).X设置批量操作行数(批量操作行数[0]).X插入并更新已存在()
	}
	return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X设置数据(值).X插入并更新已存在()
}

func (c *Core) fieldsToSequence(ctx context.Context, table string, fields []string) ([]string, error) {
	var (
		fieldSet               = gset.X创建文本并按值(fields)
		fieldsResultInSequence = make([]string, 0)
		tableFields, err       = c.db.X取表字段信息Map(ctx, table)
	)
	if err != nil {
		return nil, err
	}
		// 按顺序对字段进行排序。 md5:3edaf791b6d06284
	var fieldsOfTableInSequence = make([]string, len(tableFields))
	for _, field := range tableFields {
		fieldsOfTableInSequence[field.Index] = field.X名称
	}
	// Sort the input fields.
	for _, fieldName := range fieldsOfTableInSequence {
		if fieldSet.X是否存在(fieldName) {
			fieldsResultInSequence = append(fieldsResultInSequence, fieldName)
		}
	}
	return fieldsResultInSequence, nil
}

// X底层插入 函数用于插入或更新给定表的数据。
// 这个函数通常用于自定义接口定义，不需要手动调用。
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 示例：
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"}})
//
// 参数 `option` 的值如下：
// InsertOptionDefault：仅插入，如果数据中包含唯一键或主键，会返回错误；
// InsertOptionReplace：如果数据中包含唯一键或主键，会先从表中删除再插入新的记录；
// InsertOptionSave：   如果数据中包含唯一键或主键，会进行更新，否则插入新记录；
// InsertOptionIgnore： 如果数据中包含唯一键或主键，忽略插入操作。
// md5:e9554638335c9c80
func (c *Core) X底层插入(上下文 context.Context, 链接 Link, 表名称 string, list Map切片, option DoInsertOption) (result sql.Result, err error) {
	var (
		keys           []string      // Field names.
		values         []string      // 值持有字符串数组，例如：(?,?,?). md5:4dd91c222c15917f
		params         []interface{} // 将被提交给底层数据库驱动程序的值。 md5:d30c8d96f40663c3
		onDuplicateStr string        // onDuplicateStr 用于 "ON DUPLICATE KEY UPDATE" 语句。 md5:7056b1b5ea46e69e
	)
	// ============================================================================================
	// 按照字段对列表进行分组。不同的字段将数据分配到不同的列表中。
	// 此处使用ListMap来保持数据插入时的顺序。
	// ============================================================================================
	// md5:f3b3fbc2fd4a59f8
	var keyListMap = gmap.X创建链表mp()
	for _, item := range list {
		var (
			tmpKeys              = make([]string, 0)
			tmpKeysInSequenceStr string
		)
		for k := range item {
			tmpKeys = append(tmpKeys, k)
		}
		keys, err = c.fieldsToSequence(上下文, 表名称, tmpKeys)
		if err != nil {
			return nil, err
		}
		tmpKeysInSequenceStr = gstr.X连接(keys, ",")
		if !keyListMap.X是否存在(tmpKeysInSequenceStr) {
			keyListMap.X设置值(tmpKeysInSequenceStr, make(Map切片, 0))
		}
		tmpKeysInSequenceList := keyListMap.X取值(tmpKeysInSequenceStr).(Map切片)
		tmpKeysInSequenceList = append(tmpKeysInSequenceList, item)
		keyListMap.X设置值(tmpKeysInSequenceStr, tmpKeysInSequenceList)
	}
	if keyListMap.X取数量() > 1 {
		var (
			tmpResult    sql.Result
			sqlResult    SqlResult
			rowsAffected int64
		)
		keyListMap.X遍历(func(key, value interface{}) bool {
			tmpResult, err = c.X底层插入(上下文, 链接, 表名称, value.(Map切片), option)
			if err != nil {
				return false
			}
			rowsAffected, err = tmpResult.RowsAffected()
			if err != nil {
				return false
			}
			sqlResult.X原生sql行记录 = tmpResult
			sqlResult.X影响行数 += rowsAffected
			return true
		})
		return &sqlResult, err
	}

		// 准备批量结果指针。 md5:dfc8aa8bb292f9d5
	var (
		charL, charR = c.db.X底层取数据库安全字符()
		batchResult  = new(SqlResult)
		keysStr      = charL + strings.Join(keys, charR+","+charL) + charR
		operation    = X底层GetInsertOperationByOption(option.InsertOption)
	)
		// Upsert 子句只在 Save 操作中生效。 md5:c556e85b127111f7
	if option.InsertOption == InsertOptionSave {
		onDuplicateStr, err = c.db.FormatUpsert(keys, list, option)
		if err != nil {
			return nil, err
		}
	}
	var (
		listLength   = len(list)
		valueHolders = make([]string, 0)
	)
	for i := 0; i < listLength; i++ {
		values = values[:0]
		// 注意，映射类型是无序的，
		// 因此应该使用切片和键来检索值。
		// md5:2495d5e730dae78f
		for _, k := range keys {
			if s, ok := list[i][k].(Raw); ok {
				values = append(values, gconv.String(s))
			} else {
				values = append(values, "?")
				params = append(params, list[i][k])
			}
		}
		valueHolders = append(valueHolders, "("+gstr.X连接(values, ",")+")")
				// 批量包检查：它满足批量数量，或者它是最后一个元素。 md5:a2ef8b869c6d8888
		if len(valueHolders) == option.BatchCount || (i == listLength-1 && len(valueHolders) > 0) {
			var (
				stdSqlResult sql.Result
				affectedRows int64
			)
			stdSqlResult, err = c.db.X底层原生SQL执行(上下文, 链接, fmt.Sprintf(
				"%s INTO %s(%s) VALUES%s %s",
				operation, c.X底层添加前缀字符和引用字符(表名称), keysStr,
				gstr.X连接(valueHolders, ","),
				onDuplicateStr,
			), params...)
			if err != nil {
				return stdSqlResult, err
			}
			if affectedRows, err = stdSqlResult.RowsAffected(); err != nil {
				err = gerror.X多层错误码(gcode.CodeDbOperationError, err, `sql.Result.RowsAffected failed`)
				return stdSqlResult, err
			} else {
				batchResult.X原生sql行记录 = stdSqlResult
				batchResult.X影响行数 += affectedRows
			}
			params = params[:0]
			valueHolders = valueHolders[:0]
		}
	}
	return batchResult, nil
}

// X更新 执行表的 "UPDATE ... " 语句。
//
// 参数 `data` 可以是字符串、映射、gmap、结构体或指向结构体的指针等类型。
// 例如："uid=10000", "uid", 10000, g.Map{"uid": 10000, "name":"john"}
//
// 参数 `condition` 也可以是字符串、映射、gmap、切片或结构体及指向结构体的指针等类型。
// 常与参数 `args` 配合使用。
// 例如：
// "uid=10000",
// "uid", 10000
// "money>? AND name like ?", 99999, "vip_%"
// "status IN (?)", g.Slice{1,2,3}
// "age IN(?,?)", 18, 50
// User{ Id : 1, UserName : "john"}.
// md5:8651eb1bd7e10da0
func (c *Core) X更新(上下文 context.Context, 表名称 string, 数据 interface{}, 条件 interface{}, 参数 ...interface{}) (sql.Result, error) {
	return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X设置数据(数据).X条件(条件, 参数...).X更新()
}

// X底层更新 执行针对表的 "UPDATE ... " 语句。
// 这个函数通常用于自定义接口定义，一般不需要手动调用。
// md5:6d7e08b57dd59a0b
func (c *Core) X底层更新(上下文 context.Context, 链接 Link, 表名称 string, 值 interface{}, 条件 string, 参数 ...interface{}) (result sql.Result, err error) {
	表名称 = c.X底层添加前缀字符和引用字符(表名称)
	var (
		rv   = reflect.ValueOf(值)
		kind = rv.Kind()
	)
	if kind == reflect.Ptr {
		rv = rv.Elem()
		kind = rv.Kind()
	}
	var (
		params  []interface{}
		updates string
	)
	switch kind {
	case reflect.Map, reflect.Struct:
		var (
			fields         []string
			dataMap        map[string]interface{}
			counterHandler = func(column string, counter Counter) {
				if counter.X增减值 != 0 {
					column = c.X底层QuoteWord(column)
					var (
						columnRef = c.X底层QuoteWord(counter.X字段名称)
						columnVal = counter.X增减值
						operator  = "+"
					)
					if columnVal < 0 {
						operator = "-"
						columnVal = -columnVal
					}
					fields = append(fields, fmt.Sprintf("%s=%s%s?", column, columnRef, operator))
					params = append(params, columnVal)
				}
			}
		)
		dataMap, err = c.X底层ConvertDataForRecord(上下文, 值, 表名称)
		if err != nil {
			return nil, err
		}
				// 按照表格字段的顺序对数据键进行排序。 md5:edcdc64a514af6fa
		var (
			dataKeys       = make([]string, 0)
			keysInSequence = make([]string, 0)
		)
		for k := range dataMap {
			dataKeys = append(dataKeys, k)
		}
		keysInSequence, err = c.fieldsToSequence(上下文, 表名称, dataKeys)
		if err != nil {
			return nil, err
		}
		for _, k := range keysInSequence {
			v := dataMap[k]
			switch value := v.(type) {
			case *Counter:
				counterHandler(k, *value)

			case Counter:
				counterHandler(k, value)

			default:
				if s, ok := v.(Raw); ok {
					fields = append(fields, c.X底层QuoteWord(k)+"="+gconv.String(s))
				} else {
					fields = append(fields, c.X底层QuoteWord(k)+"=?")
					params = append(params, v)
				}
			}
		}
		updates = strings.Join(fields, ",")

	default:
		updates = gconv.String(值)
	}
	if len(updates) == 0 {
		return nil, gerror.X创建错误码(gcode.CodeMissingParameter, "data cannot be empty")
	}
	if len(params) > 0 {
		参数 = append(params, 参数...)
	}
		// 如果没有传递链接，那么它就使用主链接。 md5:02e931534071446b
	if 链接 == nil {
		if 链接, err = c.X底层MasterLink(); err != nil {
			return nil, err
		}
	}
	return c.db.X底层原生SQL执行(上下文, 链接, fmt.Sprintf(
		"UPDATE %s SET %s%s",
		表名称, updates, 条件,
	),
		参数...,
	)
}

// X删除 执行 "DELETE FROM ... " 语句针对该表。
//
// `condition` 参数可以是字符串、映射、gmap、切片、结构体或指向结构体的指针等多种类型。
// 它常与参数 `args` 一起使用。
// 例如：
// "uid=10000",
// "uid", 10000
// "money>? AND name like ?", 99999, "vip_%"
// "status IN (?)", g.Slice{1,2,3}
// "age IN(?,?)", 18, 50
// User{ Id : 1, UserName : "john"}.
// md5:c6c87830434eba7d
func (c *Core) X删除(上下文 context.Context, 表名称 string, 条件 interface{}, 参数 ...interface{}) (结果 sql.Result, 错误 error) {
	return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X条件(条件, 参数...).X删除()
}

// X底层删除 对表执行 "DELETE FROM ..." 语句。
// 此函数通常用于自定义接口定义，无需手动调用。
// md5:f902004d44b55d73
func (c *Core) X底层删除(上下文 context.Context, 链接 Link, 表名称 string, 条件 string, 参数 ...interface{}) (结果 sql.Result, 错误 error) {
	if 链接 == nil {
		if 链接, 错误 = c.X底层MasterLink(); 错误 != nil {
			return nil, 错误
		}
	}
	表名称 = c.X底层添加前缀字符和引用字符(表名称)
	return c.db.X底层原生SQL执行(上下文, 链接, fmt.Sprintf("DELETE FROM %s%s", 表名称, 条件), 参数...)
}

// X取数据库链接信息获取并返回经过过滤的`linkInfo`，这些信息可用于日志记录或跟踪目的。
// md5:5d3d4d2f55af0347
func (c *Core) X取数据库链接信息() string {
	return fmt.Sprintf(
		`%s@%s(%s:%s)/%s`,
		c.config.User, c.config.Protocol, c.config.Host, c.config.Port, c.config.Name,
	)
}

// MarshalJSON 实现了 json.Marshal 接口的MarshalJSON方法。它只是返回指针地址。
// 
// 注意，这个接口主要是为了解决 Golang 版本小于 v1.14 时的json无限循环bug而实现的。
// md5:1b2346be8e04b5fa
func (c Core) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%+v`, c)), nil
}

// writeSqlToLogger 将Sql对象输出到日志记录器。
// 仅当配置"debug"为true时，此功能才启用。
// md5:ad16123093791e59
func (c *Core) writeSqlToLogger(ctx context.Context, sql *Sql) {
	var transactionIdStr string
	if sql.IsTransaction {
		if v := ctx.Value(transactionIdForLoggerCtx); v != nil {
			transactionIdStr = fmt.Sprintf(`[txid:%d] `, v.(uint64))
		}
	}
	s := fmt.Sprintf(
		"[%3d ms] [%s] [%s] [rows:%-3d] %s%s",
		sql.End-sql.Start, sql.Group, sql.Schema, sql.RowsAffected, transactionIdStr, sql.Format,
	)
	if sql.Error != nil {
		s += "\nError: " + sql.Error.Error()
		c.logger.Error(ctx, s)
	} else {
		c.logger.X输出DEBU(ctx, s)
	}
}

// X是否存在表名 判断数据库中是否存在指定的表名。 md5:64f8bb54ba260c03
func (c *Core) X是否存在表名(表名称 string) (bool, error) {
	tables, err := c.X取表名称缓存()
	if err != nil {
		return false, err
	}
	charL, charR := c.db.X底层取数据库安全字符()
	表名称 = gstr.X过滤首尾符并含空白(表名称, charL+charR)
	for _, table := range tables {
		if table == 表名称 {
			return true, nil
		}
	}
	return false, nil
}

// X取表名称缓存 使用缓存检索并返回当前数据库中的表名。 md5:9abf0a08a0dbc629
func (c *Core) X取表名称缓存() ([]string, error) {
	var (
		ctx      = c.db.X取上下文对象()
		cacheKey = fmt.Sprintf(`Tables: %s`, c.db.X取配置组名称())
	)
	result, err := c.X取缓存对象().X取值或设置值_并发安全函数(
		ctx, cacheKey, func(ctx context.Context) (interface{}, error) {
			tableList, err := c.db.X取表名称切片(ctx)
			if err != nil {
				return false, err
			}
			return tableList, nil
		}, 0,
	)
	if err != nil {
		return nil, err
	}
	return result.X取文本切片(), nil
}

// IsSoftCreatedFieldName 检查并返回给定字段名是否为自动填充的创建时间。 md5:f4c7129bbccec8aa
func (c *Core) IsSoftCreatedFieldName(fieldName string) bool {
	if fieldName == "" {
		return false
	}
	if config := c.db.X取当前节点配置(); config.CreatedAt != "" {
		if utils.EqualFoldWithoutChars(fieldName, config.CreatedAt) {
			return true
		}
		return gstr.X切片是否存在(append([]string{config.CreatedAt}, createdFieldNames...), fieldName)
	}
	for _, v := range createdFieldNames {
		if utils.EqualFoldWithoutChars(fieldName, v) {
			return true
		}
	}
	return false
}

// X格式化Sql 在执行SQL之前格式化SQL字符串及其参数。
// 在SQL执行过程中，内部的handleArguments函数可能会被调用两次，
// 但请不必担心，这是安全且高效的。
// md5:73af1c35794cea21
func (c *Core) X格式化Sql(sql string, 参数切片 []interface{}) (新sql string, 新参数切片 []interface{}) {
	// 不要这样做，因为SQL语句中可能包含多行和注释。
	// sql = gstr.Trim(sql) 	// 删除sql字符串两侧的空白
	// sql = gstr.Replace(sql, "\n", " ") 	// 将换行符替换为单个空格
	// sql, _ = gregex.ReplaceString(`\s{2,}`, ` `, sql) 	// 替换连续两个或更多空格为单个空格
	// md5:907309db612b16e7
	return handleSliceAndStructArgsForSql(sql, 参数切片)
}
