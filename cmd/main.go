package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AlexandrKobalt/go-stats-tracker/pkg"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Use(pkg.StatsMiddleware)

	r.Post("/post", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Got POST request!")
	})

	r.Post("/getStats", func(w http.ResponseWriter, r *http.Request) {
		jsonResponse, _ := json.Marshal(pkg.GetAllStats())

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(jsonResponse)
	})

	http.ListenAndServe("localhost:8081", r)
}
