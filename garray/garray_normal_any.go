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
func X创建(并发安全 ...bool) *Array {
	return X创建并按大小(0, 0, 并发安全...)
}

// NewArray 是 New 的别名，请参阅 New。
func NewArray别名(并发安全 ...bool) *Array {
	return X创建并按大小(0, 0, 并发安全...)
}

// NewArraySize 根据给定的大小和容量创建并返回一个数组。
// 参数`safe`用于指定是否在并发安全的情况下使用数组，默认为false。
func X创建并按大小(大小 int, 上限 int, 并发安全 ...bool) *Array {
	return &Array{
		mu:    rwmutex.Create(并发安全...),
		array: make([]interface{}, 大小, 上限),
	}
}

// NewArrayRange 根据指定的范围从 `start` 到 `end` 并以步长值 `step` 创建并返回一个数组。
func X创建并按范围(起点, 终点, 步长 int, 并发安全 ...bool) *Array {
	if 步长 == 0 {
		panic(fmt.Sprintf(`invalid step value: %d`, 步长))
	}
	slice := make([]interface{}, 0)
	index := 0
	for i := 起点; i <= 终点; i += 步长 {
		slice = append(slice, i)
		index++
	}
	return X创建并从数组(slice, 并发安全...)
}

// NewFrom 是 NewArrayFrom 的别名。
// 请参阅 NewArrayFrom。
func NewFrom别名(数组 []interface{}, 并发安全 ...bool) *Array {
	return X创建并从数组(数组, 并发安全...)
}

// NewFromCopy 是 NewArrayFromCopy 的别名。
// 请参阅 NewArrayFromCopy。
func NewFromCopy别名(数组 []interface{}, 并发安全 ...bool) *Array {
	return X创建并从数组复制(数组, 并发安全...)
}

// NewArrayFrom 通过给定的切片 `array` 创建并返回一个数组。
// 参数 `safe` 用于指定是否在并发安全的情况下使用数组，默认为 false。
func X创建并从数组(数组 []interface{}, 并发安全 ...bool) *Array {
	return &Array{
		mu:    rwmutex.Create(并发安全...),
		array: 数组,
	}
}

// NewArrayFromCopy 通过复制给定切片 `array` 创建并返回一个数组。
// 参数 `safe` 用于指定是否在并发安全的情况下使用数组，默认为 false。
func X创建并从数组复制(数组 []interface{}, 并发安全 ...bool) *Array {
	newArray := make([]interface{}, len(数组))
	copy(newArray, 数组)
	return &Array{
		mu:    rwmutex.Create(并发安全...),
		array: newArray,
	}
}

// At通过指定的索引返回值。
// 如果给定的`index`超出数组范围，它将返回`nil`。
func (a *Array) X取值(索引 int) (值 interface{}) {
	值, _ = a.X取值2(索引)
	return
}

// Get 通过指定的索引返回值。
// 如果给定的 `index` 超出了数组的范围，那么 `found` 将为 false。
func (a *Array) X取值2(索引 int) (值 interface{}, 成功 bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if 索引 < 0 || 索引 >= len(a.array) {
		return nil, false
	}
	return a.array[索引], true
}

// Set将值设置到指定的索引位置。
func (a *Array) X设置值(索引 int, 值 interface{}) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if 索引 < 0 || 索引 >= len(a.array) {
		return gerror.NewCodef(gcode.CodeInvalidParameter, "index %d out of array range %d", 索引, len(a.array))
	}
	a.array[索引] = 值
	return nil
}

// SetArray 将底层的切片数组设置为给定的 `array`。
func (a *Array) X设置数组(数组 []interface{}) *Array {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.array = 数组
	return a
}

// Replace 从数组起始位置开始，使用给定的`array`替换原有数组元素。
func (a *Array) X替换(数组 []interface{}) *Array {
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
func (a *Array) X求和() (值 int) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		值 += gconv.Int(v)
	}
	return
}

// SortFunc 通过自定义函数 `less` 对数组进行排序。
func (a *Array) X排序并带函数(回调函数 func(v1, v2 interface{}) bool) *Array {
	a.mu.Lock()
	defer a.mu.Unlock()
	sort.Slice(a.array, func(i, j int) bool {
		return 回调函数(a.array[i], a.array[j])
	})
	return a
}

// InsertBefore 将 `values` 插入到 `index` 之前的位置。
func (a *Array) X插入前面(索引 int, 值 ...interface{}) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if 索引 < 0 || 索引 >= len(a.array) {
		return gerror.NewCodef(gcode.CodeInvalidParameter, "index %d out of array range %d", 索引, len(a.array))
	}
	rear := append([]interface{}{}, a.array[索引:]...)
	a.array = append(a.array[0:索引], 值...)
	a.array = append(a.array, rear...)
	return nil
}

