// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 切片类

import (
	"bytes"
	"math"
	"sort"
	"strings"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/rwmutex"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
	grand "github.com/888go/goframe/util/grand"
)

// StrArray 是一个具有丰富功能的 Go 语言字符串数组。
// 它包含一个并发安全/不安全的开关，该开关应在初始化时设置，并且之后不能更改。
// md5:60bf9d0fe402df8a
type StrArray struct {
	mu    rwmutex.RWMutex
	array []string
}

// X创建文本 创建并返回一个空数组。
// 参数 `safe` 用于指定是否在并发安全模式下使用数组，默认为 false。
// md5:1a16d6b7fa6dc90d
func X创建文本(并发安全 ...bool) *StrArray {
	return X创建文本并按大小(0, 0, 并发安全...)
}

// X创建文本并按大小 创建并返回一个给定大小和容量的数组。
// 参数 `safe` 用于指定是否在并发安全模式下使用数组，默认为 false。
// md5:d419c5b3ffb2a682
func X创建文本并按大小(大小 int, 上限 int, 并发安全 ...bool) *StrArray {
	return &StrArray{
		mu:    rwmutex.Create(并发安全...),
		array: make([]string, 大小, 上限),
	}
}

// X创建文本并从切片 根据给定的切片 `array` 创建并返回一个数组。
// 参数 `safe` 用于指定是否使用并发安全的数组，默认为 false。
// md5:719d22a529b420db
func X创建文本并从切片(切片 []string, 并发安全 ...bool) *StrArray {
	return &StrArray{
		mu:    rwmutex.Create(并发安全...),
		array: 切片,
	}
}

// X创建文本并从切片复制 根据给定切片 `array` 的副本创建并返回一个数组。
// 参数 `safe` 用于指定是否在并发安全环境下使用数组，默认为 false。
// md5:71bd55b1c0df65be
func X创建文本并从切片复制(切片 []string, 并发安全 ...bool) *StrArray {
	newArray := make([]string, len(切片))
	copy(newArray, 切片)
	return &StrArray{
		mu:    rwmutex.Create(并发安全...),
		array: newArray,
	}
}

// X取值通过指定的索引返回值。
// 如果给定的`index`超出了数组的范围，它将返回一个空字符串。
// md5:2465f6b1e3ac2863
func (a *StrArray) X取值(索引 int) (值 string) {
	值, _ = a.X取值2(索引)
	return
}

// X取值2 函数通过指定的索引返回值。
// 如果给定的 `index` 超出了数组范围，`found` 将为 false。
// md5:ab300cfc0d6dd8ee
func (a *StrArray) X取值2(索引 int) (值 string, 成功 bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if 索引 < 0 || 索引 >= len(a.array) {
		return "", false
	}
	return a.array[索引], true
}

// X设置值 设置指定索引的值。 md5:7c1d7ea9df0b722c
func (a *StrArray) X设置值(index int, value string) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if index < 0 || index >= len(a.array) {
		return gerror.X创建错误码并格式化(gcode.CodeInvalidParameter, "index %d out of array range %d", index, len(a.array))
	}
	a.array[index] = value
	return nil
}

// X设置切片 使用给定的 `array` 设置底层切片数组。 md5:160b43a5c0ec752c
func (a *StrArray) X设置切片(切片 []string) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.array = 切片
	return a
}

// X替换 从数组的起始位置开始，使用给定的 `array` 替换数组中的元素。 md5:5acead2fd9ec0761
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

// X求和 返回数组中所有值的和。 md5:b2148175a749b162
func (a *StrArray) X求和() (值 int) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		值 += gconv.X取整数(v)
	}
	return
}

// X排序递增 对数组进行升序排序。
// 参数 `reverse` 控制排序方式，如果为 true，则降序排列（默认为升序）。
// md5:35d4650a0f563ccf
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

// X排序函数 使用自定义函数 `less` 对数组进行排序。 md5:8da07d09bbd08513
func (a *StrArray) X排序函数(回调函数 func(v1, v2 string) bool) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	sort.Slice(a.array, func(i, j int) bool {
		return 回调函数(a.array[i], a.array[j])
	})
	return a
}

