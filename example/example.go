package main

import (
	"fmt"
	"log"

	c "github.com/sg3t41/gocnf"
	o "github.com/sg3t41/gocnf/option"
)

var (
	ConfigFileBasePath    = "../config"
	DevelopConfigFileName = "config.develop.yml"
	StagingConfigFileName = "config.staging.yml"
	ProductConfigFileName = "config.product.yml"
)

func main() {
	option := o.NewOption()
	option.
		SetBasePath(ConfigFileBasePath).
		SetFileName(o.Development, DevelopConfigFileName).
		SetFileName(o.Staging, StagingConfigFileName).
		SetFileName(o.Production, ProductConfigFileName).
		LoadCurrentRunMode("RUN_MODE")

	cfg, err := c.Unmarshal[Scheme](option)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Api port is %d\n", cfg.API.Port)
}
