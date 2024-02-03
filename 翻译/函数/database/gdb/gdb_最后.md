
# <翻译开始>
type TX
X事务
# <翻译结束>

# <翻译开始>
type DoCommitInput
X输入
# <翻译结束>

# <翻译开始>
Type          string
X类型
<跳到行首>
<忽略重复>
# <翻译结束>

# <翻译开始>
type DoCommitOutput
X输出
# <翻译结束>

# <翻译开始>
Result    sql.Result
X原生sql行记录
<跳到行首>
# <翻译结束>

# <翻译开始>
Records   []
X行记录数组
<跳到行首>
# <翻译结束>

# <翻译开始>
Stmt      *Stmt
X参数预处理
<跳到行首>
# <翻译结束>

# <翻译开始>
RawResult interface{}
X底层结果
<跳到行首>
# <翻译结束>

# <翻译开始>
type Driver
X驱动
# <翻译结束>

# <翻译开始>
type Link
X底层链接
# <翻译结束>

# <翻译开始>
Sql           string
SQL字符串
<跳到行首>
<忽略重复>
# <翻译结束>

# <翻译开始>
Type          string
X类型
<跳到行首>
<忽略重复>
# <翻译结束>

# <翻译开始>
Args          []interface{}
SQL参数
<跳到行首>
<忽略重复>
# <翻译结束>

# <翻译开始>
Format        string
SQL格式化后
<跳到行首>
# <翻译结束>

# <翻译开始>
Error         error
X执行错误
<跳到行首>
# <翻译结束>

# <翻译开始>
Start         int64
X开始时间戳
<跳到行首>
# <翻译结束>

# <翻译开始>
End           int64
X结束时间戳
<跳到行首>
# <翻译结束>

# <翻译开始>
Group         string
X配置组名称
<跳到行首>
# <翻译结束>

# <翻译开始>
Schema        string
X架构名称
<跳到行首>
# <翻译结束>

# <翻译开始>
IsTransaction bool
X是否为事务
<跳到行首>
<忽略重复>
# <翻译结束>

# <翻译开始>
RowsAffected  int64
X影响行数
<跳到行首>
# <翻译结束>

# <翻译开始>
type DoInsertOption
X底层输入
# <翻译结束>

# <翻译开始>
type TableField
X字段信息
# <翻译结束>

# <翻译开始>
Index int
X排序
<跳到行首>
# <翻译结束>

# <翻译开始>
Name string
X名称
<跳到行首>
# <翻译结束>

# <翻译开始>
Type string
X类型
<跳到行首>
<忽略重复>
# <翻译结束>

# <翻译开始>
Null bool
X是否为null
<跳到行首>
# <翻译结束>

# <翻译开始>
Key string
X索引信息
<跳到行首>
# <翻译结束>

# <翻译开始>
Default interface{}
X字段默认值
<跳到行首>
# <翻译结束>

# <翻译开始>
Extra string
X额外
<跳到行首>
# <翻译结束>

# <翻译开始>
Comment string
X字段注释
<跳到行首>
# <翻译结束>

# <翻译开始>
type Counter
X增减
# <翻译结束>

# <翻译开始>
Field string
X字段名称
<跳到行首>
# <翻译结束>

# <翻译开始>
Value float64
X增减值
<跳到行首>
# <翻译结束>

# <翻译开始>
Raw string
X原生sql
<跳到行首>
# <翻译结束>

# <翻译开始>
Value =
X字段值
<跳到行首>
# <翻译结束>

# <翻译开始>
Record map[string]
X行记录
<跳到行首>
# <翻译结束>

# <翻译开始>
Result []
X行记录数组
<跳到行首>
# <翻译结束>

# <翻译开始>
type InsertOption
X插入选项
# <翻译结束>
