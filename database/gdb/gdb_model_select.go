// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gdb

import (
	"context"
	"fmt"
	"reflect"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/reflection"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// 所有执行的是 "FROM ..." 语句针对该模型。
// 它从表中检索记录，并将结果作为切片类型返回。
// 如果根据给定条件从表中没有检索到任何记录，它将返回nil。
// 
// 可选参数 `where` 和 Model.Where 函数的参数相同，
// 请参阅 Model.Where。
// md5:fd88d2addfbe9655
// ff:查询
// m:
// where:查询条件
// Result:
func (m *Model) All(where ...interface{}) (Result, error) {
	var ctx = m.GetCtx()
	return m.doGetAll(ctx, false, where...)
}

// AllAndCount retrieves all records and the total count of records from the model.
// If useFieldForCount is true, it will use the fields specified in the model for counting;
// otherwise, it will use a constant value of 1 for counting.
// It returns the result as a slice of records, the total count of records, and an error if any.
// The where parameter is an optional list of conditions to use when retrieving records.
//
//
//	var model Model
//	var result Result
//	var count int
//	where := []interface{}{"name = ?", "John"}
//	result, count, err := model.AllAndCount(true)
//	if err != nil {
//	    // Handle error.
//	}
//	fmt.Println(result, count)
// ff:查询与行数
// m:
// useFieldForCount:是否用字段计数
// result:结果
// totalCount:行数
// err:错误
func (m *Model) AllAndCount(useFieldForCount bool) (result Result, totalCount int, err error) {
	// 克隆模型用于计数. md5:662b7475962d2c44
	countModel := m.Clone()

	// 如果useFieldForCount为false，将字段设置为计数的恒定值1. md5:2eea55571801d2ab
	if !useFieldForCount {
		countModel.fields = "1"
	}

	// 获取记录的总数. md5:d21517ef51fd67f3
	totalCount, err = countModel.Count()
	if err != nil {
		return
	}

	// 如果总记录数为0，就没有需要检索的记录，因此提前返回. md5:ae90d44fd00f71aa
	if totalCount == 0 {
		return
	}

	// Retrieve all records
	result, err = m.doGetAll(m.GetCtx(), false)
	return
}

// Chunk 使用给定的 `size` 和 `handler` 函数来分块迭代查询结果。 md5:4c5d0d282b8e1fe4
// ff:分割
// m:
// size:数量
// handler:处理函数
func (m *Model) Chunk(size int, handler ChunkHandler) {
	page := m.start
	if page <= 0 {
		page = 1
	}
	model := m
	for {
		model = model.Page(page, size)
		data, err := model.All()
		if err != nil {
			handler(nil, err)
			break
		}
		if len(data) == 0 {
			break
		}
		if !handler(data, err) {
			break
		}
		if len(data) < size {
			break
		}
		page++
	}
}

// 从表中获取一条记录，并将结果作为map类型返回。如果使用给定条件从表中没有检索到记录，则返回nil。
//
// 可选参数`where`与Model.Where函数的参数相同，参见Model.Where。
// md5:b48f8e0c5d07b484
// ff:查询一条
// m:
// where:条件
// Record:
func (m *Model) One(where ...interface{}) (Record, error) {
	var ctx = m.GetCtx()
	if len(where) > 0 {
		return m.Where(where[0], where[1:]...).One()
	}
	all, err := m.doGetAll(ctx, true)
	if err != nil {
		return nil, err
	}
	if len(all) > 0 {
		return all[0], nil
	}
	return nil, nil
}

// 从数据库查询并返回数据值作为切片。
// 注意，如果结果中有多个列，它会随机返回一列的值。
// 
// 如果提供了可选参数 `fieldsAndWhere`，则 fieldsAndWhere[0] 是选择的字段，
// 而 fieldsAndWhere[1:] 则被视为 where 条件字段。
// 参见 Model.Fields 和 Model.Where 函数。
// md5:1de6885dc1e83172
// ff:查询切片
// m:
// fieldsAndWhere:条件
func (m *Model) Array(fieldsAndWhere ...interface{}) ([]Value, error) {
	if len(fieldsAndWhere) > 0 {
		if len(fieldsAndWhere) > 2 {
			return m.Fields(gconv.String(fieldsAndWhere[0])).Where(fieldsAndWhere[1], fieldsAndWhere[2:]...).Array()
		} else if len(fieldsAndWhere) == 2 {
			return m.Fields(gconv.String(fieldsAndWhere[0])).Where(fieldsAndWhere[1]).Array()
		} else {
			return m.Fields(gconv.String(fieldsAndWhere[0])).Array()
		}
	}
	all, err := m.All()
	if err != nil {
		return nil, err
	}
	var (
		field string
		core  = m.db.GetCore()
		ctx   = core.injectInternalColumn(m.GetCtx())
	)
	if len(all) > 0 {
		if internalData := core.getInternalColumnFromCtx(ctx); internalData != nil {
			field = internalData.FirstResultColumn
		} else {
			return nil, gerror.NewCode(
				gcode.CodeInternalError,
				`query array error: the internal context data is missing. there's internal issue should be fixed`,
			)
		}
	}
	return all.Array(field), nil
}

// Struct retrieves one record from table and converts it into given struct.
// The parameter `pointer` should be type of *struct/**struct. If type **struct is given,
// it can create the struct internally during converting.
//
// The optional parameter `where` is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if the given parameter `pointer` pointed to a variable that has
// default value and there's no record retrieved with the given conditions from table.
//
// user := new(User)
// err  := db.Model("user").Where("id", 1).Scan(user)
//
// user := (*User)(nil)
// err  := db.Model("user").Where("id", 1).Scan(&user).
func (m *Model) doStruct(pointer interface{}, where ...interface{}) error {
	model := m
	// 自动通过结构体属性选择字段。 md5:25f031330d67c88b
	if model.fieldsEx == "" && (model.fields == "" || model.fields == "*") {
		if v, ok := pointer.(reflect.Value); ok {
			model = m.Fields(v.Interface())
		} else {
			model = m.Fields(pointer)
		}
	}
	one, err := model.One(where...)
	if err != nil {
		return err
	}
	if err = one.Struct(pointer); err != nil {
		return err
	}
	return model.doWithScanStruct(pointer)
}

// Structs retrieves records from table and converts them into given struct slice.
// The parameter `pointer` should be type of *[]struct/*[]*struct. It can create and fill the struct
// slice internally during converting.
//
// The optional parameter `where` is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if the given parameter `pointer` pointed to a variable that has
// default value and there's no record retrieved with the given conditions from table.
//
// users := ([]User)(nil)
// err   := db.Model("user").Scan(&users)
//
// users := ([]*User)(nil)
// err   := db.Model("user").Scan(&users).
func (m *Model) doStructs(pointer interface{}, where ...interface{}) error {
	model := m
	// 自动通过结构体属性选择字段。 md5:25f031330d67c88b
	if model.fieldsEx == "" && (model.fields == "" || model.fields == "*") {
		if v, ok := pointer.(reflect.Value); ok {
			model = m.Fields(
				reflect.New(
					v.Type().Elem(),
				).Interface(),
			)
		} else {
			model = m.Fields(
				reflect.New(
					reflect.ValueOf(pointer).Elem().Type().Elem(),
				).Interface(),
			)
		}
	}
	all, err := model.All(where...)
	if err != nil {
		return err
	}
	if err = all.Structs(pointer); err != nil {
		return err
	}
	return model.doWithScanStructs(pointer)
}

// Scan automatically calls Struct or Structs function according to the type of parameter `pointer`.
// It calls function doStruct if `pointer` is type of *struct/**struct.
// It calls function doStructs if `pointer` is type of *[]struct/*[]*struct.
//
// The optional parameter `where` is the same as the parameter of Model.Where function,  see Model.Where.
//
// Note that it returns sql.ErrNoRows if the given parameter `pointer` pointed to a variable that has
// default value and there's no record retrieved with the given conditions from table.
//
// user := new(User)
// err  := db.Model("user").Where("id", 1).Scan(user)
//
// user := (*User)(nil)
// err  := db.Model("user").Where("id", 1).Scan(&user)
//
// users := ([]User)(nil)
// err   := db.Model("user").Scan(&users)
//
// users := ([]*User)(nil)
// err   := db.Model("user").Scan(&users).
// ff:查询到结构体指针
// m:
// pointer:数据指针
// where:条件
func (m *Model) Scan(pointer interface{}, where ...interface{}) error {
	reflectInfo := reflection.OriginTypeAndKind(pointer)
	if reflectInfo.InputKind != reflect.Ptr {
		return gerror.NewCode(
			gcode.CodeInvalidParameter,
			`the parameter "pointer" for function Scan should type of pointer`,
		)
	}
	switch reflectInfo.OriginKind {
	case reflect.Slice, reflect.Array:
		return m.doStructs(pointer, where...)

	case reflect.Struct, reflect.Invalid:
		return m.doStruct(pointer, where...)

	default:
		return gerror.NewCode(
			gcode.CodeInvalidParameter,
			`element of parameter "pointer" for function Scan should type of struct/*struct/[]struct/[]*struct`,
		)
	}
}

// ScanAndCount scans a single record or record array that matches the given conditions and counts the total number of records that match those conditions.
// If useFieldForCount is true, it will use the fields specified in the model for counting;
// The pointer parameter is a pointer to a struct that the scanned data will be stored in.
// The pointerCount parameter is a pointer to an integer that will be set to the total number of records that match the given conditions.
// The where parameter is an optional list of conditions to use when retrieving records.
//
//
//	var count int
//	user := new(User)
//	err  := db.Model("user").Where("id", 1).ScanAndCount(user,&count,true)
//	fmt.Println(user, count)
//
// Example Join:
//
//	type User struct {
//		Id       int
//		Passport string
//		Name     string
//		Age      int
//	}
//	var users []User
//	var count int
//	db.Model(table).As("u1").
//		LeftJoin(tableName2, "u2", "u2.id=u1.id").
//		Fields("u1.passport,u1.id,u2.name,u2.age").
//		Where("u1.id<2").
//		ScanAndCount(&users, &count, false)
// ff:查询与行数到指针
// m:
// pointer:数据指针
// totalCount:行数指针
// useFieldForCount:是否用字段计数
// err:错误
func (m *Model) ScanAndCount(pointer interface{}, totalCount *int, useFieldForCount bool) (err error) {
	// 支持使用 * 的字段，例如：.Fields("a.*, b.name")。计数SQL为：select count(1) from xxx. md5:a3fc56bcc1dcba76
	countModel := m.Clone()
	// 如果useFieldForCount为false，将字段设置为计数的恒定值1. md5:2eea55571801d2ab
	if !useFieldForCount {
		countModel.fields = "1"
	}

	// 获取记录的总数. md5:d21517ef51fd67f3
	*totalCount, err = countModel.Count()
	if err != nil {
		return err
	}

	// 如果总记录数为0，就没有需要检索的记录，因此提前返回. md5:ae90d44fd00f71aa
	if *totalCount == 0 {
		return
	}
	err = m.Scan(pointer)
	return
}

// ScanList 将 `r` 转换为包含其他复杂结构体属性的切片。请注意，参数 `listPointer` 的类型应该是 `*[]struct` 或 `*[]*struct`。
// 
// 参见 Result.ScanList。
// md5:4116492a123661b5
// ff:查询到指针列表
// m:
// structSlicePointer:结构体切片指针
// bindToAttrName:绑定到结构体属性名称
// relationAttrNameAndFields:结构体属性关联
// err:错误
func (m *Model) ScanList(structSlicePointer interface{}, bindToAttrName string, relationAttrNameAndFields ...string) (err error) {
	var result Result
	out, err := checkGetSliceElementInfoForScanList(structSlicePointer, bindToAttrName)
	if err != nil {
		return err
	}
	if m.fields != defaultFields || m.fieldsEx != "" {
		// 有自定义字段。 md5:57eb1cc07145128c
		result, err = m.All()
	} else {
		// 使用反射创建的临时结构体过滤字段。 md5:6873597e9de7f128
		result, err = m.Fields(reflect.New(out.BindToAttrType).Interface()).All()
	}
	if err != nil {
		return err
	}
	var (
		relationAttrName string
		relationFields   string
	)
	switch len(relationAttrNameAndFields) {
	case 2:
		relationAttrName = relationAttrNameAndFields[0]
		relationFields = relationAttrNameAndFields[1]
	case 1:
		relationFields = relationAttrNameAndFields[0]
	}
	return doScanList(doScanListInput{
		Model:              m,
		Result:             result,
		StructSlicePointer: structSlicePointer,
		StructSliceValue:   out.SliceReflectValue,
		BindToAttrName:     bindToAttrName,
		RelationAttrName:   relationAttrName,
		RelationFields:     relationFields,
	})
}

// Value 从表中获取指定记录的值，并将结果作为接口类型返回。
// 如果根据给定条件在表中找不到记录，它将返回nil。
//
// 如果提供了可选参数 `fieldsAndWhere`，其中 fieldsAndWhere[0] 是选择的字段，
// 而 fieldsAndWhere[1:] 用作 WHERE 条件字段。
// 另请参阅 Model.Fields 和 Model.Where 函数。
// md5:e6b48ca188d3d208
// ff:查询一条值
// m:
// fieldsAndWhere:字段和条件
// Value:
func (m *Model) Value(fieldsAndWhere ...interface{}) (Value, error) {
	var (
		core = m.db.GetCore()
		ctx  = core.injectInternalColumn(m.GetCtx())
	)
	if len(fieldsAndWhere) > 0 {
		if len(fieldsAndWhere) > 2 {
			return m.Fields(gconv.String(fieldsAndWhere[0])).Where(fieldsAndWhere[1], fieldsAndWhere[2:]...).Value()
		} else if len(fieldsAndWhere) == 2 {
			return m.Fields(gconv.String(fieldsAndWhere[0])).Where(fieldsAndWhere[1]).Value()
		} else {
			return m.Fields(gconv.String(fieldsAndWhere[0])).Value()
		}
	}
	var (
		sqlWithHolder, holderArgs = m.getFormattedSqlAndArgs(ctx, queryTypeValue, true)
		all, err                  = m.doGetAllBySql(ctx, queryTypeValue, sqlWithHolder, holderArgs...)
	)
	if err != nil {
		return nil, err
	}
	if len(all) > 0 {
		if internalData := core.getInternalColumnFromCtx(ctx); internalData != nil {
			if v, ok := all[0][internalData.FirstResultColumn]; ok {
				return v, nil
			}
		} else {
			return nil, gerror.NewCode(
				gcode.CodeInternalError,
				`query value error: the internal context data is missing. there's internal issue should be fixed`,
			)
		}
	}
	return nil, nil
}

// Count 对于该模型执行 "SELECT COUNT(x) FROM ..." 语句。
// 可选参数 `where` 和 Model.Where 函数的参数相同，参见 Model.Where。
// md5:52b3d2e0e43bb2af
// ff:查询行数
// m:
// where:条件
func (m *Model) Count(where ...interface{}) (int, error) {
	var (
		core = m.db.GetCore()
		ctx  = core.injectInternalColumn(m.GetCtx())
	)
	if len(where) > 0 {
		return m.Where(where[0], where[1:]...).Count()
	}
	var (
		sqlWithHolder, holderArgs = m.getFormattedSqlAndArgs(ctx, queryTypeCount, false)
		all, err                  = m.doGetAllBySql(ctx, queryTypeCount, sqlWithHolder, holderArgs...)
	)
	if err != nil {
		return 0, err
	}
	if len(all) > 0 {
		if internalData := core.getInternalColumnFromCtx(ctx); internalData != nil {
			if v, ok := all[0][internalData.FirstResultColumn]; ok {
				return v.Int(), nil
			}
		} else {
			return 0, gerror.NewCode(
				gcode.CodeInternalError,
				`query count error: the internal context data is missing. there's internal issue should be fixed`,
			)
		}
	}
	return 0, nil
}

// CountColumn 执行对模型的 "SELECT COUNT(x) FROM ..." 语句。 md5:150abf4737c4588c
// ff:查询字段行数
// m:
// column:字段名称
func (m *Model) CountColumn(column string) (int, error) {
	if len(column) == 0 {
		return 0, nil
	}
	return m.Fields(column).Count()
}

// Min 为模型执行 "SELECT MIN(x) FROM ..." 语句。 md5:e2fc098c542503d1
// ff:查询最小值
// m:
// column:字段名称
func (m *Model) Min(column string) (float64, error) {
	if len(column) == 0 {
		return 0, nil
	}
	value, err := m.Fields(fmt.Sprintf(`MIN(%s)`, m.QuoteWord(column))).Value()
	if err != nil {
		return 0, err
	}
	return value.Float64(), err
}

// Max 对模型执行 "SELECT MAX(x) FROM ..." 语句。 md5:bb6b4d0dc51fbfaf
// ff:查询最大值
// m:
// column:字段名称
func (m *Model) Max(column string) (float64, error) {
	if len(column) == 0 {
		return 0, nil
	}
	value, err := m.Fields(fmt.Sprintf(`MAX(%s)`, m.QuoteWord(column))).Value()
	if err != nil {
		return 0, err
	}
	return value.Float64(), err
}

// Avg 对于该模型执行"SELECT AVG(x) FROM ..." 语句。 md5:9b360a11d26d6fca
// ff:查询平均值
// m:
// column:字段名称
func (m *Model) Avg(column string) (float64, error) {
	if len(column) == 0 {
		return 0, nil
	}
	value, err := m.Fields(fmt.Sprintf(`AVG(%s)`, m.QuoteWord(column))).Value()
	if err != nil {
		return 0, err
	}
	return value.Float64(), err
}

// Sum 对于该模型执行 "SELECT SUM(x) FROM ..." 语句。 md5:bcbe9e29cd168603
// ff:查询求和
// m:
// column:字段名称
func (m *Model) Sum(column string) (float64, error) {
	if len(column) == 0 {
		return 0, nil
	}
	value, err := m.Fields(fmt.Sprintf(`SUM(%s)`, m.QuoteWord(column))).Value()
	if err != nil {
		return 0, err
	}
	return value.Float64(), err
}

// Union 为模型执行 "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." 语句。 md5:97431dccd533414e
// ff:多表去重查询
// m:
// unions:Model对象
func (m *Model) Union(unions ...*Model) *Model {
	return m.db.Union(unions...)
}

// UnionAll 对模型执行 "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ..." 语句。 md5:d112aec0d1929661
// ff:多表查询
// m:
// unions:Model对象
func (m *Model) UnionAll(unions ...*Model) *Model {
	return m.db.UnionAll(unions...)
}

// Limit 设置模型的 "LIMIT" 语句。
// 参数 `limit` 可以是一个或两个数字。如果传递两个数字，它将为模型设置 "LIMIT limit[0],limit[1]" 语句；否则，它将设置 "LIMIT limit[0]" 语句。
// md5:fd06ed75a128d403
// ff:设置条数
// m:
// limit:条数或两个数字
func (m *Model) Limit(limit ...int) *Model {
	model := m.getModel()
	switch len(limit) {
	case 1:
		model.limit = limit[0]
	case 2:
		model.start = limit[0]
		model.limit = limit[1]
	}
	return model
}

// Offset 设置模型的“OFFSET”语句。它只对某些数据库（如 SQLServer、PostgreSQL 等）有意义。
// md5:5a99cab6ce558c69
// ff:
// m:
// offset:
func (m *Model) Offset(offset int) *Model {
	model := m.getModel()
	model.offset = offset
	return model
}

// Distinct 强制查询仅返回不重复的结果。 md5:ead62c0e4b4795ab
// ff:设置去重
// m:
func (m *Model) Distinct() *Model {
	model := m.getModel()
	model.distinct = "DISTINCT "
	return model
}

// Page 设置模型的分页号。
// 参数 `page` 的起始值为1，用于分页。
// 注意，这与Limit函数在"LIMIT"语句中从0开始不同。
// md5:02b920e99951ce53
// ff:设置分页
// m:
// page:第几页
// limit:条数
func (m *Model) Page(page, limit int) *Model {
	model := m.getModel()
	if page <= 0 {
		page = 1
	}
	model.start = (page - 1) * limit
	model.limit = limit
	return model
}

// Having 设置模型的having语句。
// 该函数的使用参数与Where函数相同。
// 参见Where。
// md5:b4e737511765f79f
// ff:设置分组条件
// m:
// having:条件
// args:参数
func (m *Model) Having(having interface{}, args ...interface{}) *Model {
	model := m.getModel()
	model.having = []interface{}{
		having, args,
	}
	return model
}

// doGetAll 对应于 "SELECT FROM ..." 语句，用于模型。
// 它从表中检索记录，并以切片类型返回结果。如果根据给定条件从表中没有检索到记录，则返回 nil。
// 
// 参数 `limit1` 指定当模型的 `limit` 未设置时，是否只查询一条记录。
// 可选参数 `where` 的用法与 Model.Where 函数的参数相同，参见 Model.Where。
// md5:d4f7ecca6c5aaa48
func (m *Model) doGetAll(ctx context.Context, limit1 bool, where ...interface{}) (Result, error) {
	if len(where) > 0 {
		return m.Where(where[0], where[1:]...).All()
	}
	sqlWithHolder, holderArgs := m.getFormattedSqlAndArgs(ctx, queryTypeNormal, limit1)
	return m.doGetAllBySql(ctx, queryTypeNormal, sqlWithHolder, holderArgs...)
}

// doGetAllBySql 在数据库上执行选择语句。 md5:b9498c08926ceb6a
func (m *Model) doGetAllBySql(ctx context.Context, queryType queryType, sql string, args ...interface{}) (result Result, err error) {
	if result, err = m.getSelectResultFromCache(ctx, sql, args...); err != nil || result != nil {
		return
	}

	in := &HookSelectInput{
		internalParamHookSelect: internalParamHookSelect{
			internalParamHook: internalParamHook{
				link: m.getLink(false),
			},
			handler: m.hookHandler.Select,
		},
		Model: m,
		Table: m.tables,
		Sql:   sql,
		Args:  m.mergeArguments(args),
	}
	if result, err = in.Next(ctx); err != nil {
		return
	}

	err = m.saveSelectResultToCache(ctx, queryType, result, sql, args...)
	return
}

func (m *Model) getFormattedSqlAndArgs(
	ctx context.Context, queryType queryType, limit1 bool,
) (sqlWithHolder string, holderArgs []interface{}) {
	switch queryType {
	case queryTypeCount:
		queryFields := "COUNT(1)"
		if m.fields != "" && m.fields != "*" {
// 不要在这里引用m.fields，以防出现如下的字段情况：
// DISTINCT t.user_id uid
// md5:97ff3b5639a12242
			queryFields = fmt.Sprintf(`COUNT(%s%s)`, m.distinct, m.fields)
		}
		// Raw SQL Model.
		if m.rawSql != "" {
			sqlWithHolder = fmt.Sprintf("SELECT %s FROM (%s) AS T", queryFields, m.rawSql)
			return sqlWithHolder, nil
		}
		conditionWhere, conditionExtra, conditionArgs := m.formatCondition(ctx, false, true)
		sqlWithHolder = fmt.Sprintf("SELECT %s FROM %s%s", queryFields, m.tables, conditionWhere+conditionExtra)
		if len(m.groupBy) > 0 {
			sqlWithHolder = fmt.Sprintf("SELECT COUNT(1) FROM (%s) count_alias", sqlWithHolder)
		}
		return sqlWithHolder, conditionArgs

	default:
		conditionWhere, conditionExtra, conditionArgs := m.formatCondition(ctx, limit1, false)
		// 原生SQL模型，特别适用于包含UNION/UNION ALL特性的SQL。 md5:03942fe59d08c0b4
		if m.rawSql != "" {
			sqlWithHolder = fmt.Sprintf(
				"%s%s",
				m.rawSql,
				conditionWhere+conditionExtra,
			)
			return sqlWithHolder, conditionArgs
		}
		// 请不要在 m.fields 中引用，例如：
		// 如果字段为：
		// DISTINCT t.user_id uid
		// md5:e3b773558c54f2eb
		sqlWithHolder = fmt.Sprintf(
			"SELECT %s%s FROM %s%s",
			m.distinct, m.getFieldsFiltered(), m.tables, conditionWhere+conditionExtra,
		)
		return sqlWithHolder, conditionArgs
	}
}

func (m *Model) getHolderAndArgsAsSubModel(ctx context.Context) (holder string, args []interface{}) {
	holder, args = m.getFormattedSqlAndArgs(
		ctx, queryTypeNormal, false,
	)
	args = m.mergeArguments(args)
	return
}

func (m *Model) getAutoPrefix() string {
	autoPrefix := ""
	if gstr.Contains(m.tables, " JOIN ") {
		autoPrefix = m.db.GetCore().QuoteWord(
			m.db.GetCore().guessPrimaryTableName(m.tablesInit),
		)
	}
	return autoPrefix
}

// getFieldsFiltered 检查字段和字段排除属性，过滤并返回那些将真正被提交到底层数据库驱动的字段。
// md5:e8c5bf23790637e0
func (m *Model) getFieldsFiltered() string {
	if m.fieldsEx == "" {
		// 没有过滤，包含特殊字符。 md5:f2ccc24dfd015b85
		if gstr.ContainsAny(m.fields, "()") {
			return m.fields
		}
		// No filtering.
		if !gstr.ContainsAny(m.fields, ". ") {
			return m.db.GetCore().QuoteString(m.fields)
		}
		return m.fields
	}
	var (
		fieldsArray []string
		fieldsExSet = gset.NewStrSetFrom(gstr.SplitAndTrim(m.fieldsEx, ","))
	)
	if m.fields != "*" {
		// 使用fieldEx过滤自定义字段。 md5:edee7113e1c2daf9
		fieldsArray = make([]string, 0, 8)
		for _, v := range gstr.SplitAndTrim(m.fields, ",") {
			fieldsArray = append(fieldsArray, v[gstr.PosR(v, "-")+1:])
		}
	} else {
		if gstr.Contains(m.tables, " ") {
			panic("function FieldsEx supports only single table operations")
		}
		// 使用fieldEx过滤表格字段。 md5:e15e7d68ef0a3c54
		tableFields, err := m.TableFields(m.tablesInit)
		if err != nil {
			panic(err)
		}
		if len(tableFields) == 0 {
			panic(fmt.Sprintf(`empty table fields for table "%s"`, m.tables))
		}
		fieldsArray = make([]string, len(tableFields))
		for k, v := range tableFields {
			fieldsArray[v.Index] = k
		}
	}
	newFields := ""
	for _, k := range fieldsArray {
		if fieldsExSet.Contains(k) {
			continue
		}
		if len(newFields) > 0 {
			newFields += ","
		}
		newFields += m.db.GetCore().QuoteWord(k)
	}
	return newFields
}

// formatCondition 格式化模型的where参数，并返回一个新的条件SQL及其参数。
// 注意，此函数不会更改`m`的任何属性值。
//
// 参数 `limit1` 指定如果m.limit未设置，是否限制只查询一条记录。
// md5:d251ca8a182de4ff
func (m *Model) formatCondition(
	ctx context.Context, limit1 bool, isCountStatement bool,
) (conditionWhere string, conditionExtra string, conditionArgs []interface{}) {
	var autoPrefix = m.getAutoPrefix()
	// GROUP BY.
	if m.groupBy != "" {
		conditionExtra += " GROUP BY " + m.groupBy
	}
	// WHERE
	conditionWhere, conditionArgs = m.whereBuilder.Build()
	softDeletingCondition := m.softTimeMaintainer().GetWhereConditionForDelete(ctx)
	if m.rawSql != "" && conditionWhere != "" {
		if gstr.ContainsI(m.rawSql, " WHERE ") {
			conditionWhere = " AND " + conditionWhere
		} else {
			conditionWhere = " WHERE " + conditionWhere
		}
	} else if !m.unscoped && softDeletingCondition != "" {
		if conditionWhere == "" {
			conditionWhere = fmt.Sprintf(` WHERE %s`, softDeletingCondition)
		} else {
			conditionWhere = fmt.Sprintf(` WHERE (%s) AND %s`, conditionWhere, softDeletingCondition)
		}
	} else {
		if conditionWhere != "" {
			conditionWhere = " WHERE " + conditionWhere
		}
	}
	// HAVING.
	if len(m.having) > 0 {
		havingHolder := WhereHolder{
			Where:  m.having[0],
			Args:   gconv.Interfaces(m.having[1]),
			Prefix: autoPrefix,
		}
		havingStr, havingArgs := formatWhereHolder(ctx, m.db, formatWhereHolderInput{
			WhereHolder: havingHolder,
			OmitNil:     m.option&optionOmitNilWhere > 0,
			OmitEmpty:   m.option&optionOmitEmptyWhere > 0,
			Schema:      m.schema,
			Table:       m.tables,
		})
		if len(havingStr) > 0 {
			conditionExtra += " HAVING " + havingStr
			conditionArgs = append(conditionArgs, havingArgs...)
		}
	}
	// ORDER BY.
	if !isCountStatement { // SQLServer 的 count 语句中不能包含 order by 子句. md5:a176c1f7165860e0
		if m.orderBy != "" {
			conditionExtra += " ORDER BY " + m.orderBy
		}
	}
	// LIMIT.
	if !isCountStatement {
		if m.limit != 0 {
			if m.start >= 0 {
				conditionExtra += fmt.Sprintf(" LIMIT %d,%d", m.start, m.limit)
			} else {
				conditionExtra += fmt.Sprintf(" LIMIT %d", m.limit)
			}
		} else if limit1 {
			conditionExtra += " LIMIT 1"
		}

		if m.offset >= 0 {
			conditionExtra += fmt.Sprintf(" OFFSET %d", m.offset)
		}
	}

	if m.lockInfo != "" {
		conditionExtra += " " + m.lockInfo
	}
	return
}
