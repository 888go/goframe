// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

import (
	"context"
	"fmt"
	"reflect"
	
	"github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/reflection"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
)

// All 方法执行针对模型的 "SELECT FROM ..." 语句。
// 它从表中检索记录并以切片类型返回结果。
// 如果根据给定条件未能从表中检索到任何记录，则返回 nil。
//
// 可选参数 `where` 与 Model.Where 函数的参数相同，
// 请参阅 Model.Where。
func (m *Model) X查询(查询条件 ...interface{}) (X行记录数组, error) {
	var ctx = m.X取上下文对象()
	return m.doGetAll(ctx, false, 查询条件...)
}

// AllAndCount 从模型中检索所有记录以及记录的总数。
// 如果 useFieldForCount 设为 true，它将使用模型中指定的字段进行计数；
// 否则，它将以常数值 1 进行计数。
// 它以记录切片的形式返回结果、记录的总数量以及（如有）错误信息。
// where 参数是一个可选的条件列表，在检索记录时使用这些条件。
//
// 示例：
//
//	var model Model   // 假设 Model 是自定义的数据模型
//	var result Result // 假设 Result 是存储查询结果的数据结构
//	var count int
//	where := []interface{}{"name = ?", "John"} // 设置查询条件：name 等于 "John"
//	result, count, err := model.AllAndCount(true) // 调用 AllAndCount 方法并传入参数 useFieldForCount 为 true
//	if err != nil {
//// 处理错误
//	}
//	fmt.Println(result, count) // 输出查询结果和记录总数
func (m *Model) X查询与行数(是否用字段计数 bool) (结果 X行记录数组, 行数 int, 错误 error) {
	// 对模型进行克隆以进行计数
	countModel := m.X取副本()

	// 如果useFieldForCount为false，则将字段设置为常数值1用于计数
	if !是否用字段计数 {
		countModel.fields = "1"
	}

	// 获取记录的总数
	行数, 错误 = countModel.X查询行数()
	if 错误 != nil {
		return
	}

	// 如果总记录数为0，则表示没有记录需要获取，所以提前返回
	if 行数 == 0 {
		return
	}

	// 获取所有记录
	结果, 错误 = m.doGetAll(m.X取上下文对象(), false)
	return
}

// Chunk 对查询结果进行迭代，指定 `size`（大小）和 `handler` 函数。
// 根据给定的 `size`，将查询结果分割成多个块，并对每个数据块应用 `handler` 函数进行处理。
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

