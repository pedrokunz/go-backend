package repository

type Calculator interface {
	GetPlanLimitByName(name string) (float64, error)
	GetMinuteValueByOriginAndDestination(origin, destination string) (float64, error)
}
