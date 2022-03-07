package models

import (
	"database/sql"
	"log"
	"time"
)

type Blog struct {
	Id      int
	Title   string
	Content string
	Created time.Time
}

var (
	Db *sql.DB
)

func init() {
	initPG()
}

func initPG() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp password=gwp dbname=blog sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func GetAll() (blogs []Blog) {
	return
}

func GetBlog(id int) (blog Blog) {
	return
}

func SaveBlog(blog Blog) (bg Blog) {
	return
}

func DelBlog(blog Blog) {
	return
}
