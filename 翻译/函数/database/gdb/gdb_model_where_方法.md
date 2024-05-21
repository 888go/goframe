# 备注开始
# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 如:
# //ff:取文本

# **_package.md 文件备注:
# bm= 包名,更换新的包名称 
# 如: 
# package gin //bm:gin类

# **_其他.md 文件备注:
# qm= 前面,跳转到前面进行重命名.文档内如果有多个相同的,会一起重命名.
# hm= 后面,跳转到后面进行重命名.文档内如果有多个相同的,会一起重命名.
# cz= 查找,配合前面/后面使用,
# 如:
# type Regexp struct {//qm:正则 cz:Regexp struct
#
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
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

[func (m *Model) Where(where interface{}, args ...interface{}) *Model {]
ff=条件
args=参数
where=条件

[func (m *Model) Wheref(format string, args ...interface{}) *Model {]
ff=条件格式化
args=参数
format=格式

[func (m *Model) WherePri(where interface{}, args ...interface{}) *Model {]
ff=条件并识别主键
args=参数
where=条件

[func (m *Model) WhereLT(column string, value interface{}) *Model {]
ff=条件小于
value=比较值
column=字段名

[func (m *Model) WhereLTE(column string, value interface{}) *Model {]
ff=条件小于等于
value=比较值
column=字段名

[func (m *Model) WhereGT(column string, value interface{}) *Model {]
ff=条件大于
value=比较值
column=字段名

[func (m *Model) WhereGTE(column string, value interface{}) *Model {]
ff=条件大于等于
value=比较值
column=字段名

[func (m *Model) WhereBetween(column string, min, max interface{}) *Model {]
ff=条件取范围
max=最大值
min=最小值
column=字段名

[func (m *Model) WhereLike(column string, like string) *Model {]
ff=条件模糊匹配
like=通配符条件值
column=字段名

[func (m *Model) WhereIn(column string, in interface{}) *Model {]
ff=条件包含
in=包含值
column=字段名

[func (m *Model) WhereNull(columns ...string) *Model {]
ff=条件NULL值
columns=字段名

[func (m *Model) WhereNotBetween(column string, min, max interface{}) *Model {]
ff=条件取范围以外
max=最大值
min=最小值
column=字段名

[func (m *Model) WhereNotLike(column string, like interface{}) *Model {]
ff=条件模糊匹配以外
like=通配符条件值
column=字段名

[func (m *Model) WhereNot(column string, value interface{}) *Model {]
ff=条件不等于
value=值
column=字段名

[func (m *Model) WhereNotIn(column string, in interface{}) *Model {]
ff=条件不包含
in=不包含值
column=字段名

[func (m *Model) WhereNotNull(columns ...string) *Model {]
ff=条件非Null
columns=字段名
