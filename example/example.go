package main

import (
	"fmt"
	"log"

	"github.com/sg3t41/gocnf"
	"github.com/sg3t41/gocnf/enum"
	"github.com/sg3t41/gocnf/option"
)

var (
	defaultConfigFilePath = "../config/config.local.yml"
	developConfigFilePath = "../config/config.develop.yml"
	stagingConfigFilePath = "../config/config.staging.yml"
	productConfigFilePath = "../config/config.product.yml"

	/*
	   runModeEnvKey はアプリケーションの実行モードを設定する環境変数のキーを指定します。
	   設定可能な値は "development", "staging", "production" です。

	   設定例:
	   	export RUN_MODE=development
	*/
	runModeEnvKey = "RUN_MODE"
)

func main() {
	o := option.NewOption()
	o.
		SetDefaultFilePath(defaultConfigFilePath).
		SetModeToFilePath(enum.Development, developConfigFilePath).
		SetModeToFilePath(enum.Staging, stagingConfigFilePath).
		SetModeToFilePath(enum.Production, productConfigFilePath).
		SetCurrentRunMode(runModeEnvKey)

	cfg, err := gocnf.Unmarshal[Scheme](o)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", cfg)
}
