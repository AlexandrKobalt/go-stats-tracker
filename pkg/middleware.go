package statstracker

import "net/http"

type StatisticsMiddleware struct {
	stats Statistics
	next  http.Handler
}

func NewStatisticsMiddleware(stats Statistics, next http.Handler) *StatisticsMiddleware {
	return &StatisticsMiddleware{stats: stats, next: next}
}

func (m *StatisticsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Увеличить счетчик запросов для данного URL
	m.stats.Increment(r.URL.Path)
	m.next.ServeHTTP(w, r)
}
