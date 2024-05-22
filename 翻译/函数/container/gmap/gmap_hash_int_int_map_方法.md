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

[func NewIntIntMap(safe ...bool) *IntIntMap {]
ff=创建IntInt
safe=并发安全

[func NewIntIntMapFrom(data map#左中括号#int#右中括号#int, safe ...bool) *IntIntMap {]
ff=创建IntInt并从Map
safe=并发安全
data=map值

[func (m *IntIntMap) Iterator(f func(k int, v int) bool) {]
ff=X遍历
yx=true

[func (m *IntIntMap) Clone() *IntIntMap {]
ff=取副本

[func (m *IntIntMap) Map() map#左中括号#int#右中括号#int {]
ff=取Map

[func (m *IntIntMap) MapStrAny() map#左中括号#string#右中括号#interface{} {]
ff=取MapStrAny
yx=true

[func (m *IntIntMap) MapCopy() map#左中括号#int#右中括号#int {]
ff=浅拷贝

[func (m *IntIntMap) FilterEmpty() {]
ff=删除所有空值

[func (m *IntIntMap) Set(key int, val int) {]
ff=设置值
yx=true

[func (m *IntIntMap) Sets(data map#左中括号#int#右中括号#int) {]
ff=设置值Map
data=map值

[func (m *IntIntMap) Search(key int) (value int, found bool) {]
ff=查找
key=名称

[func (m *IntIntMap) Get(key int) (value int) {]
ff=取值
key=名称

[func (m *IntIntMap) Pop() (key, value int) {]
ff=出栈
value=值
key=名称

[func (m *IntIntMap) Pops(size int) map#左中括号#int#右中括号#int {]
ff=出栈多个
size=数量

[func (m *IntIntMap) GetOrSet(key int, value int) int {]
ff=取值或设置值
key=名称

[func (m *IntIntMap) GetOrSetFunc(key int, f func() int) int {]
ff=取值或设置值_函数
key=名称

[func (m *IntIntMap) GetOrSetFuncLock(key int, f func() int) int {]
ff=取值或设置值_函数带锁
key=名称

[func (m *IntIntMap) SetIfNotExist(key int, value int) bool {]
ff=设置值并跳过已存在
key=名称

[func (m *IntIntMap) SetIfNotExistFunc(key int, f func() int) bool {]
ff=设置值并跳过已存在_函数
key=名称

[func (m *IntIntMap) SetIfNotExistFuncLock(key int, f func() int) bool {]
ff=设置值并跳过已存在_函数带锁
key=名称

[func (m *IntIntMap) Removes(keys #左中括号##右中括号#int) {]
ff=删除多个值
keys=名称

[func (m *IntIntMap) Remove(key int) (value int) {]
ff=删除
key=名称

[func (m *IntIntMap) Keys() #左中括号##右中括号#int {]
ff=取所有名称

[func (m *IntIntMap) Values() #左中括号##右中括号#int {]
ff=取所有值

[func (m *IntIntMap) Contains(key int) bool {]
ff=是否存在
key=名称

[func (m *IntIntMap) Size() int {]
ff=取数量

[func (m *IntIntMap) IsEmpty() bool {]
ff=是否为空

[func (m *IntIntMap) Clear() {]
ff=清空

[func (m *IntIntMap) Replace(data map#左中括号#int#右中括号#int) {]
ff=替换
data=map值

[func (m *IntIntMap) LockFunc(f func(m map#左中括号#int#右中括号#int)) {]
ff=遍历写锁定
f=回调函数

[func (m *IntIntMap) RLockFunc(f func(m map#左中括号#int#右中括号#int)) {]
ff=遍历读锁定
f=回调函数

[func (m *IntIntMap) Flip() {]
ff=名称值交换

[func (m *IntIntMap) Merge(other *IntIntMap) {]
ff=合并
other=map值

[func (m *IntIntMap) IsSubOf(other *IntIntMap) bool {]
ff=是否为子集
other=父集Map

[func (m *IntIntMap) Diff(other *IntIntMap) (addedKeys, removedKeys, updatedKeys #左中括号##右中括号#int) {]
ff=比较
other=map值
