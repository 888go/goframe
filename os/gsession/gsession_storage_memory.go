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
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/os/gcache"
)

// StorageMemory 实现了基于内存的 Session 存储接口。
type StorageMemory struct {
	StorageBase
// cache 是用于会话TTL（生存时间）的内存数据缓存，
// 只有在Storage在同步时不存储任何会话数据时才可用。
// 请参考StorageFile、StorageMemory和StorageRedis的具体实现。
//
// 其值的类型为`*gmap.StrAnyMap`。
	cache *缓存类.Cache
}

// NewStorageMemory 创建并返回一个用于存储session的内存文件存储对象。
func NewStorageMemory() *StorageMemory {
	return &StorageMemory{
		cache: 缓存类.X创建(),
	}
}

// RemoveAll 从存储中删除session。
func (s *StorageMemory) RemoveAll(ctx context.Context, sessionId string) error {
	_, err := s.cache.X删除并带返回值(ctx, sessionId)
	return err
}

// GetSession 通过给定的 session id 从存储中获取 session 数据，并以 *gmap.StrAnyMap 类型返回。
//
// 参数 `ttl` 指定了该 session 的生存时间（TTL），若生存时间已过，则返回 nil。
// 参数 `data` 是当前存储在内存中的旧 session 数据，如果禁用了内存存储，对于某些存储方式，此参数可能为 nil。
//
// 当每次 session 开始时，都会调用这个函数。
func (s *StorageMemory) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (*map类.StrAnyMap, error) {
	// 从管理器中检索内存会话数据。
	var (
		v   *泛型类.Var
		err error
	)
	v, err = s.cache.X取值(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	if v != nil {
		return v.X取值().(*map类.StrAnyMap), nil
	}
	return map类.X创建StrAny(true), nil
}

// SetSession 更新指定会话 ID 的数据映射。
// 在每次已标记为脏的、发生改变的会话关闭后，都会调用此函数。
// 此函数将内存中的所有会话数据映射复制到存储中。
func (s *StorageMemory) SetSession(ctx context.Context, sessionId string, sessionData *map类.StrAnyMap, ttl time.Duration) error {
	return s.cache.X设置值(ctx, sessionId, sessionData, ttl)
}

// UpdateTTL 更新指定会话ID的TTL（生存时间）。
// 此函数在非脏数据会话关闭后调用。
// 它只是将该会话ID添加到异步处理队列中。
func (s *StorageMemory) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error {
	_, err := s.cache.X更新过期时间(ctx, sessionId, ttl)
	return err
}
