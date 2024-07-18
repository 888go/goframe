// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gset提供了各种并发安全/不安全的集合。 md5:bcd5b9cf4b925a06
package gset

import (
	"bytes"

	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/internal/rwmutex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type Set struct {
	mu   rwmutex.RWMutex
	data map[interface{}]struct{}
}

// New 创建并返回一个新的集合，其中包含不重复的项目。
// 参数 `safe` 用于指定在并发安全模式下使用集合，其默认为 false。
// md5:db8312fdb3f679d3
// ff:创建
// safe:并发安全
func New(safe ...bool) *Set {
	return NewSet(safe...)
}

// NewSet 创建并返回一个新的集合，该集合包含不重复的项目。
// 另请参见 New。
// md5:3b8e2b58affe23e6
// ff:NewSet别名
// safe:
func NewSet(safe ...bool) *Set {
	return &Set{
		data: make(map[interface{}]struct{}),
		mu:   rwmutex.Create(safe...),
	}
}

// NewFrom 函数根据 `items` 创建一个新的集合。
// 参数 `items` 可以是任何类型的变量，或者是一个切片。
// md5:eab216208c4dc0bb
// ff:创建并按值
// items:值
// safe:并发安全
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

// Iterator 使用给定的回调函数 `f` 遍历只读集合，如果 `f` 返回 true，则继续遍历；否则停止。
// md5:b896360b1cf6fc88
// yx:true
// ff:X遍历
// set:
// f:
// v:
func (set *Set) Iterator(f func(v interface{}) bool) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k := range set.data {
		if !f(k) {
			break
		}
	}
}

// Add 将一个或多个项目添加到集合中。 md5:316141ff7d4b8e45
// ff:加入
// set:
// items:值s
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

// AddIfNotExist 检查项是否存在于集合中，
// 如果项不存在于集合中，则将其添加到集合并返回true，
// 否则不做任何操作并返回false。
//
// 注意，如果 `item` 为 nil，它将不做任何操作并返回false。
// md5:3d920a290d301fb9
// ff:加入值并跳过已存在
// set:
// item:值
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

// AddIfNotExistFunc 检查项目是否存在于集合中，
// 如果项目不在集合中并且函数 `f` 返回 true，那么它会将项目添加到集合并返回 true，否则不做任何操作并返回 false。
//
// 注意，如果 `item` 为 nil，它不做任何操作并返回 false。函数 `f` 在不持有写锁的情况下执行。
// md5:f80cf07184bee06f
// ff:加入值并跳过已存在_函数
// set:
// item:值
// f:
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

// AddIfNotExistFuncLock 检查项是否存在于集合中，
// 如果项不存在于集合中并且函数 `f` 返回 true，它将在集合中添加该项并返回 true，
// 否则什么也不做并返回 false。
//
// 注意，如果 `item` 为 nil，则什么也不做并返回 false。函数 `f` 在写锁保护下执行。
// md5:2a57dc990857b7b1
// ff:加入值并跳过已存在_并发安全函数
// set:
// item:值
// f:
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

// Contains 检查集合是否包含 `item`。 md5:20a3bdc6aeef1d67
// ff:是否存在
// set:
// item:值
func (set *Set) Contains(item interface{}) bool {
	var ok bool
	set.mu.RLock()
	if set.data != nil {
		_, ok = set.data[item]
	}
	set.mu.RUnlock()
	return ok
}

// Remove 从集合中删除 `item`。 md5:ab30c696cc44d190
// ff:删除
// set:
// item:值
func (set *Set) Remove(item interface{}) {
	set.mu.Lock()
	if set.data != nil {
		delete(set.data, item)
	}
	set.mu.Unlock()
}

// Size 返回集合的大小。 md5:0d55ac576b7779ee
// ff:取数量
// set:
func (set *Set) Size() int {
	set.mu.RLock()
	l := len(set.data)
	set.mu.RUnlock()
	return l
}

// Clear 删除集合中的所有项。 md5:ce349f0cd3114465
// ff:清空
// set:
func (set *Set) Clear() {
	set.mu.Lock()
	set.data = make(map[interface{}]struct{})
	set.mu.Unlock()
}

