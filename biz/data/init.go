package data

import (
	"diploma_search/biz/model"
	"embed"
	"encoding/json"
	"log"
)

var (
	Persons []model.Person
)

//go:embed mock_data.json
var f embed.FS

func initJsonData() {
	content, err := f.ReadFile("mock_data.json")
	if err != nil {
		log.Fatalln("error open file data.json : ", err)
	}
	err = json.Unmarshal(content, &Persons)
	if err != nil {
		log.Fatalln("error json Unmarshal : ", err)
	}
}

func init() {
	initJsonData()
}
