package usecase

type Error string

func (e Error) Error() string {
	return string(e)
}

var (
	ErrRepositoryIsNil Error = "repository is nil"
	ErrLoggerIsNil Error = "logger is nil"
)