// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package glog

import "github.com/fatih/color"

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

// 前景高亮度文本颜色. md5:e39d29d745b2e70e
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

// defaultLevelColor 定义了默认级别及其前缀字符串。 md5:c6a4c65e46485929
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

// getColoredStr 返回一个使用给定颜色着色的字符串。 md5:accfa5b8ad258119
func (l *Logger) getColoredStr(c int, s string) string {
	return color.New(color.Attribute(c)).Sprint(s)
}

func (l *Logger) getColorByLevel(level int) int {
	return defaultLevelColor[level]
}
