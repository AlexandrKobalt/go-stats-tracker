package statstracker

import "sync"

type MemoryStats struct {
	statsMap map[string]int
	mu       sync.RWMutex
}

func NewMemoryStats() *MemoryStats {
	return &MemoryStats{
		statsMap: make(map[string]int),
	}
}

func (s *MemoryStats) Increment(url string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.statsMap[url]++
}

func (s *MemoryStats) GetStats() []Stats {
	s.mu.RLock()
	defer s.mu.RUnlock()

	statsSlice := make([]Stats, 0, len(s.statsMap))
	for url, count := range s.statsMap {
		statsSlice = append(statsSlice, Stats{URL: url, RequestCount: count})
	}
	return statsSlice
}
