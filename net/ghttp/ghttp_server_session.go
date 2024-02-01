// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp
import (
	"github.com/888go/goframe/os/gsession"
	)
// Session 实际上是 gsession.Session 的别名，
// 并且它是与单个请求绑定的。
type Session = gsession.Session
