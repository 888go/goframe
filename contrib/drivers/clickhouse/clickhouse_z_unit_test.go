//go:build 屏蔽单元测试

// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package clickhouse

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	gdb "github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	gconv "github.com/888go/goframe/util/gconv"
	grand "github.com/888go/goframe/util/grand"
	guid "github.com/888go/goframe/util/guid"
)

const (
	sqlVisitsDDL = `
	CREATE TABLE IF NOT EXISTS visits (
	id UInt64,
	duration Float64,
	url String,
	created DateTime
	) ENGINE = MergeTree()
	PRIMARY KEY id
	ORDER BY id
`
	dimSqlDDL = `
	CREATE TABLE IF NOT EXISTS dim (
	"code" String COMMENT '编码',
	"translation" String COMMENT '译文',
	"superior" UInt64 COMMENT '上级ID',
	"row_number" UInt16 COMMENT '行号',
	"is_active" UInt8 COMMENT '是否激活',
	"is_preset" UInt8 COMMENT '是否预置',
	"category" String COMMENT '类别',
	"tree_path" Array(String) COMMENT '树路径',
	"id" UInt64 COMMENT '代理主键ID',
	"scd" UInt64 COMMENT '缓慢变化维ID',
	"version" UInt64 COMMENT 'Merge版本ID',
	"sign" Int8 COMMENT '标识位',
	"created_by" UInt64 COMMENT '创建者ID',
	"created_at" DateTime64(3,'Asia/Shanghai') COMMENT '创建时间',
	"updated_by" UInt64 COMMENT '最后修改者ID',
	"updated_at" DateTime64(3,'Asia/Shanghai') COMMENT '最后修改时间',
	"updated_tick" UInt16 COMMENT '累计修改次数'
	) ENGINE = ReplacingMergeTree("version")
	ORDER BY ("id","scd")
	COMMENT '会计准则';
`
	dimSqlDML = `
	insert into dim (code, translation, superior, row_number, is_active, is_preset, category, tree_path, id, scd, version, sign, created_by, created_at, updated_by, updated_at, updated_tick)
	values  ('CN', '{"zh_CN":"中国大陆会计准则","en_US":"Chinese mainland accounting legislation"}', 0, 1, 1, 1, 1, '[''CN'']', 607972403489804288, 0, 0, 0, 607536279118155777, '2017-09-06 00:00:00', 607536279118155777, '2017-09-06 00:00:00', 0),
			('HK', '{"zh_CN":"中国香港会计准则","en_US":"Chinese Hong Kong accounting legislation"}', 0, 2, 1, 1, 1, '[''HK'']', 607972558544834566, 0, 0, 0, 607536279118155777, '2017-09-06 00:00:00', 607536279118155777, '2017-09-06 00:00:00', 0);
`
	factSqlDDL = `
	CREATE TABLE IF NOT EXISTS fact (
	"adjustment_level" UInt64 COMMENT '调整层ID',
	"data_version" UInt64 COMMENT '数据版本ID',
	"accounting_legislation" UInt64 COMMENT '会计准则ID',
	"fiscal_year" UInt16 COMMENT '会计年度',
	"fiscal_period" UInt8 COMMENT '会计期间',
	"fiscal_year_period" UInt32 COMMENT '会计年度期间',
	"legal_entity" UInt64 COMMENT '法人主体ID',
	"cost_center" UInt64 COMMENT '成本中心ID',
	"legal_entity_partner" UInt64 COMMENT '内部关联方ID',
	"financial_posting" UInt64 COMMENT '凭证头ID',
	"line" UInt16 COMMENT '行号',
	"general_ledger_account" UInt64 COMMENT '总账科目ID',
	"debit" Decimal64(9) COMMENT '借方金额',
	"credit" Decimal64(9) COMMENT '贷方金额',
	"transaction_currency" UInt64 COMMENT '交易币种ID',
	"debit_tc" Decimal64(9) COMMENT '借方金额（交易币种）',
	"credit_tc" Decimal64(9) COMMENT '贷方金额（交易币种）',
	"posting_date" Date32 COMMENT '过账日期',
	"gc_year" UInt16 COMMENT '公历年',
	"gc_quarter" UInt8 COMMENT '公历季',
	"gc_month" UInt8 COMMENT '公历月',
	"gc_week" UInt8 COMMENT '公历周',
	"raw_info" String COMMENT '源信息',
	"summary" String COMMENT '摘要',
	"id" UInt64 COMMENT '代理主键ID',
	"version" UInt64 COMMENT 'Merge版本ID',
	"sign" Int8 COMMENT '标识位'
	) ENGINE = ReplacingMergeTree("version")
	ORDER BY ("adjustment_level","data_version","legal_entity","fiscal_year","fiscal_period","financial_posting","line")
	PARTITION BY ("adjustment_level","data_version","legal_entity","fiscal_year","fiscal_period")
	COMMENT '数据主表';
`
	factSqlDML = `
	insert into fact (adjustment_level, data_version, accounting_legislation, fiscal_year, fiscal_period, fiscal_year_period, legal_entity, cost_center, legal_entity_partner, financial_posting, line, general_ledger_account, debit, credit, transaction_currency, debit_tc, credit_tc, posting_date, gc_year, gc_quarter, gc_month, gc_week, raw_info, summary, id, version, sign)
	values  (607970943242866688, 607973669943119880, 607972403489804288, 2022, 3, 202203, 607974511316307985, 0, 607976190010986520, 607996702456025136, 1, 607985607569838111, 8674.39, 0, 607974898261823505, 8674.39, 0, '2022-03-05', 2022, 1, 3, 11, '{}', '摘要', 607992882741121073, 0, 0),
			(607970943242866688, 607973669943119880, 607972403489804288, 2022, 4, 202204, 607974511316307985, 0, 607976190010986520, 607993586419503145, 1, 607985607569838111, 9999.88, 0, 607974898261823505, 9999.88, 0, '2022-04-10', 2022, 2, 4, 18, '{}', '摘要', 607996939140599857, 0, 0);
`
	expmSqlDDL = `
	CREATE TABLE IF NOT EXISTS data_type (
		  Col1 UInt8 COMMENT '列1'
		, Col2 Nullable(String) COMMENT '列2'
		, Col3 FixedString(3) COMMENT '列3'
		, Col4 String COMMENT '列4'
		, Col5 Map(String, UInt8) COMMENT '列5'
		, Col6 Array(String) COMMENT '列6'
		, Col7 Tuple(String, UInt8, Array(Map(String, String))) COMMENT '列7'
		, Col8 DateTime COMMENT '列8'
		, Col9 UUID COMMENT '列9'
		, Col10 DateTime COMMENT '列10'
		, Col11 Decimal(9, 2) COMMENT '列11'
		, Col12 Decimal(9, 2) COMMENT '列12'
	) ENGINE = MergeTree()
	PRIMARY KEY Col4
	ORDER BY Col4
`
)

