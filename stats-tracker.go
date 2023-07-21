package statstracker

import (
	"net/http"
	"sync"
)

type RequestStats struct {
	mu      sync.Mutex
	counter map[string]int // Карта для отслеживания количества запросов по пути
}

func TrackRequestStats(stats *RequestStats) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			stats.mu.Lock()
			stats.counter[r.URL.Path]++
			stats.mu.Unlock()
			next.ServeHTTP(w, r)
		})
	}
}
