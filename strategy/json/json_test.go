package json_test

import (
	"testing"

	"github.com/sg3t41/gocnf/strategy"
	"github.com/sg3t41/gocnf/strategy/json"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	t.Run("正常系: 構造体ポインタが渡される場合", func(t *testing.T) {
		// モックデータを用意
		mockStrategy := &json.JSONStrategy{
			// 正しいJSONデータを設定
			Strategy: strategy.Strategy{},
		}

		// アンマーシャルする構造体
		type TestStruct struct {
			Key string `json:"Key"`
		}
		in := []byte(`{"Key":"value"}`)
		var out TestStruct

		// Unmarshalの実行
		err := mockStrategy.Unmarshal(in, &out)

		// 結果確認
		assert.NoError(t, err, "Unmarshalが成功するべきです")
		assert.Equal(t, "value", out.Key, "構造体に正しいデータがセットされるべきです")
	})

	t.Run("異常系: JSONパースエラー", func(t *testing.T) {
		// モックデータを用意（無効なJSON）
		mockStrategy := &json.JSONStrategy{
			// 無効なJSONを設定
			Strategy: strategy.Strategy{},
		}

		// アンマーシャルする構造体
		type TestStruct struct {
			Key string `json:"Key"`
		}
		in := []byte(`{Key:value}`)
		var out TestStruct

		// Unmarshalの実行
		err := mockStrategy.Unmarshal(in, &out)

		// 結果確認
		assert.Error(t, err, "JSONのパースエラーが発生すべきです")
		assert.Contains(t, err.Error(), "Unmarshalに失敗しました", "エラーメッセージが正しいべきです")
	})
}
