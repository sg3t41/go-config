package strategy

import (
	"fmt"
	"os"
)

type IStrategy interface {
	Load(path string) error
	Unmarshal(out any) error
}

type Strategy struct {
	Data []byte
}

func (d *Strategy) Load(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("設定ファイルの読み込みに失敗しました。 %s: %v", path, err)
	}
	d.Data = bytes
	return nil
}
