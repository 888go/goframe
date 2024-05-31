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

[func NewStrIntMap(safe ...bool) *StrIntMap {]
ff=创建StrInt
safe=并发安全

[func NewStrIntMapFrom(data map#左中括号#string#右中括号#int, safe ...bool) *StrIntMap {]
ff=创建StrInt并从Map
safe=并发安全
data=map值

[func (m *StrIntMap) Iterator(f func(k string, v int) bool) {]
ff=X遍历
yx=true

[func (m *StrIntMap) Clone() *StrIntMap {]
ff=取副本

[func (m *StrIntMap) Map() map#左中括号#string#右中括号#int {]
ff=取Map

[func (m *StrIntMap) MapStrAny() map#左中括号#string#右中括号#interface{} {]
ff=取MapStrAny
yx=true

[func (m *StrIntMap) MapCopy() map#左中括号#string#右中括号#int {]
ff=浅拷贝

[func (m *StrIntMap) FilterEmpty() {]
ff=删除所有空值

[func (m *StrIntMap) Set(key string, val int) {]
ff=设置值
yx=true

[func (m *StrIntMap) Sets(data map#左中括号#string#右中括号#int) {]
ff=设置值Map
data=map值

[func (m *StrIntMap) Search(key string) (value int, found bool) {]
ff=查找
found=成功
value=值
key=名称

[func (m *StrIntMap) Get(key string) (value int) {]
ff=取值
key=名称

[func (m *StrIntMap) Pop() (key string, value int) {]
ff=出栈
key=名称

[func (m *StrIntMap) Pops(size int) map#左中括号#string#右中括号#int {]
ff=出栈多个
size=数量

[func (m *StrIntMap) GetOrSet(key string, value int) int {]
ff=取值或设置值
value=值
key=名称

[func (m *StrIntMap) GetOrSetFunc(key string, f func() int) int {]
ff=取值或设置值_函数
key=名称

[func (m *StrIntMap) GetOrSetFuncLock(key string, f func() int) int {]
ff=取值或设置值_函数带锁
key=名称

[func (m *StrIntMap) SetIfNotExist(key string, value int) bool {]
ff=设置值并跳过已存在
value=值
key=名称

[func (m *StrIntMap) SetIfNotExistFunc(key string, f func() int) bool {]
ff=设置值并跳过已存在_函数
key=名称

[func (m *StrIntMap) SetIfNotExistFuncLock(key string, f func() int) bool {]
ff=设置值并跳过已存在_函数带锁
key=名称

[func (m *StrIntMap) Removes(keys #左中括号##右中括号#string) {]
ff=删除多个值
keys=名称

[func (m *StrIntMap) Remove(key string) (value int) {]
ff=删除
value=被删除值
key=名称

[func (m *StrIntMap) Keys() #左中括号##右中括号#string {]
ff=取所有名称

[func (m *StrIntMap) Values() #左中括号##右中括号#int {]
ff=取所有值

[func (m *StrIntMap) Contains(key string) bool {]
ff=是否存在
key=名称

[func (m *StrIntMap) Size() int {]
ff=取数量

[func (m *StrIntMap) IsEmpty() bool {]
ff=是否为空

[func (m *StrIntMap) Clear() {]
ff=清空

[func (m *StrIntMap) Replace(data map#左中括号#string#右中括号#int) {]
ff=替换
data=map值

[func (m *StrIntMap) LockFunc(f func(m map#左中括号#string#右中括号#int)) {]
ff=遍历写锁定
f=回调函数

[func (m *StrIntMap) RLockFunc(f func(m map#左中括号#string#右中括号#int)) {]
ff=遍历读锁定
f=回调函数

[func (m *StrIntMap) Flip() {]
ff=名称值交换

[func (m *StrIntMap) Merge(other *StrIntMap) {]
ff=合并
other=map值

[func (m *StrIntMap) IsSubOf(other *StrIntMap) bool {]
ff=是否为子集
other=父集Map

[func (m *StrIntMap) Diff(other *StrIntMap) (addedKeys, removedKeys, updatedKeys #左中括号##右中括号#string) {]
ff=比较
updatedKeys=更新数据的名称
removedKeys=删除的名称
addedKeys=增加的名称
other=map值
