// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

package gset

import (
	"bytes"

	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/internal/rwmutex"
	"github.com/gogf/gf/v2/util/gconv"
)

type IntSet struct {
	mu   rwmutex.RWMutex
	data map[int]struct{}
}

// NewIntSet 创建并返回一个新集合，其中包含不重复的元素。
// 参数 `safe` 用于指定是否在并发安全环境下使用集合，默认为 false。
// md5:5ede16db776ad391
func NewIntSet(safe ...bool) *IntSet {
	return &IntSet{
		mu:   rwmutex.Create(safe...),
		data: make(map[int]struct{}),
	}
}

// NewIntSetFrom 从`items`创建一个新的整数集合。. md5:473f94b321141021
func NewIntSetFrom(items []int, safe ...bool) *IntSet {
	m := make(map[int]struct{})
	for _, v := range items {
		m[v] = struct{}{}
	}
	return &IntSet{
		mu:   rwmutex.Create(safe...),
		data: m,
	}
}

// Iterator 使用给定的回调函数 `f` 遍历只读集合，如果 `f` 返回 true，则继续遍历；否则停止。
// md5:b896360b1cf6fc88
func (set *IntSet) Iterator(f func(v int) bool) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k := range set.data {
		if !f(k) {
			break
		}
	}
}

// Add 将一个或多个项目添加到集合中。. md5:316141ff7d4b8e45
func (set *IntSet) Add(item ...int) {
	set.mu.Lock()
	if set.data == nil {
		set.data = make(map[int]struct{})
	}
	for _, v := range item {
		set.data[v] = struct{}{}
	}
	set.mu.Unlock()
}

// AddIfNotExist 检查项是否存在于集合中，
// 如果项不存在于集合中，则将其添加到集合并返回true，
// 否则不做任何操作并返回false。
//
// 注意，如果 `item` 为 nil，它将不做任何操作并返回false。
// md5:3d920a290d301fb9
func (set *IntSet) AddIfNotExist(item int) bool {
	if !set.Contains(item) {
		set.mu.Lock()
		defer set.mu.Unlock()
		if set.data == nil {
			set.data = make(map[int]struct{})
		}
		if _, ok := set.data[item]; !ok {
			set.data[item] = struct{}{}
			return true
		}
	}
	return false
}

// AddIfNotExistFunc 检查项是否存在于集合中，
// 如果项不存在于集合中，且函数 `f` 返回 true，则将项添加到集合中并返回 true，否则什么都不做并返回 false。
//
// 注意，函数 `f` 在写入锁未获取的情况下执行。
// md5:7563a3cf864d8a2b
func (set *IntSet) AddIfNotExistFunc(item int, f func() bool) bool {
	if !set.Contains(item) {
		if f() {
			set.mu.Lock()
			defer set.mu.Unlock()
			if set.data == nil {
				set.data = make(map[int]struct{})
			}
			if _, ok := set.data[item]; !ok {
				set.data[item] = struct{}{}
				return true
			}
		}
	}
	return false
}

