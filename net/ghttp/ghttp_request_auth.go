// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp

import (
	"fmt"
	"net/http"
	"strings"
	
	"github.com/888go/goframe/encoding/gbase64"
)

// BasicAuth 通过给定的用户名和密码启用http基本认证功能，
// 并请求客户端进行身份验证。如果认证成功则返回 true，否则（认证失败）返回 false。
func (r *Request) BasicAuth(user, pass string, tips ...string) bool {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		r.setBasicAuth(tips...)
		return false
	}
	authArray := strings.SplitN(auth, " ", 2)
	if len(authArray) != 2 {
		r.Response.WriteStatus(http.StatusForbidden)
		return false
	}
	switch authArray[0] {
	case "Basic":
		authBytes, err := gbase64.DecodeString(authArray[1])
		if err != nil {
			r.Response.WriteStatus(http.StatusForbidden, err.Error())
			return false
		}
		authArray := strings.SplitN(string(authBytes), ":", 2)
		if len(authArray) != 2 {
			r.Response.WriteStatus(http.StatusForbidden)
			return false
		}
		if authArray[0] != user || authArray[1] != pass {
			r.setBasicAuth(tips...)
			return false
		}
		return true

	default:
		r.Response.WriteStatus(http.StatusForbidden)
		return false
	}
}

// setBasicAuth 设置HTTP基本认证提示
func (r *Request) setBasicAuth(tips ...string) {
	realm := ""
	if len(tips) > 0 && tips[0] != "" {
		realm = tips[0]
	} else {
		realm = "Need Login"
	}
	r.Response.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
	r.Response.WriteHeader(http.StatusUnauthorized)
}
