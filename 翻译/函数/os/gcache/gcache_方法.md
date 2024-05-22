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

[func Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error {]
ff=设置值
yx=true

[func SetMap(ctx context.Context, data map#左中括号#interface{}#右中括号#interface{}, duration time.Duration) error {]
ff=设置Map
duration=时长
data=值
ctx=上下文

[func SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (bool, error) {]
ff=设置值并跳过已存在
duration=时长
value=值
key=名称
ctx=上下文

[func SetIfNotExistFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (bool, error) {]
ff=设置值并跳过已存在_函数
duration=时长
f=回调函数
key=名称
ctx=上下文

[func SetIfNotExistFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (bool, error) {]
ff=设置值并跳过已存在_并发安全函数
duration=时长
f=回调函数
key=名称
ctx=上下文

[func Get(ctx context.Context, key interface{}) (*gvar.Var, error) {]
ff=取值
key=名称
ctx=上下文

[func GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (*gvar.Var, error) {]
ff=取值或设置值
duration=时长
value=值
key=名称
ctx=上下文

[func GetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (*gvar.Var, error) {]
ff=取值或设置值_函数
duration=时长
f=回调函数
key=名称
ctx=上下文

[func GetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (*gvar.Var, error) {]
ff=取值或设置值_并发安全函数
duration=时长
f=回调函数
key=名称
ctx=上下文

[func Contains(ctx context.Context, key interface{}) (bool, error) {]
ff=是否存在
key=名称
ctx=上下文

[func GetExpire(ctx context.Context, key interface{}) (time.Duration, error) {]
ff=取过期时间
key=名称
ctx=上下文

[func Remove(ctx context.Context, keys ...interface{}) (value *gvar.Var, err error) {]
ff=删除并带返回值
value=可选值
keys=名称s
ctx=上下文

[func Removes(ctx context.Context, keys #左中括号##右中括号#interface{}) error {]
ff=删除
keys=名称s
ctx=上下文

[func Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error) {]
ff=更新值
oldValue=旧值
value=值
key=名称
ctx=上下文

[func UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error) {]
ff=更新过期时间
err=错误
oldDuration=旧过期时长
duration=时长
key=名称
ctx=上下文

[func Size(ctx context.Context) (int, error) {]
ff=取数量
ctx=上下文

[func Data(ctx context.Context) (map#左中括号#interface{}#右中括号#interface{}, error) {]
ff=取所有键值Map副本
ctx=上下文

[func Keys(ctx context.Context) (#左中括号##右中括号#interface{}, error) {]
ff=取所有键
ctx=上下文

[func KeyStrings(ctx context.Context) (#左中括号##右中括号#string, error) {]
ff=取所有键文本
ctx=上下文

[func Values(ctx context.Context) (#左中括号##右中括号#interface{}, error) {]
ff=取所有值
ctx=上下文

[func MustGet(ctx context.Context, key interface{}) *gvar.Var {]
ff=取值PANI
key=名称
ctx=上下文

[func MustGetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) *gvar.Var {]
ff=取值或设置值PANI
duration=时长
value=值
key=名称
ctx=上下文

[func MustGetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) *gvar.Var {]
ff=取值或设置值_函数PANI
duration=时长
f=回调函数
key=名称
ctx=上下文

[func MustGetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) *gvar.Var {]
ff=取值或设置值_并发安全函数PANI
duration=时长
f=回调函数
key=名称
ctx=上下文

[func MustContains(ctx context.Context, key interface{}) bool {]
ff=是否存在PANI
key=名称
ctx=上下文

[func MustGetExpire(ctx context.Context, key interface{}) time.Duration {]
ff=取过期时间PANI
key=名称
ctx=上下文

[func MustSize(ctx context.Context) int {]
ff=取数量PANI
ctx=上下文

[func MustData(ctx context.Context) map#左中括号#interface{}#右中括号#interface{} {]
ff=取所有键值Map副本PANI
ctx=上下文

[func MustKeys(ctx context.Context) #左中括号##右中括号#interface{} {]
ff=取所有键PANI
ctx=上下文

[func MustKeyStrings(ctx context.Context) #左中括号##右中括号#string {]
ff=取所有键文本PANI
ctx=上下文

[func MustValues(ctx context.Context) #左中括号##右中括号#interface{} {]
ff=取所有值PANI
ctx=上下文
