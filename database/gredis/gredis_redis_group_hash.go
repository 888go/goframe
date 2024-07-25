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

// IGroupHash 管理Redis哈希操作。
// 实现参见 redis.GroupHash。 md5:25da2b7d81c6ce3b
type IGroupHash interface {
	HSet(ctx context.Context, key string, fields map[string]interface{}) (int64, error)
	HSetNX(ctx context.Context, key, field string, value interface{}) (int64, error)
	HGet(ctx context.Context, key, field string) (*gvar.Var, error)
	HStrLen(ctx context.Context, key, field string) (int64, error)
	HExists(ctx context.Context, key, field string) (int64, error)
	HDel(ctx context.Context, key string, fields ...string) (int64, error)
	HLen(ctx context.Context, key string) (int64, error)
	HIncrBy(ctx context.Context, key, field string, increment int64) (int64, error)
	HIncrByFloat(ctx context.Context, key, field string, increment float64) (float64, error)
	HMSet(ctx context.Context, key string, fields map[string]interface{}) error
	HMGet(ctx context.Context, key string, fields ...string) (gvar.Vars, error)
	HKeys(ctx context.Context, key string) ([]string, error)
	HVals(ctx context.Context, key string) (gvar.Vars, error)
	HGetAll(ctx context.Context, key string) (*gvar.Var, error)
}
