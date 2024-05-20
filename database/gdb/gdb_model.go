// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gdb

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// Model是实现ORM DAO的核心结构体。. md5:7230072d015718fc
type Model struct {
	db             DB                // 底层数据库接口。. md5:5b73fd8ce3fdaf5a
	tx             TX                // 底层的TX接口。. md5:d13e03783f7815aa
	rawSql         string            // rawSql 是原始的SQL字符串，它标志着一个基于原始SQL的模型，而不是基于表的模型。. md5:b83edc253c98f3de
	schema         string            // 自定义数据库模式。. md5:173e3cf7ad252f16
	linkType       int               // 用于标记在主节点或从节点上执行的操作。. md5:d59587510f982160
	tablesInit     string            // 模型初始化时的表格名称。. md5:7569da250a03d8b2
	tables         string            // 操作表名，可以是多个表名和别名，例如："user"、"user u"、"user u, user_detail ud"。. md5:140ed796dfa7b2e5
	fields         string            // 操作字段，多个字段通过字符'.'连接。. md5:90a8233be912ab73
	fieldsEx       string            // 排除的操作字段，多个字段使用逗号','连接。. md5:0757072228393ad3
	withArray      []interface{}     // With功能的参数。. md5:6da19be4d3cc5337
	withAll        bool              // 启用在结构体中带有 "with" 标签的所有对象的模型关联操作。. md5:68534968f6dd65cd
	extraArgs      []interface{}     // 为SQL提供的额外自定义参数，这些参数将在将SQL提交给底层驱动程序之前添加到参数前面。. md5:e6a840d23cdc5b31
	whereBuilder   *WhereBuilder     // 用于构建where操作的条件生成器。. md5:4e7d38dd793619e1
	groupBy        string            // 用于"分组 by"语句。. md5:0054c7d82c75aa83
	orderBy        string            // 用于"ORDER BY"语句。. md5:974c6823a972edbe
	having         []interface{}     // 用于 "having..." 语句。. md5:fc87b6be31414f4e
	start          int               // 用于 "select ... start, limit ..." 语句。. md5:28a92730f0f33ffe
	limit          int               // 用于 "select ... start, limit ..." 语句。. md5:28a92730f0f33ffe
	option         int               // 用于额外操作功能的选项。. md5:46fa8be84b899e8f
	offset         int               // 用于某些数据库语法的偏移语句。. md5:222cd8b108c2f2fc
	partition      string            // 分区表的分区名称。. md5:f8b787fa2b446be6
	data           interface{}       // 操作数据，可以是映射类型/切片映射/结构体/结构体指针/字符串等类型。. md5:d9d2d3cef3841513
	batch          int               // 批量操作的批次编号，用于批量插入/替换/保存操作。. md5:72e06e8a06a3dfa8
	filter         bool              // 根据表格的字段，过滤数据和键值对。. md5:6af1b96126cc53e6
	distinct       string            // 强制查询只返回唯一的结果。. md5:10ef583cb57e7d16
	lockInfo       string            // 用于更新或共享锁的锁定。. md5:a2e8bcf922a3cd09
	cacheEnabled   bool              // 启用SQL结果缓存功能，主要用于指示缓存持续时间（尤其是0）的使用。. md5:426f2265f5437a86
	cacheOption    CacheOption       // 查询语句的缓存选项。. md5:0243bce17a4463a8
	hookHandler    HookHandler       // 用于模型钩子功能的钩子函数。. md5:cb10889f174ab53d
	unscoped       bool              // 在进行选择/删除操作时，禁用软删除特性。. md5:b6cc5bc9aefe18bf
	safe           bool              // 如果为真，每次操作后都会克隆并返回一个新的模型对象；否则，它将修改当前模型的属性。. md5:b4a3ad6d8438d2de
	onDuplicate    interface{}       // onDuplicate 用于 upsert 子句。. md5:e139824b32378802
	onDuplicateEx  interface{}       // onDuplicateEx 用于在 Upsert 子句中排除某些列。. md5:f985786a6831d9ec
	onConflict     interface{}       // onConflict 用于在 Upsert 子句中处理冲突键。. md5:ec57ba30c97c0bd2
	tableAliasMap  map[string]string // 表别名到真实表名的映射，通常在JOIN语句中使用。. md5:5951bd1c3aa8b870
	softTimeOption SoftTimeOption    // SoftTimeOption 是用于自定义 Model 的软时间功能的选项。. md5:fcc19f5ef8ad45e7
}

// ModelHandler 是一个处理给定 Model 并返回一个自定义修改后的新 Model 的函数。. md5:a02af46ff8fb2568
type ModelHandler func(m *Model) *Model

// ChunkHandler 是一个函数，用于 Chunk 函数中，负责处理给定的结果和错误。
// 如果希望继续分块处理，则返回true；否则返回false以停止分块。
// md5:e7b2a1b4761ac415
type ChunkHandler func(result Result, err error) bool

