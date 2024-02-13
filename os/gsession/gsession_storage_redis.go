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
	"github.com/888go/goframe/database/gredis"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/os/gtimer"
)

// StorageRedis 使用 Redis 实现了 Session 存储接口。
type StorageRedis struct {
	StorageBase
	redis         *redis类.Redis   // Redis客户端用于会话存储。
	prefix        string          // Redis中用于session id的键前缀。
	updatingIdMap *map类.StrIntMap // 更新给定会话ID的TTL集合。
}

const (
// DefaultStorageRedisLoopInterval 是更新在最后持续时间段内会话ID的TTL（生存时间）的间隔。
	DefaultStorageRedisLoopInterval = 10 * time.Second
)

// NewStorageRedis 创建并返回一个用于session的redis存储对象。
func NewStorageRedis(redis *redis类.Redis, prefix ...string) *StorageRedis {
	if redis == nil {
		panic("redis instance for storage cannot be empty")
		return nil
	}
	s := &StorageRedis{
		redis:         redis,
		updatingIdMap: map类.X创建StrInt(true),
	}
	if len(prefix) > 0 && prefix[0] != "" {
		s.prefix = prefix[0]
	}
	// 批量及时更新会话ID的TTL（生存时间）
	定时类.X加入单例循环任务(context.Background(), DefaultStorageRedisLoopInterval, func(ctx context.Context) {
		intlog.Print(context.TODO(), "StorageRedis.timer start")
		var (
			err        error
			sessionId  string
			ttlSeconds int
		)
		for {
			if sessionId, ttlSeconds = s.updatingIdMap.X出栈(); sessionId == "" {
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

// RemoveAll 从存储中删除所有键值对。
func (s *StorageRedis) RemoveAll(ctx context.Context, sessionId string) error {
	_, err := s.redis.Del(ctx, s.sessionIdToRedisKey(sessionId))
	return err
}

// GetSession 通过给定的 session id 从存储中获取 session 数据，并以 *gmap.StrAnyMap 类型返回。
//
// 参数 `ttl` 指定了该 session 的生存时间（TTL），若生存时间已过，则返回 nil。
// 参数 `data` 是当前存储在内存中的旧 session 数据，如果禁用了内存存储，对于某些存储方式，此参数可能为 nil。
//
// 当每次 session 开始时，都会调用这个函数。
func (s *StorageRedis) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (*map类.StrAnyMap, error) {
	intlog.Printf(ctx, "StorageRedis.GetSession: %s, %v", sessionId, ttl)
	r, err := s.redis.Get(ctx, s.sessionIdToRedisKey(sessionId))
	if err != nil {
		return nil, err
	}
	content := r.X取字节集()
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
	return map类.X创建AnyStr并从Map(m, true), nil
}

// SetSession 更新指定会话 ID 的数据映射。
// 在每次已标记为脏的、发生改变的会话关闭后，都会调用此函数。
// 此函数将内存中的所有会话数据映射复制到存储中。
func (s *StorageRedis) SetSession(ctx context.Context, sessionId string, sessionData *map类.StrAnyMap, ttl time.Duration) error {
	intlog.Printf(ctx, "StorageRedis.SetSession: %s, %v, %v", sessionId, sessionData, ttl)
	content, err := json.Marshal(sessionData)
	if err != nil {
		return err
	}
	err = s.redis.SetEX(ctx, s.sessionIdToRedisKey(sessionId), content, int64(ttl.Seconds()))
	return err
}

// UpdateTTL 更新指定会话ID的TTL（生存时间）。
// 此函数在非脏数据会话关闭后调用。
// 它只是将该会话ID添加到异步处理队列中。
func (s *StorageRedis) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error {
	intlog.Printf(ctx, "StorageRedis.UpdateTTL: %s, %v", sessionId, ttl)
	if ttl >= DefaultStorageRedisLoopInterval {
		s.updatingIdMap.X设置值(sessionId, int(ttl.Seconds()))
	}
	return nil
}

// doUpdateExpireForSession 更新会话ID的TTL（生存时间）
func (s *StorageRedis) doUpdateExpireForSession(ctx context.Context, sessionId string, ttlSeconds int) error {
	intlog.Printf(ctx, "StorageRedis.doUpdateTTL: %s, %d", sessionId, ttlSeconds)
	_, err := s.redis.Expire(ctx, s.sessionIdToRedisKey(sessionId), int64(ttlSeconds))
	return err
}

// sessionIdToRedisKey 将给定的session id转换并返回其对应的redis键。
func (s *StorageRedis) sessionIdToRedisKey(sessionId string) string {
	return s.prefix + sessionId
}
