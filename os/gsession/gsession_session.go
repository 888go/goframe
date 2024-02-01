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
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	)
// Session 结构体用于存储单个会话数据，它与单个请求绑定。
// Session 结构体是与用户交互的接口，而 Storage 是为功能实现设计的底层适配器接口。
type Session struct {
	id      string          // 会话ID。如果自定义指定了id，它将获取该会话。
	ctx     context.Context // 当前会话的上下文。请注意，会话与上下文共存。
	data    *gmap.StrAnyMap // 当前会话数据，从Storage中获取。
	dirty   bool            // 用于标记会话已被修改。
	start   bool            // 用于标记会话已开始。
	manager *Manager        // 父级会话管理器。

// idFunc 是一个用于创建自定义会话 ID 的回调函数。
// 当会话开始且会话 ID 为空时，将会调用这个函数。
	idFunc func(ttl time.Duration) (id string)
}

// init 执行会话的延迟初始化。如果指定了会话ID，则从存储中获取该会话；否则，创建一个新的空会话。
func (s *Session) init() error {
	if s.start {
		return nil
	}
	var err error
	// 会话获取。
	if s.id != "" {
		// 从存储中检索已存储的会话数据。
		if s.manager.storage != nil {
			s.data, err = s.manager.storage.GetSession(s.ctx, s.id, s.manager.GetTTL())
			if err != nil && err != ErrorDisabled {
				intlog.Errorf(s.ctx, `session restoring failed for id "%s": %+v`, s.id, err)
				return err
			}
		}
	}
	// 会话ID创建
	if s.id == "" {
		if s.idFunc != nil {
			// 使用自定义的会话ID创建函数。
			s.id = s.idFunc(s.manager.ttl)
		} else {
			// 使用存储的默认会话ID创建函数。
			s.id, err = s.manager.storage.New(s.ctx, s.manager.ttl)
			if err != nil && err != ErrorDisabled {
				intlog.Errorf(s.ctx, "create session id failed: %+v", err)
				return err
			}
// 如果会话存储未实现ID生成功能，
// 则使用默认的会话ID创建函数。
			if s.id == "" {
				s.id = NewSessionId()
			}
		}
	}
	if s.data == nil {
		s.data = gmap.NewStrAnyMap(true)
	}
	s.start = true
	return nil
}

// Close 关闭当前会话并在会话管理器中更新其TTL（Time To Live，生存时间）。
// 如果此会话是脏的（即已修改），它还会将其导出到存储中。
//
// 注意：此函数必须在每次会话请求完成后调用。
func (s *Session) Close() error {
	if s.manager.storage == nil {
		return nil
	}
	if s.start && s.id != "" {
		size := s.data.Size()
		if s.dirty {
			err := s.manager.storage.SetSession(s.ctx, s.id, s.data, s.manager.ttl)
			if err != nil && err != ErrorDisabled {
				return err
			}
		} else if size > 0 {
			err := s.manager.storage.UpdateTTL(s.ctx, s.id, s.manager.ttl)
			if err != nil && err != ErrorDisabled {
				return err
			}
		}
	}
	return nil
}

// Set 将键值对设置到此会话中。
func (s *Session) Set(key string, value interface{}) (err error) {
	if err = s.init(); err != nil {
		return err
	}
	if err = s.manager.storage.Set(s.ctx, s.id, key, value, s.manager.ttl); err != nil {
		if err == ErrorDisabled {
			s.data.Set(key, value)
		} else {
			return err
		}
	}
	s.dirty = true
	return nil
}

// SetMap 批量使用 map 设置 session。
func (s *Session) SetMap(data map[string]interface{}) (err error) {
	if err = s.init(); err != nil {
		return err
	}
	if err = s.manager.storage.SetMap(s.ctx, s.id, data, s.manager.ttl); err != nil {
		if err == ErrorDisabled {
			s.data.Sets(data)
		} else {
			return err
		}
	}
	s.dirty = true
	return nil
}

// Remove 从当前会话中移除指定键及其对应的值。
func (s *Session) Remove(keys ...string) (err error) {
	if s.id == "" {
		return nil
	}
	if err = s.init(); err != nil {
		return err
	}
	for _, key := range keys {
		if err = s.manager.storage.Remove(s.ctx, s.id, key); err != nil {
			if err == ErrorDisabled {
				s.data.Remove(key)
			} else {
				return err
			}
		}
	}
	s.dirty = true
	return nil
}

// RemoveAll 从该会话中删除所有键值对。
func (s *Session) RemoveAll() (err error) {
	if s.id == "" {
		return nil
	}
	if err = s.init(); err != nil {
		return err
	}
	if err = s.manager.storage.RemoveAll(s.ctx, s.id); err != nil {
		if err != ErrorDisabled {
			return err
		}
	}
	// 从内存中移除数据。
	if s.data != nil {
		s.data.Clear()
	}
	s.dirty = true
	return nil
}

// Id 返回当前会话的 session id。
// 如果在初始化时没有传递 session id，它将创建并返回一个新的 session id。
func (s *Session) Id() (id string, err error) {
	if err = s.init(); err != nil {
		return "", err
	}
	return s.id, nil
}

