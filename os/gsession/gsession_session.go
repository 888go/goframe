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
	gvar "github.com/888go/goframe/container/gvar"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
)

// Session 结构体，用于存储单个会话数据，它与单个请求绑定。Session 结构体是与用户交互的接口，但 Storage 是底层适配器设计的接口，用于实现特定功能。
// md5:1d1b86dcb53a276e
type Session struct {
	id      string          // 会话ID。如果指定了自定义的id，则用于获取会话。 md5:7fec8def5a7d635f
	ctx     context.Context // 当前会话的上下文。请注意，会话与上下文一起存在。 md5:c61c2b0f3688fce4
	data    *gmap.StrAnyMap // 当前会话数据，从存储中获取。 md5:00861bf25945ffcd
	dirty   bool            // 用于标记会话已修改。 md5:6d1cc43a943b4c84
	start   bool            // 用于标记会话已开始。 md5:213288ad3376abfc
	manager *Manager        // 父级会话管理器。 md5:6b920db9951d3f89

	// idFunc是一个用于创建自定义会话ID的回调函数。当会话开始时，如果会话ID为空，就会调用这个函数。
	// md5:83b853b0118605bd
	idFunc func(ttl time.Duration) (id string)
}

// init函数进行懒惰初始化session，如果指定了session ID，则获取session，否则创建一个新的空session。
// md5:60a85a2954d0a427
func (s *Session) init() error {
	if s.start {
		return nil
	}
	var err error
	// Session retrieving.
	if s.id != "" {
				// 从存储中检索存储的会话数据。 md5:e6c3f2bdc143f93c
		if s.manager.storage != nil {
			s.data, err = s.manager.storage.GetSession(s.ctx, s.id, s.manager.GetTTL())
			if err != nil && err != ErrorDisabled {
				intlog.Errorf(s.ctx, `session restoring failed for id "%s": %+v`, s.id, err)
				return err
			}
		}
	}
	// Session id creation.
	if s.id == "" {
		if s.idFunc != nil {
						// 使用自定义的会话ID创建函数。 md5:4c1e5f997d31f1b3
			s.id = s.idFunc(s.manager.ttl)
		} else {
						// 使用存储的默认会话ID创建函数。 md5:11db3a1576d0231f
			s.id, err = s.manager.storage.New(s.ctx, s.manager.ttl)
			if err != nil && err != ErrorDisabled {
				intlog.Errorf(s.ctx, "create session id failed: %+v", err)
				return err
			}
			// 如果会话存储不实现ID生成功能，那么它将使用默认的会话ID创建函数。
			// md5:4f2bd0ddc795fde4
			if s.id == "" {
				s.id = NewSessionId()
			}
		}
	}
	if s.data == nil {
		s.data = gmap.X创建StrAny(true)
	}
	s.start = true
	return nil
}

// Close 方法关闭当前会话并在会话管理器中更新其TTL（生存时间）。
// 如果此会话已被修改（脏会话），它还会将该会话导出到存储中。
//
// 注意：此功能必须在每次会话请求完成后调用。
// md5:f68a83f493f4727a
func (s *Session) Close() error {
	if s.manager.storage == nil {
		return nil
	}
	if s.start && s.id != "" {
		size := s.data.X取数量()
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

// X设置值 将键值对设置到这个会话中。 md5:09e1539c4a50fcfd
func (s *Session) X设置值(key string, value interface{}) (err error) {
	if err = s.init(); err != nil {
		return err
	}
	if err = s.manager.storage.X设置值(s.ctx, s.id, key, value, s.manager.ttl); err != nil {
		if err == ErrorDisabled {
			s.data.X设置值(key, value)
		} else {
			return err
		}
	}
	s.dirty = true
	return nil
}

// SetMap 批量使用映射设置会话。 md5:f55c78b98e85ba61
func (s *Session) SetMap(data map[string]interface{}) (err error) {
	if err = s.init(); err != nil {
		return err
	}
	if err = s.manager.storage.SetMap(s.ctx, s.id, data, s.manager.ttl); err != nil {
		if err == ErrorDisabled {
			s.data.X设置值Map(data)
		} else {
			return err
		}
	}
	s.dirty = true
	return nil
}

// Remove 从本次会话中移除指定的键及其对应的值。 md5:3dc440da200c0834
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
				s.data.X删除(key)
			} else {
				return err
			}
		}
	}
	s.dirty = true
	return nil
}

