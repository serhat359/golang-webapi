package main

type NewMember struct{
	Id int64			`db:"id"`
	FirstName string	`db:"firstname"`
	LastName string		`db:"lastname"`
}

type LhScore struct{
	Id int64			`db:"id"`
	Score int			`db:"score"`
	MangaName string	`db:"manga_name"`
}

type LhReadChapter struct{
	Id int64			`db:"id"`
	MangaName string	`db:"manga_name"`
	MangaChapter string	`db:"manga_chapter"`
}

func (LhReadChapter) TableName() string {
	return "lh_read_chapter"
}