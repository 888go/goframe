
# <翻译开始>
func NewSortedArray(comparator func(a, b interface{}) int, safe
并发安全
# <翻译结束>

# <翻译开始>
func NewSortedArray(comparator
排序函数
# <翻译结束>

# <翻译开始>
func NewSortedArray
X创建排序
# <翻译结束>

# <翻译开始>
func NewSortedArraySize(cap int, comparator func(a, b interface{}) int, safe
并发安全
# <翻译结束>

# <翻译开始>
func NewSortedArraySize(cap int, comparator
排序函数
# <翻译结束>

# <翻译开始>
func NewSortedArraySize(cap
大小
# <翻译结束>

# <翻译开始>
func NewSortedArraySize
X创建排序并按大小
# <翻译结束>

# <翻译开始>
func NewSortedArrayRange(start, end, step int, comparator func(a, b interface{}) int, safe
并发安全
# <翻译结束>

# <翻译开始>
func NewSortedArrayRange(start, end, step int, comparator
排序函数
# <翻译结束>

# <翻译开始>
func NewSortedArrayRange(start, end, step
步长
# <翻译结束>

# <翻译开始>
func NewSortedArrayRange(start, end
终点
# <翻译结束>

# <翻译开始>
func NewSortedArrayRange(start
起点
# <翻译结束>

# <翻译开始>
func NewSortedArrayRange
X创建排序并按范围
# <翻译结束>

# <翻译开始>
func NewSortedArrayFrom(array []interface{}, comparator func(a, b interface{}) int, safe
并发安全
# <翻译结束>

# <翻译开始>
func NewSortedArrayFrom(array []interface{}, comparator
排序函数
# <翻译结束>

# <翻译开始>
func NewSortedArrayFrom(array
数组
# <翻译结束>

# <翻译开始>
func NewSortedArrayFrom
X创建排序并从数组
# <翻译结束>

# <翻译开始>
func NewSortedArrayFromCopy(array []interface{}, comparator func(a, b interface{}) int, safe
并发安全
# <翻译结束>

# <翻译开始>
func NewSortedArrayFromCopy(array []interface{}, comparator
排序函数
# <翻译结束>

# <翻译开始>
func NewSortedArrayFromCopy(array
数组
# <翻译结束>

# <翻译开始>
func NewSortedArrayFromCopy
X创建排序并从数组复制
# <翻译结束>

# <翻译开始>
) At(index int) (value
值
# <翻译结束>

# <翻译开始>
) At(index
索引
# <翻译结束>

# <翻译开始>
) At
X取值
# <翻译结束>

# <翻译开始>
) SetArray(array
数组
# <翻译结束>

# <翻译开始>
) SetArray
X设置数组
# <翻译结束>

# <翻译开始>
) Append(values
值
# <翻译结束>

# <翻译开始>
) Append
Append别名
# <翻译结束>

# <翻译开始>
) Get(index int) (value interface{}, found
成功
# <翻译结束>

# <翻译开始>
) Get(index int) (value
值
# <翻译结束>

# <翻译开始>
) Get(index
索引
# <翻译结束>

# <翻译开始>
) Get
X取值2
# <翻译结束>

# <翻译开始>
) Remove(index int) (value interface{}, found
成功
# <翻译结束>

# <翻译开始>
) Remove(index int) (value
值
# <翻译结束>

# <翻译开始>
) Remove(index
索引
# <翻译结束>

# <翻译开始>
) Remove
X删除
# <翻译结束>

# <翻译开始>
) RemoveValue(value
值
# <翻译结束>

# <翻译开始>
) RemoveValue
X删除值
# <翻译结束>

# <翻译开始>
) RemoveValues(values
值
# <翻译结束>

# <翻译开始>
) RemoveValues
X删除多个值
# <翻译结束>

# <翻译开始>
) PopLeft() (value interface{}, found
成功
# <翻译结束>

# <翻译开始>
) PopLeft() (value
值
# <翻译结束>

# <翻译开始>
) PopLeft
X出栈左
# <翻译结束>

# <翻译开始>
) PopRight() (value interface{}, found
成功
# <翻译结束>

# <翻译开始>
) PopRight() (value
值
# <翻译结束>

# <翻译开始>
) PopRight
X出栈右
# <翻译结束>

# <翻译开始>
) PopRand() (value interface{}, found
成功
# <翻译结束>

# <翻译开始>
) PopRand() (value
值
# <翻译结束>

# <翻译开始>
) PopRand
X出栈随机
# <翻译结束>

# <翻译开始>
) PopRands(size
数量
# <翻译结束>

# <翻译开始>
) PopRands
X出栈随机多个
# <翻译结束>

# <翻译开始>
) PopLefts(size
数量
# <翻译结束>

# <翻译开始>
) PopLefts
X出栈左多个
# <翻译结束>

# <翻译开始>
) PopRights(size
数量
# <翻译结束>

# <翻译开始>
) PopRights
X出栈右多个
# <翻译结束>

# <翻译开始>
) Range(start int, end
终点
# <翻译结束>

# <翻译开始>
) Range(start
起点
# <翻译结束>

# <翻译开始>
) Range
X取切片并按范围
# <翻译结束>

# <翻译开始>
) SubSlice(offset int, length
数量
# <翻译结束>

# <翻译开始>
) SubSlice(offset
起点
# <翻译结束>

# <翻译开始>
) SubSlice
X取切片并按数量
# <翻译结束>

# <翻译开始>
) Sum() (sum
值
# <翻译结束>

# <翻译开始>
) Sum
X求和
# <翻译结束>

