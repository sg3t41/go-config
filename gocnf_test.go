package gocnf_test

import (
	"os"
	"testing"

	"github.com/sg3t41/gocnf"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	t.Run("正常系: JSONファイルでのアンマーシャル", func(t *testing.T) {
		// モックJSONファイルを作成
		filePath := "config.json"
		err := os.WriteFile(filePath, []byte(`{"Key": "value"}`), 0644)
		if err != nil {
			t.Fatalf("ファイル作成に失敗しました: %v", err)
		}
		defer os.Remove(filePath) // テスト後にファイルを削除

		// 即席構造体を定義
		type Scheme struct {
			Key string `json:"Key"`
		}

		// gocnf のインスタンスを作成
		g := gocnf.New[Scheme](filePath)

		// Unmarshalを実行
		result, err := g.Unmarshal()

		// 結果確認
		assert.NoError(t, err, "Unmarshalが成功するべきです")
		assert.Equal(t, "value", result.Key, "構造体に正しいデータがセットされるべきです")
	})

	t.Run("正常系: YAMLファイルでのアンマーシャル", func(t *testing.T) {
		// モックYAMLファイルを作成
		filePath := "config.yml"
		err := os.WriteFile(filePath, []byte(`Key: value`), 0644)
		if err != nil {
			t.Fatalf("ファイル作成に失敗しました: %v", err)
		}
		defer os.Remove(filePath) // テスト後にファイルを削除

		// 即席構造体を定義
		type Scheme struct {
			Key string `yaml:"Key"`
		}

		// gocnf のインスタンスを作成
		g := gocnf.New[Scheme](filePath)

		// Unmarshalを実行
		result, err := g.Unmarshal()

		// 結果確認
		assert.NoError(t, err, "Unmarshalが成功するべきです")
		assert.Equal(t, "value", result.Key, "構造体に正しいデータがセットされるべきです")
	})

	t.Run("異常系: 無効なファイル拡張子", func(t *testing.T) {
		// モック無効なファイルパス
		filePath := "config.txt"
		g := gocnf.New[map[string]interface{}](filePath) // map型を使って柔軟に受け取る

		// Unmarshalを実行
		_, err := g.Unmarshal()

		// 結果確認
		assert.Error(t, err, "無効な拡張子の場合エラーが発生すべきです")
		assert.Contains(t, err.Error(), "ファイルタイプに適した戦略が存在しません", "エラーメッセージが正しいべきです")
	})

	t.Run("異常系: 拡張子がない場合", func(t *testing.T) {
		// モックYAMLファイルを作成
		filePath := "config/configfile"
		err := os.WriteFile(filePath, []byte(`Key: value`), 0644)
		if err != nil {
			t.Fatalf("ファイル作成に失敗しました: %v", err)
		}
		defer os.Remove(filePath) // テスト後にファイルを削除

		// 即席構造体を定義
		type Scheme struct {
			Key string `yaml:"Key"`
		}

		// gocnf のインスタンスを作成
		g := gocnf.New[Scheme](filePath)

		// Unmarshalを実行
		_, err = g.Unmarshal()

		// 結果確認
		assert.Error(t, err, "拡張子がない場合エラーが発生すべきです")
		assert.Contains(t, err.Error(), "ファイルの拡張子が見つかりません", "エラーメッセージが正しいべきです")
	})
}

