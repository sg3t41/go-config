package strategy

import (
	"fmt"
	"os"
)

type Strategy interface {
	Load(path string) error
	Unmarshal(out any) error
}

type DefaultStrategy struct {
	data []byte
}

func (d *DefaultStrategy) Load(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("設定ファイルの読み込みに失敗しました。 %s: %v", path, err)
	}
	d.data = bytes
	return nil
}
