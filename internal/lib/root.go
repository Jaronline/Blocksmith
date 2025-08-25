package lib

import (
	"os"
	"path/filepath"
	"strings"
)

const defaultVersion = "1.0.0"

func getDefaultName() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	name := filepath.Base(pwd)
	name = strings.ReplaceAll(strings.ToLower(name), " ", "-")
	return name, nil
}
