// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

import (
	"context"
	"fmt"
	"reflect"

	gset "github.com/888go/goframe/container/gset"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/reflection"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
)

// 所有执行的是 "FROM ..." 语句针对该模型。
// 它从表中检索记录，并将结果作为切片类型返回。
// 如果根据给定条件从表中没有检索到任何记录，它将返回nil。
// 
// 可选参数 `where` 和 Model.Where 函数的参数相同，
// 请参阅 Model.Where。
// md5:fd88d2addfbe9655
func (m *Model) X查询(查询条件 ...interface{}) (Result, error) {
	var ctx = m.X取上下文对象()
	return m.doGetAll(ctx, false, 查询条件...)
}

// X查询与行数 从模型中检索所有记录以及记录的总数量。
// 如果 useFieldForCount 为真，它将使用模型中指定的字段进行计数；
// 否则，它将使用常数值1来进行计数。
// 此方法返回结果作为一个记录切片，记录的总数量，以及可能存在的错误。
// where 参数是一个可选的条件列表，用于在检索记录时应用。
//
// 示例：
//
//	var model Model
//	var result []Record
//	var count int
//	where := []interface{}{"name = ?", "John"}
//	result, count, err := model.X查询与行数(true, where...)
//	if err != nil {
//	    // 处理错误。
//	}
//	fmt.Println(result, count)
// md5:b631bbec9e186f68
func (m *Model) X查询与行数(是否用字段计数 bool) (结果 Result, 行数 int, 错误 error) {
			// 克隆模型用于计数. md5:662b7475962d2c44
	countModel := m.X取副本()

		// 如果useFieldForCount为false，将字段设置为计数的恒定值1. md5:2eea55571801d2ab
	if !是否用字段计数 {
		countModel.fields = "1"
	}

			// 获取记录的总数. md5:d21517ef51fd67f3
	行数, 错误 = countModel.X查询行数()
	if 错误 != nil {
		return
	}

		// 如果总记录数为0，就没有需要检索的记录，因此提前返回. md5:ae90d44fd00f71aa
	if 行数 == 0 {
		return
	}

	// Retrieve all records
	结果, 错误 = m.doGetAll(m.X取上下文对象(), false)
	return
}

// X分割 使用给定的 `size` 和 `handler` 函数来分块迭代查询结果。 md5:4c5d0d282b8e1fe4
func (m *Model) X分割(数量 int, 处理函数 ChunkHandler) {
	page := m.start
	if page <= 0 {
		page = 1
	}
	model := m
	for {
		model = model.X设置分页(page, 数量)
		data, err := model.X查询()
		if err != nil {
			处理函数(nil, err)
			break
		}
		if len(data) == 0 {
			break
		}
		if !处理函数(data, err) {
			break
		}
		if len(data) < 数量 {
			break
		}
		page++
	}
}

