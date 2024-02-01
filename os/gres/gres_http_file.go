// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gres
import (
	"bytes"
	"os"
	
	"github.com/888go/goframe/errors/gerror"
	)
// Close 实现了 http.File 接口的方法。
func (f *File) Close() error {
	return nil
}

// Readdir 实现了 http.File 接口中的 Readdir 方法。
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

// Stat 实现了 http.File 接口的 Stat 方法。
func (f *File) Stat() (os.FileInfo, error) {
	return f.FileInfo(), nil
}

// Read 实现了 io.Reader 接口。
func (f *File) Read(b []byte) (n int, err error) {
	reader, err := f.getReader()
	if err != nil {
		return 0, err
	}
	if n, err = reader.Read(b); err != nil {
		err = gerror.Wrapf(err, `read content failed`)
	}
	return
}

// Seek 实现了 io.Seeker 接口。
func (f *File) Seek(offset int64, whence int) (n int64, err error) {
	reader, err := f.getReader()
	if err != nil {
		return 0, err
	}
	if n, err = reader.Seek(offset, whence); err != nil {
		err = gerror.Wrapf(err, `seek failed for offset %d, whence %d`, offset, whence)
	}
	return
}

func (f *File) getReader() (*bytes.Reader, error) {
	if f.reader == nil {
		f.reader = bytes.NewReader(f.Content())
	}
	return f.reader, nil
}
