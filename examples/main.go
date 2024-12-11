package main

import (
	"fmt"

	"github.com/sg3t41/gocnf"
)

func main() {
	gocnf := gocnf.New[Config]("../testdata/config.json")

	ymlcnf, err := gocnf.Unmarshal()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ymlcnf)

}
