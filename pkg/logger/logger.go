package logger

import (
	"log/slog"
	"os"
	"sync"
)

var lock = &sync.Mutex{}

var ins *slog.Logger

func Get() *slog.Logger {
	if ins == nil {
		lock.Lock()
		defer lock.Unlock()

		if ins == nil {
			opts := &slog.HandlerOptions{
				AddSource: true,
			}

			handler := slog.NewJSONHandler(os.Stdout, opts)

			ins = slog.New(handler)
		}
	}

	return ins
}
