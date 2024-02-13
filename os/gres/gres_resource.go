// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 资源类

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	
	"github.com/888go/goframe/container/gtree"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/text/gstr"
)

type Resource struct {
	tree *树形类.BTree
}

const (
	defaultTreeM = 100
)

// New 创建并返回一个新的资源对象。
func New() *Resource {
	return &Resource{
		tree: 树形类.NewBTree(defaultTreeM, func(v1, v2 interface{}) int {
			return strings.Compare(v1.(string), v2.(string))
		}),
	}
}

// Add 方法对`content`进行解包并将其添加到当前资源对象中。
// 不必要的参数`prefix`表示存储到当前资源对象时，每个文件的前缀。
func (r *Resource) Add(content string, prefix ...string) error {
	files, err := UnpackContent(content)
	if err != nil {
		intlog.Printf(context.TODO(), "Add resource files failed: %v", err)
		return err
	}
	namePrefix := ""
	if len(prefix) > 0 {
		namePrefix = prefix[0]
	}
	for i := 0; i < len(files); i++ {
		files[i].resource = r
		r.tree.X设置值(namePrefix+files[i].file.Name, files[i])
	}
	intlog.Printf(context.TODO(), "Add %d files to resource manager", r.tree.Size())
	return nil
}

// Load 从`path`加载、解压并将数据添加到当前资源对象中。
// 不必要的参数`prefix`表示存储到当前资源对象中的每个文件的前缀。
func (r *Resource) Load(path string, prefix ...string) error {
	realPath, err := 文件类.X查找(path)
	if err != nil {
		return err
	}
	return r.Add(文件类.X读文本(realPath), prefix...)
}

// Get 返回指定路径的文件。
func (r *Resource) Get(path string) *File {
	if path == "" {
		return nil
	}
	path = strings.ReplaceAll(path, "\\", "/")
	path = strings.ReplaceAll(path, "//", "/")
	if path != "/" {
		for path[len(path)-1] == '/' {
			path = path[:len(path)-1]
		}
	}
	result := r.tree.Get(path)
	if result != nil {
		return result.(*File)
	}
	return nil
}

// GetWithIndex 搜索指定 `path` 的文件，如果该文件是一个目录，
// 则进一步在该目录下进行索引文件的搜索。
//
// GetWithIndex 通常用于 HTTP 静态文件服务。
func (r *Resource) GetWithIndex(path string, indexFiles []string) *File {
	// 用于在前缀中替换双字符 '/'
	path = strings.ReplaceAll(path, "\\", "/")
	path = strings.ReplaceAll(path, "//", "/")
	if path != "/" {
		for path[len(path)-1] == '/' {
			path = path[:len(path)-1]
		}
	}
	if file := r.Get(path); file != nil {
		if len(indexFiles) > 0 && file.FileInfo().IsDir() {
			var f *File
			for _, name := range indexFiles {
				if f = r.Get(path + "/" + name); f != nil {
					return f
				}
			}
		}
		return file
	}
	return nil
}

// GetContent直接返回`path`的内容。
func (r *Resource) GetContent(path string) []byte {
	file := r.Get(path)
	if file != nil {
		return file.Content()
	}
	return nil
}

// Contains 检查当前资源对象中是否存在 `path`。
func (r *Resource) Contains(path string) bool {
	return r.Get(path) != nil
}

// IsEmpty 检查并返回资源管理器是否为空。
func (r *Resource) IsEmpty() bool {
	return r.tree.IsEmpty()
}

// ScanDir 返回给定路径下的文件，参数`path`应为文件夹类型。
//
// 参数`pattern`支持多个文件名模式，
// 使用','符号来分隔多个模式。
//
// 如果给定参数`recursive`为真，则会递归扫描目录。
//
// 注意，返回的文件列表中不包含给定的参数`path`所代表的目录自身。
func (r *Resource) ScanDir(path string, pattern string, recursive ...bool) []*File {
	isRecursive := false
	if len(recursive) > 0 {
		isRecursive = recursive[0]
	}
	return r.doScanDir(path, pattern, isRecursive, false)
}

