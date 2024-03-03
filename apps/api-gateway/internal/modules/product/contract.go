package product

type ProductApi interface {
	GetManyProducts(GetManyProductsIn) (*GetManyProductsOut, error)
	GetProduct(int) (*Product, error)
}
