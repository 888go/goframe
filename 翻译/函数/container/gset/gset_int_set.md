
# <翻译开始>
func NewIntSet(safe
并发安全
# <翻译结束>

# <翻译开始>
func NewIntSet
X创建整数
# <翻译结束>

# <翻译开始>
func NewIntSetFrom(items []int, safe
并发安全
# <翻译结束>

# <翻译开始>
func NewIntSetFrom(items
整数数组
# <翻译结束>

# <翻译开始>
func NewIntSetFrom
X创建整数并按值
# <翻译结束>

# <翻译开始>
) Iterator
X遍历
# <翻译结束>

# <翻译开始>
) Add(item
值s
# <翻译结束>

# <翻译开始>
) Add
X加入
# <翻译结束>

# <翻译开始>
) AddIfNotExist(item
值
# <翻译结束>

# <翻译开始>
) AddIfNotExist
X加入值并跳过已存在
# <翻译结束>

# <翻译开始>
) AddIfNotExistFunc(item
值
# <翻译结束>

# <翻译开始>
) AddIfNotExistFunc
X加入值并跳过已存在_函数
# <翻译结束>

# <翻译开始>
) AddIfNotExistFuncLock(item
值
# <翻译结束>

# <翻译开始>
) AddIfNotExistFuncLock
X加入值并跳过已存在_并发安全函数
# <翻译结束>

# <翻译开始>
) Contains(item
值
# <翻译结束>

# <翻译开始>
) Contains
X是否存在
# <翻译结束>

# <翻译开始>
) Remove(item
值
# <翻译结束>

# <翻译开始>
) Remove
X删除
# <翻译结束>

# <翻译开始>
) Size
X取数量
# <翻译结束>

# <翻译开始>
) Clear
X清空
# <翻译结束>

# <翻译开始>
) Slice
X取集合数组
# <翻译结束>

# <翻译开始>
) Join(glue
连接符
# <翻译结束>

# <翻译开始>
) Join
X取集合文本
# <翻译结束>

# <翻译开始>
) LockFunc
X写锁定_函数
# <翻译结束>

# <翻译开始>
) RLockFunc
X读锁定_函数
# <翻译结束>

# <翻译开始>
) Equal(other
待比较集合
# <翻译结束>

# <翻译开始>
) Equal
X是否相等
# <翻译结束>

# <翻译开始>
) IsSubsetOf(other
父集
# <翻译结束>

# <翻译开始>
) IsSubsetOf
X是否为子集
# <翻译结束>

# <翻译开始>
) Union(others ...*IntSet) (newSet
新集合
# <翻译结束>

# <翻译开始>
) Union(others
集合
# <翻译结束>

# <翻译开始>
) Union
X取并集
# <翻译结束>

# <翻译开始>
) Diff(others ...*IntSet) (newSet
新集合
# <翻译结束>

# <翻译开始>
) Diff(others
集合
# <翻译结束>

# <翻译开始>
) Diff
X取差集
# <翻译结束>

# <翻译开始>
) Intersect(others ...*IntSet) (newSet
新集合
# <翻译结束>

# <翻译开始>
) Intersect(others
集合
# <翻译结束>

# <翻译开始>
) Intersect
X取交集
# <翻译结束>

# <翻译开始>
) Complement(full *IntSet) (newSet
新集合
# <翻译结束>

# <翻译开始>
) Complement(full
集合
# <翻译结束>

# <翻译开始>
) Complement
X取补集
# <翻译结束>

# <翻译开始>
) Merge(others
集合s
# <翻译结束>

# <翻译开始>
) Merge
X合并
# <翻译结束>

# <翻译开始>
) Sum() (sum
总和
# <翻译结束>

# <翻译开始>
) Sum
X求和
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
) Walk(f func(item
值
# <翻译结束>

# <翻译开始>
) Walk
X遍历修改
# <翻译结束>
