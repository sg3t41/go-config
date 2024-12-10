package main

import (
	"fmt"

	"github.com/sg3t41/gocnf/config"
	"github.com/sg3t41/gocnf/strategy"
)

func main() {
	c := config.NewConfig()
	c.
		SetFilePath("../config/config.local.yml").
		SetStrategy(&strategy.YamlStrategy{})

	var scheme Config
	err := c.Unmarshal(&scheme)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(&scheme)
}
