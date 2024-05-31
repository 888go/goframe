
# <翻译开始>
) Begin(ctx context.Context) (tx TX, err
错误
# <翻译结束>

# <翻译开始>
) Begin(ctx context.Context) (tx
事务对象
# <翻译结束>

# <翻译开始>
) Begin(ctx
上下文
# <翻译结束>

# <翻译开始>
func (c *Core) Begin
X事务开启
# <翻译结束>

# <翻译开始>
func (c *Core) Transaction(ctx context.Context, f func(ctx context.Context, tx TX) error) (err
错误
# <翻译结束>

# <翻译开始>
func (c *Core) Transaction(ctx context.Context, f func(ctx context.Context, tx
事务对象
# <翻译结束>

# <翻译开始>
func (c *Core) Transaction(ctx context.Context, f func(ctx
上下文
# <翻译结束>

# <翻译开始>
func (c *Core) Transaction(ctx context.Context, f
回调函数
# <翻译结束>

# <翻译开始>
func (c *Core) Transaction(ctx
上下文
# <翻译结束>

# <翻译开始>
func (c *Core) Transaction
X事务
# <翻译结束>

# <翻译开始>
func WithTX(ctx context.Context, tx
事务对象
# <翻译结束>

# <翻译开始>
func WithTX(ctx
上下文
# <翻译结束>

# <翻译开始>
func WithTX
底层WithTX
# <翻译结束>

# <翻译开始>
func TXFromCtx(ctx
上下文
# <翻译结束>

# <翻译开始>
func TXFromCtx
X事务从上下文取对象
# <翻译结束>

# <翻译开始>
) Ctx(ctx
上下文
# <翻译结束>

# <翻译开始>
) Ctx
X设置上下文并取副本
# <翻译结束>

# <翻译开始>
) GetCtx
X取上下文对象
# <翻译结束>

# <翻译开始>
) GetDB
X取DB对象
# <翻译结束>

# <翻译开始>
) GetSqlTX
X底层取事务对象
# <翻译结束>

# <翻译开始>
) Commit
X事务提交
# <翻译结束>

# <翻译开始>
) Rollback
X事务回滚
# <翻译结束>

# <翻译开始>
) IsClosed
X是否已关闭
# <翻译结束>

# <翻译开始>
func (tx *TXCore) Begin
X事务开启
# <翻译结束>

# <翻译开始>
) SavePoint(point
事务点名称
# <翻译结束>

# <翻译开始>
) SavePoint
X保存事务点
# <翻译结束>

# <翻译开始>
) RollbackTo(point
事务点名称
# <翻译结束>

# <翻译开始>
) RollbackTo
X回滚事务点
# <翻译结束>

# <翻译开始>
func (tx *TXCore) Transaction(ctx context.Context, f func(ctx context.Context, tx TX) error) (err
错误
# <翻译结束>

# <翻译开始>
func (tx *TXCore) Transaction(ctx context.Context, f func(ctx context.Context, tx
事务对象
# <翻译结束>

# <翻译开始>
func (tx *TXCore) Transaction(ctx context.Context, f func(ctx
上下文
# <翻译结束>

# <翻译开始>
func (tx *TXCore) Transaction(ctx context.Context, f
回调函数
# <翻译结束>

# <翻译开始>
func (tx *TXCore) Transaction(ctx
上下文
# <翻译结束>

# <翻译开始>
func (tx *TXCore) Transaction
X事务
# <翻译结束>

# <翻译开始>
) Query(sql string, args ...interface{}) (result Result, err
错误
# <翻译结束>

# <翻译开始>
) Query(sql string, args ...interface{}) (result
结果
# <翻译结束>

# <翻译开始>
) Query(sql string, args
参数
# <翻译结束>

# <翻译开始>
) Query
X原生SQL查询
# <翻译结束>

# <翻译开始>
) Exec(sql string, args
参数
# <翻译结束>

# <翻译开始>
) Exec
X原生SQL执行
# <翻译结束>

# <翻译开始>
) Prepare
X原生sql取参数预处理对象
# <翻译结束>

# <翻译开始>
) GetAll(sql string, args
参数
# <翻译结束>

# <翻译开始>
) GetAll
GetAll别名
# <翻译结束>

# <翻译开始>
) GetOne(sql string, args
参数
# <翻译结束>

# <翻译开始>
) GetOne
X原生SQL查询单条记录
# <翻译结束>

# <翻译开始>
) GetStruct(obj interface{}, sql string, args
参数
# <翻译结束>

# <翻译开始>
) GetStruct(obj
结构体指针
# <翻译结束>

# <翻译开始>
) GetStruct
X原生SQL查询单条到结构体指针
# <翻译结束>

# <翻译开始>
) GetStructs(objPointerSlice interface{}, sql string, args
参数
# <翻译结束>

# <翻译开始>
) GetStructs(objPointerSlice
结构体指针
# <翻译结束>

# <翻译开始>
) GetStructs
X原生SQL查询到结构体数组指针
# <翻译结束>

# <翻译开始>
) GetScan(pointer interface{}, sql string, args
参数
# <翻译结束>

# <翻译开始>
) GetScan(pointer
结构体指针
# <翻译结束>

# <翻译开始>
) GetScan
X原生SQL查询到结构体指针
# <翻译结束>

# <翻译开始>
) GetValue(sql string, args
参数
# <翻译结束>

# <翻译开始>
) GetValue
X原生SQL查询字段值
# <翻译结束>

