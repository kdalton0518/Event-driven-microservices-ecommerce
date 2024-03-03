package customer

type CustomerApi interface {
	SignIn(*SignInRequest) (*SignInResponse, error)
	SignUp(*SignUpRequest) (*SignUpResponse, error)
	GetCustomer(string) (*Customer, error)
}
