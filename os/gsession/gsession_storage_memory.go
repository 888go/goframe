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
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcache"
)

// StorageMemory 使用内存实现了会话存储接口。 md5:1a9a78b3bd5a138b
type StorageMemory struct {
	StorageBase
	// cache是用于session TTL的内存数据缓存， 
	// 只有在Storage不存储任何会话数据时才可用（即不同步存储数据）。
	// 请参考StorageFile, StorageMemory和StorageRedis的实现。
	//
	// 其值为`*gmap.StrAnyMap`类型。
	// md5:c8273be50da58f8d
	cache *gcache.Cache
}

// NewStorageMemory 创建并返回一个用于会话的内存存储对象。 md5:9b1b616d48dd808e
func NewStorageMemory() *StorageMemory {
	return &StorageMemory{
		cache: gcache.New(),
	}
}

// RemoveAll 从存储中删除会话。 md5:488d9f9ca747e8e4
func (s *StorageMemory) RemoveAll(ctx context.Context, sessionId string) error {
	_, err := s.cache.Remove(ctx, sessionId)
	return err
}

// GetSession 从存储中根据给定的会话ID获取会话数据，返回一个指向*gmap.StrAnyMap的指针。
//
// 参数`ttl`指定了此会话的有效期，如果超过有效期，则返回nil。参数`data`是当前存储在内存中的旧会话数据，对于某些存储方式，如果禁用了内存存储，它可能会为nil。
//
// 此函数在会话启动时会被调用。
// md5:01e56ce09d5fd934
func (s *StorageMemory) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (*gmap.StrAnyMap, error) {
	// 从管理器中获取内存会话数据。 md5:9a3be5b3f3de62f6
	var (
		v   *gvar.Var
		err error
	)
	v, err = s.cache.Get(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	if v != nil {
		return v.Val().(*gmap.StrAnyMap), nil
	}
	return gmap.NewStrAnyMap(true), nil
}

// SetSession 根据指定的会话ID更新数据映射。
// 当某个被标记为脏（即发生过修改）的会话关闭后，将调用此函数。
// 该操作会将所有会话数据从内存复制到存储中。
// md5:1caa26989d884fa4
func (s *StorageMemory) SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error {
	return s.cache.Set(ctx, sessionId, sessionData, ttl)
}

// UpdateTTL 更新指定会话ID的生存时间（TTL）。
// 当一个未被修改（非脏）的会话关闭后，此函数会被调用。
// 它只是将会话ID添加到异步处理队列中。
// md5:cc5ac287cbbc0eab
func (s *StorageMemory) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error {
	_, err := s.cache.UpdateExpire(ctx, sessionId, ttl)
	return err
}
