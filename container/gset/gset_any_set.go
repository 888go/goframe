// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包gset提供了多种并发安全/不安全的集合（sets）。
// 集合，即不可重复的一组元素，元素项可以为任意类型。
// 同时，gset支持可选的并发安全参数选项，支持并发安全的场景。 
package gset

import (
	"bytes"
	
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/rwmutex"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
)

type Set struct {
	mu   rwmutex.RWMutex
	data map[interface{}]struct{}
}

// New 创建并返回一个新的集合，其中包含无重复项。
// 参数`safe`用于指定是否在并发安全的情况下使用集合，默认为false。
func New(safe ...bool) *Set {
	return NewSet(safe...)
}

// NewSet 为New别名, 创建并返回一个新的集合，其中包含不重复的项目。
// 也可以参考 New。
func NewSet(safe ...bool) *Set {
	return &Set{
		data: make(map[interface{}]struct{}),
		mu:   rwmutex.Create(safe...),
	}
}

// NewFrom 函数根据 `items` 返回一个新的集合。
// 参数 `items` 可以是任何类型的变量，也可以是一个切片。
func NewFrom(items interface{}, safe ...bool) *Set {
	m := make(map[interface{}]struct{})
	for _, v := range gconv.Interfaces(items) {
		m[v] = struct{}{}
	}
	return &Set{
		data: m,
		mu:   rwmutex.Create(safe...),
	}
}

// Iterator 使用给定的回调函数`f`对集合进行只读遍历，
// 如果`f`返回true，则继续遍历；若返回false，则停止遍历。
func (set *Set) Iterator(f func(v interface{}) bool) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k := range set.data {
		if !f(k) {
			break
		}
	}
}

// Add 向集合中添加一个或多个项目。
func (set *Set) Add(items ...interface{}) {
	set.mu.Lock()
	if set.data == nil {
		set.data = make(map[interface{}]struct{})
	}
	for _, v := range items {
		set.data[v] = struct{}{}
	}
	set.mu.Unlock()
}

// AddIfNotExist 检查项是否已存在于集合中，
// 如果项不存在于集合中，则将其添加到集合并返回 true；
// 否则，不执行任何操作并返回 false。
//
// 注意：如果 `item` 为 nil，则不执行任何操作并返回 false。
func (set *Set) AddIfNotExist(item interface{}) bool {
	if item == nil {
		return false
	}
	if !set.Contains(item) {
		set.mu.Lock()
		defer set.mu.Unlock()
		if set.data == nil {
			set.data = make(map[interface{}]struct{})
		}
		if _, ok := set.data[item]; !ok {
			set.data[item] = struct{}{}
			return true
		}
	}
	return false
}

// AddIfNotExistFunc 检查项是否存在集合中，
// 如果项不存在于集合中，并且函数 `f` 返回 true，则将该项添加到集合并返回 true；
// 否则，不做任何操作并返回 false。
//
// 注意，如果 `item` 为 nil，将不做任何操作并返回 false。函数 `f`
// 在无写入锁的情况下执行。
func (set *Set) AddIfNotExistFunc(item interface{}, f func() bool) bool {
	if item == nil {
		return false
	}
	if !set.Contains(item) {
		if f() {
			set.mu.Lock()
			defer set.mu.Unlock()
			if set.data == nil {
				set.data = make(map[interface{}]struct{})
			}
			if _, ok := set.data[item]; !ok {
				set.data[item] = struct{}{}
				return true
			}
		}
	}
	return false
}

// AddIfNotExistFuncLock 检查项目是否已存在于集合中，
// 如果该项目不存在于集合中且函数 `f` 返回 true，则将项目添加到集合并返回 true。
// 否则，不执行任何操作并返回 false。
//
// 注意，如果 `item` 为 nil，则不执行任何操作并返回 false。函数 `f`
// 在写入锁的保护下执行。
func (set *Set) AddIfNotExistFuncLock(item interface{}, f func() bool) bool {
	if item == nil {
		return false
	}
	if !set.Contains(item) {
		set.mu.Lock()
		defer set.mu.Unlock()
		if set.data == nil {
			set.data = make(map[interface{}]struct{})
		}
		if f() {
			if _, ok := set.data[item]; !ok {
				set.data[item] = struct{}{}
				return true
			}
		}
	}
	return false
}

// Contains 检查集合中是否包含 `item`。
func (set *Set) Contains(item interface{}) bool {
	var ok bool
	set.mu.RLock()
	if set.data != nil {
		_, ok = set.data[item]
	}
	set.mu.RUnlock()
	return ok
}