const (
	linkTypeMaster           = 1
	linkTypeSlave            = 2
	defaultFields            = "*"
	whereHolderOperatorWhere = 1
	whereHolderOperatorAnd   = 2
	whereHolderOperatorOr    = 3
	whereHolderTypeDefault   = "Default"
	whereHolderTypeNoArgs    = "NoArgs"
	whereHolderTypeIn        = "In"
)

// Model 根据给定的模式创建并返回一个新的ORM模型。
// 参数 `tableNameQueryOrStruct` 可以是多个表名，也可以是别名，例如：
// 1. 模型名称：
//     db.Model("user")
//     db.Model("user u")
//     db.Model("user, user_detail")
//     db.Model("user u, user_detail ud")
// 2. 带别名的模型名称：
//     db.Model("user", "u")
// 3. 带子查询的模型名称：
//     db.Model("? AS a, ? AS b", subQuery1, subQuery2)
// md5:add855a912a9b6ef
func (c *Core) Model(tableNameQueryOrStruct ...interface{}) *Model {
	var (
		ctx       = c.db.GetCtx()
		tableStr  string
		tableName string
		extraArgs []interface{}
	)
	// 使用子查询创建模型。. md5:1c8112f948bca053
	if len(tableNameQueryOrStruct) > 1 {
		conditionStr := gconv.String(tableNameQueryOrStruct[0])
		if gstr.Contains(conditionStr, "?") {
			whereHolder := WhereHolder{
				Where: conditionStr,
				Args:  tableNameQueryOrStruct[1:],
			}
			tableStr, extraArgs = formatWhereHolder(ctx, c.db, formatWhereHolderInput{
				WhereHolder: whereHolder,
				OmitNil:     false,
				OmitEmpty:   false,
				Schema:      "",
				Table:       "",
			})
		}
	}
	// Normal model creation.
	if tableStr == "" {
		tableNames := make([]string, len(tableNameQueryOrStruct))
		for k, v := range tableNameQueryOrStruct {
			if s, ok := v.(string); ok {
				tableNames[k] = s
			} else if tableName = getTableNameFromOrmTag(v); tableName != "" {
				tableNames[k] = tableName
			}
		}
		if len(tableNames) > 1 {
			tableStr = fmt.Sprintf(
				`%s AS %s`, c.QuotePrefixTableName(tableNames[0]), c.QuoteWord(tableNames[1]),
			)
		} else if len(tableNames) == 1 {
			tableStr = c.QuotePrefixTableName(tableNames[0])
		}
	}
	m := &Model{
		db:            c.db,
		schema:        c.schema,
		tablesInit:    tableStr,
		tables:        tableStr,
		fields:        defaultFields,
		start:         -1,
		offset:        -1,
		filter:        true,
		extraArgs:     extraArgs,
		tableAliasMap: make(map[string]string),
	}
	m.whereBuilder = m.Builder()
	if defaultModelSafe {
		m.safe = true
	}
	return m
}

// Raw根据原始SQL（而不是表）创建并返回一个模型。示例：
//
//	db.Raw("SELECT * FROM `user` WHERE `name` = ?", "john").Scan(&result)
// md5:0865d39f2ab854cb
func (c *Core) Raw(rawSql string, args ...interface{}) *Model {
	model := c.Model()
	model.rawSql = rawSql
	model.extraArgs = args
	return model
}

// Raw 将当前模型设置为原始SQL模型。
// 示例：
//
//	db.Raw("SELECT * FROM `user` WHERE `name` = ?", "john").Scan(&result)
//
// 参见 Core.Raw。
// md5:ced75308536ddfff
func (m *Model) Raw(rawSql string, args ...interface{}) *Model {
	model := m.db.Raw(rawSql, args...)
	model.db = m.db
	model.tx = m.tx
	return model
}

func (tx *TXCore) Raw(rawSql string, args ...interface{}) *Model {
	return tx.Model().Raw(rawSql, args...)
}

// With 根据给定对象的元数据创建并返回一个ORM模型。. md5:18604e26c0c946fb
func (c *Core) With(objects ...interface{}) *Model {
	return c.db.Model().With(objects...)
}

// 分区设置分区名称。
// 例子：
// dao.User.Ctx(ctx).Partition("p1", "p2", "p3").All() 
// 
// 这段Go代码的注释表示：使用`Partition`方法对数据进行分区操作，传入多个分区名称（如："p1", "p2", "p3"），然后在查询时指定这些分区。`Ctx(ctx)`表示使用上下文`ctx`进行操作。`All()`是获取所有满足条件的数据。
// md5:f133a577ba31c05f
func (m *Model) Partition(partitions ...string) *Model {
	model := m.getModel()
	model.partition = gstr.Join(partitions, ",")
	return model
}

// Model 类似于 Core.Model，但它是基于事务操作的。
// 请参阅 Core.Model。
// md5:2c5866afc2e5dd90
func (tx *TXCore) Model(tableNameQueryOrStruct ...interface{}) *Model {
	model := tx.db.Model(tableNameQueryOrStruct...)
	model.db = tx.db
	model.tx = tx
	return model
}

