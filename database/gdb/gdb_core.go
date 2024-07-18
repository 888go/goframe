// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

package gdb//bm:db类

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/internal/reflection"
	"github.com/gogf/gf/v2/internal/utils"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

// GetCore 返回底层的 *Core 对象。 md5:b7d2ff344b9a6a33
// ff:取Core对象
// c:
func (c *Core) GetCore() *Core {
	return c
}

// Ctx是一个链式函数，它创建并返回一个新的DB对象，它是当前DB对象的浅拷贝，并且包含了给定的上下文。
// 注意，这个返回的DB对象只能使用一次，所以不要将其分配给全局变量或长期使用的包变量。
// md5:9dfddec16d5df9f5
// ff:设置上下文并取副本
// c:
// ctx:上下文
func (c *Core) Ctx(ctx context.Context) DB {
	if ctx == nil {
		return c.db
	}
	// 它会浅复制当前的数据库，并为下一个链式操作更改其上下文。 md5:bf3a0a0f1f30a496
	var (
		err        error
		newCore    = &Core{}
		configNode = c.db.GetConfig()
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
	newCore.ctx = WithDB(ctx, newCore.db)
	newCore.ctx = c.injectInternalCtxData(newCore.ctx)
	return newCore.db
}

// GetCtx 返回当前数据库的上下文。
// 如果之前没有设置上下文，则返回 `context.Background()`。
// md5:9b56f79a5eaa891e
// ff:取上下文对象
// c:
func (c *Core) GetCtx() context.Context {
	ctx := c.ctx
	if ctx == nil {
		ctx = context.TODO()
	}
	return c.injectInternalCtxData(ctx)
}

// GetCtxTimeout 返回指定超时类型的上下文和取消函数。 md5:5d0be7078de61c6d
// ff:取超时上下文对象
// c:
// ctx:上下文
// timeoutType:超时类型
func (c *Core) GetCtxTimeout(ctx context.Context, timeoutType int) (context.Context, context.CancelFunc) {
	if ctx == nil {
		ctx = c.db.GetCtx()
	} else {
		ctx = context.WithValue(ctx, "WrappedByGetCtxTimeout", nil)
	}
	switch timeoutType {
	case ctxTimeoutTypeExec:
		if c.db.GetConfig().ExecTimeout > 0 {
			return context.WithTimeout(ctx, c.db.GetConfig().ExecTimeout)
		}
	case ctxTimeoutTypeQuery:
		if c.db.GetConfig().QueryTimeout > 0 {
			return context.WithTimeout(ctx, c.db.GetConfig().QueryTimeout)
		}
	case ctxTimeoutTypePrepare:
		if c.db.GetConfig().PrepareTimeout > 0 {
			return context.WithTimeout(ctx, c.db.GetConfig().PrepareTimeout)
		}
	default:
		panic(gerror.NewCodef(gcode.CodeInvalidParameter, "invalid context timeout type: %d", timeoutType))
	}
	return ctx, func() {}
}

// Close 方法关闭数据库并阻止新的查询开始。
// 随后，Close 会等待所有已在服务器上开始处理的查询完成。
//
// 关闭 DB 实例是很少见的操作，因为 DB 处理句柄设计为长生命周期的，
// 并且旨在多个 goroutine 间共享。
// md5:39e5c90e1da0ee5e
// ff:关闭数据库
// c:
// ctx:上下文
// err:错误
func (c *Core) Close(ctx context.Context) (err error) {
	if err = c.cache.Close(ctx); err != nil {
		return err
	}
	c.links.LockFunc(func(m map[any]any) {
		for k, v := range m {
			if db, ok := v.(*sql.DB); ok {
				err = db.Close()
				if err != nil {
					err = gerror.WrapCode(gcode.CodeDbOperationError, err, `db.Close failed`)
				}
				intlog.Printf(ctx, `close link: %s, err: %v`, k, err)
				if err != nil {
					return
				}
				delete(m, k)
			}
		}
	})
	return
}

// Master 如果配置了主从节点，则创建并返回一个与主节点的连接。
// 如果未配置主从节点，则返回默认连接。
// md5:0bd77b596cdae9a3
// ff:取主节点对象
// c:
// schema:数据库名称
func (c *Core) Master(schema ...string) (*sql.DB, error) {
	var (
		usedSchema   = gutil.GetOrDefaultStr(c.schema, schema...)
		charL, charR = c.db.GetChars()
	)
	return c.getSqlDb(true, gstr.Trim(usedSchema, charL+charR))
}

// 如果配置了主从模式，Slave 会创建并返回一个从节点连接。如果没有配置主从模式，则返回默认连接。
// md5:d92640050cf063d3
// ff:取从节点对象
// c:
// schema:数据库名称
func (c *Core) Slave(schema ...string) (*sql.DB, error) {
	var (
		usedSchema   = gutil.GetOrDefaultStr(c.schema, schema...)
		charL, charR = c.db.GetChars()
	)
	return c.getSqlDb(false, gstr.Trim(usedSchema, charL+charR))
}

// GetAll 从数据库中查询并返回数据记录。 md5:dfdcfddaa70ab1d4
// ff:GetAll别名
// c:
// ctx:上下文
// sql:
// args:参数
// Result:
func (c *Core) GetAll(ctx context.Context, sql string, args ...interface{}) (Result, error) {
	return c.db.DoSelect(ctx, nil, sql, args...)
}

// DoSelect 从数据库查询并返回数据记录。 md5:82b06146b8d539d1
// ff:底层查询
// c:
// ctx:上下文
// link:链接
// sql:
// args:参数
// result:结果
// err:错误
func (c *Core) DoSelect(ctx context.Context, link Link, sql string, args ...interface{}) (result Result, err error) {
	return c.db.DoQuery(ctx, link, sql, args...)
}

// GetOne 从数据库中查询并返回一条记录。 md5:9552f7e095f58141
// ff:原生SQL查询单条记录
// c:
// ctx:上下文
// sql:
// args:参数
// Record:
func (c *Core) GetOne(ctx context.Context, sql string, args ...interface{}) (Record, error) {
	list, err := c.db.GetAll(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	return nil, nil
}

// GetArray 从数据库查询并返回数据值作为切片。
// 注意，如果结果中有多个列，它会随机返回一列的值。
// md5:b81cd4c5e063a6f2
// ff:原生SQL查询切片
// c:
// ctx:上下文
// sql:
// args:参数
func (c *Core) GetArray(ctx context.Context, sql string, args ...interface{}) ([]Value, error) {
	all, err := c.db.DoSelect(ctx, nil, sql, args...)
	if err != nil {
		return nil, err
	}
	return all.Array(), nil
}

// doGetStruct 从数据库查询一条记录，并将其转换为给定的结构体。
// 参数 `pointer` 应该是指向结构体的指针。
// md5:9260d4f62deef626
func (c *Core) doGetStruct(ctx context.Context, pointer interface{}, sql string, args ...interface{}) error {
	one, err := c.db.GetOne(ctx, sql, args...)
	if err != nil {
		return err
	}
	return one.Struct(pointer)
}

// doGetStructs 从数据库查询记录并将其转换为给定的结构体。
// 参数 `pointer` 应为结构体切片类型：[]struct/[]*struct。
// md5:4ce864edda9b9231
func (c *Core) doGetStructs(ctx context.Context, pointer interface{}, sql string, args ...interface{}) error {
	all, err := c.db.GetAll(ctx, sql, args...)
	if err != nil {
		return err
	}
	return all.Structs(pointer)
}

// GetScan 从数据库查询一个或多个记录，并将它们转换为给定的结构体或结构体数组。
//
// 如果参数 `pointer` 是结构体指针类型，它内部会调用 GetStruct 进行转换。如果参数 `pointer` 是切片类型，它内部会调用 GetStructs 进行转换。
// md5:c1dbdab7a7c29c51
// ff:原生SQL查询到结构体指针
// c:
// ctx:上下文
// pointer:结构体指针
// sql:
// args:参数
func (c *Core) GetScan(ctx context.Context, pointer interface{}, sql string, args ...interface{}) error {
	reflectInfo := reflection.OriginTypeAndKind(pointer)
	if reflectInfo.InputKind != reflect.Ptr {
		return gerror.NewCodef(
			gcode.CodeInvalidParameter,
			"params should be type of pointer, but got: %v",
			reflectInfo.InputKind,
		)
	}
	switch reflectInfo.OriginKind {
	case reflect.Array, reflect.Slice:
		return c.db.GetCore().doGetStructs(ctx, pointer, sql, args...)

	case reflect.Struct:
		return c.db.GetCore().doGetStruct(ctx, pointer, sql, args...)
	}
	return gerror.NewCodef(
		gcode.CodeInvalidParameter,
		`in valid parameter type "%v", of which element type should be type of struct/slice`,
		reflectInfo.InputType,
	)
}

// GetValue 从数据库查询并返回字段值。
// SQL 应该只查询数据库中的一个字段，否则它将只返回结果中的一个字段。
// md5:96794360fadbc288
// ff:原生SQL查询字段值
// c:
// ctx:上下文
// sql:
// args:参数
// Value:
func (c *Core) GetValue(ctx context.Context, sql string, args ...interface{}) (Value, error) {
	one, err := c.db.GetOne(ctx, sql, args...)
	if err != nil {
		return gvar.New(nil), err
	}
	for _, v := range one {
		return v, nil
	}
	return gvar.New(nil), nil
}

// GetCount 从数据库中查询并返回计数。 md5:a8368d39f4a58979
// ff:原生SQL查询字段计数
// c:
// ctx:上下文
// sql:
// args:参数
func (c *Core) GetCount(ctx context.Context, sql string, args ...interface{}) (int, error) {
// 如果查询字段中不包含"COUNT"函数，
// 则替换SQL字符串，并在字段中添加"COUNT"函数。
// md5:624b6da82fb9facd
	if !gregex.IsMatchString(`(?i)SELECT\s+COUNT\(.+\)\s+FROM`, sql) {
		sql, _ = gregex.ReplaceString(`(?i)(SELECT)\s+(.+)\s+(FROM)`, `$1 COUNT($2) $3`, sql)
	}
	value, err := c.db.GetValue(ctx, sql, args...)
	if err != nil {
		return 0, err
	}
	return value.Int(), nil
}

// Union 执行 "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." 语句。 md5:6a2f9809c172cb31
// ff:多表去重查询
// c:
// unions:Model对象
func (c *Core) Union(unions ...*Model) *Model {
	var ctx = c.db.GetCtx()
	return c.doUnion(ctx, unionTypeNormal, unions...)
}

// UnionAll 生成 "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ... " 语句。 md5:5a15c8720fcb152f
// ff:多表查询
// c:
// unions:Model对象
func (c *Core) UnionAll(unions ...*Model) *Model {
	var ctx = c.db.GetCtx()
	return c.doUnion(ctx, unionTypeAll, unions...)
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
	return c.db.Raw(composedSqlStr, composedArgs...)
}

// PingMaster 向主节点发送请求以检查身份验证或保持连接活动。 md5:47a7df55cbee8583
// ff:向主节点发送心跳
// c:
func (c *Core) PingMaster() error {
	var ctx = c.db.GetCtx()
	if master, err := c.db.Master(); err != nil {
		return err
	} else {
		if err = master.PingContext(ctx); err != nil {
			err = gerror.WrapCode(gcode.CodeDbOperationError, err, `master.Ping failed`)
		}
		return err
	}
}

// PingSlave 调用ping命令检查从节点的认证或者维持连接。 md5:62272b38d874eda6
// ff:向从节点发送心跳
// c:
func (c *Core) PingSlave() error {
	var ctx = c.db.GetCtx()
	if slave, err := c.db.Slave(); err != nil {
		return err
	} else {
		if err = slave.PingContext(ctx); err != nil {
			err = gerror.WrapCode(gcode.CodeDbOperationError, err, `slave.Ping failed`)
		}
		return err
	}
}

// Insert does "INSERT INTO ..." statement for the table.
// If there's already one unique record of the data in the table, it returns error.
//
// The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc.
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// The parameter `batch` specifies the batch operation count when given data is slice.
// ff:插入
// c:
// ctx:上下文
// table:表名称
// data:值
// batch:批量操作行数
func (c *Core) Insert(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error) {
	if len(batch) > 0 {
		return c.Model(table).Ctx(ctx).Data(data).Batch(batch[0]).Insert()
	}
	return c.Model(table).Ctx(ctx).Data(data).Insert()
}

// InsertIgnore does "INSERT IGNORE INTO ..." statement for the table.
// If there's already one unique record of the data in the table, it ignores the inserting.
//
// The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc.
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// The parameter `batch` specifies the batch operation count when given data is slice.
// ff:插入并跳过已存在
// c:
// ctx:上下文
// table:表名称
// data:值
// batch:批量操作行数
func (c *Core) InsertIgnore(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error) {
	if len(batch) > 0 {
		return c.Model(table).Ctx(ctx).Data(data).Batch(batch[0]).InsertIgnore()
	}
	return c.Model(table).Ctx(ctx).Data(data).InsertIgnore()
}

// InsertAndGetId 执行插入操作，并返回自动生成的最后一个插入id。 md5:8d00b40a35fa48a5
// ff:插入并取ID
// c:
// ctx:上下文
// table:表名称
// data:值
// batch:批量操作行数
func (c *Core) InsertAndGetId(ctx context.Context, table string, data interface{}, batch ...int) (int64, error) {
	if len(batch) > 0 {
		return c.Model(table).Ctx(ctx).Data(data).Batch(batch[0]).InsertAndGetId()
	}
	return c.Model(table).Ctx(ctx).Data(data).InsertAndGetId()
}

// Replace does "REPLACE INTO ..." statement for the table.
// If there's already one unique record of the data in the table, it deletes the record
// and inserts a new one.
//
// The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc.
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc.
// If given data is type of slice, it then does batch replacing, and the optional parameter
// `batch` specifies the batch operation count.
// ff:插入并替换已存在
// c:
// ctx:上下文
// table:表名称
// data:值
// batch:批量操作行数
func (c *Core) Replace(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error) {
	if len(batch) > 0 {
		return c.Model(table).Ctx(ctx).Data(data).Batch(batch[0]).Replace()
	}
	return c.Model(table).Ctx(ctx).Data(data).Replace()
}

// Save does "INSERT INTO ... ON DUPLICATE KEY UPDATE..." statement for the table.
// It updates the record if there's primary or unique index in the saving data,
// or else it inserts a new record into the table.
//
// The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc.
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// If given data is type of slice, it then does batch saving, and the optional parameter
// `batch` specifies the batch operation count.
// ff:插入并更新已存在
// c:
// ctx:上下文
// table:表名称
// data:值
// batch:批量操作行数
func (c *Core) Save(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error) {
	if len(batch) > 0 {
		return c.Model(table).Ctx(ctx).Data(data).Batch(batch[0]).Save()
	}
	return c.Model(table).Ctx(ctx).Data(data).Save()
}

func (c *Core) fieldsToSequence(ctx context.Context, table string, fields []string) ([]string, error) {
	var (
		fieldSet               = gset.NewStrSetFrom(fields)
		fieldsResultInSequence = make([]string, 0)
		tableFields, err       = c.db.TableFields(ctx, table)
	)
	if err != nil {
		return nil, err
	}
	// 按顺序对字段进行排序。 md5:3edaf791b6d06284
	var fieldsOfTableInSequence = make([]string, len(tableFields))
	for _, field := range tableFields {
		fieldsOfTableInSequence[field.Index] = field.Name
	}
	// Sort the input fields.
	for _, fieldName := range fieldsOfTableInSequence {
		if fieldSet.Contains(fieldName) {
			fieldsResultInSequence = append(fieldsResultInSequence, fieldName)
		}
	}
	return fieldsResultInSequence, nil
}

// DoInsert inserts or updates data for given table.
// This function is usually used for custom interface definition, you do not need call it manually.
// The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc.
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// The parameter `option` values are as follows:
// ff:底层插入
// c:
// ctx:上下文
// link:链接
// table:表名称
// list:
// option:
// result:
// err:
func (c *Core) DoInsert(ctx context.Context, link Link, table string, list List, option DoInsertOption) (result sql.Result, err error) {
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
	var keyListMap = gmap.NewListMap()
	for _, item := range list {
		var (
			tmpKeys              = make([]string, 0)
			tmpKeysInSequenceStr string
		)
		for k := range item {
			tmpKeys = append(tmpKeys, k)
		}
		keys, err = c.fieldsToSequence(ctx, table, tmpKeys)
		if err != nil {
			return nil, err
		}
		tmpKeysInSequenceStr = gstr.Join(keys, ",")
		if !keyListMap.Contains(tmpKeysInSequenceStr) {
			keyListMap.Set(tmpKeysInSequenceStr, make(List, 0))
		}
		tmpKeysInSequenceList := keyListMap.Get(tmpKeysInSequenceStr).(List)
		tmpKeysInSequenceList = append(tmpKeysInSequenceList, item)
		keyListMap.Set(tmpKeysInSequenceStr, tmpKeysInSequenceList)
	}
	if keyListMap.Size() > 1 {
		var (
			tmpResult    sql.Result
			sqlResult    SqlResult
			rowsAffected int64
		)
		keyListMap.Iterator(func(key, value interface{}) bool {
			tmpResult, err = c.DoInsert(ctx, link, table, value.(List), option)
			if err != nil {
				return false
			}
			rowsAffected, err = tmpResult.RowsAffected()
			if err != nil {
				return false
			}
			sqlResult.Result = tmpResult
			sqlResult.Affected += rowsAffected
			return true
		})
		return &sqlResult, err
	}

	// 准备批量结果指针。 md5:dfc8aa8bb292f9d5
	var (
		charL, charR = c.db.GetChars()
		batchResult  = new(SqlResult)
		keysStr      = charL + strings.Join(keys, charR+","+charL) + charR
		operation    = GetInsertOperationByOption(option.InsertOption)
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
		valueHolders = append(valueHolders, "("+gstr.Join(values, ",")+")")
		// 批量包检查：它满足批量数量，或者它是最后一个元素。 md5:a2ef8b869c6d8888
		if len(valueHolders) == option.BatchCount || (i == listLength-1 && len(valueHolders) > 0) {
			var (
				stdSqlResult sql.Result
				affectedRows int64
			)
			stdSqlResult, err = c.db.DoExec(ctx, link, fmt.Sprintf(
				"%s INTO %s(%s) VALUES%s %s",
				operation, c.QuotePrefixTableName(table), keysStr,
				gstr.Join(valueHolders, ","),
				onDuplicateStr,
			), params...)
			if err != nil {
				return stdSqlResult, err
			}
			if affectedRows, err = stdSqlResult.RowsAffected(); err != nil {
				err = gerror.WrapCode(gcode.CodeDbOperationError, err, `sql.Result.RowsAffected failed`)
				return stdSqlResult, err
			} else {
				batchResult.Result = stdSqlResult
				batchResult.Affected += affectedRows
			}
			params = params[:0]
			valueHolders = valueHolders[:0]
		}
	}
	return batchResult, nil
}

// Update does "UPDATE ... " statement for the table.
//
// The parameter `data` can be type of string/map/gmap/struct/*struct, etc. "uid=10000", "uid", 10000, g.Map{"uid": 10000, "name":"john"}
//
// The parameter `condition` can be type of string/map/gmap/slice/struct/*struct, etc.
// It is commonly used with parameter `args`.
// "uid=10000",
// "uid", 10000
// "money>? AND name like ?", 99999, "vip_%"
// "status IN (?)", g.Slice{1,2,3}
// "age IN(?,?)", 18, 50
// User{ Id : 1, UserName : "john"}.
// ff:更新
// c:
// ctx:上下文
// table:表名称
// data:数据
// condition:条件
// args:参数
func (c *Core) Update(ctx context.Context, table string, data interface{}, condition interface{}, args ...interface{}) (sql.Result, error) {
	return c.Model(table).Ctx(ctx).Data(data).Where(condition, args...).Update()
}

// DoUpdate 执行针对表的 "UPDATE ... " 语句。
// 这个函数通常用于自定义接口定义，一般不需要手动调用。
// md5:6d7e08b57dd59a0b
// ff:底层更新
// c:
// ctx:上下文
// link:链接
// table:表名称
// data:值
// condition:条件
// args:参数
// result:
// err:
func (c *Core) DoUpdate(ctx context.Context, link Link, table string, data interface{}, condition string, args ...interface{}) (result sql.Result, err error) {
	table = c.QuotePrefixTableName(table)
	var (
		rv   = reflect.ValueOf(data)
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
				if counter.Value != 0 {
					column = c.QuoteWord(column)
					var (
						columnRef = c.QuoteWord(counter.Field)
						columnVal = counter.Value
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
		dataMap, err = c.ConvertDataForRecord(ctx, data, table)
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
		keysInSequence, err = c.fieldsToSequence(ctx, table, dataKeys)
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
					fields = append(fields, c.QuoteWord(k)+"="+gconv.String(s))
				} else {
					fields = append(fields, c.QuoteWord(k)+"=?")
					params = append(params, v)
				}
			}
		}
		updates = strings.Join(fields, ",")

	default:
		updates = gconv.String(data)
	}
	if len(updates) == 0 {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "data cannot be empty")
	}
	if len(params) > 0 {
		args = append(params, args...)
	}
	// 如果没有传递链接，那么它就使用主链接。 md5:02e931534071446b
	if link == nil {
		if link, err = c.MasterLink(); err != nil {
			return nil, err
		}
	}
	return c.db.DoExec(ctx, link, fmt.Sprintf(
		"UPDATE %s SET %s%s",
		table, updates, condition,
	),
		args...,
	)
}

