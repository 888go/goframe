// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gredis

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
)

// IGroupList 管理Redis列表操作。
// 实现了redis.GroupList接口。 md5:25ae44de5ae09921
type IGroupList interface {
	LPush(ctx context.Context, key string, values ...interface{}) (int64, error)
	LPushX(ctx context.Context, key string, element interface{}, elements ...interface{}) (int64, error)
	RPush(ctx context.Context, key string, values ...interface{}) (int64, error)
	RPushX(ctx context.Context, key string, value interface{}) (int64, error)
	LPop(ctx context.Context, key string, count ...int) (*gvar.Var, error)
	RPop(ctx context.Context, key string, count ...int) (*gvar.Var, error)
	LRem(ctx context.Context, key string, count int64, value interface{}) (int64, error)
	LLen(ctx context.Context, key string) (int64, error)
	LIndex(ctx context.Context, key string, index int64) (*gvar.Var, error)
	LInsert(ctx context.Context, key string, op LInsertOp, pivot, value interface{}) (int64, error)
	LSet(ctx context.Context, key string, index int64, value interface{}) (*gvar.Var, error)
	LRange(ctx context.Context, key string, start, stop int64) (gvar.Vars, error)
	LTrim(ctx context.Context, key string, start, stop int64) error
	BLPop(ctx context.Context, timeout int64, keys ...string) (gvar.Vars, error)
	BRPop(ctx context.Context, timeout int64, keys ...string) (gvar.Vars, error)
	RPopLPush(ctx context.Context, source, destination string) (*gvar.Var, error)
	BRPopLPush(ctx context.Context, source, destination string, timeout int64) (*gvar.Var, error)
}

// LInsertOp 定义了 LInsert 函数的操作名称。 md5:c91fe5287aa13a21
type LInsertOp string

const (
	LInsertBefore LInsertOp = "BEFORE"
	LInsertAfter  LInsertOp = "AFTER"
)
