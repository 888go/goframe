
<原文开始>
// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457
# <翻译结束>


<原文开始>
// Package gmetric provides interface definitions and simple api for metric feature.
<原文结束>

# <翻译开始>
// 包gmetric提供了指标功能的接口定义和简单API。. md5:077853b1642a75e8
# <翻译结束>


<原文开始>
// MetricType is the type of metric.
<原文结束>

# <翻译开始>
// MetricType是指标的类型。. md5:12c2b0ec2071549b
# <翻译结束>


<原文开始>
// ObservableUpDownCounter.
<原文结束>

# <翻译开始>
// 可观察的上/下计数器。. md5:d5bb673623b35d30
# <翻译结束>


<原文开始>
// MetricNamePattern is the regular expression pattern for validating metric name.
<原文结束>

# <翻译开始>
// MetricNamePattern 是用于验证指标名称的正则表达式模式。. md5:740e4cb3a2c2c55f
# <翻译结束>


<原文开始>
// Provider manages all Metric exporting.
// Be caution that the Histogram buckets could not be customized if the creation of the Histogram
// is before the creation of Provider.
<原文结束>

# <翻译开始>
// Provider 负责管理所有的 Metrics 导出。
// 注意，如果Histogram在Provider创建之前创建，那么它的桶（buckets）无法自定义。
// md5:f7dc1979ef3a25f7
# <翻译结束>


<原文开始>
	// SetAsGlobal sets current provider as global meter provider for current process,
	// which makes the following metrics creating on this Provider, especially the metrics created in runtime.
<原文结束>

# <翻译开始>
// SetAsGlobal 将当前提供者设置为当前进程的全局度量提供者，这使得后续在这个Provider上创建的指标（特别是运行时创建的指标）将使用此提供者。
// md5:ebeb141536d72db3
# <翻译结束>


<原文开始>
// MeterPerformer creates and returns the MeterPerformer that can produce kinds of metric Performer.
<原文结束>

# <翻译开始>
// MeterPerformer 创建并返回一个MeterPerformer，它可以生成各种类型的度量表演者。. md5:bb1b20c9299cdfc0
# <翻译结束>


<原文开始>
	// ForceFlush flushes all pending metrics.
	//
	// This method honors the deadline or cancellation of ctx. An appropriate
	// error will be returned in these situations. There is no guaranteed that all
	// metrics be flushed or all resources have been released in these situations.
<原文结束>

# <翻译开始>
// ForceFlush 强制刷新所有待处理的指标。
//
// 此方法会考虑 ctx 的截止时间或取消状态。在这些情况下，将返回适当的错误。但无法保证在这些情况下所有指标都已被刷新或所有资源已释放。
// md5:ea6068dfe51298a9
# <翻译结束>


<原文开始>
	// Shutdown shuts down the Provider flushing all pending metrics and
	// releasing any held computational resources.
<原文结束>

# <翻译开始>
// Shutdown 关闭提供者，刷新所有待处理的指标并释放占用的计算资源。
// md5:ac9c3ba411885fee
# <翻译结束>


<原文开始>
// MeterPerformer manages all Metric performers creating.
<原文结束>

# <翻译开始>
// MeterPerformer 管理所有指标表演者创建。. md5:5a71d1ffbf1a7a87
# <翻译结束>


<原文开始>
	// CounterPerformer creates and returns a CounterPerformer that performs
	// the operations for Counter metric.
<原文结束>

# <翻译开始>
// CounterPerformer 创建并返回一个 CounterPerformer，用于执行 Counter 指标的操作。
// md5:01ef2c0f7cee9b52
# <翻译结束>


<原文开始>
	// UpDownCounterPerformer creates and returns a UpDownCounterPerformer that performs
	// the operations for UpDownCounter metric.
<原文结束>

# <翻译开始>
// UpDownCounterPerformer 创建并返回一个UpDownCounterPerformer，用于执行UpDownCounter指标的操作。
// md5:8d0a2d6b5cb6c7e2
# <翻译结束>


