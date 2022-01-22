package usecase

import "context"

type Logger interface {
	Info(ctx context.Context, msg string)
	Debug(ctx context.Context, msg string)
	Error(ctx context.Context, msg string)
	Warn(ctx context.Context, msg string)
}
