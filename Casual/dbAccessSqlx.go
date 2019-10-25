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