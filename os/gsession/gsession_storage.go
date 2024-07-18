// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gsession

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/container/gmap"
)

// Storage是会话存储的接口定义。 md5:3c03cfdd3299edcc
type Storage interface {
// New 创建一个自定义会话ID。
// 此函数可用于自定义会话创建。
// md5:bf8b403018c5c6df
	New(ctx context.Context, ttl time.Duration) (sessionId string, err error)

// Get 通过给定的键获取并返回会话中的特定值。
// 如果键在会话中不存在，则返回nil。
// md5:2584a452a5632118
	Get(ctx context.Context, sessionId string, key string) (value interface{}, err error)

	// GetSize 从存储中获取并返回键值对的大小。 md5:2c41726f18e2cd04
	GetSize(ctx context.Context, sessionId string) (size int, err error)

	// Data 从存储中获取所有的键值对并将其作为映射返回。 md5:7160c6695dcc211b
	Data(ctx context.Context, sessionId string) (sessionData map[string]interface{}, err error)

// Set 将一个键值对设置到存储中。
// 参数 `ttl` 指定了会话 ID 的过期时间。
// md5:f141e9b5de211364
	Set(ctx context.Context, sessionId string, key string, value interface{}, ttl time.Duration) error//qm:设置值  cz:Set(ctx context.Context, sessionId string, key string, value interface{}, ttl time.Duration)  yx:true

// SetMap 批量将键值对设置为存储中的会话映射。参数 `ttl` 指定会话 ID 的过期时间。
// md5:be3d6b9412b66e49
	SetMap(ctx context.Context, sessionId string, mapData map[string]interface{}, ttl time.Duration) error

	// Remove 从存储中删除指定会话中的键值对。 md5:3887f6d1acd56ad6
	Remove(ctx context.Context, sessionId string, key string) error

	// RemoveAll 从存储中删除会话。 md5:488d9f9ca747e8e4
	RemoveAll(ctx context.Context, sessionId string) error

// GetSession 从存储中返回给定会话的数据，数据类型为 `*gmap.StrAnyMap`。
//
// 参数 `ttl` 指定了该会话的生存时间（TTL）。
// 参数 `data` 是当前存储在内存中的旧会话数据，对于某些存储，如果禁用了内存存储，此参数可能为 `nil`。
//
// 这个函数会在每次会话开始时被调用。
// 如果会话不存在或者其 TTL 已过期，它将返回 `nil`。
// md5:a495b20f42259c94
	GetSession(ctx context.Context, sessionId string, ttl time.Duration) (*gmap.StrAnyMap, error)

// SetSession 更新指定会话ID的数据。
// 在关闭已更改的会话后，都会调用此函数。这个函数将内存中的所有会话数据映射复制到存储中。
// md5:16766d7e58c61924
	SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error

// UpdateTTL 更新指定会话ID的TTL（时间到 live）。
// 在非脏会话关闭后，将调用此函数。
// md5:29eae01946af2846
	UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error
}
