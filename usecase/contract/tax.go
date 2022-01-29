package contract

type PhoneCallTaxCalculator interface {
	Calculate(origin, destination, plan string, duration float64) (withPlan float64, withoutPlan float64, err error)
}
