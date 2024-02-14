// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 静态搜索优先级：Resource > ServerPaths > ServerRoot > SearchPath
// （注释翻译：该注释描述了在Go语言代码中执行静态搜索时的资源查找顺序，优先级从高到低依次为Resource、ServerPaths、ServerRoot和SearchPath。）

package http类

import (
	"context"
	"strings"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gres"
	"github.com/888go/goframe/util/gconv"
)

// staticPathItem 是静态路径配置项的结构体。
type 静态文件配置项 struct {
	Prefix string // The router URI.
	Path   string // The static path.
}

// SetIndexFiles 设置服务器的索引文件。
func (s *X服务) X设置静态文件索引(索引 []string) {
	s.config.X静态文件索引 = 索引
}

// GetIndexFiles 从服务器获取并返回索引文件。
func (s *X服务) X取静态文件索引() []string {
	return s.config.X静态文件索引
}

// SetIndexFolder 设置启用或禁用在请求目录时列出子文件。
func (s *X服务) X设置静态文件是否列出子文件(是否 bool) {
	s.config.X静态文件是否列出子文件 = 是否
}

// SetFileServerEnabled 用于启用/禁用静态文件服务。
// 这是静态文件服务的主要开关。当调用如 SetServerRoot、AddSearchPath 和 AddStaticPath 等静态文件服务配置函数时，
// 此配置会自动启用。
func (s *X服务) X设置静态文件是否开启(开启 bool) {
	s.config.X静态文件是否开启 = 开启
}

// SetServerRoot 设置静态服务的文档根目录。
func (s *X服务) X设置静态文件根目录(根目录 string) {
	var (
		ctx      = context.TODO()
		realPath = 根目录
	)
	if !资源类.Contains(realPath) {
		if p, err := 文件类.X查找(根目录); err != nil {
			s.Logger别名().X输出并格式化FATA(ctx, `SetServerRoot failed: %+v`, err)
		} else {
			realPath = p
		}
	}
	s.Logger别名().X输出DEBU(ctx, "SetServerRoot path:", realPath)
	s.config.X静态文件额外搜索目录 = []string{strings.TrimRight(realPath, 文件类.Separator)}
	s.config.X静态文件是否开启 = true
}

// AddSearchPath 添加静态文件服务的搜索目录路径。
func (s *X服务) X静态文件添加额外搜索目录(目录 string) {
	var (
		ctx      = context.TODO()
		realPath = 目录
	)
	if !资源类.Contains(realPath) {
		if p, err := 文件类.X查找(目录); err != nil {
			s.Logger别名().X输出并格式化FATA(ctx, `AddSearchPath failed: %+v`, err)
		} else {
			realPath = p
		}
	}
	s.config.X静态文件额外搜索目录 = append(s.config.X静态文件额外搜索目录, realPath)
	s.config.X静态文件是否开启 = true
}

// AddStaticPath 设置静态文件服务的URI到静态目录路径映射。
func (s *X服务) X静态文件添加目录映射(旧路径 string, 新路径 string) {
	var (
		ctx      = context.TODO()
		realPath = 新路径
	)
	if !资源类.Contains(realPath) {
		if p, err := 文件类.X查找(新路径); err != nil {
			s.Logger别名().X输出并格式化FATA(ctx, `AddStaticPath failed: %+v`, err)
		} else {
			realPath = p
		}
	}
	addItem := 静态文件配置项{
		Prefix: 旧路径,
		Path:   realPath,
	}
	if len(s.config.X静态文件目录映射) > 0 {
		s.config.X静态文件目录映射 = append(s.config.X静态文件目录映射, addItem)
		// 按前缀长度从短到长对数组进行排序。
		array := 数组类.X创建排序(func(v1, v2 interface{}) int {
			s1 := 转换类.String(v1)
			s2 := 转换类.String(v2)
			r := len(s2) - len(s1)
			if r == 0 {
				r = strings.Compare(s1, s2)
			}
			return r
		})
		for _, v := range s.config.X静态文件目录映射 {
			array.X入栈右(v.Prefix)
		}
		// 根据先前已排序的切片，将项目添加到paths中。
		paths := make([]静态文件配置项, 0)
		for _, v := range array.X取切片() {
			for _, item := range s.config.X静态文件目录映射 {
				if strings.EqualFold(转换类.String(v), item.Prefix) {
					paths = append(paths, item)
					break
				}
			}
		}
		s.config.X静态文件目录映射 = paths
	} else {
		s.config.X静态文件目录映射 = []静态文件配置项{addItem}
	}
	s.config.X静态文件是否开启 = true
}