// RemoveAll 从该会话中删除所有键值对。 md5:6ca756339a9f18b5
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
		// 从内存中移除数据。 md5:47322b1cdcaf7596
	if s.data != nil {
		s.data.X清空()
	}
	s.dirty = true
	return nil
}

// Id 返回此会话的会话标识符。
// 如果在初始化时未传递会话标识符，则创建并返回新的会话标识符。
// md5:c1a4c6b98633e656
func (s *Session) Id() (id string, err error) {
	if err = s.init(); err != nil {
		return "", err
	}
	return s.id, nil
}

// SetId 在会话开始前设置自定义会话。如果在会话已经开始后调用，将返回错误。
// md5:cf8fd98a6cd07079
func (s *Session) SetId(id string) error {
	if s.start {
		return gerror.X创建错误码(gcode.CodeInvalidOperation, "session already started")
	}
	s.id = id
	return nil
}

// SetIdFunc 在会话开始前设置自定义会话ID生成函数。
// 如果在会话已经开始后调用它，将返回错误。
// md5:07c5962c3c68bf37
func (s *Session) SetIdFunc(f func(ttl time.Duration) string) error {
	if s.start {
		return gerror.X创建错误码(gcode.CodeInvalidOperation, "session already started")
	}
	s.idFunc = f
	return nil
}

// Data 将所有数据作为映射返回。
// 请注意，为了并发安全，它内部使用了值拷贝。
// md5:a37827aba4dd5df4
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
	return s.data.X取Map(), nil
}

// Size返回会话的大小。 md5:072795e87a3938d1
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
	return s.data.X取数量(), nil
}

// Contains 检查键是否存在于会话中。 md5:7a03d1ea75cda393
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
	return !v.X是否为Nil(), nil
}

// IsDirty 检查会话中是否有数据变更。 md5:2a726ce013b067fe
func (s *Session) IsDirty() bool {
	return s.dirty
}

// Get 通过给定的键获取 session 值。
// 如果键在 session 中不存在且提供了 `def`，则返回 `def`，
// 否则返回 nil。
// md5:893a612d87b25ee2
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
		return gvar.X创建(v), nil
	}
	if v = s.data.X取值(key); v != nil {
		return gvar.X创建(v), nil
	}
	if len(def) > 0 {
		return gvar.X创建(def[0]), nil
	}
	return nil, nil
}

// MustId 行为就像Id函数一样，但如果发生任何错误，它会引发恐慌。 md5:a51e8673adaf6727
func (s *Session) MustId() string {
	id, err := s.Id()
	if err != nil {
		panic(err)
	}
	return id
}

// MustGet执行与Get相同的功能，但如果发生任何错误，它将引发恐慌。 md5:bdc72a85510733d5
func (s *Session) MustGet(key string, def ...interface{}) *gvar.Var {
	v, err := s.Get(key, def...)
	if err != nil {
		panic(err)
	}
	return v
}

// MustSet 的功能与 Set 函数相同，但如果发生任何错误，它会直接 panic。 md5:06fa308e1636bcfa
func (s *Session) MustSet(key string, value interface{}) {
	err := s.X设置值(key, value)
	if err != nil {
		panic(err)
	}
}

// MustSetMap 行为类似于函数 SetMap，但如果发生任何错误则会引发 panic。 md5:3d54948e22292bcf
func (s *Session) MustSetMap(data map[string]interface{}) {
	err := s.SetMap(data)
	if err != nil {
		panic(err)
	}
}

// MustContains执行Contains函数的功能，但如果发生任何错误，它将引发恐慌。 md5:b9f29f0374157bc5
func (s *Session) MustContains(key string) bool {
	b, err := s.Contains(key)
	if err != nil {
		panic(err)
	}
	return b
}

// MustData 执行与函数 Data 相同的操作，但如果发生任何错误，它将引发恐慌。 md5:ae01e79f6f27c9fe
func (s *Session) MustData() map[string]interface{} {
	m, err := s.Data()
	if err != nil {
		panic(err)
	}
	return m
}

// MustSize 的行为与 Size 函数相同，但如果发生任何错误，它会直接 panic。 md5:d9d8c4724cdd0db4
func (s *Session) MustSize() int {
	size, err := s.Size()
	if err != nil {
		panic(err)
	}
	return size
}

// MustRemove 行为与函数 Remove 相同，但如果发生任何错误则会引发恐慌。 md5:76bd8c9cb1e6223b
func (s *Session) MustRemove(keys ...string) {
	err := s.Remove(keys...)
	if err != nil {
		panic(err)
	}
}
