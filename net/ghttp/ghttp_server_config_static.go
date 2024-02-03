// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 静态搜索优先级：Resource > ServerPaths > ServerRoot > SearchPath
// （注释翻译：该注释描述了在Go语言代码中执行静态搜索时的资源查找顺序，优先级从高到低依次为Resource、ServerPaths、ServerRoot和SearchPath。）

package ghttp

import (
	"context"
	"strings"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gres"
	"github.com/888go/goframe/util/gconv"
)

// staticPathItem 是静态路径配置项的结构体。
type staticPathItem struct {
	Prefix string // The router URI.
	Path   string // The static path.
}

// SetIndexFiles 设置服务器的索引文件。
func (s *Server) SetIndexFiles(indexFiles []string) {
	s.config.IndexFiles = indexFiles
}

// GetIndexFiles 从服务器获取并返回索引文件。
func (s *Server) GetIndexFiles() []string {
	return s.config.IndexFiles
}

// SetIndexFolder 设置启用或禁用在请求目录时列出子文件。
func (s *Server) SetIndexFolder(enabled bool) {
	s.config.IndexFolder = enabled
}

// SetFileServerEnabled 用于启用/禁用静态文件服务。
// 这是静态文件服务的主要开关。当调用如 SetServerRoot、AddSearchPath 和 AddStaticPath 等静态文件服务配置函数时，
// 此配置会自动启用。
func (s *Server) SetFileServerEnabled(enabled bool) {
	s.config.FileServerEnabled = enabled
}

// SetServerRoot 设置静态服务的文档根目录。
func (s *Server) SetServerRoot(root string) {
	var (
		ctx      = context.TODO()
		realPath = root
	)
	if !gres.Contains(realPath) {
		if p, err := gfile.Search(root); err != nil {
			s.Logger().Fatalf(ctx, `SetServerRoot failed: %+v`, err)
		} else {
			realPath = p
		}
	}
	s.Logger().Debug(ctx, "SetServerRoot path:", realPath)
	s.config.SearchPaths = []string{strings.TrimRight(realPath, gfile.Separator)}
	s.config.FileServerEnabled = true
}

// AddSearchPath 添加静态文件服务的搜索目录路径。
func (s *Server) AddSearchPath(path string) {
	var (
		ctx      = context.TODO()
		realPath = path
	)
	if !gres.Contains(realPath) {
		if p, err := gfile.Search(path); err != nil {
			s.Logger().Fatalf(ctx, `AddSearchPath failed: %+v`, err)
		} else {
			realPath = p
		}
	}
	s.config.SearchPaths = append(s.config.SearchPaths, realPath)
	s.config.FileServerEnabled = true
}

// AddStaticPath 设置静态文件服务的URI到静态目录路径映射。
func (s *Server) AddStaticPath(prefix string, path string) {
	var (
		ctx      = context.TODO()
		realPath = path
	)
	if !gres.Contains(realPath) {
		if p, err := gfile.Search(path); err != nil {
			s.Logger().Fatalf(ctx, `AddStaticPath failed: %+v`, err)
		} else {
			realPath = p
		}
	}
	addItem := staticPathItem{
		Prefix: prefix,
		Path:   realPath,
	}
	if len(s.config.StaticPaths) > 0 {
		s.config.StaticPaths = append(s.config.StaticPaths, addItem)
		// 按前缀长度从短到长对数组进行排序。
		array := garray.NewSortedArray(func(v1, v2 interface{}) int {
			s1 := gconv.String(v1)
			s2 := gconv.String(v2)
			r := len(s2) - len(s1)
			if r == 0 {
				r = strings.Compare(s1, s2)
			}
			return r
		})
		for _, v := range s.config.StaticPaths {
			array.Add(v.Prefix)
		}
		// 根据先前已排序的切片，将项目添加到paths中。
		paths := make([]staticPathItem, 0)
		for _, v := range array.Slice() {
			for _, item := range s.config.StaticPaths {
				if strings.EqualFold(gconv.String(v), item.Prefix) {
					paths = append(paths, item)
					break
				}
			}
		}
		s.config.StaticPaths = paths
	} else {
		s.config.StaticPaths = []staticPathItem{addItem}
	}
	s.config.FileServerEnabled = true
}
