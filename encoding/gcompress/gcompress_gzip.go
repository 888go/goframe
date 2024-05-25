// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gcompress

import (
	"bytes"
	"compress/gzip"
	"io"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
)

// Gzip 使用 gzip 算法压缩 `data`。
// 可选参数 `level` 指定压缩级别，范围从 1 到 9，1 表示无压缩，9 表示最佳压缩。
//
// 注意，如果给定的 `level` 不合法，它将返回错误。
// md5:55af48098fabf71a
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

// GzipFile 使用gzip算法将文件`src`压缩到`dst`。. md5:886a1c3d1f47c22f
func GzipFile(srcFilePath, dstFilePath string, level ...int) (err error) {
	dstFile, err := gfile.Create(dstFilePath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	return GzipPathWriter(srcFilePath, dstFile, level...)
}

// GzipPathWriter 使用gzip压缩算法将`filePath`的内容压缩并写入到`writer`中。
//
// 注意，参数`path`既可以是一个目录，也可以是一个文件。
// md5:5da525f970882d97
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

// UnGzip 使用gzip算法对`data`进行解压缩。. md5:65410de81354eedd
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

// UnGzipFile 使用gzip算法将源文件路径`src`解压缩到`dst`。. md5:d8b51242e54f12db
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
