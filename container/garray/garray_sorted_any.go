// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 数组类

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	
	"github.com/888go/goframe/internal/deepcopy"
	"github.com/888go/goframe/internal/empty"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/rwmutex"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/grand"
	"github.com/888go/goframe/util/gutil"
)

// SortedArray 是一个具有丰富特性的 Go 语言有序数组。
// 默认情况下，它使用递增顺序排列，可以通过设置自定义比较器进行更改。
// 它包含一个并发安全/不安全的切换选项，该选项应在初始化时设置，并且之后不能更改。
// 示例：
// ```go
// SortedArray 是一个Go语言实现的具备多种功能的有序数组。
// 默认情况下，数组按升序排列，通过设定自定义比较函数可以改变排序方式。
// 它具备并发安全/非安全模式切换功能，但该选项需在初始化时设定，并且设定后不可再更改。
type SortedArray struct {
	mu         rwmutex.RWMutex
	array      []interface{}
	unique     bool                       // 是否启用唯一特性(false)
	comparator func(a, b interface{}) int // 比较函数（返回值：-1表示a小于b；0表示a等于b；1表示a大于b）
}

// NewSortedArray 创建并返回一个空的排序数组。
// 参数 `safe` 用于指定是否在并发安全的情况下使用数组，默认为 false。
// 参数 `comparator` 用于比较数组中要排序的值，
// 如果它返回值 < 0，表示 `a` < `b`；此时 `a` 将被插入到 `b` 之前；
// 如果它返回值 = 0，表示 `a` = `b`；此时 `a` 将替换 `b`；
// 如果它返回值 > 0，表示 `a` > `b`；此时 `a` 将被插入到 `b` 之后。
func X创建排序(排序函数 func(a, b interface{}) int, 并发安全 ...bool) *SortedArray {
	return X创建排序并按大小(0, 排序函数, 并发安全...)
}

// NewSortedArraySize 根据给定的大小和容量创建并返回一个已排序的数组。
// 参数 `safe` 用于指定是否在并发安全的情况下使用数组，默认为 false。
func X创建排序并按大小(大小 int, 排序函数 func(a, b interface{}) int, 并发安全 ...bool) *SortedArray {
	return &SortedArray{
		mu:         rwmutex.Create(并发安全...),
		array:      make([]interface{}, 0, 大小),
		comparator: 排序函数,
	}
}

// NewSortedArrayRange 根据指定的范围从 `start` 到 `end`，并以步长值 `step` 创建并返回一个数组。
func X创建排序并按范围(起点, 终点, 步长 int, 排序函数 func(a, b interface{}) int, 并发安全 ...bool) *SortedArray {
	if 步长 == 0 {
		panic(fmt.Sprintf(`invalid step value: %d`, 步长))
	}
	slice := make([]interface{}, 0)
	index := 0
	for i := 起点; i <= 终点; i += 步长 {
		slice = append(slice, i)
		index++
	}
	return X创建排序并从数组(slice, 排序函数, 并发安全...)
}

// NewSortedArrayFrom 根据给定的切片 `array` 创建并返回一个已排序的数组。
// 参数 `safe` 用于指定是否在并发安全的情况下使用数组，默认为 false。
func X创建排序并从数组(数组 []interface{}, 排序函数 func(a, b interface{}) int, 并发安全 ...bool) *SortedArray {
	a := X创建排序并按大小(0, 排序函数, 并发安全...)
	a.array = 数组
	sort.Slice(a.array, func(i, j int) bool {
		return a.getComparator()(a.array[i], a.array[j]) < 0
	})
	return a
}

// NewSortedArrayFromCopy 函数通过复制给定切片 `array` 创建并返回一个已排序的数组。
// 参数 `safe` 用于指定是否在并发安全的情况下使用数组，默认为 false。
func X创建排序并从数组复制(数组 []interface{}, 排序函数 func(a, b interface{}) int, 并发安全 ...bool) *SortedArray {
	newArray := make([]interface{}, len(数组))
	copy(newArray, 数组)
	return X创建排序并从数组(newArray, 排序函数, 并发安全...)
}

// At通过指定的索引返回值。
// 如果给定的`index`超出数组范围，它将返回`nil`。
func (a *SortedArray) X取值(索引 int) (值 interface{}) {
	值, _ = a.X取值2(索引)
	return
}

