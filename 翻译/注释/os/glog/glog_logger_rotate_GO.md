
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
// rotateFileBySize rotates the current logging file according to the
// configured rotation size.
<原文结束>

# <翻译开始>
// rotateFileBySize 根据配置的旋转大小来旋转当前的日志文件。
# <翻译结束>


<原文开始>
// doRotateFile rotates the given logging file.
<原文结束>

# <翻译开始>
// doRotateFile 对给定的日志文件进行旋转。
# <翻译结束>


<原文开始>
// No backups, it then just removes the current logging file.
<原文结束>

# <翻译开始>
// 不进行备份，直接删除当前的日志文件。
# <翻译结束>


<原文开始>
// Else it creates new backup files.
<原文结束>

# <翻译开始>
// 否则，它会创建新的备份文件。
# <翻译结束>


<原文开始>
	// Rename the logging file by adding extra datetime information to microseconds, like:
	// access.log          -> access.20200326101301899002.log
	// access.20200326.log -> access.20200326.20200326101301899002.log
<原文结束>

# <翻译开始>
// 通过向日志文件名添加额外的微秒级日期时间信息进行重命名，例如：
// access.log          -> access.20200326101301899002.log
// access.20200326.log -> access.20200326.20200326101301899002.log
// 这段代码注释是说明一个功能，该功能可以将日志文件名进行重命名，并在原文件名基础上附加精确到微秒级别的日期时间戳。这样做的目的是为了方便管理和区分不同时间段的日志记录。
# <翻译结束>


<原文开始>
// rotateChecksTimely timely checks the backups expiration and the compression.
<原文结束>

# <翻译开始>
// rotateChecksTimely 定时检查备份的过期情况和压缩状态
# <翻译结束>


<原文开始>
// Checks whether file rotation not enabled.
<原文结束>

# <翻译开始>
// 检查文件旋转是否未启用。
# <翻译结束>


<原文开始>
// It here uses memory lock to guarantee the concurrent safety.
<原文结束>

# <翻译开始>
// 这里使用内存锁来保证并发安全性。
# <翻译结束>


<原文开始>
	// get file name regex pattern
	// access-{y-m-d}-test.log => access-$-test.log => access-\$-test\.log => access-(.+?)-test\.log
<原文结束>

# <翻译开始>
// 获取文件名正则表达式模式
// access-{y-m-d}-test.log => access-$-test.log => access-\$-test\.log => access-(.+?)-test\.log
// 原始格式的文件名中，{y-m-d}代表日期，将其转换为正则表达式模式
// 首先将大括号替换为美元符号($)，但在正则表达式中有特殊含义，因此需要转义为'\$'
// 然后将日期部分转换为一个可以匹配任何字符序列的非贪婪模式组".+?"
// 最终得到的正则表达式 "access-(.+?)-test\.log" 可以匹配类似于 "access-2022-01-01-test.log" 这样的文件名
# <翻译结束>


<原文开始>
	// =============================================================
	// Rotation of expired file checks.
	// =============================================================
<原文结束>

# <翻译开始>
// =============================================================
// 过期文件检查的轮转机制。
// =============================================================
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
// 这段注释表明该段Go语言代码是用于实现“旋转文件压缩”功能的。在日志处理、数据备份等场景中，当文件达到一定大小或满足特定条件时，会创建新的文件并将旧文件进行压缩，这个过程通常称为“文件旋转”（File Rotation）。本代码块可能涉及对已旋转的文件进行压缩操作。
# <翻译结束>


<原文开始>
// Eg: access.20200326101301899002.log.gz
<原文结束>

# <翻译开始>
// 示例：access.20200326101301899002.log.gz
// 这段Go语言代码注释的中文翻译如下：
// ```go
// 示例：access.20200326101301899002.log.gz
// 这是一个文件名示例，表示一个在2020年3月26日10时13分01秒创建的访问日志文件，
// 并且经过了gzip压缩。文件名中包含了时间戳信息用于标识记录的时间点。
# <翻译结束>


<原文开始>
			// Eg:
			// access.20200326101301899002.log
<原文结束>

# <翻译开始>
// 示例：
// access.20200326101301899002.log
// （该行代码为文件名注释，意为：这是一个日志文件的示例，文件名为“access”，后跟创建日期时间戳“20200326101301899002”，并以“.log”为扩展名。）
# <翻译结束>


<原文开始>
	// =============================================================
	// Backups count limitation and expiration checks.
	// =============================================================
<原文结束>

# <翻译开始>
// =============================================================
// 备份数量限制及过期检查
// =============================================================
# <翻译结束>


<原文开始>
		// Sorted by rotated/backup file mtime.
		// The older rotated/backup file is put in the head of array.
<原文结束>

# <翻译开始>
// 按照旋转/备份文件的修改时间进行排序。
// 较旧的旋转/备份文件被放在数组的头部。
# <翻译结束>


<原文开始>
// Backups expiration checking.
<原文结束>

# <翻译开始>
// 备份过期检查
# <翻译结束>


<原文开始>
// ignore backup file
<原文结束>

# <翻译开始>
// 忽略备份文件
# <翻译结束>


<原文开始>
// ignore not matching file
<原文结束>

# <翻译开始>
// 忽略不匹配的文件
# <翻译结束>


<原文开始>
// Update the files array.
<原文结束>

# <翻译开始>
// 更新文件数组。
# <翻译结束>

