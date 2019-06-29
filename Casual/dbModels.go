package main

type NewMember struct{
	Id int64			`db:"id"`
	FirstName string	`db:"firstname"`
	LastName string		`db:"lastname"`
}
