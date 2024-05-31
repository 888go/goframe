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

[func NewIntSet(safe ...bool) *IntSet {]
ff=创建整数
safe=并发安全

[func NewIntSetFrom(items #左中括号##右中括号#int, safe ...bool) *IntSet {]
ff=创建整数并按值
safe=并发安全
items=整数数组

[func (set *IntSet) Iterator(f func(v int) bool) {]
ff=X遍历
yx=true

[func (set *IntSet) Add(item ...int) {]
ff=加入
item=值s

[func (set *IntSet) AddIfNotExist(item int) bool {]
ff=加入值并跳过已存在
item=值

[func (set *IntSet) AddIfNotExistFunc(item int, f func() bool) bool {]
ff=加入值并跳过已存在_函数
item=值

[func (set *IntSet) AddIfNotExistFuncLock(item int, f func() bool) bool {]
ff=加入值并跳过已存在_并发安全函数
item=值

[func (set *IntSet) Contains(item int) bool {]
ff=是否存在
item=值

[func (set *IntSet) Remove(item int) {]
ff=删除
item=值

[func (set *IntSet) Size() int {]
ff=取数量

[func (set *IntSet) Clear() {]
ff=清空

[func (set *IntSet) Slice() #左中括号##右中括号#int {]
ff=取集合数组

[func (set *IntSet) Join(glue string) string {]
ff=取集合文本
glue=连接符

[func (set *IntSet) LockFunc(f func(m map#左中括号#int#右中括号#struct{})) {]
ff=写锁定_函数

[func (set *IntSet) RLockFunc(f func(m map#左中括号#int#右中括号#struct{})) {]
ff=读锁定_函数

[func (set *IntSet) Equal(other *IntSet) bool {]
ff=是否相等
other=待比较集合

[func (set *IntSet) IsSubsetOf(other *IntSet) bool {]
ff=是否为子集
other=父集

[func (set *IntSet) Union(others ...*IntSet) (newSet *IntSet) {]
ff=取并集
newSet=新集合
others=集合

[func (set *IntSet) Diff(others ...*IntSet) (newSet *IntSet) {]
ff=取差集
newSet=新集合
others=集合

[func (set *IntSet) Intersect(others ...*IntSet) (newSet *IntSet) {]
ff=取交集
newSet=新集合
others=集合

[func (set *IntSet) Complement(full *IntSet) (newSet *IntSet) {]
ff=取补集
newSet=新集合
full=集合

[func (set *IntSet) Merge(others ...*IntSet) *IntSet {]
ff=合并
others=集合s

[func (set *IntSet) Sum() (sum int) {]
ff=求和
sum=总和

[func (set *IntSet) Pop() int {]
ff=出栈

[func (set *IntSet) Pops(size int) #左中括号##右中括号#int {]
ff=出栈多个
size=数量

[func (set *IntSet) Walk(f func(item int) int) *IntSet {]
ff=遍历修改
