// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 日志类

import (
	"strings"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
)

// 注意，LEVEL_PANI和LEVEL_FATA级别不用于日志输出，而是用于前缀配置。
// md5:991a3476bacb665d
const (
	LEVEL_ALL  = LEVEL_DEBU | LEVEL_INFO | LEVEL_NOTI | LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT
	LEVEL_DEV  = LEVEL_ALL
	LEVEL_PROD = LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT
	LEVEL_NONE = 0
	LEVEL_DEBU = 1 << iota // 16
	LEVEL_INFO             // 32
	LEVEL_NOTI             // 64
	LEVEL_WARN             // 128
	LEVEL_ERRO             // 256
	LEVEL_CRIT             // 512
	LEVEL_PANI             // 1024
	LEVEL_FATA             // 2048
)

// defaultLevelPrefixes 定义了默认的日志级别及其对应的前缀字符串。 md5:9e3d8d403bb7bbce
var defaultLevelPrefixes = map[int]string{
	LEVEL_DEBU: "DEBU",
	LEVEL_INFO: "INFO",
	LEVEL_NOTI: "NOTI",
	LEVEL_WARN: "WARN",
	LEVEL_ERRO: "ERRO",
	LEVEL_CRIT: "CRIT",
	LEVEL_PANI: "PANI",
	LEVEL_FATA: "FATA",
}

// levelStringMap 定义了日志级别字符串名称到其级别的映射。 md5:d9c861b2e6837843
var levelStringMap = map[string]int{
	"ALL":      LEVEL_DEBU | LEVEL_INFO | LEVEL_NOTI | LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT,
	"DEV":      LEVEL_DEBU | LEVEL_INFO | LEVEL_NOTI | LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT,
	"DEVELOP":  LEVEL_DEBU | LEVEL_INFO | LEVEL_NOTI | LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT,
	"PROD":     LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT,
	"PRODUCT":  LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT,
	"DEBU":     LEVEL_DEBU | LEVEL_INFO | LEVEL_NOTI | LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT,
	"DEBUG":    LEVEL_DEBU | LEVEL_INFO | LEVEL_NOTI | LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT,
	"INFO":     LEVEL_INFO | LEVEL_NOTI | LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT,
	"NOTI":     LEVEL_NOTI | LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT,
	"NOTICE":   LEVEL_NOTI | LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT,
	"WARN":     LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT,
	"WARNING":  LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT,
	"ERRO":     LEVEL_ERRO | LEVEL_CRIT,
	"ERROR":    LEVEL_ERRO | LEVEL_CRIT,
	"CRIT":     LEVEL_CRIT,
	"CRITICAL": LEVEL_CRIT,
}

// X设置级别 设置日志级别。
// 注意，`LEVEL_CRIT`、`LEVEL_PANI` 和 `LEVEL_FATA` 级别无法删除，因为它们会自动添加到日志内容中。
// md5:e488e79c6c4c2e71
func (l *Logger) X设置级别(级别 int) {
	l.config.Level = 级别 | LEVEL_CRIT | LEVEL_PANI | LEVEL_FATA
}

// X取级别 返回日志级别值。 md5:8b5b1f26924bf360
func (l *Logger) X取级别() int {
	return l.config.Level
}

// X设置文本级别 通过级别字符串设置日志级别。 md5:53cbbdf23584340e
func (l *Logger) X设置文本级别(级别 string) error {
	if level, ok := levelStringMap[strings.ToUpper(级别)]; ok {
		l.config.Level = level
	} else {
		return gerror.X创建错误码并格式化(gcode.CodeInvalidParameter, `invalid level string: %s`, 级别)
	}
	return nil
}

// X设置级别前缀 为指定的日志级别设置前缀字符串。 md5:a2b7a43af150bcb7
func (l *Logger) X设置级别前缀(级别 int, 前缀 string) {
	l.config.LevelPrefixes[级别] = 前缀
}

// X设置级别前缀Map 为记录器设置级别与前缀字符串的映射关系。 md5:a80f5e3de3c222ff
func (l *Logger) X设置级别前缀Map(前缀Map map[int]string) {
	for k, v := range 前缀Map {
		l.config.LevelPrefixes[k] = v
	}
}

// X取级别前缀 返回指定级别的前缀字符串。 md5:339b86b4f84d6049
func (l *Logger) X取级别前缀(级别 int) string {
	return l.config.LevelPrefixes[级别]
}

// getLevelPrefixWithBrackets 返回指定级别的带有括号的前缀字符串。 md5:a323f3c9d4c0ab4c
func (l *Logger) getLevelPrefixWithBrackets(level int) string {
	levelStr := ""
	if s, ok := l.config.LevelPrefixes[level]; ok {
		levelStr = "[" + s + "]"
	}
	return levelStr
}
