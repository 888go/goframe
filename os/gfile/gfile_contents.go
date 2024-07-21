// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gfile

import (
	"bufio"
	"io"
	"os"

	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	// DefaultReadBuffer 是用于读取文件内容的缓冲区大小。 md5:ff60c55c31212953
	DefaultReadBuffer = 1024
)

// GetContents 以字符串形式返回路径 `path` 的文件内容。
// 如果读取失败，它将返回空字符串。
// md5:3426170b655a7b9d
// ff:读文本
// path:路径
func GetContents(path string) string {
	return string(GetBytes(path))
}

// GetBytes 将路径 `path` 对应的文件内容以 []byte 形式返回。
// 如果读取失败，则返回 nil。
// md5:be06b7ebc28d3d98
// ff:读字节集
// path:路径
func GetBytes(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	return data
}

// putContents 将二进制内容写入文件`path`。 md5:bd53892836016a1c
func putContents(path string, data []byte, flag int, perm os.FileMode) error {
	// 它支持递归地创建`path`指定的文件。 md5:4ca8118123f2e629
	dir := Dir(path)
	if !Exists(dir) {
		if err := Mkdir(dir); err != nil {
			return err
		}
	}
	// 使用给定的`flag`和`perm`打开文件。 md5:dcef8b2678320fb5
	f, err := OpenWithFlagPerm(path, flag, perm)
	if err != nil {
		return err
	}
	defer f.Close()
	// Write data.
	var n int
	if n, err = f.Write(data); err != nil {
		err = gerror.Wrapf(err, `Write data to file "%s" failed`, path)
		return err
	} else if n < len(data) {
		return io.ErrShortWrite
	}
	return nil
}

// Truncate 将路径为`path`的文件截取到给定的大小`size`。 md5:fc49c34fbdada146
// ff:截断
// path:路径
// size:长度
// err:错误
func Truncate(path string, size int) (err error) {
	err = os.Truncate(path, int64(size))
	if err != nil {
		err = gerror.Wrapf(err, `os.Truncate failed for file "%s", size "%d"`, path, size)
	}
	return
}

// PutContents 将字符串 `content` 写入到文件 `path` 中。
// 如果文件 `path` 不存在，该函数会递归创建。
// md5:155829d6dddf1340
// ff:写入文本
// path:路径
// content:文本
func PutContents(path string, content string) error {
	return putContents(path, []byte(content), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, DefaultPermOpen)
}

// PutContentsAppend 将字符串 `content` 追加到文件 `path` 中。
// 如果文件 `path` 不存在，它会递归创建该文件。
// md5:55f7095d64183741
// ff:追加文本
// path:路径
// content:文本
func PutContentsAppend(path string, content string) error {
	return putContents(path, []byte(content), os.O_WRONLY|os.O_CREATE|os.O_APPEND, DefaultPermOpen)
}

// PutBytes 将二进制 `content` 存储到 `path` 对应的文件中。
// 如果文件 `path` 不存在，它会递归创建文件。
// md5:3ac1025ef9039ab7
// ff:写入字节集
// path:路径
// content:字节集
func PutBytes(path string, content []byte) error {
	return putContents(path, content, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, DefaultPermOpen)
}

// PutBytesAppend 将二进制`content`追加到`path`指定的文件中。
// 如果该文件不存在，它会递归创建`path`指定的文件。
// md5:c5f2a5eb57487328
// ff:追加字节集
// path:路径
// content:字节集
func PutBytesAppend(path string, content []byte) error {
	return putContents(path, content, os.O_WRONLY|os.O_CREATE|os.O_APPEND, DefaultPermOpen)
}

// GetNextCharOffset 返回从`start`开始的给定`char`的文件偏移量。 md5:fd9885f76bb1a398
// ff:取字符偏移位置
// reader:
// char:待查找字符
// start:查找起点
func GetNextCharOffset(reader io.ReaderAt, char byte, start int64) int64 {
	buffer := make([]byte, DefaultReadBuffer)
	offset := start
	for {
		if n, err := reader.ReadAt(buffer, offset); n > 0 {
			for i := 0; i < n; i++ {
				if buffer[i] == char {
					return int64(i) + offset
				}
			}
			offset += int64(n)
		} else if err != nil {
			break
		}
	}
	return -1
}

