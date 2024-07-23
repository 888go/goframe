// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gres

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gogf/gf/v2/container/gtree"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
)

type Resource struct {
	tree *gtree.BTree
}

const (
	defaultTreeM = 100
)

// New 创建并返回一个新的资源对象。 md5:8c594601bad2dd64
func New() *Resource {
	return &Resource{
		tree: gtree.NewBTree(defaultTreeM, func(v1, v2 interface{}) int {
			return strings.Compare(v1.(string), v2.(string))
		}),
	}
}

// Add 解包并把`content`添加到当前资源对象中。不必要的参数`prefix`表示每个文件存储在当前资源对象中的前缀。
// md5:93345d9770c1e7fa
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
		r.tree.Set(namePrefix+files[i].file.Name, files[i])
	}
	intlog.Printf(context.TODO(), "Add %d files to resource manager", r.tree.Size())
	return nil
}

// Load 从`path`加载、解包并将数据添加到当前资源对象中。不必要的参数`prefix`表示将每个文件存储到当前资源对象中的前缀。
// md5:ab3e52fa479e7de6
func (r *Resource) Load(path string, prefix ...string) error {
	realPath, err := gfile.Search(path)
	if err != nil {
		return err
	}
	return r.Add(gfile.GetContents(realPath), prefix...)
}

// Get返回给定路径的文件。 md5:f4989a4832cde2d2
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

// GetWithIndex 在给定路径`path`下搜索文件。如果找到的是一个目录，它会在这个目录下索引文件进行搜索。
//
// GetWithIndex 通常用于HTTP静态文件服务中。
// md5:bfb61cc8920b4633
func (r *Resource) GetWithIndex(path string, indexFiles []string) *File {
	// 用于在前缀中替换双字符 '/'。 md5:2ab9f670789bab70
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

// GetContent 直接返回 `path` 的内容。 md5:50cf0f721b7b89a5
func (r *Resource) GetContent(path string) []byte {
	file := r.Get(path)
	if file != nil {
		return file.Content()
	}
	return nil
}

// Contains 检查路径 `path` 是否存在于当前资源对象中。 md5:9beb2e9c06e1e221
func (r *Resource) Contains(path string) bool {
	return r.Get(path) != nil
}

// IsEmpty 检查资源管理器是否为空，并返回结果。 md5:3aaae27781ad4e8c
func (r *Resource) IsEmpty() bool {
	return r.tree.IsEmpty()
}

// ScanDir 在给定路径下返回文件，参数 `path` 应该是一个文件夹类型。
//
// `pattern` 参数支持多个文件名模式，使用逗号 `,` 来分隔多个模式。
//
// 如果 `recursive` 参数为 true，它会递归扫描目录。
//
// 注意，返回的文件不包含给定的 `path`。
// md5:c7e8c1023db3f55f
func (r *Resource) ScanDir(path string, pattern string, recursive ...bool) []*File {
	isRecursive := false
	if len(recursive) > 0 {
		isRecursive = recursive[0]
	}
	return r.doScanDir(path, pattern, isRecursive, false)
}

// ScanDirFile 返回给定`path`下的所有子文件的绝对路径，
// 如果给定的参数`recursive`为true，它会递归地扫描目录。
//
// 注意，它只返回文件，不包括目录。
// md5:0f3154c32271652b
func (r *Resource) ScanDirFile(path string, pattern string, recursive ...bool) []*File {
	isRecursive := false
	if len(recursive) > 0 {
		isRecursive = recursive[0]
	}
	return r.doScanDir(path, pattern, isRecursive, true)
}

// doScanDir 是一个内部方法，用于扫描目录
// 并返回未排序的文件的绝对路径列表。
//
// 模式参数 `pattern` 支持多个文件名模式，
// 使用 ',' 符号来分隔多个模式。
//
// 如果给定的参数 `recursive` 为 true，则会递归地扫描目录。
// md5:9e5185e985fd2bb6
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
	// 用于检查第一个条目的类型。 md5:da747d2102d6a47c
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
		// 为了避免，例如：/i18n 和 /i18n-dir. md5:ab3565cb1db7bc63
		if !first && name[length] != '/' {
			return true
		}
		if !recursive {
			if strings.IndexByte(name[length+1:], '/') != -1 {
				return true
			}
		}
		for _, p := range patterns {
			if match, err := filepath.Match(p, gfile.Basename(name)); err == nil && match {
				files = append(files, value.(*File))
				return true
			}
		}
		return true
	})
	return files
}

// ExportOption 是 Export 函数的选项。 md5:12a5d99e83d743f7
type ExportOption struct {
	RemovePrefix string // 从资源中移除文件名的前缀。 md5:ff1e0af55baecf64
}

// Export 将指定的路径 `srcPath` 及其所有子文件递归导出并保存到指定的系统路径 `dstPath`。 md5:271f4d0f27211419
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
			name = gstr.TrimLeftStr(name, exportOption.RemovePrefix)
		}
		name = gstr.Trim(name, `\/`)
		if name == "" {
			continue
		}
		path = gfile.Join(dst, name)
		if file.FileInfo().IsDir() {
			err = gfile.Mkdir(path)
		} else {
			err = gfile.PutBytes(path, file.Content())
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// Dump 打印当前资源对象的文件。 md5:4533063269cc5df2
func (r *Resource) Dump() {
	var info os.FileInfo
	r.tree.Iterator(func(key, value interface{}) bool {
		info = value.(*File).FileInfo()
		fmt.Printf(
			"%v %8s %s\n",
			gtime.New(info.ModTime()).ISO8601(),
			gfile.FormatSize(info.Size()),
			key,
		)
		return true
	})
	fmt.Printf("TOTAL FILES: %d\n", r.tree.Size())
}
