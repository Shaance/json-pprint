package cli

import (
	"os"
)

type OSInterface interface {
	ReadFile(name string) ([]byte, error)
	Create(name string) (*os.File, error)
}

type ActualOS struct{}

func (ActualOS) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (ActualOS) Create(name string) (*os.File, error) {
	return os.Create(name)
}
