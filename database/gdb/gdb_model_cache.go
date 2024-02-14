// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

import (
	"context"
	"time"
	
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/json"
)

// CacheOption 是用于查询中模型缓存控制的选项。
type X缓存选项 struct {
// Duration 是缓存的生存时间（TTL）。
// 如果参数 `Duration` < 0，表示按照给定的 `Name` 清除缓存。
// 如果参数 `Duration` = 0，表示缓存永不过期。
// 如果参数 `Duration` > 0，表示在 `Duration` 时间后缓存过期。
	X时长 time.Duration

// Name 是缓存的一个可选的唯一名称。
// 名称用于将名称绑定到缓存，这意味着您可以在之后通过名称控制缓存，
// 例如：更改 `duration` 或清除指定名称的缓存。
	X名称 string

// Force无论查询结果是否为nil，都会缓存该查询结果。
// 它用于避免缓存穿透。
	X强制缓存 bool
}

// selectCacheItem 是用于 SELECT 语句结果的缓存项。
type selectCacheItem struct {
	Result            Result // Sql result of SELECT statement. （SQL语句中SELECT查询的结果。）
	FirstResultColumn string // 结果中的第一列名称，用于Value/Count函数。
}

// Cache 为模型设置缓存功能。它会缓存SQL查询的结果，这意味着
// 如果存在相同的SQL请求，它将直接从缓存读取并返回结果，
// 而不是提交并执行到数据库中。
//
// 注意，如果模型在事务中执行选择语句时，缓存功能是禁用的。
func (m *Model) X缓存(选项 X缓存选项) *Model {
	model := m.getModel()
	model.cacheOption = 选项
	model.cacheEnabled = true
	return model
}

// checkAndRemoveSelectCache 在缓存功能启用的情况下，检查并移除在插入/更新/删除语句中的缓存。
func (m *Model) checkAndRemoveSelectCache(ctx context.Context) {
	if m.cacheEnabled && m.cacheOption.X时长 < 0 && len(m.cacheOption.X名称) > 0 {
		var cacheKey = m.makeSelectCacheKey("")
		if _, err := m.db.X取缓存对象().X删除并带返回值(ctx, cacheKey); err != nil {
			intlog.Errorf(ctx, `%+v`, err)
		}
	}
}

func (m *Model) getSelectResultFromCache(ctx context.Context, sql string, args ...interface{}) (result Result, err error) {
	if !m.cacheEnabled || m.tx != nil {
		return
	}
	var (
		ok        bool
		cacheItem *selectCacheItem
		cacheKey  = m.makeSelectCacheKey(sql, args...)
		cacheObj  = m.db.X取缓存对象()
	)
	defer func() {
		if cacheItem != nil {
			if internalData := m.db.X取Core对象().底层_GetInternalCtxDataFromCtx(ctx); internalData != nil {
				if cacheItem.FirstResultColumn != "" {
					internalData.FirstResultColumn = cacheItem.FirstResultColumn
				}
			}
		}
	}()
	if v, _ := cacheObj.X取值(ctx, cacheKey); !v.X是否为Nil() {
		if cacheItem, ok = v.X取值().(*selectCacheItem); ok {
			// In-memory cache.
			return cacheItem.Result, nil
		}
		// 其他缓存，需要进行转换。
		if err = json.UnmarshalUseNumber(v.X取字节集(), &cacheItem); err != nil {
			return nil, err
		}
		return cacheItem.Result, nil
	}
	return
}

func (m *Model) saveSelectResultToCache(
	ctx context.Context, queryType queryType, result Result, sql string, args ...interface{},
) (err error) {
	if !m.cacheEnabled || m.tx != nil {
		return
	}
	var (
		cacheKey = m.makeSelectCacheKey(sql, args...)
		cacheObj = m.db.X取缓存对象()
	)
	if m.cacheOption.X时长 < 0 {
		if _, errCache := cacheObj.X删除并带返回值(ctx, cacheKey); errCache != nil {
			intlog.Errorf(ctx, `%+v`, errCache)
		}
		return
	}
	// 特殊处理 Value/Count 操作结果的处理器。
	if len(result) > 0 {
		switch queryType {
		case queryTypeValue, queryTypeCount:
			if internalData := m.db.X取Core对象().底层_GetInternalCtxDataFromCtx(ctx); internalData != nil {
				if result[0][internalData.FirstResultColumn].X是否为空() {
					result = nil
				}
			}
		}
	}

	// 在发生缓存穿透的情况下。
	if result.X是否为空() {
		if m.cacheOption.X强制缓存 {
			result = Result{}
		} else {
			result = nil
		}
	}
	var cacheItem = &selectCacheItem{
		Result: result,
	}
	if internalData := m.db.X取Core对象().底层_GetInternalCtxDataFromCtx(ctx); internalData != nil {
		cacheItem.FirstResultColumn = internalData.FirstResultColumn
	}
	if errCache := cacheObj.X设置值(ctx, cacheKey, cacheItem, m.cacheOption.X时长); errCache != nil {
		intlog.Errorf(ctx, `%+v`, errCache)
	}
	return
}

func (m *Model) makeSelectCacheKey(sql string, args ...interface{}) string {
	return m.db.X取Core对象().makeSelectCacheKey(
		m.cacheOption.X名称,
		m.db.X取默认数据库名称(),
		m.db.X取Core对象().guessPrimaryTableName(m.tables),
		sql,
		args...,
	)
}
