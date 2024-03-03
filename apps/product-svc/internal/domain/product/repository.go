package product

type ProductRepositoryPaginatedOut struct {
	ProductList []*Product
	TotalCount  int
}

type ProductRepository interface {
	FindMany() (*ProductRepositoryPaginatedOut, error)
	FindById(int) (*Product, error)
}
