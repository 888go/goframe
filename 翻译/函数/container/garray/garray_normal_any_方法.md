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

[func New(safe ...bool) *Array {]
ff=创建
safe=并发安全

[func NewArray(safe ...bool) *Array {]
ff=NewArray别名
safe=并发安全

[func NewArraySize(size int, cap int, safe ...bool) *Array {]
ff=创建并按大小
safe=并发安全
cap=上限
size=大小

[func NewArrayRange(start, end, step int, safe ...bool) *Array {]
ff=创建并按范围
safe=并发安全
step=步长
end=终点
start=起点

[func NewFrom(array #左中括号##右中括号#interface{}, safe ...bool) *Array {]
ff=NewFrom别名
safe=并发安全
array=数组

[func NewFromCopy(array #左中括号##右中括号#interface{}, safe ...bool) *Array {]
ff=NewFromCopy别名
safe=并发安全
array=数组

[func NewArrayFrom(array #左中括号##右中括号#interface{}, safe ...bool) *Array {]
ff=创建并从数组
safe=并发安全
array=数组

[func NewArrayFromCopy(array #左中括号##右中括号#interface{}, safe ...bool) *Array {]
ff=创建并从数组复制
safe=并发安全
array=数组

[func (a *Array) At(index int) (value interface{}) {]
ff=取值
value=值
index=索引

[func (a *Array) Get(index int) (value interface{}, found bool) {]
ff=取值2
found=成功
value=值
index=索引

[func (a *Array) Set(index int, value interface{}) error {]
ff=设置值
yx=true

[func (a *Array) SetArray(array #左中括号##右中括号#interface{}) *Array {]
ff=设置数组
array=数组

[func (a *Array) Replace(array #左中括号##右中括号#interface{}) *Array {]
ff=替换
array=数组

[func (a *Array) Sum() (sum int) {]
ff=求和
sum=值

[func (a *Array) SortFunc(less func(v1, v2 interface{}) bool) *Array {]
ff=排序并带函数
less=回调函数

[func (a *Array) InsertBefore(index int, values ...interface{}) error {]
ff=插入前面
values=值
index=索引

[func (a *Array) InsertAfter(index int, values ...interface{}) error {]
ff=插入后面
values=值
index=索引

[func (a *Array) Remove(index int) (value interface{}, found bool) {]
ff=删除
found=成功
value=值
index=索引

[func (a *Array) RemoveValue(value interface{}) bool {]
ff=删除值
value=值

[func (a *Array) RemoveValues(values ...interface{}) {]
ff=删除多个值
values=值

[func (a *Array) PushLeft(value ...interface{}) *Array {]
ff=入栈左
value=值

[func (a *Array) PushRight(value ...interface{}) *Array {]
ff=入栈右
value=值

[func (a *Array) PopRand() (value interface{}, found bool) {]
ff=出栈随机
found=成功
value=值

[func (a *Array) PopRands(size int) #左中括号##右中括号#interface{} {]
ff=出栈随机多个
size=数量

[func (a *Array) PopLeft() (value interface{}, found bool) {]
ff=出栈左
found=成功
value=值

[func (a *Array) PopRight() (value interface{}, found bool) {]
ff=出栈右
found=成功
value=值

[func (a *Array) PopLefts(size int) #左中括号##右中括号#interface{} {]
ff=出栈左多个
size=数量

[func (a *Array) PopRights(size int) #左中括号##右中括号#interface{} {]
ff=出栈右多个
size=数量

[func (a *Array) Range(start int, end ...int) #左中括号##右中括号#interface{} {]
ff=取切片并按范围
end=终点
start=起点

[func (a *Array) SubSlice(offset int, length ...int) #左中括号##右中括号#interface{} {]
ff=取切片并按数量
length=数量
offset=起点

[func (a *Array) Append(value ...interface{}) *Array {]
ff=Append别名
value=值

[func (a *Array) Len() int {]
ff=取长度

[func (a *Array) Slice() #左中括号##右中括号#interface{} {]
ff=取切片

[func (a *Array) Interfaces() #左中括号##右中括号#interface{} {]
ff=取any数组
yx=true

[func (a *Array) Clone() (newArray *Array) {]
ff=取副本
newArray=新数组

[func (a *Array) Clear() *Array {]
ff=清空

[func (a *Array) Contains(value interface{}) bool {]
ff=是否存在
value=值

[func (a *Array) Search(value interface{}) int {]
ff=查找
value=值

[func (a *Array) Unique() *Array {]
ff=去重

[func (a *Array) LockFunc(f func(array #左中括号##右中括号#interface{})) *Array {]
ff=遍历并写锁定
f=回调函数

[func (a *Array) RLockFunc(f func(array #左中括号##右中括号#interface{})) *Array {]
ff=遍历并读锁定
f=回调函数

[func (a *Array) Merge(array interface{}) *Array {]
ff=合并
array=数组

[func (a *Array) Fill(startIndex int, num int, value interface{}) error {]
ff=填充
value=值
num=填充数量
startIndex=起点

[func (a *Array) Chunk(size int) #左中括号##右中括号##左中括号##右中括号#interface{} {]
ff=分割
size=数量

[func (a *Array) Pad(size int, val interface{}) *Array {]
ff=填满
val=值
size=总数量

[func (a *Array) Rand() (value interface{}, found bool) {]
ff=取值随机
found=成功
value=值

[func (a *Array) Rands(size int) #左中括号##右中括号#interface{} {]
ff=取值随机多个
size=数量

[func (a *Array) Shuffle() *Array {]
ff=随机排序

[func (a *Array) Reverse() *Array {]
ff=倒排序

[func (a *Array) Join(glue string) string {]
ff=连接
glue=连接符

[func (a *Array) CountValues() map#左中括号#interface{}#右中括号#int {]
ff=统计

[func (a *Array) Iterator(f func(k int, v interface{}) bool) {]
ff=X遍历
yx=true

[func (a *Array) IteratorAsc(f func(k int, v interface{}) bool) {]
ff=遍历升序
f=回调函数

[func (a *Array) IteratorDesc(f func(k int, v interface{}) bool) {]
ff=遍历降序
f=回调函数

[func (a *Array) Filter(filter func(index int, value interface{}) bool) *Array {]
ff=遍历删除
filter=回调函数
value=值
index=索引

[func (a *Array) FilterNil() *Array {]
ff=删除所有nil

[func (a *Array) FilterEmpty() *Array {]
ff=删除所有空值

[func (a *Array) Walk(f func(value interface{}) interface{}) *Array {]
ff=遍历修改
f=回调函数

[func (a *Array) IsEmpty() bool {]
ff=是否为空
