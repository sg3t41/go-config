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
			Strategy: strategy.Strategy{Data: []byte(`{"Key":"value"}`)},
		}

		// アンマーシャルする構造体
		type TestStruct struct {
			Key string `json:"Key"`
		}
		var out TestStruct

		// Unmarshalの実行
		err := mockStrategy.Unmarshal(&out)

		// 結果確認
		assert.NoError(t, err, "Unmarshalが成功するべきです")
		assert.Equal(t, "value", out.Key, "構造体に正しいデータがセットされるべきです")
	})

	t.Run("異常系: 引数がポインタ型でない場合", func(t *testing.T) {
		// モックデータを用意
		mockStrategy := &json.JSONStrategy{
			// 正しいJSONデータを設定
			Strategy: strategy.Strategy{Data: []byte(`{"Key":"value"}`)},
		}

		// 引数として構造体を渡す（ポインタではない）
		type TestStruct struct {
			Key string `json:"Key"`
		}
		var out TestStruct

		// Unmarshalの実行
		err := mockStrategy.Unmarshal(out) // ポインタでない

		// 結果確認
		assert.Error(t, err, "ポインタ型でない場合エラーが発生すべきです")
		assert.Contains(t, err.Error(), "ポインタ型である必要があります", "エラーメッセージが正しいべきです")
	})

	t.Run("異常系: 引数が構造体のポインタ型でない場合", func(t *testing.T) {
		// モックデータを用意
		mockStrategy := &json.JSONStrategy{
			// 正しいJSONデータを設定
			Strategy: strategy.Strategy{Data: []byte(`{"Key":"value"}`)},
		}

		// 引数としてポインタ型でない構造体を渡す（間違った型）
		type TestStruct struct {
			Key string `json:"Key"`
		}
		out := "not a struct"

		// Unmarshalの実行
		err := mockStrategy.Unmarshal(&out) // 構造体ではない

		// 結果確認
		assert.Error(t, err, "構造体でない場合エラーが発生すべきです")
		assert.Contains(t, err.Error(), "構造体のポインタ型である必要があります", "エラーメッセージが正しいべきです")
	})

	t.Run("異常系: JSONパースエラー", func(t *testing.T) {
		// モックデータを用意（無効なJSON）
		mockStrategy := &json.JSONStrategy{
			// 無効なJSONを設定
			Strategy: strategy.Strategy{Data: []byte(`Key: value`)},
		}

		// アンマーシャルする構造体
		type TestStruct struct {
			Key string `json:"Key"`
		}
		var out TestStruct

		// Unmarshalの実行
		err := mockStrategy.Unmarshal(&out)

		// 結果確認
		assert.Error(t, err, "JSONのパースエラーが発生すべきです")
		assert.Contains(t, err.Error(), "Unmarshalに失敗しました", "エラーメッセージが正しいべきです")
	})
}

