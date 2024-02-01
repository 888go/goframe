// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gres
import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	
	"github.com/888go/goframe/internal/json"
	)
type File struct {
	file     *zip.File
	reader   *bytes.Reader
	resource *Resource
}

// Name 返回文件的名称。
func (f *File) Name() string {
	return f.file.Name
}

// Open 函数返回一个 ReadCloser 类型的对象，该对象提供了对文件内容的访问权限。
// 允许同时读取多个文件。
func (f *File) Open() (io.ReadCloser, error) {
	return f.file.Open()
}

// Content 返回文件的内容。
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

// FileInfo返回FileHeader对应的os.FileInfo对象。
func (f *File) FileInfo() os.FileInfo {
	return f.file.FileInfo()
}

// Export 递归地将所有子文件导出并保存到指定的系统路径 `dst`。
func (f *File) Export(dst string, option ...ExportOption) error {
	return f.resource.Export(f.Name(), dst, option...)
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (f File) MarshalJSON() ([]byte, error) {
	info := f.FileInfo()
	return json.Marshal(map[string]interface{}{
		"name": f.Name(),
		"size": info.Size(),
		"time": info.ModTime(),
		"file": !info.IsDir(),
	})
}
