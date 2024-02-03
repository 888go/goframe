// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb

import (
	"database/sql"
	"reflect"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/utils"
	"github.com/888go/goframe/os/gstructs"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gutil"
)

// With 启用关联查询，通过给定的属性对象指定开启。
// 常考"模型关联-静态关联"文档:https://goframe.org/pages/viewpage.action?pageId=7297190
// 例如，如果给定如下的结构体定义：
//
//	type User struct {
//		 gmeta.Meta `orm:"table:user"` // 定义表名为 user
//		 Id         int           `json:"id"`    // 用户ID
//		 Name       string        `json:"name"`   // 用户名
//		 UserDetail *UserDetail   `orm:"with:uid=id"` // 关联 UserDetail 表，通过 uid 等于 id 进行关联
//		 UserScores []*UserScores `orm:"with:uid=id"` // 关联 UserScores 表，通过 uid 等于 id 进行关联
//	}
//
// 我们可以通过以下方式在属性 `UserDetail` 和 `UserScores` 上启用模型关联操作：
// db.With(User{}.UserDetail).With(User{}.UserScores).Scan(xxx)
// 或者：
// db.With(UserDetail{}).With(UserScores{}).Scan(xxx)
// 或者：
// db.With(UserDetail{}, UserScores{}).Scan(xxx)
func (m *Model) With(objects ...interface{}) *Model {
	model := m.getModel()
	for _, object := range objects {
		if m.tables == "" {
			m.tablesInit = m.db.GetCore().QuotePrefixTableName(
				getTableNameFromOrmTag(object),
			)
			m.tables = m.tablesInit
			return model
		}
		model.withArray = append(model.withArray, object)
	}
	return model
}

// WithAll 开启在所有具有"struct"标签中包含"with"标签的对象上的模型关联操作。
// 常考"模型关联-静态关联"文档:https://goframe.org/pages/viewpage.action?pageId=7297190
func (m *Model) WithAll() *Model {
	model := m.getModel()
	model.withAll = true
	return model
}

// doWithScanStruct 处理单个结构体的模型关联操作功能。
func (m *Model) doWithScanStruct(pointer interface{}) error {
	var (
		err                 error
		allowedTypeStrArray = make([]string, 0)
	)
	currentStructFieldMap, err := gstructs.FieldMap(gstructs.FieldMapInput{
		Pointer:          pointer,
		PriorityTagArray: nil,
		RecursiveOption:  gstructs.RecursiveOptionEmbeddedNoTag,
	})
	if err != nil {
		return err
	}
	// 它会检查with数组，并自动调用ScanList完成关联查询。
	if !m.withAll {
		for _, field := range currentStructFieldMap {
			for _, withItem := range m.withArray {
				withItemReflectValueType, err := gstructs.StructType(withItem)
				if err != nil {
					return err
				}
				var (
					fieldTypeStr                = gstr.TrimAll(field.Type().String(), "*[]")
					withItemReflectValueTypeStr = gstr.TrimAll(withItemReflectValueType.String(), "*[]")
				)
				// 如果字段类型在指定的“with”类型数组中，则进行选择操作。
				if gstr.Compare(fieldTypeStr, withItemReflectValueTypeStr) == 0 {
					allowedTypeStrArray = append(allowedTypeStrArray, fieldTypeStr)
				}
			}
		}
	}
	for _, field := range currentStructFieldMap {
		var (
			fieldTypeStr    = gstr.TrimAll(field.Type().String(), "*[]")
			parsedTagOutput = m.parseWithTagInFieldStruct(field)
		)
		if parsedTagOutput.With == "" {
			continue
		}
		// 它仅处理“with”类型属性的结构体，因此会忽略其他类型的结构体。
		if !m.withAll && !gstr.InArray(allowedTypeStrArray, fieldTypeStr) {
			continue
		}
		array := gstr.SplitAndTrim(parsedTagOutput.With, "=")
		if len(array) == 1 {
// 它还支持仅使用一个列名
// 如果两个表关联时使用相同的列名。
			array = append(array, parsedTagOutput.With)
		}
		var (
			model              *Model
			fieldKeys          []string
			relatedSourceName  = array[0]
			relatedTargetName  = array[1]
			relatedTargetValue interface{}
		)
		// 从`pointer`中查找相关属性的值。
		for attributeName, attributeValue := range currentStructFieldMap {
			if utils.EqualFoldWithoutChars(attributeName, relatedTargetName) {
				relatedTargetValue = attributeValue.Value.Interface()
				break
			}
		}
		if relatedTargetValue == nil {
			return gerror.NewCodef(
				gcode.CodeInvalidParameter,
				`cannot find the target related value of name "%s" in with tag "%s" for attribute "%s.%s"`,
				relatedTargetName, parsedTagOutput.With, reflect.TypeOf(pointer).Elem(), field.Name(),
			)
		}
		bindToReflectValue := field.Value
		if bindToReflectValue.Kind() != reflect.Ptr && bindToReflectValue.CanAddr() {
			bindToReflectValue = bindToReflectValue.Addr()
		}

		// 它会自动从当前属性结构体/切片中检索结构体字段名称。
		if structType, err := gstructs.StructType(field.Value); err != nil {
			return err
		} else {
			fieldKeys = structType.FieldKeys()
		}

		// 递归并进行特性检查
		model = m.db.With(field.Value).Hook(m.hookHandler)
		if m.withAll {
			model = model.WithAll()
		} else {
			model = model.With(m.withArray...)
		}
		if parsedTagOutput.Where != "" {
			model = model.Where(parsedTagOutput.Where)
		}
		if parsedTagOutput.Order != "" {
			model = model.Order(parsedTagOutput.Order)
		}
		// 带有缓存功能。
		if m.cacheEnabled && m.cacheOption.Name == "" {
			model = model.Cache(m.cacheOption)
		}
		err = model.Fields(fieldKeys).
			Where(relatedSourceName, relatedTargetValue).
			Scan(bindToReflectValue)
		// 它在特性中忽略 sql.ErrNoRows 错误。
		if err != nil && err != sql.ErrNoRows {
			return err
		}
	}
	return nil
}

