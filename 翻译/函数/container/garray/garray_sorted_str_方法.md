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

[func NewSortedStrArray(safe ...bool) *SortedStrArray {]
ff=创建文本排序
safe=并发安全

[func NewSortedStrArrayComparator(comparator func(a, b string) int, safe ...bool) *SortedStrArray {]
ff=创建文本排序并带排序函数
safe=并发安全
comparator=排序函数

[func NewSortedStrArraySize(cap int, safe ...bool) *SortedStrArray {]
ff=创建文本排序并按大小
safe=并发安全
cap=大小

[func NewSortedStrArrayFrom(array #左中括号##右中括号#string, safe ...bool) *SortedStrArray {]
ff=创建文本排序并从数组
safe=并发安全
array=数组

[func NewSortedStrArrayFromCopy(array #左中括号##右中括号#string, safe ...bool) *SortedStrArray {]
ff=创建文本排序并从数组复制
safe=并发安全
array=数组

[func (a *SortedStrArray) SetArray(array #左中括号##右中括号#string) *SortedStrArray {]
ff=设置数组
array=数组

[func (a *SortedStrArray) At(index int) (value string) {]
ff=取值
value=值
index=索引

[func (a *SortedStrArray) Sort() *SortedStrArray {]
ff=排序递增

[func (a *SortedStrArray) Add(values ...string) *SortedStrArray {]
ff=入栈右
values=值

[func (a *SortedStrArray) Append(values ...string) *SortedStrArray {]
ff=Append别名
values=值

[func (a *SortedStrArray) Get(index int) (value string, found bool) {]
ff=取值2
found=成功
value=值
index=索引

[func (a *SortedStrArray) Remove(index int) (value string, found bool) {]
ff=删除
found=成功
value=值
index=索引

[func (a *SortedStrArray) RemoveValue(value string) bool {]
ff=删除值
value=值

[func (a *SortedStrArray) RemoveValues(values ...string) {]
ff=删除多个值
values=值

[func (a *SortedStrArray) PopLeft() (value string, found bool) {]
ff=出栈左
found=成功
value=值

[func (a *SortedStrArray) PopRight() (value string, found bool) {]
ff=出栈右
found=成功
value=值

[func (a *SortedStrArray) PopRand() (value string, found bool) {]
ff=出栈随机
found=成功
value=值

[func (a *SortedStrArray) PopRands(size int) #左中括号##右中括号#string {]
ff=出栈随机多个
size=数量

[func (a *SortedStrArray) PopLefts(size int) #左中括号##右中括号#string {]
ff=出栈左多个
size=数量

[func (a *SortedStrArray) PopRights(size int) #左中括号##右中括号#string {]
ff=出栈右多个
size=数量

[func (a *SortedStrArray) Range(start int, end ...int) #左中括号##右中括号#string {]
ff=取切片并按范围
end=终点
start=起点

[func (a *SortedStrArray) SubSlice(offset int, length ...int) #左中括号##右中括号#string {]
ff=取切片并按数量
length=数量
offset=起点

[func (a *SortedStrArray) Sum() (sum int) {]
ff=求和
sum=值

[func (a *SortedStrArray) Len() int {]
ff=取长度

[func (a *SortedStrArray) Slice() #左中括号##右中括号#string {]
ff=取切片

[func (a *SortedStrArray) Interfaces() #左中括号##右中括号#interface{} {]
ff=取any数组

[func (a *SortedStrArray) Contains(value string) bool {]
ff=是否存在
value=值

[func (a *SortedStrArray) ContainsI(value string) bool {]
ff=是否存在并忽略大小写
value=值

[func (a *SortedStrArray) Search(value string) (index int) {]
ff=查找
index=索引
value=值

[func (a *SortedStrArray) SetUnique(unique bool) *SortedStrArray {]
ff=设置去重
unique=去重

[func (a *SortedStrArray) Unique() *SortedStrArray {]
ff=去重

[func (a *SortedStrArray) Clone() (newArray *SortedStrArray) {]
ff=取副本
newArray=新数组

[func (a *SortedStrArray) Clear() *SortedStrArray {]
ff=清空

[func (a *SortedStrArray) LockFunc(f func(array #左中括号##右中括号#string)) *SortedStrArray {]
ff=遍历写锁定
f=回调函数

[func (a *SortedStrArray) RLockFunc(f func(array #左中括号##右中括号#string)) *SortedStrArray {]
ff=遍历读锁定
f=回调函数

[func (a *SortedStrArray) Merge(array interface{}) *SortedStrArray {]
ff=合并
array=数组

[func (a *SortedStrArray) Chunk(size int) #左中括号##右中括号##左中括号##右中括号#string {]
ff=分割
size=数量

[func (a *SortedStrArray) Rand() (value string, found bool) {]
ff=取值随机
found=成功
value=值

[func (a *SortedStrArray) Rands(size int) #左中括号##右中括号#string {]
ff=取值随机多个
size=数量

[func (a *SortedStrArray) Join(glue string) string {]
ff=连接
glue=连接符

[func (a *SortedStrArray) CountValues() map#左中括号#string#右中括号#int {]
ff=统计

[func (a *SortedStrArray) Iterator(f func(k int, v string) bool) {]
ff=X遍历

[func (a *SortedStrArray) IteratorAsc(f func(k int, v string) bool) {]
ff=遍历升序
f=回调函数

[func (a *SortedStrArray) IteratorDesc(f func(k int, v string) bool) {]
ff=遍历降序
f=回调函数

[func (a *SortedStrArray) Filter(filter func(index int, value string) bool) *SortedStrArray {]
ff=遍历删除
filter=回调函数
value=值
index=索引

[func (a *SortedStrArray) FilterEmpty() *SortedStrArray {]
ff=删除所有空值

[func (a *SortedStrArray) Walk(f func(value string) string) *SortedStrArray {]
ff=遍历修改
f=回调函数

[func (a *SortedStrArray) IsEmpty() bool {]
ff=是否为空
