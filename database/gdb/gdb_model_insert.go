// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

import (
	"context"
	"database/sql"
	"reflect"

	gset "github.com/888go/goframe/container/gset"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/reflection"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
	gutil "github.com/888go/goframe/util/gutil"
)

// X设置批量操作行数 为模型设置批处理操作的数量。 md5:7ae8528d1f8ac604
func (m *Model) X设置批量操作行数(数量 int) *Model {
	model := m.getModel()
	model.batch = 数量
	return model
}

// X设置数据 设置模型的操作数据。
// 参数 `data` 可以为字符串/映射/gmap/切片/结构体/**结构体指针**等类型。
// 注意，如果`data`是映射或切片类型，它将使用浅复制以避免在函数内部改变原数据。
// 例如：
// X设置数据("uid=10000")
// X设置数据("uid", 10000)
// X设置数据("uid=? AND name=?", 10000, "john")
// X设置数据(g.Map{"uid": 10000, "name":"john"})
// X设置数据(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"}})
// md5:116cf94880dfa535
func (m *Model) X设置数据(值 ...interface{}) *Model {
	var model = m.getModel()
	if len(值) > 1 {
		if s := gconv.String(值[0]); gstr.X是否包含(s, "?") {
			model.data = s
			model.extraArgs = 值[1:]
		} else {
			m := make(map[string]interface{})
			for i := 0; i < len(值); i += 2 {
				m[gconv.String(值[i])] = 值[i+1]
			}
			model.data = m
		}
	} else if len(值) == 1 {
		switch value := 值[0].(type) {
		case Result:
			model.data = value.X取Map切片()

		case Record:
			model.data = value.X取Map()

		case Map切片:
			list := make(Map切片, len(value))
			for k, v := range value {
				list[k] = gutil.MapCopy(v)
			}
			model.data = list

		case Map:
			model.data = gutil.MapCopy(value)

		default:
			reflectInfo := reflection.OriginValueAndKind(value)
			switch reflectInfo.OriginKind {
			case reflect.Slice, reflect.Array:
				if reflectInfo.OriginValue.Len() > 0 {
					// 如果`data`参数是一个DO结构体，
					// 则为这个条件添加`OmitNilData`选项，
					// 这将过滤掉`data`中的所有空值参数。
					// md5:c978d65b6ea1129a
					if isDoStruct(reflectInfo.OriginValue.Index(0).Interface()) {
						model = model.X过滤Nil数据()
						model.option |= optionOmitNilDataInternal
					}
				}
				list := make(Map切片, reflectInfo.OriginValue.Len())
				for i := 0; i < reflectInfo.OriginValue.Len(); i++ {
					list[i] = anyValueToMapBeforeToRecord(reflectInfo.OriginValue.Index(i).Interface())
				}
				model.data = list

			case reflect.Struct:
				// If the `data` parameter is a DO struct,
				// it then adds `OmitNilData` option for this condition,
				// which will filter all nil parameters in `data`.
				if isDoStruct(value) {
					model = model.X过滤Nil数据()
				}
				if v, ok := 值[0].(iInterfaces); ok {
					var (
						array = v.X取any切片()
						list  = make(Map切片, len(array))
					)
					for i := 0; i < len(array); i++ {
						list[i] = anyValueToMapBeforeToRecord(array[i])
					}
					model.data = list
				} else {
					model.data = anyValueToMapBeforeToRecord(值[0])
				}

			case reflect.Map:
				model.data = anyValueToMapBeforeToRecord(值[0])

			default:
				model.data = 值[0]
			}
		}
	}
	return model
}

// OnConflict在列冲突时设置主键或索引。对于MySQL驱动程序来说，这通常是不必要的。
// md5:30314cb75360b0e6
func (m *Model) OnConflict(onConflict ...interface{}) *Model {
	if len(onConflict) == 0 {
		return m
	}
	model := m.getModel()
	if len(onConflict) > 1 {
		model.onConflict = onConflict
	} else if len(onConflict) == 1 {
		model.onConflict = onConflict[0]
	}
	return model
}

