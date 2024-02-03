// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//

package gdb

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/reflection"
	"github.com/888go/goframe/internal/utils"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
)

// GetCore 返回底层的 *Core 对象。
func (c *Core) GetCore() *Core {
	return c
}

// Ctx 是一个链式函数，它创建并返回一个新的 DB 对象，该对象是对当前 DB 对象的浅复制，并且其中包含给定的上下文。
// 注意，返回的这个 DB 对象只能使用一次，所以不要将其赋值给全局或包级别的变量以长期使用。
func (c *Core) Ctx(ctx context.Context) DB {
	if ctx == nil {
		return c.db
	}
	// 它对当前db进行浅复制，并更改其上下文以进行下一个链式操作。
	var (
		err        error
		newCore    = &Core{}
		configNode = c.db.GetConfig()
	)
	*newCore = *c
	// 它创建一个新的DB对象（非新连接），这个对象通常是对`Core`对象的一个包装。
	newCore.db, err = driverMap[configNode.Type].New(newCore, configNode)
	if err != nil {
// 这里确实是一个严重的错误。
// 不要让它继续执行。
		panic(err)
	}
	newCore.ctx = WithDB(ctx, newCore.db)
	newCore.ctx = c.InjectInternalCtxData(newCore.ctx)
	return newCore.db
}

// GetCtx 返回当前数据库的上下文。
// 如果之前未设置上下文，则返回 `context.Background()`。
func (c *Core) GetCtx() context.Context {
	ctx := c.ctx
	if ctx == nil {
		ctx = context.TODO()
	}
	return c.InjectInternalCtxData(ctx)
}

// GetCtxTimeout 根据指定的超时类型返回上下文和取消函数。
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

