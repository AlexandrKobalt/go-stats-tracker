package pkg

import (
	"net/http"
	"time"

	"github.com/AlexandrKobalt/go-stats-tracker/internal"
)

func StatsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		next.ServeHTTP(w, r)

		endTime := time.Now()

		processTime := endTime.Sub(startTime)

		processData := internal.ProcessData{
			RequestProcessTime: processTime,
		}

		url := r.URL.Path
		stats := internal.GetStats(url)
		stats.Update(processData)
	})
}
