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

[List = #左中括号##右中括号#Map]
qm=Map数组
cz=List #等号# []Map

[ListAnyAny = #左中括号##右中括号#MapAnyAny]
qm=MapAnyAny数组
cz=ListAnyAny #等号# []MapAnyAny

[ListAnyStr = #左中括号##右中括号#MapAnyStr]
qm=MapAnyStr数组
cz=ListAnyStr #等号# []MapAnyStr

[ListAnyInt = #左中括号##右中括号#MapAnyInt]
qm=MapAnyInt数组
cz=ListAnyInt #等号# []MapAnyInt

[ListStrAny = #左中括号##右中括号#MapStrAny]
qm=MapStrAny数组
cz=ListStrAny #等号# []MapStrAny

[ListStrStr = #左中括号##右中括号#MapStrStr]
qm=MapStrStr数组
cz=ListStrStr #等号# []MapStrStr

[ListStrInt = #左中括号##右中括号#MapStrInt]
qm=MapStrInt数组
cz=ListStrInt #等号# []MapStrInt

[ListIntAny = #左中括号##右中括号#MapIntAny]
qm=MapIntAny数组
cz=ListIntAny #等号# []MapIntAny

[ListIntStr = #左中括号##右中括号#MapIntStr]
qm=MapIntStr数组
cz=ListIntStr #等号# []MapIntStr

[ListIntInt = #左中括号##右中括号#MapIntInt]
qm=MapIntInt数组
cz=ListIntInt #等号# []MapIntInt

[ListAnyBool = #左中括号##右中括号#MapAnyBool]
qm=MapAnyBool数组
cz=ListAnyBool #等号# []MapAnyBool

[ListStrBool = #左中括号##右中括号#MapStrBool]
qm=MapStrBool数组
cz=ListStrBool #等号# []MapStrBool

[ListIntBool = #左中括号##右中括号#MapIntBool]
qm=MapIntBool数组
cz=ListIntBool #等号# []MapIntBool

[Slice = #左中括号##右中括号#interface{}]
qm=Slice别名
cz=Slice #等号# []interface{}

[SliceAny = #左中括号##右中括号#interface{}]
qm=SliceAny别名
cz=SliceAny #等号# []interface{}

[SliceStr = #左中括号##右中括号#string]
qm=SliceStr别名
cz=SliceStr #等号# []string

[SliceInt = #左中括号##右中括号#int]
qm=SliceInt别名
cz=SliceInt #等号# []int

[Array = #左中括号##右中括号#interface{}]
qm=数组
cz=Array #等号# []interface{}

[ArrayAny = #左中括号##右中括号#interface{}]
qm=Any数组
cz=ArrayAny #等号# []interface{}

[ArrayStr = #左中括号##右中括号#string]
qm=文本数组
cz=ArrayStr #等号# []string

[ArrayInt = #左中括号##右中括号#int]
qm=整数数组
cz=ArrayInt #等号# []int
