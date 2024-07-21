
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Buffer size for uint32 random number.
<原文结束>

# <翻译开始>
// 用于uint32随机数的缓冲区大小。 md5:7a9c359f249cc1ff
# <翻译结束>


<原文开始>
	// bufferChan is the buffer for random bytes,
	// every item storing 4 bytes.
<原文结束>

# <翻译开始>
	// bufferChan 是用于随机字节的缓冲区，
	// 每个元素存储4个字节。
	// md5:4d20aff4f4a62f55
# <翻译结束>


<原文开始>
// asyncProducingRandomBufferBytes is a named goroutine, which uses an asynchronous goroutine
// to produce the random bytes, and a buffer chan to store the random bytes.
// So it has high performance to generate random numbers.
<原文结束>

# <翻译开始>
// asyncProducingRandomBufferBytes 是一个命名的 Goroutine，它使用异步 Goroutine
// 来生成随机字节，并通过一个缓冲通道来存储这些随机字节。
// 因此，它具有高性能地生成随机数的特性。
// md5:33995801ded223e5
# <翻译结束>


<原文开始>
			// The random buffer from system is very expensive,
			// so fully reuse the random buffer by changing
			// the step with a different number can
			// improve the performance a lot.
			// for _, step = range []int{4, 5, 6, 7} {
<原文结束>

# <翻译开始>
			// 系统提供的随机缓冲区非常昂贵，
			// 因此通过改变步长并使用不同的数字来完全重用随机缓冲区，
			// 可以大大提高性能。
			// 对于 _step 在整数切片 []int{4, 5, 6, 7} 中的每个元素：
			// md5:dc55adf04393f3bc
# <翻译结束>

