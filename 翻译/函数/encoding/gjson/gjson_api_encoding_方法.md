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

[func (j *Json) ToJson() (#左中括号##右中括号#byte, error) {]
ff=取json字节集

[func (j *Json) ToJsonString() (string, error) {]
ff=取json文本

[func (j *Json) ToJsonIndent() (#左中括号##右中括号#byte, error) {]
ff=取json字节集并格式化

[func (j *Json) ToJsonIndentString() (string, error) {]
ff=取json文本并格式化

[func (j *Json) MustToJson() #左中括号##右中括号#byte {]
ff=取json字节集PANI

[func (j *Json) MustToJsonString() string {]
ff=取json文本PANI

[func (j *Json) MustToJsonIndent() #左中括号##右中括号#byte {]
ff=取json字节集并格式化PANI

[func (j *Json) MustToJsonIndentString() string {]
ff=取json文本并格式化PANI

[func (j *Json) ToXml(rootTag ...string) (#左中括号##右中括号#byte, error) {]
ff=取xml字节集

[func (j *Json) ToXmlString(rootTag ...string) (string, error) {]
ff=取xml文本

[func (j *Json) ToXmlIndent(rootTag ...string) (#左中括号##右中括号#byte, error) {]
ff=取xml字节集并格式化

[func (j *Json) ToXmlIndentString(rootTag ...string) (string, error) {]
ff=取xml文本并格式化

[func (j *Json) MustToXml(rootTag ...string) #左中括号##右中括号#byte {]
ff=取xml字节集PANI

[func (j *Json) MustToXmlString(rootTag ...string) string {]
ff=取xml文本PANI

[func (j *Json) MustToXmlIndent(rootTag ...string) #左中括号##右中括号#byte {]
ff=取xml字节集并格式化PANI

[func (j *Json) MustToXmlIndentString(rootTag ...string) string {]
ff=取xml文本并格式化PANI

[func (j *Json) ToYaml() (#左中括号##右中括号#byte, error) {]
ff=取YAML字节集

[func (j *Json) ToYamlIndent(indent string) (#左中括号##右中括号#byte, error) {]
ff=取YAML字节集并格式化
indent=缩进

[func (j *Json) ToYamlString() (string, error) {]
ff=取YAML文本

[func (j *Json) MustToYaml() #左中括号##右中括号#byte {]
ff=取YAML字节集PANI

[func (j *Json) MustToYamlString() string {]
ff=取YAML文本PANI

[func (j *Json) ToToml() (#左中括号##右中括号#byte, error) {]
ff=取TOML字节集

[func (j *Json) ToTomlString() (string, error) {]
ff=取TOML文本

[func (j *Json) MustToToml() #左中括号##右中括号#byte {]
ff=取TOML字节集PANI

[func (j *Json) MustToTomlString() string {]
ff=取TOML文本PANI

[func (j *Json) ToIni() (#左中括号##右中括号#byte, error) {]
ff=取ini字节集

[func (j *Json) ToIniString() (string, error) {]
ff=取ini文本

[func (j *Json) MustToIni() #左中括号##右中括号#byte {]
ff=取ini字节集PANI

[func (j *Json) MustToIniString() string {]
ff=取ini文本PANI

[func (j *Json) ToProperties() (#左中括号##右中括号#byte, error) {]
ff=取properties字节集

[func (j *Json) ToPropertiesString() (string, error) {]
ff=取properties文本

[func (j *Json) MustToProperties() #左中括号##右中括号#byte {]
ff=取properties字节集PANI

[func (j *Json) MustToPropertiesString() string {]
ff=取properties文本PANI
