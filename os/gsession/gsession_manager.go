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
)

// Manager for sessions.
type Manager struct {
	ttl     time.Duration // TTL for sessions.
	storage Storage       // 会话存储的接口。 md5:75545bf4dbeae018
}

// New 创建并返回一个新的会话管理器。 md5:f41a96ed9e4273e4
func New(ttl time.Duration, storage ...Storage) *Manager {
	m := &Manager{
		ttl: ttl,
	}
	if len(storage) > 0 && storage[0] != nil {
		m.storage = storage[0]
	} else {
				// 默认使用 StorageFile。 md5:a8eedc355767eccd
		m.storage = NewStorageFile(DefaultStorageFilePath, ttl)
	}
	return m
}

// New 为给定的 session ID 创建或获取会话。
// 参数 `sessionId` 是可选的，如果未提供，则根据 Storage.New 的行为创建新的会话。
// md5:4d84930c3cbf9027
func (m *Manager) New(ctx context.Context, sessionId ...string) *Session {
	var id string
	if len(sessionId) > 0 && sessionId[0] != "" {
		id = sessionId[0]
	}
	return &Session{
		id:      id,
		ctx:     ctx,
		manager: m,
	}
}

// SetStorage 设置管理器的会话存储。 md5:9fe6b2a9a6ae9990
func (m *Manager) SetStorage(storage Storage) {
	m.storage = storage
}

// GetStorage 返回当前会话管理器的存储对象。 md5:43cdd2b5155f8389
func (m *Manager) GetStorage() Storage {
	return m.storage
}

// SetTTL 设置会话管理器的生存时间（TTL）。 md5:bba913d23693cf2a
func (m *Manager) SetTTL(ttl time.Duration) {
	m.ttl = ttl
}

// GetTTL 返回会话管理器的TTL（时间到 live，生存时间）。 md5:d0733ac8b424fbe1
func (m *Manager) GetTTL() time.Duration {
	return m.ttl
}
