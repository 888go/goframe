// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//

package db类

import (
	"context"
	"fmt"
	
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	
	"github.com/888go/goframe/net/gtrace"
)

const (
	traceInstrumentName       = "github.com/gogf/gf/v2/database/gdb"
	traceAttrDbType           = "db.type"
	traceAttrDbHost           = "db.host"
	traceAttrDbPort           = "db.port"
	traceAttrDbName           = "db.name"
	traceAttrDbUser           = "db.user"
	traceAttrDbLink           = "db.link"
	traceAttrDbGroup          = "db.group"
	traceEventDbExecution     = "db.execution"
	traceEventDbExecutionSql  = "db.execution.sql"
	traceEventDbExecutionCost = "db.execution.cost"
	traceEventDbExecutionRows = "db.execution.rows"
	traceEventDbExecutionTxID = "db.execution.txid"
	traceEventDbExecutionType = "db.execution.type"
)

// addSqlToTracing 如果启用了追踪功能，则向追踪器添加SQL信息。
func (c *Core) traceSpanEnd(ctx context.Context, span trace.Span, sql *Sql) {
	if gtrace.IsUsingDefaultProvider() || !gtrace.IsTracingInternal() {
		return
	}
	if sql.Error != nil {
		span.SetStatus(codes.Error, fmt.Sprintf(`%+v`, sql.Error))
	}
	labels := make([]attribute.KeyValue, 0)
	labels = append(labels, gtrace.CommonLabels()...)
	labels = append(labels,
		attribute.String(traceAttrDbType, c.db.X取当前节点配置().X类型),
		semconv.DBStatementKey.String(sql.Format),
	)
	if c.db.X取当前节点配置().X地址 != "" {
		labels = append(labels, attribute.String(traceAttrDbHost, c.db.X取当前节点配置().X地址))
	}
	if c.db.X取当前节点配置().X端口 != "" {
		labels = append(labels, attribute.String(traceAttrDbPort, c.db.X取当前节点配置().X端口))
	}
	if c.db.X取当前节点配置().X名称 != "" {
		labels = append(labels, attribute.String(traceAttrDbName, c.db.X取当前节点配置().X名称))
	}
	if c.db.X取当前节点配置().X账号 != "" {
		labels = append(labels, attribute.String(traceAttrDbUser, c.db.X取当前节点配置().X账号))
	}
	if filteredLink := c.db.X取Core对象().X取数据库链接信息(); filteredLink != "" {
		labels = append(labels, attribute.String(traceAttrDbLink, c.db.X取Core对象().X取数据库链接信息()))
	}
	if group := c.db.X取配置组名称(); group != "" {
		labels = append(labels, attribute.String(traceAttrDbGroup, group))
	}
	span.SetAttributes(labels...)
	events := []attribute.KeyValue{
		attribute.String(traceEventDbExecutionSql, sql.Format),
		attribute.String(traceEventDbExecutionCost, fmt.Sprintf(`%d ms`, sql.End-sql.Start)),
		attribute.String(traceEventDbExecutionRows, fmt.Sprintf(`%d`, sql.RowsAffected)),
	}
	if sql.IsTransaction {
		if v := ctx.Value(transactionIdForLoggerCtx); v != nil {
			events = append(events, attribute.String(
				traceEventDbExecutionTxID, fmt.Sprintf(`%d`, v.(uint64)),
			))
		}
	}
	events = append(events, attribute.String(traceEventDbExecutionType, sql.Type))
	span.AddEvent(traceEventDbExecution, trace.WithAttributes(events...))
}
