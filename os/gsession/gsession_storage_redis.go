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
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/os/gtimer"
)

// StorageRedis 使用Redis实现会话存储接口。 md5:df4e47711869aaf9
type StorageRedis struct {
	StorageBase
	redis         *gredis.Redis   // 用于session存储的Redis客户端。 md5:6ab8fcce48bcdda4
	prefix        string          // 会话ID的Redis键前缀。 md5:c0a31dd348ccaac3
	updatingIdMap *gmap.StrIntMap // 更新会话ID的TTL集合。 md5:5eacf4393666b98f
}

const (
// DefaultStorageRedisLoopInterval 是用于在最近一段时间内更新会话ID的TTL（生存时间）的间隔。
// md5:5adbee0aa8ff1658
	DefaultStorageRedisLoopInterval = 10 * time.Second
)

// NewStorageRedis 创建并返回一个用于session的redis存储对象。 md5:58528aab48b7daea
// ff:
// redis:
// prefix:
func NewStorageRedis(redis *gredis.Redis, prefix ...string) *StorageRedis {
	if redis == nil {
		panic("redis instance for storage cannot be empty")
		return nil
	}
	s := &StorageRedis{
		redis:         redis,
		updatingIdMap: gmap.NewStrIntMap(true),
	}
	if len(prefix) > 0 && prefix[0] != "" {
		s.prefix = prefix[0]
	}
	// 定期批量更新会话ID的TTL（时间到 live，生存时间）。 md5:81e845800fad5861
	gtimer.AddSingleton(context.Background(), DefaultStorageRedisLoopInterval, func(ctx context.Context) {
		intlog.Print(context.TODO(), "StorageRedis.timer start")
		var (
			err        error
			sessionId  string
			ttlSeconds int
		)
		for {
			if sessionId, ttlSeconds = s.updatingIdMap.Pop(); sessionId == "" {
				break
			} else {
				if err = s.doUpdateExpireForSession(context.TODO(), sessionId, ttlSeconds); err != nil {
					intlog.Errorf(context.TODO(), `%+v`, err)
				}
			}
		}
		intlog.Print(context.TODO(), "StorageRedis.timer end")
	})
	return s
}

// RemoveAll 删除存储中的所有键值对。 md5:8b06607595d19a73
// ff:
// s:
// ctx:
// sessionId:
func (s *StorageRedis) RemoveAll(ctx context.Context, sessionId string) error {
	_, err := s.redis.Del(ctx, s.sessionIdToRedisKey(sessionId))
	return err
}

// GetSession 从存储中根据给定的会话ID获取会话数据，返回一个指向*gmap.StrAnyMap的指针。
//
// 参数`ttl`指定了此会话的有效期，如果超过有效期，则返回nil。参数`data`是当前存储在内存中的旧会话数据，对于某些存储方式，如果禁用了内存存储，它可能会为nil。
//
// 此函数在会话启动时会被调用。
// md5:01e56ce09d5fd934
// ff:
// s:
// ctx:
// sessionId:
// ttl:
func (s *StorageRedis) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (*gmap.StrAnyMap, error) {
	intlog.Printf(ctx, "StorageRedis.GetSession: %s, %v", sessionId, ttl)
	r, err := s.redis.Get(ctx, s.sessionIdToRedisKey(sessionId))
	if err != nil {
		return nil, err
	}
	content := r.Bytes()
	if len(content) == 0 {
		return nil, nil
	}
	var m map[string]interface{}
	if err = json.UnmarshalUseNumber(content, &m); err != nil {
		return nil, err
	}
	if m == nil {
		return nil, nil
	}
	return gmap.NewStrAnyMapFrom(m, true), nil
}

// SetSession 根据指定的会话ID更新数据映射。
// 当某个被标记为脏（即发生过修改）的会话关闭后，将调用此函数。
// 该操作会将所有会话数据从内存复制到存储中。
// md5:1caa26989d884fa4
// ff:
// s:
// ctx:
// sessionId:
// sessionData:
// ttl:
func (s *StorageRedis) SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error {
	intlog.Printf(ctx, "StorageRedis.SetSession: %s, %v, %v", sessionId, sessionData, ttl)
	content, err := json.Marshal(sessionData)
	if err != nil {
		return err
	}
	err = s.redis.SetEX(ctx, s.sessionIdToRedisKey(sessionId), content, int64(ttl.Seconds()))
	return err
}

// UpdateTTL 更新指定会话ID的生存时间（TTL）。
// 当一个未被修改（非脏）的会话关闭后，此函数会被调用。
// 它只是将会话ID添加到异步处理队列中。
// md5:cc5ac287cbbc0eab
// ff:
// s:
// ctx:
// sessionId:
// ttl:
func (s *StorageRedis) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error {
	intlog.Printf(ctx, "StorageRedis.UpdateTTL: %s, %v", sessionId, ttl)
	if ttl >= DefaultStorageRedisLoopInterval {
		s.updatingIdMap.Set(sessionId, int(ttl.Seconds()))
	}
	return nil
}

// doUpdateExpireForSession 更新会话ID的过期时间。 md5:6a4d0561227fb192
func (s *StorageRedis) doUpdateExpireForSession(ctx context.Context, sessionId string, ttlSeconds int) error {
	intlog.Printf(ctx, "StorageRedis.doUpdateTTL: %s, %d", sessionId, ttlSeconds)
	_, err := s.redis.Expire(ctx, s.sessionIdToRedisKey(sessionId), int64(ttlSeconds))
	return err
}

// sessionIdToRedisKey 将给定的会话ID转换并返回对应的Redis键。 md5:e18b9b593a10a025
func (s *StorageRedis) sessionIdToRedisKey(sessionId string) string {
	return s.prefix + sessionId
}
