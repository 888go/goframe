// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 日志类

import (
	"context"
	"fmt"
	"strings"
	"time"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/encoding/gcompress"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gmlock"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/os/gtimer"
	"github.com/888go/goframe/text/gregex"
)

const (
	memoryLockPrefixForRotating = "glog.rotateChecksTimely:"
)

// rotateFileBySize 根据配置的旋转大小来旋转当前的日志文件。
func (l *Logger) rotateFileBySize(ctx context.Context, now time.Time) {
	if l.config.RotateSize <= 0 {
		return
	}
	if err := l.doRotateFile(ctx, l.getFilePath(now)); err != nil {
		// panic(err)
		intlog.Errorf(ctx, `%+v`, err)
	}
}

// doRotateFile 对给定的日志文件进行旋转。
func (l *Logger) doRotateFile(ctx context.Context, filePath string) error {
	memoryLockKey := "glog.doRotateFile:" + filePath
	if !内存锁类.X非阻塞写锁定(memoryLockKey) {
		return nil
	}
	defer 内存锁类.X退出写锁定(memoryLockKey)

	intlog.PrintFunc(ctx, func() string {
		return fmt.Sprintf(`start rotating file by size: %s, file: %s`, 文件类.X取大小并易读格式(filePath), filePath)
	})
	defer intlog.PrintFunc(ctx, func() string {
		return fmt.Sprintf(`done rotating file by size: %s, size: %s`, 文件类.X取大小并易读格式(filePath), filePath)
	})

	// 不进行备份，直接删除当前的日志文件。
	if l.config.RotateBackupLimit == 0 {
		if err := 文件类.X删除(filePath); err != nil {
			return err
		}
		intlog.Printf(
			ctx,
			`%d size exceeds, no backups set, remove original logging file: %s`,
			l.config.RotateSize, filePath,
		)
		return nil
	}
	// 否则，它会创建新的备份文件。
	var (
		dirPath     = 文件类.X路径取父目录(filePath)
		fileName    = 文件类.X路径取文件名且不含扩展名(filePath)
		fileExtName = 文件类.X路径取扩展名且不含点号(filePath)
		newFilePath = ""
	)
// 通过向日志文件名添加额外的微秒级日期时间信息进行重命名，例如：
// access.log          -> access.20200326101301899002.log
// access.20200326.log -> access.20200326.20200326101301899002.log
// 这段代码注释是说明一个功能，该功能可以将日志文件名进行重命名，并在原文件名基础上附加精确到微秒级别的日期时间戳。这样做的目的是为了方便管理和区分不同时间段的日志记录。
	for {
		var (
			now   = 时间类.X创建并按当前时间()
			micro = now.X取微秒() % 1000
		)
		if micro == 0 {
			micro = 101
		} else {
			for micro < 100 {
				micro *= 10
			}
		}
		newFilePath = 文件类.X路径生成(
			dirPath,
			fmt.Sprintf(
				`%s.%s%d.%s`,
				fileName, now.X取格式文本("YmdHisu"), micro, fileExtName,
			),
		)
		if !文件类.X是否存在(newFilePath) {
			break
		} else {
			intlog.Printf(ctx, `rotation file exists, continue: %s`, newFilePath)
		}
	}
	intlog.Printf(ctx, "rotating file by size from %s to %s", filePath, newFilePath)
	if err := 文件类.Rename别名(filePath, newFilePath); err != nil {
		return err
	}
	return nil
}

