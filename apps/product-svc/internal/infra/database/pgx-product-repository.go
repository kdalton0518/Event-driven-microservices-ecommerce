package database

import (
	"context"

	"github.com/buemura/event-driven-commerce/product-svc/internal/domain/product"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxProductRepository struct {
	conn *pgxpool.Pool
}

func NewPgxProductRepository(conn *pgxpool.Pool) *PgxProductRepository {
	return &PgxProductRepository{
		conn: conn,
	}
}

func (r *PgxProductRepository) FindById(id int) (*product.Product, error) {
	rows, err := r.conn.Query(context.Background(), `SELECT * FROM product WHERE id = $1`, id)
	p, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[product.Product])
	if err != nil {
		return nil, err
	}
	if len(p) == 0 {
		return nil, nil
	}
	return p[0], nil
}

func (r *PgxProductRepository) FindMany() (*product.ProductRepositoryPaginatedOut, error) {
	rows, err := r.conn.Query(context.Background(), `SELECT * FROM product`)
	p, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[product.Product])
	if err != nil {
		return nil, err
	}

	var totalCount int
	err = r.conn.QueryRow(context.Background(), `SELECT count(id) as total_count FROM product`).Scan(&totalCount)
	if err != nil {
		return nil, err
	}

	return &product.ProductRepositoryPaginatedOut{
		ProductList: p,
		TotalCount:  totalCount,
	}, nil
}