<原文开始>
	// HistogramPerformer creates and returns a HistogramPerformer that performs
	// the operations for Histogram metric.
<原文结束>

# <翻译开始>
// HistogramPerformer 创建并返回一个HistogramPerformer，用于执行Histogram指标的操作。
// md5:aef82f85510796c2
# <翻译结束>


<原文开始>
	// ObservableCounterPerformer creates and returns an ObservableCounterPerformer that performs
	// the operations for ObservableCounter metric.
<原文结束>

# <翻译开始>
// ObservableCounterPerformer 创建并返回一个 ObservableCounterPerformer，它执行 ObservableCounter 指标所需的运算。
// md5:fe300c3bfd8a3d0d
# <翻译结束>


<原文开始>
	// ObservableUpDownCounterPerformer creates and returns an ObservableUpDownCounterPerformer that performs
	// the operations for ObservableUpDownCounter metric.
<原文结束>

# <翻译开始>
// ObservableUpDownCounterPerformer 创建并返回一个 ObservableUpDownCounterPerformer，用于执行 ObservableUpDownCounter 指标的相关操作。
// md5:71a2690321ac6e11
# <翻译结束>


<原文开始>
	// ObservableGaugePerformer creates and returns an ObservableGaugePerformer that performs
	// the operations for ObservableGauge metric.
<原文结束>

# <翻译开始>
// ObservableGaugePerformer 创建并返回一个 ObservableGaugePerformer，它执行
// ObservableGauge 度量指标的相关操作。
// md5:02fbd42914a0e2c8
# <翻译结束>


<原文开始>
	// RegisterCallback registers callback on certain metrics.
	// A callback is bound to certain component and version, it is called when the associated metrics are read.
	// Multiple callbacks on the same component and version will be called by their registered sequence.
<原文结束>

# <翻译开始>
// RegisterCallback 在某些指标上注册回调函数。
// 回调函数与特定的组件和版本绑定，当关联的指标被读取时会被调用。
// 同一个组件和版本可以注册多个回调函数，它们将按照注册的顺序被调用。
// md5:89a5acee144aeb40
# <翻译结束>


<原文开始>
// MetricOption holds the basic options for creating a metric.
<原文结束>

# <翻译开始>
// MetricOption用于创建度量的基本选项。. md5:51fa64763a3de0e2
# <翻译结束>


<原文开始>
	// Help provides information about this Histogram.
	// This is an optional configuration for a metric.
<原文结束>

# <翻译开始>
// Help提供关于这个直方图的信息。
// 这是度量的可选配置。
// md5:f8de096bb00ba08e
# <翻译结束>


<原文开始>
	// Unit is the unit for metric value.
	// This is an optional configuration for a metric.
<原文结束>

# <翻译开始>
// Unit 是指标值的单位。
// 这是指标的一个可选配置。
// md5:98a35ee224664140
# <翻译结束>


<原文开始>
	// Attributes holds the constant key-value pair description metadata for this metric.
	// This is an optional configuration for a metric.
<原文结束>

# <翻译开始>
// Attributes 保存了该指标的常量键值对描述元数据。
// 这是度量的可选配置。
// md5:5610f8005466f637
# <翻译结束>


<原文开始>
	// Buckets defines the buckets into which observations are counted.
	// For Histogram metric only.
	// A histogram metric uses default buckets if no explicit buckets configured.
<原文结束>

# <翻译开始>
// Buckets 定义了观察值计数的桶。 // 仅适用于Histogram度量。 // 如果没有显式配置桶，Histogram度量将使用默认桶。
// md5:fc23bdae2b93e65c
# <翻译结束>


<原文开始>
	// Callback function for metric, which is called when metric value changes.
	// For observable metric only.
	// If an observable metric has either Callback attribute nor global callback configured, it does nothing.
<原文结束>

# <翻译开始>
// 当指标值发生变化时调用的回调函数。仅适用于可观察的指标。
// 如果可观察的指标没有配置Callback属性或全局回调，它将不会做任何事情。
// md5:ad172c2ef3bda0a0
# <翻译结束>


