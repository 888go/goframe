// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gspath 实现了文件索引和目录搜索功能。
//
// 它按照目录添加的顺序，内部高效地进行文件搜索。
// 注意：
// 如果启用了缓存功能，在添加或删除文件后，会有一个搜索延迟。
package gspath
import (
	"context"
	"os"
	"sort"
	"strings"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/text/gstr"
	)
// SPath 管理路径搜索功能。
type SPath struct {
	paths *garray.StrArray // 搜索目录的数组。
	cache *gmap.StrStrMap  // Searching cache map, it is not enabled if it's nil.
}

// SPathCacheItem 是一个用于搜索的缓存项。
type SPathCacheItem struct {
	path  string // 绝对路径（文件/目录）
	isDir bool   // 是否是目录。
}

var (
	// 对象映射的搜索路径，用于实例管理。
	pathsMap = gmap.NewStrAnyMap(true)
)

// New 创建并返回一个新的路径搜索管理器。
func New(path string, cache bool) *SPath {
	sp := &SPath{
		paths: garray.NewStrArray(true),
	}
	if cache {
		sp.cache = gmap.NewStrStrMap(true)
	}
	if len(path) > 0 {
		if _, err := sp.Add(path); err != nil {
			// intlog.Print(err)
		}
	}
	return sp
}

// Get根据给定路径创建并返回一个搜索管理器实例。
// 参数`cache`用于指定此管理器是否启用缓存功能。
// 如果启用了缓存功能，它会异步地递归扫描该路径，
// 并使用gfsnotify包更新所有子文件/文件夹到缓存中。
func Get(root string, cache bool) *SPath {
	if root == "" {
		root = "/"
	}
	return pathsMap.GetOrSetFuncLock(root, func() interface{} {
		return New(root, cache)
	}).(*SPath)
}

// Search 在路径 `root` 下搜索文件 `name`。
// 参数 `root` 应为一个绝对路径。出于性能考虑，它不会自动将 `root` 转换为绝对路径。
// 可选参数 `indexFiles` 指定了在结果是目录时要搜索的索引文件列表。
// 例如，如果结果 `filePath` 是一个目录，并且 `indexFiles` 是 [index.html, main.html]，那么它还会在 `filePath` 下搜索 [index.html, main.html]。
// 如果找到其中任意一个文件，它将返回该文件的绝对路径，否则返回 `filePath`。
func Search(root string, name string, indexFiles ...string) (filePath string, isDir bool) {
	return Get(root, false).Search(name, indexFiles...)
}

// SearchWithCache 在启用缓存功能的情况下，搜索路径`root`下的文件`name`。
// 参数`root`应为绝对路径。出于性能考虑，它不会自动将`root`转换为绝对路径。
// 可选参数`indexFiles`用于指定当结果是目录时要搜索的索引文件。
// 例如，如果结果`filePath`是一个目录，并且`indexFiles`为[index.html, main.html]，那么它还会在`filePath`下搜索[index.html, main.html]。
// 如果找到其中任何一个文件，则返回该绝对文件路径，否则返回`filePath`。
func SearchWithCache(root string, name string, indexFiles ...string) (filePath string, isDir bool) {
	return Get(root, true).Search(name, indexFiles...)
}

// Set 删除所有其他搜索目录，并为此管理器设置搜索目录。
func (sp *SPath) Set(path string) (realPath string, err error) {
	realPath = gfile.RealPath(path)
	if realPath == "" {
		realPath, _ = sp.Search(path)
		if realPath == "" {
			realPath = gfile.RealPath(gfile.Pwd() + gfile.Separator + path)
		}
	}
	if realPath == "" {
		return realPath, gerror.NewCodef(gcode.CodeInvalidParameter, `path "%s" does not exist`, path)
	}
	// 设置的路径必须是一个目录。
	if gfile.IsDir(realPath) {
		realPath = strings.TrimRight(realPath, gfile.Separator)
		if sp.paths.Search(realPath) != -1 {
			for _, v := range sp.paths.Slice() {
				sp.removeMonitorByPath(v)
			}
		}
		intlog.Print(context.TODO(), "paths clear:", sp.paths)
		sp.paths.Clear()
		if sp.cache != nil {
			sp.cache.Clear()
		}
		sp.paths.Append(realPath)
		sp.updateCacheByPath(realPath)
		sp.addMonitorByPath(realPath)
		return realPath, nil
	} else {
		return "", gerror.NewCode(gcode.CodeInvalidParameter, path+" should be a folder")
	}
}

