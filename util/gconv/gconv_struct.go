// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gconv

import (
	"reflect"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/empty"
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/internal/utils"
	"github.com/gogf/gf/v2/util/gtag"
)

// 结构体将参数的键值对映射到对应结构对象的属性。
// 第三个参数 `mapping` 不必要，表示自定义键名和属性名之间的映射规则（区分大小写）。
// 
// 注意：
// 1. `params` 可以是任何类型的 map/struct，通常为 map。
// 2. `pointer` 应该是 *struct/**struct 类型，即指向结构体对象或结构体指针。
// 3. 只有结构体对象的公共属性可以被映射。
// 4. 如果 `params` 是一个 map，其键 `params` 可以是小写。在映射过程中，它会自动将键的首字母转换为大写进行匹配。如果键不匹配，它将忽略该键。
// md5:b39a46da903b06f5
func Struct(params interface{}, pointer interface{}, paramKeyToAttrMap ...map[string]string) (err error) {
	return Scan(params, pointer, paramKeyToAttrMap...)
}

// StructTag 作为 Struct 的功能，但同时也支持优先级标签特性。这个特性用于获取 `params` 键值对中的指定标签，并将其映射到结构体属性名上。
// 参数 `priorityTag` 支持多个标签，这些标签之间可以使用逗号 `,` 进行连接。
// md5:14d47a8c22737303
func StructTag(params interface{}, pointer interface{}, priorityTag string) (err error) {
	return doStruct(params, pointer, nil, priorityTag)
}

// doStruct 是将任何数据转换为结构体的核心内部函数。. md5:43cdc6b6cc398c7c
func doStruct(
	params interface{}, pointer interface{}, paramKeyToAttrMap map[string]string, priorityTag string,
) (err error) {
	if params == nil {
		// 如果`params`为nil，则不进行转换。. md5:0520708a0e7e1c1d
		return nil
	}
	if pointer == nil {
		return gerror.NewCode(gcode.CodeInvalidParameter, "object pointer cannot be nil")
	}

	// JSON内容转换。. md5:8a29b5a7aa430047
	ok, err := doConvertWithJsonCheck(params, pointer)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}

	defer func() {
		// 捕获panic，尤其是反射操作引发的panic。. md5:dd183bf8028f513a
		if exception := recover(); exception != nil {
			if v, ok := exception.(error); ok && gerror.HasStack(v) {
				err = v
			} else {
				err = gerror.NewCodeSkipf(gcode.CodeInternalPanic, 1, "%+v", exception)
			}
		}
	}()

	var (
		paramsReflectValue      reflect.Value
		paramsInterface         interface{} // 不要直接使用`params`，因为它可能是`reflect.Value`类型. md5:f469653f5ba4e08c
		pointerReflectValue     reflect.Value
		pointerReflectKind      reflect.Kind
		pointerElemReflectValue reflect.Value // The pointed element.
	)
	if v, ok := params.(reflect.Value); ok {
		paramsReflectValue = v
	} else {
		paramsReflectValue = reflect.ValueOf(params)
	}
	paramsInterface = paramsReflectValue.Interface()
	if v, ok := pointer.(reflect.Value); ok {
		pointerReflectValue = v
		pointerElemReflectValue = v
	} else {
		pointerReflectValue = reflect.ValueOf(pointer)
		pointerReflectKind = pointerReflectValue.Kind()
		if pointerReflectKind != reflect.Ptr {
			return gerror.NewCodef(
				gcode.CodeInvalidParameter,
				"destination pointer should be type of '*struct', but got '%v'",
				pointerReflectKind,
			)
		}
		// 使用 IsNil 检查 reflect.Ptr 类型的变量是可行的。. md5:0ba920ba8a6a19cf
		if !pointerReflectValue.IsValid() || pointerReflectValue.IsNil() {
			return gerror.NewCode(
				gcode.CodeInvalidParameter,
				"destination pointer cannot be nil",
			)
		}
		pointerElemReflectValue = pointerReflectValue.Elem()
	}

// 如果`params`和`pointer`是相同类型，直接进行赋值操作。
// 为了性能优化。
// md5:87eefbed1426eef0
	if ok = doConvertWithTypeCheck(paramsReflectValue, pointerElemReflectValue); ok {
		return nil
	}

	// custom convert.
	if ok, err = callCustomConverter(paramsReflectValue, pointerReflectValue); ok {
		return err
	}

	// 通常的接口解码检查。. md5:838cb73b6b92dc54
	if ok, err = bindVarToReflectValueWithInterfaceCheck(pointerReflectValue, paramsInterface); ok {
		return err
	}

