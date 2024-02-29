package customer

import "errors"

var ErrCustomerAlreadyExists = errors.New("customer already exists")
var ErrCustomerNotFound = errors.New("customer not found")
var ErrCustomerInvalidCredential = errors.New("invalid credential")
var ErrCustomerPermissionDenied = errors.New("permission denied")
