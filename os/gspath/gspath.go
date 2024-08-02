// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gspath实现了文件夹的索引和搜索功能。
//
// 它按照目录添加顺序，内部高效地进行文件搜索。请注意：
// 如果启用了缓存功能，添加或删除文件后可能会有搜索延迟。
// md5:626b2e878f4df376
package 文件搜索类

import (
	"context"
	"os"
	"sort"
	"strings"

	garray "github.com/888go/goframe/container/garray"
	gmap "github.com/888go/goframe/container/gmap"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	gfile "github.com/888go/goframe/os/gfile"
	gstr "github.com/888go/goframe/text/gstr"
)

// SPath 管理路径搜索功能。 md5:703d47a59dea75cf
type SPath struct {
	paths *garray.StrArray // 搜索目录数组。 md5:1bb898235e976652
	cache *gmap.StrStrMap  // Searching cache map, it is not enabled if it's nil.
}

// SPathCacheItem 是用于搜索的缓存项。 md5:9b45b61130ff4d97
type SPathCacheItem struct {
	path  string // 文件或目录的绝对路径。 md5:8adda623344c67c8
	isDir bool   // Is directory or not.
}

var (
		// 用于实例管理的对象映射的搜索路径。 md5:1eefbdf17ed9097e
	pathsMap = gmap.NewStrAnyMap(true)
)

// New 创建并返回一个新的路径搜索管理器。 md5:4a9d5d03b9c2c8be
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

// Get 创建并返回一个针对给定路径的搜索管理器实例。
// 参数 `cache` 指定是否为此管理器启用缓存功能。
// 如果启用了缓存功能，它将异步且递归地扫描该路径，
// 并使用 gfsnotify 包将所有子文件/文件夹更新到缓存中。
// md5:db411c9b09cbef91
func Get(root string, cache bool) *SPath {
	if root == "" {
		root = "/"
	}
	return pathsMap.GetOrSetFuncLock(root, func() interface{} {
		return New(root, cache)
	}).(*SPath)
}

// Search 在路径 `root` 下搜索文件 `name`。
// 参数 `root` 应为绝对路径。出于性能考虑，它不会自动将 `root` 转换为绝对路径。
// 可选参数 `indexFiles` 指定在结果为目录时要搜索的索引文件。
// 例如，如果结果 `filePath` 是一个目录，并且 `indexFiles` 是 [index.html, main.html]，它将在 `filePath` 下也搜索这两个文件。
// 如果找到其中任何一个文件，它将返回该文件的绝对路径，否则返回 `filePath`。
// md5:dfc991bd35d2d178
func Search(root string, name string, indexFiles ...string) (filePath string, isDir bool) {
	return Get(root, false).Search(name, indexFiles...)
}

// SearchWithCache 在启用缓存的情况下，在路径`root`下搜索文件`name`。参数`root`应为绝对路径，出于性能考虑，它不会自动将`root`转换为绝对路径。
// 可选参数`indexFiles`用于指定当结果为目录时的索引文件。例如，如果结果`filePath`是一个目录，并且`indexFiles`为`[index.html, main.html]`，它将在`filePath`下搜索`[index.html, main.html)`。如果有找到任何文件，则返回其绝对文件路径，否则返回`filePath`。
// md5:f0b25342a4685319
func SearchWithCache(root string, name string, indexFiles ...string) (filePath string, isDir bool) {
	return Get(root, true).Search(name, indexFiles...)
}

// Set删除所有其他搜索目录，并为这个管理器设置搜索目录。 md5:6bb092ed0381b154
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
		// 设置的路径必须是一个目录。 md5:a1d52d7d0583a6ef
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

// Add 向管理器添加更多的搜索目录。
// 管理器将按照添加的顺序查找文件。
// md5:b27b49ecc2f1758a
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
		// 添加的路径必须是一个目录。 md5:3e2662f535c6872c
	if gfile.IsDir(realPath) {
		// fmt.Println("gspath:", realPath, sp.paths.Search(realPath)) 
		// 对于同一个目录，它不会重复添加。
		// md5:701deef87cf571aa
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

// Search 在管理器中搜索文件`name`。
// 可选参数`indexFiles`指定了当搜索结果为目录时，需要在该目录下查找的索引文件列表。
// 例如，如果搜索结果`filePath`是一个目录，并且`indexFiles`是[index.html, main.html]，那么它还会
// 在`filePath`目录下查找[index.html, main.html]。如果找到了其中任何一个文件，就返回其绝对路径；
// 否则，直接返回`filePath`。
// md5:8210196c6e2ae787
func (sp *SPath) Search(name string, indexFiles ...string) (filePath string, isDir bool) {
	// No cache enabled.
	if sp.cache == nil {
		sp.paths.LockFunc(func(array []string) {
			path := ""
			for _, v := range array {
				path = gfile.Join(v, name)
				if stat, err := os.Stat(path); stat != nil && !os.IsNotExist(err) {
					path = gfile.Abs(path)
										// 安全检查：结果文件路径必须在搜索目录下。 md5:b83726d297546baa
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
	// Using cache feature.
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

// Remove 从管理器的缓存文件中删除`path`。参数`path`可以是绝对路径或仅仅是相对文件名。
// md5:30f46aaaf75a1da8
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

// Paths返回所有搜索目录。 md5:0e02e3a85da8e197
func (sp *SPath) Paths() []string {
	return sp.paths.Slice()
}

// AllPaths 返回存储在管理器中的所有路径。 md5:75157edfcae7d2a0
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

// Size 返回搜索目录的数量。 md5:e115dd584d3351a2
func (sp *SPath) Size() int {
	return sp.paths.Len()
}
