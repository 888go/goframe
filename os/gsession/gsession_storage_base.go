// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gsession

import (
	"context"
	"time"
	
	"github.com/888go/goframe/container/gmap"
)

// StorageBase 是 Session 存储的一个基础实现。
type StorageBase struct{}

// New 创建一个会话ID。
// 该函数可用于自定义会话创建。
func (s *StorageBase) New(ctx context.Context, ttl time.Duration) (id string, err error) {
	return "", ErrorDisabled
}

// Get 通过给定的键获取特定会话值。
// 如果键在会话中不存在，则返回空值（nil）。
func (s *StorageBase) Get(ctx context.Context, sessionId string, key string) (value interface{}, err error) {
	return nil, ErrorDisabled
}

// Data 从存储中检索所有的键值对并以map形式返回。
func (s *StorageBase) Data(ctx context.Context, sessionId string) (sessionData map[string]interface{}, err error) {
	return nil, ErrorDisabled
}

// GetSize 从存储中检索键值对的大小。
func (s *StorageBase) GetSize(ctx context.Context, sessionId string) (size int, err error) {
	return 0, ErrorDisabled
}

// Set 将键值对会话设置到存储中。
// 参数 `ttl` 指定了会话ID的生存时间（并非针对键值对）。
func (s *StorageBase) Set(ctx context.Context, sessionId string, key string, value interface{}, ttl time.Duration) error {
	return ErrorDisabled
}

// SetMap 批量将键值对集合设置到存储中。
// 参数 `ttl` 指定的是会话ID的TTL（生存时间），而不是键值对的生存时间。
func (s *StorageBase) SetMap(ctx context.Context, sessionId string, mapData map[string]interface{}, ttl time.Duration) error {
	return ErrorDisabled
}

// Remove 从存储中删除指定键及其对应的值。
func (s *StorageBase) Remove(ctx context.Context, sessionId string, key string) error {
	return ErrorDisabled
}

// RemoveAll 从存储中删除session。
func (s *StorageBase) RemoveAll(ctx context.Context, sessionId string) error {
	return ErrorDisabled
}

// GetSession 通过给定的 session id 从存储中获取 session 数据，并以 *gmap.StrAnyMap 类型返回。
//
// 参数 `ttl` 指定了该 session 的生存时间（TTL），若生存时间已过，则返回 nil。
// 参数 `data` 是当前存储在内存中的旧 session 数据，如果禁用了内存存储，对于某些存储方式，此参数可能为 nil。
//
// 当每次 session 开始时，都会调用这个函数。
func (s *StorageBase) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (*gmap.StrAnyMap, error) {
	return nil, ErrorDisabled
}

// SetSession 更新指定会话 ID 的数据映射。
// 在每次已标记为脏的、发生改变的会话关闭后，都会调用此函数。
// 此函数将内存中的所有会话数据映射复制到存储中。
func (s *StorageBase) SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error {
	return ErrorDisabled
}

// UpdateTTL 更新指定会话ID的TTL（生存时间）。
// 此函数在非脏数据会话关闭后调用。
// 它只是将该会话ID添加到异步处理队列中。
func (s *StorageBase) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error {
	return ErrorDisabled
}
