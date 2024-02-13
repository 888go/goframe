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
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/rwmutex"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/grand"
)

// IntArray 是一个具有丰富功能的 Go 语言整数数组。
// 它包含一个并发安全/不安全切换开关，应在初始化时设置，并且此后不可更改。
type IntArray struct {
	mu    rwmutex.RWMutex
	array []int
}

// NewIntArray 创建并返回一个空数组。
// 参数 `safe` 用于指定是否在并发安全的情况下使用数组，默认为 false。
func X创建整数(并发安全 ...bool) *IntArray {
	return X创建整数并按大小(0, 0, 并发安全...)
}

// NewIntArraySize 根据给定的大小和容量创建并返回一个数组。
// 参数`safe`用于指定是否在并发安全的情况下使用数组，默认为false。
func X创建整数并按大小(大小 int, 上限 int, 并发安全 ...bool) *IntArray {
	return &IntArray{
		mu:    rwmutex.Create(并发安全...),
		array: make([]int, 大小, 上限),
	}
}

// NewIntArrayRange 根据指定范围从 `start` 到 `end` 以步长值 `step` 创建并返回一个整数数组。
func X创建整数并按范围(起点, 终点, 步长 int, 并发安全 ...bool) *IntArray {
	if 步长 == 0 {
		panic(fmt.Sprintf(`invalid step value: %d`, 步长))
	}
	slice := make([]int, 0)
	index := 0
	for i := 起点; i <= 终点; i += 步长 {
		slice = append(slice, i)
		index++
	}
	return X创建整数并从数组(slice, 并发安全...)
}

// NewIntArrayFrom 根据给定的切片 `array` 创建并返回一个数组。
// 参数 `safe` 用于指定是否在并发安全的情况下使用数组，默认为 false。
func X创建整数并从数组(数组 []int, 并发安全 ...bool) *IntArray {
	return &IntArray{
		mu:    rwmutex.Create(并发安全...),
		array: 数组,
	}
}

// NewIntArrayFromCopy 通过复制给定切片 `array` 创建并返回一个数组。
// 参数 `safe` 用于指定是否在并发安全的情况下使用数组，默认为 false。
func X创建整数并从数组复制(数组 []int, 并发安全 ...bool) *IntArray {
	newArray := make([]int, len(数组))
	copy(newArray, 数组)
	return &IntArray{
		mu:    rwmutex.Create(并发安全...),
		array: newArray,
	}
}

// At通过指定的索引返回值。
// 如果给定的`index`超出数组范围，它将返回`0`。
func (a *IntArray) X取值(索引 int) (值 int) {
	值, _ = a.X取值2(索引)
	return
}

// Get 通过指定的索引返回值。
// 如果给定的 `index` 超出了数组的范围，那么 `found` 将为 false。
func (a *IntArray) X取值2(索引 int) (值 int, 成功 bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if 索引 < 0 || 索引 >= len(a.array) {
		return 0, false
	}
	return a.array[索引], true
}

// Set将值设置到指定的索引位置。
func (a *IntArray) X设置值(索引 int, 值 int) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if 索引 < 0 || 索引 >= len(a.array) {
		return 错误类.X创建错误码并格式化(错误码类.CodeInvalidParameter, "index %d out of array range %d", 索引, len(a.array))
	}
	a.array[索引] = 值
	return nil
}

// SetArray 将底层的切片数组设置为给定的 `array`。
func (a *IntArray) X设置数组(数组 []int) *IntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.array = 数组
	return a
}

// Replace 从数组起始位置开始，使用给定的`array`替换原有数组元素。
func (a *IntArray) X替换(数组 []int) *IntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	max := len(数组)
	if max > len(a.array) {
		max = len(a.array)
	}
	for i := 0; i < max; i++ {
		a.array[i] = 数组[i]
	}
	return a
}

// Sum 返回数组中所有值的和。
func (a *IntArray) X求和() (值 int) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		值 += v
	}
	return
}

// Sort 函数用于将数组按升序排序。
// 参数 `reverse` 用于控制是否按升序（默认）或降序进行排序。
func (a *IntArray) X排序递增(降序 ...bool) *IntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	if len(降序) > 0 && 降序[0] {
		sort.Slice(a.array, func(i, j int) bool {
			return a.array[i] >= a.array[j]
		})
	} else {
		sort.Ints(a.array)
	}
	return a
}

// SortFunc 通过自定义函数 `less` 对数组进行排序。
func (a *IntArray) X排序函数(回调函数 func(v1, v2 int) bool) *IntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	sort.Slice(a.array, func(i, j int) bool {
		return 回调函数(a.array[i], a.array[j])
	})
	return a
}

