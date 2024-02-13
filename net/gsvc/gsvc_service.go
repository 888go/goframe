// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gsvc

import (
	"context"
	
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/os/gcmd"
	"github.com/888go/goframe/text/gstr"
)

// LocalService 为接口 Service 提供了一个默认的实现。
type LocalService struct {
	Head       string    // 在服务键中自定义头部字符串。
	Deployment string    // 服务部署名称，例如：开发（dev）、测试（qa）、预发布（staging）、生产（prod）等。
	Namespace  string    // 服务命名空间，用于在相同环境中标识具有相同名称的不同服务。
	Name       string    // 服务名称。
	Version    string    // 服务版本，例如：v1.0.0，v2.1.1等。
	Endpoints  Endpoints // 服务端点，格式：IP:端口，例如：192.168.1.2:8000
	Metadata   Metadata  // 此处为本服务自定义数据，可以通过环境变量或命令行使用 JSON 设置。
}

// NewServiceWithName 通过服务名创建并返回一个接口 Service 的默认实现。
func NewServiceWithName(name string) Service {
	s := &LocalService{
		Name:     name,
		Metadata: make(Metadata),
	}
	s.autoFillDefaultAttributes()
	return s
}

// NewServiceWithKV 通过键值对字符串创建并返回接口 Service 的一个默认实现。
func NewServiceWithKV(key, value string) (Service, error) {
	var (
		err   error
		array = 文本类.X分割(文本类.X过滤首尾符并含空白(key, DefaultSeparator), DefaultSeparator)
	)
	if len(array) < 6 {
		err = 错误类.X创建错误码并格式化(错误码类.CodeInvalidParameter, `invalid service key "%s"`, key)
		return nil, err
	}
	s := &LocalService{
		Head:       array[0],
		Deployment: array[1],
		Namespace:  array[2],
		Name:       array[3],
		Version:    array[4],
		Endpoints:  NewEndpoints(array[5]),
		Metadata:   make(Metadata),
	}
	s.autoFillDefaultAttributes()
	if len(value) > 0 {
		if err = json类.Unmarshal别名([]byte(value), &s.Metadata); err != nil {
			err = 错误类.X多层错误码并格式化(错误码类.CodeInvalidParameter, err, `invalid service value "%s"`, value)
			return nil, err
		}
	}
	return s, nil
}

// GetName 返回服务的名称。
// 服务必须具有名称，并且在各个服务之间应该是唯一的。
func (s *LocalService) GetName() string {
	return s.Name
}

// GetVersion 返回服务的版本。
// 建议采用类似 GNU 的版本命名方式，例如：v1.0.0、v2.0.1、v2.1.0-rc。
// 一项服务可以同时部署多个版本。
// 如果在服务中未设置版本，则服务的默认版本为 "latest"。
func (s *LocalService) GetVersion() string {
	return s.Version
}

// GetKey 格式化并返回服务的唯一密钥字符串。
// 生成的密钥通常用于键值注册服务器。
func (s *LocalService) GetKey() string {
	serviceNameUnique := s.GetPrefix()
	serviceNameUnique += DefaultSeparator + s.Endpoints.String()
	return serviceNameUnique
}

// GetValue 格式化并返回服务的值。
// 返回的结果值通常用于键值注册服务器。
func (s *LocalService) GetValue() string {
	b, err := json类.Marshal别名(s.Metadata)
	if err != nil {
		intlog.Errorf(context.TODO(), `%+v`, err)
	}
	return string(b)
}

// GetPrefix 格式化并返回键前缀字符串。
// 生成的前缀字符串通常用于键值注册服务器中进行服务搜索。
//
// 以 etcd 服务器为例，前缀字符串可以像这样使用：
// `etcdctl get /services/prod/hello.svc --prefix`
func (s *LocalService) GetPrefix() string {
	s.autoFillDefaultAttributes()
	return DefaultSeparator + 文本类.X连接(
		[]string{
			s.Head,
			s.Deployment,
			s.Namespace,
			s.Name,
			s.Version,
		},
		DefaultSeparator,
	)
}

// GetMetadata 返回服务的元数据映射。
// 元数据是一个键值对映射，用于指定服务的额外属性。
func (s *LocalService) GetMetadata() Metadata {
	return s.Metadata
}

// GetEndpoints 返回服务的 Endpoints。
// 这些 Endpoints 包含了服务的多个主机/端口信息。
func (s *LocalService) GetEndpoints() Endpoints {
	return s.Endpoints
}

func (s *LocalService) autoFillDefaultAttributes() {
	if s.Head == "" {
		s.Head = cmd类.GetOptWithEnv(EnvPrefix, DefaultHead).String()
	}
	if s.Deployment == "" {
		s.Deployment = cmd类.GetOptWithEnv(EnvDeployment, DefaultDeployment).String()
	}
	if s.Namespace == "" {
		s.Namespace = cmd类.GetOptWithEnv(EnvNamespace, DefaultNamespace).String()
	}
	if s.Name == "" {
		s.Name = cmd类.GetOptWithEnv(EnvName).String()
	}
	if s.Version == "" {
		s.Version = cmd类.GetOptWithEnv(EnvVersion, DefaultVersion).String()
	}
}
