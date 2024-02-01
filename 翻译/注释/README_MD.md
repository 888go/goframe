
<原文开始>
GoFrame

<div align=center>
<img src="https://goframe.org/statics/image/logo2.png?v=1" width="300"/>

[![Go Reference](https://pkg.go.dev/badge/github.com/gogf/gf/v2.svg)](https://pkg.go.dev/github.com/gogf/gf/v2)
[![GoFrame CI](https://github.com/gogf/gf/actions/workflows/ci-main.yml/badge.svg)](https://github.com/gogf/gf/actions/workflows/ci-main.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gogf/gf/v2)](https://goreportcard.com/report/github.com/gogf/gf/v2)
[![Code Coverage](https://codecov.io/gh/gogf/gf/branch/master/graph/badge.svg)](https://codecov.io/gh/gogf/gf)
[![Production Ready](https://img.shields.io/badge/production-ready-blue.svg)](https://github.com/gogf/gf)
[![License](https://img.shields.io/github/license/gogf/gf.svg?style=flat)](https://github.com/gogf/gf)

[![Release](https://img.shields.io/github/v/release/gogf/gf)](https://github.com/gogf/gf/releases)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/gogf/gf)](https://github.com/gogf/gf/pulls)
[![GitHub closed pull requests](https://img.shields.io/github/issues-pr-closed/gogf/gf)](https://github.com/gogf/gf/pulls?q=is%3Apr+is%3Aclosed)
[![GitHub issues](https://img.shields.io/github/issues/gogf/gf)](https://github.com/gogf/gf/issues)
[![GitHub closed issues](https://img.shields.io/github/issues-closed/gogf/gf)](https://github.com/gogf/gf/issues?q=is%3Aissue+is%3Aclosed)
![Stars](https://img.shields.io/github/stars/gogf/gf)
![Forks](https://img.shields.io/github/forks/gogf/gf)

</div>

`GoFrame` is a modular, powerful, high-performance and enterprise-class application development framework of Golang.


<原文结束>

# <翻译开始>
# GoFrame

<div align=center>
<img src="https://goframe.org/statics/image/logo2.png?v=1" width="300"/>

[![Go 语言引用](https://pkg.go.dev/badge/github.com/gogf/gf/v2.svg)](https://pkg.go.dev/github.com/gogf/gf/v2)
[![GoFrame CI](https://github.com/gogf/gf/actions/workflows/ci-main.yml/badge.svg)](https://github.com/gogf/gf/actions/workflows/ci-main.yml)
[![Go 报告卡](https://goreportcard.com/badge/github.com/gogf/gf/v2)](https://goreportcard.com/report/github.com/gogf/gf/v2)
[![代码覆盖率](https://codecov.io/gh/gogf/gf/branch/master/graph/badge.svg)](https://codecov.io/gh/gogf/gf)
[![生产环境就绪](https://img.shields.io/badge/production-ready-blue.svg)](https://github.com/gogf/gf)
[![许可证](https://img.shields.io/github/license/gogf/gf.svg?style=flat)](https://github.com/gogf/gf)

[![最新版本](https://img.shields.io/github/v/release/gogf/gf)](https://github.com/gogf/gf/releases)
[![GitHub 提交的拉取请求](https://img.shields.io/github/issues-pr/gogf/gf)](https://github.com/gogf/gf/pulls)
[![GitHub 已关闭的拉取请求](https://img.shields.io/github/issues-pr-closed/gogf/gf)](https://github.com/gogf/gf/pulls?q=is%3Apr+is%3Aclosed)
[![GitHub 问题](https://img.shields.io/github/issues/gogf/gf)](https://github.com/gogf/gf/issues)
[![GitHub 已关闭的问题](https://img.shields.io/github/issues-closed/gogf/gf)](https://github.com/gogf/gf/issues?q=is%3Aissue+is%3Aclosed)
![Stars](https://img.shields.io/github/stars/gogf/gf)
![Forks](https://img.shields.io/github/forks/gogf/gf)

</div>

`GoFrame` 是一个模块化、强大、高性能的企业级 Golang 应用开发框架。

# <翻译结束>


<原文开始>
Features

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


<原文结束>

# <翻译开始>
# 特性

- 模块化、松耦合设计
- 丰富的开箱即用组件
- 高效的自动代码生成功能
- 简单易用，详尽的文档说明
- 设计精良且具有高度可扩展性的接口化组件
- 完备的追踪与错误堆栈功能支持
- 特别研发的强大ORM组件
- 坚固的工程设计规范
- 提供便捷的开发CLI工具
- 支持OpenTelemetry可观测性特性
- 自动生成OpenAPIV3规范文档
- 更多精彩内容等待探索...准备好了吗？

# <翻译结束>


<原文开始>
Installation

Enter your repo. directory and execute following command:


<原文结束>

# <翻译开始>
# 安装

进入您的仓库目录并执行以下命令：

# <翻译结束>


<原文开始>
primary module

```bash
go get -u -v github.com/gogf/gf/v2
```


<原文结束>

# <翻译开始>
# 主要模块

```bash
go get -u -v github.com/gogf/gf/v2
```

（命令行操作：更新并获取github.com/gogf/gf/v2库的最新版本，其中-u表示更新已存在的包，-v表示输出详细信息）

# <翻译结束>


<原文开始>
cli tool

```bash
go install github.com/gogf/gf/cmd/gf/v2@latest
```


<原文结束>

# <翻译开始>
# cli 工具

```bash
go install github.com/gogf/gf/cmd/gf/v2@latest
```

（该命令用于安装Go语言开发框架gf的最新版本v2的命令行工具）

# <翻译结束>


<原文开始>
Limitation

```
golang version >= 1.18
```


<原文结束>

# <翻译开始>
# 限制

```
golang 版本 >= 1.18
```

# <翻译结束>


<原文开始>
License

`GoFrame` is licensed under the [MIT License](LICENSE), 100% free and open-source, forever.


<原文结束>

# <翻译开始>
# 许可

`GoFrame`遵循[MIT许可协议](LICENSE)进行授权，100%完全免费且开源，永久有效。

# <翻译结束>


<原文开始>
Part Of Users

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


<原文结束>

# <翻译开始>
# 用户部分

- [腾讯](https://www.tencent.com/)
- [中兴通讯](https://www.zte.com.cn/china/)
- [蚂蚁集团](https://www.antfin.com/)
- [VIVO](https://www.vivo.com/)
- [医联](https://www.medlinker.com/)
- [KuCoin](https://www.kucoin.io/)
- [乐游家](https://www.leyoujia.com/)
- [IGG](https://igg.com)
- [37游戏](https://www.37.com)
- [喜马拉雅](https://www.ximalaya.com)
- [作业帮](https://www.zybang.com/)

> 我们在此列出部分用户，如果您所在的公司或产品正在使用`GoFrame`，请在[此处](https://goframe.org/pages/viewpage.action?pageId=1114415)告知我们。

# <翻译结束>


<原文开始>
Contributors

This project exists thanks to all the people who contribute. [[Contributors](https://github.com/gogf/gf/graphs/contributors)].
<a href="https://github.com/gogf/gf/graphs/contributors"><img src="https://contributors-img.web.app/image?repo=gogf/gf" /></a>


<原文结束>

# <翻译开始>
# 贡献者

这个项目的存在要归功于所有贡献的人。[[贡献者](https://github.com/gogf/gf/graphs/contributors)]。
<img src="https://contributors-img.web.app/image?repo=gogf/gf" />

# <翻译结束>


<原文开始>
Donators

If you love `GoFrame`, why not [buy developer a cup of coffee](https://goframe.org/pages/viewpage.action?pageId=1115633)?


<原文结束>

# <翻译开始>
# 捐赠者

如果你喜欢 `GoFrame`，何不[请开发者喝杯咖啡](https://goframe.org/pages/viewpage.action?pageId=1115633)呢？

# <翻译结束>


<原文开始>
Sponsors

We appreciate any kind of sponsorship for `GoFrame` development. If you've got some interesting, please contact WeChat `389961817` / Email `john@goframe.org`.


<原文结束>

# <翻译开始>
# 赞助

我们非常感谢任何对`GoFrame`开发提供的赞助支持。如果您对此感兴趣，请通过微信`389961817`或邮箱`john@goframe.org`与我们联系。

# <翻译结束>


<原文开始>
Thanks

<a href="https://www.jetbrains.com/?from=GoFrame"><img src="https://goframe.org/download/thumbnails/1114119/jetbrains.png" height="120" alt="JetBrains"/></a>
<a href="https://www.atlassian.com/?from=GoFrame"><img src="https://goframe.org/download/attachments/1114119/atlassian.jpg" height="120" alt="Atlassian"/></a>

<原文结束>

# <翻译开始>
# 感谢

[JetBrains](https://www.jetbrains.com/?from=GoFrame)  
<img src="https://goframe.org/download/thumbnails/1114119/jetbrains.png" height="120" alt="JetBrains"/>

[Atlassian](https://www.atlassian.com/?from=GoFrame)  
<img src="https://goframe.org/download/attachments/1114119/atlassian.jpg" height="120" alt="Atlassian"/>

# <翻译结束>

