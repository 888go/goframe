// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

// Plugin 是服务器插件的接口。
type X插件配置项 interface {
	X名称() string            // Name 返回插件的名称。
	X作者() string          // Author 返回插件的作者。
	X版本() string         // Version 返回插件的版本，如 "v1.0.0"。
	X描述() string     // 描述返回插件的描述。
	Install(s *X服务) error // Install 在服务器启动之前安装插件。
	Remove() error           // Remove 在服务器关闭时移除插件。
}

// Plugin 向服务器添加插件。
func (s *X服务) X插件添加(插件 ...X插件配置项) {
	s.plugins = append(s.plugins, 插件...)
}
