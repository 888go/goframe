// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

package gset

import (
	"bytes"
	"strings"

	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/internal/rwmutex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type StrSet struct {
	mu   rwmutex.RWMutex
	data map[string]struct{}
}

// NewStrSet 创建并返回一个新集合，其中包含不重复的元素。
// 参数 `safe` 用于指定是否在并发安全环境下使用集合，默认为 false。
// md5:b4b32102d4f1da78
func NewStrSet(safe ...bool) *StrSet {
	return &StrSet{
		mu:   rwmutex.Create(safe...),
		data: make(map[string]struct{}),
	}
}

// NewStrSetFrom 从`items`创建一个新的集合。. md5:6f9a406a984403d2
func NewStrSetFrom(items []string, safe ...bool) *StrSet {
	m := make(map[string]struct{})
	for _, v := range items {
		m[v] = struct{}{}
	}
	return &StrSet{
		mu:   rwmutex.Create(safe...),
		data: m,
	}
}

// Iterator 使用给定的回调函数 `f` 遍历只读集合，如果 `f` 返回 true，则继续遍历；否则停止。
// md5:b896360b1cf6fc88
func (set *StrSet) Iterator(f func(v string) bool) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k := range set.data {
		if !f(k) {
			break
		}
	}
}

// Add 将一个或多个项目添加到集合中。. md5:316141ff7d4b8e45
func (set *StrSet) Add(item ...string) {
	set.mu.Lock()
	if set.data == nil {
		set.data = make(map[string]struct{})
	}
	for _, v := range item {
		set.data[v] = struct{}{}
	}
	set.mu.Unlock()
}

