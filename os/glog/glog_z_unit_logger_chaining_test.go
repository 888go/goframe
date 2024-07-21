// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package glog

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/text/gstr"
)

func Test_To(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		w := bytes.NewBuffer(nil)
		To(w).Error(ctx, 1, 2, 3)
		To(w).Errorf(ctx, "%d %d %d", 1, 2, 3)
		t.Assert(gstr.Count(w.String(), defaultLevelPrefixes[LEVEL_ERRO]), 2)
		t.Assert(gstr.Count(w.String(), "1 2 3"), 2)
	})
}

func Test_Path(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).Stdout(false).Error(ctx, 1, 2, 3)
		Path(path).File(file).Stdout(false).Errorf(ctx, "%d %d %d", 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, file))
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_ERRO]), 2)
		t.Assert(gstr.Count(content, "1 2 3"), 2)
	})
}

func Test_Cat(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		cat := "category"
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).Cat(cat).Stdout(false).Error(ctx, 1, 2, 3)
		Path(path).File(file).Cat(cat).Stdout(false).Errorf(ctx, "%d %d %d", 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, cat, file))
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_ERRO]), 2)
		t.Assert(gstr.Count(content, "1 2 3"), 2)
	})
}

func Test_Level(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).Level(LEVEL_PROD).Stdout(false).Debug(ctx, 1, 2, 3)
		Path(path).File(file).Level(LEVEL_PROD).Stdout(false).Debug(ctx, "%d %d %d", 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, file))
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_DEBU]), 0)
		t.Assert(gstr.Count(content, "1 2 3"), 0)
	})
}

func Test_Skip(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).Skip(10).Stdout(false).Error(ctx, 1, 2, 3)
		Path(path).File(file).Stdout(false).Errorf(ctx, "%d %d %d", 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, file))
		fmt.Println(content)
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_ERRO]), 2)
		t.Assert(gstr.Count(content, "1 2 3"), 2)
		// 断言内容(content)中"Stack"出现的次数为1次。 md5:a246e53abdc8cb5e
	})
}

func Test_Stack(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).Stack(false).Stdout(false).Error(ctx, 1, 2, 3)
		Path(path).File(file).Stdout(false).Errorf(ctx, "%d %d %d", 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, file))
		fmt.Println(content)
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_ERRO]), 2)
		t.Assert(gstr.Count(content, "1 2 3"), 2)
		// 断言内容(content)中"Stack"出现的次数为1次。 md5:a246e53abdc8cb5e
	})
}

func Test_StackWithFilter(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).StackWithFilter("none").Stdout(false).Error(ctx, 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, file))
		fmt.Println(ctx, content)
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_ERRO]), 1)
		t.Assert(gstr.Count(content, "1 2 3"), 1)
		// 断言内容(content)中"Stack"出现的次数为1次。 md5:a246e53abdc8cb5e

	})
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).StackWithFilter("/gf/").Stdout(false).Error(ctx, 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, file))
		fmt.Println(ctx, content)
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_ERRO]), 1)
		t.Assert(gstr.Count(content, "1 2 3"), 1)
		// 使用t.Assert断言gstr在content中"Stack"的计数为0。 md5:b6a4aff04f1a4b28
	})
}

func Test_Header(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).Header(true).Stdout(false).Error(ctx, 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, file))
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_ERRO]), 1)
		t.Assert(gstr.Count(content, "1 2 3"), 1)
	})
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).Header(false).Stdout(false).Error(ctx, 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, file))
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_ERRO]), 0)
		t.Assert(gstr.Count(content, "1 2 3"), 1)
	})
}

func Test_Line(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).Line(true).Stdout(false).Debug(ctx, 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, file))
		fmt.Println(content)
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_DEBU]), 1)
		t.Assert(gstr.Count(content, "1 2 3"), 1)
		// t.Assert断言字符串gstr在content中出现的".go"子串次数为1
		// t.Assert断言content中包含gfile.Separator（假设是一个路径分隔符），结果为true
		// md5:1411e0e8f0387662
	})
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).Line(false).Stdout(false).Debug(ctx, 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, file))
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_DEBU]), 1)
		t.Assert(gstr.Count(content, "1 2 3"), 1)
		// 断言 content 中 ".go" 出现的次数为 1
		// 断言 content 不包含路径分隔符
		// md5:c3f84d90ca75dcce
	})
}

func Test_Async(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).Async().Stdout(false).Debug(ctx, 1, 2, 3)
		time.Sleep(1000 * time.Millisecond)

		content := gfile.GetContents(gfile.Join(path, file))
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_DEBU]), 1)
		t.Assert(gstr.Count(content, "1 2 3"), 1)
	})

	gtest.C(t, func(t *gtest.T) {
		path := gfile.Temp(gtime.TimestampNanoStr())
		file := fmt.Sprintf(`%d.log`, gtime.TimestampNano())

		err := gfile.Mkdir(path)
		t.AssertNil(err)
		defer gfile.Remove(path)

		Path(path).File(file).Async(false).Stdout(false).Debug(ctx, 1, 2, 3)
		content := gfile.GetContents(gfile.Join(path, file))
		t.Assert(gstr.Count(content, defaultLevelPrefixes[LEVEL_DEBU]), 1)
		t.Assert(gstr.Count(content, "1 2 3"), 1)
	})
}
