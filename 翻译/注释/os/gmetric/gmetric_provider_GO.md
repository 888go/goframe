
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
// 您可以从 https://github.com/gogf/gf 获取。 md5:c14c707c81272457
# <翻译结束>


<原文开始>
// GlobalProvider hold the entry for creating Meter and Metric.
// The GlobalProvider has only one function for Meter creating, which is designed for convenient usage.
<原文结束>

# <翻译开始>
// GlobalProvider 用于创建Meter和Metric的入口。GlobalProvider 只有一个用于Meter创建的函数，设计目的是方便使用。 md5:25431281857b5d86
# <翻译结束>


<原文开始>
// Meter creates and returns the Meter by given MeterOption.
<原文结束>

# <翻译开始>
// Meter 使用MeterOption创建并返回Meter。 md5:d1e5c4bca1f6e03c
# <翻译结束>


<原文开始>
// Meter hold the functions for kinds of Metric creating.
<原文结束>

# <翻译开始>
// Meter 存放了创建各种指标功能的函数。 md5:ecfd04354d47531f
# <翻译结束>


<原文开始>
// Counter creates and returns a new Counter.
<原文结束>

# <翻译开始>
	// Counter 创建并返回一个新的 Counter.. md5:84e33be2f1339329
# <翻译结束>


<原文开始>
// UpDownCounter creates and returns a new UpDownCounter.
<原文结束>

# <翻译开始>
	// UpDownCounter 创建并返回一个新的UpDownCounter.. md5:17cdda79297f292f
# <翻译结束>


<原文开始>
// Histogram creates and returns a new Histogram.
<原文结束>

# <翻译开始>
	// Histogram 创建并返回一个新的 Histogram.. md5:8a66ea5ba65143f0
# <翻译结束>


<原文开始>
// ObservableCounter creates and returns a new ObservableCounter.
<原文结束>

# <翻译开始>
	// ObservableCounter 创建并返回一个新的 ObservableCounter。 md5:1fb1055edede2f1e
# <翻译结束>


<原文开始>
// ObservableUpDownCounter creates and returns a new ObservableUpDownCounter.
<原文结束>

# <翻译开始>
	// ObservableUpDownCounter 创建并返回一个新的ObservableUpDownCounter。 md5:a7f48b253e6c2099
# <翻译结束>


<原文开始>
// ObservableGauge creates and returns a new ObservableGauge.
<原文结束>

# <翻译开始>
	// ObservableGauge 创建并返回一个新的可观察计量表。 md5:406f093f6a405dd4
# <翻译结束>


<原文开始>
	// MustCounter creates and returns a new Counter.
	// It panics if any error occurs.
<原文结束>

# <翻译开始>
	// MustCounter 创建并返回一个新的计数器。
	// 如果发生任何错误，它将引发恐慌。 md5:8ca39b864372ccfe
# <翻译结束>


<原文开始>
	// MustUpDownCounter creates and returns a new UpDownCounter.
	// It panics if any error occurs.
<原文结束>

# <翻译开始>
	// MustUpDownCounter 创建并返回一个新的UpDownCounter。
	// 如果发生任何错误，它将引发恐慌。 md5:9bb6fb57771f0266
# <翻译结束>


<原文开始>
	// MustHistogram creates and returns a new Histogram.
	// It panics if any error occurs.
<原文结束>

# <翻译开始>
	// MustHistogram 创建并返回一个新的直方图。
	// 如果发生任何错误，它将引发恐慌。 md5:fc31a9bb5a94fd34
# <翻译结束>


<原文开始>
	// MustObservableCounter creates and returns a new ObservableCounter.
	// It panics if any error occurs.
<原文结束>

# <翻译开始>
	// MustObservableCounter 创建并返回一个新的可观察计数器。如果发生任何错误，它将 panic。 md5:d12041c97b0c5aa2
# <翻译结束>


<原文开始>
	// MustObservableUpDownCounter creates and returns a new ObservableUpDownCounter.
	// It panics if any error occurs.
<原文结束>

# <翻译开始>
	// MustObservableUpDownCounter 创建并返回一个新的 ObservableUpDownCounter。
	// 如果发生任何错误，它将引发 panic。 md5:04565499baba4d44
# <翻译结束>


<原文开始>
	// MustObservableGauge creates and returns a new ObservableGauge.
	// It panics if any error occurs.
<原文结束>

# <翻译开始>
	// MustObservableGauge 创建并返回一个新的 ObservableGauge。
	// 如果发生任何错误，它将引发恐慌。 md5:f45a16dcd373e219
# <翻译结束>


<原文开始>
	// RegisterCallback registers callback on certain metrics.
	// A callback is bound to certain component and version, it is called when the associated metrics are read.
	// Multiple callbacks on the same component and version will be called by their registered sequence.
<原文结束>

# <翻译开始>
	// RegisterCallback 在某些指标上注册回调函数。
	// 回调函数与特定的组件和版本绑定，当关联的指标被读取时会被调用。
	// 同一个组件和版本可以注册多个回调函数，它们将按照注册的顺序被调用。 md5:89a5acee144aeb40
# <翻译结束>


<原文开始>
// MustRegisterCallback performs as RegisterCallback, but it panics if any error occurs.
<原文结束>

# <翻译开始>
	// MustRegisterCallback 类似于 RegisterCallback，但是如果发生任何错误，它会直接 panic。 md5:41b35f310c8c461d
# <翻译结束>


<原文开始>
// globalProvider is the provider for global usage.
<原文结束>

# <翻译开始>
	// globalProvider 是用于全局使用的提供者。 md5:6345c47ee4cbbaf6
# <翻译结束>


<原文开始>
// GetGlobalProvider retrieves the GetGlobalProvider instance.
<原文结束>

# <翻译开始>
// GetGlobalProvider 获取GetGlobalProvider实例。 md5:88505f26db856ed4
# <翻译结束>


<原文开始>
// SetGlobalProvider registers `provider` as the global Provider,
// which means the following metrics creating will be base on the global provider.
<原文结束>

# <翻译开始>
// SetGlobalProvider 将 `provider` 注册为全局提供者，
// 这意味着后续创建的指标将基于全局提供者。 md5:ca9b6936745e89a3
# <翻译结束>

