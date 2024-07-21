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

// StorageBase是会话存储的基本实现。 md5:9a65ccca10de1608
type StorageBase struct{}

// New 创建一个会话 ID。
// 此函数可用于自定义会话创建。
// md5:ffcd61f72bd1d22b
// ff:
// s:
// ctx:
// ttl:
// id:
// err:
func (s *StorageBase) New(ctx context.Context, ttl time.Duration) (id string, err error) {
	return "", ErrorDisabled
}

// Get 使用给定的键获取会话中的特定值。
// 如果键在会话中不存在，则返回nil。
// md5:c1696c0fb72c680b
// ff:
// s:
// ctx:
// sessionId:
// key:
// value:
// err:
func (s *StorageBase) Get(ctx context.Context, sessionId string, key string) (value interface{}, err error) {
	return nil, ErrorDisabled
}

// Data 从存储中获取所有的键值对并将其作为映射返回。 md5:7160c6695dcc211b
// ff:
// s:
// ctx:
// sessionId:
// sessionData:
// err:
func (s *StorageBase) Data(ctx context.Context, sessionId string) (sessionData map[string]interface{}, err error) {
	return nil, ErrorDisabled
}

// GetSize 从存储中检索键值对的大小。 md5:9dcc1d87ddc0a989
// ff:
// s:
// ctx:
// sessionId:
// size:
// err:
func (s *StorageBase) GetSize(ctx context.Context, sessionId string) (size int, err error) {
	return 0, ErrorDisabled
}

// Set 将键值对设置到存储中。
// 参数 `ttl` 指定了会话 ID 的过期时间（而不是键值对）。
// md5:561e667e69e855f6
// yx:true
// ff:设置值
// s:
// ctx:
// sessionId:
// key:
// value:
// ttl:
func (s *StorageBase) Set(ctx context.Context, sessionId string, key string, value interface{}, ttl time.Duration) error {
	return ErrorDisabled
}

// SetMap 使用映射批量设置键值对会话到存储中。
// 参数 `ttl` 指定了会话ID的TTL（并非针对键值对）。
// md5:a1bf3a748ba4aef3
// ff:
// s:
// ctx:
// sessionId:
// mapData:
// ttl:
func (s *StorageBase) SetMap(ctx context.Context, sessionId string, mapData map[string]interface{}, ttl time.Duration) error {
	return ErrorDisabled
}

// Remove 删除存储中键及其对应的值。 md5:95ea150955b88994
// ff:
// s:
// ctx:
// sessionId:
// key:
func (s *StorageBase) Remove(ctx context.Context, sessionId string, key string) error {
	return ErrorDisabled
}

// RemoveAll 从存储中删除会话。 md5:488d9f9ca747e8e4
// ff:
// s:
// ctx:
// sessionId:
func (s *StorageBase) RemoveAll(ctx context.Context, sessionId string) error {
	return ErrorDisabled
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
func (s *StorageBase) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (*gmap.StrAnyMap, error) {
	return nil, ErrorDisabled
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
func (s *StorageBase) SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error {
	return ErrorDisabled
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
func (s *StorageBase) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error {
	return ErrorDisabled
}