// Remove 从集合中删除`item`。
func (set *Set) Remove(item interface{}) {
	set.mu.Lock()
	if set.data != nil {
		delete(set.data, item)
	}
	set.mu.Unlock()
}

// Size 返回集合的大小。
func (set *Set) Size() int {
	set.mu.RLock()
	l := len(set.data)
	set.mu.RUnlock()
	return l
}

// 清除删除集合中的所有项。
func (set *Set) Clear() {
	set.mu.Lock()
	set.data = make(map[interface{}]struct{})
	set.mu.Unlock()
}

// Slice 返回集合中所有项作为一个切片。
func (set *Set) Slice() []interface{} {
	set.mu.RLock()
	var (
		i   = 0
		ret = make([]interface{}, len(set.data))
	)
	for item := range set.data {
		ret[i] = item
		i++
	}
	set.mu.RUnlock()
	return ret
}

// Join通过字符串`glue`连接items。
func (set *Set) Join(glue string) string {
	set.mu.RLock()
	defer set.mu.RUnlock()
	if len(set.data) == 0 {
		return ""
	}
	var (
		l      = len(set.data)
		i      = 0
		buffer = bytes.NewBuffer(nil)
	)
	for k := range set.data {
		buffer.WriteString(gconv.String(k))
		if i != l-1 {
			buffer.WriteString(glue)
		}
		i++
	}
	return buffer.String()
}

// String 返回 items 作为字符串，其实现方式类似于 json.Marshal。
func (set *Set) String() string {
	if set == nil {
		return ""
	}
	set.mu.RLock()
	defer set.mu.RUnlock()
	var (
		s      string
		l      = len(set.data)
		i      = 0
		buffer = bytes.NewBuffer(nil)
	)
	buffer.WriteByte('[')
	for k := range set.data {
		s = gconv.String(k)
		if gstr.IsNumeric(s) {
			buffer.WriteString(s)
		} else {
			buffer.WriteString(`"` + gstr.QuoteMeta(s, `"\`) + `"`)
		}
		if i != l-1 {
			buffer.WriteByte(',')
		}
		i++
	}
	buffer.WriteByte(']')
	return buffer.String()
}

// LockFunc 使用回调函数`f`进行写入锁定。
func (set *Set) LockFunc(f func(m map[interface{}]struct{})) {
	set.mu.Lock()
	defer set.mu.Unlock()
	f(set.data)
}

// RLockFunc 通过回调函数 `f` 对读取进行加锁。
func (set *Set) RLockFunc(f func(m map[interface{}]struct{})) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	f(set.data)
}

// Equal 检查两个集合是否相等。
func (set *Set) Equal(other *Set) bool {
	if set == other {
		return true
	}
	set.mu.RLock()
	defer set.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()
	if len(set.data) != len(other.data) {
		return false
	}
	for key := range set.data {
		if _, ok := other.data[key]; !ok {
			return false
		}
	}
	return true
}

// IsSubsetOf 检查当前集合是否为 `other` 的子集。
func (set *Set) IsSubsetOf(other *Set) bool {
	if set == other {
		return true
	}
	set.mu.RLock()
	defer set.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()
	for key := range set.data {
		if _, ok := other.data[key]; !ok {
			return false
		}
	}
	return true
}

// Union 返回一个新的集合，该集合是 `set` 和 `others` 的并集。
// 这意味着，`newSet` 中的所有元素要么在 `set` 中，要么在 `others` 中。
func (set *Set) Union(others ...*Set) (newSet *Set) {
	newSet = NewSet()
	set.mu.RLock()
	defer set.mu.RUnlock()
	for _, other := range others {
		if set != other {
			other.mu.RLock()
		}
		for k, v := range set.data {
			newSet.data[k] = v
		}
		if set != other {
			for k, v := range other.data {
				newSet.data[k] = v
			}
		}
		if set != other {
			other.mu.RUnlock()
		}
	}

	return
}

// Diff 返回一个新的集合，它是从 `set` 到 `others` 的差集。
// 这意味着，新集合 `newSet` 中的所有元素都在 `set` 中，但不在 `others` 中。
func (set *Set) Diff(others ...*Set) (newSet *Set) {
	newSet = NewSet()
	set.mu.RLock()
	defer set.mu.RUnlock()
	for _, other := range others {
		if set == other {
			continue
		}
		other.mu.RLock()
		for k, v := range set.data {
			if _, ok := other.data[k]; !ok {
				newSet.data[k] = v
			}
		}
		other.mu.RUnlock()
	}
	return
}

// Intersect 返回一个新的集合，它是从 `set` 到 `others` 的交集。
// 这意味着，新集合 `newSet` 中的所有元素都同时存在于 `set` 和 `others` 中。
func (set *Set) Intersect(others ...*Set) (newSet *Set) {
	newSet = NewSet()
	set.mu.RLock()
	defer set.mu.RUnlock()
	for _, other := range others {
		if set != other {
			other.mu.RLock()
		}
		for k, v := range set.data {
			if _, ok := other.data[k]; ok {
				newSet.data[k] = v
			}
		}
		if set != other {
			other.mu.RUnlock()
		}
	}
	return
}

// Complement 返回一个新的集合，该集合为 `set` 在 `full` 中的补集。
// 这意味着，新集合 `newSet` 中的所有元素都在 `full` 中但不在 `set` 中。
//
// 如果给定的集合 `full` 不是 `set` 的全集，则返回 `full` 与 `set` 之间的差集。
func (set *Set) Complement(full *Set) (newSet *Set) {
	newSet = NewSet()
	set.mu.RLock()
	defer set.mu.RUnlock()
	if set != full {
		full.mu.RLock()
		defer full.mu.RUnlock()
	}
	for k, v := range full.data {
		if _, ok := set.data[k]; !ok {
			newSet.data[k] = v
		}
	}
	return
}

// Merge 将 `others` 中的元素合并到 `set` 中。
func (set *Set) Merge(others ...*Set) *Set {
	set.mu.Lock()
	defer set.mu.Unlock()
	for _, other := range others {
		if set != other {
			other.mu.RLock()
		}
		for k, v := range other.data {
			set.data[k] = v
		}
		if set != other {
			other.mu.RUnlock()
		}
	}
	return set
}

// Sum 计算项的总和。
// 注意：项应转换为 int 类型，
// 否则你将得到一个意想不到的结果。
func (set *Set) Sum() (sum int) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k := range set.data {
		sum += gconv.Int(k)
	}
	return
}

