// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

// 静态搜索优先级：资源 > 服务器路径 > 服务器根目录 > 搜索路径. md5:57bfdcb1a5b6e018

package ghttp

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/util/gconv"
)

// staticPathItem 是静态路径配置的结构体项。 md5:662c2fcf8901cd8a
type staticPathItem struct {
	Prefix string // The router URI.
	Path   string // The static path.
}

// SetIndexFiles 设置服务器的索引文件。 md5:9fa01b4418733cc9
func (s *Server) SetIndexFiles(indexFiles []string) {
	s.config.IndexFiles = indexFiles
}

// GetIndexFiles 从服务器检索并返回索引文件。 md5:4a8b46856576d1ee
func (s *Server) GetIndexFiles() []string {
	return s.config.IndexFiles
}

// SetIndexFolder 用于启用或禁用在请求目录时列出子文件的功能。 md5:68a17005c9c4e9c6
func (s *Server) SetIndexFolder(enabled bool) {
	s.config.IndexFolder = enabled
}

// SetFileServerEnabled 启用或禁用静态文件服务。
// 这是静态文件服务的主要开关。当调用如 SetServerRoot、AddSearchPath 和 AddStaticPath 等静态文件服务配置函数时，此配置会自动启用。 md5:62ef61e18a481245
func (s *Server) SetFileServerEnabled(enabled bool) {
	s.config.FileServerEnabled = enabled
}

// SetServerRoot 设置静态服务的文档根。 md5:a2b38f0b2614dd83
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

// AddSearchPath 为静态文件服务添加搜索目录路径。 md5:cd47be9ef3e2898b
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

// AddStaticPath 为静态文件服务设置URI到静态目录路径的映射。 md5:d4831b3f2ef706f4
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
		// 按照前缀长度从短到长对数组进行排序。 md5:4b9918a40946ebb8
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
		// 将项目添加到按照之前排序的切片中到paths。 md5:1bd48d981558718c
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
