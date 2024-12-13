package file

import (
	"fmt"
	"os"
)

func Load(path string) ([]byte, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("設定ファイルの読み込みに失敗しました。 %s: %v", path, err)
	}
	return bytes, nil
}