// SetArray 将底层的切片数组设置为给定的 `array`。
func (a *SortedArray) X设置数组(数组 []interface{}) *SortedArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.array = 数组
	sort.Slice(a.array, func(i, j int) bool {
		return a.getComparator()(a.array[i], a.array[j]) < 0
	})
	return a
}

// SetComparator 设置/更改排序的比较器。
// 当比较器发生改变时，它会重新对数组进行排序。
func (a *SortedArray) X设置排序函数(排序函数 func(a, b interface{}) int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.comparator = 排序函数
	sort.Slice(a.array, func(i, j int) bool {
		return a.getComparator()(a.array[i], a.array[j]) < 0
	})
}

// Sort 函数用于将数组按升序排序。
// 参数 `reverse` 控制排序方式，若 reverse 为 true，则按降序（默认为升序）排序。
func (a *SortedArray) X排序递增() *SortedArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	sort.Slice(a.array, func(i, j int) bool {
		return a.getComparator()(a.array[i], a.array[j]) < 0
	})
	return a
}

// Add 向有序数组中添加一个或多个值，数组始终保持有序。
// 它是函数 Append 的别名，请参阅 Append。
func (a *SortedArray) X入栈右(值 ...interface{}) *SortedArray {
	return a.Append别名(值...)
}