# <翻译开始>
) GetCount(sql string, args
参数
# <翻译结束>

# <翻译开始>
) GetCount
X原生SQL查询字段计数
# <翻译结束>

# <翻译开始>
) Insert(table string, data interface{}, batch
批量操作行数
# <翻译结束>

# <翻译开始>
) Insert(table string, data
值
# <翻译结束>

# <翻译开始>
) Insert(table
表名称
# <翻译结束>

# <翻译开始>
) Insert
X插入
# <翻译结束>

# <翻译开始>
) InsertIgnore(table string, data interface{}, batch
批量操作行数
# <翻译结束>

# <翻译开始>
) InsertIgnore(table string, data
值
# <翻译结束>

# <翻译开始>
) InsertIgnore(table
表名称
# <翻译结束>

# <翻译开始>
) InsertIgnore
X插入并跳过已存在
# <翻译结束>

# <翻译开始>
) InsertAndGetId(table string, data interface{}, batch
批量操作行数
# <翻译结束>

# <翻译开始>
) InsertAndGetId(table string, data
值
# <翻译结束>

# <翻译开始>
) InsertAndGetId(table
表名称
# <翻译结束>

# <翻译开始>
) InsertAndGetId
X插入并取ID
# <翻译结束>

# <翻译开始>
) Replace(table string, data interface{}, batch
批量操作行数
# <翻译结束>

# <翻译开始>
) Replace(table string, data
值
# <翻译结束>

# <翻译开始>
) Replace(table
表名称
# <翻译结束>

# <翻译开始>
) Replace
X插入并替换已存在
# <翻译结束>

# <翻译开始>
) Save(table string, data interface{}, batch
批量操作行数
# <翻译结束>

# <翻译开始>
) Save(table string, data
值
# <翻译结束>

# <翻译开始>
) Save(table
表名称
# <翻译结束>

# <翻译开始>
) Save
X插入并更新已存在
# <翻译结束>

# <翻译开始>
) Update(table string, data interface{}, condition interface{}, args
参数
# <翻译结束>

# <翻译开始>
) Update(table string, data interface{}, condition
条件
# <翻译结束>

# <翻译开始>
) Update(table string, data
值
# <翻译结束>

# <翻译开始>
) Update(table
表名称
# <翻译结束>

# <翻译开始>
) Update
X更新
# <翻译结束>

# <翻译开始>
) Delete(table string, condition interface{}, args
参数
# <翻译结束>

# <翻译开始>
) Delete(table string, condition
条件
# <翻译结束>

# <翻译开始>
) Delete(table
表名称
# <翻译结束>

# <翻译开始>
) Delete
X删除
# <翻译结束>
