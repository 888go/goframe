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

	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/internal/reflection"
	"github.com/gogf/gf/v2/internal/utils"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// Validator是用于链式操作的验证管理器。 md5:4554cd1e10f5c88e
type Validator struct {
	i18nManager                       *gi18n.Manager      // 用于错误消息翻译的国际化管理器。 md5:cc3a7d5d034e574f
	data                              interface{}         // 验证数据，可以是地图、结构体或待验证的某个值。 md5:e15200f8fa5aa3a2
	assoc                             interface{}         // 关联数据，通常是一个映射，用于联合验证。 md5:9888fdb467a95751
	rules                             interface{}         // 自定义验证数据。 md5:35e94ac262edfe24
	messages                          interface{}         // 自定义验证错误消息，可以是字符串或CustomMsg类型。 md5:c3507018b9e0da11
	ruleFuncMap                       map[string]RuleFunc // ruleFuncMap 存储当前验证器的自定义规则函数。 md5:e2e248128b108117
	useAssocInsteadOfObjectAttributes bool                // 使用`assoc`作为验证源，而不是来自`Object`的属性值。 md5:2ecc0aebe1d9f9e0
	bail                              bool                // 在第一次验证错误后停止验证。 md5:78c177ee4a5553f1
	foreach                           bool                // 它使用当前值作为数组，并验证其每个元素，以便进行下一个验证。 md5:84d43a2805a14d90
	caseInsensitive                   bool                // 大小写不敏感配置，适用于那些需要进行值比较的规则。 md5:12e4998422cd3091
}

// New 创建并返回一个新的Validator.. md5:cca3c6d267bf0323
// ff:
func New() *Validator {
	return &Validator{
		i18nManager: gi18n.Instance(),          // 使用默认的国际化管理器。 md5:89cb0f7e25a6ca81
		ruleFuncMap: make(map[string]RuleFunc), // 自定义规则函数，用于存储映射。 md5:ac4fbe8b4302ecf3
	}
}

// Run 开始根据规则和消息验证给定的数据。 md5:4345968979b93f1e
// ff:
// v:
// ctx:
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

// Clone 创建并返回一个新的Validator，它是当前对象的浅拷贝。 md5:3524ef480b75393c
// ff:
// v:
func (v *Validator) Clone() *Validator {
	newValidator := New()
	*newValidator = *v
	return newValidator
}

// I18n 设置验证器的i18n管理器。 md5:aeb8eebb20995b34
// ff:
// v:
// i18nManager:
func (v *Validator) I18n(i18nManager *gi18n.Manager) *Validator {
	if i18nManager == nil {
		return v
	}
	newValidator := v.Clone()
	newValidator.i18nManager = i18nManager
	return newValidator
}

// Bail设置在遇到第一个验证错误后停止验证的标记。 md5:219188161ae03b77
// ff:
// v:
func (v *Validator) Bail() *Validator {
	newValidator := v.Clone()
	newValidator.bail = true
	return newValidator
}

// Foreach 通知下一个验证器将当前值作为数组对待，并验证它的每个元素。
// 注意，此装饰规则仅对下一个验证规则生效一次，特别适用于单值验证。
// md5:59e49ab195827b14
// ff:
// v:
func (v *Validator) Foreach() *Validator {
	newValidator := v.Clone()
	newValidator.foreach = true
	return newValidator
}

// Ci 设置标记，表示对于需要值比较的规则进行不区分大小写的处理。 md5:a248130276497a1f
// ff:
// v:
func (v *Validator) Ci() *Validator {
	newValidator := v.Clone()
	newValidator.caseInsensitive = true
	return newValidator
}

// Data是一个链式操作函数，为当前操作设置验证数据。 md5:4bbfa1bb8271d34e
// ff:
// v:
// data:
func (v *Validator) Data(data interface{}) *Validator {
	if data == nil {
		return v
	}
	newValidator := v.Clone()
	newValidator.data = data
	return newValidator
}

// Assoc是一个链式操作函数，为当前操作设置关联验证数据。
// 可选参数`assoc`通常类型为map，用于指定并联合验证时使用的参数映射。
// 使用带有`assoc`调用此函数也会将`useAssocInsteadOfObjectAttributes`设置为true。
// md5:45823829185f6ad6
// ff:
// v:
// assoc:
func (v *Validator) Assoc(assoc interface{}) *Validator {
	if assoc == nil {
		return v
	}
	newValidator := v.Clone()
	newValidator.assoc = assoc
	newValidator.useAssocInsteadOfObjectAttributes = true
	return newValidator
}

// Rules 是一个链接操作函数，用于为当前操作设置自定义验证规则。 md5:20d3aa2d271b3575
// ff:
// v:
// rules:
func (v *Validator) Rules(rules interface{}) *Validator {
	if rules == nil {
		return v
	}
	newValidator := v.Clone()
	newValidator.rules = rules
	return newValidator
}

// Messages 是一个链式操作函数，用于为当前操作设置自定义错误消息。
// 参数 `messages` 可以为 string/[]string/map[string]string 类型。如果 `rules` 类型为 []string，它支持在错误结果中按顺序显示消息。
// md5:442bfbf7d1878c37
// ff:
// v:
// messages:
func (v *Validator) Messages(messages interface{}) *Validator {
	if messages == nil {
		return v
	}
	newValidator := v.Clone()
	newValidator.messages = messages
	return newValidator
}

// RuleFunc将一个自定义规则函数注册到当前Validator。 md5:3733cab7b3035ce3
// ff:
// v:
// rule:
// f:
func (v *Validator) RuleFunc(rule string, f RuleFunc) *Validator {
	newValidator := v.Clone()
	newValidator.ruleFuncMap[rule] = f
	return newValidator
}

// RuleFuncMap 将多个自定义规则函数注册到当前Validator。 md5:38d8a4ac760a431a
// ff:
// v:
// m:
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

// getCustomRuleFunc 获取并返回指定规则的自定义规则函数。 md5:8a82be67553011ed
func (v *Validator) getCustomRuleFunc(rule string) RuleFunc {
	ruleFunc := v.ruleFuncMap[rule]
	if ruleFunc == nil {
		ruleFunc = customRuleFuncMap[rule]
	}
	return ruleFunc
}

// checkRuleRequired 检查并返回给定的 `rule` 是否即使为 nil 或空，也是必需的。 md5:8dd4a95af0752f7f
func (v *Validator) checkRuleRequired(rule string) bool {
	// 默认所需的规则。 md5:7047f401aaa9d537
	if gstr.HasPrefix(rule, requiredRulesPrefix) {
		return true
	}
	// 所有自定义验证规则都是必填规则。 md5:58545e43bcc00d45
	if _, ok := customRuleFuncMap[rule]; ok {
		return true
	}
	return false
}
