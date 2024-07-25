// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。 md5:c14c707c81272457

// 包gmetric提供了指标功能的接口定义和简单API。 md5:077853b1642a75e8
package gmetric

import (
	"context"
)

// MetricType是指标的类型。 md5:12c2b0ec2071549b
type MetricType string

const (
	MetricTypeCounter                 MetricType = `Counter`                 // Counter.
	MetricTypeUpDownCounter           MetricType = `UpDownCounter`           // UpDownCounter.
	MetricTypeHistogram               MetricType = `Histogram`               // Histogram.
	MetricTypeObservableCounter       MetricType = `ObservableCounter`       // ObservableCounter.
	MetricTypeObservableUpDownCounter MetricType = `ObservableUpDownCounter` // 可观察的上/下计数器。 md5:d5bb673623b35d30
	MetricTypeObservableGauge         MetricType = `ObservableGauge`         // ObservableGauge.
)

const (
	// MetricNamePattern 是用于验证指标名称的正则表达式模式。 md5:740e4cb3a2c2c55f
	MetricNamePattern = `[\w\.\-\/]`
)

// Provider 负责管理所有的 Metrics 导出。
// 注意，如果Histogram在Provider创建之前创建，那么它的桶（buckets）无法自定义。 md5:f7dc1979ef3a25f7
type Provider interface {
	// SetAsGlobal 将当前提供者设置为当前进程的全局度量提供者，这使得后续在这个Provider上创建的指标（特别是运行时创建的指标）将使用此提供者。 md5:ebeb141536d72db3
	SetAsGlobal()

	// MeterPerformer 创建并返回一个MeterPerformer，它可以生成各种类型的度量表演者。 md5:bb1b20c9299cdfc0
	MeterPerformer(config MeterOption) MeterPerformer

	// ForceFlush 强制刷新所有待处理的指标。
	//
	// 此方法会考虑 ctx 的截止时间或取消状态。在这些情况下，将返回适当的错误。但无法保证在这些情况下所有指标都已被刷新或所有资源已释放。 md5:ea6068dfe51298a9
	ForceFlush(ctx context.Context) error

	// Shutdown 关闭提供者，刷新所有待处理的指标并释放占用的计算资源。 md5:ac9c3ba411885fee
	Shutdown(ctx context.Context) error
}

// MeterPerformer 管理所有指标表演者创建。 md5:5a71d1ffbf1a7a87
type MeterPerformer interface {
	// CounterPerformer 创建并返回一个 CounterPerformer，用于执行 Counter 指标的操作。 md5:01ef2c0f7cee9b52
	CounterPerformer(name string, option MetricOption) (CounterPerformer, error)

	// UpDownCounterPerformer 创建并返回一个UpDownCounterPerformer，用于执行UpDownCounter指标的操作。 md5:8d0a2d6b5cb6c7e2
	UpDownCounterPerformer(name string, option MetricOption) (UpDownCounterPerformer, error)

	// HistogramPerformer 创建并返回一个HistogramPerformer，用于执行Histogram指标的操作。 md5:aef82f85510796c2
	HistogramPerformer(name string, option MetricOption) (HistogramPerformer, error)

	// ObservableCounterPerformer 创建并返回一个 ObservableCounterPerformer，它执行 ObservableCounter 指标所需的运算。 md5:fe300c3bfd8a3d0d
	ObservableCounterPerformer(name string, option MetricOption) (ObservableCounterPerformer, error)

	// ObservableUpDownCounterPerformer 创建并返回一个 ObservableUpDownCounterPerformer，用于执行 ObservableUpDownCounter 指标的相关操作。 md5:71a2690321ac6e11
	ObservableUpDownCounterPerformer(name string, option MetricOption) (ObservableUpDownCounterPerformer, error)

	// ObservableGaugePerformer 创建并返回一个 ObservableGaugePerformer，它执行
	// ObservableGauge 度量指标的相关操作。 md5:02fbd42914a0e2c8
	ObservableGaugePerformer(name string, option MetricOption) (ObservableGaugePerformer, error)

	// RegisterCallback 在某些指标上注册回调函数。
	// 回调函数与特定的组件和版本绑定，当关联的指标被读取时会被调用。
	// 同一个组件和版本可以注册多个回调函数，它们将按照注册的顺序被调用。 md5:89a5acee144aeb40
	RegisterCallback(callback Callback, canBeCallbackMetrics ...ObservableMetric) error
}

