package json

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/sg3t41/gocnf/strategy"
)

type JSONStrategy struct {
	strategy.Strategy
}

func (js *JSONStrategy) Unmarshal(in []byte, out any) error {
	if err := json.Unmarshal(in, out); err != nil {
		return fmt.Errorf("Unmarshalに失敗しました エラー: %v", err)
	}

	return nil
}
