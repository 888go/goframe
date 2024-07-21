// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gredis

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
)

// IGroupGeneric 管理通用的 Redis 操作。
// 实现了 redis.GroupGeneric 接口。
// md5:d6eb4921760b60f4
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
	Scan(ctx context.Context, cursor uint64, option ...ScanOption) (uint64, []string, error)
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

// CopyOption 为 Copy 函数提供选项。 md5:985df0dc4c62e896
type CopyOption struct {
	DB      int  // DB 选项允许为目的地键指定一个替代的逻辑数据库索引。 md5:f7752ecd2c09888e
	REPLACE bool // REPLACE选项在将值复制到目标键之前删除目标键。 md5:7d1daa6e1cf324ab
}

type FlushOp string

const (
	FlushAsync FlushOp = "ASYNC" // ASYNC：异步刷新数据库. md5:8f0fb503842c62dc
	FlushSync  FlushOp = "SYNC"  // 同步：同步刷新数据库. md5:c995019017769085
)

// ExpireOption 提供了 Expire 函数的选项。 md5:fe605b48792fd395
type ExpireOption struct {
	NX bool // NX -- 只在键没有过期时设置过期时间. md5:753349361957bc17
	XX bool // XX -- 只在键已存在过期时间时设置过期. md5:005a0b6114104985
	GT bool // GT -- 仅当新过期时间大于当前过期时间时，才设置过期时间. md5:e25f0e8a00a61ecf
	LT bool // LT -- 只有当新的过期时间小于当前过期时间时才设置过期时间. md5:7d837833fbcaa3f3
}

// ScanOption为Scan函数提供了选项。 md5:32efa528c8a65e49
type ScanOption struct {
	Match string // Match - 定义用于筛选键的通配符风格模式。 md5:8a1fe0030e22d0f9
	Count int    // Count -- 建议每次扫描返回的键的数量。 md5:9090884e4078ad30
	Type  string // Type -- 根据键的数据类型过滤。有效的类型包括 "string"、"list"、"set"、"zset"、"hash" 和 "stream"。 md5:e1661eb01e6db304
}

// doScanOption是ScanOption的内部表示。 md5:3846dba237546aef
type doScanOption struct {
	Match *string
	Count *int
	Type  *string
}

// ToUsedOption 将ScanOption中的零值字段转换为nil。只有具有值的字段才会被保留。 md5:42a6307a3e94db33
// ff:
// scanOpt:
func (scanOpt *ScanOption) ToUsedOption() doScanOption {
	var usedOption doScanOption

	if scanOpt.Match != "" {
		usedOption.Match = &scanOpt.Match
	}
	if scanOpt.Count != 0 {
		usedOption.Count = &scanOpt.Count
	}
	if scanOpt.Type != "" {
		usedOption.Type = &scanOpt.Type
	}

	return usedOption
}
