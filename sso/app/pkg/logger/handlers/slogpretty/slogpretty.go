package slogpretty

import (
	"context"
	"encoding/json"
	"io"
	stdLog "log"
	"log/slog"
	"os"

	"github.com/fatih/color"
)

var Log *PrettyLogger

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type PrettyLogger struct {
	sl *slog.Logger
}

type PrettyHandlerOptions struct {
	SlogOpts *slog.HandlerOptions
}

type PrettyHandler struct {
	opts PrettyHandlerOptions
	slog.Handler
	l     *stdLog.Logger
	attrs []slog.Attr
}

func (opts PrettyHandlerOptions) NewPrettyHandler(
	out io.Writer,
) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, opts.SlogOpts),
		l:       stdLog.New(out, "", 0),
	}

	return h
}

func (h *PrettyHandler) Handle(_ context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())

	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	for _, a := range h.attrs {
		fields[a.Key] = a.Value.Any()
	}

	var b []byte
	var err error

	if len(fields) > 0 {
		b, err = json.MarshalIndent(fields, "", "  ")
		if err != nil {
			return err
		}
	}

	timeStr := r.Time.Format("[15:05:05.000]")
	msg := color.CyanString(r.Message)

	h.l.Println(
		timeStr,
		level,
		msg,
		color.WhiteString(string(b)),
	)

	return nil
}

func (h *PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &PrettyHandler{
		Handler: h.Handler,
		l:       h.l,
		attrs:   attrs,
	}
}

func (h *PrettyHandler) WithGroup(name string) slog.Handler {
	// TODO: implement
	return &PrettyHandler{
		Handler: h.Handler.WithGroup(name),
		l:       h.l,
	}
}

func SetupLogger(env string) {
	var opts PrettyHandlerOptions

	switch env {
	case envLocal:
		opts = PrettyHandlerOptions{
			SlogOpts: &slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		}
	case envDev:
		opts = PrettyHandlerOptions{
			SlogOpts: &slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		}
	case envProd:
		opts = PrettyHandlerOptions{
			SlogOpts: &slog.HandlerOptions{
				Level: slog.LevelInfo,
			},
		}
	}

	handler := opts.NewPrettyHandler(os.Stdout)
	Log = &PrettyLogger{}
	Log.sl = slog.New(handler)
}

func setupPrettySlog() *slog.Logger {
	opts := PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
