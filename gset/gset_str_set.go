// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//

package 集合类

import (
	"bytes"
	"strings"
	
	"github.com/888go/goframe/gset/internal/json"
	"github.com/888go/goframe/gset/internal/rwmutex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type StrSet struct {
	mu   rwmutex.RWMutex
	data map[string]struct{}
}

// NewStrSet 创建并返回一个新的不包含重复项的集合。
// 参数`safe`用于指定是否在并发安全的情况下使用集合，默认为false。
// 这里，NewStrSet函数用于创建一个字符串集合，并确保其中的元素互不重复。该函数接受一个可选参数`safe`，它是一个布尔值，表示是否需要保证在并发环境下的安全性。如果不特别指定，那么默认情况下这个集合是不保证并发安全的。
func X创建文本(并发安全 ...bool) *StrSet {
	return &StrSet{
		mu:   rwmutex.Create(并发安全...),
		data: make(map[string]struct{}),
	}
}

// NewStrSetFrom 从 `items` 中创建并返回一个新的集合。
func X创建文本并按值(值 []string, 并发安全 ...bool) *StrSet {
	m := make(map[string]struct{})
	for _, v := range 值 {
		m[v] = struct{}{}
	}
	return &StrSet{
		mu:   rwmutex.Create(并发安全...),
		data: m,
	}
}

// Iterator 使用给定的回调函数`f`对集合进行只读遍历，
// 如果`f`返回true，则继续遍历；若返回false，则停止遍历。
func (set *StrSet) X遍历(f func(v string) bool) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k := range set.data {
		if !f(k) {
			break
		}
	}
}

// Add 向集合中添加一个或多个项目。
func (set *StrSet) X加入(值s ...string) {
	set.mu.Lock()
	if set.data == nil {
		set.data = make(map[string]struct{})
	}
	for _, v := range 值s {
		set.data[v] = struct{}{}
	}
	set.mu.Unlock()
}

// AddIfNotExist 检查项目是否已存在于集合中，
// 若该项目不在集合中，则将其添加到集合并返回 true；
// 否则，不做任何操作并返回 false。
func (set *StrSet) X加入值并跳过已存在(值 string) bool {
	if !set.X是否存在(值) {
		set.mu.Lock()
		defer set.mu.Unlock()
		if set.data == nil {
			set.data = make(map[string]struct{})
		}
		if _, ok := set.data[值]; !ok {
			set.data[值] = struct{}{}
			return true
		}
	}
	return false
}

// AddIfNotExistFunc 检查项是否已存在于集合中，
// 如果该项不存在于集合中，并且函数 `f` 返回 true，则将该项添加到集合并返回 true；
// 否则，不执行任何操作并返回 false。
//
// 注意，函数 `f` 在无写入锁的情况下执行。
func (set *StrSet) X加入值并跳过已存在_函数(值 string, f func() bool) bool {
	if !set.X是否存在(值) {
		if f() {
			set.mu.Lock()
			defer set.mu.Unlock()
			if set.data == nil {
				set.data = make(map[string]struct{})
			}
			if _, ok := set.data[值]; !ok {
				set.data[值] = struct{}{}
				return true
			}
		}
	}
	return false
}

// AddIfNotExistFuncLock 检查项是否存在集合中，
// 如果该项不存在于集合中，并且函数 `f` 返回 true，则将该项添加到集合并返回 true；
// 否则，不执行任何操作并返回 false。
//
// 注意，函数 `f` 在无写入锁的情况下执行。
func (set *StrSet) X加入值并跳过已存在_并发安全函数(值 string, f func() bool) bool {
	if !set.X是否存在(值) {
		set.mu.Lock()
		defer set.mu.Unlock()
		if set.data == nil {
			set.data = make(map[string]struct{})
		}
		if f() {
			if _, ok := set.data[值]; !ok {
				set.data[值] = struct{}{}
				return true
			}
		}
	}
	return false
}

// Contains 检查集合中是否包含 `item`。
func (set *StrSet) X是否存在(值 string) bool {
	var ok bool
	set.mu.RLock()
	if set.data != nil {
		_, ok = set.data[值]
	}
	set.mu.RUnlock()
	return ok
}

// ContainsI 检查某个值是否以不区分大小写的方式存在于集合中。
// 注意：它内部会遍历整个集合，以不区分大小写的方式进行比较。
func (set *StrSet) X是否存在并忽略大小写(值 string) bool {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k := range set.data {
		if strings.EqualFold(k, 值) {
			return true
		}
	}
	return false
}

// Remove 从集合中删除`item`。
func (set *StrSet) X删除(值 string) {
	set.mu.Lock()
	if set.data != nil {
		delete(set.data, 值)
	}
	set.mu.Unlock()
}

// Size 返回集合的大小。
func (set *StrSet) X取数量() int {
	set.mu.RLock()
	l := len(set.data)
	set.mu.RUnlock()
	return l
}

// 清除删除集合中的所有项。
func (set *StrSet) X清空() {
	set.mu.Lock()
	set.data = make(map[string]struct{})
	set.mu.Unlock()
}

// Slice 返回集合中项目的切片形式。
func (set *StrSet) X取集合数组() []string {
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

// Join通过字符串`glue`连接items。
func (set *StrSet) X取集合文本(连接符 string) string {
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
			buffer.WriteString(连接符)
		}
		i++
	}
	return buffer.String()
}

// String 返回 items 作为字符串，其实现方式类似于 json.Marshal。
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

