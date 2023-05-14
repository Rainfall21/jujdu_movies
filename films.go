package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetFilms() map[int]map[string]string {
	var (
		id           int
		title        string
		description  string
		actors       string
		director     string
		price        float64
		trailer_link string
		film_link    string
		poster_link  string
	)
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "2112"
		dbname   = "web_store"
	)
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT * FROM films")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	Films := make(map[int]map[string]string)

	for rows.Next() {
		err := rows.Scan(&id, &title, &description, &actors, &director, &price, &trailer_link, &film_link, &poster_link)
		if err != nil {
			panic(err)
		}
		Films[id] = map[string]string{}
		Films[id]["title"] = title
		Films[id]["description"] = description
		Films[id]["actors"] = actors
		Films[id]["director"] = director
		Films[id]["price"] = fmt.Sprintf("%.2f", price/100)
		Films[id]["trailer_link"] = trailer_link
		Films[id]["film_link"] = film_link
		Films[id]["poster_link"] = poster_link
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return Films
}
