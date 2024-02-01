
# <翻译开始>
Handlers []Handler
X中间件
<跳到行首>
# <翻译结束>

# <翻译开始>
Flags int `json:"flags"`
X日志标识
<跳到行首>
# <翻译结束>

# <翻译开始>
TimeFormat string `json:"timeFormat"`
X时间格式
<跳到行首>
# <翻译结束>

# <翻译开始>
Path string `json:"path"`
X文件路径
<跳到行首>
# <翻译结束>

# <翻译开始>
File string `json:"file"`
X文件名格式
<跳到行首>
# <翻译结束>

# <翻译开始>
Level int `json:"level"
X级别
<跳到行首>
# <翻译结束>

# <翻译开始>
Prefix string `json:"prefix"
X前缀
<跳到行首>
# <翻译结束>

# <翻译开始>
StStatus int `json:"stStatus"`
X堆栈状态
<跳到行首>
# <翻译结束>

# <翻译开始>
StFilter string `json:"stFilter"`
X堆栈过滤
<跳到行首>
# <翻译结束>

# <翻译开始>
CtxKeys []interface{} `json:"ctxKeys"`
X上下文名称
<跳到行首>
# <翻译结束>

# <翻译开始>
HeaderPrint bool `json:"header"`
X是否输出头信息
<跳到行首>
# <翻译结束>

# <翻译开始>
StdoutPrint bool `json:"stdout"`
X是否同时输出到终端
<跳到行首>
# <翻译结束>

# <翻译开始>
LevelPrint bool `json:"levelPrint"
X是否输出级别
<跳到行首>
# <翻译结束>

# <翻译开始>
LevelPrefixes map[int]string `json:"levelPrefixes"`
X日志级别名称映射
<跳到行首>
# <翻译结束>

# <翻译开始>
RotateSize int64 `json:"rotateSize"
X文件分割大小
<跳到行首>
# <翻译结束>

# <翻译开始>
RotateExpire time.Duration `json:"rotateExpire
X文件分割周期
<跳到行首>
# <翻译结束>

# <翻译开始>
RotateBackupLimit int `json:"rotateBackupLimit"
X文件分割保留数量
<跳到行首>
# <翻译结束>

# <翻译开始>
RotateBackupExpire time.Duration `json:"rotateBackupExpi
X文件分割过期时间
<跳到行首>
# <翻译结束>

# <翻译开始>
RotateBackupCompress int `json:"rotateBackupCompress
X文件压缩级别
<跳到行首>
# <翻译结束>

# <翻译开始>
RotateCheckInterval time.Duration `json:"rotateCheckInterva
X文件分割检查间隔
<跳到行首>
# <翻译结束>

# <翻译开始>
StdoutColorDisabled bool `json:"stdoutColorDisabled
X关闭终端颜色输出
<跳到行首>
# <翻译结束>

# <翻译开始>
WriterColorEnable bool `json:"writerColorEnable"
X文件是否输出颜色
<跳到行首>
# <翻译结束>

# <翻译开始>
StSkip int `json:"stSkip"`
X堆栈偏移量
<跳到行首>
# <翻译结束>
