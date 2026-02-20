package app

import (
	"github.com/stretchr/testify/mock"
)

type MockProducer struct {
	mock.Mock
}

type MockPresenter struct {
	mock.Mock
}

func (m *MockProducer) Produce() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockPresenter) Present(data []string) error {
	args := m.Called(data)
	return args.Error(0)
}
