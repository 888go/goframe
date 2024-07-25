// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

//go:build !windows

package gproc

// 在Windows平台什么都不做，只是设置它. md5:ff86003bf629168f
func joinProcessArgs(p *Process) {}