// GetNextCharOffsetByPath 从`start`开始，返回给定`char`对应的文件偏移量。
// 它使用os.O_RDONLY标志和默认权限打开`path`文件进行读取。
// md5:c328b89cddf1bd1d
// ff:取文件字符偏移位置
// path:路径
// char:待查找字符
// start:查找起点
func GetNextCharOffsetByPath(path string, char byte, start int64) int64 {
	if f, err := OpenWithFlagPerm(path, os.O_RDONLY, DefaultPermOpen); err == nil {
		defer f.Close()
		return GetNextCharOffset(f, char, start)
	}
	return -1
}

// GetBytesTilChar returns the contents of the file as []byte
// until the next specified byte `char` position.
//
// ff:取字节集按字符位置
// reader:
// char:待查找字符
// start:查找起点
func GetBytesTilChar(reader io.ReaderAt, char byte, start int64) ([]byte, int64) {
	if offset := GetNextCharOffset(reader, char, start); offset != -1 {
		return GetBytesByTwoOffsets(reader, start, offset+1), offset
	}
	return nil, -1
}

// GetBytesTilCharByPath returns the contents of the file given by `path` as []byte
// until the next specified byte `char` position.
// It opens file of `path` for reading with os.O_RDONLY flag and default perm.
//
// ff:取文件字节集按字符位置
// path:路径
// char:待查找字符
// start:查找起点
func GetBytesTilCharByPath(path string, char byte, start int64) ([]byte, int64) {
	if f, err := OpenWithFlagPerm(path, os.O_RDONLY, DefaultPermOpen); err == nil {
		defer f.Close()
		return GetBytesTilChar(f, char, start)
	}
	return nil, -1
}

// GetBytesByTwoOffsets returns the binary content as []byte from `start` to `end`.
// it returns content range as [start, end).
// ff:取字节集按范围
// reader:
// start:起点
// end:终点
func GetBytesByTwoOffsets(reader io.ReaderAt, start int64, end int64) []byte {
	buffer := make([]byte, end-start)
	if _, err := reader.ReadAt(buffer, start); err != nil {
		return nil
	}
	return buffer
}

// GetBytesByTwoOffsetsByPath returns the binary content as []byte from `start` to `end`.
// it returns content range as [start, end).
// It opens file of `path` for reading with os.O_RDONLY flag and default perm.
// ff:取文件字节集按范围
// path:路径
// start:起点
// end:终点
func GetBytesByTwoOffsetsByPath(path string, start int64, end int64) []byte {
	if f, err := OpenWithFlagPerm(path, os.O_RDONLY, DefaultPermOpen); err == nil {
		defer f.Close()
		return GetBytesByTwoOffsets(f, start, end)
	}
	return nil
}

// ReadLines 逐行读取文件内容，将每一行作为字符串传递给回调函数 `callback`。
// 它匹配由 '\r' 或 '\n' 分隔的每一行文本，并移除任何尾随换行符。
// 
// 注意，回调函数接收到的参数可能为空值，即使最后一行没有换行符，也会将其作为非空行传递给 `callback` 函数。
// md5:462b920487edad37
// ff:逐行读文本_函数
// file:文件路径
// callback:回调函数
// line:文本
func ReadLines(file string, callback func(line string) error) error {
	f, err := Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if err = callback(scanner.Text()); err != nil {
			return err
		}
	}
	return nil
}

			// ReadLinesBytes 逐行读取文件内容，并将每一行作为 []byte 类型传递给回调函数 `callback`。
			// 它会匹配以字符 '\r' 或 '\n' 分隔的每一行文本，并去除任何尾随的换行标记。
			//
			// 注意，传递给回调函数的参数可能是一个空值，并且即使最后一行非空行没有换行标记，
			// 也会被传递给回调函数 `callback`。
			// md5:214fffa05bf19040
// ff:逐行读字节集_函数
// file:文件路径
// callback:回调函数
// bytes:字节集
func ReadLinesBytes(file string, callback func(bytes []byte) error) error {
	f, err := Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if err = callback(scanner.Bytes()); err != nil {
			return err
		}
	}
	return nil
}
