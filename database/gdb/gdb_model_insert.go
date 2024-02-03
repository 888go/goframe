// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb

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
func (m *Model) Batch(batch int) *Model {
	model := m.getModel()
	model.batch = batch
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
func (m *Model) Data(data ...interface{}) *Model {
	var model = m.getModel()
	if len(data) > 1 {
		if s := gconv.String(data[0]); gstr.Contains(s, "?") {
			model.data = s
			model.extraArgs = data[1:]
		} else {
			m := make(map[string]interface{})
			for i := 0; i < len(data); i += 2 {
				m[gconv.String(data[i])] = data[i+1]
			}
			model.data = m
		}
	} else if len(data) == 1 {
		switch value := data[0].(type) {
		case Result:
			model.data = value.List()

		case Record:
			model.data = value.Map()

		case List:
			list := make(List, len(value))
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
// 则为此条件添加`OmitNilData`选项，
// 这将会过滤掉`data`中所有为nil的参数。
					if isDoStruct(reflectInfo.OriginValue.Index(0).Interface()) {
						model = model.OmitNilData()
						model.option |= optionOmitNilDataInternal
					}
				}
				list := make(List, reflectInfo.OriginValue.Len())
				for i := 0; i < reflectInfo.OriginValue.Len(); i++ {
					list[i] = anyValueToMapBeforeToRecord(reflectInfo.OriginValue.Index(i).Interface())
				}
				model.data = list

			case reflect.Struct:
// 如果`data`参数是一个DO结构体，
// 则为此条件添加`OmitNilData`选项，
// 这将会过滤掉`data`中所有的nil参数。
				if isDoStruct(value) {
					model = model.OmitNilData()
				}
				if v, ok := data[0].(iInterfaces); ok {
					var (
						array = v.Interfaces()
						list  = make(List, len(array))
					)
					for i := 0; i < len(array); i++ {
						list[i] = anyValueToMapBeforeToRecord(array[i])
					}
					model.data = list
				} else {
					model.data = anyValueToMapBeforeToRecord(data[0])
				}

			case reflect.Map:
				model.data = anyValueToMapBeforeToRecord(data[0])

			default:
				model.data = data[0]
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
func (m *Model) OnDuplicate(onDuplicate ...interface{}) *Model {
	if len(onDuplicate) == 0 {
		return m
	}
	model := m.getModel()
	if len(onDuplicate) > 1 {
		model.onDuplicate = onDuplicate
	} else if len(onDuplicate) == 1 {
		model.onDuplicate = onDuplicate[0]
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
func (m *Model) OnDuplicateEx(onDuplicateEx ...interface{}) *Model {
	if len(onDuplicateEx) == 0 {
		return m
	}
	model := m.getModel()
	if len(onDuplicateEx) > 1 {
		model.onDuplicateEx = onDuplicateEx
	} else if len(onDuplicateEx) == 1 {
		model.onDuplicateEx = onDuplicateEx[0]
	}
	return model
}

// Insert 执行针对模型的 "INSERT INTO ..." 语句。
// 可选参数 `data` 与 Model.Data 函数的参数相同，
// 请参考 Model.Data。
func (m *Model) Insert(data ...interface{}) (result sql.Result, err error) {
	var ctx = m.GetCtx()
	if len(data) > 0 {
		return m.Data(data...).Insert()
	}
	return m.doInsertWithOption(ctx, InsertOptionDefault)
}

// InsertAndGetId 执行插入操作，并返回自动生成的最后一个插入ID。
func (m *Model) InsertAndGetId(data ...interface{}) (lastInsertId int64, err error) {
	var ctx = m.GetCtx()
	if len(data) > 0 {
		return m.Data(data...).InsertAndGetId()
	}
	result, err := m.doInsertWithOption(ctx, InsertOptionDefault)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// InsertIgnore 执行针对模型的 "INSERT IGNORE INTO ..." 语句。
// 可选参数 `data` 与 Model.Data 函数的参数相同，
// 请参阅 Model.Data。
func (m *Model) InsertIgnore(data ...interface{}) (result sql.Result, err error) {
	var ctx = m.GetCtx()
	if len(data) > 0 {
		return m.Data(data...).InsertIgnore()
	}
	return m.doInsertWithOption(ctx, InsertOptionIgnore)
}

// Replace 执行针对模型的 "REPLACE INTO ..." 语句。
// 可选参数 `data` 与 Model.Data 函数的参数相同，
// 详情请参阅 Model.Data。
func (m *Model) Replace(data ...interface{}) (result sql.Result, err error) {
	var ctx = m.GetCtx()
	if len(data) > 0 {
		return m.Data(data...).Replace()
	}
	return m.doInsertWithOption(ctx, InsertOptionReplace)
}

// Save 执行 "INSERT INTO ... ON DUPLICATE KEY UPDATE..." 语句，针对给定的 model。
// 可选参数 `data` 与 Model.Data 函数的参数相同，
// 请参阅 Model.Data。
//
// 如果保存的数据中存在主键或唯一索引，则更新记录，
// 否则会在表中插入一条新记录。
func (m *Model) Save(data ...interface{}) (result sql.Result, err error) {
	var ctx = m.GetCtx()
	if len(data) > 0 {
		return m.Data(data...).Save()
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
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "inserting into table with empty data")
	}
	var (
		list            List
		now             = gtime.Now()
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
		list = value.List()

	case Record:
		list = List{value.Map()}

	case List:
		list = value

	case Map:
		list = List{value}

	default:
// 这里使用gconv.Map简化从interface{}到map[string]interface{}的类型转换，
// 因为接下来的逻辑中会使用MapOrStructToMapDeep进行深度转换。
		reflectInfo := reflection.OriginValueAndKind(newData)
		switch reflectInfo.OriginKind {
		// 如果它是切片类型，那么将其转换为 List 类型。
		case reflect.Slice, reflect.Array:
			list = make(List, reflectInfo.OriginValue.Len())
			for i := 0; i < reflectInfo.OriginValue.Len(); i++ {
				list[i] = anyValueToMapBeforeToRecord(reflectInfo.OriginValue.Index(i).Interface())
			}

		case reflect.Map:
			list = List{anyValueToMapBeforeToRecord(value)}

		case reflect.Struct:
			if v, ok := value.(iInterfaces); ok {
				array := v.Interfaces()
				list = make(List, len(array))
				for i := 0; i < len(array); i++ {
					list[i] = anyValueToMapBeforeToRecord(array[i])
				}
			} else {
				list = List{anyValueToMapBeforeToRecord(value)}
			}

		default:
			return result, gerror.NewCodef(
				gcode.CodeInvalidParameter,
				"unsupported data list type: %v",
				reflectInfo.InputValue.Type(),
			)
		}
	}

	if len(list) < 1 {
		return result, gerror.NewCode(gcode.CodeMissingParameter, "data list cannot be empty")
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
		onDuplicateExKeySet := gset.NewStrSetFrom(onDuplicateExKeys)
		if m.onDuplicate != nil {
			switch m.onDuplicate.(type) {
			case Raw, *Raw:
				option.OnDuplicateStr = gconv.String(m.onDuplicate)

			default:
				reflectInfo := reflection.OriginValueAndKind(m.onDuplicate)
				switch reflectInfo.OriginKind {
				case reflect.String:
					option.OnDuplicateMap = make(map[string]interface{})
					for _, v := range gstr.SplitAndTrim(reflectInfo.OriginValue.String(), ",") {
						if onDuplicateExKeySet.Contains(v) {
							continue
						}
						option.OnDuplicateMap[v] = v
					}

				case reflect.Map:
					option.OnDuplicateMap = make(map[string]interface{})
					for k, v := range gconv.Map(m.onDuplicate) {
						if onDuplicateExKeySet.Contains(k) {
							continue
						}
						option.OnDuplicateMap[k] = v
					}

				case reflect.Slice, reflect.Array:
					option.OnDuplicateMap = make(map[string]interface{})
					for _, v := range gconv.Strings(m.onDuplicate) {
						if onDuplicateExKeySet.Contains(v) {
							continue
						}
						option.OnDuplicateMap[v] = v
					}

				default:
					return option, gerror.NewCodef(
						gcode.CodeInvalidParameter,
						`unsupported OnDuplicate parameter type "%s"`,
						reflect.TypeOf(m.onDuplicate),
					)
				}
			}
		} else if onDuplicateExKeySet.Size() > 0 {
			option.OnDuplicateMap = make(map[string]interface{})
			for _, v := range columnNames {
				if onDuplicateExKeySet.Contains(v) {
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
		return gstr.SplitAndTrim(reflectInfo.OriginValue.String(), ","), nil

	case reflect.Map:
		return gutil.Keys(onDuplicateEx), nil

	case reflect.Slice, reflect.Array:
		return gconv.Strings(onDuplicateEx), nil

	default:
		return nil, gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`unsupported OnDuplicateEx parameter type "%s"`,
			reflect.TypeOf(onDuplicateEx),
		)
	}
}

func (m *Model) getBatch() int {
	return m.batch
}
