package gocnf

import (
	"fmt"
	"github.com/sg3t41/gocnf/config"
	"github.com/sg3t41/gocnf/pkg/filetype"
	"github.com/sg3t41/gocnf/strategy"
)

type GoCnf[T any] struct {
	FileType filetype.ConfigFileType
	FilePath string
}

func (gc GoCnf[T]) Unmarshal() (*T, error) {
	s, _ := getStrategy(gc.FileType)
	c := config.NewConfig()
	c.
		SetFilePath(gc.FilePath).
		SetStrategy(s)

	var t T
	err := c.Unmarshal(&t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func getStrategy(ft filetype.ConfigFileType) (strategy.Strategy, error) {
	switch ft {
	case filetype.YAML:
		return &strategy.YamlStrategy{}, nil
	// add to
	default:
		return nil, fmt.Errorf("ファイルタイプに適した戦略が存在しません。")
	}
}
