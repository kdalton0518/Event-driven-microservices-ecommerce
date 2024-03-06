package customer

type CustomerService struct {
	cApi CustomerApi
}

func NewCustomerService(cApi CustomerApi) *CustomerService {
	return &CustomerService{
		cApi: cApi,
	}
}

func (s *CustomerService) SignIn(in *SignInRequest) (*SignInResponse, error) {
	return s.cApi.SignIn(in)
}

func (s *CustomerService) SignUp(in *SignUpRequest) (*SignUpResponse, error) {
	return s.cApi.SignUp(in)
}

func (s *CustomerService) GetCustomer(id string) (*Customer, error) {
	return s.cApi.GetCustomer(id)
}