// 从表中获取一条记录，并将结果作为map类型返回。如果使用给定条件从表中没有检索到记录，则返回nil。
//
// 可选参数`where`与Model.Where函数的参数相同，参见Model.Where。
// md5:b48f8e0c5d07b484
func (m *Model) X查询一条(条件 ...interface{}) (Record, error) {
	var ctx = m.X取上下文对象()
	if len(条件) > 0 {
		return m.X条件(条件[0], 条件[1:]...).X查询一条()
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
func (m *Model) X查询切片(条件 ...interface{}) ([]Value, error) {
	if len(条件) > 0 {
		if len(条件) > 2 {
			return m.X字段保留过滤(gconv.String(条件[0])).X条件(条件[1], 条件[2:]...).X查询切片()
		} else if len(条件) == 2 {
			return m.X字段保留过滤(gconv.String(条件[0])).X条件(条件[1]).X查询切片()
		} else {
			return m.X字段保留过滤(gconv.String(条件[0])).X查询切片()
		}
	}
	all, err := m.X查询()
	if err != nil {
		return nil, err
	}
	var (
		field string
		core  = m.db.X取Core对象()
		ctx   = core.injectInternalColumn(m.X取上下文对象())
	)
	if len(all) > 0 {
		if internalData := core.getInternalColumnFromCtx(ctx); internalData != nil {
			field = internalData.FirstResultColumn
		} else {
			return nil, gerror.X创建错误码(
				gcode.CodeInternalError,
				`query array error: the internal context data is missing. there's internal issue should be fixed`,
			)
		}
	}
	return all.X取字段切片(field), nil
}

// Struct 从表中检索一条记录，并将其转换为给定的结构体。
// 参数 `pointer` 应为 *struct 或 **struct 类型。如果给出 **struct 类型，
// 则在转换过程中可以在内部创建该结构体。
//
// 可选参数 `where` 与 Model.Where 函数的参数相同，
// 详情请参阅 Model.Where。
//
// 注意，如果给定参数 `pointer` 指向一个具有默认值的变量，并且根据给定条件从表中没有检索到任何记录，
// 则它将返回 sql.ErrNoRows 错误。
//
// 示例：
// user := new(User)
// err  := db.Model("user").Where("id", 1).Scan(user)
//
// user := (*User)(nil)
// err  := db.Model("user").Where("id", 1).Scan(&user)
// md5:473a4005864a522f
func (m *Model) doStruct(pointer interface{}, where ...interface{}) error {
	model := m
		// 自动通过结构体属性选择字段。 md5:25f031330d67c88b
	if model.fieldsEx == "" && (model.fields == "" || model.fields == "*") {
		if v, ok := pointer.(reflect.Value); ok {
			model = m.X字段保留过滤(v.Interface())
		} else {
			model = m.X字段保留过滤(pointer)
		}
	}
	one, err := model.X查询一条(where...)
	if err != nil {
		return err
	}
	if err = one.X取结构体指针(pointer); err != nil {
		return err
	}
	return model.doWithScanStruct(pointer)
}

// Structs 从表中检索记录并将其转换为给定的结构体切片。
// 参数 `pointer` 应该是类型为 *[]struct 或 *[]*struct。在转换过程中，它可以内部创建和填充结构体切片。
//
// 可选参数 `where` 和 Model.Where 函数的参数相同，参见 Model.Where。
//
// 注意，如果给定的 `pointer` 指向一个具有默认值的变量，并且根据给定条件从表中没有检索到任何记录，则返回 sql.ErrNoRows。
//
// 示例：
// users := ([]User)(nil)
// err   := db.Model("user").Scan(&users)
//
// users := ([]*User)(nil)
// err   := db.Model("user").Scan(&users)
// md5:bd3102709ae8c192
func (m *Model) doStructs(pointer interface{}, where ...interface{}) error {
	model := m
		// 自动通过结构体属性选择字段。 md5:25f031330d67c88b
	if model.fieldsEx == "" && (model.fields == "" || model.fields == "*") {
		if v, ok := pointer.(reflect.Value); ok {
			model = m.X字段保留过滤(
				reflect.New(
					v.Type().Elem(),
				).Interface(),
			)
		} else {
			model = m.X字段保留过滤(
				reflect.New(
					reflect.ValueOf(pointer).Elem().Type().Elem(),
				).Interface(),
			)
		}
	}
	all, err := model.X查询(where...)
	if err != nil {
		return err
	}
	if err = all.X取切片结构体指针(pointer); err != nil {
		return err
	}
	return model.doWithScanStructs(pointer)
}

// X查询到结构体指针会根据参数`pointer`的类型自动调用Struct或Structs函数。
// 如果`pointer`是类型`*struct`或`**struct`，它会调用doStruct函数。
// 如果`pointer`是类型`*[]struct`或`*[]*struct`，它会调用doStructs函数。
// 
// 可选参数`where`与Model.Where函数的参数相同，参见Model.Where。
// 
// 注意，如果给定的`pointer`指向一个具有默认值的变量，并且在表中没有满足条件的记录，它将返回sql.ErrNoRows错误。
// 
// 示例：
// user := new(User)
// err := db.Model("user").Where("id", 1).X查询到结构体指针(user)
// 
// user := (*User)(nil)
// err := db.Model("user").Where("id", 1).X查询到结构体指针(&user)
// 
// users := ([]User)(nil)
// err := db.Model("user").X查询到结构体指针(&users)
// 
// users := ([]*User)(nil)
// err := db.Model("user").X查询到结构体指针(&users)
// md5:a6df07ddafe5975a
func (m *Model) X查询到结构体指针(数据指针 interface{}, 条件 ...interface{}) error {
	reflectInfo := reflection.OriginTypeAndKind(数据指针)
	if reflectInfo.InputKind != reflect.Ptr {
		return gerror.X创建错误码(
			gcode.CodeInvalidParameter,
			`the parameter "pointer" for function Scan should type of pointer`,
		)
	}
	switch reflectInfo.OriginKind {
	case reflect.Slice, reflect.Array:
		return m.doStructs(数据指针, 条件...)

	case reflect.Struct, reflect.Invalid:
		return m.doStruct(数据指针, 条件...)

	default:
		return gerror.X创建错误码(
			gcode.CodeInvalidParameter,
			`element of parameter "pointer" for function Scan should type of struct/*struct/[]struct/[]*struct`,
		)
	}
}

// X查询与行数到指针 扫描与给定条件匹配的单条记录或记录数组，并计算符合这些条件的总记录数。
// 如果 useFieldForCount 为 true，它将使用模型中指定的字段进行计数；
// pointer 参数是一个指向结构体的指针，用于存储扫描到的数据。
// pointerCount 参数是一个指向整数的指针，将被设置为符合给定条件的总记录数。
// where 参数是可选的条件列表，用于在检索记录时使用。
//
// 示例：
//
//	var count int
//	user := new(User)
//	err  := db.Model("user").Where("id", 1).X查询与行数到指针(user, &count, true)
//	fmt.Println(user, count)
//
// 示例（联接）：
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
//		X查询与行数到指针(&users, &count, false)
// md5:984fa8f0e50708f4
func (m *Model) X查询与行数到指针(数据指针 interface{}, 行数指针 *int, 是否用字段计数 bool) (错误 error) {
		// 支持使用 * 的字段，例如：.Fields("a.*, b.name")。计数SQL为：select count(1) from xxx. md5:a3fc56bcc1dcba76
	countModel := m.X取副本()
		// 如果useFieldForCount为false，将字段设置为计数的恒定值1. md5:2eea55571801d2ab
	if !是否用字段计数 {
		countModel.fields = "1"
	}

			// 获取记录的总数. md5:d21517ef51fd67f3
	*行数指针, 错误 = countModel.X查询行数()
	if 错误 != nil {
		return 错误
	}

		// 如果总记录数为0，就没有需要检索的记录，因此提前返回. md5:ae90d44fd00f71aa
	if *行数指针 == 0 {
		return
	}
	错误 = m.X查询到结构体指针(数据指针)
	return
}

// X查询到指针列表 将 `r` 转换为包含其他复杂结构体属性的切片。请注意，参数 `listPointer` 的类型应该是 `*[]struct` 或 `*[]*struct`。
// 
// 参见 Result.X查询到指针列表。
// md5:4116492a123661b5
func (m *Model) X查询到指针列表(结构体切片指针 interface{}, 绑定到结构体属性名称 string, 结构体属性关联 ...string) (错误 error) {
	var result Result
	out, 错误 := checkGetSliceElementInfoForScanList(结构体切片指针, 绑定到结构体属性名称)
	if 错误 != nil {
		return 错误
	}
	if m.fields != defaultFields || m.fieldsEx != "" {
				// 有自定义字段。 md5:57eb1cc07145128c
		result, 错误 = m.X查询()
	} else {
				// 使用反射创建的临时结构体过滤字段。 md5:6873597e9de7f128
		result, 错误 = m.X字段保留过滤(reflect.New(out.BindToAttrType).Interface()).X查询()
	}
	if 错误 != nil {
		return 错误
	}
	var (
		relationAttrName string
		relationFields   string
	)
	switch len(结构体属性关联) {
	case 2:
		relationAttrName = 结构体属性关联[0]
		relationFields = 结构体属性关联[1]
	case 1:
		relationFields = 结构体属性关联[0]
	}
	return doScanList(doScanListInput{
		Model:              m,
		Result:             result,
		StructSlicePointer: 结构体切片指针,
		StructSliceValue:   out.SliceReflectValue,
		BindToAttrName:     绑定到结构体属性名称,
		RelationAttrName:   relationAttrName,
		RelationFields:     relationFields,
	})
}

// X查询一条值 从表中获取指定记录的值，并将结果作为接口类型返回。
// 如果根据给定条件在表中找不到记录，它将返回nil。
//
// 如果提供了可选参数 `fieldsAndWhere`，其中 fieldsAndWhere[0] 是选择的字段，
// 而 fieldsAndWhere[1:] 用作 WHERE 条件字段。
// 另请参阅 Model.Fields 和 Model.Where 函数。
// md5:e6b48ca188d3d208
func (m *Model) X查询一条值(字段和条件 ...interface{}) (Value, error) {
	var (
		core = m.db.X取Core对象()
		ctx  = core.injectInternalColumn(m.X取上下文对象())
	)
	if len(字段和条件) > 0 {
		if len(字段和条件) > 2 {
			return m.X字段保留过滤(gconv.String(字段和条件[0])).X条件(字段和条件[1], 字段和条件[2:]...).X查询一条值()
		} else if len(字段和条件) == 2 {
			return m.X字段保留过滤(gconv.String(字段和条件[0])).X条件(字段和条件[1]).X查询一条值()
		} else {
			return m.X字段保留过滤(gconv.String(字段和条件[0])).X查询一条值()
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
			return nil, gerror.X创建错误码(
				gcode.CodeInternalError,
				`query value error: the internal context data is missing. there's internal issue should be fixed`,
			)
		}
	}
	return nil, nil
}

// X查询行数 对于该模型执行 "SELECT COUNT(x) FROM ..." 语句。
// 可选参数 `where` 和 Model.Where 函数的参数相同，参见 Model.Where。
// md5:52b3d2e0e43bb2af
func (m *Model) X查询行数(条件 ...interface{}) (int, error) {
	var (
		core = m.db.X取Core对象()
		ctx  = core.injectInternalColumn(m.X取上下文对象())
	)
	if len(条件) > 0 {
		return m.X条件(条件[0], 条件[1:]...).X查询行数()
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
				return v.X取整数(), nil
			}
		} else {
			return 0, gerror.X创建错误码(
				gcode.CodeInternalError,
				`query count error: the internal context data is missing. there's internal issue should be fixed`,
			)
		}
	}
	return 0, nil
}

