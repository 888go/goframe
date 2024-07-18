// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package garray

import (
	"bytes"
	"fmt"
	"math"
	"sort"

	"github.com/gogf/gf/v2/internal/deepcopy"
	"github.com/gogf/gf/v2/internal/empty"
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/internal/rwmutex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/gutil"
)

// SortedArray 是一个功能丰富的 Go 语言排序数组。
// 默认情况下，它使用递增顺序，但可以通过设置自定义比较器进行更改。
// 它包含一个并发安全/不安全的开关，该开关应在初始化时设置并且不能更改。
// md5:48308289d58755e8
type SortedArray struct {
	mu         rwmutex.RWMutex
	array      []interface{}
	unique     bool                       // 是否启用唯一功能（false）. md5:e1a1e6b26151e91d
	comparator func(a, b interface{}) int // 比较函数（返回值：-1 表示 a < b；0 表示 a == b；1 表示 a > b）. md5:2be44acd57b55d6a
}

// NewSortedArray 创建并返回一个空的排序数组。
// 参数 `safe` 用于指定是否在并发安全环境下使用数组，默认为 false。
// 参数 `comparator` 用于比较数组中值的排序，
// 若返回值小于 0，表示 `a` 小于 `b`；`a` 将插入到 `b` 之前；
// 若返回值等于 0，表示 `a` 等于 `b`；`a` 将被 `b` 替换；
// 若返回值大于 0，表示 `a` 大于 `b`；`a` 将插入到 `b` 之后。
// md5:72443a89d087c135
// ff:创建排序
// comparator:排序函数
// a:
// b:
// safe:并发安全
func NewSortedArray(comparator func(a, b interface{}) int, safe ...bool) *SortedArray {
	return NewSortedArraySize(0, comparator, safe...)
}

// NewSortedArraySize 创建并返回一个给定大小和容量的排序数组。
// 参数 `safe` 用于指定是否在并发安全模式下使用数组，默认为 false。
// md5:827c3309faba84ac
// ff:创建排序并按大小
// cap:大小
// comparator:排序函数
// a:
// b:
// safe:并发安全
func NewSortedArraySize(cap int, comparator func(a, b interface{}) int, safe ...bool) *SortedArray {
	return &SortedArray{
		mu:         rwmutex.Create(safe...),
		array:      make([]interface{}, 0, cap),
		comparator: comparator,
	}
}

// NewSortedArrayRange 创建并返回一个从 `start` 到 `end` 的范围，步长为 `step` 的数组。
// md5:93c103a8dc8cf9d7
// ff:创建排序并按范围
// start:起点
// end:终点
// step:步长
// comparator:排序函数
// a:
// b:
// safe:并发安全
func NewSortedArrayRange(start, end, step int, comparator func(a, b interface{}) int, safe ...bool) *SortedArray {
	if step == 0 {
		panic(fmt.Sprintf(`invalid step value: %d`, step))
	}
	slice := make([]interface{}, 0)
	index := 0
	for i := start; i <= end; i += step {
		slice = append(slice, i)
		index++
	}
	return NewSortedArrayFrom(slice, comparator, safe...)
}

// NewSortedArrayFrom 根据给定的切片 `array` 创建并返回一个排序数组。
// 参数 `safe` 用于指定是否使用并发安全的数组，默认为 false。
// md5:764ff7e74cab303e
// ff:创建排序并从切片
// array:切片
// comparator:排序函数
// a:
// b:
// safe:并发安全
func NewSortedArrayFrom(array []interface{}, comparator func(a, b interface{}) int, safe ...bool) *SortedArray {
	a := NewSortedArraySize(0, comparator, safe...)
	a.array = array
	sort.Slice(a.array, func(i, j int) bool {
		return a.getComparator()(a.array[i], a.array[j]) < 0
	})
	return a
}

// NewSortedArrayFromCopy 根据给定切片 `array` 的副本创建并返回一个已排序的数组。
// 参数 `safe` 用于指定是否在并发安全模式下使用数组，默认为 false。
// md5:ec79e11f360050f4
// ff:创建排序并从切片复制
// array:切片
// comparator:排序函数
// a:
// b:
// safe:并发安全
func NewSortedArrayFromCopy(array []interface{}, comparator func(a, b interface{}) int, safe ...bool) *SortedArray {
	newArray := make([]interface{}, len(array))
	copy(newArray, array)
	return NewSortedArrayFrom(newArray, comparator, safe...)
}

