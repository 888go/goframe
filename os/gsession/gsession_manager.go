// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package session类

import (
	"context"
	"time"
)

// 会话管理器。
type Manager struct {
	ttl     time.Duration // TTL for sessions.
	storage Storage       // Storage interface 用于会话存储。
}

// New 创建并返回一个新的会话管理器。
func New(ttl time.Duration, storage ...Storage) *Manager {
	m := &Manager{
		ttl: ttl,
	}
	if len(storage) > 0 && storage[0] != nil {
		m.storage = storage[0]
	} else {
		// 默认情况下，它使用 StorageFile。
		m.storage = NewStorageFile(DefaultStorageFilePath, ttl)
	}
	return m
}

// New 根据给定的会话ID创建或获取会话。
// 参数 `sessionId` 是可选的，如果不传递，则根据 Storage.New 创建一个新的会话。
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

// SetStorage 为 manager 设置 session 存储。
func (m *Manager) SetStorage(storage Storage) {
	m.storage = storage
}

// GetStorage 返回当前管理器的会话存储。
func (m *Manager) GetStorage() Storage {
	return m.storage
}

// 设置会话管理器的TTL（生存时间）。
func (m *Manager) SetTTL(ttl time.Duration) {
	m.ttl = ttl
}

// GetTTL 返回会话管理器的TTL（生存时间）
func (m *Manager) GetTTL() time.Duration {
	return m.ttl
}
