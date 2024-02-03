// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gsel

type builderRoundRobin struct{}

func NewBuilderRoundRobin() Builder {
	return &builderRoundRobin{}
}

func (*builderRoundRobin) Name() string {
	return "BalancerRoundRobin"
}

func (*builderRoundRobin) Build() Selector {
	return NewSelectorRoundRobin()
}
