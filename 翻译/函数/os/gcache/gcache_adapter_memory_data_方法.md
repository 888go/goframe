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

[func (d *adapterMemoryData) Update(key interface{}, value interface{}) (oldValue interface{}, exist bool, err error) {]
ff=更新值
err=错误
exist=是否已存在
oldValue=旧值
value=值
key=名称

[func (d *adapterMemoryData) UpdateExpire(key interface{}, expireTime int64) (oldDuration time.Duration, err error) {]
ff=更新过期时间
err=错误
oldDuration=旧过期时长
expireTime=时长
key=名称

[func (d *adapterMemoryData) Remove(keys ...interface{}) (removedKeys #左中括号##右中括号#interface{}, value interface{}, err error) {]
ff=删除并带返回值
err=错误
value=值
removedKeys=被删除名称
keys=名称

[func (d *adapterMemoryData) Data() (map#左中括号#interface{}#右中括号#interface{}, error) {]
ff=取所有键值Map副本

[func (d *adapterMemoryData) Keys() (#左中括号##右中括号#interface{}, error) {]
ff=取所有键

[func (d *adapterMemoryData) Values() (#左中括号##右中括号#interface{}, error) {]
ff=取所有值

[func (d *adapterMemoryData) Size() (size int, err error) {]
ff=取数量
err=错误
size=数量

[func (d *adapterMemoryData) Clear() error {]
ff=清空

[func (d *adapterMemoryData) Get(key interface{}) (item adapterMemoryItem, ok bool) {]
ff=取值
ok=成功
key=名称

[func (d *adapterMemoryData) Set(key interface{}, value adapterMemoryItem) {]
ff=设置值
yx=true

[func (d *adapterMemoryData) SetMap(data map#左中括号#interface{}#右中括号#interface{}, expireTime int64) error {]
ff=设置Map
