// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 日志类

import (
	"github.com/fatih/color"
)

const (
	COLOR_BLACK = 30 + iota
	COLOR_RED
	COLOR_GREEN
	COLOR_YELLOW
	COLOR_BLUE
	COLOR_MAGENTA
	COLOR_CYAN
	COLOR_WHITE
)

// 前景高亮度文本颜色
const (
	COLOR_HI_BLACK = 90 + iota
	COLOR_HI_RED
	COLOR_HI_GREEN
	COLOR_HI_YELLOW
	COLOR_HI_BLUE
	COLOR_HI_MAGENTA
	COLOR_HI_CYAN
	COLOR_HI_WHITE
)

// defaultLevelColor 定义默认日志级别及其映射前缀字符串。
var defaultLevelColor = map[int]int{
	LEVEL_DEBU: COLOR_YELLOW,
	LEVEL_INFO: COLOR_GREEN,
	LEVEL_NOTI: COLOR_CYAN,
	LEVEL_WARN: COLOR_MAGENTA,
	LEVEL_ERRO: COLOR_RED,
	LEVEL_CRIT: COLOR_HI_RED,
	LEVEL_PANI: COLOR_HI_RED,
	LEVEL_FATA: COLOR_HI_RED,
}

// getColoredStr 根据给定的颜色返回一个带颜色的字符串。
func (l *Logger) getColoredStr(c int, s string) string {
	return color.New(color.Attribute(c)).Sprint(s)
}

func (l *Logger) getColorByLevel(level int) int {
	return defaultLevelColor[level]
}
