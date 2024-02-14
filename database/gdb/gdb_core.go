// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//

package db类

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
func (c *Core) X取Core对象() *Core {
	return c
}

// Ctx 是一个链式函数，它创建并返回一个新的 DB 对象，该对象是对当前 DB 对象的浅复制，并且其中包含给定的上下文。
// 注意，返回的这个 DB 对象只能使用一次，所以不要将其赋值给全局或包级别的变量以长期使用。
func (c *Core) X设置上下文并取副本(上下文 context.Context) DB {
	if 上下文 == nil {
		return c.db
	}
	// 它对当前db进行浅复制，并更改其上下文以进行下一个链式操作。
	var (
		err        error
		newCore    = &Core{}
		configNode = c.db.X取当前节点配置()
	)
	*newCore = *c
	// 它创建一个新的DB对象（非新连接），这个对象通常是对`Core`对象的一个包装。
	newCore.db, err = driverMap[configNode.X类型].New(newCore, configNode)
	if err != nil {
// 这里确实是一个严重的错误。
// 不要让它继续执行。
		panic(err)
	}
	newCore.ctx = 底层WithDB(上下文, newCore.db)
	newCore.ctx = c.底层_InjectInternalCtxData(newCore.ctx)
	return newCore.db
}

// GetCtx 返回当前数据库的上下文。
// 如果之前未设置上下文，则返回 `context.Background()`。
func (c *Core) X取上下文对象() context.Context {
	ctx := c.ctx
	if ctx == nil {
		ctx = context.TODO()
	}
	return c.底层_InjectInternalCtxData(ctx)
}

// GetCtxTimeout 根据指定的超时类型返回上下文和取消函数。
func (c *Core) X取超时上下文对象(上下文 context.Context, 超时类型 int) (context.Context, context.CancelFunc) {
	if 上下文 == nil {
		上下文 = c.db.X取上下文对象()
	} else {
		上下文 = context.WithValue(上下文, "WrappedByGetCtxTimeout", nil)
	}
	switch 超时类型 {
	case ctxTimeoutTypeExec:
		if c.db.X取当前节点配置().X执行超时时长 > 0 {
			return context.WithTimeout(上下文, c.db.X取当前节点配置().X执行超时时长)
		}
	case ctxTimeoutTypeQuery:
		if c.db.X取当前节点配置().X查询超时时长 > 0 {
			return context.WithTimeout(上下文, c.db.X取当前节点配置().X查询超时时长)
		}
	case ctxTimeoutTypePrepare:
		if c.db.X取当前节点配置().X预准备SQL超时时长 > 0 {
			return context.WithTimeout(上下文, c.db.X取当前节点配置().X预准备SQL超时时长)
		}
	default:
		panic(错误类.X创建错误码并格式化(错误码类.CodeInvalidParameter, "invalid context timeout type: %d", 超时类型))
	}
	return 上下文, func() {}
}

