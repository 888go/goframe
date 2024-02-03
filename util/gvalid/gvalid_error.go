// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gvalid

import (
	"strings"
	
	"github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/text/gstr"
)

// Error 是验证结果的错误信息。
type Error interface {
	Code() gcode.Code
	Current() error
	Error() string
	FirstItem() (key string, messages map[string]error)
	FirstRule() (rule string, err error)
	FirstError() (err error)
	Items() (items []map[string]map[string]error)
	Map() map[string]error
	Maps() map[string]map[string]error
	String() string
	Strings() (errs []string)
}

// validationError 是验证结果的验证错误。
type validationError struct {
	code      gcode.Code                  // Error code.
	rules     []fieldRule                 // 按照顺序排列的规则，仅用于保留错误序列。
	errors    map[string]map[string]error // 错误映射：map[字段]map[规则]消息
	firstKey  string                      // 第一条错误规则键（默认为空）
	firstItem map[string]error            // 第一条错误规则值（默认为nil）
}

// newValidationError 创建并返回一个验证错误。
func newValidationError(code gcode.Code, rules []fieldRule, fieldRuleErrorMap map[string]map[string]error) *validationError {
	for field, ruleErrorMap := range fieldRuleErrorMap {
		for rule, err := range ruleErrorMap {
			if !gerror.HasStack(err) {
				ruleErrorMap[rule] = gerror.NewWithOption(gerror.Option{
					Stack: false,
					Text:  gstr.Trim(err.Error()),
					Code:  code,
				})
			}
		}
		fieldRuleErrorMap[field] = ruleErrorMap
	}
	// 过滤重复的序列规则。
	var ruleNameSet = gset.NewStrSet()
	for i := 0; i < len(rules); {
		if !ruleNameSet.AddIfNotExist(rules[i].Name) {
			// 删除重复的规则。
			rules = append(rules[:i], rules[i+1:]...)
			continue
		}
		i++
	}
	return &validationError{
		code:   code,
		rules:  rules,
		errors: fieldRuleErrorMap,
	}
}

// newValidationErrorByStr 通过字符串创建并返回一个验证错误。
func newValidationErrorByStr(key string, err error) *validationError {
	return newValidationError(
		gcode.CodeInternalError,
		nil,
		map[string]map[string]error{
			internalErrorMapKey: {
				key: err,
			},
		},
	)
}

// Code 返回当前验证错误的错误代码。
func (e *validationError) Code() gcode.Code {
	if e == nil {
		return gcode.CodeNil
	}
	return e.code
}

// Map 返回第一个错误消息作为映射（map）。
func (e *validationError) Map() map[string]error {
	if e == nil {
		return map[string]error{}
	}
	_, m := e.FirstItem()
	return m
}

// Maps 将所有错误消息以映射形式返回。
func (e *validationError) Maps() map[string]map[string]error {
	if e == nil {
		return nil
	}
	return e.errors
}

// Items 函数尝试按顺序检索并返回错误项数组，如果无法按顺序获取，
// 则返回无特定顺序的错误项数组。
func (e *validationError) Items() (items []map[string]map[string]error) {
	if e == nil {
		return []map[string]map[string]error{}
	}
	items = make([]map[string]map[string]error, 0)
	// By sequence.
	if len(e.rules) > 0 {
		for _, v := range e.rules {
			if errorItemMap, ok := e.errors[v.Name]; ok {
				items = append(items, map[string]map[string]error{
					v.Name: errorItemMap,
				})
			}
		}
		return items
	}
	// No sequence.
	for name, errorRuleMap := range e.errors {
		items = append(items, map[string]map[string]error{
			name: errorRuleMap,
		})
	}
	return
}

// FirstItem 返回第一个验证规则错误的字段名称和错误消息。
func (e *validationError) FirstItem() (key string, messages map[string]error) {
	if e == nil {
		return "", map[string]error{}
	}
	if e.firstItem != nil {
		return e.firstKey, e.firstItem
	}
	// By sequence.
	if len(e.rules) > 0 {
		for _, v := range e.rules {
			if errorItemMap, ok := e.errors[v.Name]; ok {
				e.firstKey = v.Name
				e.firstItem = errorItemMap
				return v.Name, errorItemMap
			}
		}
	}
	// No sequence.
	for k, m := range e.errors {
		e.firstKey = k
		e.firstItem = m
		return k, m
	}
	return "", nil
}

// FirstRule 返回第一个错误规则及其消息字符串。
func (e *validationError) FirstRule() (rule string, err error) {
	if e == nil {
		return "", nil
	}
	// By sequence.
	if len(e.rules) > 0 {
		for _, v := range e.rules {
			if errorItemMap, ok := e.errors[v.Name]; ok {
				for _, ruleItem := range strings.Split(v.Rule, "|") {
					array := strings.Split(ruleItem, ":")
					ruleItem = strings.TrimSpace(array[0])
					if err, ok = errorItemMap[ruleItem]; ok {
						return ruleItem, err
					}
				}
			}
		}
	}
	// No sequence.
	for _, errorItemMap := range e.errors {
		for k, v := range errorItemMap {
			return k, v
		}
	}
	return "", nil
}

// FirstError 返回第一个错误信息作为字符串。
// 注意，如果没有顺序，返回的消息可能会不同。
func (e *validationError) FirstError() (err error) {
	if e == nil {
		return nil
	}
	_, err = e.FirstRule()
	return
}

// Current 是 FirstError 的别名，实现了 gerror.iCurrent 接口。
func (e *validationError) Current() error {
	return e.FirstError()
}

// String 将所有错误消息作为字符串返回，多个错误消息之间使用字符 ';' 连接。
func (e *validationError) String() string {
	if e == nil {
		return ""
	}
	return strings.Join(e.Strings(), "; ")
}

// Error 实现了 error 接口的 Error 方法。
func (e *validationError) Error() string {
	if e == nil {
		return ""
	}
	return e.String()
}

// Strings 将所有错误消息作为字符串数组返回。
func (e *validationError) Strings() (errs []string) {
	if e == nil {
		return []string{}
	}
	errs = make([]string, 0)
	// By sequence.
	if len(e.rules) > 0 {
		for _, v := range e.rules {
			if errorItemMap, ok := e.errors[v.Name]; ok {
				// 验证错误检查。
				for _, ruleItem := range strings.Split(v.Rule, "|") {
					ruleItem = strings.TrimSpace(strings.Split(ruleItem, ":")[0])
					if err, ok := errorItemMap[ruleItem]; ok {
						errs = append(errs, err.Error())
					}
				}
				// 内部错误检查
				for k := range internalErrKeyMap {
					if err, ok := errorItemMap[k]; ok {
						errs = append(errs, err.Error())
					}
				}
			}
		}
		return errs
	}
	// No sequence.
	for _, errorItemMap := range e.errors {
		for _, err := range errorItemMap {
			errs = append(errs, err.Error())
		}
	}
	return
}
