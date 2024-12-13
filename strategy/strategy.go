package strategy

import (
	"github.com/sg3t41/gocnf/util/file"
)

type IStrategy interface {
	Load(path string) ([]byte, error)
	Unmarshal(in []byte, out any) error
}

type Strategy struct{}

func (d *Strategy) Load(path string) ([]byte, error) {
	return file.Load(path)
}
