// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package redis类

import (
	"context"

	gvar "github.com/888go/goframe/container/gvar"
)

// IGroupString 管理 Redis 字符串操作。
// 实现了 redis.GroupString。
// md5:1b6f861ea35b113e
type IGroupString interface {
	Set(ctx context.Context, key string, value interface{}, option ...SetOption) (*gvar.Var, error)
	SetNX(ctx context.Context, key string, value interface{}) (bool, error)
	SetEX(ctx context.Context, key string, value interface{}, ttlInSeconds int64) error
	Get(ctx context.Context, key string) (*gvar.Var, error)
	GetDel(ctx context.Context, key string) (*gvar.Var, error)
	GetEX(ctx context.Context, key string, option ...GetEXOption) (*gvar.Var, error)
	GetSet(ctx context.Context, key string, value interface{}) (*gvar.Var, error)
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
	MGet(ctx context.Context, keys ...string) (map[string]*gvar.Var, error)
}

// TTLOption 为与TTL（生存时间）相关的功能提供额外的选项。 md5:806d2a9bb0ba8e2c
type TTLOption struct {
	EX      *int64 // EX seconds -- 设置指定的过期时间，以秒为单位。 md5:c9b922a2aed5b01f
	PX      *int64 // PX 毫秒 - 设置指定的过期时间，以毫秒为单位。 md5:7355e125322e02a7
	EXAT    *int64 // EXAT 时间戳-秒 - 设置键将在多少秒后过期的指定Unix时间。 md5:e9439dac9f4a9efa
	PXAT    *int64 // PXAT 时间戳-毫秒 -- 设置键将在指定的Unix时间（以毫秒为单位）过期。 md5:0342540fbf12c73c
	KeepTTL bool   // 保留与键相关联的生存时间。 md5:b504515edfac8f85
}

// SetOption为Set函数提供额外选项。 md5:c920a4cbd42cf35d
type SetOption struct {
	TTLOption
	NX bool // 只有在键不存在时才设置它。 md5:29aba97b33955575
	XX bool // 仅当键已存在时才设置该键。 md5:919ca2b47e92656f

	// 返回key存储的旧字符串，如果key不存在则返回nil。
	// 如果key存储的值不是字符串，将返回错误并中止SET操作。
	// md5:0932521a97426d54
	Get bool
}

// GetEXOption 为 GetEx 函数提供额外选项。 md5:853b343734200902
type GetEXOption struct {
	TTLOption
	Persist bool // Persist -- 删除与键关联的过期时间。 md5:68ca7740bd8bdb7c
}
