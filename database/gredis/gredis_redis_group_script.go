// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gredis

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
)

// IGroupScript 管理 redis 脚本操作。
// 实现了 redis.GroupScript 接口。
// md5:e1fe409124b24bad
type IGroupScript interface {
	Eval(ctx context.Context, script string, numKeys int64, keys []string, args []interface{}) (*gvar.Var, error)
	EvalSha(ctx context.Context, sha1 string, numKeys int64, keys []string, args []interface{}) (*gvar.Var, error)
	ScriptLoad(ctx context.Context, script string) (string, error)
	ScriptExists(ctx context.Context, sha1 string, sha1s ...string) (map[string]bool, error)
	ScriptFlush(ctx context.Context, option ...ScriptFlushOption) error
	ScriptKill(ctx context.Context) error
}

// ScriptFlushOption 是 ScriptFlush 函数的选项。. md5:dc79d8420cb1e3bb
type ScriptFlushOption struct {
	SYNC  bool // SYNC 将缓存同步刷新。. md5:199c25e221e4fff0
	ASYNC bool // ASYNC 异步刷新缓存。. md5:cc79b38230559523
}
