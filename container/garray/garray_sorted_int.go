// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 切片类

import (
	"bytes"
	"fmt"
	"math"
	"sort"

	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/rwmutex"
	gconv "github.com/888go/goframe/util/gconv"
	grand "github.com/888go/goframe/util/grand"
)

// SortedIntArray 是一个具有丰富功能的 Go 语言排序整数数组。
// 默认情况下，它使用升序排序，但可以通过设置自定义比较器来改变。
// 它包含一个并发安全/不安全的开关，应在初始化时设置，并且之后不能更改。
// md5:5106a295d33aeafe
type SortedIntArray struct {
	mu         rwmutex.RWMutex
	array      []int
	unique     bool               // 是否启用唯一功能（false）. md5:e1a1e6b26151e91d
	comparator func(a, b int) int // 比较函数（返回值：-1 表示 a < b；0 表示 a == b；1 表示 a > b）. md5:2be44acd57b55d6a
}

// X创建整数排序 创建并返回一个空的排序数组。
// 参数 `safe` 用于指定是否在并发安全模式下使用数组，默认为 false。
// md5:29a3eb47ff7212a7
func X创建整数排序(并发安全 ...bool) *SortedIntArray {
	return X创建整数排序并按大小(0, 并发安全...)
}

// X创建整数排序并带排序函数 创建并返回一个空的已排序整数数组，使用指定的比较器。
// 参数 `safe` 用于指定是否在并发安全模式下使用数组，默认为 false。
// md5:9049bb870ab7f29a
func X创建整数排序并带排序函数(排序函数 func(a, b int) int, 并发安全 ...bool) *SortedIntArray {
	array := X创建整数排序(并发安全...)
	array.comparator = 排序函数
	return array
}

// X创建整数排序并按大小 创建并返回一个给定大小和容量的有序整数数组。
// 参数 `safe` 用于指定是否使用并发安全的数组，默认为false。
// md5:8ae629092f076b83
func X创建整数排序并按大小(大小 int, 并发安全 ...bool) *SortedIntArray {
	return &SortedIntArray{
		mu:         rwmutex.Create(并发安全...),
		array:      make([]int, 0, 大小),
		comparator: defaultComparatorInt,
	}
}

// X创建整数排序并按范围 根据范围 `start` 到 `end` 以及步长 `step` 创建并返回一个排序后的整数数组。
// md5:7835c3ca423dbb6b
func X创建整数排序并按范围(起点, 终点, 步长 int, 并发安全 ...bool) *SortedIntArray {
	if 步长 == 0 {
		panic(fmt.Sprintf(`invalid step value: %d`, 步长))
	}
	slice := make([]int, 0)
	index := 0
	for i := 起点; i <= 终点; i += 步长 {
		slice = append(slice, i)
		index++
	}
	return X创建整数排序并从切片(slice, 并发安全...)
}

// X创建整数排序并从切片 创建并返回一个已排序的整数数组，使用给定的切片 `array`。
// 参数 `safe` 用于指定是否在并发安全模式下使用数组，默认为 false。
// md5:31b8ab8b3ab6bf88
func X创建整数排序并从切片(切片 []int, 并发安全 ...bool) *SortedIntArray {
	a := X创建整数排序并按大小(0, 并发安全...)
	a.array = 切片
	sort.Ints(a.array)
	return a
}

// X创建整数排序并从切片复制 从给定切片 `array` 的副本创建并返回一个排序的整数数组。
// 参数 `safe` 用于指定是否在并发安全模式下使用数组，默认为 false。
// md5:b520af8a99d30fcc
func X创建整数排序并从切片复制(切片 []int, 并发安全 ...bool) *SortedIntArray {
	newArray := make([]int, len(切片))
	copy(newArray, 切片)
	return X创建整数排序并从切片(newArray, 并发安全...)
}

// X取值 函数返回指定索引处的值。
// 如果给定的 `index` 超出了数组的范围，它将返回 `0`。
// md5:f1565bd13293ecb5
func (a *SortedIntArray) X取值(索引 int) (值 int) {
	值, _ = a.X取值2(索引)
	return
}

// X设置切片 使用给定的 `array` 设置底层切片数组。 md5:160b43a5c0ec752c
func (a *SortedIntArray) X设置切片(切片 []int) *SortedIntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.array = 切片
	quickSortInt(a.array, a.getComparator())
	return a
}

