package main

import (
	"./Models"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	configuration, err := loadConfiguration()
	if err != nil {
		log.Fatalln(err.Error())
	}
	context := GetContext(configuration)
	router := RegisterRoutes(context, configuration.Server.Prefix)
	log.Printf("API server is up and listening on http://localhost:%s%s\n", configuration.Server.Port, configuration.Server.Prefix)
	err = http.ListenAndServe(":" + configuration.Server.Port, router)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func loadConfiguration() (Models.Configuration, error) {
	file, _ := os.Open("appSettings.json")
	defer file.Close()
	configuration := Models.Configuration{}
	err := json.NewDecoder(file).Decode(&configuration)
	if err != nil {
		return Models.Configuration{}, err
	}
	return configuration, nil
}
