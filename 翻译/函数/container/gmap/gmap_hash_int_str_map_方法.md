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

[func NewIntStrMap(safe ...bool) *IntStrMap {]
ff=创建IntStr
safe=并发安全

[func NewIntStrMapFrom(data map#左中括号#int#右中括号#string, safe ...bool) *IntStrMap {]
ff=创建IntStr并从Map
safe=并发安全
data=map值

[func (m *IntStrMap) Iterator(f func(k int, v string) bool) {]
ff=X遍历

[func (m *IntStrMap) Clone() *IntStrMap {]
ff=取副本

[func (m *IntStrMap) Map() map#左中括号#int#右中括号#string {]
ff=取Map

[func (m *IntStrMap) MapStrAny() map#左中括号#string#右中括号#interface{} {]
ff=取MapStrAny

[func (m *IntStrMap) MapCopy() map#左中括号#int#右中括号#string {]
ff=浅拷贝

[func (m *IntStrMap) FilterEmpty() {]
ff=删除所有空值

[func (m *IntStrMap) Set(key int, val string) {]
ff=设置值

[func (m *IntStrMap) Sets(data map#左中括号#int#右中括号#string) {]
ff=设置值Map
data=map值

[func (m *IntStrMap) Search(key int) (value string, found bool) {]
ff=查找
key=名称

[func (m *IntStrMap) Get(key int) (value string) {]
ff=取值
key=名称

[func (m *IntStrMap) Pop() (key int, value string) {]
ff=出栈
key=名称

[func (m *IntStrMap) Pops(size int) map#左中括号#int#右中括号#string {]
ff=出栈多个
size=数量

[func (m *IntStrMap) GetOrSet(key int, value string) string {]
ff=取值或设置值
key=名称

[func (m *IntStrMap) GetOrSetFunc(key int, f func() string) string {]
ff=取值或设置值_函数
key=名称

[func (m *IntStrMap) GetOrSetFuncLock(key int, f func() string) string {]
ff=取值或设置值_函数带锁
key=名称

[func (m *IntStrMap) SetIfNotExist(key int, value string) bool {]
ff=设置值并跳过已存在
key=名称

[func (m *IntStrMap) SetIfNotExistFunc(key int, f func() string) bool {]
ff=设置值并跳过已存在_函数
key=名称

[func (m *IntStrMap) SetIfNotExistFuncLock(key int, f func() string) bool {]
ff=设置值并跳过已存在_函数带锁
key=名称

[func (m *IntStrMap) Removes(keys #左中括号##右中括号#int) {]
ff=删除多个值
keys=名称

[func (m *IntStrMap) Remove(key int) (value string) {]
ff=删除
key=名称

[func (m *IntStrMap) Keys() #左中括号##右中括号#int {]
ff=取所有名称

[func (m *IntStrMap) Values() #左中括号##右中括号#string {]
ff=取所有值

[func (m *IntStrMap) Contains(key int) bool {]
ff=是否存在
key=名称

[func (m *IntStrMap) Size() int {]
ff=取数量

[func (m *IntStrMap) IsEmpty() bool {]
ff=是否为空

[func (m *IntStrMap) Clear() {]
ff=清空

[func (m *IntStrMap) Replace(data map#左中括号#int#右中括号#string) {]
ff=替换
data=map值

[func (m *IntStrMap) LockFunc(f func(m map#左中括号#int#右中括号#string)) {]
ff=遍历写锁定
f=回调函数

[func (m *IntStrMap) RLockFunc(f func(m map#左中括号#int#右中括号#string)) {]
ff=遍历读锁定
f=回调函数

[func (m *IntStrMap) Flip() {]
ff=名称值交换

[func (m *IntStrMap) Merge(other *IntStrMap) {]
ff=合并
other=map值

[func (m *IntStrMap) IsSubOf(other *IntStrMap) bool {]
ff=是否为子集
other=父集Map

[func (m *IntStrMap) Diff(other *IntStrMap) (addedKeys, removedKeys, updatedKeys #左中括号##右中括号#int) {]
ff=比较
other=map值
