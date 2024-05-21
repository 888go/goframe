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

[func (b *WhereBuilder) WhereOr(where interface{}, args ...interface{}) *WhereBuilder {]
ff=条件或
args=参数
where=条件

[func (b *WhereBuilder) WhereOrf(format string, args ...interface{}) *WhereBuilder {]
ff=条件或格式化
args=参数
format=格式

[func (b *WhereBuilder) WhereOrNot(column string, value interface{}) *WhereBuilder {]
ff=条件或不等于
value=值
column=字段名

[func (b *WhereBuilder) WhereOrLT(column string, value interface{}) *WhereBuilder {]
ff=条件或小于
value=比较值
column=字段名

[func (b *WhereBuilder) WhereOrLTE(column string, value interface{}) *WhereBuilder {]
ff=条件或小于等于
value=比较值
column=字段名

[func (b *WhereBuilder) WhereOrGT(column string, value interface{}) *WhereBuilder {]
ff=条件或大于
value=比较值
column=字段名

[func (b *WhereBuilder) WhereOrGTE(column string, value interface{}) *WhereBuilder {]
ff=条件或大于等于
value=比较值
column=字段名

[func (b *WhereBuilder) WhereOrBetween(column string, min, max interface{}) *WhereBuilder {]
ff=条件或取范围
max=最大值
min=最小值
column=字段名

[func (b *WhereBuilder) WhereOrLike(column string, like interface{}) *WhereBuilder {]
ff=条件或模糊匹配
like=通配符条件值
column=字段名

[func (b *WhereBuilder) WhereOrIn(column string, in interface{}) *WhereBuilder {]
ff=条件或包含
in=包含值
column=字段名

[func (b *WhereBuilder) WhereOrNull(columns ...string) *WhereBuilder {]
ff=条件或NULL值
columns=字段名

[func (b *WhereBuilder) WhereOrNotBetween(column string, min, max interface{}) *WhereBuilder {]
ff=条件或取范围以外
max=最大值
min=最小值
column=字段名

[func (b *WhereBuilder) WhereOrNotLike(column string, like interface{}) *WhereBuilder {]
ff=条件或模糊匹配以外
like=通配符条件值
column=字段名

[func (b *WhereBuilder) WhereOrNotIn(column string, in interface{}) *WhereBuilder {]
ff=条件或不包含
in=不包含值
column=字段名

[func (b *WhereBuilder) WhereOrNotNull(columns ...string) *WhereBuilder {]
ff=条件或非Null
columns=字段名
