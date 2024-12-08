package gocnf

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/sg3t41/gocnf/option"
	"gopkg.in/yaml.v3"
)

func Unmarshal[C any](opt *option.Option) (*C, error) {
	var path string
	// RunModeとファイルパスの設定がされていない場合はFilePathを使用する
	modeToFilePath := opt.ModeToFilePath[opt.RunMode]
	if modeToFilePath != "" {
		path = modeToFilePath
		log.Printf("【INFO】実行モード[%s] で設定ファイル [%s] を読み込みます。", opt.RunMode, modeToFilePath)
	} else {
		if opt.DefaultFilePath == "" {
			return nil, fmt.Errorf("設定ファイルパスが未指定です。SetFilePath()を呼び出して設定してください。")
		}
		path = opt.DefaultFilePath
		log.Printf("【INFO】デフォルトに設定されたファイルを読み込みます。 パス: [%s]", opt.DefaultFilePath)
	}

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
