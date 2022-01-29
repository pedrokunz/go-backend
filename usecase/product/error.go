package product

type Error string

func (e Error) Error() string {
	return string(e)
}

var (
	ErrProductNotFound Error = "product not found"
	ErrProductAlreadyExists Error = "product already exists"
	ErrProductInvalidError Error = "product invalid"
	ErrProductCreateError Error = "product create error"
	ErrProductUpdateError Error = "product update error"
	ErrProductDeleteError Error = "product delete error"
	ErrProductListError Error = "product list error"
)