// ScanDirFile 返回给定 `path` 下所有子文件的绝对路径，
// 若给定参数 `recursive` 为 true，则会递归扫描目录。
//
// 注意，该函数仅返回文件，不包括目录。
func (r *Resource) ScanDirFile(path string, pattern string, recursive ...bool) []*File {
	isRecursive := false
	if len(recursive) > 0 {
		isRecursive = recursive[0]
	}
	return r.doScanDir(path, pattern, isRecursive, true)
}

// doScanDir 是一个内部方法，用于扫描目录并返回一个未排序的文件绝对路径列表。
//
// 参数`pattern`支持多个文件名模式，使用','符号分隔多个模式。
//
// 如果给定参数`recursive`为真，则会递归地扫描目录。
func (r *Resource) doScanDir(path string, pattern string, recursive bool, onlyFile bool) []*File {
	path = strings.ReplaceAll(path, "\\", "/")
	path = strings.ReplaceAll(path, "//", "/")
	if path != "/" {
		for path[len(path)-1] == '/' {
			path = path[:len(path)-1]
		}
	}
	var (
		name     = ""
		files    = make([]*File, 0)
		length   = len(path)
		patterns = strings.Split(pattern, ",")
	)
	for i := 0; i < len(patterns); i++ {
		patterns[i] = strings.TrimSpace(patterns[i])
	}
	// 用于对第一个条目的类型检查。
	first := true
	r.tree.IteratorFrom(path, true, func(key, value interface{}) bool {
		if first {
			if !value.(*File).FileInfo().IsDir() {
				return false
			}
			first = false
		}
		if onlyFile && value.(*File).FileInfo().IsDir() {
			return true
		}
		name = key.(string)
		if len(name) <= length {
			return true
		}
		if path != name[:length] {
			return false
		}
		// 为避免出现诸如/i18n和/i18n-dir这样的情况
		if !first && name[length] != '/' {
			return true
		}
		if !recursive {
			if strings.IndexByte(name[length+1:], '/') != -1 {
				return true
			}
		}
		for _, p := range patterns {
			if match, err := filepath.Match(p, 文件类.X路径取文件名(name)); err == nil && match {
				files = append(files, value.(*File))
				return true
			}
		}
		return true
	})
	return files
}

// ExportOption 是函数 Export 的选项。
type ExportOption struct {
	RemovePrefix string // 从资源中移除文件名前缀
}

// Export 递归地导出并保存指定路径`srcPath`及其所有子文件到指定系统路径`dstPath`。
func (r *Resource) Export(src, dst string, option ...ExportOption) error {
	var (
		err          error
		name         string
		path         string
		exportOption ExportOption
		files        []*File
	)

	if r.Get(src).FileInfo().IsDir() {
		files = r.doScanDir(src, "*", true, false)
	} else {
		files = append(files, r.Get(src))
	}

	if len(option) > 0 {
		exportOption = option[0]
	}
	for _, file := range files {
		name = file.Name()
		if exportOption.RemovePrefix != "" {
			name = 文本类.X过滤首字符(name, exportOption.RemovePrefix)
		}
		name = 文本类.X过滤首尾符并含空白(name, `\/`)
		if name == "" {
			continue
		}
		path = 文件类.X路径生成(dst, name)
		if file.FileInfo().IsDir() {
			err = 文件类.X创建目录(path)
		} else {
			err = 文件类.X写入字节集(path, file.Content())
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// Dump 打印当前资源对象的文件。
func (r *Resource) Dump() {
	var info os.FileInfo
	r.tree.X遍历(func(key, value interface{}) bool {
		info = value.(*File).FileInfo()
		fmt.Printf(
			"%v %8s %s\n",
			时间类.X创建(info.ModTime()).X取文本时间ISO8601(),
			文件类.X字节长度转易读格式(info.Size()),
			key,
		)
		return true
	})
	fmt.Printf("TOTAL FILES: %d\n", r.tree.Size())
}
