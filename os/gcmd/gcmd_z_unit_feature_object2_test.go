// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。 md5:81db3d7bd1ed4da8

package cmd类_test

import (
	"context"
	"os"
	"testing"

	"github.com/888go/goframe/frame/g"
	gcmd "github.com/888go/goframe/os/gcmd"
	gctx "github.com/888go/goframe/os/gctx"
	gtest "github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gtag"
)

type commandBuild struct {
	g.Meta               `name:"build" root:"build" args:"true" brief:"{commandBuildBrief}" dc:"{commandBuildDc}" eg:"{commandBuildEg}" ad:"{commandBuildAd}"`
	nodeNameInConfigFile string // configFileNodeName 是配置文件中编译器配置的节点名称。 md5:3f5ccccd42d1daca
	packedGoFileName     string // packedGoFileName 用于指定将公共文件夹打包成单个 Go 文件时的文件名。 md5:ecbe7ef3574629cf
}

const (
	commandBuildBrief = `cross-building go project for lots of platforms`
	commandBuildEg    = `
gf build main.go
gf build main.go --pack public,template
gf build main.go --cgo
gf build main.go -m none 
gf build main.go -n my-app -a all -s all
gf build main.go -n my-app -a amd64,386 -s linux -p .
gf build main.go -n my-app -v 1.0 -a amd64,386 -s linux,windows,darwin -p ./docker/bin
`
	commandBuildDc = `
The "build" command is most commonly used command, which is designed as a powerful wrapper for 
"go build" command for convenience cross-compiling usage. 
It provides much more features for building binary:
1. Cross-Compiling for many platforms and architectures.
2. Configuration file support for compiling.
3. Build-In Variables.
`
	commandBuildAd = `
PLATFORMS
    darwin    amd64,arm64
    freebsd   386,amd64,arm
    linux     386,amd64,arm,arm64,ppc64,ppc64le,mips,mipsle,mips64,mips64le
    netbsd    386,amd64,arm
    openbsd   386,amd64,arm
    windows   386,amd64
`
			//golang.google.cn/doc/install/source. md5:f8fa931d443d3f23
	commandBuildPlatforms = `
    darwin    amd64
    darwin    arm64
    ios       amd64
    ios       arm64
    freebsd   386
    freebsd   amd64
    freebsd   arm
    linux     386
    linux     amd64
    linux     arm
    linux     arm64
    linux     ppc64
    linux     ppc64le
    linux     mips
    linux     mipsle
    linux     mips64
    linux     mips64le
    netbsd    386
    netbsd    amd64
    netbsd    arm
    openbsd   386
    openbsd   amd64
    openbsd   arm
    windows   386
    windows   amd64
	android   arm
	dragonfly amd64
	plan9     386
	plan9     amd64
	solaris   amd64
`
	commandBuildBriefPack = `
destination file path for packed file. if extension of the filename is ".go" and "-n" option is given, 
it enables packing SRC to go file, or else it packs SRC into a binary file.

`
	commandGenDaoBriefJsonCase = `
generated json tag case for model struct, cases are as follows:
| Case            | Example            |
|---------------- |--------------------|
| Camel           | AnyKindOfString    | 
| CamelLower      | anyKindOfString    | default
| Snake           | any_kind_of_string |
| SnakeScreaming  | ANY_KIND_OF_STRING |
| SnakeFirstUpper | rgb_code_md5       |
| Kebab           | any-kind-of-string |
| KebabScreaming  | ANY-KIND-OF-STRING |
`
)

func init() {
	gtag.Sets(map[string]string{
		`commandBuildBrief`:          commandBuildBrief,
		`commandBuildDc`:             commandBuildDc,
		`commandBuildEg`:             commandBuildEg,
		`commandBuildAd`:             commandBuildAd,
		`commandBuildBriefPack`:      commandBuildBriefPack,
		`commandGenDaoBriefJsonCase`: commandGenDaoBriefJsonCase,
	})
}

type commandBuildInput struct {
	g.Meta   `name:"build" config:"gfcli.build"`
	Name     string `short:"n" name:"name"     brief:"output binary name"`
	Version  string `short:"v" name:"version"  brief:"output binary version"`
	Arch     string `short:"a" name:"arch"     brief:"output binary architecture, multiple arch separated with ','"`
	System   string `short:"s" name:"system"   brief:"output binary system, multiple os separated with ','"`
	Output   string `short:"o" name:"output"   brief:"output binary path, used when building single binary file"`
	Path     string `short:"p" name:"path"     brief:"output binary directory path, default is './bin'" d:"./bin"`
	Extra    string `short:"e" name:"extra"    brief:"extra custom \"go build\" options"`
	Mod      string `short:"m" name:"mod"      brief:"like \"-mod\" option of \"go build\", use \"-m none\" to disable go module"`
	Cgo      bool   `short:"c" name:"cgo"      brief:"enable or disable cgo feature, it's disabled in default" orphan:"true"`
	JsonCase string `short:"j" name:"jsonCase" brief:"{commandGenDaoBriefJsonCase}" d:"CamelLower"`
	Pack     string `name:"pack" brief:"{commandBuildBriefPack}"`
}

type commandBuildOutput struct{}

func (c commandBuild) Index(ctx context.Context, in commandBuildInput) (out *commandBuildOutput, err error) {
	return
}

func TestNewFromObject(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx = gctx.X创建()
		)
		cmd, err := gcmd.NewFromObject(commandBuild{
			nodeNameInConfigFile: "gfcli.build",
			packedGoFileName:     "build_pack_data.go",
		})
		t.AssertNil(err)

		os.Args = []string{"build", "-h"}
		err = cmd.RunWithError(ctx)
		t.AssertNil(err)
	})
}
