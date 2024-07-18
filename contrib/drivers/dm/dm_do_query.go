// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package dm

// TODO I originally wanted to only convert keywords in select
// 但是我发现 DoQuery 中会对 sql 会对 " " 达梦的安全字符 进行 / 转义，最后还是导致达梦无法正常解析
// However, I found that DoQuery() will perform / escape on sql with " " Dameng's safe characters, which ultimately caused Dameng to be unable to parse normally.
// But processing in DoFilter() is OK
// func (d *Driver) DoQuery(ctx context.Context, link gdb.Link, sql string, args ...interface{}) (gdb.Result, error) {
// 	l, r := d.GetChars()
// 	new := gstr.ReplaceI(sql, "INDEX", l+"INDEX"+r)
// 	g.Dump("new:", new)
// 	return d.Core.DoQuery(
// 		ctx,
// 		link,
// 		new,
// 		args,
// 	)
// }
