// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gredis
import (
	"context"
	"time"
	
	"github.com/888go/goframe/container/gvar"
	)
// IGroupGeneric 管理通用的 Redis 操作。
// 实现请参考 redis.GroupGeneric。
type IGroupGeneric interface {
	Copy(ctx context.Context, source, destination string, option ...CopyOption) (int64, error)
	Exists(ctx context.Context, keys ...string) (int64, error)
	Type(ctx context.Context, key string) (string, error)
	Unlink(ctx context.Context, keys ...string) (int64, error)
	Rename(ctx context.Context, key, newKey string) error
	RenameNX(ctx context.Context, key, newKey string) (int64, error)
	Move(ctx context.Context, key string, db int) (int64, error)
	Del(ctx context.Context, keys ...string) (int64, error)
	RandomKey(ctx context.Context) (string, error)
	DBSize(ctx context.Context) (int64, error)
	Keys(ctx context.Context, pattern string) ([]string, error)
	FlushDB(ctx context.Context, option ...FlushOp) error
	FlushAll(ctx context.Context, option ...FlushOp) error
	Expire(ctx context.Context, key string, seconds int64, option ...ExpireOption) (int64, error)
	ExpireAt(ctx context.Context, key string, time time.Time, option ...ExpireOption) (int64, error)
	ExpireTime(ctx context.Context, key string) (*gvar.Var, error)
	TTL(ctx context.Context, key string) (int64, error)
	Persist(ctx context.Context, key string) (int64, error)
	PExpire(ctx context.Context, key string, milliseconds int64, option ...ExpireOption) (int64, error)
	PExpireAt(ctx context.Context, key string, time time.Time, option ...ExpireOption) (int64, error)
	PExpireTime(ctx context.Context, key string) (*gvar.Var, error)
	PTTL(ctx context.Context, key string) (int64, error)
}

// CopyOption 提供函数 Copy 的选项。
type CopyOption struct {
	DB      int  // DB选项允许为目标键指定一个替代的逻辑数据库索引。
	REPLACE bool // REPLACE 选项在复制值到目标键之前会先移除该目标键。
}

type FlushOp string

const (
	FlushAsync FlushOp = "ASYNC" // 异步: 异步地刷新数据库
	FlushSync  FlushOp = "SYNC"  // SYNC: 同步刷新数据库
)

// ExpireOption 提供了用于 Expire 函数的选项。
type ExpireOption struct {
	NX bool // NX -- 当键未设置过期时间时，才设置过期时间
	XX bool // XX -- 只在键已存在有效期时设置过期时间
	GT bool // GT -- 当新的过期时间大于当前过期时间时，才设置过期时间
	LT bool // LT -- 当新的过期时间小于当前过期时间时，才设置过期时间
}
