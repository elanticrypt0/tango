package tango_helpers

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

func OpenFile(file string) []byte {
	if _, err := os.Stat(file); err != nil {
		log.Fatalln("File dosen't exists: " + file)
	}
	filedata, err := os.ReadFile(file)
	if err != nil {
		log.Fatalln("Can not open file: " + file)
	}
	return filedata
}

func ReadAndParseToml[T any](file string, stru *T) {
	tomlData := string(OpenFile(file))
	_, err := toml.Decode(tomlData, &stru)
	if err != nil {
		log.Fatalln("Cannot parse file: %s\n" + file)
	}
}

func ReadAndParseJson[T any](file string, stru *T) {
	fileData := strings.NewReader(string(OpenFile(file)))
	jsonParser := json.NewDecoder(fileData)
	jsonParser.Decode(&stru)
}
