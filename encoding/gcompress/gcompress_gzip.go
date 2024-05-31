// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gcompress

import (
	"bytes"
	"compress/gzip"
	"io"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
)

// Gzip compresses `data` using gzip algorithm.
// The optional parameter `level` specifies the compression level from
// 1 to 9 which means from none to the best compression.
//
// Note that it returns error if given `level` is invalid.

// ff:Gzip压缩字节集
// level:可选压缩级别
// data:字节集
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

// GzipFile compresses the file `src` to `dst` using gzip algorithm.

// ff:Gzip压缩文件
// err:错误
// level:可选压缩级别
// dstFilePath:压缩文件路径
// srcFilePath:文件路径
func GzipFile(srcFilePath, dstFilePath string, level ...int) (err error) {
	dstFile, err := gfile.Create(dstFilePath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	return GzipPathWriter(srcFilePath, dstFile, level...)
}

// GzipPathWriter compresses `filePath` to `writer` using gzip compressing algorithm.
//
// Note that the parameter `path` can be either a directory or a file.

// ff:Gzip压缩文件到Writer
// level:可选压缩级别
// writer:
// filePath:文件路径
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

// UnGzip decompresses `data` with gzip algorithm.

// ff:Gzip解压字节集
// data:gzip字节集
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

// UnGzipFile decompresses srcFilePath `src` to `dst` using gzip algorithm.

// ff:Gzip解压文件
// dstFilePath:文件路径
// srcFilePath:gzip文件路径
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
