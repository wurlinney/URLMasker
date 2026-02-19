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
		return fmt.Errorf("ошибка записи в файл: %v", err)
	}
	defer file.Close()

	for _, str := range data {
		if _, err := fmt.Fprintln(file, str); err != nil {
			return fmt.Errorf("ошибка записи в файл: %v", err)
		}
	}

	return nil
}