// At 通过指定的索引返回值。
// 如果给定的`index`超出了数组的范围，它将返回`nil`。
// md5:09a7e6585d2eba1a
// ff:取值
// a:
// index:索引
// value:值
func (a *SortedArray) At(index int) (value interface{}) {
	value, _ = a.Get(index)
	return
}

// SetArray 使用给定的 `array` 设置底层切片数组。 md5:160b43a5c0ec752c
// ff:设置切片
// a:
// array:切片
func (a *SortedArray) SetArray(array []interface{}) *SortedArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.array = array
	sort.Slice(a.array, func(i, j int) bool {
		return a.getComparator()(a.array[i], a.array[j]) < 0
	})
	return a
}

// SetComparator 设置/更改排序的比较器。
// 当比较器更改时，它会重新对数组进行排序。
// md5:1323d8fba2b97b75
// ff:设置排序函数
// a:
// comparator:排序函数
// a:
// b:
func (a *SortedArray) SetComparator(comparator func(a, b interface{}) int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.comparator = comparator
	sort.Slice(a.array, func(i, j int) bool {
		return a.getComparator()(a.array[i], a.array[j]) < 0
	})
}

// Sort 对数组进行升序排序。
// 参数 `reverse` 控制排序方式，如果为 true，则降序排列（默认为升序）。
// md5:35d4650a0f563ccf
// ff:排序递增
// a:
func (a *SortedArray) Sort() *SortedArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	sort.Slice(a.array, func(i, j int) bool {
		return a.getComparator()(a.array[i], a.array[j]) < 0
	})
	return a
}

// Add 将一个或多个值添加到已排序的数组中，数组始终保持排序。它是Append函数的别名，请参阅Append。
// md5:34facedfc7e1b731
// ff:入栈右
// a:
// values:值
func (a *SortedArray) Add(values ...interface{}) *SortedArray {
	return a.Append(values...)
}

// Append 向已排序的数组中添加一个或多个值，数组将始终保持排序状态。 md5:f839b377c2c77f6b
// ff:Append别名
// a:
// values:值
func (a *SortedArray) Append(values ...interface{}) *SortedArray {
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
		a.array = append(a.array[:index], append([]interface{}{value}, a.array[index:]...)...)
	}
	return a
}

// Get 函数通过指定的索引返回值。
// 如果给定的 `index` 超出了数组范围，`found` 将为 false。
// md5:ab300cfc0d6dd8ee
// ff:取值2
// a:
// index:索引
// value:值
// found:成功
func (a *SortedArray) Get(index int) (value interface{}, found bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if index < 0 || index >= len(a.array) {
		return nil, false
	}
	return a.array[index], true
}

// Remove 函数通过索引移除一个元素。
// 如果给定的 `index` 超出了数组范围，`found` 将为 false。
// md5:feaf958654838c25
// ff:删除
// a:
// index:索引
// value:值
// found:成功
func (a *SortedArray) Remove(index int) (value interface{}, found bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.doRemoveWithoutLock(index)
}

// doRemoveWithoutLock 不使用锁移除一个项目。 md5:a6a1746903fd131c
func (a *SortedArray) doRemoveWithoutLock(index int) (value interface{}, found bool) {
	if index < 0 || index >= len(a.array) {
		return nil, false
	}
	// 在删除时确定数组边界，以提高删除效率。 md5:bc969ee880edf699
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
// ff:删除值
// a:
// value:值
func (a *SortedArray) RemoveValue(value interface{}) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if i, r := a.binSearch(value, false); r == 0 {
		_, res := a.doRemoveWithoutLock(i)
		return res
	}
	return false
}