// X插入前面 将`values`插入到`index`的前面。 md5:f5f3b46cd17ba885
func (a *StrArray) X插入前面(索引 int, 值 ...string) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if 索引 < 0 || 索引 >= len(a.array) {
		return gerror.X创建错误码并格式化(gcode.CodeInvalidParameter, "index %d out of array range %d", 索引, len(a.array))
	}
	rear := append([]string{}, a.array[索引:]...)
	a.array = append(a.array[0:索引], 值...)
	a.array = append(a.array, rear...)
	return nil
}

// X插入后面 将 `values` 插入到 `index` 后面。 md5:b90b80fa75b6b6e0
func (a *StrArray) X插入后面(索引 int, 值 ...string) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if 索引 < 0 || 索引 >= len(a.array) {
		return gerror.X创建错误码并格式化(gcode.CodeInvalidParameter, "index %d out of array range %d", 索引, len(a.array))
	}
	rear := append([]string{}, a.array[索引+1:]...)
	a.array = append(a.array[0:索引+1], 值...)
	a.array = append(a.array, rear...)
	return nil
}

// X删除 函数通过索引移除一个元素。
// 如果给定的 `index` 超出了数组范围，`found` 将为 false。
// md5:feaf958654838c25
func (a *StrArray) X删除(索引 int) (值 string, 成功 bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.doRemoveWithoutLock(索引)
}

// doRemoveWithoutLock 不使用锁移除一个项目。 md5:a6a1746903fd131c
func (a *StrArray) doRemoveWithoutLock(index int) (value string, found bool) {
	if index < 0 || index >= len(a.array) {
		return "", false
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
func (a *StrArray) X删除值(值 string) bool {
	if i := a.X查找(值); i != -1 {
		_, found := a.X删除(i)
		return found
	}
	return false
}

// X删除多个值 根据`values`移除多个项目。 md5:fbdf68fa6a8cdd26
func (a *StrArray) X删除多个值(值 ...string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	for _, value := range 值 {
		if i := a.doSearchWithoutLock(value); i != -1 {
			a.doRemoveWithoutLock(i)
		}
	}
}

// X入栈左 将一个或多个项目推送到数组的开头。 md5:9062afab48970bed
func (a *StrArray) X入栈左(值 ...string) *StrArray {
	a.mu.Lock()
	a.array = append(值, a.array...)
	a.mu.Unlock()
	return a
}

// X入栈右 将一个或多个元素添加到数组的末尾。
// 它等同于 Append。
// md5:bb33f2edfdfd9896
func (a *StrArray) X入栈右(值 ...string) *StrArray {
	a.mu.Lock()
	a.array = append(a.array, 值...)
	a.mu.Unlock()
	return a
}

// X出栈左 从数组的开头弹出并返回一个项目。
// 注意，如果数组为空，`found` 为 false。
// md5:68f14002d84594a4
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

// X出栈右 从数组的末尾弹出并返回一个元素。
// 注意，如果数组为空，则 `found` 为 false。
// md5:207fa7c7c4a04a10
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

// X出栈随机 从数组中随机弹出并返回一个元素。
// 注意，如果数组为空，`found` 将为 false。
// md5:29338267db400401
func (a *StrArray) X出栈随机() (值 string, 成功 bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.doRemoveWithoutLock(grand.X整数(len(a.array)))
}

// X出栈随机多个 随机地从数组中弹出并返回 `size` 个元素。
// 如果给定的 `size` 大于数组的大小，它将返回数组的所有元素。
// 注意，如果给定的 `size` 小于等于 0 或数组为空，它将返回 nil。
// md5:9fd270d3d3021d32
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
		array[i], _ = a.doRemoveWithoutLock(grand.X整数(len(a.array)))
	}
	return array
}

// X出栈左多个 从数组开始处弹出并返回 `size` 个元素。
// 如果给定的 `size` 大于数组的长度，它将返回数组中的所有元素。
// 请注意，如果给定的 `size` 小于等于 0 或数组为空，它将返回 nil。
// md5:3ecbe066336a9849
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

// X出栈右多个 从数组末尾弹出并返回 `size` 个元素。
// 如果给定的 `size` 大于数组的大小，它将返回数组中的所有元素。
// 注意，如果给定的 `size` 小于等于 0 或数组为空，它将返回 nil。
// md5:4f44f32fbb68fb50
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

