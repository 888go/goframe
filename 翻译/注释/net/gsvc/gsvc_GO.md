
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
// Package gsvc provides service registry and discovery definition.
<原文结束>

# <翻译开始>
// 包gsvc提供了服务注册和发现的定义。 md5:d3c854663f57d96a
# <翻译结束>


<原文开始>
// Registry interface for service.
<原文结束>

# <翻译开始>
// 服务注册接口。 md5:ee4e4676007285d8
# <翻译结束>


<原文开始>
// Registrar interface for service registrar.
<原文结束>

# <翻译开始>
// 服务注册器的Registrar接口。 md5:524b709abe0ef38a
# <翻译结束>


<原文开始>
	// Register registers `service` to Registry.
	// Note that it returns a new Service if it changes the input Service with custom one.
<原文结束>

# <翻译开始>
// Register将`service`注册到Registry中。
// 请注意，如果它使用自定义的服务修改了输入服务，它将返回一个新的Service。
// md5:f5dad972d06c6ae4
# <翻译结束>


<原文开始>
// Deregister off-lines and removes `service` from the Registry.
<原文结束>

# <翻译开始>
// 从注册表中注销离线服务，并移除 `service`。 md5:dff133d3dba9309d
# <翻译结束>


<原文开始>
// Discovery interface for service discovery.
<原文结束>

# <翻译开始>
// 服务发现的接口。 md5:c536eba406d0d0be
# <翻译结束>


<原文开始>
// Search searches and returns services with specified condition.
<原文结束>

# <翻译开始>
// Search 搜索并返回符合指定条件的服务。 md5:62e529e326dae7b7
# <翻译结束>


<原文开始>
	// Watch watches specified condition changes.
	// The `key` is the prefix of service key.
<原文结束>

# <翻译开始>
// Watch 监视指定条件的变化。
// `key` 是服务键的前缀。
// md5:5f440541abba5ceb
# <翻译结束>


<原文开始>
// Watcher interface for service.
<原文结束>

# <翻译开始>
// 服务的观察者接口。 md5:fbe479c3a1bcdea0
# <翻译结束>


<原文开始>
	// Proceed proceeds watch in blocking way.
	// It returns all complete services that watched by `key` if any change.
<原文结束>

# <翻译开始>
// Proceed以阻塞方式继续监控。如果有任何变化，它将返回所有由`key`监控的完成服务。
// md5:9d737841a3f4691f
# <翻译结束>


<原文开始>
// Close closes the watcher.
<原文结束>

# <翻译开始>
// Close 关闭监听器。 md5:c20cd2708e199b34
# <翻译结束>


<原文开始>
// Service interface for service definition.
<原文结束>

# <翻译开始>
// 服务接口，用于服务定义。 md5:e9dc084a14b9ab3c
# <翻译结束>


<原文开始>
	// GetName returns the name of the service.
	// The name is necessary for a service, and should be unique among services.
<原文结束>

# <翻译开始>
// GetName 返回服务的名称。
// 服务名称是必需的，并且在所有服务中应具有唯一性。
// md5:0eed35abab84da2e
# <翻译结束>


<原文开始>
	// GetVersion returns the version of the service.
	// It is suggested using GNU version naming like: v1.0.0, v2.0.1, v2.1.0-rc.
	// A service can have multiple versions deployed at once.
	// If no version set in service, the default version of service is "latest".
<原文结束>

# <翻译开始>
// GetVersion 返回服务的版本。建议使用类似于 v1.0.0、v2.0.1、v2.1.0-rc 的 GNU 版本命名法。一个服务可以同时部署多个版本。如果服务中未设置版本，那么默认版本为 "latest"。
// md5:17ba2c1584ea1bed
# <翻译结束>


<原文开始>
	// GetKey formats and returns a unique key string for service.
	// The result key is commonly used for key-value registrar server.
<原文结束>

