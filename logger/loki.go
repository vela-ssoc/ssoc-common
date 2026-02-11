package logger

import (
	"context"
	"log/slog"
)

type lokiHandler struct {
	opts *slog.HandlerOptions
}

func (l *lokiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return l.opts.Level.Level() <= level
}

func (l *lokiHandler) Handle(ctx context.Context, record slog.Record) error {
	// TODO implement me
	panic("implement me")
}

func (l *lokiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// TODO implement me
	panic("implement me")
}

func (l *lokiHandler) WithGroup(name string) slog.Handler {
	// TODO implement me
	panic("implement me")
}
