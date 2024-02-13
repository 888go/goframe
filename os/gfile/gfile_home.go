// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件类

import (
	"bytes"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
	
	"github.com/888go/goframe/errors/gerror"
)

// Home 返回当前用户主目录的绝对路径。
// 可选参数 `names` 指定了子文件夹或子文件，
// 这些名称将与当前系统分隔符连接，并与路径一起返回。
func X取用户目录(可选子目录或子文件 ...string) (string, error) {
	path, err := getHomePath()
	if err != nil {
		return "", err
	}
	for _, name := range 可选子目录或子文件 {
		path += Separator + name
	}
	return path, nil
}

// getHomePath 返回当前用户主目录的绝对路径。
func getHomePath() (string, error) {
	u, err := user.Current()
	if nil == err {
		return u.HomeDir, nil
	}
	if runtime.GOOS == "windows" {
		return homeWindows()
	}
	return homeUnix()
}

// homeUnix在Unix系统上获取并返回用户的主目录。
func homeUnix() (string, error) {
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		err = 错误类.X多层错误并格式化(err, `retrieve home directory failed`)
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", 错误类.X创建("blank output when reading home directory")
	}

	return result, nil
}

// homeWindows 在Windows系统上获取并返回用户的主目录。
func homeWindows() (string, error) {
	var (
		drive = os.Getenv("HOMEDRIVE")
		path  = os.Getenv("HOMEPATH")
		home  = drive + path
	)
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", 错误类.X创建("environment keys HOMEDRIVE, HOMEPATH and USERPROFILE are empty")
	}

	return home, nil
}
