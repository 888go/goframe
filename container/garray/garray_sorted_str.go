// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package garray

import (
	"bytes"
	"math"
	"sort"
	"strings"

	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/internal/rwmutex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

// SortedStrArray 是一个具有丰富功能的 Go 语言排序字符串数组。
// 默认情况下，它使用递增顺序，可以通过设置自定义比较器来改变。
// 它包含一个并发安全/非安全开关，应在初始化时设置并且不能更改。
// md5:5738650138eb500e
type SortedStrArray struct {
	mu         rwmutex.RWMutex
	array      []string
	unique     bool                  // 是否启用唯一功能（false）. md5:e1a1e6b26151e91d
	comparator func(a, b string) int // 比较函数（返回值：-1 表示 a < b；0 表示 a == b；1 表示 a > b）. md5:2be44acd57b55d6a
}

// NewSortedStrArray 创建并返回一个空的有序数组。
// 参数 `safe` 用于指定是否使用并发安全的数组，
// 默认情况下为 false。
// md5:99d40d03c8301a35
// 翻译提示:func 新建排序字符串数组(safe ...bool) *排序字符串数组 {}
func NewSortedStrArray(safe ...bool) *SortedStrArray {
	return NewSortedStrArraySize(0, safe...)
}

// NewSortedStrArrayComparator 创建并返回一个使用指定比较器的空排序数组。
// 参数 `safe` 用于指定是否在并发安全模式下使用数组，默认为 false。
// md5:94666fc48998c3d3
// 翻译提示:func 新建排序字符串数组比较器(comparator func(a, b string) int, 是否安全 ...bool) *排序字符串数组 {}
func NewSortedStrArrayComparator(comparator func(a, b string) int, safe ...bool) *SortedStrArray {
	array := NewSortedStrArray(safe...)
	array.comparator = comparator
	return array
}

// NewSortedStrArraySize 创建并返回一个已排序的数组，给定大小和容量。
// 参数 `safe` 用于指定是否使用并发安全的数组，默认为 false。
// md5:dd1a2b286a0cce79
// 翻译提示:func 新建排序字符串数组(capacity int, 是否线程安全 ...bool) *排序字符串数组 {}
func NewSortedStrArraySize(cap int, safe ...bool) *SortedStrArray {
	return &SortedStrArray{
		mu:         rwmutex.Create(safe...),
		array:      make([]string, 0, cap),
		comparator: defaultComparatorStr,
	}
}

// NewSortedStrArrayFrom 创建并返回一个已排序的数组，使用给定的切片 `array`。
// 参数 `safe` 用于指定是否在并发安全模式下使用数组，默认为 false。
// md5:992cfc5cad0e6f7a
// 翻译提示:func 新建排序字符串数组从(array 字符串切片, safe ...布尔) *排序字符串数组 {}
func NewSortedStrArrayFrom(array []string, safe ...bool) *SortedStrArray {
	a := NewSortedStrArraySize(0, safe...)
	a.array = array
	quickSortStr(a.array, a.getComparator())
	return a
}

// NewSortedStrArrayFromCopy 根据给定的切片`array`的副本创建并返回一个排序后的数组。
// 参数`safe`用于指定是否使用并发安全的数组，默认为false。
// md5:e8cbae9d3604f7fc
// 翻译提示:func 新建已排序字符串数组从复制(array []string, 安全性 ...bool) *已排序字符串数组 {
//     return &已排序字符串数组{array: array, safe: gfutil.BoolValue(safe[0], true)}
// }
func NewSortedStrArrayFromCopy(array []string, safe ...bool) *SortedStrArray {
	newArray := make([]string, len(array))
	copy(newArray, array)
	return NewSortedStrArrayFrom(newArray, safe...)
}

// SetArray 使用给定的 `array` 设置底层切片数组。. md5:160b43a5c0ec752c
// 翻译提示:func (a *排序Str数组) 设置数组(array []string) *排序Str数组 {}
func (a *SortedStrArray) SetArray(array []string) *SortedStrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.array = array
	quickSortStr(a.array, a.getComparator())
	return a
}

// At通过指定的索引返回值。
// 如果给定的`index`超出了数组的范围，它将返回一个空字符串。
// md5:2465f6b1e3ac2863
// 翻译提示:func (a *SortedStrArray) 获取(index int) (元素 string) {} 
func (a *SortedStrArray) At(index int) (value string) {
	value, _ = a.Get(index)
	return
}

