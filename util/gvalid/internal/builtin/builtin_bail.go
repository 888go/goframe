// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package builtin

// RuleBail 实现了 `bail` 规则：
// 当此字段的验证失败时，停止进行后续验证。
//
// 格式：bail md5:be07eaed27ee8cf4
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
