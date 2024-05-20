
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
// defaultCacheExpire is the expire time for file content caching in seconds.
<原文结束>

# <翻译开始>
// defaultCacheExpire是文件内容缓存的过期时间（以秒为单位）。. md5:93f4150c6283fef8
# <翻译结束>


<原文开始>
// commandEnvKeyForCache is the configuration key for command argument or environment configuring cache expire duration.
<原文结束>

# <翻译开始>
// commandEnvKeyForCache 是用于配置缓存过期持续时间的命令行参数或环境变量的配置键。. md5:e8e411869780802b
# <翻译结束>


<原文开始>
// Default expire time for file content caching.
<原文结束>

# <翻译开始>
// 默认的文件内容缓存过期时间。. md5:848c5089a9dc23eb
# <翻译结束>


<原文开始>
// internalCache is the memory cache for internal usage.
<原文结束>

# <翻译开始>
// internalCache是内部使用的内存缓存。. md5:5cd10c891525ec8d
# <翻译结束>


<原文开始>
// GetContentsWithCache returns string content of given file by `path` from cache.
// If there's no content in the cache, it will read it from disk file specified by `path`.
// The parameter `expire` specifies the caching time for this file content in seconds.
<原文结束>

# <翻译开始>
// GetContentsWithCache 通过`path`从缓存中返回给定文件的字符串内容。如果缓存中没有内容，它将从指定的磁盘文件（由`path`提供）中读取。参数`expire`指定了该文件内容的缓存过期时间（以秒为单位）。
// md5:ee3ca4011fe59d23
# <翻译结束>


<原文开始>
// GetBytesWithCache returns []byte content of given file by `path` from cache.
// If there's no content in the cache, it will read it from disk file specified by `path`.
// The parameter `expire` specifies the caching time for this file content in seconds.
<原文结束>

# <翻译开始>
// GetBytesWithCache 通过`path`从缓存中返回给定文件的[]byte内容。
// 如果缓存中没有内容，它将从由`path`指定的磁盘文件中读取。
// 参数`expire`以秒为单位指定该文件内容的缓存时间。
// md5:8b877378627c94a2
# <翻译结束>


<原文开始>
			// Adding this `path` to gfsnotify,
			// it will clear its cache if there's any changes of the file.
<原文结束>

# <翻译开始>
// 将此`path`添加到gfsnotify，
// 如果文件有任何更改，它将清除其缓存。
// md5:d6795c29773b5d37
# <翻译结束>