// 如果必要，它会自动创建结构体对象。
// 例如，如果`pointer`是**User（双星号表示指针），那么`elem`就是*User，即User类型的指针。
// md5:172757349701f610
	if pointerElemReflectValue.Kind() == reflect.Ptr {
		if !pointerElemReflectValue.IsValid() || pointerElemReflectValue.IsNil() {
			e := reflect.New(pointerElemReflectValue.Type().Elem())
			pointerElemReflectValue.Set(e)
			defer func() {
				if err != nil {
					// 如果转换失败，它将`pointer`重置为nil。. md5:52f95bfcfceeefc0
					pointerReflectValue.Elem().Set(reflect.Zero(pointerReflectValue.Type().Elem()))
				}
			}()
		}
// 如果v, ok := pointerElemReflectValue.Interface().(iUnmarshalValue); ok {
// 	return v.UnmarshalValue(params)
// }
// 请注意，这里是`pointerElemReflectValue`而不是`pointerReflectValue`。
// md5:722eb6b1c6132d70
		if ok, err = bindVarToReflectValueWithInterfaceCheck(pointerElemReflectValue, paramsInterface); ok {
			return err
		}
		// 获取其元素，可能是最后的结构体。. md5:4a887dcf759fad9d
		pointerElemReflectValue = pointerElemReflectValue.Elem()
	}

// paramsMap 是一个类型为 map[string]interface{} 的变量，用于存储参数。
// 不要在這裡使用 MapDeep。
// md5:96735ea71b035d62
	paramsMap := doMapConvert(paramsInterface, recursiveTypeAuto, true)
	if paramsMap == nil {
		return gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`convert params from "%#v" to "map[string]interface{}" failed`,
			params,
		)
	}

	// 由于参数为空，无需进行任何操作。. md5:958747d8f67e1e73
	if len(paramsMap) == 0 {
		return nil
	}

	// 用于后续转换的信息存储。. md5:5cb67597e7ff966f
	type toBeConvertedFieldInfo struct {
		Value          any    // 从输入中通过标签名或字段名找到值。. md5:ec1aa28d82e3ec74
		FieldIndex     int    // 关联的反射字段索引。. md5:f8941ca77d57dd95
		FieldOrTagName string // 按优先级标签的字段名称或字段标签名称。. md5:2a10a58537738182
	}

	var (
		priorityTagArray                []string
		elemFieldName                   string
		elemFieldType                   reflect.StructField
		elemFieldValue                  reflect.Value
		elemType                        = pointerElemReflectValue.Type()
		toBeConvertedFieldNameToInfoMap = map[string]toBeConvertedFieldInfo{} // key=elemFieldName
	)

	if priorityTag != "" {
		priorityTagArray = append(utils.SplitAndTrim(priorityTag, ","), gtag.StructTagPriority...)
	} else {
		priorityTagArray = gtag.StructTagPriority
	}

	for i := 0; i < pointerElemReflectValue.NumField(); i++ {
		elemFieldType = elemType.Field(i)
		elemFieldName = elemFieldType.Name
		// 只转换为公共属性。. md5:4fc00fe51391895a
		if !utils.IsLetterUpper(elemFieldName[0]) {
			continue
		}

		var fieldTagName = getTagNameFromField(elemFieldType, priorityTagArray)
		// 也许它嵌入了struct。. md5:e77a8f08191e1bd2
		if elemFieldType.Anonymous {
// 定义一个名为Name的结构体，其中包含两个字段：LastName和FirstName，它们都有`json`标签进行标记
// ```
// type Name struct {
//    LastName  string `json:"lastName"`
//    FirstName string `json:"firstName"`
// }
// ```
// 
// 定义一个User结构体，其中包含一个嵌套的Name结构体，并使用`json:"name"`对整个嵌套结构进行标记
// ```
// type User struct {
//     Name `json:"name"`
//     // ...
// }
// ```
// 
// 只有当Name结构体中包含fieldTag（字段标签）时，才会记录这些信息
// md5:d42e389449351045
			if fieldTagName != "" {
				toBeConvertedFieldNameToInfoMap[elemFieldName] = toBeConvertedFieldInfo{
					FieldIndex:     elemFieldType.Index[0],
					FieldOrTagName: fieldTagName,
				}
			}

			elemFieldValue = pointerElemReflectValue.Field(i)
			// 如果接口属性为nil，则忽略它。. md5:5bbafbaa5b14794d
			if elemFieldValue.Kind() == reflect.Interface {
				elemFieldValue = elemFieldValue.Elem()
				if !elemFieldValue.IsValid() {
					continue
				}
			}
			if err = doStruct(paramsMap, elemFieldValue, paramKeyToAttrMap, priorityTag); err != nil {
				return err
			}
		} else {
			// 使用原生的elemFieldName名称作为字段标签. md5:80bfd9b406ef430f
			if fieldTagName == "" {
				fieldTagName = elemFieldName
			}
			toBeConvertedFieldNameToInfoMap[elemFieldName] = toBeConvertedFieldInfo{
				FieldIndex:     elemFieldType.Index[0],
				FieldOrTagName: fieldTagName,
			}
		}
	}

	// 没有需要转换的内容。. md5:68441f55873cce91
	if len(toBeConvertedFieldNameToInfoMap) == 0 {
		return nil
	}

	// 在参数值中搜索该字段。. md5:761e9f220df7696c
	var paramsValue any
	for fieldName, fieldInfo := range toBeConvertedFieldNameToInfoMap {
		if paramsValue, ok = paramsMap[fieldInfo.FieldOrTagName]; ok {
			fieldInfo.Value = paramsValue
			toBeConvertedFieldNameToInfoMap[fieldName] = fieldInfo
		}
	}

