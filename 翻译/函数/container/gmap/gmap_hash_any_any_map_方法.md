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

[func NewAnyAnyMap(safe ...bool) *AnyAnyMap {]
ff=创建AnyAny
safe=并发安全

[func NewAnyAnyMapFrom(data map#左中括号#interface{}#右中括号#interface{}, safe ...bool) *AnyAnyMap {]
ff=创建AnyAny并从Map
safe=并发安全
data=map值

[func (m *AnyAnyMap) Iterator(f func(k interface{}, v interface{}) bool) {]
ff=X遍历
yx=true

[func (m *AnyAnyMap) Clone(safe ...bool) *AnyAnyMap {]
ff=取副本
safe=并发安全

[func (m *AnyAnyMap) Map() map#左中括号#interface{}#右中括号#interface{} {]
ff=取Map

[func (m *AnyAnyMap) MapCopy() map#左中括号#interface{}#右中括号#interface{} {]
ff=浅拷贝

[func (m *AnyAnyMap) MapStrAny() map#左中括号#string#右中括号#interface{} {]
ff=取MapStrAny
yx=true

[func (m *AnyAnyMap) FilterEmpty() {]
ff=删除所有空值

[func (m *AnyAnyMap) FilterNil() {]
ff=删除所有nil值

[func (m *AnyAnyMap) Set(key interface{}, value interface{}) {]
ff=设置值
yx=true

[func (m *AnyAnyMap) Sets(data map#左中括号#interface{}#右中括号#interface{}) {]
ff=设置值Map
data=map值

[func (m *AnyAnyMap) Search(key interface{}) (value interface{}, found bool) {]
ff=查找
found=成功
value=值
key=名称

[func (m *AnyAnyMap) Get(key interface{}) (value interface{}) {]
ff=取值
value=值
key=名称

[func (m *AnyAnyMap) Pop() (key, value interface{}) {]
ff=出栈
value=值
key=名称

[func (m *AnyAnyMap) Pops(size int) map#左中括号#interface{}#右中括号#interface{} {]
ff=出栈多个
size=数量

[func (m *AnyAnyMap) GetOrSet(key interface{}, value interface{}) interface{} {]
ff=取值或设置值
value=值
key=名称

[func (m *AnyAnyMap) GetOrSetFunc(key interface{}, f func() interface{}) interface{} {]
ff=取值或设置值_函数
f=回调函数
key=名称

[func (m *AnyAnyMap) GetOrSetFuncLock(key interface{}, f func() interface{}) interface{} {]
ff=取值或设置值_函数带锁
f=回调函数
key=名称

[func (m *AnyAnyMap) GetVar(key interface{}) *gvar.Var {]
ff=取值泛型类
key=名称

[func (m *AnyAnyMap) GetVarOrSet(key interface{}, value interface{}) *gvar.Var {]
ff=取值或设置值泛型类
value=值
key=名称

[func (m *AnyAnyMap) GetVarOrSetFunc(key interface{}, f func() interface{}) *gvar.Var {]
ff=取值或设置值泛型类_函数
f=回调函
key=名称

[func (m *AnyAnyMap) GetVarOrSetFuncLock(key interface{}, f func() interface{}) *gvar.Var {]
ff=取值或设置值泛型类_函数带锁
f=回调函数
key=名称

[func (m *AnyAnyMap) SetIfNotExist(key interface{}, value interface{}) bool {]
ff=设置值并跳过已存在
value=值
key=名称

[func (m *AnyAnyMap) SetIfNotExistFunc(key interface{}, f func() interface{}) bool {]
ff=设置值并跳过已存在_函数
f=回调函数
key=名称

[func (m *AnyAnyMap) SetIfNotExistFuncLock(key interface{}, f func() interface{}) bool {]
ff=设置值并跳过已存在_函数带锁
f=回调函数
key=名称

[func (m *AnyAnyMap) Remove(key interface{}) (value interface{}) {]
ff=删除
value=值
key=名称

[func (m *AnyAnyMap) Removes(keys #左中括号##右中括号#interface{}) {]
ff=删除多个值
keys=名称

[func (m *AnyAnyMap) Keys() #左中括号##右中括号#interface{} {]
ff=取所有名称

[func (m *AnyAnyMap) Values() #左中括号##右中括号#interface{} {]
ff=取所有值

[func (m *AnyAnyMap) Contains(key interface{}) bool {]
ff=是否存在
key=名称

[func (m *AnyAnyMap) Size() int {]
ff=取数量

[func (m *AnyAnyMap) IsEmpty() bool {]
ff=是否为空

[func (m *AnyAnyMap) Clear() {]
ff=清空

[func (m *AnyAnyMap) Replace(data map#左中括号#interface{}#右中括号#interface{}) {]
ff=替换
data=map值

[func (m *AnyAnyMap) LockFunc(f func(m map#左中括号#interface{}#右中括号#interface{})) {]
ff=遍历写锁定
f=回调函数

[func (m *AnyAnyMap) RLockFunc(f func(m map#左中括号#interface{}#右中括号#interface{})) {]
ff=遍历读锁定
f=回调函数

[func (m *AnyAnyMap) Flip() {]
ff=名称值交换

[func (m *AnyAnyMap) Merge(other *AnyAnyMap) {]
ff=合并
other=map值

[func (m *AnyAnyMap) IsSubOf(other *AnyAnyMap) bool {]
ff=是否为子集
other=父集Map

[func (m *AnyAnyMap) Diff(other *AnyAnyMap) (addedKeys, removedKeys, updatedKeys #左中括号##右中括号#interface{}) {]
ff=比较
updatedKeys=更新数据的名称
removedKeys=删除的名称
addedKeys=增加的名称
other=map值