// Close 关闭数据库并阻止新的查询开始。
// Close 会等待在服务器上已经开始处理的所有查询完成。
//
// 很少会关闭一个 DB，因为 DB 处理程序旨在长期存在并被多个 goroutine 共享。
func (c *Core) Close(ctx context.Context) (err error) {
	if err = c.cache.Close(ctx); err != nil {
		return err
	}
	c.links.LockFunc(func(m map[string]interface{}) {
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

// Master在主从配置的情况下，创建并从主节点返回一个连接。如果未配置主从，则返回默认连接。
func (c *Core) Master(schema ...string) (*sql.DB, error) {
	var (
		usedSchema   = gutil.GetOrDefaultStr(c.schema, schema...)
		charL, charR = c.db.GetChars()
	)
	return c.getSqlDb(true, gstr.Trim(usedSchema, charL+charR))
}

// Slave在主从配置的情况下，创建并返回从节点的连接。如果未配置主从，则返回默认连接。
func (c *Core) Slave(schema ...string) (*sql.DB, error) {
	var (
		usedSchema   = gutil.GetOrDefaultStr(c.schema, schema...)
		charL, charR = c.db.GetChars()
	)
	return c.getSqlDb(false, gstr.Trim(usedSchema, charL+charR))
}

// GetAll 从数据库查询并返回数据记录。
func (c *Core) GetAll(ctx context.Context, sql string, args ...interface{}) (Result, error) {
	return c.db.DoSelect(ctx, nil, sql, args...)
}

// DoSelect 从数据库查询并返回数据记录。
func (c *Core) DoSelect(ctx context.Context, link Link, sql string, args ...interface{}) (result Result, err error) {
	return c.db.DoQuery(ctx, link, sql, args...)
}

// GetOne 从数据库查询并返回一条记录。
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
// 注意，如果结果中有多个列，则它会随机返回其中一列的值。
func (c *Core) GetArray(ctx context.Context, sql string, args ...interface{}) ([]Value, error) {
	all, err := c.db.DoSelect(ctx, nil, sql, args...)
	if err != nil {
		return nil, err
	}
	return all.Array(), nil
}

// doGetStruct 从数据库查询一条记录并将其转换为给定的结构体。
// 参数 `pointer` 应该是指向结构体的指针。
func (c *Core) doGetStruct(ctx context.Context, pointer interface{}, sql string, args ...interface{}) error {
	one, err := c.db.GetOne(ctx, sql, args...)
	if err != nil {
		return err
	}
	return one.Struct(pointer)
}

// doGetStructs 从数据库查询记录并将其转换为给定的结构体。
// 参数 `pointer` 应为结构体切片类型：[]struct 或 []*struct。
func (c *Core) doGetStructs(ctx context.Context, pointer interface{}, sql string, args ...interface{}) error {
	all, err := c.db.GetAll(ctx, sql, args...)
	if err != nil {
		return err
	}
	return all.Structs(pointer)
}

// GetScan 从数据库查询一个或多个记录，并将它们转换为给定的结构体或结构体数组。
//
// 如果参数`pointer`是结构体指针类型，它会内部调用 GetStruct 进行转换。如果参数 `pointer` 是切片类型，则会内部调用 GetStructs 进行转换。
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
// SQL语句应当只查询数据库中的一个字段，否则它将仅返回结果中的一个字段。
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

// GetCount 从数据库查询并返回计数。
func (c *Core) GetCount(ctx context.Context, sql string, args ...interface{}) (int, error) {
// 如果查询字段中不包含函数"COUNT"，
// 则替换sql字符串，并在字段中添加"COUNT"函数。
// 这段代码的注释是说，当SQL查询语句中的字段部分未使用“COUNT”函数时，会对原始的sql字符串进行替换处理，将“COUNT”函数添加到字段表达式中。
	if !gregex.IsMatchString(`(?i)SELECT\s+COUNT\(.+\)\s+FROM`, sql) {
		sql, _ = gregex.ReplaceString(`(?i)(SELECT)\s+(.+)\s+(FROM)`, `$1 COUNT($2) $3`, sql)
	}
	value, err := c.db.GetValue(ctx, sql, args...)
	if err != nil {
		return 0, err
	}
	return value.Int(), nil
}

// Union 执行 "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." 语句。
func (c *Core) Union(unions ...*Model) *Model {
	var ctx = c.db.GetCtx()
	return c.doUnion(ctx, unionTypeNormal, unions...)
}

// UnionAll 执行 "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ..." 语句。
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

// PingMaster 用于向主节点发送心跳以检查身份验证或保持连接存活。
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

// PingSlave 向从节点发送ping请求，用于检查身份验证或保持连接活跃。
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

// Insert 执行针对该表的 "INSERT INTO ..." 语句。
// 如果表中已经存在一条相同数据的唯一记录，则返回错误。
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 示例：
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// 当给定的数据为切片时，参数 `batch` 指定了批量操作的数量。
func (c *Core) Insert(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error) {
	if len(batch) > 0 {
		return c.Model(table).Ctx(ctx).Data(data).Batch(batch[0]).Insert()
	}
	return c.Model(table).Ctx(ctx).Data(data).Insert()
}

// InsertIgnore 执行针对表的 "INSERT IGNORE INTO ..." 语句。
// 如果表中已存在一条相同的数据记录，它将忽略插入操作。
//
// 参数 `data` 可以为 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 例如：
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// 当给定数据为切片时，参数 `batch` 指定批量操作的数量。
func (c *Core) InsertIgnore(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error) {
	if len(batch) > 0 {
		return c.Model(table).Ctx(ctx).Data(data).Batch(batch[0]).InsertIgnore()
	}
	return c.Model(table).Ctx(ctx).Data(data).InsertIgnore()
}

// InsertAndGetId 执行插入操作，并返回自动生成的最后一个插入ID。
func (c *Core) InsertAndGetId(ctx context.Context, table string, data interface{}, batch ...int) (int64, error) {
	if len(batch) > 0 {
		return c.Model(table).Ctx(ctx).Data(data).Batch(batch[0]).InsertAndGetId()
	}
	return c.Model(table).Ctx(ctx).Data(data).InsertAndGetId()
}

// Replace 执行针对该表的 "REPLACE INTO ..." 语句。
// 如果表中已存在一条唯一数据记录，它会先删除这条记录，然后插入一条新的记录。
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 示例：
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 若给定的数据是切片类型，它将执行批量替换操作，可选参数
// `batch` 指定了批量操作的数量。
func (c *Core) Replace(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error) {
	if len(batch) > 0 {
		return c.Model(table).Ctx(ctx).Data(data).Batch(batch[0]).Replace()
	}
	return c.Model(table).Ctx(ctx).Data(data).Replace()
}

// Save 执行针对表的 "INSERT INTO ... ON DUPLICATE KEY UPDATE..." 语句。
// 如果保存数据中存在主键或唯一索引，它将更新记录，否则将在表中插入新的记录。
//
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 例如：
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// 如果给定的数据是切片类型，那么它将执行批量保存操作，可选参数
// `batch` 指定了批量操作的数量。
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
	// 按照顺序对字段进行排序。
	var fieldsOfTableInSequence = make([]string, len(tableFields))
	for _, field := range tableFields {
		fieldsOfTableInSequence[field.Index] = field.Name
	}
	// 对输入字段进行排序。
	for _, fieldName := range fieldsOfTableInSequence {
		if fieldSet.Contains(fieldName) {
			fieldsResultInSequence = append(fieldsResultInSequence, fieldName)
		}
	}
	return fieldsResultInSequence, nil
}

// DoInsert 插入或更新给定表中的数据。
// 该函数通常用于自定义接口定义，您无需手动调用它。
// 参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型的。
// 示例：
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
// 参数 `option` 的取值如下：
// InsertOptionDefault：仅插入，如果数据中存在唯一/主键，则返回错误；
// InsertOptionReplace：如果数据中存在唯一/主键，先从表中删除，再插入新的数据；
// InsertOptionSave：如果数据中存在唯一/主键，则更新该记录，否则插入新记录；
// InsertOptionIgnore：如果数据中存在唯一/主键，则忽略插入操作。
func (c *Core) DoInsert(ctx context.Context, link Link, table string, list List, option DoInsertOption) (result sql.Result, err error) {
	var (
		keys           []string      // Field names.
		values         []string      // 值持有者字符串数组，例如：(?,?,?)
		params         []interface{} // 这些值将会被提交到底层数据库驱动中。
		onDuplicateStr string        // onDuplicateStr 用于 "ON DUPLICATE KEY UPDATE" 语句中。
	)
// ============================================================================================
// 根据字段对列表进行分组。不同的字段将数据分到不同的列表中。
// 这里使用ListMap来保持数据插入时的顺序。
// ============================================================================================
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

	// 准备批量结果指针。
	var (
		charL, charR = c.db.GetChars()
		batchResult  = new(SqlResult)
		keysStr      = charL + strings.Join(keys, charR+","+charL) + charR
		operation    = GetInsertOperationByOption(option.InsertOption)
	)
	if option.InsertOption == InsertOptionSave {
		onDuplicateStr = c.formatOnDuplicate(keys, option)
	}
	var (
		listLength  = len(list)
		valueHolder = make([]string, 0)
	)
	for i := 0; i < listLength; i++ {
		values = values[:0]
// 请注意，map类型是无序的，
// 所以应当使用切片+键来获取值。
		for _, k := range keys {
			if s, ok := list[i][k].(Raw); ok {
				values = append(values, gconv.String(s))
			} else {
				values = append(values, "?")
				params = append(params, list[i][k])
			}
		}
		valueHolder = append(valueHolder, "("+gstr.Join(values, ",")+")")
		// 批量校验包：满足批量数量，或者已是最后一个元素。
		if len(valueHolder) == option.BatchCount || (i == listLength-1 && len(valueHolder) > 0) {
			var (
				stdSqlResult sql.Result
				affectedRows int64
			)
			stdSqlResult, err = c.db.DoExec(ctx, link, fmt.Sprintf(
				"%s INTO %s(%s) VALUES%s %s",
				operation, c.QuotePrefixTableName(table), keysStr,
				gstr.Join(valueHolder, ","),
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
			valueHolder = valueHolder[:0]
		}
	}
	return batchResult, nil
}

func (c *Core) formatOnDuplicate(columns []string, option DoInsertOption) string {
	var onDuplicateStr string
	if option.OnDuplicateStr != "" {
		onDuplicateStr = option.OnDuplicateStr
	} else if len(option.OnDuplicateMap) > 0 {
		for k, v := range option.OnDuplicateMap {
			if len(onDuplicateStr) > 0 {
				onDuplicateStr += ","
			}
			switch v.(type) {
			case Raw, *Raw:
				onDuplicateStr += fmt.Sprintf(
					"%s=%s",
					c.QuoteWord(k),
					v,
				)
			default:
				onDuplicateStr += fmt.Sprintf(
					"%s=VALUES(%s)",
					c.QuoteWord(k),
					c.QuoteWord(gconv.String(v)),
				)
			}
		}
	} else {
		for _, column := range columns {
			// 如果是保存操作，不自动更新创建时间。
			if c.isSoftCreatedFieldName(column) {
				continue
			}
			if len(onDuplicateStr) > 0 {
				onDuplicateStr += ","
			}
			onDuplicateStr += fmt.Sprintf(
				"%s=VALUES(%s)",
				c.QuoteWord(column),
				c.QuoteWord(column),
			)
		}
	}
	return InsertOnDuplicateKeyUpdate + " " + onDuplicateStr
}

// Update 执行针对该表的 "UPDATE ... " 语句。
//
// 参数 `data` 可以为 string/map/gmap/struct/*struct 等类型。
// 例如："uid=10000"、"uid", 10000、g.Map{"uid": 10000, "name":"john"}
//
// 参数 `condition` 可以为 string/map/gmap/slice/struct/*struct 等类型，通常与参数 `args` 一起使用。
// 例如：
// "uid=10000"
// "uid", 10000
// "money>? AND name like ?", 99999, "vip_%"
// "status IN (?)", g.Slice{1,2,3}
// "age IN(?,?)", 18, 50
// User{ Id : 1, UserName : "john"}
// 注：这里的 `g.Map` 和 `g.Slice` 是一种特定的 Go 语言数据结构（可能是自定义类型），分别代表映射和切片。
func (c *Core) Update(ctx context.Context, table string, data interface{}, condition interface{}, args ...interface{}) (sql.Result, error) {
	return c.Model(table).Ctx(ctx).Data(data).Where(condition, args...).Update()
}

// DoUpdate 执行针对该表的 "UPDATE ... " 语句。
// 该函数通常用于自定义接口定义，您无需手动调用它。
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
		// 按照表格字段的顺序对数据键进行排序。
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
	// 如果没有传递链接，则使用主链接。
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

// Delete 执行针对该表的 "DELETE FROM ... " 语句。
//
// 参数 `condition` 可以是字符串、映射（map/gmap）、切片、结构体或指针类型等。
// 常常与参数 `args` 一起使用。例如：
// "uid=10000"，
// "uid", 10000
// "money>? AND name like ?", 99999, "vip_%"
// "status IN (?)", g.Slice{1,2,3}
// "age IN(?,?)", 18, 50
// User{ Id : 1, UserName : "john"}
// 中文注释：
// Delete 方法用于执行对该数据表执行 "DELETE FROM ... " SQL语句。
//
// 参数 `condition` 支持多种类型，如字符串、字典(map/gmap)、数组、结构体、结构体指针等。
// 通常会结合参数 `args` 使用，例如以下示例：
// "uid=10000"，（条件为 uid 等于 10000）
// "uid", 10000
// "money>? AND name like ?", 99999, "vip_%" （条件为 money 大于 99999 并且 name 字段匹配 "vip_%"）
// "status IN (?)", g.Slice{1,2,3} （条件为 status 字段在数组 [1,2,3] 中）
// "age IN(?,?)", 18, 50 （条件为 age 字段在范围 18 到 50 内）
// User{ Id : 1, UserName : "john"} （根据结构体定义的字段作为条件）
func (c *Core) Delete(ctx context.Context, table string, condition interface{}, args ...interface{}) (result sql.Result, err error) {
	return c.Model(table).Ctx(ctx).Where(condition, args...).Delete()
}

// DoDelete 执行针对表的 "DELETE FROM ..." 语句。
// 该函数通常用于自定义接口定义，无需手动调用。
func (c *Core) DoDelete(ctx context.Context, link Link, table string, condition string, args ...interface{}) (result sql.Result, err error) {
	if link == nil {
		if link, err = c.MasterLink(); err != nil {
			return nil, err
		}
	}
	table = c.QuotePrefixTableName(table)
	return c.db.DoExec(ctx, link, fmt.Sprintf("DELETE FROM %s%s", table, condition), args...)
}

// FilteredLink 获取并返回可用于日志记录或跟踪目的的已过滤`linkInfo`。
func (c *Core) FilteredLink() string {
	return fmt.Sprintf(
		`%s@%s(%s:%s)/%s`,
		c.config.User, c.config.Protocol, c.config.Host, c.config.Port, c.config.Name,
	)
}

// MarshalJSON 实现了 json.Marshal 接口的MarshalJSON方法。
// 它仅仅是返回指针地址。
//
// 注意，这个接口主要为了应对 Go 语言版本小于 v1.14 时存在的一个 json 无限循环bug而实现的。
func (c Core) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%+v`, c)), nil
}

// writeSqlToLogger 将 Sql 对象输出到日志器。
// 仅当配置项 "debug" 为 true 时，此功能才被启用。
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

// HasTable 判断给定的表名是否存在于数据库中。
func (c *Core) HasTable(name string) (bool, error) {
	tables, err := c.GetTablesWithCache()
	if err != nil {
		return false, err
	}
	for _, table := range tables {
		if table == name {
			return true, nil
		}
	}
	return false, nil
}

// GetTablesWithCache 使用缓存获取并返回当前数据库的表名。
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

// isSoftCreatedFieldName 检查并返回给定的字段名是否为自动填充的创建时间。
func (c *Core) isSoftCreatedFieldName(fieldName string) bool {
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

// FormatSqlBeforeExecuting 在执行SQL之前，对SQL字符串及其参数进行格式化处理。
// 在SQL过程中，内部函数handleArguments可能被调用两次，
// 但请不用担心，这是安全且高效的。
func (c *Core) FormatSqlBeforeExecuting(sql string, args []interface{}) (newSql string, newArgs []interface{}) {
// **不要**这样做，因为SQL中可能包含多行和注释。
// 删除sql的首尾空格
// 将sql中的换行符("\n")替换为空格
// 使用正则表达式将sql中连续出现2个或以上空格的情况替换为单个空格，并返回处理后的sql（_用于忽略错误信息）
	return handleArguments(sql, args)
}