// 首先，根据自定义的映射规则进行搜索。
// 如果找到了可能的直接赋值关系，减少后续映射搜索的数量。
// md5:50dd567944f99367
	var fieldInfo toBeConvertedFieldInfo
	for paramKey, fieldName := range paramKeyToAttrMap {
		// 防止设置不存在的字段. md5:408a34ea9e6a0539
		fieldInfo, ok = toBeConvertedFieldNameToInfoMap[fieldName]
		if ok {
			// 防止不存在的值被设置。. md5:16a6e1bcb81b8eb9
			if paramsValue, ok = paramsMap[paramKey]; ok {
				fieldInfo.Value = paramsValue
				toBeConvertedFieldNameToInfoMap[fieldName] = fieldInfo
			}
		}
	}

	var (
		paramKey   string
		paramValue any
		fieldName  string
		// 表示这些值已被使用，不能重复使用。. md5:66845c8e5a8adbe8
		usedParamsKeyOrTagNameMap = map[string]struct{}{}
	)
	for fieldName, fieldInfo = range toBeConvertedFieldNameToInfoMap {
		// 如果非空，标签或elemFieldName的名称匹配. md5:dcf5990abe97052c
		if fieldInfo.Value != nil {
			if err = bindVarToStructAttrWithFieldIndex(
				pointerElemReflectValue, fieldName, fieldInfo.FieldIndex, fieldInfo.Value, paramKeyToAttrMap,
			); err != nil {
				return err
			}
			usedParamsKeyOrTagNameMap[fieldInfo.FieldOrTagName] = struct{}{}
			continue
		}

		// 如果value为nil，搜索时会使用模糊匹配来转换键和值。. md5:30209602b5ceef13
		paramKey, paramValue = fuzzyMatchingFieldName(fieldName, paramsMap, usedParamsKeyOrTagNameMap)
		if paramValue != nil {
			if err = bindVarToStructAttrWithFieldIndex(
				pointerElemReflectValue, fieldName, fieldInfo.FieldIndex, paramValue, paramKeyToAttrMap,
			); err != nil {
				return err
			}
			usedParamsKeyOrTagNameMap[paramKey] = struct{}{}
		}
	}
	return nil
}

func getTagNameFromField(field reflect.StructField, priorityTags []string) string {
	for _, tag := range priorityTags {
		value, ok := field.Tag.Lookup(tag)
		if ok {
// 如果标签字符串中还有其他内容，
// 它会使用以逗号','分隔的第一部分。
// 例如：
// `orm:"id, priority"`
// `orm:"name, with:uid=id"` 
// 
// 这段注释说明了一个Go语言中的ORM（对象关系映射）相关代码。它解释了当解析一个包含多个属性的标签字符串时，程序会选择以逗号分隔的第一个属性作为主要处理的部分。如果标签格式为`attribute1, attribute2`，则只会使用`attribute1`。另一个例子展示了如何在`name`属性中使用额外的条件，即`with:uid=id`。
// md5:fab9db8addb2ccc4
			array := strings.Split(value, ",")
			// json:",omitempty"
			trimmedTagName := strings.TrimSpace(array[0])
			return trimmedTagName
		}
	}
	return ""
}

