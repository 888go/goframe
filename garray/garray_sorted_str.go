// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package garray

import (
	"bytes"
	"math"
	"sort"
	"strings"
	
	"github.com/888go/goframe/garray/internal/json"
	"github.com/888go/goframe/garray/internal/rwmutex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

// SortedStrArray 是一个具有丰富特性的 Go 语言有序字符串数组。
// 默认情况下，它按升序排列，但可以通过设置自定义比较器来改变排序方式。
// 它包含一个并发安全/非安全切换选项，该选项应在初始化时设定，并且一旦设定后不可更改。
type SortedStrArray struct {
	mu         rwmutex.RWMutex
	array      []string
	unique     bool                  // 是否启用唯一特性(false)
	comparator func(a, b string) int // 比较函数（返回值：-1表示a小于b；0表示a等于b；1表示a大于b）
}

// NewSortedStrArray 创建并返回一个空的已排序数组。
// 参数`safe`用于指定是否在并发安全的情况下使用数组，默认为false。
func NewSortedStrArray(safe ...bool) *SortedStrArray {
	return NewSortedStrArraySize(0, safe...)
}

// NewSortedStrArrayComparator 创建并返回一个根据指定比较器排序的空数组。
// 参数`safe`用于指定是否在并发安全的情况下使用数组，默认为false。
func NewSortedStrArrayComparator(comparator func(a, b string) int, safe ...bool) *SortedStrArray {
	array := NewSortedStrArray(safe...)
	array.comparator = comparator
	return array
}

// NewSortedStrArraySize 根据给定的大小和容量创建并返回一个已排序的数组。
// 参数 `safe` 用于指定是否在并发安全环境下使用该数组，默认为 false。
func NewSortedStrArraySize(cap int, safe ...bool) *SortedStrArray {
	return &SortedStrArray{
		mu:         rwmutex.Create(safe...),
		array:      make([]string, 0, cap),
		comparator: defaultComparatorStr,
	}
}

// NewSortedStrArrayFrom 通过给定的切片`array`创建并返回一个已排序的数组。
// 参数`safe`用于指定是否在并发安全的情况下使用数组，默认为false。
func NewSortedStrArrayFrom(array []string, safe ...bool) *SortedStrArray {
	a := NewSortedStrArraySize(0, safe...)
	a.array = array
	quickSortStr(a.array, a.getComparator())
	return a
}

// NewSortedStrArrayFromCopy 通过复制给定切片 `array` 创建并返回一个已排序的数组。
// 参数 `safe` 用于指定是否在并发安全的情况下使用数组，默认为 false。
func NewSortedStrArrayFromCopy(array []string, safe ...bool) *SortedStrArray {
	newArray := make([]string, len(array))
	copy(newArray, array)
	return NewSortedStrArrayFrom(newArray, safe...)
}

// SetArray 将底层的切片数组设置为给定的 `array`。
func (a *SortedStrArray) SetArray(array []string) *SortedStrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.array = array
	quickSortStr(a.array, a.getComparator())
	return a
}

// At 通过指定的索引返回值。
// 如果给定的 `index` 超出了数组的范围，它将返回一个空字符串。
func (a *SortedStrArray) At(index int) (value string) {
	value, _ = a.Get(index)
	return
}

// Sort 函数用于将数组按升序排列。
// 参数 `reverse` 用于控制排序方式，若为 true 则按降序排列（默认为升序）。
func (a *SortedStrArray) Sort() *SortedStrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	quickSortStr(a.array, a.getComparator())
	return a
}

// Add 向有序数组中添加一个或多个值，数组始终保持有序。
// 它是函数 Append 的别名，请参阅 Append。
func (a *SortedStrArray) Add(values ...string) *SortedStrArray {
	return a.Append(values...)
}