# <翻译开始>
) Len
X取长度
# <翻译结束>

# <翻译开始>
) Slice
X取切片
# <翻译结束>

# <翻译开始>
) Contains(value
值
# <翻译结束>

# <翻译开始>
) Contains
X是否存在
# <翻译结束>

# <翻译开始>
) Search(value interface{}) (index
索引
# <翻译结束>

# <翻译开始>
) Search(value
值
# <翻译结束>

# <翻译开始>
) Search
X查找
# <翻译结束>

# <翻译开始>
) Unique
X去重
# <翻译结束>

# <翻译开始>
) Clone() (newArray
新数组
# <翻译结束>

# <翻译开始>
) Clone
X取副本
# <翻译结束>

# <翻译开始>
) Clear
X清空
# <翻译结束>

# <翻译开始>
) LockFunc(f func(array
数组
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
) RLockFunc(f func(array
数组
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
) Merge(array
数组
# <翻译结束>

# <翻译开始>
) Merge
X合并
# <翻译结束>

# <翻译开始>
) Chunk(size
数量
# <翻译结束>

# <翻译开始>
) Chunk
X分割
# <翻译结束>

# <翻译开始>
) Rand() (value interface{}, found
成功
# <翻译结束>

# <翻译开始>
) Rand() (value
值
# <翻译结束>

# <翻译开始>
) Rand
X取值随机
# <翻译结束>

# <翻译开始>
) Rands(size
数量
# <翻译结束>

# <翻译开始>
) Rands
X取值随机多个
# <翻译结束>

# <翻译开始>
) Join(glue
连接符
# <翻译结束>

# <翻译开始>
) Join
X连接
# <翻译结束>

# <翻译开始>
) CountValues
X统计
# <翻译结束>

# <翻译开始>
) X遍历(f
回调函数
# <翻译结束>

# <翻译开始>
) IteratorAsc(f
回调函数
# <翻译结束>

# <翻译开始>
) IteratorAsc
X遍历升序
# <翻译结束>

# <翻译开始>
) IteratorDesc(f
回调函数
# <翻译结束>

# <翻译开始>
) IteratorDesc
X遍历降序
# <翻译结束>

# <翻译开始>
) FilterNil
X删除所有nil
# <翻译结束>

# <翻译开始>
) Filter(filter func(index int, value
值
# <翻译结束>

# <翻译开始>
) Filter(filter func(index
索引
# <翻译结束>

# <翻译开始>
) Filter(filter
回调函数
# <翻译结束>

# <翻译开始>
) Filter
X遍历删除
# <翻译结束>

# <翻译开始>
) FilterEmpty
X删除所有空值
# <翻译结束>

# <翻译开始>
) Walk(f func(value
值
# <翻译结束>

# <翻译开始>
) Walk(f
回调函数
# <翻译结束>

# <翻译开始>
) Walk
X遍历修改
# <翻译结束>

# <翻译开始>
) IsEmpty
X是否为空
# <翻译结束>

# <翻译开始>
) Add(values
值
# <翻译结束>

# <翻译开始>
) Add
X入栈右
# <翻译结束>

# <翻译开始>
) SetUnique(unique
去重
# <翻译结束>

# <翻译开始>
) SetUnique
X设置去重
# <翻译结束>

# <翻译开始>
) SetComparator(comparator
排序函数
# <翻译结束>

# <翻译开始>
) SetComparator
X设置排序函数
# <翻译结束>

# <翻译开始>
) Sort
X排序递增
# <翻译结束>

# <追加函数开始>
func (a *SortedArray) X取文本() string {
return a.String()
}
# <追加函数结束>
