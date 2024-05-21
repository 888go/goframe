
# <翻译开始>
) SetMap(ctx context.Context, data map[interface{}]interface{}, duration
时长
# <翻译结束>

# <翻译开始>
) SetMap(ctx context.Context, data
值
# <翻译结束>

# <翻译开始>
) SetMap(ctx
上下文
# <翻译结束>

# <翻译开始>
) SetMap
X设置Map
# <翻译结束>

# <翻译开始>
) X设置值(ctx context.Context, key interface{}, value interface{}, duration
时长
# <翻译结束>

# <翻译开始>
) X设置值(ctx context.Context, key interface{}, value
值
# <翻译结束>

# <翻译开始>
) X设置值(ctx context.Context, key
名称
# <翻译结束>

# <翻译开始>
) X设置值(ctx
上下文
# <翻译结束>

# <翻译开始>
) SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration
时长
# <翻译结束>

# <翻译开始>
) SetIfNotExist(ctx context.Context, key interface{}, value
值
# <翻译结束>

# <翻译开始>
) SetIfNotExist(ctx context.Context, key
名称
# <翻译结束>

# <翻译开始>
) SetIfNotExist(ctx
上下文
# <翻译结束>

# <翻译开始>
) SetIfNotExist
X设置值并跳过已存在
# <翻译结束>

# <翻译开始>
) SetIfNotExistFunc(ctx context.Context, key interface{}, f Func, duration
时长
# <翻译结束>

# <翻译开始>
) SetIfNotExistFunc(ctx context.Context, key interface{}, f
回调函数
# <翻译结束>

# <翻译开始>
) SetIfNotExistFunc(ctx context.Context, key
名称
# <翻译结束>

# <翻译开始>
) SetIfNotExistFunc(ctx
上下文
# <翻译结束>

# <翻译开始>
) SetIfNotExistFunc
X设置值并跳过已存在_函数
# <翻译结束>

# <翻译开始>
) SetIfNotExistFuncLock(ctx context.Context, key interface{}, f Func, duration
时长
# <翻译结束>

# <翻译开始>
) SetIfNotExistFuncLock(ctx context.Context, key interface{}, f
回调函数
# <翻译结束>

# <翻译开始>
) SetIfNotExistFuncLock(ctx context.Context, key
名称
# <翻译结束>

# <翻译开始>
) SetIfNotExistFuncLock(ctx
上下文
# <翻译结束>

# <翻译开始>
) SetIfNotExistFuncLock
X设置值并跳过已存在_并发安全函数
# <翻译结束>

# <翻译开始>
) Get(ctx context.Context, key
名称
# <翻译结束>

# <翻译开始>
) Get(ctx
上下文
# <翻译结束>

# <翻译开始>
) Get
X取值
# <翻译结束>

# <翻译开始>
) GetOrSet(ctx context.Context, key interface{}, value interface{}, duration
时长
# <翻译结束>

# <翻译开始>
) GetOrSet(ctx context.Context, key interface{}, value
值
# <翻译结束>

# <翻译开始>
) GetOrSet(ctx context.Context, key
名称
# <翻译结束>

# <翻译开始>
) GetOrSet(ctx
上下文
# <翻译结束>

# <翻译开始>
) GetOrSet
X取值或设置值
# <翻译结束>

# <翻译开始>
) GetOrSetFunc(ctx context.Context, key interface{}, f Func, duration
时长
# <翻译结束>

# <翻译开始>
) GetOrSetFunc(ctx context.Context, key interface{}, f
回调函数
# <翻译结束>

# <翻译开始>
) GetOrSetFunc(ctx context.Context, key
名称
# <翻译结束>

# <翻译开始>
) GetOrSetFunc(ctx
上下文
# <翻译结束>

# <翻译开始>
) GetOrSetFunc
X取值或设置值_函数
# <翻译结束>

# <翻译开始>
) GetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration
时长
# <翻译结束>

# <翻译开始>
) GetOrSetFuncLock(ctx context.Context, key interface{}, f
回调函数
# <翻译结束>

# <翻译开始>
) GetOrSetFuncLock(ctx context.Context, key
名称
# <翻译结束>

# <翻译开始>
) GetOrSetFuncLock(ctx
上下文
# <翻译结束>

# <翻译开始>
) GetOrSetFuncLock
X取值或设置值_并发安全函数
# <翻译结束>

# <翻译开始>
) Contains(ctx context.Context, key
名称
# <翻译结束>

# <翻译开始>
) Contains(ctx
上下文
# <翻译结束>

# <翻译开始>
) Contains
X是否存在
# <翻译结束>

# <翻译开始>
) GetExpire(ctx context.Context, key
名称
# <翻译结束>

# <翻译开始>
) GetExpire(ctx
上下文
# <翻译结束>

# <翻译开始>
) GetExpire
X取过期时间
# <翻译结束>

# <翻译开始>
) Remove(ctx context.Context, keys
名称s
# <翻译结束>

# <翻译开始>
) Remove(ctx
上下文
# <翻译结束>

# <翻译开始>
) Remove
X删除并带返回值
# <翻译结束>

# <翻译开始>
) Update(ctx context.Context, key interface{}, value interface{}) (oldValue *泛型类.Var, exist bool, err
错误
# <翻译结束>

# <翻译开始>
) Update(ctx context.Context, key interface{}, value interface{}) (oldValue *泛型类.Var, exist
是否已存在
# <翻译结束>

# <翻译开始>
) Update(ctx context.Context, key interface{}, value interface{}) (oldValue
旧值
# <翻译结束>

# <翻译开始>
) Update(ctx context.Context, key interface{}, value
值
# <翻译结束>

# <翻译开始>
) Update(ctx context.Context, key
名称
# <翻译结束>

# <翻译开始>
) Update(ctx
上下文
# <翻译结束>

# <翻译开始>
) Update
X更新值
# <翻译结束>

# <翻译开始>
) UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err
错误
# <翻译结束>

# <翻译开始>
) UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration
旧过期时长
# <翻译结束>

# <翻译开始>
) UpdateExpire(ctx context.Context, key interface{}, duration
时长
# <翻译结束>

# <翻译开始>
) UpdateExpire(ctx context.Context, key
名称
# <翻译结束>

# <翻译开始>
) UpdateExpire(ctx
上下文
# <翻译结束>

# <翻译开始>
) UpdateExpire
X更新过期时间
# <翻译结束>

# <翻译开始>
) Size(ctx context.Context) (size int, err
错误
# <翻译结束>

# <翻译开始>
) Size(ctx context.Context) (size
数量
# <翻译结束>

# <翻译开始>
) Size(ctx
上下文
# <翻译结束>

# <翻译开始>
) Size
X取数量
# <翻译结束>

# <翻译开始>
) Data(ctx
上下文
# <翻译结束>

# <翻译开始>
) Data
X取所有键值Map副本
# <翻译结束>

# <翻译开始>
) Keys(ctx
上下文
# <翻译结束>

# <翻译开始>
) Keys
X取所有键
# <翻译结束>

# <翻译开始>
) Values(ctx
上下文
# <翻译结束>

# <翻译开始>
) Values
X取所有值
# <翻译结束>

# <翻译开始>
) Clear
X清空
# <翻译结束>

# <翻译开始>
) Close
X关闭
# <翻译结束>

# <翻译开始>
func NewAdapterMemory(lruCap
淘汰数量
# <翻译结束>

# <翻译开始>
func NewAdapterMemory
X创建内存适配器
# <翻译结束>

# <翻译开始>
) X清空(ctx
上下文
# <翻译结束>

# <翻译开始>
) X关闭(ctx
上下文
# <翻译结束>
