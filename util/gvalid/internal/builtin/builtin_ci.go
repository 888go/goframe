// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

// RuleCi 实现了 `ci` 规则：
// 对于那些需要进行值比较（如：same, different, in, not-in 等）的规则提供不区分大小写的配置。
//
// 格式：ci
type RuleCi struct{}

func init() {
	Register(RuleCi{})
}

func (r RuleCi) Name() string {
	return "ci"
}

func (r RuleCi) Message() string {
	return ""
}

func (r RuleCi) Run(in RunInput) error {
	return nil
}
