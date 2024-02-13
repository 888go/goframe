// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 日志类

import (
	"bytes"
	"fmt"
	"testing"
	"time"
	
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
)

func Test_To(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		w := bytes.NewBuffer(nil)
		X重定向输出(w).Error(ctx, 1, 2, 3)
		X重定向输出(w).X输出并格式化ERR(ctx, "%d %d %d", 1, 2, 3)
		t.Assert(文本类.X统计次数(w.String(), defaultLevelPrefixes[LEVEL_ERRO]), 2)
		t.Assert(文本类.X统计次数(w.String(), "1 2 3"), 2)
	})
}

func Test_Path(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		path := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, 时间类.X取时间戳纳秒())

		err := 文件类.X创建目录(path)
		t.AssertNil(err)
		defer 文件类.X删除(path)

		X文件路径(path).X文件名格式(file).X是否同时输出到终端(false).Error(ctx, 1, 2, 3)
		X文件路径(path).X文件名格式(file).X是否同时输出到终端(false).X输出并格式化ERR(ctx, "%d %d %d", 1, 2, 3)
		content := 文件类.X读文本(文件类.X路径生成(path, file))
		t.Assert(文本类.X统计次数(content, defaultLevelPrefixes[LEVEL_ERRO]), 2)
		t.Assert(文本类.X统计次数(content, "1 2 3"), 2)
	})
}

func Test_Cat(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cat := "category"
		path := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, 时间类.X取时间戳纳秒())

		err := 文件类.X创建目录(path)
		t.AssertNil(err)
		defer 文件类.X删除(path)

		X文件路径(path).X文件名格式(file).X文件分类(cat).X是否同时输出到终端(false).Error(ctx, 1, 2, 3)
		X文件路径(path).X文件名格式(file).X文件分类(cat).X是否同时输出到终端(false).X输出并格式化ERR(ctx, "%d %d %d", 1, 2, 3)
		content := 文件类.X读文本(文件类.X路径生成(path, cat, file))
		t.Assert(文本类.X统计次数(content, defaultLevelPrefixes[LEVEL_ERRO]), 2)
		t.Assert(文本类.X统计次数(content, "1 2 3"), 2)
	})
}

func Test_Level(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		path := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, 时间类.X取时间戳纳秒())

		err := 文件类.X创建目录(path)
		t.AssertNil(err)
		defer 文件类.X删除(path)

		X文件路径(path).X文件名格式(file).X级别(LEVEL_PROD).X是否同时输出到终端(false).X输出DEBU(ctx, 1, 2, 3)
		X文件路径(path).X文件名格式(file).X级别(LEVEL_PROD).X是否同时输出到终端(false).X输出DEBU(ctx, "%d %d %d", 1, 2, 3)
		content := 文件类.X读文本(文件类.X路径生成(path, file))
		t.Assert(文本类.X统计次数(content, defaultLevelPrefixes[LEVEL_DEBU]), 0)
		t.Assert(文本类.X统计次数(content, "1 2 3"), 0)
	})
}

func Test_Skip(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		path := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, 时间类.X取时间戳纳秒())

		err := 文件类.X创建目录(path)
		t.AssertNil(err)
		defer 文件类.X删除(path)

		X文件路径(path).X文件名格式(file).X堆栈偏移量(10).X是否同时输出到终端(false).Error(ctx, 1, 2, 3)
		X文件路径(path).X文件名格式(file).X是否同时输出到终端(false).X输出并格式化ERR(ctx, "%d %d %d", 1, 2, 3)
		content := 文件类.X读文本(文件类.X路径生成(path, file))
		fmt.Println(content)
		t.Assert(文本类.X统计次数(content, defaultLevelPrefixes[LEVEL_ERRO]), 2)
		t.Assert(文本类.X统计次数(content, "1 2 3"), 2)
		// t.Assert(gstr.Count(content, "Stack"), 1) // 翻译：// t断言content中"Stack"出现的次数为1次
	})
}

func Test_Stack(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		path := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, 时间类.X取时间戳纳秒())

		err := 文件类.X创建目录(path)
		t.AssertNil(err)
		defer 文件类.X删除(path)

		X文件路径(path).X文件名格式(file).X堆栈选项(false).X是否同时输出到终端(false).Error(ctx, 1, 2, 3)
		X文件路径(path).X文件名格式(file).X是否同时输出到终端(false).X输出并格式化ERR(ctx, "%d %d %d", 1, 2, 3)
		content := 文件类.X读文本(文件类.X路径生成(path, file))
		fmt.Println(content)
		t.Assert(文本类.X统计次数(content, defaultLevelPrefixes[LEVEL_ERRO]), 2)
		t.Assert(文本类.X统计次数(content, "1 2 3"), 2)
		// t.Assert(gstr.Count(content, "Stack"), 1) // 翻译：// t断言content中"Stack"出现的次数为1次
	})
}

