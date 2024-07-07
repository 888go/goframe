# 备注开始
# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 如://ff:取文本
#
# yx=true,此方法优先翻译
# 如: //yx=true

# **_package.md 文件备注:
# bm= 包名,更换新的包名称 
# 如: package gin //bm:gin类

# **_其他.md 文件备注:
# qm= 前面,跳转到前面进行重命名.文档内如果有多个相同的,会一起重命名.
# hm= 后面,跳转到后面进行重命名.文档内如果有多个相同的,会一起重命名.
# cz= 查找,配合前面/后面使用,
# zz= 正则查找,配合前面/后面使用, 有设置正则查找,就不用设置上面的查找
# 如: type Regexp struct {//qm:正则 cz:Regexp struct
#
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# 如:
# type Regexp struct {//th:type Regexp222 struct
#
# cf= 重复,用于重命名多次,
# 如: 
# 一个文档内有2个"One(result interface{}) error"需要重命名.
# 但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

# **_追加.md 文件备注:
# 在代码内追加代码,如:
# //zj:前面一行的代码,如果为空,追加到末尾行
# func (re *Regexp) X取文本() string { 
# re.F.String()
# }
# //zj:
# 备注结束

[func NewSortedIntArray(safe ...bool) *SortedIntArray {]
ff=创建整数排序
safe=并发安全

[func NewSortedIntArrayComparator(comparator func(a, b int) int, safe ...bool) *SortedIntArray {]
ff=创建整数排序并带排序函数
safe=并发安全
comparator=排序函数

[func NewSortedIntArraySize(cap int, safe ...bool) *SortedIntArray {]
ff=创建整数排序并按大小
safe=并发安全
cap=大小

[func NewSortedIntArrayRange(start, end, step int, safe ...bool) *SortedIntArray {]
ff=创建整数排序并按范围
safe=并发安全
step=步长
end=终点
start=起点

[func NewSortedIntArrayFrom(array #左中括号##右中括号#int, safe ...bool) *SortedIntArray {]
ff=创建整数排序并从切片
safe=并发安全
array=切片

[func NewSortedIntArrayFromCopy(array #左中括号##右中括号#int, safe ...bool) *SortedIntArray {]
ff=创建整数排序并从切片复制
safe=并发安全
array=切片

[func (a *SortedIntArray) At(index int) (value int) {]
ff=取值
value=值
index=索引

[func (a *SortedIntArray) SetArray(array #左中括号##右中括号#int) *SortedIntArray {]
ff=设置切片
array=切片

[func (a *SortedIntArray) Sort() *SortedIntArray {]
ff=排序递增

[func (a *SortedIntArray) Add(values ...int) *SortedIntArray {]
ff=入栈右
values=值

[func (a *SortedIntArray) Append(values ...int) *SortedIntArray {]
ff=Append别名
values=值

[func (a *SortedIntArray) Get(index int) (value int, found bool) {]
ff=取值2
found=成功
value=值
index=索引

[func (a *SortedIntArray) Remove(index int) (value int, found bool) {]
ff=删除
found=成功
value=值
index=索引

[func (a *SortedIntArray) RemoveValue(value int) bool {]
ff=删除值
value=值

[func (a *SortedIntArray) RemoveValues(values ...int) {]
ff=删除多个值
values=值

[func (a *SortedIntArray) PopLeft() (value int, found bool) {]
ff=出栈左
found=成功
value=值

[func (a *SortedIntArray) PopRight() (value int, found bool) {]
ff=出栈右
found=成功
value=值

[func (a *SortedIntArray) PopRand() (value int, found bool) {]
ff=出栈随机
found=成功
value=值

[func (a *SortedIntArray) PopRands(size int) #左中括号##右中括号#int {]
ff=出栈随机多个
size=数量

[func (a *SortedIntArray) PopLefts(size int) #左中括号##右中括号#int {]
ff=出栈左多个
size=数量

[func (a *SortedIntArray) PopRights(size int) #左中括号##右中括号#int {]
ff=出栈右多个
size=数量

[func (a *SortedIntArray) Range(start int, end ...int) #左中括号##右中括号#int {]
ff=取切片并按范围
end=终点
start=起点

[func (a *SortedIntArray) SubSlice(offset int, length ...int) #左中括号##右中括号#int {]
ff=取切片并按数量
length=数量
offset=起点

[func (a *SortedIntArray) Len() int {]
ff=取长度

[func (a *SortedIntArray) Sum() (sum int) {]
ff=求和
sum=值

[func (a *SortedIntArray) Slice() #左中括号##右中括号#int {]
ff=取切片

[func (a *SortedIntArray) Interfaces() #左中括号##右中括号#interface{} {]
ff=取any切片
yx=true

[func (a *SortedIntArray) Contains(value int) bool {]
ff=是否存在
value=值

[func (a *SortedIntArray) Search(value int) (index int) {]
ff=查找
index=索引
value=值

[func (a *SortedIntArray) SetUnique(unique bool) *SortedIntArray {]
ff=设置去重
unique=去重

[func (a *SortedIntArray) Unique() *SortedIntArray {]
ff=去重

[func (a *SortedIntArray) Clone() (newArray *SortedIntArray) {]
ff=取副本
newArray=新切片

[func (a *SortedIntArray) Clear() *SortedIntArray {]
ff=清空

[func (a *SortedIntArray) LockFunc(f func(array #左中括号##右中括号#int)) *SortedIntArray {]
ff=遍历写锁定
f=回调函数

[func (a *SortedIntArray) RLockFunc(f func(array #左中括号##右中括号#int)) *SortedIntArray {]
ff=遍历读锁定
f=回调函数

[func (a *SortedIntArray) Merge(array interface{}) *SortedIntArray {]
ff=合并
array=切片

[func (a *SortedIntArray) Chunk(size int) #左中括号##右中括号##左中括号##右中括号#int {]
ff=分割
size=数量

[func (a *SortedIntArray) Rand() (value int, found bool) {]
ff=取值随机
found=成功
value=值

[func (a *SortedIntArray) Rands(size int) #左中括号##右中括号#int {]
ff=取值随机多个
size=数量

[func (a *SortedIntArray) Join(glue string) string {]
ff=连接
glue=连接符

[func (a *SortedIntArray) CountValues() map#左中括号#int#右中括号#int {]
ff=统计

[func (a *SortedIntArray) Iterator(f func(k int, v int) bool) {]
ff=X遍历
yx=true

[func (a *SortedIntArray) IteratorAsc(f func(k int, v int) bool) {]
ff=遍历升序
f=回调函数

[func (a *SortedIntArray) IteratorDesc(f func(k int, v int) bool) {]
ff=遍历降序
f=回调函数

[func (a *SortedIntArray) Filter(filter func(index int, value int) bool) *SortedIntArray {]
ff=遍历删除
filter=回调函数
value=值
index=索引

[func (a *SortedIntArray) FilterEmpty() *SortedIntArray {]
ff=删除所有空值

[func (a *SortedIntArray) Walk(f func(value int) int) *SortedIntArray {]
ff=遍历修改
f=回调函数

[func (a *SortedIntArray) IsEmpty() bool {]
ff=是否为空
