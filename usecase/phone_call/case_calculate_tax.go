package phone_call

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

func (s *usecaseCalculateTax) Perform(input *CalculateTaxInput) (*CalculateTaxOutput, error) {	
	withPlan, withoutPlan, err := s.taxCalculator.Calculate(input.Origin, input.Destination, input.Plan, input.Duration)
	return &CalculateTaxOutput{
		WithPlan:    withPlan,
		WithoutPlan: withoutPlan,
	}, err
}
