// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 转换类

import (
	"reflect"
	"strings"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/empty"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/utils"
	"github.com/888go/goframe/os/gstructs"
)

// Struct 将参数键值对映射到相应结构体对象的属性上。
// 第三个参数 `mapping` 是可选的，表示自定义键名与属性名（大小写敏感）之间的映射规则。
// 注意：
// 1. `params` 可以是任何类型的 map 或 struct，通常为 map 类型。
// 2. `pointer` 应为 *struct/**struct 类型，即指向结构体对象或结构体指针的指针。
// 3. 只有结构体对象的公共属性可以被映射。
// 4. 如果 `params` 是一个 map，map 的键可以是小写的。
//    在映射过程中，会自动将键的首字母转换为大写进行匹配。
//    如果不匹配，则忽略该 map 键。
func Struct(params interface{}, pointer interface{}, paramKeyToAttrMap ...map[string]string) (err error) {
	return Scan(params, pointer, paramKeyToAttrMap...)
}

// StructTag 结构体在 Struct 的基础上增加了支持优先级标签功能，该功能用于获取 `params` 键值对中指定的标签，并映射到结构体属性名称。  
// 参数 `priorityTag` 支持多个标签，多个标签之间可以通过字符 ',' 连接。
func StructTag(params interface{}, pointer interface{}, priorityTag string) (err error) {
	return doStruct(params, pointer, nil, priorityTag)
}

// doStructWithJsonCheck 检查给定的 `params` 是否为 JSON 格式，如果是，则使用 json.Unmarshal 进行转换。
func doStructWithJsonCheck(params interface{}, pointer interface{}) (err error, ok bool) {
	switch r := params.(type) {
	case []byte:
		if json.Valid(r) {
			if rv, ok := pointer.(reflect.Value); ok {
				if rv.Kind() == reflect.Ptr {
					if rv.IsNil() {
						return nil, false
					}
					return json.UnmarshalUseNumber(r, rv.Interface()), true
				} else if rv.CanAddr() {
					return json.UnmarshalUseNumber(r, rv.Addr().Interface()), true
				}
			} else {
				return json.UnmarshalUseNumber(r, pointer), true
			}
		}
	case string:
		if paramsBytes := []byte(r); json.Valid(paramsBytes) {
			if rv, ok := pointer.(reflect.Value); ok {
				if rv.Kind() == reflect.Ptr {
					if rv.IsNil() {
						return nil, false
					}
					return json.UnmarshalUseNumber(paramsBytes, rv.Interface()), true
				} else if rv.CanAddr() {
					return json.UnmarshalUseNumber(paramsBytes, rv.Addr().Interface()), true
				}
			} else {
				return json.UnmarshalUseNumber(paramsBytes, pointer), true
			}
		}
	default:
		// `params` 可能是一个实现了 Interface 接口函数的结构体，例如：gvar.Var。
		if v, ok := params.(iInterface); ok {
			return doStructWithJsonCheck(v.Interface(), pointer)
		}
	}
	return nil, false
}

// doStruct 是用于将任何数据转换为结构体的核心内部函数。
func doStruct(params interface{}, pointer interface{}, paramKeyToAttrMap map[string]string, priorityTag string) (err error) {
	if params == nil {
		// 如果`params`为nil，则不进行转换。
		return nil
	}
	if pointer == nil {
		return 错误类.X创建错误码(错误码类.CodeInvalidParameter, "object pointer cannot be nil")
	}

	defer func() {
		// 捕获 panic，特别是反射操作引发的 panic。
		if exception := recover(); exception != nil {
			if v, ok := exception.(error); ok && 错误类.X判断是否带堆栈(v) {
				err = v
			} else {
				err = 错误类.X创建错误码并跳过堆栈与格式化(错误码类.CodeInternalPanic, 1, "%+v", exception)
			}
		}
	}()

	// JSON内容转换
	err, ok := doStructWithJsonCheck(params, pointer)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}

	var (
		paramsReflectValue      reflect.Value
		paramsInterface         interface{} // **请勿**直接使用 `params`，因为它可能为 `reflect.Value` 类型
		pointerReflectValue     reflect.Value
		pointerReflectKind      reflect.Kind
		pointerElemReflectValue reflect.Value // 指向的元素。
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
			return 错误类.X创建错误码并格式化(错误码类.CodeInvalidParameter, "object pointer should be type of '*struct', but got '%v'", pointerReflectKind)
		}
		// 对reflect.Ptr类型的变量使用IsNil方法是可行的。
		if !pointerReflectValue.IsValid() || pointerReflectValue.IsNil() {
			return 错误类.X创建错误码(错误码类.CodeInvalidParameter, "object pointer cannot be nil")
		}
		pointerElemReflectValue = pointerReflectValue.Elem()
	}

	// 首先尝试自定义转换
	if ok, err = callCustomConverter(paramsReflectValue, pointerReflectValue); ok {
		return err
	}

