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
# re.F.String()
# }
# //zj:
# 备注结束

[func NewStrArray(safe ...bool) *StrArray {]
ff=创建文本
safe=并发安全

[func NewStrArraySize(size int, cap int, safe ...bool) *StrArray {]
ff=创建文本并按大小
safe=并发安全
cap=上限
size=大小

[func NewStrArrayFrom(array #左中括号##右中括号#string, safe ...bool) *StrArray {]
ff=创建文本并从数组
safe=并发安全
array=数组

[func NewStrArrayFromCopy(array #左中括号##右中括号#string, safe ...bool) *StrArray {]
ff=创建文本并从数组复制
safe=并发安全
array=数组

[func (a *StrArray) At(index int) (value string) {]
ff=取值
value=值
index=索引

[func (a *StrArray) Get(index int) (value string, found bool) {]
ff=取值2
found=成功
value=值
index=索引

[func (a *StrArray) Set(index int, value string) error {]
ff=设置值
yx=true

[func (a *StrArray) SetArray(array #左中括号##右中括号#string) *StrArray {]
ff=设置数组
array=数组

[func (a *StrArray) Replace(array #左中括号##右中括号#string) *StrArray {]
ff=替换
array=数组

[func (a *StrArray) Sum() (sum int) {]
ff=求和
sum=值

[func (a *StrArray) Sort(reverse ...bool) *StrArray {]
ff=排序递增
reverse=降序

[func (a *StrArray) SortFunc(less func(v1, v2 string) bool) *StrArray {]
ff=排序函数
less=回调函数

[func (a *StrArray) InsertBefore(index int, values ...string) error {]
ff=插入前面
values=值
index=索引

[func (a *StrArray) InsertAfter(index int, values ...string) error {]
ff=插入后面
values=值
index=索引

[func (a *StrArray) Remove(index int) (value string, found bool) {]
ff=删除
found=成功
value=值
index=索引

[func (a *StrArray) RemoveValue(value string) bool {]
ff=删除值
value=值

[func (a *StrArray) RemoveValues(values ...string) {]
ff=删除多个值
values=值

[func (a *StrArray) PushLeft(value ...string) *StrArray {]
ff=入栈左
value=值

[func (a *StrArray) PushRight(value ...string) *StrArray {]
ff=入栈右
value=值

[func (a *StrArray) PopLeft() (value string, found bool) {]
ff=出栈左
found=成功
value=值

[func (a *StrArray) PopRight() (value string, found bool) {]
ff=出栈右
found=成功
value=值

[func (a *StrArray) PopRand() (value string, found bool) {]
ff=出栈随机
found=成功
value=值

[func (a *StrArray) PopRands(size int) #左中括号##右中括号#string {]
ff=出栈随机多个
size=数量

[func (a *StrArray) PopLefts(size int) #左中括号##右中括号#string {]
ff=出栈左多个
size=数量

[func (a *StrArray) PopRights(size int) #左中括号##右中括号#string {]
ff=出栈右多个
size=数量

[func (a *StrArray) Range(start int, end ...int) #左中括号##右中括号#string {]
ff=取切片并按范围
end=终点
start=起点

[func (a *StrArray) SubSlice(offset int, length ...int) #左中括号##右中括号#string {]
ff=取切片并按数量
length=数量
offset=起点

[func (a *StrArray) Append(value ...string) *StrArray {]
ff=Append别名
value=值

[func (a *StrArray) Len() int {]
ff=取长度

[func (a *StrArray) Slice() #左中括号##右中括号#string {]
ff=取切片

[func (a *StrArray) Interfaces() #左中括号##右中括号#interface{} {]
ff=取any数组
yx=true

[func (a *StrArray) Clone() (newArray *StrArray) {]
ff=取副本
newArray=新数组

[func (a *StrArray) Clear() *StrArray {]
ff=清空

[func (a *StrArray) Contains(value string) bool {]
ff=是否存在
value=值

[func (a *StrArray) ContainsI(value string) bool {]
ff=是否存在并忽略大小写
value=值

[func (a *StrArray) Search(value string) int {]
ff=查找
value=值

[func (a *StrArray) Unique() *StrArray {]
ff=去重

[func (a *StrArray) LockFunc(f func(array #左中括号##右中括号#string)) *StrArray {]
ff=遍历写锁定
f=回调函数

[func (a *StrArray) RLockFunc(f func(array #左中括号##右中括号#string)) *StrArray {]
ff=遍历读锁定
f=回调函数

[func (a *StrArray) Merge(array interface{}) *StrArray {]
ff=合并
array=数组

[func (a *StrArray) Fill(startIndex int, num int, value string) error {]
ff=填充
value=值
num=填充数量
startIndex=起点

[func (a *StrArray) Chunk(size int) #左中括号##右中括号##左中括号##右中括号#string {]
ff=分割
size=数量

[func (a *StrArray) Pad(size int, value string) *StrArray {]
ff=填满
value=值
size=总数量

[func (a *StrArray) Rand() (value string, found bool) {]
ff=取值随机
found=成功
value=值

[func (a *StrArray) Rands(size int) #左中括号##右中括号#string {]
ff=取值随机多个
size=数量

[func (a *StrArray) Shuffle() *StrArray {]
ff=随机排序

[func (a *StrArray) Reverse() *StrArray {]
ff=倒排序

[func (a *StrArray) Join(glue string) string {]
ff=连接
glue=连接符

[func (a *StrArray) CountValues() map#左中括号#string#右中括号#int {]
ff=统计

[func (a *StrArray) Iterator(f func(k int, v string) bool) {]
ff=X遍历
yx=true

[func (a *StrArray) IteratorAsc(f func(k int, v string) bool) {]
ff=遍历升序
f=回调函数

[func (a *StrArray) IteratorDesc(f func(k int, v string) bool) {]
ff=遍历降序
f=回调函数

[func (a *StrArray) Filter(filter func(index int, value string) bool) *StrArray {]
ff=遍历删除
filter=回调函数
value=值
index=索引

[func (a *StrArray) FilterEmpty() *StrArray {]
ff=删除所有空值

[func (a *StrArray) Walk(f func(value string) string) *StrArray {]
ff=遍历修改
f=回调函数

[func (a *StrArray) IsEmpty() bool {]
ff=是否为空