// X排序递增 按照递增顺序对数组进行排序。
// 参数 `reverse` 控制排序方式，如果为真，则按递减顺序排序（默认为递增排序）。
// md5:13939809cd029411
func (a *SortedIntArray) X排序递增() *SortedIntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	quickSortInt(a.array, a.getComparator())
	return a
}

// X入栈右 将一个或多个值添加到已排序的数组中，数组始终保持排序。它是Append函数的别名，请参阅Append。
// md5:34facedfc7e1b731
func (a *SortedIntArray) X入栈右(值 ...int) *SortedIntArray {
	return a.Append别名(值...)
}

// Append别名 向已排序的数组中添加一个或多个值，数组将始终保持排序状态。 md5:f839b377c2c77f6b
func (a *SortedIntArray) Append别名(值 ...int) *SortedIntArray {
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
		rear := append([]int{}, a.array[index:]...)
		a.array = append(a.array[0:index], value)
		a.array = append(a.array, rear...)
	}
	return a
}

// X取值2 函数通过指定的索引返回值。
// 如果给定的 `index` 超出了数组范围，`found` 将为 false。
// md5:ab300cfc0d6dd8ee
func (a *SortedIntArray) X取值2(索引 int) (值 int, 成功 bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if 索引 < 0 || 索引 >= len(a.array) {
		return 0, false
	}
	return a.array[索引], true
}

// X删除 函数通过索引移除一个元素。
// 如果给定的 `index` 超出了数组范围，`found` 将为 false。
// md5:feaf958654838c25
func (a *SortedIntArray) X删除(索引 int) (值 int, 成功 bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.doRemoveWithoutLock(索引)
}

// doRemoveWithoutLock 不使用锁移除一个项目。 md5:a6a1746903fd131c
func (a *SortedIntArray) doRemoveWithoutLock(index int) (value int, found bool) {
	if index < 0 || index >= len(a.array) {
		return 0, false
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

// X删除值 函数根据值删除一个元素。
// 如果值在数组中找到，它将返回 true，否则如果未找到则返回 false。
// md5:c49c7706ce703d00
func (a *SortedIntArray) X删除值(值 int) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if i, r := a.binSearch(值, false); r == 0 {
		_, res := a.doRemoveWithoutLock(i)
		return res
	}
	return false
}

// X删除多个值 通过 `values` 删除一个项目。 md5:05e01eb00e998269
func (a *SortedIntArray) X删除多个值(值 ...int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	for _, value := range 值 {
		if i, r := a.binSearch(value, false); r == 0 {
			a.doRemoveWithoutLock(i)
		}
	}
}

// X出栈左 从数组的开头弹出并返回一个项目。
// 注意，如果数组为空，`found` 为 false。
// md5:68f14002d84594a4
func (a *SortedIntArray) X出栈左() (值 int, 成功 bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if len(a.array) == 0 {
		return 0, false
	}
	值 = a.array[0]
	a.array = a.array[1:]
	return 值, true
}

// X出栈右 从数组的末尾弹出并返回一个元素。
// 注意，如果数组为空，则 `found` 为 false。
// md5:207fa7c7c4a04a10
func (a *SortedIntArray) X出栈右() (值 int, 成功 bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	index := len(a.array) - 1
	if index < 0 {
		return 0, false
	}
	值 = a.array[index]
	a.array = a.array[:index]
	return 值, true
}

// X出栈随机 从数组中随机弹出并返回一个元素。
// 注意，如果数组为空，`found` 将为 false。
// md5:29338267db400401
func (a *SortedIntArray) X出栈随机() (值 int, 成功 bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.doRemoveWithoutLock(grand.X整数(len(a.array)))
}

// X出栈随机多个 随机地从数组中弹出并返回 `size` 个元素。
// 如果给定的 `size` 大于数组的大小，它将返回数组的所有元素。
// 注意，如果给定的 `size` 小于等于 0 或数组为空，它将返回 nil。
// md5:9fd270d3d3021d32
func (a *SortedIntArray) X出栈随机多个(数量 int) []int {
	a.mu.Lock()
	defer a.mu.Unlock()
	if 数量 <= 0 || len(a.array) == 0 {
		return nil
	}
	if 数量 >= len(a.array) {
		数量 = len(a.array)
	}
	array := make([]int, 数量)
	for i := 0; i < 数量; i++ {
		array[i], _ = a.doRemoveWithoutLock(grand.X整数(len(a.array)))
	}
	return array
}

// X出栈左多个 从数组开始处弹出并返回 `size` 个元素。
// 如果给定的 `size` 大于数组的长度，它将返回数组中的所有元素。
// 请注意，如果给定的 `size` 小于等于 0 或数组为空，它将返回 nil。
// md5:3ecbe066336a9849
func (a *SortedIntArray) X出栈左多个(数量 int) []int {
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

// X出栈右多个 从数组末尾弹出并返回 `size` 个元素。
// 如果给定的 `size` 大于数组的大小，它将返回数组中的所有元素。
// 注意，如果给定的 `size` 小于等于 0 或数组为空，它将返回 nil。
// md5:4f44f32fbb68fb50
func (a *SortedIntArray) X出栈右多个(数量 int) []int {
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

// X取切片并按范围通过范围选择并返回项目，就像数组[start:end]一样。
// 请注意，如果在并发安全使用中，它将返回切片的副本；否则返回底层数据的指针。
// 
// 如果`end`为负数，则偏移量将从数组末尾开始。
// 如果省略`end`，则序列将包含从`start`到数组结尾的所有内容。
// md5:8b71690536bb9ec5
func (a *SortedIntArray) X取切片并按范围(起点 int, 终点 ...int) []int {
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
	array := ([]int)(nil)
	if a.mu.IsSafe() {
		array = make([]int, offsetEnd-起点)
		copy(array, a.array[起点:offsetEnd])
	} else {
		array = a.array[起点:offsetEnd]
	}
	return array
}

// X取切片并按数量 返回数组中指定的一段元素切片。
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
func (a *SortedIntArray) X取切片并按数量(起点 int, 数量 ...int) []int {
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
		s := make([]int, size)
		copy(s, a.array[起点:])
		return s
	} else {
		return a.array[起点:end]
	}
}

// X取长度 返回数组的长度。 md5:593b37501e98da95
func (a *SortedIntArray) X取长度() int {
	a.mu.RLock()
	length := len(a.array)
	a.mu.RUnlock()
	return length
}

// X求和 返回数组中所有值的和。 md5:b2148175a749b162
func (a *SortedIntArray) X求和() (值 int) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		值 += v
	}
	return
}

