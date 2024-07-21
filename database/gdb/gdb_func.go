// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gdb

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/internal/empty"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/internal/reflection"
	"github.com/gogf/gf/v2/internal/utils"
	"github.com/gogf/gf/v2/os/gstructs"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/gogf/gf/v2/util/gtag"
	"github.com/gogf/gf/v2/util/gutil"
)

// iString 是用于字符串的类型断言接口。 md5:a51d442c33ac2cf4
type iString interface {
	String() string
}

// iIterator是Iterator的类型断言API。 md5:9e146fc0e640273e
type iIterator interface {
	Iterator(f func(key, value interface{}) bool)
}

// iInterfaces是Interfaces类型的断言API。 md5:843230a1ccff49f5
type iInterfaces interface {
	Interfaces() []interface{}
}

// iNil 用于类型断言接口以检查是否为Nil。 md5:49ddfde26b501402
type iNil interface {
	IsNil() bool
}

// iTableName 是用于获取结构体对应表名的接口。 md5:f3583f3a54701536
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
	// quoteWordReg是用于单词检查的正则表达式对象。 md5:99c41eabb9d23388
	quoteWordReg = regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`)

	// structTagPriority 是用于结构体转换为 ORM 字段映射的标签优先级。 md5:6e5a8632b6c8e48f
	structTagPriority = append([]string{OrmTagForStruct}, gtag.StructTagPriority...)
)

// WithDB 将给定的db对象注入到context中并返回一个新的context。 md5:e414408e96157a02
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

// DBFromCtx 从上下文中获取并返回DB对象。 md5:90c01e951db89218
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

// ToSQL 将给定闭包函数中的最后一个 SQL 语句格式化并返回，但并不会真正执行。
// 注意，所有后续的 SQL 语句都应该使用由函数 `f` 传递的上下文对象。
// md5:3fe82285d68728a0
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

// CatchSQL捕获并返回在给定闭包函数中执行的所有SQL语句。
// 注意，所有以下SQL语句都应使用由`f`函数传递的context对象。
// md5:1088111f1248173d
func CatchSQL(ctx context.Context, f func(ctx context.Context) error) (sqlArray []string, err error) {
	var manager = &CatchSQLManager{
		SQLArray: garray.NewStrArray(),
		DoCommit: true,
	}
	ctx = context.WithValue(ctx, ctxKeyCatchSQL, manager)
	err = f(ctx)
	return manager.SQLArray.Slice(), err
}

// isDoStruct 检查并返回给定类型是否为 DO 结构体。 md5:c235f077b3b3fcc5
func isDoStruct(object interface{}) bool {
	// 它通过结构体名称如 "XxxForDao" 进行检查，以兼容旧版本。
	// TODO: 未来某个时候移除这个兼容性代码。
	// md5:c8dc9518a3014aca
	reflectType := reflect.TypeOf(object)
	if gstr.HasSuffix(reflectType.String(), modelForDaoSuffix) {
		return true
	}
	// 它通过结构体元数据检查version中的DO结构。 md5:1651dfa7d2770eb0
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

// getTableNameFromOrmTag 从结构体对象中检索并返回表名。 md5:9da9a9980775dc38
func getTableNameFromOrmTag(object interface{}) string {
	var tableName string
	// 使用接口值。 md5:d769d38046ef266c
	if r, ok := object.(iTableName); ok {
		tableName = r.TableName()
	}
	// User meta data 标签 "orm"。 md5:7269c54b8f9aa97f
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
	// 使用蛇形命名的结构体名称。 md5:02b76586ae24bd54
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

// ListItemValues 从所有具有键为 `key` 的映射或结构体元素中检索并返回。请注意，参数 `list` 应该是包含映射或结构体元素的切片，否则将返回一个空切片。
// 
// 参数 `list` 支持以下类型：
// []map[string]interface{}
// []map[string]子映射
// []struct
// []struct:子结构体
// 如果提供了可选参数 `subKey`，子映射/子结构体才有意义。请参阅 gutil.ListItemValues。
// md5:e67327bcbcd82096
func ListItemValues(list interface{}, key interface{}, subKey ...interface{}) (values []interface{}) {
	return gutil.ListItemValues(list, key, subKey...)
}

// ListItemValuesUnique 获取并返回所有结构体/映射中键为`key`的唯一元素。
// 注意，参数`list`应为切片类型，且包含的元素为映射或结构体，
// 否则将返回一个空切片。
// 参见gutil.ListItemValuesUnique。
// md5:aa00cb15fafa41ba
func ListItemValuesUnique(list interface{}, key string, subKey ...interface{}) []interface{} {
	return gutil.ListItemValuesUnique(list, key, subKey...)
}

// GetInsertOperationByOption 根据给定的 `option` 参数返回合适的插入操作选项。 md5:19b87dd1244d55ec
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
	convertedMap := gconv.Map(value, gconv.MapOption{
		Tags:      structTagPriority,
		OmitEmpty: true, // 为了与从v2.6.0版本兼容。 md5:c3311f248eacbf1e
	})
	if gutil.OriginValueAndKind(value).OriginKind != reflect.Struct {
		return convertedMap
	}
	// 这里将结构体/映射切片的所有属性转换为JSON字符串。 md5:0355d0233b27ae20
	for k, v := range convertedMap {
		originValueAndKind := gutil.OriginValueAndKind(v)
		switch originValueAndKind.OriginKind {
		// 检查map中的项目切片项。 md5:e8ef9cf68f78aa92
		case reflect.Array, reflect.Slice:
			mapItemValue := originValueAndKind.OriginValue
			if mapItemValue.Len() == 0 {
				break
			}
			// 检查切片元素类型为结构体/映射类型。 md5:8ae4ace59a5f1070
			switch mapItemValue.Index(0).Kind() {
			case reflect.Struct, reflect.Map:
				mapItemJsonBytes, err := json.Marshal(v)
				if err != nil {
					// Do not eat any error.
					intlog.Error(context.TODO(), err)
				}
				convertedMap[k] = mapItemJsonBytes
			}
		}
	}
	return convertedMap
}

// MapOrStructToMapDeep 递归地将 `value` 转换为映射类型（如果属性结构体被嵌入）。
// 参数 `value` 应该是 *map、map、*struct 或 struct 类型。
// 它支持结构体中的嵌入式结构体定义。
// md5:3d9a4c7ad65d9fe1
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

// doQuoteTableName 在表名前添加前缀字符串和引号。它处理的表名类型包括：
// "user", "user u", "user,user_detail", "user u, user_detail ut", "user as u, user_detail as ut",
// "user.user u", "``user``.``user`` u"。
// 
// 注意，它会自动检查表名前缀是否已添加。如果已添加，则不进行任何操作；否则，会在表名前添加前缀，并返回带有前缀的新表名。
// md5:fc5ea60c27043ac8
func doQuoteTableName(table, prefix, charLeft, charRight string) string {
	var (
		index  int
		chars  = charLeft + charRight
		array1 = gstr.SplitAndTrim(table, ",")
	)
	for k1, v1 := range array1 {
		array2 := gstr.SplitAndTrim(v1, " ")
		// 去除安全字符。 md5:13d69b898c5635c6
		array2[0] = gstr.Trim(array2[0], chars)
		// 检查是否具有数据库名称。 md5:8d7ee8b347dfbbac
		array3 := gstr.Split(gstr.Trim(array2[0]), ".")
		for k, v := range array3 {
			array3[k] = gstr.Trim(v, chars)
		}
		index = len(array3) - 1
		// 如果表名已经包含前缀，则跳过添加前缀。 md5:047ee36460a1c519
		if len(array3[index]) <= len(prefix) || array3[index][:len(prefix)] != prefix {
			array3[index] = prefix + array3[index]
		}
		array2[0] = gstr.Join(array3, ".")
		// 添加安全字符。 md5:b299af34cac147b9
		array2[0] = doQuoteString(array2[0], charLeft, charRight)
		array1[k1] = gstr.Join(array2, " ")
	}
	return gstr.Join(array1, ",")
}

// doQuoteWord 检查给定的字符串 `s` 是否为一个单词，如果是，则使用 `charLeft` 和 `charRight` 对其进行引用并返回被引用的字符串；否则原样返回 `s`，不做任何改变。
// md5:ac0c8a621b951784
func doQuoteWord(s, charLeft, charRight string) string {
	if quoteWordReg.MatchString(s) && !gstr.ContainsAny(s, charLeft+charRight) {
		return charLeft + s + charRight
	}
	return s
}

// doQuoteString 使用引号字符对字符串进行引用。
// 例如，如果引用字符是 '`'：
// "null"                             => "NULL"
// "user"                             => "`user`"
// "user u"                           => "`user` u"
// "user,user_detail"                 => "`user`,`user_detail`"
// "user u, user_detail ut"           => "`user` u,`user_detail` ut"
// "user.user u, user.user_detail ut" => "`user`.`user` u,`user`.`user_detail` ut"
// "u.id, u.name, u.age"              => "`u`.`id`,`u`.`name`,`u`.`age`"
// "u.id asc"                         => "`u`.`id` asc"
// md5:556d2b4db186afe8
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
			// 注释：
			// mysql：u.uid
			// mssql：使用双点表示法 Database..Table
			// md5:66df82a8563f168b
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

