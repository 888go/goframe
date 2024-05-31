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

[func Valid(data interface{}) bool {]
ff=是否为有效json
data=值

[func Marshal(v interface{}) (marshaledBytes #左中括号##右中括号#byte, err error) {]
ff=Marshal别名

[func MarshalIndent(v interface{}, prefix, indent string) (marshaledBytes #左中括号##右中括号#byte, err error) {]
ff=MarshalIndent别名

[func Unmarshal(data #左中括号##右中括号#byte, v interface{}) (err error) {]
ff=Unmarshal别名

[func Encode(value interface{}) (#左中括号##右中括号#byte, error) {]
ff=变量到json字节集
value=值

[func MustEncode(value interface{}) #左中括号##右中括号#byte {]
ff=变量到json字节集PANI
value=值

[func EncodeString(value interface{}) (string, error) {]
ff=变量到json文本
value=值

[func MustEncodeString(value interface{}) string {]
ff=变量到json文本PANI
value=值

[func Decode(data interface{}, options ...Options) (interface{}, error) {]
ff=Json格式到变量
options=选项
data=值

[func DecodeTo(data interface{}, v interface{}, options ...Options) (err error) {]
ff=Json格式到变量指针
err=错误
options=选项
v=变量指针
data=值

[func DecodeToJson(data interface{}, options ...Options) (*Json, error) {]
ff=解码到json
options=选项
data=值
