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

[func (m *Model) WhereOr(where interface{}, args ...interface{}) *Model {]
ff=条件或
args=参数
where=条件

[func (m *Model) WhereOrf(format string, args ...interface{}) *Model {]
ff=条件或格式化
args=参数
format=格式

[func (m *Model) WhereOrLT(column string, value interface{}) *Model {]
ff=条件或小于
value=比较值
column=字段名

[func (m *Model) WhereOrLTE(column string, value interface{}) *Model {]
ff=条件或小于等于
value=比较值
column=字段名

[func (m *Model) WhereOrGT(column string, value interface{}) *Model {]
ff=条件或大于
value=比较值
column=字段名

[func (m *Model) WhereOrGTE(column string, value interface{}) *Model {]
ff=条件或大于等于
value=比较值
column=字段名

[func (m *Model) WhereOrBetween(column string, min, max interface{}) *Model {]
ff=条件或取范围
max=最大值
min=最小值
column=字段名

[func (m *Model) WhereOrLike(column string, like interface{}) *Model {]
ff=条件或模糊匹配
like=通配符条件值
column=字段名

[func (m *Model) WhereOrIn(column string, in interface{}) *Model {]
ff=条件或包含
in=包含值
column=字段名

[func (m *Model) WhereOrNull(columns ...string) *Model {]
ff=条件或NULL值
columns=字段名

[func (m *Model) WhereOrNotBetween(column string, min, max interface{}) *Model {]
ff=条件或取范围以外
max=最大值
min=最小值
column=字段名

[func (m *Model) WhereOrNotLike(column string, like interface{}) *Model {]
ff=条件或模糊匹配以外
like=通配符条件值
column=字段名

[func (m *Model) WhereOrNot(column string, value interface{}) *Model {]
ff=条件或不等于
value=值
column=字段名

[func (m *Model) WhereOrNotIn(column string, in interface{}) *Model {]
ff=条件或不包含
in=不包含值
column=字段名

[func (m *Model) WhereOrNotNull(columns ...string) *Model {]
ff=条件或非Null
columns=字段名