// X查询字段行数 执行对模型的 "SELECT COUNT(x) FROM ..." 语句。 md5:150abf4737c4588c
func (m *Model) X查询字段行数(字段名称 string) (int, error) {
	if len(字段名称) == 0 {
		return 0, nil
	}
	return m.X字段保留过滤(字段名称).X查询行数()
}

// X查询最小值 为模型执行 "SELECT MIN(x) FROM ..." 语句。 md5:e2fc098c542503d1
func (m *Model) X查询最小值(字段名称 string) (float64, error) {
	if len(字段名称) == 0 {
		return 0, nil
	}
	value, err := m.X字段保留过滤(fmt.Sprintf(`MIN(%s)`, m.X底层QuoteWord(字段名称))).X查询一条值()
	if err != nil {
		return 0, err
	}
	return value.X取小数64位(), err
}

// X查询最大值 对模型执行 "SELECT MAX(x) FROM ..." 语句。 md5:bb6b4d0dc51fbfaf
func (m *Model) X查询最大值(字段名称 string) (float64, error) {
	if len(字段名称) == 0 {
		return 0, nil
	}
	value, err := m.X字段保留过滤(fmt.Sprintf(`MAX(%s)`, m.X底层QuoteWord(字段名称))).X查询一条值()
	if err != nil {
		return 0, err
	}
	return value.X取小数64位(), err
}