// 如果`params`和`pointer`是相同类型，则直接进行赋值操作。
// 为了提升性能。
	if pointerElemReflectValue.IsValid() {
		switch {
// 例如：
// UploadFile  => 上传文件
// *UploadFile => 指向UploadFile的指针
		case pointerElemReflectValue.Type() == paramsReflectValue.Type():
			pointerElemReflectValue.Set(paramsReflectValue)
			return nil

// 例如：
// UploadFile  => *UploadFile
// （此代码注释翻译为：）
// 示例：
// UploadFile  => 指向UploadFile类型的指针
		case pointerElemReflectValue.Kind() == reflect.Ptr && pointerElemReflectValue.Elem().IsValid() &&
			pointerElemReflectValue.Elem().Type() == paramsReflectValue.Type():
			pointerElemReflectValue.Elem().Set(paramsReflectValue)
			return nil

// 示例：
// *UploadFile  => 上传文件
		case paramsReflectValue.Kind() == reflect.Ptr && paramsReflectValue.Elem().IsValid() &&
			pointerElemReflectValue.Type() == paramsReflectValue.Elem().Type():
			pointerElemReflectValue.Set(paramsReflectValue.Elem())
			return nil
		}
	}

	// 正常的接口解码检查。
	if err, ok = bindVarToReflectValueWithInterfaceCheck(pointerReflectValue, paramsInterface); ok {
		return err
	}

// 如果有必要，它会自动创建结构体对象。
// 例如，如果 `pointer` 是 **User 类型，那么 `elem` 就是 *User 类型，即指向 User 的指针。
	if pointerElemReflectValue.Kind() == reflect.Ptr {
		if !pointerElemReflectValue.IsValid() || pointerElemReflectValue.IsNil() {
			e := reflect.New(pointerElemReflectValue.Type().Elem())
			pointerElemReflectValue.Set(e)
			defer func() {
				if err != nil {
					// 如果转换失败，则将`pointer`重置为nil。
					pointerReflectValue.Elem().Set(reflect.Zero(pointerReflectValue.Type().Elem()))
				}
			}()
		}
// 如果v, ok := 将pointerElemReflectValue.Interface().(iUnmarshalValue)进行类型断言并赋值；如果ok为真（即转换成功），
// 则返回v.UnmarshalValue(params)的结果
// 注意：这里使用的是`pointerElemReflectValue`而非`pointerReflectValue`
// 这段代码的中文注释翻译如下：
// ```go
// 若能将pointerElemReflectValue.Interface()转换为iUnmarshalValue类型，并将转换结果赋值给v和ok，且ok为真，
// 则调用v的UnmarshalValue方法处理params并返回其结果。
// 需要注意的是，此处使用的是`pointerElemReflectValue`变量，而不是`pointerReflectValue`变量。
		if err, ok = bindVarToReflectValueWithInterfaceCheck(pointerElemReflectValue, paramsInterface); ok {
			return err
		}
		// 获取其元素，最后可能是一个结构体。
		pointerElemReflectValue = pointerElemReflectValue.Elem()
	}

