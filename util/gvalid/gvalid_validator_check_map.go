// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gvalid
import (
	"context"
	"errors"
	"reflect"
	"strings"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/internal/reflection"
	"github.com/888go/goframe/util/gconv"
	)

func (v *Validator) doCheckMap(ctx context.Context, params interface{}) Error {
	if params == nil {
		return nil
	}
	var (
		checkRules    = make([]fieldRule, 0)
		customMessage = make(CustomMsg) // map[规则键]错误信息。
		errorMaps     = make(map[string]map[string]error)
	)
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
				for k, ruleItem := range ruleArray {
// 如果自定义消息的长度小于规则的长度，
// 剩余的规则将使用默认错误消息。
					if len(msgArray) <= k {
						continue
					}
					if len(msgArray[k]) == 0 {
						continue
					}
					array := strings.Split(ruleItem, ":")
					if _, ok := customMessage[name]; !ok {
						customMessage[name] = make(map[string]string)
					}
					customMessage[name].(map[string]string)[strings.TrimSpace(array[0])] = strings.TrimSpace(msgArray[k])
				}
			}
			checkRules = append(checkRules, fieldRule{
				Name: name,
				Rule: rule,
			})
		}

	// 无序规则：map[field]rule
// （这段代码注释表明，该处定义了一个无序的映射关系，其中键（key）为field，值（value）为rule。在Go语言中，"map[field]rule"代表一个映射类型，其键和值分别为field类型和rule类型，且这个映射中的元素没有特定顺序。）
	case map[string]string:
		for name, rule := range assertValue {
			checkRules = append(checkRules, fieldRule{
				Name: name,
				Rule: rule,
			})
		}
	}
	inputParamMap := gconv.Map(params)
	if inputParamMap == nil {
		return newValidationErrorByStr(
			internalParamsErrRuleName,
			errors.New("invalid params type: convert to map failed"),
		)
	}
	if msg, ok := v.messages.(CustomMsg); ok && len(msg) > 0 {
		if len(customMessage) > 0 {
			for k, v := range msg {
				customMessage[k] = v
			}
		} else {
			customMessage = msg
		}
	}
	var (
		value     interface{}
		validator = v.Clone()
	)

// 它递归地检查结构体，如果其属性是一个嵌入式结构体。
// 忽略来自父级的inputParamMap、assoc、rules和messages。
	validator.assoc = nil
	validator.rules = nil
	validator.messages = nil
	for _, item := range inputParamMap {
		originTypeAndKind := reflection.OriginTypeAndKind(item)
		switch originTypeAndKind.OriginKind {
		case reflect.Map, reflect.Struct, reflect.Slice, reflect.Array:
			v.doCheckValueRecursively(ctx, doCheckValueRecursivelyInput{
				Value:     item,
				Type:      originTypeAndKind.InputType,
				Kind:      originTypeAndKind.OriginKind,
				ErrorMaps: errorMaps,
			})
		}
		// Bail feature.
		if v.bail && len(errorMaps) > 0 {
			break
		}
	}
	if v.bail && len(errorMaps) > 0 {
		return newValidationError(gcode.CodeValidationFailed, nil, errorMaps)
	}

	// 下面的逻辑与 CheckStruct 的部分功能相同，但不支持顺序检查。
	for _, checkRuleItem := range checkRules {
		if len(checkRuleItem.Rule) == 0 {
			continue
		}
		value = nil
		if valueItem, ok := inputParamMap[checkRuleItem.Name]; ok {
			value = valueItem
		}
		// 它在循环中检查每一条规则及其对应的值。
		if validatedError := v.doCheckValue(ctx, doCheckValueInput{
			Name:      checkRuleItem.Name,
			Value:     value,
			ValueType: reflect.TypeOf(value),
			Rule:      checkRuleItem.Rule,
			Messages:  customMessage[checkRuleItem.Name],
			DataRaw:   params,
			DataMap:   inputParamMap,
		}); validatedError != nil {
			_, errorItem := validatedError.FirstItem()
// ===========================================================
// 仅在map和结构体验证中：
// 如果值为nil或空字符串且没有required*规则，
// 它将清除错误消息。
// ===========================================================
			if gconv.String(value) == "" {
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
			for ruleKey, ruleError := range errorItem {
				errorMaps[checkRuleItem.Name][ruleKey] = ruleError
			}
			if v.bail {
				break
			}
		}
	}
	if len(errorMaps) > 0 {
		return newValidationError(gcode.CodeValidationFailed, checkRules, errorMaps)
	}
	return nil
}
