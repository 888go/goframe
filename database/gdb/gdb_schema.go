// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb

// Schema 是一个模式对象，通过它可以创建一个 Model。
type Schema struct {
	DB
}

// Schema 创建并返回一个模式（schema）。
func (c *Core) Schema(schema string) *Schema {
// 不要更改原始数据库的模式，
// 而是在这里创建一个新的数据库并更改其模式。
	db, err := NewByGroup(c.GetGroup())
	if err != nil {
		panic(err)
	}
	core := db.GetCore()
	// 不同的模式共享一些相同的对象。
	core.logger = c.logger
	core.cache = c.cache
	core.schema = schema
	return &Schema{
		DB: db,
	}
}
