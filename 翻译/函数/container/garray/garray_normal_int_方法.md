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
# //zj:
# func (re *Regexp) X取文本() string { 
#    re.F.String()
# }
# //zj:
# 备注结束

[func NewIntArray(safe ...bool) *IntArray {]
ff=创建整数
safe=并发安全

[func NewIntArraySize(size int, cap int, safe ...bool) *IntArray {]
ff=创建整数并按大小
safe=并发安全
cap=上限
size=大小

[func NewIntArrayRange(start, end, step int, safe ...bool) *IntArray {]
ff=创建整数并按范围
safe=并发安全
step=步长
end=终点
start=起点

[func NewIntArrayFrom(array #左中括号##右中括号#int, safe ...bool) *IntArray {]
ff=创建整数并从切片
safe=并发安全
array=切片

[func NewIntArrayFromCopy(array #左中括号##右中括号#int, safe ...bool) *IntArray {]
ff=创建整数并从切片复制
safe=并发安全
array=切片

[func (a *IntArray) At(index int) (value int) {]
ff=取值
value=值
index=索引

[func (a *IntArray) Get(index int) (value int, found bool) {]
ff=取值2
found=成功
value=值
index=索引

[func (a *IntArray) Set(index int, value int) error {]
ff=设置值
yx=true

[func (a *IntArray) SetArray(array #左中括号##右中括号#int) *IntArray {]
ff=设置切片
array=切片

[func (a *IntArray) Replace(array #左中括号##右中括号#int) *IntArray {]
ff=替换
array=切片

[func (a *IntArray) Sum() (sum int) {]
ff=求和
sum=值

[func (a *IntArray) Sort(reverse ...bool) *IntArray {]
ff=排序递增
reverse=降序

[func (a *IntArray) SortFunc(less func(v1, v2 int) bool) *IntArray {]
ff=排序函数
less=回调函数

[func (a *IntArray) InsertBefore(index int, values ...int) error {]
ff=插入前面
values=值
index=索引

[func (a *IntArray) InsertAfter(index int, values ...int) error {]
ff=插入后面
values=值

[func (a *IntArray) Remove(index int) (value int, found bool) {]
ff=删除
found=成功
value=值
index=索引

[func (a *IntArray) RemoveValue(value int) bool {]
ff=删除值
value=值

[func (a *IntArray) RemoveValues(values ...int) {]
ff=删除多个值
values=值

[func (a *IntArray) PushLeft(value ...int) *IntArray {]
ff=入栈左
value=值

[func (a *IntArray) PushRight(value ...int) *IntArray {]
ff=入栈右
value=值

[func (a *IntArray) PopLeft() (value int, found bool) {]
ff=出栈左
found=成功
value=值

[func (a *IntArray) PopRight() (value int, found bool) {]
ff=出栈右
found=成功
value=值

[func (a *IntArray) PopRand() (value int, found bool) {]
ff=出栈随机
found=成功
value=值

[func (a *IntArray) PopRands(size int) #左中括号##右中括号#int {]
ff=出栈随机多个
size=数量

[func (a *IntArray) PopLefts(size int) #左中括号##右中括号#int {]
ff=出栈左多个
size=数量

[func (a *IntArray) PopRights(size int) #左中括号##右中括号#int {]
ff=出栈右多个
size=数量

[func (a *IntArray) Range(start int, end ...int) #左中括号##右中括号#int {]
ff=取切片并按范围
end=终点
start=起点

[func (a *IntArray) SubSlice(offset int, length ...int) #左中括号##右中括号#int {]
ff=取切片并按数量
length=数量
offset=起点

[func (a *IntArray) Append(value ...int) *IntArray {]
ff=Append别名
value=值

[func (a *IntArray) Len() int {]
ff=取长度

[func (a *IntArray) Slice() #左中括号##右中括号#int {]
ff=取切片

[func (a *IntArray) Interfaces() #左中括号##右中括号#interface{} {]
ff=取any切片
yx=true

[func (a *IntArray) Clone() (newArray *IntArray) {]
ff=取副本
newArray=新切片

[func (a *IntArray) Clear() *IntArray {]
ff=清空

[func (a *IntArray) Contains(value int) bool {]
ff=是否存在
value=值

[func (a *IntArray) Search(value int) int {]
ff=查找
value=值

[func (a *IntArray) Unique() *IntArray {]
ff=去重

[func (a *IntArray) LockFunc(f func(array #左中括号##右中括号#int)) *IntArray {]
ff=遍历写锁定
f=回调函数

[func (a *IntArray) RLockFunc(f func(array #左中括号##右中括号#int)) *IntArray {]
ff=遍历读锁定
f=回调函数

[func (a *IntArray) Merge(array interface{}) *IntArray {]
ff=合并
array=切片

[func (a *IntArray) Fill(startIndex int, num int, value int) error {]
ff=填充
value=值
num=填充数量
startIndex=起点

[func (a *IntArray) Chunk(size int) #左中括号##右中括号##左中括号##右中括号#int {]
ff=分割
size=数量

[func (a *IntArray) Pad(size int, value int) *IntArray {]
ff=填满
value=值
size=总数量

[func (a *IntArray) Rand() (value int, found bool) {]
ff=取值随机
found=成功
value=值

[func (a *IntArray) Rands(size int) #左中括号##右中括号#int {]
ff=取值随机多个
size=数量

[func (a *IntArray) Shuffle() *IntArray {]
ff=随机排序

[func (a *IntArray) Reverse() *IntArray {]
ff=倒排序

[func (a *IntArray) Join(glue string) string {]
ff=连接
glue=连接符

[func (a *IntArray) CountValues() map#左中括号#int#右中括号#int {]
ff=统计

[func (a *IntArray) Iterator(f func(k int, v int) bool) {]
ff=X遍历
yx=true

[func (a *IntArray) IteratorAsc(f func(k int, v int) bool) {]
ff=遍历升序
f=回调函数

[func (a *IntArray) IteratorDesc(f func(k int, v int) bool) {]
ff=遍历降序
f=回调函数

[func (a *IntArray) Filter(filter func(index int, value int) bool) *IntArray {]
ff=遍历删除
filter=回调函数
value=值
index=索引

[func (a *IntArray) FilterEmpty() *IntArray {]
ff=删除所有零值

[func (a *IntArray) Walk(f func(value int) int) *IntArray {]
ff=遍历修改
f=回调函数

[func (a *IntArray) IsEmpty() bool {]
ff=是否为空
