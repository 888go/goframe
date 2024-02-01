// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gredis
import (
	"context"
	
	"github.com/888go/goframe/container/gvar"
	)
// IGroupHash 管理 Redis hash 操作。
// 实现细节参见 redis.GroupHash。
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
