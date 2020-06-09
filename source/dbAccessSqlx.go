package main

import (
	_ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)

func getNewMembers() []NewMember {
	db := getDB()
	defer db.Close()

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
	defer db.Close()

	score := getLhScore(db, mangaName)
	chapters := selectChapters(getLhReadChapters(db, mangaName))

	scorePoint := 0
	if score != nil {
		scorePoint = score.Score
	}

	return LhInfoData{Score: scorePoint, ReadChapters: chapters}
}

func getLhScore(db *sqlx.DB, mangaName string) *LhScore {
	lhScores := []LhScore{}
	err := db.Select(&lhScores, "SELECT * FROM lh_score WHERE manga_name = $1", mangaName)
	if err != nil {
		panic(err)
	}

	var returnVal *LhScore = nil

	if len(lhScores) > 0 {
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
	defer db.Close()

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
	rows, err := db.Query(query, args...)
	if err != nil {
		panic(err)
	}

	lhScores := make(map[string]LhScore)
	for rows.Next() {
		singleScore := LhScore{}
		err = rows.Scan(&singleScore.Id, &singleScore.Score, &singleScore.MangaName)
		if err != nil {
			panic(err)
		}
		lhScores[singleScore.MangaName] = singleScore
	}

	dic := make(map[string]int)

	for _, mangaName := range mangaNames {
		score := 0
		lhScore, exists := lhScores[mangaName]
		if exists {
			score = lhScore.Score
		}

		dic[mangaName] = score
	}

	return dic
}

func GetChapterStatus(chapters []string) map[string]bool {
	db := getDB()
	defer db.Close()

	arg := map[string]interface{}{
		"chapters": chapters,
	}
	query, args, err := sqlx.Named("SELECT manga_chapter FROM lh_read_chapter WHERE manga_chapter IN (:chapters)", arg)
	if err != nil {
		panic(err)
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		panic(err)
	}
	query = db.Rebind(query)
	rows, err := db.Query(query, args...)
	if err != nil {
		panic(err)
	}

	lhChapters := make(map[string]bool)
	for rows.Next() {
		singleChapter := ""
		err = rows.Scan(&singleChapter)
		if err != nil {
			panic(err)
		}

		if singleChapter != "" {
			lhChapters[singleChapter] = true
		}
	}

	dic := make(map[string]bool)

	for _, chapter := range chapters {
		score := false

		_, exists := lhChapters[chapter]
		if exists {
			score = true
		}

		dic[chapter] = score
	}

	return dic
}