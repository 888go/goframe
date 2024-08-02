// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

import (
	"database/sql"
	"reflect"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/utils"
	"github.com/888go/goframe/os/gstructs"
	gstr "github.com/888go/goframe/text/gstr"
	gutil "github.com/888go/goframe/util/gutil"
)

// With 创建并返回一个基于给定对象元数据的 ORM 模型。它还为给定的 `object` 启用模型关联操作功能。
// 可以多次调用此函数，以向模型中添加一个或多个对象，并启用它们的模式关联操作功能。
// 例如，如果给出的结构体定义如下：
// 
// ```
// type User struct {
//     gmeta.Meta `orm:"table:user"`
//     Id         int           `json:"id"`
//     Name       string        `json:"name"`
//     UserDetail *UserDetail   `orm:"with:uid=id"`
//     UserScores []*UserScores `orm:"with:uid=id"`
// }
// ```
// 
// 我们可以通过以下方式在 `UserDetail` 和 `UserScores` 属性上启用模型关联操作：
// 
// ```
// db.With(User{}.UserDetail).With(User{}.UserScores).Scan(xxx)
// ```
// 
// 或者：
// 
// ```
// db.With(UserDetail{}).With(UserScores{}).Scan(xxx)
// ```
// 
// 或者：
// 
// ```
// db.With(UserDetail{}, UserScores{}).Scan(xxx)
// ```
// md5:c9498702475d54a9
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

// WithAll 启用对结构体中带有 "with" 标签的所有对象进行模型关联操作。 md5:83d3591315f0add0
func (m *Model) WithAll() *Model {
	model := m.getModel()
	model.withAll = true
	return model
}

// doWithScanStruct 处理单个结构体的模型关联操作功能。 md5:64dcc9bfd0382aa8
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
		// 它会检查with数组，并自动调用ScanList来完成关联查询。 md5:cb83f16b7131ad65
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
								// 如果字段类型在指定的"with"类型数组中，它会执行选择操作。 md5:b425357c98d952c8
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
				// 它仅处理带有"type"属性的"with"类型结构体，因此会忽略其他类型的结构体。 md5:c1f385406b699f00
		if !m.withAll && !gstr.InArray(allowedTypeStrArray, fieldTypeStr) {
			continue
		}
		array := gstr.SplitAndTrim(parsedTagOutput.With, "=")
		if len(array) == 1 {
			// 它还支持仅使用一个列名
			// 如果两个表使用相同的列名进行关联。
			// md5:c924339d8b4eddbc
			array = append(array, parsedTagOutput.With)
		}
		var (
			model              *Model
			fieldKeys          []string
			relatedSourceName  = array[0]
			relatedTargetName  = array[1]
			relatedTargetValue interface{}
		)
				// 从`pointer`中找到相关的属性值。 md5:b2da611599aed2d2
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

				// 它会自动从当前属性结构体/切片中获取字段名。 md5:09af2856a6801ffd
		if structType, err := gstructs.StructType(field.Value); err != nil {
			return err
		} else {
			fieldKeys = structType.FieldKeys()
		}

				// 递归实现并带有特性检查。 md5:9ddeb46ca8a2b86d
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
		// With cache feature.
		if m.cacheEnabled && m.cacheOption.Name == "" {
			model = model.Cache(m.cacheOption)
		}
		err = model.Fields(fieldKeys).
			Where(relatedSourceName, relatedTargetValue).
			Scan(bindToReflectValue)
				// 它在该特性中忽略sql.ErrNoRows错误。 md5:4b82d692c0646927
		if err != nil && err != sql.ErrNoRows {
			return err
		}
	}
	return nil
}

// doWithScanStructs 处理结构切片的模型关联操作功能。
// 参见 doWithScanStruct。
// md5:6219b8feabf0e7d9
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
		// 它会检查with数组，并自动调用ScanList来完成关联查询。 md5:cb83f16b7131ad65
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
								// 如果字段类型在指定的数组类型中，它将执行选择操作。 md5:afefe105662c6d79
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
			// 它支持仅使用一个列名的情况，
			// 当两个表通过相同的列名关联时。
			// md5:18222f22ecbee1ef
			array = append(array, parsedTagOutput.With)
		}
		var (
			model              *Model
			fieldKeys          []string
			relatedSourceName  = array[0]
			relatedTargetName  = array[1]
			relatedTargetValue interface{}
		)
				// 从`pointer`中查找相关属性的值切片。 md5:e729db1e29dfb929
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
				// 如果相关值为空，它什么也不做，只是返回。 md5:e4acb6a4c5d73f8f
		if gutil.IsEmpty(relatedTargetValue) {
			return nil
		}
				// 它会自动从当前属性结构体/切片中获取字段名。 md5:09af2856a6801ffd
		if structType, err := gstructs.StructType(field.Value); err != nil {
			return err
		} else {
			fieldKeys = structType.FieldKeys()
		}
				// 递归实现并带有特性检查。 md5:9ddeb46ca8a2b86d
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
		// With cache feature.
		if m.cacheEnabled && m.cacheOption.Name == "" {
			model = model.Cache(m.cacheOption)
		}
		err = model.Fields(fieldKeys).
			Where(relatedSourceName, relatedTargetValue).
			ScanList(pointer, fieldName, parsedTagOutput.With)
				// 它在该特性中忽略sql.ErrNoRows错误。 md5:4b82d692c0646927
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