<原文开始>
// Metric models a single sample value with its metadata being exported.
<原文结束>

# <翻译开始>
// Metric 表示一个带有其元数据的单个样本值，用于导出。. md5:9886e1ff863b8aaa
# <翻译结束>


<原文开始>
// Info returns the basic information of a Metric.
<原文结束>

# <翻译开始>
// Info 返回一个Metric的基本信息。. md5:d521e5fdeb6e591f
# <翻译结束>


<原文开始>
// MetricInfo exports information of the Metric.
<原文结束>

# <翻译开始>
// MetricInfo 导出Metric的信息。. md5:757372004c377817
# <翻译结束>


<原文开始>
// Key returns the unique string key of the metric.
<原文结束>

# <翻译开始>
// Key 返回指标的唯一字符串键。. md5:944b693fe2dee89f
# <翻译结束>


<原文开始>
// Name returns the name of the metric.
<原文结束>

# <翻译开始>
// Name 返回度量指标的名称。. md5:d3300c9b003bf9e9
# <翻译结束>


<原文开始>
// Help returns the help description of the metric.
<原文结束>

# <翻译开始>
// Help 返回度量指标的帮助描述。. md5:af48acbd53473872
# <翻译结束>


<原文开始>
// Unit returns the unit name of the metric.
<原文结束>

# <翻译开始>
// Unit 返回度量的单位名称。. md5:985c39bd14c3f833
# <翻译结束>


<原文开始>
// Type returns the type of the metric.
<原文结束>

# <翻译开始>
// Type返回度量的类型。. md5:5ceeadaee6d55192
# <翻译结束>


<原文开始>
// Attributes returns the constant attribute slice of the metric.
<原文结束>

# <翻译开始>
// Attributes 返回度量指标的常量属性切片。. md5:ca84de54457cab11
# <翻译结束>


<原文开始>
// InstrumentInfo returns the instrument info of the metric.
<原文结束>

# <翻译开始>
// InstrumentInfo 返回指标的仪表盘信息。. md5:674bd5d4b4646fe0
# <翻译结束>


<原文开始>
// InstrumentInfo exports the instrument information of a metric.
<原文结束>

# <翻译开始>
// InstrumentInfo导出指标的仪器信息。. md5:4789d30f855066d1
# <翻译结束>


<原文开始>
// Name returns the instrument name of the metric.
<原文结束>

# <翻译开始>
// Name返回度量的仪器名称。. md5:a45a5bfd2d0859c1
# <翻译结束>


<原文开始>
// Version returns the instrument version of the metric.
<原文结束>

# <翻译开始>
// Version 返回指标的仪器版本。. md5:de1a566af4fc0f5a
# <翻译结束>


<原文开始>
// Counter is a Metric that represents a single numerical value that can ever
// goes up.
<原文结束>

# <翻译开始>
// Counter是一个度量指标，它表示一个只能不断增加的单个数值。
// md5:109ed10af8f638da
# <翻译结束>


<原文开始>
// CounterPerformer performs operations for Counter metric.
<原文结束>

# <翻译开始>
// CounterPerformer 为 Counter 指标执行操作。. md5:6dbc874005efa26f
# <翻译结束>


<原文开始>
	// Inc increments the counter by 1. Use Add to increment it by arbitrary
	// non-negative values.
<原文结束>

# <翻译开始>
// Inc 通过1递增计数器。使用Add方法可以按任意非负值进行递增。
// md5:63d13b20691c041f
# <翻译结束>


<原文开始>
// Add adds the given value to the counter. It panics if the value is < 0.
<原文结束>

# <翻译开始>
// Add将给定的值添加到计数器中。如果值小于0，则会引发恐慌。. md5:d0e547634cd9c23f
# <翻译结束>


<原文开始>
// UpDownCounter is a Metric that represents a single numerical value that can ever
// goes up or down.
<原文结束>

