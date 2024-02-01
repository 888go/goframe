
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// LocalService provides a default implements for interface Service.
<原文结束>

# <翻译开始>
// LocalService 为接口 Service 提供了一个默认的实现。
# <翻译结束>


<原文开始>
// Service custom head string in service key.
<原文结束>

# <翻译开始>
// 在服务键中自定义头部字符串。
# <翻译结束>


<原文开始>
// Service deployment name, eg: dev, qa, staging, prod, etc.
<原文结束>

# <翻译开始>
// 服务部署名称，例如：开发（dev）、测试（qa）、预发布（staging）、生产（prod）等。
# <翻译结束>


<原文开始>
// Service Namespace, to indicate different services in the same environment with the same Name.
<原文结束>

# <翻译开始>
// 服务命名空间，用于在相同环境中标识具有相同名称的不同服务。
# <翻译结束>


<原文开始>
// Service version, eg: v1.0.0, v2.1.1, etc.
<原文结束>

# <翻译开始>
// 服务版本，例如：v1.0.0，v2.1.1等。
# <翻译结束>


<原文开始>
// Service Endpoints, pattern: IP:port, eg: 192.168.1.2:8000.
<原文结束>

# <翻译开始>
// 服务端点，格式：IP:端口，例如：192.168.1.2:8000
# <翻译结束>


<原文开始>
// Custom data for this service, which can be set using JSON by environment or command-line.
<原文结束>

# <翻译开始>
// 此处为本服务自定义数据，可以通过环境变量或命令行使用 JSON 设置。
# <翻译结束>


<原文开始>
// NewServiceWithName creates and returns a default implements for interface Service by service name.
<原文结束>

# <翻译开始>
// NewServiceWithName 通过服务名创建并返回一个接口 Service 的默认实现。
# <翻译结束>


<原文开始>
// NewServiceWithKV creates and returns a default implements for interface Service by key-value pair string.
<原文结束>

# <翻译开始>
// NewServiceWithKV 通过键值对字符串创建并返回接口 Service 的一个默认实现。
# <翻译结束>


<原文开始>
// GetName returns the name of the service.
// The name is necessary for a service, and should be unique among services.
<原文结束>

# <翻译开始>
// GetName 返回服务的名称。
// 服务必须具有名称，并且在各个服务之间应该是唯一的。
# <翻译结束>


<原文开始>
// GetVersion returns the version of the service.
// It is suggested using GNU version naming like: v1.0.0, v2.0.1, v2.1.0-rc.
// A service can have multiple versions deployed at once.
// If no version set in service, the default version of service is "latest".
<原文结束>

# <翻译开始>
// GetVersion 返回服务的版本。
// 建议采用类似 GNU 的版本命名方式，例如：v1.0.0、v2.0.1、v2.1.0-rc。
// 一项服务可以同时部署多个版本。
// 如果在服务中未设置版本，则服务的默认版本为 "latest"。
# <翻译结束>


<原文开始>
// GetKey formats and returns a unique key string for service.
// The result key is commonly used for key-value registrar server.
<原文结束>

# <翻译开始>
// GetKey 格式化并返回服务的唯一密钥字符串。
// 生成的密钥通常用于键值注册服务器。
# <翻译结束>


<原文开始>
// GetValue formats and returns the value of the service.
// The result value is commonly used for key-value registrar server.
<原文结束>

# <翻译开始>
// GetValue 格式化并返回服务的值。
// 返回的结果值通常用于键值注册服务器。
# <翻译结束>


<原文开始>
// GetPrefix formats and returns the key prefix string.
// The result prefix string is commonly used in key-value registrar server
// for service searching.
//
// Take etcd server for example, the prefix string is used like:
// `etcdctl get /services/prod/hello.svc --prefix`
<原文结束>

# <翻译开始>
// GetPrefix 格式化并返回键前缀字符串。
// 生成的前缀字符串通常用于键值注册服务器中进行服务搜索。
//
// 以 etcd 服务器为例，前缀字符串可以像这样使用：
// `etcdctl get /services/prod/hello.svc --prefix`
# <翻译结束>


<原文开始>
// GetMetadata returns the Metadata map of service.
// The Metadata is key-value pair map specifying extra attributes of a service.
<原文结束>

# <翻译开始>
// GetMetadata 返回服务的元数据映射。
// 元数据是一个键值对映射，用于指定服务的额外属性。
# <翻译结束>


<原文开始>
// GetEndpoints returns the Endpoints of service.
// The Endpoints contain multiple host/port information of service.
<原文结束>

# <翻译开始>
// GetEndpoints 返回服务的 Endpoints。
// 这些 Endpoints 包含了服务的多个主机/端口信息。
# <翻译结束>


<原文开始>
// Name for the service.
<原文结束>

# <翻译开始>
// 服务名称。
# <翻译结束>

