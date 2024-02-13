// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

// DriverWrapper 是一个驱动程序包装器，用于通过嵌入式驱动扩展功能。
type DriverWrapper struct {
	driver X驱动
}

// New 创建并返回一个用于 mysql 的数据库对象。
// 它实现了 gdb.Driver 接口，以便进行额外的数据库驱动安装。
func (d *DriverWrapper) New(core *Core, node *X配置项) (DB, error) {
	db, err := d.driver.New(core, node)
	if err != nil {
		return nil, err
	}
	return &DriverWrapperDB{
		DB: db,
	}, nil
}

// newDriverWrapper 创建并返回一个驱动程序包装器。
func newDriverWrapper(driver X驱动) X驱动 {
	return &DriverWrapper{
		driver: driver,
	}
}
