package main

import (
	"log"
	"net/http"
	"fmt"
	"reflect"
)

type A struct {
	Foo string
	Foo2 int
}

func main() {
	router := NewRouter()
	server := http.ListenAndServeTLS(":24672", "server.crt", "server.key", router) // 24672 for development port
	log.Fatal(server)
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