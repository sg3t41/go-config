package main

import (
	"fmt"

	c "github.com/sg3t41/gocnf"
	o "github.com/sg3t41/gocnf/option"
)

var (
	ConfigFileBasePath    = "/config"
	DevelopConfigFileName = "config.develop.yml"
	StagingConfigFileName = "config.staging.yml"
	ProductConfigFileName = "config.product.yml"

	/*
	   RunModeEnvKey はアプリケーションの実行モードを設定する環境変数のキーを指定します。
	   設定可能な値は "development", "staging", "production" です。

	   設定例:
	   	export RUN_MODE=development
	*/
	RunModeEnvKey = "RUN_MODE"
)

func main() {
	option := o.NewOption()
	option.
		SetBasePath(ConfigFileBasePath).
		SetFileName(o.Development, DevelopConfigFileName).
		SetFileName(o.Staging, StagingConfigFileName).
		SetFileName(o.Production, ProductConfigFileName).
		LoadCurrentRunMode(RunModeEnvKey)

	cfg, _ := c.Unmarshal[Scheme](option)
	fmt.Printf("Key1: %s\n", cfg.App1.Key1)
}
