package phone_call

import (
	"github.com/pedrokunz/go_backend/entity/tax"
	"github.com/pedrokunz/go_backend/usecase/repository"
)

type CalculateTaxInput struct {
	Plan        string
	Duration    float64
	Origin      string
	Destination string
}

type CalculateTaxOutput struct {
	WithPlan    float64
	WithoutPlan float64
}

type usecaseCalculateTax struct {
	taxCalculator *tax.Calculator
}

func NewCalculateTax(repo repository.CalculatorRepository) (*usecaseCalculateTax, error) {
	taxCalculator, err := tax.NewTaxCalculator(repo)
	if err != nil {
		return nil, err
	}

	return &usecaseCalculateTax{
		taxCalculator: taxCalculator,
	}, nil
}

func (s *usecaseCalculateTax) Perform(input *CalculateTaxInput) (*CalculateTaxOutput, error) {
	withPlan, withoutPlan, err := s.taxCalculator.Calculate(input.Origin, input.Destination, input.Plan, input.Duration)
	return &CalculateTaxOutput{
		WithPlan:    withPlan,
		WithoutPlan: withoutPlan,
	}, err
}