// X取切片并按范围通过范围选择并返回项目，就像数组[start:end]一样。
// 请注意，如果在并发安全使用中，它将返回切片的副本；否则返回底层数据的指针。
// 
// 如果`end`为负数，则偏移量将从数组末尾开始。
// 如果省略`end`，则序列将包含从`start`到数组结尾的所有内容。
// md5:8b71690536bb9ec5
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

// Append别名 是 PushRight 的别名，详情请参阅 PushRight。 md5:2f083a022f7fd9c3
func (a *StrArray) Append别名(值 ...string) *StrArray {
	a.mu.Lock()
	a.array = append(a.array, 值...)
	a.mu.Unlock()
	return a
}

// X取长度 返回数组的长度。 md5:593b37501e98da95
func (a *StrArray) X取长度() int {
	a.mu.RLock()
	length := len(a.array)
	a.mu.RUnlock()
	return length
}

// X取切片 返回数组的底层数据。
// 注意，如果在并发安全的使用情况下，它会返回底层数据的副本，否则返回底层数据的指针。
// md5:111cbee45795a58b
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

// X取any切片 将当前数组作为 []interface{} 返回。 md5:f7a2e3459e185314
func (a *StrArray) X取any切片() []interface{} {
	a.mu.RLock()
	defer a.mu.RUnlock()
	array := make([]interface{}, len(a.array))
	for k, v := range a.array {
		array[k] = v
	}
	return array
}

// X取副本 返回一个新的数组，它是当前数组的副本。 md5:52ada4030c562295
func (a *StrArray) X取副本() (新切片 *StrArray) {
	a.mu.RLock()
	array := make([]string, len(a.array))
	copy(array, a.array)
	a.mu.RUnlock()
	return X创建文本并从切片(array, a.mu.IsSafe())
}

// X清空 删除当前数组中的所有项目。 md5:3d9c6d68a5719979
func (a *StrArray) X清空() *StrArray {
	a.mu.Lock()
	if len(a.array) > 0 {
		a.array = make([]string, 0)
	}
	a.mu.Unlock()
	return a
}

// X是否存在 检查值是否存在于数组中。 md5:f209e1f30dd53cb2
func (a *StrArray) X是否存在(值 string) bool {
	return a.X查找(值) != -1
}

// X是否存在并忽略大小写 检查数组中是否存在某个值（忽略大小写）。
// 注意，它内部会遍历整个数组以进行不区分大小写的比较。
// md5:faf76a65365aa0ac
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

// X查找 在数组中搜索 `value`，返回 `value` 的索引，
// 如果不存在则返回 -1。
// md5:787617bfeade8f93
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

// X去重 去除数组中的重复元素。
// 例如：[1,1,2,3,2] -> [1,2,3]
// md5:5083aa414231fd30
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

// X遍历写锁定 通过回调函数 `f` 实现写入锁定。 md5:d45a130fa9aa0af2
func (a *StrArray) X遍历写锁定(回调函数 func(array []string)) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	回调函数(a.array)
	return a
}

// X遍历读锁定 通过回调函数 `f` 实现读取锁定。 md5:a45deee1e6f17c88
func (a *StrArray) X遍历读锁定(回调函数 func(array []string)) *StrArray {
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
func (a *StrArray) X合并(切片 interface{}) *StrArray {
	return a.Append别名(gconv.X取文本切片(切片)...)
}

// X填充 使用`value`值填充数组，从`startIndex`参数开始的num个条目。
// md5:0a7d3daa806b72ca
func (a *StrArray) X填充(起点 int, 填充数量 int, 值 string) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if 起点 < 0 || 起点 > len(a.array) {
		return gerror.X创建错误码并格式化(gcode.CodeInvalidParameter, "index %d out of array range %d", 起点, len(a.array))
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

// X分割 将一个数组分割成多个子数组，每个子数组的大小由 `size` 决定。最后一个子数组可能包含少于 `size` 个元素。
// md5:0f1f74ff34633d24
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

// X填满 用`value`将数组填充到指定的长度。
// 如果大小为正数，则在右侧填充数组，如果为负数，则在左侧填充。
// 如果`size`的绝对值小于或等于数组的长度，则不进行填充。
// md5:fbe08b371c540418
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

// X取值随机 随机从数组中返回一个元素（不进行删除）。 md5:e152d2c5bc15ecd7
func (a *StrArray) X取值随机() (值 string, 成功 bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if len(a.array) == 0 {
		return "", false
	}
	return a.array[grand.X整数(len(a.array))], true
}

// X取值随机多个 随机从数组中返回 `size` 个元素（不删除）。 md5:09ad7802f8190e3c
func (a *StrArray) X取值随机多个(数量 int) []string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if 数量 <= 0 || len(a.array) == 0 {
		return nil
	}
	array := make([]string, 数量)
	for i := 0; i < 数量; i++ {
		array[i] = a.array[grand.X整数(len(a.array))]
	}
	return array
}

// 随机打乱数组。 md5:5897797461d9f11a
func (a *StrArray) X随机排序() *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, v := range grand.X整数切片(len(a.array)) {
		a.array[i], a.array[v] = a.array[v], a.array[i]
	}
	return a
}

