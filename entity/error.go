package entity

type TaxError string

func (e TaxError) Error() string {
	return string(e)
}

var (
	ErrTaxNotFound TaxError = "tax not found"
)
