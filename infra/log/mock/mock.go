package mock

import "context"

type Logger struct {
	onInfo func(ctx context.Context, msg string)
	onDebug func(ctx context.Context, msg string)
	onError func(ctx context.Context, msg string)
	onWarn func(ctx context.Context, msg string)
}

func (l *Logger) Info(ctx context.Context, msg string) {
	l.onInfo(ctx, msg)
}

func (l *Logger) Debug(ctx context.Context, msg string) {
	l.onDebug(ctx, msg)
}

func (l *Logger) Error(ctx context.Context, msg string) {
	l.onError(ctx, msg)
}

func (l *Logger) Warn(ctx context.Context, msg string) {
	l.onWarn(ctx, msg)
}
