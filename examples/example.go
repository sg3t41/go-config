package main

import (
	"log"

	"github.com/sg3t41/gocnf"

	"github.com/sg3t41/gocnf/testdata/json"
)

const CONFIG_FILE_PATH = "../testdata/json/config.json"

func main() {
	gocnf := gocnf.New[testdatajson.Scheme](CONFIG_FILE_PATH)

	settings, err := gocnf.Unmarshal()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(settings)
}
