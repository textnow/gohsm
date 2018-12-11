package states

import (
	"github.com/Enflick/gohsm"
)

type ServiceKey int

const (
	TEST_STRING ServiceKey = 1 + iota
)

type SimpleService struct {
	hsm.Service
}

func ToSimpleService(ctx hsm.Service) *SimpleService {
	sc := SimpleService{
		Service: ctx,
	}
	return &sc
}

func NewSimpleService(ctx hsm.Service, test string) *SimpleService {
	sc := &SimpleService{
		Service: ctx,
	}

	// Initial save into map in the HSM context
	sc.Set(TEST_STRING, test)

	return sc
}

func (sc *SimpleService) GetTest() string {
	return sc.Service.Get(TEST_STRING).(string)
}

func (sc *SimpleService) SetTest(value string) {
	sc.Set(TEST_STRING, value)
}
