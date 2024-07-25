// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gres

import (
	"archive/zip"
	"bytes"
	"io"
	"os"

	"github.com/gogf/gf/v2/internal/json"
)

type File struct {
	file     *zip.File
	reader   *bytes.Reader
	resource *Resource
}

// Name返回文件的名称。 md5:a3ea90169ca420db
func (f *File) Name() string {
	return f.file.Name
}

// Open 返回一个 ReadCloser，可以用来访问文件的内容。多个文件可以并发读取。 md5:884ff7d72298ecd8
func (f *File) Open() (io.ReadCloser, error) {
	return f.file.Open()
}

// Content 返回文件的内容。 md5:66d562dea01e0ea1
func (f *File) Content() []byte {
	reader, err := f.Open()
	if err != nil {
		return nil
	}
	defer reader.Close()
	buffer := bytes.NewBuffer(nil)
	if _, err = io.Copy(buffer, reader); err != nil {
		return nil
	}
	return buffer.Bytes()
}

// FileInfo 返回一个表示FileHeader的os.FileInfo。 md5:da797c4560c42771
func (f *File) FileInfo() os.FileInfo {
	return f.file.FileInfo()
}

// Export 将所有子文件递归地导出并保存到指定的系统路径 `dst`。 md5:e85b8976b49230e6
func (f *File) Export(dst string, option ...ExportOption) error {
	return f.resource.Export(f.Name(), dst, option...)
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (f File) MarshalJSON() ([]byte, error) {
	info := f.FileInfo()
	return json.Marshal(map[string]interface{}{
		"name": f.Name(),
		"size": info.Size(),
		"time": info.ModTime(),
		"file": !info.IsDir(),
	})
}
