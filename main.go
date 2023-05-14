package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	handleFunc()
}

func handleFunc() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/login/", login)
	router.GET("/cart/", cart)
	router.GET("/checkout/", checkout)
	router.GET("/films/:id", film)
	router.NotFound = http.HandlerFunc(notFound)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func film(w http.ResponseWriter, r *http.Request, id httprouter.Params) {
	filmId, _ := strconv.Atoi(id.ByName("id"))
	films := GetFilms()
	_, ok := films[filmId]
	if !ok {
		notFound(w, r)
	} else {
		tmplt, err := template.ParseFiles("templates/film.html", "templates/header.html", "templates/footer.html")
		if err != nil {
			fmt.Println(w, err.Error())
		}
		tmplt.ExecuteTemplate(w, "film", films[filmId])
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	tmplt, err := template.ParseFiles("templates/404.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Println(w, err.Error())
	}
	tmplt.ExecuteTemplate(w, "404", nil)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	films := GetFilms()
	tmplt, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Println(w, err.Error())
	}
	tmplt.ExecuteTemplate(w, "index", films)
}

func login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmplt, err := template.ParseFiles("templates/login.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Println(w, err.Error())
	}
	tmplt.ExecuteTemplate(w, "login", nil)
}

func cart(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmplt, err := template.ParseFiles("templates/cart.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Println(w, err.Error())
	}
	tmplt.ExecuteTemplate(w, "cart", nil)
}

func checkout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmplt, err := template.ParseFiles("templates/checkout.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Println(w, err.Error())
	}
	tmplt.ExecuteTemplate(w, "checkout", nil)
}
