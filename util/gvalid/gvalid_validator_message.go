// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gvalid

import (
	"context"

	"github.com/gogf/gf/v2/util/gvalid/internal/builtin"
)

// getErrorMessageByRule 根据指定的规则获取并返回错误信息。
// 首先从自定义消息映射中检索消息，然后检查国际化（i18n）管理器。
// 如果在自定义消息映射和国际化管理器中都未找到，则返回默认的错误消息。
// md5:2d034704bcfee175
func (v *Validator) getErrorMessageByRule(ctx context.Context, ruleKey string, customMsgMap map[string]string) string {
	content := customMsgMap[ruleKey]
	if content != "" {
		// I18n translation.
		i18nContent := v.i18nManager.GetContent(ctx, content)
		if i18nContent != "" {
			return i18nContent
		}
		return content
	}

		// 根据特定规则获取默认消息。 md5:2b6271804cddc298
	content = v.i18nManager.GetContent(ctx, ruleMessagePrefixForI18n+ruleKey)
	if content == "" {
		content = defaultErrorMessages[ruleKey]
	}
	// Builtin rule message.
	if content == "" {
		if builtinRule := builtin.GetRule(ruleKey); builtinRule != nil {
			content = builtinRule.Message()
		}
	}
		// 如果没有配置的规则消息，它将使用默认的消息。 md5:e64aad97d87b44ca
	if content == "" {
		content = v.i18nManager.GetContent(ctx, ruleMessagePrefixForI18n+internalDefaultRuleName)
	}
		// 如果没有配置的规则消息，它将使用默认的消息。 md5:e64aad97d87b44ca
	if content == "" {
		content = defaultErrorMessages[internalDefaultRuleName]
	}
	return content
}
