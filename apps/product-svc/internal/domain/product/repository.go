package product

type ProductRepository interface {
	FindMany() ([]*Product, error)
	FindById(int) (*Product, error)
}
