package yaml

import (
	"fmt"

	"github.com/sg3t41/gocnf/strategy"
	"gopkg.in/yaml.v3"
)

type YamlStrategy struct {
	strategy.Strategy
}

func (ys *YamlStrategy) Unmarshal(in []byte, out any) error {
	if err := yaml.Unmarshal(in, out); err != nil {
		return fmt.Errorf("Unmarshalに失敗しました エラー: %v", err)
	}

	return nil
}
