package logger

import (
	"io"
	"log/slog"
	"time"

	"github.com/lmittmann/tint"
)

func NewTint(w io.Writer, opts ...*slog.HandlerOptions) slog.Handler {
	opt := &tint.Options{
		AddSource:   true,
		Level:       slog.LevelInfo,
		TimeFormat:  time.RFC3339,
		ReplaceAttr: tintReplaceAttr,
	}
	if len(opts) != 0 && opts[0] != nil {
		f := opts[0]
		opt.AddSource = f.AddSource
		opt.Level = f.Level
		if fr := f.ReplaceAttr; fr != nil {
			opt.ReplaceAttr = fr
		}
	}

	return tint.NewHandler(w, opt)
}

func tintReplaceAttr(groups []string, attr slog.Attr) slog.Attr {
	if attr.Key == "error" {
		return tint.Attr(1, attr)
	}

	return attr
}
