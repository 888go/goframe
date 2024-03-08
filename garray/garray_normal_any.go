// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package garray

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/888go/goframe/garray/internal/deepcopy"
	"github.com/888go/goframe/garray/internal/empty"
	"github.com/888go/goframe/garray/internal/json"
	"github.com/888go/goframe/garray/internal/rwmutex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

// Array 是一个具有丰富特性的 Go 语言数组。
// 它包含一个并发安全/不安全切换开关，该开关应在初始化时设置，并且之后不可更改。
type Array struct {
	mu    rwmutex.RWMutex
	array []interface{}
}

// New 创建并返回一个空数组。
// 参数 `safe` 用于指定是否在并发安全的情况下使用数组，默认为 false。
func New(safe ...bool) *Array {
	return NewArraySize(0, 0, safe...)
}

// NewArray 是 New 的别名，请参阅 New。
func NewArray(safe ...bool) *Array {
	return NewArraySize(0, 0, safe...)
}

// NewArraySize 根据给定的大小和容量创建并返回一个数组。
// 参数`safe`用于指定是否在并发安全的情况下使用数组，默认为false。
func NewArraySize(size int, cap int, safe ...bool) *Array {
	return &Array{
		mu:    rwmutex.Create(safe...),
		array: make([]interface{}, size, cap),
	}
}

// NewArrayRange 根据指定的范围从 `start` 到 `end` 并以步长值 `step` 创建并返回一个数组。
func NewArrayRange(start, end, step int, safe ...bool) *Array {
	if step == 0 {
		panic(fmt.Sprintf(`invalid step value: %d`, step))
	}
	slice := make([]interface{}, 0)
	index := 0
	for i := start; i <= end; i += step {
		slice = append(slice, i)
		index++
	}
	return NewArrayFrom(slice, safe...)
}

// NewFrom 是 NewArrayFrom 的别名。
// 请参阅 NewArrayFrom。
func NewFrom(array []interface{}, safe ...bool) *Array {
	return NewArrayFrom(array, safe...)
}

// NewFromCopy 是 NewArrayFromCopy 的别名。
// 请参阅 NewArrayFromCopy。
func NewFromCopy(array []interface{}, safe ...bool) *Array {
	return NewArrayFromCopy(array, safe...)
}

// NewArrayFrom 通过给定的切片 `array` 创建并返回一个数组。
// 参数 `safe` 用于指定是否在并发安全的情况下使用数组，默认为 false。
func NewArrayFrom(array []interface{}, safe ...bool) *Array {
	return &Array{
		mu:    rwmutex.Create(safe...),
		array: array,
	}
}

// NewArrayFromCopy 通过复制给定切片 `array` 创建并返回一个数组。
// 参数 `safe` 用于指定是否在并发安全的情况下使用数组，默认为 false。
func NewArrayFromCopy(array []interface{}, safe ...bool) *Array {
	newArray := make([]interface{}, len(array))
	copy(newArray, array)
	return &Array{
		mu:    rwmutex.Create(safe...),
		array: newArray,
	}
}

// At通过指定的索引返回值。
// 如果给定的`index`超出数组范围，它将返回`nil`。
func (a *Array) At(index int) (value interface{}) {
	value, _ = a.Get(index)
	return
}

// Get 通过指定的索引返回值。
// 如果给定的 `index` 超出了数组的范围，那么 `found` 将为 false。
func (a *Array) Get(index int) (value interface{}, found bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if index < 0 || index >= len(a.array) {
		return nil, false
	}
	return a.array[index], true
}

// Set将值设置到指定的索引位置。
func (a *Array) Set(index int, value interface{}) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if index < 0 || index >= len(a.array) {
		return gerror.NewCodef(gcode.CodeInvalidParameter, "index %d out of array range %d", index, len(a.array))
	}
	a.array[index] = value
	return nil
}

// SetArray 将底层的切片数组设置为给定的 `array`。
func (a *Array) SetArray(array []interface{}) *Array {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.array = array
	return a
}

