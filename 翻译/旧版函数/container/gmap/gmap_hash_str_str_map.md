
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
) X设置值(key string, val
值
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
) Search(key string) (value string, found
成功
# <翻译结束>

# <翻译开始>
) Search(key string) (value
值
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
) Get(key string) (value
值
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
) Pop() (key, value
值
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
) GetOrSet(key string, value
值
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
) SetIfNotExist(key string, value
值
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
) Remove(key string) (value
被删除值
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
) Diff(other *StrStrMap) (addedKeys, removedKeys, updatedKeys
更新数据的名称
# <翻译结束>

# <翻译开始>
) Diff(other *StrStrMap) (addedKeys, removedKeys
删除的名称
# <翻译结束>

# <翻译开始>
) Diff(other *StrStrMap) (addedKeys
增加的名称
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
func NewStrStrMap(safe
并发安全
# <翻译结束>

# <翻译开始>
func NewStrStrMap
X创建StrStr
# <翻译结束>

# <翻译开始>
func NewStrStrMapFrom(data map[string]string, safe
并发安全
# <翻译结束>

# <翻译开始>
func NewStrStrMapFrom(data
map值
# <翻译结束>

# <翻译开始>
func NewStrStrMapFrom
X创建StrStr并从Map
# <翻译结束>

# <翻译开始>
) X取值或设置值_函数(名称 string, f
回调函数
# <翻译结束>

# <翻译开始>
) X取值或设置值_函数带锁(名称 string, f
回调函数
# <翻译结束>

# <翻译开始>
) X设置值并跳过已存在_函数(名称 string, f
回调函数
# <翻译结束>

# <翻译开始>
) X设置值并跳过已存在_函数带锁(名称 string, f
回调函数
# <翻译结束>
