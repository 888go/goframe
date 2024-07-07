
# <翻译开始>
) GetCore
X取Core对象
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
) GetCtxTimeout(ctx context.Context, timeoutType
超时类型
# <翻译结束>

# <翻译开始>
) GetCtxTimeout(ctx
上下文
# <翻译结束>

# <翻译开始>
) GetCtxTimeout
X取超时上下文对象
# <翻译结束>

# <翻译开始>
) Close(ctx context.Context) (err
错误
# <翻译结束>

# <翻译开始>
) Close(ctx
上下文
# <翻译结束>

# <翻译开始>
) Close
X关闭数据库
# <翻译结束>

# <翻译开始>
) Master(schema
数据库名称
# <翻译结束>

# <翻译开始>
) Master
X取主节点对象
# <翻译结束>

# <翻译开始>
) Slave(schema
数据库名称
# <翻译结束>

# <翻译开始>
) Slave
X取从节点对象
# <翻译结束>

# <翻译开始>
) GetAll(ctx context.Context, sql string, args
参数
# <翻译结束>

# <翻译开始>
) GetAll(ctx
上下文
# <翻译结束>

# <翻译开始>
) GetAll
GetAll别名
# <翻译结束>

# <翻译开始>
) DoSelect(ctx context.Context, link Link, sql string, args ...interface{}) (result Result, err
错误
# <翻译结束>

# <翻译开始>
) DoSelect(ctx context.Context, link Link, sql string, args ...interface{}) (result
结果
# <翻译结束>

# <翻译开始>
) DoSelect(ctx context.Context, link Link, sql string, args
参数
# <翻译结束>

# <翻译开始>
) DoSelect(ctx context.Context, link
链接
# <翻译结束>

# <翻译开始>
) DoSelect(ctx
上下文
# <翻译结束>

# <翻译开始>
) DoSelect
X底层查询
# <翻译结束>

# <翻译开始>
) GetOne(ctx context.Context, sql string, args
参数
# <翻译结束>

# <翻译开始>
) GetOne(ctx
上下文
# <翻译结束>

# <翻译开始>
) GetOne
X原生SQL查询单条记录
# <翻译结束>

# <翻译开始>
) GetArray(ctx context.Context, sql string, args
参数
# <翻译结束>

# <翻译开始>
) GetArray(ctx
上下文
# <翻译结束>

# <翻译开始>
) GetArray
X原生SQL查询切片
# <翻译结束>

# <翻译开始>
) GetScan(ctx context.Context, pointer interface{}, sql string, args
参数
# <翻译结束>

# <翻译开始>
) GetScan(ctx context.Context, pointer
结构体指针
# <翻译结束>

# <翻译开始>
) GetScan(ctx
上下文
# <翻译结束>

# <翻译开始>
) GetScan
X原生SQL查询到结构体指针
# <翻译结束>

# <翻译开始>
) GetValue(ctx context.Context, sql string, args
参数
# <翻译结束>

# <翻译开始>
) GetValue(ctx
上下文
# <翻译结束>

# <翻译开始>
) GetValue
X原生SQL查询字段值
# <翻译结束>

# <翻译开始>
) GetCount(ctx context.Context, sql string, args
参数
# <翻译结束>

# <翻译开始>
) GetCount(ctx
上下文
# <翻译结束>

# <翻译开始>
) GetCount
X原生SQL查询字段计数
# <翻译结束>

# <翻译开始>
) Union(unions
Model对象
# <翻译结束>

# <翻译开始>
) Union
X多表去重查询
# <翻译结束>

# <翻译开始>
) UnionAll(unions
Model对象
# <翻译结束>

# <翻译开始>
) UnionAll
X多表查询
# <翻译结束>

# <翻译开始>
) PingMaster
X向主节点发送心跳
# <翻译结束>

# <翻译开始>
) PingSlave
X向从节点发送心跳
# <翻译结束>

# <翻译开始>
) Insert(ctx context.Context, table string, data interface{}, batch
批量操作行数
# <翻译结束>

# <翻译开始>
) Insert(ctx context.Context, table string, data
值
# <翻译结束>

# <翻译开始>
) Insert(ctx context.Context, table
表名称
# <翻译结束>

# <翻译开始>
) Insert(ctx
上下文
# <翻译结束>

# <翻译开始>
) Insert
X插入
# <翻译结束>

# <翻译开始>
) InsertIgnore(ctx context.Context, table string, data interface{}, batch
批量操作行数
# <翻译结束>

# <翻译开始>
) InsertIgnore(ctx context.Context, table string, data
值
# <翻译结束>

# <翻译开始>
) InsertIgnore(ctx context.Context, table
表名称
# <翻译结束>

# <翻译开始>
) InsertIgnore(ctx
上下文
# <翻译结束>

# <翻译开始>
) InsertIgnore
X插入并跳过已存在
# <翻译结束>

# <翻译开始>
) InsertAndGetId(ctx context.Context, table string, data interface{}, batch
批量操作行数
# <翻译结束>

# <翻译开始>
) InsertAndGetId(ctx context.Context, table string, data
值
# <翻译结束>

# <翻译开始>
) InsertAndGetId(ctx context.Context, table
表名称
# <翻译结束>

# <翻译开始>
) InsertAndGetId(ctx
上下文
# <翻译结束>

# <翻译开始>
) InsertAndGetId
X插入并取ID
# <翻译结束>

