package tax

import (
	"github.com/pedrokunz/go_backend/usecase/repository"
	"math"
)

type Calculator struct {
	repository repository.Calculator
}

func (c *Calculator) Calculate(origin, destination, plan string, duration float64) (withPlan float64, withoutPlan float64, err error) {
	planLimit, err := c.repository.GetPlanLimitByName(plan)
	if err != nil {
		return 0, 0, ErrPlanNotFound
	}

	minuteValue, err := c.repository.GetMinuteValueByOriginAndDestination(origin, destination)
	if minuteValue == 0 || err != nil {
		return 0, 0, ErrInvalidOriginOrDestination
	}

	if duration < planLimit {
		withPlan = 0
	} else {
		exceededMinutes := duration - planLimit
		exceedingMinuteValue := minuteValue * 1.1
		withPlan = toFixed(exceededMinutes*exceedingMinuteValue, 2)
	}

	withoutPlan = toFixed(minuteValue*duration, 2)

	return withPlan, withoutPlan, nil
}

func NewTaxCalculator(repository repository.Calculator) (*Calculator, error) {
	if repository == nil {
		return nil, ErrInvalidRepository
	}

	return &Calculator{
		repository: repository,
	}, nil
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
