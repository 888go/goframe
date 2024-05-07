// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 切片类

import (
	"bytes"
	"math"
	"sort"
	"strings"

	"github.com/888go/goframe/garray/internal/json"
	"github.com/888go/goframe/garray/internal/rwmutex"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

// StrArray 是一个具有丰富特性的 Go 语言字符串切片。
// 它包含一个并发安全/不安全的开关，在初始化时应设置该开关，之后不可更改。
// ```go
// StrArray 是一个功能丰富的 Golang 字符串切片类型。
// 其中包含一个并发安全模式切换选项，应在初始化时设定，并且一旦设定后不可再更改。
type StrArray struct {
	mu    rwmutex.RWMutex
	array []string
}

// NewStrArray 创建并返回一个空字符串切片。
// 参数`safe`用于指定是否在并发安全的情况下使用切片，默认为false。
func X创建文本(并发安全 ...bool) *StrArray {
	return X创建文本并按大小(0, 0, 并发安全...)
}

// NewStrArraySize 根据给定的大小和容量创建并返回一个切片。
// 参数 `safe` 用于指定是否在并发安全的情况下使用切片，默认为 false。
func X创建文本并按大小(大小 int, 上限 int, 并发安全 ...bool) *StrArray {
	return &StrArray{
		mu:    rwmutex.Create(并发安全...),
		array: make([]string, 大小, 上限),
	}
}

// NewStrArrayFrom 根据给定的切片 `array` 创建并返回一个切片。
// 参数 `safe` 用于指定是否在并发安全的情况下使用切片，默认为 false。
func X创建文本并从切片(切片 []string, 并发安全 ...bool) *StrArray {
	return &StrArray{
		mu:    rwmutex.Create(并发安全...),
		array: 切片,
	}
}

// NewStrArrayFromCopy 通过复制给定切片 `array` 创建并返回一个切片。
// 参数 `safe` 用于指定是否在并发安全环境下使用该切片，默认为 false。
func X创建文本并从切片复制(切片 []string, 并发安全 ...bool) *StrArray {
	newArray := make([]string, len(切片))
	copy(newArray, 切片)
	return &StrArray{
		mu:    rwmutex.Create(并发安全...),
		array: newArray,
	}
}

// At 通过指定的索引返回值。
// 如果给定的 `index` 超出了切片的范围，它将返回一个空字符串。
func (a *StrArray) X取值(索引 int) (值 string) {
	值, _ = a.X取值2(索引)
	return
}

// Get 通过指定的索引返回值。
// 如果给定的 `index` 超出了切片的范围，那么 `found` 将为 false。
func (a *StrArray) X取值2(索引 int) (值 string, 成功 bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if 索引 < 0 || 索引 >= len(a.array) {
		return "", false
	}
	return a.array[索引], true
}

// Set将值设置到指定的索引位置。
func (a *StrArray) X设置值(索引 int, 值 string) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if 索引 < 0 || 索引 >= len(a.array) {
		return gerror.NewCodef(gcode.CodeInvalidParameter, "index %d out of array range %d", 索引, len(a.array))
	}
	a.array[索引] = 值
	return nil
}

// SetArray 将底层的切片切片设置为给定的 `array`。
func (a *StrArray) X设置切片(切片 []string) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.array = 切片
	return a
}

// Replace 从切片起始位置开始，使用给定的`array`替换原有切片元素。
func (a *StrArray) X替换(切片 []string) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	max := len(切片)
	if max > len(a.array) {
		max = len(a.array)
	}
	for i := 0; i < max; i++ {
		a.array[i] = 切片[i]
	}
	return a
}

// Sum 返回切片中所有值的和。
func (a *StrArray) X求和() (值 int) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		值 += gconv.Int(v)
	}
	return
}

// Sort 函数用于将切片按升序排序。
// 参数 `reverse` 控制排序方式，若 reverse 为 true，则按降序（默认为升序）排序。
func (a *StrArray) X排序递增(降序 ...bool) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	if len(降序) > 0 && 降序[0] {
		sort.Slice(a.array, func(i, j int) bool {
			return strings.Compare(a.array[i], a.array[j]) >= 0
		})
	} else {
		sort.Strings(a.array)
	}
	return a
}

