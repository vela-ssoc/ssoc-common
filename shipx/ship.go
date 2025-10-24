package shipx

import (
	"context"
	"log/slog"

	"github.com/vela-ssoc/ssoc-common/logger"
	"github.com/xgfone/ship/v5"
)

func NewLog(h slog.Handler, skip ...int) ship.Logger {
	num := 6
	if len(skip) > 0 {
		num = skip[0]
	}

	sh := logger.Skip(h, num)
	log := slog.New(sh)

	return &shipLog{log: log}
}

type shipLog struct {
	log *slog.Logger
}

func (s *shipLog) Tracef(msg string, args ...any) {
	s.logf(slog.LevelDebug, msg, args...)
}

func (s *shipLog) Debugf(msg string, args ...any) {
	s.logf(slog.LevelDebug, msg, args...)
}

func (s *shipLog) Infof(msg string, args ...any) {
	s.logf(slog.LevelInfo, msg, args...)
}

func (s *shipLog) Warnf(msg string, args ...any) {
	s.logf(slog.LevelWarn, msg, args...)
}

func (s *shipLog) Errorf(msg string, args ...any) {
	s.logf(slog.LevelError, msg, args...)
}

func (s *shipLog) logf(level slog.Level, msg string, args ...any) {
	ctx := context.Background()
	if !s.log.Enabled(ctx, level) {
		return
	}

	s.log.Log(ctx, level, msg, args...)
}
