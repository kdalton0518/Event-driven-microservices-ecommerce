package database

import (
	"context"
	"customer-svc/internal/domain/customer"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxCustomerRepository struct {
	conn *pgxpool.Pool
}

func NewPgxCustomerRepository(conn *pgxpool.Pool) *PgxCustomerRepository {
	return &PgxCustomerRepository{
		conn: conn,
	}
}

func (r *PgxCustomerRepository) FindById(id string) (*customer.Customer, error) {
	rows, err := r.conn.Query(context.Background(), `SELECT * FROM customer WHERE id = $1`, id)
	c, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[customer.Customer])
	if err != nil {
		return nil, err
	}
	if len(c) == 0 {
		return nil, nil
	}
	return c[0], nil
}

func (r *PgxCustomerRepository) FindByEmail(email string) (*customer.Customer, error) {
	rows, err := r.conn.Query(context.Background(), `SELECT * FROM customer WHERE email = $1`, email)
	c, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[customer.Customer])
	if err != nil {
		return nil, err
	}
	if len(c) == 0 {
		return nil, nil
	}
	return c[0], nil
}

func (r *PgxCustomerRepository) Save(cust *customer.Customer) (*customer.Customer, error) {
	_, err := r.conn.Query(
		context.Background(),
		`INSERT INTO customer (id, name, email, password) VALUES ($1, $2, $3, $4)`,
		cust.ID, cust.Name, cust.Email, cust.Password,
	)
	if err != nil {
		return nil, err
	}
	return cust, nil
}
