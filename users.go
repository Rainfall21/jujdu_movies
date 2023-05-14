package main

// import (
// 	"database/sql"
// 	"fmt"
// )

// func LogInUser(email string, password string) string {

// 	const (
// 		host        = "localhost"
// 		port        = 5432
// 		user        = "postgres"
// 		SQLpassword = "2112"
// 		dbname      = "web_store"
// 	)

// 	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		host, port, user, SQLpassword, dbname)
// 	db, err := sql.Open("postgres", conn)
// 	if err != nil {
// 		panic(err)
// 	}
// 	rows, err := db.Query("SELECT * FROM users WHERE email = ?", email)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if rows == nil {

// 	}
// 	defer db.Close()

// 	Films := make(map[int]map[string]string)

// 	for rows.Next() {
// 	// 	err := rows.Scan(&id, &title, &description, &actors, &director, &price, &trailer_link, &film_link, &poster_link)
// 	// 	if err != nil {
// 	// 		panic(err)
// 	// 	}
// 	// 	Films[id] = map[string]string{}
// 	// 	Films[id]["title"] = title
// 	// 	Films[id]["description"] = description
// 	// 	Films[id]["actors"] = actors
// 	// 	Films[id]["director"] = director
// 	// 	Films[id]["price"] = fmt.Sprintf("%.2f", price/100)
// 	// 	Films[id]["trailer_link"] = trailer_link
// 	// 	Films[id]["film_link"] = film_link
// 	// 	Films[id]["poster_link"] = poster_link
// 	// }
// 	// err = rows.Err()
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// err = db.Ping()
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// return Films
// }
