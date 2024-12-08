# gocnf

[![GitHub release](https://img.shields.io/github/v/release/sg3t41/gocnf?include_prereleases)](https://github.com/sg3t41/gocnf/releases)
![Go version](https://img.shields.io/github/go-mod/go-version/sg3t41/gocnf?style=flat-square)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/db340dce37434e5bbef6b2261eb8fb8d)](https://app.codacy.com/gh/sg3t41/gocnf/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/sg3t41/gocnf)](https://goreportcard.com/report/github.com/sg3t41/gocnf)
[![Go Reference](https://pkg.go.dev/badge/github.com/sg3t41/gocnf/v2.svg)](https://pkg.go.dev/github.com/sg3t41/gocnf)

`gocnf` - Goアプリケーションで設定ファイルを読み込むためのライブラリです。

## Feature

サポートファイル: `yaml`

※順次追加予定

## Install

```bash
go get github.com/sg3t41/gocnf
```

## Usage

```yaml
app1:
  key1: value1
  key2: value2

app2:
  key3:
    - value3
    - value4

alias:
  key4: &key4 value5

app3:
  key4: *key4

data1: null
data2:
data3: 0
```

### Load data

```go
//scheme.go
package main

type App1 struct {
	Key1 string `yaml:"key1"`
	Key2 string `yaml:"key2"`
}

type App2 struct {
	Key3 []string `yaml:"key3"`
}

type App3 struct {
	Key4 string `yaml:"key4"`
}

type Scheme struct {
	App1  App1    `yaml:"app1"`
	App2  App2    `yaml:"app2"`
	App3  App3    `yaml:"app3"`
	Data1 *string `yaml:"data1"`
	Data2 *string `yaml:"data2"`
	Data3 int     `yaml:"data3"`
}


```

```go
package main

import (
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
	// fmt.Printf("Key1: %s\n", cfg.App1.Key1)
}
```