// rotateChecksTimely 定时检查备份的过期情况和压缩状态
func (l *Logger) rotateChecksTimely(ctx context.Context) {
	defer 定时类.X加入单次任务(ctx, l.config.RotateCheckInterval, l.rotateChecksTimely)

	// 检查文件旋转是否未启用。
	if l.config.RotateSize <= 0 && l.config.RotateExpire == 0 {
		intlog.Printf(
			ctx,
			"logging rotation ignore checks: RotateSize: %d, RotateExpire: %s",
			l.config.RotateSize, l.config.RotateExpire.String(),
		)
		return
	}

	// 这里使用内存锁来保证并发安全性。
	memoryLockKey := memoryLockPrefixForRotating + l.config.Path
	if !内存锁类.X非阻塞写锁定(memoryLockKey) {
		return
	}
	defer 内存锁类.X退出写锁定(memoryLockKey)

	var (
		now        = time.Now()
		pattern    = "*.log, *.gz"
		files, err = 文件类.X枚举(l.config.Path, pattern, true)
	)
	if err != nil {
		intlog.Errorf(ctx, `%+v`, err)
	}
	intlog.Printf(ctx, "logging rotation start checks: %+v", files)
// 获取文件名正则表达式模式
// access-{y-m-d}-test.log => access-$-test.log => access-\$-test\.log => access-(.+?)-test\.log
// 原始格式的文件名中，{y-m-d}代表日期，将其转换为正则表达式模式
// 首先将大括号替换为美元符号($)，但在正则表达式中有特殊含义，因此需要转义为'\$'
// 然后将日期部分转换为一个可以匹配任何字符序列的非贪婪模式组".+?"
// 最终得到的正则表达式 "access-(.+?)-test\.log" 可以匹配类似于 "access-2022-01-01-test.log" 这样的文件名
	fileNameRegexPattern, _ := 正则类.X替换文本(`{.+?}`, "$", l.config.File)
	fileNameRegexPattern = 正则类.X转义特殊符号(fileNameRegexPattern)
	fileNameRegexPattern = strings.ReplaceAll(fileNameRegexPattern, "\\$", "(.+?)")
// =============================================================
// 过期文件检查的轮转机制。
// =============================================================
	if l.config.RotateExpire > 0 {
		var (
			mtime         time.Time
			subDuration   time.Duration
			expireRotated bool
		)
		for _, file := range files {
			// 忽略备份文件
			if 正则类.X是否匹配文本(`.+\.\d{20}\.log`, 文件类.X路径取文件名(file)) {
				continue
			}
			// 忽略不匹配的文件
			if !正则类.X是否匹配文本(fileNameRegexPattern, file) {
				continue
			}
			mtime = 文件类.X取修改时间秒(file)
			subDuration = now.Sub(mtime)
			if subDuration > l.config.RotateExpire {
				func() {
					memoryLockFileKey := memoryLockPrefixForPrintingToFile + file
					if !内存锁类.X非阻塞写锁定(memoryLockFileKey) {
						return
					}
					defer 内存锁类.X退出写锁定(memoryLockFileKey)
					expireRotated = true
					intlog.Printf(
						ctx,
						`%v - %v = %v > %v, rotation expire logging file: %s`,
						now, mtime, subDuration, l.config.RotateExpire, file,
					)
					if err = l.doRotateFile(ctx, file); err != nil {
						intlog.Errorf(ctx, `%+v`, err)
					}
				}()
			}
		}
		if expireRotated {
			// 更新文件数组。
			files, err = 文件类.X枚举(l.config.Path, pattern, true)
			if err != nil {
				intlog.Errorf(ctx, `%+v`, err)
			}
		}
	}

// =============================================================
// 旋转文件压缩。
// =============================================================
// 这段注释表明该段Go语言代码是用于实现“旋转文件压缩”功能的。在日志处理、数据备份等场景中，当文件达到一定大小或满足特定条件时，会创建新的文件并将旧文件进行压缩，这个过程通常称为“文件旋转”（File Rotation）。本代码块可能涉及对已旋转的文件进行压缩操作。
	needCompressFileArray := 数组类.X创建文本()
	if l.config.RotateBackupCompress > 0 {
		for _, file := range files {
			// 示例：access.20200326101301899002.log.gz
// 这段Go语言代码注释的中文翻译如下：
// ```go
// 示例：access.20200326101301899002.log.gz
// 这是一个文件名示例，表示一个在2020年3月26日10时13分01秒创建的访问日志文件，
// 并且经过了gzip压缩。文件名中包含了时间戳信息用于标识记录的时间点。
			if 文件类.X路径取扩展名且不含点号(file) == "gz" {
				continue
			}
			// 忽略不匹配的文件
			originalLoggingFilePath, _ := 正则类.X替换文本(`\.\d{20}`, "", file)
			if !正则类.X是否匹配文本(fileNameRegexPattern, originalLoggingFilePath) {
				continue
			}
// 示例：
// access.20200326101301899002.log
// （该行代码为文件名注释，意为：这是一个日志文件的示例，文件名为“access”，后跟创建日期时间戳“20200326101301899002”，并以“.log”为扩展名。）
			if 正则类.X是否匹配文本(`.+\.\d{20}\.log`, 文件类.X路径取文件名(file)) {
				needCompressFileArray.Append别名(file)
			}
		}
		if needCompressFileArray.X取长度() > 0 {
			needCompressFileArray.X遍历(func(_ int, path string) bool {
				err := 压缩类.Gzip压缩文件(path, path+".gz")
				if err == nil {
					intlog.Printf(ctx, `compressed done, remove original logging file: %s`, path)
					if err = 文件类.X删除(path); err != nil {
						intlog.Print(ctx, err)
					}
				} else {
					intlog.Print(ctx, err)
				}
				return true
			})
			// 更新文件数组。
			files, err = 文件类.X枚举(l.config.Path, pattern, true)
			if err != nil {
				intlog.Errorf(ctx, `%+v`, err)
			}
		}
	}

// =============================================================
// 备份数量限制及过期检查
// =============================================================
	backupFiles := 数组类.X创建排序(func(a, b interface{}) int {
// 按照旋转/备份文件的修改时间进行排序。
// 较旧的旋转/备份文件被放在数组的头部。
		var (
			file1  = a.(string)
			file2  = b.(string)
			result = 文件类.X取修改时间戳毫秒(file1) - 文件类.X取修改时间戳毫秒(file2)
		)
		if result <= 0 {
			return -1
		}
		return 1
	})
	if l.config.RotateBackupLimit > 0 || l.config.RotateBackupExpire > 0 {
		for _, file := range files {
			// 忽略不匹配的文件
			originalLoggingFilePath, _ := 正则类.X替换文本(`\.\d{20}`, "", file)
			if !正则类.X是否匹配文本(fileNameRegexPattern, originalLoggingFilePath) {
				continue
			}
			if 正则类.X是否匹配文本(`.+\.\d{20}\.log`, 文件类.X路径取文件名(file)) {
				backupFiles.X入栈右(file)
			}
		}
		intlog.Printf(ctx, `calculated backup files array: %+v`, backupFiles)
		diff := backupFiles.X取长度() - l.config.RotateBackupLimit
		for i := 0; i < diff; i++ {
			path, _ := backupFiles.X出栈左()
			intlog.Printf(ctx, `remove exceeded backup limit file: %s`, path)
			if err := 文件类.X删除(path.(string)); err != nil {
				intlog.Errorf(ctx, `%+v`, err)
			}
		}
		// 备份过期检查
		if l.config.RotateBackupExpire > 0 {
			var (
				mtime       time.Time
				subDuration time.Duration
			)
			backupFiles.X遍历(func(_ int, v interface{}) bool {
				path := v.(string)
				mtime = 文件类.X取修改时间秒(path)
				subDuration = now.Sub(mtime)
				if subDuration > l.config.RotateBackupExpire {
					intlog.Printf(
						ctx,
						`%v - %v = %v > %v, remove expired backup file: %s`,
						now, mtime, subDuration, l.config.RotateBackupExpire, path,
					)
					if err := 文件类.X删除(path); err != nil {
						intlog.Errorf(ctx, `%+v`, err)
					}
					return true
				} else {
					return false
				}
			})
		}
	}
}