// One 从表中检索一条记录并以map类型返回结果。
// 如果根据给定条件没有从表中检索到任何记录，则返回nil。
//
// 可选参数`where`与Model.Where函数的参数相同，
// 请参阅Model.Where。
func (m *Model) X查询一条(条件 ...interface{}) (X行记录, error) {
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

// 从数据库查询并以切片形式返回数据值。
// 注意，如果结果中有多个列，则随机返回其中一个列的值。
//
// 如果提供可选参数`fieldsAndWhere`，则fieldsAndWhere[0]表示选定的字段，
// 而fieldsAndWhere[1:]被视为where条件字段。
// 同时参阅Model.Fields和Model.Where函数。
func (m *Model) X查询数组(条件 ...interface{}) ([]X字段值, error) {
	if len(条件) > 0 {
		if len(条件) > 2 {
			return m.X字段保留过滤(转换类.String(条件[0])).X条件(条件[1], 条件[2:]...).X查询数组()
		} else if len(条件) == 2 {
			return m.X字段保留过滤(转换类.String(条件[0])).X条件(条件[1]).X查询数组()
		} else {
			return m.X字段保留过滤(转换类.String(条件[0])).X查询数组()
		}
	}
	all, err := m.X查询()
	if err != nil {
		return nil, err
	}
	var field string
	if len(all) > 0 {
		if internalData := m.db.X取Core对象().底层_GetInternalCtxDataFromCtx(m.X取上下文对象()); internalData != nil {
			field = internalData.FirstResultColumn
		} else {
			return nil, 错误类.X创建错误码(
				错误码类.CodeInternalError,
				`query array error: the internal context data is missing. there's internal issue should be fixed`,
			)
		}
	}
	return all.X取字段数组(field), nil
}

// Struct 从表中获取一条记录并将其转换为给定的结构体。
// 参数 `pointer` 应为 *struct 或 **struct 类型。如果提供的是 **struct 类型，
// 它会在转换过程中自动创建该结构体。
//
// 可选参数 `where` 与 Model.Where 函数的参数相同，详情请参阅 Model.Where。
//
// 注意，如果给定的 `pointer` 指向一个具有默认值的变量，并且根据给定条件未能从表中检索到记录，
// 则返回 sql.ErrNoRows 错误。
//
// 示例：
// user := new(User)
// err  := db.Model("user").Where("id", 1).Scan(user)
//
// user := (*User)(nil)
// err  := db.Model("user").Where("id", 1).Scan(&user)
func (m *Model) doStruct(pointer interface{}, where ...interface{}) error {
	model := m
	// 根据结构体属性自动选择字段
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
// 参数 `pointer` 应为 *[]struct 或 *[]*struct 类型，它在转换过程中可以内部创建并填充结构体切片。
//
// 可选参数 `where` 与 Model.Where 函数的参数相同，请参阅 Model.Where。
//
// 注意，如果给定的参数 `pointer` 指向一个具有默认值的变量，并且根据给定条件没有从表中检索到任何记录，则会返回 sql.ErrNoRows 错误。
//
// 示例：
// users := ([]User)(nil)
// err   := db.Model("user").Scan(&users)
//
// users := ([]*User)(nil)
// err   := db.Model("user").Scan(&users)
func (m *Model) doStructs(pointer interface{}, where ...interface{}) error {
	model := m
	// 根据结构体属性自动选择字段
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
	if err = all.X取数组结构体指针(pointer); err != nil {
		return err
	}
	return model.doWithScanStructs(pointer)
}

// Scan 根据参数 `pointer` 的类型自动调用 Struct 或 Structs 函数。
// 若 `pointer` 类型为 *struct/**struct，则调用 doStruct 函数。
// 若 `pointer` 类型为 *[]struct/*[]*struct，则调用 doStructs 函数。
// 可选参数 `where` 与 Model.Where 函数的参数相同，具体可参考 Model.Where。
// 注意：如果给定的参数 `pointer` 指向一个具有默认值的变量，并且根据给定条件未能从表中检索到任何记录时，
// 此函数将返回 sql.ErrNoRows 错误。
// 示例：
// 创建一个新的 User 实例
// user := new(User)
// 使用查询条件从 "user" 表中扫描数据并赋值给 user
// err  := db.Model("user").Where("id", 1).Scan(user)
// 创建一个空指针类型的 User 实例
// user := (*User)(nil)
// 使用查询条件从 "user" 表中扫描数据并赋值给 user
// err  := db.Model("user").Where("id", 1).Scan(&user)
// 创建一个空切片类型的 User 实例
// users := ([]User)(nil)
// 从 "user" 表中扫描数据并将其填充到 users 切片
// err   := db.Model("user").Scan(&users)
// 创建一个空指针切片类型的 User 实例
// users := ([]*User)(nil)
// 从 "user" 表中扫描数据并将其填充到 users 指针切片
// err   := db.Model("user").Scan(&users)
func (m *Model) X查询到结构体指针(数据指针 interface{}, 条件 ...interface{}) error {
	reflectInfo := reflection.OriginTypeAndKind(数据指针)
	if reflectInfo.InputKind != reflect.Ptr {
		return 错误类.X创建错误码(
			错误码类.CodeInvalidParameter,
			`the parameter "pointer" for function Scan should type of pointer`,
		)
	}
	switch reflectInfo.OriginKind {
	case reflect.Slice, reflect.Array:
		return m.doStructs(数据指针, 条件...)

	case reflect.Struct, reflect.Invalid:
		return m.doStruct(数据指针, 条件...)

	default:
		return 错误类.X创建错误码(
			错误码类.CodeInvalidParameter,
			`element of parameter "pointer" for function Scan should type of struct/*struct/[]struct/[]*struct`,
		)
	}
}

// ScanAndCount 扫描满足给定条件的单个记录或记录数组，并计算匹配这些条件的记录总数。
// 如果 useFieldForCount 为 true，则会使用模型中指定的字段进行计数；
// pointer 参数是一个指向结构体的指针，扫描的数据将存储在这个结构体中。
// pointerCount 参数是一个指向整数的指针，该整数将被设置为匹配给定条件的总记录数。
// where 参数是可选的检索记录时使用的条件列表。
// 示例：
//
//	var count int
//	user := new(User)
//	err  := db.Model("user").Where("id", 1).ScanAndCount(user, &count, true)
//	fmt.Println(user, count)
// 示例（联接查询）：
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
//		ScanAndCount(&users, &count, false) // 对于计数不考虑特定字段，因此设为 false
// 此函数用于根据提供的条件执行数据库查询，同时获取满足条件的记录数量。当 `useFieldForCount` 设为 `true` 时，计数会基于模型中定义的某些字段；否则，计数的是所有满足条件的记录条数。查询结果可以填充到传入的结构体实例（单例或切片）中，同时返回满足条件的记录总数。
func (m *Model) X查询与行数到指针(数据指针 interface{}, 行数指针 *int, 是否用字段计数 bool) (错误 error) {
	// 支持使用 * 通配符的 Fields，例如：.Fields("a.*, b.name")。Count SQL 为：从 xxx 中选择 count(1)
	countModel := m.X取副本()
	// 如果useFieldForCount为false，则将字段设置为常数值1用于计数
	if !是否用字段计数 {
		countModel.fields = "1"
	}

	// 获取记录的总数
	*行数指针, 错误 = countModel.X查询行数()
	if 错误 != nil {
		return 错误
	}

	// 如果总记录数为0，则表示没有记录需要获取，所以提前返回
	if *行数指针 == 0 {
		return
	}
	错误 = m.X查询到结构体指针(数据指针)
	return
}

// ScanList 将 `r` 转换为包含其他复杂结构体属性的结构体切片。
// 注意，参数 `listPointer` 应该是指向 []struct 或 []*struct 类型的指针。
// 参考关联模型: https://goframe.org/pages/viewpage.action?pageId=1114326
//
// 参见 Result.ScanList。
func (m *Model) X查询到指针列表(结构体切片指针 interface{}, 绑定到结构体属性名称 string, 结构体属性关联 ...string) (错误 error) {
	var result X行记录数组
	out, 错误 := checkGetSliceElementInfoForScanList(结构体切片指针, 绑定到结构体属性名称)
	if 错误 != nil {
		return 错误
	}
	if m.fields != defaultFields || m.fieldsEx != "" {
		// 存在自定义字段。
		result, 错误 = m.X查询()
	} else {
		// 使用reflect.New创建的临时结构体来过滤字段。
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

// Value 从表中检索指定记录的值并以 interface 类型返回结果。
// 如果根据给定条件在表中未找到记录，则返回 nil。
//
// 如果提供了可选参数 `fieldsAndWhere`，则 fieldsAndWhere[0] 表示选择的字段，
// 而 fieldsAndWhere[1:] 将被视为 where 条件字段。
// 请参阅 Model.Fields 和 Model.Where 函数。
func (m *Model) X查询一条值(字段和条件 ...interface{}) (X字段值, error) {
	var ctx = m.X取上下文对象()
	if len(字段和条件) > 0 {
		if len(字段和条件) > 2 {
			return m.X字段保留过滤(转换类.String(字段和条件[0])).X条件(字段和条件[1], 字段和条件[2:]...).X查询一条值()
		} else if len(字段和条件) == 2 {
			return m.X字段保留过滤(转换类.String(字段和条件[0])).X条件(字段和条件[1]).X查询一条值()
		} else {
			return m.X字段保留过滤(转换类.String(字段和条件[0])).X查询一条值()
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
		if internalData := m.db.X取Core对象().底层_GetInternalCtxDataFromCtx(ctx); internalData != nil {
			if v, ok := all[0][internalData.FirstResultColumn]; ok {
				return v, nil
			}
		} else {
			return nil, 错误类.X创建错误码(
				错误码类.CodeInternalError,
				`query value error: the internal context data is missing. there's internal issue should be fixed`,
			)
		}
	}
	return nil, nil
}

// Count 对模型执行 "SELECT COUNT(x) FROM ..." 语句。
// 可选参数 `where` 与 Model.Where 函数的参数相同，
// 请参阅 Model.Where。
func (m *Model) X查询行数(条件 ...interface{}) (int, error) {
	var ctx = m.X取上下文对象()
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
		if internalData := m.db.X取Core对象().底层_GetInternalCtxDataFromCtx(ctx); internalData != nil {
			if v, ok := all[0][internalData.FirstResultColumn]; ok {
				return v.X取整数(), nil
			}
		} else {
			return 0, 错误类.X创建错误码(
				错误码类.CodeInternalError,
				`query count error: the internal context data is missing. there's internal issue should be fixed`,
			)
		}
	}
	return 0, nil
}

// CountColumn 对模型执行 "SELECT COUNT(x) FROM ..." 语句。
func (m *Model) X查询字段行数(字段名称 string) (int, error) {
	if len(字段名称) == 0 {
		return 0, nil
	}
	return m.X字段保留过滤(字段名称).X查询行数()
}

// Min 为该模型执行“SELECT MIN(x) FROM ...”语句。
func (m *Model) X查询最小值(字段名称 string) (float64, error) {
	if len(字段名称) == 0 {
		return 0, nil
	}
	value, err := m.X字段保留过滤(fmt.Sprintf(`MIN(%s)`, m.底层QuoteWord(字段名称))).X查询一条值()
	if err != nil {
		return 0, err
	}
	return value.X取小数64位(), err
}

// Max 为给定的模型执行“SELECT MAX(x) FROM ...”语句。
func (m *Model) X查询最大值(字段名称 string) (float64, error) {
	if len(字段名称) == 0 {
		return 0, nil
	}
	value, err := m.X字段保留过滤(fmt.Sprintf(`MAX(%s)`, m.底层QuoteWord(字段名称))).X查询一条值()
	if err != nil {
		return 0, err
	}
	return value.X取小数64位(), err
}

// Avg 对模型执行 "SELECT AVG(x) FROM ..." 语句，计算平均值。
func (m *Model) X查询平均值(字段名称 string) (float64, error) {
	if len(字段名称) == 0 {
		return 0, nil
	}
	value, err := m.X字段保留过滤(fmt.Sprintf(`AVG(%s)`, m.底层QuoteWord(字段名称))).X查询一条值()
	if err != nil {
		return 0, err
	}
	return value.X取小数64位(), err
}

// Sum 对模型执行 "SELECT SUM(x) FROM ..." 语句，计算求和。
func (m *Model) X查询求和(字段名称 string) (float64, error) {
	if len(字段名称) == 0 {
		return 0, nil
	}
	value, err := m.X字段保留过滤(fmt.Sprintf(`SUM(%s)`, m.底层QuoteWord(字段名称))).X查询一条值()
	if err != nil {
		return 0, err
	}
	return value.X取小数64位(), err
}

// Union 为给定的模型执行 "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." 类似的SQL语句查询。
func (m *Model) X多表去重查询(Model对象 ...*Model) *Model {
	return m.db.X多表去重查询(Model对象...)
}

// UnionAll 对模型执行“(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ...”语句。
func (m *Model) X多表查询(Model对象 ...*Model) *Model {
	return m.db.X多表查询(Model对象...)
}

// Limit 为模型设置 "LIMIT" 语句。
// 参数 `limit` 可以是一个或两个数字，如果传入两个数字，
// 则为模型设置 "LIMIT limit[0], limit[1]" 语句，否则设置 "LIMIT limit[0]" 语句。
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

// Offset 设置模型的 "OFFSET" 语句。
// 它只对某些数据库有意义，如 SQLServer、PostgreSQL 等。
func (m *Model) Offset(offset int) *Model {
	model := m.getModel()
	model.offset = offset
	return model
}

// Distinct 用于强制查询只返回不重复的结果。
func (m *Model) X设置去重() *Model {
	model := m.getModel()
	model.distinct = "DISTINCT "
	return model
}

// Page 为模型设置分页号。
// 参数 `page` 的分页从1开始计数。
// 注意，它与 Limit 函数从0开始为 "LIMIT" 语句设置偏移量有所不同。
func (m *Model) X设置分页(第几页, 条数 int) *Model {
	model := m.getModel()
	if 第几页 <= 0 {
		第几页 = 1
	}
	model.start = (第几页 - 1) * 条数
	model.limit = 条数
	return model
}

// Having 设置模型的 having 子句。
// 该函数的使用参数与 Where 函数相同。
// 参见 Where。
func (m *Model) X设置分组条件(条件 interface{}, 参数 ...interface{}) *Model {
	model := m.getModel()
	model.having = []interface{}{
		条件, 参数,
	}
	return model
}

// doGetAll 执行针对模型的 "SELECT FROM ..." 语句。
// 它从表中检索记录并以切片类型返回结果。
// 如果根据给定条件没有从表中检索到任何记录，则返回 nil。
//
// 参数 `limit1` 指定在 m.limit 未设置时是否限制仅查询一条记录。
// 可选参数 `where` 与 Model.Where 函数的参数相同，
// 请参阅 Model.Where。
func (m *Model) doGetAll(ctx context.Context, limit1 bool, where ...interface{}) (X行记录数组, error) {
	if len(where) > 0 {
		return m.X条件(where[0], where[1:]...).X查询()
	}
	sqlWithHolder, holderArgs := m.getFormattedSqlAndArgs(ctx, queryTypeNormal, limit1)
	return m.doGetAllBySql(ctx, queryTypeNormal, sqlWithHolder, holderArgs...)
}

// doGetAllBySql 对数据库执行 select 语句。
func (m *Model) doGetAllBySql(ctx context.Context, queryType queryType, sql string, args ...interface{}) (result X行记录数组, err error) {
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
// **注意**：在此处不要引用m.fields，以防字段类似以下情况：
// DISTINCT t.user_id uid
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
		// 原生SQL模型，特别适用于包含UNION/UNION ALL特性的SQL语句。
		if m.rawSql != "" {
			sqlWithHolder = fmt.Sprintf(
				"%s%s",
				m.rawSql,
				conditionWhere+conditionExtra,
			)
			return sqlWithHolder, conditionArgs
		}
// **不要**对m.fields进行引用，特别是在处理类似以下字段时：
// DISTINCT t.user_id AS uid
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
	if 文本类.X是否包含(m.tables, " JOIN ") {
		autoPrefix = m.db.X取Core对象().X底层QuoteWord(
			m.db.X取Core对象().guessPrimaryTableName(m.tablesInit),
		)
	}
	return autoPrefix
}

// getFieldsFiltered 检查 fields 和 fieldsEx 属性，进行过滤并返回真正将提交到底层数据库驱动的字段。
func (m *Model) getFieldsFiltered() string {
	if m.fieldsEx == "" {
		// 不进行过滤，包含特殊字符。
		if 文本类.X是否包含Any(m.fields, "()") {
			return m.fields
		}
		// No filtering.
		if !文本类.X是否包含Any(m.fields, ". ") {
			return m.db.X取Core对象().X底层QuoteString(m.fields)
		}
		return m.fields
	}
	var (
		fieldsArray []string
		fieldsExSet = 集合类.X创建文本并按值(文本类.X分割并忽略空值(m.fieldsEx, ","))
	)
	if m.fields != "*" {
		// 使用fieldEx过滤自定义字段。
		fieldsArray = make([]string, 0, 8)
		for _, v := range 文本类.X分割并忽略空值(m.fields, ",") {
			fieldsArray = append(fieldsArray, v[文本类.X倒找(v, "-")+1:])
		}
	} else {
		if 文本类.X是否包含(m.tables, " ") {
			panic("function FieldsEx supports only single table operations")
		}
		// 使用fieldEx过滤表字段。
		tableFields, err := m.X取表字段信息Map(m.tablesInit)
		if err != nil {
			panic(err)
		}
		if len(tableFields) == 0 {
			panic(fmt.Sprintf(`empty table fields for table "%s"`, m.tables))
		}
		fieldsArray = make([]string, len(tableFields))
		for k, v := range tableFields {
			fieldsArray[v.X排序] = k
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

// formatCondition 格式化模型的 where 条件参数，并返回一个新的条件 SQL 语句及其参数。
// 注意，此函数不会改变 `m` 的任何属性值。
//
// 参数 `limit1` 指定在 `m.limit` 未设置时是否限制查询仅一条记录。
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
	softDeletingCondition := m.getConditionForSoftDeleting()
	if m.rawSql != "" && conditionWhere != "" {
		if 文本类.X是否包含并忽略大小写(m.rawSql, " WHERE ") {
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
			Args:   转换类.X取any数组(m.having[1]),
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
	if !isCountStatement { // SQL Server 中的 count 语句不能包含 order by 子句
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
