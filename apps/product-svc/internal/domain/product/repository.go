package product

type ProductRepositoryPaginatedOut struct {
	ProductList []*Product
	TotalCount  int
}

type ProductRepository interface {
	FindMany(*GetManyProductsIn) (*ProductRepositoryPaginatedOut, error)
	FindById(int) (*Product, error)
	Update(*Product) (*Product, error)
}
