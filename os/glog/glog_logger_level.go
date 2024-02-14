// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 日志类

import (
	"strings"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
)

// 注意，LEVEL_PANI 和 LEVEL_FATA 级别并不用于日志输出，
// 而是用于前缀配置。
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

// defaultLevelPrefixes 定义了默认的日志级别及其对应的前缀字符串。
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

// levelStringMap 定义了日志级别字符串名称与其对应的级别映射。
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

// SetLevel 设置日志记录级别。
// 注意，对于日志内容，级别 `LEVEL_CRIT | LEVEL_PANI | LEVEL_FATA` 不能被移除，
// 这些级别会自动添加到现有级别中。
func (l *Logger) X设置级别(级别 int) {
	l.config.X级别 = 级别 | LEVEL_CRIT | LEVEL_PANI | LEVEL_FATA
}

// GetLevel 返回日志等级值。
func (l *Logger) X取级别() int {
	return l.config.X级别
}

// SetLevelStr 通过级别字符串设置日志记录级别。
func (l *Logger) X设置文本级别(级别 string) error {
	if level, ok := levelStringMap[strings.ToUpper(级别)]; ok {
		l.config.X级别 = level
	} else {
		return 错误类.X创建错误码并格式化(错误码类.CodeInvalidParameter, `invalid level string: %s`, 级别)
	}
	return nil
}

// SetLevelPrefix 为指定等级设置前缀字符串。
func (l *Logger) X设置级别前缀(级别 int, 前缀 string) {
	l.config.X日志级别名称映射[级别] = 前缀
}

// SetLevelPrefixes 为日志器设置级别到前缀字符串的映射。
func (l *Logger) X设置级别前缀Map(前缀Map map[int]string) {
	for k, v := range 前缀Map {
		l.config.X日志级别名称映射[k] = v
	}
}

// GetLevelPrefix 返回指定级别的前缀字符串。
func (l *Logger) X取级别前缀(级别 int) string {
	return l.config.X日志级别名称映射[级别]
}

// getLevelPrefixWithBrackets 根据指定级别返回带括号的前缀字符串。
func (l *Logger) getLevelPrefixWithBrackets(level int) string {
	levelStr := ""
	if s, ok := l.config.X日志级别名称映射[level]; ok {
		levelStr = "[" + s + "]"
	}
	return levelStr
}
