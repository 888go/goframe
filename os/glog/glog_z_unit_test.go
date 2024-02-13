// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 日志类_test

import (
	"bytes"
	"context"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
)

func TestCase(t *testing.T) {
	defaultLog := 日志类.X取默认日志类().X取副本()
	defer 日志类.X设置默认日志类(defaultLog)

	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertNE(日志类.X取单例对象(), nil)
	})
}

func TestDefaultLogger(t *testing.T) {
	defaultLog := 日志类.X取默认日志类().X取副本()
	defer 日志类.X设置默认日志类(defaultLog)

	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertNE(defaultLog, nil)
		log := 日志类.X创建()
		日志类.X设置默认日志类(log)
		t.AssertEQ(日志类.X取默认日志类(), defaultLog)
		t.AssertEQ(日志类.Expose别名(), defaultLog)
	})
}

func TestAPI(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		日志类.X输出(ctx, "Print")
		日志类.X输出并格式化(ctx, "%s", "Printf")
		日志类.X输出INFO(ctx, "Info")
		日志类.X输出并格式化INFO(ctx, "%s", "Infof")
		日志类.X输出DEBU(ctx, "Debug")
		日志类.X输出并格式化DEBU(ctx, "%s", "Debugf")
		日志类.X输出NOTI(ctx, "Notice")
		日志类.X输出并格式化NOTI(ctx, "%s", "Noticef")
		日志类.X输出WARN(ctx, "Warning")
		日志类.X输出并格式化WARN(ctx, "%s", "Warningf")
		日志类.Error(ctx, "Error")
		日志类.X输出并格式化ERR(ctx, "%s", "Errorf")
		日志类.X输出CRIT(ctx, "Critical")
		日志类.X输出并格式化CRIT(ctx, "%s", "Criticalf")
	})
}

func TestChaining(t *testing.T) {
	defaultLog := 日志类.X取默认日志类().X取副本()
	defer 日志类.X设置默认日志类(defaultLog)

	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertNE(日志类.X文件分类("module"), nil)
		t.AssertNE(日志类.X文件名格式("test.log"), nil)
		t.AssertNE(日志类.X级别(日志类.LEVEL_ALL), nil)
		t.AssertNE(日志类.X文本级别("all"), nil)
		t.AssertNE(日志类.X堆栈偏移量(1), nil)
		t.AssertNE(日志类.X堆栈选项(false), nil)
		t.AssertNE(日志类.X堆栈过滤("none"), nil)
		t.AssertNE(日志类.X是否同时输出到终端(false), nil)
		t.AssertNE(日志类.X是否输出头信息(false), nil)
		t.AssertNE(日志类.X是否输出源文件路径与行号(false), nil)
		t.AssertNE(日志类.X是否异步输出(false), nil)
	})
}

func Test_SetFile(t *testing.T) {
	defaultLog := 日志类.X取默认日志类().X取副本()
	defer 日志类.X设置默认日志类(defaultLog)
	单元测试类.C(t, func(t *单元测试类.T) {
		日志类.X设置文件名格式("test.log")
	})
}

func Test_SetTimeFormat(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		w := bytes.NewBuffer(nil)
		l := 日志类.X创建并按writer(w)

		l.X设置时间格式("2006-01-02T15:04:05.000Z07:00")
		l.X输出DEBU(ctx, "test")

		t.AssertGE(len(strings.Split(w.String(), "[DEBU]")), 1)
		datetime := strings.Trim(strings.Split(w.String(), "[DEBU]")[0], " ")

		_, err := time.Parse("2006-01-02T15:04:05.000Z07:00", datetime)
		t.AssertNil(err)
		_, err = time.Parse("2006-01-02 15:04:05.000", datetime)
		t.AssertNE(err, nil)
		_, err = time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", datetime)
		t.AssertNE(err, nil)
	})
}

