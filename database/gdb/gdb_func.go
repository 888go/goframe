// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb
import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/internal/empty"
	"github.com/888go/goframe/internal/reflection"
	"github.com/888go/goframe/internal/utils"
	"github.com/888go/goframe/os/gstructs"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gmeta"
	"github.com/888go/goframe/util/gutil"
	)
// iString 是 String 类型断言的 API。
type iString interface {
	String() string
}

// iIterator 是 Iterator 的类型断言 API。
type iIterator interface {
	Iterator(f func(key, value interface{}) bool)
}

// iInterfaces 是 Interfaces 的类型断言 API。
type iInterfaces interface {
	Interfaces() []interface{}
}

// iNil 是用于类型断言的 IsNil 方法。
type iNil interface {
	IsNil() bool
}

// iTableName 是用于为结构体获取表名的接口。
type iTableName interface {
	TableName() string
}

const (
	OrmTagForStruct    = "orm"
	OrmTagForTable     = "table"
	OrmTagForWith      = "with"
	OrmTagForWithWhere = "where"
	OrmTagForWithOrder = "order"
	OrmTagForDo        = "do"
)

var (
	// quoteWordReg 是用于单词检查的正则表达式对象。
	quoteWordReg = regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`)

	// structTagPriority 结构体标签，用于在ORM字段映射时进行结构体转换的标记。
	structTagPriority = append([]string{OrmTagForStruct}, gconv.StructTagPriority...)
)

// WithDB 将给定的 db 对象注入到上下文中并返回一个新的上下文。
func WithDB(ctx context.Context, db DB) context.Context {
	if db == nil {
		return ctx
	}
	dbCtx := db.GetCtx()
	if ctxDb := DBFromCtx(dbCtx); ctxDb != nil {
		return dbCtx
	}
	ctx = context.WithValue(ctx, ctxKeyForDB, db)
	return ctx
}

// DBFromCtx 从context中获取并返回DB对象。
func DBFromCtx(ctx context.Context) DB {
	if ctx == nil {
		return nil
	}
	v := ctx.Value(ctxKeyForDB)
	if v != nil {
		return v.(DB)
	}
	return nil
}

// ToSQL 格式化并返回给定闭包函数中的最后一个 SQL 语句，
// 但**并不会真正执行它**。
// 注意，所有后续的 SQL 语句都应使用通过 `f` 函数传递的上下文对象。
func ToSQL(ctx context.Context, f func(ctx context.Context) error) (sql string, err error) {
	var manager = &CatchSQLManager{
		SQLArray: garray.NewStrArray(),
		DoCommit: false,
	}
	ctx = context.WithValue(ctx, ctxKeyCatchSQL, manager)
	err = f(ctx)
	sql, _ = manager.SQLArray.PopRight()
	return
}

// CatchSQL 在给定闭包函数中捕获并返回所有已执行的SQL语句。
// 注意，所有后续SQL语句都应使用通过`f`函数传递的上下文对象。
func CatchSQL(ctx context.Context, f func(ctx context.Context) error) (sqlArray []string, err error) {
	var manager = &CatchSQLManager{
		SQLArray: garray.NewStrArray(),
		DoCommit: true,
	}
	ctx = context.WithValue(ctx, ctxKeyCatchSQL, manager)
	err = f(ctx)
	return manager.SQLArray.Slice(), err
}

// isDoStruct 检查并返回给定类型是否为DO结构体。
func isDoStruct(object interface{}) bool {
// 它通过检查结构体名称（如 "XxxForDao"），以兼容旧版本。
// TODO：未来删除这些兼容代码。
	reflectType := reflect.TypeOf(object)
	if gstr.HasSuffix(reflectType.String(), modelForDaoSuffix) {
		return true
	}
	// 它通过结构体元数据在指定版本中检查DO（Data Object）结构体。
	if ormTag := gmeta.Get(object, OrmTagForStruct); !ormTag.IsEmpty() {
		match, _ := gregex.MatchString(
			fmt.Sprintf(`%s\s*:\s*([^,]+)`, OrmTagForDo),
			ormTag.String(),
		)
		if len(match) > 1 {
			return gconv.Bool(match[1])
		}
	}
	return false
}

// getTableNameFromOrmTag 从结构体对象中获取并返回表名。
func getTableNameFromOrmTag(object interface{}) string {
	var tableName string
	// 使用接口值。
	if r, ok := object.(iTableName); ok {
		tableName = r.TableName()
	}
	// 用户元数据标签 "orm"。
	if tableName == "" {
		if ormTag := gmeta.Get(object, OrmTagForStruct); !ormTag.IsEmpty() {
			match, _ := gregex.MatchString(
				fmt.Sprintf(`%s\s*:\s*([^,]+)`, OrmTagForTable),
				ormTag.String(),
			)
			if len(match) > 1 {
				tableName = match[1]
			}
		}
	}
	// 使用蛇形命名法为结构体命名。
	if tableName == "" {
		if t, err := gstructs.StructType(object); err != nil {
			panic(err)
		} else {
			tableName = gstr.CaseSnakeFirstUpper(
				gstr.StrEx(t.String(), "."),
			)
		}
	}
	return tableName
}

// ListItemValues 函数用于获取并返回所有以 `key` 为键的项结构体或映射元素的值。
// 注意，参数 `list` 应为包含映射或结构体元素的切片类型，否则将返回一个空切片。
//
// 参数 `list` 支持以下类型：
// []map[string]interface{}
// []map[string]子映射
// []struct
// []struct:子结构体
// 注意，只有在提供可选参数 `subKey` 的情况下，子映射/子结构体才有意义。
// 参见 gutil.ListItemValues。
func ListItemValues(list interface{}, key interface{}, subKey ...interface{}) (values []interface{}) {
	return gutil.ListItemValues(list, key, subKey...)
}

// ListItemValuesUnique 函数用于获取并返回所有具有键 `key` 的结构体或映射中的唯一元素。
// 注意，参数 `list` 应为包含映射或结构体元素的切片类型，否则将返回一个空切片。
// 参见 gutil.ListItemValuesUnique。
func ListItemValuesUnique(list interface{}, key string, subKey ...interface{}) []interface{} {
	return gutil.ListItemValuesUnique(list, key, subKey...)
}

// GetInsertOperationByOption 根据给定的参数 `option` 返回合适的插入选项。
func GetInsertOperationByOption(option InsertOption) string {
	var operator string
	switch option {
	case InsertOptionReplace:
		operator = InsertOperationReplace
	case InsertOptionIgnore:
		operator = InsertOperationIgnore
	default:
		operator = InsertOperationInsert
	}
	return operator
}

func anyValueToMapBeforeToRecord(value interface{}) map[string]interface{} {
	return gconv.Map(value, gconv.MapOption{
		Tags:      structTagPriority,
		OmitEmpty: true, // 为了与 v2.6.0 及其之前的旧版本兼容。
	})
}

// DaToMapDeep 已废弃，请改用 MapOrStructToMapDeep。
func DaToMapDeep(value interface{}) map[string]interface{} {
	return MapOrStructToMapDeep(value, true)
}

// MapOrStructToMapDeep 递归地将`value`转换为map类型（如果属性结构体是嵌入的）。
// 参数`value`应为*map、map、*struct或struct类型。
// 它支持对结构体的嵌入式结构体定义。
func MapOrStructToMapDeep(value interface{}, omitempty bool) map[string]interface{} {
	m := gconv.Map(value, gconv.MapOption{
		Tags:      structTagPriority,
		OmitEmpty: omitempty,
	})
	for k, v := range m {
		switch v.(type) {
		case time.Time, *time.Time, gtime.Time, *gtime.Time, gjson.Json, *gjson.Json:
			m[k] = v
		}
	}
	return m
}

// doQuoteTableName 为表名添加前缀字符串和引用字符。它处理诸如以下格式的表名：
// "user", "user u", "user,user_detail", "user u, user_detail ut", "user as u, user_detail as ut",
// "user.user u", "`user`.`user` u"。
// 注意，此函数会自动检查表名是否已包含前缀，如果已包含，则不对表名做任何操作；
// 否则，将前缀添加到表名中，并返回带有前缀的新表名。
func doQuoteTableName(table, prefix, charLeft, charRight string) string {
	var (
		index  int
		chars  = charLeft + charRight
		array1 = gstr.SplitAndTrim(table, ",")
	)
	for k1, v1 := range array1 {
		array2 := gstr.SplitAndTrim(v1, " ")
		// 剔除安全字符。
		array2[0] = gstr.Trim(array2[0], chars)
		// 检查它是否包含数据库名。
		array3 := gstr.Split(gstr.Trim(array2[0]), ".")
		for k, v := range array3 {
			array3[k] = gstr.Trim(v, chars)
		}
		index = len(array3) - 1
		// 如果表名已经包含前缀，则跳过添加前缀的操作。
		if len(array3[index]) <= len(prefix) || array3[index][:len(prefix)] != prefix {
			array3[index] = prefix + array3[index]
		}
		array2[0] = gstr.Join(array3, ".")
		// 添加安全字符。
		array2[0] = doQuoteString(array2[0], charLeft, charRight)
		array1[k1] = gstr.Join(array2, " ")
	}
	return gstr.Join(array1, ",")
}

// doQuoteWord 检查给定的字符串 `s` 是否为单词，如果是，则使用 `charLeft` 和 `charRight` 对其进行引号括起，
// 并返回引号包含的字符串；否则原样返回 `s` 不做任何改变。
func doQuoteWord(s, charLeft, charRight string) string {
	if quoteWordReg.MatchString(s) && !gstr.ContainsAny(s, charLeft+charRight) {
		return charLeft + s + charRight
	}
	return s
}

// doQuoteString 用引号字符对字符串进行引用。
// 例如，如果引号字符是 '`'：
// "null"                             => "NULL"
// "user"                             => "`user`"
// "user u"                           => "`user` u"
// "user,user_detail"                 => "`user`,`user_detail`"
// "user u, user_detail ut"           => "`user` u,`user_detail` ut"
// "user.user u, user.user_detail ut" => "`user`.`user` u,`user`.`user_detail` ut"
// "u.id, u.name, u.age"              => "`u`.`id`,`u`.`name`,`u`.`age`"
// "u.id asc"                         => "`u`.`id` asc"
// 此函数用于将输入的字符串中的字段名用指定的引号（本例中为反引号 `）包裹起来，以符合SQL语句中字段名引用的规范。对于包含点号（.）的情况，会分别对表名和字段名进行引号包裹，同时保留原有的空格和逗号等分隔符及排序关键字（如“asc”）。
func doQuoteString(s, charLeft, charRight string) string {
	array1 := gstr.SplitAndTrim(s, ",")
	for k1, v1 := range array1 {
		array2 := gstr.SplitAndTrim(v1, " ")
		array3 := gstr.Split(gstr.Trim(array2[0]), ".")
		if len(array3) == 1 {
			if strings.EqualFold(array3[0], "NULL") {
				array3[0] = doQuoteWord(array3[0], "", "")
			} else {
				array3[0] = doQuoteWord(array3[0], charLeft, charRight)
			}
		} else if len(array3) >= 2 {
			array3[0] = doQuoteWord(array3[0], charLeft, charRight)
// 注意：
// mysql: u.uid
// mssql 双点表示法：Database..Table
// （注释翻译如下）
// 注释：
// 在 MySQL 中，使用 `u.uid` 表示用户ID
// 在 MSSQL 中，双点（..）用于表示特定的表引用方式，即“Database..Table”，表示从指定数据库引用某个表
			array3[len(array3)-1] = doQuoteWord(array3[len(array3)-1], charLeft, charRight)
		}
		array2[0] = gstr.Join(array3, ".")
		array1[k1] = gstr.Join(array2, " ")
	}
	return gstr.Join(array1, ",")
}

func getFieldsFromStructOrMap(structOrMap interface{}) (fields []string) {
	fields = []string{}
	if utils.IsStruct(structOrMap) {
		structFields, _ := gstructs.Fields(gstructs.FieldsInput{
			Pointer:         structOrMap,
			RecursiveOption: gstructs.RecursiveOptionEmbeddedNoTag,
		})
		var ormTagValue string
		for _, structField := range structFields {
			ormTagValue = structField.Tag(OrmTagForStruct)
			ormTagValue = gstr.Split(gstr.Trim(ormTagValue), ",")[0]
			if ormTagValue != "" && gregex.IsMatchString(regularFieldNameRegPattern, ormTagValue) {
				fields = append(fields, ormTagValue)
			} else {
				fields = append(fields, structField.Name())
			}
		}
	} else {
		fields = gutil.Keys(structOrMap)
	}
	return
}

// GetPrimaryKeyCondition 通过主键字段名称返回一个新的条件表达式。
// 可选参数 `where` 形如：
// 123                             => 主键=123
// []int{1, 2, 3}                  => 主键 IN(1,2,3)
// "john"                          => 主键='john'
// []string{"john", "smith"}       => 主键 IN('john','smith')
// g.Map{"id": g.Slice{1,2,3}}     => id IN(1,2,3)
// g.Map{"id": 1, "name": "john"}  => id=1 AND name='john'
// 等等
//
// 注意，如果 `primary` 为空或者 `where` 参数长度大于1，则直接返回给定的 `where` 参数。
func GetPrimaryKeyCondition(primary string, where ...interface{}) (newWhereCondition []interface{}) {
	if len(where) == 0 {
		return nil
	}
	if primary == "" {
		return where
	}
	if len(where) == 1 {
		var (
			rv   = reflect.ValueOf(where[0])
			kind = rv.Kind()
		)
		if kind == reflect.Ptr {
			rv = rv.Elem()
			kind = rv.Kind()
		}
		switch kind {
		case reflect.Map, reflect.Struct:
			// 忽略参数`primary`。
			break

		default:
			return []interface{}{map[string]interface{}{
				primary: where[0],
			}}
		}
	}
	return where
}

type formatWhereHolderInput struct {
	WhereHolder
	OmitNil   bool
	OmitEmpty bool
	Schema    string
	Table     string // Table 用于内部字段的映射和筛选。
}

func isKeyValueCanBeOmitEmpty(omitEmpty bool, whereType string, key, value interface{}) bool {
	if !omitEmpty {
		return false
	}
// 示例：
// Where("id", []int{}).All()             -> SELECT xxx FROM xxx WHERE 0=1
// 当条件字段为"id"且传入一个空的整数切片时，生成SQL语句为：从xxx表中选择所有列，但WHERE子句恒为假（0=1）
// Where("name", "").All()                -> SELECT xxx FROM xxx WHERE `name`=''
// 当条件字段为"name"且传入空字符串时，生成SQL语句为：从xxx表中选择所有列，其中"name"字段为空
// OmitEmpty().Where("id", []int{}).All() -> SELECT xxx FROM xxx
// 使用OmitEmpty()方法后，在条件字段为"id"且传入一个空的整数切片时，忽略无效的条件，生成SQL语句为：从xxx表中选择所有列，不带任何WHERE子句
// OmitEmpty().Where("name", "").All()    -> SELECT xxx FROM xxx
// 使用OmitEmpty()方法后，在条件字段为"name"且传入空字符串时，忽略无效的条件，生成SQL语句为：从xxx表中选择所有列，不带任何WHERE子句
// OmitEmpty().Where("1").All()           -> SELECT xxx FROM xxx WHERE 1
// 使用OmitEmpty()方法后，当条件为"1"时，生成SQL语句为：从xxx表中选择所有列，并在WHERE子句中使用恒为真的条件（1）
// 以上代码中的注释是描述了一个用于生成SQL查询语句的方法链式调用。根据不同的输入参数，会生成不同WHERE条件的SQL查询语句，而OmitEmpty()方法的作用是忽略无效或空的条件表达式。
	switch whereType {
	case whereHolderTypeNoArgs:
		return false

	case whereHolderTypeIn:
		return gutil.IsEmpty(value)

	default:
		if gstr.Count(gconv.String(key), "?") == 0 && gutil.IsEmpty(value) {
			return true
		}
	}
	return false
}

// formatWhereHolder 格式化 where 语句及其参数，用于 `Where` 和 `Having` 语句。
func formatWhereHolder(ctx context.Context, db DB, in formatWhereHolderInput) (newWhere string, newArgs []interface{}) {
	var (
		buffer      = bytes.NewBuffer(nil)
		reflectInfo = reflection.OriginValueAndKind(in.Where)
	)
	switch reflectInfo.OriginKind {
	case reflect.Array, reflect.Slice:
		newArgs = formatWhereInterfaces(db, gconv.Interfaces(in.Where), buffer, newArgs)

	case reflect.Map:
		for key, value := range MapOrStructToMapDeep(in.Where, true) {
			if in.OmitNil && empty.IsNil(value) {
				continue
			}
			if in.OmitEmpty && empty.IsEmpty(value) {
				continue
			}
			newArgs = formatWhereKeyValue(formatWhereKeyValueInput{
				Db:     db,
				Buffer: buffer,
				Args:   newArgs,
				Key:    key,
				Value:  value,
				Prefix: in.Prefix,
				Type:   in.Type,
			})
		}

	case reflect.Struct:
// 如果`where`参数是`DO`结构体，那么它会为这个条件添加`OmitNil`选项，
// 这将会过滤`where`中所有为nil的参数。
		if isDoStruct(in.Where) {
			in.OmitNil = true
		}
// 如果`where`结构体实现了`iIterator`接口，
// 那么它会使用其Iterate函数遍历自身的键值对。
// 例如，ListMap和TreeMap是有序映射，
// 它们实现了`iIterator`接口，对于where条件查询时更加友好。
		if iterator, ok := in.Where.(iIterator); ok {
			iterator.Iterator(func(key, value interface{}) bool {
				ketStr := gconv.String(key)
				if in.OmitNil && empty.IsNil(value) {
					return true
				}
				if in.OmitEmpty && empty.IsEmpty(value) {
					return true
				}
				newArgs = formatWhereKeyValue(formatWhereKeyValueInput{
					Db:        db,
					Buffer:    buffer,
					Args:      newArgs,
					Key:       ketStr,
					Value:     value,
					OmitEmpty: in.OmitEmpty,
					Prefix:    in.Prefix,
					Type:      in.Type,
				})
				return true
			})
			break
		}
		// 自动映射并过滤结构体属性。
		var (
			reflectType = reflectInfo.OriginValue.Type()
			structField reflect.StructField
			data        = MapOrStructToMapDeep(in.Where, true)
		)
		// 如果`Prefix`已给出，它会检查并检索表名。
		if in.Prefix != "" {
			hasTable, _ := db.GetCore().HasTable(in.Prefix)
			if hasTable {
				in.Table = in.Prefix
			} else {
				ormTagTableName := getTableNameFromOrmTag(in.Where)
				if ormTagTableName != "" {
					in.Table = ormTagTableName
				}
			}
		}
		// 如果提供了`Table`，则对字段进行映射和过滤。
		if in.Table != "" {
			data, _ = db.GetCore().mappingAndFilterData(ctx, in.Schema, in.Table, data, true)
		}
		// 在Where语句中按顺序放置结构体属性。
		var ormTagValue string
		for i := 0; i < reflectType.NumField(); i++ {
			structField = reflectType.Field(i)
			// 如果已指定，则使用来自`orm`标签的值作为字段名。
			ormTagValue = structField.Tag.Get(OrmTagForStruct)
			ormTagValue = gstr.Split(gstr.Trim(ormTagValue), ",")[0]
			if ormTagValue == "" {
				ormTagValue = structField.Name
			}
			foundKey, foundValue := gutil.MapPossibleItemByKey(data, ormTagValue)
			if foundKey != "" {
				if in.OmitNil && empty.IsNil(foundValue) {
					continue
				}
				if in.OmitEmpty && empty.IsEmpty(foundValue) {
					continue
				}
				newArgs = formatWhereKeyValue(formatWhereKeyValueInput{
					Db:        db,
					Buffer:    buffer,
					Args:      newArgs,
					Key:       foundKey,
					Value:     foundValue,
					OmitEmpty: in.OmitEmpty,
					Prefix:    in.Prefix,
					Type:      in.Type,
				})
			}
		}

	default:
		// Where filter.
		var omitEmptyCheckValue interface{}
		if len(in.Args) == 1 {
			omitEmptyCheckValue = in.Args[0]
		} else {
			omitEmptyCheckValue = in.Args
		}
		if isKeyValueCanBeOmitEmpty(in.OmitEmpty, in.Type, in.Where, omitEmptyCheckValue) {
			return
		}
		// Usually a string.
		whereStr := gstr.Trim(gconv.String(in.Where))
// `whereStr` 是否是由键值对构成的字段名？
// 例如：
// Where("id", 1) // 指定id为1的条件
// Where("id", g.Slice{1,2,3}) // 指定id在[1, 2, 3]范围内的条件
		if gregex.IsMatchString(regularFieldNameWithoutDotRegPattern, whereStr) && len(in.Args) == 1 {
			newArgs = formatWhereKeyValue(formatWhereKeyValueInput{
				Db:        db,
				Buffer:    buffer,
				Args:      newArgs,
				Key:       whereStr,
				Value:     in.Args[0],
				OmitEmpty: in.OmitEmpty,
				Prefix:    in.Prefix,
				Type:      in.Type,
			})
			in.Args = in.Args[:0]
			break
		}
		// 如果第一部分是列名，它会自动为该列添加前缀。
		if in.Prefix != "" {
			array := gstr.Split(whereStr, " ")
			if ok, _ := db.GetCore().HasField(ctx, in.Table, array[0]); ok {
				whereStr = in.Prefix + "." + whereStr
			}
		}
// 正常字符串及参数占位符处理
// 示例：
// Where("id in(?) and name=?", g.Slice{1,2,3}, "john")
// 表示：当id在(1,2,3)中且name为"john"时的条件语句
		i := 0
		for {
			if i >= len(in.Args) {
				break
			}
// ===============================================================
// 子查询，始终与字符串条件一起使用。
// ===============================================================
			if subModel, ok := in.Args[i].(*Model); ok {
				index := -1
				whereStr, _ = gregex.ReplaceStringFunc(`(\?)`, whereStr, func(s string) string {
					index++
					if i+len(newArgs) == index {
						sqlWithHolder, holderArgs := subModel.getHolderAndArgsAsSubModel(ctx)
						in.Args = gutil.SliceInsertAfter(in.Args, i, holderArgs...)
						// 自动添加括号
						return "(" + sqlWithHolder + ")"
					}
					return s
				})
				in.Args = gutil.SliceDelete(in.Args, i)
				continue
			}
			i++
		}
		buffer.WriteString(whereStr)
	}

	if buffer.Len() == 0 {
		return "", in.Args
	}
	if len(in.Args) > 0 {
		newArgs = append(newArgs, in.Args...)
	}
	newWhere = buffer.String()
	if len(newArgs) > 0 {
		if gstr.Pos(newWhere, "?") == -1 {
			if gregex.IsMatchString(lastOperatorRegPattern, newWhere) {
				// 示例：Where/And/Or("uid>=", 1)
// （译注：在Go语言中，这段代码可能是用于构建SQL查询条件的方法调用，表示查询条件为“uid大于等于1”）
// Where: 设置或添加查询条件，如“uid>=”
// And: 在已有的查询条件下追加一个与（AND）关系的条件，此处表示“并且uid大于等于1”
// Or: 在已有的查询条件下追加一个或（OR）关系的条件，但根据示例实际未使用到OR操作
// 整体来看，这段代码片段是展示如何通过链式调用构建复杂查询条件的一种方式。
				newWhere += "?"
			} else if gregex.IsMatchString(regularFieldNameRegPattern, newWhere) {
				newWhere = db.GetCore().QuoteString(newWhere)
				if len(newArgs) > 0 {
					if utils.IsArray(newArgs[0]) {
// 示例：
// Where("id", []int{1,2,3}) // 根据id为1、2、3进行查询
// Where("user.id", []int{1,2,3}) // 根据user表中的id为1、2、3进行查询
// 以上Go语言代码的注释翻译成中文如下：
// ```go
// 例如：
// Where("id", []int{1,2,3}) // 用于指定id字段分别在1、2、3时的条件查询
// Where("user.id", []int{1,2,3}) // 用于指定user表中id字段分别在1、2、3时的条件查询
						newWhere += " IN (?)"
					} else if empty.IsNil(newArgs[0]) {
// 示例：
// Where("id", nil) // 根据id查询
// Where("user.id", nil) // 根据user表中的id字段查询
						newWhere += " IS NULL"
						newArgs = nil
					} else {
// 示例：
// Where/And/Or("uid", 1) // 指定条件：uid 等于 1
// Where/And/Or("user.uid", 1) // 指定条件：user.uid 等于 1
// 这段 Go 代码的注释翻译成中文为：
// ```go
// 示例：
// Where/And/Or 函数用于设置查询条件，例如：
// Where/And/Or("uid", 1) // 设置条件为 uid 等于 1
// Where/And/Or("user.uid", 1) // 设置条件为 user 表中的 uid 字段等于 1
// 这里的 `Where`、`And` 和 `Or` 都是可能的方法名，表示SQL语句中的“WHERE”子句以及逻辑连接符“AND”、“OR”，用于构建查询条件。在实际的数据库操作中，这些方法通常用于链式调用，形成更复杂的查询表达式。
						newWhere += "=?"
					}
				}
			}
		}
	}
	return handleArguments(newWhere, newArgs)
}

// formatWhereInterfaces 将 `where` 格式化为 []interface{} 类型的切片。
func formatWhereInterfaces(db DB, where []interface{}, buffer *bytes.Buffer, newArgs []interface{}) []interface{} {
	if len(where) == 0 {
		return newArgs
	}
	if len(where)%2 != 0 {
		buffer.WriteString(gstr.Join(gconv.Strings(where), ""))
		return newArgs
	}
	var str string
	for i := 0; i < len(where); i += 2 {
		str = gconv.String(where[i])
		if buffer.Len() > 0 {
			buffer.WriteString(" AND " + db.GetCore().QuoteWord(str) + "=?")
		} else {
			buffer.WriteString(db.GetCore().QuoteWord(str) + "=?")
		}
		if s, ok := where[i+1].(Raw); ok {
			buffer.WriteString(gconv.String(s))
		} else {
			newArgs = append(newArgs, where[i+1])
		}
	}
	return newArgs
}

type formatWhereKeyValueInput struct {
	Db        DB            // Db 是当前操作所基于的底层 DB 对象。
	Buffer    *bytes.Buffer // Buffer 是当前操作中不包含 Args 的 SQL 语句字符串。
	Args      []interface{} // Args 是当前操作的完整参数。
	Key       string        // 字段名称，例如："id"、"name"等。
	Value     interface{}   // 字段值，可以是任意类型。
	Type      string        // Where 类型中的值。
	OmitEmpty bool          // 如果`value`为空，则忽略当前条件键。
	Prefix    string        // 字段前缀，例如："user"、"order"等。
}

// formatWhereKeyValue 处理参数映射中的每一组键值对。
func formatWhereKeyValue(in formatWhereKeyValueInput) (newArgs []interface{}) {
	var (
		quotedKey   = in.Db.GetCore().QuoteWord(in.Key)
		holderCount = gstr.Count(quotedKey, "?")
	)
	if isKeyValueCanBeOmitEmpty(in.OmitEmpty, in.Type, quotedKey, in.Value) {
		return in.Args
	}
	if in.Prefix != "" && !gstr.Contains(quotedKey, ".") {
		quotedKey = in.Prefix + "." + quotedKey
	}
	if in.Buffer.Len() > 0 {
		in.Buffer.WriteString(" AND ")
	}
// 如果值的类型是切片，并且在键字符串中只有一个 '?' 占位符，
// 那么它会根据其参数个数自动添加 '?' 占位符，并将其转换为 "IN" 语句。
	var (
		reflectValue = reflect.ValueOf(in.Value)
		reflectKind  = reflectValue.Kind()
	)
	switch reflectKind {
	// Slice argument.
	case reflect.Slice, reflect.Array:
		if holderCount == 0 {
			in.Buffer.WriteString(quotedKey + " IN(?)")
			in.Args = append(in.Args, in.Value)
		} else {
			if holderCount != reflectValue.Len() {
				in.Buffer.WriteString(quotedKey)
				in.Args = append(in.Args, in.Value)
			} else {
				in.Buffer.WriteString(quotedKey)
				in.Args = append(in.Args, gconv.Interfaces(in.Value)...)
			}
		}

	default:
		if in.Value == nil || empty.IsNil(reflectValue) {
			if gregex.IsMatchString(regularFieldNameRegPattern, in.Key) {
				// 键是一个单独的字段名称。
				in.Buffer.WriteString(quotedKey + " IS NULL")
			} else {
				// 密钥可能包含操作字符。
				in.Buffer.WriteString(quotedKey)
			}
		} else {
			// 它还支持 "LIKE" 语句，我们认为它是一个运算符。
			quotedKey = gstr.Trim(quotedKey)
			if gstr.Pos(quotedKey, "?") == -1 {
				like := " LIKE"
				if len(quotedKey) > len(like) && gstr.Equal(quotedKey[len(quotedKey)-len(like):], like) {
					// 示例：Where(g.Map{"name like": "john%"})
// （注：此代码片段使用了golang编写的数据库操作语句，其中"Where"表示SQL中的WHERE子句，用于设置查询条件。这里传入了一个g.Map类型的参数，它是一个键值对映射，其中"name like"是SQL的模糊查询条件，"john%"代表查询名字以"john"开头的所有记录。）
					in.Buffer.WriteString(quotedKey + " ?")
				} else if gregex.IsMatchString(lastOperatorRegPattern, quotedKey) {
					// 示例：Where(g.Map{"age > ": 16})
// （译注：此处代码为Go语言中使用g.Map进行条件筛选的示例，其中"g.Map"是一个自定义的映射类型，"age > "表示年龄大于，整体即表示筛选出年龄大于16的记录。）
					in.Buffer.WriteString(quotedKey + " ?")
				} else if gregex.IsMatchString(regularFieldNameRegPattern, in.Key) {
					// key 是一个普通的字段名称。
					in.Buffer.WriteString(quotedKey + "=?")
				} else {
// 这个键不是常规的字段名称。
// 例如：Where(g.Map{"age > 16": nil})
// 相关问题：https://github.com/gogf/gf/issues/765
					if empty.IsEmpty(in.Value) {
						in.Buffer.WriteString(quotedKey)
						break
					} else {
						in.Buffer.WriteString(quotedKey + "=?")
					}
				}
			} else {
				in.Buffer.WriteString(quotedKey)
			}
			if s, ok := in.Value.(Raw); ok {
				in.Buffer.WriteString(gconv.String(s))
			} else {
				in.Args = append(in.Args, in.Value)
			}
		}
	}
	return in.Args
}

// handleArguments 是一个重要的函数，它在将 SQL 及其所有参数提交给底层驱动之前，负责处理这些 SQL 和参数。
func handleArguments(sql string, args []interface{}) (newSql string, newArgs []interface{}) {
	newSql = sql
	// insertHolderCount 用于计算 '?' 占位符的插入位置。
	insertHolderCount := 0
	// 处理切片参数。
	if len(args) > 0 {
		for index, arg := range args {
			reflectInfo := reflection.OriginValueAndKind(arg)
			switch reflectInfo.OriginKind {
			case reflect.Slice, reflect.Array:
// 它不会分割[]byte类型的数据。
// 例如：table.Where("name = ?", []byte("john"))
// 翻译为：
// 此处的处理不会对[]byte类型的值进行分割操作。
// 举例说明：在调用table.Where方法时，可以传入一个如"name = ?"的条件字符串以及一个[]byte类型的值，如：[]byte("john")。
				if _, ok := arg.([]byte); ok {
					newArgs = append(newArgs, arg)
					continue
				}

				if reflectInfo.OriginValue.Len() == 0 {
// 当传入空切片作为参数时，它会将SQL转换为一个永假的SQL语句。
// 例如：
// Query("select * from xxx where id in(?)", g.Slice{}) 将转换为 -> select * from xxx where 0=1
// Where("id in(?)", g.Slice{}) 将转换为 -> WHERE 0=1
					if gstr.Contains(newSql, "?") {
						whereKeyWord := " WHERE "
						if p := gstr.PosI(newSql, whereKeyWord); p == -1 {
							return "0=1", []interface{}{}
						} else {
							return gstr.SubStr(newSql, 0, p+len(whereKeyWord)) + "0=1", []interface{}{}
						}
					}
				} else {
					for i := 0; i < reflectInfo.OriginValue.Len(); i++ {
						newArgs = append(newArgs, reflectInfo.OriginValue.Index(i).Interface())
					}
				}

// 如果'?'占位符的数量等于切片的长度，
// 则它不会执行参数分割逻辑。
// 例如：db.Query("SELECT ?+?", g.Slice{1, 2})
				if len(args) == 1 && gstr.Count(newSql, "?") == reflectInfo.OriginValue.Len() {
					break
				}
				// counter 用于计算 '?' 占位符的插入位置。
				var (
					counter  = 0
					replaced = false
				)
				newSql, _ = gregex.ReplaceStringFunc(`\?`, newSql, func(s string) string {
					if replaced {
						return s
					}
					counter++
					if counter == index+insertHolderCount+1 {
						replaced = true
						insertHolderCount += reflectInfo.OriginValue.Len() - 1
						return "?" + strings.Repeat(",?", reflectInfo.OriginValue.Len()-1)
					}
					return s
				})

			// 特殊结构体处理。
			case reflect.Struct:
				switch arg.(type) {
				// 底层驱动程序支持 time.Time 类型（注：time.Time 是 Go 语言中的时间类型）。
				case time.Time, *time.Time:
					newArgs = append(newArgs, arg)
					continue

				case gtime.Time:
					newArgs = append(newArgs, arg.(gtime.Time).Time)
					continue

				case *gtime.Time:
					newArgs = append(newArgs, arg.(*gtime.Time).Time)
					continue

				default:
// 如果结构体实现了 String 接口，它会默认将该结构体转换为字符串。
					if v, ok := arg.(iString); ok {
						newArgs = append(newArgs, v.String())
						continue
					}
				}
				newArgs = append(newArgs, arg)

			default:
				newArgs = append(newArgs, arg)
			}
		}
	}
	return
}

// FormatSqlWithArgs 将参数绑定到sql字符串，并返回一个完整的sql字符串，仅用于调试。
func FormatSqlWithArgs(sql string, args []interface{}) string {
	index := -1
	newQuery, _ := gregex.ReplaceStringFunc(
		`(\?|:v\d+|\$\d+|@p\d+)`,
		sql,
		func(s string) string {
			index++
			if len(args) > index {
				if args[index] == nil {
					return "null"
				}
				// 类型为Raw的参数不需要进行特殊处理
				if v, ok := args[index].(Raw); ok {
					return gconv.String(v)
				}
				reflectInfo := reflection.OriginValueAndKind(args[index])
				if reflectInfo.OriginKind == reflect.Ptr &&
					(reflectInfo.OriginValue.IsNil() || !reflectInfo.OriginValue.IsValid()) {
					return "null"
				}
				switch reflectInfo.OriginKind {
				case reflect.String, reflect.Map, reflect.Slice, reflect.Array:
					return `'` + gstr.QuoteMeta(gconv.String(args[index]), `'`) + `'`

				case reflect.Struct:
					if t, ok := args[index].(time.Time); ok {
						return `'` + t.Format(`2006-01-02 15:04:05`) + `'`
					}
					return `'` + gstr.QuoteMeta(gconv.String(args[index]), `'`) + `'`
				}
				return gconv.String(args[index])
			}
			return s
		})
	return newQuery
}
