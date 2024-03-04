package database

import (
	"context"

	"github.com/buemura/event-driven-commerce/order-svc/internal/domain/order"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxOrderRepository struct {
	conn *pgxpool.Pool
}

func NewPgxOrderRepository(conn *pgxpool.Pool) *PgxOrderRepository {
	return &PgxOrderRepository{
		conn: conn,
	}
}

func (r *PgxOrderRepository) FindById(id string) (*order.Order, error) {
	rows, err := r.conn.Query(context.Background(), `SELECT * FROM "order" WHERE id = $1`, id)
	o, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[order.Order])
	if err != nil {
		return nil, err
	}
	if len(o) == 0 {
		return nil, nil
	}
	return o[0], nil
}

func (r *PgxOrderRepository) FindMany(in *order.GetManyOrdersIn) (*order.OrderRepositoryPaginatedOut, error) {
	limit := in.Items
	offset := (in.Page - 1) * in.Items

	rows, err := r.conn.Query(context.Background(), `SELECT * FROM "order" LIMIT $1 OFFSET $2`, limit, offset)
	o, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[order.Order])
	if err != nil {
		return nil, err
	}

	var totalCount int
	err = r.conn.QueryRow(context.Background(), `SELECT count(id) as total_count FROM "order"`).Scan(&totalCount)
	if err != nil {
		return nil, err
	}

	return &order.OrderRepositoryPaginatedOut{
		OrderList:  o,
		TotalCount: totalCount,
	}, nil
}

func (r *PgxOrderRepository) Save(o *order.Order) (*order.Order, error) {
	_, err := r.conn.Exec(
		context.Background(),
		`
		INSERT INTO "order" (id, customer_id, product_id_list, total_price, status, payment_method, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		`,
		o.ID, o.CustomerId, o.ProductIdList, o.TotalPrice, o.Status, o.PaymentMethod, o.CreatedAt, o.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return o, nil
}
