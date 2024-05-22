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

[Handlers             #左中括号##右中括号#Handler      `json:"-"`]
qm=中间件
cz=Handlers []Handler

[Flags                int            `json:"flags"`]
qm=日志标识
cz=Flags int `json:"flags"`

[TimeFormat           string         `json:"timeFormat"`]
qm=时间格式
cz=TimeFormat string `json:"timeFormat"`

[Path                 string         `json:"path"`]
qm=文件路径
cz=Path string `json:"path"`

[File                 string         `json:"file"`]
qm=文件名格式
cz=File string `json:"file"`

[Level                int            `json:"level"`]
qm=级别
cz=Level int `json:"level"

[Prefix               string         `json:"prefix"`]
qm=前缀
cz=Prefix string `json:"prefix"

[StSkip               int            `json:"stSkip"`]
qm=堆栈偏移量
cz=StSkip int `json:"stSkip"`

[StStatus             int            `json:"stStatus"`]
qm=堆栈状态
cz=StStatus int `json:"stStatus"`

[StFilter             string         `json:"stFilter"`]
qm=堆栈过滤
cz=StFilter string `json:"stFilter"`

[CtxKeys              #左中括号##右中括号#interface{}  `json:"ctxKeys"`]
qm=上下文名称
cz=CtxKeys []interface{} `json:"ctxKeys"`

[HeaderPrint          bool           `json:"header"`]
qm=是否输出头信息
cz=HeaderPrint bool `json:"header"`

[StdoutPrint          bool           `json:"stdout"`]
qm=是否同时输出到终端
cz=StdoutPrint bool `json:"stdout"`

[LevelPrint           bool           `json:"levelPrint"`]
qm=是否输出级别
cz=LevelPrint bool `json:"levelPrint"

[LevelPrefixes        map#左中括号#int#右中括号#string `json:"levelPrefixes"`]
qm=日志级别名称映射
cz=LevelPrefixes map[int]string `json:"levelPrefixes"`

[RotateSize           int64          `json:"rotateSize"`]
qm=文件分割大小
cz=RotateSize int64 `json:"rotateSize"

[RotateExpire         time.Duration  `json:"rotateExpire"`]
qm=文件分割周期
cz=RotateExpire time.Duration `json:"rotateExpire

[RotateBackupLimit    int            `json:"rotateBackupLimit"`]
qm=文件分割保留数量
cz=RotateBackupLimit int `json:"rotateBackupLimit"

[RotateBackupExpire   time.Duration  `json:"rotateBackupExpire"`]
qm=文件分割过期时间
cz=RotateBackupExpire time.Duration `json:"rotateBackupExpi

[RotateBackupCompress int            `json:"rotateBackupCompress"`]
qm=文件压缩级别
cz=RotateBackupCompress int `json:"rotateBackupCompress

[RotateCheckInterval  time.Duration  `json:"rotateCheckInterval"`]
qm=文件分割检查间隔
cz=RotateCheckInterval time.Duration `json:"rotateCheckInterva

[StdoutColorDisabled  bool           `json:"stdoutColorDisabled"`]
qm=关闭终端颜色输出
cz=StdoutColorDisabled bool `json:"stdoutColorDisabled

[WriterColorEnable    bool           `json:"writerColorEnable"`]
qm=文件是否输出颜色
cz=WriterColorEnable bool `json:"writerColorEnable"
