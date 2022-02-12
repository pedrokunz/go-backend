package entity

import (
	"github.com/pedrokunz/go_backend/usecase/contract"
	"math"
)

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

	ErrInvalidOriginOrDestination TaxError = "invalid origin or destination"
)

type taxCalculator struct{}

func (c *taxCalculator) Calculate(origin, destination, plan string, duration float64) (withPlan float64, withoutPlan float64, err error) {
	planLimit := plans[plan]
	minuteByOriginAndDestinationValue := originDestinationMinuteValues[originDestinationKey{origin, destination}]
	if minuteByOriginAndDestinationValue == 0 {
		return 0, 0, ErrInvalidOriginOrDestination
	}

	if duration < planLimit {
		withPlan = 0
	} else {
		exceededMinutes := duration - planLimit
		exceedingMinuteValue := minuteByOriginAndDestinationValue * 1.1
		withPlan = toFixed(exceededMinutes*exceedingMinuteValue, 2)
	}

	withoutPlan = toFixed(minuteByOriginAndDestinationValue*duration, 2)

	return withPlan, withoutPlan, nil
}

func NewTaxCalculator() contract.PhoneCallTaxCalculator {
	return &taxCalculator{}
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
