// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package builtin

// RuleCi 实现了 `ci` 规则：
// 为需要值比较的规则（如：相同、不同、属于、不属于等）提供了大小写不敏感的配置。
//
// 格式：ci md5:2a5b8056a85341d7
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
