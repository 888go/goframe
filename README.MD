# # GoFrame

<div align=center>
<img src="https://goframe.org/statics/image/logo2.png?v=1" width="300"/>

[![Go 参考手册](https://pkg.go.dev/badge/github.com/gogf/gf/v2.svg)](https://pkg.go.dev/github.com/gogf/gf/v2)
[![GoFrame CI](https://github.com/gogf/gf/actions/workflows/ci-main.yml/badge.svg)](https://github.com/gogf/gf/actions/workflows/ci-main.yml)
[![Go 项目质量报告](https://goreportcard.com/badge/github.com/gogf/gf/v2)](https://goreportcard.com/report/github.com/gogf/gf/v2)
[![代码覆盖率](https://codecov.io/gh/gogf/gf/branch/master/graph/badge.svg)](https://codecov.io/gh/gogf/gf)
[![生产环境就绪](https://img.shields.io/badge/生产环境就绪-蓝色.svg?style=flat)](https://github.com/gogf/gf)
[![许可证](https://img.shields.io/github/license/gogf/gf.svg?style=flat)](https://github.com/gogf/gf)

[![发布版本](https://img.shields.io/github/v/release/gogf/gf?style=flat)](https://github.com/gogf/gf/releases)
[![GitHub 提交的拉取请求](https://img.shields.io/github/issues-pr/gogf/gf?style=flat)](https://github.com/gogf/gf/pulls)
[![GitHub 已关闭的拉取请求](https://img.shields.io/github/issues-pr-closed/gogf/gf?style=flat)](https://github.com/gogf/gf/pulls?q=is%3Apr+is%3Aclosed)
[![GitHub 问题](https://img.shields.io/github/issues/gogf/gf?style=flat)](https://github.com/gogf/gf/issues)
[![GitHub 已关闭的问题](https://img.shields.io/github/issues-closed/gogf/gf?style=flat)](https://github.com/gogf/gf/issues?q=is%3Aissue+is%3Aclosed)
![Stars](https://img.shields.io/github/stars/gogf/gf?style=flat)
![Forks](https://img.shields.io/github/forks/gogf/gf?style=flat)

</div>

`GoFrame` 是一个模块化、强大、高性能且适用于企业级应用开发的 Golang 框架。

# GoFrame

<div align=center>
<img src="https://goframe.org/statics/image/logo2.png?v=1" width="300"/>

[![Go 参考手册](https://pkg.go.dev/badge/github.com/gogf/gf/v2.svg)](https://pkg.go.dev/github.com/gogf/gf/v2)
[![GoFrame CI](https://github.com/gogf/gf/actions/workflows/ci-main.yml/badge.svg)](https://github.com/gogf/gf/actions/workflows/ci-main.yml)
[![Go 项目质量报告](https://goreportcard.com/badge/github.com/gogf/gf/v2)](https://goreportcard.com/report/github.com/gogf/gf/v2)
[![代码覆盖率](https://codecov.io/gh/gogf/gf/branch/master/graph/badge.svg)](https://codecov.io/gh/gogf/gf)
[![生产环境就绪](https://img.shields.io/badge/生产环境就绪-蓝色.svg?style=flat)](https://github.com/gogf/gf)
[![许可证](https://img.shields.io/github/license/gogf/gf.svg?style=flat)](https://github.com/gogf/gf)

[![发布版本](https://img.shields.io/github/v/release/gogf/gf?style=flat)](https://github.com/gogf/gf/releases)
[![GitHub 提交的拉取请求](https://img.shields.io/github/issues-pr/gogf/gf?style=flat)](https://github.com/gogf/gf/pulls)
[![GitHub 已关闭的拉取请求](https://img.shields.io/github/issues-pr-closed/gogf/gf?style=flat)](https://github.com/gogf/gf/pulls?q=is%3Apr+is%3Aclosed)
[![GitHub 问题](https://img.shields.io/github/issues/gogf/gf?style=flat)](https://github.com/gogf/gf/issues)
[![GitHub 已关闭的问题](https://img.shields.io/github/issues-closed/gogf/gf?style=flat)](https://github.com/gogf/gf/issues?q=is%3Aissue+is%3Aclosed)
![Stars](https://img.shields.io/github/stars/gogf/gf?style=flat)
![Forks](https://img.shields.io/github/forks/gogf/gf?style=flat)

</div>

`GoFrame` 是一个模块化、强大、高性能且适用于企业级应用开发的 Golang 框架。

// md5:a1eeb4ca8cc38c34

- modular, loosely coupled design
- rich components, out-of-the-box
- automatic codes generating for efficiency
- simple and easy to use, detailed documentation
- interface designed components, with high scalability
- fully supported tracing and error stack feature
- specially developed and powerful ORM component
- robust engineering design specifications
- convenient development CLI tool provide
- OpenTelemetry observability features support
- OpenAPIV3 documentation generating, automatically
- much, much more...ready to explore?

# # 安装

进入你的仓库目录，然后执行以下命令：

# 安装

进入你的仓库目录，然后执行以下命令：

// md5:e3515ec73e31e013

```bash
go get -u -v github.com/gogf/gf/v2
```

## # 命令行工具

```bash
go install github.com/gogf/gf/cmd/gf/v2@latest
```

此命令用于安装 `gf` 工具的最新版本，`gf` 是一个由 Go 语言编写的命令行工具，它属于 `gogf/gf` 项目中的 `v2` 版本。通过运行这个命令，您可以获取并安装该工具，以便在您的开发环境中使用。

# 命令行工具

```bash
go install github.com/gogf/gf/cmd/gf/v2@latest
```

此命令用于安装 `gf` 工具的最新版本，`gf` 是一个由 Go 语言编写的命令行工具，它属于 `gogf/gf` 项目中的 `v2` 版本。通过运行这个命令，您可以获取并安装该工具，以便在您的开发环境中使用。

// md5:7c864919a81959c9

```
golang version >= 1.18
```

# Documentation

- Chinese Official Site(中文官网): [https://goframe.org](https://goframe.org/display/gf)
- Chinese Pages Document(中文镜像文档): [https://pages.goframe.org](https://pages.goframe.org)
- Chinese Offline Document(中文离线文档): [https://github.com/gogf/goframe.org-pdf](https://github.com/gogf/goframe.org-pdf)
- GoDoc API: [https://pkg.go.dev/github.com/gogf/gf/v2](https://pkg.go.dev/github.com/gogf/gf/v2)

# # 许可

`GoFrame`遵循[MIT许可证](LICENSE)，它完全免费且开源，永远是这样。

# 许可

`GoFrame`遵循[MIT许可证](LICENSE)，它完全免费且开源，永远是这样。

// md5:13de2ffc1c07d37f

- [Tencent](https://www.tencent.com/)
- [ZTE](https://www.zte.com.cn/china/)
- [Ant Financial Services](https://www.antfin.com/)
- [VIVO](https://www.vivo.com/)
- [MedLinker](https://www.medlinker.com/)
- [KuCoin](https://www.kucoin.io/)
- [LeYouJia](https://www.leyoujia.com/)
- [IGG](https://igg.com)
- [37](https://www.37.com)
- [XiMaLaYa](https://www.ximalaya.com)
- [ZYBang](https://www.zybang.com/)

> We list part of the users here, if your company or products are using `GoFrame`, please let us know [here](https://goframe.org/pages/viewpage.action?pageId=1114415).

# # 贡献者

这个项目的成立得益于所有贡献者的努力。[[贡献者列表](https://github.com/gogf/gf/graphs/contributors)]。
[![贡献者画像](https://contributors-img.web.app/image?repo=gogf/gf)](https://github.com/gogf/gf/graphs/contributors)

# 贡献者

这个项目的成立得益于所有贡献者的努力。[[贡献者列表](https://github.com/gogf/gf/graphs/contributors)]。
[![贡献者画像](https://contributors-img.web.app/image?repo=gogf/gf)](https://github.com/gogf/gf/graphs/contributors)

// md5:6fa6efb14ea1c6b8

If you love `GoFrame`, why not [buy developer a cup of coffee](https://goframe.org/pages/viewpage.action?pageId=1115633)?

# # 赞助商

我们非常感谢对`GoFrame`开发的任何形式的赞助。如果您有任何有趣的想法或提议，请通过微信`389961817`或电子邮件`john@goframe.org`与我们联系。

# 赞助商

我们非常感谢对`GoFrame`开发的任何形式的赞助。如果您有任何有趣的想法或提议，请通过微信`389961817`或电子邮件`john@goframe.org`与我们联系。

// md5:3bce56c5fb676712

<a href="https://www.jetbrains.com/?from=GoFrame"><img src="https://goframe.org/download/thumbnails/1114119/jetbrains.png" height="120" alt="JetBrains"/></a>
<a href="https://www.atlassian.com/?from=GoFrame"><img src="https://goframe.org/download/attachments/1114119/atlassian.jpg" height="120" alt="Atlassian"/></a>
