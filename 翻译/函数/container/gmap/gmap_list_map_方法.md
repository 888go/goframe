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

[func NewListMap(safe ...bool) *ListMap {]
ff=创建链表mp
safe=并发安全

[func NewListMapFrom(data map#左中括号#interface{}#右中括号#interface{}, safe ...bool) *ListMap {]
ff=创建链表Map并从Map
safe=并发安全
data=map值

[func (m *ListMap) Iterator(f func(key, value interface{}) bool) {]
ff=X遍历
yx=true

[func (m *ListMap) IteratorAsc(f func(key interface{}, value interface{}) bool) {]
ff=遍历升序
f=回调函数

[func (m *ListMap) IteratorDesc(f func(key interface{}, value interface{}) bool) {]
ff=遍历降序
f=回调函数

[func (m *ListMap) Clone(safe ...bool) *ListMap {]
ff=取副本
safe=并发安全

[func (m *ListMap) Clear() {]
ff=清空

[func (m *ListMap) Replace(data map#左中括号#interface{}#右中括号#interface{}) {]
ff=替换
data=map值

[func (m *ListMap) Map() map#左中括号#interface{}#右中括号#interface{} {]
ff=取Map

[func (m *ListMap) MapStrAny() map#左中括号#string#右中括号#interface{} {]
ff=取MapStrAny
yx=true

[func (m *ListMap) FilterEmpty() {]
ff=删除所有空值

[func (m *ListMap) Set(key interface{}, value interface{}) {]
ff=设置值
yx=true

[func (m *ListMap) Sets(data map#左中括号#interface{}#右中括号#interface{}) {]
ff=设置值Map
data=map值

[func (m *ListMap) Search(key interface{}) (value interface{}, found bool) {]
ff=查找
found=成功
value=值
key=名称

[func (m *ListMap) Get(key interface{}) (value interface{}) {]
ff=取值
value=值
key=名称

[func (m *ListMap) Pop() (key, value interface{}) {]
ff=出栈
value=值
key=名称

[func (m *ListMap) Pops(size int) map#左中括号#interface{}#右中括号#interface{} {]
ff=出栈多个
size=数量

[func (m *ListMap) GetOrSet(key interface{}, value interface{}) interface{} {]
ff=取值或设置值
value=值
key=名称

[func (m *ListMap) GetOrSetFunc(key interface{}, f func() interface{}) interface{} {]
ff=取值或设置值_函数
f=回调函数
key=名称

[func (m *ListMap) GetOrSetFuncLock(key interface{}, f func() interface{}) interface{} {]
ff=取值或设置值_函数带锁
f=回调函数
key=名称

[func (m *ListMap) GetVar(key interface{}) *gvar.Var {]
ff=取值泛型类
key=名称

[func (m *ListMap) GetVarOrSet(key interface{}, value interface{}) *gvar.Var {]
ff=取值或设置值泛型类
value=值
key=名称

[func (m *ListMap) GetVarOrSetFunc(key interface{}, f func() interface{}) *gvar.Var {]
ff=取值或设置值泛型类_函数
f=回调函
key=名称

[func (m *ListMap) GetVarOrSetFuncLock(key interface{}, f func() interface{}) *gvar.Var {]
ff=取值或设置值泛型类_函数带锁
f=回调函数
key=名称

[func (m *ListMap) SetIfNotExist(key interface{}, value interface{}) bool {]
ff=设置值并跳过已存在
value=值
key=名称

[func (m *ListMap) SetIfNotExistFunc(key interface{}, f func() interface{}) bool {]
ff=设置值并跳过已存在_函数
f=回调函数
key=名称

[func (m *ListMap) SetIfNotExistFuncLock(key interface{}, f func() interface{}) bool {]
ff=设置值并跳过已存在_函数带锁
f=回调函数
key=名称

[func (m *ListMap) Remove(key interface{}) (value interface{}) {]
ff=删除
value=值
key=名称

[func (m *ListMap) Removes(keys #左中括号##右中括号#interface{}) {]
ff=删除多个值
keys=名称

[func (m *ListMap) Keys() #左中括号##右中括号#interface{} {]
ff=取所有名称

[func (m *ListMap) Values() #左中括号##右中括号#interface{} {]
ff=取所有值

[func (m *ListMap) Contains(key interface{}) (ok bool) {]
ff=是否存在
key=名称

[func (m *ListMap) Size() (size int) {]
ff=取数量
size=数量

[func (m *ListMap) IsEmpty() bool {]
ff=是否为空

[func (m *ListMap) Flip() {]
ff=名称值交换

[func (m *ListMap) Merge(other *ListMap) {]
ff=合并
other=map值
