package entity

type originDestinationKey struct {
	Origin      string
	Destination string
}

var (
	//TODO move to repository interface
	plans = map[string]float64{
		"FaleMais30":  30.0,
		"FaleMais60":  60.0,
		"FaleMais120": 120.0,
	}

	//TODO move to repository interface
	originDestinationMinuteValues = map[originDestinationKey]float64{
		{"011", "016"}: 1.90,
		{"016", "011"}: 2.90,
		{"011", "017"}: 1.70,
		{"017", "011"}: 2.70,
		{"011", "018"}: 0.90,
		{"018", "011"}: 1.90,
	}
)

type taxCalculator struct{}

func (c *taxCalculator) Calculate(origin, destination, plan string, duration float64) (withPlan float64, withoutPlan float64, err error) {
	if plans[plan] >= duration {
		withPlan = 0
	}

	withoutPlan = originDestinationMinuteValues[originDestinationKey{origin, destination}] * duration

	return withPlan, withoutPlan, nil
}

func NewTaxCalculator() *taxCalculator {
	return &taxCalculator{}
}
