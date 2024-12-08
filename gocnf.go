package gocnf

import (
	"fmt"
	"os"
	"reflect"

	"github.com/sg3t41/gocnf/option"
	"gopkg.in/yaml.v3"
)

func Unmarshal[C any](opt *option.Option) (*C, error) {
	base := opt.File.BasePath
	filename := opt.File.Name[opt.RunMode]
	path := string(base) + "/" + filename

	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("設定ファイルの読み込みに失敗しました。 %s: %v", path, err)
	}

	kind := reflect.TypeOf((*C)(nil)).Elem().Kind()
	if kind != reflect.Struct {
		return nil, fmt.Errorf("gocnf.Unmarshalのジェネリックは構造体である必要があります。 要求された型: %s", kind)
	}

	var c C
	if err := yaml.Unmarshal(bytes, &c); err != nil {
		return nil, fmt.Errorf("Unmarshalに失敗しました エラー: %v", err)
	}

	return &c, nil
}
