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
	"github.com/888go/goframe/database/gredis"
	"github.com/888go/goframe/internal/intlog"
)

// StorageRedisHashTable 实现了使用 Redis 哈希表的 Session 存储接口。
type StorageRedisHashTable struct {
	StorageBase
	redis  *gredis.Redis // Redis客户端用于会话存储。
	prefix string        // Redis中用于session id的键前缀。
}

// NewStorageRedisHashTable 创建并返回一个用于存储session的redis哈希表存储对象。
func NewStorageRedisHashTable(redis *gredis.Redis, prefix ...string) *StorageRedisHashTable {
	if redis == nil {
		panic("redis instance for storage cannot be empty")
		return nil
	}
	s := &StorageRedisHashTable{
		redis: redis,
	}
	if len(prefix) > 0 && prefix[0] != "" {
		s.prefix = prefix[0]
	}
	return s
}

// Get 通过给定的键获取 session 值。
// 如果键在 session 中不存在，则返回 nil。
func (s *StorageRedisHashTable) Get(ctx context.Context, sessionId string, key string) (value interface{}, err error) {
	v, err := s.redis.HGet(ctx, s.sessionIdToRedisKey(sessionId), key)
	if err != nil {
		return nil, err
	}
	if v.IsNil() {
		return nil, nil
	}
	return v.String(), nil
}

// Data 从存储中检索所有的键值对并以map形式返回。
func (s *StorageRedisHashTable) Data(ctx context.Context, sessionId string) (data map[string]interface{}, err error) {
	m, err := s.redis.HGetAll(ctx, s.sessionIdToRedisKey(sessionId))
	if err != nil {
		return nil, err
	}
	return m.Map(), nil
}

// GetSize 从存储中检索键值对的大小。
func (s *StorageRedisHashTable) GetSize(ctx context.Context, sessionId string) (size int, err error) {
	v, err := s.redis.HLen(ctx, s.sessionIdToRedisKey(sessionId))
	return int(v), err
}

// Set 将键值对会话设置到存储中。
// 参数 `ttl` 指定了会话ID的生存时间（并非针对键值对）。
func (s *StorageRedisHashTable) Set(ctx context.Context, sessionId string, key string, value interface{}, ttl time.Duration) error {
	_, err := s.redis.HSet(ctx, s.sessionIdToRedisKey(sessionId), map[string]interface{}{
		key: value,
	})
	return err
}

// SetMap 批量将键值对集合设置到存储中。
// 参数 `ttl` 指定的是会话ID的TTL（生存时间），而不是键值对的生存时间。
func (s *StorageRedisHashTable) SetMap(ctx context.Context, sessionId string, data map[string]interface{}, ttl time.Duration) error {
	err := s.redis.HMSet(ctx, s.sessionIdToRedisKey(sessionId), data)
	return err
}

// Remove 从存储中删除指定键及其对应的值。
func (s *StorageRedisHashTable) Remove(ctx context.Context, sessionId string, key string) error {
	_, err := s.redis.HDel(ctx, s.sessionIdToRedisKey(sessionId), key)
	return err
}

// RemoveAll 从存储中删除所有键值对。
func (s *StorageRedisHashTable) RemoveAll(ctx context.Context, sessionId string) error {
	_, err := s.redis.Del(ctx, s.sessionIdToRedisKey(sessionId))
	return err
}

// GetSession 通过给定的 session id 从存储中获取 session 数据，并以 *gmap.StrAnyMap 类型返回。
//
// 参数 `ttl` 指定了该 session 的生存时间（TTL），若生存时间已过，则返回 nil。
// 参数 `data` 是当前存储在内存中的旧 session 数据，如果禁用了内存存储，对于某些存储方式，此参数可能为 nil。
//
// 当每次 session 开始时，都会调用这个函数。
func (s *StorageRedisHashTable) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (*gmap.StrAnyMap, error) {
	intlog.Printf(ctx, "StorageRedisHashTable.GetSession: %s, %v", sessionId, ttl)
	v, err := s.redis.Exists(ctx, s.sessionIdToRedisKey(sessionId))
	if err != nil {
		return nil, err
	}
	if v > 0 {
// 它并不在内存中存储会话数据，因此返回一个空映射。
// 每次都直接通过 Redis 服务器获取会话数据项。
		return gmap.NewStrAnyMap(true), nil
	}
	return nil, nil
}

// SetSession 更新指定会话 ID 的数据映射。
// 在每次已标记为脏的、发生改变的会话关闭后，都会调用此函数。
// 此函数将内存中的所有会话数据映射复制到存储中。
func (s *StorageRedisHashTable) SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error {
	intlog.Printf(ctx, "StorageRedisHashTable.SetSession: %s, %v", sessionId, ttl)
	_, err := s.redis.Expire(ctx, s.sessionIdToRedisKey(sessionId), int64(ttl.Seconds()))
	return err
}

// UpdateTTL 更新指定会话ID的TTL（生存时间）。
// 此函数在非脏数据会话关闭后调用。
// 它只是将该会话ID添加到异步处理队列中。
func (s *StorageRedisHashTable) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error {
	intlog.Printf(ctx, "StorageRedisHashTable.UpdateTTL: %s, %v", sessionId, ttl)
	_, err := s.redis.Expire(ctx, s.sessionIdToRedisKey(sessionId), int64(ttl.Seconds()))
	return err
}

// sessionIdToRedisKey 将给定的session id转换并返回其对应的redis键。
func (s *StorageRedisHashTable) sessionIdToRedisKey(sessionId string) string {
	return s.prefix + sessionId
}
