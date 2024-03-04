package product

import "errors"

var ErrProductNotFound = errors.New("product not found")
var ErrProductInsufficientQuantity = errors.New("product insufficient quantity")
