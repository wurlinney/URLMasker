package app

import (
	"bufio"
	"fmt"
	"os"
)

type FileProducer struct {
	Filename string
}

func (fprod *FileProducer) Produce() ([]string, error) {
	var res []string
	file, err := os.Open(fprod.Filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ошибка открытия файла: %w", err)
	}

	return res, nil
}

func NewProducer(filename string) *FileProducer {
	if filename != "" {
		return &FileProducer{Filename: filename}
	}
	return nil
}
