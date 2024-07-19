// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp

// Plugin is the interface for server plugin.
type Plugin interface {
	Name() string//qm:名称  cz:Name() string              // Name returns the name of the plugin.
	Author() string//qm:作者  cz:Author() string            // Author returns the author of the plugin.
	Version() string//qm:X版本  cz:Version() string           // Version returns the version of the plugin, like "v1.0.0".
	Description() string//qm:描述  cz:Description()       // Description returns the description of the plugin.
	Install(s *Server) error // Install installs the plugin BEFORE the server starts.
	Remove() error           // Remove removes the plugin when server shuts down.
}

// Plugin adds plugin to the server.
// ff:插件添加
// s:
// plugin:插件
func (s *Server) Plugin(plugin ...Plugin) {
	s.plugins = append(s.plugins, plugin...)
}
