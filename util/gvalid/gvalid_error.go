// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gvalid//bm:效验类

import (
	"strings"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
)

// Error是验证结果的错误。 md5:333865ca9d205dfa
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
	Strings() (errs []string)//qm:取文本切片  zz:Strings\(\) +\(errs +\[\]string\) *  yx:true
}

// validationError 是验证结果的验证错误。 md5:b67f2d45170f86ce
type validationError struct {
	code      gcode.Code                  // Error code.
	rules     []fieldRule                 // 按顺序的规则，仅用于保持错误顺序。 md5:865d75142a03d16d
	errors    map[string]map[string]error // 错误信息映射：map字段到map规则到消息. md5:57934a019c99d928
	firstKey  string                      // 第一个错误规则键（默认为空）。 md5:19b132d9be7a2e96
	firstItem map[string]error            // 第一个错误规则的值（默认为nil）。 md5:282d9086842ac373
}

// newValidationError 创建并返回一个验证错误。 md5:60829ca804e6f83e
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
	// 过滤重复序列规则。 md5:7a7958b11e315baa
	var ruleNameSet = gset.NewStrSet()
	for i := 0; i < len(rules); {
		if !ruleNameSet.AddIfNotExist(rules[i].Name) {
			// Delete repeated rule.
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

// newValidationErrorByStr 通过字符串创建并返回一个验证错误。 md5:f8649a2f7b8f4b7c
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

// Code 返回当前验证错误的错误代码。 md5:e3c1f143cc6ab020
// ff:
// e:
func (e *validationError) Code() gcode.Code {
	if e == nil {
		return gcode.CodeNil
	}
	return e.code
}

// Map 返回第一个错误消息作为映射。 md5:a50660d08282062c
// ff:
// e:
func (e *validationError) Map() map[string]error {
	if e == nil {
		return map[string]error{}
	}
	_, m := e.FirstItem()
	return m
}

// Maps返回所有的错误消息作为映射。 md5:3018cad54a77010b
// ff:
// e:
func (e *validationError) Maps() map[string]map[string]error {
	if e == nil {
		return nil
	}
	return e.errors
}

// Items 如果可能，按顺序检索并返回错误项数组，否则返回无序的错误项。
// md5:cb51d4d0fa07a635
// ff:
// e:
// items:
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

// FirstItem 返回第一个验证规则错误的字段名称和错误消息。 md5:f1a0ce09f39c751b
// ff:
// e:
// key:
// messages:
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

// FirstRule 返回第一个错误规则及其消息字符串。 md5:ba540411a8e82a5d
// ff:
// e:
// rule:
// err:
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

// FirstError 返回第一个错误消息作为字符串。
// 注意，如果没有错误序列，返回的消息可能会有所不同。
// md5:194a5e5551fbb1e3
// ff:
// e:
// err:
func (e *validationError) FirstError() (err error) {
	if e == nil {
		return nil
	}
	_, err = e.FirstRule()
	return
}

// Current是FirstError的别名，实现了gerror.iCurrent接口。 md5:0a09fda4e8417f2c
// ff:
// e:
func (e *validationError) Current() error {
	return e.FirstError()
}

// String 返回所有错误信息作为一个字符串，多个错误消息使用分号 ';' 连接。 md5:d6ac7d8c7c8a6a03
// ff:
// e:
func (e *validationError) String() string {
	if e == nil {
		return ""
	}
	return strings.Join(e.Strings(), "; ")
}

// Error 实现了 error 接口的 Error 方法。 md5:6b9d58fee5a72399
// ff:
// e:
func (e *validationError) Error() string {
	if e == nil {
		return ""
	}
	return e.String()
}

// Strings 将所有的错误消息返回为字符串数组。 md5:63f084a27bc91b14
// ff:
// e:
// errs:
func (e *validationError) Strings() (errs []string) {
	if e == nil {
		return []string{}
	}
	errs = make([]string, 0)
	// By sequence.
	if len(e.rules) > 0 {
		for _, v := range e.rules {
			if errorItemMap, ok := e.errors[v.Name]; ok {
				// 验证错误检查。 md5:f68965da177b50ef
				for _, ruleItem := range strings.Split(v.Rule, "|") {
					ruleItem = strings.TrimSpace(strings.Split(ruleItem, ":")[0])
					if err, ok := errorItemMap[ruleItem]; ok {
						errs = append(errs, err.Error())
					}
				}
				// internal error checks.
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
