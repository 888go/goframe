// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

import (
	"context"
	"fmt"
	
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
)

// Model 是核心结构体，实现了 ORM 的 DAO（数据访问对象）。
type Model struct {
	db            DB                // 基础数据库接口。
	tx            X事务                // 基础的TX接口。
	rawSql        string            // rawSql 是原始SQL字符串，用于标记基于原始SQL的模型，而非基于表的模型。
	schema        string            // 自定义数据库模式
	linkType      int               // 标记用于在主服务器或从服务器上执行操作。
	tablesInit    string            // 在模型初始化时的表格名称。
	tables        string            // 操作表名，可以是多个表名及别名，例如："user"、"user u"、"user u, user_detail ud"。
	fields        string            // 操作字段，多个字段使用字符 ',' 连接。
	fieldsEx      string            // 排除的操作字段，多个字段使用逗号（char ',')连接。
	withArray     []interface{}     // With功能的参数。
	withAll       bool              // 在结构体中具有"with"标签的所有对象上启用模型关联操作。
	extraArgs     []interface{}     // 在SQL提交给底层驱动之前，额外自定义的SQL参数，这些参数将被追加到原有参数之前。
	whereBuilder  *X组合条件     // 条件构造器，用于where操作。
	groupBy       string            // 用于“group by”语句。
	orderBy       string            // 用于 "order by" 语句。
	having        []interface{}     // 用于 "having..." 语句。
	start         int               // 用于 "select ... start, limit ..." 语句。
	limit         int               // 用于 "select ... start, limit ..." 语句。
	option        int               // 用于额外操作功能的选项。
	offset        int               // 为某些数据库语法提供的偏移量语句。
	partition     string            // 分区表分区名称。
	data          interface{}       // Data 用于操作的数据，其类型可以是 map/[]map/struct/*struct/string 等等。
	batch         int               // 批量插入/替换/保存操作的批次号。
	filter        bool              // 根据表格字段过滤并筛选出符合条件的键值对数据。
	distinct      string            // 强制查询只返回不重复的结果。
	lockInfo      string            // 加锁以便进行更新或共享锁操作。
	cacheEnabled  bool              // 启用SQL结果缓存功能，主要用于指示缓存持续时间（尤其是0）的使用情况。
	cacheOption   X缓存选项       // 查询语句的缓存选项。
	hookHandler   HookHandler       // 钩子函数，用于模型钩子功能。
	unscoped      bool              // 禁用在选择/删除操作时的软删除功能。
	safe          bool              // 如果为true，则在操作完成后克隆并返回一个新的模型对象；否则，它会改变当前模型的属性。
	onDuplicate   interface{}       // onDuplicate 用于 ON "DUPLICATE KEY UPDATE" 语句。
	onDuplicateEx interface{}       // onDuplicateEx 用于在 "DUPLICATE KEY UPDATE" 语句中排除某些列。
	tableAliasMap map[string]string // 表别名到真实表名的映射，通常用于连接语句中。
}

// ModelHandler 是一个函数，用于处理给定的 Model，并返回一个经过自定义修改的新 Model。
type ModelHandler func(m *Model) *Model

// ChunkHandler 是一个在 Chunk 函数中使用的函数，用于处理给定的 Result 和错误。
// 如果希望继续分块处理，则返回 true；否则返回 false 以停止分块处理。
type ChunkHandler func(result X行记录数组, err error) bool

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

