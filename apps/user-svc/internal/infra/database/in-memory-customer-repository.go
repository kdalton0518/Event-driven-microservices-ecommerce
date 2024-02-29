package database

import (
	"user-svc/internal/domain/customer"
)

type InMemoryCustomerRepo struct {
	cus []customer.Customer
}

func NewInMemoryCustomerRepo(cus []customer.Customer) *InMemoryCustomerRepo {
	return &InMemoryCustomerRepo{
		cus: cus,
	}
}

func (r *InMemoryCustomerRepo) FindById(id string) (*customer.Customer, error) {
	var c *customer.Customer
	for _, v := range r.cus {
		if v.ID == id {
			c = &v
			break
		}
	}
	return c, nil
}

func (r *InMemoryCustomerRepo) FindByEmail(email string) (*customer.Customer, error) {
	var c *customer.Customer
	for _, v := range r.cus {
		if v.Email == email {
			c = &v
			break
		}
	}
	return c, nil
}

func (r *InMemoryCustomerRepo) Save(cust *customer.Customer) (*customer.Customer, error) {
	r.cus = append(r.cus, *cust)
	return cust, nil
}
