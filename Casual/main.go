package main

import (
	"log"
	"net/http"
	"fmt"
	"reflect"
	"os"
	"encoding/json"
	"strconv"
)

type A struct {
	Foo string
	Foo2 int
}

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

func testReflection(){
	fmt.Println("Exiting the program");

	obj := A{ Foo:"asd", Foo2: 3 }

	newval := reflect.ValueOf(obj)
	objType := newval.Type()

	for i := 0; i < newval.NumField(); i++ {

		fmt.Print(objType.Field(i).Name)
		fmt.Print(": ")
		fmt.Print(newval.Field(i))
		fmt.Printf(", (%s)", objType.Field(i).Type)
		fmt.Println()
	}
}