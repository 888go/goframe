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

[func (c *Core) Query(ctx context.Context, sql string, args ...interface{}) (result Result, err error) {]
err=错误
result=结果
args=参数
ctx=上下文

[func (c *Core) DoQuery(ctx context.Context, link Link, sql string, args ...interface{}) (result Result, err error) {]
err=错误
result=结果
args=参数
link=链接
ctx=上下文

[func (c *Core) Exec(ctx context.Context, sql string, args ...interface{}) (result sql.Result, err error) {]
err=错误
result=结果
args=参数
ctx=上下文

[func (c *Core) DoExec(ctx context.Context, link Link, sql string, args ...interface{}) (result sql.Result, err error) {]
err=错误
result=结果
args=参数
link=链接
ctx=上下文

[func (c *Core) Prepare(ctx context.Context, sql string, execOnMaster ...bool) (*Stmt, error) {]
execOnMaster=是否主节点执行
ctx=上下文

[func (c *Core) DoPrepare(ctx context.Context, link Link, sql string) (stmt *Stmt, err error) {]
err=错误
stmt=参数预处理
link=链接
ctx=上下文

[func (c *Core) RowsToResult(ctx context.Context, rows *sql.Rows) (Result, error) {]
ff=原生sql记录到行记录切片对象
rows=底层数据记录
ctx=上下文
