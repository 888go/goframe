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

[func (c *Core) Begin(ctx context.Context) (tx TX, err error) {]
ff=事务开启
err=错误
tx=事务对象
ctx=上下文

[func (c *Core) Transaction(ctx context.Context, f func(ctx context.Context, tx TX) error) (err error) {]
ff=事务
err=错误
f=回调函数
ctx=上下文

[func WithTX(ctx context.Context, tx TX) context.Context {]
ff=底层WithTX
tx=事务对象
ctx=上下文

[func TXFromCtx(ctx context.Context, group string) TX {]
ff=事务从上下文取对象
ctx=上下文

[func (tx *TXCore) Ctx(ctx context.Context) TX {]
ff=设置上下文并取副本
ctx=上下文

[func (tx *TXCore) GetCtx() context.Context {]
ff=取上下文对象

[func (tx *TXCore) GetDB() DB {]
ff=取DB对象

[func (tx *TXCore) GetSqlTX() *sql.Tx {]
ff=底层取事务对象

[func (tx *TXCore) Commit() error {]
ff=事务提交

[func (tx *TXCore) Rollback() error {]
ff=事务回滚

[func (tx *TXCore) IsClosed() bool {]
ff=是否已关闭

[func (tx *TXCore) Begin() error {]
ff=事务开启

[func (tx *TXCore) SavePoint(point string) error {]
ff=保存事务点
point=事务点名称

[func (tx *TXCore) RollbackTo(point string) error {]
ff=回滚事务点
point=事务点名称

[func (tx *TXCore) Transaction(ctx context.Context, f func(ctx context.Context, tx TX) error) (err error) {]
ff=事务
err=错误
f=回调函数
ctx=上下文

[func (tx *TXCore) Query(sql string, args ...interface{}) (result Result, err error) {]
ff=原生SQL查询
err=错误
result=结果
args=参数

[func (tx *TXCore) Exec(sql string, args ...interface{}) (sql.Result, error) {]
ff=原生SQL执行
args=参数

[func (tx *TXCore) Prepare(sql string) (*Stmt, error) {]
ff=原生sql取参数预处理对象

[func (tx *TXCore) GetAll(sql string, args ...interface{}) (Result, error) {]
ff=GetAll别名
args=参数

[func (tx *TXCore) GetOne(sql string, args ...interface{}) (Record, error) {]
ff=原生SQL查询单条记录
args=参数

[func (tx *TXCore) GetStruct(obj interface{}, sql string, args ...interface{}) error {]
ff=原生SQL查询单条到结构体指针
args=参数
obj=结构体指针

[func (tx *TXCore) GetStructs(objPointerSlice interface{}, sql string, args ...interface{}) error {]
ff=原生SQL查询到结构体数组指针
args=参数
objPointerSlice=结构体指针

[func (tx *TXCore) GetScan(pointer interface{}, sql string, args ...interface{}) error {]
ff=原生SQL查询到结构体指针
args=参数
pointer=结构体指针

[func (tx *TXCore) GetValue(sql string, args ...interface{}) (Value, error) {]
ff=原生SQL查询字段值
args=参数

[func (tx *TXCore) GetCount(sql string, args ...interface{}) (int64, error) {]
ff=原生SQL查询字段计数
args=参数

[func (tx *TXCore) Insert(table string, data interface{}, batch ...int) (sql.Result, error) {]
ff=插入
batch=批量操作行数
data=值
table=表名称

[func (tx *TXCore) InsertIgnore(table string, data interface{}, batch ...int) (sql.Result, error) {]
ff=插入并跳过已存在
batch=批量操作行数
data=值
table=表名称

[func (tx *TXCore) InsertAndGetId(table string, data interface{}, batch ...int) (int64, error) {]
ff=插入并取ID
batch=批量操作行数
data=值
table=表名称

[func (tx *TXCore) Replace(table string, data interface{}, batch ...int) (sql.Result, error) {]
ff=插入并替换已存在
batch=批量操作行数
data=值
table=表名称

[func (tx *TXCore) Save(table string, data interface{}, batch ...int) (sql.Result, error) {]
ff=插入并更新已存在
batch=批量操作行数
data=值
table=表名称

[func (tx *TXCore) Update(table string, data interface{}, condition interface{}, args ...interface{}) (sql.Result, error) {]
ff=更新
args=参数
condition=条件
data=值
table=表名称

[func (tx *TXCore) Delete(table string, condition interface{}, args ...interface{}) (sql.Result, error) {]
ff=删除
args=参数
condition=条件
table=表名称
