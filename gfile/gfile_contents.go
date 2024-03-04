// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gfile

import (
	"bufio"
	"io"
	"os"
	
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	// DefaultReadBuffer 是用于读取文件内容的缓冲区大小。
	DefaultReadBuffer = 1024
)

// GetContents 函数返回字符串形式的 `path` 文件内容。
// 若读取文件失败，则返回一个空字符串。
func GetContents(path string) string {
	return string(GetBytes(path))
}

// GetBytes 函数返回 `path` 文件内容的 []byte 表示形式。
// 如果读取文件失败，则返回 nil。
func GetBytes(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	return data
}

// putContents 将二进制内容写入到指定 `path` 的文件中。
func putContents(path string, data []byte, flag int, perm os.FileMode) error {
	// 它支持递归地创建 `path` 指定的文件。
	dir := Dir(path)
	if !Exists(dir) {
		if err := Mkdir(dir); err != nil {
			return err
		}
	}
	// 使用给定的`flag`和`perm`打开文件。
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

// Truncate 函数通过给定的 `size` 截断指定路径 `path` 下的文件至相应大小。
func Truncate(path string, size int) (err error) {
	err = os.Truncate(path, int64(size))
	if err != nil {
		err = gerror.Wrapf(err, `os.Truncate failed for file "%s", size "%d"`, path, size)
	}
	return
}

// PutContents 将字符串 `content` 写入到 `path` 指定的文件中。
// 如果该文件不存在，会递归创建包含 `path` 的所有目录及文件。
func PutContents(path string, content string) error {
	return putContents(path, []byte(content), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, DefaultPermOpen)
}

// PutContentsAppend 将字符串 `content` 追加到 `path` 指定的文件中。
// 如果该文件不存在，会递归创建 `path` 指定的文件。
func PutContentsAppend(path string, content string) error {
	return putContents(path, []byte(content), os.O_WRONLY|os.O_CREATE|os.O_APPEND, DefaultPermOpen)
}

// PutBytes 函数将二进制内容 `content` 写入到指定路径 `path` 的文件中。
// 如果目标文件不存在，它会递归创建该路径及其所有父目录，然后写入文件。
func PutBytes(path string, content []byte) error {
	return putContents(path, content, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, DefaultPermOpen)
}

// PutBytesAppend 将二进制`content`追加到`path`文件中。
// 如果`path`文件不存在，会递归创建该文件。
func PutBytesAppend(path string, content []byte) error {
	return putContents(path, content, os.O_WRONLY|os.O_CREATE|os.O_APPEND, DefaultPermOpen)
}

// GetNextCharOffset 函数从 `start` 开始，返回给定字符 `char` 对应的文件偏移量。
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

// GetNextCharOffsetByPath 函数从给定的 `start` 位置开始，返回文件中指定 `char` 字符的文件偏移量。
// 它以 os.O_RDONLY 标志和默认权限打开 `path` 指定的文件进行读取。
func GetNextCharOffsetByPath(path string, char byte, start int64) int64 {
	if f, err := OpenWithFlagPerm(path, os.O_RDONLY, DefaultPermOpen); err == nil {
		defer f.Close()
		return GetNextCharOffset(f, char, start)
	}
	return -1
}

// GetBytesTilChar 返回文件内容作为 []byte，直到遇到下一个指定的字节 `char` 位置为止。
//
// 注意：返回的值中包含最后位置的字符。
func GetBytesTilChar(reader io.ReaderAt, char byte, start int64) ([]byte, int64) {
	if offset := GetNextCharOffset(reader, char, start); offset != -1 {
		return GetBytesByTwoOffsets(reader, start, offset+1), offset
	}
	return nil, -1
}

// GetBytesTilCharByPath 函数通过给定的 `path` 返回文件内容，直到遇到指定字节 `char` 的位置为止。
// 它以 os.O_RDONLY 标志和默认权限打开 `path` 指定的文件进行读取。
//
// 注意：返回的结果包含最后一个位置的字符。
func GetBytesTilCharByPath(path string, char byte, start int64) ([]byte, int64) {
	if f, err := OpenWithFlagPerm(path, os.O_RDONLY, DefaultPermOpen); err == nil {
		defer f.Close()
		return GetBytesTilChar(f, char, start)
	}
	return nil, -1
}

// GetBytesByTwoOffsets 函数从 `start` 位置到 `end` 位置返回二进制内容作为 []byte 类型。
// 注意：返回的值不包含结束位置的字符，也就是说，
// 它返回的内容范围是 [start, end)。
func GetBytesByTwoOffsets(reader io.ReaderAt, start int64, end int64) []byte {
	buffer := make([]byte, end-start)
	if _, err := reader.ReadAt(buffer, start); err != nil {
		return nil
	}
	return buffer
}

// GetBytesByTwoOffsetsByPath 根据路径返回从 `start` 到 `end` 的二进制内容作为 []byte。
// 注意：返回的值不包含结束位置的字符，这意味着它返回的内容范围是 [start, end)。
// 它以 os.O_RDONLY 标志和默认权限打开 `path` 指定的文件进行读取。
func GetBytesByTwoOffsetsByPath(path string, start int64, end int64) []byte {
	if f, err := OpenWithFlagPerm(path, os.O_RDONLY, DefaultPermOpen); err == nil {
		defer f.Close()
		return GetBytesByTwoOffsets(f, start, end)
	}
	return nil
}

// ReadLines 逐行读取文件内容，并将每行字符串作为参数传递给回调函数 `callback`。
// 它按 '\r' 或 '\n' 字符分割每一行文本，同时去掉末尾的换行标记。
//
// 注意，传递给回调函数的参数可能为空值，即使最后一行非空行没有换行标记，也会将其传递给回调函数 `callback`。
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

// ReadLinesBytes 逐行读取文件内容，并以 []byte 的形式将每一行传递给回调函数 `callback`。
// 它匹配由 '\r' 或 '\n' 分隔的每行文本，同时去除末尾的换行符。
//
// 注意，传递给回调函数的参数可能为空值，即使最后一行非空行没有换行符，也会将其传递给回调函数 `callback`。
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
