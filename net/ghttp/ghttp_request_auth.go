// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"fmt"
	"net/http"
	"strings"

	gbase64 "github.com/888go/goframe/encoding/gbase64"
)

// X账号密码认证 启用HTTP基本认证功能，使用给定的凭证（passport）和密码，并要求客户端进行身份验证。如果认证成功，返回true；否则，如果认证失败，返回false。
// md5:8ea275597053f51d
func (r *Request) X账号密码认证(账号, 密码 string, 可选提示 ...string) bool {
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
		authBytes, err := gbase64.X文本解码到字节集(authArray[1])
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

// setBasicAuth 设置HTTP基本认证信息。 md5:7bb0db8710e057f8
func (r *Request) setBasicAuth(tips ...string) {
	realm := ""
	if len(tips) > 0 && tips[0] != "" {
		realm = tips[0]
	} else {
		realm = "Need Login"
	}
	r.X响应.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
	r.X响应.WriteHeader(http.StatusUnauthorized)
}