// Append 向已排序的数组中添加一个或多个值，数组将始终保持有序。
func (a *SortedStrArray) Append(values ...string) *SortedStrArray {
	if len(values) == 0 {
		return a
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	for _, value := range values {
		index, cmp := a.binSearch(value, false)
		if a.unique && cmp == 0 {
			continue
		}
		if index < 0 {
			a.array = append(a.array, value)
			continue
		}
		if cmp > 0 {
			index++
		}
		rear := append([]string{}, a.array[index:]...)
		a.array = append(a.array[0:index], value)
		a.array = append(a.array, rear...)
	}
	return a
}

// Get 通过指定的索引返回值。
// 如果给定的 `index` 超出了数组的范围，那么 `found` 将为 false。
func (a *SortedStrArray) Get(index int) (value string, found bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if index < 0 || index >= len(a.array) {
		return "", false
	}
	return a.array[index], true
}

// Remove 通过索引移除一个元素。
// 如果给定的 `index` 超出了数组范围，`found` 将为 false。
func (a *SortedStrArray) Remove(index int) (value string, found bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.doRemoveWithoutLock(index)
}

// doRemoveWithoutLock 在没有加锁的情况下通过索引移除一个项。
func (a *SortedStrArray) doRemoveWithoutLock(index int) (value string, found bool) {
	if index < 0 || index >= len(a.array) {
		return "", false
	}
	// 确定删除时的数组边界以提高删除效率
	if index == 0 {
		value := a.array[0]
		a.array = a.array[1:]
		return value, true
	} else if index == len(a.array)-1 {
		value := a.array[index]
		a.array = a.array[:index]
		return value, true
	}
// 如果这是一个非边界删除，
// 那么它将涉及创建一个数组，
// 因此，删除操作效率较低。
	value = a.array[index]
	a.array = append(a.array[:index], a.array[index+1:]...)
	return value, true
}

// RemoveValue 通过值移除一个元素。
// 若在数组中找到该值，则返回 true，否则（未找到时）返回 false。
func (a *SortedStrArray) RemoveValue(value string) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if i, r := a.binSearch(value, false); r == 0 {
		_, res := a.doRemoveWithoutLock(i)
		return res
	}
	return false
}

// RemoveValues 通过 `values` 移除项目。
func (a *SortedStrArray) RemoveValues(values ...string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	for _, value := range values {
		if i, r := a.binSearch(value, false); r == 0 {
			a.doRemoveWithoutLock(i)
		}
	}
}

// PopLeft 从数组开头弹出并返回一个元素。
// 注意，如果数组为空，则 `found` 为 false。
func (a *SortedStrArray) PopLeft() (value string, found bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if len(a.array) == 0 {
		return "", false
	}
	value = a.array[0]
	a.array = a.array[1:]
	return value, true
}

// PopRight从数组的末尾弹出并返回一个元素。
// 注意，如果数组为空，则`found`为false。
func (a *SortedStrArray) PopRight() (value string, found bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	index := len(a.array) - 1
	if index < 0 {
		return "", false
	}
	value = a.array[index]
	a.array = a.array[:index]
	return value, true
}

// PopRand 随机地从数组中弹出并返回一个元素。
// 注意，如果数组为空，则 `found` 为 false。
func (a *SortedStrArray) PopRand() (value string, found bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.doRemoveWithoutLock(grand.Intn(len(a.array)))
}

// PopRands 随机地从数组中弹出并返回 `size` 个元素。
// 若给定的 `size` 大于数组的大小，则返回数组中的所有元素。
// 注意，如果给定的 `size` 小于等于0或者数组为空，它将返回 nil。
func (a *SortedStrArray) PopRands(size int) []string {
	a.mu.Lock()
	defer a.mu.Unlock()
	if size <= 0 || len(a.array) == 0 {
		return nil
	}
	if size >= len(a.array) {
		size = len(a.array)
	}
	array := make([]string, size)
	for i := 0; i < size; i++ {
		array[i], _ = a.doRemoveWithoutLock(grand.Intn(len(a.array)))
	}
	return array
}

// PopLefts 从数组开头弹出并返回 `size` 个元素。
// 如果给定的 `size` 大于数组的大小，则返回数组中的所有元素。
// 注意，如果给定的 `size` 小于等于0或者数组为空，则返回nil。
func (a *SortedStrArray) PopLefts(size int) []string {
	a.mu.Lock()
	defer a.mu.Unlock()
	if size <= 0 || len(a.array) == 0 {
		return nil
	}
	if size >= len(a.array) {
		array := a.array
		a.array = a.array[:0]
		return array
	}
	value := a.array[0:size]
	a.array = a.array[size:]
	return value
}

// PopRights 从数组末尾弹出并返回 `size` 个元素。
// 若给定的 `size` 大于数组的大小，则返回数组中所有元素。
// 注意，如果给定的 `size` 小于等于0或者数组为空，它将返回nil。
func (a *SortedStrArray) PopRights(size int) []string {
	a.mu.Lock()
	defer a.mu.Unlock()
	if size <= 0 || len(a.array) == 0 {
		return nil
	}
	index := len(a.array) - size
	if index <= 0 {
		array := a.array
		a.array = a.array[:0]
		return array
	}
	value := a.array[index:]
	a.array = a.array[:index]
	return value
}

