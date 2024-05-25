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

[func NewStrSet(safe ...bool) *StrSet {]
ff=创建文本
safe=并发安全

[func NewStrSetFrom(items #左中括号##右中括号#string, safe ...bool) *StrSet {]
ff=创建文本并按值
safe=并发安全
items=值

[func (set *StrSet) Iterator(f func(v string) bool) {]
ff=X遍历
yx=true

[func (set *StrSet) Add(item ...string) {]
ff=加入
item=值s

[func (set *StrSet) AddIfNotExist(item string) bool {]
ff=加入值并跳过已存在
item=值

[func (set *StrSet) AddIfNotExistFunc(item string, f func() bool) bool {]
ff=加入值并跳过已存在_函数
item=值

[func (set *StrSet) AddIfNotExistFuncLock(item string, f func() bool) bool {]
ff=加入值并跳过已存在_并发安全函数
item=值

[func (set *StrSet) Contains(item string) bool {]
ff=是否存在
item=值

[func (set *StrSet) ContainsI(item string) bool {]
ff=是否存在并忽略大小写
item=值

[func (set *StrSet) Remove(item string) {]
ff=删除
item=值

[func (set *StrSet) Size() int {]
ff=取数量

[func (set *StrSet) Clear() {]
ff=清空

[func (set *StrSet) Slice() #左中括号##右中括号#string {]
ff=取集合切片

[func (set *StrSet) Join(glue string) string {]
ff=取集合文本
glue=连接符

[func (set *StrSet) LockFunc(f func(m map#左中括号#string#右中括号#struct{})) {]
ff=写锁定_函数

[func (set *StrSet) RLockFunc(f func(m map#左中括号#string#右中括号#struct{})) {]
ff=读锁定_函数

[func (set *StrSet) Equal(other *StrSet) bool {]
ff=是否相等
other=待比较集合

[func (set *StrSet) IsSubsetOf(other *StrSet) bool {]
ff=是否为子集
other=父集

[func (set *StrSet) Union(others ...*StrSet) (newSet *StrSet) {]
ff=取并集
newSet=新集合
others=集合

[func (set *StrSet) Diff(others ...*StrSet) (newSet *StrSet) {]
ff=取差集
newSet=新集合
others=集合

[func (set *StrSet) Intersect(others ...*StrSet) (newSet *StrSet) {]
ff=取交集
newSet=新集合
others=集合

[func (set *StrSet) Complement(full *StrSet) (newSet *StrSet) {]
ff=取补集
newSet=新集合
full=集合

[func (set *StrSet) Merge(others ...*StrSet) *StrSet {]
ff=合并
others=集合s

[func (set *StrSet) Sum() (sum int) {]
ff=求和
sum=总和

[func (set *StrSet) Pop() string {]
ff=出栈

[func (set *StrSet) Pops(size int) #左中括号##右中括号#string {]
ff=出栈多个
size=数量

[func (set *StrSet) Walk(f func(item string) string) *StrSet {]
ff=遍历修改