func clickhouseConfigDB() gdb.DB {
	connect, err := gdb.X创建DB对象(gdb.ConfigNode{
		Host:  "127.0.0.1",
		Port:  "9000",
		User:  "default",
		Name:  "default",
		Type:  "clickhouse",
		Debug: false,
	})
	gtest.AssertNil(err)
	gtest.AssertNE(connect, nil)
	return connect
}

func clickhouseLink() gdb.DB {
	connect, err := gdb.X创建DB对象(gdb.ConfigNode{
		Link: "clickhouse:default:@tcp(127.0.0.1:9000)/default?dial_timeout=200ms&max_execution_time=60",
	})
	gtest.AssertNil(err)
	gtest.AssertNE(connect, nil)
	return connect
}

func createClickhouseTableVisits(connect gdb.DB) error {
	_, err := connect.X原生SQL执行(context.Background(), sqlVisitsDDL)
	return err
}

func createClickhouseTableDim(connect gdb.DB) error {
	_, err := connect.X原生SQL执行(context.Background(), dimSqlDDL)
	return err
}

func createClickhouseTableFact(connect gdb.DB) error {
	_, err := connect.X原生SQL执行(context.Background(), factSqlDDL)
	return err
}

func createClickhouseExampleTable(connect gdb.DB) error {
	_, err := connect.X原生SQL执行(context.Background(), expmSqlDDL)
	return err
}

