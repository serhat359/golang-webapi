package main

import (
    _ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)

func getNewMembers() []NewMember {
	db := getDB()
	
	members := []NewMember{}
	db.Select(&members, "SELECT * FROM member")
	
	return members
}

func getDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", getConnectionString())
    if err != nil {
        panic(err)
	}

	return db
}

func getLhInfo(mangaName string) LhInfoData {
	db := getDB()

	score := getLhScore(db, mangaName)
	chapters := selectChapters(getLhReadChapters(db, mangaName))

	scorePoint := 0
	if score != nil {
		scorePoint = score.Score
	}

	return LhInfoData{ Score: scorePoint, ReadChapters: chapters }
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

func GetScores(mangaNames []string) map[string]int {
	db := getDB()
	
	lhScores := []LhScore{}
	arg := map[string]interface{}{
		"mangas": mangaNames,
	}
	query, args, err := sqlx.Named("SELECT id, score, manga_name FROM lh_score WHERE manga_name IN (:mangas)", arg)
	if err != nil {
        panic(err)
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
        panic(err)
	}
	query = db.Rebind(query)
	rows,err := db.Query(query, args...)
	if err != nil {
        panic(err)
	}
	
	singleScore := LhScore{}
	for rows.Next() {
		err = rows.Scan(&singleScore.Id, &singleScore.Score, &singleScore.MangaName)
		if err != nil {
			panic(err)
		}
		lhScores = append(lhScores, singleScore)
	}

	dic := make(map[string]int, len(mangaNames))

	for _, mangaName := range mangaNames {
		score := 0
		for _, row := range lhScores{
			if row.MangaName == mangaName{
				score = row.Score
				break
			}
		}

		dic[mangaName] = score
	}

	return dic
}