package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

const GET string = "GET"
const POST string = "POST"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	route(router, GET, "/", "Index", Index)
	route(router, GET, "/todos/{todoId}", "TodoShow", TodoShow)
	route(router, GET, "/list", "RowList", RowList)
	route(router, GET, "/LhInfo/{mangaName}", "LhInfo", LhInfo)

	route(router, POST, "/LhInfoAddChapter", "LhInfoAddChapter", LhInfoAddChapter)
	route(router, POST, "/LhInfoSetScore", "LhInfoSetScore", LhInfoSetScore)
	route(router, POST, "/LhGetScoreBatch", "LhGetScoreBatch", LhGetScoreBatch)

	return router
}

func route(router *mux.Router, method, path, name string, handler http.HandlerFunc){
	router.
		Methods(method).
		Path(path).
		Name(name).
		Handler(handler)
}