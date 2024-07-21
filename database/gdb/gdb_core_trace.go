// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

package gdb

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/gogf/gf/v2/net/gtrace"
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

// addSqlToTracing 如果启用追踪，则将SQL信息添加到追踪器中。 md5:a591b35a0635d132
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
		attribute.String(traceAttrDbType, c.db.GetConfig().Type),
		semconv.DBStatement(sql.Format),
	)
	if c.db.GetConfig().Host != "" {
		labels = append(labels, attribute.String(traceAttrDbHost, c.db.GetConfig().Host))
	}
	if c.db.GetConfig().Port != "" {
		labels = append(labels, attribute.String(traceAttrDbPort, c.db.GetConfig().Port))
	}
	if c.db.GetConfig().Name != "" {
		labels = append(labels, attribute.String(traceAttrDbName, c.db.GetConfig().Name))
	}
	if c.db.GetConfig().User != "" {
		labels = append(labels, attribute.String(traceAttrDbUser, c.db.GetConfig().User))
	}
	if filteredLink := c.db.GetCore().FilteredLink(); filteredLink != "" {
		labels = append(labels, attribute.String(traceAttrDbLink, c.db.GetCore().FilteredLink()))
	}
	if group := c.db.GetGroup(); group != "" {
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
	events = append(events, attribute.String(traceEventDbExecutionType, string(sql.Type)))
	span.AddEvent(traceEventDbExecution, trace.WithAttributes(events...))
}