// InsertBefore 将 `values` 插入到 `index` 之前的位置。
func (a *IntArray) X插入前面(索引 int, 值 ...int) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if 索引 < 0 || 索引 >= len(a.array) {
		return 错误类.X创建错误码并格式化(错误码类.CodeInvalidParameter, "index %d out of array range %d", 索引, len(a.array))
	}
	rear := append([]int{}, a.array[索引:]...)
	a.array = append(a.array[0:索引], 值...)
	a.array = append(a.array, rear...)
	return nil
}

// InsertAfter 在`index`之后插入`value`。
func (a *IntArray) X插入后面(index int, 值 ...int) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if index < 0 || index >= len(a.array) {
		return 错误类.X创建错误码并格式化(错误码类.CodeInvalidParameter, "index %d out of array range %d", index, len(a.array))
	}
	rear := append([]int{}, a.array[index+1:]...)
	a.array = append(a.array[0:index+1], 值...)
	a.array = append(a.array, rear...)
	return nil
}

// Remove 通过索引移除一个元素。
// 如果给定的 `index` 超出了数组范围，`found` 将为 false。
func (a *IntArray) X删除(索引 int) (值 int, 成功 bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.doRemoveWithoutLock(索引)
}

// doRemoveWithoutLock 在没有加锁的情况下通过索引移除一个项。
func (a *IntArray) doRemoveWithoutLock(index int) (value int, found bool) {
	if index < 0 || index >= len(a.array) {
		return 0, false
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
func (a *IntArray) X删除值(值 int) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if i := a.doSearchWithoutLock(值); i != -1 {
		a.doRemoveWithoutLock(i)
		return true
	}
	return false
}

// RemoveValues 通过 `values` 移除多个项目。
func (a *IntArray) X删除多个值(值 ...int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	for _, value := range 值 {
		if i := a.doSearchWithoutLock(value); i != -1 {
			a.doRemoveWithoutLock(i)
		}
	}
}

// PushLeft 将一个或多个元素推送到数组的起始位置。
func (a *IntArray) X入栈左(值 ...int) *IntArray {
	a.mu.Lock()
	a.array = append(值, a.array...)
	a.mu.Unlock()
	return a
}

// PushRight将一个或多个元素推送到数组的末尾。
// 它等同于Append。
func (a *IntArray) X入栈右(值 ...int) *IntArray {
	a.mu.Lock()
	a.array = append(a.array, 值...)
	a.mu.Unlock()
	return a
}

// PopLeft 从数组开头弹出并返回一个元素。
// 注意，如果数组为空，则 `found` 为 false。
func (a *IntArray) X出栈左() (值 int, 成功 bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if len(a.array) == 0 {
		return 0, false
	}
	值 = a.array[0]
	a.array = a.array[1:]
	return 值, true
}

// PopRight从数组的末尾弹出并返回一个元素。
// 注意，如果数组为空，则`found`为false。
func (a *IntArray) X出栈右() (值 int, 成功 bool) {
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

// PopRand 随机地从数组中弹出并返回一个元素。
// 注意，如果数组为空，则 `found` 为 false。
func (a *IntArray) X出栈随机() (值 int, 成功 bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.doRemoveWithoutLock(随机类.X整数(len(a.array)))
}

// PopRands 随机地从数组中弹出并返回 `size` 个元素。
// 若给定的 `size` 大于数组的大小，则返回数组中的所有元素。
// 注意，如果给定的 `size` 小于等于0或者数组为空，它将返回 nil。
func (a *IntArray) X出栈随机多个(数量 int) []int {
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
		array[i], _ = a.doRemoveWithoutLock(随机类.X整数(len(a.array)))
	}
	return array
}

// PopLefts 从数组开头弹出并返回 `size` 个元素。
// 如果给定的 `size` 大于数组的大小，则返回数组中的所有元素。
// 注意，如果给定的 `size` 小于等于0或者数组为空，则返回nil。
func (a *IntArray) X出栈左多个(数量 int) []int {
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
// 若给定的 `size` 大于数组的大小，则返回数组中所有元素。
// 注意，如果给定的 `size` 小于等于0或者数组为空，它将返回nil。
func (a *IntArray) X出栈右多个(数量 int) []int {
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
func (a *IntArray) X取切片并按范围(起点 int, 终点 ...int) []int {
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
func (a *IntArray) X取切片并按数量(起点 int, 数量 ...int) []int {
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

// Append 是 PushRight 的别名，请参阅 PushRight。
func (a *IntArray) Append别名(值 ...int) *IntArray {
	a.mu.Lock()
	a.array = append(a.array, 值...)
	a.mu.Unlock()
	return a
}

// Len 返回数组的长度。
func (a *IntArray) X取长度() int {
	a.mu.RLock()
	length := len(a.array)
	a.mu.RUnlock()
	return length
}

// Slice 返回数组的基础数据。
// 注意，如果它在并发安全的使用场景下，会返回基础数据的一个副本，
// 否则，则返回指向基础数据的指针。
func (a *IntArray) X取切片() []int {
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

// Interfaces 函数将当前数组转换为 []interface{} 类型并返回。
func (a *IntArray) X取any数组() []interface{} {
	a.mu.RLock()
	defer a.mu.RUnlock()
	array := make([]interface{}, len(a.array))
	for k, v := range a.array {
		array[k] = v
	}
	return array
}

// Clone 返回一个新的数组，它是当前数组的一个副本。
func (a *IntArray) X取副本() (新数组 *IntArray) {
	a.mu.RLock()
	array := make([]int, len(a.array))
	copy(array, a.array)
	a.mu.RUnlock()
	return X创建整数并从数组(array, a.mu.IsSafe())
}

// 清空删除当前数组中的所有元素。
func (a *IntArray) X清空() *IntArray {
	a.mu.Lock()
	if len(a.array) > 0 {
		a.array = make([]int, 0)
	}
	a.mu.Unlock()
	return a
}

// Contains 检查某个值是否存在于数组中。
func (a *IntArray) X是否存在(值 int) bool {
	return a.X查找(值) != -1
}

// Search 在数组中通过 `value` 进行搜索，返回 `value` 的索引，
// 若不存在，则返回 -1。
func (a *IntArray) X查找(值 int) int {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.doSearchWithoutLock(值)
}

func (a *IntArray) doSearchWithoutLock(value int) int {
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
func (a *IntArray) X去重() *IntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	if len(a.array) == 0 {
		return a
	}
	var (
		ok          bool
		temp        int
		uniqueSet   = make(map[int]struct{})
		uniqueArray = make([]int, 0, len(a.array))
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
func (a *IntArray) X遍历写锁定(回调函数 func(数组 []int)) *IntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	回调函数(a.array)
	return a
}

// RLockFunc 通过回调函数`f`锁定读取操作。
func (a *IntArray) X遍历读锁定(回调函数 func(数组 []int)) *IntArray {
	a.mu.RLock()
	defer a.mu.RUnlock()
	回调函数(a.array)
	return a
}

// Merge 将`array`合并到当前数组中。
// 参数`array`可以是任何garray类型或切片类型。
// Merge 和 Append 的区别在于，Append 仅支持特定类型的切片作为参数，
// 而 Merge 支持更多类型的参数。
func (a *IntArray) X合并(数组 interface{}) *IntArray {
	return a.Append别名(转换类.X取整数数组(数组)...)
}

// Fill 用 `value` 值填充数组，填充 num 个条目，
// 键（索引）从 `startIndex` 参数开始。
func (a *IntArray) X填充(起点 int, 填充数量 int, 值 int) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if 起点 < 0 || 起点 > len(a.array) {
		return 错误类.X创建错误码并格式化(错误码类.CodeInvalidParameter, "index %d out of array range %d", 起点, len(a.array))
	}
	for i := 起点; i < 起点+填充数量; i++ {
		if i > len(a.array)-1 {
			a.array = append(a.array, 值)
		} else {
			a.array[i] = 值
		}
	}
	return nil
}

// Chunk 函数将一个数组分割成多个子数组，
// 每个子数组的大小由参数 `size` 确定。
// 最后一个子数组可能包含少于 size 个元素。
func (a *IntArray) X分割(数量 int) [][]int {
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

// Pad 通过 `value` 值对数组进行填充，以达到指定长度。
// 如果 size 为正数，则在数组右侧进行填充；若为负数，则在左侧填充。
// 若 `size` 的绝对值小于或等于数组的长度，则不进行填充操作。
func (a *IntArray) X填满(总数量 int, 值 int) *IntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	if 总数量 == 0 || (总数量 > 0 && 总数量 < len(a.array)) || (总数量 < 0 && 总数量 > -len(a.array)) {
		return a
	}
	n := 总数量
	if 总数量 < 0 {
		n = -总数量
	}
	n -= len(a.array)
	tmp := make([]int, n)
	for i := 0; i < n; i++ {
		tmp[i] = 值
	}
	if 总数量 > 0 {
		a.array = append(a.array, tmp...)
	} else {
		a.array = append(tmp, a.array...)
	}
	return a
}

// Rand 随机地从数组中返回一个元素（不删除）。
func (a *IntArray) X取值随机() (值 int, 成功 bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if len(a.array) == 0 {
		return 0, false
	}
	return a.array[随机类.X整数(len(a.array))], true
}

// Rands 随机返回数组中的 `size` 个元素（不删除）。
func (a *IntArray) X取值随机多个(数量 int) []int {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if 数量 <= 0 || len(a.array) == 0 {
		return nil
	}
	array := make([]int, 数量)
	for i := 0; i < 数量; i++ {
		array[i] = a.array[随机类.X整数(len(a.array))]
	}
	return array
}

// Shuffle 随机地对数组进行洗牌。
func (a *IntArray) X随机排序() *IntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, v := range 随机类.X整数数组(len(a.array)) {
		a.array[i], a.array[v] = a.array[v], a.array[i]
	}
	return a
}

// Reverse 将数组元素按逆序排列生成新数组。
func (a *IntArray) X倒排序() *IntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, j := 0, len(a.array)-1; i < j; i, j = i+1, j-1 {
		a.array[i], a.array[j] = a.array[j], a.array[i]
	}
	return a
}

// Join 通过字符串 `glue` 连接数组元素。
func (a *IntArray) X连接(连接符 string) string {
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
func (a *IntArray) X统计() map[int]int {
	m := make(map[int]int)
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		m[v]++
	}
	return m
}

// Iterator 是 IteratorAsc 的别名。
func (a *IntArray) X遍历(回调函数 func(k int, v int) bool) {
	a.X遍历升序(回调函数)
}

// IteratorAsc 以升序遍历给定数组，并使用回调函数 `f` 进行只读操作。
// 如果 `f` 返回 true，则继续迭代；若返回 false，则停止遍历。
func (a *IntArray) X遍历升序(回调函数 func(k int, v int) bool) {
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
func (a *IntArray) X遍历降序(回调函数 func(k int, v int) bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for i := len(a.array) - 1; i >= 0; i-- {
		if !回调函数(i, a.array[i]) {
			break
		}
	}
}

// String 方法将当前数组以字符串形式返回，其实现方式类似于 json.Marshal。
func (a *IntArray) String() string {
	if a == nil {
		return ""
	}
	return "[" + a.X连接(",") + "]"
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
// 注意：此处接收者不使用指针。
func (a IntArray) MarshalJSON() ([]byte, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return json.Marshal(a.array)
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func (a *IntArray) UnmarshalJSON(b []byte) error {
	if a.array == nil {
		a.array = make([]int, 0)
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	if err := json.UnmarshalUseNumber(b, &a.array); err != nil {
		return err
	}
	return nil
}

// UnmarshalValue 实现了一个接口，该接口用于为数组设置任意类型的值。
func (a *IntArray) UnmarshalValue(value interface{}) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	switch value.(type) {
	case string, []byte:
		return json.UnmarshalUseNumber(转换类.X取字节集(value), &a.array)
	default:
		a.array = 转换类.SliceIne别名(value)
	}
	return nil
}

// Filter 对数组进行迭代，并通过自定义回调函数进行元素过滤。
// 如果回调函数 `filter` 返回 true，则从数组中移除该元素；
// 否则不做任何处理并继续迭代。
func (a *IntArray) X遍历删除(回调函数 func(索引 int, 值 int) bool) *IntArray {
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

// FilterEmpty 移除数组中的所有零值。
func (a *IntArray) X删除所有零值() *IntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i := 0; i < len(a.array); {
		if a.array[i] == 0 {
			a.array = append(a.array[:i], a.array[i+1:]...)
		} else {
			i++
		}
	}
	return a
}

// Walk 对数组中的每一项应用用户提供的函数 `f`。
func (a *IntArray) X遍历修改(回调函数 func(值 int) int) *IntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, v := range a.array {
		a.array[i] = 回调函数(v)
	}
	return a
}

// IsEmpty 检查数组是否为空。
func (a *IntArray) X是否为空() bool {
	return a.X取长度() == 0
}

// DeepCopy 实现接口，用于当前类型的深度复制。
func (a *IntArray) DeepCopy() interface{} {
	if a == nil {
		return nil
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	newSlice := make([]int, len(a.array))
	copy(newSlice, a.array)
	return X创建整数并从数组(newSlice, a.mu.IsSafe())
}

func (a *IntArray) X取文本() string {
return a.String()
}