// 模糊匹配规则：
// 不区分大小写，不考虑符号地匹配字段名和参数键。
// md5:22c4645c8af23d0d
func fuzzyMatchingFieldName(
	fieldName string,
	paramsMap map[string]any,
	usedParamsKeyMap map[string]struct{},
) (string, any) {
	fieldName = utils.RemoveSymbols(fieldName)
	for paramKey, paramValue := range paramsMap {
		if _, ok := usedParamsKeyMap[paramKey]; ok {
			continue
		}
		removeParamKeyUnderline := utils.RemoveSymbols(paramKey)
		if strings.EqualFold(fieldName, removeParamKeyUnderline) {
			return paramKey, paramValue
		}
	}
	return "", nil
}

// bindVarToStructAttrWithFieldIndex 通过名称将值设置给结构体对象的属性。. md5:884feed9b741e07a
func bindVarToStructAttrWithFieldIndex(
	structReflectValue reflect.Value, attrName string,
	fieldIndex int, value interface{}, paramKeyToAttrMap map[string]string,
) (err error) {
	structFieldValue := structReflectValue.Field(fieldIndex)
	if !structFieldValue.IsValid() {
		return nil
	}
	// CanSet 检查该属性是否可以公开访问。. md5:fafe4f3a8bd7621f
	if !structFieldValue.CanSet() {
		return nil
	}
	defer func() {
		if exception := recover(); exception != nil {
			if err = bindVarToReflectValue(structFieldValue, value, paramKeyToAttrMap); err != nil {
				err = gerror.Wrapf(err, `error binding value to attribute "%s"`, attrName)
			}
		}
	}()
	// Directly converting.
	if empty.IsNil(value) {
		structFieldValue.Set(reflect.Zero(structFieldValue.Type()))
	} else {
// 尝试调用自定义转换器。
// 问题：https://github.com/gogf/gf/issues/3099
// md5:e874679d6ecc39f0
		var (
			customConverterInput reflect.Value
			ok                   bool
		)
		if customConverterInput, ok = value.(reflect.Value); !ok {
			customConverterInput = reflect.ValueOf(value)
		}

		if ok, err = callCustomConverter(customConverterInput, structFieldValue); ok || err != nil {
			return
		}

// 对某些类型进行特殊处理：
// - 重写stdlib中time.Time类型的默认类型转换逻辑。
// md5:39ca7f7684bdc13c
		var structFieldTypeName = structFieldValue.Type().String()
		switch structFieldTypeName {
		case "time.Time", "*time.Time":
			doConvertWithReflectValueSet(structFieldValue, doConvertInput{
				FromValue:  value,
				ToTypeName: structFieldTypeName,
				ReferValue: structFieldValue,
			})
			return
// 在递归中保持时区一致
// 问题：https://github.com/gogf/gf/issues/2980
// md5:1d09e937a28bf051
		case "*gtime.Time", "gtime.Time":
			doConvertWithReflectValueSet(structFieldValue, doConvertInput{
				FromValue:  value,
				ToTypeName: structFieldTypeName,
				ReferValue: structFieldValue,
			})
			return
		}

		// 公共接口检查。. md5:0e7cc3af409e672f
		if ok, err = bindVarToReflectValueWithInterfaceCheck(structFieldValue, value); ok {
			return err
		}

		// Default converting.
		doConvertWithReflectValueSet(structFieldValue, doConvertInput{
			FromValue:  value,
			ToTypeName: structFieldTypeName,
			ReferValue: structFieldValue,
		})
	}
	return nil
}

// bindVarToReflectValueWithInterfaceCheck 使用通用接口检查进行绑定。. md5:ede209e9eacebf79
func bindVarToReflectValueWithInterfaceCheck(reflectValue reflect.Value, value interface{}) (bool, error) {
	var pointer interface{}
	if reflectValue.Kind() != reflect.Ptr && reflectValue.CanAddr() {
		reflectValueAddr := reflectValue.Addr()
		if reflectValueAddr.IsNil() || !reflectValueAddr.IsValid() {
			return false, nil
		}
		// 不是指针，但可以处理地址，因此它可以被反序列化。. md5:52a739dbed72b8c0
		pointer = reflectValue.Addr().Interface()
	} else {
		if reflectValue.IsNil() || !reflectValue.IsValid() {
			return false, nil
		}
		pointer = reflectValue.Interface()
	}
	// UnmarshalValue.
	if v, ok := pointer.(iUnmarshalValue); ok {
		return ok, v.UnmarshalValue(value)
	}
	// UnmarshalText.
	if v, ok := pointer.(iUnmarshalText); ok {
		var valueBytes []byte
		if b, ok := value.([]byte); ok {
			valueBytes = b
		} else if s, ok := value.(string); ok {
			valueBytes = []byte(s)
		} else if f, ok := value.(iString); ok {
			valueBytes = []byte(f.String())
		}
		if len(valueBytes) > 0 {
			return ok, v.UnmarshalText(valueBytes)
		}
	}
	// UnmarshalJSON.
	if v, ok := pointer.(iUnmarshalJSON); ok {
		var valueBytes []byte
		if b, ok := value.([]byte); ok {
			valueBytes = b
		} else if s, ok := value.(string); ok {
			valueBytes = []byte(s)
		} else if f, ok := value.(iString); ok {
			valueBytes = []byte(f.String())
		}

		if len(valueBytes) > 0 {
			// 如果它不是一个有效的JSON字符串，那么就在它的两边添加字符 `"` 以使其成为有效JSON字符串。. md5:d6a38f1500604604
			if !json.Valid(valueBytes) {
				newValueBytes := make([]byte, len(valueBytes)+2)
				newValueBytes[0] = '"'
				newValueBytes[len(newValueBytes)-1] = '"'
				copy(newValueBytes[1:], valueBytes)
				valueBytes = newValueBytes
			}
			return ok, v.UnmarshalJSON(valueBytes)
		}
	}
	if v, ok := pointer.(iSet); ok {
		v.Set(value)
		return ok, nil
	}
	return false, nil
}