func Test_SetLevel(t *testing.T) {
	defaultLog := 日志类.X取默认日志类().X取副本()
	defer 日志类.X设置默认日志类(defaultLog)
	单元测试类.C(t, func(t *单元测试类.T) {
		日志类.X设置级别(日志类.LEVEL_ALL)
		t.Assert(日志类.X取级别()&日志类.LEVEL_ALL, 日志类.LEVEL_ALL)
	})
}

func Test_SetAsync(t *testing.T) {
	defaultLog := 日志类.X取默认日志类().X取副本()
	defer 日志类.X设置默认日志类(defaultLog)
	单元测试类.C(t, func(t *单元测试类.T) {
		日志类.X设置异步输出(false)
	})
}

func Test_SetStdoutPrint(t *testing.T) {
	defaultLog := 日志类.X取默认日志类().X取副本()
	defer 日志类.X设置默认日志类(defaultLog)
	单元测试类.C(t, func(t *单元测试类.T) {
		日志类.X设置是否同时输出到终端(false)
	})
}

func Test_SetHeaderPrint(t *testing.T) {
	defaultLog := 日志类.X取默认日志类().X取副本()
	defer 日志类.X设置默认日志类(defaultLog)
	单元测试类.C(t, func(t *单元测试类.T) {
		日志类.X设置是否输出头信息(false)
	})
}

func Test_SetPrefix(t *testing.T) {
	defaultLog := 日志类.X取默认日志类().X取副本()
	defer 日志类.X设置默认日志类(defaultLog)
	单元测试类.C(t, func(t *单元测试类.T) {
		日志类.X设置前缀("log_prefix")
	})
}

func Test_SetConfigWithMap(t *testing.T) {
	defaultLog := 日志类.X取默认日志类().X取副本()
	defer 日志类.X设置默认日志类(defaultLog)
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(日志类.X设置配置Map(map[string]interface{}{
			"level": "all",
		}), nil)
	})
}

func Test_SetPath(t *testing.T) {
	defaultLog := 日志类.X取默认日志类().X取副本()
	defer 日志类.X设置默认日志类(defaultLog)
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(日志类.X设置文件路径("/var/log"), nil)
		t.Assert(日志类.X取文件路径(), "/var/log")
	})
}

func Test_SetWriter(t *testing.T) {
	defaultLog := 日志类.X取默认日志类().X取副本()
	defer 日志类.X设置默认日志类(defaultLog)
	单元测试类.C(t, func(t *单元测试类.T) {
		日志类.X设置Writer(os.Stdout)
		t.Assert(日志类.X取Writer(), os.Stdout)
	})
}

func Test_SetFlags(t *testing.T) {
	defaultLog := 日志类.X取默认日志类().X取副本()
	defer 日志类.X设置默认日志类(defaultLog)
	单元测试类.C(t, func(t *单元测试类.T) {
		日志类.X设置额外标识(日志类.F_ASYNC)
		t.Assert(日志类.X取标识(), 日志类.F_ASYNC)
	})
}

func Test_SetCtxKeys(t *testing.T) {
	defaultLog := 日志类.X取默认日志类().X取副本()
	defer 日志类.X设置默认日志类(defaultLog)
	单元测试类.C(t, func(t *单元测试类.T) {
		日志类.X设置上下文名称("SpanId", "TraceId")
		t.Assert(日志类.X取上下文名称(), []string{"SpanId", "TraceId"})
	})
}

func Test_PrintStack(t *testing.T) {
	defaultLog := 日志类.X取默认日志类().X取副本()
	defer 日志类.X设置默认日志类(defaultLog)
	单元测试类.C(t, func(t *单元测试类.T) {
		日志类.X输出堆栈信息(ctx, 1)
	})
}

func Test_SetStack(t *testing.T) {
	defaultLog := 日志类.X取默认日志类().X取副本()
	defer 日志类.X设置默认日志类(defaultLog)
	单元测试类.C(t, func(t *单元测试类.T) {
		日志类.X设置堆栈跟踪(true)
		t.Assert(日志类.X取堆栈信息(1), "")
	})
}