// Delete does "DELETE FROM ... " statement for the table.
//
// The parameter `condition` can be type of string/map/gmap/slice/struct/*struct, etc.
// It is commonly used with parameter `args`.
// "uid=10000",
// "uid", 10000
// "money>? AND name like ?", 99999, "vip_%"
// "status IN (?)", g.Slice{1,2,3}
// "age IN(?,?)", 18, 50
// User{ Id : 1, UserName : "john"}.
// ff:删除
// c:
// ctx:上下文
// table:表名称
// condition:条件
// args:参数
// result:结果
// err:错误
func (c *Core) Delete(ctx context.Context, table string, condition interface{}, args ...interface{}) (result sql.Result, err error) {
	return c.Model(table).Ctx(ctx).Where(condition, args...).Delete()
}

// DoDelete 对表执行 "DELETE FROM ..." 语句。
// 此函数通常用于自定义接口定义，无需手动调用。
// md5:f902004d44b55d73
// ff:底层删除
// c:
// ctx:上下文
// link:链接
// table:表名称
// condition:条件
// args:参数
// result:结果
// err:错误
func (c *Core) DoDelete(ctx context.Context, link Link, table string, condition string, args ...interface{}) (result sql.Result, err error) {
	if link == nil {
		if link, err = c.MasterLink(); err != nil {
			return nil, err
		}
	}
	table = c.QuotePrefixTableName(table)
	return c.db.DoExec(ctx, link, fmt.Sprintf("DELETE FROM %s%s", table, condition), args...)
}

