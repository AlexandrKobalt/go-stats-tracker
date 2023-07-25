package main

import (
	"fmt"
	"net/http"

	api "github.com/AlexandrKobalt/go-stats-tracker/pkg/api"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Use(api.StatsMiddleware)

	r.Post("/post", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Got POST request!")
	})

	r.Post("/getStats", func(w http.ResponseWriter, r *http.Request) {
		stats, err := api.GetAllStats()
		if err != nil {
			panic("it's just impossible, I can't believe it")
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(stats)
	})

	http.ListenAndServe("localhost:8081", r)
}