// doWithScanStructs 处理结构体切片的模型关联操作特性。
// 也可参考 doWithScanStruct。
func (m *Model) doWithScanStructs(pointer interface{}) error {
	if v, ok := pointer.(reflect.Value); ok {
		pointer = v.Interface()
	}

	var (
		err                 error
		allowedTypeStrArray = make([]string, 0)
	)
	currentStructFieldMap, err := gstructs.FieldMap(gstructs.FieldMapInput{
		Pointer:          pointer,
		PriorityTagArray: nil,
		RecursiveOption:  gstructs.RecursiveOptionEmbeddedNoTag,
	})
	if err != nil {
		return err
	}
	// 它会检查with数组，并自动调用ScanList完成关联查询。
	if !m.withAll {
		for _, field := range currentStructFieldMap {
			for _, withItem := range m.withArray {
				withItemReflectValueType, err := gstructs.StructType(withItem)
				if err != nil {
					return err
				}
				var (
					fieldTypeStr                = gstr.TrimAll(field.Type().String(), "*[]")
					withItemReflectValueTypeStr = gstr.TrimAll(withItemReflectValueType.String(), "*[]")
				)
				// 如果字段类型在指定的类型数组中，它将执行选择操作。
				if gstr.Compare(fieldTypeStr, withItemReflectValueTypeStr) == 0 {
					allowedTypeStrArray = append(allowedTypeStrArray, fieldTypeStr)
				}
			}
		}
	}

	for fieldName, field := range currentStructFieldMap {
		var (
			fieldTypeStr    = gstr.TrimAll(field.Type().String(), "*[]")
			parsedTagOutput = m.parseWithTagInFieldStruct(field)
		)
		if parsedTagOutput.With == "" {
			continue
		}
		if !m.withAll && !gstr.InArray(allowedTypeStrArray, fieldTypeStr) {
			continue
		}
		array := gstr.SplitAndTrim(parsedTagOutput.With, "=")
		if len(array) == 1 {
// 如果两个表使用相同的列名关联，则它支持仅使用一个列名。
			array = append(array, parsedTagOutput.With)
		}
		var (
			model              *Model
			fieldKeys          []string
			relatedSourceName  = array[0]
			relatedTargetName  = array[1]
			relatedTargetValue interface{}
		)
		// 从`pointer`中找到相关属性的值切片。
		for attributeName := range currentStructFieldMap {
			if utils.EqualFoldWithoutChars(attributeName, relatedTargetName) {
				relatedTargetValue = ListItemValuesUnique(pointer, attributeName)
				break
			}
		}
		if relatedTargetValue == nil {
			return gerror.NewCodef(
				gcode.CodeInvalidParameter,
				`cannot find the related value for attribute name "%s" of with tag "%s"`,
				relatedTargetName, parsedTagOutput.With,
			)
		}
		// 如果相关值为空，则此函数不做任何操作，仅返回。
		if gutil.IsEmpty(relatedTargetValue) {
			return nil
		}
		// 它会自动从当前属性结构体/切片中检索结构体字段名称。
		if structType, err := gstructs.StructType(field.Value); err != nil {
			return err
		} else {
			fieldKeys = structType.FieldKeys()
		}
		// 递归并进行特性检查
		model = m.db.With(field.Value).Hook(m.hookHandler)
		if m.withAll {
			model = model.WithAll()
		} else {
			model = model.With(m.withArray...)
		}
		if parsedTagOutput.Where != "" {
			model = model.Where(parsedTagOutput.Where)
		}
		if parsedTagOutput.Order != "" {
			model = model.Order(parsedTagOutput.Order)
		}
		// 带有缓存功能。
		if m.cacheEnabled && m.cacheOption.Name == "" {
			model = model.Cache(m.cacheOption)
		}
		err = model.Fields(fieldKeys).
			Where(relatedSourceName, relatedTargetValue).
			ScanList(pointer, fieldName, parsedTagOutput.With)
		// 它在特性中忽略 sql.ErrNoRows 错误。
		if err != nil && err != sql.ErrNoRows {
			return err
		}
	}
	return nil
}

type parseWithTagInFieldStructOutput struct {
	With  string
	Where string
	Order string
}

func (m *Model) parseWithTagInFieldStruct(field gstructs.Field) (output parseWithTagInFieldStructOutput) {
	var (
		ormTag = field.Tag(OrmTagForStruct)
		data   = make(map[string]string)
		array  []string
		key    string
	)
	for _, v := range gstr.SplitAndTrim(ormTag, " ") {
		array = gstr.Split(v, ":")
		if len(array) == 2 {
			key = array[0]
			data[key] = gstr.Trim(array[1])
		} else {
			data[key] += " " + gstr.Trim(v)
		}
	}
	for k, v := range data {
		data[k] = gstr.TrimRight(v, ",")
	}
	output.With = data[OrmTagForWith]
	output.Where = data[OrmTagForWithWhere]
	output.Order = data[OrmTagForWithOrder]
	return
}
