// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

// RuleForeach 实现了 `foreach` 规则：
// 它指示接下来的验证应将当前值视为一个数组，并对其中的每个元素进行验证。
//
// 格式：foreach
type RuleForeach struct{}

func init() {
	Register(RuleForeach{})
}

func (r RuleForeach) Name() string {
	return "foreach"
}

func (r RuleForeach) Message() string {
	return ""
}

func (r RuleForeach) Run(in RunInput) error {
	return nil
}
