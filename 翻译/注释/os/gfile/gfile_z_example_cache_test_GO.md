
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
	// It reads the file content with cache duration of one minute,
	// which means it reads from cache after then without any IO operations within on minute.
<原文结束>

# <翻译开始>
// 它使用1分钟的缓存过期时间读取文件内容，
// 这意味着在接下来的一分钟内，如果没有进行任何IO操作，它将从缓存中读取。
// md5:2d9221dfe7c2f44a
# <翻译结束>


<原文开始>
// write new contents will clear its cache
<原文结束>

# <翻译开始>
// 写入新内容将清除其缓存. md5:cdefd2fa84d5ae75
# <翻译结束>


<原文开始>
// There's some delay for cache clearing after file content change.
<原文结束>

# <翻译开始>
// 文件内容更改后，清除缓存会有一些延迟。 md5:7f776df808d0e69c
# <翻译结束>


<原文开始>
	// May Output:
	// goframe example content
	// new goframe example content
<原文结束>

# <翻译开始>
	// May Output:
	// goframe example content
	// new goframe example content
# <翻译结束>