// Range 函数通过范围选择并返回数组中的元素，类似于 array[start:end]。
// 注意：在并发安全的使用场景下，它会返回一个原数据的副本；否则，返回的是底层数据的指针。
//
// 如果 `end` 为负数，则偏移量将从数组末尾开始计算。
// 如果省略了 `end`，则序列将包含从 start 开始直到数组末尾的所有元素。
func (a *SortedStrArray) Range(start int, end ...int) []string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	offsetEnd := len(a.array)
	if len(end) > 0 && end[0] < offsetEnd {
		offsetEnd = end[0]
	}
	if start > offsetEnd {
		return nil
	}
	if start < 0 {
		start = 0
	}
	array := ([]string)(nil)
	if a.mu.IsSafe() {
		array = make([]string, offsetEnd-start)
		copy(array, a.array[start:offsetEnd])
	} else {
		array = a.array[start:offsetEnd]
	}
	return array
}

// SubSlice 返回数组中由 `offset` 和 `size` 参数指定的元素子序列，并将其作为切片。
// 若在并发安全场景下使用，返回该切片的副本；否则返回指向切片的指针。
//
// 如果 offset 非负，则序列从数组该偏移位置开始。
// 如果 offset 为负，则序列从数组末尾向前偏移该距离的位置开始。
//
// 如果提供了 length 并且为正数，则序列将包含最多该数量的元素。
// 若数组长度小于 length，则序列仅包含数组中可获得的元素。
// 如果 length 为负数，则序列将在数组末尾向前停在该距离的位置。
// 如果未提供 length，则序列包含从 offset 开始直到数组末尾的所有元素。
//
// 若有任何可能穿越数组左边界的情况，函数将失败。
func (a *SortedStrArray) SubSlice(offset int, length ...int) []string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	size := len(a.array)
	if len(length) > 0 {
		size = length[0]
	}
	if offset > len(a.array) {
		return nil
	}
	if offset < 0 {
		offset = len(a.array) + offset
		if offset < 0 {
			return nil
		}
	}
	if size < 0 {
		offset += size
		size = -size
		if offset < 0 {
			return nil
		}
	}
	end := offset + size
	if end > len(a.array) {
		end = len(a.array)
		size = len(a.array) - offset
	}
	if a.mu.IsSafe() {
		s := make([]string, size)
		copy(s, a.array[offset:])
		return s
	} else {
		return a.array[offset:end]
	}
}

// Sum 返回数组中所有值的和。
func (a *SortedStrArray) Sum() (sum int) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		sum += gconv.Int(v)
	}
	return
}

// Len 返回数组的长度。
func (a *SortedStrArray) Len() int {
	a.mu.RLock()
	length := len(a.array)
	a.mu.RUnlock()
	return length
}

// Slice 返回数组的基础数据。
// 注意，如果它在并发安全的使用场景下，会返回基础数据的一个副本，
// 否则，则返回指向基础数据的指针。
func (a *SortedStrArray) Slice() []string {
	array := ([]string)(nil)
	if a.mu.IsSafe() {
		a.mu.RLock()
		defer a.mu.RUnlock()
		array = make([]string, len(a.array))
		copy(array, a.array)
	} else {
		array = a.array
	}
	return array
}

// Interfaces 函数将当前数组转换为 []interface{} 类型并返回。
func (a *SortedStrArray) Interfaces() []interface{} {
	a.mu.RLock()
	defer a.mu.RUnlock()
	array := make([]interface{}, len(a.array))
	for k, v := range a.array {
		array[k] = v
	}
	return array
}

// Contains 检查某个值是否存在于数组中。
func (a *SortedStrArray) Contains(value string) bool {
	return a.Search(value) != -1
}

// ContainsI 检查一个值是否以不区分大小写的方式存在于数组中。
// 注意，它在内部会遍历整个数组并进行不区分大小写的比较。
func (a *SortedStrArray) ContainsI(value string) bool {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if len(a.array) == 0 {
		return false
	}
	for _, v := range a.array {
		if strings.EqualFold(v, value) {
			return true
		}
	}
	return false
}

// Search 在数组中通过 `value` 进行搜索，返回 `value` 的索引，
// 若不存在，则返回 -1。
func (a *SortedStrArray) Search(value string) (index int) {
	if i, r := a.binSearch(value, true); r == 0 {
		return i
	}
	return -1
}

