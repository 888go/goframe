// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457

// 包 consts 定义了框架中所有包共享的常量。 md5:29270e2ac9710640
package consts

const (
	ConfigNodeNameDatabase        = "database"
	ConfigNodeNameLogger          = "logger"
	ConfigNodeNameRedis           = "redis"
	ConfigNodeNameViewer          = "viewer"
	ConfigNodeNameServer          = "server"     // 通用版本配置项名称。 md5:2873fdceb78be9aa
	ConfigNodeNameServerSecondary = "httpserver" // 从v2版本开始支持新的配置项名称. md5:b302e38bb84071a9

	// StackFilterKeyForGoFrame 是用于过滤所有 GoFrame 模块路径的栈追踪键。
	// 例如：.../pkg/mod/github.com/gogf/gf/v2@v2.0.0-20211011134327-54dd11f51122/debug/gdebug/gdebug_caller.go
	// md5:3b4239b7b32969a5
	StackFilterKeyForGoFrame = "github.com/gogf/gf/"
)