# <翻译开始>
) Replace(ctx context.Context, table string, data interface{}, batch
批量操作行数
# <翻译结束>

# <翻译开始>
) Replace(ctx context.Context, table string, data
值
# <翻译结束>

# <翻译开始>
) Replace(ctx context.Context, table
表名称
# <翻译结束>

# <翻译开始>
) Replace(ctx
上下文
# <翻译结束>

# <翻译开始>
) Replace
X插入并替换已存在
# <翻译结束>

# <翻译开始>
) Save(ctx context.Context, table string, data interface{}, batch
批量操作行数
# <翻译结束>

# <翻译开始>
) Save(ctx context.Context, table string, data
值
# <翻译结束>

# <翻译开始>
) Save(ctx context.Context, table
表名称
# <翻译结束>

# <翻译开始>
) Save(ctx
上下文
# <翻译结束>

# <翻译开始>
) Save
X插入并更新已存在
# <翻译结束>

# <翻译开始>
) DoInsert(ctx context.Context, link Link, table
表名称
# <翻译结束>

# <翻译开始>
) DoInsert(ctx context.Context, link
链接
# <翻译结束>

# <翻译开始>
) DoInsert(ctx
上下文
# <翻译结束>

# <翻译开始>
) DoInsert
X底层插入
# <翻译结束>

# <翻译开始>
) Update(ctx context.Context, table string, data interface{}, condition interface{}, args
参数
# <翻译结束>

# <翻译开始>
) Update(ctx context.Context, table string, data interface{}, condition
条件
# <翻译结束>

# <翻译开始>
) Update(ctx context.Context, table string, data
数据
# <翻译结束>

# <翻译开始>
) Update(ctx context.Context, table
表名称
# <翻译结束>

# <翻译开始>
) Update(ctx
上下文
# <翻译结束>

# <翻译开始>
) Update
X更新
# <翻译结束>

# <翻译开始>
) DoUpdate(ctx context.Context, link Link, table string, data interface{}, condition string, args
参数
# <翻译结束>

# <翻译开始>
) DoUpdate(ctx context.Context, link Link, table string, data interface{}, condition
条件
# <翻译结束>

# <翻译开始>
) DoUpdate(ctx context.Context, link Link, table string, data
值
# <翻译结束>

# <翻译开始>
) DoUpdate(ctx context.Context, link Link, table
表名称
# <翻译结束>

# <翻译开始>
) DoUpdate(ctx context.Context, link
链接
# <翻译结束>

# <翻译开始>
) DoUpdate(ctx
上下文
# <翻译结束>

# <翻译开始>
) DoUpdate
X底层更新
# <翻译结束>

# <翻译开始>
 Delete(ctx context.Context, table string, condition interface{}, args ...interface{}) (result sql.Result, err
错误
# <翻译结束>

# <翻译开始>
Delete(ctx context.Context, table string, condition interface{}, args ...interface{}) (result
结果
# <翻译结束>

# <翻译开始>
) Delete(ctx context.Context, table string, condition interface{}, args
参数
# <翻译结束>

# <翻译开始>
 Delete(ctx context.Context, table string, condition
条件
# <翻译结束>

# <翻译开始>
Delete(ctx context.Context, table
表名称
# <翻译结束>

# <翻译开始>
) Delete(ctx
上下文
# <翻译结束>

# <翻译开始>
) Delete
X删除
# <翻译结束>

# <翻译开始>
) DoDelete(ctx context.Context, link Link, table string, condition string, args ...interface{}) (result sql.Result, err
错误
# <翻译结束>

# <翻译开始>
) DoDelete(ctx context.Context, link Link, table string, condition string, args ...interface{}) (result
结果
# <翻译结束>

# <翻译开始>
) DoDelete(ctx context.Context, link Link, table string, condition string, args
参数
# <翻译结束>

# <翻译开始>
DoDelete(ctx context.Context, link Link, table string, condition
条件
# <翻译结束>

# <翻译开始>
DoDelete(ctx context.Context, link Link, table
表名称
# <翻译结束>

# <翻译开始>
DoDelete(ctx context.Context, link
链接
# <翻译结束>

# <翻译开始>
DoDelete(ctx
上下文
# <翻译结束>

# <翻译开始>
) DoDelete
X底层删除
# <翻译结束>

# <翻译开始>
) FilteredLink
X取数据库链接信息
# <翻译结束>

# <翻译开始>
) HasTable(name
表名称
# <翻译结束>

# <翻译开始>
) HasTable
X是否存在表名
# <翻译结束>

# <翻译开始>
) GetTablesWithCache
X取表名称缓存
# <翻译结束>

# <翻译开始>
FormatSqlBeforeExecuting(sql string, args []interface{}) (newSql string, newArgs
新参数切片
# <翻译结束>

# <翻译开始>
FormatSqlBeforeExecuting(sql string, args []interface{}) (newSql
新sql
# <翻译结束>

# <翻译开始>
) FormatSqlBeforeExecuting(sql string, args
参数切片
# <翻译结束>

# <翻译开始>
) FormatSqlBeforeExecuting
X格式化Sql
# <翻译结束>
