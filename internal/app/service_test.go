package app

import (
	"errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"testing"
)

var (
	Input = []string{
		"Hello! See my website http://site.com http //",
	}
	Expected = []string{
		"Hello! See my website http://******** http //",
	}
)

func TestServiceSuccess(t *testing.T) {
	testPresenter := new(MockPresenter)
	testProducer := new(MockProducer)
	testService := NewService(testProducer, testPresenter)

	testProducer.On("Produce").Return(Input, nil)
	testPresenter.On("Present", Expected).Return(nil)

	err := testService.Run()
	require.NoError(t, err, "Выполнение успешно")
	testProducer.AssertExpectations(t)
	testPresenter.AssertExpectations(t)
}

func TestServiceWithProducerError(t *testing.T) {
	testPresenter := new(MockPresenter)
	testProducer := new(MockProducer)
	testService := NewService(testProducer, testPresenter)

	expectedErr := errors.New("Error with producer's work")
	testProducer.On("Produce").Return(([]string)(nil), expectedErr)
	testPresenter.On("Present", mock.Anything).Return(([]string)(nil))

	err := testService.Run()
	assert.Error(t, err, "Ошибка")
	assert.Contains(t, err.Error(), "ошибка при работе сервиса")
	testProducer.AssertExpectations(t)
}

func TestServiceWithPresenterError(t *testing.T) {
	testPresenter := new(MockPresenter)
	testProducer := new(MockProducer)
	testService := NewService(testProducer, testPresenter)

	expectedErr := errors.New("Error with presenter's work")

	testProducer.On("Produce").Return(Input, nil)
	testPresenter.On("Present", Expected).Return(expectedErr)

	err := testService.Run()
	assert.Error(t, err, "Ошибка")
	assert.Contains(t, err.Error(), "ошибка при работе сервиса")
	testProducer.AssertExpectations(t)
	testPresenter.AssertExpectations(t)
}

func TestServiceNil(t *testing.T) {
	testService := &Service{}

	err := testService.Run()
	assert.Error(t, err, "Ошибка, зависимости не инициализированы")
	assert.Contains(t, err.Error(), "ошибка при работе сервиса")
}
