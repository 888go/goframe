// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package session类

import (
	"context"
	"time"
	
	"github.com/888go/goframe/container/gmap"
)

// Storage 是会话存储的接口定义。
type Storage interface {
// New 创建一个自定义会话ID。
// 该函数可用于创建自定义会话。
	New(ctx context.Context, ttl time.Duration) (sessionId string, err error)

// Get 函数通过给定的键获取并返回特定的 session 值。
// 如果键在 session 中不存在，则返回 nil。
	Get(ctx context.Context, sessionId string, key string) (value interface{}, err error)

	// GetSize 从存储中获取并返回键值对的大小。
	GetSize(ctx context.Context, sessionId string) (size int, err error)

	// Data 从存储中检索所有的键值对并以map形式返回。
	Data(ctx context.Context, sessionId string) (sessionData map[string]interface{}, err error)

// Set 将一个键值对会话数据设置到存储中。
// 参数 `ttl` 指定该会话ID的生存时间（TTL，Time To Live）。
	X设置值(ctx context.Context, sessionId string, key string, value interface{}, ttl time.Duration) error

// SetMap 批量将键值对形式的session设置到存储中。
// 参数 `ttl` 指定该session id的有效期（TTL，Time To Live）。
	SetMap(ctx context.Context, sessionId string, mapData map[string]interface{}, ttl time.Duration) error

	// Remove 从存储中删除指定会话的键值对。
	Remove(ctx context.Context, sessionId string, key string) error

	// RemoveAll 从存储中删除session。
	RemoveAll(ctx context.Context, sessionId string) error

// GetSession 从存储中获取给定会话的 session 数据，并以 `*gmap.StrAnyMap` 类型返回。
//
// 参数 `ttl` 指定了本次会话的生存时间（TTL）。
// 参数 `data` 是当前存储在内存中的旧会话数据，
// 如果禁用了内存存储，对于某些存储方式而言，这个参数可能会是 nil。
//
// 在每次会话启动时都会调用此函数。
// 如果会话不存在或者其 TTL 已经过期，则返回 nil。
	GetSession(ctx context.Context, sessionId string, ttl time.Duration) (*map类.StrAnyMap, error)

// SetSession 更新指定会话 ID 的数据。
// 当发生更改且变为脏状态的会话关闭后，将调用此函数。
// 此函数将内存中所有会话数据映射复制到存储中。
	SetSession(ctx context.Context, sessionId string, sessionData *map类.StrAnyMap, ttl time.Duration) error

// UpdateTTL 更新指定会话ID的TTL（生存时间）。
// 此函数在非脏数据会话关闭后调用。
	UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error
}
