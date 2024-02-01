
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
// Package gsvc provides service registry and discovery definition.
<原文结束>

# <翻译开始>
// Package gsvc 提供了服务注册与发现的定义。
# <翻译结束>


<原文开始>
// Registry interface for service.
<原文结束>

# <翻译开始>
// Registry interface 用于服务。
# <翻译结束>


<原文开始>
// Registrar interface for service registrar.
<原文结束>

# <翻译开始>
// Registrar 是服务注册器的接口。
# <翻译结束>


<原文开始>
	// Register registers `service` to Registry.
	// Note that it returns a new Service if it changes the input Service with custom one.
<原文结束>

# <翻译开始>
// Register 将 `service` 注册到 Registry。
// 注意，如果使用自定义服务更改了输入的服务，它将返回一个新的 Service。
# <翻译结束>


<原文开始>
// Deregister off-lines and removes `service` from the Registry.
<原文结束>

# <翻译开始>
// 从注册中心注销并移除`service`的离线信息。
# <翻译结束>


<原文开始>
// Discovery interface for service discovery.
<原文结束>

# <翻译开始>
// Discovery 接口用于服务发现。
# <翻译结束>


<原文开始>
// Search searches and returns services with specified condition.
<原文结束>

# <翻译开始>
// Search 按照指定条件搜索并返回服务。
# <翻译结束>


<原文开始>
	// Watch watches specified condition changes.
	// The `key` is the prefix of service key.
<原文结束>

# <翻译开始>
// Watch 监视指定条件的变化。
// `key` 是服务键的前缀。
# <翻译结束>


<原文开始>
// Watcher interface for service.
<原文结束>

# <翻译开始>
// Watcher接口用于服务。
# <翻译结束>


<原文开始>
	// Proceed proceeds watch in blocking way.
	// It returns all complete services that watched by `key` if any change.
<原文结束>

# <翻译开始>
// Proceed 以阻塞方式继续观察。
// 如果有任何变化，它将返回由`key`观察到的所有已完成服务。
# <翻译结束>


<原文开始>
// Service interface for service definition.
<原文结束>

# <翻译开始>
// Service 接口，用于定义服务。
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
// GetVersion 返回服务的版本信息。
// 建议使用GNU版本命名规则，例如：v1.0.0、v2.0.1、v2.1.0-rc。
// 一个服务可以同时部署多个版本。
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
// 生成的前缀字符串通常用于键值注册服务器中的服务搜索。
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
// Endpoint interface for service.
<原文结束>

# <翻译开始>
// Endpoint 是服务的接口。
# <翻译结束>


<原文开始>
// Host returns the IPv4/IPv6 address of a service.
<原文结束>

# <翻译开始>
// Host 返回一个服务的 IPv4/IPv6 地址。
# <翻译结束>


<原文开始>
// Port returns the port of a service.
<原文结束>

# <翻译开始>
// Port 返回一个服务的端口号。
# <翻译结束>


<原文开始>
// String formats and returns the Endpoint as a string.
<原文结束>

# <翻译开始>
// String 格式化并返回 Endpoint 作为字符串。
# <翻译结束>


<原文开始>
// Endpoints are composed by multiple Endpoint.
<原文结束>

# <翻译开始>
// Endpoints是由多个Endpoint组成的。
# <翻译结束>


<原文开始>
// Metadata stores custom key-value pairs.
<原文结束>

# <翻译开始>
// Metadata 用于存储自定义键值对。
# <翻译结束>


<原文开始>
// SearchInput is the input for service searching.
<原文结束>

# <翻译开始>
// SearchInput 是用于服务搜索的输入参数。
# <翻译结束>


<原文开始>
// Filter by metadata if there are multiple result.
<原文结束>

# <翻译开始>
// 如果有多个结果，通过元数据进行过滤。
# <翻译结束>


<原文开始>
// Schema is the schema of service.
<原文结束>

# <翻译开始>
// Schema 是服务的架构。
# <翻译结束>


<原文开始>
// DefaultHead is the default head of service.
<原文结束>