// paramsMap 是用于参数的 map[string]interface{} 类型变量。
// 在此处不要使用 MapDeep。
	paramsMap := doMapConvert(paramsInterface, recursiveTypeAuto, true)
	if paramsMap == nil {
		return 错误类.X创建错误码并格式化(
			错误码类.CodeInvalidParameter,
			`convert params from "%#v" to "map[string]interface{}" failed`,
			params,
		)
	}

	// 由于参数为空，无需执行任何操作。
	if len(paramsMap) == 0 {
		return nil
	}

// 它只执行同一属性的一次转换。
// doneMap 用于检查重复转换，其键是结构体的实际属性名称。
	doneMap := make(map[string]struct{})

// attrMap 的键是结构体的属性名，
// 值是用于后续比较时的替换名称，目的是为了提升性能。
	var (
		tempName       string
		elemFieldType  reflect.StructField
		elemFieldValue reflect.Value
		elemType       = pointerElemReflectValue.Type()
// 将属性名称映射到去除符号后的名称，
// 以便在后续逻辑中快速索引和比较。
		attrToCheckNameMap = make(map[string]string)
	)
	for i := 0; i < pointerElemReflectValue.NumField(); i++ {
		elemFieldType = elemType.Field(i)
		// 只对公开属性进行转换
		if !utils.IsLetterUpper(elemFieldType.Name[0]) {
			continue
		}
		// 可能这是一个结构体 /* 或是结构体嵌入.
		if elemFieldType.Anonymous {
			elemFieldValue = pointerElemReflectValue.Field(i)
			// 如果接口属性为nil，则忽略它。
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
			tempName = elemFieldType.Name
			attrToCheckNameMap[tempName] = utils.RemoveSymbols(tempName)
		}
	}
	if len(attrToCheckNameMap) == 0 {
		return nil
	}

// `attrToTagCheckNameMap` 的键是结构体的属性名称，
// 而值则是用于后续比较时替换的标签名称，目的是为了提高性能。
	var (
		attrToTagCheckNameMap = make(map[string]string)
		priorityTagArray      []string
	)
	if priorityTag != "" {
		priorityTagArray = append(utils.SplitAndTrim(priorityTag, ","), StructTagPriority...)
	} else {
		priorityTagArray = StructTagPriority
	}
	tagToAttrNameMap, err := gstructs.TagMapName(pointerElemReflectValue, priorityTagArray)
	if err != nil {
		return err
	}
	for tagName, attributeName := range tagToAttrNameMap {
// 如果标签字符串中还有其他内容，
// 它会使用通过逗号（,）分割得到的第一部分。
// 例如：
// orm:"id, priority" // 使用id和priority
// orm:"name, with:uid=id" // 使用name和with:uid=id中的name部分
		attrToTagCheckNameMap[attributeName] = utils.RemoveSymbols(strings.Split(tagName, ",")[0])
// 如果tag和attribute值同时存在于`paramsMap`中，
// 则优先使用tag值，并在`paramsMap`中覆盖原有的attribute值。
		if paramsMap[tagName] != nil && paramsMap[attributeName] != nil {
			paramsMap[attributeName] = paramsMap[tagName]
		}
	}

	// 根据自定义参数键到属性名映射来转换值。
	err = doStructBaseOnParamKeyToAttrMap(
		pointerElemReflectValue,
		paramsMap,
		paramKeyToAttrMap,
		doneMap,
	)
	if err != nil {
		return err
	}
	// 已经完成了所有属性值的赋值，接下来无事可做。
	if len(doneMap) == len(attrToCheckNameMap) {
		return nil
	}

	// 根据精确属性名称转换值。
	err = doStructBaseOnAttribute(
		pointerElemReflectValue,
		paramsMap,
		paramKeyToAttrMap,
		doneMap,
		attrToCheckNameMap,
	)
	if err != nil {
		return err
	}
	// 已经完成了所有属性值的赋值，接下来无事可做。
	if len(doneMap) == len(attrToCheckNameMap) {
		return nil
	}

	// 根据参数映射转换值。
	err = doStructBaseOnParamMap(
		pointerElemReflectValue,
		paramsMap,
		paramKeyToAttrMap,
		doneMap,
		attrToCheckNameMap,
		attrToTagCheckNameMap,
		tagToAttrNameMap,
	)
	if err != nil {
		return err
	}
	return nil
}

