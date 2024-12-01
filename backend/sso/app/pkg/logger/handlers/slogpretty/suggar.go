package slogpretty

import "fmt"

func (h *PrettyLogger) Info(msg string) {
	h.sl.Info(msg)
}

func (h *PrettyLogger) Error(msg string) {
	h.sl.Error(msg)
}

func (h *PrettyLogger) Debug(msg string, args ...any) {
	h.sl.Debug(msg)
}

func (h *PrettyLogger) Infof(msg string, args ...any) {
	h.sl.Info(fmt.Sprintf(msg, args))
}

func (h *PrettyLogger) Errorf(msg string, args ...any) {
	h.sl.Error(fmt.Sprintf(msg, args))
}

func (h *PrettyLogger) Debugf(msg string, args ...any) {
	h.sl.Debug(fmt.Sprintf(msg, args))
}
