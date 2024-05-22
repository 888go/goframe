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
#    re.F.String()
# }
# //zj:
# 备注结束

[func (r *Response) Write(content ...interface{}) {]
ff=写响应缓冲区
content=内容

[func (r *Response) WriteExit(content ...interface{}) {]
ff=写响应缓冲区并退出
content=内容

[func (r *Response) WriteOver(content ...interface{}) {]
ff=写覆盖响应缓冲区
content=内容

[func (r *Response) WriteOverExit(content ...interface{}) {]
ff=写覆盖响应缓冲区并退出
content=内容

[func (r *Response) Writef(format string, params ...interface{}) {]
ff=写响应缓冲区并格式化
params=内容
format=格式

[func (r *Response) WritefExit(format string, params ...interface{}) {]
ff=写响应缓冲区并退出与格式化
params=内容
format=格式

[func (r *Response) Writeln(content ...interface{}) {]
ff=写响应缓冲区并换行
content=内容

[func (r *Response) WritelnExit(content ...interface{}) {]
ff=写响应缓冲区并退出与换行
content=内容

[func (r *Response) Writefln(format string, params ...interface{}) {]
ff=写响应缓冲区并格式化与换行
params=内容
format=格式

[func (r *Response) WriteflnExit(format string, params ...interface{}) {]
ff=写响应缓冲区并退出与格式化换行
params=内容
format=格式

[func (r *Response) WriteJson(content interface{}) {]
ff=写响应缓冲区JSON
content=内容

[func (r *Response) WriteJsonExit(content interface{}) {]
ff=写响应缓冲区JSON并退出
content=内容

[func (r *Response) WriteJsonP(content interface{}) {]
ff=写响应缓冲区JSONP
content=内容

[func (r *Response) WriteJsonPExit(content interface{}) {]
ff=写响应缓冲区JSONP并退出
content=内容

[func (r *Response) WriteXml(content interface{}, rootTag ...string) {]
ff=写响应缓冲区XML
rootTag=根标记
content=内容

[func (r *Response) WriteXmlExit(content interface{}, rootTag ...string) {]
ff=写响应缓冲区XML并退出
rootTag=根标记
content=内容

[func (r *Response) WriteStatus(status int, content ...interface{}) {]
ff=写响应缓冲区与HTTP状态码
content=内容
status=状态码

[func (r *Response) WriteStatusExit(status int, content ...interface{}) {]
ff=写响应缓冲区与HTTP状态码并退出
content=内容
status=状态码