// FilteredLink获取并返回经过过滤的`linkInfo`，这些信息可用于日志记录或跟踪目的。
// md5:5d3d4d2f55af0347
// ff:取数据库链接信息
// c:
func (c *Core) FilteredLink() string {
	return fmt.Sprintf(
		`%s@%s(%s:%s)/%s`,
		c.config.User, c.config.Protocol, c.config.Host, c.config.Port, c.config.Name,
	)
}

// MarshalJSON 实现了 json.Marshal 接口的MarshalJSON方法。它只是返回指针地址。
// 
// 注意，这个接口主要是为了解决 Golang 版本小于 v1.14 时的json无限循环bug而实现的。
// md5:1b2346be8e04b5fa
// ff:
// c:
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
		c.logger.Debug(ctx, s)
	}
}

// HasTable 判断数据库中是否存在指定的表名。 md5:64f8bb54ba260c03
// ff:是否存在表名
// c:
// name:表名称
func (c *Core) HasTable(name string) (bool, error) {
	tables, err := c.GetTablesWithCache()
	if err != nil {
		return false, err
	}
	charL, charR := c.db.GetChars()
	name = gstr.Trim(name, charL+charR)
	for _, table := range tables {
		if table == name {
			return true, nil
		}
	}
	return false, nil
}

