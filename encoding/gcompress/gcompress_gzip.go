// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcompress
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
func Gzip(data []byte, level ...int) ([]byte, error) {
	var (
		writer *gzip.Writer
		buf    bytes.Buffer
		err    error
	)
	if len(level) > 0 {
		writer, err = gzip.NewWriterLevel(&buf, level[0])
		if err != nil {
			err = gerror.Wrapf(err, `gzip.NewWriterLevel failed for level "%d"`, level[0])
			return nil, err
		}
	} else {
		writer = gzip.NewWriter(&buf)
	}
	if _, err = writer.Write(data); err != nil {
		err = gerror.Wrap(err, `writer.Write failed`)
		return nil, err
	}
	if err = writer.Close(); err != nil {
		err = gerror.Wrap(err, `writer.Close failed`)
		return nil, err
	}
	return buf.Bytes(), nil
}

// GzipFile 使用gzip算法将文件`src`压缩到`dst`。
func GzipFile(srcFilePath, dstFilePath string, level ...int) (err error) {
	dstFile, err := gfile.Create(dstFilePath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	return GzipPathWriter(srcFilePath, dstFile, level...)
}

// GzipPathWriter 使用gzip压缩算法将`filePath`压缩并写入到`writer`。
//
// 注意，参数`path`可以是一个目录或一个文件。
func GzipPathWriter(filePath string, writer io.Writer, level ...int) error {
	var (
		gzipWriter *gzip.Writer
		err        error
	)
	srcFile, err := gfile.Open(filePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	if len(level) > 0 {
		gzipWriter, err = gzip.NewWriterLevel(writer, level[0])
		if err != nil {
			return gerror.Wrap(err, `gzip.NewWriterLevel failed`)
		}
	} else {
		gzipWriter = gzip.NewWriter(writer)
	}
	defer gzipWriter.Close()

	if _, err = io.Copy(gzipWriter, srcFile); err != nil {
		err = gerror.Wrap(err, `io.Copy failed`)
		return err
	}
	return nil
}

// UnGzip 使用gzip算法对`data`进行解压缩
func UnGzip(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		err = gerror.Wrap(err, `gzip.NewReader failed`)
		return nil, err
	}
	if _, err = io.Copy(&buf, reader); err != nil {
		err = gerror.Wrap(err, `io.Copy failed`)
		return nil, err
	}
	if err = reader.Close(); err != nil {
		err = gerror.Wrap(err, `reader.Close failed`)
		return buf.Bytes(), err
	}
	return buf.Bytes(), nil
}

// UnGzipFile 使用gzip算法将源文件路径`src`解压缩到`dst`。
func UnGzipFile(srcFilePath, dstFilePath string) error {
	srcFile, err := gfile.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstFile, err := gfile.Create(dstFilePath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	reader, err := gzip.NewReader(srcFile)
	if err != nil {
		err = gerror.Wrap(err, `gzip.NewReader failed`)
		return err
	}
	defer reader.Close()

	if _, err = io.Copy(dstFile, reader); err != nil {
		err = gerror.Wrap(err, `io.Copy failed`)
		return err
	}
	return nil
}
