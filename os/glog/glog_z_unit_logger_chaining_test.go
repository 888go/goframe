// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 日志类

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	gfile "github.com/888go/goframe/os/gfile"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
)

func Test_To(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		w := bytes.NewBuffer(nil)
		X重定向输出(w).Error(ctx, 1, 2, 3)
		X重定向输出(w).X输出并格式化ERR(ctx, "%d %d %d", 1, 2, 3)
		t.Assert(gstr.X统计次数(w.String(), defaultLevelPrefixes[LEVEL_ERRO]), 2)
		t.Assert(gstr.X统计次数(w.String(), "1 2 3"), 2)
	})
}

func Test_Path(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, gtime.X取时间戳纳秒())

		err := gfile.X创建目录(path)
		t.AssertNil(err)
		defer gfile.X删除(path)

		X文件路径(path).X文件名格式(file).X是否同时输出到终端(false).Error(ctx, 1, 2, 3)
		X文件路径(path).X文件名格式(file).X是否同时输出到终端(false).X输出并格式化ERR(ctx, "%d %d %d", 1, 2, 3)
		content := gfile.X读文本(gfile.X路径生成(path, file))
		t.Assert(gstr.X统计次数(content, defaultLevelPrefixes[LEVEL_ERRO]), 2)
		t.Assert(gstr.X统计次数(content, "1 2 3"), 2)
	})
}

func Test_Cat(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cat := "category"
		path := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, gtime.X取时间戳纳秒())

		err := gfile.X创建目录(path)
		t.AssertNil(err)
		defer gfile.X删除(path)

		X文件路径(path).X文件名格式(file).X文件分类(cat).X是否同时输出到终端(false).Error(ctx, 1, 2, 3)
		X文件路径(path).X文件名格式(file).X文件分类(cat).X是否同时输出到终端(false).X输出并格式化ERR(ctx, "%d %d %d", 1, 2, 3)
		content := gfile.X读文本(gfile.X路径生成(path, cat, file))
		t.Assert(gstr.X统计次数(content, defaultLevelPrefixes[LEVEL_ERRO]), 2)
		t.Assert(gstr.X统计次数(content, "1 2 3"), 2)
	})
}

func Test_Level(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, gtime.X取时间戳纳秒())

		err := gfile.X创建目录(path)
		t.AssertNil(err)
		defer gfile.X删除(path)

		X文件路径(path).X文件名格式(file).X级别(LEVEL_PROD).X是否同时输出到终端(false).X输出DEBU(ctx, 1, 2, 3)
		X文件路径(path).X文件名格式(file).X级别(LEVEL_PROD).X是否同时输出到终端(false).X输出DEBU(ctx, "%d %d %d", 1, 2, 3)
		content := gfile.X读文本(gfile.X路径生成(path, file))
		t.Assert(gstr.X统计次数(content, defaultLevelPrefixes[LEVEL_DEBU]), 0)
		t.Assert(gstr.X统计次数(content, "1 2 3"), 0)
	})
}

func Test_Skip(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, gtime.X取时间戳纳秒())

		err := gfile.X创建目录(path)
		t.AssertNil(err)
		defer gfile.X删除(path)

		X文件路径(path).X文件名格式(file).X堆栈偏移量(10).X是否同时输出到终端(false).Error(ctx, 1, 2, 3)
		X文件路径(path).X文件名格式(file).X是否同时输出到终端(false).X输出并格式化ERR(ctx, "%d %d %d", 1, 2, 3)
		content := gfile.X读文本(gfile.X路径生成(path, file))
		fmt.Println(content)
		t.Assert(gstr.X统计次数(content, defaultLevelPrefixes[LEVEL_ERRO]), 2)
		t.Assert(gstr.X统计次数(content, "1 2 3"), 2)
				// 断言内容(content)中"Stack"出现的次数为1次。 md5:a246e53abdc8cb5e
	})
}

func Test_Stack(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, gtime.X取时间戳纳秒())

		err := gfile.X创建目录(path)
		t.AssertNil(err)
		defer gfile.X删除(path)

		X文件路径(path).X文件名格式(file).X堆栈选项(false).X是否同时输出到终端(false).Error(ctx, 1, 2, 3)
		X文件路径(path).X文件名格式(file).X是否同时输出到终端(false).X输出并格式化ERR(ctx, "%d %d %d", 1, 2, 3)
		content := gfile.X读文本(gfile.X路径生成(path, file))
		fmt.Println(content)
		t.Assert(gstr.X统计次数(content, defaultLevelPrefixes[LEVEL_ERRO]), 2)
		t.Assert(gstr.X统计次数(content, "1 2 3"), 2)
				// 断言内容(content)中"Stack"出现的次数为1次。 md5:a246e53abdc8cb5e
	})
}

