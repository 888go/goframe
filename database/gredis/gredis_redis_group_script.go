// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gredis

import (
	"context"
	
	"github.com/888go/goframe/container/gvar"
)

// IGroupScript 管理 Redis 脚本操作。
// 实现细节参考 redis.GroupScript。
type IGroupScript interface {
	Eval(ctx context.Context, script string, numKeys int64, keys []string, args []interface{}) (*gvar.Var, error)
	EvalSha(ctx context.Context, sha1 string, numKeys int64, keys []string, args []interface{}) (*gvar.Var, error)
	ScriptLoad(ctx context.Context, script string) (string, error)
	ScriptExists(ctx context.Context, sha1 string, sha1s ...string) (map[string]bool, error)
	ScriptFlush(ctx context.Context, option ...ScriptFlushOption) error
	ScriptKill(ctx context.Context) error
}

// ScriptFlushOption 提供了函数 ScriptFlush 的选项。
type ScriptFlushOption struct {
	SYNC  bool // SYNC 同步刷新缓存。
	ASYNC bool // ASYNC 异步刷新缓存。
}