// Append 向已排序的数组中添加一个或多个值，数组将始终保持有序。
func (a *SortedArray) Append别名(值 ...interface{}) *SortedArray {
	if len(值) == 0 {
		return a
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	for _, value := range 值 {
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

// Get 通过指定的索引返回值。
// 如果给定的 `index` 超出了数组的范围，那么 `found` 将为 false。
func (a *SortedArray) X取值2(索引 int) (值 interface{}, 成功 bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if 索引 < 0 || 索引 >= len(a.array) {
		return nil, false
	}
	return a.array[索引], true
}

// Remove 通过索引移除一个元素。
// 如果给定的 `index` 超出了数组范围，`found` 将为 false。
func (a *SortedArray) X删除(索引 int) (值 interface{}, 成功 bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.doRemoveWithoutLock(索引)
}

// doRemoveWithoutLock 在没有加锁的情况下通过索引移除一个项。
func (a *SortedArray) doRemoveWithoutLock(index int) (value interface{}, found bool) {
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
func (a *SortedArray) X删除值(值 interface{}) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if i, r := a.binSearch(值, false); r == 0 {
		_, res := a.doRemoveWithoutLock(i)
		return res
	}
	return false
}

// RemoveValues 通过 `values` 移除项目。
func (a *SortedArray) X删除多个值(值 ...interface{}) {
	a.mu.Lock()
	defer a.mu.Unlock()
	for _, value := range 值 {
		if i, r := a.binSearch(value, false); r == 0 {
			a.doRemoveWithoutLock(i)
		}
	}
}

// PopLeft 从数组开头弹出并返回一个元素。
// 注意，如果数组为空，则 `found` 为 false。
func (a *SortedArray) X出栈左() (值 interface{}, 成功 bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if len(a.array) == 0 {
		return nil, false
	}
	值 = a.array[0]
	a.array = a.array[1:]
	return 值, true
}

// PopRight从数组的末尾弹出并返回一个元素。
// 注意，如果数组为空，则`found`为false。
func (a *SortedArray) X出栈右() (值 interface{}, 成功 bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	index := len(a.array) - 1
	if index < 0 {
		return nil, false
	}
	值 = a.array[index]
	a.array = a.array[:index]
	return 值, true
}

// PopRand 随机地从数组中弹出并返回一个元素。
// 注意，如果数组为空，则 `found` 为 false。
func (a *SortedArray) X出栈随机() (值 interface{}, 成功 bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.doRemoveWithoutLock(随机类.X整数(len(a.array)))
}

// PopRands 随机地从数组中弹出并返回 `size` 个元素。
func (a *SortedArray) X出栈随机多个(数量 int) []interface{} {
	a.mu.Lock()
	defer a.mu.Unlock()
	if 数量 <= 0 || len(a.array) == 0 {
		return nil
	}
	if 数量 >= len(a.array) {
		数量 = len(a.array)
	}
	array := make([]interface{}, 数量)
	for i := 0; i < 数量; i++ {
		array[i], _ = a.doRemoveWithoutLock(随机类.X整数(len(a.array)))
	}
	return array
}

// PopLefts 从数组开头弹出并返回 `size` 个元素。
func (a *SortedArray) X出栈左多个(数量 int) []interface{} {
	a.mu.Lock()
	defer a.mu.Unlock()
	if 数量 <= 0 || len(a.array) == 0 {
		return nil
	}
	if 数量 >= len(a.array) {
		array := a.array
		a.array = a.array[:0]
		return array
	}
	value := a.array[0:数量]
	a.array = a.array[数量:]
	return value
}

// PopRights 从数组末尾弹出并返回 `size` 个元素。
func (a *SortedArray) X出栈右多个(数量 int) []interface{} {
	a.mu.Lock()
	defer a.mu.Unlock()
	if 数量 <= 0 || len(a.array) == 0 {
		return nil
	}
	index := len(a.array) - 数量
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
func (a *SortedArray) X取切片并按范围(起点 int, 终点 ...int) []interface{} {
	a.mu.RLock()
	defer a.mu.RUnlock()
	offsetEnd := len(a.array)
	if len(终点) > 0 && 终点[0] < offsetEnd {
		offsetEnd = 终点[0]
	}
	if 起点 > offsetEnd {
		return nil
	}
	if 起点 < 0 {
		起点 = 0
	}
	array := ([]interface{})(nil)
	if a.mu.IsSafe() {
		array = make([]interface{}, offsetEnd-起点)
		copy(array, a.array[起点:offsetEnd])
	} else {
		array = a.array[起点:offsetEnd]
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
func (a *SortedArray) X取切片并按数量(起点 int, 数量 ...int) []interface{} {
	a.mu.RLock()
	defer a.mu.RUnlock()
	size := len(a.array)
	if len(数量) > 0 {
		size = 数量[0]
	}
	if 起点 > len(a.array) {
		return nil
	}
	if 起点 < 0 {
		起点 = len(a.array) + 起点
		if 起点 < 0 {
			return nil
		}
	}
	if size < 0 {
		起点 += size
		size = -size
		if 起点 < 0 {
			return nil
		}
	}
	end := 起点 + size
	if end > len(a.array) {
		end = len(a.array)
		size = len(a.array) - 起点
	}
	if a.mu.IsSafe() {
		s := make([]interface{}, size)
		copy(s, a.array[起点:])
		return s
	} else {
		return a.array[起点:end]
	}
}

// Sum 返回数组中所有值的和。
func (a *SortedArray) X求和() (值 int) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		值 += 转换类.X取整数(v)
	}
	return
}

// Len 返回数组的长度。
func (a *SortedArray) X取长度() int {
	a.mu.RLock()
	length := len(a.array)
	a.mu.RUnlock()
	return length
}

// Slice 返回数组的基础数据。
// 注意，如果它在并发安全的使用场景下，会返回基础数据的一个副本，
// 否则，则返回指向基础数据的指针。
func (a *SortedArray) X取切片() []interface{} {
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

// Interfaces 函数将当前数组转换为 []interface{} 类型并返回。
func (a *SortedArray) X取any数组() []interface{} {
	return a.X取切片()
}

// Contains 检查某个值是否存在于数组中。
func (a *SortedArray) X是否存在(值 interface{}) bool {
	return a.X查找(值) != -1
}

// Search 在数组中通过 `value` 进行搜索，返回 `value` 的索引，
// 若不存在，则返回 -1。
func (a *SortedArray) X查找(值 interface{}) (索引 int) {
	if i, r := a.binSearch(值, true); r == 0 {
		return i
	}
	return -1
}

// 二分查找
// 返回最后比较的索引以及查找结果
// 若`result`等于0，表示在`index`位置的值等于`value`
// 若`result`小于0，表示在`index`位置的值小于`value`
// 若`result`大于0，表示在`index`位置的值大于`value`
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

// SetUnique 将唯一标志设置到数组中，
// 意味着该数组不包含任何重复的元素。
// 同时进行唯一性检查，移除所有重复的项。
func (a *SortedArray) X设置去重(去重 bool) *SortedArray {
	oldUnique := a.unique
	a.unique = 去重
	if 去重 && oldUnique != 去重 {
		a.X去重()
	}
	return a
}

// Unique 对数组进行去重，清除重复的元素。
func (a *SortedArray) X去重() *SortedArray {
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
func (a *SortedArray) X取副本() (新数组 *SortedArray) {
	a.mu.RLock()
	array := make([]interface{}, len(a.array))
	copy(array, a.array)
	a.mu.RUnlock()
	return X创建排序并从数组(array, a.comparator, a.mu.IsSafe())
}

// 清空删除当前数组中的所有元素。
func (a *SortedArray) X清空() *SortedArray {
	a.mu.Lock()
	if len(a.array) > 0 {
		a.array = make([]interface{}, 0)
	}
	a.mu.Unlock()
	return a
}

// LockFunc 通过回调函数`f`进行写入锁定。
func (a *SortedArray) X遍历写锁定(回调函数 func(数组 []interface{})) *SortedArray {
	a.mu.Lock()
	defer a.mu.Unlock()

	// 保持数组始终有序。
	defer sort.Slice(a.array, func(i, j int) bool {
		return a.getComparator()(a.array[i], a.array[j]) < 0
	})

	回调函数(a.array)
	return a
}

// RLockFunc 通过回调函数`f`锁定读取操作。
func (a *SortedArray) X遍历读锁定(回调函数 func(数组 []interface{})) *SortedArray {
	a.mu.RLock()
	defer a.mu.RUnlock()
	回调函数(a.array)
	return a
}

// Merge 将`array`合并到当前数组中。
// 参数`array`可以是任何garray类型或切片类型。
// Merge 和 Append 的区别在于，Append 仅支持特定类型的切片作为参数，
// 而 Merge 支持更多类型的参数。
func (a *SortedArray) X合并(数组 interface{}) *SortedArray {
	return a.X入栈右(转换类.X取any数组(数组)...)
}

// Chunk 函数将一个数组分割成多个子数组，
// 每个子数组的大小由参数 `size` 确定。
// 最后一个子数组可能包含少于 size 个元素。
func (a *SortedArray) X分割(数量 int) [][]interface{} {
	if 数量 < 1 {
		return nil
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	length := len(a.array)
	chunks := int(math.Ceil(float64(length) / float64(数量)))
	var n [][]interface{}
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * 数量
		if end > length {
			end = length
		}
		n = append(n, a.array[i*数量:end])
		i++
	}
	return n
}

// Rand 随机地从数组中返回一个元素（不删除）。
func (a *SortedArray) X取值随机() (值 interface{}, 成功 bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if len(a.array) == 0 {
		return nil, false
	}
	return a.array[随机类.X整数(len(a.array))], true
}

// Rands 随机返回数组中的 `size` 个元素（不删除）。
func (a *SortedArray) X取值随机多个(数量 int) []interface{} {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if 数量 <= 0 || len(a.array) == 0 {
		return nil
	}
	array := make([]interface{}, 数量)
	for i := 0; i < 数量; i++ {
		array[i] = a.array[随机类.X整数(len(a.array))]
	}
	return array
}

// Join 通过字符串 `glue` 连接数组元素。
func (a *SortedArray) X连接(连接符 string) string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if len(a.array) == 0 {
		return ""
	}
	buffer := bytes.NewBuffer(nil)
	for k, v := range a.array {
		buffer.WriteString(转换类.String(v))
		if k != len(a.array)-1 {
			buffer.WriteString(连接符)
		}
	}
	return buffer.String()
}

// CountValues 计算数组中所有值出现的次数。
func (a *SortedArray) X统计() map[interface{}]int {
	m := make(map[interface{}]int)
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		m[v]++
	}
	return m
}

// Iterator 是 IteratorAsc 的别名。
func (a *SortedArray) X遍历(回调函数 func(k int, v interface{}) bool) {
	a.X遍历升序(回调函数)
}

// IteratorAsc 以升序遍历给定数组，并使用回调函数 `f` 进行只读操作。
// 如果 `f` 返回 true，则继续迭代；若返回 false，则停止遍历。
func (a *SortedArray) X遍历升序(回调函数 func(k int, v interface{}) bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for k, v := range a.array {
		if !回调函数(k, v) {
			break
		}
	}
}

// IteratorDesc 函数以降序遍历给定的数组，并使用指定回调函数 `f` 进行只读操作。
// 若 `f` 返回 true，则继续迭代；若返回 false，则停止迭代。
func (a *SortedArray) X遍历降序(回调函数 func(k int, v interface{}) bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for i := len(a.array) - 1; i >= 0; i-- {
		if !回调函数(i, a.array[i]) {
			break
		}
	}
}

// String 方法将当前数组以字符串形式返回，其实现方式类似于 json.Marshal。
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
		s = 转换类.String(v)
		if 文本类.X是否为数字(s) {
			buffer.WriteString(s)
		} else {
			buffer.WriteString(`"` + 文本类.X转义并按字符(s, `"\`) + `"`)
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
func (a SortedArray) MarshalJSON() ([]byte, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return json.Marshal(a.array)
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
// 注意，该比较器默认设置为字符串比较器。
func (a *SortedArray) UnmarshalJSON(b []byte) error {
	if a.comparator == nil {
		a.array = make([]interface{}, 0)
		a.comparator = 工具类.X比较文本
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

// UnmarshalValue 实现了一个接口，该接口用于为数组设置任何类型的值。
// 注意，默认情况下比较器设置为字符串比较器。
func (a *SortedArray) UnmarshalValue(value interface{}) (err error) {
	if a.comparator == nil {
		a.comparator = 工具类.X比较文本
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	switch value.(type) {
	case string, []byte:
		err = json.UnmarshalUseNumber(转换类.X取字节集(value), &a.array)
	default:
		a.array = 转换类.SliceAny别名(value)
	}
	if a.comparator != nil && a.array != nil {
		sort.Slice(a.array, func(i, j int) bool {
			return a.comparator(a.array[i], a.array[j]) < 0
		})
	}
	return err
}

// FilterNil 移除数组中所有的 nil 值。
func (a *SortedArray) X删除所有nil() *SortedArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i := 0; i < len(a.array); {
		if empty.X是否为Nil(a.array[i]) {
			a.array = append(a.array[:i], a.array[i+1:]...)
		} else {
			break
		}
	}
	for i := len(a.array) - 1; i >= 0; {
		if empty.X是否为Nil(a.array[i]) {
			a.array = append(a.array[:i], a.array[i+1:]...)
		} else {
			break
		}
	}
	return a
}

// Filter 对数组进行迭代，并通过自定义回调函数进行元素过滤。
// 如果回调函数 `filter` 返回 true，则从数组中移除该元素；
// 否则不做任何处理并继续迭代。
func (a *SortedArray) X遍历删除(回调函数 func(索引 int, 值 interface{}) bool) *SortedArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i := 0; i < len(a.array); {
		if 回调函数(i, a.array[i]) {
			a.array = append(a.array[:i], a.array[i+1:]...)
		} else {
			i++
		}
	}
	return a
}

// FilterEmpty 用于移除数组中所有空值。
// 下列值被认为是空值：0, nil, false, "", 以及长度为0的slice、map或chan。
func (a *SortedArray) X删除所有空值() *SortedArray {
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

// Walk 对数组中的每一项应用用户提供的函数 `f`。
func (a *SortedArray) X遍历修改(回调函数 func(值 interface{}) interface{}) *SortedArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	// 保持数组始终有序。
	defer sort.Slice(a.array, func(i, j int) bool {
		return a.getComparator()(a.array[i], a.array[j]) < 0
	})
	for i, v := range a.array {
		a.array[i] = 回调函数(v)
	}
	return a
}

// IsEmpty 检查数组是否为空。
func (a *SortedArray) X是否为空() bool {
	return a.X取长度() == 0
}

// getComparator 返回之前设置的比较器，如果之前未设置，则会引发panic。
func (a *SortedArray) getComparator() func(a, b interface{}) int {
	if a.comparator == nil {
		panic("comparator is missing for sorted array")
	}
	return a.comparator
}

// DeepCopy 实现接口，用于当前类型的深度复制。
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
	return X创建排序并从数组(newSlice, a.comparator, a.mu.IsSafe())
}

func (a *SortedArray) X取文本() string {
return a.String()
}
