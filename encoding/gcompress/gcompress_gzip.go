// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 压缩类

import (
	"bytes"
	"compress/gzip"
	"io"
	
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/os/gfile"
)

// Gzip 使用gzip算法压缩`data`。
// 可选参数`level`指定压缩级别，范围从1到9，
// 这意味着从无压缩到最佳压缩。
//
// 注意，如果给定的`level`无效，将会返回错误。
func Gzip压缩字节集(字节集 []byte, 可选压缩级别 ...int) ([]byte, error) {
	var (
		writer *gzip.Writer
		buf    bytes.Buffer
		err    error
	)
	if len(可选压缩级别) > 0 {
		writer, err = gzip.NewWriterLevel(&buf, 可选压缩级别[0])
		if err != nil {
			err = 错误类.X多层错误并格式化(err, `gzip.NewWriterLevel failed for level "%d"`, 可选压缩级别[0])
			return nil, err
		}
	} else {
		writer = gzip.NewWriter(&buf)
	}
	if _, err = writer.Write(字节集); err != nil {
		err = 错误类.X多层错误(err, `writer.Write failed`)
		return nil, err
	}
	if err = writer.Close(); err != nil {
		err = 错误类.X多层错误(err, `writer.Close failed`)
		return nil, err
	}
	return buf.Bytes(), nil
}

// GzipFile 使用gzip算法将文件`src`压缩到`dst`。
func Gzip压缩文件(文件路径, 压缩文件路径 string, 可选压缩级别 ...int) (错误 error) {
	dstFile, 错误 := 文件类.X创建文件与目录(压缩文件路径)
	if 错误 != nil {
		return 错误
	}
	defer dstFile.Close()

	return Gzip压缩文件到Writer(文件路径, dstFile, 可选压缩级别...)
}

// GzipPathWriter 使用gzip压缩算法将`filePath`压缩并写入到`writer`。
//
// 注意，参数`path`可以是一个目录或一个文件。
func Gzip压缩文件到Writer(文件路径 string, writer io.Writer, 可选压缩级别 ...int) error {
	var (
		gzipWriter *gzip.Writer
		err        error
	)
	srcFile, err := 文件类.X打开并按只读模式(文件路径)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	if len(可选压缩级别) > 0 {
		gzipWriter, err = gzip.NewWriterLevel(writer, 可选压缩级别[0])
		if err != nil {
			return 错误类.X多层错误(err, `gzip.NewWriterLevel failed`)
		}
	} else {
		gzipWriter = gzip.NewWriter(writer)
	}
	defer gzipWriter.Close()

	if _, err = io.Copy(gzipWriter, srcFile); err != nil {
		err = 错误类.X多层错误(err, `io.Copy failed`)
		return err
	}
	return nil
}

// UnGzip 使用gzip算法对`data`进行解压缩
func Gzip解压字节集(gzip字节集 []byte) ([]byte, error) {
	var buf bytes.Buffer
	reader, err := gzip.NewReader(bytes.NewReader(gzip字节集))
	if err != nil {
		err = 错误类.X多层错误(err, `gzip.NewReader failed`)
		return nil, err
	}
	if _, err = io.Copy(&buf, reader); err != nil {
		err = 错误类.X多层错误(err, `io.Copy failed`)
		return nil, err
	}
	if err = reader.Close(); err != nil {
		err = 错误类.X多层错误(err, `reader.Close failed`)
		return buf.Bytes(), err
	}
	return buf.Bytes(), nil
}

// UnGzipFile 使用gzip算法将源文件路径`src`解压缩到`dst`。
func Gzip解压文件(gzip文件路径, 文件路径 string) error {
	srcFile, err := 文件类.X打开并按只读模式(gzip文件路径)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstFile, err := 文件类.X创建文件与目录(文件路径)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	reader, err := gzip.NewReader(srcFile)
	if err != nil {
		err = 错误类.X多层错误(err, `gzip.NewReader failed`)
		return err
	}
	defer reader.Close()

	if _, err = io.Copy(dstFile, reader); err != nil {
		err = 错误类.X多层错误(err, `io.Copy failed`)
		return err
	}
	return nil
}
