package internal

import (
	"fmt"
	"time"
)

// Data storage
type RouteStats struct {
	TotalRequestsCount int           `json:"totalRequestsCount"`
	RequestsFrequency  float64       `json:"requestsFrequency"`  // per second
	AverageProcessTime time.Duration `json:"averageProcessTime"` // in ms
	LastRequestTime    time.Time     `json:"lastRequestTime"`
}

// Interpreting data for the user
type RouteStatsOutput struct {
	TotalRequestsCount string `json:"totalRequestsCount"`
	RequestsFrequency  string `json:"requestsFrequency"`  // per second
	AverageProcessTime string `json:"averageProcessTime"` // in ms
	LastRequestTime    string `json:"lastRequestTime"`
}

type ProcessData struct {
	RequestProcessTime time.Duration
}

var routeStats map[string]*RouteStats
var initTime time.Time

func init() {
	routeStats = make(map[string]*RouteStats)
	initTime = time.Now()
}

func GetAllStats() map[string]RouteStatsOutput {
	destMap := make(map[string]RouteStatsOutput)

	for key, value := range routeStats {
		destMap[key] = convertToOutput(*value)
	}

	return destMap
}

func convertToOutput(stats RouteStats) RouteStatsOutput {
	output := RouteStatsOutput{
		TotalRequestsCount: fmt.Sprint(stats.TotalRequestsCount),
		RequestsFrequency:  fmt.Sprintf("%.2f per second", stats.RequestsFrequency),
		AverageProcessTime: fmt.Sprintf("%.2s ms", fmt.Sprint(stats.AverageProcessTime.Milliseconds())),
		LastRequestTime:    stats.LastRequestTime.Format("2006-01-02 15:04:05"),
	}

	return output
}

func GetStats(url string) *RouteStats {
	stats, exists := routeStats[url]
	if !exists {
		stats = &RouteStats{}
		routeStats[url] = stats
	}

	return stats
}

func (stats *RouteStats) Update(processData ProcessData) {
	stats.updateTotalRequestsCount()
	stats.updateLastRequestTime()
	stats.updateRequestsFrequency()
	stats.updateAverageProcessTime(processData.RequestProcessTime)
}

func (stats *RouteStats) updateTotalRequestsCount() {
	stats.TotalRequestsCount++
}

func (stats *RouteStats) updateRequestsFrequency() {
	convertedLastRequestTime := stats.LastRequestTime
	upTime := convertedLastRequestTime.Sub(initTime).Seconds()
	stats.RequestsFrequency = float64(stats.TotalRequestsCount) / upTime
}

func (stats *RouteStats) updateAverageProcessTime(value time.Duration) {
	averageValue := (stats.AverageProcessTime + value) / 2
	stats.AverageProcessTime = averageValue
}

func (stats *RouteStats) updateLastRequestTime() {
	stats.LastRequestTime = time.Now()
}
