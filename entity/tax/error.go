package tax

type Error string

func (e Error) Error() string {
	return string(e)
}

var (
	ErrNotFound                   Error = "tax not found"
	ErrInvalidOriginOrDestination Error = "invalid origin or destination"
	ErrPlanNotFound               Error = "tax plan not found"
	ErrInvalidRepository          Error = "invalid repository"
)