// Slice 返回集合中的所有项目作为切片。 md5:d07c46cf5dee2602
// ff:取集合切片
// set:
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

// Join 使用字符串 `glue` 连接多个项目。 md5:c8699391999ac788
// ff:取集合文本
// set:
// glue:连接符
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

// String 将 items 转换为字符串，其实现方式类似于 json.Marshal。 md5:cedb10711c2e5dac
// ff:
// set:
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

// LockFunc 使用回调函数 `f` 为写入操作加锁。 md5:85d746d8a49edab7
// ff:写锁定_函数
// set:
// f:
// m:
func (set *Set) LockFunc(f func(m map[interface{}]struct{})) {
	set.mu.Lock()
	defer set.mu.Unlock()
	f(set.data)
}

// RLockFunc 使用回调函数 `f` 进行读取锁定。 md5:5fe2bf1a85ce319e
// ff:读锁定_函数
// set:
// f:
// m:
func (set *Set) RLockFunc(f func(m map[interface{}]struct{})) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	f(set.data)
}

// Equal 检查两个集合是否相等。 md5:105ea4dd39b57fe8
// ff:是否相等
// set:
// other:待比较集合
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

// IsSubsetOf 检查当前集合是否为 `other` 的子集。 md5:333e392219846e17
// ff:是否为子集
// set:
// other:父集
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

// Union 返回一个新集合，它是`set`和`others`的并集。
// 意味着，新集合`newSet`中的所有项目都在`set`中或在`others`中。
// md5:81f60d9140026203
// ff:取并集
// set:
// others:集合
// newSet:新集合
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

// Diff 返回一个新的集合，它是 `set` 与 `others` 的差集。
// 意味着，新集合 `newSet` 中的所有项都在 `set` 中但不在 `others` 中。
// md5:0fe9ba09d007ac00
// ff:取差集
// set:
// others:集合
// newSet:新集合
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

// Intersect 返回一个新集合，该集合是将`set`与`others`进行交集运算的结果。
// 这意味着，新集合`newSet`中的所有元素都既存在于`set`中也存在于`others`中。
// md5:4db6ae5026f8dedc
// ff:取交集
// set:
// others:集合
// newSet:新集合
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

// Complement 返回一个新的集合，该集合是`set`在`full`中的补集。
// 换句话说，`newSet`中的所有元素都在`full`中但不在`set`中。
//
// 如果给定的集合`full`不是`set`的全集，它将返回`full`和`set`之间的差集。
// md5:7e76900d6f20af06
// ff:取补集
// set:
// full:集合
// newSet:新集合
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

// Merge 将 `others` 集合中的项目合并到 `set` 中。 md5:788b02e300c6f440
// ff:合并
// set:
// others:集合s
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

// Sum sums items.
// or you'd get a result that you unexpected.
// ff:求和
// set:
// sum:总和
func (set *Set) Sum() (sum int) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k := range set.data {
		sum += gconv.Int(k)
	}
	return
}

// Pop 随机从集合中弹出一个元素。 md5:7e1906e951f13db1
// ff:出栈
// set:
func (set *Set) Pop() interface{} {
	set.mu.Lock()
	defer set.mu.Unlock()
	for k := range set.data {
		delete(set.data, k)
		return k
	}
	return nil
}

// Pops 从集合中随机弹出 `size` 个元素。
// 如果 size == -1，它将返回所有元素。
// md5:c687f88e0a2df8f2
// ff:出栈多个
// set:
// size:数量
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

// Walk应用用户提供的函数`f`到集合中的每一项。 md5:d6ceaae555e8a9e6
// ff:遍历修改
// set:
// f:
// item:
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

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
// ff:
// set:
func (set Set) MarshalJSON() ([]byte, error) {
	return json.Marshal(set.Slice())
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
// ff:
// set:
// b:
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

// UnmarshalValue 是一个接口实现，用于将任何类型的价值设置为集合。 md5:b119247f684920ad
// ff:
// set:
// value:
// err:
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

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
// ff:
// set:
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