func doStructBaseOnParamKeyToAttrMap(
	pointerElemReflectValue reflect.Value,
	paramsMap map[string]interface{},
	paramKeyToAttrMap map[string]string,
	doneAttrMap map[string]struct{},
) error {
	if len(paramKeyToAttrMap) == 0 {
		return nil
	}
	for paramKey, attrName := range paramKeyToAttrMap {
		paramValue, ok := paramsMap[paramKey]
		if !ok {
			continue
		}
		// 如果属性名称已经经过转换检查，那么跳过它。
		if _, ok = doneAttrMap[attrName]; ok {
			continue
		}
		// Mark it done.
		doneAttrMap[attrName] = struct{}{}
		if err := bindVarToStructAttr(
			pointerElemReflectValue, attrName, paramValue, paramKeyToAttrMap,
		); err != nil {
			return err
		}
	}
	return nil
}

func doStructBaseOnAttribute(
	pointerElemReflectValue reflect.Value,
	paramsMap map[string]interface{},
	paramKeyToAttrMap map[string]string,
	doneAttrMap map[string]struct{},
	attrToCheckNameMap map[string]string,
) error {
	var customMappingAttrMap = make(map[string]struct{})
	if len(paramKeyToAttrMap) > 0 {
		// 如果属性名在`paramKeyToAttrMap`中指定，它将忽略这些属性名。
		for paramName := range paramsMap {
			if passedAttrKey, ok := paramKeyToAttrMap[paramName]; ok {
				customMappingAttrMap[passedAttrKey] = struct{}{}
			}
		}
	}
	for attrName := range attrToCheckNameMap {
		// 通过精确属性名称获取的值。
		paramValue, ok := paramsMap[attrName]
		if !ok {
			continue
		}
		// 如果属性名存在于自定义的paramKeyToAttrMap中，则忽略该转换操作。
		if _, ok = customMappingAttrMap[attrName]; ok {
			continue
		}
		// 如果属性名称已经经过转换检查，那么跳过它。
		if _, ok = doneAttrMap[attrName]; ok {
			continue
		}
		// Mark it done.
		doneAttrMap[attrName] = struct{}{}
		if err := bindVarToStructAttr(
			pointerElemReflectValue, attrName, paramValue, paramKeyToAttrMap,
		); err != nil {
			return err
		}
	}
	return nil
}

func doStructBaseOnParamMap(
	pointerElemReflectValue reflect.Value,
	paramsMap map[string]interface{},
	paramKeyToAttrMap map[string]string,
	doneAttrMap map[string]struct{},
	attrToCheckNameMap map[string]string,
	attrToTagCheckNameMap map[string]string,
	tagToAttrNameMap map[string]string,
) error {
	var (
		attrName  string
		checkName string
	)
	for paramName, paramValue := range paramsMap {
// 它首先将`paramName`视为准确的标签名称，
// 然后从`tagToAttrNameMap`中检索属性名称。
		attrName = tagToAttrNameMap[paramName]
		if attrName == "" {
			checkName = utils.RemoveSymbols(paramName)
// 循环查找匹配的属性名，支持大小写不敏感以及包含'-'/'_'/'.'/' '等字符的情况

// 将参数与结构体标签名称进行匹配。
// `attrKey` 是该结构体的属性名称。
			for attrKey, cmpKey := range attrToTagCheckNameMap {
				if strings.EqualFold(checkName, cmpKey) {
					attrName = attrKey
					break
				}
			}
		}

		// 将参数与结构体属性进行匹配。
		if attrName == "" {
			for attrKey, cmpKey := range attrToCheckNameMap {
// 示例：
// UserName 等价于 user_name
// User-Name 等价于 username
// username 等价于 userName
// 等等。
// 这段Go代码注释描述了不同形式的字符串表示，它们在某种上下文中被视为等价：
// - `UserName` 和 `user_name` 是等价的；
// - `User-Name` 和 `username` 也是等价的；
// - `username` 和 `userName` 同样视为等价。
// 这通常出现在将驼峰命名（camelCase）和下划线命名（snake_case）互相转换的场景中。
				if strings.EqualFold(checkName, cmpKey) {
					attrName = attrKey
					break
				}
			}
		}

		// 没有找到匹配项，放弃该属性的转换。
		if attrName == "" {
			continue
		}
		// 如果属性名称已经经过转换检查，那么跳过它。
		if _, ok := doneAttrMap[attrName]; ok {
			continue
		}
		// Mark it done.
		doneAttrMap[attrName] = struct{}{}
		if err := bindVarToStructAttr(
			pointerElemReflectValue, attrName, paramValue, paramKeyToAttrMap,
		); err != nil {
			return err
		}
	}
	return nil
}

