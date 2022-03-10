package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=pguser dbname=db password=password sslmode=disable")
	if err != nil {
		panic(err)
	}
}

type Master struct {
	id    int
	email string
	key   string
	start string
	end   string
}

func main() {
	fmt.Println("///////////表示///////////")
	rows, err := Db.Query("SELECT * FROM masterkey ORDER BY id")
	if err != nil {
		return
	}
	for rows.Next() {
		m := Master{}
		rows.Scan(&m.id, &m.email, &m.key, &m.start, &m.end)
		fmt.Println(m)
	}
	fmt.Println()
	fmt.Println("///////////更新///////////")
	_, err = Db.Exec("INSERT INTO masterkey VALUES ($1, $2, $3, $4, $5)", 2, "test@fun.ac.jp", "aaaaaaaaaaaaa", "2020-11-17", time.Now())
	if err != nil {
		return
	}
	rows, err = Db.Query("SELECT * FROM masterkey ORDER BY id")
	if err != nil {
		return
	}
	for rows.Next() {
		m := Master{}
		rows.Scan(&m.id, &m.email, &m.key, &m.start, &m.end)
		fmt.Println(m)
	}
	fmt.Println()
}
