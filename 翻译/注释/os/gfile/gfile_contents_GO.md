
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
// DefaultReadBuffer is the buffer size for reading file content.
<原文结束>

# <翻译开始>
	// DefaultReadBuffer 是用于读取文件内容的缓冲区大小。 md5:ff60c55c31212953
# <翻译结束>


<原文开始>
// GetContents returns the file content of `path` as string.
// It returns en empty string if it fails reading.
<原文结束>

# <翻译开始>
// GetContents 以字符串形式返回路径 `path` 的文件内容。
// 如果读取失败，它将返回空字符串。
// md5:3426170b655a7b9d
# <翻译结束>


<原文开始>
// GetBytes returns the file content of `path` as []byte.
// It returns nil if it fails reading.
<原文结束>

# <翻译开始>
// GetBytes 将路径 `path` 对应的文件内容以 []byte 形式返回。
// 如果读取失败，则返回 nil。
// md5:be06b7ebc28d3d98
# <翻译结束>


<原文开始>
// putContents puts binary content to file of `path`.
<原文结束>

# <翻译开始>
// putContents 将二进制内容写入文件`path`。 md5:bd53892836016a1c
# <翻译结束>


<原文开始>
// It supports creating file of `path` recursively.
<原文结束>

# <翻译开始>
	// 它支持递归地创建`path`指定的文件。 md5:4ca8118123f2e629
# <翻译结束>


<原文开始>
// Opening file with given `flag` and `perm`.
<原文结束>

# <翻译开始>
	// 使用给定的`flag`和`perm`打开文件。 md5:dcef8b2678320fb5
# <翻译结束>


<原文开始>
// Truncate truncates file of `path` to given size by `size`.
<原文结束>

# <翻译开始>
// Truncate 将路径为`path`的文件截取到给定的大小`size`。 md5:fc49c34fbdada146
# <翻译结束>


<原文开始>
// PutContents puts string `content` to file of `path`.
// It creates file of `path` recursively if it does not exist.
<原文结束>

# <翻译开始>
// PutContents 将字符串 `content` 写入到文件 `path` 中。
// 如果文件 `path` 不存在，该函数会递归创建。
// md5:155829d6dddf1340
# <翻译结束>


<原文开始>
// PutContentsAppend appends string `content` to file of `path`.
// It creates file of `path` recursively if it does not exist.
<原文结束>

# <翻译开始>
// PutContentsAppend 将字符串 `content` 追加到文件 `path` 中。
// 如果文件 `path` 不存在，它会递归创建该文件。
// md5:55f7095d64183741
# <翻译结束>


<原文开始>
// PutBytes puts binary `content` to file of `path`.
// It creates file of `path` recursively if it does not exist.
<原文结束>

# <翻译开始>
// PutBytes 将二进制 `content` 存储到 `path` 对应的文件中。
// 如果文件 `path` 不存在，它会递归创建文件。
// md5:3ac1025ef9039ab7
# <翻译结束>


<原文开始>
// PutBytesAppend appends binary `content` to file of `path`.
// It creates file of `path` recursively if it does not exist.
<原文结束>

# <翻译开始>
// PutBytesAppend 将二进制`content`追加到`path`指定的文件中。
// 如果该文件不存在，它会递归创建`path`指定的文件。
// md5:c5f2a5eb57487328
# <翻译结束>


<原文开始>
// GetNextCharOffset returns the file offset for given `char` starting from `start`.
<原文结束>

# <翻译开始>
// GetNextCharOffset 返回从`start`开始的给定`char`的文件偏移量。 md5:fd9885f76bb1a398
# <翻译结束>


<原文开始>
// GetNextCharOffsetByPath returns the file offset for given `char` starting from `start`.
// It opens file of `path` for reading with os.O_RDONLY flag and default perm.
<原文结束>

# <翻译开始>
// GetNextCharOffsetByPath 从`start`开始，返回给定`char`对应的文件偏移量。
// 它使用os.O_RDONLY标志和默认权限打开`path`文件进行读取。
// md5:c328b89cddf1bd1d
# <翻译结束>