func Test_StackWithFilter(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		path := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, 时间类.X取时间戳纳秒())

		err := 文件类.X创建目录(path)
		t.AssertNil(err)
		defer 文件类.X删除(path)

		X文件路径(path).X文件名格式(file).X堆栈过滤("none").X是否同时输出到终端(false).Error(ctx, 1, 2, 3)
		content := 文件类.X读文本(文件类.X路径生成(path, file))
		fmt.Println(ctx, content)
		t.Assert(文本类.X统计次数(content, defaultLevelPrefixes[LEVEL_ERRO]), 1)
		t.Assert(文本类.X统计次数(content, "1 2 3"), 1)
		// t.Assert(gstr.Count(content, "Stack"), 1) // 翻译：// t断言content中"Stack"出现的次数为1次

	})
	单元测试类.C(t, func(t *单元测试类.T) {
		path := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, 时间类.X取时间戳纳秒())

		err := 文件类.X创建目录(path)
		t.AssertNil(err)
		defer 文件类.X删除(path)

		X文件路径(path).X文件名格式(file).X堆栈过滤("/gf/").X是否同时输出到终端(false).Error(ctx, 1, 2, 3)
		content := 文件类.X读文本(文件类.X路径生成(path, file))
		fmt.Println(ctx, content)
		t.Assert(文本类.X统计次数(content, defaultLevelPrefixes[LEVEL_ERRO]), 1)
		t.Assert(文本类.X统计次数(content, "1 2 3"), 1)
		// t.Assert(gstr.Count(content, "Stack"), 0) // 翻译：// t断言content中"Stack"出现的次数为0
	})
}

func Test_Header(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		path := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, 时间类.X取时间戳纳秒())

		err := 文件类.X创建目录(path)
		t.AssertNil(err)
		defer 文件类.X删除(path)

		X文件路径(path).X文件名格式(file).X是否输出头信息(true).X是否同时输出到终端(false).Error(ctx, 1, 2, 3)
		content := 文件类.X读文本(文件类.X路径生成(path, file))
		t.Assert(文本类.X统计次数(content, defaultLevelPrefixes[LEVEL_ERRO]), 1)
		t.Assert(文本类.X统计次数(content, "1 2 3"), 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		path := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, 时间类.X取时间戳纳秒())

		err := 文件类.X创建目录(path)
		t.AssertNil(err)
		defer 文件类.X删除(path)

		X文件路径(path).X文件名格式(file).X是否输出头信息(false).X是否同时输出到终端(false).Error(ctx, 1, 2, 3)
		content := 文件类.X读文本(文件类.X路径生成(path, file))
		t.Assert(文本类.X统计次数(content, defaultLevelPrefixes[LEVEL_ERRO]), 0)
		t.Assert(文本类.X统计次数(content, "1 2 3"), 1)
	})
}

func Test_Line(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		path := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, 时间类.X取时间戳纳秒())

		err := 文件类.X创建目录(path)
		t.AssertNil(err)
		defer 文件类.X删除(path)

		X文件路径(path).X文件名格式(file).X是否输出源文件路径与行号(true).X是否同时输出到终端(false).X输出DEBU(ctx, 1, 2, 3)
		content := 文件类.X读文本(文件类.X路径生成(path, file))
		fmt.Println(content)
		t.Assert(文本类.X统计次数(content, defaultLevelPrefixes[LEVEL_DEBU]), 1)
		t.Assert(文本类.X统计次数(content, "1 2 3"), 1)
// t.Assert(gstr.Count(content, ".go"), 1)
// 检查content字符串中".go"出现的次数是否为1，将结果断言为真
// t.Assert(gstr.Contains(content, gfile.Separator), true)
// 断言content字符串中是否包含系统文件分隔符（gfile.Separator），结果应为真
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		path := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, 时间类.X取时间戳纳秒())

		err := 文件类.X创建目录(path)
		t.AssertNil(err)
		defer 文件类.X删除(path)

		X文件路径(path).X文件名格式(file).X是否输出源文件路径与行号(false).X是否同时输出到终端(false).X输出DEBU(ctx, 1, 2, 3)
		content := 文件类.X读文本(文件类.X路径生成(path, file))
		t.Assert(文本类.X统计次数(content, defaultLevelPrefixes[LEVEL_DEBU]), 1)
		t.Assert(文本类.X统计次数(content, "1 2 3"), 1)
// t.Assert(gstr.Count(content, ".go"), 1) // 断言content字符串中".go"子串出现的次数为1次
// t.Assert(gstr.Contains(content, gfile.Separator), false) // 断言content字符串中不包含gfile.Separator定义的分隔符
	})
}

func Test_Async(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		path := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, 时间类.X取时间戳纳秒())

		err := 文件类.X创建目录(path)
		t.AssertNil(err)
		defer 文件类.X删除(path)

		X文件路径(path).X文件名格式(file).X是否异步输出().X是否同时输出到终端(false).X输出DEBU(ctx, 1, 2, 3)
		time.Sleep(1000 * time.Millisecond)

		content := 文件类.X读文本(文件类.X路径生成(path, file))
		t.Assert(文本类.X统计次数(content, defaultLevelPrefixes[LEVEL_DEBU]), 1)
		t.Assert(文本类.X统计次数(content, "1 2 3"), 1)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		path := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, 时间类.X取时间戳纳秒())

		err := 文件类.X创建目录(path)
		t.AssertNil(err)
		defer 文件类.X删除(path)

		X文件路径(path).X文件名格式(file).X是否异步输出(false).X是否同时输出到终端(false).X输出DEBU(ctx, 1, 2, 3)
		content := 文件类.X读文本(文件类.X路径生成(path, file))
		t.Assert(文本类.X统计次数(content, defaultLevelPrefixes[LEVEL_DEBU]), 1)
		t.Assert(文本类.X统计次数(content, "1 2 3"), 1)
	})
}