// bindVarToStructAttr 通过名称将值设置到结构体对象的属性中。
func bindVarToStructAttr(
	structReflectValue reflect.Value,
	attrName string, value interface{}, paramKeyToAttrMap map[string]string,
) (err error) {
	structFieldValue := structReflectValue.FieldByName(attrName)
	if !structFieldValue.IsValid() {
		return nil
	}
	// CanSet 检查属性是否可公开访问。
	if !structFieldValue.CanSet() {
		return nil
	}
	defer func() {
		if exception := recover(); exception != nil {
			if err = bindVarToReflectValue(structFieldValue, value, paramKeyToAttrMap); err != nil {
				err = 错误类.X多层错误并格式化(err, `error binding value to attribute "%s"`, attrName)
			}
		}
	}()
	// 直接转换
	if empty.X是否为Nil(value) {
		structFieldValue.Set(reflect.Zero(structFieldValue.Type()))
	} else {
// 尝试调用自定义转换器。
// 问题：https://github.com/gogf/gf/issues/3099
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

// 特殊处理某些类型：
// - 覆盖标准库中time.Time类型的默认转换逻辑。
		var structFieldTypeName = structFieldValue.Type().String()
		switch structFieldTypeName {
		case "time.Time", "*time.Time":
			doConvertWithReflectValueSet(structFieldValue, doConvertInput{
				FromValue:  value,
				ToTypeName: structFieldTypeName,
				ReferValue: structFieldValue,
			})
			return
// 在递归中保持时区的一致性
// 问题：https://github.com/gogf/gf/issues/2980
		case "*gtime.Time", "gtime.Time":
			doConvertWithReflectValueSet(structFieldValue, doConvertInput{
				FromValue:  value,
				ToTypeName: structFieldTypeName,
				ReferValue: structFieldValue,
			})
			return
		}

		// 常用接口检查。
		if err, ok = bindVarToReflectValueWithInterfaceCheck(structFieldValue, value); ok {
			return err
		}

		// 默认转换。
		doConvertWithReflectValueSet(structFieldValue, doConvertInput{
			FromValue:  value,
			ToTypeName: structFieldTypeName,
			ReferValue: structFieldValue,
		})
	}
	return nil
}

// bindVarToReflectValueWithInterfaceCheck 通过通用接口检查进行绑定。
func bindVarToReflectValueWithInterfaceCheck(reflectValue reflect.Value, value interface{}) (error, bool) {
	var pointer interface{}
	if reflectValue.Kind() != reflect.Ptr && reflectValue.CanAddr() {
		reflectValueAddr := reflectValue.Addr()
		if reflectValueAddr.IsNil() || !reflectValueAddr.IsValid() {
			return nil, false
		}
		// 不是指针类型，但可以获取其地址，因此它可以被反序列化。
		pointer = reflectValue.Addr().Interface()
	} else {
		if reflectValue.IsNil() || !reflectValue.IsValid() {
			return nil, false
		}
		pointer = reflectValue.Interface()
	}
	// UnmarshalValue.
	if v, ok := pointer.(iUnmarshalValue); ok {
		return v.UnmarshalValue(value), ok
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
			return v.UnmarshalText(valueBytes), ok
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
			// 如果这不是一个有效的JSON字符串，那么会在其两边添加字符`"`以使其成为有效的JSON字符串。
			if !json.Valid(valueBytes) {
				newValueBytes := make([]byte, len(valueBytes)+2)
				newValueBytes[0] = '"'
				newValueBytes[len(newValueBytes)-1] = '"'
				copy(newValueBytes[1:], valueBytes)
				valueBytes = newValueBytes
			}
			return v.UnmarshalJSON(valueBytes), ok
		}
	}
	if v, ok := pointer.(iSet); ok {
		v.X设置值(value)
		return nil, ok
	}
	return nil, false
}