// 二分查找
// 返回最后比较的索引以及查找结果
// 若`result`等于0，表示在`index`位置的值等于`value`
// 若`result`小于0，表示在`index`位置的值小于`value`
// 若`result`大于0，表示在`index`位置的值大于`value`
func (a *SortedStrArray) binSearch(value string, lock bool) (index int, result int) {
	if lock {
		a.mu.RLock()
		defer a.mu.RUnlock()
	}
	if len(a.array) == 0 {
		return -1, -2
	}
	min := 0
	max := len(a.array) - 1
	mid := 0
	cmp := -2
	for min <= max {
		mid = min + int((max-min)/2)
		cmp = a.getComparator()(value, a.array[mid])
		switch {
		case cmp < 0:
			max = mid - 1
		case cmp > 0:
			min = mid + 1
		default:
			return mid, cmp
		}
	}
	return mid, cmp
}

// SetUnique 将唯一标志设置到数组中，
// 这意味着该数组不包含任何重复的元素。
// 同时，它还会进行唯一性检查，并移除所有重复的项。
func (a *SortedStrArray) SetUnique(unique bool) *SortedStrArray {
	oldUnique := a.unique
	a.unique = unique
	if unique && oldUnique != unique {
		a.Unique()
	}
	return a
}

// Unique 对数组进行去重，清除重复的元素。
func (a *SortedStrArray) Unique() *SortedStrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	if len(a.array) == 0 {
		return a
	}
	i := 0
	for {
		if i == len(a.array)-1 {
			break
		}
		if a.getComparator()(a.array[i], a.array[i+1]) == 0 {
			a.array = append(a.array[:i+1], a.array[i+1+1:]...)
		} else {
			i++
		}
	}
	return a
}

// Clone 返回一个新的数组，它是当前数组的一个副本。
func (a *SortedStrArray) Clone() (newArray *SortedStrArray) {
	a.mu.RLock()
	array := make([]string, len(a.array))
	copy(array, a.array)
	a.mu.RUnlock()
	return NewSortedStrArrayFrom(array, a.mu.IsSafe())
}

// 清空删除当前数组中的所有元素。
func (a *SortedStrArray) Clear() *SortedStrArray {
	a.mu.Lock()
	if len(a.array) > 0 {
		a.array = make([]string, 0)
	}
	a.mu.Unlock()
	return a
}

// LockFunc 通过回调函数`f`进行写入锁定。
func (a *SortedStrArray) LockFunc(f func(array []string)) *SortedStrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	f(a.array)
	return a
}

// RLockFunc 通过回调函数`f`锁定读取操作。
func (a *SortedStrArray) RLockFunc(f func(array []string)) *SortedStrArray {
	a.mu.RLock()
	defer a.mu.RUnlock()
	f(a.array)
	return a
}

// Merge 将`array`合并到当前数组中。
// 参数`array`可以是任何garray类型或切片类型。
// Merge 和 Append 的区别在于，Append 仅支持特定类型的切片作为参数，
// 而 Merge 支持更多类型的参数。
func (a *SortedStrArray) Merge(array interface{}) *SortedStrArray {
	return a.Add(gconv.Strings(array)...)
}

// Chunk 函数将一个数组分割成多个子数组，
// 每个子数组的大小由参数 `size` 确定。
// 最后一个子数组可能包含少于 size 个元素。
func (a *SortedStrArray) Chunk(size int) [][]string {
	if size < 1 {
		return nil
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	length := len(a.array)
	chunks := int(math.Ceil(float64(length) / float64(size)))
	var n [][]string
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		n = append(n, a.array[i*size:end])
		i++
	}
	return n
}

// Rand 随机地从数组中返回一个元素（不删除）。
func (a *SortedStrArray) Rand() (value string, found bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if len(a.array) == 0 {
		return "", false
	}
	return a.array[grand.Intn(len(a.array))], true
}

// Rands 随机返回数组中的 `size` 个元素（不删除）。
func (a *SortedStrArray) Rands(size int) []string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if size <= 0 || len(a.array) == 0 {
		return nil
	}
	array := make([]string, size)
	for i := 0; i < size; i++ {
		array[i] = a.array[grand.Intn(len(a.array))]
	}
	return array
}

// Join 通过字符串 `glue` 连接数组元素。
func (a *SortedStrArray) Join(glue string) string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if len(a.array) == 0 {
		return ""
	}
	buffer := bytes.NewBuffer(nil)
	for k, v := range a.array {
		buffer.WriteString(v)
		if k != len(a.array)-1 {
			buffer.WriteString(glue)
		}
	}
	return buffer.String()
}

// CountValues 计算数组中所有值出现的次数。
func (a *SortedStrArray) CountValues() map[string]int {
	m := make(map[string]int)
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		m[v]++
	}
	return m
}

