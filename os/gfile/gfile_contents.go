// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件类

import (
	"bufio"
	"io"
	"os"
	
	"github.com/888go/goframe/errors/gerror"
)

var (
	// DefaultReadBuffer 是用于读取文件内容的缓冲区大小。
	DefaultReadBuffer = 1024
)

// GetContents 函数返回字符串形式的 `path` 文件内容。
// 若读取文件失败，则返回一个空字符串。
func X读文本(路径 string) string {
	return string(X读字节集(路径))
}

// GetBytes 函数返回 `path` 文件内容的 []byte 表示形式。
// 如果读取文件失败，则返回 nil。
func X读字节集(路径 string) []byte {
	data, err := os.ReadFile(路径)
	if err != nil {
		return nil
	}
	return data
}

// putContents 将二进制内容写入到指定 `path` 的文件中。
func putContents(path string, data []byte, flag int, perm os.FileMode) error {
	// 它支持递归地创建 `path` 指定的文件。
	dir := X路径取父目录(path)
	if !X是否存在(dir) {
		if err := X创建目录(dir); err != nil {
			return err
		}
	}
	// 使用给定的`flag`和`perm`打开文件。
	f, err := OpenWithFlagPerm别名(path, flag, perm)
	if err != nil {
		return err
	}
	defer f.Close()
	// Write data.
	var n int
	if n, err = f.Write(data); err != nil {
		err = 错误类.X多层错误并格式化(err, `Write data to file "%s" failed`, path)
		return err
	} else if n < len(data) {
		return io.ErrShortWrite
	}
	return nil
}

// Truncate 函数通过给定的 `size` 截断指定路径 `path` 下的文件至相应大小。
func X截断(路径 string, 长度 int) (错误 error) {
	错误 = os.Truncate(路径, int64(长度))
	if 错误 != nil {
		错误 = 错误类.X多层错误并格式化(错误, `os.Truncate failed for file "%s", size "%d"`, 路径, 长度)
	}
	return
}

// PutContents 将字符串 `content` 写入到 `path` 指定的文件中。
// 如果该文件不存在，会递归创建包含 `path` 的所有目录及文件。
func X写入文本(路径 string, 文本 string) error {
	return putContents(路径, []byte(文本), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, DefaultPermOpen)
}

// PutContentsAppend 将字符串 `content` 追加到 `path` 指定的文件中。
// 如果该文件不存在，会递归创建 `path` 指定的文件。
func X追加文本(路径 string, 文本 string) error {
	return putContents(路径, []byte(文本), os.O_WRONLY|os.O_CREATE|os.O_APPEND, DefaultPermOpen)
}

// PutBytes 函数将二进制内容 `content` 写入到指定路径 `path` 的文件中。
// 如果目标文件不存在，它会递归创建该路径及其所有父目录，然后写入文件。
func X写入字节集(路径 string, 字节集 []byte) error {
	return putContents(路径, 字节集, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, DefaultPermOpen)
}

// PutBytesAppend 将二进制`content`追加到`path`文件中。
// 如果`path`文件不存在，会递归创建该文件。
func X追加字节集(路径 string, 字节集 []byte) error {
	return putContents(路径, 字节集, os.O_WRONLY|os.O_CREATE|os.O_APPEND, DefaultPermOpen)
}

// GetNextCharOffset 函数从 `start` 开始，返回给定字符 `char` 对应的文件偏移量。
func X取字符偏移位置(reader io.ReaderAt, 待查找字符 byte, 查找起点 int64) int64 {
	buffer := make([]byte, DefaultReadBuffer)
	offset := 查找起点
	for {
		if n, err := reader.ReadAt(buffer, offset); n > 0 {
			for i := 0; i < n; i++ {
				if buffer[i] == 待查找字符 {
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
func X取文件字符偏移位置(路径 string, 待查找字符 byte, 查找起点 int64) int64 {
	if f, err := OpenWithFlagPerm别名(路径, os.O_RDONLY, DefaultPermOpen); err == nil {
		defer f.Close()
		return X取字符偏移位置(f, 待查找字符, 查找起点)
	}
	return -1
}

// GetBytesTilChar 返回文件内容作为 []byte，直到遇到下一个指定的字节 `char` 位置为止。
//
// 注意：返回的值中包含最后位置的字符。
func X取字节集按字符位置(reader io.ReaderAt, 待查找字符 byte, 查找起点 int64) ([]byte, int64) {
	if offset := X取字符偏移位置(reader, 待查找字符, 查找起点); offset != -1 {
		return X取字节集按范围(reader, 查找起点, offset+1), offset
	}
	return nil, -1
}

// GetBytesTilCharByPath 函数通过给定的 `path` 返回文件内容，直到遇到指定字节 `char` 的位置为止。
// 它以 os.O_RDONLY 标志和默认权限打开 `path` 指定的文件进行读取。
//
// 注意：返回的结果包含最后一个位置的字符。
func X取文件字节集按字符位置(路径 string, 待查找字符 byte, 查找起点 int64) ([]byte, int64) {
	if f, err := OpenWithFlagPerm别名(路径, os.O_RDONLY, DefaultPermOpen); err == nil {
		defer f.Close()
		return X取字节集按字符位置(f, 待查找字符, 查找起点)
	}
	return nil, -1
}

// GetBytesByTwoOffsets 函数从 `start` 位置到 `end` 位置返回二进制内容作为 []byte 类型。
// 注意：返回的值不包含结束位置的字符，也就是说，
// 它返回的内容范围是 [start, end)。
func X取字节集按范围(reader io.ReaderAt, 起点 int64, 终点 int64) []byte {
	buffer := make([]byte, 终点-起点)
	if _, err := reader.ReadAt(buffer, 起点); err != nil {
		return nil
	}
	return buffer
}

// GetBytesByTwoOffsetsByPath 根据路径返回从 `start` 到 `end` 的二进制内容作为 []byte。
// 注意：返回的值不包含结束位置的字符，这意味着它返回的内容范围是 [start, end)。
// 它以 os.O_RDONLY 标志和默认权限打开 `path` 指定的文件进行读取。
func X取文件字节集按范围(路径 string, 起点 int64, 终点 int64) []byte {
	if f, err := OpenWithFlagPerm别名(路径, os.O_RDONLY, DefaultPermOpen); err == nil {
		defer f.Close()
		return X取字节集按范围(f, 起点, 终点)
	}
	return nil
}

// ReadLines 逐行读取文件内容，并将每行字符串作为参数传递给回调函数 `callback`。
// 它按 '\r' 或 '\n' 字符分割每一行文本，同时去掉末尾的换行标记。
//
// 注意，传递给回调函数的参数可能为空值，即使最后一行非空行没有换行标记，也会将其传递给回调函数 `callback`。
func X逐行读文本_函数(文件路径 string, 回调函数 func(文本 string) error) error {
	f, err := X打开并按只读模式(文件路径)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if err = 回调函数(scanner.Text()); err != nil {
			return err
		}
	}
	return nil
}

// ReadLinesBytes 逐行读取文件内容，并以 []byte 的形式将每一行传递给回调函数 `callback`。
// 它匹配由 '\r' 或 '\n' 分隔的每行文本，同时去除末尾的换行符。
//
// 注意，传递给回调函数的参数可能为空值，即使最后一行非空行没有换行符，也会将其传递给回调函数 `callback`。
func X逐行读字节集_函数(文件路径 string, 回调函数 func(字节集 []byte) error) error {
	f, err := X打开并按只读模式(文件路径)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if err = 回调函数(scanner.Bytes()); err != nil {
			return err
		}
	}
	return nil
}