# <翻译开始>
// UpDownCounter 是一个度量指标，表示可以向上或向下变动的单个数值。
// md5:040100a8543058fb
# <翻译结束>


<原文开始>
// UpDownCounterPerformer performs operations for UpDownCounter metric.
<原文结束>

# <翻译开始>
//UpDownCounterPerformer为UpDownCounter指标执行操作。. md5:84129079abbaa297
# <翻译结束>


<原文开始>
// Dec decrements the Gauge by 1. Use Sub to decrement it by arbitrary values.
<原文结束>

# <翻译开始>
// Dec 减少标度器的值 by 1。使用 Sub 函数可以减少任意数值。. md5:b095db72185faf82
# <翻译结束>


<原文开始>
// Histogram counts individual observations from an event or sample stream in
// configurable static buckets (or in dynamic sparse buckets as part of the
// experimental Native Histograms, see below for more details). Similar to a
// Summary, it also provides a sum of observations and an observation count.
<原文结束>

# <翻译开始>
// 直方图在可配置的静态桶中对来自事件或样本流的单个观测值进行计数（或者作为实验性Native直方图的一部分，使用动态稀疏桶，详情请见下文）。与Summary相似，它也提供了观测值的总和及观测计数。
// md5:fff6786ef225e994
# <翻译结束>


<原文开始>
// Buckets returns the bucket slice of the Histogram.
<原文结束>

# <翻译开始>
// Buckets 返回Histogram的桶数组。. md5:b0cdc9def1273944
# <翻译结束>


<原文开始>
// HistogramPerformer performs operations for Histogram metric.
<原文结束>

# <翻译开始>
// HistogramPerformer 为 Histogram 指标执行操作。. md5:ff12684d951e55fc
# <翻译结束>


<原文开始>
	// Record adds a single value to the histogram.
	// The value is usually positive or zero.
<原文结束>

# <翻译开始>
// Record 将一个值添加到直方图中。这个值通常是正数或零。
// md5:01f7e56a23c793fb
# <翻译结束>


<原文开始>
// ObservableCounter is an instrument used to asynchronously
// record float64 measurements once per collection cycle. Observations are only
// made within a callback for this instrument. The value observed is assumed
// the to be the cumulative sum of the count.
<原文结束>

# <翻译开始>
// ObservableCounter 是一种用于异步记录浮点数测量值的工具，
// 每次收集周期记录一次。观测操作仅在此工具的回调中进行。
// 观测到的值被认为是计数的累计总和。
// md5:6e0cd2e0e8e7f991
# <翻译结束>


<原文开始>
// ObservableUpDownCounter is used to synchronously record float64 measurements during a computational
// operation.
<原文结束>

# <翻译开始>
// ObservableUpDownCounter 用于在计算操作中同步记录浮点数测量值。
// md5:35e24209bc52ccd8
# <翻译结束>


<原文开始>
// ObservableGauge is an instrument used to asynchronously record
// instantaneous float64 measurements once per collection cycle. Observations
// are only made within a callback for this instrument.
<原文结束>

# <翻译开始>
// ObservableGauge是一种用于异步记录每次收集周期内的浮点64位测量值的工具。对于这种工具，只在回调中进行观察。
// md5:8dc864280c4807f0
# <翻译结束>


<原文开始>
// ObservableCounterPerformer is performer for observable ObservableCounter.
<原文结束>

# <翻译开始>
// ObservableCounterPerformer 是可观察计数器的表演者。. md5:5bd18cf71c0c3331
# <翻译结束>


<原文开始>
// ObservableUpDownCounterPerformer is performer for observable ObservableUpDownCounter.
<原文结束>

# <翻译开始>
// ObservableUpDownCounterPerformer 是可观测的 ObservableUpDownCounter 的执行者。. md5:4bf423dbe75e51b8
# <翻译结束>


<原文开始>
// ObservableGaugePerformer is performer for observable ObservableGauge.
<原文结束>