// SetId 在会话开始前设置自定义会话ID。
// 如果在会话开始后调用，将返回错误。
func (s *Session) SetId(id string) error {
	if s.start {
		return gerror.NewCode(gcode.CodeInvalidOperation, "session already started")
	}
	s.id = id
	return nil
}

// SetIdFunc 在会话开始前设置自定义的 session id 生成函数。 
// 如果在会话开始后调用，将返回错误。
func (s *Session) SetIdFunc(f func(ttl time.Duration) string) error {
	if s.start {
		return gerror.NewCode(gcode.CodeInvalidOperation, "session already started")
	}
	s.idFunc = f
	return nil
}

// Data 返回所有数据作为映射。
// 注意，为了保证并发安全，内部使用了值复制。
func (s *Session) Data() (sessionData map[string]interface{}, err error) {
	if s.id == "" {
		return map[string]interface{}{}, nil
	}
	if err = s.init(); err != nil {
		return nil, err
	}
	sessionData, err = s.manager.storage.Data(s.ctx, s.id)
	if err != nil && err != ErrorDisabled {
		intlog.Errorf(s.ctx, `%+v`, err)
	}
	if sessionData != nil {
		return sessionData, nil
	}
	return s.data.Map(), nil
}

// Size 返回会话的大小。
func (s *Session) Size() (size int, err error) {
	if s.id == "" {
		return 0, nil
	}
	if err = s.init(); err != nil {
		return 0, err
	}
	size, err = s.manager.storage.GetSize(s.ctx, s.id)
	if err != nil && err != ErrorDisabled {
		intlog.Errorf(s.ctx, `%+v`, err)
	}
	if size > 0 {
		return size, nil
	}
	return s.data.Size(), nil
}

// Contains 检查键是否存在于会话中。
func (s *Session) Contains(key string) (ok bool, err error) {
	if s.id == "" {
		return false, nil
	}
	if err = s.init(); err != nil {
		return false, err
	}
	v, err := s.Get(key)
	if err != nil {
		return false, err
	}
	return !v.IsNil(), nil
}

// IsDirty 检查会话中是否存在任何数据更改。
func (s *Session) IsDirty() bool {
	return s.dirty
}

// Get 通过给定的键从会话中检索值。
// 如果提供了 `def`，当键在会话中不存在时，它将返回 `def`，
// 否则返回 nil。
func (s *Session) Get(key string, def ...interface{}) (value *gvar.Var, err error) {
	if s.id == "" {
		return nil, nil
	}
	if err = s.init(); err != nil {
		return nil, err
	}
	v, err := s.manager.storage.Get(s.ctx, s.id, key)
	if err != nil && err != ErrorDisabled {
		intlog.Errorf(s.ctx, `%+v`, err)
		return nil, err
	}
	if v != nil {
		return gvar.New(v), nil
	}
	if v = s.data.Get(key); v != nil {
		return gvar.New(v), nil
	}
	if len(def) > 0 {
		return gvar.New(def[0]), nil
	}
	return nil, nil
}

// MustId 函数表现如同 Id 函数，但是当发生任何错误时，它会触发panic。
func (s *Session) MustId() string {
	id, err := s.Id()
	if err != nil {
		panic(err)
	}
	return id
}

// MustGet 函数表现如同 Get 函数，但当发生任何错误时，它会触发panic。
func (s *Session) MustGet(key string, def ...interface{}) *gvar.Var {
	v, err := s.Get(key, def...)
	if err != nil {
		panic(err)
	}
	return v
}

// MustSet 函数表现如同 Set 函数，但是当发生任何错误时它会引发panic。
func (s *Session) MustSet(key string, value interface{}) {
	err := s.Set(key, value)
	if err != nil {
		panic(err)
	}
}

// MustSetMap表现如同函数 SetMap，但是当发生任何错误时它会触发panic。
func (s *Session) MustSetMap(data map[string]interface{}) {
	err := s.SetMap(data)
	if err != nil {
		panic(err)
	}
}

// MustContains 执行与 Contains 函数相同的功能，但是当发生任何错误时，它会触发panic。
func (s *Session) MustContains(key string) bool {
	b, err := s.Contains(key)
	if err != nil {
		panic(err)
	}
	return b
}

// MustData 的行为类似于函数 Data，但是当发生任何错误时，它会触发 panic。
func (s *Session) MustData() map[string]interface{} {
	m, err := s.Data()
	if err != nil {
		panic(err)
	}
	return m
}

// MustSize 执行与 Size 函数相同的功能，但是当发生任何错误时，它会触发 panic。
func (s *Session) MustSize() int {
	size, err := s.Size()
	if err != nil {
		panic(err)
	}
	return size
}

// MustRemove 的行为与 Remove 函数相同，但是当发生任何错误时，它会触发 panic。
func (s *Session) MustRemove(keys ...string) {
	err := s.Remove(keys...)
	if err != nil {
		panic(err)
	}
}
