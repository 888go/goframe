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
# zz= 正则查找,配合前面/后面使用, 有设置正则查找,就不用设置上面的查找
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
# //zj:前面一行的代码,如果为空,追加到末尾行
# func (re *Regexp) X取文本() string { 
# re.F.String()
# }
# //zj:
# 备注结束

[Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error]
qm=设置值
cz=Set(
yx=true

[SetMap(ctx context.Context, data map#左中括号#interface{}#右中括号#interface{}, duration time.Duration) error]
qm=设置Map
cz=SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) error

[SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (ok bool, err error)]
qm=设置值并跳过已存在
cz=SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (ok bool, err error)

[SetIfNotExistFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (ok bool, err error)]
qm=设置值并跳过已存在_函数
cz=SetIfNotExistFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (ok bool, err error)

[SetIfNotExistFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (ok bool, err error)]
qm=设置值并跳过已存在_并发安全函数
cz=SetIfNotExistFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (ok bool, err error)

[Get(ctx context.Context, key interface{}) (*gvar.Var, error)]
qm=取值
cz=Get(ctx context.Context, key interface{})

[GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (result *gvar.Var, err error)]
qm=取值或设置值
cz=GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration)

[GetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (result *gvar.Var, err error)]
qm=取值或设置值_函数
cz=GetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration)

[GetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (result *gvar.Var, err error)]
qm=取值或设置值_并发安全函数
cz=GetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration)

[Contains(ctx context.Context, key interface{}) (bool, error)]
qm=是否存在
cz=Contains(ctx context.Context, key interface{}) (bool, error)

[Size(ctx context.Context) (size int, err error)]
qm=取数量
cz=Size(ctx context.Context) (size int, err error)

[Data(ctx context.Context) (data map#左中括号#interface{}#右中括号#interface{}, err error)]
qm=取所有键值Map副本
cz=Data(ctx context.Context) (data map[interface{}]interface{}, err error)

[Keys(ctx context.Context) (keys #左中括号##右中括号#interface{}, err error)]
qm=取所有键
cz=Keys(ctx context.Context) (keys []interface{}, err error)

[Values(ctx context.Context) (values #左中括号##右中括号#interface{}, err error)]
qm=取所有值
cz=Values(ctx context.Context) (values []interface{}, err error)

[Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error)]
qm=更新值
cz=Update(ctx context.Context, key interface{}, value interface{})

[UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error)]
qm=更新过期时间
cz=UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error)

[GetExpire(ctx context.Context, key interface{}) (time.Duration, error)]
qm=取过期时间
cz=GetExpire(ctx context.Context, key interface{}) (time.Duration, error)

[Remove(ctx context.Context, keys ...interface{}) (lastValue *gvar.Var, err error)]
qm=删除并带返回值
cz=Remove(ctx context.Context, keys ...interface{})

[Clear(ctx context.Context) error]
qm=清空
cz=Clear(ctx context.Context) error

[Close(ctx context.Context) error]
qm=关闭
cz=Close(ctx context.Context) error