// X查询平均值 对于该模型执行"SELECT AVG(x) FROM ..." 语句。 md5:9b360a11d26d6fca
func (m *Model) X查询平均值(字段名称 string) (float64, error) {
	if len(字段名称) == 0 {
		return 0, nil
	}
	value, err := m.X字段保留过滤(fmt.Sprintf(`AVG(%s)`, m.X底层QuoteWord(字段名称))).X查询一条值()
	if err != nil {
		return 0, err
	}
	return value.X取小数64位(), err
}

// X查询求和 对于该模型执行 "SELECT SUM(x) FROM ..." 语句。 md5:bcbe9e29cd168603
func (m *Model) X查询求和(字段名称 string) (float64, error) {
	if len(字段名称) == 0 {
		return 0, nil
	}
	value, err := m.X字段保留过滤(fmt.Sprintf(`SUM(%s)`, m.X底层QuoteWord(字段名称))).X查询一条值()
	if err != nil {
		return 0, err
	}
	return value.X取小数64位(), err
}

// X多表去重查询 为模型执行 "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." 语句。 md5:97431dccd533414e
func (m *Model) X多表去重查询(Model对象 ...*Model) *Model {
	return m.db.X多表去重查询(Model对象...)
}

// X多表查询 对模型执行 "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ..." 语句。 md5:d112aec0d1929661
func (m *Model) X多表查询(Model对象 ...*Model) *Model {
	return m.db.X多表查询(Model对象...)
}