// GetPrimaryKeyCondition 通过主键字段名返回一个新的WHERE条件。可选参数`where`的格式如下：
// 123                            => primary=123
// []int{1, 2, 3}                  => primary IN(1,2,3)
// "john"                          => primary='john'
// []string{"john", "smith"}       => primary IN('john','smith')
// g.Map{"id": g.Slice{1,2,3}}     => id IN(1,2,3)
// g.Map{"id": 1, "name": "john"}  => id=1 AND name='john'
// 等等。
//
// 注意，如果主键为空或者`where`参数长度大于1，它会直接返回给定的`where`参数。
// md5:748dfa9c0f0d93b5
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
			// 忽略参数 `primary`。 md5:be747ae45f6887e1
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
	Table     string // Table 用于内部字段映射和过滤。 md5:6f0adcc9f806782a
}

func isKeyValueCanBeOmitEmpty(omitEmpty bool, whereType string, key, value interface{}) bool {
	if !omitEmpty {
		return false
	}
	// 例如：
	// Where("id", []int{}) .All()             -> 选择xxx FROM xxx WHERE 0=1
	// Where("name", "") .All()                -> 选择xxx FROM xxx WHERE `name`= ''
	// OmitEmpty() .Where("id", []int{}) .All() -> 选择xxx FROM xxx
	// OmitEmpty() .Where("name", "") .All()    -> 选择xxx FROM xxx
	// OmitEmpty() .Where("1") .All()           -> 选择xxx FROM xxx WHERE 1
	// md5:13a1baa59d83a9fe
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

// formatWhereHolder 格式化 WHERE 和 HAVING 语句及其参数。 md5:bd64f5b4ad435946
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
		// 这将会过滤掉`where`中的所有nil参数。
		// md5:dc90f650b5a33b25
		if isDoStruct(in.Where) {
			in.OmitNil = true
		}
		// 如果`where`结构体实现了`iIterator`接口，
		// 则使用其Iterate函数来遍历键值对。
		// 例如，ListMap和TreeMap是有序映射，
		// 它们实现了`iIterator`接口，并且对where条件的索引友好。
		// md5:d2bd42ea2a41d114
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
		// 自动映射和过滤结构体属性。 md5:8dc7e982c45f4e9d
		var (
			reflectType = reflectInfo.OriginValue.Type()
			structField reflect.StructField
			data        = MapOrStructToMapDeep(in.Where, true)
		)
		// 如果提供了`Prefix`，则检查并获取表名。 md5:68af0bee501f583b
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
		// 如果提供了`Table`，则对字段进行映射和过滤。 md5:369e7f3da1245a29
		if in.Table != "" {
			data, _ = db.GetCore().mappingAndFilterData(ctx, in.Schema, in.Table, data, true)
		}
		// 将结构体属性按顺序放入Where语句中。 md5:e18e7534d834dd8a
		var ormTagValue string
		for i := 0; i < reflectType.NumField(); i++ {
			structField = reflectType.Field(i)
			// 如果指定了，使用`orm`标签的值作为字段名。 md5:0e761199d5be562b
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
		// `whereStr` 是用作键值条件的字段名称吗？
		// 例如：
		// Where("id", 1)
		// Where("id", g.Slice{1,2,3}) 
		// 
		// 这段Go代码中的注释是在询问`whereStr`是否是一个用作键值对条件的字段名。它举例说明了如何使用`where`函数，其中第一个参数是字段名（如"id"），第二个参数可以是单个值（如1）或一个包含多个值的切片（如g.Slice{1,2,3}）。
		// md5:3e3e293b8d2b6e27
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
		// 如果第一部分是列名，它会自动为列添加前缀。 md5:8174a130580bbf74
		if in.Prefix != "" {
			array := gstr.Split(whereStr, " ")
			if ok, _ := db.GetCore().HasField(ctx, in.Table, array[0]); ok {
				whereStr = in.Prefix + "." + whereStr
			}
		}
		// 普通字符串和参数占位符处理。
		// 例如：
		// Where("id in(?) and name=?", g.Slice{1,2,3}, "john")
		// md5:8a2be53569a9ada1
		i := 0
		for {
			if i >= len(in.Args) {
				break
			}
			// ===============================================================
			// 子查询，总是与字符串条件一起使用。
			// ===============================================================
			// md5:3cd7047ec77cba30
			if subModel, ok := in.Args[i].(*Model); ok {
				index := -1
				whereStr, _ = gregex.ReplaceStringFunc(`(\?)`, whereStr, func(s string) string {
					index++
					if i+len(newArgs) == index {
						sqlWithHolder, holderArgs := subModel.getHolderAndArgsAsSubModel(ctx)
						in.Args = gutil.SliceInsertAfter(in.Args, i, holderArgs...)
						// 自动添加括号。 md5:4b202bb8e8a55e8b
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
// 
// 翻译为中文：
// 
// 示例：Where/And/Or("uid>=", 1) 
// 
// 翻译为中文：
// 
// 例如：Where/And/Or("uid>=", 1). md5:97d0042da730e39a
				newWhere += "?"
			} else if gregex.IsMatchString(regularFieldNameRegPattern, newWhere) {
				newWhere = db.GetCore().QuoteString(newWhere)
				if len(newArgs) > 0 {
					if utils.IsArray(newArgs[0]) {
						// 例如：
						// Where("id", []int{1,2,3})
						// Where("user.id", []int{1,2,3})
						// 
						// 这段代码的注释表示示例用法，其中`Where`是一个函数，它接受两个参数：一个是要查询的字段（如"id"或"user.id"），另一个是一组值（如包含1, 2, 3的整数切片）。这通常用于在数据库查询中设置条件，比如筛选id为1, 2或3的记录，或者在"user"表中的"id"字段匹配这些值。
						// md5:5688161e5a37e690
						newWhere += " IN (?)"
					} else if empty.IsNil(newArgs[0]) {
						// 例如：
						// Where("id", nil) 						// 指定 "id" 字段的查询条件为 nil（空）
						// Where("user.id", nil) 						// 指定 "user.id" 字段的查询条件为 nil（空）
						// md5:7b874349f1af2dd8
						newWhere += " IS NULL"
						newArgs = nil
					} else {
						// 例如：
						// Where/And/Or("uid", 1) 						// 指定 "uid" 字段等于 1 的条件
						// Where/And/Or("user.uid", 1) 						// 指定 "user" 对象下的 "uid" 字段等于 1 的条件
						// md5:0809c46d1c195714
						newWhere += "=?"
					}
				}
			}
		}
	}
	return handleSliceAndStructArgsForSql(newWhere, newArgs)
}

// formatWhereInterfaces 将 `where` 格式化为 []interface{}。 md5:6fb34f9561771cdc
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
	Db        DB            // Db 是当前操作的底层数据库对象。 md5:f4b4c46633cb4235
	Buffer    *bytes.Buffer // Buffer是当前操作的SQL语句字符串，不包含Args。 md5:fc54d627bfb62054
	Args      []interface{} // Args是当前操作的全部参数。 md5:e962690161726419
	Key       string        // 字段名称，例如："id"，"name"等。 md5:26a2c4cbd9f18aa7
	Value     interface{}   // 字段值，可以是任何类型。 md5:1edf6819770a85b8
	Type      string        // Where类型中的值。 md5:8becc6ca3981308b
	OmitEmpty bool          // 如果`value`为空，忽略当前条件键。 md5:a7b83f4b09b6f499
	Prefix    string        // 字段前缀，例如："用户"，"订单"等。 md5:53032e7ee552fd8f
}

// formatWhereKeyValue 处理参数映射中的每一组键值对。 md5:5d6f9d3dee346d1a
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
	// 如果值是切片类型，并且键字符串中只有一个'?'占位符，它会根据参数数量自动添加占位符字符，并将其转换为"In"语句。
	// md5:10f5c168c92db7c7
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
				// 键是一个单个字段名。 md5:637e2145bd0e0f57
				in.Buffer.WriteString(quotedKey + " IS NULL")
			} else {
				// 键可能包含操作字符。 md5:1e21edc23189eeb0
				in.Buffer.WriteString(quotedKey)
			}
		} else {
			// 它还支持"LIKE"语句，我们将其视为一种运算符。 md5:71ae1896f0d3b4fd
			quotedKey = gstr.Trim(quotedKey)
			if gstr.Pos(quotedKey, "?") == -1 {
				like := " LIKE"
				if len(quotedKey) > len(like) && gstr.Equal(quotedKey[len(quotedKey)-len(like):], like) {
					// 例如：Where(g.Map{"name like": "john%"}) 
// 
// 这段Go语言的注释表示一个示例用法，其中`Where`是一个函数，它接受一个映射（Map）作为参数，这个映射的键值对是`"name like"` 和 `"john%"`。这意味着在查询时，将对"name"字段进行模糊匹配，查找以"john"开头的记录。 md5:a6037088e14ea97a
					in.Buffer.WriteString(quotedKey + " ?")
				} else if gregex.IsMatchString(lastOperatorRegPattern, quotedKey) {
					// 例如：Where(g.Map{"age > ": 16}) 
// 
// 这段Go语言代码的注释表示这是一个示例（Eg），它使用了一个谓词（Where）和一个映射（Map），这个映射中键值对为 "age > " : 16，意思是筛选出年龄大于16的项。 md5:2b3b5668547eafe7
					in.Buffer.WriteString(quotedKey + " ?")
				} else if gregex.IsMatchString(regularFieldNameRegPattern, in.Key) {
					// 键是一个常规的字段名。 md5:6088ec7a69f84698
					in.Buffer.WriteString(quotedKey + "=?")
				} else {
					// 键不是一个常规的字段名。
					// 例如：Where(g.Map{"age > 16": nil})
					// 问题链接：https:					//github.com/gogf/gf/issues/765
					// md5:79107f0b28e8b612
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

// handleSliceAndStructArgsForSql 是一个重要的函数，它在将 sql 和其所有参数提交给底层驱动程序之前处理它们。
// md5:a6e05a5f78b51a2b
func handleSliceAndStructArgsForSql(
	oldSql string, oldArgs []interface{},
) (newSql string, newArgs []interface{}) {
	newSql = oldSql
	if len(oldArgs) == 0 {
		return
	}
	// insertHolderCount 用于计算 "?" 持有者的插入位置。 md5:878313cc9b1fa5c3
	insertHolderCount := 0
	// 处理切片和结构体类型的参数项。 md5:ce55342863e73d8f
	for index, oldArg := range oldArgs {
		argReflectInfo := reflection.OriginValueAndKind(oldArg)
		switch argReflectInfo.OriginKind {
		case reflect.Slice, reflect.Array:
			// 它不会分割 []byte 类型。
			// 例如：table.Where("name = ?", []byte("john"))
			// md5:05dcc823e289de42
			if _, ok := oldArg.([]byte); ok {
				newArgs = append(newArgs, oldArg)
				continue
			}
			var (
				valueHolderCount = gstr.Count(newSql, "?")
				argSliceLength   = argReflectInfo.OriginValue.Len()
			)
			if argSliceLength == 0 {
				// 空切片参数，它将SQL转换为一个假的SQL。
				// 示例：
				// Query("select * from xxx where id in(?)", g.Slice{}) -> select * from xxx where 0=1
				// Where("id in(?)", g.Slice{}) -> WHERE 0=1
				// 
				// 这里的注释说明了当使用空切片（`g.Slice{}`）作为参数时，Go的某些函数（如`Query`和`Where`）会将SQL中的条件改变为等价于`false`的形式，例如将`in`条件替换为`0=1`，从而使得查询结果为空。
				// md5:020597c0f38437e4
				if gstr.Contains(newSql, "?") {
					whereKeyWord := " WHERE "
					if p := gstr.PosI(newSql, whereKeyWord); p == -1 {
						return "0=1", []interface{}{}
					} else {
						return gstr.SubStr(newSql, 0, p+len(whereKeyWord)) + "0=1", []interface{}{}
					}
				}
			} else {
				// 示例：
				// Query("SELECT ?+?", g.Slice{1,2}) 				// 查询语句，参数为1和2
				// WHERE("id=?", g.Slice{1,2}) 				// WHERE子句，参数为1和2
				// md5:4f9ae718d40ffb8b
				for i := 0; i < argSliceLength; i++ {
					newArgs = append(newArgs, argReflectInfo.OriginValue.Index(i).Interface())
				}
			}

			// 如果 "?" 占位符的数量等于切片的长度，
			// 则不执行参数分割的逻辑。
			// 例如：db.Query("SELECT ?+?", g.Slice{1, 2})
			// md5:aac31c8c27bdcf7d
			if len(oldArgs) == 1 && valueHolderCount == argSliceLength {
				break
			}

			// counter 用于查找 '?' 占位符的插入位置。 md5:22bff4ac2bdd0f47
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
					insertHolderCount += argSliceLength - 1
					return "?" + strings.Repeat(",?", argSliceLength-1)
				}
				return s
			})

		// 特殊的结构体处理。 md5:8911bc2424fd10eb
		case reflect.Struct:
			switch oldArg.(type) {
			// 基础驱动程序支持time.Time类型。 md5:9143055892307413
			case time.Time, *time.Time:
				newArgs = append(newArgs, oldArg)
				continue

			case gtime.Time:
				newArgs = append(newArgs, oldArg.(gtime.Time).Time)
				continue

			case *gtime.Time:
				newArgs = append(newArgs, oldArg.(*gtime.Time).Time)
				continue

			default:
				// 如果结构体实现了String接口，它将默认将结构体转换为字符串。
				// md5:59ba6afad009bc6a
				if v, ok := oldArg.(iString); ok {
					newArgs = append(newArgs, v.String())
					continue
				}
			}
			newArgs = append(newArgs, oldArg)

		default:
			newArgs = append(newArgs, oldArg)
		}
	}
	return
}

// FormatSqlWithArgs 将参数绑定到SQL字符串并返回一个完整的SQL字符串，仅用于调试。
// md5:1453466956e418ba
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
				// 类型为Raw的参数不需要特殊处理. md5:aa477b3ebc58b939
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

// FormatMultiLineSqlToSingle 将多行SQL模板字符串格式化为单行。 md5:cb1180487fd5c495
func FormatMultiLineSqlToSingle(sql string) (string, error) {
	var err error
	// 格式化SQL模板字符串。 md5:77bcc0fc1c095ebd
	sql, err = gregex.ReplaceString(`[\n\r\s]+`, " ", gstr.Trim(sql))
	if err != nil {
		return "", err
	}
	sql, err = gregex.ReplaceString(`\s{2,}`, " ", gstr.Trim(sql))
	if err != nil {
		return "", err
	}
	return sql, nil
}
