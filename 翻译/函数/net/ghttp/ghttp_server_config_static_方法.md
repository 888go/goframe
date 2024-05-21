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

[func (s *Server) SetIndexFiles(indexFiles #左中括号##右中括号#string) {]
ff=设置静态文件索引
indexFiles=索引

[func (s *Server) GetIndexFiles() #左中括号##右中括号#string {]
ff=取静态文件索引

[func (s *Server) SetIndexFolder(enabled bool) {]
ff=设置静态文件是否列出子文件
enabled=是否

[func (s *Server) SetFileServerEnabled(enabled bool) {]
ff=设置静态文件是否开启
enabled=开启

[func (s *Server) SetServerRoot(root string) {]
ff=设置静态文件根目录
root=根目录

[func (s *Server) AddSearchPath(path string) {]
ff=静态文件添加额外搜索目录
path=目录

[func (s *Server) AddStaticPath(prefix string, path string) {]
ff=静态文件添加目录映射
path=新路径
prefix=旧路径