# <翻译开始>
// GetKey 格式化并返回服务的唯一键字符串。这个结果键通常用于键值注册服务器。
// md5:b5f03cb2d5e4ec9b
# <翻译结束>


<原文开始>
	// GetValue formats and returns the value of the service.
	// The result value is commonly used for key-value registrar server.
<原文结束>

# <翻译开始>
// GetValue 格式化并返回服务的值。
// 返回的结果值通常用于键值注册服务器。
// md5:4272ce46aece6a65
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
// GetPrefix 格式化并返回键的前缀字符串。
// 该前缀字符串常用于键值注册服务器中进行服务搜索。
//
// 以 etcd 服务器为例，前缀字符串的使用方式如下：
// `etcdctl get /services/prod/hello.svc --prefix`
// md5:85bdda3bcce1a496
# <翻译结束>


<原文开始>
	// GetMetadata returns the Metadata map of service.
	// The Metadata is key-value pair map specifying extra attributes of a service.
<原文结束>

# <翻译开始>
// GetMetadata 返回服务的元数据映射（Metadata）。元数据是一个键值对映射，用于指定服务的额外属性。
// md5:eb70e1f52cd809f7
# <翻译结束>


<原文开始>
	// GetEndpoints returns the Endpoints of service.
	// The Endpoints contain multiple host/port information of service.
<原文结束>

# <翻译开始>
// GetEndpoints 返回服务的 Endpoints。Endpoints 包含服务的多个主机/端口信息。
// md5:7c1a630ab5b68b23
# <翻译结束>


<原文开始>
// Endpoint interface for service.
<原文结束>

# <翻译开始>
// 服务的端点接口。 md5:949bbba900f0fdec
# <翻译结束>


<原文开始>
// Host returns the IPv4/IPv6 address of a service.
<原文结束>

# <翻译开始>
// Host 返回服务的 IPv4/IPv6 地址。 md5:c70938f835a0f6e4
# <翻译结束>


<原文开始>
// Port returns the port of a service.
<原文结束>

# <翻译开始>
// Port 返回服务的端口。 md5:1650bc955f20ce4c
# <翻译结束>


<原文开始>
// String formats and returns the Endpoint as a string.
<原文结束>

# <翻译开始>
// String 将Endpoint格式化并作为字符串返回。 md5:03761a672c1719e3
# <翻译结束>


<原文开始>
// Endpoints are composed by multiple Endpoint.
<原文结束>

# <翻译开始>
// 终点（Endpoints）由多个Endpoint组成。 md5:253a91bd2341b2c5
# <翻译结束>


<原文开始>
// Metadata stores custom key-value pairs.
<原文结束>

# <翻译开始>
// Metadata 存储自定义的键值对。 md5:4a32feeda30a366c
# <翻译结束>


<原文开始>
// SearchInput is the input for service searching.
<原文结束>

# <翻译开始>
// SearchInput 是用于服务搜索的输入参数。 md5:af7e87c98cbdc120
# <翻译结束>


<原文开始>
// Search by service name.
<原文结束>

# <翻译开始>
// 通过服务名称进行搜索。 md5:2d5eea224c638808
# <翻译结束>


<原文开始>
// Search by service version.
<原文结束>

# <翻译开始>
// 通过服务版本搜索。 md5:b84dd2ba1226889a
# <翻译结束>


<原文开始>
// Filter by metadata if there are multiple result.
<原文结束>

# <翻译开始>
// 如果有多个结果，根据元数据进行过滤。 md5:96b1e7f8b8bddffa
# <翻译结束>


<原文开始>
// Schema is the schema of service.
<原文结束>

# <翻译开始>
// Schema 是服务的架构。 md5:6a1b7f8839c75ab9
# <翻译结束>


<原文开始>
// DefaultHead is the default head of service.
<原文结束>

# <翻译开始>
// DefaultHead 是服务的默认头部。 md5:d7d162114616b5b2
# <翻译结束>


<原文开始>
// DefaultDeployment is the default deployment of service.
<原文结束>

