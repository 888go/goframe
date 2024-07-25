// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// gspath 包实现了对文件夹的索引和搜索功能。
// md5:04299e0152ee648b

package gspath

import (
	"runtime"
	"strings"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gfsnotify"
	"github.com/gogf/gf/v2/text/gstr"
)

// updateCacheByPath 递归地在`path`下添加所有文件。 md5:ef869f8f30af135a
func (sp *SPath) updateCacheByPath(path string) {
	if sp.cache == nil {
		return
	}
	sp.addToCache(path, path)
}

// formatCacheName 根据以下规则格式化`name`：
// 1. 分隔符统一为字符'/'。
// 2. 名称应以'/'开头（类似于HTTP URI）。
// md5:ed5316ca14ce4d4c
func (sp *SPath) formatCacheName(name string) string {
	if runtime.GOOS != "linux" {
		name = gstr.Replace(name, "\\", "/")
	}
	return "/" + strings.Trim(name, "./")
}

// nameFromPath 将 `filePath` 转换为缓存名称。 md5:5e0f623421b9d54d
func (sp *SPath) nameFromPath(filePath, rootPath string) string {
	name := gstr.Replace(filePath, rootPath, "")
	name = sp.formatCacheName(name)
	return name
}

// makeCacheValue 将 `filePath` 格式化为缓存值。 md5:ac703cee872ac9d4
func (sp *SPath) makeCacheValue(filePath string, isDir bool) string {
	if isDir {
		return filePath + "_D_"
	}
	return filePath + "_F_"
}

// parseCacheValue 解析缓存值为文件路径和类型。 md5:c7d4d6cc498a746f
func (sp *SPath) parseCacheValue(value string) (filePath string, isDir bool) {
	if value[len(value)-2 : len(value)-1][0] == 'F' {
		return value[:len(value)-3], false
	}
	return value[:len(value)-3], true
}

// addToCache 将一个项目添加到缓存中。
// 如果 `filePath` 是一个目录，它还会递归地将所有子文件/目录添加到缓存中。
// md5:836028ec6822544d
func (sp *SPath) addToCache(filePath, rootPath string) {
	// Add itself firstly.
	idDir := gfile.IsDir(filePath)
	sp.cache.SetIfNotExist(
		sp.nameFromPath(filePath, rootPath), sp.makeCacheValue(filePath, idDir),
	)
		// 如果是一个目录，它会添加其下所有的子文件和子目录。 md5:d133c73c85e80b5b
	if idDir {
		if files, err := gfile.ScanDir(filePath, "*", true); err == nil {
						// fmt.Println("将文件路径", filePath, "和文件列表添加到缓存:", files). md5:787f23087852cffe
			for _, path := range files {
				sp.cache.SetIfNotExist(sp.nameFromPath(path, rootPath), sp.makeCacheValue(path, gfile.IsDir(path)))
			}
		}
	}
}

// addMonitorByPath 递归地添加 gfsnotify 监控。
// 当目录下的文件被更新时，缓存也会同时被更新。
// 注意，由于监听是递归添加的，如果你删除一个目录，该目录下的所有文件（包括目录本身）都会触发删除事件，这意味着如果删除了一个包含 N 个文件的目录，总共会产生 N+1 个事件。
// md5:0142c351fd8dd58f
func (sp *SPath) addMonitorByPath(path string) {
	if sp.cache == nil {
		return
	}
	_, _ = gfsnotify.Add(path, func(event *gfsnotify.Event) {
				// 这个Go语言注释的中文翻译是：使用glog库的Debug级别记录event的字符串表示。 md5:9b7b21454414b499
		switch {
		case event.IsRemove():
			sp.cache.Remove(sp.nameFromPath(event.Path, path))

		case event.IsRename():
			if !gfile.Exists(event.Path) {
				sp.cache.Remove(sp.nameFromPath(event.Path, path))
			}

		case event.IsCreate():
			sp.addToCache(event.Path, path)
		}
	}, true)
}

// removeMonitorByPath 递归地移除对 `path` 的 gfsnotify 监控。 md5:8401f0941bb9504c
func (sp *SPath) removeMonitorByPath(path string) {
	if sp.cache == nil {
		return
	}
	_ = gfsnotify.Remove(path)
}