// Replace 从数组起始位置开始，使用给定的`array`替换原有数组元素。
func (a *Array) Replace(array []interface{}) *Array {
	a.mu.Lock()
	defer a.mu.Unlock()
	max := len(array)
	if max > len(a.array) {
		max = len(a.array)
	}
	for i := 0; i < max; i++ {
		a.array[i] = array[i]
	}
	return a
}

// Sum 返回数组中所有值的和。
func (a *Array) Sum() (sum int) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		sum += gconv.Int(v)
	}
	return
}

// SortFunc 通过自定义函数 `less` 对数组进行排序。
func (a *Array) SortFunc(less func(v1, v2 interface{}) bool) *Array {
	a.mu.Lock()
	defer a.mu.Unlock()
	sort.Slice(a.array, func(i, j int) bool {
		return less(a.array[i], a.array[j])
	})
	return a
}

// InsertBefore 将 `values` 插入到 `index` 之前的位置。
func (a *Array) InsertBefore(index int, values ...interface{}) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if index < 0 || index >= len(a.array) {
		return gerror.NewCodef(gcode.CodeInvalidParameter, "index %d out of array range %d", index, len(a.array))
	}
	rear := append([]interface{}{}, a.array[index:]...)
	a.array = append(a.array[0:index], values...)
	a.array = append(a.array, rear...)
	return nil
}

// InsertAfter 在`index`之后插入`values`。
func (a *Array) InsertAfter(index int, values ...interface{}) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if index < 0 || index >= len(a.array) {
		return gerror.NewCodef(gcode.CodeInvalidParameter, "index %d out of array range %d", index, len(a.array))
	}
	rear := append([]interface{}{}, a.array[index+1:]...)
	a.array = append(a.array[0:index+1], values...)
	a.array = append(a.array, rear...)
	return nil
}

// Remove 通过索引移除一个元素。
// 如果给定的 `index` 超出了数组范围，`found` 将为 false。
func (a *Array) Remove(index int) (value interface{}, found bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.doRemoveWithoutLock(index)
}