func dropClickhouseTableVisits(conn gdb.DB) {
	sqlStr := fmt.Sprintf("DROP TABLE IF EXISTS `visits`")
	_, _ = conn.X原生SQL执行(context.Background(), sqlStr)
}

func dropClickhouseTableDim(conn gdb.DB) {
	sqlStr := fmt.Sprintf("DROP TABLE IF EXISTS `dim`")
	_, _ = conn.X原生SQL执行(context.Background(), sqlStr)
}

func dropClickhouseTableFact(conn gdb.DB) {
	sqlStr := fmt.Sprintf("DROP TABLE IF EXISTS `fact`")
	_, _ = conn.X原生SQL执行(context.Background(), sqlStr)
}

func dropClickhouseExampleTable(conn gdb.DB) {
	sqlStr := fmt.Sprintf("DROP TABLE IF EXISTS `data_type`")
	_, _ = conn.X原生SQL执行(context.Background(), sqlStr)
}

func TestDriverClickhouse_Create(t *testing.T) {
	gtest.AssertNil(createClickhouseTableVisits(clickhouseConfigDB()))
}

func TestDriverClickhouse_New(t *testing.T) {
	connect := clickhouseConfigDB()
	gtest.AssertNE(connect, nil)
	gtest.AssertNil(connect.X向主节点发送心跳())
	gtest.AssertNil(connect.X向从节点发送心跳())
}

func TestDriverClickhouse_OpenLink_Ping(t *testing.T) {
	connect := clickhouseConfigDB()
	gtest.AssertNE(connect, nil)
	gtest.AssertNil(connect.X向主节点发送心跳())
}

func TestDriverClickhouse_Tables(t *testing.T) {
	connect := clickhouseConfigDB()
	gtest.AssertEQ(createClickhouseTableVisits(connect), nil)
	defer dropClickhouseTableVisits(connect)
	tables, err := connect.X取表名称切片(context.Background())
	gtest.AssertNil(err)
	gtest.AssertNE(len(tables), 0)
}

func TestDriverClickhouse_TableFields_Use_Config(t *testing.T) {
	connect := clickhouseConfigDB()
	gtest.AssertNil(createClickhouseTableVisits(connect))
	defer dropClickhouseTableVisits(connect)
	field, err := connect.X取表字段信息Map(context.Background(), "visits")
	gtest.AssertNil(err)
	gtest.AssertEQ(len(field), 4)
	gtest.AssertNQ(field, nil)
}

func TestDriverClickhouse_TableFields_Use_Link(t *testing.T) {
	connect := clickhouseLink()
	gtest.AssertNil(createClickhouseTableVisits(connect))
	defer dropClickhouseTableVisits(connect)
	field, err := connect.X取表字段信息Map(context.Background(), "visits")
	gtest.AssertNil(err)
	gtest.AssertEQ(len(field), 4)
	gtest.AssertNQ(field, nil)
}

func TestDriverClickhouse_Transaction(t *testing.T) {
	connect := clickhouseConfigDB()
	defer dropClickhouseTableVisits(connect)
	gtest.AssertNE(connect.X事务(context.Background(), func(ctx context.Context, tx gdb.TX) error {
		return nil
	}), nil)
}

func TestDriverClickhouse_InsertIgnore(t *testing.T) {
	connect := clickhouseConfigDB()
	_, err := connect.X插入并跳过已存在(context.Background(), "", nil)
	gtest.AssertEQ(err, errUnsupportedInsertIgnore)
}

func TestDriverClickhouse_InsertAndGetId(t *testing.T) {
	connect := clickhouseConfigDB()
	_, err := connect.X插入并取ID(context.Background(), "", nil)
	gtest.AssertEQ(err, errUnsupportedInsertGetId)
}

func TestDriverClickhouse_InsertOne(t *testing.T) {
	connect := clickhouseConfigDB()
	gtest.AssertEQ(createClickhouseTableVisits(connect), nil)
	defer dropClickhouseTableVisits(connect)
	_, err := connect.X创建Model对象("visits").X设置数据(g.Map{
		"duration": float64(grand.X整数(999)),
		"url":      gconv.String(grand.X整数(999)),
		"created":  time.Now(),
	}).X插入()
	gtest.AssertNil(err)
}

