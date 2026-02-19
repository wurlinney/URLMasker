package app

import (
	"fmt"
)

type Producer interface {
	Produce() ([]string, error)
}

type Presenter interface {
	Present([]string) error
}

type Service struct {
	prod Producer
	pres Presenter
}

func (s *Service) Run() error {
	data, err := s.prod.Produce()
	if err != nil {
		return fmt.Errorf("ошибка при работе сервиса: %v", err)
	}
	var res []string
	for _, str := range data {
		maskedStr := s.DoMaskForLink(str)
		res = append(res, maskedStr)
	}
	if err := s.pres.Present(res); err != nil {
		return fmt.Errorf("ошибка при работе сервиса: %v", err)
	}

	return nil
}

func NewService(prod Producer, pres Presenter) *Service {
	return &Service{prod, pres}

}
