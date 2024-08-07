// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 效验类

import (
	"context"
	"errors"
	"reflect"
	"strings"

	gcode "github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/internal/reflection"
	gconv "github.com/888go/goframe/util/gconv"
)

func (v *Validator) doCheckMap(ctx context.Context, params interface{}) Error {
	if params == nil {
		return nil
	}
	var (
		checkRules    = make([]fieldRule, 0)
		customMessage = make(CustomMsg) // map[RuleKey]ErrorMsg.
		errorMaps     = make(map[string]map[string]error)
	)
	switch assertValue := v.rules.(type) {
	// 序列标签：[]序列标签
	// 序列中错误结果的顺序是有意义的。
	// md5:3ffc642de1ce88d6
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
					// 如果自定义消息的长度小于规则的长度，那么剩余的规则将使用默认的错误消息。
					// md5:ada20f4d064fc46a
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

			// 无序列规则：map[field]rule. md5:5142e13fc1107ce4
	case map[string]string:
		for name, rule := range assertValue {
			checkRules = append(checkRules, fieldRule{
				Name: name,
				Rule: rule,
			})
		}
	}
	inputParamMap := gconv.X取Map(params)
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

	// 该函数会递归地检查结构体，以确定其属性是否为嵌入式结构体。
	// 从父结构体中忽略输入参数映射（inputParamMap）、关联（assoc）、规则（rules）和消息（messages）。
	// md5:ac90de50afcf3ac6
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

		// 以下逻辑与CheckStruct的部分相同，但不支持序列化。 md5:98c6e4bd2e64ed87
	for _, checkRuleItem := range checkRules {
		if len(checkRuleItem.Rule) == 0 {
			continue
		}
		value = nil
		if valueItem, ok := inputParamMap[checkRuleItem.Name]; ok {
			value = valueItem
		}
				// 它在循环中检查每个规则及其值。 md5:5ab8f96747fbcec4
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
			// 仅在映射和结构体验证中：
			// 如果值为nil或空字符串，并且没有required*规则，
			// 它会清除错误消息。
			// ===========================================================
			// md5:561b52db8297e035
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
