// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gset提供了各种并发安全/不安全的集合。 md5:bcd5b9cf4b925a06
package 集合类

import (
	"bytes"

	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/rwmutex"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
)

type Set struct {
	mu   rwmutex.RWMutex
	data map[interface{}]struct{}
}

// X创建 创建并返回一个新的集合，其中包含不重复的项目。
// 参数 `safe` 用于指定在并发安全模式下使用集合，其默认为 false。
// md5:db8312fdb3f679d3
func X创建(并发安全 ...bool) *Set {
	return NewSet别名(并发安全...)
}

// NewSet别名 创建并返回一个新的集合，该集合包含不重复的项目。
// 另请参见 New。
// md5:3b8e2b58affe23e6
func NewSet别名(safe ...bool) *Set {
	return &Set{
		data: make(map[interface{}]struct{}),
		mu:   rwmutex.Create(safe...),
	}
}

// X创建并按值 函数根据 `items` 创建一个新的集合。
// 参数 `items` 可以是任何类型的变量，或者是一个切片。
// md5:eab216208c4dc0bb
func X创建并按值(值 interface{}, 并发安全 ...bool) *Set {
	m := make(map[interface{}]struct{})
	for _, v := range gconv.X取any切片(值) {
		m[v] = struct{}{}
	}
	return &Set{
		data: m,
		mu:   rwmutex.Create(并发安全...),
	}
}

// X遍历 使用给定的回调函数 `f` 遍历只读集合，如果 `f` 返回 true，则继续遍历；否则停止。
// md5:b896360b1cf6fc88
func (set *Set) X遍历(f func(v interface{}) bool) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k := range set.data {
		if !f(k) {
			break
		}
	}
}

// X加入 将一个或多个项目添加到集合中。 md5:316141ff7d4b8e45
func (set *Set) X加入(值s ...interface{}) {
	set.mu.Lock()
	if set.data == nil {
		set.data = make(map[interface{}]struct{})
	}
	for _, v := range 值s {
		set.data[v] = struct{}{}
	}
	set.mu.Unlock()
}

// X加入值并跳过已存在 检查项是否存在于集合中，
// 如果项不存在于集合中，则将其添加到集合并返回true，
// 否则不做任何操作并返回false。
//
// 注意，如果 `item` 为 nil，它将不做任何操作并返回false。
// md5:3d920a290d301fb9
func (set *Set) X加入值并跳过已存在(值 interface{}) bool {
	if 值 == nil {
		return false
	}
	if !set.X是否存在(值) {
		set.mu.Lock()
		defer set.mu.Unlock()
		if set.data == nil {
			set.data = make(map[interface{}]struct{})
		}
		if _, ok := set.data[值]; !ok {
			set.data[值] = struct{}{}
			return true
		}
	}
	return false
}

// X加入值并跳过已存在_函数 检查项目是否存在于集合中，
// 如果项目不在集合中并且函数 `f` 返回 true，那么它会将项目添加到集合并返回 true，否则不做任何操作并返回 false。
//
// 注意，如果 `item` 为 nil，它不做任何操作并返回 false。函数 `f` 在不持有写锁的情况下执行。
// md5:f80cf07184bee06f
func (set *Set) X加入值并跳过已存在_函数(值 interface{}, f func() bool) bool {
	if 值 == nil {
		return false
	}
	if !set.X是否存在(值) {
		if f() {
			set.mu.Lock()
			defer set.mu.Unlock()
			if set.data == nil {
				set.data = make(map[interface{}]struct{})
			}
			if _, ok := set.data[值]; !ok {
				set.data[值] = struct{}{}
				return true
			}
		}
	}
	return false
}

