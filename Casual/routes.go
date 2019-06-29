package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	routes := GetRoutes()
	for _, route := range routes {

		handler := Logger(route.HandlerFunc, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func GetRoutes() []Route {
	return []Route{
		Route{
			"Index",
			"GET",
			"/",
			Index,
		},
		Route{
			"TodoIndex",
			"GET",
			"/todos",
			TodoIndex,
		},
		Route{
			"TodoShow",
			"GET",
			"/todos/{todoId}",
			TodoShow,
		},
		Route{
			"RowList",
			"GET",
			"/list",
			RowList,
		},
	}
}
