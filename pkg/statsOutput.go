package pkg

import (
	"fmt"

	"github.com/AlexandrKobalt/go-stats-tracker/internal"
)

func GetAllStats() map[string]internal.RouteStatsOutput {
	destMap := make(map[string]internal.RouteStatsOutput)

	for key, value := range internal.GetCurrentRouteStats() {
		destMap[key] = convertToOutput(value)
	}

	return destMap
}

func convertToOutput(stats internal.RouteStats) internal.RouteStatsOutput {
	output := internal.RouteStatsOutput{
		TotalRequestsCount: fmt.Sprint(stats.TotalRequestsCount),
		RequestsFrequency:  fmt.Sprintf("%.2f per second", stats.RequestsFrequency),
		AverageProcessTime: fmt.Sprintf("%.2s ms", fmt.Sprint(stats.AverageProcessTime.Milliseconds())),
		LastRequestTime:    stats.LastRequestTime.Format("2006-01-02 15:04:05"),
	}

	return output
}
