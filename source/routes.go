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
	allRoutes := GetRoutes()
	for _, route := range allRoutes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
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
		Route{
			"LhInfo",
			"GET",
			"/LhInfo/{mangaName}",
			LhInfo,
		},
		Route{
			"LhInfoAddChapter",
			"POST",
			"/LhInfoAddChapter",
			LhInfoAddChapter,
		},
		Route{
			"LhInfoSetScore",
			"POST",
			"/LhInfoSetScore",
			LhInfoSetScore,
		},
		Route{
			"LhGetScoreBatch",
			"POST",
			"/LhGetScoreBatch",
			LhGetScoreBatch,
		},
	}
}
