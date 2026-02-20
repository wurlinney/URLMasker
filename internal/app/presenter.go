package app

import (
	"fmt"
	"os"
)

type FilePresenter struct {
	Filename string
}

func (fpres *FilePresenter) Present(data []string) error {
	file, err := os.Create(fpres.Filename)
	if err != nil {
		return fmt.Errorf("ошибка записи в файл: %w", err)
	}
	defer file.Close()

	for _, str := range data {
		if _, err := fmt.Fprintln(file, str); err != nil {
			return fmt.Errorf("ошибка записи в файл: %w", err)
		}
	}

	return nil
}

func NewPresenter(filename string) *FilePresenter {
	return &FilePresenter{Filename: filename}
}
