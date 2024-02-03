// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// gring包提供了一种并发安全/不安全的循环链表（环形列表）。
package gring

import (
	"container/ring"
	
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/internal/rwmutex"
)

// Ring 是一个环形结构的结构体。
type Ring struct {
	mu    *rwmutex.RWMutex
	ring  *ring.Ring  // Underlying ring.
	len   *gtype.Int  // 长度（已使用大小）
	cap   *gtype.Int  // Capability(>=len) 表示能力（或功能）至少为 len。
	dirty *gtype.Bool // Dirty, which means the len and cap should be recalculated. It's marked dirty when the size of ring changes.
}

// internalRingItem 存储环形元素的值。
type internalRingItem struct {
	Value interface{}
}

// New 创建并返回一个容量为`cap`的Ring结构体。
// 可选参数`sage`用于指定该结构体是否在并发场景下安全使用，默认为false（不安全）。
func New(cap int, safe ...bool) *Ring {
	return &Ring{
		mu:    rwmutex.New(safe...),
		ring:  ring.New(cap),
		len:   gtype.NewInt(),
		cap:   gtype.NewInt(cap),
		dirty: gtype.NewBool(),
	}
}

// Val 返回当前位置项的值。
func (r *Ring) Val() interface{} {
	var value interface{}
	r.mu.RLock()
	if r.ring.Value != nil {
		value = r.ring.Value.(internalRingItem).Value
	}
	r.mu.RUnlock()
	return value
}

// Len 返回环形结构的大小。
func (r *Ring) Len() int {
	r.checkAndUpdateLenAndCap()
	return r.len.Val()
}

// Cap 返回环形缓冲区的容量。
func (r *Ring) Cap() int {
	r.checkAndUpdateLenAndCap()
	return r.cap.Val()
}

// 当ring脏时，检查并更新ring的长度(len)和容量(cap)。
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

// Set 将值设置为当前位置的项。
func (r *Ring) Set(value interface{}) *Ring {
	r.mu.Lock()
	if r.ring.Value == nil {
		r.len.Add(1)
	}
	r.ring.Value = internalRingItem{Value: value}
	r.mu.Unlock()
	return r
}

// Put 将 `value` 设置为环形结构当前项的值，并将位置移动到下一个项。
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

// Move 函数将循环队列中的元素向后（n < 0）或向前（n >= 0）移动 n % r.Len() 个位置，
// 并返回该移动后的位置上的元素。注意，r 不得为空。
func (r *Ring) Move(n int) *Ring {
	r.mu.Lock()
	r.ring = r.ring.Move(n)
	r.mu.Unlock()
	return r
}

// Prev 返回上一个环形元素。r不能为空。
func (r *Ring) Prev() *Ring {
	r.mu.Lock()
	r.ring = r.ring.Prev()
	r.mu.Unlock()
	return r
}

// Next 返回下一个环形元素。r 必须非空。
func (r *Ring) Next() *Ring {
	r.mu.Lock()
	r.ring = r.ring.Next()
	r.mu.Unlock()
	return r
}

// Link 将环 r 与环 s 连接，使得 r.Next() 指向 s，并返回连接前 r.Next() 的原始值。
// r 必须非空。
//
// 如果 r 和 s 指向同一个环，将它们连接会从环中移除 r 和 s 之间的元素。被移除的元素形成一个子环，结果是对该子环的一个引用（如果未移除任何元素，则结果仍然是原始的 r.Next() 值，而非 nil）。
//
// 如果 r 和 s 指向不同的环，将它们连接会创建一个新的单个环，在 r 后面插入 s 中的所有元素。结果指向在插入后 s 的最后一个元素之后的那个元素。
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

// Unlink 从环形链表 r 中移除 n % r.Len() 个元素，从 r.Next() 开始移除。
// 若 n % r.Len() 等于 0，则 r 保持不变。
// 返回值为被移除的子环。r 必须非空。
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

// RLockIteratorNext 在 RWMutex.RLock 的保护下，以给定的回调函数 `f` 进行正向迭代并进行读取锁定。
// 若 `f` 返回 true，则继续进行迭代；若返回 false，则停止迭代。
// 这段代码的中文注释如下：
// ```go
// RLockIteratorNext 函数在读写互斥锁（RWMutex）的读锁状态下，通过给定的回调函数 `f` 实现向前遍历并加读锁。
// 当 `f` 返回值为 true 时，将继续进行遍历操作；若返回 false，则停止遍历。
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

// RLockIteratorPrev 在RWMutex.RLock的保护下，以给定回调函数`f`向后遍历并加写锁。
// 如果`f`返回true，则继续迭代；若返回false，则停止遍历。
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

// SliceNext 从当前位置开始向前复制所有项目值，并以切片形式返回。
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

// SlicePrev 从当前位置开始向后返回所有项值的切片副本。
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
