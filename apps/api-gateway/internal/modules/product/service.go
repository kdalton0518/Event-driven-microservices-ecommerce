package product

type ProductService struct {
	pApi ProductApi
}

func NewProductService(pApi ProductApi) *ProductService {
	return &ProductService{
		pApi: pApi,
	}
}

func (s *ProductService) GetManyProducts(opts GetManyProductsIn) (*GetManyProductsOut, error) {
	return s.pApi.GetManyProducts(opts)
}

func (s *ProductService) GetProduct(id int) (*Product, error) {
	return s.pApi.GetProduct(id)
}
