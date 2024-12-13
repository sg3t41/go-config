package gocnf

import (
	"fmt"
	"reflect"

	"github.com/sg3t41/gocnf/config"
	"github.com/sg3t41/gocnf/strategy"
	"github.com/sg3t41/gocnf/strategy/json"
	"github.com/sg3t41/gocnf/strategy/yaml"
	"github.com/sg3t41/gocnf/util/file"
)

func Unmarshal[T any](filePath string) (*T, error) {
	if isPtrToStruct[T]() {
		return nil, fmt.Errorf("構造体のポインター型を指定してください。")
	}

	strategy, err := getStrategy(filePath)
	if err != nil {
		return nil, err
	}

	c := config.NewConfig()
	c.
		SetFilePath(filePath).
		SetStrategy(strategy)

	var out T
	if err := c.Unmarshal(&out); err != nil {
		return nil, err
	}

	return &out, nil
}

func isPtrToStruct[T any]() bool {
	// Tの型情報を取得
	t := reflect.TypeOf((*T)(nil)).Elem()
	// Tがポインター型 かつ ポインターが指す先が構造体 であればtrue
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

func getStrategy(path string) (strategy.IStrategy, error) {
	switch file.Ext(path) {
	case ".yaml", ".yml":
		return &yaml.YamlStrategy{}, nil
	case ".json":
		return &json.JSONStrategy{}, nil
	// add to
	default:
		return nil, fmt.Errorf("ファイルタイプに適した戦略が存在しません。")
	}
}
