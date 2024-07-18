// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457

package gmetric

// noopHistogramPerformer 是一个实现了HistogramPerformer接口但没有任何实际操作的实现者。 md5:7c57c6cdb0ed19f2
type noopHistogramPerformer struct{}

// newNoopHistogramPerformer 创建并返回一个没有真正操作的 HistogramPerformer。 md5:77291cead93f5830
func newNoopHistogramPerformer() HistogramPerformer {
	return noopHistogramPerformer{}
}

// Record 向直方图中添加一个值。该值通常为正数或零。 md5:9623642cd5abf1d5
// ff:
// noopHistogramPerformer:
// increment:
// option:
func (noopHistogramPerformer) Record(increment float64, option ...Option) {}