// X设置插入冲突更新字段 设置在列发生冲突时执行的操作。
// 在MySQL中，这用于 "ON DUPLICATE KEY UPDATE" 语句。
// 在PgSQL中，这用于 "ON CONFLICT (id) DO UPDATE SET" 语句。
// 参数 `onDuplicate` 可以是字符串/Raw/*Raw/映射/切片类型。
// 示例：
//
// X设置插入冲突更新字段("nickname, age")
// X设置插入冲突更新字段("nickname", "age")
// 
// X设置插入冲突更新字段(g.Map{
// 	  "nickname": gdb.Raw("CONCAT('name_', VALUES(`nickname`))"),
// })
//
// X设置插入冲突更新字段(g.Map{
// 	  "nickname": "passport",
// })
// md5:fa9214f9681b4e5d
func (m *Model) X设置插入冲突更新字段(字段名称 ...interface{}) *Model {
	if len(字段名称) == 0 {
		return m
	}
	model := m.getModel()
	if len(字段名称) > 1 {
		model.onDuplicate = 字段名称
	} else if len(字段名称) == 1 {
		model.onDuplicate = 字段名称[0]
	}
	return model
}

// X设置插入冲突不更新字段 设置在发生列冲突时排除的列，用于操作。
// 在 MySQL 中，这用于 "ON DUPLICATE KEY UPDATE" 语句。
// 在 PgSQL 中，这用于 "ON CONFLICT (id) DO UPDATE SET" 语句。
// 参数 `onDuplicateEx` 可以是字符串、映射或切片类型。
// 示例：
//
// X设置插入冲突不更新字段("passport, password")
// X设置插入冲突不更新字段("passport", "password")
//
//	X设置插入冲突不更新字段(g.Map{
//		  "passport": "",
//		  "password": "",
//	})
// md5:6fa8981bef042b71
func (m *Model) X设置插入冲突不更新字段(字段名称 ...interface{}) *Model {
	if len(字段名称) == 0 {
		return m
	}
	model := m.getModel()
	if len(字段名称) > 1 {
		model.onDuplicateEx = 字段名称
	} else if len(字段名称) == 1 {
		model.onDuplicateEx = 字段名称[0]
	}
	return model
}

// X插入 为模型执行 "INSERT INTO ..." 语句。
// 可选参数 `data` 等同于 Model.Data 函数的参数，参见 Model.Data。
// md5:9a6427cabf3ec194
func (m *Model) X插入(值 ...interface{}) (结果 sql.Result, 错误 error) {
	var ctx = m.X取上下文对象()
	if len(值) > 0 {
		return m.X设置数据(值...).X插入()
	}
	return m.doInsertWithOption(ctx, InsertOptionDefault)
}

// X插入并取ID 执行插入操作，并返回自动生成的最后一个插入id。 md5:8d00b40a35fa48a5
func (m *Model) X插入并取ID(值 ...interface{}) (最后插入ID int64, 错误 error) {
	var ctx = m.X取上下文对象()
	if len(值) > 0 {
		return m.X设置数据(值...).X插入并取ID()
	}
	result, 错误 := m.doInsertWithOption(ctx, InsertOptionDefault)
	if 错误 != nil {
		return 0, 错误
	}
	return result.LastInsertId()
}

// X插入并跳过已存在 为模型执行 "INSERT IGNORE INTO..." 语句。
// 可选参数 `data` 和 Model.Data 函数的参数相同，详情请参考 Model.Data。
// md5:d6d8007d779bd324
func (m *Model) X插入并跳过已存在(值 ...interface{}) (结果 sql.Result, 错误 error) {
	var ctx = m.X取上下文对象()
	if len(值) > 0 {
		return m.X设置数据(值...).X插入并跳过已存在()
	}
	return m.doInsertWithOption(ctx, InsertOptionIgnore)
}

// X插入并替换已存在 执行 "REPLACE INTO ..." 语句用于模型。
// 可选参数 `data` 与 Model.Data 函数的参数相同，
// 请参阅 Model.Data。
// md5:d5596c2470b6bcf4
func (m *Model) X插入并替换已存在(值 ...interface{}) (结果 sql.Result, 错误 error) {
	var ctx = m.X取上下文对象()
	if len(值) > 0 {
		return m.X设置数据(值...).X插入并替换已存在()
	}
	return m.doInsertWithOption(ctx, InsertOptionReplace)
}