// AddIfNotExistFuncLock 检查项是否存在于集合中，
// 如果该项不存在于集合中并且函数 `f` 返回 true，那么它会将该项添加到集合中并返回 true；
// 否则，它不做任何操作并返回 false。
//
// 注意，函数 `f` 的执行不在写入锁的保护下进行。
// md5:48d67b0145855ed9
func (set *IntSet) AddIfNotExistFuncLock(item int, f func() bool) bool {
	if !set.Contains(item) {
		set.mu.Lock()
		defer set.mu.Unlock()
		if set.data == nil {
			set.data = make(map[int]struct{})
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

// Contains 检查集合是否包含 `item`。. md5:20a3bdc6aeef1d67
func (set *IntSet) Contains(item int) bool {
	var ok bool
	set.mu.RLock()
	if set.data != nil {
		_, ok = set.data[item]
	}
	set.mu.RUnlock()
	return ok
}

// Remove 从集合中删除 `item`。. md5:ab30c696cc44d190
func (set *IntSet) Remove(item int) {
	set.mu.Lock()
	if set.data != nil {
		delete(set.data, item)
	}
	set.mu.Unlock()
}

// Size 返回集合的大小。. md5:0d55ac576b7779ee
func (set *IntSet) Size() int {
	set.mu.RLock()
	l := len(set.data)
	set.mu.RUnlock()
	return l
}

// Clear 删除集合中的所有项。. md5:ce349f0cd3114465
func (set *IntSet) Clear() {
	set.mu.Lock()
	set.data = make(map[int]struct{})
	set.mu.Unlock()
}

// Slice 返回集合中的元素作为切片。. md5:f5bc80ac01ae812b
func (set *IntSet) Slice() []int {
	set.mu.RLock()
	var (
		i   = 0
		ret = make([]int, len(set.data))
	)
	for k := range set.data {
		ret[i] = k
		i++
	}
	set.mu.RUnlock()
	return ret
}

// Join 使用字符串 `glue` 连接多个项目。. md5:c8699391999ac788
func (set *IntSet) Join(glue string) string {
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

// String 将 items 转换为字符串，其实现方式类似于 json.Marshal。. md5:cedb10711c2e5dac
func (set *IntSet) String() string {
	if set == nil {
		return ""
	}
	return "[" + set.Join(",") + "]"
}

// LockFunc 使用回调函数 `f` 为写入操作加锁。. md5:85d746d8a49edab7
func (set *IntSet) LockFunc(f func(m map[int]struct{})) {
	set.mu.Lock()
	defer set.mu.Unlock()
	f(set.data)
}

// RLockFunc 使用回调函数 `f` 进行读取锁定。. md5:5fe2bf1a85ce319e
func (set *IntSet) RLockFunc(f func(m map[int]struct{})) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	f(set.data)
}

// Equal 检查两个集合是否相等。. md5:105ea4dd39b57fe8
func (set *IntSet) Equal(other *IntSet) bool {
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

// IsSubsetOf 检查当前集合是否为 `other` 的子集。. md5:333e392219846e17
func (set *IntSet) IsSubsetOf(other *IntSet) bool {
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

// Union 返回一个新集合，它是`set`和`other`的并集。
// 意味着，`newSet`中的所有项目都在`set`中或在`other`中。
// md5:420e241c3c12e8e6
func (set *IntSet) Union(others ...*IntSet) (newSet *IntSet) {
	newSet = NewIntSet()
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

// Diff 返回一个新的集合，它是 `set` 与 `other` 之间的差集。
// 这意味着，`newSet` 中的所有项目都在 `set` 中，但不在 `other` 中。
// md5:6779e6e007651b53
func (set *IntSet) Diff(others ...*IntSet) (newSet *IntSet) {
	newSet = NewIntSet()
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

// Intersect 返回一个新的集合，这个集合是 `set` 和 `other` 的交集。
// 这意味着，`newSet` 中的所有元素都既存在于 `set` 中也存在于 `other` 中。
// md5:327d3fcc12f06583
func (set *IntSet) Intersect(others ...*IntSet) (newSet *IntSet) {
	newSet = NewIntSet()
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

// Complement 返回一个新的集合，该集合是`set`在`full`中的补集。
// 换句话说，`newSet`中的所有元素都在`full`中但不在`set`中。
//
// 如果给定的集合`full`不是`set`的全集，它将返回`full`和`set`之间的差集。
// md5:7e76900d6f20af06
func (set *IntSet) Complement(full *IntSet) (newSet *IntSet) {
	newSet = NewIntSet()
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

// Merge 将 `others` 集合中的项目合并到 `set` 中。. md5:788b02e300c6f440
func (set *IntSet) Merge(others ...*IntSet) *IntSet {
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

// Sum 计算项目总和。
// 注意：项目应该转换为整数类型，
// 否则你可能会得到意想不到的结果。
// md5:979b37fbf86a5233
func (set *IntSet) Sum() (sum int) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k := range set.data {
		sum += k
	}
	return
}

// Pop 随机从集合中弹出一个元素。. md5:7e1906e951f13db1
func (set *IntSet) Pop() int {
	set.mu.Lock()
	defer set.mu.Unlock()
	for k := range set.data {
		delete(set.data, k)
		return k
	}
	return 0
}

// Pops 从集合中随机弹出 `size` 个元素。
// 如果 size == -1，它将返回所有元素。
// md5:c687f88e0a2df8f2
func (set *IntSet) Pops(size int) []int {
	set.mu.Lock()
	defer set.mu.Unlock()
	if size > len(set.data) || size == -1 {
		size = len(set.data)
	}
	if size <= 0 {
		return nil
	}
	index := 0
	array := make([]int, size)
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

// Walk应用用户提供的函数`f`到集合中的每一项。. md5:d6ceaae555e8a9e6
func (set *IntSet) Walk(f func(item int) int) *IntSet {
	set.mu.Lock()
	defer set.mu.Unlock()
	m := make(map[int]struct{}, len(set.data))
	for k, v := range set.data {
		m[f(k)] = v
	}
	set.data = m
	return set
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。. md5:43c3b36e60a18f9a
func (set IntSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(set.Slice())
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。. md5:f6766b88cf3d63c2
func (set *IntSet) UnmarshalJSON(b []byte) error {
	set.mu.Lock()
	defer set.mu.Unlock()
	if set.data == nil {
		set.data = make(map[int]struct{})
	}
	var array []int
	if err := json.UnmarshalUseNumber(b, &array); err != nil {
		return err
	}
	for _, v := range array {
		set.data[v] = struct{}{}
	}
	return nil
}

// UnmarshalValue 是一个接口实现，用于将任何类型的价值设置为集合。. md5:b119247f684920ad
func (set *IntSet) UnmarshalValue(value interface{}) (err error) {
	set.mu.Lock()
	defer set.mu.Unlock()
	if set.data == nil {
		set.data = make(map[int]struct{})
	}
	var array []int
	switch value.(type) {
	case string, []byte:
		err = json.UnmarshalUseNumber(gconv.Bytes(value), &array)
	default:
		array = gconv.SliceInt(value)
	}
	for _, v := range array {
		set.data[v] = struct{}{}
	}
	return
}

// DeepCopy实现当前类型的深拷贝接口。. md5:9cfbcb08109f6ce1
func (set *IntSet) DeepCopy() interface{} {
	if set == nil {
		return nil
	}
	set.mu.RLock()
	defer set.mu.RUnlock()
	var (
		slice = make([]int, len(set.data))
		index = 0
	)
	for k := range set.data {
		slice[index] = k
		index++
	}
	return NewIntSetFrom(slice, set.mu.IsSafe())
}
