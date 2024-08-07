// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gsvc提供了服务注册和发现的定义。 md5:d3c854663f57d96a
package gsvc

import (
	"context"
	"time"

	gerror "github.com/888go/goframe/errors/gerror"
)

// 服务注册接口。 md5:ee4e4676007285d8
type Registry interface {
	Registrar
	Discovery
}

// 服务注册器的Registrar接口。 md5:524b709abe0ef38a
type Registrar interface {
	// Register将`service`注册到Registry中。
	// 请注意，如果它使用自定义的服务修改了输入服务，它将返回一个新的Service。
	// md5:f5dad972d06c6ae4
	Register(ctx context.Context, service Service) (registered Service, err error)

		// 从注册表中注销离线服务，并移除 `service`。 md5:dff133d3dba9309d
	Deregister(ctx context.Context, service Service) error
}

// 服务发现的接口。 md5:c536eba406d0d0be
type Discovery interface {
		// Search 搜索并返回符合指定条件的服务。 md5:62e529e326dae7b7
	Search(ctx context.Context, in SearchInput) (result []Service, err error)

	// Watch 监视指定条件的变化。
	// `key` 是服务键的前缀。
	// md5:5f440541abba5ceb
	Watch(ctx context.Context, key string) (watcher Watcher, err error)
}

// 服务的观察者接口。 md5:fbe479c3a1bcdea0
type Watcher interface {
	// Proceed以阻塞方式继续监控。如果有任何变化，它将返回所有由`key`监控的完成服务。
	// md5:9d737841a3f4691f
	Proceed() (services []Service, err error)

		// Close 关闭监听器。 md5:c20cd2708e199b34
	Close() error
}

// 服务接口，用于服务定义。 md5:e9dc084a14b9ab3c
type Service interface {
	// GetName 返回服务的名称。
	// 服务名称是必需的，并且在所有服务中应具有唯一性。
	// md5:0eed35abab84da2e
	GetName() string

	// GetVersion 返回服务的版本。建议使用类似于 v1.0.0、v2.0.1、v2.1.0-rc 的 GNU 版本命名法。一个服务可以同时部署多个版本。如果服务中未设置版本，那么默认版本为 "latest"。
	// md5:17ba2c1584ea1bed
	GetVersion() string

	// GetKey 格式化并返回服务的唯一键字符串。这个结果键通常用于键值注册服务器。
	// md5:b5f03cb2d5e4ec9b
	GetKey() string

	// GetValue 格式化并返回服务的值。
	// 返回的结果值通常用于键值注册服务器。
	// md5:4272ce46aece6a65
	GetValue() string

	// GetPrefix 格式化并返回键的前缀字符串。
	// 该前缀字符串常用于键值注册服务器中进行服务搜索。
	//
	// 以 etcd 服务器为例，前缀字符串的使用方式如下：
	// `etcdctl get /services/prod/hello.svc --prefix`
	// md5:85bdda3bcce1a496
	GetPrefix() string

	// GetMetadata 返回服务的元数据映射（Metadata）。元数据是一个键值对映射，用于指定服务的额外属性。
	// md5:eb70e1f52cd809f7
	GetMetadata() Metadata

	// GetEndpoints 返回服务的 Endpoints。Endpoints 包含服务的多个主机/端口信息。
	// md5:7c1a630ab5b68b23
	GetEndpoints() Endpoints
}

// 服务的端点接口。 md5:949bbba900f0fdec
type Endpoint interface {
		// Host 返回服务的 IPv4/IPv6 地址。 md5:c70938f835a0f6e4
	Host() string

		// Port 返回服务的端口。 md5:1650bc955f20ce4c
	Port() int

		// String 将Endpoint格式化并作为字符串返回。 md5:03761a672c1719e3
	String() string
}

// 终点（Endpoints）由多个Endpoint组成。 md5:253a91bd2341b2c5
type Endpoints []Endpoint

// Metadata 存储自定义的键值对。 md5:4a32feeda30a366c
type Metadata map[string]interface{}

// SearchInput 是用于服务搜索的输入参数。 md5:af7e87c98cbdc120
type SearchInput struct {
	Prefix   string   // Search by key prefix.
	Name     string   // 通过服务名称进行搜索。 md5:2d5eea224c638808
	Version  string   // 通过服务版本搜索。 md5:b84dd2ba1226889a
	Metadata Metadata // 如果有多个结果，根据元数据进行过滤。 md5:96b1e7f8b8bddffa
}

const (
	Schema                    = `service`            // Schema 是服务的架构。 md5:6a1b7f8839c75ab9
	DefaultHead               = `service`            // DefaultHead 是服务的默认头部。 md5:d7d162114616b5b2
	DefaultDeployment         = `default`            // DefaultDeployment 是服务的默认部署。 md5:65707ae0cef5a070
	DefaultNamespace          = `default`            // DefaultNamespace 是服务的默认命名空间。 md5:4f63daab89c93e9b
	DefaultVersion            = `latest`             // DefaultVersion 是服务的默认版本。 md5:d4b3ea889260c002
	EnvPrefix                 = `GF_GSVC_PREFIX`     // EnvPrefix 是环境变量的前缀。 md5:f9c76700eaf507b7
	EnvDeployment             = `GF_GSVC_DEPLOYMENT` // EnvDeployment 是环境变量部署。 md5:debd2d8465c7fe29
	EnvNamespace              = `GF_GSVC_NAMESPACE`  // EnvNamespace 是环境变量的命名空间。 md5:fc532cfcc9e08d5d
	EnvName                   = `GF_GSVC_Name`       // EnvName 是环境变量的名称。 md5:40377ca0a4d5ab27
	EnvVersion                = `GF_GSVC_VERSION`    // EnvVersion 是环境变量的版本。 md5:c2bc20cec15238b8
	MDProtocol                = `protocol`           // MDProtocol是协议的元数据键。 md5:1263117efae5845d
	MDInsecure                = `insecure`           // MDInsecure是关于不安全的元数据键。 md5:f869d79b8547b1a2
	MDWeight                  = `weight`             // MDWeight 是用于权重的元数据键。 md5:48954b59a72a74ce
	DefaultProtocol           = `http`               // DefaultProtocol 是服务的默认协议。 md5:49c5538b8f79edc9
	DefaultSeparator          = "/"                  // DefaultSeparator是服务的默认分隔符。 md5:c24fd1db9709d190
	EndpointHostPortDelimiter = ":"                  // EndpointHostPortDelimiter是主机和端口之间的分隔符。 md5:305b988da4318717
	defaultTimeout            = 5 * time.Second      // defaultTimeout 是服务注册的默认超时时间。 md5:18422c70857b9cf0
	EndpointsDelimiter        = ","                  // EndpointsDelimiter 是端点的分隔符。 md5:53e996ee7d6bcd61
)

var defaultRegistry Registry

// SetRegistry 设置默认的Registry实现为你自定义实现的接口。 md5:4b1340106280e0dd
func SetRegistry(registry Registry) {
	if registry == nil {
		panic(gerror.X创建(`invalid Registry value "nil" given`))
	}
	defaultRegistry = registry
}

// GetRegistry 返回之前设置的默认Registry。如果没有设置Registry，它将返回nil。
// md5:efe24f47351d5419
func GetRegistry() Registry {
	return defaultRegistry
}
