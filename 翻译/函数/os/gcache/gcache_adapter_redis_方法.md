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

[func (c *AdapterRedis) Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (err error) {]
ff=设置值
yx=true

[func (c *AdapterRedis) SetMap(ctx context.Context, data map#左中括号#interface{}#右中括号#interface{}, duration time.Duration) error {]
ff=设置Map
duration=时长
data=值
ctx=上下文

[func (c *AdapterRedis) SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (bool, error) {]
ff=设置值并跳过已存在
duration=时长
value=值
key=名称
ctx=上下文

[func (c *AdapterRedis) SetIfNotExistFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (ok bool, err error) {]
ff=设置值并跳过已存在_函数
err=错误
ok=成功
duration=时长
f=回调函数
key=名称
ctx=上下文

[func (c *AdapterRedis) SetIfNotExistFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (ok bool, err error) {]
ff=设置值并跳过已存在_并发安全函数
err=错误
ok=成功
duration=时长
f=回调函数
key=名称
ctx=上下文

[func (c *AdapterRedis) Get(ctx context.Context, key interface{}) (*gvar.Var, error) {]
ff=取值
key=名称
ctx=上下文

[func (c *AdapterRedis) GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (result *gvar.Var, err error) {]
ff=取值或设置值
result=结果
duration=时长
value=值
key=名称
ctx=上下文

[func (c *AdapterRedis) GetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (result *gvar.Var, err error) {]
ff=取值或设置值_函数
result=结果
duration=时长
f=回调函数
key=名称
ctx=上下文

[func (c *AdapterRedis) GetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (result *gvar.Var, err error) {]
ff=取值或设置值_并发安全函数
result=结果
duration=时长
f=回调函数
key=名称
ctx=上下文

[func (c *AdapterRedis) Contains(ctx context.Context, key interface{}) (bool, error) {]
ff=是否存在
key=名称
ctx=上下文

[func (c *AdapterRedis) Size(ctx context.Context) (size int, err error) {]
ff=取数量
err=错误
size=数量
ctx=上下文

[func (c *AdapterRedis) Data(ctx context.Context) (map#左中括号#interface{}#右中括号#interface{}, error) {]
ff=取所有键值Map副本
ctx=上下文

[func (c *AdapterRedis) Keys(ctx context.Context) (#左中括号##右中括号#interface{}, error) {]
ff=取所有键
ctx=上下文

[func (c *AdapterRedis) Values(ctx context.Context) (#左中括号##右中括号#interface{}, error) {]
ff=取所有值
ctx=上下文

[func (c *AdapterRedis) Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error) {]
ff=更新值
oldValue=旧值
value=值
key=名称
ctx=上下文

[func (c *AdapterRedis) UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error) {]
ff=更新过期时间
err=错误
oldDuration=旧过期时长
duration=时长
key=名称
ctx=上下文

[func (c *AdapterRedis) GetExpire(ctx context.Context, key interface{}) (time.Duration, error) {]
ff=取过期时间
key=名称
ctx=上下文

[func (c *AdapterRedis) Remove(ctx context.Context, keys ...interface{}) (lastValue *gvar.Var, err error) {]
ff=删除并带返回值
lastValue=最后一个删除值
keys=名称s
ctx=上下文

[func (c *AdapterRedis) Clear(ctx context.Context) (err error) {]
ff=清空
err=错误
ctx=上下文

[func (c *AdapterRedis) Close(ctx context.Context) error {]
ff=关闭
ctx=上下文
