// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gins

import (
	"fmt"

	"github.com/888go/goframe/internal/instance"
	gclient "github.com/888go/goframe/net/gclient"
)

// HttpClient 返回指定名称的 http 客户端实例。 md5:4ad3009bae654769
func HttpClient(name ...interface{}) *gclient.Client {
	var instanceKey = fmt.Sprintf("%s.%v", frameCoreComponentNameHttpClient, name)
	return instance.GetOrSetFuncLock(instanceKey, func() interface{} {
		return gclient.X创建()
	}).(*gclient.Client)
}
