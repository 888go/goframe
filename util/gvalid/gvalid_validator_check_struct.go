// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gvalid
import (
	"context"
	"reflect"
	"strings"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/internal/empty"
	"github.com/888go/goframe/os/gstructs"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gmeta"
	"github.com/888go/goframe/util/gutil"
	)

func (v *Validator) doCheckStruct(ctx context.Context, object interface{}) Error {
	var (
		errorMaps           = make(map[string]map[string]error) // Returning error.
		fieldToAliasNameMap = make(map[string]string)           // 字段名称到别名名称的映射。
		resultSequenceRules = make([]fieldRule, 0)
		isEmptyData         = empty.IsEmpty(v.data)
		isEmptyAssoc        = empty.IsEmpty(v.assoc)
	)
	fieldMap, err := gstructs.FieldMap(gstructs.FieldMapInput{
		Pointer:          object,
		PriorityTagArray: aliasNameTagPriority,
		RecursiveOption:  gstructs.RecursiveOptionEmbedded,
	})
	if err != nil {
		return newValidationErrorByStr(internalObjectErrRuleName, err)
	}

	// 在这里必须使用gstructs.TagFields而不是gstructs.FieldMap，以确保错误顺序的正确性。
	tagFields, err := gstructs.TagFields(object, structTagPriority)
	if err != nil {
		return newValidationErrorByStr(internalObjectErrRuleName, err)
	}
	// 如果没有结构体标签和验证规则，它将不做任何操作并快速返回。
	if len(tagFields) == 0 && v.messages == nil && isEmptyData && isEmptyAssoc {
		return nil
	}

	var (
		inputParamMap  map[string]interface{}
		checkRules     = make([]fieldRule, 0)
		nameToRuleMap  = make(map[string]string) // 仅用于内部搜索索引目的。
		customMessage  = make(CustomMsg)         // 自定义规则错误消息映射。
		checkValueData = v.assoc                 // 准备就绪，等待验证的数据。
	)
	if checkValueData == nil {
		checkValueData = object
	}
	switch assertValue := v.rules.(type) {
// 序列标签: []序列标签
// 序列对错误结果有顺序要求。
	case []string:
		for _, tag := range assertValue {
			name, rule, msg := ParseTagValue(tag)
			if len(name) == 0 {
				continue
			}
			if len(msg) > 0 {
				var (
					msgArray  = strings.Split(msg, "|")
					ruleArray = strings.Split(rule, "|")
				)
				for k, ruleKey := range ruleArray {
// 如果自定义消息的长度小于规则的长度，
// 剩余的规则将使用默认错误消息。
					if len(msgArray) <= k {
						continue
					}
					if len(msgArray[k]) == 0 {
						continue
					}
					array := strings.Split(ruleKey, ":")
					if _, ok := customMessage[name]; !ok {
						customMessage[name] = make(map[string]string)
					}
					customMessage[name].(map[string]string)[strings.TrimSpace(array[0])] = strings.TrimSpace(msgArray[k])
				}
			}
			nameToRuleMap[name] = rule
			checkRules = append(checkRules, fieldRule{
				Name: name,
				Rule: rule,
			})
		}

// map类型规则不支持序列。
// 格式：map[key]rule
	case map[string]string:
		nameToRuleMap = assertValue
		for name, rule := range assertValue {
			checkRules = append(checkRules, fieldRule{
				Name: name,
				Rule: rule,
			})
		}
	}
	// 如果没有结构体标签和验证规则，它将不做任何操作并快速返回。
	if len(tagFields) == 0 && len(checkRules) == 0 && isEmptyData && isEmptyAssoc {
		return nil
	}
	// 输入参数映射处理。
	if v.assoc == nil || !v.useAssocInsteadOfObjectAttributes {
		inputParamMap = make(map[string]interface{})
	} else {
		inputParamMap = gconv.Map(v.assoc)
	}
	// 检查并使用结构体别名标签扩展参数映射。
	if !v.useAssocInsteadOfObjectAttributes {
		for nameOrTag, field := range fieldMap {
			inputParamMap[nameOrTag] = field.Value.Interface()
			if nameOrTag != field.Name() {
				inputParamMap[field.Name()] = field.Value.Interface()
			}
		}
	}

// 将自定义验证规则与结构体标签中的规则进行合并。
// 自定义规则具有最高优先级，可以覆盖结构体标签中的规则。
	for _, field := range tagFields {
		var (
			isMeta          bool
			fieldName       = field.Name()                  // Attribute name.
			name, rule, msg = ParseTagValue(field.TagValue) // `name`与用于验证的`attribute alias`不同。
		)
		if len(name) == 0 {
			if value, ok := fieldToAliasNameMap[fieldName]; ok {
				// 如果属性存在别名标签，则它使用该属性的别名名称。
				name = value
			} else {
				// 如果不使用属性名称直接作为键，则使用它
				name = fieldName
			}
		} else {
			// 它使用了验证规则中的别名名称。
			fieldToAliasNameMap[fieldName] = name
		}
// 这里通过别名扩展params映射。
// 注意，变量`name`可能是别名或属性名。
		if _, ok := inputParamMap[name]; !ok {
			if !v.useAssocInsteadOfObjectAttributes {
				inputParamMap[name] = field.Value.Interface()
			} else {
				if name != fieldName {
					if foundKey, foundValue := gutil.MapPossibleItemByKey(inputParamMap, fieldName); foundKey != "" {
						inputParamMap[name] = foundValue
					}
				}
			}
		}

		if _, ok := nameToRuleMap[name]; !ok {
			if _, ok = nameToRuleMap[fieldName]; ok {
// 如果存在别名名称，
// 则使用别名名称作为键，并移除字段名称键。
				nameToRuleMap[name] = nameToRuleMap[fieldName]
				delete(nameToRuleMap, fieldName)
				for index, checkRuleItem := range checkRules {
					if fieldName == checkRuleItem.Name {
						checkRuleItem.Name = name
						checkRules[index] = checkRuleItem
						break
					}
				}
			} else {
				nameToRuleMap[name] = rule
				if fieldValue := field.Value.Interface(); fieldValue != nil {
					_, isMeta = fieldValue.(gmeta.Meta)
				}
				checkRules = append(checkRules, fieldRule{
					Name:      name,
					Rule:      rule,
					IsMeta:    isMeta,
					FieldKind: field.OriginalKind(),
					FieldType: field.Type(),
				})
			}
		} else {
			// 输入的规则可以覆盖结构体标签中的规则。
			continue
		}

		if len(msg) > 0 {
			var (
				msgArray  = strings.Split(msg, "|")
				ruleArray = strings.Split(rule, "|")
			)
			for k, ruleKey := range ruleArray {
				// If length of custom messages is lesser than length of rules,
				// the rest rules use the default error messages.
				if len(msgArray) <= k {
					continue
				}
				if len(msgArray[k]) == 0 {
					continue
				}
				array := strings.Split(ruleKey, ":")
				if _, ok := customMessage[name]; !ok {
					customMessage[name] = make(map[string]string)
				}
				customMessage[name].(map[string]string)[strings.TrimSpace(array[0])] = strings.TrimSpace(msgArray[k])
			}
		}
	}

// 自定义错误消息，
// 这些错误消息具有比`rules`和结构体标签更高的优先级。
	if msg, ok := v.messages.(CustomMsg); ok && len(msg) > 0 {
		for k, msgName := range msg {
			if aliasName, ok := fieldToAliasNameMap[k]; ok {
				// 覆盖字段名称的键。
				customMessage[aliasName] = msgName
			} else {
				customMessage[k] = msgName
			}
		}
	}

	// 用于临时存储值的变量。
	var value interface{}

	// 它递归检查结构体，如果其属性是结构体/结构体切片。
	for _, field := range fieldMap {
		// 没有验证接口实现了check。
		if _, ok := field.Value.Interface().(iNoValidation); ok {
			continue
		}
		// 不进行字段标签验证检查。
		if _, ok := field.TagLookup(noValidationTagName); ok {
			continue
		}
		if field.IsEmbedded() {
			if err = v.doCheckStruct(ctx, field.Value); err != nil {
				// 它将错误合并到单个错误映射中。
				for k, m := range err.(*validationError).errors {
					errorMaps[k] = m
				}
			}
		} else {
// `field.TagValue` 是 `field.Name()` 的别名。
// 例如，来自结构体标签 `p` 中的值。
			if field.TagValue != "" {
				fieldToAliasNameMap[field.Name()] = field.TagValue
			}
			switch field.OriginalKind() {
			case reflect.Map, reflect.Struct, reflect.Slice, reflect.Array:
				// 递归检查属性切片/映射。
				value = getPossibleValueFromMap(
					inputParamMap, field.Name(), fieldToAliasNameMap[field.Name()],
				)
				if value == nil {
					switch field.Kind() {
					case reflect.Map, reflect.Ptr, reflect.Slice, reflect.Array:
						// Nothing to do.
						continue
					}
				}
				v.doCheckValueRecursively(ctx, doCheckValueRecursivelyInput{
					Value:               value,
					Kind:                field.OriginalKind(),
					Type:                field.Type().Type,
					ErrorMaps:           errorMaps,
					ResultSequenceRules: &resultSequenceRules,
				})
			}
		}
		if v.bail && len(errorMaps) > 0 {
			break
		}
	}
	if v.bail && len(errorMaps) > 0 {
		return newValidationError(gcode.CodeValidationFailed, resultSequenceRules, errorMaps)
	}

	// 下面的逻辑与 CheckMap 的部分功能相同，但增加了对序列的支持。
	for _, checkRuleItem := range checkRules {
		if !checkRuleItem.IsMeta {
			value = getPossibleValueFromMap(
				inputParamMap, checkRuleItem.Name, fieldToAliasNameMap[checkRuleItem.Name],
			)
		}
		// 根据映射字段类型检查空的json字符串。
		if value != nil {
			switch checkRuleItem.FieldKind {
			case reflect.Struct, reflect.Map:
				if gconv.String(value) == emptyJsonObjectStr {
					value = ""
				}
			case reflect.Slice, reflect.Array:
				if gconv.String(value) == emptyJsonArrayStr {
					value = ""
				}
			}
		}
		// 它在循环中检查每一条规则及其对应的值。
		if validatedError := v.doCheckValue(ctx, doCheckValueInput{
			Name:      checkRuleItem.Name,
			Value:     value,
			ValueType: checkRuleItem.FieldType,
			Rule:      checkRuleItem.Rule,
			Messages:  customMessage[checkRuleItem.Name],
			DataRaw:   checkValueData,
			DataMap:   inputParamMap,
		}); validatedError != nil {
			_, errorItem := validatedError.FirstItem()
// ============================================================
// 仅在map和struct验证中：
// 如果值为nil或空字符串且没有required*规则，
// 它将清除错误消息。
// ============================================================
			if !checkRuleItem.IsMeta && (value == nil || gconv.String(value) == "") {
				required := false
				// rule => error
				for ruleKey := range errorItem {
					if required = v.checkRuleRequired(ruleKey); required {
						break
					}
				}
				if !required {
					continue
				}
			}
			if _, ok := errorMaps[checkRuleItem.Name]; !ok {
				errorMaps[checkRuleItem.Name] = make(map[string]error)
			}
			for ruleKey, errorItemMsgMap := range errorItem {
				errorMaps[checkRuleItem.Name][ruleKey] = errorItemMsgMap
			}
			// Bail feature.
			if v.bail {
				break
			}
		}
	}
	if len(errorMaps) > 0 {
		return newValidationError(
			gcode.CodeValidationFailed,
			append(checkRules, resultSequenceRules...),
			errorMaps,
		)
	}
	return nil
}

func getPossibleValueFromMap(inputParamMap map[string]interface{}, fieldName, aliasName string) (value interface{}) {
	_, value = gutil.MapPossibleItemByKey(inputParamMap, fieldName)
	if value == nil && aliasName != "" {
		_, value = gutil.MapPossibleItemByKey(inputParamMap, aliasName)
	}
	return
}