// X插入并更新已存在 执行 "INSERT INTO ... ON DUPLICATE KEY UPDATE..." 语句，针对指定的模型。
// 可选参数 `data` 与 Model.Data 函数的参数相同，请参阅 Model.Data。
//
// 如果保存的数据中包含主键或唯一索引，它将更新记录；
// 否则，它会向表中插入一条新记录。
// md5:9d87bd779f8f5acd
func (m *Model) X插入并更新已存在(值 ...interface{}) (结果 sql.Result, 错误 error) {
	var ctx = m.X取上下文对象()
	if len(值) > 0 {
		return m.X设置数据(值...).X插入并更新已存在()
	}
	return m.doInsertWithOption(ctx, InsertOptionSave)
}

// doInsertWithOption 使用option参数插入数据。 md5:49dfb820e896850a
func (m *Model) doInsertWithOption(ctx context.Context, insertOption InsertOption) (result sql.Result, err error) {
	defer func() {
		if err == nil {
			m.checkAndRemoveSelectCache(ctx)
		}
	}()
	if m.data == nil {
		return nil, gerror.X创建错误码(gcode.CodeMissingParameter, "inserting into table with empty data")
	}
	var (
		list                             Map切片
		stm                              = m.softTimeMaintainer()
		fieldNameCreate, fieldTypeCreate = stm.GetFieldNameAndTypeForCreate(ctx, "", m.tablesInit)
		fieldNameUpdate, fieldTypeUpdate = stm.GetFieldNameAndTypeForUpdate(ctx, "", m.tablesInit)
		fieldNameDelete, fieldTypeDelete = stm.GetFieldNameAndTypeForDelete(ctx, "", m.tablesInit)
	)
			// m.data 已经通过 Data 函数转换为了 List/Map 类型. md5:cce9527c9f06deb0
	newData, err := m.filterDataForInsertOrUpdate(m.data)
	if err != nil {
		return nil, err
	}
		// 它将任何数据转换为List类型以便插入。 md5:8e4e33863c8e1d24
	switch value := newData.(type) {
	case Map切片:
		list = value

	case Map:
		list = Map切片{value}
	}

	if len(list) < 1 {
		return result, gerror.X创建错误码(gcode.CodeMissingParameter, "data list cannot be empty")
	}

		// 自动处理创建/更新时间。 md5:c45a07308954de68
	if !m.unscoped && (fieldNameCreate != "" || fieldNameUpdate != "") {
		for k, v := range list {
			if fieldNameCreate != "" {
				fieldCreateValue := stm.GetValueByFieldTypeForCreateOrUpdate(ctx, fieldTypeCreate, false)
				if fieldCreateValue != nil {
					v[fieldNameCreate] = fieldCreateValue
				}
			}
			if fieldNameUpdate != "" {
				fieldUpdateValue := stm.GetValueByFieldTypeForCreateOrUpdate(ctx, fieldTypeUpdate, false)
				if fieldUpdateValue != nil {
					v[fieldNameUpdate] = fieldUpdateValue
				}
			}
			if fieldNameDelete != "" {
				fieldDeleteValue := stm.GetValueByFieldTypeForCreateOrUpdate(ctx, fieldTypeDelete, true)
				if fieldDeleteValue != nil {
					v[fieldNameDelete] = fieldDeleteValue
				}
			}
			list[k] = v
		}
	}
		// 格式化DoInsertOption，特别是针对“ON DUPLICATE KEY UPDATE”语句。 md5:e668e4c647415360
	columnNames := make([]string, 0, len(list[0]))
	for k := range list[0] {
		columnNames = append(columnNames, k)
	}
	doInsertOption, err := m.formatDoInsertOption(insertOption, columnNames)
	if err != nil {
		return result, err
	}

	in := &HookInsertInput{
		internalParamHookInsert: internalParamHookInsert{
			internalParamHook: internalParamHook{
				link: m.getLink(true),
			},
			handler: m.hookHandler.Insert,
		},
		Model:  m,
		Table:  m.tables,
		Data:   list,
		Option: doInsertOption,
	}
	return in.Next(ctx)
}