func TestDriverClickhouse_InsertOneAutoDateTimeWrite(t *testing.T) {
	connect, err := gdb.X创建DB对象(gdb.ConfigNode{
		Host:      "127.0.0.1",
		Port:      "9000",
		User:      "default",
		Name:      "default",
		Type:      "clickhouse",
		Debug:     false,
		CreatedAt: "created",
	})
	gtest.AssertNil(err)
	gtest.AssertNE(connect, nil)
	gtest.AssertEQ(createClickhouseTableVisits(connect), nil)
	defer dropClickhouseTableVisits(connect)
	beforeInsertTime := time.Now()
	_, err = connect.X创建Model对象("visits").X设置数据(g.Map{
		"duration": float64(grand.X整数(999)),
		"url":      gconv.String(grand.X整数(999)),
	}).X插入()
	gtest.AssertNil(err)
		// 查询插入的数据以获取时间字段的值. md5:7044ac81cbd8f93f
	data, err := connect.X创建Model对象("visits").X查询一条()
	gtest.AssertNil(err)
		// 从插入的数据中获取时间值. md5:f519e5efebe5e810
	createdTime := data["created"].X取时间类()
			// 断言时间字段的值等于或晚于beforeInsertTime. md5:1ec7235b45d129a3
	gtest.AssertGE(createdTime.Unix(), beforeInsertTime.Unix())
}

func TestDriverClickhouse_InsertMany(t *testing.T) {
	connect := clickhouseConfigDB()
	gtest.AssertEQ(createClickhouseTableVisits(connect), nil)
	defer dropClickhouseTableVisits(connect)
	tx, err := connect.X事务开启(context.Background())
	gtest.AssertEQ(err, errUnsupportedBegin)
	gtest.AssertNil(tx)
}

func TestDriverClickhouse_Insert(t *testing.T) {
	connect := clickhouseConfigDB()
	gtest.AssertEQ(createClickhouseTableVisits(connect), nil)
	defer dropClickhouseTableVisits(connect)
	type insertItem struct {
		Id       uint64    `orm:"id"`
		Duration float64   `orm:"duration"`
		Url      string    `orm:"url"`
		Created  time.Time `orm:"created"`
	}
	var (
		insertUrl = "https://goframe.org"
		total     = 0
		item      = insertItem{
			Duration: 1,
			Url:      insertUrl,
			Created:  time.Now(),
		}
	)
	_, err := connect.X创建Model对象("visits").X设置数据(item).X插入()
	gtest.AssertNil(err)
	_, err = connect.X创建Model对象("visits").X设置数据(item).X插入并更新已存在()
	gtest.AssertNil(err)
	total, err = connect.X创建Model对象("visits").X查询行数()
	gtest.AssertNil(err)
	gtest.AssertEQ(total, 2)
	var list []*insertItem
	for i := 0; i < 50; i++ {
		list = append(list, &insertItem{
			Duration: float64(grand.X整数(999)),
			Url:      insertUrl,
			Created:  time.Now(),
		})
	}
	_, err = connect.X创建Model对象("visits").X设置数据(list).X插入()
	gtest.AssertNil(err)
	_, err = connect.X创建Model对象("visits").X设置数据(list).X插入并更新已存在()
	gtest.AssertNil(err)
	total, err = connect.X创建Model对象("visits").X查询行数()
	gtest.AssertNil(err)
	gtest.AssertEQ(total, 102)
}

func TestDriverClickhouse_Insert_Use_Exec(t *testing.T) {
	connect := clickhouseConfigDB()
	gtest.AssertEQ(createClickhouseTableFact(connect), nil)
	defer dropClickhouseTableFact(connect)
	_, err := connect.X原生SQL执行(context.Background(), factSqlDML)
	gtest.AssertNil(err)
}

