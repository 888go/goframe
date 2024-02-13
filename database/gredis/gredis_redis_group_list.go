// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package redis类

import (
	"context"
	
	"github.com/888go/goframe/container/gvar"
)

// IGroupList 管理 Redis 列表操作。
// 实现细节参见 redis.GroupList。
type IGroupList interface {
	LPush(ctx context.Context, key string, values ...interface{}) (int64, error)
	LPushX(ctx context.Context, key string, element interface{}, elements ...interface{}) (int64, error)
	RPush(ctx context.Context, key string, values ...interface{}) (int64, error)
	RPushX(ctx context.Context, key string, value interface{}) (int64, error)
	LPop(ctx context.Context, key string, count ...int) (*泛型类.Var, error)
	RPop(ctx context.Context, key string, count ...int) (*泛型类.Var, error)
	LRem(ctx context.Context, key string, count int64, value interface{}) (int64, error)
	LLen(ctx context.Context, key string) (int64, error)
	LIndex(ctx context.Context, key string, index int64) (*泛型类.Var, error)
	LInsert(ctx context.Context, key string, op LInsertOp, pivot, value interface{}) (int64, error)
	LSet(ctx context.Context, key string, index int64, value interface{}) (*泛型类.Var, error)
	LRange(ctx context.Context, key string, start, stop int64) (泛型类.Vars, error)
	LTrim(ctx context.Context, key string, start, stop int64) error
	BLPop(ctx context.Context, timeout int64, keys ...string) (泛型类.Vars, error)
	BRPop(ctx context.Context, timeout int64, keys ...string) (泛型类.Vars, error)
	RPopLPush(ctx context.Context, source, destination string) (*泛型类.Var, error)
	BRPopLPush(ctx context.Context, source, destination string, timeout int64) (*泛型类.Var, error)
}

// LInsertOp 定义了函数 LInsert 的操作名称。
type LInsertOp string

const (
	LInsertBefore LInsertOp = "BEFORE"
	LInsertAfter  LInsertOp = "AFTER"
)
