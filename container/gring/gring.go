// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gring提供了一个并发安全/不安全的环形列表（圆形队列）。 md5:bc78eee87d7b5c4b
package gring

import (
	"container/ring"

	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/internal/rwmutex"
)

// Ring是一个环形结构的结构体。 md5:f371ac74ef187b03
type Ring struct {
	mu    *rwmutex.RWMutex
	ring  *ring.Ring  // Underlying ring.
	len   *gtype.Int  // 已使用长度。 md5:0093c138cefea3f0
	cap   *gtype.Int  // Capability(>=len).
	dirty *gtype.Bool // Dirty, which means the len and cap should be recalculated. It's marked dirty when the size of ring changes.
}

// internalRingItem 用于存储环形元素的值。 md5:d1394bf8a3b910df
type internalRingItem struct {
	Value interface{}
}

// New 创建并返回一个具有`cap`个元素的Ring结构。
// 可选参数`safe`指定是否在并发安全环境下使用该结构，默认为false。
// md5:70892e7ec9ed75d6
func New(cap int, safe ...bool) *Ring {
	return &Ring{
		mu:    rwmutex.New(safe...),
		ring:  ring.New(cap),
		len:   gtype.NewInt(),
		cap:   gtype.NewInt(cap),
		dirty: gtype.NewBool(),
	}
}

// Val 返回当前位置的项的值。 md5:b1027c8df14f08d2
func (r *Ring) Val() interface{} {
	var value interface{}
	r.mu.RLock()
	if r.ring.Value != nil {
		value = r.ring.Value.(internalRingItem).Value
	}
	r.mu.RUnlock()
	return value
}

// Len 返回环的大小。 md5:c4ff976cf0b72c58
func (r *Ring) Len() int {
	r.checkAndUpdateLenAndCap()
	return r.len.Val()
}

// Cap 返回环形缓冲区的容量。 md5:2ac015d8e20dce37
func (r *Ring) Cap() int {
	r.checkAndUpdateLenAndCap()
	return r.cap.Val()
}

