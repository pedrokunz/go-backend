package phone_call

import "github.com/pedrokunz/go_backend/usecase/contract"

type usecaseCalculateTax struct {
	taxCalculator contract.PhoneCallTaxCalculator
}

func NewCalculateTax(taxCalculator contract.PhoneCallTaxCalculator) *usecaseCalculateTax {
	return &usecaseCalculateTax{
		taxCalculator: taxCalculator,
	}
}