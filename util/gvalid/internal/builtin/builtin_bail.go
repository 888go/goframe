// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

// RuleBail 实现了 `bail` 规则：
// 当该字段的验证失败时，停止对其他字段的验证。
//
// 格式：bail
type RuleBail struct{}

func init() {
	Register(RuleBail{})
}

func (r RuleBail) Name() string {
	return "bail"
}

func (r RuleBail) Message() string {
	return ""
}

func (r RuleBail) Run(in RunInput) error {
	return nil
}
