// 版权所有 GoFrame gf 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一个。

// Package consts 定义了在整个框架包中共享的常量。
package consts

const (
	ConfigNodeNameDatabase        = "database"
	ConfigNodeNameLogger          = "logger"
	ConfigNodeNameRedis           = "redis"
	ConfigNodeNameViewer          = "viewer"
	ConfigNodeNameServer          = "server"     // 通用版本配置项名称。
	ConfigNodeNameServerSecondary = "httpserver" // 新版本配置项名称支持，从v2开始

// StackFilterKeyForGoFrame 是用于所有 GoFrame 模块路径的堆栈过滤键。
// 例如：.../pkg/mod/github.com/gogf/gf/v2@v2.0.0-20211011134327-54dd11f51122/debug/gdebug/gdebug_caller.go
	StackFilterKeyForGoFrame = "github.com/gogf/gf/"
)