// LockFunc 使用回调函数`f`进行写入锁定。
func (set *StrSet) X写锁定_函数(f func(m map[string]struct{})) {
	set.mu.Lock()
	defer set.mu.Unlock()
	f(set.data)
}

// RLockFunc 通过回调函数 `f` 对读取进行加锁。
func (set *StrSet) X读锁定_函数(f func(m map[string]struct{})) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	f(set.data)
}

// Equal 检查两个集合是否相等。
func (set *StrSet) X是否相等(待比较集合 *StrSet) bool {
	if set == 待比较集合 {
		return true
	}
	set.mu.RLock()
	defer set.mu.RUnlock()
	待比较集合.mu.RLock()
	defer 待比较集合.mu.RUnlock()
	if len(set.data) != len(待比较集合.data) {
		return false
	}
	for key := range set.data {
		if _, ok := 待比较集合.data[key]; !ok {
			return false
		}
	}
	return true
}

// IsSubsetOf 检查当前集合是否为 `other` 的子集。
func (set *StrSet) X是否为子集(父集 *StrSet) bool {
	if set == 父集 {
		return true
	}
	set.mu.RLock()
	defer set.mu.RUnlock()
	父集.mu.RLock()
	defer 父集.mu.RUnlock()
	for key := range set.data {
		if _, ok := 父集.data[key]; !ok {
			return false
		}
	}
	return true
}

// Union 返回一个新的集合，该集合是 `set` 和 `other` 的并集。
// 这意味着，`newSet` 中的所有元素都在 `set` 或者 `other` 中。
func (set *StrSet) X取并集(集合 ...*StrSet) (新集合 *StrSet) {
	新集合 = X创建文本()
	set.mu.RLock()
	defer set.mu.RUnlock()
	for _, other := range 集合 {
		if set != other {
			other.mu.RLock()
		}
		for k, v := range set.data {
			新集合.data[k] = v
		}
		if set != other {
			for k, v := range other.data {
				新集合.data[k] = v
			}
		}
		if set != other {
			other.mu.RUnlock()
		}
	}

	return
}

// Diff 返回一个新的集合，这个集合是 `set` 与 `other` 的差集。
// 这意味着，新集合 `newSet` 中的所有元素都在 `set` 中，但不在 `other` 中。
func (set *StrSet) X取差集(集合 ...*StrSet) (新集合 *StrSet) {
	新集合 = X创建文本()
	set.mu.RLock()
	defer set.mu.RUnlock()
	for _, other := range 集合 {
		if set == other {
			continue
		}
		other.mu.RLock()
		for k, v := range set.data {
			if _, ok := other.data[k]; !ok {
				新集合.data[k] = v
			}
		}
		other.mu.RUnlock()
	}
	return
}

// Intersect 返回一个新的集合，它是从 `set` 到 `other` 的交集。
// 这意味着，`newSet` 中的所有元素都在 `set` 中，并且也在 `other` 中。
func (set *StrSet) X取交集(集合 ...*StrSet) (新集合 *StrSet) {
	新集合 = X创建文本()
	set.mu.RLock()
	defer set.mu.RUnlock()
	for _, other := range 集合 {
		if set != other {
			other.mu.RLock()
		}
		for k, v := range set.data {
			if _, ok := other.data[k]; ok {
				新集合.data[k] = v
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
func (set *StrSet) X取补集(集合 *StrSet) (新集合 *StrSet) {
	新集合 = X创建文本()
	set.mu.RLock()
	defer set.mu.RUnlock()
	if set != 集合 {
		集合.mu.RLock()
		defer 集合.mu.RUnlock()
	}
	for k, v := range 集合.data {
		if _, ok := set.data[k]; !ok {
			新集合.data[k] = v
		}
	}
	return
}

// Merge 将 `others` 中的元素合并到 `set` 中。
func (set *StrSet) X合并(集合s ...*StrSet) *StrSet {
	set.mu.Lock()
	defer set.mu.Unlock()
	for _, other := range 集合s {
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
func (set *StrSet) X求和() (总和 int) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k := range set.data {
		总和 += gconv.Int(k)
	}
	return
}

// Pop 随机地从集合中弹出一个元素。
func (set *StrSet) X出栈() string {
	set.mu.Lock()
	defer set.mu.Unlock()
	for k := range set.data {
		delete(set.data, k)
		return k
	}
	return ""
}

// Pops 随机地从集合中弹出 `size` 个元素。
// 如果 size == -1，则返回所有元素。
func (set *StrSet) X出栈多个(数量 int) []string {
	set.mu.Lock()
	defer set.mu.Unlock()
	if 数量 > len(set.data) || 数量 == -1 {
		数量 = len(set.data)
	}
	if 数量 <= 0 {
		return nil
	}
	index := 0
	array := make([]string, 数量)
	for k := range set.data {
		delete(set.data, k)
		array[index] = k
		index++
		if index == 数量 {
			break
		}
	}
	return array
}

// Walk 对集合中的每一个元素应用用户提供的函数 `f`。
func (set *StrSet) X遍历修改(f func(值 string) string) *StrSet {
	set.mu.Lock()
	defer set.mu.Unlock()
	m := make(map[string]struct{}, len(set.data))
	for k, v := range set.data {
		m[f(k)] = v
	}
	set.data = m
	return set
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (set StrSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(set.X取集合数组())
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
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

// UnmarshalValue 是一个接口实现，用于为 set 设置任意类型的值。
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

// DeepCopy 实现接口，用于当前类型的深度复制。
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
	return X创建文本并按值(slice, set.mu.IsSafe())
}
