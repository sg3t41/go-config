package gocnf

import (
	"testing"

	"github.com/sg3t41/gocnf/option"
	"github.com/stretchr/testify/assert"
)

type Endpoint struct {
	Path   string `yaml:"path"`
	Method string `yaml:"method"`
}

type API struct {
	Port      int        `yaml:"port"`
	Debug     bool       `yaml:"debug"`
	Endpoints []Endpoint `yaml:"endpoints"`
}

type Frontend struct {
	Port        int      `yaml:"port"`
	EnableHTTPS bool     `yaml:"enable_https"`
	Domains     []string `yaml:"domains"`
}

type Replica struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Database struct {
	Host     string    `yaml:"host"`
	Port     int       `yaml:"port"`
	Username string    `yaml:"username"`
	Password string    `yaml:"password"`
	Replicas []Replica `yaml:"replicas"`
}

type Upstream struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Proxy struct {
	Host      string     `yaml:"host"`
	Port      int        `yaml:"port"`
	SSL       bool       `yaml:"ssl"`
	Upstreams []Upstream `yaml:"upstreams"`
}

type Scheme struct {
	API      API      `yaml:"api"`
	Frontend Frontend `yaml:"frontend"`
	Database Database `yaml:"database"`
	Proxy    Proxy    `yaml:"proxy"`
}

var (
	ConfigFileBasePath    = "./config"
	DevelopConfigFileName = "config.develop.yml"
	StagingConfigFileName = "config.staging.yml"
	ProductConfigFileName = "config.product.yml"
)

func Test_Unmarshal(t *testing.T) {
	t.Run("正常動作確認", func(t *testing.T) {
		opt := option.NewOption()
		opt.
			SetBasePath(ConfigFileBasePath).
			SetFileName(option.Development, DevelopConfigFileName).
			SetFileName(option.Staging, StagingConfigFileName).
			SetFileName(option.Production, ProductConfigFileName).
			LoadCurrentRunMode("RUN_MODE")

		c, err := Unmarshal[Scheme](opt)
		// エラーが発生しないことを確認
		assert.NoError(t, err, "設定ファイル読み込み中にエラーが発生しました")

		// 設定ファイルから読み込んだポート番号が期待通りか確認
		assert.Equal(t, 8080, c.API.Port, "ポートの値が一致しません。")
	})

	t.Run("異常動作確認:存在しないファイルパスを指定", func(t *testing.T) {
		undefinedFilePath := "./undefined/config.yml"
		opt := option.NewOption()
		opt.
			SetBasePath(undefinedFilePath).
			SetFileName(option.Development, DevelopConfigFileName).
			SetFileName(option.Staging, StagingConfigFileName).
			SetFileName(option.Production, ProductConfigFileName).
			LoadCurrentRunMode("RUN_MODE")

		c, err := Unmarshal[Scheme](opt)

		assert.Error(t, err, "設定ファイルが存在しない場合はエラーが発生するべきです。")
		assert.Nil(t, c, "設定ファイルが読み込めない場合、構造体はnilであるべきです。")
	})

	t.Run("異常動作確認:ジェネリクスに文字列(構造体以外)を指定", func(t *testing.T) {
		opt := option.NewOption()
		c, err := Unmarshal[string](opt)

		assert.Error(t, err, "構造体以外の型を指定した場合はエラーが発生するべきです。")
		assert.Nil(t, c, "設定ファイルが読み込めない場合、構造体はnilであるべきです。")
	})
}
