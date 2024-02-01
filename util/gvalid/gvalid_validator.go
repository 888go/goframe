// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gvalid
import (
	"context"
	"errors"
	"reflect"
	
	"github.com/888go/goframe/i18n/gi18n"
	"github.com/888go/goframe/internal/reflection"
	"github.com/888go/goframe/internal/utils"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	)
// Validator 是用于链接操作的验证管理器。
type Validator struct {
	i18nManager                       *gi18n.Manager      // I18n 错误信息翻译管理器。
	data                              interface{}         // 验证数据，它可以是一个映射、结构体或需要进行验证的特定值。
	assoc                             interface{}         // 关联数据，通常是一个映射，用于联合验证。
	rules                             interface{}         // 自定义验证数据。
	messages                          interface{}         // 自定义验证错误消息，可以是字符串或CustomMsg类型的。
	ruleFuncMap                       map[string]RuleFunc // ruleFuncMap 用于存储当前 Validator 的自定义规则函数。
	useAssocInsteadOfObjectAttributes bool                // 使用 `assoc` 作为其验证源，而不是从 `Object` 中获取属性值。
	bail                              bool                // 在出现第一个验证错误后停止验证。
	foreach                           bool                // 它使用当前值作为数组并对其中的每个元素进行验证，用于告知接下来的校验。
	caseInsensitive                   bool                // 对于那些需要值比较的规则，提供不区分大小写的配置。
}

// New 创建并返回一个新的验证器。
func New() *Validator {
	return &Validator{
		i18nManager: gi18n.Instance(),          // 使用默认的国际化管理器。
		ruleFuncMap: make(map[string]RuleFunc), // 自定义规则函数存储映射。
	}
}

// Run 开始使用规则和消息验证给定的数据。
func (v *Validator) Run(ctx context.Context) Error {
	if v.data == nil {
		return newValidationErrorByStr(
			internalParamsErrRuleName,
			errors.New(`no data passed for validation`),
		)
	}

	originValueAndKind := reflection.OriginValueAndKind(v.data)
	switch originValueAndKind.OriginKind {
	case reflect.Map:
		isMapValidation := false
		if v.rules == nil {
			isMapValidation = true
		} else if utils.IsMap(v.rules) || utils.IsSlice(v.rules) {
			isMapValidation = true
		}
		if isMapValidation {
			return v.doCheckMap(ctx, v.data)
		}

	case reflect.Struct:
		isStructValidation := false
		if v.rules == nil {
			isStructValidation = true
		} else if utils.IsMap(v.rules) || utils.IsSlice(v.rules) {
			isStructValidation = true
		}
		if isStructValidation {
			return v.doCheckStruct(ctx, v.data)
		}
	}

	return v.doCheckValue(ctx, doCheckValueInput{
		Name:      "",
		Value:     v.data,
		ValueType: reflect.TypeOf(v.data),
		Rule:      gconv.String(v.rules),
		Messages:  v.messages,
		DataRaw:   v.assoc,
		DataMap:   gconv.Map(v.assoc),
	})
}

// Clone 创建并返回一个新的 Validator，它是当前 Validator 的浅复制。
func (v *Validator) Clone() *Validator {
	newValidator := New()
	*newValidator = *v
	return newValidator
}

// I18n 为验证器设置国际化管理器。
func (v *Validator) I18n(i18nManager *gi18n.Manager) *Validator {
	if i18nManager == nil {
		return v
	}
	newValidator := v.Clone()
	newValidator.i18nManager = i18nManager
	return newValidator
}

// Bail 设置标记，在出现第一个验证错误后停止验证。
func (v *Validator) Bail() *Validator {
	newValidator := v.Clone()
	newValidator.bail = true
	return newValidator
}

// Foreach 注释：对于当前值作为数组进行处理，并对其每个元素进行验证。
// 注意，此装饰规则仅对下一次验证规则生效一次，特别是针对单个值的验证。
func (v *Validator) Foreach() *Validator {
	newValidator := v.Clone()
	newValidator.foreach = true
	return newValidator
}

// Ci 设置标记，用于那些需要值比较的规则，实现大小写不敏感。
func (v *Validator) Ci() *Validator {
	newValidator := v.Clone()
	newValidator.caseInsensitive = true
	return newValidator
}

// Data 是一个链式操作函数，用于为当前操作设置验证数据。
func (v *Validator) Data(data interface{}) *Validator {
	if data == nil {
		return v
	}
	newValidator := v.Clone()
	newValidator.data = data
	return newValidator
}

// Assoc 是一个链式操作函数，用于为当前操作设置关联验证数据。
// 可选参数 `assoc` 通常为 map 类型，用于指定联合验证中使用的参数映射。
// 当调用该函数并传入 `assoc` 参数时，同时会将 `useAssocInsteadOfObjectAttributes` 设置为 true。
func (v *Validator) Assoc(assoc interface{}) *Validator {
	if assoc == nil {
		return v
	}
	newValidator := v.Clone()
	newValidator.assoc = assoc
	newValidator.useAssocInsteadOfObjectAttributes = true
	return newValidator
}

// Rules 是一个链式操作函数，用于为当前操作设置自定义验证规则。
func (v *Validator) Rules(rules interface{}) *Validator {
	if rules == nil {
		return v
	}
	newValidator := v.Clone()
	newValidator.rules = rules
	return newValidator
}

// Messages 是一个链式操作函数，用于为当前操作设置自定义错误消息。
// 参数 `messages` 可以是 string/[]string/map[string]string 类型。如果 `rules` 类型为 []string，则支持在错误结果中按顺序展示消息。
// 更详细的翻译：
// ```go
// Messages 函数提供链式操作功能，允许为当前执行的操作设定个性化的错误信息。
// 其参数 `messages` 的类型可以是字符串(string)、字符串切片([]string)或字符串到字符串的映射(map[string]string)。
// 特别地，当 `rules` 参数为字符串切片([]string)类型时，Messages 函数能够支持按照规则数组中的顺序，在生成的错误结果中逐条展示相应的错误消息。
func (v *Validator) Messages(messages interface{}) *Validator {
	if messages == nil {
		return v
	}
	newValidator := v.Clone()
	newValidator.messages = messages
	return newValidator
}

// RuleFunc 将一个自定义规则函数注册到当前的 Validator。
func (v *Validator) RuleFunc(rule string, f RuleFunc) *Validator {
	newValidator := v.Clone()
	newValidator.ruleFuncMap[rule] = f
	return newValidator
}

// RuleFuncMap 将多个自定义规则函数注册到当前的 Validator。
func (v *Validator) RuleFuncMap(m map[string]RuleFunc) *Validator {
	if m == nil {
		return v
	}
	newValidator := v.Clone()
	for k, v := range m {
		newValidator.ruleFuncMap[k] = v
	}
	return newValidator
}

// getCustomRuleFunc 根据指定规则获取并返回自定义规则函数。
func (v *Validator) getCustomRuleFunc(rule string) RuleFunc {
	ruleFunc := v.ruleFuncMap[rule]
	if ruleFunc == nil {
		ruleFunc = customRuleFuncMap[rule]
	}
	return ruleFunc
}

// checkRuleRequired 检查并返回给定的 `rule` 是否为必需，即使它是 nil 或为空。
func (v *Validator) checkRuleRequired(rule string) bool {
	// 默认必需的规则。
	if gstr.HasPrefix(rule, requiredRulesPrefix) {
		return true
	}
	// 所有自定义验证规则均为必填规则。
	if _, ok := customRuleFuncMap[rule]; ok {
		return true
	}
	return false
}
