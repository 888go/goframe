// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gins

import (
	"fmt"
	
	"github.com/888go/goframe/internal/instance"
	"github.com/888go/goframe/net/gclient"
)

// HttpClient 返回具有指定名称的 http 客户端实例。
func HttpClient(name ...interface{}) *gclient.Client {
	var instanceKey = fmt.Sprintf("%s.%v", frameCoreComponentNameHttpClient, name)
	return instance.GetOrSetFuncLock(instanceKey, func() interface{} {
		return gclient.New()
	}).(*gclient.Client)
}
