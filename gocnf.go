package gocnf

import (
	"fmt"
	"path/filepath"

	"github.com/sg3t41/gocnf/config"
	"github.com/sg3t41/gocnf/strategy"
	"github.com/sg3t41/gocnf/strategy/json"
	"github.com/sg3t41/gocnf/strategy/yaml"
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
	strategy, err := getStrategy(gc.FilePath)
	if err != nil {
		return nil, err
	}

	c := config.NewConfig()
	c.
		SetFilePath(gc.FilePath).
		SetStrategy(strategy)

	var out T
	if err := c.Unmarshal(&out); err != nil {
		return nil, err
	}

	return &out, nil
}

func getStrategy(confFilePath string) (strategy.IStrategy, error) {
	// 設定ファイルの拡張子を取得する
	ext := filepath.Ext(confFilePath)
	if ext == "" {
		return nil, fmt.Errorf("ファイルの拡張子が見つかりません。")
	}

	// 設定ファイルの拡張子によってストラテジを決定する
	switch ext {
	case ".yaml":
	case ".yml":
		return &yaml.YamlStrategy{}, nil
	case ".json":
		return &json.JSONStrategy{}, nil
	// add to
	default:
		return nil, fmt.Errorf("ファイルタイプに適した戦略が存在しません。")
	}
	return nil, fmt.Errorf("ファイルタイプに適した戦略が存在しません。")
}
