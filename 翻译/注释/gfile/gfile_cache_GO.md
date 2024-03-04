
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
// defaultCacheExpire is the expire time for file content caching in seconds.
<原文结束>

# <翻译开始>
// defaultCacheExpire 是文件内容缓存的默认过期时间，单位为秒。
# <翻译结束>


<原文开始>
// commandEnvKeyForCache is the configuration key for command argument or environment configuring cache expire duration.
<原文结束>

# <翻译开始>
// commandEnvKeyForCache 是用于配置缓存过期时间的命令行参数或环境变量的配置键。
# <翻译结束>


<原文开始>
// Default expire time for file content caching.
<原文结束>

# <翻译开始>
// 默认的文件内容缓存过期时间。
# <翻译结束>


<原文开始>
// internalCache is the memory cache for internal usage.
<原文结束>

# <翻译开始>
// internalCache 是内部使用的内存缓存。
# <翻译结束>


<原文开始>
// GetContentsWithCache returns string content of given file by `path` from cache.
// If there's no content in the cache, it will read it from disk file specified by `path`.
// The parameter `expire` specifies the caching time for this file content in seconds.
<原文结束>

# <翻译开始>
// GetContentsWithCache 通过`path`从缓存返回指定文件的字符串内容。
// 如果缓存中没有内容，则会从由`path`指定的磁盘文件中读取内容。
// 参数`expire`指定了此文件内容在缓存中的有效期，单位为秒。
# <翻译结束>


<原文开始>
// GetBytesWithCache returns []byte content of given file by `path` from cache.
// If there's no content in the cache, it will read it from disk file specified by `path`.
// The parameter `expire` specifies the caching time for this file content in seconds.
<原文结束>

# <翻译开始>
// GetBytesWithCache 函数通过 `path` 从缓存中获取指定文件的 []byte 内容。
// 如果缓存中没有内容，会从由 `path` 指定的磁盘文件中读取内容。
// 参数 `expire` 指定了此文件内容在缓存中的有效期，单位为秒。
# <翻译结束>


<原文开始>
			// Adding this `path` to gfsnotify,
			// it will clear its cache if there's any changes of the file.
<原文结束>

# <翻译开始>
// 将此`path`添加到gfsnotify，
// 若该文件有任何变化，将会清除其缓存。
# <翻译结束>