// Pop 随机地从集合中弹出一个元素。
func (set *Set) Pop() interface{} {
	set.mu.Lock()
	defer set.mu.Unlock()
	for k := range set.data {
		delete(set.data, k)
		return k
	}
	return nil
}

// Pops 随机地从集合中弹出 `size` 个元素。
// 如果 size == -1，则返回所有元素。
func (set *Set) Pops(size int) []interface{} {
	set.mu.Lock()
	defer set.mu.Unlock()
	if size > len(set.data) || size == -1 {
		size = len(set.data)
	}
	if size <= 0 {
		return nil
	}
	index := 0
	array := make([]interface{}, size)
	for k := range set.data {
		delete(set.data, k)
		array[index] = k
		index++
		if index == size {
			break
		}
	}
	return array
}

// Walk 对集合中的每一个元素应用用户提供的函数 `f`。
func (set *Set) Walk(f func(item interface{}) interface{}) *Set {
	set.mu.Lock()
	defer set.mu.Unlock()
	m := make(map[interface{}]struct{}, len(set.data))
	for k, v := range set.data {
		m[f(k)] = v
	}
	set.data = m
	return set
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (set Set) MarshalJSON() ([]byte, error) {
	return json.Marshal(set.Slice())
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func (set *Set) UnmarshalJSON(b []byte) error {
	set.mu.Lock()
	defer set.mu.Unlock()
	if set.data == nil {
		set.data = make(map[interface{}]struct{})
	}
	var array []interface{}
	if err := json.UnmarshalUseNumber(b, &array); err != nil {
		return err
	}
	for _, v := range array {
		set.data[v] = struct{}{}
	}
	return nil
}

// UnmarshalValue 是一个接口实现，用于为 set 设置任意类型的值。
func (set *Set) UnmarshalValue(value interface{}) (err error) {
	set.mu.Lock()
	defer set.mu.Unlock()
	if set.data == nil {
		set.data = make(map[interface{}]struct{})
	}
	var array []interface{}
	switch value.(type) {
	case string, []byte:
		err = json.UnmarshalUseNumber(gconv.Bytes(value), &array)
	default:
		array = gconv.SliceAny(value)
	}
	for _, v := range array {
		set.data[v] = struct{}{}
	}
	return
}

// DeepCopy 实现接口，用于当前类型的深度复制。
func (set *Set) DeepCopy() interface{} {
	if set == nil {
		return nil
	}
	set.mu.RLock()
	defer set.mu.RUnlock()
	data := make([]interface{}, 0)
	for k := range set.data {
		data = append(data, k)
	}
	return NewFrom(data, set.mu.IsSafe())
}
