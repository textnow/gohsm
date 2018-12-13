package hsm

import (
	"go.uber.org/zap"
)

// Assertion causes an runtime panic if expression is not true
func Assertion(logger *zap.Logger, expression bool, message string) {
	if !expression {
		logger.Panic(message)
	}
}
