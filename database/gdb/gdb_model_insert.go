// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

import (
	"context"
	"database/sql"
	"reflect"
	
	"github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/reflection"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
)

// Batch 设置模型的批量操作数量。
func (m *Model) X设置批量操作行数(数量 int) *Model {
	model := m.getModel()
	model.batch = 数量
	return model
}

// Data 设置模型操作的数据。
// 参数 `data` 可以为 string、map、gmap、slice、struct、*struct 等类型。
// 注意，如果 `data` 为 map 或 slice 类型，会采用浅值复制的方式来避免在函数内部对原数据进行修改。
// 示例：
// Data("uid=10000")
// Data("uid", 10000)
// Data("uid=? AND name=?", 10000, "john")
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"}})
func (m *Model) X设置数据(值 ...interface{}) *Model {
	var model = m.getModel()
	if len(值) > 1 {
		if s := 转换类.String(值[0]); 文本类.X是否包含(s, "?") {
			model.data = s
			model.extraArgs = 值[1:]
		} else {
			m := make(map[string]interface{})
			for i := 0; i < len(值); i += 2 {
				m[转换类.String(值[i])] = 值[i+1]
			}
			model.data = m
		}
	} else if len(值) == 1 {
		switch value := 值[0].(type) {
		case Result:
			model.data = value.X取Map数组()

		case Record:
			model.data = value.X取Map()

		case Map数组:
			list := make(Map数组, len(value))
			for k, v := range value {
				list[k] = 工具类.MapCopy(v)
			}
			model.data = list

		case Map:
			model.data = 工具类.MapCopy(value)

		default:
			reflectInfo := reflection.OriginValueAndKind(value)
			switch reflectInfo.OriginKind {
			case reflect.Slice, reflect.Array:
				if reflectInfo.OriginValue.Len() > 0 {
// 如果`data`参数是一个DO结构体，
// 则为此条件添加`OmitNilData`选项，
// 这将会过滤掉`data`中所有为nil的参数。
					if isDoStruct(reflectInfo.OriginValue.Index(0).Interface()) {
						model = model.X过滤Nil数据()
						model.option |= optionOmitNilDataInternal
					}
				}
				list := make(Map数组, reflectInfo.OriginValue.Len())
				for i := 0; i < reflectInfo.OriginValue.Len(); i++ {
					list[i] = anyValueToMapBeforeToRecord(reflectInfo.OriginValue.Index(i).Interface())
				}
				model.data = list

			case reflect.Struct:
// 如果`data`参数是一个DO结构体，
// 则为此条件添加`OmitNilData`选项，
// 这将会过滤掉`data`中所有的nil参数。
				if isDoStruct(value) {
					model = model.X过滤Nil数据()
				}
				if v, ok := 值[0].(iInterfaces); ok {
					var (
						array = v.X取any数组()
						list  = make(Map数组, len(array))
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

// OnDuplicate 设置在列冲突发生时的操作。
// 在 MySQL 中，此方法用于“ON DUPLICATE KEY UPDATE”语句。
// 参数 `onDuplicate` 可以为 string/Raw/*Raw/map/slice 类型。
// 示例：
//
// OnDuplicate("nickname, age") // 设置当主键重复时更新nickname和age字段
// OnDuplicate("nickname", "age") // 同上，以逗号分隔多个字段名
//
//	OnDuplicate(g.Map{
//		  "nickname": gdb.Raw("CONCAT('name_', VALUES(`nickname`))"), // 使用原始SQL表达式更新nickname字段
//	})
//
//	OnDuplicate(g.Map{
//		  "nickname": "passport", // 当主键重复时，将nickname字段的值设置为passport字段的值
//	}).
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

// OnDuplicateEx 设置在列冲突发生时操作的排除列。
// 在 MySQL 中，此函数用于 "ON DUPLICATE KEY UPDATE" 语句。
// 参数 `onDuplicateEx` 可以是字符串、映射或切片类型。
// 示例：
//
// OnDuplicateEx("passport, password") // 传入一个包含列名的字符串
// OnDuplicateEx("passport", "password") // 分别指定列名参数
//
//	OnDuplicateEx(g.Map{
//		  "passport": "",
//		  "password": "",
//	}) // 通过映射传入选定列名和其对应的更新值（此处为空字符串）
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

// Insert 执行针对模型的 "INSERT INTO ..." 语句。
// 可选参数 `data` 与 Model.Data 函数的参数相同，
// 请参考 Model.Data。
func (m *Model) X插入(值 ...interface{}) (结果 sql.Result, 错误 error) {
	var ctx = m.X取上下文对象()
	if len(值) > 0 {
		return m.X设置数据(值...).X插入()
	}
	return m.doInsertWithOption(ctx, InsertOptionDefault)
}

// InsertAndGetId 执行插入操作，并返回自动生成的最后一个插入ID。
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

// InsertIgnore 执行针对模型的 "INSERT IGNORE INTO ..." 语句。
// 可选参数 `data` 与 Model.Data 函数的参数相同，
// 请参阅 Model.Data。
func (m *Model) X插入并跳过已存在(值 ...interface{}) (结果 sql.Result, 错误 error) {
	var ctx = m.X取上下文对象()
	if len(值) > 0 {
		return m.X设置数据(值...).X插入并跳过已存在()
	}
	return m.doInsertWithOption(ctx, InsertOptionIgnore)
}

// Replace 执行针对模型的 "REPLACE INTO ..." 语句。
// 可选参数 `data` 与 Model.Data 函数的参数相同，
// 详情请参阅 Model.Data。
func (m *Model) X插入并替换已存在(值 ...interface{}) (结果 sql.Result, 错误 error) {
	var ctx = m.X取上下文对象()
	if len(值) > 0 {
		return m.X设置数据(值...).X插入并替换已存在()
	}
	return m.doInsertWithOption(ctx, InsertOptionReplace)
}

// Save 执行 "INSERT INTO ... ON DUPLICATE KEY UPDATE..." 语句，针对给定的 model。
// 可选参数 `data` 与 Model.Data 函数的参数相同，
// 请参阅 Model.Data。
//
// 如果保存的数据中存在主键或唯一索引，则更新记录，
// 否则会在表中插入一条新记录。
func (m *Model) X插入并更新已存在(值 ...interface{}) (结果 sql.Result, 错误 error) {
	var ctx = m.X取上下文对象()
	if len(值) > 0 {
		return m.X设置数据(值...).X插入并更新已存在()
	}
	return m.doInsertWithOption(ctx, InsertOptionSave)
}

// doInsertWithOption 使用选项参数插入数据。
func (m *Model) doInsertWithOption(ctx context.Context, insertOption InsertOption) (result sql.Result, err error) {
	defer func() {
		if err == nil {
			m.checkAndRemoveSelectCache(ctx)
		}
	}()
	if m.data == nil {
		return nil, 错误类.X创建错误码(错误码类.CodeMissingParameter, "inserting into table with empty data")
	}
	var (
		list            Map数组
		now             = 时间类.X创建并按当前时间()
		fieldNameCreate = m.getSoftFieldNameCreated("", m.tablesInit)
		fieldNameUpdate = m.getSoftFieldNameUpdated("", m.tablesInit)
	)
	newData, err := m.filterDataForInsertOrUpdate(m.data)
	if err != nil {
		return nil, err
	}
	// 它将任何数据转换为 List 类型以便插入。
	switch value := newData.(type) {
	case Result:
		list = value.X取Map数组()

	case Record:
		list = Map数组{value.X取Map()}

	case Map数组:
		list = value

	case Map:
		list = Map数组{value}

	default:
// 这里使用gconv.Map简化从interface{}到map[string]interface{}的类型转换，
// 因为接下来的逻辑中会使用MapOrStructToMapDeep进行深度转换。
		reflectInfo := reflection.OriginValueAndKind(newData)
		switch reflectInfo.OriginKind {
		// 如果它是切片类型，那么将其转换为 List 类型。
		case reflect.Slice, reflect.Array:
			list = make(Map数组, reflectInfo.OriginValue.Len())
			for i := 0; i < reflectInfo.OriginValue.Len(); i++ {
				list[i] = anyValueToMapBeforeToRecord(reflectInfo.OriginValue.Index(i).Interface())
			}

		case reflect.Map:
			list = Map数组{anyValueToMapBeforeToRecord(value)}

		case reflect.Struct:
			if v, ok := value.(iInterfaces); ok {
				array := v.X取any数组()
				list = make(Map数组, len(array))
				for i := 0; i < len(array); i++ {
					list[i] = anyValueToMapBeforeToRecord(array[i])
				}
			} else {
				list = Map数组{anyValueToMapBeforeToRecord(value)}
			}

		default:
			return result, 错误类.X创建错误码并格式化(
				错误码类.CodeInvalidParameter,
				"unsupported data list type: %v",
				reflectInfo.InputValue.Type(),
			)
		}
	}

	if len(list) < 1 {
		return result, 错误类.X创建错误码(错误码类.CodeMissingParameter, "data list cannot be empty")
	}

	// 自动处理创建/更新时间。
	if !m.unscoped && (fieldNameCreate != "" || fieldNameUpdate != "") {
		for k, v := range list {
			if fieldNameCreate != "" {
				v[fieldNameCreate] = now
			}
			if fieldNameUpdate != "" {
				v[fieldNameUpdate] = now
			}
			list[k] = v
		}
	}
	// 格式化DoInsertOption，特别是用于“ON DUPLICATE KEY UPDATE”语句。
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
	if insertOption == InsertOptionSave {
		onDuplicateExKeys, err := m.formatOnDuplicateExKeys(m.onDuplicateEx)
		if err != nil {
			return option, err
		}
		onDuplicateExKeySet := 集合类.X创建文本并按值(onDuplicateExKeys)
		if m.onDuplicate != nil {
			switch m.onDuplicate.(type) {
			case Raw, *Raw:
				option.OnDuplicateStr = 转换类.String(m.onDuplicate)

			default:
				reflectInfo := reflection.OriginValueAndKind(m.onDuplicate)
				switch reflectInfo.OriginKind {
				case reflect.String:
					option.OnDuplicateMap = make(map[string]interface{})
					for _, v := range 文本类.X分割并忽略空值(reflectInfo.OriginValue.String(), ",") {
						if onDuplicateExKeySet.X是否存在(v) {
							continue
						}
						option.OnDuplicateMap[v] = v
					}

				case reflect.Map:
					option.OnDuplicateMap = make(map[string]interface{})
					for k, v := range 转换类.X取Map(m.onDuplicate) {
						if onDuplicateExKeySet.X是否存在(k) {
							continue
						}
						option.OnDuplicateMap[k] = v
					}

				case reflect.Slice, reflect.Array:
					option.OnDuplicateMap = make(map[string]interface{})
					for _, v := range 转换类.X取文本数组(m.onDuplicate) {
						if onDuplicateExKeySet.X是否存在(v) {
							continue
						}
						option.OnDuplicateMap[v] = v
					}

				default:
					return option, 错误类.X创建错误码并格式化(
						错误码类.CodeInvalidParameter,
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
		return 文本类.X分割并忽略空值(reflectInfo.OriginValue.String(), ","), nil

	case reflect.Map:
		return 工具类.X取所有名称(onDuplicateEx), nil

	case reflect.Slice, reflect.Array:
		return 转换类.X取文本数组(onDuplicateEx), nil

	default:
		return nil, 错误类.X创建错误码并格式化(
			错误码类.CodeInvalidParameter,
			`unsupported OnDuplicateEx parameter type "%s"`,
			reflect.TypeOf(onDuplicateEx),
		)
	}
}

func (m *Model) getBatch() int {
	return m.batch
}
