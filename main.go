package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/unrolled/render.v1"
	"net/http"
)

func main() {
	db := NewDB()
	renderer := render.New(render.Options{})

	router := httprouter.New()

	// Home Route
	router.GET("/", HomeHandler(db, renderer))

	// New Book Routes
	router.GET("/add", AddBookPageHandler(db, renderer))
	router.POST("/add", AddBookPostHandler(db, renderer))

	router.ServeFiles("/static/*filepath", http.Dir("static"))

	http.ListenAndServe(":8080", router)

}
