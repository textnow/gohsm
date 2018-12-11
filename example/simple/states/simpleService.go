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
	ss := SimpleService{
		Service: ctx,
	}
	return &ss
}

func NewSimpleService(ctx hsm.Service, test string) *SimpleService {
	ss := &SimpleService{
		Service: ctx,
	}

	// Initial save into map in the HSM context
	ss.Set(TEST_STRING, test)

	return ss
}

func (ss *SimpleService) GetTest() string {
	test := ss.Service.Get(TEST_STRING)
	if test == nil {
		return "";
	}
	return test.(string)
}

func (ss *SimpleService) SetTest(value string) {
	ss.Set(TEST_STRING, value)
}