// With 的行为类似于 Core.With，但它是在事务上操作。
// 参见 Core.With。
// md5:37000d6ea41561fc
func (tx *TXCore) With(object interface{}) *Model {
	return tx.Model().With(object)
}

// Ctx 设置当前操作的上下文。. md5:77d589f34a65753b
func (m *Model) Ctx(ctx context.Context) *Model {
	if ctx == nil {
		return m
	}
	model := m.getModel()
	model.db = model.db.Ctx(ctx)
	if m.tx != nil {
		model.tx = model.tx.Ctx(ctx)
	}
	return model
}

// GetCtx返回当前Model的上下文。
// 如果之前没有设置上下文，则返回`context.Background()`。
// md5:48edd9b438a38523
func (m *Model) GetCtx() context.Context {
	if m.tx != nil && m.tx.GetCtx() != nil {
		return m.tx.GetCtx()
	}
	return m.db.GetCtx()
}

// As 设置当前表的别名名称。. md5:c28e3f79c6fe2e48
func (m *Model) As(as string) *Model {
	if m.tables != "" {
		model := m.getModel()
		split := " JOIN "
		if gstr.ContainsI(model.tables, split) {
			// For join table.
			array := gstr.Split(model.tables, split)
			array[len(array)-1], _ = gregex.ReplaceString(`(.+) ON`, fmt.Sprintf(`$1 AS %s ON`, as), array[len(array)-1])
			model.tables = gstr.Join(array, split)
		} else {
			// For base table.
			model.tables = gstr.TrimRight(model.tables) + " AS " + as
		}
		return model
	}
	return m
}

// DB 为当前操作设置或更改 db 对象。. md5:1761cc3b00f1d6bb
func (m *Model) DB(db DB) *Model {
	model := m.getModel()
	model.db = db
	return model
}

// TX 设置或更改当前操作的事务。. md5:7171a26d8d2d8431
func (m *Model) TX(tx TX) *Model {
	model := m.getModel()
	model.db = tx.GetDB()
	model.tx = tx
	return model
}

// Schema 设置当前操作的模式。. md5:723e31c5f24ff604
func (m *Model) Schema(schema string) *Model {
	model := m.getModel()
	model.schema = schema
	return model
}

// Clone 创建并返回一个新的模型，它是当前模型的克隆。请注意，它使用深拷贝进行克隆。
// md5:27e973f2f4fb42b3
func (m *Model) Clone() *Model {
	newModel := (*Model)(nil)
	if m.tx != nil {
		newModel = m.tx.Model(m.tablesInit)
	} else {
		newModel = m.db.Model(m.tablesInit)
	}
	// Basic attributes copy.
	*newModel = *m
	// WhereBuilder 的复制方法，注意属性是指针。. md5:c9aa75059eb72059
	newModel.whereBuilder = m.whereBuilder.Clone()
	newModel.whereBuilder.model = newModel
	// 浅复制切片属性。. md5:d03df5f661b330b7
	if n := len(m.extraArgs); n > 0 {
		newModel.extraArgs = make([]interface{}, n)
		copy(newModel.extraArgs, m.extraArgs)
	}
	if n := len(m.withArray); n > 0 {
		newModel.withArray = make([]interface{}, n)
		copy(newModel.withArray, m.withArray)
	}
	return newModel
}

// Master 在主节点上标记以下操作。. md5:86cff0c5fb8d6d5d
func (m *Model) Master() *Model {
	model := m.getModel()
	model.linkType = linkTypeMaster
	return model
}

// Slave 在 slave 节点上标记以下操作。
// 请注意，只有在配置了 slave 节点的情况下，此注释才有意义。
// md5:3d6dbca5bafb9cdf
func (m *Model) Slave() *Model {
	model := m.getModel()
	model.linkType = linkTypeSlave
	return model
}

// Safe 标记此模型为安全或不安全。如果 safe 为 true，那么在执行完操作后，它会克隆并返回一个新的模型对象；
// 否则，它将直接修改当前模型的属性。
// md5:56aecad30556ca98
func (m *Model) Safe(safe ...bool) *Model {
	if len(safe) > 0 {
		m.safe = safe[0]
	} else {
		m.safe = true
	}
	return m
}

// Args 为模型操作设置自定义参数。. md5:6cf507acdf0e2401
func (m *Model) Args(args ...interface{}) *Model {
	model := m.getModel()
	model.extraArgs = append(model.extraArgs, args)
	return model
}

// Handler calls each of `handlers` on current Model and returns a new Model.
// ModelHandler 是一个处理给定 Model 并返回一个自定义修改后的新 Model 的函数。. md5:a02af46ff8fb2568
func (m *Model) Handler(handlers ...ModelHandler) *Model {
	model := m.getModel()
	for _, handler := range handlers {
		model = handler(model)
	}
	return model
}
