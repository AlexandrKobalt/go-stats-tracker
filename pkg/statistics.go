package statstracker

type Stats struct {
	URL          string
	RequestCount int
}

type Statistics interface {
	Increment(url string)
	GetStats() []Stats
}
