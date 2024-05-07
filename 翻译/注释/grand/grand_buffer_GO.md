
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// Buffer size for uint32 random number.
<原文结束>

# <翻译开始>
// Buffer大小用于uint32随机数。
# <翻译结束>


<原文开始>
	// bufferChan is the buffer for random bytes,
	// every item storing 4 bytes.
<原文结束>

# <翻译开始>
// bufferChan 是用于存储随机字节的缓冲通道，
// 每个元素存储 4 字节。
# <翻译结束>


<原文开始>
// asyncProducingRandomBufferBytes is a named goroutine, which uses an asynchronous goroutine
// to produce the random bytes, and a buffer chan to store the random bytes.
// So it has high performance to generate random numbers.
<原文结束>

# <翻译开始>
// asyncProducingRandomBufferBytes 是一个命名的 goroutine，它使用异步 goroutine 生成随机字节，并通过一个缓冲通道（chan）来存储这些随机字节。
// 因此，它在生成随机数时具有较高的性能。
# <翻译结束>


<原文开始>
			// The random buffer from system is very expensive,
			// so fully reuse the random buffer by changing
			// the step with a different number can
			// improve the performance a lot.
			// for _, step = range []int{4, 5, 6, 7} {
<原文结束>

# <翻译开始>
// 系统提供的随机缓冲区代价非常高昂，
// 因此，通过改变步进值（使用不同的数字）来充分复用随机缓冲区，
// 可以显著提升性能。
// 对以下整数切片进行遍历：[4, 5, 6, 7] {
# <翻译结束>

