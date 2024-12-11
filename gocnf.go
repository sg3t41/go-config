package gocnf

import (
	"fmt"
	"path/filepath"

	"github.com/sg3t41/gocnf/config"
	"github.com/sg3t41/gocnf/strategy"
)

type gocnf[T any] struct {
	FilePath string
}

func New[T any](filePath string) *gocnf[T] {
	return &gocnf[T]{
		FilePath: filePath,
	}
}

func (gc gocnf[T]) Unmarshal() (*T, error) {
	s, _ := getStrategy(gc.FilePath)
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

func getStrategy(confFilePath string) (strategy.Strategy, error) {
	// 設定ファイルの拡張子を取得する
	ext := filepath.Ext(confFilePath)
	if ext == "" {
		return nil, fmt.Errorf("ファイルの拡張子が見つかりません。")
	}

	// 設定ファイルの拡張子によってストラテジを決定する
	switch ext {
	case ".yaml":
	case ".yml":
		return &strategy.YamlStrategy{}, nil
	case ".json":
		return &strategy.JSONStrategy{}, nil
	// add to
	default:
		return nil, fmt.Errorf("ファイルタイプに適した戦略が存在しません。")
	}
	return nil, fmt.Errorf("ファイルタイプに適した戦略が存在しません。")
}
