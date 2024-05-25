
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
// CallbackItem is the global callback item registered.
<原文结束>

# <翻译开始>
// CallbackItem 是已注册的全局回调项。 md5:afee0a9b376754c0
# <翻译结束>


<原文开始>
// Callback on certain metrics.
<原文结束>

# <翻译开始>
// 在特定指标上回调。 md5:1b8f919fb6f66516
# <翻译结束>


<原文开始>
// MeterOption is the option that the meter holds.
<原文结束>

# <翻译开始>
// MeterOption 是度量器所持有的选项。 md5:7cbdb72b93713b5d
# <翻译结束>


<原文开始>
// Provider is the Provider that the callback item is bound to.
<原文结束>

# <翻译开始>
// Provider 是回调项绑定到的 Provider。 md5:5ad4c5696a74f008
# <翻译结束>


<原文开始>
// RegisterCallback registers callback on certain metrics.
// A callback is bound to certain component and version, it is called when the associated metrics are read.
// Multiple callbacks on the same component and version will be called by their registered sequence.
<原文结束>

# <翻译开始>
// RegisterCallback 在特定指标上注册回调。
// 回调与特定组件和版本绑定，当关联的指标被读取时会被调用。
// 同一组件和版本上的多个回调将按照它们注册的顺序被调用。
// md5:a7b0f2e948a5cd42
# <翻译结束>


<原文开始>
// MustRegisterCallback performs as RegisterCallback, but it panics if any error occurs.
<原文结束>

# <翻译开始>
// MustRegisterCallback 类似于 RegisterCallback，但是如果发生任何错误，它会直接 panic。 md5:41b35f310c8c461d
# <翻译结束>


<原文开始>
// GetRegisteredCallbacks retrieves and returns the registered global callbacks.
// It truncates the callback slice is the callbacks are returned.
<原文结束>

# <翻译开始>
// GetRegisteredCallbacks 获取并返回已注册的全局回调函数。
// 如果返回回调函数，会截断回调函数切片。
// md5:22e1858dfd047cb0
# <翻译结束>