// X设置条数 设置模型的 "LIMIT" 语句。
// 参数 `limit` 可以是一个或两个数字。如果传递两个数字，它将为模型设置 "LIMIT limit[0],limit[1]" 语句；否则，它将设置 "LIMIT limit[0]" 语句。
// md5:fd06ed75a128d403
func (m *Model) X设置条数(条数或两个数字 ...int) *Model {
	model := m.getModel()
	switch len(条数或两个数字) {
	case 1:
		model.limit = 条数或两个数字[0]
	case 2:
		model.start = 条数或两个数字[0]
		model.limit = 条数或两个数字[1]
	}
	return model
}

// Offset 设置模型的“OFFSET”语句。它只对某些数据库（如 SQLServer、PostgreSQL 等）有意义。
// md5:5a99cab6ce558c69
func (m *Model) Offset(offset int) *Model {
	model := m.getModel()
	model.offset = offset
	return model
}

// X设置去重 强制查询仅返回不重复的结果。 md5:ead62c0e4b4795ab
func (m *Model) X设置去重() *Model {
	model := m.getModel()
	model.distinct = "DISTINCT "
	return model
}

// X设置分页 设置模型的分页号。
// 参数 `page` 的起始值为1，用于分页。
// 注意，这与Limit函数在"LIMIT"语句中从0开始不同。
// md5:02b920e99951ce53
func (m *Model) X设置分页(第几页, 条数 int) *Model {
	model := m.getModel()
	if 第几页 <= 0 {
		第几页 = 1
	}
	model.start = (第几页 - 1) * 条数
	model.limit = 条数
	return model
}

// X设置分组条件 设置模型的having语句。
// 该函数的使用参数与Where函数相同。
// 参见Where。
// md5:b4e737511765f79f
func (m *Model) X设置分组条件(条件 interface{}, 参数 ...interface{}) *Model {
	model := m.getModel()
	model.having = []interface{}{
		条件, 参数,
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
		return m.X条件(where[0], where[1:]...).X查询()
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
	if gstr.X是否包含(m.tables, " JOIN ") {
		autoPrefix = m.db.X取Core对象().X底层QuoteWord(
			m.db.X取Core对象().guessPrimaryTableName(m.tablesInit),
		)
	}
	return autoPrefix
}

// getFieldsFiltered 检查字段和字段排除属性，过滤并返回那些将真正被提交到底层数据库驱动的字段。
// md5:e8c5bf23790637e0
func (m *Model) getFieldsFiltered() string {
	if m.fieldsEx == "" {
				// 没有过滤，包含特殊字符。 md5:f2ccc24dfd015b85
		if gstr.X是否包含Any(m.fields, "()") {
			return m.fields
		}
		// No filtering.
		if !gstr.X是否包含Any(m.fields, ". ") {
			return m.db.X取Core对象().X底层QuoteString(m.fields)
		}
		return m.fields
	}
	var (
		fieldsArray []string
		fieldsExSet = gset.X创建文本并按值(gstr.X分割并忽略空值(m.fieldsEx, ","))
	)
	if m.fields != "*" {
				// 使用fieldEx过滤自定义字段。 md5:edee7113e1c2daf9
		fieldsArray = make([]string, 0, 8)
		for _, v := range gstr.X分割并忽略空值(m.fields, ",") {
			fieldsArray = append(fieldsArray, v[gstr.X倒找(v, "-")+1:])
		}
	} else {
		if gstr.X是否包含(m.tables, " ") {
			panic("function FieldsEx supports only single table operations")
		}
				// 使用fieldEx过滤表格字段。 md5:e15e7d68ef0a3c54
		tableFields, err := m.X取表字段信息Map(m.tablesInit)
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
		if fieldsExSet.X是否存在(k) {
			continue
		}
		if len(newFields) > 0 {
			newFields += ","
		}
		newFields += m.db.X取Core对象().X底层QuoteWord(k)
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
	conditionWhere, conditionArgs = m.whereBuilder.X生成条件字符串及参数()
	softDeletingCondition := m.softTimeMaintainer().GetWhereConditionForDelete(ctx)
	if m.rawSql != "" && conditionWhere != "" {
		if gstr.X是否包含并忽略大小写(m.rawSql, " WHERE ") {
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
			Args:   gconv.X取any切片(m.having[1]),
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