# <翻译开始>
// `ObservableGaugePerformer` 是用于可观察 `ObservableGauge` 的执行者。. md5:c3b3497dd317d711
# <翻译结束>


<原文开始>
// ObservableMetric is an instrument used to asynchronously record
// instantaneous float64 measurements once per collection cycle.
<原文结束>

# <翻译开始>
// ObservableMetric是一种用于异步记录每次收集周期内的浮点64测量值的工具。
// md5:c0d995ca6881db3b
# <翻译结束>


<原文开始>
// MetricInitializer manages the initialization for Metric.
// It is called internally in metric interface implements.
<原文结束>

# <翻译开始>
// MetricInitializer 管理Metric的初始化。
// 它在实现metric接口时被内部调用。
// md5:ae21a8f190b60a04
# <翻译结束>


<原文开始>
	// Init initializes the Metric in Provider creation.
	// It sets the metric performer which really takes action.
<原文结束>

# <翻译开始>
// Init 在创建提供者时初始化指标。
// 它设置了真正执行操作的指标执行者。
// md5:99dffc37e63f54c1
# <翻译结束>


<原文开始>
// PerformerExporter exports internal Performer of Metric.
// It is called internally in metric interface implements.
<原文结束>

# <翻译开始>
// PerformerExporter 导出内部的 Metric Performer。
// 它在 metric 接口的实现中被内部调用。
// md5:203c3a063f4a9d3c
# <翻译结束>


<原文开始>
	// Performer exports internal Performer of Metric.
	// This is usually used by metric implements.
<原文结束>

# <翻译开始>
// Performer导出Metric的内部实现Performer。
// 这通常被metric实现者使用。
// md5:5c2fb8da83b93435
# <翻译结束>


<原文开始>
// MetricCallback is automatically called when metric reader starts reading the metric value.
<原文结束>

# <翻译开始>
// MetricCallback在metric reader开始读取指标值时自动调用。. md5:fbdeeeecfd4a085e
# <翻译结束>


<原文开始>
// Callback is a function registered with a Meter that makes observations for
// the set of instruments it is registered with. The Observer parameter is used
// to record measurement observations for these instruments.
<原文结束>

# <翻译开始>
// Callback 是一个函数，它向Meter注册并为与其关联的一组仪器做出观察。
// 参数 Observer 用于记录这些仪器的测量观察值。
// md5:b00a662ef19a704e
# <翻译结束>


<原文开始>
// Observer sets the value for certain initialized Metric.
<原文结束>

# <翻译开始>
// Observer 为特定初始化的 Metric 设置值。. md5:6cd07eda38c4e418
# <翻译结束>


<原文开始>
	// Observe observes the value for certain initialized Metric.
	// It adds the value to total result if the observed Metrics is type of Counter.
	// It sets the value as the result if the observed Metrics is type of Gauge.
<原文结束>

# <翻译开始>
// Observe 观察已初始化的 Metric 的值。如果观察到的 Metrics 类型为 Counter，它会将值添加到总结果中。如果观察到的 Metrics 类型为 Gauge，它会将值设置为结果。
// md5:c1a20cb2ada29935
# <翻译结束>


<原文开始>
// MetricObserver sets the value for bound Metric.
<原文结束>

# <翻译开始>
// MetricObserver 设置绑定Metric的值。. md5:9a9b8c76fa6b893c
# <翻译结束>


<原文开始>
// metrics stores all created Metric by current package.
<原文结束>

# <翻译开始>
// metrics 存储了当前包创建的所有 Metric。. md5:f312f8edd5130178
# <翻译结束>


<原文开始>
// IsEnabled returns whether the metrics feature is enabled.
<原文结束>

# <翻译开始>
// IsEnabled 返回指标功能是否启用。. md5:0e3e632a639d3774
# <翻译结束>


<原文开始>
// GetAllMetrics returns all Metric that created by current package.
<原文结束>

# <翻译开始>
// GetAllMetrics 返回当前包创建的所有指标。. md5:bcdcca5acd76b678
# <翻译结束>

