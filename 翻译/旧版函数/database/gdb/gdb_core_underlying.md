
# <翻译开始>
) Query(ctx context.Context, sql string, args ...interface{}) (result Result, err
错误
# <翻译结束>

# <翻译开始>
) Query(ctx context.Context, sql string, args ...interface{}) (result
结果
# <翻译结束>

# <翻译开始>
) Query(ctx context.Context, sql string, args
参数
# <翻译结束>

# <翻译开始>
) Query(ctx
上下文
# <翻译结束>

# <翻译开始>
) Query
X原生SQL查询
# <翻译结束>

# <翻译开始>
) DoQuery(ctx context.Context, link Link, sql string, args ...interface{}) (result Result, err
错误
# <翻译结束>

# <翻译开始>
) DoQuery(ctx context.Context, link Link, sql string, args ...interface{}) (result
结果
# <翻译结束>

# <翻译开始>
) DoQuery(ctx context.Context, link Link, sql string, args
参数
# <翻译结束>

# <翻译开始>
) DoQuery(ctx context.Context, link
链接
# <翻译结束>

# <翻译开始>
) DoQuery(ctx
上下文
# <翻译结束>

# <翻译开始>
) DoQuery
X底层原生SQL查询
# <翻译结束>

# <翻译开始>
) Exec(ctx context.Context, sql string, args ...interface{}) (result sql.Result, err
错误
# <翻译结束>

# <翻译开始>
) Exec(ctx context.Context, sql string, args ...interface{}) (result
结果
# <翻译结束>

# <翻译开始>
) Exec(ctx context.Context, sql string, args
参数
# <翻译结束>

# <翻译开始>
) Exec(ctx
上下文
# <翻译结束>

# <翻译开始>
) Exec
X原生SQL执行
# <翻译结束>

# <翻译开始>
) DoExec(ctx context.Context, link Link, sql string, args ...interface{}) (result sql.Result, err
错误
# <翻译结束>

# <翻译开始>
) DoExec(ctx context.Context, link Link, sql string, args ...interface{}) (result
结果
# <翻译结束>

# <翻译开始>
) DoExec(ctx context.Context, link Link, sql string, args
参数
# <翻译结束>

# <翻译开始>
) DoExec(ctx context.Context, link
链接
# <翻译结束>

# <翻译开始>
) DoExec(ctx
上下文
# <翻译结束>

# <翻译开始>
) DoExec
X底层原生SQL执行
# <翻译结束>

# <翻译开始>
) DoFilter
X底层DoFilter
# <翻译结束>

# <翻译开始>
) DoCommit
X底层DoCommit
# <翻译结束>

# <翻译开始>
) Prepare(ctx context.Context, sql string, execOnMaster
是否主节点执行
# <翻译结束>

# <翻译开始>
) Prepare(ctx
上下文
# <翻译结束>

# <翻译开始>
) Prepare
X原生sql取参数预处理对象
# <翻译结束>

# <翻译开始>
) DoPrepare(ctx context.Context, link Link, sql string) (stmt *Stmt, err
错误
# <翻译结束>

# <翻译开始>
) DoPrepare(ctx context.Context, link Link, sql string) (stmt
参数预处理
# <翻译结束>

# <翻译开始>
 DoPrepare(ctx context.Context, link
链接
# <翻译结束>

# <翻译开始>
) DoPrepare(ctx
上下文
# <翻译结束>

# <翻译开始>
) DoPrepare
X底层原生sql参数预处理对象
# <翻译结束>

# <翻译开始>
) RowsToResult(ctx context.Context, rows
底层数据记录
# <翻译结束>

# <翻译开始>
) RowsToResult(ctx
上下文
# <翻译结束>

# <翻译开始>
) RowsToResult
X原生sql记录到行记录切片对象
# <翻译结束>
