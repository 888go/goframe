
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
// rotateFileBySize rotates the current logging file according to the
// configured rotation size.
<原文结束>

# <翻译开始>
// rotateFileBySize 根据配置的旋转大小，旋转当前日志文件。
// md5:e07365ed108ab9ed
# <翻译结束>


<原文开始>
// doRotateFile rotates the given logging file.
<原文结束>

# <翻译开始>
// doRotateFile 旋转给定的日志文件。 md5:c1b732cfe2000ccf
# <翻译结束>


<原文开始>
// No backups, it then just removes the current logging file.
<原文结束>

# <翻译开始>
	// 无备份情况下，它将直接删除当前的日志文件。 md5:66cbeaeb716f06ee
# <翻译结束>


<原文开始>
// Else it creates new backup files.
<原文结束>

# <翻译开始>
	// 否则，它将创建新的备份文件。 md5:98bfbd0a3d10fcb0
# <翻译结束>


<原文开始>
	// Rename the logging file by adding extra datetime information to microseconds, like:
	// access.log          -> access.20200326101301899002.log
	// access.20200326.log -> access.20200326.20200326101301899002.log
<原文结束>

# <翻译开始>
	// 通过在日志文件名中添加额外的日期时间信息（到微秒级别），重命名日志文件，例如：
	// access.log            -> access.20200326101301899002.log
	// access.20200326.log   -> access.20200326.20200326101301899002.log
	// md5:96d2e4456a2a561d
# <翻译结束>


<原文开始>
// rotateChecksTimely timely checks the backups expiration and the compression.
<原文结束>

# <翻译开始>
// timelyChecksTimely检查备份的过期和压缩。 md5:0502efeb887ae657
# <翻译结束>


<原文开始>
// Checks whether file rotation not enabled.
<原文结束>

# <翻译开始>
	// 检查文件旋转是否未启用。 md5:22b3a5305aaec48c
# <翻译结束>


<原文开始>
// It here uses memory lock to guarantee the concurrent safety.
<原文结束>

# <翻译开始>
	// 此处使用内存锁来保证并发安全性。 md5:a621f4c111c27699
# <翻译结束>


<原文开始>
	// get file name regex pattern
	// access-{y-m-d}-test.log => access-$-test.log => access-\$-test\.log => access-(.+?)-test\.log
<原文结束>

# <翻译开始>
	// 获取文件名正则表达式模式
	// access-{yyyy-mm-dd}-test.log => access-${}-test.log => access-\$\-test\.log => access-(\w+)-test\.log
	// 
	// 这段注释说明了一个正则表达式规则，用于从文件名中提取部分。原始格式是`access-yyyy-mm-dd-test.log`，经过转换后，它首先替换`{}`为`-`（`access-yyyy-mm-dd-test.log` => `access-yyyy-mm-dd-test.log`），然后替换`-`为`\`（`access-yyyy-mm-dd-test.log` => `access-$-test.log`），再进一步替换`\`为`\`（`access-$-test.log` => `access-\$-test\.log`），最后使用正向前瞻断言匹配一个或多个任意字符但不包括`-`（`access-\$-test\.log` => `access-(.+?)-test\.log`），这样就可以匹配如`access-2021-08-31-test.log`这样的文件名。
	// md5:e9cbde6eccd06a32
# <翻译结束>


<原文开始>
	// =============================================================
	// Rotation of expired file checks.
	// =============================================================
<原文结束>

# <翻译开始>
	// =============================================================
	// 无效文件检查的旋转
	// =============================================================
	// md5:2ac41d9c8ed6dcd1
# <翻译结束>


<原文开始>
// ignore not matching file
<原文结束>

# <翻译开始>
			// 忽略不匹配的文件. md5:a1b51f5b82391575
# <翻译结束>


<原文开始>
// Update the files array.
<原文结束>

# <翻译开始>
			// 更新文件数组。 md5:eb6c80314da4cb7a
# <翻译结束>


<原文开始>
	// =============================================================
	// Rotated file compression.
	// =============================================================
<原文结束>

# <翻译开始>
	// =============================================================
	// 旋转文件压缩。
	// =============================================================
	// md5:c028a879a3e48da1
# <翻译结束>


<原文开始>
// Eg: access.20200326101301899002.log.gz
<原文结束>

# <翻译开始>
			// 例如：access.20200326101301899002.log.gz. md5:e037aa543e2a446f
# <翻译结束>


<原文开始>
			// Eg:
			// access.20200326101301899002.log
<原文结束>

# <翻译开始>
			// 示例：
			// access.20200326101301899002.log
			// 
			// 这个注释没有明确的翻译需求，因为它本身就是表示一个文件名样例，其中包含了日期和可能的访问记录序列号。如果需要解释其结构含义，可以这样翻译：
			// 
			// 示例文件名：
			// 访问日志文件，格式为"access.日期(YYYYMMDDHHMMSS).序列号.log"
			// 例如：access.2020年03月26日10时13分01秒899002序列号.log
			// md5:08ddd9e8cc49fee7
# <翻译结束>


<原文开始>
	// =============================================================
	// Backups count limitation and expiration checks.
	// =============================================================
<原文结束>

# <翻译开始>
	// =============================================================
	// 备份数量限制和过期检查。
	// =============================================================
	// md5:7edc3bfeec7fde2e
# <翻译结束>


<原文开始>
		// Sorted by rotated/backup file mtime.
		// The older rotated/backup file is put in the head of array.
<原文结束>

# <翻译开始>
		// 按照旋转/备份文件的mtime（修改时间）排序。
		// 老的旋转/备份文件被放在数组的头部。
		// md5:7ead56b6a771900f
# <翻译结束>


<原文开始>
// Backups expiration checking.
<原文结束>

# <翻译开始>
		// 备份过期检查。 md5:f974bc9ca93e7536
# <翻译结束>

