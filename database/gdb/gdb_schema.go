// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

// Schema 是一个模式对象，从中可以创建一个模型。 md5:59296199d3fdabf5
type Schema struct {
	DB
}

// X切换数据库 创建并返回一个模式。 md5:2857e60cfa18e20e
func (c *Core) X切换数据库(数据库名 string) *Schema {
	// 不要更改原始数据库的模式，
	// 这里会创建一个新的数据库并改变其模式。
	// md5:a0cc2eeb4148cd74
	db, err := X创建DB对象并按配置组(c.X取配置组名称())
	if err != nil {
		panic(err)
	}
	core := db.X取Core对象()
		// 不同的模式共享一些相同的对象。 md5:dcaf41f78fadc582
	core.logger = c.logger
	core.cache = c.cache
	core.schema = 数据库名
	return &Schema{
		DB: db,
	}
}
