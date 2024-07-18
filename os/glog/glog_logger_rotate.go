// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package glog

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/encoding/gcompress"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gmlock"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gogf/gf/v2/text/gregex"
)

const (
	memoryLockPrefixForRotating = "glog.rotateChecksTimely:"
)

// rotateFileBySize 根据配置的旋转大小，旋转当前日志文件。
// md5:e07365ed108ab9ed
func (l *Logger) rotateFileBySize(ctx context.Context, now time.Time) {
	if l.config.RotateSize <= 0 {
		return
	}
	if err := l.doRotateFile(ctx, l.getFilePath(now)); err != nil {
		// panic(err)
		intlog.Errorf(ctx, `%+v`, err)
	}
}

// doRotateFile 旋转给定的日志文件。 md5:c1b732cfe2000ccf
func (l *Logger) doRotateFile(ctx context.Context, filePath string) error {
	memoryLockKey := "glog.doRotateFile:" + filePath
	if !gmlock.TryLock(memoryLockKey) {
		return nil
	}
	defer gmlock.Unlock(memoryLockKey)

	intlog.PrintFunc(ctx, func() string {
		return fmt.Sprintf(`start rotating file by size: %s, file: %s`, gfile.SizeFormat(filePath), filePath)
	})
	defer intlog.PrintFunc(ctx, func() string {
		return fmt.Sprintf(`done rotating file by size: %s, size: %s`, gfile.SizeFormat(filePath), filePath)
	})

	// 无备份情况下，它将直接删除当前的日志文件。 md5:66cbeaeb716f06ee
	if l.config.RotateBackupLimit == 0 {
		if err := gfile.Remove(filePath); err != nil {
			return err
		}
		intlog.Printf(
			ctx,
			`%d size exceeds, no backups set, remove original logging file: %s`,
			l.config.RotateSize, filePath,
		)
		return nil
	}
	// 否则，它将创建新的备份文件。 md5:98bfbd0a3d10fcb0
	var (
		dirPath     = gfile.Dir(filePath)
		fileName    = gfile.Name(filePath)
		fileExtName = gfile.ExtName(filePath)
		newFilePath = ""
	)
// 通过在日志文件名中添加额外的日期时间信息（到微秒级别），重命名日志文件，例如：
// access.log            -> access.20200326101301899002.log
// access.20200326.log   -> access.20200326.20200326101301899002.log
// md5:96d2e4456a2a561d
	for {
		var (
			now   = gtime.Now()
			micro = now.Microsecond() % 1000
		)
		if micro == 0 {
			micro = 101
		} else {
			for micro < 100 {
				micro *= 10
			}
		}
		newFilePath = gfile.Join(
			dirPath,
			fmt.Sprintf(
				`%s.%s%d.%s`,
				fileName, now.Format("YmdHisu"), micro, fileExtName,
			),
		)
		if !gfile.Exists(newFilePath) {
			break
		} else {
			intlog.Printf(ctx, `rotation file exists, continue: %s`, newFilePath)
		}
	}
	intlog.Printf(ctx, "rotating file by size from %s to %s", filePath, newFilePath)
	if err := gfile.Rename(filePath, newFilePath); err != nil {
		return err
	}
	return nil
}

