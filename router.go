package main

import (
	"net/http"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/users response: " + r.URL.Query().Get("page")))
	})
	mux.HandleFunc("DELETE /users/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/subtree path"))
	})
	mux.HandleFunc("PUT /users/{id...}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/dynamic path id :" + strings.Split(r.PathValue("id"), "/")[0]))
	})
	mux.HandleFunc("POST /users/vikash", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("vikash the user path"))
	})
	mux.HandleFunc("DELETE /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("/root file"))
	})
	http.ListenAndServe(":8000", mux)
}
