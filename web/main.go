package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	r.Get("/{articleId}", getArticle)

	http.ListenAndServe(":8080", r)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	articleId := chi.URLParam(r, "articleId")

	// resMap := map[string]string{}
	// resMap["articleId"] = articleId
	resMap := map[string]string{
		"articleId": articleId,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resMap)
}
