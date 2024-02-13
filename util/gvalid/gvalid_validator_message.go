// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 效验类

import (
	"context"
	
	"github.com/888go/goframe/util/gvalid/internal/builtin"
)

// getErrorMessageByRule 根据指定规则获取并返回错误消息。
// 首先从自定义消息映射中检索消息，然后检查 i18n 管理器，
// 如果在自定义消息映射和 i18n 管理器中都未找到，则返回默认错误消息。
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

	// 根据特定规则获取默认消息。
	content = v.i18nManager.GetContent(ctx, ruleMessagePrefixForI18n+ruleKey)
	if content == "" {
		content = defaultErrorMessages[ruleKey]
	}
	// 内置规则消息
	if content == "" {
		if builtinRule := builtin.GetRule(ruleKey); builtinRule != nil {
			content = builtinRule.Message()
		}
	}
	// 如果没有配置规则消息，则使用默认消息。
	if content == "" {
		content = v.i18nManager.GetContent(ctx, ruleMessagePrefixForI18n+internalDefaultRuleName)
	}
	// 如果没有配置规则消息，则使用默认消息。
	if content == "" {
		content = defaultErrorMessages[internalDefaultRuleName]
	}
	return content
}
