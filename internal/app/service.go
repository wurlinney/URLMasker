package app

import (
	"fmt"
	"urlmasker/internal/urlmasker"
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
	if s.prod == nil || s.pres == nil {
		return fmt.Errorf("ошибка при работе сервиса")
	}
	data, err := s.prod.Produce()
	if err != nil {
		return fmt.Errorf("ошибка при работе сервиса: %w", err)
	}
	res := make([]string, 0, len(data))
	for _, str := range data {
		maskedStr := s.MaskURLInService(str)
		res = append(res, maskedStr)
	}
	if err := s.pres.Present(res); err != nil {
		return fmt.Errorf("ошибка при работе сервиса: %w", err)
	}

	return nil
}

func NewService(prod Producer, pres Presenter) *Service {
	return &Service{prod, pres}

}

func (s *Service) MaskURLInService(input string) string {
	return urlmasker.DoMaskForLink(input)
}
