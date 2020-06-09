package main

import (
	"log"
	"net/http"
	"os"
	"encoding/json"
	"strconv"
)

type Configuration struct {
    Port int `json:"port"`
}

func main(){
	runServer()
}

func runServer() {
	router := NewRouter()
	port := getPort()
	server := http.ListenAndServeTLS(port, "server.crt", "server.key", router) // 24672 for development port, 36219 for actual service
	log.Fatal(server)
}

func getPort() string {
	file, _ := os.Open("conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}
	return ":" + strconv.Itoa(configuration.Port)
}
