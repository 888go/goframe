// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gspath 实现了文件索引和对文件夹的搜索功能。

package gspath
import (
	"runtime"
	"strings"
	
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gfsnotify"
	"github.com/888go/goframe/text/gstr"
	)
// updateCacheByPath 递归地将`path`路径下所有文件添加到缓存中。
func (sp *SPath) updateCacheByPath(path string) {
	if sp.cache == nil {
		return
	}
	sp.addToCache(path, path)
}

// formatCacheName 根据以下规则格式化 `name`：
// 1. 统一分隔符为字符 '/'。
// 2. 名称应以 '/' 开头（类似于 HTTP URI）。
func (sp *SPath) formatCacheName(name string) string {
	if runtime.GOOS != "linux" {
		name = gstr.Replace(name, "\\", "/")
	}
	return "/" + strings.Trim(name, "./")
}

// nameFromPath将`filePath`转换为缓存名称。
func (sp *SPath) nameFromPath(filePath, rootPath string) string {
	name := gstr.Replace(filePath, rootPath, "")
	name = sp.formatCacheName(name)
	return name
}

// makeCacheValue将`filePath`格式化为缓存值。
func (sp *SPath) makeCacheValue(filePath string, isDir bool) string {
	if isDir {
		return filePath + "_D_"
	}
	return filePath + "_F_"
}

// parseCacheValue 将缓存值解析为文件路径和类型。
func (sp *SPath) parseCacheValue(value string) (filePath string, isDir bool) {
	if value[len(value)-2 : len(value)-1][0] == 'F' {
		return value[:len(value)-3], false
	}
	return value[:len(value)-3], true
}

// addToCache 将一个项目添加到缓存中。
// 如果 `filePath` 是一个目录，它还会递归地将该目录下的所有子文件和子目录
// 一并添加到缓存中。
func (sp *SPath) addToCache(filePath, rootPath string) {
	// 首先对其自身进行加法操作。
	idDir := gfile.IsDir(filePath)
	sp.cache.SetIfNotExist(
		sp.nameFromPath(filePath, rootPath), sp.makeCacheValue(filePath, idDir),
	)
	// 如果它是一个目录，那么它会添加其所有子文件/子目录。
	if idDir {
		if files, err := gfile.ScanDir(filePath, "*", true); err == nil {
			// 输出到控制台：gspath添加到缓存中: filePath, files
			for _, path := range files {
				sp.cache.SetIfNotExist(sp.nameFromPath(path, rootPath), sp.makeCacheValue(path, gfile.IsDir(path)))
			}
		}
	}
}

// addMonitorByPath 递归地添加gfsnotify监控。
// 当目录下的文件被更新时，缓存会同时得到更新。
// 注意，由于监听器是递归添加的，如果你删除了一个目录，那么该目录下的所有文件（包括目录本身）
// 也会生成删除事件。这意味着，如果一个目录被删除且其下有N个文件，则总共会产生N+1个事件。
func (sp *SPath) addMonitorByPath(path string) {
	if sp.cache == nil {
		return
	}
	_, _ = gfsnotify.Add(path, func(event *gfsnotify.Event) {
		// glog.Debug(event.String()) // 使用glog库输出debug级别的日志，内容为event对象转换为字符串后的结果
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

// removeMonitorByPath 递归地移除对`path`的gfsnotify监控。
func (sp *SPath) removeMonitorByPath(path string) {
	if sp.cache == nil {
		return
	}
	_ = gfsnotify.Remove(path)
}
