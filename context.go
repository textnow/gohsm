package hsm

import "go.uber.org/zap"

type Context interface {
	Logger() *zap.Logger
}

type DefaultContext struct {
	logger *zap.Logger
}

func NewDefaultContext(logger *zap.Logger) (Context) {
	return &DefaultContext{
		logger: logger,
	}
}

func (c *DefaultContext) Logger() *zap.Logger {
	return c.logger
}