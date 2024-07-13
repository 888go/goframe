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

[func (c *Core) GetDB() DB {]
ff=取DB对象

[func (c *Core) GetLink(ctx context.Context, master bool, schema string) (Link, error) {]
ff=取数据库链接对象
master=主节点
ctx=上下文

[func (c *Core) MasterLink(schema ...string) (Link, error) {]
ff=底层MasterLink

[func (c *Core) SlaveLink(schema ...string) (Link, error) {]
ff=底层SlaveLink

[func (c *Core) QuoteWord(s string) string {]
ff=底层QuoteWord

[func (c *Core) QuoteString(s string) string {]
ff=底层QuoteString

[func (c *Core) QuotePrefixTableName(table string) string {]
ff=底层添加前缀字符和引用字符
table=表名称

[func (c *Core) GetChars() (charLeft string, charRight string) {]
ff=底层取数据库安全字符
charRight=右字符
charLeft=左字符

[func (c *Core) Tables(ctx context.Context, schema ...string) (tables #左中括号##右中括号#string, err error) {]
ff=取表名称切片
err=错误
tables=表名称切片
ctx=上下文

[func (c *Core) TableFields(ctx context.Context, table string, schema ...string) (fields map#左中括号#string#右中括号#*TableField, err error) {]
ff=取表字段信息Map
err=错误
fields=字段信息Map
table=表名称
ctx=上下文

[func (c *Core) ClearTableFields(ctx context.Context, table string, schema ...string) (err error) {]
ff=删除表字段缓存
err=错误
table=表名称
ctx=上下文

[func (c *Core) ClearTableFieldsAll(ctx context.Context) (err error) {]
ff=删除表字段所有缓存
err=错误
ctx=上下文

[func (c *Core) ClearCache(ctx context.Context, table string) (err error) {]
ff=删除表查询缓存
err=错误
table=表名称
ctx=上下文

[func (c *Core) ClearCacheAll(ctx context.Context) (err error) {]
ff=删除所有表查询缓存
err=错误
ctx=上下文

[func (c *Core) HasField(ctx context.Context, table, field string, schema ...string) (bool, error) {]
ff=是否存在字段
field=字段名称
table=表名称
ctx=上下文