// bindVarToReflectValue 将 `value` 设置为反射值对象 `structFieldValue`。. md5:c78b60ec569060eb
func bindVarToReflectValue(
	structFieldValue reflect.Value, value interface{}, paramKeyToAttrMap map[string]string,
) (err error) {
	// JSON内容转换。. md5:8a29b5a7aa430047
	ok, err := doConvertWithJsonCheck(value, structFieldValue)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}

	kind := structFieldValue.Kind()
	// 使用`Set`接口实现转换，对于某些类型。. md5:51e8e3ad23771259
	switch kind {
	case reflect.Slice, reflect.Array, reflect.Ptr, reflect.Interface:
		if !structFieldValue.IsNil() {
			if v, ok := structFieldValue.Interface().(iSet); ok {
				v.Set(value)
				return nil
			}
		}
	}

	// 使用反射按类型进行转换。. md5:e3c406f111505fd2
	switch kind {
	case reflect.Map:
		return doMapToMap(value, structFieldValue, paramKeyToAttrMap)

	case reflect.Struct:
		// 递归转换结构体属性。. md5:ae6513ef6e56f654
		if err = doStruct(value, structFieldValue, nil, ""); err != nil {
			// 请注意这里存在反射转换机制。. md5:84599bf48af19237
			structFieldValue.Set(reflect.ValueOf(value).Convert(structFieldValue.Type()))
		}

// 注意，切片元素的类型可能是结构体，
// 因此它内部使用了一个名为Struct的函数来进行转换。
// md5:b8519d4d1a736c40
	case reflect.Slice, reflect.Array:
		var (
			reflectArray reflect.Value
			reflectValue = reflect.ValueOf(value)
		)
		if reflectValue.Kind() == reflect.Slice || reflectValue.Kind() == reflect.Array {
			reflectArray = reflect.MakeSlice(structFieldValue.Type(), reflectValue.Len(), reflectValue.Len())
			if reflectValue.Len() > 0 {
				var (
					elemType     = reflectArray.Index(0).Type()
					elemTypeName string
					converted    bool
				)
				for i := 0; i < reflectValue.Len(); i++ {
					converted = false
					elemTypeName = elemType.Name()
					if elemTypeName == "" {
						elemTypeName = elemType.String()
					}
					var elem reflect.Value
					if elemType.Kind() == reflect.Ptr {
						elem = reflect.New(elemType.Elem()).Elem()
					} else {
						elem = reflect.New(elemType).Elem()
					}
					if elem.Kind() == reflect.Struct {
						if err = doStruct(reflectValue.Index(i).Interface(), elem, nil, ""); err == nil {
							converted = true
						}
					}
					if !converted {
						doConvertWithReflectValueSet(elem, doConvertInput{
							FromValue:  reflectValue.Index(i).Interface(),
							ToTypeName: elemTypeName,
							ReferValue: elem,
						})
					}
					if elemType.Kind() == reflect.Ptr {
						// 在将`elem`设置为数组之前，如果必要的话进行指针转换。. md5:1466632fc1d552e6
						elem = elem.Addr()
					}
					reflectArray.Index(i).Set(elem)
				}
			}
		} else {
			var (
				elem         reflect.Value
				elemType     = structFieldValue.Type().Elem()
				elemTypeName = elemType.Name()
				converted    bool
			)
			switch reflectValue.Kind() {
			case reflect.String:
				// Value is empty string.
				if reflectValue.IsZero() {
					var elemKind = elemType.Kind()
					// 尝试找到切片元素的原始类型类别。. md5:903e45eb4bc9a592
					if elemKind == reflect.Ptr {
						elemKind = elemType.Elem().Kind()
					}
					switch elemKind {
					case reflect.String:
						// 空字符串不能赋值给字符串切片。. md5:7015d8a83525c473
						return nil
					}
				}
			}
			if elemTypeName == "" {
				elemTypeName = elemType.String()
			}
			if elemType.Kind() == reflect.Ptr {
				elem = reflect.New(elemType.Elem()).Elem()
			} else {
				elem = reflect.New(elemType).Elem()
			}
			if elem.Kind() == reflect.Struct {
				if err = doStruct(value, elem, nil, ""); err == nil {
					converted = true
				}
			}
			if !converted {
				doConvertWithReflectValueSet(elem, doConvertInput{
					FromValue:  value,
					ToTypeName: elemTypeName,
					ReferValue: elem,
				})
			}
			if elemType.Kind() == reflect.Ptr {
				// 在将`elem`设置为数组之前，如果必要的话进行指针转换。. md5:1466632fc1d552e6
				elem = elem.Addr()
			}
			reflectArray = reflect.MakeSlice(structFieldValue.Type(), 1, 1)
			reflectArray.Index(0).Set(elem)
		}
		structFieldValue.Set(reflectArray)

	case reflect.Ptr:
		if structFieldValue.IsNil() || structFieldValue.IsZero() {
			// 如果是空指针或空列表，它会创建一个新的。. md5:a005c5e6ed40f663
			item := reflect.New(structFieldValue.Type().Elem())
			if ok, err = bindVarToReflectValueWithInterfaceCheck(item, value); ok {
				structFieldValue.Set(item)
				return err
			}
			elem := item.Elem()
			if err = bindVarToReflectValue(elem, value, paramKeyToAttrMap); err == nil {
				structFieldValue.Set(elem.Addr())
			}
		} else {
			// 非空指针，它会给它赋值。. md5:2bd4c15a81dcbdcf
			return bindVarToReflectValue(structFieldValue.Elem(), value, paramKeyToAttrMap)
		}

	// 它主要且特别地处理了nil值的接口。. md5:0c8e2dd31d82d96e
	case reflect.Interface:
		if value == nil {
			// Specially.
			structFieldValue.Set(reflect.ValueOf((*interface{})(nil)))
		} else {
			// 请注意这里存在反射转换机制。. md5:84599bf48af19237
			structFieldValue.Set(reflect.ValueOf(value).Convert(structFieldValue.Type()))
		}

	default:
		defer func() {
			if exception := recover(); exception != nil {
				err = gerror.NewCodef(
					gcode.CodeInternalPanic,
					`cannot convert value "%+v" to type "%s":%+v`,
					value,
					structFieldValue.Type().String(),
					exception,
				)
			}
		}()
// 此处使用反射将`value`转换为属性的类型，然后将结果值赋给该属性。
// 如果常规的Go转换规则不允许转换，此操作可能会失败并引发恐慌。
// md5:931b86f723a12b7c
		structFieldValue.Set(reflect.ValueOf(value).Convert(structFieldValue.Type()))
	}
	return nil
}
