// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghttp

// Plugin是服务器插件的接口。 md5:f625001233bdb03a
type Plugin interface {
	Name() string//qm:名称  cz:Name() string              // Name 返回插件的名称。 md5:2db3d61fbe59c133
	Author() string//qm:作者  cz:Author() string            // Author返回插件的作者。 md5:1b32b0dc8d2dde4d
	Version() string//qm:X版本  cz:Version() string           // Version 返回插件的版本号，例如 "v1.0.0"。 md5:d00cdffa2274f882
	Description() string//qm:描述  cz:Description()       // Description 返回插件的描述。 md5:b3440b3816c4f5af
	Install(s *Server) error // Install 在服务器启动之前安装插件。 md5:8ba67273b69cfbb3
	Remove() error           // Remove 在服务器关闭时移除插件。 md5:873b1d7c56ba7013
}

// Plugin 向服务器添加插件。 md5:381028a4fb84b7ca
// ff:插件添加
// s:
// plugin:插件
func (s *Server) Plugin(plugin ...Plugin) {
	s.plugins = append(s.plugins, plugin...)
}
