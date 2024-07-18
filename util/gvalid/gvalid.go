// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gvalid实现了强大且实用的数据/表单验证功能。 md5:e037cf7a2dd78c4c
package gvalid

import (
	"context"
	"reflect"
	"regexp"
	"strings"

	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/util/gtag"
)

// CustomMsg is the custom error message type,
type CustomMsg = map[string]interface{}

// fieldRule 定义了指定字段的别名名称和规则字符串。 md5:54ed236de61abbbc
type fieldRule struct {
	Name      string       // 字段的别名名称。 md5:c25526bbdb1ad925
	Rule      string       // 规则字符串格式，例如："max:6". md5:347084141516f2db
	IsMeta    bool         // 是否此规则来自gmeta.Meta，这表明它是整个结构体规则。 md5:32e5c4d3da4488ad
	FieldKind reflect.Kind // 原始的结构体字段类型，用于参数类型检查。 md5:425bcc151df78ab4
	FieldType reflect.Type // 结构体字段的类型，用于参数类型检查。 md5:04ad279a1f0f2f8d
}

// iNoValidation 是一个接口，用于标记当前结构体不被 `gvalid` 包验证。 md5:b80ca2c2ec6a9a3e
type iNoValidation interface {
	NoValidation()
}

const (
	singleRulePattern         = `^([\w-]+):{0,1}(.*)` // 单个验证规则的正则表达式模式。 md5:38a42987a367551c
	internalRulesErrRuleName  = "InvalidRules"        // 用于内部无效规则验证错误的规则名称。 md5:4f1137b2b0bcbf87
	internalParamsErrRuleName = "InvalidParams"       // 用于内部无效参数验证错误的规则名称。 md5:891768620779d9d5
	internalObjectErrRuleName = "InvalidObject"       // 内部无效对象验证错误的规则名称。 md5:d86405319c6f33d2
	internalErrorMapKey       = "__InternalError__"   // error map 中用于内部错误的键。 md5:e69cd3c42d326301
	internalDefaultRuleName   = "__default__"         // 如果为指定的错误规则找不到i18n消息，则为i18n错误消息格式设置的默认规则名称。 md5:c390e0867660c4ca
	ruleMessagePrefixForI18n  = "gf.gvalid.rule."     // i18n内容中每个规则配置的前缀字符串。 md5:0f1f87e48f3229f2
	noValidationTagName       = gtag.NoValidation     // 结构体属性缺少验证标签名称。 md5:eb3058b8dac711c4
	ruleNameRegex             = "regex"               // "regex"规则的名称. md5:4a58a8d172eb6158
	ruleNameNotRegex          = "not-regex"           // "not-regex" 规则的名称. md5:f7723458e5697b5d
	ruleNameForeach           = "foreach"             // "foreach"规则的名称. md5:3d97e3f1ec27986c
	ruleNameBail              = "bail"                // "bail"规则的名称. md5:e9d4d005416cc4b3
	ruleNameCi                = "ci"                  // the name for rule "ci"
	emptyJsonArrayStr         = "[]"                  // 空的json字符串，用于数组类型。 md5:977af1a23874089e
	emptyJsonObjectStr        = "{}"                  // 对象类型的空json字符串。 md5:5c45918837cd2fe1
	requiredRulesPrefix       = "required"            // requiredRulesPrefix 指定必须验证的规则前缀，即使值为空（nil或空字符串）。 md5:be7bfaed0613daec
)

var (
// defaultErrorMessages 是默认的错误信息。
// 注意，这些信息是从 ./i18n/en/validation.toml 文件同步而来的。
// md5:373f31d6c37a48f9
	defaultErrorMessages = map[string]string{
		internalDefaultRuleName: "The {field} value `{value}` is invalid",
	}

	// structTagPriority 指定结构体标签的验证优先级数组。 md5:f41cb86d701dc7f1
	structTagPriority = []string{gtag.Valid, gtag.ValidShort}

	// aliasNameTagPriority 指定别名标签优先级数组。 md5:8c51a2951426a6c4
	aliasNameTagPriority = []string{gtag.Param, gtag.ParamShort}

	// 所有内部错误键。 md5:d6981aa171db620c
	internalErrKeyMap = map[string]string{
		internalRulesErrRuleName:  internalRulesErrRuleName,
		internalParamsErrRuleName: internalParamsErrRuleName,
		internalObjectErrRuleName: internalObjectErrRuleName,
	}
// 单个规则的正则表达式对象
// 它仅编译一次，可重复使用。
// md5:5d3b8b54080f71ba
	ruleRegex, _ = regexp.Compile(singleRulePattern)

// decorativeRuleMap 定义了所有仅具有标记规则，既没有功能意义也没有错误信息的规则。
// md5:d98db5ea3aaff41f
	decorativeRuleMap = map[string]bool{
		ruleNameForeach: true,
		ruleNameBail:    true,
		ruleNameCi:      true,
	}
)

// ParseTagValue 解析一个序列标签到字段、规则和错误消息。
// 序列标签的格式为：[别名@]规则[...#消息...]
// md5:c1a14088e6940223
// ff:
// tag:
// field:
// rule:
// msg:
func ParseTagValue(tag string) (field, rule, msg string) {
	// Complete sequence tag.
	// Example: name@required|length:2,20|password3|same:password1#||密码强度不足|两次密码不一致
	match, _ := gregex.MatchString(`\s*((\w+)\s*@){0,1}\s*([^#]+)\s*(#\s*(.*)){0,1}\s*`, tag)
	if len(match) > 5 {
		msg = strings.TrimSpace(match[5])
		rule = strings.TrimSpace(match[3])
		field = strings.TrimSpace(match[2])
	} else {
		intlog.Errorf(context.TODO(), `invalid validation tag value: %s`, tag)
	}
	return
}

// GetTags 返回验证标签。 md5:58fb30086314fe05
// ff:
func GetTags() []string {
	return structTagPriority
}
