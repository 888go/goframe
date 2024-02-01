// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gsvc 提供了服务注册与发现的定义。
package gsvc
import (
	"context"
	"time"
	
	"github.com/888go/goframe/errors/gerror"
	)
// Registry interface 用于服务。
type Registry interface {
	Registrar
	Discovery
}

// Registrar 是服务注册器的接口。
type Registrar interface {
// Register 将 `service` 注册到 Registry。
// 注意，如果使用自定义服务更改了输入的服务，它将返回一个新的 Service。
	Register(ctx context.Context, service Service) (registered Service, err error)

	// 从注册中心注销并移除`service`的离线信息。
	Deregister(ctx context.Context, service Service) error
}

// Discovery 接口用于服务发现。
type Discovery interface {
	// Search 按照指定条件搜索并返回服务。
	Search(ctx context.Context, in SearchInput) (result []Service, err error)

// Watch 监视指定条件的变化。
// `key` 是服务键的前缀。
	Watch(ctx context.Context, key string) (watcher Watcher, err error)
}

// Watcher接口用于服务。
type Watcher interface {
// Proceed 以阻塞方式继续观察。
// 如果有任何变化，它将返回由`key`观察到的所有已完成服务。
	Proceed() (services []Service, err error)

	// Close 关闭监视器。
	Close() error
}

// Service 接口，用于定义服务。
type Service interface {
// GetName 返回服务的名称。
// 服务必须具有名称，并且在各个服务之间应该是唯一的。
	GetName() string

// GetVersion 返回服务的版本信息。
// 建议使用GNU版本命名规则，例如：v1.0.0、v2.0.1、v2.1.0-rc。
// 一个服务可以同时部署多个版本。
// 如果在服务中未设置版本，则服务的默认版本为 "latest"。
	GetVersion() string

// GetKey 格式化并返回服务的唯一密钥字符串。
// 生成的密钥通常用于键值注册服务器。
	GetKey() string

// GetValue 格式化并返回服务的值。
// 返回的结果值通常用于键值注册服务器。
	GetValue() string

// GetPrefix 格式化并返回键前缀字符串。
// 生成的前缀字符串通常用于键值注册服务器中的服务搜索。
//
// 以 etcd 服务器为例，前缀字符串可以像这样使用：
// `etcdctl get /services/prod/hello.svc --prefix`
	GetPrefix() string

// GetMetadata 返回服务的元数据映射。
// 元数据是一个键值对映射，用于指定服务的额外属性。
	GetMetadata() Metadata

// GetEndpoints 返回服务的 Endpoints。
// 这些 Endpoints 包含了服务的多个主机/端口信息。
	GetEndpoints() Endpoints
}

// Endpoint 是服务的接口。
type Endpoint interface {
	// Host 返回一个服务的 IPv4/IPv6 地址。
	Host() string

	// Port 返回一个服务的端口号。
	Port() int

	// String 格式化并返回 Endpoint 作为字符串。
	String() string
}

// Endpoints是由多个Endpoint组成的。
type Endpoints []Endpoint

// Metadata 用于存储自定义键值对。
type Metadata map[string]interface{}

// SearchInput 是用于服务搜索的输入参数。
type SearchInput struct {
	Prefix   string   // 通过键前缀搜索。
	Name     string   // 通过服务名搜索
	Version  string   // 通过服务版本进行搜索。
	Metadata Metadata // 如果有多个结果，通过元数据进行过滤。
}

const (
	Schema                    = `service`            // Schema 是服务的架构。
	DefaultHead               = `service`            // DefaultHead 是服务的默认头部。
	DefaultDeployment         = `default`            // DefaultDeployment 是服务的默认部署。
	DefaultNamespace          = `default`            // DefaultNamespace 是服务的默认命名空间。
	DefaultVersion            = `latest`             // DefaultVersion 是服务的默认版本。
	EnvPrefix                 = `GF_GSVC_PREFIX`     // EnvPrefix 是环境变量前缀。
	EnvDeployment             = `GF_GSVC_DEPLOYMENT` // EnvDeployment 是环境变量部署。
	EnvNamespace              = `GF_GSVC_NAMESPACE`  // EnvNamespace 是环境变量的命名空间。
	EnvName                   = `GF_GSVC_Name`       // EnvName 是环境变量名称。
	EnvVersion                = `GF_GSVC_VERSION`    // EnvVersion 是环境变量版本。
	MDProtocol                = `protocol`           // MDProtocol 是用于协议的元数据键。
	MDInsecure                = `insecure`           // MDInsecure 是用于表示不安全的元数据键。
	MDWeight                  = `weight`             // MDWeight 是用于权重的元数据键。
	DefaultProtocol           = `http`               // DefaultProtocol 是服务的默认协议。
	DefaultSeparator          = "/"                  // DefaultSeparator 是服务的默认分隔符。
	EndpointHostPortDelimiter = ":"                  // EndpointHostPortDelimiter 是主机和端口之间的分隔符。
	defaultTimeout            = 5 * time.Second      // defaultTimeout 是服务注册表的默认超时时间。
	EndpointsDelimiter        = ","                  // EndpointsDelimiter 是 endpoints 的分隔符。
)

var defaultRegistry Registry

// SetRegistry 设置默认的注册中心为自定义实现的接口。
func SetRegistry(registry Registry) {
	if registry == nil {
		panic(gerror.New(`invalid Registry value "nil" given`))
	}
	defaultRegistry = registry
}

// GetRegistry 返回先前设置的默认 Registry。
// 如果未设置 Registry，则返回 nil。
func GetRegistry() Registry {
	return defaultRegistry
}