<原文开始>
// GetBytesTilChar returns the contents of the file as []byte
// until the next specified byte `char` position.
//
// Note: Returned value contains the character of the last position.
<原文结束>

# <翻译开始>
// GetBytesTilChar 读取文件内容，直到遇到指定的字节`char`位置，然后返回该位置之前的所有内容作为[]byte。
// 
// 注意：返回值包含最后一个位置的字符。
// md5:d7db409d8f51ccd8
# <翻译结束>


<原文开始>
// GetBytesTilCharByPath returns the contents of the file given by `path` as []byte
// until the next specified byte `char` position.
// It opens file of `path` for reading with os.O_RDONLY flag and default perm.
//
// Note: Returned value contains the character of the last position.
<原文结束>

# <翻译开始>
// GetBytesTilCharByPath 根据给定的`path`返回文件内容，直到遇到下一个指定的字节`char`位置为止。
// 它使用 os.O_RDONLY 标志和默认权限以只读方式打开`path`指定的文件。
//
// 注意：返回的值包含最后一个位置的字符。
// md5:aefddbb30f37dff7
# <翻译结束>


<原文开始>
// GetBytesByTwoOffsets returns the binary content as []byte from `start` to `end`.
// Note: Returned value does not contain the character of the last position, which means
// it returns content range as [start, end).
<原文结束>

# <翻译开始>
// GetBytesByTwoOffsets 从`start`到`end`返回二进制内容作为[]byte。
// 注意：返回的值不包含最后一个位置的字符，意味着它返回的内容范围是[start, end)。
// md5:fb50a5776f6863f6
# <翻译结束>


<原文开始>
// GetBytesByTwoOffsetsByPath returns the binary content as []byte from `start` to `end`.
// Note: Returned value does not contain the character of the last position, which means
// it returns content range as [start, end).
// It opens file of `path` for reading with os.O_RDONLY flag and default perm.
<原文结束>

# <翻译开始>
// GetBytesByTwoOffsetsByPath 通过"path"从`start`到`end`获取二进制内容并返回为[]byte。
// 注意：返回值不包含最后一个位置的字符，意味着它返回的内容范围是[start, end)。
// 它使用os.O_RDONLY标志和默认权限打开`path`文件进行读取。
// md5:26e574b11f4b1e0f
# <翻译结束>


<原文开始>
// ReadLines reads file content line by line, which is passed to the callback function `callback` as string.
// It matches each line of text, separated by chars '\r' or '\n', stripped any trailing end-of-line marker.
//
// Note that the parameter passed to callback function might be an empty value, and the last non-empty line
// will be passed to callback function `callback` even if it has no newline marker.
<原文结束>

# <翻译开始>
// ReadLines 逐行读取文件内容，将每一行作为字符串传递给回调函数 `callback`。
// 它匹配由 '\r' 或 '\n' 分隔的每一行文本，并移除任何尾随换行符。
// 
// 注意，回调函数接收到的参数可能为空值，即使最后一行没有换行符，也会将其作为非空行传递给 `callback` 函数。
// md5:462b920487edad37
# <翻译结束>


<原文开始>
// ReadLinesBytes reads file content line by line, which is passed to the callback function `callback` as []byte.
// It matches each line of text, separated by chars '\r' or '\n', stripped any trailing end-of-line marker.
//
// Note that the parameter passed to callback function might be an empty value, and the last non-empty line
// will be passed to callback function `callback` even if it has no newline marker.
<原文结束>

# <翻译开始>
// ReadLinesBytes 逐行读取文件内容，并将每一行作为 []byte 类型传递给回调函数 `callback`。
// 它会匹配以字符 '\r' 或 '\n' 分隔的每一行文本，并去除任何尾随的换行标记。
//
// 注意，传递给回调函数的参数可能是一个空值，并且即使最后一行非空行没有换行标记，
// 也会被传递给回调函数 `callback`。
// md5:214fffa05bf19040
# <翻译结束>