// Iterator 是 IteratorAsc 的别名。
func (a *SortedStrArray) Iterator(f func(k int, v string) bool) {
	a.IteratorAsc(f)
}

// IteratorAsc 以升序遍历给定数组，并使用回调函数 `f` 进行只读操作。
// 如果 `f` 返回 true，则继续迭代；若返回 false，则停止遍历。
func (a *SortedStrArray) IteratorAsc(f func(k int, v string) bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for k, v := range a.array {
		if !f(k, v) {
			break
		}
	}
}

// IteratorDesc 函数以降序遍历给定的数组，并使用指定回调函数 `f` 进行只读操作。
// 若 `f` 返回 true，则继续迭代；若返回 false，则停止迭代。
func (a *SortedStrArray) IteratorDesc(f func(k int, v string) bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for i := len(a.array) - 1; i >= 0; i-- {
		if !f(i, a.array[i]) {
			break
		}
	}
}

// String 方法将当前数组以字符串形式返回，其实现方式类似于 json.Marshal。
func (a *SortedStrArray) String() string {
	if a == nil {
		return ""
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	buffer := bytes.NewBuffer(nil)
	buffer.WriteByte('[')
	for k, v := range a.array {
		buffer.WriteString(`"` + gstr.QuoteMeta(v, `"\`) + `"`)
		if k != len(a.array)-1 {
			buffer.WriteByte(',')
		}
	}
	buffer.WriteByte(']')
	return buffer.String()
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
// 注意：此处接收者不使用指针。
func (a SortedStrArray) MarshalJSON() ([]byte, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return json.Marshal(a.array)
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func (a *SortedStrArray) UnmarshalJSON(b []byte) error {
	if a.comparator == nil {
		a.array = make([]string, 0)
		a.comparator = defaultComparatorStr
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	if err := json.UnmarshalUseNumber(b, &a.array); err != nil {
		return err
	}
	if a.array != nil {
		sort.Strings(a.array)
	}
	return nil
}

// UnmarshalValue 实现了一个接口，该接口用于为数组设置任意类型的值。
func (a *SortedStrArray) UnmarshalValue(value interface{}) (err error) {
	if a.comparator == nil {
		a.comparator = defaultComparatorStr
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	switch value.(type) {
	case string, []byte:
		err = json.UnmarshalUseNumber(gconv.Bytes(value), &a.array)
	default:
		a.array = gconv.SliceStr(value)
	}
	if a.array != nil {
		sort.Strings(a.array)
	}
	return err
}

// Filter 对数组进行迭代，并通过自定义回调函数进行元素过滤。
// 如果回调函数 `filter` 返回 true，则从数组中移除该元素；
// 否则不做任何处理并继续迭代。
func (a *SortedStrArray) Filter(filter func(index int, value string) bool) *SortedStrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i := 0; i < len(a.array); {
		if filter(i, a.array[i]) {
			a.array = append(a.array[:i], a.array[i+1:]...)
		} else {
			i++
		}
	}
	return a
}

// FilterEmpty 从数组中移除所有空字符串值。
func (a *SortedStrArray) FilterEmpty() *SortedStrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i := 0; i < len(a.array); {
		if a.array[i] == "" {
			a.array = append(a.array[:i], a.array[i+1:]...)
		} else {
			break
		}
	}
	for i := len(a.array) - 1; i >= 0; {
		if a.array[i] == "" {
			a.array = append(a.array[:i], a.array[i+1:]...)
		} else {
			break
		}
	}
	return a
}

// Walk 对数组中的每一项应用用户提供的函数 `f`。
func (a *SortedStrArray) Walk(f func(value string) string) *SortedStrArray {
	a.mu.Lock()
	defer a.mu.Unlock()

	// 保持数组始终有序。
	defer quickSortStr(a.array, a.getComparator())

	for i, v := range a.array {
		a.array[i] = f(v)
	}
	return a
}

// IsEmpty 检查数组是否为空。
func (a *SortedStrArray) IsEmpty() bool {
	return a.Len() == 0
}

// getComparator 返回之前设置的比较器，如果没有设置过，则返回一个默认的比较器。
func (a *SortedStrArray) getComparator() func(a, b string) int {
	if a.comparator == nil {
		return defaultComparatorStr
	}
	return a.comparator
}

// DeepCopy 实现接口，用于当前类型的深度复制。
func (a *SortedStrArray) DeepCopy() interface{} {
	if a == nil {
		return nil
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	newSlice := make([]string, len(a.array))
	copy(newSlice, a.array)
	return NewSortedStrArrayFrom(newSlice, a.mu.IsSafe())
}