// MetricOption用于创建度量的基本选项。 md5:51fa64763a3de0e2
type MetricOption struct {
	// Help提供关于这个直方图的信息。
	// 这是度量的可选配置。 md5:f8de096bb00ba08e
	Help string

	// Unit 是指标值的单位。
	// 这是指标的一个可选配置。 md5:98a35ee224664140
	Unit string

	// Attributes 保存了该指标的常量键值对描述元数据。
	// 这是度量的可选配置。 md5:5610f8005466f637
	Attributes Attributes

	// Buckets 定义了观察值计数的桶。 	// 仅适用于Histogram度量。 	// 如果没有显式配置桶，Histogram度量将使用默认桶。 md5:fc23bdae2b93e65c
	Buckets []float64

	// 当指标值发生变化时调用的回调函数。仅适用于可观察的指标。
	// 如果可观察的指标没有配置Callback属性或全局回调，它将不会做任何事情。 md5:ad172c2ef3bda0a0
	Callback MetricCallback
}

// Metric 表示一个带有其元数据的单个样本值，用于导出。 md5:9886e1ff863b8aaa
type Metric interface {
	// Info 返回一个Metric的基本信息。 md5:d521e5fdeb6e591f
	Info() MetricInfo
}

// MetricInfo 导出Metric的信息。 md5:757372004c377817
type MetricInfo interface {
	Key() string                // Key 返回指标的唯一字符串键。 md5:944b693fe2dee89f
	Name() string               // Name 返回度量指标的名称。 md5:d3300c9b003bf9e9
	Help() string               // Help 返回度量指标的帮助描述。 md5:af48acbd53473872
	Unit() string               // Unit 返回度量的单位名称。 md5:985c39bd14c3f833
	Type() MetricType           // Type返回度量的类型。 md5:5ceeadaee6d55192
	Attributes() Attributes     // Attributes 返回度量指标的常量属性切片。 md5:ca84de54457cab11
	Instrument() InstrumentInfo // InstrumentInfo 返回指标的仪表盘信息。 md5:674bd5d4b4646fe0
}

// InstrumentInfo导出指标的仪器信息。 md5:4789d30f855066d1
type InstrumentInfo interface {
	Name() string    // Name返回度量的仪器名称。 md5:a45a5bfd2d0859c1
	Version() string // Version 返回指标的仪器版本。 md5:de1a566af4fc0f5a
}

// Counter是一个度量指标，它表示一个只能不断增加的单个数值。 md5:109ed10af8f638da
type Counter interface {
	Metric
	CounterPerformer
}

// CounterPerformer 为 Counter 指标执行操作。 md5:6dbc874005efa26f
type CounterPerformer interface {
	// Inc 通过1递增计数器。使用Add方法可以按任意非负值进行递增。 md5:63d13b20691c041f
	Inc(ctx context.Context, option ...Option)

	// Add将给定的值添加到计数器中。如果值小于0，则会引发恐慌。 md5:d0e547634cd9c23f
	Add(ctx context.Context, increment float64, option ...Option)
}

// UpDownCounter 是一个度量指标，表示可以向上或向下变动的单个数值。 md5:040100a8543058fb
type UpDownCounter interface {
	Metric
	UpDownCounterPerformer
}

// UpDownCounterPerformer为UpDownCounter指标执行操作。 md5:84129079abbaa297
type UpDownCounterPerformer interface {
	// Inc 通过1递增计数器。使用Add方法可以按任意非负值进行递增。 md5:63d13b20691c041f
	Inc(ctx context.Context, option ...Option)

	// Dec 减少标度器的值 by 1。使用 Sub 函数可以减少任意数值。 md5:b095db72185faf82
	Dec(ctx context.Context, option ...Option)

	// Add将给定的值添加到计数器中。如果值小于0，则会引发恐慌。 md5:d0e547634cd9c23f
	Add(ctx context.Context, increment float64, option ...Option)
}

// 直方图在可配置的静态桶中对来自事件或样本流的单个观测值进行计数（或者作为实验性Native直方图的一部分，使用动态稀疏桶，详情请见下文）。与Summary相似，它也提供了观测值的总和及观测计数。 md5:fff6786ef225e994
type Histogram interface {
	Metric
	HistogramPerformer

	// Buckets 返回Histogram的桶数组。 md5:b0cdc9def1273944
	Buckets() []float64
}