// Sort 按照递增顺序对数组进行排序。
// 参数 `reverse` 控制排序方式，如果为真，则按递减顺序排序（默认为递增排序）。
// md5:13939809cd029411
// 翻译提示:func (a *排序Str数组) 排序() *排序Str数组 {
//     return a
// }
func (a *SortedStrArray) Sort() *SortedStrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	quickSortStr(a.array, a.getComparator())
	return a
}

// Add 将一个或多个值添加到已排序的数组中，数组始终保持排序。它是Append函数的别名，请参阅Append。
// md5:34facedfc7e1b731
// 翻译提示:func (a *排序字符串数组) 添加(值 ...string) *排序字符串数组 {}
func (a *SortedStrArray) Add(values ...string) *SortedStrArray {
	return a.Append(values...)
}

// Append 向已排序的数组中添加一个或多个值，数组将始终保持排序状态。. md5:f839b377c2c77f6b
// 翻译提示:func (a *排序Str数组) 添加(值 ...string) *排序Str数组 {}
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

// Get 函数通过指定的索引返回值。
// 如果给定的 `index` 超出了数组范围，`found` 将为 false。
// md5:ab300cfc0d6dd8ee
// 翻译提示:func (a *SortedStrArray) 获取(index int) (元素 string, 是否存在 bool) {}
func (a *SortedStrArray) Get(index int) (value string, found bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if index < 0 || index >= len(a.array) {
		return "", false
	}
	return a.array[index], true
}

// Remove 函数通过索引移除一个元素。
// 如果给定的 `index` 超出了数组范围，`found` 将为 false。
// md5:feaf958654838c25
// 翻译提示:func (a *SortedStrArray) 移除(index int) (值 string, 是否找到 bool) {}
func (a *SortedStrArray) Remove(index int) (value string, found bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.doRemoveWithoutLock(index)
}

// doRemoveWithoutLock 不使用锁移除一个项目。. md5:a6a1746903fd131c
func (a *SortedStrArray) doRemoveWithoutLock(index int) (value string, found bool) {
	if index < 0 || index >= len(a.array) {
		return "", false
	}
	// 在删除时确定数组边界，以提高删除效率。. md5:bc969ee880edf699
	if index == 0 {
		value := a.array[0]
		a.array = a.array[1:]
		return value, true
	} else if index == len(a.array)-1 {
		value := a.array[index]
		a.array = a.array[:index]
		return value, true
	}
// 如果是一个非边界删除，
// 它将涉及创建一个数组，
// 那么删除操作效率较低。
// md5:6a664196d66bc968
	value = a.array[index]
	a.array = append(a.array[:index], a.array[index+1:]...)
	return value, true
}

// RemoveValue 函数根据值删除一个元素。
// 如果值在数组中找到，它将返回 true，否则如果未找到则返回 false。
// md5:c49c7706ce703d00
// 翻译提示:func (a *SortedStrArray) 删除值(value 字符串) bool {}
func (a *SortedStrArray) RemoveValue(value string) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if i, r := a.binSearch(value, false); r == 0 {
		_, res := a.doRemoveWithoutLock(i)
		return res
	}
	return false
}

// RemoveValues 通过 `values` 删除一个项目。. md5:05e01eb00e998269
// 翻译提示:func (a *排序字符串数组) 移除值(values ...string) {}
func (a *SortedStrArray) RemoveValues(values ...string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	for _, value := range values {
		if i, r := a.binSearch(value, false); r == 0 {
			a.doRemoveWithoutLock(i)
		}
	}
}

// PopLeft 从数组的开头弹出并返回一个项目。
// 注意，如果数组为空，`found` 为 false。
// md5:68f14002d84594a4
// 翻译提示:func (a *排序后字符串数组) 删除左元素() (元素值 string, 是否找到 bool) {}
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

// PopRight 从数组的末尾弹出并返回一个元素。
// 注意，如果数组为空，则 `found` 为 false。
// md5:207fa7c7c4a04a10
// 翻译提示:func (a *排序Str数组) 从右弹出() (值 string, 是否找到 bool) {}
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

// PopRand 从数组中随机弹出并返回一个元素。
// 注意，如果数组为空，`found` 将为 false。
// md5:29338267db400401
// 翻译提示:func (a *排序Str数组) PopRandom() (元素 string, 是否找到 bool) {}
func (a *SortedStrArray) PopRand() (value string, found bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.doRemoveWithoutLock(grand.Intn(len(a.array)))
}

// PopRands 随机地从数组中弹出并返回 `size` 个元素。
// 如果给定的 `size` 大于数组的大小，它将返回数组的所有元素。
// 注意，如果给定的 `size` 小于等于 0 或数组为空，它将返回 nil。
// md5:9fd270d3d3021d32
// 翻译提示:func (a *排序后字符串数组) 随机弹出(size int) []string {}
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

// PopLefts 从数组开始处弹出并返回 `size` 个元素。
// 如果给定的 `size` 大于数组的长度，它将返回数组中的所有元素。
// 请注意，如果给定的 `size` 小于等于 0 或数组为空，它将返回 nil。
// md5:3ecbe066336a9849
// 翻译提示:func (a *排序Str数组) 弹出左部(size int) []string {}
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
// 如果给定的 `size` 大于数组的大小，它将返回数组中的所有元素。
// 注意，如果给定的 `size` 小于等于 0 或数组为空，它将返回 nil。
// md5:4f44f32fbb68fb50
// 翻译提示:func (a *排序字符串数组) 弹出右侧元素(size int) []string {}
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

// Range通过范围选择并返回项目，就像数组[start:end]一样。
// 请注意，如果在并发安全使用中，它将返回切片的副本；否则返回底层数据的指针。
//
// 如果`end`为负数，则偏移量将从数组末尾开始。
// 如果省略`end`，则序列将包含从`start`到数组结尾的所有内容。
// md5:8b71690536bb9ec5
// 翻译提示:func (a *排序字符串数组) 范围(start int, end ...int) []string {}
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

// SubSlice 返回数组中指定的一段元素切片。
// 如果在并发安全的使用场景下，它将返回切片的一个副本；否则返回切片的指针。
//
// 如果偏移量（offset）为非负数，序列将从数组的该位置开始。
// 如果偏移量为负数，序列将从数组末尾向前该距离的位置开始。
//
// 如果提供了长度（size）且为正数，那么序列将包含最多这么多元素。
// 如果数组比指定的长度短，则序列只包含可用的数组元素。
// 如果长度为负数，则序列将在距离数组末尾该数量的元素处停止。
// 如果省略长度参数，那么序列将从偏移量开始直到数组末尾的所有元素。
//
// 如果切片范围的起始位置超出数组左侧边界，操作将失败。
// md5:f87ecd35d1dd7ac8
// 翻译提示:func (a *排序字符串数组) 子切片(开始位置 int, 长度 ...int) []string {}
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

// Sum 返回数组中所有值的和。. md5:b2148175a749b162
// 翻译提示:func (a *SortedStrArray) 汇总() (总和 int) {}
func (a *SortedStrArray) Sum() (sum int) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		sum += gconv.Int(v)
	}
	return
}

// Len 返回数组的长度。. md5:593b37501e98da95
// 翻译提示:func (a *排序Str数组) 长度() int {}
func (a *SortedStrArray) Len() int {
	a.mu.RLock()
	length := len(a.array)
	a.mu.RUnlock()
	return length
}

// Slice 返回数组的底层数据。
// 注意，如果在并发安全的使用情况下，它会返回底层数据的副本，否则返回底层数据的指针。
// md5:111cbee45795a58b
// 翻译提示:func (a *排序后字符串数组) 获取切片() []string {}
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

// Interfaces 将当前数组作为 []interface{} 返回。. md5:f7a2e3459e185314
// 翻译提示:func (a *排序Str数组) 接口值() []interface{}
func (a *SortedStrArray) Interfaces() []interface{} {
	a.mu.RLock()
	defer a.mu.RUnlock()
	array := make([]interface{}, len(a.array))
	for k, v := range a.array {
		array[k] = v
	}
	return array
}

// Contains 检查值是否存在于数组中。. md5:f209e1f30dd53cb2
// 翻译提示:func (a *排序Str数组) 包含(value 字符串) bool {}
func (a *SortedStrArray) Contains(value string) bool {
	return a.Search(value) != -1
}

// ContainsI 检查数组中是否存在某个值（忽略大小写）。
// 注意，它内部会遍历整个数组以进行不区分大小写的比较。
// md5:faf76a65365aa0ac
// 翻译提示:func (a *SortedStrArray) 包含I(元素 string) bool {}
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

// Search 在数组中搜索 `value`，返回 `value` 的索引，
// 如果不存在则返回 -1。
// md5:787617bfeade8f93
// 翻译提示:func (a *排序字符串数组) 搜索(value 字符串) (索引 int) {}
func (a *SortedStrArray) Search(value string) (index int) {
	if i, r := a.binSearch(value, true); r == 0 {
		return i
	}
	return -1
}

// 二分查找。
// 它返回最后比较的索引和结果。
// 如果 `result` 等于 0，表示索引处的值等于 `value`。
// 如果 `result` 小于 0，表示索引处的值小于 `value`。
// 如果 `result` 大于 0，表示索引处的值大于 `value`。
// md5:869c6a1ccba79c7a
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

// SetUnique 将唯一标记设置到数组中，
// 表示数组不包含任何重复的元素。
// 它还执行唯一性检查，移除所有重复的项。
// md5:52bf11e8153b2459
// 翻译提示:func (a *排序Str数组) 设置唯一(unique bool) *排序Str数组 {}
func (a *SortedStrArray) SetUnique(unique bool) *SortedStrArray {
	oldUnique := a.unique
	a.unique = unique
	if unique && oldUnique != unique {
		a.Unique()
	}
	return a
}

// Unique 函数用于清除非唯一元素，确保数组中的每个元素都是唯一的。. md5:6dfd767cdbb67ed2
// 翻译提示:func (a *排序字符串数组) 去重() *排序字符串数组 {}
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

// Clone 返回一个新的数组，它是当前数组的副本。. md5:52ada4030c562295
// 翻译提示:func (a *排序Str数组) 克隆() (新数组 *排序Str数组) {}
func (a *SortedStrArray) Clone() (newArray *SortedStrArray) {
	a.mu.RLock()
	array := make([]string, len(a.array))
	copy(array, a.array)
	a.mu.RUnlock()
	return NewSortedStrArrayFrom(array, a.mu.IsSafe())
}

// Clear 删除当前数组中的所有项目。. md5:3d9c6d68a5719979
// 翻译提示:func (a *排序字符串数组) 清空() *排序字符串数组 {}
func (a *SortedStrArray) Clear() *SortedStrArray {
	a.mu.Lock()
	if len(a.array) > 0 {
		a.array = make([]string, 0)
	}
	a.mu.Unlock()
	return a
}

// LockFunc 通过回调函数 `f` 实现写入锁定。. md5:d45a130fa9aa0af2
// 翻译提示:func (a *排序Str数组) 加锁函数(f func(数组 []string)) *排序Str数组 {}
func (a *SortedStrArray) LockFunc(f func(array []string)) *SortedStrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	f(a.array)
	return a
}

// RLockFunc 通过回调函数 `f` 实现读取锁定。. md5:a45deee1e6f17c88
// 翻译提示:func (a *排序Str数组) 读锁函数(f func(数组 []string)) *排序Str数组 {}
func (a *SortedStrArray) RLockFunc(f func(array []string)) *SortedStrArray {
	a.mu.RLock()
	defer a.mu.RUnlock()
	f(a.array)
	return a
}

// Merge 将 `array` 合并到当前数组中。
// 参数 `array` 可以是任何 garray 或切片类型。
// Merge 和 Append 的区别在于，Append 只支持特定的切片类型，
// 而 Merge 支持更多种类的参数类型。
// md5:465caccda38e84f8
// 翻译提示:func (a *SortedStrArray) 合并数组(array interface{})
func (a *SortedStrArray) Merge(array interface{}) *SortedStrArray {
	return a.Add(gconv.Strings(array)...)
}

// Chunk 将一个数组分割成多个子数组，每个子数组的大小由 `size` 决定。最后一个子数组可能包含少于 `size` 个元素。
// md5:0f1f74ff34633d24
// 翻译提示:func (a *排序后字符串数组) 分块(size int) [][]string {}
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

// Rand 随机从数组中返回一个元素（不进行删除）。. md5:e152d2c5bc15ecd7
// 翻译提示:func (a *排序Str数组) 随机获取() (值 string, 是否找到 bool) {}
func (a *SortedStrArray) Rand() (value string, found bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if len(a.array) == 0 {
		return "", false
	}
	return a.array[grand.Intn(len(a.array))], true
}

// Rands 随机从数组中返回 `size` 个元素（不删除）。. md5:09ad7802f8190e3c
// 翻译提示:func (a *排序后字符串数组) 随机取样(size int) []string {}
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

// Join 使用字符串 `glue` 连接数组元素。. md5:ec3894b049af1251
// 翻译提示:func (a *排序字符串数组) Concatenate(分隔符 string) string {}
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

// CountValues 计算数组中所有值出现的次数。. md5:95b4772dcb002365
// 翻译提示:func (a *排序Str数组) 计算值出现次数() map[string]int {}
func (a *SortedStrArray) CountValues() map[string]int {
	m := make(map[string]int)
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		m[v]++
	}
	return m
}

// Iterator 是 IteratorAsc 的别名。. md5:1bfdea306db62845
// 翻译提示:func (a *排序Str数组) 迭代器(f func(key int, value string) bool) {}
func (a *SortedStrArray) Iterator(f func(k int, v string) bool) {
	a.IteratorAsc(f)
}

// IteratorAsc 遍历数组，按照给定的回调函数 `f` 以升序进行只读访问。如果 `f` 返回 true，则继续遍历；否则停止。
// md5:8a125e2dd8982d48
// 翻译提示:func (a *SortedStrArray) 顺序迭代器(f func(index int, value string) bool) {}
func (a *SortedStrArray) IteratorAsc(f func(k int, v string) bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for k, v := range a.array {
		if !f(k, v) {
			break
		}
	}
}

// IteratorDesc 以降序遍历数组，并使用给定的回调函数`f`进行只读迭代。
// 如果`f`返回true，则继续遍历；如果返回false，则停止遍历。
// md5:ea0a3805bccce0f7
// 翻译提示:func (a *SortedStrArray) 降序迭代器(f func(index int, value string) bool) {}
func (a *SortedStrArray) IteratorDesc(f func(k int, v string) bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for i := len(a.array) - 1; i >= 0; i-- {
		if !f(i, a.array[i]) {
			break
		}
	}
}

// String 将当前数组转换为字符串，其实现方式类似于 json.Marshal。. md5:feda8f29233cde8d
// 翻译提示:func (a *排序后字符串数组) 字符串() string {}
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

// MarshalJSON实现了json.Marshal接口的MarshalJSON方法。
// 注意，这里不要使用指针作为接收者。
// md5:b4f76062b07a5263
// 翻译提示:func (a 排序字符串数组) MarshalJSON() ([]字节, 错误) {}
func (a SortedStrArray) MarshalJSON() ([]byte, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return json.Marshal(a.array)
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。. md5:f6766b88cf3d63c2
// 翻译提示:func (a *排序字符串数组) 解析JSON(b []byte) error {}
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

// UnmarshalValue 是一个接口实现，用于为数组设置任何类型的数据值。. md5:35211e747ab939ab
// 翻译提示:func (a *排序Str数组) 解码Value(value interface{})
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

// Filter 遍历数组，并使用自定义回调函数过滤元素。
// 如果回调函数`filter`返回true，它将从数组中移除该元素，否则不做任何操作并继续遍历。
// md5:d33873cfb9f1bb38
// 翻译提示:func (a *排序字符串数组) 过滤(filter func(index int, value string) bool) *排序字符串数组 {}
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

// FilterEmpty 函数移除数组中的所有空字符串值。. md5:2b2e8cd6c844936a
// 翻译提示:func (a *SortedStrArray) 过滤空字符串() *SortedStrArray {
//     return a
// }
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

// Walk 将用户提供的函数 `f` 应用到数组的每个元素上。. md5:51e35ea7c2c6525c
// 翻译提示:func (a *排序字符串数组) 遍历(f func(元素 string) string) *排序字符串数组 {}
func (a *SortedStrArray) Walk(f func(value string) string) *SortedStrArray {
	a.mu.Lock()
	defer a.mu.Unlock()

	// 保持数组始终排序。. md5:b2ef189f10478e96
	defer quickSortStr(a.array, a.getComparator())

	for i, v := range a.array {
		a.array[i] = f(v)
	}
	return a
}

// IsEmpty 检查数组是否为空。. md5:fb6684351506a02d
// 翻译提示:func (a *排序Str数组) 是否为空() bool {}
func (a *SortedStrArray) IsEmpty() bool {
	return a.Len() == 0
}

// getComparator 如果之前已经设置过比较器，则返回它，
// 否则返回一个默认的比较器。
// md5:8f22547cd8cea6eb
func (a *SortedStrArray) getComparator() func(a, b string) int {
	if a.comparator == nil {
		return defaultComparatorStr
	}
	return a.comparator
}

// DeepCopy实现当前类型的深拷贝接口。. md5:9cfbcb08109f6ce1
// 翻译提示:func (a *排序后字符串数组) 深度复制() interface{}
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
