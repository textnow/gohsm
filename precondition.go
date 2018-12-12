package hsm

import (
	"go.uber.org/zap"
)

// Precondition causes a panic if expression is not true
func Precondition(logger *zap.Logger, expression bool, message string) {
	if !expression {
		logger.Panic(message)
		panic(message)
	}
}
