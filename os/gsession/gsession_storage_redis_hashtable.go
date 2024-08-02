// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package session类

import (
	"context"
	"time"

	gmap "github.com/888go/goframe/container/gmap"
	gredis "github.com/888go/goframe/database/gredis"
	"github.com/888go/goframe/internal/intlog"
)

// StorageRedisHashTable 是使用 Redis 哈希表实现的会话存储接口。 md5:4479b82640ee5fc6
type StorageRedisHashTable struct {
	StorageBase
	redis  *gredis.Redis // 用于session存储的Redis客户端。 md5:6ab8fcce48bcdda4
	prefix string        // 会话ID的Redis键前缀。 md5:c0a31dd348ccaac3
}

// NewStorageRedisHashTable 创建并返回一个用于会话的redis哈希表存储对象。 md5:7d5ec78a44d3be11
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

// Get 通过给定的键获取会话值。
// 如果该键不存在于会话中，它将返回nil。
// md5:dd25fb53030b0080
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

// Data 从存储中获取所有的键值对并将其作为映射返回。 md5:7160c6695dcc211b
func (s *StorageRedisHashTable) Data(ctx context.Context, sessionId string) (data map[string]interface{}, err error) {
	m, err := s.redis.HGetAll(ctx, s.sessionIdToRedisKey(sessionId))
	if err != nil {
		return nil, err
	}
	return m.Map(), nil
}

// GetSize 从存储中检索键值对的大小。 md5:9dcc1d87ddc0a989
func (s *StorageRedisHashTable) GetSize(ctx context.Context, sessionId string) (size int, err error) {
	v, err := s.redis.HLen(ctx, s.sessionIdToRedisKey(sessionId))
	return int(v), err
}

// Set 将键值对设置到存储中。
// 参数 `ttl` 指定了会话 ID 的过期时间（而不是键值对）。
// md5:561e667e69e855f6
func (s *StorageRedisHashTable) Set(ctx context.Context, sessionId string, key string, value interface{}, ttl time.Duration) error {
	_, err := s.redis.HSet(ctx, s.sessionIdToRedisKey(sessionId), map[string]interface{}{
		key: value,
	})
	return err
}

// SetMap 使用映射批量设置键值对会话到存储中。
// 参数 `ttl` 指定了会话ID的TTL（并非针对键值对）。
// md5:a1bf3a748ba4aef3
func (s *StorageRedisHashTable) SetMap(ctx context.Context, sessionId string, data map[string]interface{}, ttl time.Duration) error {
	err := s.redis.HMSet(ctx, s.sessionIdToRedisKey(sessionId), data)
	return err
}

// Remove 删除存储中键及其对应的值。 md5:95ea150955b88994
func (s *StorageRedisHashTable) Remove(ctx context.Context, sessionId string, key string) error {
	_, err := s.redis.HDel(ctx, s.sessionIdToRedisKey(sessionId), key)
	return err
}

// RemoveAll 删除存储中的所有键值对。 md5:8b06607595d19a73
func (s *StorageRedisHashTable) RemoveAll(ctx context.Context, sessionId string) error {
	_, err := s.redis.Del(ctx, s.sessionIdToRedisKey(sessionId))
	return err
}

// GetSession 从存储中根据给定的会话ID获取会话数据，返回一个指向*gmap.StrAnyMap的指针。
//
// 参数`ttl`指定了此会话的有效期，如果超过有效期，则返回nil。参数`data`是当前存储在内存中的旧会话数据，对于某些存储方式，如果禁用了内存存储，它可能会为nil。
//
// 此函数在会话启动时会被调用。
// md5:01e56ce09d5fd934
func (s *StorageRedisHashTable) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (*gmap.StrAnyMap, error) {
	intlog.Printf(ctx, "StorageRedisHashTable.GetSession: %s, %v", sessionId, ttl)
	v, err := s.redis.Exists(ctx, s.sessionIdToRedisKey(sessionId))
	if err != nil {
		return nil, err
	}
	if v > 0 {
		// 它不将会话数据存储在内存中，因此返回一个空的映射。
		// 每次都是直接通过Redis服务器获取会话数据项。
		// md5:780013e56e130612
		return gmap.NewStrAnyMap(true), nil
	}
	return nil, nil
}

// SetSession 根据指定的会话ID更新数据映射。
// 当某个被标记为脏（即发生过修改）的会话关闭后，将调用此函数。
// 该操作会将所有会话数据从内存复制到存储中。
// md5:1caa26989d884fa4
func (s *StorageRedisHashTable) SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error {
	intlog.Printf(ctx, "StorageRedisHashTable.SetSession: %s, %v", sessionId, ttl)
	_, err := s.redis.Expire(ctx, s.sessionIdToRedisKey(sessionId), int64(ttl.Seconds()))
	return err
}

// UpdateTTL 更新指定会话ID的生存时间（TTL）。
// 当一个未被修改（非脏）的会话关闭后，此函数会被调用。
// 它只是将会话ID添加到异步处理队列中。
// md5:cc5ac287cbbc0eab
func (s *StorageRedisHashTable) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error {
	intlog.Printf(ctx, "StorageRedisHashTable.UpdateTTL: %s, %v", sessionId, ttl)
	_, err := s.redis.Expire(ctx, s.sessionIdToRedisKey(sessionId), int64(ttl.Seconds()))
	return err
}

// sessionIdToRedisKey 将给定的会话ID转换并返回对应的Redis键。 md5:e18b9b593a10a025
func (s *StorageRedisHashTable) sessionIdToRedisKey(sessionId string) string {
	return s.prefix + sessionId
}