// Model 根据给定的模式创建并返回一个新的 ORM 模型。
// 参数 `tableNameQueryOrStruct` 可以是多个表名，也可以包含别名，如下所示：
// 1. 表名示例：
//     db.Model("user")                 // 单个表名
//     db.Model("user u")               // 带别名的表名
//     db.Model("user, user_detail")    // 多个表名
//     db.Model("user u, user_detail ud") // 多个带别名的表名
// 2. 包含别名的表名示例：
//     db.Model("user", "u")         // 表名和对应的别名
// 3. 使用子查询作为表名的示例：
//     db.Model("? AS a, ? AS b", subQuery1, subQuery2) // 使用子查询表达式作为模型，并为子查询结果设置别名
func (c *Core) X创建Model对象(表名或结构体 ...interface{}) *Model {
	var (
		ctx       = c.db.X取上下文对象()
		tableStr  string
		tableName string
		extraArgs []interface{}
	)
	// 使用子查询创建模型
	if len(表名或结构体) > 1 {
		conditionStr := 转换类.String(表名或结构体[0])
		if 文本类.X是否包含(conditionStr, "?") {
			whereHolder := WhereHolder{
				Where: conditionStr,
				Args:  表名或结构体[1:],
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
	// 正常模型创建。
	if tableStr == "" {
		tableNames := make([]string, len(表名或结构体))
		for k, v := range 表名或结构体 {
			if s, ok := v.(string); ok {
				tableNames[k] = s
			} else if tableName = getTableNameFromOrmTag(v); tableName != "" {
				tableNames[k] = tableName
			}
		}
		if len(tableNames) > 1 {
			tableStr = fmt.Sprintf(
				`%s AS %s`, c.X底层添加前缀字符和引用字符(tableNames[0]), c.X底层QuoteWord(tableNames[1]),
			)
		} else if len(tableNames) == 1 {
			tableStr = c.X底层添加前缀字符和引用字符(tableNames[0])
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
	m.whereBuilder = m.X创建组合条件()
	if defaultModelSafe {
		m.safe = true
	}
	return m
}

// Raw 根据原始 SQL（非表）创建并返回一个模型。
// 通常用于嵌入原始sql语句,如:
// g.Model("user").WhereLT("created_at", gdb.Raw("now()")).All()  // SELECT * FROM `user` WHERE `created_at`<now()
// 参考文档:https://goframe.org/pages/viewpage.action?pageId=111911590&showComments=true

// 也可以直接直接执行原始sql,示例：
// db.Raw("SELECT * FROM `user` WHERE `name` = ?", "john").Scan(&result)
// 上述代码表示，通过执行原始SQL语句（根据"name"为"john"的条件查询user表中所有列），并使用Scan方法将查询结果绑定到result变量中。
func (c *Core) X原生SQL(原生Sql string, 参数 ...interface{}) *Model {
	model := c.X创建Model对象()
	model.rawSql = 原生Sql
	model.extraArgs = 参数
	return model
}

// Raw 将当前模型设置为原始SQL模型。
// 通常用于嵌入原始sql语句,如:
// g.Model("user").WhereLT("created_at", gdb.Raw("now()")).All()  // SELECT * FROM `user` WHERE `created_at`<now()
// 参考文档:https://goframe.org/pages/viewpage.action?pageId=111911590&showComments=true
//
// 也可以直接直接执行原始sql,示例:
// db.Raw("SELECT * FROM `user` WHERE `name` = ?", "john").Scan(&result)
// 请参阅Core.Raw。
func (m *Model) X原生SQL(原生Sql string, 参数 ...interface{}) *Model {
	model := m.db.X原生SQL(原生Sql, 参数...)
	model.db = m.db
	model.tx = m.tx
	return model
}

func (tx *X基础事务) X原生SQL(原生Sql string, 参数 ...interface{}) *Model {
	return tx.X创建Model对象().X原生SQL(原生Sql, 参数...)
}

// With 根据给定对象的元数据创建并返回一个ORM模型。
//
// 原注释未提及with使用方法, 以下摘自Model对象示例,仅供参考.
// With 启用关联查询，通过给定的属性对象指定开启。
// 常考"模型关联-静态关联"文档:https://goframe.org/pages/viewpage.action?pageId=7297190
// 例如，如果给定如下的结构体定义：
//
//	type User struct {
//		 gmeta.Meta `orm:"table:user"` // 定义表名为 user
//		 Id         int           `json:"id"`    // 用户ID
//		 Name       string        `json:"name"`   // 用户名
//		 UserDetail *UserDetail   `orm:"with:uid=id"` // 关联 UserDetail 表，通过 uid 等于 id 进行关联
//		 UserScores []*UserScores `orm:"with:uid=id"` // 关联 UserScores 表，通过 uid 等于 id 进行关联
//	}
//
// 我们可以通过以下方式在属性 `UserDetail` 和 `UserScores` 上启用模型关联操作：
// db.With(User{}.UserDetail).With(User{}.UserScores).Scan(xxx)
// 或者：
// db.With(UserDetail{}).With(UserScores{}).Scan(xxx)
// 或者：
// db.With(UserDetail{}, UserScores{}).Scan(xxx)
func (c *Core) X关联对象(关联结构体 ...interface{}) *Model {
	return c.db.X创建Model对象().X关联对象(关联结构体...)
}

// Partition sets Partition name.
// Example:
// dao.User.Ctx(ctx).Partition（"p1","p2","p3").All()
func (m *Model) X设置分区名称(分区名称 ...string) *Model {
	model := m.getModel()
	model.partition = 文本类.X连接(分区名称, ",")
	return model
}

// Model 类似于 Core.Model，但其在事务上进行操作。
// 请参阅 Core.Model。
func (tx *X基础事务) X创建Model对象(表名或结构体 ...interface{}) *Model {
	model := tx.db.X创建Model对象(表名或结构体...)
	model.db = tx.db
	model.tx = tx
	return model
}

// With 类似于 Core.With，但其操作针对事务。
// 请参阅 Core.With。
//
// 原注释未提及with使用方法, 以下摘自Model对象示例,仅供参考.
// With 启用关联查询，通过给定的属性对象指定开启。
// 常考"模型关联-静态关联"文档:https://goframe.org/pages/viewpage.action?pageId=7297190
// 例如，如果给定如下的结构体定义：
//	type User struct {
//		 gmeta.Meta `orm:"table:user"` // 定义表名为 user
//		 Id         int           `json:"id"`    // 用户ID
//		 Name       string        `json:"name"`   // 用户名
//		 UserDetail *UserDetail   `orm:"with:uid=id"` // 关联 UserDetail 表，通过 uid 等于 id 进行关联
//		 UserScores []*UserScores `orm:"with:uid=id"` // 关联 UserScores 表，通过 uid 等于 id 进行关联
//	}
//
// 我们可以通过以下方式在属性 `UserDetail` 和 `UserScores` 上启用模型关联操作：
// db.With(User{}.UserDetail).With(User{}.UserScores).Scan(xxx)
// 或者：
// db.With(UserDetail{}).With(UserScores{}).Scan(xxx)
// 或者：
// db.With(UserDetail{}, UserScores{}).Scan(xxx)
func (tx *X基础事务) X关联对象(关联结构体 interface{}) *Model {
	return tx.X创建Model对象().X关联对象(关联结构体)
}

// Ctx 设置当前操作的上下文。
func (m *Model) X设置上下文并取副本(上下文 context.Context) *Model {
	if 上下文 == nil {
		return m
	}
	model := m.getModel()
	model.db = model.db.X设置上下文并取副本(上下文)
	if m.tx != nil {
		model.tx = model.tx.X设置上下文并取副本(上下文)
	}
	return model
}

// GetCtx 返回当前 Model 的上下文。
// 若此前未设置过上下文，则返回 `context.Background()`。
func (m *Model) X取上下文对象() context.Context {
	if m.tx != nil && m.tx.X取上下文对象() != nil {
		return m.tx.X取上下文对象()
	}
	return m.db.X取上下文对象()
}

// As 为当前表设置别名名称。
func (m *Model) X设置表别名(别名 string) *Model {
	if m.tables != "" {
		model := m.getModel()
		split := " JOIN "
		if 文本类.X是否包含并忽略大小写(model.tables, split) {
			// For join table.
			array := 文本类.X分割(model.tables, split)
			array[len(array)-1], _ = 正则类.X替换文本(`(.+) ON`, fmt.Sprintf(`$1 AS %s ON`, 别名), array[len(array)-1])
			model.tables = 文本类.X连接(array, split)
		} else {
			// For base table.
			model.tables = 文本类.X过滤尾字符并含空白(model.tables) + " AS " + 别名
		}
		return model
	}
	return m
}

// DB 设置/更改当前操作的数据库对象。
func (m *Model) X设置DB对象(DB对象 DB) *Model {
	model := m.getModel()
	model.db = DB对象
	return model
}

// TX 设置/更改当前操作的事务。
func (m *Model) X设置事务对象(事务对象 X事务) *Model {
	model := m.getModel()
	model.db = 事务对象.X取DB对象()
	model.tx = 事务对象
	return model
}

// 设置当前操作的模式。
func (m *Model) X切换数据库(数据库名 string) *Model {
	model := m.getModel()
	model.schema = 数据库名
	return model
}

// Clone 创建并返回一个新的模型，该模型是当前模型的克隆版本。
// 注意，它使用深度复制进行克隆。
func (m *Model) X取副本() *Model {
	newModel := (*Model)(nil)
	if m.tx != nil {
		newModel = m.tx.X创建Model对象(m.tablesInit)
	} else {
		newModel = m.db.X创建Model对象(m.tablesInit)
	}
	// 基础属性复制
	*newModel = *m
	// WhereBuilder 复制，注意属性指针。
	newModel.whereBuilder = m.whereBuilder.X取副本()
	newModel.whereBuilder.model = newModel
	// 浅复制切片属性。
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

// Master 标识以下操作将在主节点上执行。
func (m *Model) X取主节点对象() *Model {
	model := m.getModel()
	model.linkType = linkTypeMaster
	return model
}

// Slave 标记在从属节点上执行的后续操作。
// 注意，只有在配置了从属节点时才有意义。
func (m *Model) X取从节点对象() *Model {
	model := m.getModel()
	model.linkType = linkTypeSlave
	return model
}

// Safe 用于标记该模型为安全或不安全。如果 safe 为 true，则在每次操作完成后都会克隆并返回一个新的模型对象；否则，它会改变当前模型的属性。
func (m *Model) X链式安全(开启 ...bool) *Model {
	if len(开启) > 0 {
		m.safe = 开启[0]
	} else {
		m.safe = true
	}
	return m
}

// Args 设置模型操作的自定义参数。
func (m *Model) 底层Args(参数 ...interface{}) *Model {
	model := m.getModel()
	model.extraArgs = append(model.extraArgs, 参数)
	return model
}

// Handler calls each of `handlers` on current Model and returns a new Model.
// ModelHandler 是一个函数，用于处理给定的 Model，并返回一个经过自定义修改的新 Model。
func (m *Model) X处理函数(处理函数 ...ModelHandler) *Model {
	model := m.getModel()
	for _, handler := range 处理函数 {
		model = handler(model)
	}
	return model
}