// doRemoveWithoutLock 在没有加锁的情况下通过索引移除一个项。
func (a *Array) doRemoveWithoutLock(index int) (value interface{}, found bool) {
	if index < 0 || index >= len(a.array) {
		return nil, false
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
func (a *Array) RemoveValue(value interface{}) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if i := a.doSearchWithoutLock(value); i != -1 {
		a.doRemoveWithoutLock(i)
		return true
	}
	return false
}

// RemoveValues 通过 `values` 移除多个项目。
func (a *Array) RemoveValues(values ...interface{}) {
	a.mu.Lock()
	defer a.mu.Unlock()
	for _, value := range values {
		if i := a.doSearchWithoutLock(value); i != -1 {
			a.doRemoveWithoutLock(i)
		}
	}
}

// PushLeft 将一个或多个元素推送到数组的起始位置。
func (a *Array) PushLeft(value ...interface{}) *Array {
	a.mu.Lock()
	a.array = append(value, a.array...)
	a.mu.Unlock()
	return a
}

// PushRight将一个或多个元素推送到数组的末尾。
// 它等同于Append。
func (a *Array) PushRight(value ...interface{}) *Array {
	a.mu.Lock()
	a.array = append(a.array, value...)
	a.mu.Unlock()
	return a
}

// PopRand 随机地从数组中弹出并返回一个元素。
// 注意，如果数组为空，则 `found` 为 false。
func (a *Array) PopRand() (value interface{}, found bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.doRemoveWithoutLock(grand.Intn(len(a.array)))
}

// PopRands 随机地从数组中弹出并返回 `size` 个元素。
func (a *Array) PopRands(size int) []interface{} {
	a.mu.Lock()
	defer a.mu.Unlock()
	if size <= 0 || len(a.array) == 0 {
		return nil
	}
	if size >= len(a.array) {
		size = len(a.array)
	}
	array := make([]interface{}, size)
	for i := 0; i < size; i++ {
		array[i], _ = a.doRemoveWithoutLock(grand.Intn(len(a.array)))
	}
	return array
}

// PopLeft 从数组开头弹出并返回一个元素。
// 注意，如果数组为空，则 `found` 为 false。
func (a *Array) PopLeft() (value interface{}, found bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if len(a.array) == 0 {
		return nil, false
	}
	value = a.array[0]
	a.array = a.array[1:]
	return value, true
}

// PopRight从数组的末尾弹出并返回一个元素。
// 注意，如果数组为空，则`found`为false。
func (a *Array) PopRight() (value interface{}, found bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	index := len(a.array) - 1
	if index < 0 {
		return nil, false
	}
	value = a.array[index]
	a.array = a.array[:index]
	return value, true
}

// PopLefts 从数组开头弹出并返回 `size` 个元素。
func (a *Array) PopLefts(size int) []interface{} {
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
func (a *Array) PopRights(size int) []interface{} {
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
func (a *Array) Range(start int, end ...int) []interface{} {
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
	array := ([]interface{})(nil)
	if a.mu.IsSafe() {
		array = make([]interface{}, offsetEnd-start)
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
func (a *Array) SubSlice(offset int, length ...int) []interface{} {
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
		s := make([]interface{}, size)
		copy(s, a.array[offset:])
		return s
	} else {
		return a.array[offset:end]
	}
}

// Append 是 PushRight 的别名，请参阅 PushRight。
func (a *Array) Append(value ...interface{}) *Array {
	a.PushRight(value...)
	return a
}

// Len 返回数组的长度。
func (a *Array) Len() int {
	a.mu.RLock()
	length := len(a.array)
	a.mu.RUnlock()
	return length
}

// Slice 返回数组的基础数据。
// 注意，如果它在并发安全的使用场景下，会返回基础数据的一个副本，
// 否则，则返回指向基础数据的指针。
func (a *Array) Slice() []interface{} {
	if a.mu.IsSafe() {
		a.mu.RLock()
		defer a.mu.RUnlock()
		array := make([]interface{}, len(a.array))
		copy(array, a.array)
		return array
	} else {
		return a.array
	}
}

// Interfaces 函数将当前数组转换为 []interface{} 类型并返回。
func (a *Array) Interfaces() []interface{} {
	return a.Slice()
}

// Clone 返回一个新的数组，它是当前数组的一个副本。
func (a *Array) Clone() (newArray *Array) {
	a.mu.RLock()
	array := make([]interface{}, len(a.array))
	copy(array, a.array)
	a.mu.RUnlock()
	return NewArrayFrom(array, a.mu.IsSafe())
}

// 清空删除当前数组中的所有元素。
func (a *Array) Clear() *Array {
	a.mu.Lock()
	if len(a.array) > 0 {
		a.array = make([]interface{}, 0)
	}
	a.mu.Unlock()
	return a
}

// Contains 检查某个值是否存在于数组中。
func (a *Array) Contains(value interface{}) bool {
	return a.Search(value) != -1
}

// Search 在数组中通过 `value` 进行搜索，返回 `value` 的索引，
// 若不存在，则返回 -1。
func (a *Array) Search(value interface{}) int {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.doSearchWithoutLock(value)
}

func (a *Array) doSearchWithoutLock(value interface{}) int {
	if len(a.array) == 0 {
		return -1
	}
	result := -1
	for index, v := range a.array {
		if v == value {
			result = index
			break
		}
	}
	return result
}

// Unique 函数用于对数组去重，清除重复的元素。
// 示例：[1,1,2,3,2] -> [1,2,3]
func (a *Array) Unique() *Array {
	a.mu.Lock()
	defer a.mu.Unlock()
	if len(a.array) == 0 {
		return a
	}
	var (
		ok          bool
		temp        interface{}
		uniqueSet   = make(map[interface{}]struct{})
		uniqueArray = make([]interface{}, 0, len(a.array))
	)
	for i := 0; i < len(a.array); i++ {
		temp = a.array[i]
		if _, ok = uniqueSet[temp]; ok {
			continue
		}
		uniqueSet[temp] = struct{}{}
		uniqueArray = append(uniqueArray, temp)
	}
	a.array = uniqueArray
	return a
}

// LockFunc 通过回调函数`f`进行写入锁定。
func (a *Array) LockFunc(f func(array []interface{})) *Array {
	a.mu.Lock()
	defer a.mu.Unlock()
	f(a.array)
	return a
}

// RLockFunc 通过回调函数`f`锁定读取操作。
func (a *Array) RLockFunc(f func(array []interface{})) *Array {
	a.mu.RLock()
	defer a.mu.RUnlock()
	f(a.array)
	return a
}

// Merge 将`array`合并到当前数组中。
// 参数`array`可以是任何garray类型或切片类型。
// Merge 和 Append 的区别在于，Append 仅支持特定类型的切片作为参数，
// 而 Merge 支持更多类型的参数。
func (a *Array) Merge(array interface{}) *Array {
	return a.Append(gconv.Interfaces(array)...)
}

// Fill 用 `value` 值填充数组，填充 num 个条目，
// 键（索引）从 `startIndex` 参数开始。
func (a *Array) Fill(startIndex int, num int, value interface{}) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if startIndex < 0 || startIndex > len(a.array) {
		return gerror.NewCodef(gcode.CodeInvalidParameter, "index %d out of array range %d", startIndex, len(a.array))
	}
	for i := startIndex; i < startIndex+num; i++ {
		if i > len(a.array)-1 {
			a.array = append(a.array, value)
		} else {
			a.array[i] = value
		}
	}
	return nil
}

// Chunk 函数将一个数组分割成多个子数组，
// 每个子数组的大小由参数 `size` 确定。
// 最后一个子数组可能包含少于 size 个元素。
func (a *Array) Chunk(size int) [][]interface{} {
	if size < 1 {
		return nil
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	length := len(a.array)
	chunks := int(math.Ceil(float64(length) / float64(size)))
	var n [][]interface{}
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

// Pad 通过 `value` 值对数组进行填充，以达到指定长度。
// 如果 size 为正数，则在数组右侧进行填充；若为负数，则在左侧填充。
// 若 `size` 的绝对值小于或等于数组的长度，则不进行填充操作。
func (a *Array) Pad(size int, val interface{}) *Array {
	a.mu.Lock()
	defer a.mu.Unlock()
	if size == 0 || (size > 0 && size < len(a.array)) || (size < 0 && size > -len(a.array)) {
		return a
	}
	n := size
	if size < 0 {
		n = -size
	}
	n -= len(a.array)
	tmp := make([]interface{}, n)
	for i := 0; i < n; i++ {
		tmp[i] = val
	}
	if size > 0 {
		a.array = append(a.array, tmp...)
	} else {
		a.array = append(tmp, a.array...)
	}
	return a
}

// Rand 随机地从数组中返回一个元素（不删除）。
func (a *Array) Rand() (value interface{}, found bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if len(a.array) == 0 {
		return nil, false
	}
	return a.array[grand.Intn(len(a.array))], true
}

// Rands 随机返回数组中的 `size` 个元素（不删除）。
func (a *Array) Rands(size int) []interface{} {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if size <= 0 || len(a.array) == 0 {
		return nil
	}
	array := make([]interface{}, size)
	for i := 0; i < size; i++ {
		array[i] = a.array[grand.Intn(len(a.array))]
	}
	return array
}

// Shuffle 随机地对数组进行洗牌。
func (a *Array) Shuffle() *Array {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, v := range grand.Perm(len(a.array)) {
		a.array[i], a.array[v] = a.array[v], a.array[i]
	}
	return a
}

// Reverse 将数组元素按逆序排列生成新数组。
func (a *Array) Reverse() *Array {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, j := 0, len(a.array)-1; i < j; i, j = i+1, j-1 {
		a.array[i], a.array[j] = a.array[j], a.array[i]
	}
	return a
}

// Join 通过字符串 `glue` 连接数组元素。
func (a *Array) Join(glue string) string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if len(a.array) == 0 {
		return ""
	}
	buffer := bytes.NewBuffer(nil)
	for k, v := range a.array {
		buffer.WriteString(gconv.String(v))
		if k != len(a.array)-1 {
			buffer.WriteString(glue)
		}
	}
	return buffer.String()
}

// CountValues 计算数组中所有值出现的次数。
func (a *Array) CountValues() map[interface{}]int {
	m := make(map[interface{}]int)
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		m[v]++
	}
	return m
}

// Iterator 是 IteratorAsc 的别名。
func (a *Array) Iterator(f func(k int, v interface{}) bool) {
	a.IteratorAsc(f)
}

// IteratorAsc 以升序遍历给定数组，并使用回调函数 `f` 进行只读操作。
// 如果 `f` 返回 true，则继续迭代；若返回 false，则停止遍历。
func (a *Array) IteratorAsc(f func(k int, v interface{}) bool) {
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
func (a *Array) IteratorDesc(f func(k int, v interface{}) bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for i := len(a.array) - 1; i >= 0; i-- {
		if !f(i, a.array[i]) {
			break
		}
	}
}

// String 方法将当前数组以字符串形式返回，其实现方式类似于 json.Marshal。
func (a *Array) String() string {
	if a == nil {
		return ""
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	buffer := bytes.NewBuffer(nil)
	buffer.WriteByte('[')
	s := ""
	for k, v := range a.array {
		s = gconv.String(v)
		if gstr.IsNumeric(s) {
			buffer.WriteString(s)
		} else {
			buffer.WriteString(`"` + gstr.QuoteMeta(s, `"\`) + `"`)
		}
		if k != len(a.array)-1 {
			buffer.WriteByte(',')
		}
	}
	buffer.WriteByte(']')
	return buffer.String()
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
// 注意：此处接收者不使用指针。
func (a Array) MarshalJSON() ([]byte, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return json.Marshal(a.array)
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func (a *Array) UnmarshalJSON(b []byte) error {
	if a.array == nil {
		a.array = make([]interface{}, 0)
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	if err := json.UnmarshalUseNumber(b, &a.array); err != nil {
		return err
	}
	return nil
}

// UnmarshalValue 实现了一个接口，该接口用于为数组设置任意类型的值。
func (a *Array) UnmarshalValue(value interface{}) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	switch value.(type) {
	case string, []byte:
		return json.UnmarshalUseNumber(gconv.Bytes(value), &a.array)
	default:
		a.array = gconv.SliceAny(value)
	}
	return nil
}

// Filter 对数组进行迭代，并通过自定义回调函数进行元素过滤。
// 如果回调函数 `filter` 返回 true，则从数组中移除该元素；
// 否则不做任何处理并继续迭代。
func (a *Array) Filter(filter func(index int, value interface{}) bool) *Array {
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

// FilterNil 移除数组中所有的 nil 值。
func (a *Array) FilterNil() *Array {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i := 0; i < len(a.array); {
		if empty.IsNil(a.array[i]) {
			a.array = append(a.array[:i], a.array[i+1:]...)
		} else {
			i++
		}
	}
	return a
}

// FilterEmpty 用于移除数组中所有空值。
// 下列值被认为是空值：0, nil, false, "", 以及长度为0的slice、map或chan。
func (a *Array) FilterEmpty() *Array {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i := 0; i < len(a.array); {
		if empty.IsEmpty(a.array[i]) {
			a.array = append(a.array[:i], a.array[i+1:]...)
		} else {
			i++
		}
	}
	return a
}

// Walk 对数组中的每一项应用用户提供的函数 `f`。
func (a *Array) Walk(f func(value interface{}) interface{}) *Array {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, v := range a.array {
		a.array[i] = f(v)
	}
	return a
}

// IsEmpty 检查数组是否为空。
func (a *Array) IsEmpty() bool {
	return a.Len() == 0
}

// DeepCopy 实现接口，用于当前类型的深度复制。
func (a *Array) DeepCopy() interface{} {
	if a == nil {
		return nil
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	newSlice := make([]interface{}, len(a.array))
	for i, v := range a.array {
		newSlice[i] = deepcopy.Copy(v)
	}
	return NewArrayFrom(newSlice, a.mu.IsSafe())
}