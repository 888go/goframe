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
# re.F.String()
# }
# //zj:
# 备注结束

[func (c *Core) GetFieldTypeStr(ctx context.Context, fieldName, table, schema string) string {]
ff=取字段类型
schema=数据库名称
table=表名称
fieldName=字段名称
ctx=上下文

[func (c *Core) GetFieldType(ctx context.Context, fieldName, table, schema string) *TableField {]
ff=取字段信息对象
schema=数据库名称
table=表名称
fieldName=字段名称
ctx=上下文

[func (c *Core) ConvertDataForRecord(ctx context.Context, value interface{}, table string) (map#左中括号#string#右中括号#interface{}, error) {]
ff=底层ConvertDataForRecord
table=表名称
value=值
ctx=上下文

[func (c *Core) ConvertValueForField(ctx context.Context, fieldType string, fieldValue interface{}) (interface{}, error) {]
ff=底层ConvertValueForField

[func (c *Core) CheckLocalTypeForField(ctx context.Context, fieldType string, fieldValue interface{}) (LocalType, error) {]
ff=底层CheckLocalTypeForField

[func (c *Core) ConvertValueForLocal(]
ff=底层ConvertValueForLocal
