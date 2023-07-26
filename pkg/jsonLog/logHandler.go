package jsonlog

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

func ProcessLogs(logDir string, keyword string) ([]byte, error) {
	stats := make(map[string]map[string]int)

	err := filepath.Walk(logDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
			fileStats, err := processLogFile(path, keyword)
			if err != nil {
				return err
			}
			for k, v := range fileStats {
				if _, found := stats[keyword]; !found {
					stats[keyword] = make(map[string]int)
				}
				stats[keyword][k] += v
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	// Возвращаем []byte с данными JSON
	return json.MarshalIndent(stats, "", "  ")
}

func processLogFile(file string, keyword string) (map[string]int, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var logs []map[string]interface{}
	if err := json.Unmarshal(data, &logs); err != nil {
		return nil, err
	}

	fileStats := make(map[string]int)

	for _, log := range logs {
		if name, ok := log["name"].(string); ok && name == keyword {
			errorName, _ := log["error"].(string)
			fileStats[errorName]++
		}
	}

	return fileStats, nil
}
