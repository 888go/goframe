
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
// DefaultReadBuffer is the buffer size for reading file content.
<原文结束>

# <翻译开始>
// DefaultReadBuffer 是用于读取文件内容的缓冲区大小。
# <翻译结束>


<原文开始>
// GetContents returns the file content of `path` as string.
// It returns en empty string if it fails reading.
<原文结束>

# <翻译开始>
// GetContents 函数返回字符串形式的 `path` 文件内容。
// 若读取文件失败，则返回一个空字符串。
# <翻译结束>


<原文开始>
// GetBytes returns the file content of `path` as []byte.
// It returns nil if it fails reading.
<原文结束>

# <翻译开始>
// GetBytes 函数返回 `path` 文件内容的 []byte 表示形式。
// 如果读取文件失败，则返回 nil。
# <翻译结束>


<原文开始>
// putContents puts binary content to file of `path`.
<原文结束>

# <翻译开始>
// putContents 将二进制内容写入到指定 `path` 的文件中。
# <翻译结束>


<原文开始>
// It supports creating file of `path` recursively.
<原文结束>

# <翻译开始>
// 它支持递归地创建 `path` 指定的文件。
# <翻译结束>


<原文开始>
// Opening file with given `flag` and `perm`.
<原文结束>

# <翻译开始>
// 使用给定的`flag`和`perm`打开文件。
# <翻译结束>


<原文开始>
// Truncate truncates file of `path` to given size by `size`.
<原文结束>

# <翻译开始>
// Truncate 函数通过给定的 `size` 截断指定路径 `path` 下的文件至相应大小。
# <翻译结束>


<原文开始>
// PutContents puts string `content` to file of `path`.
// It creates file of `path` recursively if it does not exist.
<原文结束>

# <翻译开始>
// PutContents 将字符串 `content` 写入到 `path` 指定的文件中。
// 如果该文件不存在，会递归创建包含 `path` 的所有目录及文件。
# <翻译结束>


<原文开始>
// PutContentsAppend appends string `content` to file of `path`.
// It creates file of `path` recursively if it does not exist.
<原文结束>

# <翻译开始>
// PutContentsAppend 将字符串 `content` 追加到 `path` 指定的文件中。
// 如果该文件不存在，会递归创建 `path` 指定的文件。
# <翻译结束>


<原文开始>
// PutBytes puts binary `content` to file of `path`.
// It creates file of `path` recursively if it does not exist.
<原文结束>

# <翻译开始>
// PutBytes 函数将二进制内容 `content` 写入到指定路径 `path` 的文件中。
// 如果目标文件不存在，它会递归创建该路径及其所有父目录，然后写入文件。
# <翻译结束>


<原文开始>
// PutBytesAppend appends binary `content` to file of `path`.
// It creates file of `path` recursively if it does not exist.
<原文结束>

# <翻译开始>
// PutBytesAppend 将二进制`content`追加到`path`文件中。
// 如果`path`文件不存在，会递归创建该文件。
# <翻译结束>


<原文开始>
// GetNextCharOffset returns the file offset for given `char` starting from `start`.
<原文结束>

# <翻译开始>
// GetNextCharOffset 函数从 `start` 开始，返回给定字符 `char` 对应的文件偏移量。
# <翻译结束>


<原文开始>
// GetNextCharOffsetByPath returns the file offset for given `char` starting from `start`.
// It opens file of `path` for reading with os.O_RDONLY flag and default perm.
<原文结束>

# <翻译开始>
// GetNextCharOffsetByPath 函数从给定的 `start` 位置开始，返回文件中指定 `char` 字符的文件偏移量。
// 它以 os.O_RDONLY 标志和默认权限打开 `path` 指定的文件进行读取。
# <翻译结束>


<原文开始>
// GetBytesTilChar returns the contents of the file as []byte
// until the next specified byte `char` position.
//
// Note: Returned value contains the character of the last position.
<原文结束>

# <翻译开始>
// GetBytesTilChar 返回文件内容作为 []byte，直到遇到下一个指定的字节 `char` 位置为止。
//
// 注意：返回的值中包含最后位置的字符。
# <翻译结束>


<原文开始>
// GetBytesTilCharByPath returns the contents of the file given by `path` as []byte
// until the next specified byte `char` position.
// It opens file of `path` for reading with os.O_RDONLY flag and default perm.
//
// Note: Returned value contains the character of the last position.
<原文结束>

# <翻译开始>
// GetBytesTilCharByPath 函数通过给定的 `path` 返回文件内容，直到遇到指定字节 `char` 的位置为止。
// 它以 os.O_RDONLY 标志和默认权限打开 `path` 指定的文件进行读取。
//
// 注意：返回的结果包含最后一个位置的字符。
# <翻译结束>


<原文开始>
// GetBytesByTwoOffsets returns the binary content as []byte from `start` to `end`.
// Note: Returned value does not contain the character of the last position, which means
// it returns content range as [start, end).
<原文结束>

# <翻译开始>
// GetBytesByTwoOffsets 函数从 `start` 位置到 `end` 位置返回二进制内容作为 []byte 类型。
// 注意：返回的值不包含结束位置的字符，也就是说，
// 它返回的内容范围是 [start, end)。
# <翻译结束>


<原文开始>
// GetBytesByTwoOffsetsByPath returns the binary content as []byte from `start` to `end`.
// Note: Returned value does not contain the character of the last position, which means
// it returns content range as [start, end).
// It opens file of `path` for reading with os.O_RDONLY flag and default perm.
<原文结束>

# <翻译开始>
// GetBytesByTwoOffsetsByPath 根据路径返回从 `start` 到 `end` 的二进制内容作为 []byte。
// 注意：返回的值不包含结束位置的字符，这意味着它返回的内容范围是 [start, end)。
// 它以 os.O_RDONLY 标志和默认权限打开 `path` 指定的文件进行读取。
# <翻译结束>


<原文开始>
// ReadLines reads file content line by line, which is passed to the callback function `callback` as string.
// It matches each line of text, separated by chars '\r' or '\n', stripped any trailing end-of-line marker.
//
// Note that the parameter passed to callback function might be an empty value, and the last non-empty line
// will be passed to callback function `callback` even if it has no newline marker.
<原文结束>

# <翻译开始>
// ReadLines 逐行读取文件内容，并将每行字符串作为参数传递给回调函数 `callback`。
// 它按 '\r' 或 '\n' 字符分割每一行文本，同时去掉末尾的换行标记。
//
// 注意，传递给回调函数的参数可能为空值，即使最后一行非空行没有换行标记，也会将其传递给回调函数 `callback`。
# <翻译结束>


<原文开始>
// ReadLinesBytes reads file content line by line, which is passed to the callback function `callback` as []byte.
// It matches each line of text, separated by chars '\r' or '\n', stripped any trailing end-of-line marker.
//
// Note that the parameter passed to callback function might be an empty value, and the last non-empty line
// will be passed to callback function `callback` even if it has no newline marker.
<原文结束>

# <翻译开始>
// ReadLinesBytes 逐行读取文件内容，并以 []byte 的形式将每一行传递给回调函数 `callback`。
// 它匹配由 '\r' 或 '\n' 分隔的每行文本，同时去除末尾的换行符。
//
// 注意，传递给回调函数的参数可能为空值，即使最后一行非空行没有换行符，也会将其传递给回调函数 `callback`。
# <翻译结束>