// bindVarToReflectValue 将 `value` 绑定到反射值对象 `structFieldValue`。
func bindVarToReflectValue(
	structFieldValue reflect.Value, value interface{}, paramKeyToAttrMap map[string]string,
) (err error) {
	// JSON内容转换
	err, ok := doStructWithJsonCheck(value, structFieldValue)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}

	kind := structFieldValue.Kind()
	// 通过实现`Set`接口进行转换，针对某些类型。
	switch kind {
	case reflect.Slice, reflect.Array, reflect.Ptr, reflect.Interface:
		if !structFieldValue.IsNil() {
			if v, ok := structFieldValue.Interface().(iSet); ok {
				v.X设置值(value)
				return nil
			}
		}
	}

	// 通过 kind 使用反射进行转换。
	switch kind {
	case reflect.Map:
		return doMapToMap(value, structFieldValue, paramKeyToAttrMap)

	case reflect.Struct:
		// 递归地转换结构体属性。
		if err = doStruct(value, structFieldValue, nil, ""); err != nil {
			// 注意这里存在反射转换机制。
			structFieldValue.Set(reflect.ValueOf(value).Convert(structFieldValue.Type()))
		}

// 注意，切片元素可能为结构体类型，
// 因此内部使用Struct函数进行转换。
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
						// 在将`elem`设置到数组之前，如有必要，进行指针转换。
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
				// Value为空字符串。
				if reflectValue.IsZero() {
					var elemKind = elemType.Kind()
					// 尝试查找切片元素的原始类型种类。
					if elemKind == reflect.Ptr {
						elemKind = elemType.Elem().Kind()
					}
					switch elemKind {
					case reflect.String:
						// 空字符串不能被赋值给字符串切片。
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
				// 在将`elem`设置到数组之前，如有必要，进行指针转换。
				elem = elem.Addr()
			}
			reflectArray = reflect.MakeSlice(structFieldValue.Type(), 1, 1)
			reflectArray.Index(0).Set(elem)
		}
		structFieldValue.Set(reflectArray)

	case reflect.Ptr:
		if structFieldValue.IsNil() || structFieldValue.IsZero() {
			// 如果是空指针或为空，它会创建一个新的。
			item := reflect.New(structFieldValue.Type().Elem())
			if err, ok = bindVarToReflectValueWithInterfaceCheck(item, value); ok {
				structFieldValue.Set(item)
				return err
			}
			elem := item.Elem()
			if err = bindVarToReflectValue(elem, value, paramKeyToAttrMap); err == nil {
				structFieldValue.Set(elem.Addr())
			}
		} else {
			// 非空指针，用于给它赋值。
			return bindVarToReflectValue(structFieldValue.Elem(), value, paramKeyToAttrMap)
		}

	// 它主要用于并特别处理接口的 nil 值情况。
	case reflect.Interface:
		if value == nil {
			// Specially.
			structFieldValue.Set(reflect.ValueOf((*interface{})(nil)))
		} else {
			// 注意这里存在反射转换机制。
			structFieldValue.Set(reflect.ValueOf(value).Convert(structFieldValue.Type()))
		}

	default:
		defer func() {
			if exception := recover(); exception != nil {
				err = 错误类.X创建错误码并格式化(
					错误码类.CodeInternalPanic,
					`cannot convert value "%+v" to type "%s":%+v`,
					value,
					structFieldValue.Type().String(),
					exception,
				)
			}
		}()
// 这里使用 reflect 将 `value` 转换为属性的类型，并将转换后的结果值赋给该属性。如果按照 Go 语言通常的转换规则无法进行转换，则可能会导致失败并引发 panic。
		structFieldValue.Set(reflect.ValueOf(value).Convert(structFieldValue.Type()))
	}
	return nil
}
