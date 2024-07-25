// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gvalid

import (
	"context"
	"reflect"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/internal/empty"
	"github.com/gogf/gf/v2/os/gstructs"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/gogf/gf/v2/util/gutil"
)

func (v *Validator) doCheckStruct(ctx context.Context, object interface{}) Error {
	var (
		errorMaps           = make(map[string]map[string]error) // Returning error.
		fieldToAliasNameMap = make(map[string]string)           // 字段名称到别名映射的地图。 md5:cd4fee2326581f83
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

	// 此处必须使用gstructs.TagFields而不是gstructs.FieldMap，以确保错误顺序。 md5:7f18271929bfd060
	tagFields, err := gstructs.TagFields(object, structTagPriority)
	if err != nil {
		return newValidationErrorByStr(internalObjectErrRuleName, err)
	}
	// 如果没有结构体标签和验证规则，它什么也不做，快速返回。 md5:62043578db8966a6
	if len(tagFields) == 0 && v.messages == nil && isEmptyData && isEmptyAssoc {
		return nil
	}

	var (
		inputParamMap  map[string]interface{}
		checkRules     = make([]fieldRule, 0)
		nameToRuleMap  = make(map[string]string) // 只是为了内部搜索索引目的。 md5:d89a7eb74470680a
		customMessage  = make(CustomMsg)         // 自定义规则错误消息映射。 md5:1489786362b1ed0d
		checkValueData = v.assoc                 // 准备进行验证的数据。 md5:8be527aae62e9c7b
	)
	if checkValueData == nil {
		checkValueData = object
	}
	switch assertValue := v.rules.(type) {
	// 序列标签：[]序列标签
	// 序列中错误结果的顺序是有意义的。 md5:3ffc642de1ce88d6
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
					// 如果自定义消息的长度小于规则的长度，那么剩余的规则将使用默认的错误消息。 md5:ada20f4d064fc46a
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

	// 地图类型规则不支持序列。
	// 格式：map[key]rule md5:e0dce9966e6f4666
	case map[string]string:
		nameToRuleMap = assertValue
		for name, rule := range assertValue {
			checkRules = append(checkRules, fieldRule{
				Name: name,
				Rule: rule,
			})
		}
	}
	// 如果没有结构体标签和验证规则，它什么也不做，快速返回。 md5:62043578db8966a6
	if len(tagFields) == 0 && len(checkRules) == 0 && isEmptyData && isEmptyAssoc {
		return nil
	}
	// 处理输入参数映射。 md5:4e321b50a9e44d75
	if v.assoc == nil || !v.useAssocInsteadOfObjectAttributes {
		inputParamMap = make(map[string]interface{})
	} else {
		inputParamMap = gconv.Map(v.assoc)
	}
	// 检查并使用结构体别名标签扩展参数映射。 md5:d6fb47b89d8c4795
	if !v.useAssocInsteadOfObjectAttributes {
		for nameOrTag, field := range fieldMap {
			inputParamMap[nameOrTag] = field.Value.Interface()
			if nameOrTag != field.Name() {
				inputParamMap[field.Name()] = field.Value.Interface()
			}
		}
	}

	// 将自定义验证规则与结构体标签中的规则合并。
	// 自定义规则具有最高优先级，可以覆盖结构体标签中的规则。 md5:327ef56dd9382e55
	for _, field := range tagFields {
		var (
			isMeta          bool
			fieldName       = field.Name()                  // Attribute name.
			name, rule, msg = ParseTagValue(field.TagValue) // `name`与`attribute alias`不同，后者仅用于验证。 md5:01baece1f454b49d
		)
		if len(name) == 0 {
			if value, ok := fieldToAliasNameMap[fieldName]; ok {
				// 如果属性存在别名标签，它将使用属性的别名名称。 md5:6b80790d910fc981
				name = value
			} else {
				// 否则，它直接使用属性名称。 md5:4b45cf8fb32210b5
				name = fieldName
			}
		} else {
			// 它使用了验证规则中的别名名称。 md5:e12f1e225f531883
			fieldToAliasNameMap[fieldName] = name
		}
		// 这里使用别名名称扩展params映射。
		// 注意变量`name`可能是别名名称或属性名称。 md5:67115358a00d1d8c
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
				// 如果有别名名称，
				// 使用别名名称作为其键，删除字段名称键。 md5:4454688e693edd27
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
			// 输入的规则可以覆盖结构标签中的规则。 md5:2f6a125f9ce31c45
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
	// 其优先级高于 `rules` 和结构体标签。 md5:9ed0fde7c514e9ef
	if msg, ok := v.messages.(CustomMsg); ok && len(msg) > 0 {
		for k, msgName := range msg {
			if aliasName, ok := fieldToAliasNameMap[k]; ok {
				// 替换字段名称的键。 md5:b77535a09a21fa98
				customMessage[aliasName] = msgName
			} else {
				customMessage[k] = msgName
			}
		}
	}

	// 临时变量，用于存储值。 md5:5c2a7c202b7bf486
	var value interface{}

	// 它会递归地检查结构体，以确定其属性是否为结构体或结构体切片。 md5:5461f7e14c7ea93f
	for _, field := range fieldMap {
		// 没有验证接口实现检查。 md5:ac37246e66f75369
		if _, ok := field.Value.Interface().(iNoValidation); ok {
			continue
		}
		// 不进行验证字段标签检查。 md5:70499912c8f45a5d
		if _, ok := field.TagLookup(noValidationTagName); ok {
			continue
		}
		if field.IsEmbedded() {
			//嵌入的结构体的属性被视为其父结构体的直接属性。 md5:157af1f27cab37d6
			if err = v.doCheckStruct(ctx, field.Value); err != nil {
				// 它将错误合并为单个错误映射。 md5:56fe32c627a507ee
				for k, m := range err.(*validationError).errors {
					errorMaps[k] = m
				}
			}
		} else {
			// `field.TagValue`是field.Name()的别名。
			// 例如，从结构体标签`p`获取的值。 md5:0b34b40285c5eb31
			if field.TagValue != "" {
				fieldToAliasNameMap[field.Name()] = field.TagValue
			}
			switch field.OriginalKind() {
			case reflect.Map, reflect.Struct, reflect.Slice, reflect.Array:
				// 递归检查属性切片/映射。 md5:da454f5b75f8a7ba
				value = getPossibleValueFromMap(
					inputParamMap, field.Name(), fieldToAliasNameMap[field.Name()],
				)
				if empty.IsNil(value) {
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

	// 下面的逻辑与CheckMap中的一些逻辑相同，但增加了序列支持。 md5:8a807868be68e3c3
	for _, checkRuleItem := range checkRules {
		if !checkRuleItem.IsMeta {
			value = getPossibleValueFromMap(
				inputParamMap, checkRuleItem.Name, fieldToAliasNameMap[checkRuleItem.Name],
			)
		}
		// 根据映射字段类型检查空的json字符串。 md5:e4223594884df6e0
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
		// 它在循环中检查每个规则及其值。 md5:5ab8f96747fbcec4
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
			// 仅在映射和结构体验证中：
			// 如果值为nil或空字符串，并且没有required*规则，
			// 则清除错误信息。
			// ============================================================ md5:4632db837be49942
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
