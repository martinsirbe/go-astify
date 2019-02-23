package main

import (
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/martinsirbe/go-astify/json"
)

func main() {
	jsonFile, err := os.Open("test_data/test.json")
	if err != nil {
		log.WithError(err).Fatal("failed to open the json file")
	}
	defer func() {
		if err := jsonFile.Close(); err != nil {
			log.WithError(err).Fatal("failed to close JSON file")
		}
	}()

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.WithError(err).Fatal("failed to read JSON file")
	}

	tokenizer := json.NewTokenizer([]rune(string(jsonBytes)))
	parser := json.NewParser(tokenizer)

	ast, err := parser.Parse()
	if err != nil {
		log.WithError(err).Fatal("failed to parse the JSON file")
	}

	log.Infof("ast - %+v", ast)
}
