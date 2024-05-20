
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// LocalService provides a default implements for interface Service.
<原文结束>

# <翻译开始>
// LocalService 提供了接口 Service 的默认实现。. md5:dd822c79690b7ca2
# <翻译结束>


<原文开始>
// Service custom head string in service key.
<原文结束>

# <翻译开始>
// 服务键中的自定义头部字符串。. md5:c33fc15669872592
# <翻译结束>


<原文开始>
// Service deployment name, eg: dev, qa, staging, prod, etc.
<原文结束>

# <翻译开始>
// 服务部署名称，例如：开发、测试、预发布、生产等。. md5:f4fb3f36ce08f902
# <翻译结束>


<原文开始>
// Service Namespace, to indicate different services in the same environment with the same Name.
<原文结束>

# <翻译开始>
// 服务命名空间，用于在相同环境中标识具有相同名称的不同服务。. md5:9ef35b3338765912
# <翻译结束>


<原文开始>
// Service version, eg: v1.0.0, v2.1.1, etc.
<原文结束>

# <翻译开始>
// 服务版本，例如：v1.0.0，v2.1.1 等等。. md5:ba8bedccf2112f4a
# <翻译结束>


<原文开始>
// Service Endpoints, pattern: IP:port, eg: 192.168.1.2:8000.
<原文结束>

# <翻译开始>
// 服务端点，格式：IP:port，例如：192.168.1.2:8000。. md5:8898f7c8bc80d33e
# <翻译结束>


<原文开始>
// Custom data for this service, which can be set using JSON by environment or command-line.
<原文结束>

# <翻译开始>
// 此服务的自定义数据，可以通过环境或命令行使用JSON进行设置。. md5:9eb9ae52ba435daf
# <翻译结束>


<原文开始>
// NewServiceWithName creates and returns a default implements for interface Service by service name.
<原文结束>

# <翻译开始>
// NewServiceWithName 根据服务名称创建并返回一个默认实现的 Service 接口实例。. md5:24c38960f0d05e58
# <翻译结束>


<原文开始>
// NewServiceWithKV creates and returns a default implements for interface Service by key-value pair string.
<原文结束>

# <翻译开始>
// NewServiceWithKV 使用键值对字符串创建并返回一个默认实现 Service 接口的服务。. md5:8789fc514001694f
# <翻译结束>


<原文开始>
// GetName returns the name of the service.
// The name is necessary for a service, and should be unique among services.
<原文结束>

# <翻译开始>
// GetName 返回服务的名称。
// 名称对于服务是必需的，应在所有服务中保持唯一。
// md5:c0cc1fa5e19d9a6c
# <翻译结束>


<原文开始>
// GetVersion returns the version of the service.
// It is suggested using GNU version naming like: v1.0.0, v2.0.1, v2.1.0-rc.
// A service can have multiple versions deployed at once.
// If no version set in service, the default version of service is "latest".
<原文结束>

# <翻译开始>
// GetVersion 返回服务的版本号。
// 建议使用GNU版本命名方式，例如：v1.0.0, v2.0.1, v2.1.0-rc。
// 服务可以同时部署多个版本。
// 如果服务中未设置版本，那么服务的默认版本为 "latest"。
// md5:bf857eeaf16711ca
# <翻译结束>


<原文开始>
// GetKey formats and returns a unique key string for service.
// The result key is commonly used for key-value registrar server.
<原文结束>

# <翻译开始>
// GetKey 格式化并返回服务的唯一键字符串。
// 生成的结果键通常用于键值注册服务器。
// md5:8651d9bc2f308934
# <翻译结束>


<原文开始>
// GetValue formats and returns the value of the service.
// The result value is commonly used for key-value registrar server.
<原文结束>

# <翻译开始>
// GetValue 格式化并返回服务的值。结果值通常用于键值注册服务器。
// md5:81a88bc4bcc73037
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
// 结果前缀字符串通常用于服务注册服务器中的服务搜索。
//
// 以 etcd 服务器为例，前缀字符串的用法如下：
// `etcdctl get /services/prod/hello.svc --prefix`
// md5:3c443e018050694a
# <翻译结束>


<原文开始>
// GetMetadata returns the Metadata map of service.
// The Metadata is key-value pair map specifying extra attributes of a service.
<原文结束>

# <翻译开始>
// GetMetadata 返回服务的元数据地图。
// 元数据是一个键值对映射，用于指定服务的额外属性。
// md5:42fd4200585681c1
# <翻译结束>


<原文开始>
// GetEndpoints returns the Endpoints of service.
// The Endpoints contain multiple host/port information of service.
<原文结束>

# <翻译开始>
// GetEndpoints 返回服务的端点信息。
// 端点包含服务的多个主机/端口信息。
// md5:164bdc2d3a7db5e0
# <翻译结束>

