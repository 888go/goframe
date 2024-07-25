// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package builtin

// RuleForeach 实现了 `foreach` 规则：
// 它告诉下一个验证使用当前值作为数组，并验证其每个元素。
//
// 格式：foreach md5:ee948d3f16a2f23d
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
