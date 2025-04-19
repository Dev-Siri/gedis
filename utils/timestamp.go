package utils

import (
	"fmt"
	"time"
)

func GetCurrentTimestamp() string {
	currentTime := time.Now()
	day := currentTime.Day()
	month := int(currentTime.Month())
	year := currentTime.Year() % 100
	hour := currentTime.Hour()
	minute := currentTime.Minute()
	second := currentTime.Second()

	timestamp := fmt.Sprintf("%02d%02d%d-%02d:%02d:%02d", day, month, year, hour, minute, second)
	return timestamp
}