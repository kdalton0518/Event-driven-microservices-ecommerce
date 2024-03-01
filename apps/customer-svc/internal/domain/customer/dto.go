package customer

type CreateCustomerIn struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninCustomerIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninCustomerOut struct {
	AccessToken string      `json:"accessToken"`
	Customer    CustomerOut `json:"customer"`
}

type CustomerOut struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
