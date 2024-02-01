// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package swaggerui 提供了通过资源管理器打包的 Swagger UI 静态文件。
//
// 文件来源：
// https://github.com/Redocly/redoc
// https://www.jsdelivr.com/package/npm/redoc
//
// 打包命令：
// gf pack redoc.standalone.js swaggerui-redoc.go -n=swaggerui -p=/goframe/swaggerui
// （该注释说明了这个Go语言包的作用是提供Swagger UI静态文件，这些文件来源于Redocly的GitHub仓库和jsDelivr的npm包。同时给出了使用gf工具进行文件打包的具体命令，将redoc.standalone.js打包到名为swaggerui-redoc.go的文件中，并设置了包名和路径前缀。）
package swaggerui
