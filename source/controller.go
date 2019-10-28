package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Json(w http.ResponseWriter, o interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(o)
	if err != nil {
		panic(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func RowList(w http.ResponseWriter, r *http.Request) {
	testConnection()

	members := getNewMembers()

	Json(w, members)
}

func LhInfo(w http.ResponseWriter, r *http.Request){
	enableCors(&w)

	testConnection()

	vars := mux.Vars(r)
	mangaName := vars["mangaName"]

	lhInfo := getLhInfo(mangaName)

	Json(w, lhInfo)
}

func LhInfoAddChapter(w http.ResponseWriter, r *http.Request){
	enableCors(&w)

	testConnection()

	decoder := json.NewDecoder(r.Body)

	chapter := LhReadChapter{}
	decoder.Decode(&chapter)

	SaveChapter(&chapter)

	Json(w, "")
}

func LhInfoSetScore(w http.ResponseWriter, r *http.Request){
	enableCors(&w)

	testConnection()

	decoder := json.NewDecoder(r.Body)
	score := LhScore{}
	decoder.Decode(&score)

	SaveScore(&score)

	Json(w, "")
}

func LhGetScoreBatch(w http.ResponseWriter, r *http.Request){
	enableCors(&w)

	testConnection()

	decoder := json.NewDecoder(r.Body)
	mangas := []LhMangaChapterData{}
	decoder.Decode(&mangas)

	mangaNames := selectMangaNames(mangas)
	scores := GetScores(mangaNames)

	chapters := selectMangaChapters(mangas)
	chapterStatus := GetChapterStatus(chapters)

	type ResultStruct struct {
		Scores map[string]int
		ChapterStatus map[string]bool
	}

	Json(w, ResultStruct{ Scores: scores, ChapterStatus: chapterStatus })
}