// Add 向管理器添加更多搜索目录。
// 管理器将按照添加顺序搜索文件。
func (sp *SPath) Add(path string) (realPath string, err error) {
	realPath = gfile.RealPath(path)
	if realPath == "" {
		realPath, _ = sp.Search(path)
		if realPath == "" {
			realPath = gfile.RealPath(gfile.Pwd() + gfile.Separator + path)
		}
	}
	if realPath == "" {
		return realPath, gerror.NewCodef(gcode.CodeInvalidParameter, `path "%s" does not exist`, path)
	}
	// 添加的路径必须是一个目录。
	if gfile.IsDir(realPath) {
// fmt.Println("gspath:", realPath, sp.paths.Search(realPath)) // 输出gspath:（realPath的值），以及sp.paths在realPath路径下搜索的结果
// 同一目录不会被重复添加两次。
		if sp.paths.Search(realPath) < 0 {
			realPath = strings.TrimRight(realPath, gfile.Separator)
			sp.paths.Append(realPath)
			sp.updateCacheByPath(realPath)
			sp.addMonitorByPath(realPath)
		}
		return realPath, nil
	} else {
		return "", gerror.NewCode(gcode.CodeInvalidParameter, path+" should be a folder")
	}
}

// Search 在manager中搜索文件`name`。
// 可选参数`indexFiles`指定了当结果为目录时需要搜索的索引文件。
// 例如，如果结果`filePath`是一个目录，并且`indexFiles`是[index.html, main.html]，
// 它还会在`filePath`下搜索[index.html, main.html]。如果有任何一项被找到，它将返回该绝对文件路径，
// 否则返回`filePath`。
func (sp *SPath) Search(name string, indexFiles ...string) (filePath string, isDir bool) {
	// No cache enabled.
	if sp.cache == nil {
		sp.paths.LockFunc(func(array []string) {
			path := ""
			for _, v := range array {
				path = gfile.Join(v, name)
				if stat, err := os.Stat(path); stat != nil && !os.IsNotExist(err) {
					path = gfile.Abs(path)
					// 安全检查：结果文件路径必须在搜索目录下。
					if len(path) >= len(v) && path[:len(v)] == v {
						filePath = path
						isDir = stat.IsDir()
						break
					}
				}
			}
		})
		if len(indexFiles) > 0 && isDir {
			if name == "/" {
				name = ""
			}
			path := ""
			for _, file := range indexFiles {
				path = filePath + gfile.Separator + file
				if gfile.Exists(path) {
					filePath = path
					isDir = false
					break
				}
			}
		}
		return
	}
	// 使用缓存功能。
	name = sp.formatCacheName(name)
	if v := sp.cache.Get(name); v != "" {
		filePath, isDir = sp.parseCacheValue(v)
		if len(indexFiles) > 0 && isDir {
			if name == "/" {
				name = ""
			}
			for _, file := range indexFiles {
				if v = sp.cache.Get(name + "/" + file); v != "" {
					return sp.parseCacheValue(v)
				}
			}
		}
	}
	return
}

// Remove 从管理器的缓存文件中删除指定`path`。
// 参数`path`可以是绝对路径，也可以只是一个相对文件名。
func (sp *SPath) Remove(path string) {
	if sp.cache == nil {
		return
	}
	if gfile.Exists(path) {
		for _, v := range sp.paths.Slice() {
			name := gstr.Replace(path, v, "")
			name = sp.formatCacheName(name)
			sp.cache.Remove(name)
		}
	} else {
		name := sp.formatCacheName(path)
		sp.cache.Remove(name)
	}
}

// Paths 返回所有搜索目录。
func (sp *SPath) Paths() []string {
	return sp.paths.Slice()
}

// AllPaths 返回缓存在manager中的所有路径。
func (sp *SPath) AllPaths() []string {
	if sp.cache == nil {
		return nil
	}
	paths := sp.cache.Keys()
	if len(paths) > 0 {
		sort.Strings(paths)
	}
	return paths
}

// Size 返回搜索目录的数量。
func (sp *SPath) Size() int {
	return sp.paths.Len()
}
