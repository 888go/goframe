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

[func (c *Core) Ctx(ctx context.Context) DB {]
ctx=上下文

[func (c *Core) GetCtxTimeout(ctx context.Context, timeoutType int) (context.Context, context.CancelFunc) {]
ff=取超时上下文对象
timeoutType=超时类型
ctx=上下文

[func (c *Core) Close(ctx context.Context) (err error) {]
err=错误
ctx=上下文

[func (c *Core) Master(schema ...string) (*sql.DB, error) {]
schema=数据库名称

[func (c *Core) Slave(schema ...string) (*sql.DB, error) {]
schema=数据库名称

[func (c *Core) GetAll(ctx context.Context, sql string, args ...interface{}) (Result, error) {]
args=参数
ctx=上下文

[func (c *Core) DoSelect(ctx context.Context, link Link, sql string, args ...interface{}) (result Result, err error) {]
err=错误
result=结果
args=参数
link=链接
ctx=上下文

[func (c *Core) GetOne(ctx context.Context, sql string, args ...interface{}) (Record, error) {]
args=参数
ctx=上下文

[func (c *Core) GetArray(ctx context.Context, sql string, args ...interface{}) (#左中括号##右中括号#Value, error) {]
args=参数
ctx=上下文

[func (c *Core) GetScan(ctx context.Context, pointer interface{}, sql string, args ...interface{}) error {]
args=参数
pointer=结构体指针
ctx=上下文

[func (c *Core) GetValue(ctx context.Context, sql string, args ...interface{}) (Value, error) {]
args=参数
ctx=上下文

[func (c *Core) GetCount(ctx context.Context, sql string, args ...interface{}) (int, error) {]
args=参数
ctx=上下文

[func (c *Core) Union(unions ...*Model) *Model {]
unions=Model对象

[func (c *Core) UnionAll(unions ...*Model) *Model {]
unions=Model对象

[func (c *Core) Insert(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error) {]
batch=批量操作行数
data=值
table=表名称
ctx=上下文

[func (c *Core) InsertIgnore(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error) {]
batch=批量操作行数
data=值
table=表名称
ctx=上下文

[func (c *Core) InsertAndGetId(ctx context.Context, table string, data interface{}, batch ...int) (int64, error) {]
batch=批量操作行数
data=值
table=表名称
ctx=上下文

[func (c *Core) Replace(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error) {]
batch=批量操作行数
data=值
table=表名称
ctx=上下文

[func (c *Core) Save(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error) {]
batch=批量操作行数
data=值
table=表名称
ctx=上下文

[func (c *Core) DoInsert(ctx context.Context, link Link, table string, list List, option DoInsertOption) (result sql.Result, err error) {]
table=表名称
link=链接
ctx=上下文

[func (c *Core) Update(ctx context.Context, table string, data interface{}, condition interface{}, args ...interface{}) (sql.Result, error) {]
args=参数
condition=条件
data=数据
table=表名称
ctx=上下文

[func (c *Core) DoUpdate(ctx context.Context, link Link, table string, data interface{}, condition string, args ...interface{}) (result sql.Result, err error) {]
args=参数
condition=条件
data=值
table=表名称
link=链接
ctx=上下文

[func (c *Core) Delete(ctx context.Context, table string, condition interface{}, args ...interface{}) (result sql.Result, err error) {]
err=错误
result=结果
args=参数
condition=条件
table=表名称
ctx=上下文

[func (c *Core) DoDelete(ctx context.Context, link Link, table string, condition string, args ...interface{}) (result sql.Result, err error) {]
err=错误
result=结果
args=参数
condition=条件
table=表名称
link=链接
ctx=上下文

[func (c *Core) FilteredLink() string {]
ff=取数据库链接信息

[func (c *Core) HasTable(name string) (bool, error) {]
ff=是否存在表名
name=表名称

[func (c *Core) GetTablesWithCache() (#左中括号##右中括号#string, error) {]
ff=取表名称缓存

[func (c *Core) FormatSqlBeforeExecuting(sql string, args #左中括号##右中括号#interface{}) (newSql string, newArgs #左中括号##右中括号#interface{}) {]
ff=格式化Sql
newArgs=新参数切片
newSql=新sql
args=参数切片
