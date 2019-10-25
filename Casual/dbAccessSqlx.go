package main

import (
    _ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)

func getNewMembers() []NewMember {
	db, err := sqlx.Connect("postgres", getConnectionString())
    if err != nil {
        panic(err)
	}
	
	members := []NewMember{}
	db.Select(&members, "SELECT * FROM member")
	
	return members
}

func getLhInfo(mangaName string) LhInfoData {
	db, err := sqlx.Connect("postgres", getConnectionString())
    if err != nil {
        panic(err)
	}

	score := getLhScore(db, mangaName)
	chapters := getLhReadChapters(db, mangaName)

	return LhInfoData{ Score: score, ReadChapters: chapters }
}

func getLhScore(db *sqlx.DB, mangaName string) *LhScore {
	lhScores := []LhScore{}
	err := db.Select(&lhScores, "SELECT * FROM lh_score WHERE manga_name = $1", mangaName)
	if err != nil {
        panic(err)
	}

	var returnVal *LhScore = nil

	if len(lhScores) > 0{
		returnVal = &lhScores[0]
	}

	return returnVal
}

func getLhReadChapters(db *sqlx.DB, mangaName string) []LhReadChapter {
	readChapters := []LhReadChapter{}

	err := db.Select(&readChapters, "SELECT * FROM lh_read_chapter WHERE manga_name = $1", mangaName)
	if err != nil {
        panic(err)
	}

	return readChapters
}