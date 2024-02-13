// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gvalid 提供了强大且实用的数据/表单验证功能。
package 效验类

import (
	"context"
	"reflect"
	"regexp"
	"strings"
	
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/util/gtag"
)

// CustomMsg 是自定义错误消息类型，
// 形如：map[field] => string|map[rule]string
// （注释解释：此代码定义了一个名为 CustomMsg 的自定义错误消息类型，它是一个映射结构，键为字段名，值可以是字符串或另一个映射结构，其中内部映射的键为规则名，值为字符串。这个类型用于根据特定字段和规则生成或存储自定义的错误消息。）
type CustomMsg = map[string]interface{}

// fieldRule 定义了特定字段的别名名称和规则字符串。
type fieldRule struct {
	Name      string       // 字段的别名名称。
	Rule      string       // Rule 字符串格式如："max:6"
	IsMeta    bool         // 这个规则是否来自gmeta.Meta，将其标记为整个结构体规则。
	FieldKind reflect.Kind // 原始的结构体字段类型，用于参数类型检查。
	FieldType reflect.Type // 结构体字段的类型，用于参数类型的检查。
}

// iNoValidation 是一个接口，用于标记当前结构体未通过 `gvalid` 包进行验证。
type iNoValidation interface {
	NoValidation()
}

const (
	singleRulePattern         = `^([\w-]+):{0,1}(.*)` // 正则表达式模式，用于单个验证规则。
	internalRulesErrRuleName  = "InvalidRules"        // 规则名称，用于内部无效规则验证错误。
	internalParamsErrRuleName = "InvalidParams"       // 内部无效参数验证错误的规则名称。
	internalObjectErrRuleName = "InvalidObject"       // 内部无效对象验证错误的规则名称。
	internalErrorMapKey       = "__InternalError__"   // 错误映射键，用于内部错误。
	internalDefaultRuleName   = "__default__"         // 如果未找到指定错误规则的国际化消息，则为此提供默认规则名称，用于国际化错误消息格式。
	ruleMessagePrefixForI18n  = "gf.gvalid.rule."     // 在i18n内容中，每一项规则配置前缀字符串。
	noValidationTagName       = gtag.NoValidation     // 对结构体属性没有验证标签名称。
	ruleNameRegex             = "regex"               // 正则规则的名称为 "regex"
	ruleNameNotRegex          = "not-regex"           // "not-regex"规则的名称
	ruleNameForeach           = "foreach"             // “foreach”规则的名称
	ruleNameBail              = "bail"                // "bail"规则的名称
	ruleNameCi                = "ci"                  // "ci"规则的名称
	emptyJsonArrayStr         = "[]"                  // 空的json字符串，用于数组类型。
	emptyJsonObjectStr        = "{}"                  // 空的JSON字符串，用于对象类型。
	requiredRulesPrefix       = "required"            // requiredRulesPrefix 指定即使值为空（nil 或空）也必须进行验证的规则前缀。
)

var (
// defaultErrorMessages 是默认的错误消息集合。
// 注意，这些消息是从 ./i18n/en/validation.toml 文件同步过来的。
	defaultErrorMessages = map[string]string{
		internalDefaultRuleName: "The {field} value `{value}` is invalid",
	}

	// structTagPriority 指定了验证标签优先级数组。
	structTagPriority = []string{gtag.Valid, gtag.ValidShort}

	// aliasNameTagPriority 指定了别名标签优先级数组。
	aliasNameTagPriority = []string{gtag.Param, gtag.ParamShort}

	// 所有内部错误键。
	internalErrKeyMap = map[string]string{
		internalRulesErrRuleName:  internalRulesErrRuleName,
		internalParamsErrRuleName: internalParamsErrRuleName,
		internalObjectErrRuleName: internalObjectErrRuleName,
	}
// 正则表达式对象，用于单个规则
// 该对象仅编译一次，且可用于重复使用。
	ruleRegex, _ = regexp.Compile(singleRulePattern)

// decorativeRuleMap 定义了所有仅作为标记规则的规则，这些规则既没有功能含义，也没有错误消息。
	decorativeRuleMap = map[string]bool{
		ruleNameForeach: true,
		ruleNameBail:    true,
		ruleNameCi:      true,
	}
)

// ParseTagValue解析一个序列标签到字段、规则和错误消息。
// 序列标签格式类似：[别名@]规则[...#消息...]
func ParseTagValue(tag string) (field, rule, msg string) {
	// Complete sequence tag.
	// Example: name@required|length:2,20|password3|same:password1#||密码强度不足|两次密码不一致
	match, _ := 正则类.X匹配文本(`\s*((\w+)\s*@){0,1}\s*([^#]+)\s*(#\s*(.*)){0,1}\s*`, tag)
	if len(match) > 5 {
		msg = strings.TrimSpace(match[5])
		rule = strings.TrimSpace(match[3])
		field = strings.TrimSpace(match[2])
	} else {
		intlog.Errorf(context.TODO(), `invalid validation tag value: %s`, tag)
	}
	return
}

// GetTags 返回验证标签。
func GetTags() []string {
	return structTagPriority
}
