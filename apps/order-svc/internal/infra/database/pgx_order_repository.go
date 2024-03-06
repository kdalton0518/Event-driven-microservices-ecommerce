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
	rows, err := r.conn.Query(
		context.Background(),
		`
		SELECT 
			o.*, 
			ARRAY_AGG(ROW(op.product_id, op.price, op.quantity)) AS product_list
		FROM "order" o 
		INNER JOIN order_product op ON o.id = op.order_id 
		WHERE o.id = $1
		GROUP BY o.id`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderList []*order.Order
	for rows.Next() {
		var order order.Order

		if err := rows.Scan(
			&order.ID, &order.CustomerId, &order.TotalPrice, &order.Status, &order.PaymentMethod,
			&order.CreatedAt, &order.UpdatedAt, &order.ProductList,
		); err != nil {
			return nil, err
		}

		orderList = append(orderList, &order)
	}

	if len(orderList) == 0 {
		return nil, nil
	}
	return orderList[0], nil
}

func (r *PgxOrderRepository) FindMany(in *order.GetManyOrdersIn) (*order.OrderRepositoryPaginatedOut, error) {
	limit := in.Items
	offset := (in.Page - 1) * in.Items

	rows, err := r.conn.Query(
		context.Background(),
		`
		SELECT 
			o.*,
			ARRAY_AGG(ROW(op.product_id, op.price, op.quantity)) AS product_list
		FROM "order" o
		INNER JOIN order_product op ON o.id = op.order_id
		GROUP BY o.id
		LIMIT $1
		OFFSET $2`,
		limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderList []*order.Order
	for rows.Next() {
		var order order.Order

		if err := rows.Scan(
			&order.ID, &order.CustomerId, &order.TotalPrice, &order.Status, &order.PaymentMethod,
			&order.CreatedAt, &order.UpdatedAt, &order.ProductList,
		); err != nil {
			return nil, err
		}

		orderList = append(orderList, &order)
	}

	var totalCount int
	err = r.conn.QueryRow(context.Background(), `SELECT count(id) as total_count FROM "order"`).Scan(&totalCount)
	if err != nil {
		return nil, err
	}

	return &order.OrderRepositoryPaginatedOut{
		OrderList:  orderList,
		TotalCount: totalCount,
	}, nil
}

func (r *PgxOrderRepository) Save(o *order.Order) (*order.Order, error) {
	tx, err := r.conn.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	batch := &pgx.Batch{}

	batch.Queue(`
		INSERT INTO "order" (id, customer_id, total_price, status, payment_method, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		o.ID, o.CustomerId, o.TotalPrice, o.Status, o.PaymentMethod, o.CreatedAt, o.UpdatedAt,
	)

	for _, p := range o.ProductList {
		batch.Queue(`
			INSERT INTO order_product (order_id, product_id, price, quantity)
			VALUES ($1, $2, $3, $4)`,
			o.ID, p.ID, p.Price, p.Quantity,
		)
	}

	bRes := tx.SendBatch(context.Background(), batch)
	if _, err := bRes.Exec(); err != nil {
		return nil, err
	}
	if err := bRes.Close(); err != nil {
		return nil, err
	}
	if err := tx.Commit(context.Background()); err != nil {
		return nil, err
	}
	return o, nil
}
