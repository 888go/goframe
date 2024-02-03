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
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gvalid/internal/builtin"
)

type doCheckValueInput struct {
	Name      string                 // Name 指定参数 `value` 的名称。
	Value     interface{}            // Value 指定待验证规则的值。
	ValueType reflect.Type           // ValueType 指定值的类型，主要用于获取值的类型标识。
	Rule      string                 // Rule 指定验证规则字符串，如 "required", "required|between:1,100" 等。
	Messages  interface{}            // Messages 指定了该规则从输入参数（通常为 map 或 slice 类型）获取的自定义错误消息。
	DataRaw   interface{}            // DataRaw 指定传递给验证器的 `原始数据`，其类型可以是 map 或 struct，也可以是 nil 值。
	DataMap   map[string]interface{} // DataMap 指定了从 `dataRaw` 转换而来的映射（map）。它通常用于内部实现
// ```go
// DataMap 代表由 `dataRaw` 转化而来的数据映射，主要用于内部使用
}

// doCheckValue 对单个键值执行真正的规则验证。
func (v *Validator) doCheckValue(ctx context.Context, in doCheckValueInput) Error {
	// 如果没有验证规则，它将不做任何操作并迅速返回。
	if in.Rule == "" {
		return nil
	}
	// 它将值转换为字符串，然后进行验证。
	var (
		// 不要进行修剪操作，因为空格也是值的一部分。
		ruleErrorMap = make(map[string]error)
	)
	// 自定义错误消息处理。
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
// 处理规则中的字符' | '，
// 这使得该规则被分割为多个规则。
	ruleItems := strings.Split(strings.TrimSpace(in.Rule), "|")
	for i := 0; ; {
		array := strings.Split(ruleItems[i], ":")
		if builtin.GetRule(array[0]) == nil && v.getCustomRuleFunc(array[0]) == nil {
// ================================== 特殊规则 ==================================
// 特殊的 `regex` 和 `not-regex` 规则。
// 如果模式中包含特殊字符（如 ':'、'|' 等），则合并正则表达式模式。
// ================================== 特殊规则 ==================================
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
			results     = ruleRegex.FindStringSubmatch(ruleItems[index]) // 分割单个规则。
			ruleKey     = gstr.Trim(results[1])                          // rule key 类似于规则 "max: 6" 中的 "max"
			rulePattern = gstr.Trim(results[2])                          // rule pattern 是规则中的模式部分，例如在规则 "max:6" 中的 "6"
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

		// 忽略已标记规则的执行逻辑。
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
			// 因为此处标记了 `foreach`，所以它会将值转换为切片。
			foreachValues = gconv.Interfaces(in.Value)
			// 重置`foreach`规则，因为它只为下一条规则生效一次。
			hasForeachRule = false
		}

		for _, value := range foreachValues {
			switch {
			// 自定义验证规则。
			case customRuleFunc != nil:
				err = customRuleFunc(ctx, RuleFuncInput{
					Rule:      ruleItems[index],
					Message:   message,
					Field:     in.Name,
					ValueType: in.ValueType,
					Value:     gvar.New(value),
					Data:      gvar.New(in.DataRaw),
				})

			// 内置验证规则。
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
				// 这里永远不会执行到。
			}

			// Error handling.
			if err != nil {
				// 错误变量替换用于错误消息。
				if errMsg := err.Error(); gstr.Contains(errMsg, "{") {
					errMsg = gstr.ReplaceByMap(errMsg, map[string]string{
						"{field}":     in.Name,             // `value`的字段名称。
						"{value}":     gconv.String(value), // 当前验证中的值。
						"{pattern}":   rulePattern,         // 规则的可变部分。
						"{attribute}": in.Name,             // 与 `{field}` 相同。已被弃用。
					})
					errMsg, _ = gregex.ReplaceString(`\s{2,}`, ` `, errMsg)
					err = errors.New(errMsg)
				}
				// 该错误应包含堆栈信息以指示错误位置。
				if !gerror.HasStack(err) {
					err = gerror.NewCode(gcode.CodeValidationFailed, err.Error())
				}
				// 错误应具有错误代码 `gcode.CodeValidationFailed`。
				if gerror.Code(err) == gcode.CodeNil {
					if e, ok := err.(*gerror.Error); ok {
						e.SetCode(gcode.CodeValidationFailed)
					}
				}
				ruleErrorMap[ruleKey] = err

// 如果遇到错误且存在中断规则，
// 则不再继续验证剩余规则。
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
	Value               interface{}                 // 需要验证的值。
	Type                reflect.Type                // 需要递归验证的结构体/映射/切片类型。
	Kind                reflect.Kind                // 在接下来的switch case中，需要断言的结构体/映射/切片类型。
	ErrorMaps           map[string]map[string]error // 验证失败的错误映射。
	ResultSequenceRules *[]fieldRule                // 验证失败的规则按顺序排列。
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
		// 忽略来自父级的数据、关联、规则和消息。
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
			// 它将错误合并成单个错误映射。
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
