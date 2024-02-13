// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package redis类

import (
	"context"
	
	"github.com/888go/goframe/container/gvar"
)

// IGroupString 管理 Redis 字符串操作。
// 实现细节请参考 redis.GroupString。
type IGroupString interface {
	X设置值(ctx context.Context, key string, value interface{}, option ...SetOption) (*泛型类.Var, error)
	SetNX(ctx context.Context, key string, value interface{}) (bool, error)
	SetEX(ctx context.Context, key string, value interface{}, ttlInSeconds int64) error
	Get(ctx context.Context, key string) (*泛型类.Var, error)
	GetDel(ctx context.Context, key string) (*泛型类.Var, error)
	GetEX(ctx context.Context, key string, option ...GetEXOption) (*泛型类.Var, error)
	GetSet(ctx context.Context, key string, value interface{}) (*泛型类.Var, error)
	StrLen(ctx context.Context, key string) (int64, error)
	Append(ctx context.Context, key string, value string) (int64, error)
	SetRange(ctx context.Context, key string, offset int64, value string) (int64, error)
	GetRange(ctx context.Context, key string, start, end int64) (string, error)
	Incr(ctx context.Context, key string) (int64, error)
	IncrBy(ctx context.Context, key string, increment int64) (int64, error)
	IncrByFloat(ctx context.Context, key string, increment float64) (float64, error)
	Decr(ctx context.Context, key string) (int64, error)
	DecrBy(ctx context.Context, key string, decrement int64) (int64, error)
	MSet(ctx context.Context, keyValueMap map[string]interface{}) error
	MSetNX(ctx context.Context, keyValueMap map[string]interface{}) (bool, error)
	MGet(ctx context.Context, keys ...string) (map[string]*泛型类.Var, error)
}

// TTLOption 提供了与TTL相关函数的额外选项。
type TTLOption struct {
	EX      *int64 // EX 秒数 -- 设置指定的过期时间，单位为秒。
	PX      *int64 // PX milliseconds -- 设置指定的过期时间，单位为毫秒。
	EXAT    *int64 // EXAT 时间戳-秒 -- 设置键在指定的 Unix 时间（单位：秒）时过期。
	PXAT    *int64 // PXAT 时间戳-毫秒 -- 设置键的过期时间，以毫秒为单位，指定 Unix 时间。
	KeepTTL bool   // 保留与键关联的生存时间（TTL）。
}

// SetOption为Set函数提供额外的选项。
type SetOption struct {
	TTLOption
	NX bool // 如果键尚未存在，则设置该键。
	XX bool // 只有在键已存在的时候才设置该键。

// 如果键存在，则返回该键存储的旧字符串，否则返回nil。
// 若键中存储的值不是字符串，则会返回错误并中止SET操作。
	Get bool
}

// GetEXOption为GetEx函数提供额外的选项。
type GetEXOption struct {
	TTLOption
	Persist bool // Persist -- 移除与该键关联的生存时间（TTL）。
}
