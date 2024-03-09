package database

import (
	"context"
	"time"

	"github.com/buemura/event-driven-commerce/payment-svc/internal/domain/payment"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxPaymentRepository struct {
	conn *pgxpool.Pool
}

func NewPgxPaymentRepository() *PgxPaymentRepository {
	return &PgxPaymentRepository{
		conn: Conn,
	}
}

func (r *PgxPaymentRepository) FindById(id string) (*payment.Payment, error) {
	rows, err := r.conn.Query(context.Background(), `SELECT * FROM payment WHERE id = $1`, id)
	p, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByPos[payment.Payment])
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *PgxPaymentRepository) FindByOrderId(id string) ([]*payment.Payment, error) {
	rows, err := r.conn.Query(context.Background(), `SELECT * FROM payment WHERE order_id = $1`, id)
	p, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[payment.Payment])
	if err != nil {
		return nil, err
	}
	if len(p) == 0 {
		return nil, nil
	}
	return p, nil
}

func (r *PgxPaymentRepository) FindPendingByOrderId(id string) (*payment.Payment, error) {
	rows, err := r.conn.Query(context.Background(), `SELECT * FROM payment WHERE order_id = $1 AND status = $2`, id, payment.PaymentPending)
	p, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[payment.Payment])
	if err != nil {
		return nil, err
	}
	if len(p) == 0 {
		return nil, nil
	}
	return p[0], nil
}

func (r *PgxPaymentRepository) Save(p *payment.Payment) (*payment.Payment, error) {
	_, err := r.conn.Exec(
		context.Background(),
		`
		INSERT INTO payment (id, order_id, status, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5)`,
		p.ID, p.OrderId, p.Status, p.CreatedAt, p.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *PgxPaymentRepository) Update(id, status string) error {
	_, err := r.conn.Exec(
		context.Background(),
		`
		UPDATE payment SET status = $1, updated_at = $2
		WHERE id = $3
		`,
		status, time.Now(), id,
	)

	if err != nil {
		return err
	}
	return nil
}
