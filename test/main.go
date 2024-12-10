package main

import (
	"fmt"

	"github.com/sg3t41/gocnf"
	"github.com/sg3t41/gocnf/pkg/filetype"
)

func main() {
	gc := gocnf.GoCnf[Scheme]{
		FileType: filetype.YAML,
		FilePath: "../config/config.local.yml",
	}

	cnf, err := gc.Unmarshal()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cnf.API.Port)
}
