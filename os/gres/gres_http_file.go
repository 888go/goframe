// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 资源类

import (
	"bytes"
	"os"

	gerror "github.com/888go/goframe/errors/gerror"
)

// Close 实现了 http.File 接口。 md5:ed68e1aedf92d678
func (f *File) Close() error {
	return nil
}

// Readdir 实现了 http.File 接口的 Readdir 方法。 md5:d47ea30bb07047c4
func (f *File) Readdir(count int) ([]os.FileInfo, error) {
	files := f.resource.ScanDir(f.Name(), "*", false)
	if len(files) > 0 {
		if count <= 0 || count > len(files) {
			count = len(files)
		}
		infos := make([]os.FileInfo, count)
		for k, v := range files {
			infos[k] = v.FileInfo()
		}
		return infos, nil
	}
	return nil, nil
}

// Stat 实现了 http.File 接口中的 Stat 方法。 md5:f25fb28810e2d18b
func (f *File) Stat() (os.FileInfo, error) {
	return f.FileInfo(), nil
}

// Read implements the io.Reader接口。 md5:ef2823b98664212a
func (f *File) Read(b []byte) (n int, err error) {
	reader, err := f.getReader()
	if err != nil {
		return 0, err
	}
	if n, err = reader.Read(b); err != nil {
		err = gerror.X多层错误并格式化(err, `read content failed`)
	}
	return
}

// Seek 实现了 io.Seeker 接口。 md5:891a8f7e89e80191
func (f *File) Seek(offset int64, whence int) (n int64, err error) {
	reader, err := f.getReader()
	if err != nil {
		return 0, err
	}
	if n, err = reader.Seek(offset, whence); err != nil {
		err = gerror.X多层错误并格式化(err, `seek failed for offset %d, whence %d`, offset, whence)
	}
	return
}

func (f *File) getReader() (*bytes.Reader, error) {
	if f.reader == nil {
		f.reader = bytes.NewReader(f.Content())
	}
	return f.reader, nil
}