// SortFunc 通过自定义函数 `less` 对切片进行排序。
func (a *StrArray) X排序函数(回调函数 func(v1, v2 string) bool) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	sort.Slice(a.array, func(i, j int) bool {
		return 回调函数(a.array[i], a.array[j])
	})
	return a
}

// InsertBefore 将 `values` 插入到 `index` 之前的位置。
func (a *StrArray) X插入前面(索引 int, 值 ...string) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if 索引 < 0 || 索引 >= len(a.array) {
		return gerror.NewCodef(gcode.CodeInvalidParameter, "index %d out of array range %d", 索引, len(a.array))
	}
	rear := append([]string{}, a.array[索引:]...)
	a.array = append(a.array[0:索引], 值...)
	a.array = append(a.array, rear...)
	return nil
}

// InsertAfter 在`index`之后插入`values`。
func (a *StrArray) X插入后面(索引 int, 值 ...string) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if 索引 < 0 || 索引 >= len(a.array) {
		return gerror.NewCodef(gcode.CodeInvalidParameter, "index %d out of array range %d", 索引, len(a.array))
	}
	rear := append([]string{}, a.array[索引+1:]...)
	a.array = append(a.array[0:索引+1], 值...)
	a.array = append(a.array, rear...)
	return nil
}

// Remove 通过索引移除一个元素。
// 如果给定的 `index` 超出了切片范围，`found` 将为 false。
func (a *StrArray) X删除(索引 int) (值 string, 成功 bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.doRemoveWithoutLock(索引)
}

// doRemoveWithoutLock 在没有加锁的情况下通过索引移除一个项。
func (a *StrArray) doRemoveWithoutLock(index int) (value string, found bool) {
	if index < 0 || index >= len(a.array) {
		return "", false
	}
	// 确定删除时的切片边界以提高删除效率
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
	// 那么它将涉及创建一个切片，
	// 因此，删除操作效率较低。
	value = a.array[index]
	a.array = append(a.array[:index], a.array[index+1:]...)
	return value, true
}

// RemoveValue 通过值移除一个元素。
// 若在切片中找到该值，则返回 true，否则（未找到时）返回 false。
func (a *StrArray) X删除值(值 string) bool {
	if i := a.X查找(值); i != -1 {
		_, found := a.X删除(i)
		return found
	}
	return false
}

// RemoveValues 通过 `values` 移除多个项目。
func (a *StrArray) X删除多个值(值 ...string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	for _, value := range 值 {
		if i := a.doSearchWithoutLock(value); i != -1 {
			a.doRemoveWithoutLock(i)
		}
	}
}

// PushLeft 将一个或多个元素推送到切片的起始位置。
func (a *StrArray) X入栈左(值 ...string) *StrArray {
	a.mu.Lock()
	a.array = append(值, a.array...)
	a.mu.Unlock()
	return a
}

// PushRight将一个或多个元素推送到切片的末尾。
// 它等同于Append。
func (a *StrArray) X入栈右(值 ...string) *StrArray {
	a.mu.Lock()
	a.array = append(a.array, 值...)
	a.mu.Unlock()
	return a
}

// PopLeft 从切片开头弹出并返回一个元素。
// 注意，如果切片为空，则 `found` 为 false。
func (a *StrArray) X出栈左() (值 string, 成功 bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if len(a.array) == 0 {
		return "", false
	}
	值 = a.array[0]
	a.array = a.array[1:]
	return 值, true
}

// PopRight从切片的末尾弹出并返回一个元素。
// 注意，如果切片为空，则`found`为false。
func (a *StrArray) X出栈右() (值 string, 成功 bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	index := len(a.array) - 1
	if index < 0 {
		return "", false
	}
	值 = a.array[index]
	a.array = a.array[:index]
	return 值, true
}

// PopRand 随机地从切片中弹出并返回一个元素。
// 注意，如果切片为空，则 `found` 为 false。
func (a *StrArray) X出栈随机() (值 string, 成功 bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.doRemoveWithoutLock(grand.Intn(len(a.array)))
}