// GetTablesWithCache 使用缓存检索并返回当前数据库中的表名。 md5:9abf0a08a0dbc629
// ff:取表名称缓存
// c:
func (c *Core) GetTablesWithCache() ([]string, error) {
	var (
		ctx      = c.db.GetCtx()
		cacheKey = fmt.Sprintf(`Tables: %s`, c.db.GetGroup())
	)
	result, err := c.GetCache().GetOrSetFuncLock(
		ctx, cacheKey, func(ctx context.Context) (interface{}, error) {
			tableList, err := c.db.Tables(ctx)
			if err != nil {
				return false, err
			}
			return tableList, nil
		}, 0,
	)
	if err != nil {
		return nil, err
	}
	return result.Strings(), nil
}

// IsSoftCreatedFieldName 检查并返回给定字段名是否为自动填充的创建时间。 md5:f4c7129bbccec8aa
// ff:
// c:
// fieldName:
func (c *Core) IsSoftCreatedFieldName(fieldName string) bool {
	if fieldName == "" {
		return false
	}
	if config := c.db.GetConfig(); config.CreatedAt != "" {
		if utils.EqualFoldWithoutChars(fieldName, config.CreatedAt) {
			return true
		}
		return gstr.InArray(append([]string{config.CreatedAt}, createdFieldNames...), fieldName)
	}
	for _, v := range createdFieldNames {
		if utils.EqualFoldWithoutChars(fieldName, v) {
			return true
		}
	}
	return false
}

// FormatSqlBeforeExecuting 在执行SQL之前格式化SQL字符串及其参数。
// 在SQL执行过程中，内部的handleArguments函数可能会被调用两次，
// 但请不必担心，这是安全且高效的。
// md5:73af1c35794cea21
// ff:格式化Sql
// c:
// sql:
// args:参数切片
// newSql:新sql
// newArgs:新参数切片
func (c *Core) FormatSqlBeforeExecuting(sql string, args []interface{}) (newSql string, newArgs []interface{}) {
// 不要这样做，因为SQL语句中可能包含多行和注释。
// sql = gstr.Trim(sql) // 删除sql字符串两侧的空白
// sql = gstr.Replace(sql, "\n", " ") // 将换行符替换为单个空格
// sql, _ = gregex.ReplaceString(`\s{2,}`, ` `, sql) // 替换连续两个或更多空格为单个空格
// md5:907309db612b16e7
	return handleSliceAndStructArgsForSql(sql, args)
}
