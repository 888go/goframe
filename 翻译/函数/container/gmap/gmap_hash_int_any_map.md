
# <翻译开始>
) X遍历(f
回调函数
# <翻译结束>

# <翻译开始>
) Clone
X取副本
# <翻译结束>

# <翻译开始>
) Map
X取Map
# <翻译结束>

# <翻译开始>
) MapCopy
X浅拷贝
# <翻译结束>

# <翻译开始>
) FilterEmpty
X删除所有空值
# <翻译结束>

# <翻译开始>
) FilterNil
X删除所有nil值
# <翻译结束>

# <翻译开始>
) X设置值(key
名称
# <翻译结束>

# <翻译开始>
) Sets(data
map值
# <翻译结束>

# <翻译开始>
) Sets
X设置值Map
# <翻译结束>

# <翻译开始>
) Search(key
名称
# <翻译结束>

# <翻译开始>
) Search
X查找
# <翻译结束>

# <翻译开始>
) Get(key
名称
# <翻译结束>

# <翻译开始>
) Get
X取值
# <翻译结束>

# <翻译开始>
) Pop() (key
名称
# <翻译结束>

# <翻译开始>
) Pop
X出栈
# <翻译结束>

# <翻译开始>
) Pops(size
数量
# <翻译结束>

# <翻译开始>
) Pops
X出栈多个
# <翻译结束>

# <翻译开始>
) GetOrSet(key
名称
# <翻译结束>

# <翻译开始>
) GetOrSet
X取值或设置值
# <翻译结束>

# <翻译开始>
) GetOrSetFunc(key
名称
# <翻译结束>

# <翻译开始>
) GetOrSetFunc
X取值或设置值_函数
# <翻译结束>

# <翻译开始>
) GetOrSetFuncLock(key
名称
# <翻译结束>

# <翻译开始>
) GetOrSetFuncLock
X取值或设置值_函数带锁
# <翻译结束>

# <翻译开始>
) GetVar(key
名称
# <翻译结束>

# <翻译开始>
) GetVar
X取值泛型类
# <翻译结束>

# <翻译开始>
) GetVarOrSet(key
名称
# <翻译结束>

# <翻译开始>
) GetVarOrSet
X取值或设置值泛型类
# <翻译结束>

# <翻译开始>
GetVarOrSetFunc(key
名称
# <翻译结束>

# <翻译开始>
) GetVarOrSetFunc
X取值或设置值泛型类_函数
# <翻译结束>

# <翻译开始>
) GetVarOrSetFuncLock(key
名称
# <翻译结束>

# <翻译开始>
) GetVarOrSetFuncLock
X取值或设置值泛型类_函数带锁
# <翻译结束>

# <翻译开始>
) SetIfNotExist(key
名称
# <翻译结束>

# <翻译开始>
) SetIfNotExist
X设置值并跳过已存在
# <翻译结束>

# <翻译开始>
) SetIfNotExistFunc(key
名称
# <翻译结束>

# <翻译开始>
) SetIfNotExistFunc
X设置值并跳过已存在_函数
# <翻译结束>

# <翻译开始>
) SetIfNotExistFuncLock(key
名称
# <翻译结束>

# <翻译开始>
) SetIfNotExistFuncLock
X设置值并跳过已存在_函数带锁
# <翻译结束>

# <翻译开始>
) Remove(key
名称
# <翻译结束>

# <翻译开始>
) Remove
X删除
# <翻译结束>

# <翻译开始>
) Removes(keys
名称
# <翻译结束>

# <翻译开始>
) Removes
X删除多个值
# <翻译结束>

# <翻译开始>
) Keys
X取所有名称
# <翻译结束>

# <翻译开始>
) Values
X取所有值
# <翻译结束>

# <翻译开始>
) Contains(key
名称
# <翻译结束>

# <翻译开始>
) Contains
X是否存在
# <翻译结束>

# <翻译开始>
) Size
X取数量
# <翻译结束>

# <翻译开始>
) IsEmpty
X是否为空
# <翻译结束>

# <翻译开始>
) Clear
X清空
# <翻译结束>

# <翻译开始>
) Replace(data
map值
# <翻译结束>

# <翻译开始>
) Replace
X替换
# <翻译结束>

# <翻译开始>
) LockFunc(f
回调函数
# <翻译结束>

# <翻译开始>
) LockFunc
X遍历写锁定
# <翻译结束>

# <翻译开始>
) RLockFunc(f
回调函数
# <翻译结束>

# <翻译开始>
) RLockFunc
X遍历读锁定
# <翻译结束>

# <翻译开始>
) Flip
X名称值交换
# <翻译结束>

# <翻译开始>
) Merge(other
map值
# <翻译结束>

# <翻译开始>
) Merge
X合并
# <翻译结束>

# <翻译开始>
) IsSubOf(other
父集Map
# <翻译结束>

# <翻译开始>
) IsSubOf
X是否为子集
# <翻译结束>

# <翻译开始>
) Diff(other
map值
# <翻译结束>

# <翻译开始>
) Diff
X比较
# <翻译结束>

# <翻译开始>
func NewIntAnyMap(safe
并发安全
# <翻译结束>

# <翻译开始>
func NewIntAnyMap
X创建IntAny
# <翻译结束>

# <翻译开始>
func NewIntAnyMapFrom(data map[int]interface{}, safe
并发安全
# <翻译结束>

# <翻译开始>
func NewIntAnyMapFrom(data
map值
# <翻译结束>

# <翻译开始>
func NewIntAnyMapFrom
X创建IntAny并从Map
# <翻译结束>

# <翻译开始>
) X设置值(名称 int, val
值
# <翻译结束>

# <翻译开始>
) X查找(名称 int) (value interface{}, found
成功
# <翻译结束>

# <翻译开始>
) X查找(名称 int) (value
值
# <翻译结束>

# <翻译开始>
) X取值(名称 int) (value
值
# <翻译结束>

# <翻译开始>
) X出栈() (名称 int, value
值
# <翻译结束>

# <翻译开始>
) X取值或设置值(名称 int, value
值
# <翻译结束>

# <翻译开始>
) X取值或设置值_函数(名称 int, f
回调函数
# <翻译结束>

# <翻译开始>
) X取值或设置值_函数带锁(名称 int, f
回调函数
# <翻译结束>

# <翻译开始>
) X取值或设置值泛型类(名称 int, value
值
# <翻译结束>

# <翻译开始>
) X取值或设置值泛型类_函数(名称 int, f
回调函数
# <翻译结束>

# <翻译开始>
) X取值或设置值泛型类_函数带锁(名称 int, f
回调函数
# <翻译结束>

# <翻译开始>
) X设置值并跳过已存在(名称 int, value
值
# <翻译结束>

# <翻译开始>
) X设置值并跳过已存在_函数(名称 int, f
回调函数
# <翻译结束>

# <翻译开始>
) X设置值并跳过已存在_函数带锁(名称 int, f
回调函数
# <翻译结束>

# <翻译开始>
) X删除(名称 int) (value
值
# <翻译结束>

# <翻译开始>
) X比较(map值 *IntAnyMap) (addedKeys, removedKeys, updatedKeys
更新数据的名称
# <翻译结束>

# <翻译开始>
) X比较(map值 *IntAnyMap) (addedKeys, removedKeys
删除的名称
# <翻译结束>

# <翻译开始>
) X比较(map值 *IntAnyMap) (addedKeys
增加的名称
# <翻译结束>