// RemoveValues 通过 `values` 删除一个项目。 md5:05e01eb00e998269
// ff:删除多个值
// a:
// values:值
func (a *SortedArray) RemoveValues(values ...interface{}) {
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
// ff:出栈左
// a:
// value:值
// found:成功
func (a *SortedArray) PopLeft() (value interface{}, found bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if len(a.array) == 0 {
		return nil, false
	}
	value = a.array[0]
	a.array = a.array[1:]
	return value, true
}

// PopRight 从数组的末尾弹出并返回一个元素。
// 注意，如果数组为空，则 `found` 为 false。
// md5:207fa7c7c4a04a10
// ff:出栈右
// a:
// value:值
// found:成功
func (a *SortedArray) PopRight() (value interface{}, found bool) {
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

// PopRand 从数组中随机弹出并返回一个元素。
// 注意，如果数组为空，`found` 将为 false。
// md5:29338267db400401
// ff:出栈随机
// a:
// value:值
// found:成功
func (a *SortedArray) PopRand() (value interface{}, found bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.doRemoveWithoutLock(grand.Intn(len(a.array)))
}

// PopRands 随机地从数组中弹出并返回 `size` 个元素。 md5:3e1b1cbd52abd4cf
// ff:出栈随机多个
// a:
// size:数量
func (a *SortedArray) PopRands(size int) []interface{} {
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

// PopLefts 从数组开头弹出并返回 `size` 个元素。 md5:4a903258f1fe1dd4
// ff:出栈左多个
// a:
// size:数量
func (a *SortedArray) PopLefts(size int) []interface{} {
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

// PopRights 从数组末尾移除并返回 `size` 个元素。 md5:0b04e6ad99e5349b
// ff:出栈右多个
// a:
// size:数量
func (a *SortedArray) PopRights(size int) []interface{} {
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
// ff:取切片并按范围
// a:
// start:起点
// end:终点
func (a *SortedArray) Range(start int, end ...int) []interface{} {
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
// ff:取切片并按数量
// a:
// offset:起点
// length:数量
func (a *SortedArray) SubSlice(offset int, length ...int) []interface{} {
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

// Sum 返回数组中所有值的和。 md5:b2148175a749b162
// ff:求和
// a:
// sum:值
func (a *SortedArray) Sum() (sum int) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		sum += gconv.Int(v)
	}
	return
}

// Len 返回数组的长度。 md5:593b37501e98da95
// ff:取长度
// a:
func (a *SortedArray) Len() int {
	a.mu.RLock()
	length := len(a.array)
	a.mu.RUnlock()
	return length
}

// Slice 返回数组的底层数据。
// 注意，如果在并发安全的使用情况下，它会返回底层数据的副本，否则返回底层数据的指针。
// md5:111cbee45795a58b
// ff:取切片
// a:
func (a *SortedArray) Slice() []interface{} {
	var array []interface{}
	if a.mu.IsSafe() {
		a.mu.RLock()
		defer a.mu.RUnlock()
		array = make([]interface{}, len(a.array))
		copy(array, a.array)
	} else {
		array = a.array
	}
	return array
}

// Interfaces 将当前数组作为 []interface{} 返回。 md5:f7a2e3459e185314
// yx:true
// ff:取any切片
// a:
func (a *SortedArray) Interfaces() []interface{} {
	return a.Slice()
}

// Contains 检查值是否存在于数组中。 md5:f209e1f30dd53cb2
// ff:是否存在
// a:
// value:值
func (a *SortedArray) Contains(value interface{}) bool {
	return a.Search(value) != -1
}

// Search 在数组中搜索 `value`，返回 `value` 的索引，
// 如果不存在则返回 -1。
// md5:787617bfeade8f93
// ff:查找
// a:
// value:值
// index:索引
func (a *SortedArray) Search(value interface{}) (index int) {
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
func (a *SortedArray) binSearch(value interface{}, lock bool) (index int, result int) {
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
		mid = min + (max-min)/2
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
// 这意味着它不包含任何重复的项目。
// 它还会进行唯一性检查，删除所有重复项。
// md5:acbac75bf944670c
// ff:设置去重
// a:
// unique:去重
func (a *SortedArray) SetUnique(unique bool) *SortedArray {
	oldUnique := a.unique
	a.unique = unique
	if unique && oldUnique != unique {
		a.Unique()
	}
	return a
}

// Unique 函数用于清除非唯一元素，确保数组中的每个元素都是唯一的。 md5:6dfd767cdbb67ed2
// ff:去重
// a:
func (a *SortedArray) Unique() *SortedArray {
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

// Clone 返回一个新的数组，它是当前数组的副本。 md5:52ada4030c562295
// ff:取副本
// a:
// newArray:新切片
func (a *SortedArray) Clone() (newArray *SortedArray) {
	a.mu.RLock()
	array := make([]interface{}, len(a.array))
	copy(array, a.array)
	a.mu.RUnlock()
	return NewSortedArrayFrom(array, a.comparator, a.mu.IsSafe())
}

// Clear 删除当前数组中的所有项目。 md5:3d9c6d68a5719979
// ff:清空
// a:
func (a *SortedArray) Clear() *SortedArray {
	a.mu.Lock()
	if len(a.array) > 0 {
		a.array = make([]interface{}, 0)
	}
	a.mu.Unlock()
	return a
}

// LockFunc 通过回调函数 `f` 实现写入锁定。 md5:d45a130fa9aa0af2
// ff:遍历写锁定
// a:
// f:回调函数
// array:
func (a *SortedArray) LockFunc(f func(array []interface{})) *SortedArray {
	a.mu.Lock()
	defer a.mu.Unlock()

	// 保持数组始终排序。 md5:b2ef189f10478e96
	defer sort.Slice(a.array, func(i, j int) bool {
		return a.getComparator()(a.array[i], a.array[j]) < 0
	})

	f(a.array)
	return a
}

// RLockFunc 通过回调函数 `f` 实现读取锁定。 md5:a45deee1e6f17c88
// ff:遍历读锁定
// a:
// f:回调函数
// array:
func (a *SortedArray) RLockFunc(f func(array []interface{})) *SortedArray {
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
// ff:合并
// a:
// array:切片
func (a *SortedArray) Merge(array interface{}) *SortedArray {
	return a.Add(gconv.Interfaces(array)...)
}

// Chunk 将一个数组分割成多个子数组，每个子数组的大小由 `size` 决定。最后一个子数组可能包含少于 `size` 个元素。
// md5:0f1f74ff34633d24
// ff:分割
// a:
// size:数量
func (a *SortedArray) Chunk(size int) [][]interface{} {
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

// Rand 随机从数组中返回一个元素（不进行删除）。 md5:e152d2c5bc15ecd7
// ff:取值随机
// a:
// value:值
// found:成功
func (a *SortedArray) Rand() (value interface{}, found bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if len(a.array) == 0 {
		return nil, false
	}
	return a.array[grand.Intn(len(a.array))], true
}

// Rands 随机从数组中返回 `size` 个元素（不删除）。 md5:09ad7802f8190e3c
// ff:取值随机多个
// a:
// size:数量
func (a *SortedArray) Rands(size int) []interface{} {
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

// Join 使用字符串 `glue` 连接数组元素。 md5:ec3894b049af1251
// ff:连接
// a:
// glue:连接符
func (a *SortedArray) Join(glue string) string {
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

// CountValues 计算数组中所有值出现的次数。 md5:95b4772dcb002365
// ff:统计
// a:
func (a *SortedArray) CountValues() map[interface{}]int {
	m := make(map[interface{}]int)
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		m[v]++
	}
	return m
}

// Iterator 是 IteratorAsc 的别名。 md5:1bfdea306db62845
// yx:true
// ff:X遍历
// a:
// f:
// k:
// v:
func (a *SortedArray) Iterator(f func(k int, v interface{}) bool) {
	a.IteratorAsc(f)
}

// IteratorAsc 遍历数组，按照给定的回调函数 `f` 以升序进行只读访问。如果 `f` 返回 true，则继续遍历；否则停止。
// md5:8a125e2dd8982d48
// ff:遍历升序
// a:
// f:回调函数
// k:
// v:
func (a *SortedArray) IteratorAsc(f func(k int, v interface{}) bool) {
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
// ff:遍历降序
// a:
// f:回调函数
// k:
// v:
func (a *SortedArray) IteratorDesc(f func(k int, v interface{}) bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for i := len(a.array) - 1; i >= 0; i-- {
		if !f(i, a.array[i]) {
			break
		}
	}
}

// String 将当前数组转换为字符串，其实现方式类似于 json.Marshal。 md5:feda8f29233cde8d
// ff:
// a:
func (a *SortedArray) String() string {
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

// MarshalJSON实现了json.Marshal接口的MarshalJSON方法。
// 注意，这里不要使用指针作为接收者。
// md5:b4f76062b07a5263
// ff:
// a:
func (a SortedArray) MarshalJSON() ([]byte, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return json.Marshal(a.array)
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。
// 注意，默认情况下，比较器被设置为字符串比较器。
// md5:8af2d4d6f742bb31
// ff:
// a:
// b:
func (a *SortedArray) UnmarshalJSON(b []byte) error {
	if a.comparator == nil {
		a.array = make([]interface{}, 0)
		a.comparator = gutil.ComparatorString
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	if err := json.UnmarshalUseNumber(b, &a.array); err != nil {
		return err
	}
	if a.comparator != nil && a.array != nil {
		sort.Slice(a.array, func(i, j int) bool {
			return a.comparator(a.array[i], a.array[j]) < 0
		})
	}
	return nil
}

// UnmarshalValue 是一个接口实现，用于为数组设置任何类型的价值。
// 注意，比较器默认设置为字符串比较器。
// md5:5c9d5d1af1e97ec8
// ff:
// a:
// value:
// err:
func (a *SortedArray) UnmarshalValue(value interface{}) (err error) {
	if a.comparator == nil {
		a.comparator = gutil.ComparatorString
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	switch value.(type) {
	case string, []byte:
		err = json.UnmarshalUseNumber(gconv.Bytes(value), &a.array)
	default:
		a.array = gconv.SliceAny(value)
	}
	if a.comparator != nil && a.array != nil {
		sort.Slice(a.array, func(i, j int) bool {
			return a.comparator(a.array[i], a.array[j]) < 0
		})
	}
	return err
}

// FilterNil 删除数组中的所有空值（nil）。 md5:df6d66c2056b4815
// ff:删除所有nil
// a:
func (a *SortedArray) FilterNil() *SortedArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i := 0; i < len(a.array); {
		if empty.IsNil(a.array[i]) {
			a.array = append(a.array[:i], a.array[i+1:]...)
		} else {
			break
		}
	}
	for i := len(a.array) - 1; i >= 0; {
		if empty.IsNil(a.array[i]) {
			a.array = append(a.array[:i], a.array[i+1:]...)
		} else {
			break
		}
	}
	return a
}

// Filter 遍历数组，并使用自定义回调函数过滤元素。
// 如果回调函数`filter`返回true，它将从数组中移除该元素，否则不做任何操作并继续遍历。
// md5:d33873cfb9f1bb38
// ff:遍历删除
// a:
// filter:回调函数
// index:索引
// value:值
func (a *SortedArray) Filter(filter func(index int, value interface{}) bool) *SortedArray {
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

// FilterEmpty 移除数组中的所有空值。
// 被认为是空的值包括：0，nil，false，""，切片、映射（map）或通道（channel）的长度为0。
// md5:da01f627cd0962db
// ff:删除所有空值
// a:
func (a *SortedArray) FilterEmpty() *SortedArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i := 0; i < len(a.array); {
		if empty.IsEmpty(a.array[i]) {
			a.array = append(a.array[:i], a.array[i+1:]...)
		} else {
			break
		}
	}
	for i := len(a.array) - 1; i >= 0; {
		if empty.IsEmpty(a.array[i]) {
			a.array = append(a.array[:i], a.array[i+1:]...)
		} else {
			break
		}
	}
	return a
}

// Walk 将用户提供的函数 `f` 应用到数组的每个元素上。 md5:51e35ea7c2c6525c
// ff:遍历修改
// a:
// f:回调函数
// value:
func (a *SortedArray) Walk(f func(value interface{}) interface{}) *SortedArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	// 保持数组始终排序。 md5:b2ef189f10478e96
	defer sort.Slice(a.array, func(i, j int) bool {
		return a.getComparator()(a.array[i], a.array[j]) < 0
	})
	for i, v := range a.array {
		a.array[i] = f(v)
	}
	return a
}

// IsEmpty 检查数组是否为空。 md5:fb6684351506a02d
// ff:是否为空
// a:
func (a *SortedArray) IsEmpty() bool {
	return a.Len() == 0
}

// getComparator 如果之前已设置比较器，则返回该比较器，否则将引发恐慌。
// md5:03eac9fd6d838369
func (a *SortedArray) getComparator() func(a, b interface{}) int {
	if a.comparator == nil {
		panic("comparator is missing for sorted array")
	}
	return a.comparator
}

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
// ff:
// a:
func (a *SortedArray) DeepCopy() interface{} {
	if a == nil {
		return nil
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	newSlice := make([]interface{}, len(a.array))
	for i, v := range a.array {
		newSlice[i] = deepcopy.Copy(v)
	}
	return NewSortedArrayFrom(newSlice, a.comparator, a.mu.IsSafe())
}
