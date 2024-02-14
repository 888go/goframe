// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"fmt"
	"net/http"
	"strings"
	
	"github.com/888go/goframe/encoding/gbase64"
)

// BasicAuth 通过给定的用户名和密码启用http基本认证功能，
// 并请求客户端进行身份验证。如果认证成功则返回 true，否则（认证失败）返回 false。
func (r *X请求) X账号密码认证(账号, 密码 string, 可选提示 ...string) bool {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		r.setBasicAuth(可选提示...)
		return false
	}
	authArray := strings.SplitN(auth, " ", 2)
	if len(authArray) != 2 {
		r.X响应.X写响应缓冲区与HTTP状态码(http.StatusForbidden)
		return false
	}
	switch authArray[0] {
	case "Basic":
		authBytes, err := 编码base64类.X文本解码到字节集(authArray[1])
		if err != nil {
			r.X响应.X写响应缓冲区与HTTP状态码(http.StatusForbidden, err.Error())
			return false
		}
		authArray := strings.SplitN(string(authBytes), ":", 2)
		if len(authArray) != 2 {
			r.X响应.X写响应缓冲区与HTTP状态码(http.StatusForbidden)
			return false
		}
		if authArray[0] != 账号 || authArray[1] != 密码 {
			r.setBasicAuth(可选提示...)
			return false
		}
		return true

	default:
		r.X响应.X写响应缓冲区与HTTP状态码(http.StatusForbidden)
		return false
	}
}

// setBasicAuth 设置HTTP基本认证提示
func (r *X请求) setBasicAuth(tips ...string) {
	realm := ""
	if len(tips) > 0 && tips[0] != "" {
		realm = tips[0]
	} else {
		realm = "Need Login"
	}
	r.X响应.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
	r.X响应.WriteHeader(http.StatusUnauthorized)
}