// HistogramPerformer 为 Histogram 指标执行操作。 md5:ff12684d951e55fc
type HistogramPerformer interface {
	// Record 将一个值添加到直方图中。这个值通常是正数或零。 md5:01f7e56a23c793fb
	Record(increment float64, option ...Option)
}

// ObservableCounter 是一种用于异步记录浮点数测量值的工具，
// 每次收集周期记录一次。观测操作仅在此工具的回调中进行。
// 观测到的值被认为是计数的累计总和。 md5:6e0cd2e0e8e7f991
type ObservableCounter interface {
	Metric
	ObservableCounterPerformer
}

// ObservableUpDownCounter 用于在计算操作中同步记录浮点数测量值。 md5:35e24209bc52ccd8
type ObservableUpDownCounter interface {
	Metric
	ObservableUpDownCounterPerformer
}

// ObservableGauge是一种用于异步记录每次收集周期内的浮点64位测量值的工具。对于这种工具，只在回调中进行观察。 md5:8dc864280c4807f0
type ObservableGauge interface {
	Metric
	ObservableGaugePerformer
}

type (
	// ObservableCounterPerformer 是可观察计数器的表演者。 md5:5bd18cf71c0c3331
	ObservableCounterPerformer = ObservableMetric

	// ObservableUpDownCounterPerformer 是可观测的 ObservableUpDownCounter 的执行者。 md5:4bf423dbe75e51b8
	ObservableUpDownCounterPerformer = ObservableMetric

	// `ObservableGaugePerformer` 是用于可观察 `ObservableGauge` 的执行者。 md5:c3b3497dd317d711
	ObservableGaugePerformer = ObservableMetric
)

// ObservableMetric是一种用于异步记录每次收集周期内的浮点64测量值的工具。 md5:c0d995ca6881db3b
type ObservableMetric interface {
	observable()
}

// MetricInitializer 管理Metric的初始化。
// 它在实现metric接口时被内部调用。 md5:ae21a8f190b60a04
type MetricInitializer interface {
	// Init 在创建提供者时初始化指标。
	// 它设置了真正执行操作的指标执行者。 md5:99dffc37e63f54c1
	Init(provider Provider) error
}

// PerformerExporter 导出内部的 Metric Performer。
// 它在 metric 接口的实现中被内部调用。 md5:203c3a063f4a9d3c
type PerformerExporter interface {
	// Performer导出Metric的内部实现Performer。
	// 这通常被metric实现者使用。 md5:5c2fb8da83b93435
	Performer() any
}

// MetricCallback在metric reader开始读取指标值时自动调用。 md5:fbdeeeecfd4a085e
type MetricCallback func(ctx context.Context, obs MetricObserver) error

// Callback 是一个函数，它向Meter注册并为与其关联的一组仪器做出观察。
// 参数 Observer 用于记录这些仪器的测量观察值。 md5:b00a662ef19a704e
type Callback func(ctx context.Context, obs Observer) error

// Observer 为特定初始化的 Metric 设置值。 md5:6cd07eda38c4e418
type Observer interface {
	// Observe 观察已初始化的 Metric 的值。如果观察到的 Metrics 类型为 Counter，它会将值添加到总结果中。如果观察到的 Metrics 类型为 Gauge，它会将值设置为结果。 md5:c1a20cb2ada29935
	Observe(m ObservableMetric, value float64, option ...Option)
}

// MetricObserver 设置绑定Metric的值。 md5:9a9b8c76fa6b893c
type MetricObserver interface {
	// Observe 观察已初始化的 Metric 的值。如果观察到的 Metrics 类型为 Counter，它会将值添加到总结果中。如果观察到的 Metrics 类型为 Gauge，它会将值设置为结果。 md5:c1a20cb2ada29935
	Observe(value float64, option ...Option)
}

var (
	// metrics 存储了当前包创建的所有 Metric。 md5:f312f8edd5130178
	allMetrics = make([]Metric, 0)
)

// IsEnabled 返回指标功能是否启用。 md5:0e3e632a639d3774
func IsEnabled() bool {
	return globalProvider != nil
}

// GetAllMetrics 返回当前包创建的所有指标。 md5:bcdcca5acd76b678
func GetAllMetrics() []Metric {
	return allMetrics
}