// 检查并更新ring的长度和容量，当ring被修改时。 md5:264d6fdc8ef33d31
func (r *Ring) checkAndUpdateLenAndCap() {
	if !r.dirty.Val() {
		return
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	totalLen := 0
	emptyLen := 0
	if r.ring != nil {
		if r.ring.Value == nil {
			emptyLen++
		}
		totalLen++
		for p := r.ring.Next(); p != r.ring; p = p.Next() {
			if p.Value == nil {
				emptyLen++
			}
			totalLen++
		}
	}
	r.cap.Set(totalLen)
	r.len.Set(totalLen - emptyLen)
	r.dirty.Set(false)
}

// Set 将值设置为当前位置的项目。 md5:7140c77dfa3aa5dc
func (r *Ring) Set(value interface{}) *Ring {
	r.mu.Lock()
	if r.ring.Value == nil {
		r.len.Add(1)
	}
	r.ring.Value = internalRingItem{Value: value}
	r.mu.Unlock()
	return r
}

// Put 将`value`设置为环形列表的当前项，并将位置移动到下一项。 md5:737e9a607801eee9
func (r *Ring) Put(value interface{}) *Ring {
	r.mu.Lock()
	if r.ring.Value == nil {
		r.len.Add(1)
	}
	r.ring.Value = internalRingItem{Value: value}
	r.ring = r.ring.Next()
	r.mu.Unlock()
	return r
}

// Move 函数根据给定的 n 值，将环(ring)中的元素向后（n < 0）或向前（n >= 0）移动 n % r.Len() 个位置，并返回移动后所在位置的元素。环(r)不能为空。
// md5:92f786c9a5a8b8cd
func (r *Ring) Move(n int) *Ring {
	r.mu.Lock()
	r.ring = r.ring.Move(n)
	r.mu.Unlock()
	return r
}

// Prev 返回前一个环形元素。r 不能为空。 md5:574e755d8883dc1f
func (r *Ring) Prev() *Ring {
	r.mu.Lock()
	r.ring = r.ring.Prev()
	r.mu.Unlock()
	return r
}

// Next 返回下一个环元素。r必须不为空。 md5:f16b811ce23ee06b
func (r *Ring) Next() *Ring {
	r.mu.Lock()
	r.ring = r.ring.Next()
	r.mu.Unlock()
	return r
}

// Link 将环 r 与环 s 连接起来，使得 r.Next() 变为 s，并返回 r.Next() 的原始值。
// r 不应为空。
// 
// 如果 r 和 s 指向相同的环，链接它们会移除 r 和 s 之间的元素。移除的元素形成一个子环，结果是对这个子环的引用（如果没有移除任何元素，结果仍然是 r.Next() 的原始值，而不是 nil）。
// 
// 如果 r 和 s 指向不同的环，链接它们会在 r 后插入 s 的元素，创建一个单一的环。结果指向插入后 s 的最后一个元素之后的元素。
// md5:faa73e3f5f43468a
func (r *Ring) Link(s *Ring) *Ring {
	r.mu.Lock()
	s.mu.Lock()
	r.ring.Link(s.ring)
	s.mu.Unlock()
	r.mu.Unlock()
	r.dirty.Set(true)
	s.dirty.Set(true)
	return r
}

// Unlink 从环 r 中移除 n % r.Len() 个元素，开始于 r.Next()。
// 如果 n % r.Len() == 0，那么 r 保持不变。
// 结果是被移除的子环。r 必须非空。
// md5:00909914d8e87d32
func (r *Ring) Unlink(n int) *Ring {
	r.mu.Lock()
	resultRing := r.ring.Unlink(n)
	r.dirty.Set(true)
	r.mu.Unlock()
	resultGRing := New(resultRing.Len())
	resultGRing.ring = resultRing
	resultGRing.dirty.Set(true)
	return resultGRing
}

// RLockIteratorNext 在RWMutex的RLock范围内向前迭代并加锁读取。
// 如果提供的回调函数`f`返回true，那么将继续迭代；如果返回false，则停止迭代。
// md5:8cb144956023168f
func (r *Ring) RLockIteratorNext(f func(value interface{}) bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if r.ring.Value != nil && !f(r.ring.Value.(internalRingItem).Value) {
		return
	}
	for p := r.ring.Next(); p != r.ring; p = p.Next() {
		if p.Value == nil || !f(p.Value.(internalRingItem).Value) {
			break
		}
	}
}

// RLockIteratorPrev 逆序迭代并锁定（写入）RWMutex.RLock。
// 使用给定的回调函数 `f`。如果 `f` 返回 true，那么继续迭代；否则停止。
// md5:31f0f0af041e234a
func (r *Ring) RLockIteratorPrev(f func(value interface{}) bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if r.ring.Value != nil && !f(r.ring.Value.(internalRingItem).Value) {
		return
	}
	for p := r.ring.Prev(); p != r.ring; p = p.Prev() {
		if p.Value == nil || !f(p.Value.(internalRingItem).Value) {
			break
		}
	}
}

// SliceNext 返回一个从当前位置开始向前的所有项值的切片副本。 md5:54ba7b6ac01a38f8
func (r *Ring) SliceNext() []interface{} {
	s := make([]interface{}, 0)
	r.mu.RLock()
	if r.ring.Value != nil {
		s = append(s, r.ring.Value.(internalRingItem).Value)
	}
	for p := r.ring.Next(); p != r.ring; p = p.Next() {
		if p.Value == nil {
			break
		}
		s = append(s, p.Value.(internalRingItem).Value)
	}
	r.mu.RUnlock()
	return s
}

// SlicePrev 从当前位置向前返回所有项目值的副本作为切片。 md5:632f85c2939f2e91
func (r *Ring) SlicePrev() []interface{} {
	s := make([]interface{}, 0)
	r.mu.RLock()
	if r.ring.Value != nil {
		s = append(s, r.ring.Value.(internalRingItem).Value)
	}
	for p := r.ring.Prev(); p != r.ring; p = p.Prev() {
		if p.Value == nil {
			break
		}
		s = append(s, p.Value.(internalRingItem).Value)
	}
	r.mu.RUnlock()
	return s
}
