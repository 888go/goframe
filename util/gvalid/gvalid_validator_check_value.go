// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gvalid

import (
	"context"
	"errors"
	"reflect"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gvalid/internal/builtin"
)

type doCheckValueInput struct {
	Name      string                 // Name 指定了参数 `value` 的名称，它可能是参数的自定义标签名。 md5:31bb3221a3724b81
	Value     interface{}            // Value 指定要验证的规则的值。 md5:7ee2f1438ebf6afb
	ValueType reflect.Type           // ValueType 指定了值的类型，主要用于获取值类型的ID。 md5:299dde27f54dddba
	Rule      string                 // Rule 指定验证规则字符串，如 "required", "required|between:1,100" 等。 md5:2ae63fb98a26734b
	Messages  interface{}            // Messages从参数输入（通常为map或slice类型）中指定该规则的自定义错误消息。 md5:6aa72de28be7f730
	DataRaw   interface{}            // DataRaw 指定传递给Validator的`原始数据`。它可能是map/结构体类型或nil值。 md5:ea0ae3bdb176793b
	DataMap   map[string]interface{} // DataMap 指定了从 `dataRaw` 转换而来的映射。它通常在内部使用. md5:50285f5b09df4771
}

// doCheckValue 对单个键值对执行实际的规则验证。 md5:9032f66341668b1c
func (v *Validator) doCheckValue(ctx context.Context, in doCheckValueInput) Error {
	// 如果没有验证规则，它什么也不做并迅速返回。 md5:bc52d29571b990f7
	if in.Rule == "" {
		return nil
	}
	// 它将值转换为字符串，然后进行验证。 md5:2687e35bf141700c
	var (
		// 不要删除空白，因为空白也是值的一部分。 md5:149754fbc3e60837
		ruleErrorMap = make(map[string]error)
	)
	// 自定义错误消息处理。 md5:034ef969034ce61c
	var (
		msgArray     = make([]string, 0)
		customMsgMap = make(map[string]string)
	)
	switch messages := in.Messages.(type) {
	case string:
		msgArray = strings.Split(messages, "|")

	default:
		for k, message := range gconv.Map(in.Messages) {
			customMsgMap[k] = gconv.String(message)
		}
	}
	// 处理规则中的字符'|'，这使得该规则被分成多个子规则。
	// md5:11aa0a7f39f13bef
	ruleItems := strings.Split(strings.TrimSpace(in.Rule), "|")
	for i := 0; ; {
		array := strings.Split(ruleItems[i], ":")
		if builtin.GetRule(array[0]) == nil && v.getCustomRuleFunc(array[0]) == nil {
			// =========================== 特殊 ===========================
			// 对于特殊的正则表达式 (`regex`) 和非正则表达式 (`not-regex`) 规则。
			// 如果模式中包含特殊字符，如 ':' 或 '|'，则合并正则表达式模式。
			// =========================== 特殊 ===========================
			// md5:8f3bcac9a314de33
			var (
				ruleNameRegexLengthMatch    bool
				ruleNameNotRegexLengthMatch bool
			)
			if i > 0 {
				ruleItem := ruleItems[i-1]
				if len(ruleItem) >= len(ruleNameRegex) && ruleItem[:len(ruleNameRegex)] == ruleNameRegex {
					ruleNameRegexLengthMatch = true
				}
				if len(ruleItem) >= len(ruleNameNotRegex) && ruleItem[:len(ruleNameNotRegex)] == ruleNameNotRegex {
					ruleNameNotRegexLengthMatch = true
				}
			}
			if i > 0 && (ruleNameRegexLengthMatch || ruleNameNotRegexLengthMatch) {
				ruleItems[i-1] += "|" + ruleItems[i]
				ruleItems = append(ruleItems[:i], ruleItems[i+1:]...)
			} else {
				return newValidationErrorByStr(
					internalRulesErrRuleName,
					errors.New(internalRulesErrRuleName+": "+ruleItems[i]),
				)
			}
		} else {
			i++
		}
		if i == len(ruleItems) {
			break
		}
	}
	var (
		hasBailRule        = v.bail
		hasForeachRule     = v.foreach
		hasCaseInsensitive = v.caseInsensitive
	)
	for index := 0; index < len(ruleItems); {
		var (
			err         error
			results     = ruleRegex.FindStringSubmatch(ruleItems[index]) // split single rule.
			ruleKey     = gstr.Trim(results[1])                          // 规则键，如规则 "max: 6" 中的 "max". md5:b9eff8d7691a084c
			rulePattern = gstr.Trim(results[2])                          // 规则模式类似于 "6" 在规则 "max:6" 中. md5:7766c1e829f5f940
		)

		if !hasBailRule && ruleKey == ruleNameBail {
			hasBailRule = true
		}
		if !hasForeachRule && ruleKey == ruleNameForeach {
			hasForeachRule = true
		}
		if !hasCaseInsensitive && ruleKey == ruleNameCi {
			hasCaseInsensitive = true
		}

		// 忽略标记规则的执行逻辑。 md5:34f3e7a7cffba70b
		if decorativeRuleMap[ruleKey] {
			index++
			continue
		}

		if len(msgArray) > index {
			customMsgMap[ruleKey] = strings.TrimSpace(msgArray[index])
		}

		var (
			message        = v.getErrorMessageByRule(ctx, ruleKey, customMsgMap)
			customRuleFunc = v.getCustomRuleFunc(ruleKey)
			builtinRule    = builtin.GetRule(ruleKey)
			foreachValues  = []interface{}{in.Value}
		)
		if hasForeachRule {
			// 由于它标记了 `foreach`，所以它会将值转换为切片。 md5:9f599bb9b2fe0bba
			foreachValues = gconv.Interfaces(in.Value)
			// 重置 `foreach` 规则，因为它只对下一条规则生效一次。 md5:8c7dd94030559037
			hasForeachRule = false
		}

		for _, value := range foreachValues {
			switch {
			// 自定义验证规则。 md5:fbd7800af1a73578
			case customRuleFunc != nil:
				err = customRuleFunc(ctx, RuleFuncInput{
					Rule:      ruleItems[index],
					Message:   message,
					Field:     in.Name,
					ValueType: in.ValueType,
					Value:     gvar.New(value),
					Data:      gvar.New(in.DataRaw),
				})

			// 内置验证规则。 md5:4f4f87cac993a840
			case customRuleFunc == nil && builtinRule != nil:
				err = builtinRule.Run(builtin.RunInput{
					RuleKey:     ruleKey,
					RulePattern: rulePattern,
					Field:       in.Name,
					ValueType:   in.ValueType,
					Value:       gvar.New(value),
					Data:        gvar.New(in.DataRaw),
					Message:     message,
					Option: builtin.RunOption{
						CaseInsensitive: hasCaseInsensitive,
					},
				})

			default:
				// 它永远不会出现在这里。 md5:1b17e9ac7d650245
			}

			// Error handling.
			if err != nil {
				// 用于错误信息的错误变量替换。 md5:c424d98305e44662
				if errMsg := err.Error(); gstr.Contains(errMsg, "{") {
					errMsg = gstr.ReplaceByMap(errMsg, map[string]string{
						"{field}":     in.Name,             // `value` 的字段名称。 md5:c75900d2041a10e5
						"{value}":     gconv.String(value), // 当前验证的值。 md5:17abd56cedea072f
						"{pattern}":   rulePattern,         // 规则的变量部分。 md5:1463434d04a94902
						"{attribute}": in.Name,             // 与 `{field}` 相同。此用法已废弃。 md5:0ceaca304a2589af
					})
					errMsg, _ = gregex.ReplaceString(`\s{2,}`, ` `, errMsg)
					err = errors.New(errMsg)
				}
				// 错误应该包含堆栈信息，以指示错误的位置。 md5:bef4a94931ed384c
				if !gerror.HasStack(err) {
					err = gerror.NewCode(gcode.CodeValidationFailed, err.Error())
				}
				// 错误应该有错误代码，该代码为 `gcode.CodeValidationFailed`。 md5:b54af62f83c4db11
				if gerror.Code(err) == gcode.CodeNil {
					// TODO 使用接口可能更好？. md5:04cb382580755c3a
					if e, ok := err.(*gerror.Error); ok {
						e.SetCode(gcode.CodeValidationFailed)
					}
				}
				ruleErrorMap[ruleKey] = err

				// 如果存在错误并且有放弃规则，
				// 则不再继续验证剩余的规则。
				// md5:746db6c03bb62206
				if hasBailRule {
					goto CheckDone
				}
			}
		}
		index++
	}

CheckDone:
	if len(ruleErrorMap) > 0 {
		return newValidationError(
			gcode.CodeValidationFailed,
			[]fieldRule{{Name: in.Name, Rule: in.Rule}},
			map[string]map[string]error{
				in.Name: ruleErrorMap,
			},
		)
	}
	return nil
}