// Close 关闭数据库并阻止新的查询开始。
// Close 会等待在服务器上已经开始处理的所有查询完成。
//
// 很少会关闭一个 DB，因为 DB 处理程序旨在长期存在并被多个 goroutine 共享。
func (c *Core) X关闭数据库(上下文 context.Context) (错误 error) {
	if 错误 = c.cache.X关闭(上下文); 错误 != nil {
		return 错误
	}
	c.links.X遍历写锁定(func(m map[string]interface{}) {
		for k, v := range m {
			if db, ok := v.(*sql.DB); ok {
				错误 = db.Close()
				if 错误 != nil {
					错误 = 错误类.X多层错误码(错误码类.CodeDbOperationError, 错误, `db.Close failed`)
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

// Master在主从配置的情况下，创建并从主节点返回一个连接。如果未配置主从，则返回默认连接。
func (c *Core) X取主节点对象(数据库名称 ...string) (*sql.DB, error) {
	var (
		usedSchema   = 工具类.X取文本值或取默认值(c.schema, 数据库名称...)
		charL, charR = c.db.X底层取数据库安全字符()
	)
	return c.getSqlDb(true, 文本类.X过滤首尾符并含空白(usedSchema, charL+charR))
}

// Slave在主从配置的情况下，创建并返回从节点的连接。如果未配置主从，则返回默认连接。
func (c *Core) X取从节点对象(数据库名称 ...string) (*sql.DB, error) {
	var (
		usedSchema   = 工具类.X取文本值或取默认值(c.schema, 数据库名称...)
		charL, charR = c.db.X底层取数据库安全字符()
	)
	return c.getSqlDb(false, 文本类.X过滤首尾符并含空白(usedSchema, charL+charR))
}

// GetAll 从数据库查询并返回数据记录。
func (c *Core) GetAll别名(上下文 context.Context, sql string, 参数 ...interface{}) (Result, error) {
	return c.db.X底层查询(上下文, nil, sql, 参数...)
}

// DoSelect 从数据库查询并返回数据记录。
func (c *Core) X底层查询(上下文 context.Context, 链接 Link, sql string, 参数 ...interface{}) (结果 Result, 错误 error) {
	return c.db.X底层原生SQL查询(上下文, 链接, sql, 参数...)
}

// GetOne 从数据库查询并返回一条记录。
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

// GetArray 从数据库查询并返回数据值作为切片。
// 注意，如果结果中有多个列，则它会随机返回其中一列的值。
func (c *Core) X原生SQL查询数组(上下文 context.Context, sql string, 参数 ...interface{}) ([]Value, error) {
	all, err := c.db.X底层查询(上下文, nil, sql, 参数...)
	if err != nil {
		return nil, err
	}
	return all.X取字段数组(), nil
}

// doGetStruct 从数据库查询一条记录并将其转换为给定的结构体。
// 参数 `pointer` 应该是指向结构体的指针。
func (c *Core) doGetStruct(ctx context.Context, pointer interface{}, sql string, args ...interface{}) error {
	one, err := c.db.X原生SQL查询单条记录(ctx, sql, args...)
	if err != nil {
		return err
	}
	return one.X取结构体指针(pointer)
}

// doGetStructs 从数据库查询记录并将其转换为给定的结构体。
// 参数 `pointer` 应为结构体切片类型：[]struct 或 []*struct。
func (c *Core) doGetStructs(ctx context.Context, pointer interface{}, sql string, args ...interface{}) error {
	all, err := c.db.GetAll别名(ctx, sql, args...)
	if err != nil {
		return err
	}
	return all.X取数组结构体指针(pointer)
}

// GetScan 从数据库查询一个或多个记录，并将它们转换为给定的结构体或结构体数组。
//
// 如果参数`pointer`是结构体指针类型，它会内部调用 GetStruct 进行转换。如果参数 `pointer` 是切片类型，则会内部调用 GetStructs 进行转换。
func (c *Core) X原生SQL查询到结构体指针(上下文 context.Context, 结构体指针 interface{}, sql string, 参数 ...interface{}) error {
	reflectInfo := reflection.OriginTypeAndKind(结构体指针)
	if reflectInfo.InputKind != reflect.Ptr {
		return 错误类.X创建错误码并格式化(
			错误码类.CodeInvalidParameter,
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
	return 错误类.X创建错误码并格式化(
		错误码类.CodeInvalidParameter,
		`in valid parameter type "%v", of which element type should be type of struct/slice`,
		reflectInfo.InputType,
	)
}

// GetValue 从数据库查询并返回字段值。
// SQL语句应当只查询数据库中的一个字段，否则它将仅返回结果中的一个字段。
func (c *Core) X原生SQL查询字段值(上下文 context.Context, sql string, 参数 ...interface{}) (Value, error) {
	one, err := c.db.X原生SQL查询单条记录(上下文, sql, 参数...)
	if err != nil {
		return 泛型类.X创建(nil), err
	}
	for _, v := range one {
		return v, nil
	}
	return 泛型类.X创建(nil), nil
}

// GetCount 从数据库查询并返回计数。
func (c *Core) X原生SQL查询字段计数(上下文 context.Context, sql string, 参数 ...interface{}) (int, error) {
// 如果查询字段中不包含函数"COUNT"，
// 则替换sql字符串，并在字段中添加"COUNT"函数。
// 这段代码的注释是说，当SQL查询语句中的字段部分未使用“COUNT”函数时，会对原始的sql字符串进行替换处理，将“COUNT”函数添加到字段表达式中。
	if !正则类.X是否匹配文本(`(?i)SELECT\s+COUNT\(.+\)\s+FROM`, sql) {
		sql, _ = 正则类.X替换文本(`(?i)(SELECT)\s+(.+)\s+(FROM)`, `$1 COUNT($2) $3`, sql)
	}
	value, err := c.db.X原生SQL查询字段值(上下文, sql, 参数...)
	if err != nil {
		return 0, err
	}
	return value.X取整数(), nil
}

// Union 执行 "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." 语句。
func (c *Core) X多表去重查询(Model对象 ...*Model) *Model {
	var ctx = c.db.X取上下文对象()
	return c.doUnion(ctx, unionTypeNormal, Model对象...)
}

// UnionAll 执行 "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ..." 语句。
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

// PingMaster 用于向主节点发送心跳以检查身份验证或保持连接存活。
func (c *Core) X向主节点发送心跳() error {
	var ctx = c.db.X取上下文对象()
	if master, err := c.db.X取主节点对象(); err != nil {
		return err
	} else {
		if err = master.PingContext(ctx); err != nil {
			err = 错误类.X多层错误码(错误码类.CodeDbOperationError, err, `master.Ping failed`)
		}
		return err
	}
}

// PingSlave 向从节点发送ping请求，用于检查身份验证或保持连接活跃。
func (c *Core) X向从节点发送心跳() error {
	var ctx = c.db.X取上下文对象()
	if slave, err := c.db.X取从节点对象(); err != nil {
		return err
	} else {
		if err = slave.PingContext(ctx); err != nil {
			err = 错误类.X多层错误码(错误码类.CodeDbOperationError, err, `slave.Ping failed`)
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
func (c *Core) X插入(上下文 context.Context, 表名称 string, 值 interface{}, 批量操作行数 ...int) (sql.Result, error) {
	if len(批量操作行数) > 0 {
		return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X设置数据(值).X设置批量操作行数(批量操作行数[0]).X插入()
	}
	return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X设置数据(值).X插入()
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
func (c *Core) X插入并跳过已存在(上下文 context.Context, 表名称 string, 值 interface{}, 批量操作行数 ...int) (sql.Result, error) {
	if len(批量操作行数) > 0 {
		return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X设置数据(值).X设置批量操作行数(批量操作行数[0]).X插入并跳过已存在()
	}
	return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X设置数据(值).X插入并跳过已存在()
}

// InsertAndGetId 执行插入操作，并返回自动生成的最后一个插入ID。
func (c *Core) X插入并取ID(上下文 context.Context, 表名称 string, 值 interface{}, 批量操作行数 ...int) (int64, error) {
	if len(批量操作行数) > 0 {
		return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X设置数据(值).X设置批量操作行数(批量操作行数[0]).X插入并取ID()
	}
	return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X设置数据(值).X插入并取ID()
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
func (c *Core) X插入并替换已存在(上下文 context.Context, 表名称 string, 值 interface{}, 批量操作行数 ...int) (sql.Result, error) {
	if len(批量操作行数) > 0 {
		return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X设置数据(值).X设置批量操作行数(批量操作行数[0]).X插入并替换已存在()
	}
	return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X设置数据(值).X插入并替换已存在()
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
func (c *Core) X插入并更新已存在(上下文 context.Context, 表名称 string, 值 interface{}, 批量操作行数 ...int) (sql.Result, error) {
	if len(批量操作行数) > 0 {
		return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X设置数据(值).X设置批量操作行数(批量操作行数[0]).X插入并更新已存在()
	}
	return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X设置数据(值).X插入并更新已存在()
}

func (c *Core) fieldsToSequence(ctx context.Context, table string, fields []string) ([]string, error) {
	var (
		fieldSet               = 集合类.X创建文本并按值(fields)
		fieldsResultInSequence = make([]string, 0)
		tableFields, err       = c.db.X取表字段信息Map(ctx, table)
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
		if fieldSet.X是否存在(fieldName) {
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
func (c *Core) X底层插入(上下文 context.Context, 链接 Link, 表名称 string, list Map数组, option DoInsertOption) (result sql.Result, err error) {
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
	var keyListMap = map类.X创建链表mp()
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
		tmpKeysInSequenceStr = 文本类.X连接(keys, ",")
		if !keyListMap.X是否存在(tmpKeysInSequenceStr) {
			keyListMap.X设置值(tmpKeysInSequenceStr, make(Map数组, 0))
		}
		tmpKeysInSequenceList := keyListMap.X取值(tmpKeysInSequenceStr).(Map数组)
		tmpKeysInSequenceList = append(tmpKeysInSequenceList, item)
		keyListMap.X设置值(tmpKeysInSequenceStr, tmpKeysInSequenceList)
	}
	if keyListMap.X取数量() > 1 {
		var (
			tmpResult    sql.Result
			sqlResult    Sql执行结果
			rowsAffected int64
		)
		keyListMap.X遍历(func(key, value interface{}) bool {
			tmpResult, err = c.X底层插入(上下文, 链接, 表名称, value.(Map数组), option)
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

	// 准备批量结果指针。
	var (
		charL, charR = c.db.X底层取数据库安全字符()
		batchResult  = new(Sql执行结果)
		keysStr      = charL + strings.Join(keys, charR+","+charL) + charR
		operation    = 底层GetInsertOperationByOption(option.InsertOption)
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
				values = append(values, 转换类.String(s))
			} else {
				values = append(values, "?")
				params = append(params, list[i][k])
			}
		}
		valueHolder = append(valueHolder, "("+文本类.X连接(values, ",")+")")
		// 批量校验包：满足批量数量，或者已是最后一个元素。
		if len(valueHolder) == option.BatchCount || (i == listLength-1 && len(valueHolder) > 0) {
			var (
				stdSqlResult sql.Result
				affectedRows int64
			)
			stdSqlResult, err = c.db.X底层原生SQL执行(上下文, 链接, fmt.Sprintf(
				"%s INTO %s(%s) VALUES%s %s",
				operation, c.X底层添加前缀字符和引用字符(表名称), keysStr,
				文本类.X连接(valueHolder, ","),
				onDuplicateStr,
			), params...)
			if err != nil {
				return stdSqlResult, err
			}
			if affectedRows, err = stdSqlResult.RowsAffected(); err != nil {
				err = 错误类.X多层错误码(错误码类.CodeDbOperationError, err, `sql.Result.RowsAffected failed`)
				return stdSqlResult, err
			} else {
				batchResult.X原生sql行记录 = stdSqlResult
				batchResult.X影响行数 += affectedRows
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
					c.X底层QuoteWord(k),
					v,
				)
			default:
				onDuplicateStr += fmt.Sprintf(
					"%s=VALUES(%s)",
					c.X底层QuoteWord(k),
					c.X底层QuoteWord(转换类.String(v)),
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
				c.X底层QuoteWord(column),
				c.X底层QuoteWord(column),
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
func (c *Core) X更新(上下文 context.Context, 表名称 string, 数据 interface{}, 条件 interface{}, 参数 ...interface{}) (sql.Result, error) {
	return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X设置数据(数据).X条件(条件, 参数...).X更新()
}

// DoUpdate 执行针对该表的 "UPDATE ... " 语句。
// 该函数通常用于自定义接口定义，您无需手动调用它。
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
			counterHandler = func(column string, counter X增减) {
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
		// 按照表格字段的顺序对数据键进行排序。
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
			case *X增减:
				counterHandler(k, *value)

			case X增减:
				counterHandler(k, value)

			default:
				if s, ok := v.(Raw); ok {
					fields = append(fields, c.X底层QuoteWord(k)+"="+转换类.String(s))
				} else {
					fields = append(fields, c.X底层QuoteWord(k)+"=?")
					params = append(params, v)
				}
			}
		}
		updates = strings.Join(fields, ",")

	default:
		updates = 转换类.String(值)
	}
	if len(updates) == 0 {
		return nil, 错误类.X创建错误码(错误码类.CodeMissingParameter, "data cannot be empty")
	}
	if len(params) > 0 {
		参数 = append(params, 参数...)
	}
	// 如果没有传递链接，则使用主链接。
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
func (c *Core) X删除(上下文 context.Context, 表名称 string, 条件 interface{}, 参数 ...interface{}) (结果 sql.Result, 错误 error) {
	return c.X创建Model对象(表名称).X设置上下文并取副本(上下文).X条件(条件, 参数...).X删除()
}

// DoDelete 执行针对表的 "DELETE FROM ..." 语句。
// 该函数通常用于自定义接口定义，无需手动调用。
func (c *Core) X底层删除(上下文 context.Context, 链接 Link, 表名称 string, 条件 string, 参数 ...interface{}) (结果 sql.Result, 错误 error) {
	if 链接 == nil {
		if 链接, 错误 = c.X底层MasterLink(); 错误 != nil {
			return nil, 错误
		}
	}
	表名称 = c.X底层添加前缀字符和引用字符(表名称)
	return c.db.X底层原生SQL执行(上下文, 链接, fmt.Sprintf("DELETE FROM %s%s", 表名称, 条件), 参数...)
}

// FilteredLink 获取并返回可用于日志记录或跟踪目的的已过滤`linkInfo`。
func (c *Core) X取数据库链接信息() string {
	return fmt.Sprintf(
		`%s@%s(%s:%s)/%s`,
		c.config.X账号, c.config.X协议, c.config.X地址, c.config.X端口, c.config.X名称,
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
		c.logger.X输出DEBU(ctx, s)
	}
}

// HasTable 判断给定的表名是否存在于数据库中。
func (c *Core) X是否存在表名(表名称 string) (bool, error) {
	tables, err := c.X取表名称缓存()
	if err != nil {
		return false, err
	}
	for _, table := range tables {
		if table == 表名称 {
			return true, nil
		}
	}
	return false, nil
}

// GetTablesWithCache 使用缓存获取并返回当前数据库的表名。
func (c *Core) X取表名称缓存() ([]string, error) {
	var (
		ctx      = c.db.X取上下文对象()
		cacheKey = fmt.Sprintf(`Tables: %s`, c.db.X取配置组名称())
	)
	result, err := c.X取缓存对象().X取值或设置值_并发安全函数(
		ctx, cacheKey, func(ctx context.Context) (interface{}, error) {
			tableList, err := c.db.X取表名称数组(ctx)
			if err != nil {
				return false, err
			}
			return tableList, nil
		}, 0,
	)
	if err != nil {
		return nil, err
	}
	return result.X取文本数组(), nil
}

// isSoftCreatedFieldName 检查并返回给定的字段名是否为自动填充的创建时间。
func (c *Core) isSoftCreatedFieldName(fieldName string) bool {
	if fieldName == "" {
		return false
	}
	if config := c.db.X取当前节点配置(); config.X创建时间字段名 != "" {
		if utils.EqualFoldWithoutChars(fieldName, config.X创建时间字段名) {
			return true
		}
		return 文本类.X数组是否存在(append([]string{config.X创建时间字段名}, createdFieldNames...), fieldName)
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
func (c *Core) X格式化Sql(sql string, 参数数组 []interface{}) (新sql string, 新参数数组 []interface{}) {
// **不要**这样做，因为SQL中可能包含多行和注释。
// 删除sql的首尾空格
// 将sql中的换行符("\n")替换为空格
// 使用正则表达式将sql中连续出现2个或以上空格的情况替换为单个空格，并返回处理后的sql（_用于忽略错误信息）
	return handleArguments(sql, 参数数组)
}