# <翻译开始>
// DefaultDeployment 是服务的默认部署。 md5:65707ae0cef5a070
# <翻译结束>


<原文开始>
// DefaultNamespace is the default namespace of service.
<原文结束>

# <翻译开始>
// DefaultNamespace 是服务的默认命名空间。 md5:4f63daab89c93e9b
# <翻译结束>


<原文开始>
// DefaultVersion is the default version of service.
<原文结束>

# <翻译开始>
// DefaultVersion 是服务的默认版本。 md5:d4b3ea889260c002
# <翻译结束>


<原文开始>
// EnvPrefix is the environment variable prefix.
<原文结束>

# <翻译开始>
// EnvPrefix 是环境变量的前缀。 md5:f9c76700eaf507b7
# <翻译结束>


<原文开始>
// EnvDeployment is the environment variable deployment.
<原文结束>

# <翻译开始>
// EnvDeployment 是环境变量部署。 md5:debd2d8465c7fe29
# <翻译结束>


<原文开始>
// EnvNamespace is the environment variable namespace.
<原文结束>

# <翻译开始>
// EnvNamespace 是环境变量的命名空间。 md5:fc532cfcc9e08d5d
# <翻译结束>


<原文开始>
// EnvName is the environment variable name.
<原文结束>

# <翻译开始>
// EnvName 是环境变量的名称。 md5:40377ca0a4d5ab27
# <翻译结束>


<原文开始>
// EnvVersion is the environment variable version.
<原文结束>

# <翻译开始>
// EnvVersion 是环境变量的版本。 md5:c2bc20cec15238b8
# <翻译结束>


<原文开始>
// MDProtocol is the metadata key for protocol.
<原文结束>

# <翻译开始>
// MDProtocol是协议的元数据键。 md5:1263117efae5845d
# <翻译结束>


<原文开始>
// MDInsecure is the metadata key for insecure.
<原文结束>

# <翻译开始>
// MDInsecure是关于不安全的元数据键。 md5:f869d79b8547b1a2
# <翻译结束>


<原文开始>
// MDWeight is the metadata key for weight.
<原文结束>

# <翻译开始>
// MDWeight 是用于权重的元数据键。 md5:48954b59a72a74ce
# <翻译结束>


<原文开始>
// DefaultProtocol is the default protocol of service.
<原文结束>

# <翻译开始>
// DefaultProtocol 是服务的默认协议。 md5:49c5538b8f79edc9
# <翻译结束>


<原文开始>
// DefaultSeparator is the default separator of service.
<原文结束>

# <翻译开始>
// DefaultSeparator是服务的默认分隔符。 md5:c24fd1db9709d190
# <翻译结束>


<原文开始>
// EndpointHostPortDelimiter is the delimiter of host and port.
<原文结束>

# <翻译开始>
// EndpointHostPortDelimiter是主机和端口之间的分隔符。 md5:305b988da4318717
# <翻译结束>


<原文开始>
// defaultTimeout is the default timeout for service registry.
<原文结束>

# <翻译开始>
// defaultTimeout 是服务注册的默认超时时间。 md5:18422c70857b9cf0
# <翻译结束>


<原文开始>
// EndpointsDelimiter is the delimiter of endpoints.
<原文结束>

# <翻译开始>
// EndpointsDelimiter 是端点的分隔符。 md5:53e996ee7d6bcd61
# <翻译结束>


<原文开始>
// SetRegistry sets the default Registry implements as your own implemented interface.
<原文结束>

# <翻译开始>
// SetRegistry 设置默认的Registry实现为你自定义实现的接口。 md5:4b1340106280e0dd
# <翻译结束>


<原文开始>
// GetRegistry returns the default Registry that is previously set.
// It returns nil if no Registry is set.
<原文结束>

# <翻译开始>
// GetRegistry 返回之前设置的默认Registry。如果没有设置Registry，它将返回nil。
// md5:efe24f47351d5419
# <翻译结束>