// AddIfNotExist 检查项是否存在于集合中，
// 如果项不存在于集合中，它会将项添加到集合中并返回 true，否则什么都不做并返回 false。
// md5:9cff508c42cffd55
func (set *StrSet) AddIfNotExist(item string) bool {
	if !set.Contains(item) {
		set.mu.Lock()
		defer set.mu.Unlock()
		if set.data == nil {
			set.data = make(map[string]struct{})
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
func (set *StrSet) AddIfNotExistFunc(item string, f func() bool) bool {
	if !set.Contains(item) {
		if f() {
			set.mu.Lock()
			defer set.mu.Unlock()
			if set.data == nil {
				set.data = make(map[string]struct{})
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
func (set *StrSet) AddIfNotExistFuncLock(item string, f func() bool) bool {
	if !set.Contains(item) {
		set.mu.Lock()
		defer set.mu.Unlock()
		if set.data == nil {
			set.data = make(map[string]struct{})
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
func (set *StrSet) Contains(item string) bool {
	var ok bool
	set.mu.RLock()
	if set.data != nil {
		_, ok = set.data[item]
	}
	set.mu.RUnlock()
	return ok
}

// ContainsI 检查集合中是否存在某个值（忽略大小写）。
// 注意，它内部会遍历整个集合以进行不区分大小写的比较。
// md5:851e1bbfa6da1bae
func (set *StrSet) ContainsI(item string) bool {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k := range set.data {
		if strings.EqualFold(k, item) {
			return true
		}
	}
	return false
}

// Remove 从集合中删除 `item`。. md5:ab30c696cc44d190
func (set *StrSet) Remove(item string) {
	set.mu.Lock()
	if set.data != nil {
		delete(set.data, item)
	}
	set.mu.Unlock()
}

// Size 返回集合的大小。. md5:0d55ac576b7779ee
func (set *StrSet) Size() int {
	set.mu.RLock()
	l := len(set.data)
	set.mu.RUnlock()
	return l
}

// Clear 删除集合中的所有项。. md5:ce349f0cd3114465
func (set *StrSet) Clear() {
	set.mu.Lock()
	set.data = make(map[string]struct{})
	set.mu.Unlock()
}

// Slice 返回集合中的元素作为切片。. md5:f5bc80ac01ae812b
func (set *StrSet) Slice() []string {
	set.mu.RLock()
	var (
		i   = 0
		ret = make([]string, len(set.data))
	)
	for item := range set.data {
		ret[i] = item
		i++
	}

	set.mu.RUnlock()
	return ret
}

// Join 使用字符串 `glue` 连接多个项目。. md5:c8699391999ac788
func (set *StrSet) Join(glue string) string {
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
		buffer.WriteString(k)
		if i != l-1 {
			buffer.WriteString(glue)
		}
		i++
	}
	return buffer.String()
}

// String 将 items 转换为字符串，其实现方式类似于 json.Marshal。. md5:cedb10711c2e5dac
func (set *StrSet) String() string {
	if set == nil {
		return ""
	}
	set.mu.RLock()
	defer set.mu.RUnlock()
	var (
		l      = len(set.data)
		i      = 0
		buffer = bytes.NewBuffer(nil)
	)
	buffer.WriteByte('[')
	for k := range set.data {
		buffer.WriteString(`"` + gstr.QuoteMeta(k, `"\`) + `"`)
		if i != l-1 {
			buffer.WriteByte(',')
		}
		i++
	}
	buffer.WriteByte(']')
	return buffer.String()
}

// LockFunc 使用回调函数 `f` 为写入操作加锁。. md5:85d746d8a49edab7
func (set *StrSet) LockFunc(f func(m map[string]struct{})) {
	set.mu.Lock()
	defer set.mu.Unlock()
	f(set.data)
}

// RLockFunc 使用回调函数 `f` 进行读取锁定。. md5:5fe2bf1a85ce319e
func (set *StrSet) RLockFunc(f func(m map[string]struct{})) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	f(set.data)
}

// Equal 检查两个集合是否相等。. md5:105ea4dd39b57fe8
func (set *StrSet) Equal(other *StrSet) bool {
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
func (set *StrSet) IsSubsetOf(other *StrSet) bool {
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
func (set *StrSet) Union(others ...*StrSet) (newSet *StrSet) {
	newSet = NewStrSet()
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
func (set *StrSet) Diff(others ...*StrSet) (newSet *StrSet) {
	newSet = NewStrSet()
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
func (set *StrSet) Intersect(others ...*StrSet) (newSet *StrSet) {
	newSet = NewStrSet()
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
func (set *StrSet) Complement(full *StrSet) (newSet *StrSet) {
	newSet = NewStrSet()
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
func (set *StrSet) Merge(others ...*StrSet) *StrSet {
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
func (set *StrSet) Sum() (sum int) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k := range set.data {
		sum += gconv.Int(k)
	}
	return
}

// Pop 随机从集合中弹出一个元素。. md5:7e1906e951f13db1
func (set *StrSet) Pop() string {
	set.mu.Lock()
	defer set.mu.Unlock()
	for k := range set.data {
		delete(set.data, k)
		return k
	}
	return ""
}

// Pops 从集合中随机弹出 `size` 个元素。
// 如果 size == -1，它将返回所有元素。
// md5:c687f88e0a2df8f2
func (set *StrSet) Pops(size int) []string {
	set.mu.Lock()
	defer set.mu.Unlock()
	if size > len(set.data) || size == -1 {
		size = len(set.data)
	}
	if size <= 0 {
		return nil
	}
	index := 0
	array := make([]string, size)
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
func (set *StrSet) Walk(f func(item string) string) *StrSet {
	set.mu.Lock()
	defer set.mu.Unlock()
	m := make(map[string]struct{}, len(set.data))
	for k, v := range set.data {
		m[f(k)] = v
	}
	set.data = m
	return set
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。. md5:43c3b36e60a18f9a
func (set StrSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(set.Slice())
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。. md5:f6766b88cf3d63c2
func (set *StrSet) UnmarshalJSON(b []byte) error {
	set.mu.Lock()
	defer set.mu.Unlock()
	if set.data == nil {
		set.data = make(map[string]struct{})
	}
	var array []string
	if err := json.UnmarshalUseNumber(b, &array); err != nil {
		return err
	}
	for _, v := range array {
		set.data[v] = struct{}{}
	}
	return nil
}

// UnmarshalValue 是一个接口实现，用于将任何类型的价值设置为集合。. md5:b119247f684920ad
func (set *StrSet) UnmarshalValue(value interface{}) (err error) {
	set.mu.Lock()
	defer set.mu.Unlock()
	if set.data == nil {
		set.data = make(map[string]struct{})
	}
	var array []string
	switch value.(type) {
	case string, []byte:
		err = json.UnmarshalUseNumber(gconv.Bytes(value), &array)
	default:
		array = gconv.SliceStr(value)
	}
	for _, v := range array {
		set.data[v] = struct{}{}
	}
	return
}

// DeepCopy实现当前类型的深拷贝接口。. md5:9cfbcb08109f6ce1
func (set *StrSet) DeepCopy() interface{} {
	if set == nil {
		return nil
	}
	set.mu.RLock()
	defer set.mu.RUnlock()
	var (
		slice = make([]string, len(set.data))
		index = 0
	)
	for k := range set.data {
		slice[index] = k
		index++
	}
	return NewStrSetFrom(slice, set.mu.IsSafe())
}