func (m *Model) formatDoInsertOption(insertOption InsertOption, columnNames []string) (option DoInsertOption, err error) {
	option = DoInsertOption{
		InsertOption: insertOption,
		BatchCount:   m.getBatch(),
	}
	if insertOption != InsertOptionSave {
		return
	}

	onConflictKeys, err := m.formatOnConflictKeys(m.onConflict)
	if err != nil {
		return option, err
	}
	option.OnConflict = onConflictKeys

	onDuplicateExKeys, err := m.formatOnDuplicateExKeys(m.onDuplicateEx)
	if err != nil {
		return option, err
	}
	onDuplicateExKeySet := gset.X创建文本并按值(onDuplicateExKeys)
	if m.onDuplicate != nil {
		switch m.onDuplicate.(type) {
		case Raw, *Raw:
			option.OnDuplicateStr = gconv.String(m.onDuplicate)

		default:
			reflectInfo := reflection.OriginValueAndKind(m.onDuplicate)
			switch reflectInfo.OriginKind {
			case reflect.String:
				option.OnDuplicateMap = make(map[string]interface{})
				for _, v := range gstr.X分割并忽略空值(reflectInfo.OriginValue.String(), ",") {
					if onDuplicateExKeySet.X是否存在(v) {
						continue
					}
					option.OnDuplicateMap[v] = v
				}

			case reflect.Map:
				option.OnDuplicateMap = make(map[string]interface{})
				for k, v := range gconv.X取Map(m.onDuplicate) {
					if onDuplicateExKeySet.X是否存在(k) {
						continue
					}
					option.OnDuplicateMap[k] = v
				}

			case reflect.Slice, reflect.Array:
				option.OnDuplicateMap = make(map[string]interface{})
				for _, v := range gconv.X取文本切片(m.onDuplicate) {
					if onDuplicateExKeySet.X是否存在(v) {
						continue
					}
					option.OnDuplicateMap[v] = v
				}

			default:
				return option, gerror.X创建错误码并格式化(
					gcode.CodeInvalidParameter,
					`unsupported OnDuplicate parameter type "%s"`,
					reflect.TypeOf(m.onDuplicate),
				)
			}
		}
	} else if onDuplicateExKeySet.X取数量() > 0 {
		option.OnDuplicateMap = make(map[string]interface{})
		for _, v := range columnNames {
			if onDuplicateExKeySet.X是否存在(v) {
				continue
			}
			option.OnDuplicateMap[v] = v
		}
	}
	return
}

func (m *Model) formatOnDuplicateExKeys(onDuplicateEx interface{}) ([]string, error) {
	if onDuplicateEx == nil {
		return nil, nil
	}

	reflectInfo := reflection.OriginValueAndKind(onDuplicateEx)
	switch reflectInfo.OriginKind {
	case reflect.String:
		return gstr.X分割并忽略空值(reflectInfo.OriginValue.String(), ","), nil

	case reflect.Map:
		return gutil.X取所有名称(onDuplicateEx), nil

	case reflect.Slice, reflect.Array:
		return gconv.X取文本切片(onDuplicateEx), nil

	default:
		return nil, gerror.X创建错误码并格式化(
			gcode.CodeInvalidParameter,
			`unsupported OnDuplicateEx parameter type "%s"`,
			reflect.TypeOf(onDuplicateEx),
		)
	}
}

func (m *Model) formatOnConflictKeys(onConflict interface{}) ([]string, error) {
	if onConflict == nil {
		return nil, nil
	}

	reflectInfo := reflection.OriginValueAndKind(onConflict)
	switch reflectInfo.OriginKind {
	case reflect.String:
		return gstr.X分割并忽略空值(reflectInfo.OriginValue.String(), ","), nil

	case reflect.Slice, reflect.Array:
		return gconv.X取文本切片(onConflict), nil

	default:
		return nil, gerror.X创建错误码并格式化(
			gcode.CodeInvalidParameter,
			`unsupported onConflict parameter type "%s"`,
			reflect.TypeOf(onConflict),
		)
	}
}

func (m *Model) getBatch() int {
	return m.batch
}
