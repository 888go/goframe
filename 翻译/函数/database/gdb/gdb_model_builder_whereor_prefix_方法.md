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

[func (b *WhereBuilder) WhereOrPrefix(prefix string, where interface{}, args ...interface{}) *WhereBuilder {]
ff=条件或并带前缀
args=参数
where=条件
prefix=字段前缀

[func (b *WhereBuilder) WhereOrPrefixNot(prefix string, column string, value interface{}) *WhereBuilder {]
ff=条件或不等于并带前缀
value=值
column=字段名
prefix=字段前缀

[func (b *WhereBuilder) WhereOrPrefixLT(prefix string, column string, value interface{}) *WhereBuilder {]
ff=条件或小于并带前缀
value=比较值
column=字段名
prefix=字段前缀

[func (b *WhereBuilder) WhereOrPrefixLTE(prefix string, column string, value interface{}) *WhereBuilder {]
ff=条件或小于等于并带前缀
value=比较值
column=字段名
prefix=字段前缀

[func (b *WhereBuilder) WhereOrPrefixGT(prefix string, column string, value interface{}) *WhereBuilder {]
ff=条件或大于并带前缀
value=比较值
column=字段名
prefix=字段前缀

[func (b *WhereBuilder) WhereOrPrefixGTE(prefix string, column string, value interface{}) *WhereBuilder {]
ff=条件或大于等于并带前缀
value=比较值
column=字段名
prefix=字段前缀

[func (b *WhereBuilder) WhereOrPrefixBetween(prefix string, column string, min, max interface{}) *WhereBuilder {]
ff=条件或取范围并带前缀
max=最大值
min=最小值
column=字段名
prefix=字段前缀

[func (b *WhereBuilder) WhereOrPrefixLike(prefix string, column string, like interface{}) *WhereBuilder {]
ff=条件或模糊匹配并带前缀
like=通配符条件值
column=字段名
prefix=字段前缀

[func (b *WhereBuilder) WhereOrPrefixIn(prefix string, column string, in interface{}) *WhereBuilder {]
ff=条件或包含并带前缀
in=包含值
column=字段名
prefix=字段前缀

[func (b *WhereBuilder) WhereOrPrefixNull(prefix string, columns ...string) *WhereBuilder {]
ff=条件或NULL值并带前缀
columns=字段名
prefix=字段前缀

[func (b *WhereBuilder) WhereOrPrefixNotBetween(prefix string, column string, min, max interface{}) *WhereBuilder {]
ff=条件或取范围以外并带前缀
max=最大值
min=最小值
column=字段名
prefix=字段前缀

[func (b *WhereBuilder) WhereOrPrefixNotLike(prefix string, column string, like interface{}) *WhereBuilder {]
ff=条件或模糊匹配以外并带前缀
like=通配符条件值
column=字段名
prefix=字段前缀

[func (b *WhereBuilder) WhereOrPrefixNotIn(prefix string, column string, in interface{}) *WhereBuilder {]
ff=条件或不包含并带前缀
in=不包含值
column=字段名
prefix=字段前缀

[func (b *WhereBuilder) WhereOrPrefixNotNull(prefix string, columns ...string) *WhereBuilder {]
ff=条件或非Null并带前缀
columns=字段名
prefix=字段前缀