// InsertAfter 在`index`之后插入`values`。
func (a *Array) X插入后面(索引 int, 值 ...interface{}) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if 索引 < 0 || 索引 >= len(a.array) {
		return gerror.NewCodef(gcode.CodeInvalidParameter, "index %d out of array range %d", 索引, len(a.array))
	}
	rear := append([]interface{}{}, a.array[索引+1:]...)
	a.array = append(a.array[0:索引+1], 值...)
	a.array = append(a.array, rear...)
	return nil
}

// Remove 通过索引移除一个元素。
// 如果给定的 `index` 超出了数组范围，`found` 将为 false。
func (a *Array) X删除(索引 int) (值 interface{}, 成功 bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.doRemoveWithoutLock(索引)
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
func (a *Array) X删除值(值 interface{}) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if i := a.doSearchWithoutLock(值); i != -1 {
		a.doRemoveWithoutLock(i)
		return true
	}
	return false
}

// RemoveValues 通过 `values` 移除多个项目。
func (a *Array) X删除多个值(值 ...interface{}) {
	a.mu.Lock()
	defer a.mu.Unlock()
	for _, value := range 值 {
		if i := a.doSearchWithoutLock(value); i != -1 {
			a.doRemoveWithoutLock(i)
		}
	}
}

// PushLeft 将一个或多个元素推送到数组的起始位置。
func (a *Array) X入栈左(值 ...interface{}) *Array {
	a.mu.Lock()
	a.array = append(值, a.array...)
	a.mu.Unlock()
	return a
}

// PushRight将一个或多个元素推送到数组的末尾。
// 它等同于Append。
func (a *Array) X入栈右(值 ...interface{}) *Array {
	a.mu.Lock()
	a.array = append(a.array, 值...)
	a.mu.Unlock()
	return a
}

// PopRand 随机地从数组中弹出并返回一个元素。
// 注意，如果数组为空，则 `found` 为 false。
func (a *Array) X出栈随机() (值 interface{}, 成功 bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.doRemoveWithoutLock(grand.Intn(len(a.array)))
}

// PopRands 随机地从数组中弹出并返回 `size` 个元素。
func (a *Array) X出栈随机多个(数量 int) []interface{} {
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
		array[i], _ = a.doRemoveWithoutLock(grand.Intn(len(a.array)))
	}
	return array
}

