# 备注开始
# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 如:
# //ff:取文本

# **_package.md 文件备注:
# bm= 包名,更换新的包名称 
# 如: 
# package gin //bm:gin类

# **_其他.md 文件备注:
# qm= 前面,跳转到前面进行重命名.文档内如果有多个相同的,会一起重命名.
# hm= 后面,跳转到后面进行重命名.文档内如果有多个相同的,会一起重命名.
# cz= 查找,配合前面/后面使用,
# 如:
# type Regexp struct {//qm:正则 cz:Regexp struct
#
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# type Regexp struct {//th:type Regexp222 struct
#
# cf= 重复,用于重命名多次,
# 如: 
# 一个文档内有2个"One(result interface{}) error"需要重命名.
# 但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

# **_追加.md 文件备注:
# 在代码内追加代码,如:
# //zj:
# func (re *Regexp) X取文本() string { 
#    re.F.String()
# }
# //zj:
# 备注结束

[func NewSortedArray(comparator func(a, b interface{}) int, safe ...bool) *SortedArray {]
ff=创建排序
safe=并发安全
comparator=排序函数

[func NewSortedArraySize(cap int, comparator func(a, b interface{}) int, safe ...bool) *SortedArray {]
ff=创建排序并按大小
safe=并发安全
comparator=排序函数
cap=大小

[func NewSortedArrayRange(start, end, step int, comparator func(a, b interface{}) int, safe ...bool) *SortedArray {]
ff=创建排序并按范围
safe=并发安全
comparator=排序函数
step=步长
end=终点
start=起点

[func NewSortedArrayFrom(array #左中括号##右中括号#interface{}, comparator func(a, b interface{}) int, safe ...bool) *SortedArray {]
ff=创建排序并从数组
safe=并发安全
comparator=排序函数
array=数组

[func NewSortedArrayFromCopy(array #左中括号##右中括号#interface{}, comparator func(a, b interface{}) int, safe ...bool) *SortedArray {]
ff=创建排序并从数组复制
safe=并发安全
comparator=排序函数
array=数组

[func (a *SortedArray) At(index int) (value interface{}) {]
ff=取值
value=值
index=索引

[func (a *SortedArray) SetArray(array #左中括号##右中括号#interface{}) *SortedArray {]
ff=设置数组
array=数组

[func (a *SortedArray) SetComparator(comparator func(a, b interface{}) int) {]
ff=设置排序函数
comparator=排序函数

[func (a *SortedArray) Sort() *SortedArray {]
ff=排序递增

[func (a *SortedArray) Add(values ...interface{}) *SortedArray {]
ff=入栈右
values=值

[func (a *SortedArray) Append(values ...interface{}) *SortedArray {]
ff=Append别名
values=值

[func (a *SortedArray) Get(index int) (value interface{}, found bool) {]
ff=取值2
found=成功
value=值
index=索引

[func (a *SortedArray) Remove(index int) (value interface{}, found bool) {]
ff=删除
found=成功
value=值
index=索引

[func (a *SortedArray) RemoveValue(value interface{}) bool {]
ff=删除值
value=值

[func (a *SortedArray) RemoveValues(values ...interface{}) {]
ff=删除多个值
values=值

[func (a *SortedArray) PopLeft() (value interface{}, found bool) {]
ff=出栈左
found=成功
value=值

[func (a *SortedArray) PopRight() (value interface{}, found bool) {]
ff=出栈右
found=成功
value=值

[func (a *SortedArray) PopRand() (value interface{}, found bool) {]
ff=出栈随机
found=成功
value=值

[func (a *SortedArray) PopRands(size int) #左中括号##右中括号#interface{} {]
ff=出栈随机多个
size=数量

[func (a *SortedArray) PopLefts(size int) #左中括号##右中括号#interface{} {]
ff=出栈左多个
size=数量

[func (a *SortedArray) PopRights(size int) #左中括号##右中括号#interface{} {]
ff=出栈右多个
size=数量

[func (a *SortedArray) Range(start int, end ...int) #左中括号##右中括号#interface{} {]
ff=取切片并按范围
end=终点
start=起点

[func (a *SortedArray) SubSlice(offset int, length ...int) #左中括号##右中括号#interface{} {]
ff=取切片并按数量
length=数量
offset=起点

[func (a *SortedArray) Sum() (sum int) {]
ff=求和
sum=值

[func (a *SortedArray) Len() int {]
ff=取长度

[func (a *SortedArray) Slice() #左中括号##右中括号#interface{} {]
ff=取切片

[func (a *SortedArray) Interfaces() #左中括号##右中括号#interface{} {]
ff=取any数组

[func (a *SortedArray) Contains(value interface{}) bool {]
ff=是否存在
value=值

[func (a *SortedArray) Search(value interface{}) (index int) {]
ff=查找
index=索引
value=值

[func (a *SortedArray) SetUnique(unique bool) *SortedArray {]
ff=设置去重
unique=去重

[func (a *SortedArray) Unique() *SortedArray {]
ff=去重

[func (a *SortedArray) Clone() (newArray *SortedArray) {]
ff=取副本
newArray=新数组

[func (a *SortedArray) Clear() *SortedArray {]
ff=清空

[func (a *SortedArray) LockFunc(f func(array #左中括号##右中括号#interface{})) *SortedArray {]
ff=遍历写锁定
f=回调函数

[func (a *SortedArray) RLockFunc(f func(array #左中括号##右中括号#interface{})) *SortedArray {]
ff=遍历读锁定
f=回调函数

[func (a *SortedArray) Merge(array interface{}) *SortedArray {]
ff=合并
array=数组

[func (a *SortedArray) Chunk(size int) #左中括号##右中括号##左中括号##右中括号#interface{} {]
ff=分割
size=数量

[func (a *SortedArray) Rand() (value interface{}, found bool) {]
ff=取值随机
found=成功
value=值

[func (a *SortedArray) Rands(size int) #左中括号##右中括号#interface{} {]
ff=取值随机多个
size=数量

[func (a *SortedArray) Join(glue string) string {]
ff=连接
glue=连接符

[func (a *SortedArray) CountValues() map#左中括号#interface{}#右中括号#int {]
ff=统计

[func (a *SortedArray) Iterator(f func(k int, v interface{}) bool) {]
ff=X遍历

[func (a *SortedArray) IteratorAsc(f func(k int, v interface{}) bool) {]
ff=遍历升序
f=回调函数

[func (a *SortedArray) IteratorDesc(f func(k int, v interface{}) bool) {]
ff=遍历降序
f=回调函数

[func (a *SortedArray) FilterNil() *SortedArray {]
ff=删除所有nil

[func (a *SortedArray) Filter(filter func(index int, value interface{}) bool) *SortedArray {]
ff=遍历删除
filter=回调函数
value=值
index=索引

[func (a *SortedArray) FilterEmpty() *SortedArray {]
ff=删除所有空值

[func (a *SortedArray) Walk(f func(value interface{}) interface{}) *SortedArray {]
ff=遍历修改
f=回调函数

[func (a *SortedArray) IsEmpty() bool {]
ff=是否为空
