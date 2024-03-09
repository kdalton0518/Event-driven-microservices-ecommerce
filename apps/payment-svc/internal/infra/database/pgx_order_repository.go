package database

import (
	"context"

	"github.com/buemura/event-driven-commerce/payment-svc/internal/domain/order"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxOrderRepository struct {
	conn *pgxpool.Pool
}

func NewPgxOrderRepository() *PgxOrderRepository {
	return &PgxOrderRepository{
		conn: Conn,
	}
}

func (r *PgxOrderRepository) FindById(id string) (*order.Order, error) {
	rows, err := r.conn.Query(
		context.Background(),
		`
		SELECT 
			o.*, 
			ARRAY_AGG(ROW(p.id, p.order_id, p.status, p.created_at, p.updated_at)) AS payment_order_list
		FROM "order" o 
		INNER JOIN payment p ON o.id = p.order_id 
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
			&order.ID, &order.Amount, &order.Status, &order.PaymentMethod,
			&order.CreatedAt, &order.UpdatedAt, &order.PaymentOrderList,
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

func (r *PgxOrderRepository) Save(o *order.Order) (*order.Order, error) {
	_, err := r.conn.Query(
		context.Background(),
		`
		INSERT INTO "order" (id, amount, status, payment_method, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6)`,
		o.ID, o.Amount, o.Status, o.PaymentMethod, o.CreatedAt, o.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return o, nil
}
