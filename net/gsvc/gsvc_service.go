// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gsvc

import (
	"context"

	gjson "github.com/888go/goframe/encoding/gjson"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	gcmd "github.com/888go/goframe/os/gcmd"
	gstr "github.com/888go/goframe/text/gstr"
)

// LocalService 提供了接口 Service 的默认实现。 md5:dd822c79690b7ca2
type LocalService struct {
	Head       string    // 服务键中的自定义头部字符串。 md5:c33fc15669872592
	Deployment string    // 服务部署名称，例如：开发、测试、预发布、生产等。 md5:f4fb3f36ce08f902
	Namespace  string    // 服务命名空间，用于在相同环境中标识具有相同名称的不同服务。 md5:9ef35b3338765912
	Name       string    // Name for the service.
	Version    string    // 服务版本，例如：v1.0.0，v2.1.1 等等。 md5:ba8bedccf2112f4a
	Endpoints  Endpoints // 服务端点，格式：IP:port，例如：192.168.1.2:8000。 md5:8898f7c8bc80d33e
	Metadata   Metadata  // 此服务的自定义数据，可以通过环境或命令行使用JSON进行设置。 md5:9eb9ae52ba435daf
}

// NewServiceWithName 根据服务名称创建并返回一个默认实现的 Service 接口实例。 md5:24c38960f0d05e58
func NewServiceWithName(name string) Service {
	s := &LocalService{
		Name:     name,
		Metadata: make(Metadata),
	}
	s.autoFillDefaultAttributes()
	return s
}

// NewServiceWithKV 使用键值对字符串创建并返回一个默认实现 Service 接口的服务。 md5:8789fc514001694f
func NewServiceWithKV(key, value string) (Service, error) {
	var (
		err   error
		array = gstr.X分割(gstr.X过滤首尾符并含空白(key, DefaultSeparator), DefaultSeparator)
	)
	if len(array) < 6 {
		err = gerror.X创建错误码并格式化(gcode.CodeInvalidParameter, `invalid service key "%s"`, key)
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
		if err = gjson.Unmarshal别名([]byte(value), &s.Metadata); err != nil {
			err = gerror.X多层错误码并格式化(gcode.CodeInvalidParameter, err, `invalid service value "%s"`, value)
			return nil, err
		}
	}
	return s, nil
}

// GetName 返回服务的名称。
// 名称对于服务是必需的，应在所有服务中保持唯一。
// md5:c0cc1fa5e19d9a6c
func (s *LocalService) GetName() string {
	return s.Name
}

// GetVersion 返回服务的版本号。
// 建议使用GNU版本命名方式，例如：v1.0.0, v2.0.1, v2.1.0-rc。
// 服务可以同时部署多个版本。
// 如果服务中未设置版本，那么服务的默认版本为 "latest"。
// md5:bf857eeaf16711ca
func (s *LocalService) GetVersion() string {
	return s.Version
}

// GetKey 格式化并返回服务的唯一键字符串。
// 生成的结果键通常用于键值注册服务器。
// md5:8651d9bc2f308934
func (s *LocalService) GetKey() string {
	serviceNameUnique := s.GetPrefix()
	serviceNameUnique += DefaultSeparator + s.Endpoints.String()
	return serviceNameUnique
}

// GetValue 格式化并返回服务的值。结果值通常用于键值注册服务器。
// md5:81a88bc4bcc73037
func (s *LocalService) GetValue() string {
	b, err := gjson.Marshal别名(s.Metadata)
	if err != nil {
		intlog.Errorf(context.TODO(), `%+v`, err)
	}
	return string(b)
}

// GetPrefix 格式化并返回键前缀字符串。
// 结果前缀字符串通常用于服务注册服务器中的服务搜索。
//
// 以 etcd 服务器为例，前缀字符串的用法如下：
// `etcdctl get /services/prod/hello.svc --prefix`
// md5:3c443e018050694a
func (s *LocalService) GetPrefix() string {
	s.autoFillDefaultAttributes()
	return DefaultSeparator + gstr.X连接(
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

// GetMetadata 返回服务的元数据地图。
// 元数据是一个键值对映射，用于指定服务的额外属性。
// md5:42fd4200585681c1
func (s *LocalService) GetMetadata() Metadata {
	return s.Metadata
}

// GetEndpoints 返回服务的端点信息。
// 端点包含服务的多个主机/端口信息。
// md5:164bdc2d3a7db5e0
func (s *LocalService) GetEndpoints() Endpoints {
	return s.Endpoints
}

func (s *LocalService) autoFillDefaultAttributes() {
	if s.Head == "" {
		s.Head = gcmd.GetOptWithEnv(EnvPrefix, DefaultHead).String()
	}
	if s.Deployment == "" {
		s.Deployment = gcmd.GetOptWithEnv(EnvDeployment, DefaultDeployment).String()
	}
	if s.Namespace == "" {
		s.Namespace = gcmd.GetOptWithEnv(EnvNamespace, DefaultNamespace).String()
	}
	if s.Name == "" {
		s.Name = gcmd.GetOptWithEnv(EnvName).String()
	}
	if s.Version == "" {
		s.Version = gcmd.GetOptWithEnv(EnvVersion, DefaultVersion).String()
	}
}