func Test_SetLevelStr(t *testing.T) {
	defaultLog := 日志类.X取默认日志类().X取副本()
	defer 日志类.X设置默认日志类(defaultLog)
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(日志类.X设置文本级别("all"), nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		l := 日志类.X创建()
		t.AssertNE(l.X设置文本级别("test"), nil)
	})
}

func Test_SetLevelPrefix(t *testing.T) {
	defaultLog := 日志类.X取默认日志类().X取副本()
	defer 日志类.X设置默认日志类(defaultLog)
	单元测试类.C(t, func(t *单元测试类.T) {
		日志类.X设置级别前缀(日志类.LEVEL_ALL, "LevelPrefix")
		t.Assert(日志类.X取级别前缀(日志类.LEVEL_ALL), "LevelPrefix")
	})
}

func Test_SetLevelPrefixes(t *testing.T) {
	defaultLog := 日志类.X取默认日志类().X取副本()
	defer 日志类.X设置默认日志类(defaultLog)
	单元测试类.C(t, func(t *单元测试类.T) {
		日志类.X设置级别前缀Map(map[int]string{
			日志类.LEVEL_ALL: "ALL_Prefix",
		})
	})
}

func Test_SetHandlers(t *testing.T) {
	defaultLog := 日志类.X取默认日志类().X取副本()
	defer 日志类.X设置默认日志类(defaultLog)
	单元测试类.C(t, func(t *单元测试类.T) {
		日志类.X设置中间件(func(ctx context.Context, in *日志类.HandlerInput) {
		})
	})
}

func Test_SetWriterColorEnable(t *testing.T) {
	defaultLog := 日志类.X取默认日志类().X取副本()
	defer 日志类.X设置默认日志类(defaultLog)
	单元测试类.C(t, func(t *单元测试类.T) {
		日志类.X设置文件是否输出颜色(true)
	})
}

func Test_Instance(t *testing.T) {
	defaultLog := 日志类.X取默认日志类().X取副本()
	defer 日志类.X设置默认日志类(defaultLog)
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertNE(日志类.X取单例对象("gf"), nil)
	})
}

func Test_GetConfig(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		config := 日志类.X取默认日志类().X取配置项()
		t.Assert(config.Path, "")
		t.Assert(config.StdoutPrint, true)
	})
}

func Test_Write(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		l := 日志类.X创建()
		len, err := l.Write([]byte("GoFrame"))
		t.AssertNil(err)
		t.Assert(len, 7)
	})
}

func Test_Chaining_To(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		l := 日志类.X取默认日志类().X取副本()
		logTo := l.X重定向输出(os.Stdout)
		t.AssertNE(logTo, nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		l := 日志类.X创建()
		logTo := l.X重定向输出(os.Stdout)
		t.AssertNE(logTo, nil)
	})
}

func Test_Chaining_Path(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		l := 日志类.X取默认日志类().X取副本()
		logPath := l.X文件路径("./")
		t.AssertNE(logPath, nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		l := 日志类.X创建()
		logPath := l.X文件路径("./")
		t.AssertNE(logPath, nil)
	})
}

func Test_Chaining_Cat(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		l := 日志类.X创建()
		logCat := l.X文件分类(".gf")
		t.AssertNE(logCat, nil)
	})
}

func Test_Chaining_Level(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		l := 日志类.X创建()
		logLevel := l.X级别(日志类.LEVEL_ALL)
		t.AssertNE(logLevel, nil)
	})
}

func Test_Chaining_LevelStr(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		l := 日志类.X创建()
		logLevelStr := l.X文本级别("all")
		t.AssertNE(logLevelStr, nil)
	})
}

func Test_Chaining_Skip(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		l := 日志类.X创建()
		logSkip := l.X堆栈偏移量(1)
		t.AssertNE(logSkip, nil)
	})
}

