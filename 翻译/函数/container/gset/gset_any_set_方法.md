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

[func New(safe ...bool) *Set {]
safe=并发安全

[func NewSet(safe ...bool) *Set {]
ff=NewSet别名

[func NewFrom(items interface{}, safe ...bool) *Set {]
ff=创建并按值
safe=并发安全
items=值

[func (set *Set) Iterator(f func(v interface{}) bool) {]
ff=X遍历

[func (set *Set) Add(items ...interface{}) {]
ff=加入
items=值s

[func (set *Set) AddIfNotExist(item interface{}) bool {]
ff=加入值并跳过已存在
item=值

[func (set *Set) AddIfNotExistFunc(item interface{}, f func() bool) bool {]
ff=加入值并跳过已存在_函数
item=值

[func (set *Set) AddIfNotExistFuncLock(item interface{}, f func() bool) bool {]
ff=加入值并跳过已存在_并发安全函数
item=值

[func (set *Set) Contains(item interface{}) bool {]
ff=是否存在
item=值

[func (set *Set) Remove(item interface{}) {]
ff=删除
item=值

[func (set *Set) Size() int {]
ff=取数量

[func (set *Set) Clear() {]
ff=清空

[func (set *Set) Slice() #左中括号##右中括号#interface{} {]
ff=取集合数组

[func (set *Set) Join(glue string) string {]
ff=取集合文本
glue=连接符

[func (set *Set) LockFunc(f func(m map#左中括号#interface{}#右中括号#struct{})) {]
ff=写锁定_函数

[func (set *Set) RLockFunc(f func(m map#左中括号#interface{}#右中括号#struct{})) {]
ff=读锁定_函数

[func (set *Set) Equal(other *Set) bool {]
ff=是否相等
other=待比较集合

[func (set *Set) IsSubsetOf(other *Set) bool {]
ff=是否为子集
other=父集

[func (set *Set) Union(others ...*Set) (newSet *Set) {]
ff=取并集
newSet=新集合
others=集合

[func (set *Set) Diff(others ...*Set) (newSet *Set) {]
ff=取差集
newSet=新集合
others=集合

[func (set *Set) Intersect(others ...*Set) (newSet *Set) {]
ff=取交集
newSet=新集合
others=集合

[func (set *Set) Complement(full *Set) (newSet *Set) {]
ff=取补集
newSet=新集合
full=集合

[func (set *Set) Merge(others ...*Set) *Set {]
ff=合并
others=集合s

[func (set *Set) Sum() (sum int) {]
ff=求和
sum=总和

[func (set *Set) Pop() interface{} {]
ff=出栈

[func (set *Set) Pops(size int) #左中括号##右中括号#interface{} {]
ff=出栈多个
size=数量

[func (set *Set) Walk(f func(item interface{}) interface{}) *Set {]
ff=遍历修改
