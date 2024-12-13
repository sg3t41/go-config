package main

import (
	"log"

	"github.com/sg3t41/gocnf"
	testdata "github.com/sg3t41/gocnf/testdata/yaml"
)

const CONFIG_FILE_PATH = "../../testdata/yaml/config.yml"

func main() {
	settings, err := gocnf.Unmarshal[testdata.Scheme](CONFIG_FILE_PATH)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(settings)
}
