@Echo Off
color 2e
Title 命令
:begin

cls
Echo. 请选择需要的操作
Echo     0 新建项目 go mod init *项目名称*
Echo     1 安装第三方包 go mod download
Echo     2 复制第三方包到目录 go mod vendor
echo     3 更新第三方包 go mod tidy
echo     4 更新最新第三方包 go get -u *包地址*
echo     5 更新全部最新第三方包 go get -u -all
echo     6 无黑框编译 go build -ldflags -H=windowsgui 
echo     7 常规编译 go build 
echo     8 七牛代理 go env -w GOPROXY=*代理地址*
echo     9 设置不走proxy的私有仓库 go env -w GOPRIVATE=*包地址*
echo     10 自建代理
echo     11 查看go变量 go env 
echo     12 新建工作区或加入工作区
echo     13 切换到x86
echo     14 切换到x64
Set /P Choice= 　　　　　请选择要进行的操作数字 ，然后按回车：


If not "%Choice%"=="" (
  If "%Choice%"=="0" goto 新建项目
  If "%Choice%"=="1" goto 安装第三方包
  If "%Choice%"=="2" goto 复制第三方包到目录
  If "%Choice%"=="3" goto 更新第三方包
  If "%Choice%"=="4" goto 更新最新第三方包
  If "%Choice%"=="5" goto 更新全部最新第三方包
  If "%Choice%"=="6" goto 无黑框编译
  If "%Choice%"=="7" goto 常规编译
  If "%Choice%"=="8" goto 七牛代理
  If "%Choice%"=="9" goto 设置不走proxy的私有仓库
  If "%Choice%"=="10" goto 自建代理
  If "%Choice%"=="11" goto 查看go变量
  If "%Choice%"=="12" goto 新建工作区或加入工作区
  If "%Choice%"=="13" goto 切换到x86
  If "%Choice%"=="14" goto 切换到x64
)

:新建项目
@Echo on
set input=
set /p input=请输入项目名称:
go mod init %input%
pause
@Echo Off
goto :begin

:安装第三方包
@Echo on
go mod download
pause
@Echo Off
goto :begin

:复制第三方包到目录
@Echo on
go mod vendor
pause
@Echo Off
goto :begin

:更新第三方包
@Echo on
go mod tidy
pause
@Echo Off
goto :begin

:更新最新第三方包
@Echo on
set input=
set /p input=输入第三方包地址:
go get -u %input%
pause
@Echo Off
goto :begin

:更新全部最新第三方包
@Echo on
go get -u -all
pause
@Echo Off
goto :begin

:无黑框编译
@Echo on
go build -ldflags -H=windowsgui 
pause
@Echo Off
goto :begin

:常规编译
@Echo on
go build 
pause
@Echo Off
goto :begin

:七牛代理
@Echo on
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
pause
@Echo Off
goto :begin

:自建代理
@Echo on
set input=
set /p input=请输入ip地址, 如127.0.0.1:
go env -w GO111MODULE=on
go env -w GOPROXY=http://%input%,direct
pause
@Echo Off
goto :begin


:设置不走proxy的私有仓库
@Echo on
set input=
set /p input=请输入不走proxy的私有仓库, 如 git.mycompany.com,github.com/my/private
go env -w GO111MODULE=on
go env -w GOPRIVATE="%input%"
pause
@Echo Off
goto :begin


:查看go变量
@Echo on
go env 
pause
@Echo Off
goto :begin

:新建工作区或加入工作区
@Echo on
set input=
set /p input=请输入需要加入新工作区的目录名称, 如demo:
@Echo Off

if not exist "%cd%\go.work" (
echo go.work不存在！新建工作区后加入%input%
go work init %input%
)else (
echo go.work存在！加入工作区%input%
go work use %input%
)

pause
goto :begin


pause>nul
goto :begin


:切换到x86
@Echo on
go env -w GOARCH=386
pause
@Echo Off
goto :begin

:切换到x64
@Echo on
go env -w GOARCH=amd64
pause
@Echo Off
goto :begin

