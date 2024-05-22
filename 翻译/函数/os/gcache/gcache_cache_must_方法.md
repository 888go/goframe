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

[func (c *Cache) MustGet(ctx context.Context, key interface{}) *gvar.Var {]
ff=取值PANI
key=名称
ctx=上下文

[func (c *Cache) MustGetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) *gvar.Var {]
ff=取值或设置值PANI
duration=时长
value=值
key=名称
ctx=上下文

[func (c *Cache) MustGetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) *gvar.Var {]
ff=取值或设置值_函数PANI
duration=时长
f=回调函数
key=名称
ctx=上下文

[func (c *Cache) MustGetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) *gvar.Var {]
ff=取值或设置值_并发安全函数PANI
duration=时长
f=回调函数
key=名称
ctx=上下文

[func (c *Cache) MustContains(ctx context.Context, key interface{}) bool {]
ff=是否存在PANI
key=名称
ctx=上下文

[func (c *Cache) MustGetExpire(ctx context.Context, key interface{}) time.Duration {]
ff=取过期时间PANI
key=名称
ctx=上下文

[func (c *Cache) MustSize(ctx context.Context) int {]
ff=取数量PANI
ctx=上下文

[func (c *Cache) MustData(ctx context.Context) map#左中括号#interface{}#右中括号#interface{} {]
ff=取所有键值Map副本PANI
ctx=上下文

[func (c *Cache) MustKeys(ctx context.Context) #左中括号##右中括号#interface{} {]
ff=取所有键PANI
ctx=上下文

[func (c *Cache) MustKeyStrings(ctx context.Context) #左中括号##右中括号#string {]
ff=取所有键文本PANI
ctx=上下文

[func (c *Cache) MustValues(ctx context.Context) #左中括号##右中括号#interface{} {]
ff=取所有值PANI
ctx=上下文
