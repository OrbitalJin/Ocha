package utils

import (
	"os"
	"path/filepath"
)

func Read(path string) (string, string, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return "", "", err
	}
	name := filepath.Base(path)
	content := string(dat)
	return name, content, nil
}