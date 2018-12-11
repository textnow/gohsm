package hsm

import "go.uber.org/zap"

type Service interface {
	Logger() *zap.Logger
	Set(interface{}, interface{})
	Get(interface{}) (interface{})
}

type DefaultService struct {
	logger *zap.Logger
	props map[interface{}]interface{}
}

func NewDefaultService(logger *zap.Logger) (*DefaultService) {
	return &DefaultService{
		logger: logger,
		props: make(map[interface{}]interface{}),
	}
}

func (c *DefaultService) Logger() *zap.Logger {
	return c.logger
}

func (c *DefaultService) Set(key interface{}, value interface{}) {
	c.props[key] = value
}

func (c *DefaultService) Get(key interface{}) (interface{}) {
	return c.props[key]
}