// timelyChecksTimely检查备份的过期和压缩。 md5:0502efeb887ae657
func (l *Logger) rotateChecksTimely(ctx context.Context) {
	defer gtimer.AddOnce(ctx, l.config.RotateCheckInterval, l.rotateChecksTimely)

	// 检查文件旋转是否未启用。 md5:22b3a5305aaec48c
	if l.config.RotateSize <= 0 && l.config.RotateExpire == 0 {
		intlog.Printf(
			ctx,
			"logging rotation ignore checks: RotateSize: %d, RotateExpire: %s",
			l.config.RotateSize, l.config.RotateExpire.String(),
		)
		return
	}

	// 此处使用内存锁来保证并发安全性。 md5:a621f4c111c27699
	memoryLockKey := memoryLockPrefixForRotating + l.config.Path
	if !gmlock.TryLock(memoryLockKey) {
		return
	}
	defer gmlock.Unlock(memoryLockKey)

	var (
		now        = time.Now()
		pattern    = "*.log, *.gz"
		files, err = gfile.ScanDirFile(l.config.Path, pattern, true)
	)
	if err != nil {
		intlog.Errorf(ctx, `%+v`, err)
	}
	intlog.Printf(ctx, "logging rotation start checks: %+v", files)
// 获取文件名正则表达式模式
// access-{yyyy-mm-dd}-test.log => access-${}-test.log => access-\$\-test\.log => access-(\w+)-test\.log
// 
// 这段注释说明了一个正则表达式规则，用于从文件名中提取部分。原始格式是`access-yyyy-mm-dd-test.log`，经过转换后，它首先替换`{}`为`-`（`access-yyyy-mm-dd-test.log` => `access-yyyy-mm-dd-test.log`），然后替换`-`为`\`（`access-yyyy-mm-dd-test.log` => `access-$-test.log`），再进一步替换`\`为`\`（`access-$-test.log` => `access-\$-test\.log`），最后使用正向前瞻断言匹配一个或多个任意字符但不包括`-`（`access-\$-test\.log` => `access-(.+?)-test\.log`），这样就可以匹配如`access-2021-08-31-test.log`这样的文件名。
// md5:e9cbde6eccd06a32
	fileNameRegexPattern, _ := gregex.ReplaceString(`{.+?}`, "$", l.config.File)
	fileNameRegexPattern = gregex.Quote(fileNameRegexPattern)
	fileNameRegexPattern = strings.ReplaceAll(fileNameRegexPattern, "\\$", "(.+?)")
// =============================================================
// 无效文件检查的旋转
// =============================================================
// md5:2ac41d9c8ed6dcd1
	if l.config.RotateExpire > 0 {
		var (
			mtime         time.Time
			subDuration   time.Duration
			expireRotated bool
		)
		for _, file := range files {
			// ignore backup file
			if gregex.IsMatchString(`.+\.\d{20}\.log`, gfile.Basename(file)) || gfile.ExtName(file) == "gz" {
				continue
			}
			// 忽略不匹配的文件. md5:a1b51f5b82391575
			if !gregex.IsMatchString(fileNameRegexPattern, file) {
				continue
			}
			mtime = gfile.MTime(file)
			subDuration = now.Sub(mtime)
			if subDuration > l.config.RotateExpire {
				func() {
					memoryLockFileKey := memoryLockPrefixForPrintingToFile + file
					if !gmlock.TryLock(memoryLockFileKey) {
						return
					}
					defer gmlock.Unlock(memoryLockFileKey)
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
			// 更新文件数组。 md5:eb6c80314da4cb7a
			files, err = gfile.ScanDirFile(l.config.Path, pattern, true)
			if err != nil {
				intlog.Errorf(ctx, `%+v`, err)
			}
		}
	}

// =============================================================
// 旋转文件压缩。
// =============================================================
// md5:c028a879a3e48da1
	needCompressFileArray := garray.NewStrArray()
	if l.config.RotateBackupCompress > 0 {
		for _, file := range files {
			// 例如：access.20200326101301899002.log.gz. md5:e037aa543e2a446f
			if gfile.ExtName(file) == "gz" {
				continue
			}
			// 忽略不匹配的文件. md5:a1b51f5b82391575
			originalLoggingFilePath, _ := gregex.ReplaceString(`\.\d{20}`, "", file)
			if !gregex.IsMatchString(fileNameRegexPattern, originalLoggingFilePath) {
				continue
			}
// 示例：
// access.20200326101301899002.log
// 
// 这个注释没有明确的翻译需求，因为它本身就是表示一个文件名样例，其中包含了日期和可能的访问记录序列号。如果需要解释其结构含义，可以这样翻译：
// 
// 示例文件名：
// 访问日志文件，格式为"access.日期(YYYYMMDDHHMMSS).序列号.log"
// 例如：access.2020年03月26日10时13分01秒899002序列号.log
// md5:08ddd9e8cc49fee7
			if gregex.IsMatchString(`.+\.\d{20}\.log`, gfile.Basename(file)) {
				needCompressFileArray.Append(file)
			}
		}
		if needCompressFileArray.Len() > 0 {
			needCompressFileArray.Iterator(func(_ int, path string) bool {
				err := gcompress.GzipFile(path, path+".gz")
				if err == nil {
					intlog.Printf(ctx, `compressed done, remove original logging file: %s`, path)
					if err = gfile.Remove(path); err != nil {
						intlog.Print(ctx, err)
					}
				} else {
					intlog.Print(ctx, err)
				}
				return true
			})
			// 更新文件数组。 md5:eb6c80314da4cb7a
			files, err = gfile.ScanDirFile(l.config.Path, pattern, true)
			if err != nil {
				intlog.Errorf(ctx, `%+v`, err)
			}
		}
	}

// =============================================================
// 备份数量限制和过期检查。
// =============================================================
// md5:7edc3bfeec7fde2e
	backupFiles := garray.NewSortedArray(func(a, b interface{}) int {
// 按照旋转/备份文件的mtime（修改时间）排序。
// 老的旋转/备份文件被放在数组的头部。
// md5:7ead56b6a771900f
		var (
			file1  = a.(string)
			file2  = b.(string)
			result = gfile.MTimestampMilli(file1) - gfile.MTimestampMilli(file2)
		)
		if result <= 0 {
			return -1
		}
		return 1
	})
	if l.config.RotateBackupLimit > 0 || l.config.RotateBackupExpire > 0 {
		for _, file := range files {
			// 忽略不匹配的文件. md5:a1b51f5b82391575
			originalLoggingFilePath, _ := gregex.ReplaceString(`\.\d{20}`, "", file)
			if !gregex.IsMatchString(fileNameRegexPattern, originalLoggingFilePath) {
				continue
			}
			if gregex.IsMatchString(`.+\.\d{20}\.log`, gfile.Basename(file)) {
				backupFiles.Add(file)
			}
		}
		intlog.Printf(ctx, `calculated backup files array: %+v`, backupFiles)
		diff := backupFiles.Len() - l.config.RotateBackupLimit
		for i := 0; i < diff; i++ {
			path, _ := backupFiles.PopLeft()
			intlog.Printf(ctx, `remove exceeded backup limit file: %s`, path)
			if err := gfile.Remove(path.(string)); err != nil {
				intlog.Errorf(ctx, `%+v`, err)
			}
		}
		// 备份过期检查。 md5:f974bc9ca93e7536
		if l.config.RotateBackupExpire > 0 {
			var (
				mtime       time.Time
				subDuration time.Duration
			)
			backupFiles.Iterator(func(_ int, v interface{}) bool {
				path := v.(string)
				mtime = gfile.MTime(path)
				subDuration = now.Sub(mtime)
				if subDuration > l.config.RotateBackupExpire {
					intlog.Printf(
						ctx,
						`%v - %v = %v > %v, remove expired backup file: %s`,
						now, mtime, subDuration, l.config.RotateBackupExpire, path,
					)
					if err := gfile.Remove(path); err != nil {
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
