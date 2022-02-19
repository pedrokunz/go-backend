package tax

import (
	"errors"
	"github.com/pedrokunz/go_backend/usecase/phone_call"
)

type originDestinationKey struct {
	Origin      string
	Destination string
}

type Mock struct {
	plans                         map[string]float64
	originDestinationMinuteValues map[originDestinationKey]float64
}

func New() phone_call.CalculatorRepository {
	return &Mock{
		plans: map[string]float64{
			"FaleMais30":  30.0,
			"FaleMais60":  60.0,
			"FaleMais120": 120.0,
		},
		originDestinationMinuteValues: map[originDestinationKey]float64{
			{"011", "016"}: 1.90,
			{"016", "011"}: 2.90,
			{"011", "017"}: 1.70,
			{"017", "011"}: 2.70,
			{"011", "018"}: 0.90,
			{"018", "011"}: 1.90,
		},
	}
}

func (m *Mock) GetPlanLimitByName(name string) (float64, error) {
	planLimit, ok := m.plans[name]
	if !ok {
		return 0, errors.New("plan not found")
	}

	return planLimit, nil
}

func (m *Mock) GetMinuteValueByOriginAndDestination(origin, destination string) (float64, error) {
	minuteValue, ok := m.originDestinationMinuteValues[originDestinationKey{origin, destination}]
	if !ok {
		return 0, errors.New("minute value not found")
	}

	return minuteValue, nil
}
