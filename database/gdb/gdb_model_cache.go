// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gdb

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/internal/intlog"
)

// CacheOption是查询中用于模型缓存控制的选项。 md5:8a833b8335c45455
type CacheOption struct {
	// Duration 是缓存的过期时间。
	// 如果参数 `Duration` 小于 0，表示使用给定的 `Name` 清除缓存。
	// 如果参数 `Duration` 等于 0，表示永不过期。
	// 如果参数 `Duration` 大于 0，表示在 `Duration` 秒后过期。
	// md5:28707300732ac411
	Duration time.Duration

	// Name 是一个可选的唯一名称，用于标识缓存。
	// 通过 Name 可以将一个名称与缓存绑定，这意味着您之后可以根据指定的名称来控制该缓存，
	// 例如更改缓存的 `持续时间` 或者清除指定名称的缓存。
	// md5:8c2eeafa42d36067
	Name string

	// 强制缓存查询结果，无论结果是否为nil。
	// 这用于防止缓存穿透。
	// md5:78fc7d8520d64954
	Force bool
}

// selectCacheItem是用于SELECT语句结果的缓存项。 md5:73fb34eaa64ea7d1
type selectCacheItem struct {
	Result            Result // SELECT语句的SQL结果。 md5:1f098617a374fffc
	FirstResultColumn string // 结果的第一列名称，用于Value/Count函数。 md5:2c091aca88ae5aa3
}

// Cache 为模型设置缓存功能。它将 SQL 的结果缓存起来，这意味着如果有相同的 SQL 请求，
// 它会直接从缓存中读取并返回结果，而不会真正提交并执行到数据库中。
//
// 注意，如果模型在事务中执行 SELECT 语句，缓存功能将被禁用。
// md5:5d7ea513a485f3ad
func (m *Model) Cache(option CacheOption) *Model {
	model := m.getModel()
	model.cacheOption = option
	model.cacheEnabled = true
	return model
}

// checkAndRemoveSelectCache 检查并移除插入/更新/删除语句中的缓存，如果启用了缓存功能。
// md5:7247a2e1e2e19e4b
func (m *Model) checkAndRemoveSelectCache(ctx context.Context) {
	if m.cacheEnabled && m.cacheOption.Duration < 0 && len(m.cacheOption.Name) > 0 {
		var cacheKey = m.makeSelectCacheKey("")
		if _, err := m.db.GetCache().Remove(ctx, cacheKey); err != nil {
			intlog.Errorf(ctx, `%+v`, err)
		}
	}
}

func (m *Model) getSelectResultFromCache(ctx context.Context, sql string, args ...interface{}) (result Result, err error) {
	if !m.cacheEnabled || m.tx != nil {
		return
	}
	var (
		cacheItem *selectCacheItem
		cacheKey  = m.makeSelectCacheKey(sql, args...)
		cacheObj  = m.db.GetCache()
		core      = m.db.GetCore()
	)
	defer func() {
		if cacheItem != nil {
			if internalData := core.getInternalColumnFromCtx(ctx); internalData != nil {
				if cacheItem.FirstResultColumn != "" {
					internalData.FirstResultColumn = cacheItem.FirstResultColumn
				}
			}
		}
	}()
	if v, _ := cacheObj.Get(ctx, cacheKey); !v.IsNil() {
		if err = v.Scan(&cacheItem); err != nil {
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
		cacheObj = m.db.GetCache()
	)
	if m.cacheOption.Duration < 0 {
		if _, errCache := cacheObj.Remove(ctx, cacheKey); errCache != nil {
			intlog.Errorf(ctx, `%+v`, errCache)
		}
		return
	}
		// 对Value/Count操作结果的特殊处理器。 md5:beba69dc2347fa3a
	if len(result) > 0 {
		var core = m.db.GetCore()
		switch queryType {
		case queryTypeValue, queryTypeCount:
			if internalData := core.getInternalColumnFromCtx(ctx); internalData != nil {
				if result[0][internalData.FirstResultColumn].IsEmpty() {
					result = nil
				}
			}
		}
	}

		// 针对缓存穿透的情况。 md5:1464372279e61a7d
	if result.IsEmpty() {
		if m.cacheOption.Force {
			result = Result{}
		} else {
			result = nil
		}
	}
	var (
		core      = m.db.GetCore()
		cacheItem = &selectCacheItem{
			Result: result,
		}
	)
	if internalData := core.getInternalColumnFromCtx(ctx); internalData != nil {
		cacheItem.FirstResultColumn = internalData.FirstResultColumn
	}
	if errCache := cacheObj.Set(ctx, cacheKey, cacheItem, m.cacheOption.Duration); errCache != nil {
		intlog.Errorf(ctx, `%+v`, errCache)
	}
	return
}

func (m *Model) makeSelectCacheKey(sql string, args ...interface{}) string {
	return m.db.GetCore().makeSelectCacheKey(
		m.cacheOption.Name,
		m.db.GetSchema(),
		m.db.GetCore().guessPrimaryTableName(m.tables),
		sql,
		args...,
	)
}
