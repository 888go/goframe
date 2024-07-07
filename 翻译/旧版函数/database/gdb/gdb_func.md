
# <翻译开始>
WithDB(ctx context.Context, db
DB对象
# <翻译结束>

# <翻译开始>
WithDB(ctx
上下文
# <翻译结束>

# <翻译开始>
func WithDB
底层WithDB
# <翻译结束>

# <翻译开始>
DBFromCtx(ctx
上下文
# <翻译结束>

# <翻译开始>
func DBFromCtx
X上下文取DB对象
# <翻译结束>

# <翻译开始>
func ToSQL(ctx context.Context, f func(ctx context.Context) error) (sql string, err
错误
# <翻译结束>

# <翻译开始>
func ToSQL(ctx context.Context, f func(ctx
上下文
# <翻译结束>

# <翻译开始>
func ToSQL(ctx context.Context, f
回调函数
# <翻译结束>

# <翻译开始>
func ToSQL(ctx
上下文
# <翻译结束>

# <翻译开始>
func ToSQL
X捕捉最后一条SQL语句
# <翻译结束>

# <翻译开始>
func CatchSQL(ctx context.Context, f func(ctx context.Context) error) (sqlArray []string, err
错误
# <翻译结束>

# <翻译开始>
func CatchSQL(ctx context.Context, f func(ctx context.Context) error) (sqlArray
sql切片
# <翻译结束>

# <翻译开始>
func CatchSQL(ctx context.Context, f func(ctx
上下文
# <翻译结束>

# <翻译开始>
func CatchSQL(ctx context.Context, f
回调函数
# <翻译结束>

# <翻译开始>
func CatchSQL(ctx
上下文
# <翻译结束>

# <翻译开始>
func CatchSQL
X捕捉SQL语句
# <翻译结束>

# <翻译开始>
func ListItemValues(list interface{}, key interface{}, subKey ...interface{}) (values
切片值
# <翻译结束>

# <翻译开始>
func ListItemValues(list interface{}, key interface{}, subKey
子名称
# <翻译结束>

# <翻译开始>
func ListItemValues(list interface{}, key
名称
# <翻译结束>

# <翻译开始>
func ListItemValues(list
结构体切片或Map切片
# <翻译结束>

# <翻译开始>
func ListItemValues
X取结构体切片或Map切片值
# <翻译结束>

# <翻译开始>
func ListItemValuesUnique(list interface{}, key string, subKey
子名称
# <翻译结束>

# <翻译开始>
func ListItemValuesUnique(list interface{}, key
名称
# <翻译结束>

# <翻译开始>
func ListItemValuesUnique(list
结构体切片或Map切片
# <翻译结束>

# <翻译开始>
func ListItemValuesUnique
X取结构体切片或Map切片值并去重
# <翻译结束>

# <翻译开始>
func GetInsertOperationByOption(option
选项
# <翻译结束>

# <翻译开始>
func GetInsertOperationByOption
底层GetInsertOperationByOption
# <翻译结束>

# <翻译开始>
func DaToMapDeep
作废DaToMapDeep
# <翻译结束>

# <翻译开始>
func MapOrStructToMapDeep(value
待转换值
# <翻译结束>

# <翻译开始>
func MapOrStructToMapDeep
X转换到Map
# <翻译结束>

# <翻译开始>
func GetPrimaryKeyCondition
底层GetPrimaryKeyCondition
# <翻译结束>

# <翻译开始>
func FormatSqlWithArgs(sql string, args
参数切片
# <翻译结束>

# <翻译开始>
func FormatSqlWithArgs
X格式化Sql
# <翻译结束>