// PopRands 随机地从切片中弹出并返回 `size` 个元素。
// 若给定的 `size` 大于切片的大小，则返回切片中的所有元素。
// 注意，如果给定的 `size` 小于等于0或者切片为空，它将返回 nil。
func (a *StrArray) X出栈随机多个(数量 int) []string {
	a.mu.Lock()
	defer a.mu.Unlock()
	if 数量 <= 0 || len(a.array) == 0 {
		return nil
	}
	if 数量 >= len(a.array) {
		数量 = len(a.array)
	}
	array := make([]string, 数量)
	for i := 0; i < 数量; i++ {
		array[i], _ = a.doRemoveWithoutLock(grand.Intn(len(a.array)))
	}
	return array
}

// PopLefts 从切片开头弹出并返回 `size` 个元素。
// 如果给定的 `size` 大于切片的大小，则返回切片中的所有元素。
// 注意，如果给定的 `size` 小于等于0或者切片为空，则返回nil。
func (a *StrArray) X出栈左多个(数量 int) []string {
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

// PopRights 从切片末尾弹出并返回 `size` 个元素。
// 若给定的 `size` 大于切片的大小，则返回切片中所有元素。
// 注意，如果给定的 `size` 小于等于0或者切片为空，它将返回nil。
func (a *StrArray) X出栈右多个(数量 int) []string {
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

// Range 函数通过范围选择并返回切片中的元素，类似于 array[start:end]。
// 注意：在并发安全的使用场景下，它会返回一个原数据的副本；否则，返回的是底层数据的指针。
//
// 如果 `end` 为负数，则偏移量将从切片末尾开始计算。
// 如果省略了 `end`，则序列将包含从 start 开始直到切片末尾的所有元素。
func (a *StrArray) X取切片并按范围(起点 int, 终点 ...int) []string {
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
	array := ([]string)(nil)
	if a.mu.IsSafe() {
		array = make([]string, offsetEnd-起点)
		copy(array, a.array[起点:offsetEnd])
	} else {
		array = a.array[起点:offsetEnd]
	}
	return array
}

// SubSlice 返回切片中由 `offset` 和 `size` 参数指定的元素子序列，并将其作为切片。
// 若在并发安全场景下使用，返回该切片的副本；否则返回指向切片的指针。
//
// 如果 offset 非负，则序列从切片该偏移位置开始。
// 如果 offset 为负，则序列从切片末尾向前偏移该距离的位置开始。
//
// 如果提供了 length 并且为正数，则序列将包含最多该数量的元素。
// 若切片长度小于 length，则序列仅包含切片中可获得的元素。
// 如果 length 为负数，则序列将在切片末尾向前停在该距离的位置。
// 如果未提供 length，则序列包含从 offset 开始直到切片末尾的所有元素。
//
// 若有任何可能穿越切片左边界的情况，函数将失败。
func (a *StrArray) X取切片并按数量(起点 int, 数量 ...int) []string {
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
		s := make([]string, size)
		copy(s, a.array[起点:])
		return s
	}
	return a.array[起点:end]
}

// Append 是 PushRight 的别名，请参阅 PushRight。
func (a *StrArray) Append别名(值 ...string) *StrArray {
	a.mu.Lock()
	a.array = append(a.array, 值...)
	a.mu.Unlock()
	return a
}

// Len 返回切片的长度。
func (a *StrArray) X取长度() int {
	a.mu.RLock()
	length := len(a.array)
	a.mu.RUnlock()
	return length
}

// Slice 返回切片的基础数据。
// 注意，如果它在并发安全的使用场景下，会返回基础数据的一个副本，
// 否则，则返回指向基础数据的指针。
func (a *StrArray) X取切片() []string {
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

// Interfaces 函数将当前切片转换为 []interface{} 类型并返回。
func (a *StrArray) Interfaces() []interface{} {
	a.mu.RLock()
	defer a.mu.RUnlock()
	array := make([]interface{}, len(a.array))
	for k, v := range a.array {
		array[k] = v
	}
	return array
}

// Clone 返回一个新的切片，它是当前切片的一个副本。
func (a *StrArray) X取副本() (新切片 *StrArray) {
	a.mu.RLock()
	array := make([]string, len(a.array))
	copy(array, a.array)
	a.mu.RUnlock()
	return X创建文本并从切片(array, a.mu.IsSafe())
}

// 清空删除当前切片中的所有元素。
func (a *StrArray) X清空() *StrArray {
	a.mu.Lock()
	if len(a.array) > 0 {
		a.array = make([]string, 0)
	}
	a.mu.Unlock()
	return a
}

// Contains 检查某个值是否存在于切片中。
func (a *StrArray) X是否存在(值 string) bool {
	return a.X查找(值) != -1
}

// ContainsI 检查一个值是否以不区分大小写的方式存在于切片中。
// 注意，它在内部会遍历整个切片并进行不区分大小写的比较。
func (a *StrArray) X是否存在并忽略大小写(值 string) bool {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if len(a.array) == 0 {
		return false
	}
	for _, v := range a.array {
		if strings.EqualFold(v, 值) {
			return true
		}
	}
	return false
}

// Search 在切片中通过 `value` 进行搜索，返回 `value` 的索引，
// 若不存在，则返回 -1。
func (a *StrArray) X查找(值 string) int {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.doSearchWithoutLock(值)
}

func (a *StrArray) doSearchWithoutLock(value string) int {
	if len(a.array) == 0 {
		return -1
	}
	result := -1
	for index, v := range a.array {
		if strings.Compare(v, value) == 0 {
			result = index
			break
		}
	}
	return result
}

// Unique 函数用于对切片去重，清除重复的元素。
// 示例：[1,1,2,3,2] -> [1,2,3]
func (a *StrArray) X去重() *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	if len(a.array) == 0 {
		return a
	}
	var (
		ok          bool
		temp        string
		uniqueSet   = make(map[string]struct{})
		uniqueArray = make([]string, 0, len(a.array))
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
func (a *StrArray) X遍历写锁定(回调函数 func(切片 []string)) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	回调函数(a.array)
	return a
}

// RLockFunc 通过回调函数`f`锁定读取操作。
func (a *StrArray) X遍历读锁定(回调函数 func(切片 []string)) *StrArray {
	a.mu.RLock()
	defer a.mu.RUnlock()
	回调函数(a.array)
	return a
}

// Merge 将`array`合并到当前切片中。
// 参数`array`可以是任何garray类型或切片类型。
// Merge 和 Append 的区别在于，Append 仅支持特定类型的切片作为参数，
// 而 Merge 支持更多类型的参数。
func (a *StrArray) X合并(切片 interface{}) *StrArray {
	return a.Append别名(gconv.Strings(切片)...)
}

// Fill 用 `value` 值填充切片，填充 num 个条目，
// 键（索引）从 `startIndex` 参数开始。
func (a *StrArray) X填充(起点 int, 填充数量 int, 值 string) error {
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

// Chunk 函数将一个切片分割成多个子切片，
// 每个子切片的大小由参数 `size` 确定。
// 最后一个子切片可能包含少于 size 个元素。
func (a *StrArray) X分割(数量 int) [][]string {
	if 数量 < 1 {
		return nil
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	length := len(a.array)
	chunks := int(math.Ceil(float64(length) / float64(数量)))
	var n [][]string
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

// Pad 通过 `value` 值对切片进行填充，以达到指定长度。
// 如果 size 为正数，则在切片右侧进行填充；若为负数，则在左侧填充。
// 若 `size` 的绝对值小于或等于切片的长度，则不进行填充操作。
func (a *StrArray) X填满(总数量 int, 值 string) *StrArray {
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
	tmp := make([]string, n)
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

// Rand 随机地从切片中返回一个元素（不删除）。
func (a *StrArray) X取值随机() (值 string, 成功 bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if len(a.array) == 0 {
		return "", false
	}
	return a.array[grand.Intn(len(a.array))], true
}

// Rands 随机返回切片中的 `size` 个元素（不删除）。
func (a *StrArray) X取值随机多个(数量 int) []string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if 数量 <= 0 || len(a.array) == 0 {
		return nil
	}
	array := make([]string, 数量)
	for i := 0; i < 数量; i++ {
		array[i] = a.array[grand.Intn(len(a.array))]
	}
	return array
}

// Shuffle 随机地对切片进行洗牌。
func (a *StrArray) X随机排序() *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, v := range grand.Perm(len(a.array)) {
		a.array[i], a.array[v] = a.array[v], a.array[i]
	}
	return a
}

// Reverse 将切片元素按逆序排列生成新切片。
func (a *StrArray) X倒排序() *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, j := 0, len(a.array)-1; i < j; i, j = i+1, j-1 {
		a.array[i], a.array[j] = a.array[j], a.array[i]
	}
	return a
}

// Join 通过字符串 `glue` 连接切片元素。
func (a *StrArray) X连接(连接符 string) string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if len(a.array) == 0 {
		return ""
	}
	buffer := bytes.NewBuffer(nil)
	for k, v := range a.array {
		buffer.WriteString(v)
		if k != len(a.array)-1 {
			buffer.WriteString(连接符)
		}
	}
	return buffer.String()
}

// CountValues 计算切片中所有值出现的次数。
func (a *StrArray) X统计() map[string]int {
	m := make(map[string]int)
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		m[v]++
	}
	return m
}

// Iterator 是 IteratorAsc 的别名。
func (a *StrArray) X遍历(回调函数 func(k int, v string) bool) {
	a.X遍历升序(回调函数)
}

// IteratorAsc 以升序遍历给定切片，并使用回调函数 `f` 进行只读操作。
// 如果 `f` 返回 true，则继续迭代；若返回 false，则停止遍历。
func (a *StrArray) X遍历升序(回调函数 func(k int, v string) bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for k, v := range a.array {
		if !回调函数(k, v) {
			break
		}
	}
}

// IteratorDesc 函数以降序遍历给定的切片，并使用指定回调函数 `f` 进行只读操作。
// 若 `f` 返回 true，则继续迭代；若返回 false，则停止迭代。
func (a *StrArray) X遍历降序(回调函数 func(k int, v string) bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for i := len(a.array) - 1; i >= 0; i-- {
		if !回调函数(i, a.array[i]) {
			break
		}
	}
}

// String 方法将当前切片以字符串形式返回，其实现方式类似于 json.Marshal。
func (a *StrArray) String() string {
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
func (a StrArray) MarshalJSON() ([]byte, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return json.Marshal(a.array)
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func (a *StrArray) UnmarshalJSON(b []byte) error {
	if a.array == nil {
		a.array = make([]string, 0)
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	if err := json.UnmarshalUseNumber(b, &a.array); err != nil {
		return err
	}
	return nil
}

// UnmarshalValue 实现了一个接口，该接口用于为切片设置任意类型的值。
func (a *StrArray) UnmarshalValue(value interface{}) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	switch value.(type) {
	case string, []byte:
		return json.UnmarshalUseNumber(gconv.Bytes(value), &a.array)
	default:
		a.array = gconv.SliceStr(value)
	}
	return nil
}

// Filter 对切片进行迭代，并通过自定义回调函数进行元素过滤。
// 如果回调函数 `filter` 返回 true，则从切片中移除该元素；
// 否则不做任何处理并继续迭代。
func (a *StrArray) X遍历删除(回调函数 func(索引 int, 值 string) bool) *StrArray {
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

// FilterEmpty 从切片中移除所有空字符串值。
func (a *StrArray) X删除所有空值() *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i := 0; i < len(a.array); {
		if a.array[i] == "" {
			a.array = append(a.array[:i], a.array[i+1:]...)
		} else {
			i++
		}
	}
	return a
}

// Walk 对切片中的每一项应用用户提供的函数 `f`。
func (a *StrArray) X遍历修改(回调函数 func(值 string) string) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, v := range a.array {
		a.array[i] = 回调函数(v)
	}
	return a
}

// IsEmpty 检查切片是否为空。
func (a *StrArray) X是否为空() bool {
	return a.X取长度() == 0
}

// DeepCopy 实现接口，用于当前类型的深度复制。
func (a *StrArray) DeepCopy() interface{} {
	if a == nil {
		return nil
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	newSlice := make([]string, len(a.array))
	copy(newSlice, a.array)
	return X创建文本并从切片(newSlice, a.mu.IsSafe())
}

func (a *StrArray) X取文本() string {
	return a.String()
}

func (a *StrArray) X取any切片() []interface{} {
	return a.Interfaces()
}
