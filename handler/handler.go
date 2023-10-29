package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	year := time.Now().Year()

	t, _ := template.ParseFiles(
		"templates/index.html",
		"templates/header.html",
		"templates/footer.html",
	)
	t.Execute(w, map[string]string{"Year": strconv.Itoa(year)})
}

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/images/fav/favicon.ico")
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)
	json.NewEncoder(w).Encode(map[string]any{
		"code":    404,
		"message": "nothing to see here",
	})
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(405)
	json.NewEncoder(w).Encode(map[string]any{
		"code":    405,
		"message": "not allowed my friend",
	})
}