func TestDriverClickhouse_Delete(t *testing.T) {
	connect := clickhouseConfigDB()
	gtest.AssertEQ(createClickhouseTableVisits(connect), nil)
	defer dropClickhouseTableVisits(connect)
	_, err := connect.X创建Model对象("visits").X条件("created >", "2021-01-01 00:00:00").X删除()
	gtest.AssertNil(err)
	_, err = connect.X创建Model对象("visits").
		X条件("created >", "2021-01-01 00:00:00").
		X条件("duration > ", 0).
		X条件("url is not null").
		X删除()
	gtest.AssertNil(err)
}

func TestDriverClickhouse_Update(t *testing.T) {
	connect := clickhouseConfigDB()
	gtest.AssertEQ(createClickhouseTableVisits(connect), nil)
	defer dropClickhouseTableVisits(connect)
	_, err := connect.X创建Model对象("visits").X条件("created > ", "2021-01-01 15:15:15").X设置数据(g.Map{
		"created": time.Now().Format("2006-01-02 15:04:05"),
	}).X更新()
	gtest.AssertNil(err)
	_, err = connect.X创建Model对象("visits").
		X条件("created > ", "2021-01-01 15:15:15").
		X条件("duration > ", 0).
		X条件("url is not null").
		X设置数据(g.Map{
			"created": time.Now().Format("2006-01-02 15:04:05"),
		}).X更新()
}

func TestDriverClickhouse_Replace(t *testing.T) {
	connect := clickhouseConfigDB()
	_, err := connect.X插入并替换已存在(context.Background(), "", nil)
	gtest.AssertEQ(err, errUnsupportedReplace)
}

func TestDriverClickhouse_DoFilter(t *testing.T) {
	rawSQL := "select * from visits where 1 = 1"
	this := Driver{}
	replaceSQL, _, err := this.X底层DoFilter(context.Background(), nil, rawSQL, []interface{}{1})
	gtest.AssertNil(err)
	gtest.AssertEQ(rawSQL, replaceSQL)

		// 此SQL无法运行，因为没有WHERE子句，Clickhouse会报告错误. md5:50770b7fc72b157f
	rawSQL = "update visit set url = '1'"
	replaceSQL, _, err = this.X底层DoFilter(context.Background(), nil, rawSQL, []interface{}{1})
	gtest.AssertNil(err)

		// 此SQL无法运行，因为没有WHERE子句，Clickhouse会报告错误. md5:50770b7fc72b157f
	rawSQL = "delete from visit"
	replaceSQL, _, err = this.X底层DoFilter(context.Background(), nil, rawSQL, []interface{}{1})
	gtest.AssertNil(err)

	ctx := this.injectNeedParsedSql(context.Background())
	rawSQL = "UPDATE visit SET url = '1' WHERE url = '0'"
	replaceSQL, _, err = this.X底层DoFilter(ctx, nil, rawSQL, []interface{}{1})
	gtest.AssertNil(err)
	gtest.AssertEQ(replaceSQL, "ALTER TABLE visit UPDATE url = '1' WHERE url = '0'")

	rawSQL = "DELETE FROM visit WHERE url = '0'"
	replaceSQL, _, err = this.X底层DoFilter(ctx, nil, rawSQL, []interface{}{1})
	gtest.AssertNil(err)
	gtest.AssertEQ(replaceSQL, "ALTER TABLE visit DELETE WHERE url = '0'")
}

func TestDriverClickhouse_Select(t *testing.T) {
	connect := clickhouseConfigDB()
	gtest.AssertNil(createClickhouseTableVisits(connect))
	defer dropClickhouseTableVisits(connect)
	_, err := connect.X创建Model对象("visits").X设置数据(g.Map{
		"url":      "goframe.org",
		"duration": float64(1),
	}).X插入()
	gtest.AssertNil(err)
	temp, err := connect.X创建Model对象("visits").X条件("url", "goframe.org").X条件("duration >= ", 1).X查询一条()
	gtest.AssertNil(err)
	gtest.AssertEQ(temp.X是否为空(), false)
	_, err = connect.X创建Model对象("visits").X设置数据(g.Map{
		"url":      "goframe.org",
		"duration": float64(2),
	}).X插入()
	gtest.AssertNil(err)
	data, err := connect.X创建Model对象("visits").X条件("url", "goframe.org").X条件("duration >= ", 1).X查询()
	gtest.AssertNil(err)
	gtest.AssertEQ(len(data), 2)
}

