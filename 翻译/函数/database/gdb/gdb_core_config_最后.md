
# <翻译开始>
type Config
X配置
# <翻译结束>

# <翻译开始>
type ConfigGroup
X配置组
# <翻译结束>

# <翻译开始>
type ConfigNode
X配置项
# <翻译结束>

# <翻译开始>
Host                 string        `json:"host"`
X地址
<跳到行首>
# <翻译结束>

# <翻译开始>
Port                 string        `json:"port"`
X端口
<跳到行首>
# <翻译结束>

# <翻译开始>
User                 string        `json:"user"
X账号
<跳到行首>
# <翻译结束>

# <翻译开始>
Pass                 string        `json:"pass"
X密码
<跳到行首>
# <翻译结束>

# <翻译开始>
Name                 string        `json:"name"
X默认数据库名称
<跳到行首>
# <翻译结束>

# <翻译开始>
Type                 string        `json:"type"`
X数据库类型
<跳到行首>
# <翻译结束>

# <翻译开始>
Link                 string        `json:"link"`
X自定义链接信息
<跳到行首>
# <翻译结束>

# <翻译开始>
Extra                string        `json:"extra"
X额外配置
<跳到行首>
# <翻译结束>

# <翻译开始>
Role                 string        `json:"role"`
X节点角色
<跳到行首>
# <翻译结束>

# <翻译开始>
Debug                bool          `json:"debug"
X调试模式
<跳到行首>
# <翻译结束>

# <翻译开始>
Prefix               string        `json:"prefix"`
X表前缀
<跳到行首>
# <翻译结束>

# <翻译开始>
DryRun               bool          `json:"dryRun"`
X空跑特性
<跳到行首>
# <翻译结束>

# <翻译开始>
Weight               int           `json:"weight"
X负载均衡权重
<跳到行首>
# <翻译结束>

# <翻译开始>
Charset              string        `json:"charset"
X字符集
<跳到行首>
# <翻译结束>

# <翻译开始>
Protocol             string        `json:"protocol"`
X协议
<跳到行首>
# <翻译结束>

# <翻译开始>
Timezone             string        `json:"timezone"
X时区
<跳到行首>
# <翻译结束>

# <翻译开始>
Namespace            string        `json:"namespace"
X命名空间
<跳到行首>
# <翻译结束>

# <翻译开始>
MaxIdleConnCount     int           `json:"maxIdle"`
X最大闲置连接数
<跳到行首>
# <翻译结束>

# <翻译开始>
MaxOpenConnCount     int           `json:"maxOpen"`
X最大打开连接数
<跳到行首>
# <翻译结束>

# <翻译开始>
MaxConnLifeTime      time.Duration `json:"maxLifeTime"`
X最大空闲时长
<跳到行首>
# <翻译结束>

# <翻译开始>
QueryTimeout         time.Duration `json:"queryTimeout"
X查询超时时长
<跳到行首>
# <翻译结束>

# <翻译开始>
ExecTimeout          time.Duration `json:"execTimeout"
X执行超时时长
<跳到行首>
# <翻译结束>

# <翻译开始>
TranTimeout          time.Duration `json:"tranTimeout"
X事务超时时长
<跳到行首>
# <翻译结束>

# <翻译开始>
PrepareTimeout       time.Duration `json:"prepareTimeout"
X预准备SQL超时时长
<跳到行首>
# <翻译结束>

# <翻译开始>
CreatedAt            string        `json:"createdAt"
X创建时间字段名
<跳到行首>
# <翻译结束>

# <翻译开始>
UpdatedAt            string        `json:"updatedAt
X更新时间字段名
<跳到行首>
# <翻译结束>

# <翻译开始>
DeletedAt            string        `json:"deletedAt"`
X软删除时间字段名
<跳到行首>
# <翻译结束>

# <翻译开始>
TimeMaintainDisabled bool          `json:"timeMaintainDisabled"
X禁用时间自动更新特性
<跳到行首>
# <翻译结束>