func Test_Chaining_Stack(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		l := 日志类.X创建()
		logStack := l.X堆栈选项(true)
		t.AssertNE(logStack, nil)
	})
}

func Test_Chaining_StackWithFilter(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		l := 日志类.X创建()
		logStackWithFilter := l.X堆栈过滤("gtest")
		t.AssertNE(logStackWithFilter, nil)
	})
}

func Test_Chaining_Stdout(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		l := 日志类.X创建()
		logStdout := l.X是否同时输出到终端(true)
		t.AssertNE(logStdout, nil)
	})
}

func Test_Chaining_Header(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		l := 日志类.X创建()
		logHeader := l.X是否输出头信息(true)
		t.AssertNE(logHeader, nil)
	})
}

func Test_Chaining_Line(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		l := 日志类.X创建()
		logLine := l.X是否输出源文件路径与行号(true)
		t.AssertNE(logLine, nil)
	})
}

func Test_Chaining_Async(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		l := 日志类.X创建()
		logAsync := l.X是否异步输出(true)
		t.AssertNE(logAsync, nil)
	})
}

func Test_Config_SetDebug(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		l := 日志类.X创建()
		l.X设置debug(false)
	})
}

func Test_Config_AppendCtxKeys(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		l := 日志类.X创建()
		l.AppendCtxKeys("Trace-Id", "Span-Id", "Test")
		l.AppendCtxKeys("Trace-Id-New", "Span-Id-New", "Test")
	})
}

func Test_Config_SetPath(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		l := 日志类.X创建()
		t.AssertNE(l.X设置文件路径(""), nil)
	})
}

func Test_Config_SetStdoutColorDisabled(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		l := 日志类.X创建()
		l.X设置关闭终端颜色输出(false)
	})
}

func Test_Ctx(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		w := bytes.NewBuffer(nil)
		l := 日志类.X创建并按writer(w)
		l.X设置上下文名称("Trace-Id", "Span-Id", "Test")
		ctx := context.WithValue(context.Background(), "Trace-Id", "1234567890")
		ctx = context.WithValue(ctx, "Span-Id", "abcdefg")

		l.X输出(ctx, 1, 2, 3)
		t.Assert(文本类.X统计次数(w.String(), "1234567890"), 1)
		t.Assert(文本类.X统计次数(w.String(), "abcdefg"), 1)
		t.Assert(文本类.X统计次数(w.String(), "1 2 3"), 1)
	})
}

func Test_Ctx_Config(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		w := bytes.NewBuffer(nil)
		l := 日志类.X创建并按writer(w)
		m := map[string]interface{}{
			"CtxKeys": g.SliceStr别名{"Trace-Id", "Span-Id", "Test"},
		}
		var nilMap map[string]interface{}

		err := l.X设置配置Map(m)
		t.AssertNil(err)
		err = l.X设置配置Map(nilMap)
		t.AssertNE(err, nil)

		ctx := context.WithValue(context.Background(), "Trace-Id", "1234567890")
		ctx = context.WithValue(ctx, "Span-Id", "abcdefg")

		l.X输出(ctx, 1, 2, 3)
		t.Assert(文本类.X统计次数(w.String(), "1234567890"), 1)
		t.Assert(文本类.X统计次数(w.String(), "abcdefg"), 1)
		t.Assert(文本类.X统计次数(w.String(), "1 2 3"), 1)
	})
}

func Test_Concurrent(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		c := 1000
		l := 日志类.X创建()
		s := "@1234567890#"
		f := "test.log"
		p := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		t.Assert(l.X设置文件路径(p), nil)
		defer 文件类.X删除(p)
		wg := sync.WaitGroup{}
		ch := make(chan struct{})
		for i := 0; i < c; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				<-ch
				l.X文件名格式(f).X是否同时输出到终端(false).X输出(ctx, s)
			}()
		}
		close(ch)
		wg.Wait()
		content := 文件类.X读文本(文件类.X路径生成(p, f))
		t.Assert(文本类.X统计次数(content, s), c)
	})
}
