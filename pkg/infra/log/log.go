package log

import (
	gokitlog "github.com/go-kit/log"
)

var (
	loggersToReload []ReloadableHandler
)

type ConcreteLogger struct {
	ctx []interface{}
	gokitlog.SwapLogger
}

// Reload reloads all loggers.
func Reload() error {
	for _, logger := range loggersToReload {
		if err := logger.Reload(); err != nil {
			return err
		}
	}

	return nil
}