func Test_StackWithFilter(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, gtime.X取时间戳纳秒())

		err := gfile.X创建目录(path)
		t.AssertNil(err)
		defer gfile.X删除(path)

		X文件路径(path).X文件名格式(file).X堆栈过滤("none").X是否同时输出到终端(false).Error(ctx, 1, 2, 3)
		content := gfile.X读文本(gfile.X路径生成(path, file))
		fmt.Println(ctx, content)
		t.Assert(gstr.X统计次数(content, defaultLevelPrefixes[LEVEL_ERRO]), 1)
		t.Assert(gstr.X统计次数(content, "1 2 3"), 1)
				// 断言内容(content)中"Stack"出现的次数为1次。 md5:a246e53abdc8cb5e

	})
	gtest.C(t, func(t *gtest.T) {
		path := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, gtime.X取时间戳纳秒())

		err := gfile.X创建目录(path)
		t.AssertNil(err)
		defer gfile.X删除(path)

		X文件路径(path).X文件名格式(file).X堆栈过滤("/gf/").X是否同时输出到终端(false).Error(ctx, 1, 2, 3)
		content := gfile.X读文本(gfile.X路径生成(path, file))
		fmt.Println(ctx, content)
		t.Assert(gstr.X统计次数(content, defaultLevelPrefixes[LEVEL_ERRO]), 1)
		t.Assert(gstr.X统计次数(content, "1 2 3"), 1)
				// 使用t.Assert断言gstr在content中"Stack"的计数为0。 md5:b6a4aff04f1a4b28
	})
}

func Test_Header(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, gtime.X取时间戳纳秒())

		err := gfile.X创建目录(path)
		t.AssertNil(err)
		defer gfile.X删除(path)

		X文件路径(path).X文件名格式(file).X是否输出头信息(true).X是否同时输出到终端(false).Error(ctx, 1, 2, 3)
		content := gfile.X读文本(gfile.X路径生成(path, file))
		t.Assert(gstr.X统计次数(content, defaultLevelPrefixes[LEVEL_ERRO]), 1)
		t.Assert(gstr.X统计次数(content, "1 2 3"), 1)
	})
	gtest.C(t, func(t *gtest.T) {
		path := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, gtime.X取时间戳纳秒())

		err := gfile.X创建目录(path)
		t.AssertNil(err)
		defer gfile.X删除(path)

		X文件路径(path).X文件名格式(file).X是否输出头信息(false).X是否同时输出到终端(false).Error(ctx, 1, 2, 3)
		content := gfile.X读文本(gfile.X路径生成(path, file))
		t.Assert(gstr.X统计次数(content, defaultLevelPrefixes[LEVEL_ERRO]), 0)
		t.Assert(gstr.X统计次数(content, "1 2 3"), 1)
	})
}

func Test_Line(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, gtime.X取时间戳纳秒())

		err := gfile.X创建目录(path)
		t.AssertNil(err)
		defer gfile.X删除(path)

		X文件路径(path).X文件名格式(file).X是否输出源文件路径与行号(true).X是否同时输出到终端(false).X输出DEBU(ctx, 1, 2, 3)
		content := gfile.X读文本(gfile.X路径生成(path, file))
		fmt.Println(content)
		t.Assert(gstr.X统计次数(content, defaultLevelPrefixes[LEVEL_DEBU]), 1)
		t.Assert(gstr.X统计次数(content, "1 2 3"), 1)
		// t.Assert断言字符串gstr在content中出现的".go"子串次数为1
		// t.Assert断言content中包含gfile.Separator（假设是一个路径分隔符），结果为true
		// md5:1411e0e8f0387662
	})
	gtest.C(t, func(t *gtest.T) {
		path := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, gtime.X取时间戳纳秒())

		err := gfile.X创建目录(path)
		t.AssertNil(err)
		defer gfile.X删除(path)

		X文件路径(path).X文件名格式(file).X是否输出源文件路径与行号(false).X是否同时输出到终端(false).X输出DEBU(ctx, 1, 2, 3)
		content := gfile.X读文本(gfile.X路径生成(path, file))
		t.Assert(gstr.X统计次数(content, defaultLevelPrefixes[LEVEL_DEBU]), 1)
		t.Assert(gstr.X统计次数(content, "1 2 3"), 1)
		// 断言 content 中 ".go" 出现的次数为 1
		// 断言 content 不包含路径分隔符
		// md5:c3f84d90ca75dcce
	})
}

func Test_Async(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, gtime.X取时间戳纳秒())

		err := gfile.X创建目录(path)
		t.AssertNil(err)
		defer gfile.X删除(path)

		X文件路径(path).X文件名格式(file).X是否异步输出().X是否同时输出到终端(false).X输出DEBU(ctx, 1, 2, 3)
		time.Sleep(1000 * time.Millisecond)

		content := gfile.X读文本(gfile.X路径生成(path, file))
		t.Assert(gstr.X统计次数(content, defaultLevelPrefixes[LEVEL_DEBU]), 1)
		t.Assert(gstr.X统计次数(content, "1 2 3"), 1)
	})

	gtest.C(t, func(t *gtest.T) {
		path := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		file := fmt.Sprintf(`%d.log`, gtime.X取时间戳纳秒())

		err := gfile.X创建目录(path)
		t.AssertNil(err)
		defer gfile.X删除(path)

		X文件路径(path).X文件名格式(file).X是否异步输出(false).X是否同时输出到终端(false).X输出DEBU(ctx, 1, 2, 3)
		content := gfile.X读文本(gfile.X路径生成(path, file))
		t.Assert(gstr.X统计次数(content, defaultLevelPrefixes[LEVEL_DEBU]), 1)
		t.Assert(gstr.X统计次数(content, "1 2 3"), 1)
	})
}