// X倒排序 函数将数组元素反转顺序。 md5:cc34cd0a2fa08e1c
func (a *StrArray) X倒排序() *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, j := 0, len(a.array)-1; i < j; i, j = i+1, j-1 {
		a.array[i], a.array[j] = a.array[j], a.array[i]
	}
	return a
}

// X连接 使用字符串 `glue` 连接数组元素。 md5:ec3894b049af1251
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

// X统计 计算数组中所有值出现的次数。 md5:95b4772dcb002365
func (a *StrArray) X统计() map[string]int {
	m := make(map[string]int)
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		m[v]++
	}
	return m
}

// X遍历 是 IteratorAsc 的别名。 md5:1bfdea306db62845
func (a *StrArray) X遍历(f func(k int, v string) bool) {
	a.X遍历升序(f)
}

// X遍历升序 遍历数组，按照给定的回调函数 `f` 以升序进行只读访问。如果 `f` 返回 true，则继续遍历；否则停止。
// md5:8a125e2dd8982d48
func (a *StrArray) X遍历升序(回调函数 func(k int, v string) bool) {
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
func (a *StrArray) X遍历降序(回调函数 func(k int, v string) bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for i := len(a.array) - 1; i >= 0; i-- {
		if !回调函数(i, a.array[i]) {
			break
		}
	}
}

// String 将当前数组转换为字符串，其实现方式类似于 json.Marshal。 md5:feda8f29233cde8d
func (a *StrArray) String() string {
	if a == nil {
		return ""
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	buffer := bytes.NewBuffer(nil)
	buffer.WriteByte('[')
	for k, v := range a.array {
		buffer.WriteString(`"` + gstr.X转义并按字符(v, `"\`) + `"`)
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
func (a StrArray) MarshalJSON() ([]byte, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return json.Marshal(a.array)
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
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

// UnmarshalValue 是一个接口实现，用于为数组设置任何类型的数据值。 md5:35211e747ab939ab
func (a *StrArray) UnmarshalValue(value interface{}) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	switch value.(type) {
	case string, []byte:
		return json.UnmarshalUseNumber(gconv.X取字节集(value), &a.array)
	default:
		a.array = gconv.SliceStr别名(value)
	}
	return nil
}

// X遍历删除 遍历数组，并使用自定义回调函数过滤元素。
// 如果回调函数`filter`返回true，它将从数组中移除该元素，否则不做任何操作并继续遍历。
// md5:d33873cfb9f1bb38
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

// X删除所有空值 函数移除数组中的所有空字符串值。 md5:2b2e8cd6c844936a
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

// X遍历修改 将用户提供的函数 `f` 应用到数组的每个元素上。 md5:51e35ea7c2c6525c
func (a *StrArray) X遍历修改(回调函数 func(value string) string) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, v := range a.array {
		a.array[i] = 回调函数(v)
	}
	return a
}

// X是否为空 检查数组是否为空。 md5:fb6684351506a02d
func (a *StrArray) X是否为空() bool {
	return a.X取长度() == 0
}

// DeepCopy实现当前类型的深拷贝接口。 md5:9cfbcb08109f6ce1
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


// zj:
func (a *StrArray) X取文本() string {
	return a.String()
}

//zj: