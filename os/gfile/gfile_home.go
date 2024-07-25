// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gfile

import (
	"bytes"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
)

// Home 返回当前用户主目录的绝对路径。
// 可选参数 `names` 指定了要与当前系统分隔符连接的子文件夹/子文件，将与路径一起返回。 md5:0b575ba0215ebb2d
func Home(names ...string) (string, error) {
	path, err := getHomePath()
	if err != nil {
		return "", err
	}
	for _, name := range names {
		path += Separator + name
	}
	return path, nil
}

// getHomePath 返回当前用户主目录的绝对路径。 md5:81fe93a610949935
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

// homeUnix 在Unix系统中获取并返回家目录。 md5:8a70582364ac0874
func homeUnix() (string, error) {
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		err = gerror.Wrapf(err, `retrieve home directory failed`)
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", gerror.New("blank output when reading home directory")
	}

	return result, nil
}

// homeWindows 获取并返回Windows系统中的主目录。 md5:211bc3c03e9cf044
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
		return "", gerror.New("environment keys HOMEDRIVE, HOMEPATH and USERPROFILE are empty")
	}

	return home, nil
}