// X取切片 返回数组的底层数据。
// 注意，如果在并发安全的使用情况下，它会返回底层数据的副本，否则返回底层数据的指针。
// md5:111cbee45795a58b
func (a *SortedIntArray) X取切片() []int {
	array := ([]int)(nil)
	if a.mu.IsSafe() {
		a.mu.RLock()
		defer a.mu.RUnlock()
		array = make([]int, len(a.array))
		copy(array, a.array)
	} else {
		array = a.array
	}
	return array
}

// X取any切片 将当前数组作为 []interface{} 返回。 md5:f7a2e3459e185314
func (a *SortedIntArray) X取any切片() []interface{} {
	a.mu.RLock()
	defer a.mu.RUnlock()
	array := make([]interface{}, len(a.array))
	for k, v := range a.array {
		array[k] = v
	}
	return array
}

// X是否存在 检查值是否存在于数组中。 md5:f209e1f30dd53cb2
func (a *SortedIntArray) X是否存在(值 int) bool {
	return a.X查找(值) != -1
}

// X查找 在数组中搜索 `value`，返回 `value` 的索引，
// 如果不存在则返回 -1。
// md5:787617bfeade8f93
func (a *SortedIntArray) X查找(值 int) (索引 int) {
	if i, r := a.binSearch(值, true); r == 0 {
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
func (a *SortedIntArray) binSearch(value int, lock bool) (index int, result int) {
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

// X设置去重 将唯一标记设置到数组中，
// 表示数组不包含任何重复的元素。
// 它还执行唯一性检查，移除所有重复的项。
// md5:52bf11e8153b2459
func (a *SortedIntArray) X设置去重(去重 bool) *SortedIntArray {
	oldUnique := a.unique
	a.unique = 去重
	if 去重 && oldUnique != 去重 {
		a.X去重()
	}
	return a
}

// X去重 函数用于清除非唯一元素，确保数组中的每个元素都是唯一的。 md5:6dfd767cdbb67ed2
func (a *SortedIntArray) X去重() *SortedIntArray {
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

// X取副本 返回一个新的数组，它是当前数组的副本。 md5:52ada4030c562295
func (a *SortedIntArray) X取副本() (新切片 *SortedIntArray) {
	a.mu.RLock()
	array := make([]int, len(a.array))
	copy(array, a.array)
	a.mu.RUnlock()
	return X创建整数排序并从切片(array, a.mu.IsSafe())
}

// X清空 删除当前数组中的所有项目。 md5:3d9c6d68a5719979
func (a *SortedIntArray) X清空() *SortedIntArray {
	a.mu.Lock()
	if len(a.array) > 0 {
		a.array = make([]int, 0)
	}
	a.mu.Unlock()
	return a
}

// X遍历写锁定 通过回调函数 `f` 实现写入锁定。 md5:d45a130fa9aa0af2
func (a *SortedIntArray) X遍历写锁定(回调函数 func(array []int)) *SortedIntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	回调函数(a.array)
	return a
}

// X遍历读锁定 通过回调函数 `f` 实现读取锁定。 md5:a45deee1e6f17c88
func (a *SortedIntArray) X遍历读锁定(回调函数 func(array []int)) *SortedIntArray {
	a.mu.RLock()
	defer a.mu.RUnlock()
	回调函数(a.array)
	return a
}

// X合并 将 `array` 合并到当前数组中。
// 参数 `array` 可以是任何 garray 或切片类型。
// X合并 和 Append 的区别在于，Append 只支持特定的切片类型，
// 而 X合并 支持更多种类的参数类型。
// md5:465caccda38e84f8
func (a *SortedIntArray) X合并(切片 interface{}) *SortedIntArray {
	return a.X入栈右(gconv.X取整数切片(切片)...)
}

// X分割 将一个数组分割成多个子数组，每个子数组的大小由 `size` 决定。最后一个子数组可能包含少于 `size` 个元素。
// md5:0f1f74ff34633d24
func (a *SortedIntArray) X分割(数量 int) [][]int {
	if 数量 < 1 {
		return nil
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	length := len(a.array)
	chunks := int(math.Ceil(float64(length) / float64(数量)))
	var n [][]int
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

// X取值随机 随机从数组中返回一个元素（不进行删除）。 md5:e152d2c5bc15ecd7
func (a *SortedIntArray) X取值随机() (值 int, 成功 bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if len(a.array) == 0 {
		return 0, false
	}
	return a.array[grand.X整数(len(a.array))], true
}

// X取值随机多个 随机从数组中返回 `size` 个元素（不删除）。 md5:09ad7802f8190e3c
func (a *SortedIntArray) X取值随机多个(数量 int) []int {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if 数量 <= 0 || len(a.array) == 0 {
		return nil
	}
	array := make([]int, 数量)
	for i := 0; i < 数量; i++ {
		array[i] = a.array[grand.X整数(len(a.array))]
	}
	return array
}

// X连接 使用字符串 `glue` 连接数组元素。 md5:ec3894b049af1251
func (a *SortedIntArray) X连接(连接符 string) string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if len(a.array) == 0 {
		return ""
	}
	buffer := bytes.NewBuffer(nil)
	for k, v := range a.array {
		buffer.WriteString(gconv.String(v))
		if k != len(a.array)-1 {
			buffer.WriteString(连接符)
		}
	}
	return buffer.String()
}

// X统计 计算数组中所有值出现的次数。 md5:95b4772dcb002365
func (a *SortedIntArray) X统计() map[int]int {
	m := make(map[int]int)
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		m[v]++
	}
	return m
}

// X遍历 是 IteratorAsc 的别名。 md5:1bfdea306db62845
func (a *SortedIntArray) X遍历(f func(k int, v int) bool) {
	a.X遍历升序(f)
}

// X遍历升序 遍历数组，按照给定的回调函数 `f` 以升序进行只读访问。如果 `f` 返回 true，则继续遍历；否则停止。
// md5:8a125e2dd8982d48
func (a *SortedIntArray) X遍历升序(回调函数 func(k int, v int) bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for k, v := range a.array {
		if !回调函数(k, v) {
			break
		}
	}
}

// X遍历降序 以降序遍历数组，并使用给定的回调函数`f`进行只读迭代。
// 如果`f`返回true，则继续遍历；如果返回false，则停止遍历。
// md5:ea0a3805bccce0f7
func (a *SortedIntArray) X遍历降序(回调函数 func(k int, v int) bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for i := len(a.array) - 1; i >= 0; i-- {
		if !回调函数(i, a.array[i]) {
			break
		}
	}
}

// String 将当前数组转换为字符串，其实现方式类似于 json.Marshal。 md5:feda8f29233cde8d
func (a *SortedIntArray) String() string {
	if a == nil {
		return ""
	}
	return "[" + a.X连接(",") + "]"
}

// MarshalJSON实现了json.Marshal接口的MarshalJSON方法。
// 注意，这里不要使用指针作为接收者。
// md5:b4f76062b07a5263
func (a SortedIntArray) MarshalJSON() ([]byte, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return json.Marshal(a.array)
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
func (a *SortedIntArray) UnmarshalJSON(b []byte) error {
	if a.comparator == nil {
		a.array = make([]int, 0)
		a.comparator = defaultComparatorInt
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	if err := json.UnmarshalUseNumber(b, &a.array); err != nil {
		return err
	}
	if a.array != nil {
		sort.Ints(a.array)
	}
	return nil
}

// UnmarshalValue 是一个接口实现，用于为数组设置任何类型的数据值。 md5:35211e747ab939ab
func (a *SortedIntArray) UnmarshalValue(value interface{}) (err error) {
	if a.comparator == nil {
		a.comparator = defaultComparatorInt
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	switch value.(type) {
	case string, []byte:
		err = json.UnmarshalUseNumber(gconv.X取字节集(value), &a.array)
	default:
		a.array = gconv.SliceIne别名(value)
	}
	if a.array != nil {
		sort.Ints(a.array)
	}
	return err
}

// X遍历删除 遍历数组，并使用自定义回调函数过滤元素。
// 如果回调函数`filter`返回true，它将从数组中移除该元素，否则不做任何操作并继续遍历。
// md5:d33873cfb9f1bb38
func (a *SortedIntArray) X遍历删除(回调函数 func(索引 int, 值 int) bool) *SortedIntArray {
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

// X删除所有空值 移除数组中的所有零值。 md5:9155868b39243912
func (a *SortedIntArray) X删除所有空值() *SortedIntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i := 0; i < len(a.array); {
		if a.array[i] == 0 {
			a.array = append(a.array[:i], a.array[i+1:]...)
		} else {
			break
		}
	}
	for i := len(a.array) - 1; i >= 0; {
		if a.array[i] == 0 {
			a.array = append(a.array[:i], a.array[i+1:]...)
		} else {
			break
		}
	}
	return a
}

// X遍历修改 将用户提供的函数 `f` 应用到数组的每个元素上。 md5:51e35ea7c2c6525c
func (a *SortedIntArray) X遍历修改(回调函数 func(value int) int) *SortedIntArray {
	a.mu.Lock()
	defer a.mu.Unlock()

		// 保持数组始终排序。 md5:b2ef189f10478e96
	defer quickSortInt(a.array, a.getComparator())

	for i, v := range a.array {
		a.array[i] = 回调函数(v)
	}
	return a
}

// X是否为空 检查数组是否为空。 md5:fb6684351506a02d
func (a *SortedIntArray) X是否为空() bool {
	return a.X取长度() == 0
}

// getComparator 如果之前已经设置过比较器，则返回它，
// 否则返回一个默认的比较器。
// md5:8f22547cd8cea6eb
func (a *SortedIntArray) getComparator() func(a, b int) int {
	if a.comparator == nil {
		return defaultComparatorInt
	}
	return a.comparator
}

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
func (a *SortedIntArray) DeepCopy() interface{} {
	if a == nil {
		return nil
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	newSlice := make([]int, len(a.array))
	copy(newSlice, a.array)
	return X创建整数排序并从切片(newSlice, a.mu.IsSafe())
}


// zj:
func (a *SortedIntArray) X取文本() string {
	return a.String()
}

//zj: