package customer

type CustomerRepository interface {
	FindById(id string) (*Customer, error)
	FindByEmail(email string) (*Customer, error)
	Save(customer *Customer) (*Customer, error)
}