# <翻译开始>
// DefaultHead 是服务的默认头部。
# <翻译结束>


<原文开始>
// DefaultDeployment is the default deployment of service.
<原文结束>

# <翻译开始>
// DefaultDeployment 是服务的默认部署。
# <翻译结束>


<原文开始>
// DefaultNamespace is the default namespace of service.
<原文结束>

# <翻译开始>
// DefaultNamespace 是服务的默认命名空间。
# <翻译结束>


<原文开始>
// DefaultVersion is the default version of service.
<原文结束>

# <翻译开始>
// DefaultVersion 是服务的默认版本。
# <翻译结束>


<原文开始>
// EnvPrefix is the environment variable prefix.
<原文结束>

# <翻译开始>
// EnvPrefix 是环境变量前缀。
# <翻译结束>


<原文开始>
// EnvDeployment is the environment variable deployment.
<原文结束>

# <翻译开始>
// EnvDeployment 是环境变量部署。
# <翻译结束>


<原文开始>
// EnvNamespace is the environment variable namespace.
<原文结束>

# <翻译开始>
// EnvNamespace 是环境变量的命名空间。
# <翻译结束>


<原文开始>
// EnvName is the environment variable name.
<原文结束>

# <翻译开始>
// EnvName 是环境变量名称。
# <翻译结束>


<原文开始>
// EnvVersion is the environment variable version.
<原文结束>

# <翻译开始>
// EnvVersion 是环境变量版本。
# <翻译结束>


<原文开始>
// MDProtocol is the metadata key for protocol.
<原文结束>

# <翻译开始>
// MDProtocol 是用于协议的元数据键。
# <翻译结束>


<原文开始>
// MDInsecure is the metadata key for insecure.
<原文结束>

# <翻译开始>
// MDInsecure 是用于表示不安全的元数据键。
# <翻译结束>


<原文开始>
// MDWeight is the metadata key for weight.
<原文结束>

# <翻译开始>
// MDWeight 是用于权重的元数据键。
# <翻译结束>


<原文开始>
// DefaultProtocol is the default protocol of service.
<原文结束>

# <翻译开始>
// DefaultProtocol 是服务的默认协议。
# <翻译结束>


<原文开始>
// DefaultSeparator is the default separator of service.
<原文结束>

# <翻译开始>
// DefaultSeparator 是服务的默认分隔符。
# <翻译结束>


<原文开始>
// EndpointHostPortDelimiter is the delimiter of host and port.
<原文结束>

# <翻译开始>
// EndpointHostPortDelimiter 是主机和端口之间的分隔符。
# <翻译结束>


<原文开始>
// defaultTimeout is the default timeout for service registry.
<原文结束>

# <翻译开始>
// defaultTimeout 是服务注册表的默认超时时间。
# <翻译结束>


<原文开始>
// EndpointsDelimiter is the delimiter of endpoints.
<原文结束>

# <翻译开始>
// EndpointsDelimiter 是 endpoints 的分隔符。
# <翻译结束>


<原文开始>
// SetRegistry sets the default Registry implements as your own implemented interface.
<原文结束>

# <翻译开始>
// SetRegistry 设置默认的注册中心为自定义实现的接口。
# <翻译结束>


<原文开始>
// GetRegistry returns the default Registry that is previously set.
// It returns nil if no Registry is set.
<原文结束>

# <翻译开始>
// GetRegistry 返回先前设置的默认 Registry。
// 如果未设置 Registry，则返回 nil。
# <翻译结束>


<原文开始>
// Close closes the watcher.
<原文结束>

# <翻译开始>
// Close 关闭监视器。
# <翻译结束>


<原文开始>
// Search by key prefix.
<原文结束>

# <翻译开始>
// 通过键前缀搜索。
# <翻译结束>


<原文开始>
// Search by service name.
<原文结束>

# <翻译开始>
// 通过服务名搜索
# <翻译结束>


<原文开始>
// Search by service version.
<原文结束>

# <翻译开始>
// 通过服务版本进行搜索。
# <翻译结束>