type doCheckValueRecursivelyInput struct {
	Value               interface{}                 // Value to be validated.
	Type                reflect.Type                // 将要递归验证的结构体/映射/切片类型。 md5:ae6984d7ba567001
	Kind                reflect.Kind                // 要在接下来的开关语句中进行断言的结构体/映射/切片类型。 md5:b683235f95d7aae1
	ErrorMaps           map[string]map[string]error // 验证失败的错误映射。 md5:e0888bbfb505d641
	ResultSequenceRules *[]fieldRule                // 依次验证失败的规则。 md5:5e8c03560ecc4a22
}

func (v *Validator) doCheckValueRecursively(ctx context.Context, in doCheckValueRecursivelyInput) {
	switch in.Kind {
	case reflect.Ptr:
		v.doCheckValueRecursively(ctx, doCheckValueRecursivelyInput{
			Value:               in.Value,
			Type:                in.Type.Elem(),
			Kind:                in.Type.Elem().Kind(),
			ErrorMaps:           in.ErrorMaps,
			ResultSequenceRules: in.ResultSequenceRules,
		})

	case reflect.Struct:
		// 忽略父级的数据、关联、规则和消息。 md5:27ad0097eee0432e
		var (
			validator           = v.Clone()
			toBeValidatedObject interface{}
		)
		if in.Type.Kind() == reflect.Ptr {
			toBeValidatedObject = reflect.New(in.Type.Elem()).Interface()
		} else {
			toBeValidatedObject = reflect.New(in.Type).Interface()
		}
		validator.assoc = nil
		validator.rules = nil
		validator.messages = nil
		if err := validator.Data(toBeValidatedObject).Assoc(in.Value).Run(ctx); err != nil {
			// 它将错误合并为单个错误映射。 md5:56fe32c627a507ee
			for k, m := range err.(*validationError).errors {
				in.ErrorMaps[k] = m
			}
			if in.ResultSequenceRules != nil {
				*in.ResultSequenceRules = append(*in.ResultSequenceRules, err.(*validationError).rules...)
			}
		}

	case reflect.Map:
		var (
			dataMap     = gconv.Map(in.Value)
			mapTypeElem = in.Type.Elem()
			mapTypeKind = mapTypeElem.Kind()
		)
		for _, item := range dataMap {
			v.doCheckValueRecursively(ctx, doCheckValueRecursivelyInput{
				Value:               item,
				Type:                mapTypeElem,
				Kind:                mapTypeKind,
				ErrorMaps:           in.ErrorMaps,
				ResultSequenceRules: in.ResultSequenceRules,
			})
			// Bail feature.
			if v.bail && len(in.ErrorMaps) > 0 {
				break
			}
		}

	case reflect.Slice, reflect.Array:
		var array []interface{}
		if gjson.Valid(in.Value) {
			array = gconv.Interfaces(gconv.Bytes(in.Value))
		} else {
			array = gconv.Interfaces(in.Value)
		}
		if len(array) == 0 {
			return
		}
		for _, item := range array {
			v.doCheckValueRecursively(ctx, doCheckValueRecursivelyInput{
				Value:               item,
				Type:                in.Type.Elem(),
				Kind:                in.Type.Elem().Kind(),
				ErrorMaps:           in.ErrorMaps,
				ResultSequenceRules: in.ResultSequenceRules,
			})
			// Bail feature.
			if v.bail && len(in.ErrorMaps) > 0 {
				break
			}
		}
	}
}
