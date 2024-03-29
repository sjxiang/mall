// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	orderFieldNames          = builder.RawFieldNames(&Order{})
	orderRows                = strings.Join(orderFieldNames, ",")
	orderRowsExpectAutoSet   = strings.Join(stringx.Remove(orderFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	orderRowsWithPlaceHolder = strings.Join(stringx.Remove(orderFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheOrderIdPrefix = "cache:order:id:"
)

type (
	orderModel interface {
		Insert(ctx context.Context, data *Order) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Order, error)
		Update(ctx context.Context, data *Order) error
		Delete(ctx context.Context, id int64) error
	}

	defaultOrderModel struct {
		sqlc.CachedConn
		table string
	}

	Order struct {
		Id         int64     `db:"id"`          // 主键
		CreateAt   time.Time `db:"create_at"`   // 创建时间
		CreateBy   string    `db:"create_by"`   // 创建者
		UpdateAt   time.Time `db:"update_at"`   // 更新时间
		UpdateBy   string    `db:"update_by"`   // 更新者
		Version    int64     `db:"version"`     // 乐观锁版本号
		IsDel      int64     `db:"is_del"`      // 是否删除：0正常1删除
		UserId     int64     `db:"user_id"`     // 用户id
		OrderId    int64     `db:"order_id"`    // 订单id
		TradeId    string    `db:"trade_id"`    // 交易单号
		PayChannel int64     `db:"pay_channel"` // 支付方式
		Status     int64     `db:"status"`      // 订单状态:100创建订单/待支付 200已支付 300交易关闭 400完成
		PayAmount  int64     `db:"pay_amount"`  // 支付金额（分）
		PayTime    time.Time `db:"pay_time"`    // 支付时间
	}
)

func newOrderModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultOrderModel {
	return &defaultOrderModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`order`",
	}
}

func (m *defaultOrderModel) Delete(ctx context.Context, id int64) error {
	orderIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, orderIdKey)
	return err
}

func (m *defaultOrderModel) FindOne(ctx context.Context, id int64) (*Order, error) {
	orderIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, id)
	var resp Order
	err := m.QueryRowCtx(ctx, &resp, orderIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOrderModel) Insert(ctx context.Context, data *Order) (sql.Result, error) {
	orderIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, orderRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.CreateBy, data.UpdateBy, data.Version, data.IsDel, data.UserId, data.OrderId, data.TradeId, data.PayChannel, data.Status, data.PayAmount, data.PayTime)
	}, orderIdKey)
	return ret, err
}

func (m *defaultOrderModel) Update(ctx context.Context, data *Order) error {
	orderIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, orderRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.CreateBy, data.UpdateBy, data.Version, data.IsDel, data.UserId, data.OrderId, data.TradeId, data.PayChannel, data.Status, data.PayAmount, data.PayTime, data.Id)
	}, orderIdKey)
	return err
}

func (m *defaultOrderModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheOrderIdPrefix, primary)
}

func (m *defaultOrderModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultOrderModel) tableName() string {
	return m.table
}
