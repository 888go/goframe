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

[func NewStrAnyMap(safe ...bool) *StrAnyMap {]
ff=创建StrAny
safe=并发安全

[func NewStrAnyMapFrom(data map#左中括号#string#右中括号#interface{}, safe ...bool) *StrAnyMap {]
ff=创建AnyStr并从Map
safe=并发安全
data=map值

[func (m *StrAnyMap) Iterator(f func(k string, v interface{}) bool) {]
ff=X遍历
yx=true

[func (m *StrAnyMap) Clone() *StrAnyMap {]
ff=取副本

[func (m *StrAnyMap) Map() map#左中括号#string#右中括号#interface{} {]
ff=取Map

[func (m *StrAnyMap) MapStrAny() map#左中括号#string#右中括号#interface{} {]
ff=取MapStrAny
yx=true

[func (m *StrAnyMap) MapCopy() map#左中括号#string#右中括号#interface{} {]
ff=浅拷贝

[func (m *StrAnyMap) FilterEmpty() {]
ff=删除所有空值

[func (m *StrAnyMap) FilterNil() {]
ff=删除所有nil值

[func (m *StrAnyMap) Set(key string, val interface{}) {]
ff=设置值
yx=true

[func (m *StrAnyMap) Sets(data map#左中括号#string#右中括号#interface{}) {]
ff=设置值Map
data=map值

[func (m *StrAnyMap) Search(key string) (value interface{}, found bool) {]
ff=查找
key=名称

[func (m *StrAnyMap) Get(key string) (value interface{}) {]
ff=取值
key=名称

[func (m *StrAnyMap) Pop() (key string, value interface{}) {]
ff=出栈
key=名称

[func (m *StrAnyMap) Pops(size int) map#左中括号#string#右中括号#interface{} {]
ff=出栈多个
size=数量

[func (m *StrAnyMap) GetOrSet(key string, value interface{}) interface{} {]
ff=取值或设置值
key=名称

[func (m *StrAnyMap) GetOrSetFunc(key string, f func() interface{}) interface{} {]
ff=取值或设置值_函数
key=名称

[func (m *StrAnyMap) GetOrSetFuncLock(key string, f func() interface{}) interface{} {]
ff=取值或设置值_函数带锁
key=名称

[func (m *StrAnyMap) GetVar(key string) *gvar.Var {]
ff=取值泛型类
key=名称

[func (m *StrAnyMap) GetVarOrSet(key string, value interface{}) *gvar.Var {]
ff=取值或设置值泛型类
key=名称

[func (m *StrAnyMap) GetVarOrSetFunc(key string, f func() interface{}) *gvar.Var {]
ff=取值或设置值泛型类_函数
key=名称

[func (m *StrAnyMap) GetVarOrSetFuncLock(key string, f func() interface{}) *gvar.Var {]
ff=取值或设置值泛型类_函数带锁
key=名称

[func (m *StrAnyMap) SetIfNotExist(key string, value interface{}) bool {]
ff=设置值并跳过已存在
key=名称

[func (m *StrAnyMap) SetIfNotExistFunc(key string, f func() interface{}) bool {]
ff=设置值并跳过已存在_函数
key=名称

[func (m *StrAnyMap) SetIfNotExistFuncLock(key string, f func() interface{}) bool {]
ff=设置值并跳过已存在_函数带锁
key=名称

[func (m *StrAnyMap) Removes(keys #左中括号##右中括号#string) {]
ff=删除多个值
keys=名称

[func (m *StrAnyMap) Remove(key string) (value interface{}) {]
ff=删除
key=名称

[func (m *StrAnyMap) Keys() #左中括号##右中括号#string {]
ff=取所有名称

[func (m *StrAnyMap) Values() #左中括号##右中括号#interface{} {]
ff=取所有值

[func (m *StrAnyMap) Contains(key string) bool {]
ff=是否存在
key=名称

[func (m *StrAnyMap) Size() int {]
ff=取数量

[func (m *StrAnyMap) IsEmpty() bool {]
ff=是否为空

[func (m *StrAnyMap) Clear() {]
ff=清空

[func (m *StrAnyMap) Replace(data map#左中括号#string#右中括号#interface{}) {]
ff=替换
data=map值

[func (m *StrAnyMap) LockFunc(f func(m map#左中括号#string#右中括号#interface{})) {]
ff=遍历写锁定
f=回调函数

[func (m *StrAnyMap) RLockFunc(f func(m map#左中括号#string#右中括号#interface{})) {]
ff=遍历读锁定
f=回调函数

[func (m *StrAnyMap) Flip() {]
ff=名称值交换

[func (m *StrAnyMap) Merge(other *StrAnyMap) {]
ff=合并
other=map值

[func (m *StrAnyMap) IsSubOf(other *StrAnyMap) bool {]
ff=是否为子集
other=父集Map

[func (m *StrAnyMap) Diff(other *StrAnyMap) (addedKeys, removedKeys, updatedKeys #左中括号##右中括号#string) {]
ff=比较
other=map值