// PopLeft 从数组开头弹出并返回一个元素。
// 注意，如果数组为空，则 `found` 为 false。
func (a *Array) X出栈左() (值 interface{}, 成功 bool) {
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
func (a *Array) X出栈右() (值 interface{}, 成功 bool) {
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

// PopLefts 从数组开头弹出并返回 `size` 个元素。
func (a *Array) X出栈左多个(数量 int) []interface{} {
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
func (a *Array) X出栈右多个(数量 int) []interface{} {
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
func (a *Array) X取切片并按范围(起点 int, 终点 ...int) []interface{} {
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
func (a *Array) X取切片并按数量(起点 int, 数量 ...int) []interface{} {
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

// Append 是 PushRight 的别名，请参阅 PushRight。
func (a *Array) Append别名(值 ...interface{}) *Array {
	a.X入栈右(值...)
	return a
}

// Len 返回数组的长度。
func (a *Array) X取长度() int {
	a.mu.RLock()
	length := len(a.array)
	a.mu.RUnlock()
	return length
}

// Slice 返回数组的基础数据。
// 注意，如果它在并发安全的使用场景下，会返回基础数据的一个副本，
// 否则，则返回指向基础数据的指针。
func (a *Array) X取切片() []interface{} {
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
	return a.X取切片()
}

// Clone 返回一个新的数组，它是当前数组的一个副本。
func (a *Array) X取副本() (新数组 *Array) {
	a.mu.RLock()
	array := make([]interface{}, len(a.array))
	copy(array, a.array)
	a.mu.RUnlock()
	return X创建并从数组(array, a.mu.IsSafe())
}

// 清空删除当前数组中的所有元素。
func (a *Array) X清空() *Array {
	a.mu.Lock()
	if len(a.array) > 0 {
		a.array = make([]interface{}, 0)
	}
	a.mu.Unlock()
	return a
}

// Contains 检查某个值是否存在于数组中。
func (a *Array) X是否存在(值 interface{}) bool {
	return a.X查找(值) != -1
}

// Search 在数组中通过 `value` 进行搜索，返回 `value` 的索引，
// 若不存在，则返回 -1。
func (a *Array) X查找(值 interface{}) int {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.doSearchWithoutLock(值)
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
func (a *Array) X去重() *Array {
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
func (a *Array) X遍历并写锁定(回调函数 func(数组 []interface{})) *Array {
	a.mu.Lock()
	defer a.mu.Unlock()
	回调函数(a.array)
	return a
}

// RLockFunc 通过回调函数`f`锁定读取操作。
func (a *Array) X遍历并读锁定(回调函数 func(数组 []interface{})) *Array {
	a.mu.RLock()
	defer a.mu.RUnlock()
	回调函数(a.array)
	return a
}

// Merge 将`array`合并到当前数组中。
// 参数`array`可以是任何garray类型或切片类型。
// Merge 和 Append 的区别在于，Append 仅支持特定类型的切片作为参数，
// 而 Merge 支持更多类型的参数。
func (a *Array) X合并(数组 interface{}) *Array {
	return a.Append别名(gconv.Interfaces(数组)...)
}

// Fill 用 `value` 值填充数组，填充 num 个条目，
// 键（索引）从 `startIndex` 参数开始。
func (a *Array) X填充(起点 int, 填充数量 int, 值 interface{}) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if 起点 < 0 || 起点 > len(a.array) {
		return gerror.NewCodef(gcode.CodeInvalidParameter, "index %d out of array range %d", 起点, len(a.array))
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
func (a *Array) X分割(数量 int) [][]interface{} {
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

// Pad 通过 `value` 值对数组进行填充，以达到指定长度。
// 如果 size 为正数，则在数组右侧进行填充；若为负数，则在左侧填充。
// 若 `size` 的绝对值小于或等于数组的长度，则不进行填充操作。
func (a *Array) X填满(总数量 int, 值 interface{}) *Array {
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
	tmp := make([]interface{}, n)
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
func (a *Array) X取值随机() (值 interface{}, 成功 bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if len(a.array) == 0 {
		return nil, false
	}
	return a.array[grand.Intn(len(a.array))], true
}

// Rands 随机返回数组中的 `size` 个元素（不删除）。
func (a *Array) X取值随机多个(数量 int) []interface{} {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if 数量 <= 0 || len(a.array) == 0 {
		return nil
	}
	array := make([]interface{}, 数量)
	for i := 0; i < 数量; i++ {
		array[i] = a.array[grand.Intn(len(a.array))]
	}
	return array
}

// Shuffle 随机地对数组进行洗牌。
func (a *Array) X随机排序() *Array {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, v := range grand.Perm(len(a.array)) {
		a.array[i], a.array[v] = a.array[v], a.array[i]
	}
	return a
}

// Reverse 将数组元素按逆序排列生成新数组。
func (a *Array) X倒排序() *Array {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, j := 0, len(a.array)-1; i < j; i, j = i+1, j-1 {
		a.array[i], a.array[j] = a.array[j], a.array[i]
	}
	return a
}

// Join 通过字符串 `glue` 连接数组元素。
func (a *Array) X连接(连接符 string) string {
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

// CountValues 计算数组中所有值出现的次数。
func (a *Array) X统计() map[interface{}]int {
	m := make(map[interface{}]int)
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		m[v]++
	}
	return m
}

// Iterator 是 IteratorAsc 的别名。
func (a *Array) X遍历(回调函数 func(k int, v interface{}) bool) {
	a.X遍历升序(回调函数)
}

// IteratorAsc 以升序遍历给定数组，并使用回调函数 `f` 进行只读操作。
// 如果 `f` 返回 true，则继续迭代；若返回 false，则停止遍历。
func (a *Array) X遍历升序(回调函数 func(k int, v interface{}) bool) {
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
func (a *Array) X遍历降序(回调函数 func(k int, v interface{}) bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for i := len(a.array) - 1; i >= 0; i-- {
		if !回调函数(i, a.array[i]) {
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
func (a *Array) X遍历删除(回调函数 func(索引 int, 值 interface{}) bool) *Array {
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

// FilterNil 移除数组中所有的 nil 值。
func (a *Array) X删除所有nil() *Array {
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
func (a *Array) X删除所有空值() *Array {
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
func (a *Array) X遍历修改(回调函数 func(值 interface{}) interface{}) *Array {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, v := range a.array {
		a.array[i] = 回调函数(v)
	}
	return a
}

// IsEmpty 检查数组是否为空。
func (a *Array) X是否为空() bool {
	return a.X取长度() == 0
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
	return X创建并从数组(newSlice, a.mu.IsSafe())
}

func (a *Array) X取文本() string {
	return a.String()
}

func (a *Array) X取any数组() []interface{} {
return a.Interfaces()
}
