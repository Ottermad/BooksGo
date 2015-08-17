package main

import (
	"database/sql"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/unrolled/render.v1"
	//"log"
	"net/http"
)

func HomeHandler(db *sql.DB, renderer *render.Render) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		books := ReadBooks(db)
		renderer.HTML(w, http.StatusOK, "index", books)
	}
}

func AddBookPageHandler(db *sql.DB, renderer *render.Render) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		renderer.HTML(w, http.StatusOK, "add_book", nil)
	}
}

func AddBookPostHandler(db *sql.DB, renderer *render.Render) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		req.ParseForm()
		InsertBook(db, string(req.Form["title"][0]), string(req.Form["author"][0]))
		http.Redirect(w, req, "/", 302)
	}
}
