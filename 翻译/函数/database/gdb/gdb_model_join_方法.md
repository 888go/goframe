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

[func (m *Model) LeftJoin(tableOrSubQueryAndJoinConditions ...string) *Model {]
ff=左连接
tableOrSubQueryAndJoinConditions=表或子查询和连接条件

[func (m *Model) RightJoin(tableOrSubQueryAndJoinConditions ...string) *Model {]
ff=右连接
tableOrSubQueryAndJoinConditions=表或子查询和连接条件

[func (m *Model) InnerJoin(tableOrSubQueryAndJoinConditions ...string) *Model {]
ff=内连接
tableOrSubQueryAndJoinConditions=表或子查询和连接条件

[func (m *Model) LeftJoinOnField(table, field string) *Model {]
ff=左连接相同字段
field=相同字段名
table=表名

[func (m *Model) RightJoinOnField(table, field string) *Model {]
ff=右连接相同字段
field=相同字段名
table=表名

[func (m *Model) InnerJoinOnField(table, field string) *Model {]
ff=内连接相同字段
field=相同字段名
table=表名

[func (m *Model) LeftJoinOnFields(table, firstField, operator, secondField string) *Model {]
ff=左连接带比较运算符
secondField=第二个字段
operator=比较运算符
firstField=第一个字段
table=表名

[func (m *Model) RightJoinOnFields(table, firstField, operator, secondField string) *Model {]
ff=右连接带比较运算符
secondField=第二个字段
operator=比较运算符
firstField=第一个字段
table=表名

[func (m *Model) InnerJoinOnFields(table, firstField, operator, secondField string) *Model {]
ff=内连接带比较运算符
secondField=第二个字段
operator=比较运算符
firstField=第一个字段
table=表名