// X加入值并跳过已存在_并发安全函数 检查项是否存在于集合中，
// 如果项不存在于集合中并且函数 `f` 返回 true，它将在集合中添加该项并返回 true，
// 否则什么也不做并返回 false。
//
// 注意，如果 `item` 为 nil，则什么也不做并返回 false。函数 `f` 在写锁保护下执行。
// md5:2a57dc990857b7b1
func (set *Set) X加入值并跳过已存在_并发安全函数(值 interface{}, f func() bool) bool {
	if 值 == nil {
		return false
	}
	if !set.X是否存在(值) {
		set.mu.Lock()
		defer set.mu.Unlock()
		if set.data == nil {
			set.data = make(map[interface{}]struct{})
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

// X是否存在 检查集合是否包含 `item`。 md5:20a3bdc6aeef1d67
func (set *Set) X是否存在(值 interface{}) bool {
	var ok bool
	set.mu.RLock()
	if set.data != nil {
		_, ok = set.data[值]
	}
	set.mu.RUnlock()
	return ok
}

// X删除 从集合中删除 `item`。 md5:ab30c696cc44d190
func (set *Set) X删除(值 interface{}) {
	set.mu.Lock()
	if set.data != nil {
		delete(set.data, 值)
	}
	set.mu.Unlock()
}

// X取数量 返回集合的大小。 md5:0d55ac576b7779ee
func (set *Set) X取数量() int {
	set.mu.RLock()
	l := len(set.data)
	set.mu.RUnlock()
	return l
}

// X清空 删除集合中的所有项。 md5:ce349f0cd3114465
func (set *Set) X清空() {
	set.mu.Lock()
	set.data = make(map[interface{}]struct{})
	set.mu.Unlock()
}

// X取集合切片 返回集合中的所有项目作为切片。 md5:d07c46cf5dee2602
func (set *Set) X取集合切片() []interface{} {
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

// X取集合文本 使用字符串 `glue` 连接多个项目。 md5:c8699391999ac788
func (set *Set) X取集合文本(连接符 string) string {
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
			buffer.WriteString(连接符)
		}
		i++
	}
	return buffer.String()
}

// String 将 items 转换为字符串，其实现方式类似于 json.Marshal。 md5:cedb10711c2e5dac
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
		if gstr.X是否为数字(s) {
			buffer.WriteString(s)
		} else {
			buffer.WriteString(`"` + gstr.X转义并按字符(s, `"\`) + `"`)
		}
		if i != l-1 {
			buffer.WriteByte(',')
		}
		i++
	}
	buffer.WriteByte(']')
	return buffer.String()
}

// X写锁定_函数 使用回调函数 `f` 为写入操作加锁。 md5:85d746d8a49edab7
func (set *Set) X写锁定_函数(f func(m map[interface{}]struct{})) {
	set.mu.Lock()
	defer set.mu.Unlock()
	f(set.data)
}

// X读锁定_函数 使用回调函数 `f` 进行读取锁定。 md5:5fe2bf1a85ce319e
func (set *Set) X读锁定_函数(f func(m map[interface{}]struct{})) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	f(set.data)
}

// X是否相等 检查两个集合是否相等。 md5:105ea4dd39b57fe8
func (set *Set) X是否相等(待比较集合 *Set) bool {
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

// X是否为子集 检查当前集合是否为 `other` 的子集。 md5:333e392219846e17
func (set *Set) X是否为子集(父集 *Set) bool {
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

// X取并集 返回一个新集合，它是`set`和`others`的并集。
// 意味着，新集合`newSet`中的所有项目都在`set`中或在`others`中。
// md5:81f60d9140026203
func (set *Set) X取并集(集合 ...*Set) (新集合 *Set) {
	新集合 = NewSet别名()
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

// X取差集 返回一个新的集合，它是 `set` 与 `others` 的差集。
// 意味着，新集合 `newSet` 中的所有项都在 `set` 中但不在 `others` 中。
// md5:0fe9ba09d007ac00
func (set *Set) X取差集(集合 ...*Set) (新集合 *Set) {
	新集合 = NewSet别名()
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

// X取交集 返回一个新集合，该集合是将`set`与`others`进行交集运算的结果。
// 这意味着，新集合`newSet`中的所有元素都既存在于`set`中也存在于`others`中。
// md5:4db6ae5026f8dedc
func (set *Set) X取交集(集合 ...*Set) (新集合 *Set) {
	新集合 = NewSet别名()
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

// X取补集 返回一个新的集合，该集合是`set`在`full`中的补集。
// 换句话说，`newSet`中的所有元素都在`full`中但不在`set`中。
//
// 如果给定的集合`full`不是`set`的全集，它将返回`full`和`set`之间的差集。
// md5:7e76900d6f20af06
func (set *Set) X取补集(集合 *Set) (新集合 *Set) {
	新集合 = NewSet别名()
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

// X合并 将 `others` 集合中的项目合并到 `set` 中。 md5:788b02e300c6f440
func (set *Set) X合并(集合s ...*Set) *Set {
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

// X求和 计算项目总和。
// 注意：项目应该转换为整数类型，
// 否则你可能会得到意想不到的结果。
// md5:979b37fbf86a5233
func (set *Set) X求和() (总和 int) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k := range set.data {
		总和 += gconv.X取整数(k)
	}
	return
}

// X出栈 随机从集合中弹出一个元素。 md5:7e1906e951f13db1
func (set *Set) X出栈() interface{} {
	set.mu.Lock()
	defer set.mu.Unlock()
	for k := range set.data {
		delete(set.data, k)
		return k
	}
	return nil
}

// X出栈多个 从集合中随机弹出 `size` 个元素。
// 如果 size == -1，它将返回所有元素。
// md5:c687f88e0a2df8f2
func (set *Set) X出栈多个(数量 int) []interface{} {
	set.mu.Lock()
	defer set.mu.Unlock()
	if 数量 > len(set.data) || 数量 == -1 {
		数量 = len(set.data)
	}
	if 数量 <= 0 {
		return nil
	}
	index := 0
	array := make([]interface{}, 数量)
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

// X遍历修改应用用户提供的函数`f`到集合中的每一项。 md5:d6ceaae555e8a9e6
func (set *Set) X遍历修改(f func(item interface{}) interface{}) *Set {
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
func (set Set) MarshalJSON() ([]byte, error) {
	return json.Marshal(set.X取集合切片())
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
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
func (set *Set) UnmarshalValue(value interface{}) (err error) {
	set.mu.Lock()
	defer set.mu.Unlock()
	if set.data == nil {
		set.data = make(map[interface{}]struct{})
	}
	var array []interface{}
	switch value.(type) {
	case string, []byte:
		err = json.UnmarshalUseNumber(gconv.X取字节集(value), &array)
	default:
		array = gconv.SliceAny别名(value)
	}
	for _, v := range array {
		set.data[v] = struct{}{}
	}
	return
}

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
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
	return X创建并按值(data, set.mu.IsSafe())
}