func TestDriverClickhouse_Exec_OPTIMIZE(t *testing.T) {
	connect := clickhouseConfigDB()
	gtest.AssertNil(createClickhouseTableVisits(connect))
	defer dropClickhouseTableVisits(connect)
	sqlStr := "OPTIMIZE table visits"
	_, err := connect.X原生SQL执行(context.Background(), sqlStr)
	gtest.AssertNil(err)
}

func TestDriverClickhouse_ExecInsert(t *testing.T) {
	connect := clickhouseConfigDB()
	gtest.AssertEQ(createClickhouseTableDim(connect), nil)
	defer dropClickhouseTableDim(connect)
	_, err := connect.X原生SQL执行(context.Background(), dimSqlDML)
	gtest.AssertNil(err)
}

func TestDriverClickhouse_NilTime(t *testing.T) {
	connect := clickhouseConfigDB()
	gtest.AssertNil(createClickhouseExampleTable(connect))
	defer dropClickhouseExampleTable(connect)
	type testNilTime struct {
		Col1  uint8
		Col2  string
		Col3  string
		Col4  string
		Col5  map[string]uint8
		Col6  []string
		Col7  []interface{}
		Col8  *time.Time
		Col9  uuid.UUID
		Col10 *gtime.Time
		Col11 decimal.Decimal
		Col12 *decimal.Decimal
	}
	insertData := []*testNilTime{}
	money := decimal.NewFromFloat(1.12)
	strMoney, _ := decimal.NewFromString("99999.999")
	for i := 0; i < 10000; i++ {
		insertData = append(insertData, &testNilTime{
			Col4: "Inc.",
			Col9: uuid.New(),
			Col7: []interface{}{ // 定义了一个元组，包含三个元素：一个字符串（String）、一个无符号八位整数（UInt8）和一个数组，该数组的每个元素都是一个映射（Map），映射的键和值都是字符串（String）。 md5:21034fd2e2a7f9b3
				"String Value", uint8(5), []map[string]string{
					{"key": "value"},
					{"key": "value"},
					{"key": "value"},
				}},
			Col11: money,
			Col12: &strMoney,
		})
	}
	_, err := connect.X创建Model对象("data_type").X设置数据(insertData).X插入()
	gtest.AssertNil(err)
	count, err := connect.X创建Model对象("data_type").X条件("Col4", "Inc.").X查询行数()
	gtest.AssertNil(err)
	gtest.AssertEQ(count, 10000)

	data, err := connect.X创建Model对象("data_type").X条件("Col4", "Inc.").X查询一条()
	gtest.AssertNil(err)
	gtest.AssertNE(data, nil)
	g.X调试输出(data)
	gtest.AssertEQ(data["Col11"].String(), "1.12")
	gtest.AssertEQ(data["Col12"].String(), "99999.99")
}

func TestDriverClickhouse_BatchInsert(t *testing.T) {
	// 这是来自
	// https:	//github.com/ClickHouse/clickhouse-go/blob/v2/examples/std/batch/main.go 示例的注释
	// md5:c0ed3f953d0aaab1
	connect := clickhouseConfigDB()
	gtest.AssertNil(createClickhouseExampleTable(connect))
	defer dropClickhouseExampleTable(connect)
	insertData := []g.Map{}
	for i := 0; i < 10000; i++ {
		insertData = append(insertData, g.Map{
			"Col1": uint8(42),
			"Col2": "ClickHouse",
			"Col3": "Inc",
			"Col4": guid.X生成(),
			"Col5": map[string]uint8{"key": 1},             // Map(String, UInt8)
			"Col6": []string{"Q", "W", "E", "R", "T", "Y"}, // Array(String)
			"Col7": []interface{}{ // 定义了一个元组，包含三个元素：一个字符串（String）、一个无符号八位整数（UInt8）和一个数组，该数组的每个元素都是一个映射（Map），映射的键和值都是字符串（String）。 md5:21034fd2e2a7f9b3
				"String Value", uint8(5), []map[string]string{
					{"key": "value"},
					{"key": "value"},
					{"key": "value"},
				},
			},
			"Col8":  gtime.X创建并按当前时间(),
			"Col9":  uuid.New(),
			"Col10": nil,
		})
	}
	_, err := connect.X创建Model对象("data_type").X设置数据(insertData).X插入()
	gtest.AssertNil(err)
	count, err := connect.X创建Model对象("data_type").X条件("Col2", "ClickHouse").X条件("Col3", "Inc").X查询行数()
	gtest.AssertNil(err)
	gtest.AssertEQ(count, 10000)
}

func TestDriverClickhouse_Open(t *testing.T) {
	// 链接
	// 数据源管理(DSM)
	// 点击house协议连接字符串：	//用户名:密码@主机1:9000,主机2:9000/数据库名?拨号超时=200毫秒&最大执行时间=60秒
	// md5:c1d7a1212d7e0483
	link := "clickhouse://default@127.0.0.1:9000,127.0.0.1:9000/default?dial_timeout=200ms&max_execution_time=60"
	db, err := gdb.X创建DB对象(gdb.ConfigNode{
		Link: link,
		Type: "clickhouse",
	})
	gtest.AssertNil(err)
	gtest.AssertNil(db.X向主节点发送心跳())
}

func TestDriverClickhouse_TableFields(t *testing.T) {
	connect := clickhouseConfigDB()
	gtest.AssertNil(createClickhouseExampleTable(connect))
	defer dropClickhouseExampleTable(connect)
	dataTypeTable, err := connect.X取表字段信息Map(context.Background(), "data_type")
	gtest.AssertNil(err)
	gtest.AssertNE(dataTypeTable, nil)

	var result = map[string][]interface{}{
		"Col1":  {0, "Col1", "UInt8", false, "", "", "", "列1"},
		"Col2":  {1, "Col2", "String", true, "", "", "", "列2"},
		"Col3":  {2, "Col3", "FixedString(3)", false, "", "", "", "列3"},
		"Col4":  {3, "Col4", "String", false, "", "", "", "列4"},
		"Col5":  {4, "Col5", "Map(String, UInt8)", false, "", "", "", "列5"},
		"Col6":  {5, "Col6", "Array(String)", false, "", "", "", "列6"},
		"Col7":  {6, "Col7", "Tuple(String, UInt8, Array(Map(String, String)))", false, "", "", "", "列7"},
		"Col8":  {7, "Col8", "DateTime", false, "", "", "", "列8"},
		"Col9":  {8, "Col9", "UUID", false, "", "", "", "列9"},
		"Col10": {9, "Col10", "DateTime", false, "", "", "", "列10"},
		"Col11": {10, "Col11", "Decimal(9, 2)", false, "", "", "", "列11"},
		"Col12": {11, "Col12", "Decimal(9, 2)", false, "", "", "", "列12"},
	}
	for k, v := range result {
		_, ok := dataTypeTable[k]
		gtest.AssertEQ(ok, true)
		gtest.AssertEQ(dataTypeTable[k].Index, v[0])
		gtest.AssertEQ(dataTypeTable[k].X名称, v[1])
		gtest.AssertEQ(dataTypeTable[k].X类型, v[2])
		gtest.AssertEQ(dataTypeTable[k].Null, v[3])
		gtest.AssertEQ(dataTypeTable[k].Key, v[4])
		gtest.AssertEQ(dataTypeTable[k].Default, v[5])
		gtest.AssertEQ(dataTypeTable[k].Comment, v[7])
	}
}

func TestDriverClickhouse_TableFields_HasField(t *testing.T) {
	connect := clickhouseConfigDB()
	gtest.AssertNil(createClickhouseExampleTable(connect))
	defer dropClickhouseExampleTable(connect)
	// 未修复前：panic: runtime error: index out of range [12] with length 12
	b, err := connect.X取Core对象().X是否存在字段(context.Background(), "data_type", "Col1")
	gtest.AssertNil(err)
	gtest.AssertEQ(b, true)
}
